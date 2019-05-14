package raft

import (
	"bytes"
	"errors"
	"fmt"
	pb "github.com/coreos/etcd/raft/raftpb"
	"math"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"
)

const None uint64 = 0
const noLimit = math.MaxUint64
const (
	StateFollower StateType = iota
	StateCandidate
	StateLeader
	StatePreCandidate
	numStates
)

type ReadOnlyOption int

const (
	ReadOnlySafe ReadOnlyOption = iota
	ReadOnlyLeaseBased
)
const (
	campaignPreElection CampaignType = "CampaignPreElection"
	campaignElection    CampaignType = "CampaignElection"
	campaignTransfer    CampaignType = "CampaignTransfer"
)

type lockedRand struct {
	mu   sync.Mutex
	rand *rand.Rand
}

func (r *lockedRand) Intn(n int) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.mu.Lock()
	v := r.rand.Intn(n)
	r.mu.Unlock()
	return v
}

var globalRand = &lockedRand{rand: rand.New(rand.NewSource(time.Now().UnixNano()))}

type CampaignType string
type StateType uint64

var stmap = [...]string{"StateFollower", "StateCandidate", "StateLeader", "StatePreCandidate"}

func (st StateType) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return stmap[uint64(st)]
}

type Config struct {
	ID                        uint64
	peers                     []uint64
	learners                  []uint64
	ElectionTick              int
	HeartbeatTick             int
	Storage                   Storage
	Applied                   uint64
	MaxSizePerMsg             uint64
	MaxInflightMsgs           int
	CheckQuorum               bool
	PreVote                   bool
	ReadOnlyOption            ReadOnlyOption
	Logger                    Logger
	DisableProposalForwarding bool
}

func (c *Config) validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.ID == None {
		return errors.New("cannot use none as id")
	}
	if c.HeartbeatTick <= 0 {
		return errors.New("heartbeat tick must be greater than 0")
	}
	if c.ElectionTick <= c.HeartbeatTick {
		return errors.New("election tick must be greater than heartbeat tick")
	}
	if c.Storage == nil {
		return errors.New("storage cannot be nil")
	}
	if c.MaxInflightMsgs <= 0 {
		return errors.New("max inflight messages must be greater than 0")
	}
	if c.Logger == nil {
		c.Logger = raftLogger
	}
	if c.ReadOnlyOption == ReadOnlyLeaseBased && !c.CheckQuorum {
		return errors.New("CheckQuorum must be enabled when ReadOnlyOption is ReadOnlyLeaseBased")
	}
	return nil
}

type raft struct {
	id                        uint64
	Term                      uint64
	Vote                      uint64
	readStates                []ReadState
	raftLog                   *raftLog
	maxInflight               int
	maxMsgSize                uint64
	prs                       map[uint64]*Progress
	learnerPrs                map[uint64]*Progress
	state                     StateType
	isLearner                 bool
	votes                     map[uint64]bool
	msgs                      []pb.Message
	lead                      uint64
	leadTransferee            uint64
	pendingConf               bool
	readOnly                  *readOnly
	electionElapsed           int
	heartbeatElapsed          int
	checkQuorum               bool
	preVote                   bool
	heartbeatTimeout          int
	electionTimeout           int
	randomizedElectionTimeout int
	disableProposalForwarding bool
	tick                      func()
	step                      stepFunc
	logger                    Logger
}

