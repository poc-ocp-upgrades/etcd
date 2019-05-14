package etcdserver

import (
	"encoding/json"
	"expvar"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/contention"
	"github.com/coreos/etcd/pkg/pbutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/rafthttp"
	"github.com/coreos/etcd/wal"
	"github.com/coreos/etcd/wal/walpb"
	"github.com/coreos/pkg/capnslog"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numberOfCatchUpEntries = 5000
	maxSizePerMsg          = 1 * 1024 * 1024
	maxInflightMsgs        = 4096 / 8
)

var (
	raftStatusMu sync.Mutex
	raftStatus   func() raft.Status
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	raft.SetLogger(capnslog.NewPackageLogger("github.com/coreos/etcd", "raft"))
	expvar.Publish("raft.status", expvar.Func(func() interface{} {
		raftStatusMu.Lock()
		defer raftStatusMu.Unlock()
		return raftStatus()
	}))
}

type RaftTimer interface {
	Index() uint64
	Term() uint64
}
type apply struct {
	entries  []raftpb.Entry
	snapshot raftpb.Snapshot
	notifyc  chan struct{}
}
type raftNode struct {
	index  uint64
	term   uint64
	lead   uint64
	tickMu *sync.Mutex
	raftNodeConfig
	msgSnapC   chan raftpb.Message
	applyc     chan apply
	readStateC chan raft.ReadState
	ticker     *time.Ticker
	td         *contention.TimeoutDetector
	stopped    chan struct{}
	done       chan struct{}
}
type raftNodeConfig struct {
	isIDRemoved func(id uint64) bool
	raft.Node
	raftStorage *raft.MemoryStorage
	storage     Storage
	heartbeat   time.Duration
	transport   rafthttp.Transporter
}