func newRaft(c *Config) *raft {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.validate(); err != nil {
		panic(err.Error())
	}
	raftlog := newLog(c.Storage, c.Logger)
	hs, cs, err := c.Storage.InitialState()
	if err != nil {
		panic(err)
	}
	peers := c.peers
	learners := c.learners
	if len(cs.Nodes) > 0 || len(cs.Learners) > 0 {
		if len(peers) > 0 || len(learners) > 0 {
			panic("cannot specify both newRaft(peers, learners) and ConfState.(Nodes, Learners)")
		}
		peers = cs.Nodes
		learners = cs.Learners
	}
	r := &raft{id: c.ID, lead: None, isLearner: false, raftLog: raftlog, maxMsgSize: c.MaxSizePerMsg, maxInflight: c.MaxInflightMsgs, prs: make(map[uint64]*Progress), learnerPrs: make(map[uint64]*Progress), electionTimeout: c.ElectionTick, heartbeatTimeout: c.HeartbeatTick, logger: c.Logger, checkQuorum: c.CheckQuorum, preVote: c.PreVote, readOnly: newReadOnly(c.ReadOnlyOption), disableProposalForwarding: c.DisableProposalForwarding}
	for _, p := range peers {
		r.prs[p] = &Progress{Next: 1, ins: newInflights(r.maxInflight)}
	}
	for _, p := range learners {
		if _, ok := r.prs[p]; ok {
			panic(fmt.Sprintf("node %x is in both learner and peer list", p))
		}
		r.learnerPrs[p] = &Progress{Next: 1, ins: newInflights(r.maxInflight), IsLearner: true}
		if r.id == p {
			r.isLearner = true
		}
	}
	if !isHardStateEqual(hs, emptyState) {
		r.loadState(hs)
	}
	if c.Applied > 0 {
		raftlog.appliedTo(c.Applied)
	}
	r.becomeFollower(r.Term, None)
	var nodesStrs []string
	for _, n := range r.nodes() {
		nodesStrs = append(nodesStrs, fmt.Sprintf("%x", n))
	}
	r.logger.Infof("newRaft %x [peers: [%s], term: %d, commit: %d, applied: %d, lastindex: %d, lastterm: %d]", r.id, strings.Join(nodesStrs, ","), r.Term, r.raftLog.committed, r.raftLog.applied, r.raftLog.lastIndex(), r.raftLog.lastTerm())
	return r
}
func (r *raft) hasLeader() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.lead != None
}
func (r *raft) softState() *SoftState {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &SoftState{Lead: r.lead, RaftState: r.state}
}
func (r *raft) hardState() pb.HardState {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pb.HardState{Term: r.Term, Vote: r.Vote, Commit: r.raftLog.committed}
}
func (r *raft) quorum() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(r.prs)/2 + 1
}
func (r *raft) nodes() []uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nodes := make([]uint64, 0, len(r.prs)+len(r.learnerPrs))
	for id := range r.prs {
		nodes = append(nodes, id)
	}
	for id := range r.learnerPrs {
		nodes = append(nodes, id)
	}
	sort.Sort(uint64Slice(nodes))
	return nodes
}
func (r *raft) send(m pb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.From = r.id
	if m.Type == pb.MsgVote || m.Type == pb.MsgVoteResp || m.Type == pb.MsgPreVote || m.Type == pb.MsgPreVoteResp {
		if m.Term == 0 {
			panic(fmt.Sprintf("term should be set when sending %s", m.Type))
		}
	} else {
		if m.Term != 0 {
			panic(fmt.Sprintf("term should not be set when sending %s (was %d)", m.Type, m.Term))
		}
		if m.Type != pb.MsgProp && m.Type != pb.MsgReadIndex {
			m.Term = r.Term
		}
	}
	r.msgs = append(r.msgs, m)
}
func (r *raft) getProgress(id uint64) *Progress {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr, ok := r.prs[id]; ok {
		return pr
	}
	return r.learnerPrs[id]
}
func (r *raft) sendAppend(to uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr := r.getProgress(to)
	if pr.IsPaused() {
		return
	}
	m := pb.Message{}
	m.To = to
	term, errt := r.raftLog.term(pr.Next - 1)
	ents, erre := r.raftLog.entries(pr.Next, r.maxMsgSize)
	if errt != nil || erre != nil {
		if !pr.RecentActive {
			r.logger.Debugf("ignore sending snapshot to %x since it is not recently active", to)
			return
		}
		m.Type = pb.MsgSnap
		snapshot, err := r.raftLog.snapshot()
		if err != nil {
			if err == ErrSnapshotTemporarilyUnavailable {
				r.logger.Debugf("%x failed to send snapshot to %x because snapshot is temporarily unavailable", r.id, to)
				return
			}
			panic(err)
		}
		if IsEmptySnap(snapshot) {
			panic("need non-empty snapshot")
		}
		m.Snapshot = snapshot
		sindex, sterm := snapshot.Metadata.Index, snapshot.Metadata.Term
		r.logger.Debugf("%x [firstindex: %d, commit: %d] sent snapshot[index: %d, term: %d] to %x [%s]", r.id, r.raftLog.firstIndex(), r.raftLog.committed, sindex, sterm, to, pr)
		pr.becomeSnapshot(sindex)
		r.logger.Debugf("%x paused sending replication messages to %x [%s]", r.id, to, pr)
	} else {
		m.Type = pb.MsgApp
		m.Index = pr.Next - 1
		m.LogTerm = term
		m.Entries = ents
		m.Commit = r.raftLog.committed
		if n := len(m.Entries); n != 0 {
			switch pr.State {
			case ProgressStateReplicate:
				last := m.Entries[n-1].Index
				pr.optimisticUpdate(last)
				pr.ins.add(last)
			case ProgressStateProbe:
				pr.pause()
			default:
				r.logger.Panicf("%x is sending append in unhandled state %s", r.id, pr.State)
			}
		}
	}
	r.send(m)
}
func (r *raft) sendHeartbeat(to uint64, ctx []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	commit := min(r.getProgress(to).Match, r.raftLog.committed)
	m := pb.Message{To: to, Type: pb.MsgHeartbeat, Commit: commit, Context: ctx}
	r.send(m)
}
func (r *raft) forEachProgress(f func(id uint64, pr *Progress)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for id, pr := range r.prs {
		f(id, pr)
	}
	for id, pr := range r.learnerPrs {
		f(id, pr)
	}
}
func (r *raft) bcastAppend() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.forEachProgress(func(id uint64, _ *Progress) {
		if id == r.id {
			return
		}
		r.sendAppend(id)
	})
}
func (r *raft) bcastHeartbeat() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lastCtx := r.readOnly.lastPendingRequestCtx()
	if len(lastCtx) == 0 {
		r.bcastHeartbeatWithCtx(nil)
	} else {
		r.bcastHeartbeatWithCtx([]byte(lastCtx))
	}
}
func (r *raft) bcastHeartbeatWithCtx(ctx []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.forEachProgress(func(id uint64, _ *Progress) {
		if id == r.id {
			return
		}
		r.sendHeartbeat(id, ctx)
	})
}
func (r *raft) maybeCommit() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mis := make(uint64Slice, 0, len(r.prs))
	for _, p := range r.prs {
		mis = append(mis, p.Match)
	}
	sort.Sort(sort.Reverse(mis))
	mci := mis[r.quorum()-1]
	return r.raftLog.maybeCommit(mci, r.Term)
}
func (r *raft) reset(term uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.Term != term {
		r.Term = term
		r.Vote = None
	}
	r.lead = None
	r.electionElapsed = 0
	r.heartbeatElapsed = 0
	r.resetRandomizedElectionTimeout()
	r.abortLeaderTransfer()
	r.votes = make(map[uint64]bool)
	r.forEachProgress(func(id uint64, pr *Progress) {
		*pr = Progress{Next: r.raftLog.lastIndex() + 1, ins: newInflights(r.maxInflight), IsLearner: pr.IsLearner}
		if id == r.id {
			pr.Match = r.raftLog.lastIndex()
		}
	})
	r.pendingConf = false
	r.readOnly = newReadOnly(r.readOnly.option)
}
func (r *raft) appendEntry(es ...pb.Entry) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	li := r.raftLog.lastIndex()
	for i := range es {
		es[i].Term = r.Term
		es[i].Index = li + 1 + uint64(i)
	}
	r.raftLog.append(es...)
	r.getProgress(r.id).maybeUpdate(r.raftLog.lastIndex())
	r.maybeCommit()
}
func (r *raft) tickElection() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.electionElapsed++
	if r.promotable() && r.pastElectionTimeout() {
		r.electionElapsed = 0
		r.Step(pb.Message{From: r.id, Type: pb.MsgHup})
	}
}
func (r *raft) tickHeartbeat() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.heartbeatElapsed++
	r.electionElapsed++
	if r.electionElapsed >= r.electionTimeout {
		r.electionElapsed = 0
		if r.checkQuorum {
			r.Step(pb.Message{From: r.id, Type: pb.MsgCheckQuorum})
		}
		if r.state == StateLeader && r.leadTransferee != None {
			r.abortLeaderTransfer()
		}
	}
	if r.state != StateLeader {
		return
	}
	if r.heartbeatElapsed >= r.heartbeatTimeout {
		r.heartbeatElapsed = 0
		r.Step(pb.Message{From: r.id, Type: pb.MsgBeat})
	}
}
func (r *raft) becomeFollower(term uint64, lead uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.step = stepFollower
	r.reset(term)
	r.tick = r.tickElection
	r.lead = lead
	r.state = StateFollower
	r.logger.Infof("%x became follower at term %d", r.id, r.Term)
}
func (r *raft) becomeCandidate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.state == StateLeader {
		panic("invalid transition [leader -> candidate]")
	}
	r.step = stepCandidate
	r.reset(r.Term + 1)
	r.tick = r.tickElection
	r.Vote = r.id
	r.state = StateCandidate
	r.logger.Infof("%x became candidate at term %d", r.id, r.Term)
}
func (r *raft) becomePreCandidate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.state == StateLeader {
		panic("invalid transition [leader -> pre-candidate]")
	}
	r.step = stepCandidate
	r.votes = make(map[uint64]bool)
	r.tick = r.tickElection
	r.state = StatePreCandidate
	r.logger.Infof("%x became pre-candidate at term %d", r.id, r.Term)
}
func (r *raft) becomeLeader() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.state == StateFollower {
		panic("invalid transition [follower -> leader]")
	}
	r.step = stepLeader
	r.reset(r.Term)
	r.tick = r.tickHeartbeat
	r.lead = r.id
	r.state = StateLeader
	ents, err := r.raftLog.entries(r.raftLog.committed+1, noLimit)
	if err != nil {
		r.logger.Panicf("unexpected error getting uncommitted entries (%v)", err)
	}
	nconf := numOfPendingConf(ents)
	if nconf > 1 {
		panic("unexpected multiple uncommitted config entry")
	}
	if nconf == 1 {
		r.pendingConf = true
	}
	r.appendEntry(pb.Entry{Data: nil})
	r.logger.Infof("%x became leader at term %d", r.id, r.Term)
}
func (r *raft) campaign(t CampaignType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var term uint64
	var voteMsg pb.MessageType
	if t == campaignPreElection {
		r.becomePreCandidate()
		voteMsg = pb.MsgPreVote
		term = r.Term + 1
	} else {
		r.becomeCandidate()
		voteMsg = pb.MsgVote
		term = r.Term
	}
	if r.quorum() == r.poll(r.id, voteRespMsgType(voteMsg), true) {
		if t == campaignPreElection {
			r.campaign(campaignElection)
		} else {
			r.becomeLeader()
		}
		return
	}
	for id := range r.prs {
		if id == r.id {
			continue
		}
		r.logger.Infof("%x [logterm: %d, index: %d] sent %s request to %x at term %d", r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), voteMsg, id, r.Term)
		var ctx []byte
		if t == campaignTransfer {
			ctx = []byte(t)
		}
		r.send(pb.Message{Term: term, To: id, Type: voteMsg, Index: r.raftLog.lastIndex(), LogTerm: r.raftLog.lastTerm(), Context: ctx})
	}
}
func (r *raft) poll(id uint64, t pb.MessageType, v bool) (granted int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if v {
		r.logger.Infof("%x received %s from %x at term %d", r.id, t, id, r.Term)
	} else {
		r.logger.Infof("%x received %s rejection from %x at term %d", r.id, t, id, r.Term)
	}
	if _, ok := r.votes[id]; !ok {
		r.votes[id] = v
	}
	for _, vv := range r.votes {
		if vv {
			granted++
		}
	}
	return granted
}
func (r *raft) Step(m pb.Message) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case m.Term == 0:
	case m.Term > r.Term:
		if m.Type == pb.MsgVote || m.Type == pb.MsgPreVote {
			force := bytes.Equal(m.Context, []byte(campaignTransfer))
			inLease := r.checkQuorum && r.lead != None && r.electionElapsed < r.electionTimeout
			if !force && inLease {
				r.logger.Infof("%x [logterm: %d, index: %d, vote: %x] ignored %s from %x [logterm: %d, index: %d] at term %d: lease is not expired (remaining ticks: %d)", r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), r.Vote, m.Type, m.From, m.LogTerm, m.Index, r.Term, r.electionTimeout-r.electionElapsed)
				return nil
			}
		}
		switch {
		case m.Type == pb.MsgPreVote:
		case m.Type == pb.MsgPreVoteResp && !m.Reject:
		default:
			r.logger.Infof("%x [term: %d] received a %s message with higher term from %x [term: %d]", r.id, r.Term, m.Type, m.From, m.Term)
			if m.Type == pb.MsgApp || m.Type == pb.MsgHeartbeat || m.Type == pb.MsgSnap {
				r.becomeFollower(m.Term, m.From)
			} else {
				r.becomeFollower(m.Term, None)
			}
		}
	case m.Term < r.Term:
		if r.checkQuorum && (m.Type == pb.MsgHeartbeat || m.Type == pb.MsgApp) {
			r.send(pb.Message{To: m.From, Type: pb.MsgAppResp})
		} else {
			r.logger.Infof("%x [term: %d] ignored a %s message with lower term from %x [term: %d]", r.id, r.Term, m.Type, m.From, m.Term)
		}
		return nil
	}
	switch m.Type {
	case pb.MsgHup:
		if r.state != StateLeader {
			ents, err := r.raftLog.slice(r.raftLog.applied+1, r.raftLog.committed+1, noLimit)
			if err != nil {
				r.logger.Panicf("unexpected error getting unapplied entries (%v)", err)
			}
			if n := numOfPendingConf(ents); n != 0 && r.raftLog.committed > r.raftLog.applied {
				r.logger.Warningf("%x cannot campaign at term %d since there are still %d pending configuration changes to apply", r.id, r.Term, n)
				return nil
			}
			r.logger.Infof("%x is starting a new election at term %d", r.id, r.Term)
			if r.preVote {
				r.campaign(campaignPreElection)
			} else {
				r.campaign(campaignElection)
			}
		} else {
			r.logger.Debugf("%x ignoring MsgHup because already leader", r.id)
		}
	case pb.MsgVote, pb.MsgPreVote:
		if r.isLearner {
			r.logger.Infof("%x [logterm: %d, index: %d, vote: %x] ignored %s from %x [logterm: %d, index: %d] at term %d: learner can not vote", r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), r.Vote, m.Type, m.From, m.LogTerm, m.Index, r.Term)
			return nil
		}
		if (r.Vote == None || m.Term > r.Term || r.Vote == m.From) && r.raftLog.isUpToDate(m.Index, m.LogTerm) {
			r.logger.Infof("%x [logterm: %d, index: %d, vote: %x] cast %s for %x [logterm: %d, index: %d] at term %d", r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), r.Vote, m.Type, m.From, m.LogTerm, m.Index, r.Term)
			r.send(pb.Message{To: m.From, Term: m.Term, Type: voteRespMsgType(m.Type)})
			if m.Type == pb.MsgVote {
				r.electionElapsed = 0
				r.Vote = m.From
			}
		} else {
			r.logger.Infof("%x [logterm: %d, index: %d, vote: %x] rejected %s from %x [logterm: %d, index: %d] at term %d", r.id, r.raftLog.lastTerm(), r.raftLog.lastIndex(), r.Vote, m.Type, m.From, m.LogTerm, m.Index, r.Term)
			r.send(pb.Message{To: m.From, Term: r.Term, Type: voteRespMsgType(m.Type), Reject: true})
		}
	default:
		r.step(r, m)
	}
	return nil
}