func newRaftNode(cfg raftNodeConfig) *raftNode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := &raftNode{tickMu: new(sync.Mutex), raftNodeConfig: cfg, td: contention.NewTimeoutDetector(2 * cfg.heartbeat), readStateC: make(chan raft.ReadState, 1), msgSnapC: make(chan raftpb.Message, maxInFlightMsgSnap), applyc: make(chan apply), stopped: make(chan struct{}), done: make(chan struct{})}
	if r.heartbeat == 0 {
		r.ticker = &time.Ticker{}
	} else {
		r.ticker = time.NewTicker(r.heartbeat)
	}
	return r
}
func (r *raftNode) tick() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.tickMu.Lock()
	r.Tick()
	r.tickMu.Unlock()
}
func (r *raftNode) start(rh *raftReadyHandler) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	internalTimeout := time.Second
	go func() {
		defer r.onStop()
		islead := false
		for {
			select {
			case <-r.ticker.C:
				r.tick()
			case rd := <-r.Ready():
				if rd.SoftState != nil {
					newLeader := rd.SoftState.Lead != raft.None && atomic.LoadUint64(&r.lead) != rd.SoftState.Lead
					if newLeader {
						leaderChanges.Inc()
					}
					if rd.SoftState.Lead == raft.None {
						hasLeader.Set(0)
					} else {
						hasLeader.Set(1)
					}
					atomic.StoreUint64(&r.lead, rd.SoftState.Lead)
					islead = rd.RaftState == raft.StateLeader
					if islead {
						isLeader.Set(1)
					} else {
						isLeader.Set(0)
					}
					rh.updateLeadership(newLeader)
					r.td.Reset()
				}
				if len(rd.ReadStates) != 0 {
					select {
					case r.readStateC <- rd.ReadStates[len(rd.ReadStates)-1]:
					case <-time.After(internalTimeout):
						plog.Warningf("timed out sending read state")
					case <-r.stopped:
						return
					}
				}
				notifyc := make(chan struct{}, 1)
				ap := apply{entries: rd.CommittedEntries, snapshot: rd.Snapshot, notifyc: notifyc}
				updateCommittedIndex(&ap, rh)
				select {
				case r.applyc <- ap:
				case <-r.stopped:
					return
				}
				if islead {
					r.transport.Send(r.processMessages(rd.Messages))
				}
				if err := r.storage.Save(rd.HardState, rd.Entries); err != nil {
					plog.Fatalf("raft save state and entries error: %v", err)
				}
				if !raft.IsEmptyHardState(rd.HardState) {
					proposalsCommitted.Set(float64(rd.HardState.Commit))
				}
				if !raft.IsEmptySnap(rd.Snapshot) {
					if err := r.storage.SaveSnap(rd.Snapshot); err != nil {
						plog.Fatalf("raft save snapshot error: %v", err)
					}
					notifyc <- struct{}{}
					r.raftStorage.ApplySnapshot(rd.Snapshot)
					plog.Infof("raft applied incoming snapshot at index %d", rd.Snapshot.Metadata.Index)
				}
				r.raftStorage.Append(rd.Entries)
				if !islead {
					msgs := r.processMessages(rd.Messages)
					notifyc <- struct{}{}
					waitApply := false
					for _, ent := range rd.CommittedEntries {
						if ent.Type == raftpb.EntryConfChange {
							waitApply = true
							break
						}
					}
					if waitApply {
						select {
						case notifyc <- struct{}{}:
						case <-r.stopped:
							return
						}
					}
					r.transport.Send(msgs)
				} else {
					notifyc <- struct{}{}
				}
				r.Advance()
			case <-r.stopped:
				return
			}
		}
	}()
}
func updateCommittedIndex(ap *apply, rh *raftReadyHandler) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ci uint64
	if len(ap.entries) != 0 {
		ci = ap.entries[len(ap.entries)-1].Index
	}
	if ap.snapshot.Metadata.Index > ci {
		ci = ap.snapshot.Metadata.Index
	}
	if ci != 0 {
		rh.updateCommittedIndex(ci)
	}
}
func (r *raftNode) processMessages(ms []raftpb.Message) []raftpb.Message {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sentAppResp := false
	for i := len(ms) - 1; i >= 0; i-- {
		if r.isIDRemoved(ms[i].To) {
			ms[i].To = 0
		}
		if ms[i].Type == raftpb.MsgAppResp {
			if sentAppResp {
				ms[i].To = 0
			} else {
				sentAppResp = true
			}
		}
		if ms[i].Type == raftpb.MsgSnap {
			select {
			case r.msgSnapC <- ms[i]:
			default:
			}
			ms[i].To = 0
		}
		if ms[i].Type == raftpb.MsgHeartbeat {
			ok, exceed := r.td.Observe(ms[i].To)
			if !ok {
				plog.Warningf("failed to send out heartbeat on time (exceeded the %v timeout for %v)", r.heartbeat, exceed)
				plog.Warningf("server is likely overloaded")
				heartbeatSendFailures.Inc()
			}
		}
	}
	return ms
}
func (r *raftNode) apply() chan apply {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.applyc
}
func (r *raftNode) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.stopped <- struct{}{}
	<-r.done
}
func (r *raftNode) onStop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.Stop()
	r.ticker.Stop()
	r.transport.Stop()
	if err := r.storage.Close(); err != nil {
		plog.Panicf("raft close storage error: %v", err)
	}
	close(r.done)
}
func (r *raftNode) pauseSending() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := r.transport.(rafthttp.Pausable)
	p.Pause()
}
func (r *raftNode) resumeSending() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := r.transport.(rafthttp.Pausable)
	p.Resume()
}
func (r *raftNode) advanceTicks(ticks int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < ticks; i++ {
		r.tick()
	}
}
func startNode(cfg ServerConfig, cl *membership.RaftCluster, ids []types.ID) (id types.ID, n raft.Node, s *raft.MemoryStorage, w *wal.WAL) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	member := cl.MemberByName(cfg.Name)
	metadata := pbutil.MustMarshal(&pb.Metadata{NodeID: uint64(member.ID), ClusterID: uint64(cl.ID())})
	if w, err = wal.Create(cfg.WALDir(), metadata); err != nil {
		plog.Fatalf("create wal error: %v", err)
	}
	peers := make([]raft.Peer, len(ids))
	for i, id := range ids {
		ctx, err := json.Marshal((*cl).Member(id))
		if err != nil {
			plog.Panicf("marshal member should never fail: %v", err)
		}
		peers[i] = raft.Peer{ID: uint64(id), Context: ctx}
	}
	id = member.ID
	plog.Infof("starting member %s in cluster %s", id, cl.ID())
	s = raft.NewMemoryStorage()
	c := &raft.Config{ID: uint64(id), ElectionTick: cfg.ElectionTicks, HeartbeatTick: 1, Storage: s, MaxSizePerMsg: maxSizePerMsg, MaxInflightMsgs: maxInflightMsgs, CheckQuorum: true}
	n = raft.StartNode(c, peers)
	raftStatusMu.Lock()
	raftStatus = n.Status
	raftStatusMu.Unlock()
	return id, n, s, w
}
func restartNode(cfg ServerConfig, snapshot *raftpb.Snapshot) (types.ID, *membership.RaftCluster, raft.Node, *raft.MemoryStorage, *wal.WAL) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var walsnap walpb.Snapshot
	if snapshot != nil {
		walsnap.Index, walsnap.Term = snapshot.Metadata.Index, snapshot.Metadata.Term
	}
	w, id, cid, st, ents := readWAL(cfg.WALDir(), walsnap)
	plog.Infof("restarting member %s in cluster %s at commit index %d", id, cid, st.Commit)
	cl := membership.NewCluster("")
	cl.SetID(cid)
	s := raft.NewMemoryStorage()
	if snapshot != nil {
		s.ApplySnapshot(*snapshot)
	}
	s.SetHardState(st)
	s.Append(ents)
	c := &raft.Config{ID: uint64(id), ElectionTick: cfg.ElectionTicks, HeartbeatTick: 1, Storage: s, MaxSizePerMsg: maxSizePerMsg, MaxInflightMsgs: maxInflightMsgs, CheckQuorum: true}
	n := raft.RestartNode(c)
	raftStatusMu.Lock()
	raftStatus = n.Status
	raftStatusMu.Unlock()
	return id, cl, n, s, w
}
func restartAsStandaloneNode(cfg ServerConfig, snapshot *raftpb.Snapshot) (types.ID, *membership.RaftCluster, raft.Node, *raft.MemoryStorage, *wal.WAL) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var walsnap walpb.Snapshot
	if snapshot != nil {
		walsnap.Index, walsnap.Term = snapshot.Metadata.Index, snapshot.Metadata.Term
	}
	w, id, cid, st, ents := readWAL(cfg.WALDir(), walsnap)
	for i, ent := range ents {
		if ent.Index > st.Commit {
			plog.Infof("discarding %d uncommitted WAL entries ", len(ents)-i)
			ents = ents[:i]
			break
		}
	}
	toAppEnts := createConfigChangeEnts(getIDs(snapshot, ents), uint64(id), st.Term, st.Commit)
	ents = append(ents, toAppEnts...)
	err := w.Save(raftpb.HardState{}, toAppEnts)
	if err != nil {
		plog.Fatalf("%v", err)
	}
	if len(ents) != 0 {
		st.Commit = ents[len(ents)-1].Index
	}
	plog.Printf("forcing restart of member %s in cluster %s at commit index %d", id, cid, st.Commit)
	cl := membership.NewCluster("")
	cl.SetID(cid)
	s := raft.NewMemoryStorage()
	if snapshot != nil {
		s.ApplySnapshot(*snapshot)
	}
	s.SetHardState(st)
	s.Append(ents)
	c := &raft.Config{ID: uint64(id), ElectionTick: cfg.ElectionTicks, HeartbeatTick: 1, Storage: s, MaxSizePerMsg: maxSizePerMsg, MaxInflightMsgs: maxInflightMsgs, CheckQuorum: true}
	n := raft.RestartNode(c)
	raftStatus = n.Status
	return id, cl, n, s, w
}
func getIDs(snap *raftpb.Snapshot, ents []raftpb.Entry) []uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ids := make(map[uint64]bool)
	if snap != nil {
		for _, id := range snap.Metadata.ConfState.Nodes {
			ids[id] = true
		}
	}
	for _, e := range ents {
		if e.Type != raftpb.EntryConfChange {
			continue
		}
		var cc raftpb.ConfChange
		pbutil.MustUnmarshal(&cc, e.Data)
		switch cc.Type {
		case raftpb.ConfChangeAddNode:
			ids[cc.NodeID] = true
		case raftpb.ConfChangeRemoveNode:
			delete(ids, cc.NodeID)
		case raftpb.ConfChangeUpdateNode:
		default:
			plog.Panicf("ConfChange Type should be either ConfChangeAddNode or ConfChangeRemoveNode!")
		}
	}
	sids := make(types.Uint64Slice, 0, len(ids))
	for id := range ids {
		sids = append(sids, id)
	}
	sort.Sort(sids)
	return []uint64(sids)
}
func createConfigChangeEnts(ids []uint64, self uint64, term, index uint64) []raftpb.Entry {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ents := make([]raftpb.Entry, 0)
	next := index + 1
	found := false
	for _, id := range ids {
		if id == self {
			found = true
			continue
		}
		cc := &raftpb.ConfChange{Type: raftpb.ConfChangeRemoveNode, NodeID: id}
		e := raftpb.Entry{Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(cc), Term: term, Index: next}
		ents = append(ents, e)
		next++
	}
	if !found {
		m := membership.Member{ID: types.ID(self), RaftAttributes: membership.RaftAttributes{PeerURLs: []string{"http://localhost:2380"}}}
		ctx, err := json.Marshal(m)
		if err != nil {
			plog.Panicf("marshal member should never fail: %v", err)
		}
		cc := &raftpb.ConfChange{Type: raftpb.ConfChangeAddNode, NodeID: self, Context: ctx}
		e := raftpb.Entry{Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(cc), Term: term, Index: next}
		ents = append(ents, e)
	}
	return ents
}