type stepFunc func(r *raft, m pb.Message)

func stepLeader(r *raft, m pb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch m.Type {
	case pb.MsgBeat:
		r.bcastHeartbeat()
		return
	case pb.MsgCheckQuorum:
		if !r.checkQuorumActive() {
			r.logger.Warningf("%x stepped down to follower since quorum is not active", r.id)
			r.becomeFollower(r.Term, None)
		}
		return
	case pb.MsgProp:
		if len(m.Entries) == 0 {
			r.logger.Panicf("%x stepped empty MsgProp", r.id)
		}
		if _, ok := r.prs[r.id]; !ok {
			return
		}
		if r.leadTransferee != None {
			r.logger.Debugf("%x [term %d] transfer leadership to %x is in progress; dropping proposal", r.id, r.Term, r.leadTransferee)
			return
		}
		for i, e := range m.Entries {
			if e.Type == pb.EntryConfChange {
				if r.pendingConf {
					r.logger.Infof("propose conf %s ignored since pending unapplied configuration", e.String())
					m.Entries[i] = pb.Entry{Type: pb.EntryNormal}
				}
				r.pendingConf = true
			}
		}
		r.appendEntry(m.Entries...)
		r.bcastAppend()
		return
	case pb.MsgReadIndex:
		if r.quorum() > 1 {
			if r.raftLog.zeroTermOnErrCompacted(r.raftLog.term(r.raftLog.committed)) != r.Term {
				return
			}
			switch r.readOnly.option {
			case ReadOnlySafe:
				r.readOnly.addRequest(r.raftLog.committed, m)
				r.bcastHeartbeatWithCtx(m.Entries[0].Data)
			case ReadOnlyLeaseBased:
				ri := r.raftLog.committed
				if m.From == None || m.From == r.id {
					r.readStates = append(r.readStates, ReadState{Index: r.raftLog.committed, RequestCtx: m.Entries[0].Data})
				} else {
					r.send(pb.Message{To: m.From, Type: pb.MsgReadIndexResp, Index: ri, Entries: m.Entries})
				}
			}
		} else {
			r.readStates = append(r.readStates, ReadState{Index: r.raftLog.committed, RequestCtx: m.Entries[0].Data})
		}
		return
	}
	pr := r.getProgress(m.From)
	if pr == nil {
		r.logger.Debugf("%x no progress available for %x", r.id, m.From)
		return
	}
	switch m.Type {
	case pb.MsgAppResp:
		pr.RecentActive = true
		if m.Reject {
			r.logger.Debugf("%x received msgApp rejection(lastindex: %d) from %x for index %d", r.id, m.RejectHint, m.From, m.Index)
			if pr.maybeDecrTo(m.Index, m.RejectHint) {
				r.logger.Debugf("%x decreased progress of %x to [%s]", r.id, m.From, pr)
				if pr.State == ProgressStateReplicate {
					pr.becomeProbe()
				}
				r.sendAppend(m.From)
			}
		} else {
			oldPaused := pr.IsPaused()
			if pr.maybeUpdate(m.Index) {
				switch {
				case pr.State == ProgressStateProbe:
					pr.becomeReplicate()
				case pr.State == ProgressStateSnapshot && pr.needSnapshotAbort():
					r.logger.Debugf("%x snapshot aborted, resumed sending replication messages to %x [%s]", r.id, m.From, pr)
					pr.becomeProbe()
				case pr.State == ProgressStateReplicate:
					pr.ins.freeTo(m.Index)
				}
				if r.maybeCommit() {
					r.bcastAppend()
				} else if oldPaused {
					r.sendAppend(m.From)
				}
				if m.From == r.leadTransferee && pr.Match == r.raftLog.lastIndex() {
					r.logger.Infof("%x sent MsgTimeoutNow to %x after received MsgAppResp", r.id, m.From)
					r.sendTimeoutNow(m.From)
				}
			}
		}
	case pb.MsgHeartbeatResp:
		pr.RecentActive = true
		pr.resume()
		if pr.State == ProgressStateReplicate && pr.ins.full() {
			pr.ins.freeFirstOne()
		}
		if pr.Match < r.raftLog.lastIndex() {
			r.sendAppend(m.From)
		}
		if r.readOnly.option != ReadOnlySafe || len(m.Context) == 0 {
			return
		}
		ackCount := r.readOnly.recvAck(m)
		if ackCount < r.quorum() {
			return
		}
		rss := r.readOnly.advance(m)
		for _, rs := range rss {
			req := rs.req
			if req.From == None || req.From == r.id {
				r.readStates = append(r.readStates, ReadState{Index: rs.index, RequestCtx: req.Entries[0].Data})
			} else {
				r.send(pb.Message{To: req.From, Type: pb.MsgReadIndexResp, Index: rs.index, Entries: req.Entries})
			}
		}
	case pb.MsgSnapStatus:
		if pr.State != ProgressStateSnapshot {
			return
		}
		if !m.Reject {
			pr.becomeProbe()
			r.logger.Debugf("%x snapshot succeeded, resumed sending replication messages to %x [%s]", r.id, m.From, pr)
		} else {
			pr.snapshotFailure()
			pr.becomeProbe()
			r.logger.Debugf("%x snapshot failed, resumed sending replication messages to %x [%s]", r.id, m.From, pr)
		}
		pr.pause()
	case pb.MsgUnreachable:
		if pr.State == ProgressStateReplicate {
			pr.becomeProbe()
		}
		r.logger.Debugf("%x failed to send message to %x because it is unreachable [%s]", r.id, m.From, pr)
	case pb.MsgTransferLeader:
		if pr.IsLearner {
			r.logger.Debugf("%x is learner. Ignored transferring leadership", r.id)
			return
		}
		leadTransferee := m.From
		lastLeadTransferee := r.leadTransferee
		if lastLeadTransferee != None {
			if lastLeadTransferee == leadTransferee {
				r.logger.Infof("%x [term %d] transfer leadership to %x is in progress, ignores request to same node %x", r.id, r.Term, leadTransferee, leadTransferee)
				return
			}
			r.abortLeaderTransfer()
			r.logger.Infof("%x [term %d] abort previous transferring leadership to %x", r.id, r.Term, lastLeadTransferee)
		}
		if leadTransferee == r.id {
			r.logger.Debugf("%x is already leader. Ignored transferring leadership to self", r.id)
			return
		}
		r.logger.Infof("%x [term %d] starts to transfer leadership to %x", r.id, r.Term, leadTransferee)
		r.electionElapsed = 0
		r.leadTransferee = leadTransferee
		if pr.Match == r.raftLog.lastIndex() {
			r.sendTimeoutNow(leadTransferee)
			r.logger.Infof("%x sends MsgTimeoutNow to %x immediately as %x already has up-to-date log", r.id, leadTransferee, leadTransferee)
		} else {
			r.sendAppend(leadTransferee)
		}
	}
}
func stepCandidate(r *raft, m pb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var myVoteRespType pb.MessageType
	if r.state == StatePreCandidate {
		myVoteRespType = pb.MsgPreVoteResp
	} else {
		myVoteRespType = pb.MsgVoteResp
	}
	switch m.Type {
	case pb.MsgProp:
		r.logger.Infof("%x no leader at term %d; dropping proposal", r.id, r.Term)
		return
	case pb.MsgApp:
		r.becomeFollower(r.Term, m.From)
		r.handleAppendEntries(m)
	case pb.MsgHeartbeat:
		r.becomeFollower(r.Term, m.From)
		r.handleHeartbeat(m)
	case pb.MsgSnap:
		r.becomeFollower(m.Term, m.From)
		r.handleSnapshot(m)
	case myVoteRespType:
		gr := r.poll(m.From, m.Type, !m.Reject)
		r.logger.Infof("%x [quorum:%d] has received %d %s votes and %d vote rejections", r.id, r.quorum(), gr, m.Type, len(r.votes)-gr)
		switch r.quorum() {
		case gr:
			if r.state == StatePreCandidate {
				r.campaign(campaignElection)
			} else {
				r.becomeLeader()
				r.bcastAppend()
			}
		case len(r.votes) - gr:
			r.becomeFollower(r.Term, None)
		}
	case pb.MsgTimeoutNow:
		r.logger.Debugf("%x [term %d state %v] ignored MsgTimeoutNow from %x", r.id, r.Term, r.state, m.From)
	}
}
func stepFollower(r *raft, m pb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch m.Type {
	case pb.MsgProp:
		if r.lead == None {
			r.logger.Infof("%x no leader at term %d; dropping proposal", r.id, r.Term)
			return
		} else if r.disableProposalForwarding {
			r.logger.Infof("%x not forwarding to leader %x at term %d; dropping proposal", r.id, r.lead, r.Term)
			return
		}
		m.To = r.lead
		r.send(m)
	case pb.MsgApp:
		r.electionElapsed = 0
		r.lead = m.From
		r.handleAppendEntries(m)
	case pb.MsgHeartbeat:
		r.electionElapsed = 0
		r.lead = m.From
		r.handleHeartbeat(m)
	case pb.MsgSnap:
		r.electionElapsed = 0
		r.lead = m.From
		r.handleSnapshot(m)
	case pb.MsgTransferLeader:
		if r.lead == None {
			r.logger.Infof("%x no leader at term %d; dropping leader transfer msg", r.id, r.Term)
			return
		}
		m.To = r.lead
		r.send(m)
	case pb.MsgTimeoutNow:
		if r.promotable() {
			r.logger.Infof("%x [term %d] received MsgTimeoutNow from %x and starts an election to get leadership.", r.id, r.Term, m.From)
			r.campaign(campaignTransfer)
		} else {
			r.logger.Infof("%x received MsgTimeoutNow from %x but is not promotable", r.id, m.From)
		}
	case pb.MsgReadIndex:
		if r.lead == None {
			r.logger.Infof("%x no leader at term %d; dropping index reading msg", r.id, r.Term)
			return
		}
		m.To = r.lead
		r.send(m)
	case pb.MsgReadIndexResp:
		if len(m.Entries) != 1 {
			r.logger.Errorf("%x invalid format of MsgReadIndexResp from %x, entries count: %d", r.id, m.From, len(m.Entries))
			return
		}
		r.readStates = append(r.readStates, ReadState{Index: m.Index, RequestCtx: m.Entries[0].Data})
	}
}
func (r *raft) handleAppendEntries(m pb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.Index < r.raftLog.committed {
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: r.raftLog.committed})
		return
	}
	if mlastIndex, ok := r.raftLog.maybeAppend(m.Index, m.LogTerm, m.Commit, m.Entries...); ok {
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: mlastIndex})
	} else {
		r.logger.Debugf("%x [logterm: %d, index: %d] rejected msgApp [logterm: %d, index: %d] from %x", r.id, r.raftLog.zeroTermOnErrCompacted(r.raftLog.term(m.Index)), m.Index, m.LogTerm, m.Index, m.From)
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: m.Index, Reject: true, RejectHint: r.raftLog.lastIndex()})
	}
}
func (r *raft) handleHeartbeat(m pb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.raftLog.commitTo(m.Commit)
	r.send(pb.Message{To: m.From, Type: pb.MsgHeartbeatResp, Context: m.Context})
}
func (r *raft) handleSnapshot(m pb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sindex, sterm := m.Snapshot.Metadata.Index, m.Snapshot.Metadata.Term
	if r.restore(m.Snapshot) {
		r.logger.Infof("%x [commit: %d] restored snapshot [index: %d, term: %d]", r.id, r.raftLog.committed, sindex, sterm)
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: r.raftLog.lastIndex()})
	} else {
		r.logger.Infof("%x [commit: %d] ignored snapshot [index: %d, term: %d]", r.id, r.raftLog.committed, sindex, sterm)
		r.send(pb.Message{To: m.From, Type: pb.MsgAppResp, Index: r.raftLog.committed})
	}
}
func (r *raft) restore(s pb.Snapshot) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.Metadata.Index <= r.raftLog.committed {
		return false
	}
	if r.raftLog.matchTerm(s.Metadata.Index, s.Metadata.Term) {
		r.logger.Infof("%x [commit: %d, lastindex: %d, lastterm: %d] fast-forwarded commit to snapshot [index: %d, term: %d]", r.id, r.raftLog.committed, r.raftLog.lastIndex(), r.raftLog.lastTerm(), s.Metadata.Index, s.Metadata.Term)
		r.raftLog.commitTo(s.Metadata.Index)
		return false
	}
	if !r.isLearner {
		for _, id := range s.Metadata.ConfState.Learners {
			if id == r.id {
				r.logger.Errorf("%x can't become learner when restores snapshot [index: %d, term: %d]", r.id, s.Metadata.Index, s.Metadata.Term)
				return false
			}
		}
	}
	r.logger.Infof("%x [commit: %d, lastindex: %d, lastterm: %d] starts to restore snapshot [index: %d, term: %d]", r.id, r.raftLog.committed, r.raftLog.lastIndex(), r.raftLog.lastTerm(), s.Metadata.Index, s.Metadata.Term)
	r.raftLog.restore(s)
	r.prs = make(map[uint64]*Progress)
	r.learnerPrs = make(map[uint64]*Progress)
	r.restoreNode(s.Metadata.ConfState.Nodes, false)
	r.restoreNode(s.Metadata.ConfState.Learners, true)
	return true
}
func (r *raft) restoreNode(nodes []uint64, isLearner bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, n := range nodes {
		match, next := uint64(0), r.raftLog.lastIndex()+1
		if n == r.id {
			match = next - 1
			r.isLearner = isLearner
		}
		r.setProgress(n, match, next, isLearner)
		r.logger.Infof("%x restored progress of %x [%s]", r.id, n, r.getProgress(n))
	}
}
func (r *raft) promotable() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := r.prs[r.id]
	return ok
}
func (r *raft) addNode(id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.addNodeOrLearnerNode(id, false)
}
func (r *raft) addLearner(id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.addNodeOrLearnerNode(id, true)
}
func (r *raft) addNodeOrLearnerNode(id uint64, isLearner bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.pendingConf = false
	pr := r.getProgress(id)
	if pr == nil {
		r.setProgress(id, 0, r.raftLog.lastIndex()+1, isLearner)
	} else {
		if isLearner && !pr.IsLearner {
			r.logger.Infof("%x ignored addLeaner: do not support changing %x from raft peer to learner.", r.id, id)
			return
		}
		if isLearner == pr.IsLearner {
			return
		}
		delete(r.learnerPrs, id)
		pr.IsLearner = false
		r.prs[id] = pr
	}
	if r.id == id {
		r.isLearner = isLearner
	}
	pr = r.getProgress(id)
	pr.RecentActive = true
}
func (r *raft) removeNode(id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.delProgress(id)
	r.pendingConf = false
	if len(r.prs) == 0 && len(r.learnerPrs) == 0 {
		return
	}
	if r.maybeCommit() {
		r.bcastAppend()
	}
	if r.state == StateLeader && r.leadTransferee == id {
		r.abortLeaderTransfer()
	}
}
func (r *raft) resetPendingConf() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.pendingConf = false
}
func (r *raft) setProgress(id, match, next uint64, isLearner bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !isLearner {
		delete(r.learnerPrs, id)
		r.prs[id] = &Progress{Next: next, Match: match, ins: newInflights(r.maxInflight)}
		return
	}
	if _, ok := r.prs[id]; ok {
		panic(fmt.Sprintf("%x unexpected changing from voter to learner for %x", r.id, id))
	}
	r.learnerPrs[id] = &Progress{Next: next, Match: match, ins: newInflights(r.maxInflight), IsLearner: true}
}
func (r *raft) delProgress(id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delete(r.prs, id)
	delete(r.learnerPrs, id)
}
func (r *raft) loadState(state pb.HardState) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if state.Commit < r.raftLog.committed || state.Commit > r.raftLog.lastIndex() {
		r.logger.Panicf("%x state.commit %d is out of range [%d, %d]", r.id, state.Commit, r.raftLog.committed, r.raftLog.lastIndex())
	}
	r.raftLog.committed = state.Commit
	r.Term = state.Term
	r.Vote = state.Vote
}
func (r *raft) pastElectionTimeout() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.electionElapsed >= r.randomizedElectionTimeout
}
func (r *raft) resetRandomizedElectionTimeout() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.randomizedElectionTimeout = r.electionTimeout + globalRand.Intn(r.electionTimeout)
}
func (r *raft) checkQuorumActive() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var act int
	r.forEachProgress(func(id uint64, pr *Progress) {
		if id == r.id {
			act++
			return
		}
		if pr.RecentActive && !pr.IsLearner {
			act++
		}
		pr.RecentActive = false
	})
	return act >= r.quorum()
}
func (r *raft) sendTimeoutNow(to uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.send(pb.Message{To: to, Type: pb.MsgTimeoutNow})
}
func (r *raft) abortLeaderTransfer() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.leadTransferee = None
}
func numOfPendingConf(ents []pb.Entry) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := 0
	for i := range ents {
		if ents[i].Type == pb.EntryConfChange {
			n++
		}
	}
	return n
}
