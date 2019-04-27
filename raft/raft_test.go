package raft

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"testing"
	pb "github.com/coreos/etcd/raft/raftpb"
)

func nextEnts(r *raft, s *MemoryStorage) (ents []pb.Entry) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Append(r.raftLog.unstableEntries())
	r.raftLog.stableTo(r.raftLog.lastIndex(), r.raftLog.lastTerm())
	ents = r.raftLog.nextEnts()
	r.raftLog.appliedTo(r.raftLog.committed)
	return ents
}

type stateMachine interface {
	Step(m pb.Message) error
	readMessages() []pb.Message
}

func (r *raft) readMessages() []pb.Message {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	msgs := r.msgs
	r.msgs = make([]pb.Message, 0)
	return msgs
}
func TestProgressBecomeProbe(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	match := uint64(1)
	tests := []struct {
		p	*Progress
		wnext	uint64
	}{{&Progress{State: ProgressStateReplicate, Match: match, Next: 5, ins: newInflights(256)}, 2}, {&Progress{State: ProgressStateSnapshot, Match: match, Next: 5, PendingSnapshot: 10, ins: newInflights(256)}, 11}, {&Progress{State: ProgressStateSnapshot, Match: match, Next: 5, PendingSnapshot: 0, ins: newInflights(256)}, 2}}
	for i, tt := range tests {
		tt.p.becomeProbe()
		if tt.p.State != ProgressStateProbe {
			t.Errorf("#%d: state = %s, want %s", i, tt.p.State, ProgressStateProbe)
		}
		if tt.p.Match != match {
			t.Errorf("#%d: match = %d, want %d", i, tt.p.Match, match)
		}
		if tt.p.Next != tt.wnext {
			t.Errorf("#%d: next = %d, want %d", i, tt.p.Next, tt.wnext)
		}
	}
}
func TestProgressBecomeReplicate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &Progress{State: ProgressStateProbe, Match: 1, Next: 5, ins: newInflights(256)}
	p.becomeReplicate()
	if p.State != ProgressStateReplicate {
		t.Errorf("state = %s, want %s", p.State, ProgressStateReplicate)
	}
	if p.Match != 1 {
		t.Errorf("match = %d, want 1", p.Match)
	}
	if w := p.Match + 1; p.Next != w {
		t.Errorf("next = %d, want %d", p.Next, w)
	}
}
func TestProgressBecomeSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &Progress{State: ProgressStateProbe, Match: 1, Next: 5, ins: newInflights(256)}
	p.becomeSnapshot(10)
	if p.State != ProgressStateSnapshot {
		t.Errorf("state = %s, want %s", p.State, ProgressStateSnapshot)
	}
	if p.Match != 1 {
		t.Errorf("match = %d, want 1", p.Match)
	}
	if p.PendingSnapshot != 10 {
		t.Errorf("pendingSnapshot = %d, want 10", p.PendingSnapshot)
	}
}
func TestProgressUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prevM, prevN := uint64(3), uint64(5)
	tests := []struct {
		update	uint64
		wm	uint64
		wn	uint64
		wok	bool
	}{{prevM - 1, prevM, prevN, false}, {prevM, prevM, prevN, false}, {prevM + 1, prevM + 1, prevN, true}, {prevM + 2, prevM + 2, prevN + 1, true}}
	for i, tt := range tests {
		p := &Progress{Match: prevM, Next: prevN}
		ok := p.maybeUpdate(tt.update)
		if ok != tt.wok {
			t.Errorf("#%d: ok= %v, want %v", i, ok, tt.wok)
		}
		if p.Match != tt.wm {
			t.Errorf("#%d: match= %d, want %d", i, p.Match, tt.wm)
		}
		if p.Next != tt.wn {
			t.Errorf("#%d: next= %d, want %d", i, p.Next, tt.wn)
		}
	}
}
func TestProgressMaybeDecr(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		state		ProgressStateType
		m		uint64
		n		uint64
		rejected	uint64
		last		uint64
		w		bool
		wn		uint64
	}{{ProgressStateReplicate, 5, 10, 5, 5, false, 10}, {ProgressStateReplicate, 5, 10, 4, 4, false, 10}, {ProgressStateReplicate, 5, 10, 9, 9, true, 6}, {ProgressStateProbe, 0, 0, 0, 0, false, 0}, {ProgressStateProbe, 0, 10, 5, 5, false, 10}, {ProgressStateProbe, 0, 10, 9, 9, true, 9}, {ProgressStateProbe, 0, 2, 1, 1, true, 1}, {ProgressStateProbe, 0, 1, 0, 0, true, 1}, {ProgressStateProbe, 0, 10, 9, 2, true, 3}, {ProgressStateProbe, 0, 10, 9, 0, true, 1}}
	for i, tt := range tests {
		p := &Progress{State: tt.state, Match: tt.m, Next: tt.n}
		if g := p.maybeDecrTo(tt.rejected, tt.last); g != tt.w {
			t.Errorf("#%d: maybeDecrTo= %t, want %t", i, g, tt.w)
		}
		if gm := p.Match; gm != tt.m {
			t.Errorf("#%d: match= %d, want %d", i, gm, tt.m)
		}
		if gn := p.Next; gn != tt.wn {
			t.Errorf("#%d: next= %d, want %d", i, gn, tt.wn)
		}
	}
}
func TestProgressIsPaused(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		state	ProgressStateType
		paused	bool
		w	bool
	}{{ProgressStateProbe, false, false}, {ProgressStateProbe, true, true}, {ProgressStateReplicate, false, false}, {ProgressStateReplicate, true, false}, {ProgressStateSnapshot, false, true}, {ProgressStateSnapshot, true, true}}
	for i, tt := range tests {
		p := &Progress{State: tt.state, Paused: tt.paused, ins: newInflights(256)}
		if g := p.IsPaused(); g != tt.w {
			t.Errorf("#%d: paused= %t, want %t", i, g, tt.w)
		}
	}
}
func TestProgressResume(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &Progress{Next: 2, Paused: true}
	p.maybeDecrTo(1, 1)
	if p.Paused {
		t.Errorf("paused= %v, want false", p.Paused)
	}
	p.Paused = true
	p.maybeUpdate(2)
	if p.Paused {
		t.Errorf("paused= %v, want false", p.Paused)
	}
}
func TestProgressResumeByHeartbeatResp(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 5, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	r.prs[2].Paused = true
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgBeat})
	if !r.prs[2].Paused {
		t.Errorf("paused = %v, want true", r.prs[2].Paused)
	}
	r.prs[2].becomeReplicate()
	r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgHeartbeatResp})
	if r.prs[2].Paused {
		t.Errorf("paused = %v, want false", r.prs[2].Paused)
	}
}
func TestProgressPaused(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 5, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
	ms := r.readMessages()
	if len(ms) != 1 {
		t.Errorf("len(ms) = %d, want 1", len(ms))
	}
}
func TestLeaderElection(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testLeaderElection(t, false)
}
func TestLeaderElectionPreVote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testLeaderElection(t, true)
}
func testLeaderElection(t *testing.T, preVote bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cfg func(*Config)
	if preVote {
		cfg = preVoteConfig
	}
	tests := []struct {
		*network
		state	StateType
		expTerm	uint64
	}{{newNetworkWithConfig(cfg, nil, nil, nil), StateLeader, 1}, {newNetworkWithConfig(cfg, nil, nil, nopStepper), StateLeader, 1}, {newNetworkWithConfig(cfg, nil, nopStepper, nopStepper), StateCandidate, 1}, {newNetworkWithConfig(cfg, nil, nopStepper, nopStepper, nil), StateCandidate, 1}, {newNetworkWithConfig(cfg, nil, nopStepper, nopStepper, nil, nil), StateLeader, 1}, {newNetworkWithConfig(cfg, nil, entsWithConfig(cfg, 1), entsWithConfig(cfg, 1), entsWithConfig(cfg, 1, 1), nil), StateFollower, 1}}
	for i, tt := range tests {
		tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
		sm := tt.network.peers[1].(*raft)
		var expState StateType
		var expTerm uint64
		if tt.state == StateCandidate && preVote {
			expState = StatePreCandidate
			expTerm = 0
		} else {
			expState = tt.state
			expTerm = tt.expTerm
		}
		if sm.state != expState {
			t.Errorf("#%d: state = %s, want %s", i, sm.state, expState)
		}
		if g := sm.Term; g != expTerm {
			t.Errorf("#%d: term = %d, want %d", i, g, expTerm)
		}
	}
}
func TestLearnerElectionTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	n1 := newTestLearnerRaft(1, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	n2 := newTestLearnerRaft(2, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	n1.becomeFollower(1, None)
	n2.becomeFollower(1, None)
	setRandomizedElectionTimeout(n2, n2.electionTimeout)
	for i := 0; i < n2.electionTimeout; i++ {
		n2.tick()
	}
	if n2.state != StateFollower {
		t.Errorf("peer 2 state: %s, want %s", n2.state, StateFollower)
	}
}
func TestLearnerPromotion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	n1 := newTestLearnerRaft(1, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	n2 := newTestLearnerRaft(2, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	n1.becomeFollower(1, None)
	n2.becomeFollower(1, None)
	nt := newNetwork(n1, n2)
	if n1.state == StateLeader {
		t.Error("peer 1 state is leader, want not", n1.state)
	}
	setRandomizedElectionTimeout(n1, n1.electionTimeout)
	for i := 0; i < n1.electionTimeout; i++ {
		n1.tick()
	}
	if n1.state != StateLeader {
		t.Errorf("peer 1 state: %s, want %s", n1.state, StateLeader)
	}
	if n2.state != StateFollower {
		t.Errorf("peer 2 state: %s, want %s", n2.state, StateFollower)
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgBeat})
	n1.addNode(2)
	n2.addNode(2)
	if n2.isLearner {
		t.Error("peer 2 is learner, want not")
	}
	setRandomizedElectionTimeout(n2, n2.electionTimeout)
	for i := 0; i < n2.electionTimeout; i++ {
		n2.tick()
	}
	nt.send(pb.Message{From: 2, To: 2, Type: pb.MsgBeat})
	if n1.state != StateFollower {
		t.Errorf("peer 1 state: %s, want %s", n1.state, StateFollower)
	}
	if n2.state != StateLeader {
		t.Errorf("peer 2 state: %s, want %s", n2.state, StateLeader)
	}
}
func TestLearnerCannotVote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	n2 := newTestLearnerRaft(2, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	n2.becomeFollower(1, None)
	n2.Step(pb.Message{From: 1, To: 2, Term: 2, Type: pb.MsgVote, LogTerm: 11, Index: 11})
	if len(n2.msgs) != 0 {
		t.Errorf("expect learner not to vote, but received %v messages", n2.msgs)
	}
}
func TestLeaderCycle(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testLeaderCycle(t, false)
}
func TestLeaderCyclePreVote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testLeaderCycle(t, true)
}
func testLeaderCycle(t *testing.T, preVote bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cfg func(*Config)
	if preVote {
		cfg = preVoteConfig
	}
	n := newNetworkWithConfig(cfg, nil, nil, nil)
	for campaignerID := uint64(1); campaignerID <= 3; campaignerID++ {
		n.send(pb.Message{From: campaignerID, To: campaignerID, Type: pb.MsgHup})
		for _, peer := range n.peers {
			sm := peer.(*raft)
			if sm.id == campaignerID && sm.state != StateLeader {
				t.Errorf("preVote=%v: campaigning node %d state = %v, want StateLeader", preVote, sm.id, sm.state)
			} else if sm.id != campaignerID && sm.state != StateFollower {
				t.Errorf("preVote=%v: after campaign of node %d, "+"node %d had state = %v, want StateFollower", preVote, campaignerID, sm.id, sm.state)
			}
		}
	}
}
func TestLeaderElectionOverwriteNewerLogs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testLeaderElectionOverwriteNewerLogs(t, false)
}
func TestLeaderElectionOverwriteNewerLogsPreVote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testLeaderElectionOverwriteNewerLogs(t, true)
}
func testLeaderElectionOverwriteNewerLogs(t *testing.T, preVote bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cfg func(*Config)
	if preVote {
		cfg = preVoteConfig
	}
	n := newNetworkWithConfig(cfg, entsWithConfig(cfg, 1), entsWithConfig(cfg, 1), entsWithConfig(cfg, 2), votedWithConfig(cfg, 3, 2), votedWithConfig(cfg, 3, 2))
	n.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	sm1 := n.peers[1].(*raft)
	if sm1.state != StateFollower {
		t.Errorf("state = %s, want StateFollower", sm1.state)
	}
	if sm1.Term != 2 {
		t.Errorf("term = %d, want 2", sm1.Term)
	}
	n.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	if sm1.state != StateLeader {
		t.Errorf("state = %s, want StateLeader", sm1.state)
	}
	if sm1.Term != 3 {
		t.Errorf("term = %d, want 3", sm1.Term)
	}
	for i := range n.peers {
		sm := n.peers[i].(*raft)
		entries := sm.raftLog.allEntries()
		if len(entries) != 2 {
			t.Fatalf("node %d: len(entries) == %d, want 2", i, len(entries))
		}
		if entries[0].Term != 1 {
			t.Errorf("node %d: term at index 1 == %d, want 1", i, entries[0].Term)
		}
		if entries[1].Term != 3 {
			t.Errorf("node %d: term at index 2 == %d, want 3", i, entries[1].Term)
		}
	}
}
func TestVoteFromAnyState(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testVoteFromAnyState(t, pb.MsgVote)
}
func TestPreVoteFromAnyState(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testVoteFromAnyState(t, pb.MsgPreVote)
}
func testVoteFromAnyState(t *testing.T, vt pb.MessageType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for st := StateType(0); st < numStates; st++ {
		r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
		r.Term = 1
		switch st {
		case StateFollower:
			r.becomeFollower(r.Term, 3)
		case StatePreCandidate:
			r.becomePreCandidate()
		case StateCandidate:
			r.becomeCandidate()
		case StateLeader:
			r.becomeCandidate()
			r.becomeLeader()
		}
		origTerm := r.Term
		newTerm := r.Term + 1
		msg := pb.Message{From: 2, To: 1, Type: vt, Term: newTerm, LogTerm: newTerm, Index: 42}
		if err := r.Step(msg); err != nil {
			t.Errorf("%s,%s: Step failed: %s", vt, st, err)
		}
		if len(r.msgs) != 1 {
			t.Errorf("%s,%s: %d response messages, want 1: %+v", vt, st, len(r.msgs), r.msgs)
		} else {
			resp := r.msgs[0]
			if resp.Type != voteRespMsgType(vt) {
				t.Errorf("%s,%s: response message is %s, want %s", vt, st, resp.Type, voteRespMsgType(vt))
			}
			if resp.Reject {
				t.Errorf("%s,%s: unexpected rejection", vt, st)
			}
		}
		if vt == pb.MsgVote {
			if r.state != StateFollower {
				t.Errorf("%s,%s: state %s, want %s", vt, st, r.state, StateFollower)
			}
			if r.Term != newTerm {
				t.Errorf("%s,%s: term %d, want %d", vt, st, r.Term, newTerm)
			}
			if r.Vote != 2 {
				t.Errorf("%s,%s: vote %d, want 2", vt, st, r.Vote)
			}
		} else {
			if r.state != st {
				t.Errorf("%s,%s: state %s, want %s", vt, st, r.state, st)
			}
			if r.Term != origTerm {
				t.Errorf("%s,%s: term %d, want %d", vt, st, r.Term, origTerm)
			}
			if r.Vote != None && r.Vote != 1 {
				t.Errorf("%s,%s: vote %d, want %d or 1", vt, st, r.Vote, None)
			}
		}
	}
}
func TestLogReplication(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		*network
		msgs		[]pb.Message
		wcommitted	uint64
	}{{newNetwork(nil, nil, nil), []pb.Message{{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}}}, 2}, {newNetwork(nil, nil, nil), []pb.Message{{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}}, {From: 1, To: 2, Type: pb.MsgHup}, {From: 1, To: 2, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}}}, 4}}
	for i, tt := range tests {
		tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
		for _, m := range tt.msgs {
			tt.send(m)
		}
		for j, x := range tt.network.peers {
			sm := x.(*raft)
			if sm.raftLog.committed != tt.wcommitted {
				t.Errorf("#%d.%d: committed = %d, want %d", i, j, sm.raftLog.committed, tt.wcommitted)
			}
			ents := []pb.Entry{}
			for _, e := range nextEnts(sm, tt.network.storage[j]) {
				if e.Data != nil {
					ents = append(ents, e)
				}
			}
			props := []pb.Message{}
			for _, m := range tt.msgs {
				if m.Type == pb.MsgProp {
					props = append(props, m)
				}
			}
			for k, m := range props {
				if !bytes.Equal(ents[k].Data, m.Entries[0].Data) {
					t.Errorf("#%d.%d: data = %d, want %d", i, j, ents[k].Data, m.Entries[0].Data)
				}
			}
		}
	}
}
func TestLearnerLogReplication(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	n1 := newTestLearnerRaft(1, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	n2 := newTestLearnerRaft(2, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	nt := newNetwork(n1, n2)
	n1.becomeFollower(1, None)
	n2.becomeFollower(1, None)
	setRandomizedElectionTimeout(n1, n1.electionTimeout)
	for i := 0; i < n1.electionTimeout; i++ {
		n1.tick()
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgBeat})
	if n1.state != StateLeader {
		t.Errorf("peer 1 state: %s, want %s", n1.state, StateLeader)
	}
	if !n2.isLearner {
		t.Error("peer 2 state: not learner, want yes")
	}
	nextCommitted := n1.raftLog.committed + 1
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
	if n1.raftLog.committed != nextCommitted {
		t.Errorf("peer 1 wants committed to %d, but still %d", nextCommitted, n1.raftLog.committed)
	}
	if n1.raftLog.committed != n2.raftLog.committed {
		t.Errorf("peer 2 wants committed to %d, but still %d", n1.raftLog.committed, n2.raftLog.committed)
	}
	match := n1.getProgress(2).Match
	if match != n2.raftLog.committed {
		t.Errorf("progresss 2 of leader 1 wants match %d, but got %d", n2.raftLog.committed, match)
	}
}
func TestSingleNodeCommit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := newNetwork(nil)
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
	sm := tt.peers[1].(*raft)
	if sm.raftLog.committed != 3 {
		t.Errorf("committed = %d, want %d", sm.raftLog.committed, 3)
	}
}
func TestCannotCommitWithoutNewTermEntry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := newNetwork(nil, nil, nil, nil, nil)
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	tt.cut(1, 3)
	tt.cut(1, 4)
	tt.cut(1, 5)
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
	sm := tt.peers[1].(*raft)
	if sm.raftLog.committed != 1 {
		t.Errorf("committed = %d, want %d", sm.raftLog.committed, 1)
	}
	tt.recover()
	tt.ignore(pb.MsgApp)
	tt.send(pb.Message{From: 2, To: 2, Type: pb.MsgHup})
	sm = tt.peers[2].(*raft)
	if sm.raftLog.committed != 1 {
		t.Errorf("committed = %d, want %d", sm.raftLog.committed, 1)
	}
	tt.recover()
	tt.send(pb.Message{From: 2, To: 2, Type: pb.MsgBeat})
	tt.send(pb.Message{From: 2, To: 2, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
	if sm.raftLog.committed != 5 {
		t.Errorf("committed = %d, want %d", sm.raftLog.committed, 5)
	}
}
func TestCommitWithoutNewTermEntry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := newNetwork(nil, nil, nil, nil, nil)
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	tt.cut(1, 3)
	tt.cut(1, 4)
	tt.cut(1, 5)
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
	sm := tt.peers[1].(*raft)
	if sm.raftLog.committed != 1 {
		t.Errorf("committed = %d, want %d", sm.raftLog.committed, 1)
	}
	tt.recover()
	tt.send(pb.Message{From: 2, To: 2, Type: pb.MsgHup})
	if sm.raftLog.committed != 4 {
		t.Errorf("committed = %d, want %d", sm.raftLog.committed, 4)
	}
}
func TestDuelingCandidates(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	b := newTestRaft(2, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	c := newTestRaft(3, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	nt := newNetwork(a, b, c)
	nt.cut(1, 3)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	sm := nt.peers[1].(*raft)
	if sm.state != StateLeader {
		t.Errorf("state = %s, want %s", sm.state, StateLeader)
	}
	sm = nt.peers[3].(*raft)
	if sm.state != StateCandidate {
		t.Errorf("state = %s, want %s", sm.state, StateCandidate)
	}
	nt.recover()
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	wlog := &raftLog{storage: &MemoryStorage{ents: []pb.Entry{{}, {Data: nil, Term: 1, Index: 1}}}, committed: 1, unstable: unstable{offset: 2}}
	tests := []struct {
		sm	*raft
		state	StateType
		term	uint64
		raftLog	*raftLog
	}{{a, StateFollower, 2, wlog}, {b, StateFollower, 2, wlog}, {c, StateFollower, 2, newLog(NewMemoryStorage(), raftLogger)}}
	for i, tt := range tests {
		if g := tt.sm.state; g != tt.state {
			t.Errorf("#%d: state = %s, want %s", i, g, tt.state)
		}
		if g := tt.sm.Term; g != tt.term {
			t.Errorf("#%d: term = %d, want %d", i, g, tt.term)
		}
		base := ltoa(tt.raftLog)
		if sm, ok := nt.peers[1+uint64(i)].(*raft); ok {
			l := ltoa(sm.raftLog)
			if g := diffu(base, l); g != "" {
				t.Errorf("#%d: diff:\n%s", i, g)
			}
		} else {
			t.Logf("#%d: empty log", i)
		}
	}
}
func TestDuelingPreCandidates(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfgA := newTestConfig(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	cfgB := newTestConfig(2, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	cfgC := newTestConfig(3, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	cfgA.PreVote = true
	cfgB.PreVote = true
	cfgC.PreVote = true
	a := newRaft(cfgA)
	b := newRaft(cfgB)
	c := newRaft(cfgC)
	nt := newNetwork(a, b, c)
	nt.cut(1, 3)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	sm := nt.peers[1].(*raft)
	if sm.state != StateLeader {
		t.Errorf("state = %s, want %s", sm.state, StateLeader)
	}
	sm = nt.peers[3].(*raft)
	if sm.state != StateFollower {
		t.Errorf("state = %s, want %s", sm.state, StateFollower)
	}
	nt.recover()
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	wlog := &raftLog{storage: &MemoryStorage{ents: []pb.Entry{{}, {Data: nil, Term: 1, Index: 1}}}, committed: 1, unstable: unstable{offset: 2}}
	tests := []struct {
		sm	*raft
		state	StateType
		term	uint64
		raftLog	*raftLog
	}{{a, StateLeader, 1, wlog}, {b, StateFollower, 1, wlog}, {c, StateFollower, 1, newLog(NewMemoryStorage(), raftLogger)}}
	for i, tt := range tests {
		if g := tt.sm.state; g != tt.state {
			t.Errorf("#%d: state = %s, want %s", i, g, tt.state)
		}
		if g := tt.sm.Term; g != tt.term {
			t.Errorf("#%d: term = %d, want %d", i, g, tt.term)
		}
		base := ltoa(tt.raftLog)
		if sm, ok := nt.peers[1+uint64(i)].(*raft); ok {
			l := ltoa(sm.raftLog)
			if g := diffu(base, l); g != "" {
				t.Errorf("#%d: diff:\n%s", i, g)
			}
		} else {
			t.Logf("#%d: empty log", i)
		}
	}
}
func TestCandidateConcede(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := newNetwork(nil, nil, nil)
	tt.isolate(1)
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	tt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	tt.recover()
	tt.send(pb.Message{From: 3, To: 3, Type: pb.MsgBeat})
	data := []byte("force follower")
	tt.send(pb.Message{From: 3, To: 3, Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
	tt.send(pb.Message{From: 3, To: 3, Type: pb.MsgBeat})
	a := tt.peers[1].(*raft)
	if g := a.state; g != StateFollower {
		t.Errorf("state = %s, want %s", g, StateFollower)
	}
	if g := a.Term; g != 1 {
		t.Errorf("term = %d, want %d", g, 1)
	}
	wantLog := ltoa(&raftLog{storage: &MemoryStorage{ents: []pb.Entry{{}, {Data: nil, Term: 1, Index: 1}, {Term: 1, Index: 2, Data: data}}}, unstable: unstable{offset: 3}, committed: 2})
	for i, p := range tt.peers {
		if sm, ok := p.(*raft); ok {
			l := ltoa(sm.raftLog)
			if g := diffu(wantLog, l); g != "" {
				t.Errorf("#%d: diff:\n%s", i, g)
			}
		} else {
			t.Logf("#%d: empty log", i)
		}
	}
}
func TestSingleNodeCandidate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := newNetwork(nil)
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	sm := tt.peers[1].(*raft)
	if sm.state != StateLeader {
		t.Errorf("state = %d, want %d", sm.state, StateLeader)
	}
}
func TestSingleNodePreCandidate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := newNetworkWithConfig(preVoteConfig, nil)
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	sm := tt.peers[1].(*raft)
	if sm.state != StateLeader {
		t.Errorf("state = %d, want %d", sm.state, StateLeader)
	}
}
func TestOldMessages(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := newNetwork(nil, nil, nil)
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	tt.send(pb.Message{From: 2, To: 2, Type: pb.MsgHup})
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	tt.send(pb.Message{From: 2, To: 1, Type: pb.MsgApp, Term: 2, Entries: []pb.Entry{{Index: 3, Term: 2}}})
	tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
	ilog := &raftLog{storage: &MemoryStorage{ents: []pb.Entry{{}, {Data: nil, Term: 1, Index: 1}, {Data: nil, Term: 2, Index: 2}, {Data: nil, Term: 3, Index: 3}, {Data: []byte("somedata"), Term: 3, Index: 4}}}, unstable: unstable{offset: 5}, committed: 4}
	base := ltoa(ilog)
	for i, p := range tt.peers {
		if sm, ok := p.(*raft); ok {
			l := ltoa(sm.raftLog)
			if g := diffu(base, l); g != "" {
				t.Errorf("#%d: diff:\n%s", i, g)
			}
		} else {
			t.Logf("#%d: empty log", i)
		}
	}
}
func TestProposal(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		*network
		success	bool
	}{{newNetwork(nil, nil, nil), true}, {newNetwork(nil, nil, nopStepper), true}, {newNetwork(nil, nopStepper, nopStepper), false}, {newNetwork(nil, nopStepper, nopStepper, nil), false}, {newNetwork(nil, nopStepper, nopStepper, nil, nil), true}}
	for j, tt := range tests {
		send := func(m pb.Message) {
			defer func() {
				if !tt.success {
					e := recover()
					if e != nil {
						t.Logf("#%d: err: %s", j, e)
					}
				}
			}()
			tt.send(m)
		}
		data := []byte("somedata")
		send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
		send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: data}}})
		wantLog := newLog(NewMemoryStorage(), raftLogger)
		if tt.success {
			wantLog = &raftLog{storage: &MemoryStorage{ents: []pb.Entry{{}, {Data: nil, Term: 1, Index: 1}, {Term: 1, Index: 2, Data: data}}}, unstable: unstable{offset: 3}, committed: 2}
		}
		base := ltoa(wantLog)
		for i, p := range tt.peers {
			if sm, ok := p.(*raft); ok {
				l := ltoa(sm.raftLog)
				if g := diffu(base, l); g != "" {
					t.Errorf("#%d: diff:\n%s", i, g)
				}
			} else {
				t.Logf("#%d: empty log", i)
			}
		}
		sm := tt.network.peers[1].(*raft)
		if g := sm.Term; g != 1 {
			t.Errorf("#%d: term = %d, want %d", j, g, 1)
		}
	}
}
func TestProposalByProxy(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	data := []byte("somedata")
	tests := []*network{newNetwork(nil, nil, nil), newNetwork(nil, nil, nopStepper)}
	for j, tt := range tests {
		tt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
		tt.send(pb.Message{From: 2, To: 2, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
		wantLog := &raftLog{storage: &MemoryStorage{ents: []pb.Entry{{}, {Data: nil, Term: 1, Index: 1}, {Term: 1, Data: data, Index: 2}}}, unstable: unstable{offset: 3}, committed: 2}
		base := ltoa(wantLog)
		for i, p := range tt.peers {
			if sm, ok := p.(*raft); ok {
				l := ltoa(sm.raftLog)
				if g := diffu(base, l); g != "" {
					t.Errorf("#%d: diff:\n%s", i, g)
				}
			} else {
				t.Logf("#%d: empty log", i)
			}
		}
		sm := tt.peers[1].(*raft)
		if g := sm.Term; g != 1 {
			t.Errorf("#%d: term = %d, want %d", j, g, 1)
		}
	}
}
func TestCommit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		matches	[]uint64
		logs	[]pb.Entry
		smTerm	uint64
		w	uint64
	}{{[]uint64{1}, []pb.Entry{{Index: 1, Term: 1}}, 1, 1}, {[]uint64{1}, []pb.Entry{{Index: 1, Term: 1}}, 2, 0}, {[]uint64{2}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 2}}, 2, 2}, {[]uint64{1}, []pb.Entry{{Index: 1, Term: 2}}, 2, 1}, {[]uint64{2, 1, 1}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 2}}, 1, 1}, {[]uint64{2, 1, 1}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 1}}, 2, 0}, {[]uint64{2, 1, 2}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 2}}, 2, 2}, {[]uint64{2, 1, 2}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 1}}, 2, 0}, {[]uint64{2, 1, 1, 1}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 2}}, 1, 1}, {[]uint64{2, 1, 1, 1}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 1}}, 2, 0}, {[]uint64{2, 1, 1, 2}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 2}}, 1, 1}, {[]uint64{2, 1, 1, 2}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 1}}, 2, 0}, {[]uint64{2, 1, 2, 2}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 2}}, 2, 2}, {[]uint64{2, 1, 2, 2}, []pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 1}}, 2, 0}}
	for i, tt := range tests {
		storage := NewMemoryStorage()
		storage.Append(tt.logs)
		storage.hardState = pb.HardState{Term: tt.smTerm}
		sm := newTestRaft(1, []uint64{1}, 5, 1, storage)
		for j := 0; j < len(tt.matches); j++ {
			sm.setProgress(uint64(j)+1, tt.matches[j], tt.matches[j]+1, false)
		}
		sm.maybeCommit()
		if g := sm.raftLog.committed; g != tt.w {
			t.Errorf("#%d: committed = %d, want %d", i, g, tt.w)
		}
	}
}
func TestPastElectionTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		elapse		int
		wprobability	float64
		round		bool
	}{{5, 0, false}, {10, 0.1, true}, {13, 0.4, true}, {15, 0.6, true}, {18, 0.9, true}, {20, 1, false}}
	for i, tt := range tests {
		sm := newTestRaft(1, []uint64{1}, 10, 1, NewMemoryStorage())
		sm.electionElapsed = tt.elapse
		c := 0
		for j := 0; j < 10000; j++ {
			sm.resetRandomizedElectionTimeout()
			if sm.pastElectionTimeout() {
				c++
			}
		}
		got := float64(c) / 10000.0
		if tt.round {
			got = math.Floor(got*10+0.5) / 10.0
		}
		if got != tt.wprobability {
			t.Errorf("#%d: probability = %v, want %v", i, got, tt.wprobability)
		}
	}
}
func TestStepIgnoreOldTermMsg(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	called := false
	fakeStep := func(r *raft, m pb.Message) {
		called = true
	}
	sm := newTestRaft(1, []uint64{1}, 10, 1, NewMemoryStorage())
	sm.step = fakeStep
	sm.Term = 2
	sm.Step(pb.Message{Type: pb.MsgApp, Term: sm.Term - 1})
	if called {
		t.Errorf("stepFunc called = %v , want %v", called, false)
	}
}
func TestHandleMsgApp(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		m	pb.Message
		wIndex	uint64
		wCommit	uint64
		wReject	bool
	}{{pb.Message{Type: pb.MsgApp, Term: 2, LogTerm: 3, Index: 2, Commit: 3}, 2, 0, true}, {pb.Message{Type: pb.MsgApp, Term: 2, LogTerm: 3, Index: 3, Commit: 3}, 2, 0, true}, {pb.Message{Type: pb.MsgApp, Term: 2, LogTerm: 1, Index: 1, Commit: 1}, 2, 1, false}, {pb.Message{Type: pb.MsgApp, Term: 2, LogTerm: 0, Index: 0, Commit: 1, Entries: []pb.Entry{{Index: 1, Term: 2}}}, 1, 1, false}, {pb.Message{Type: pb.MsgApp, Term: 2, LogTerm: 2, Index: 2, Commit: 3, Entries: []pb.Entry{{Index: 3, Term: 2}, {Index: 4, Term: 2}}}, 4, 3, false}, {pb.Message{Type: pb.MsgApp, Term: 2, LogTerm: 2, Index: 2, Commit: 4, Entries: []pb.Entry{{Index: 3, Term: 2}}}, 3, 3, false}, {pb.Message{Type: pb.MsgApp, Term: 2, LogTerm: 1, Index: 1, Commit: 4, Entries: []pb.Entry{{Index: 2, Term: 2}}}, 2, 2, false}, {pb.Message{Type: pb.MsgApp, Term: 1, LogTerm: 1, Index: 1, Commit: 3}, 2, 1, false}, {pb.Message{Type: pb.MsgApp, Term: 1, LogTerm: 1, Index: 1, Commit: 3, Entries: []pb.Entry{{Index: 2, Term: 2}}}, 2, 2, false}, {pb.Message{Type: pb.MsgApp, Term: 2, LogTerm: 2, Index: 2, Commit: 3}, 2, 2, false}, {pb.Message{Type: pb.MsgApp, Term: 2, LogTerm: 2, Index: 2, Commit: 4}, 2, 2, false}}
	for i, tt := range tests {
		storage := NewMemoryStorage()
		storage.Append([]pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 2}})
		sm := newTestRaft(1, []uint64{1}, 10, 1, storage)
		sm.becomeFollower(2, None)
		sm.handleAppendEntries(tt.m)
		if sm.raftLog.lastIndex() != tt.wIndex {
			t.Errorf("#%d: lastIndex = %d, want %d", i, sm.raftLog.lastIndex(), tt.wIndex)
		}
		if sm.raftLog.committed != tt.wCommit {
			t.Errorf("#%d: committed = %d, want %d", i, sm.raftLog.committed, tt.wCommit)
		}
		m := sm.readMessages()
		if len(m) != 1 {
			t.Fatalf("#%d: msg = nil, want 1", i)
		}
		if m[0].Reject != tt.wReject {
			t.Errorf("#%d: reject = %v, want %v", i, m[0].Reject, tt.wReject)
		}
	}
}
func TestHandleHeartbeat(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	commit := uint64(2)
	tests := []struct {
		m	pb.Message
		wCommit	uint64
	}{{pb.Message{From: 2, To: 1, Type: pb.MsgHeartbeat, Term: 2, Commit: commit + 1}, commit + 1}, {pb.Message{From: 2, To: 1, Type: pb.MsgHeartbeat, Term: 2, Commit: commit - 1}, commit}}
	for i, tt := range tests {
		storage := NewMemoryStorage()
		storage.Append([]pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 2}, {Index: 3, Term: 3}})
		sm := newTestRaft(1, []uint64{1, 2}, 5, 1, storage)
		sm.becomeFollower(2, 2)
		sm.raftLog.commitTo(commit)
		sm.handleHeartbeat(tt.m)
		if sm.raftLog.committed != tt.wCommit {
			t.Errorf("#%d: committed = %d, want %d", i, sm.raftLog.committed, tt.wCommit)
		}
		m := sm.readMessages()
		if len(m) != 1 {
			t.Fatalf("#%d: msg = nil, want 1", i)
		}
		if m[0].Type != pb.MsgHeartbeatResp {
			t.Errorf("#%d: type = %v, want MsgHeartbeatResp", i, m[0].Type)
		}
	}
}
func TestHandleHeartbeatResp(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := NewMemoryStorage()
	storage.Append([]pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 2}, {Index: 3, Term: 3}})
	sm := newTestRaft(1, []uint64{1, 2}, 5, 1, storage)
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.raftLog.commitTo(sm.raftLog.lastIndex())
	sm.Step(pb.Message{From: 2, Type: pb.MsgHeartbeatResp})
	msgs := sm.readMessages()
	if len(msgs) != 1 {
		t.Fatalf("len(msgs) = %d, want 1", len(msgs))
	}
	if msgs[0].Type != pb.MsgApp {
		t.Errorf("type = %v, want MsgApp", msgs[0].Type)
	}
	sm.Step(pb.Message{From: 2, Type: pb.MsgHeartbeatResp})
	msgs = sm.readMessages()
	if len(msgs) != 1 {
		t.Fatalf("len(msgs) = %d, want 1", len(msgs))
	}
	if msgs[0].Type != pb.MsgApp {
		t.Errorf("type = %v, want MsgApp", msgs[0].Type)
	}
	sm.Step(pb.Message{From: 2, Type: pb.MsgAppResp, Index: msgs[0].Index + uint64(len(msgs[0].Entries))})
	sm.readMessages()
	sm.Step(pb.Message{From: 2, Type: pb.MsgHeartbeatResp})
	msgs = sm.readMessages()
	if len(msgs) != 0 {
		t.Fatalf("len(msgs) = %d, want 0: %+v", len(msgs), msgs)
	}
}
func TestRaftFreesReadOnlyMem(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sm := newTestRaft(1, []uint64{1, 2}, 5, 1, NewMemoryStorage())
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.raftLog.commitTo(sm.raftLog.lastIndex())
	ctx := []byte("ctx")
	sm.Step(pb.Message{From: 2, Type: pb.MsgReadIndex, Entries: []pb.Entry{{Data: ctx}}})
	msgs := sm.readMessages()
	if len(msgs) != 1 {
		t.Fatalf("len(msgs) = %d, want 1", len(msgs))
	}
	if msgs[0].Type != pb.MsgHeartbeat {
		t.Fatalf("type = %v, want MsgHeartbeat", msgs[0].Type)
	}
	if !bytes.Equal(msgs[0].Context, ctx) {
		t.Fatalf("Context = %v, want %v", msgs[0].Context, ctx)
	}
	if len(sm.readOnly.readIndexQueue) != 1 {
		t.Fatalf("len(readIndexQueue) = %v, want 1", len(sm.readOnly.readIndexQueue))
	}
	if len(sm.readOnly.pendingReadIndex) != 1 {
		t.Fatalf("len(pendingReadIndex) = %v, want 1", len(sm.readOnly.pendingReadIndex))
	}
	if _, ok := sm.readOnly.pendingReadIndex[string(ctx)]; !ok {
		t.Fatalf("can't find context %v in pendingReadIndex ", ctx)
	}
	sm.Step(pb.Message{From: 2, Type: pb.MsgHeartbeatResp, Context: ctx})
	if len(sm.readOnly.readIndexQueue) != 0 {
		t.Fatalf("len(readIndexQueue) = %v, want 0", len(sm.readOnly.readIndexQueue))
	}
	if len(sm.readOnly.pendingReadIndex) != 0 {
		t.Fatalf("len(pendingReadIndex) = %v, want 0", len(sm.readOnly.pendingReadIndex))
	}
	if _, ok := sm.readOnly.pendingReadIndex[string(ctx)]; ok {
		t.Fatalf("found context %v in pendingReadIndex, want none", ctx)
	}
}
func TestMsgAppRespWaitReset(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sm := newTestRaft(1, []uint64{1, 2, 3}, 5, 1, NewMemoryStorage())
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.bcastAppend()
	sm.readMessages()
	sm.Step(pb.Message{From: 2, Type: pb.MsgAppResp, Index: 1})
	if sm.raftLog.committed != 1 {
		t.Fatalf("expected committed to be 1, got %d", sm.raftLog.committed)
	}
	sm.readMessages()
	sm.Step(pb.Message{From: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	msgs := sm.readMessages()
	if len(msgs) != 1 {
		t.Fatalf("expected 1 message, got %d: %+v", len(msgs), msgs)
	}
	if msgs[0].Type != pb.MsgApp || msgs[0].To != 2 {
		t.Errorf("expected MsgApp to node 2, got %v to %d", msgs[0].Type, msgs[0].To)
	}
	if len(msgs[0].Entries) != 1 || msgs[0].Entries[0].Index != 2 {
		t.Errorf("expected to send entry 2, but got %v", msgs[0].Entries)
	}
	sm.Step(pb.Message{From: 3, Type: pb.MsgAppResp, Index: 1})
	msgs = sm.readMessages()
	if len(msgs) != 1 {
		t.Fatalf("expected 1 message, got %d: %+v", len(msgs), msgs)
	}
	if msgs[0].Type != pb.MsgApp || msgs[0].To != 3 {
		t.Errorf("expected MsgApp to node 3, got %v to %d", msgs[0].Type, msgs[0].To)
	}
	if len(msgs[0].Entries) != 1 || msgs[0].Entries[0].Index != 2 {
		t.Errorf("expected to send entry 2, but got %v", msgs[0].Entries)
	}
}
func TestRecvMsgVote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testRecvMsgVote(t, pb.MsgVote)
}
func TestRecvMsgPreVote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testRecvMsgVote(t, pb.MsgPreVote)
}
func testRecvMsgVote(t *testing.T, msgType pb.MessageType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		state		StateType
		index, logTerm	uint64
		voteFor		uint64
		wreject		bool
	}{{StateFollower, 0, 0, None, true}, {StateFollower, 0, 1, None, true}, {StateFollower, 0, 2, None, true}, {StateFollower, 0, 3, None, false}, {StateFollower, 1, 0, None, true}, {StateFollower, 1, 1, None, true}, {StateFollower, 1, 2, None, true}, {StateFollower, 1, 3, None, false}, {StateFollower, 2, 0, None, true}, {StateFollower, 2, 1, None, true}, {StateFollower, 2, 2, None, false}, {StateFollower, 2, 3, None, false}, {StateFollower, 3, 0, None, true}, {StateFollower, 3, 1, None, true}, {StateFollower, 3, 2, None, false}, {StateFollower, 3, 3, None, false}, {StateFollower, 3, 2, 2, false}, {StateFollower, 3, 2, 1, true}, {StateLeader, 3, 3, 1, true}, {StatePreCandidate, 3, 3, 1, true}, {StateCandidate, 3, 3, 1, true}}
	max := func(a, b uint64) uint64 {
		if a > b {
			return a
		}
		return b
	}
	for i, tt := range tests {
		sm := newTestRaft(1, []uint64{1}, 10, 1, NewMemoryStorage())
		sm.state = tt.state
		switch tt.state {
		case StateFollower:
			sm.step = stepFollower
		case StateCandidate, StatePreCandidate:
			sm.step = stepCandidate
		case StateLeader:
			sm.step = stepLeader
		}
		sm.Vote = tt.voteFor
		sm.raftLog = &raftLog{storage: &MemoryStorage{ents: []pb.Entry{{}, {Index: 1, Term: 2}, {Index: 2, Term: 2}}}, unstable: unstable{offset: 3}}
		term := max(sm.raftLog.lastTerm(), tt.logTerm)
		sm.Term = term
		sm.Step(pb.Message{Type: msgType, Term: term, From: 2, Index: tt.index, LogTerm: tt.logTerm})
		msgs := sm.readMessages()
		if g := len(msgs); g != 1 {
			t.Fatalf("#%d: len(msgs) = %d, want 1", i, g)
			continue
		}
		if g := msgs[0].Type; g != voteRespMsgType(msgType) {
			t.Errorf("#%d, m.Type = %v, want %v", i, g, voteRespMsgType(msgType))
		}
		if g := msgs[0].Reject; g != tt.wreject {
			t.Errorf("#%d, m.Reject = %v, want %v", i, g, tt.wreject)
		}
	}
}
func TestStateTransition(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		from	StateType
		to	StateType
		wallow	bool
		wterm	uint64
		wlead	uint64
	}{{StateFollower, StateFollower, true, 1, None}, {StateFollower, StatePreCandidate, true, 0, None}, {StateFollower, StateCandidate, true, 1, None}, {StateFollower, StateLeader, false, 0, None}, {StatePreCandidate, StateFollower, true, 0, None}, {StatePreCandidate, StatePreCandidate, true, 0, None}, {StatePreCandidate, StateCandidate, true, 1, None}, {StatePreCandidate, StateLeader, true, 0, 1}, {StateCandidate, StateFollower, true, 0, None}, {StateCandidate, StatePreCandidate, true, 0, None}, {StateCandidate, StateCandidate, true, 1, None}, {StateCandidate, StateLeader, true, 0, 1}, {StateLeader, StateFollower, true, 1, None}, {StateLeader, StatePreCandidate, false, 0, None}, {StateLeader, StateCandidate, false, 1, None}, {StateLeader, StateLeader, true, 0, 1}}
	for i, tt := range tests {
		func() {
			defer func() {
				if r := recover(); r != nil {
					if tt.wallow {
						t.Errorf("%d: allow = %v, want %v", i, false, true)
					}
				}
			}()
			sm := newTestRaft(1, []uint64{1}, 10, 1, NewMemoryStorage())
			sm.state = tt.from
			switch tt.to {
			case StateFollower:
				sm.becomeFollower(tt.wterm, tt.wlead)
			case StatePreCandidate:
				sm.becomePreCandidate()
			case StateCandidate:
				sm.becomeCandidate()
			case StateLeader:
				sm.becomeLeader()
			}
			if sm.Term != tt.wterm {
				t.Errorf("%d: term = %d, want %d", i, sm.Term, tt.wterm)
			}
			if sm.lead != tt.wlead {
				t.Errorf("%d: lead = %d, want %d", i, sm.lead, tt.wlead)
			}
		}()
	}
}
func TestAllServerStepdown(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		state	StateType
		wstate	StateType
		wterm	uint64
		windex	uint64
	}{{StateFollower, StateFollower, 3, 0}, {StatePreCandidate, StateFollower, 3, 0}, {StateCandidate, StateFollower, 3, 0}, {StateLeader, StateFollower, 3, 1}}
	tmsgTypes := [...]pb.MessageType{pb.MsgVote, pb.MsgApp}
	tterm := uint64(3)
	for i, tt := range tests {
		sm := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
		switch tt.state {
		case StateFollower:
			sm.becomeFollower(1, None)
		case StatePreCandidate:
			sm.becomePreCandidate()
		case StateCandidate:
			sm.becomeCandidate()
		case StateLeader:
			sm.becomeCandidate()
			sm.becomeLeader()
		}
		for j, msgType := range tmsgTypes {
			sm.Step(pb.Message{From: 2, Type: msgType, Term: tterm, LogTerm: tterm})
			if sm.state != tt.wstate {
				t.Errorf("#%d.%d state = %v , want %v", i, j, sm.state, tt.wstate)
			}
			if sm.Term != tt.wterm {
				t.Errorf("#%d.%d term = %v , want %v", i, j, sm.Term, tt.wterm)
			}
			if uint64(sm.raftLog.lastIndex()) != tt.windex {
				t.Errorf("#%d.%d index = %v , want %v", i, j, sm.raftLog.lastIndex(), tt.windex)
			}
			if uint64(len(sm.raftLog.allEntries())) != tt.windex {
				t.Errorf("#%d.%d len(ents) = %v , want %v", i, j, len(sm.raftLog.allEntries()), tt.windex)
			}
			wlead := uint64(2)
			if msgType == pb.MsgVote {
				wlead = None
			}
			if sm.lead != wlead {
				t.Errorf("#%d, sm.lead = %d, want %d", i, sm.lead, None)
			}
		}
	}
}
func TestLeaderStepdownWhenQuorumActive(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sm := newTestRaft(1, []uint64{1, 2, 3}, 5, 1, NewMemoryStorage())
	sm.checkQuorum = true
	sm.becomeCandidate()
	sm.becomeLeader()
	for i := 0; i < sm.electionTimeout+1; i++ {
		sm.Step(pb.Message{From: 2, Type: pb.MsgHeartbeatResp, Term: sm.Term})
		sm.tick()
	}
	if sm.state != StateLeader {
		t.Errorf("state = %v, want %v", sm.state, StateLeader)
	}
}
func TestLeaderStepdownWhenQuorumLost(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sm := newTestRaft(1, []uint64{1, 2, 3}, 5, 1, NewMemoryStorage())
	sm.checkQuorum = true
	sm.becomeCandidate()
	sm.becomeLeader()
	for i := 0; i < sm.electionTimeout+1; i++ {
		sm.tick()
	}
	if sm.state != StateFollower {
		t.Errorf("state = %v, want %v", sm.state, StateFollower)
	}
}
func TestLeaderSupersedingWithCheckQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	b := newTestRaft(2, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	c := newTestRaft(3, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	a.checkQuorum = true
	b.checkQuorum = true
	c.checkQuorum = true
	nt := newNetwork(a, b, c)
	setRandomizedElectionTimeout(b, b.electionTimeout+1)
	for i := 0; i < b.electionTimeout; i++ {
		b.tick()
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	if a.state != StateLeader {
		t.Errorf("state = %s, want %s", a.state, StateLeader)
	}
	if c.state != StateFollower {
		t.Errorf("state = %s, want %s", c.state, StateFollower)
	}
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	if c.state != StateCandidate {
		t.Errorf("state = %s, want %s", c.state, StateCandidate)
	}
	for i := 0; i < b.electionTimeout; i++ {
		b.tick()
	}
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	if c.state != StateLeader {
		t.Errorf("state = %s, want %s", c.state, StateLeader)
	}
}
func TestLeaderElectionWithCheckQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	b := newTestRaft(2, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	c := newTestRaft(3, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	a.checkQuorum = true
	b.checkQuorum = true
	c.checkQuorum = true
	nt := newNetwork(a, b, c)
	setRandomizedElectionTimeout(a, a.electionTimeout+1)
	setRandomizedElectionTimeout(b, b.electionTimeout+2)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	if a.state != StateLeader {
		t.Errorf("state = %s, want %s", a.state, StateLeader)
	}
	if c.state != StateFollower {
		t.Errorf("state = %s, want %s", c.state, StateFollower)
	}
	setRandomizedElectionTimeout(a, a.electionTimeout+1)
	setRandomizedElectionTimeout(b, b.electionTimeout+2)
	for i := 0; i < a.electionTimeout; i++ {
		a.tick()
	}
	for i := 0; i < b.electionTimeout; i++ {
		b.tick()
	}
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	if a.state != StateFollower {
		t.Errorf("state = %s, want %s", a.state, StateFollower)
	}
	if c.state != StateLeader {
		t.Errorf("state = %s, want %s", c.state, StateLeader)
	}
}
func TestFreeStuckCandidateWithCheckQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	b := newTestRaft(2, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	c := newTestRaft(3, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	a.checkQuorum = true
	b.checkQuorum = true
	c.checkQuorum = true
	nt := newNetwork(a, b, c)
	setRandomizedElectionTimeout(b, b.electionTimeout+1)
	for i := 0; i < b.electionTimeout; i++ {
		b.tick()
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(1)
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	if b.state != StateFollower {
		t.Errorf("state = %s, want %s", b.state, StateFollower)
	}
	if c.state != StateCandidate {
		t.Errorf("state = %s, want %s", c.state, StateCandidate)
	}
	if c.Term != b.Term+1 {
		t.Errorf("term = %d, want %d", c.Term, b.Term+1)
	}
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	if b.state != StateFollower {
		t.Errorf("state = %s, want %s", b.state, StateFollower)
	}
	if c.state != StateCandidate {
		t.Errorf("state = %s, want %s", c.state, StateCandidate)
	}
	if c.Term != b.Term+2 {
		t.Errorf("term = %d, want %d", c.Term, b.Term+2)
	}
	nt.recover()
	nt.send(pb.Message{From: 1, To: 3, Type: pb.MsgHeartbeat, Term: a.Term})
	if a.state != StateFollower {
		t.Errorf("state = %s, want %s", a.state, StateFollower)
	}
	if c.Term != a.Term {
		t.Errorf("term = %d, want %d", c.Term, a.Term)
	}
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	if c.state != StateLeader {
		t.Errorf("peer 3 state: %s, want %s", c.state, StateLeader)
	}
}
func TestNonPromotableVoterWithCheckQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
	b := newTestRaft(2, []uint64{1}, 10, 1, NewMemoryStorage())
	a.checkQuorum = true
	b.checkQuorum = true
	nt := newNetwork(a, b)
	setRandomizedElectionTimeout(b, b.electionTimeout+1)
	b.delProgress(2)
	if b.promotable() {
		t.Fatalf("promotable = %v, want false", b.promotable())
	}
	for i := 0; i < b.electionTimeout; i++ {
		b.tick()
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	if a.state != StateLeader {
		t.Errorf("state = %s, want %s", a.state, StateLeader)
	}
	if b.state != StateFollower {
		t.Errorf("state = %s, want %s", b.state, StateFollower)
	}
	if b.lead != 1 {
		t.Errorf("lead = %d, want 1", b.lead)
	}
}
func TestReadOnlyOptionSafe(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	b := newTestRaft(2, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	c := newTestRaft(3, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	nt := newNetwork(a, b, c)
	setRandomizedElectionTimeout(b, b.electionTimeout+1)
	for i := 0; i < b.electionTimeout; i++ {
		b.tick()
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	if a.state != StateLeader {
		t.Fatalf("state = %s, want %s", a.state, StateLeader)
	}
	tests := []struct {
		sm		*raft
		proposals	int
		wri		uint64
		wctx		[]byte
	}{{a, 10, 11, []byte("ctx1")}, {b, 10, 21, []byte("ctx2")}, {c, 10, 31, []byte("ctx3")}, {a, 10, 41, []byte("ctx4")}, {b, 10, 51, []byte("ctx5")}, {c, 10, 61, []byte("ctx6")}}
	for i, tt := range tests {
		for j := 0; j < tt.proposals; j++ {
			nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
		}
		nt.send(pb.Message{From: tt.sm.id, To: tt.sm.id, Type: pb.MsgReadIndex, Entries: []pb.Entry{{Data: tt.wctx}}})
		r := tt.sm
		if len(r.readStates) == 0 {
			t.Errorf("#%d: len(readStates) = 0, want non-zero", i)
		}
		rs := r.readStates[0]
		if rs.Index != tt.wri {
			t.Errorf("#%d: readIndex = %d, want %d", i, rs.Index, tt.wri)
		}
		if !bytes.Equal(rs.RequestCtx, tt.wctx) {
			t.Errorf("#%d: requestCtx = %v, want %v", i, rs.RequestCtx, tt.wctx)
		}
		r.readStates = nil
	}
}
func TestReadOnlyOptionLease(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	b := newTestRaft(2, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	c := newTestRaft(3, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	a.readOnly.option = ReadOnlyLeaseBased
	b.readOnly.option = ReadOnlyLeaseBased
	c.readOnly.option = ReadOnlyLeaseBased
	a.checkQuorum = true
	b.checkQuorum = true
	c.checkQuorum = true
	nt := newNetwork(a, b, c)
	setRandomizedElectionTimeout(b, b.electionTimeout+1)
	for i := 0; i < b.electionTimeout; i++ {
		b.tick()
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	if a.state != StateLeader {
		t.Fatalf("state = %s, want %s", a.state, StateLeader)
	}
	tests := []struct {
		sm		*raft
		proposals	int
		wri		uint64
		wctx		[]byte
	}{{a, 10, 11, []byte("ctx1")}, {b, 10, 21, []byte("ctx2")}, {c, 10, 31, []byte("ctx3")}, {a, 10, 41, []byte("ctx4")}, {b, 10, 51, []byte("ctx5")}, {c, 10, 61, []byte("ctx6")}}
	for i, tt := range tests {
		for j := 0; j < tt.proposals; j++ {
			nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
		}
		nt.send(pb.Message{From: tt.sm.id, To: tt.sm.id, Type: pb.MsgReadIndex, Entries: []pb.Entry{{Data: tt.wctx}}})
		r := tt.sm
		rs := r.readStates[0]
		if rs.Index != tt.wri {
			t.Errorf("#%d: readIndex = %d, want %d", i, rs.Index, tt.wri)
		}
		if !bytes.Equal(rs.RequestCtx, tt.wctx) {
			t.Errorf("#%d: requestCtx = %v, want %v", i, rs.RequestCtx, tt.wctx)
		}
		r.readStates = nil
	}
}
func TestReadOnlyForNewLeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nodeConfigs := []struct {
		id		uint64
		committed	uint64
		applied		uint64
		compact_index	uint64
	}{{1, 1, 1, 0}, {2, 2, 2, 2}, {3, 2, 2, 2}}
	peers := make([]stateMachine, 0)
	for _, c := range nodeConfigs {
		storage := NewMemoryStorage()
		storage.Append([]pb.Entry{{Index: 1, Term: 1}, {Index: 2, Term: 1}})
		storage.SetHardState(pb.HardState{Term: 1, Commit: c.committed})
		if c.compact_index != 0 {
			storage.Compact(c.compact_index)
		}
		cfg := newTestConfig(c.id, []uint64{1, 2, 3}, 10, 1, storage)
		cfg.Applied = c.applied
		raft := newRaft(cfg)
		peers = append(peers, raft)
	}
	nt := newNetwork(peers...)
	nt.ignore(pb.MsgApp)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	sm := nt.peers[1].(*raft)
	if sm.state != StateLeader {
		t.Fatalf("state = %s, want %s", sm.state, StateLeader)
	}
	var windex uint64 = 4
	wctx := []byte("ctx")
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgReadIndex, Entries: []pb.Entry{{Data: wctx}}})
	if len(sm.readStates) != 0 {
		t.Fatalf("len(readStates) = %d, want zero", len(sm.readStates))
	}
	nt.recover()
	for i := 0; i < sm.heartbeatTimeout; i++ {
		sm.tick()
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	if sm.raftLog.committed != 4 {
		t.Fatalf("committed = %d, want 4", sm.raftLog.committed)
	}
	lastLogTerm := sm.raftLog.zeroTermOnErrCompacted(sm.raftLog.term(sm.raftLog.committed))
	if lastLogTerm != sm.Term {
		t.Fatalf("last log term = %d, want %d", lastLogTerm, sm.Term)
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgReadIndex, Entries: []pb.Entry{{Data: wctx}}})
	if len(sm.readStates) != 1 {
		t.Fatalf("len(readStates) = %d, want 1", len(sm.readStates))
	}
	rs := sm.readStates[0]
	if rs.Index != windex {
		t.Fatalf("readIndex = %d, want %d", rs.Index, windex)
	}
	if !bytes.Equal(rs.RequestCtx, wctx) {
		t.Fatalf("requestCtx = %v, want %v", rs.RequestCtx, wctx)
	}
}
func TestLeaderAppResp(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		index		uint64
		reject		bool
		wmatch		uint64
		wnext		uint64
		wmsgNum		int
		windex		uint64
		wcommitted	uint64
	}{{3, true, 0, 3, 0, 0, 0}, {2, true, 0, 2, 1, 1, 0}, {2, false, 2, 4, 2, 2, 2}, {0, false, 0, 3, 0, 0, 0}}
	for i, tt := range tests {
		sm := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
		sm.raftLog = &raftLog{storage: &MemoryStorage{ents: []pb.Entry{{}, {Index: 1, Term: 0}, {Index: 2, Term: 1}}}, unstable: unstable{offset: 3}}
		sm.becomeCandidate()
		sm.becomeLeader()
		sm.readMessages()
		sm.Step(pb.Message{From: 2, Type: pb.MsgAppResp, Index: tt.index, Term: sm.Term, Reject: tt.reject, RejectHint: tt.index})
		p := sm.prs[2]
		if p.Match != tt.wmatch {
			t.Errorf("#%d match = %d, want %d", i, p.Match, tt.wmatch)
		}
		if p.Next != tt.wnext {
			t.Errorf("#%d next = %d, want %d", i, p.Next, tt.wnext)
		}
		msgs := sm.readMessages()
		if len(msgs) != tt.wmsgNum {
			t.Errorf("#%d msgNum = %d, want %d", i, len(msgs), tt.wmsgNum)
		}
		for j, msg := range msgs {
			if msg.Index != tt.windex {
				t.Errorf("#%d.%d index = %d, want %d", i, j, msg.Index, tt.windex)
			}
			if msg.Commit != tt.wcommitted {
				t.Errorf("#%d.%d commit = %d, want %d", i, j, msg.Commit, tt.wcommitted)
			}
		}
	}
}
func TestBcastBeat(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	offset := uint64(1000)
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: offset, Term: 1, ConfState: pb.ConfState{Nodes: []uint64{1, 2, 3}}}}
	storage := NewMemoryStorage()
	storage.ApplySnapshot(s)
	sm := newTestRaft(1, nil, 10, 1, storage)
	sm.Term = 1
	sm.becomeCandidate()
	sm.becomeLeader()
	for i := 0; i < 10; i++ {
		sm.appendEntry(pb.Entry{Index: uint64(i) + 1})
	}
	sm.prs[2].Match, sm.prs[2].Next = 5, 6
	sm.prs[3].Match, sm.prs[3].Next = sm.raftLog.lastIndex(), sm.raftLog.lastIndex()+1
	sm.Step(pb.Message{Type: pb.MsgBeat})
	msgs := sm.readMessages()
	if len(msgs) != 2 {
		t.Fatalf("len(msgs) = %v, want 2", len(msgs))
	}
	wantCommitMap := map[uint64]uint64{2: min(sm.raftLog.committed, sm.prs[2].Match), 3: min(sm.raftLog.committed, sm.prs[3].Match)}
	for i, m := range msgs {
		if m.Type != pb.MsgHeartbeat {
			t.Fatalf("#%d: type = %v, want = %v", i, m.Type, pb.MsgHeartbeat)
		}
		if m.Index != 0 {
			t.Fatalf("#%d: prevIndex = %d, want %d", i, m.Index, 0)
		}
		if m.LogTerm != 0 {
			t.Fatalf("#%d: prevTerm = %d, want %d", i, m.LogTerm, 0)
		}
		if wantCommitMap[m.To] == 0 {
			t.Fatalf("#%d: unexpected to %d", i, m.To)
		} else {
			if m.Commit != wantCommitMap[m.To] {
				t.Fatalf("#%d: commit = %d, want %d", i, m.Commit, wantCommitMap[m.To])
			}
			delete(wantCommitMap, m.To)
		}
		if len(m.Entries) != 0 {
			t.Fatalf("#%d: len(entries) = %d, want 0", i, len(m.Entries))
		}
	}
}
func TestRecvMsgBeat(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		state	StateType
		wMsg	int
	}{{StateLeader, 2}, {StateCandidate, 0}, {StateFollower, 0}}
	for i, tt := range tests {
		sm := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
		sm.raftLog = &raftLog{storage: &MemoryStorage{ents: []pb.Entry{{}, {Index: 1, Term: 0}, {Index: 2, Term: 1}}}}
		sm.Term = 1
		sm.state = tt.state
		switch tt.state {
		case StateFollower:
			sm.step = stepFollower
		case StateCandidate:
			sm.step = stepCandidate
		case StateLeader:
			sm.step = stepLeader
		}
		sm.Step(pb.Message{From: 1, To: 1, Type: pb.MsgBeat})
		msgs := sm.readMessages()
		if len(msgs) != tt.wMsg {
			t.Errorf("%d: len(msgs) = %d, want %d", i, len(msgs), tt.wMsg)
		}
		for _, m := range msgs {
			if m.Type != pb.MsgHeartbeat {
				t.Errorf("%d: msg.type = %v, want %v", i, m.Type, pb.MsgHeartbeat)
			}
		}
	}
}
func TestLeaderIncreaseNext(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	previousEnts := []pb.Entry{{Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}}
	tests := []struct {
		state	ProgressStateType
		next	uint64
		wnext	uint64
	}{{ProgressStateReplicate, 2, uint64(len(previousEnts) + 1 + 1 + 1)}, {ProgressStateProbe, 2, 2}}
	for i, tt := range tests {
		sm := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
		sm.raftLog.append(previousEnts...)
		sm.becomeCandidate()
		sm.becomeLeader()
		sm.prs[2].State = tt.state
		sm.prs[2].Next = tt.next
		sm.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
		p := sm.prs[2]
		if p.Next != tt.wnext {
			t.Errorf("#%d next = %d, want %d", i, p.Next, tt.wnext)
		}
	}
}
func TestSendAppendForProgressProbe(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	r.readMessages()
	r.prs[2].becomeProbe()
	for i := 0; i < 3; i++ {
		if i == 0 {
			r.appendEntry(pb.Entry{Data: []byte("somedata")})
			r.sendAppend(2)
			msg := r.readMessages()
			if len(msg) != 1 {
				t.Errorf("len(msg) = %d, want %d", len(msg), 1)
			}
			if msg[0].Index != 0 {
				t.Errorf("index = %d, want %d", msg[0].Index, 0)
			}
		}
		if !r.prs[2].Paused {
			t.Errorf("paused = %v, want true", r.prs[2].Paused)
		}
		for j := 0; j < 10; j++ {
			r.appendEntry(pb.Entry{Data: []byte("somedata")})
			r.sendAppend(2)
			if l := len(r.readMessages()); l != 0 {
				t.Errorf("len(msg) = %d, want %d", l, 0)
			}
		}
		for j := 0; j < r.heartbeatTimeout; j++ {
			r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgBeat})
		}
		if !r.prs[2].Paused {
			t.Errorf("paused = %v, want true", r.prs[2].Paused)
		}
		msg := r.readMessages()
		if len(msg) != 1 {
			t.Errorf("len(msg) = %d, want %d", len(msg), 1)
		}
		if msg[0].Type != pb.MsgHeartbeat {
			t.Errorf("type = %v, want %v", msg[0].Type, pb.MsgHeartbeat)
		}
	}
	r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgHeartbeatResp})
	msg := r.readMessages()
	if len(msg) != 1 {
		t.Errorf("len(msg) = %d, want %d", len(msg), 1)
	}
	if msg[0].Index != 0 {
		t.Errorf("index = %d, want %d", msg[0].Index, 0)
	}
	if !r.prs[2].Paused {
		t.Errorf("paused = %v, want true", r.prs[2].Paused)
	}
}
func TestSendAppendForProgressReplicate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	r.readMessages()
	r.prs[2].becomeReplicate()
	for i := 0; i < 10; i++ {
		r.appendEntry(pb.Entry{Data: []byte("somedata")})
		r.sendAppend(2)
		msgs := r.readMessages()
		if len(msgs) != 1 {
			t.Errorf("len(msg) = %d, want %d", len(msgs), 1)
		}
	}
}
func TestSendAppendForProgressSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	r.readMessages()
	r.prs[2].becomeSnapshot(10)
	for i := 0; i < 10; i++ {
		r.appendEntry(pb.Entry{Data: []byte("somedata")})
		r.sendAppend(2)
		msgs := r.readMessages()
		if len(msgs) != 0 {
			t.Errorf("len(msg) = %d, want %d", len(msgs), 0)
		}
	}
}
func TestRecvMsgUnreachable(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	previousEnts := []pb.Entry{{Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}}
	s := NewMemoryStorage()
	s.Append(previousEnts)
	r := newTestRaft(1, []uint64{1, 2}, 10, 1, s)
	r.becomeCandidate()
	r.becomeLeader()
	r.readMessages()
	r.prs[2].Match = 3
	r.prs[2].becomeReplicate()
	r.prs[2].optimisticUpdate(5)
	r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgUnreachable})
	if r.prs[2].State != ProgressStateProbe {
		t.Errorf("state = %s, want %s", r.prs[2].State, ProgressStateProbe)
	}
	if wnext := r.prs[2].Match + 1; r.prs[2].Next != wnext {
		t.Errorf("next = %d, want %d", r.prs[2].Next, wnext)
	}
}
func TestRestore(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: 11, Term: 11, ConfState: pb.ConfState{Nodes: []uint64{1, 2, 3}}}}
	storage := NewMemoryStorage()
	sm := newTestRaft(1, []uint64{1, 2}, 10, 1, storage)
	if ok := sm.restore(s); !ok {
		t.Fatal("restore fail, want succeed")
	}
	if sm.raftLog.lastIndex() != s.Metadata.Index {
		t.Errorf("log.lastIndex = %d, want %d", sm.raftLog.lastIndex(), s.Metadata.Index)
	}
	if mustTerm(sm.raftLog.term(s.Metadata.Index)) != s.Metadata.Term {
		t.Errorf("log.lastTerm = %d, want %d", mustTerm(sm.raftLog.term(s.Metadata.Index)), s.Metadata.Term)
	}
	sg := sm.nodes()
	if !reflect.DeepEqual(sg, s.Metadata.ConfState.Nodes) {
		t.Errorf("sm.Nodes = %+v, want %+v", sg, s.Metadata.ConfState.Nodes)
	}
	if ok := sm.restore(s); ok {
		t.Fatal("restore succeed, want fail")
	}
}
func TestRestoreWithLearner(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: 11, Term: 11, ConfState: pb.ConfState{Nodes: []uint64{1, 2}, Learners: []uint64{3}}}}
	storage := NewMemoryStorage()
	sm := newTestLearnerRaft(3, []uint64{1, 2}, []uint64{3}, 10, 1, storage)
	if ok := sm.restore(s); !ok {
		t.Error("restore fail, want succeed")
	}
	if sm.raftLog.lastIndex() != s.Metadata.Index {
		t.Errorf("log.lastIndex = %d, want %d", sm.raftLog.lastIndex(), s.Metadata.Index)
	}
	if mustTerm(sm.raftLog.term(s.Metadata.Index)) != s.Metadata.Term {
		t.Errorf("log.lastTerm = %d, want %d", mustTerm(sm.raftLog.term(s.Metadata.Index)), s.Metadata.Term)
	}
	sg := sm.nodes()
	if len(sg) != len(s.Metadata.ConfState.Nodes)+len(s.Metadata.ConfState.Learners) {
		t.Errorf("sm.Nodes = %+v, length not equal with %+v", sg, s.Metadata.ConfState)
	}
	for _, n := range s.Metadata.ConfState.Nodes {
		if sm.prs[n].IsLearner {
			t.Errorf("sm.Node %x isLearner = %s, want %t", n, sm.prs[n], false)
		}
	}
	for _, n := range s.Metadata.ConfState.Learners {
		if !sm.learnerPrs[n].IsLearner {
			t.Errorf("sm.Node %x isLearner = %s, want %t", n, sm.prs[n], true)
		}
	}
	if ok := sm.restore(s); ok {
		t.Error("restore succeed, want fail")
	}
}
func TestRestoreInvalidLearner(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: 11, Term: 11, ConfState: pb.ConfState{Nodes: []uint64{1, 2}, Learners: []uint64{3}}}}
	storage := NewMemoryStorage()
	sm := newTestRaft(3, []uint64{1, 2, 3}, 10, 1, storage)
	if sm.isLearner {
		t.Errorf("%x is learner, want not", sm.id)
	}
	if ok := sm.restore(s); ok {
		t.Error("restore succeed, want fail")
	}
}
func TestRestoreLearnerPromotion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: 11, Term: 11, ConfState: pb.ConfState{Nodes: []uint64{1, 2, 3}}}}
	storage := NewMemoryStorage()
	sm := newTestLearnerRaft(3, []uint64{1, 2}, []uint64{3}, 10, 1, storage)
	if !sm.isLearner {
		t.Errorf("%x is not learner, want yes", sm.id)
	}
	if ok := sm.restore(s); !ok {
		t.Error("restore fail, want succeed")
	}
	if sm.isLearner {
		t.Errorf("%x is learner, want not", sm.id)
	}
}
func TestLearnerReceiveSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: 11, Term: 11, ConfState: pb.ConfState{Nodes: []uint64{1}, Learners: []uint64{2}}}}
	n1 := newTestLearnerRaft(1, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	n2 := newTestLearnerRaft(2, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	n1.restore(s)
	n1.raftLog.appliedTo(n1.raftLog.committed)
	nt := newNetwork(n1, n2)
	setRandomizedElectionTimeout(n1, n1.electionTimeout)
	for i := 0; i < n1.electionTimeout; i++ {
		n1.tick()
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgBeat})
	if n2.raftLog.committed != n1.raftLog.committed {
		t.Errorf("peer 2 must commit to %d, but %d", n1.raftLog.committed, n2.raftLog.committed)
	}
}
func TestRestoreIgnoreSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	previousEnts := []pb.Entry{{Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}}
	commit := uint64(1)
	storage := NewMemoryStorage()
	sm := newTestRaft(1, []uint64{1, 2}, 10, 1, storage)
	sm.raftLog.append(previousEnts...)
	sm.raftLog.commitTo(commit)
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: commit, Term: 1, ConfState: pb.ConfState{Nodes: []uint64{1, 2}}}}
	if ok := sm.restore(s); ok {
		t.Errorf("restore = %t, want %t", ok, false)
	}
	if sm.raftLog.committed != commit {
		t.Errorf("commit = %d, want %d", sm.raftLog.committed, commit)
	}
	s.Metadata.Index = commit + 1
	if ok := sm.restore(s); ok {
		t.Errorf("restore = %t, want %t", ok, false)
	}
	if sm.raftLog.committed != commit+1 {
		t.Errorf("commit = %d, want %d", sm.raftLog.committed, commit+1)
	}
}
func TestProvideSnap(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: 11, Term: 11, ConfState: pb.ConfState{Nodes: []uint64{1, 2}}}}
	storage := NewMemoryStorage()
	sm := newTestRaft(1, []uint64{1}, 10, 1, storage)
	sm.restore(s)
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.prs[2].Next = sm.raftLog.firstIndex()
	sm.Step(pb.Message{From: 2, To: 1, Type: pb.MsgAppResp, Index: sm.prs[2].Next - 1, Reject: true})
	msgs := sm.readMessages()
	if len(msgs) != 1 {
		t.Fatalf("len(msgs) = %d, want 1", len(msgs))
	}
	m := msgs[0]
	if m.Type != pb.MsgSnap {
		t.Errorf("m.Type = %v, want %v", m.Type, pb.MsgSnap)
	}
}
func TestIgnoreProvidingSnap(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: 11, Term: 11, ConfState: pb.ConfState{Nodes: []uint64{1, 2}}}}
	storage := NewMemoryStorage()
	sm := newTestRaft(1, []uint64{1}, 10, 1, storage)
	sm.restore(s)
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.prs[2].Next = sm.raftLog.firstIndex() - 1
	sm.prs[2].RecentActive = false
	sm.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
	msgs := sm.readMessages()
	if len(msgs) != 0 {
		t.Errorf("len(msgs) = %d, want 0", len(msgs))
	}
}
func TestRestoreFromSnapMsg(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: 11, Term: 11, ConfState: pb.ConfState{Nodes: []uint64{1, 2}}}}
	m := pb.Message{Type: pb.MsgSnap, From: 1, Term: 2, Snapshot: s}
	sm := newTestRaft(2, []uint64{1, 2}, 10, 1, NewMemoryStorage())
	sm.Step(m)
	if sm.lead != uint64(1) {
		t.Errorf("sm.lead = %d, want 1", sm.lead)
	}
}
func TestSlowNodeRestore(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(3)
	for j := 0; j <= 100; j++ {
		nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	}
	lead := nt.peers[1].(*raft)
	nextEnts(lead, nt.storage[1])
	nt.storage[1].CreateSnapshot(lead.raftLog.applied, &pb.ConfState{Nodes: lead.nodes()}, nil)
	nt.storage[1].Compact(lead.raftLog.applied)
	nt.recover()
	for {
		nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgBeat})
		if lead.prs[3].RecentActive {
			break
		}
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	follower := nt.peers[3].(*raft)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	if follower.raftLog.committed != lead.raftLog.committed {
		t.Errorf("follower.committed = %d, want %d", follower.raftLog.committed, lead.raftLog.committed)
	}
}
func TestStepConfig(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	index := r.raftLog.lastIndex()
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Type: pb.EntryConfChange}}})
	if g := r.raftLog.lastIndex(); g != index+1 {
		t.Errorf("index = %d, want %d", g, index+1)
	}
	if !r.pendingConf {
		t.Errorf("pendingConf = %v, want true", r.pendingConf)
	}
}
func TestStepIgnoreConfig(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Type: pb.EntryConfChange}}})
	index := r.raftLog.lastIndex()
	pendingConf := r.pendingConf
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Type: pb.EntryConfChange}}})
	wents := []pb.Entry{{Type: pb.EntryNormal, Term: 1, Index: 3, Data: nil}}
	ents, err := r.raftLog.entries(index+1, noLimit)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if !reflect.DeepEqual(ents, wents) {
		t.Errorf("ents = %+v, want %+v", ents, wents)
	}
	if r.pendingConf != pendingConf {
		t.Errorf("pendingConf = %v, want %v", r.pendingConf, pendingConf)
	}
}
func TestRecoverPendingConfig(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		entType		pb.EntryType
		wpending	bool
	}{{pb.EntryNormal, false}, {pb.EntryConfChange, true}}
	for i, tt := range tests {
		r := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
		r.appendEntry(pb.Entry{Type: tt.entType})
		r.becomeCandidate()
		r.becomeLeader()
		if r.pendingConf != tt.wpending {
			t.Errorf("#%d: pendingConf = %v, want %v", i, r.pendingConf, tt.wpending)
		}
	}
}
func TestRecoverDoublePendingConfig(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Errorf("expect panic, but nothing happens")
			}
		}()
		r := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
		r.appendEntry(pb.Entry{Type: pb.EntryConfChange})
		r.appendEntry(pb.Entry{Type: pb.EntryConfChange})
		r.becomeCandidate()
		r.becomeLeader()
	}()
}
func TestAddNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1}, 10, 1, NewMemoryStorage())
	r.pendingConf = true
	r.addNode(2)
	if r.pendingConf {
		t.Errorf("pendingConf = %v, want false", r.pendingConf)
	}
	nodes := r.nodes()
	wnodes := []uint64{1, 2}
	if !reflect.DeepEqual(nodes, wnodes) {
		t.Errorf("nodes = %v, want %v", nodes, wnodes)
	}
}
func TestAddLearner(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1}, 10, 1, NewMemoryStorage())
	r.pendingConf = true
	r.addLearner(2)
	if r.pendingConf {
		t.Errorf("pendingConf = %v, want false", r.pendingConf)
	}
	nodes := r.nodes()
	wnodes := []uint64{1, 2}
	if !reflect.DeepEqual(nodes, wnodes) {
		t.Errorf("nodes = %v, want %v", nodes, wnodes)
	}
	if !r.learnerPrs[2].IsLearner {
		t.Errorf("node 2 is learner %t, want %t", r.prs[2].IsLearner, true)
	}
}
func TestAddNodeCheckQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1}, 10, 1, NewMemoryStorage())
	r.pendingConf = true
	r.checkQuorum = true
	r.becomeCandidate()
	r.becomeLeader()
	for i := 0; i < r.electionTimeout-1; i++ {
		r.tick()
	}
	r.addNode(2)
	r.tick()
	if r.state != StateLeader {
		t.Errorf("state = %v, want %v", r.state, StateLeader)
	}
	for i := 0; i < r.electionTimeout; i++ {
		r.tick()
	}
	if r.state != StateFollower {
		t.Errorf("state = %v, want %v", r.state, StateFollower)
	}
}
func TestRemoveNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 10, 1, NewMemoryStorage())
	r.pendingConf = true
	r.removeNode(2)
	if r.pendingConf {
		t.Errorf("pendingConf = %v, want false", r.pendingConf)
	}
	w := []uint64{1}
	if g := r.nodes(); !reflect.DeepEqual(g, w) {
		t.Errorf("nodes = %v, want %v", g, w)
	}
	r.removeNode(1)
	w = []uint64{}
	if g := r.nodes(); !reflect.DeepEqual(g, w) {
		t.Errorf("nodes = %v, want %v", g, w)
	}
}
func TestRemoveLearner(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestLearnerRaft(1, []uint64{1}, []uint64{2}, 10, 1, NewMemoryStorage())
	r.pendingConf = true
	r.removeNode(2)
	if r.pendingConf {
		t.Errorf("pendingConf = %v, want false", r.pendingConf)
	}
	w := []uint64{1}
	if g := r.nodes(); !reflect.DeepEqual(g, w) {
		t.Errorf("nodes = %v, want %v", g, w)
	}
	r.removeNode(1)
	w = []uint64{}
	if g := r.nodes(); !reflect.DeepEqual(g, w) {
		t.Errorf("nodes = %v, want %v", g, w)
	}
}
func TestPromotable(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	id := uint64(1)
	tests := []struct {
		peers	[]uint64
		wp	bool
	}{{[]uint64{1}, true}, {[]uint64{1, 2, 3}, true}, {[]uint64{}, false}, {[]uint64{2, 3}, false}}
	for i, tt := range tests {
		r := newTestRaft(id, tt.peers, 5, 1, NewMemoryStorage())
		if g := r.promotable(); g != tt.wp {
			t.Errorf("#%d: promotable = %v, want %v", i, g, tt.wp)
		}
	}
}
func TestRaftNodes(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		ids	[]uint64
		wids	[]uint64
	}{{[]uint64{1, 2, 3}, []uint64{1, 2, 3}}, {[]uint64{3, 2, 1}, []uint64{1, 2, 3}}}
	for i, tt := range tests {
		r := newTestRaft(1, tt.ids, 10, 1, NewMemoryStorage())
		if !reflect.DeepEqual(r.nodes(), tt.wids) {
			t.Errorf("#%d: nodes = %+v, want %+v", i, r.nodes(), tt.wids)
		}
	}
}
func TestCampaignWhileLeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCampaignWhileLeader(t, false)
}
func TestPreCampaignWhileLeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCampaignWhileLeader(t, true)
}
func testCampaignWhileLeader(t *testing.T, preVote bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := newTestConfig(1, []uint64{1}, 5, 1, NewMemoryStorage())
	cfg.PreVote = preVote
	r := newRaft(cfg)
	if r.state != StateFollower {
		t.Errorf("expected new node to be follower but got %s", r.state)
	}
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	if r.state != StateLeader {
		t.Errorf("expected single-node election to become leader but got %s", r.state)
	}
	term := r.Term
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	if r.state != StateLeader {
		t.Errorf("expected to remain leader but got %s", r.state)
	}
	if r.Term != term {
		t.Errorf("expected to remain in term %v but got %v", term, r.Term)
	}
}
func TestCommitAfterRemoveNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := NewMemoryStorage()
	r := newTestRaft(1, []uint64{1, 2}, 5, 1, s)
	r.becomeCandidate()
	r.becomeLeader()
	cc := pb.ConfChange{Type: pb.ConfChangeRemoveNode, NodeID: 2}
	ccData, err := cc.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	r.Step(pb.Message{Type: pb.MsgProp, Entries: []pb.Entry{{Type: pb.EntryConfChange, Data: ccData}}})
	if ents := nextEnts(r, s); len(ents) > 0 {
		t.Fatalf("unexpected committed entries: %v", ents)
	}
	ccIndex := r.raftLog.lastIndex()
	r.Step(pb.Message{Type: pb.MsgProp, Entries: []pb.Entry{{Type: pb.EntryNormal, Data: []byte("hello")}}})
	r.Step(pb.Message{Type: pb.MsgAppResp, From: 2, Index: ccIndex})
	ents := nextEnts(r, s)
	if len(ents) != 2 {
		t.Fatalf("expected two committed entries, got %v", ents)
	}
	if ents[0].Type != pb.EntryNormal || ents[0].Data != nil {
		t.Fatalf("expected ents[0] to be empty, but got %v", ents[0])
	}
	if ents[1].Type != pb.EntryConfChange {
		t.Fatalf("expected ents[1] to be EntryConfChange, got %v", ents[1])
	}
	r.removeNode(2)
	ents = nextEnts(r, s)
	if len(ents) != 1 || ents[0].Type != pb.EntryNormal || string(ents[0].Data) != "hello" {
		t.Fatalf("expected one committed EntryNormal, got %v", ents)
	}
}
func TestLeaderTransferToUpToDateNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	lead := nt.peers[1].(*raft)
	if lead.lead != 1 {
		t.Fatalf("after election leader is %x, want 1", lead.lead)
	}
	nt.send(pb.Message{From: 2, To: 1, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateFollower, 2)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	nt.send(pb.Message{From: 1, To: 2, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateLeader, 1)
}
func TestLeaderTransferToUpToDateNodeFromFollower(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	lead := nt.peers[1].(*raft)
	if lead.lead != 1 {
		t.Fatalf("after election leader is %x, want 1", lead.lead)
	}
	nt.send(pb.Message{From: 2, To: 2, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateFollower, 2)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateLeader, 1)
}
func TestLeaderTransferWithCheckQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	for i := 1; i < 4; i++ {
		r := nt.peers[uint64(i)].(*raft)
		r.checkQuorum = true
		setRandomizedElectionTimeout(r, r.electionTimeout+i)
	}
	f := nt.peers[2].(*raft)
	for i := 0; i < f.electionTimeout; i++ {
		f.tick()
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	lead := nt.peers[1].(*raft)
	if lead.lead != 1 {
		t.Fatalf("after election leader is %x, want 1", lead.lead)
	}
	nt.send(pb.Message{From: 2, To: 1, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateFollower, 2)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	nt.send(pb.Message{From: 1, To: 2, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateLeader, 1)
}
func TestLeaderTransferToSlowFollower(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defaultLogger.EnableDebug()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(3)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	nt.recover()
	lead := nt.peers[1].(*raft)
	if lead.prs[3].Match != 1 {
		t.Fatalf("node 1 has match %x for node 3, want %x", lead.prs[3].Match, 1)
	}
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateFollower, 3)
}
func TestLeaderTransferAfterSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(3)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	lead := nt.peers[1].(*raft)
	nextEnts(lead, nt.storage[1])
	nt.storage[1].CreateSnapshot(lead.raftLog.applied, &pb.ConfState{Nodes: lead.nodes()}, nil)
	nt.storage[1].Compact(lead.raftLog.applied)
	nt.recover()
	if lead.prs[3].Match != 1 {
		t.Fatalf("node 1 has match %x for node 3, want %x", lead.prs[3].Match, 1)
	}
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgHeartbeatResp})
	checkLeaderTransferState(t, lead, StateFollower, 3)
}
func TestLeaderTransferToSelf(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	lead := nt.peers[1].(*raft)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateLeader, 1)
}
func TestLeaderTransferToNonExistingNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	lead := nt.peers[1].(*raft)
	nt.send(pb.Message{From: 4, To: 1, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateLeader, 1)
}
func TestLeaderTransferTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(3)
	lead := nt.peers[1].(*raft)
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	if lead.leadTransferee != 3 {
		t.Fatalf("wait transferring, leadTransferee = %v, want %v", lead.leadTransferee, 3)
	}
	for i := 0; i < lead.heartbeatTimeout; i++ {
		lead.tick()
	}
	if lead.leadTransferee != 3 {
		t.Fatalf("wait transferring, leadTransferee = %v, want %v", lead.leadTransferee, 3)
	}
	for i := 0; i < lead.electionTimeout-lead.heartbeatTimeout; i++ {
		lead.tick()
	}
	checkLeaderTransferState(t, lead, StateLeader, 1)
}
func TestLeaderTransferIgnoreProposal(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(3)
	lead := nt.peers[1].(*raft)
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	if lead.leadTransferee != 3 {
		t.Fatalf("wait transferring, leadTransferee = %v, want %v", lead.leadTransferee, 3)
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
	if lead.prs[1].Match != 1 {
		t.Fatalf("node 1 has match %x, want %x", lead.prs[1].Match, 1)
	}
}
func TestLeaderTransferReceiveHigherTermVote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(3)
	lead := nt.peers[1].(*raft)
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	if lead.leadTransferee != 3 {
		t.Fatalf("wait transferring, leadTransferee = %v, want %v", lead.leadTransferee, 3)
	}
	nt.send(pb.Message{From: 2, To: 2, Type: pb.MsgHup, Index: 1, Term: 2})
	checkLeaderTransferState(t, lead, StateFollower, 2)
}
func TestLeaderTransferRemoveNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.ignore(pb.MsgTimeoutNow)
	lead := nt.peers[1].(*raft)
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	if lead.leadTransferee != 3 {
		t.Fatalf("wait transferring, leadTransferee = %v, want %v", lead.leadTransferee, 3)
	}
	lead.removeNode(3)
	checkLeaderTransferState(t, lead, StateLeader, 1)
}
func TestLeaderTransferBack(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(3)
	lead := nt.peers[1].(*raft)
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	if lead.leadTransferee != 3 {
		t.Fatalf("wait transferring, leadTransferee = %v, want %v", lead.leadTransferee, 3)
	}
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateLeader, 1)
}
func TestLeaderTransferSecondTransferToAnotherNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(3)
	lead := nt.peers[1].(*raft)
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	if lead.leadTransferee != 3 {
		t.Fatalf("wait transferring, leadTransferee = %v, want %v", lead.leadTransferee, 3)
	}
	nt.send(pb.Message{From: 2, To: 1, Type: pb.MsgTransferLeader})
	checkLeaderTransferState(t, lead, StateFollower, 2)
}
func TestLeaderTransferSecondTransferToSameNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt := newNetwork(nil, nil, nil)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(3)
	lead := nt.peers[1].(*raft)
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	if lead.leadTransferee != 3 {
		t.Fatalf("wait transferring, leadTransferee = %v, want %v", lead.leadTransferee, 3)
	}
	for i := 0; i < lead.heartbeatTimeout; i++ {
		lead.tick()
	}
	nt.send(pb.Message{From: 3, To: 1, Type: pb.MsgTransferLeader})
	for i := 0; i < lead.electionTimeout-lead.heartbeatTimeout; i++ {
		lead.tick()
	}
	checkLeaderTransferState(t, lead, StateLeader, 1)
}
func checkLeaderTransferState(t *testing.T, r *raft, state StateType, lead uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.state != state || r.lead != lead {
		t.Fatalf("after transferring, node has state %v lead %v, want state %v lead %v", r.state, r.lead, state, lead)
	}
	if r.leadTransferee != None {
		t.Fatalf("after transferring, node has leadTransferee %v, want leadTransferee %v", r.leadTransferee, None)
	}
}
func TestTransferNonMember(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{2, 3, 4}, 5, 1, NewMemoryStorage())
	r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgTimeoutNow})
	r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgVoteResp})
	r.Step(pb.Message{From: 3, To: 1, Type: pb.MsgVoteResp})
	if r.state != StateFollower {
		t.Fatalf("state is %s, want StateFollower", r.state)
	}
}
func TestNodeWithSmallerTermCanCompleteElection(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	n1 := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	n2 := newTestRaft(2, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	n3 := newTestRaft(3, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	n1.becomeFollower(1, None)
	n2.becomeFollower(1, None)
	n3.becomeFollower(1, None)
	n1.preVote = true
	n2.preVote = true
	n3.preVote = true
	nt := newNetwork(n1, n2, n3)
	nt.cut(1, 3)
	nt.cut(2, 3)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	sm := nt.peers[1].(*raft)
	if sm.state != StateLeader {
		t.Errorf("peer 1 state: %s, want %s", sm.state, StateLeader)
	}
	sm = nt.peers[2].(*raft)
	if sm.state != StateFollower {
		t.Errorf("peer 2 state: %s, want %s", sm.state, StateFollower)
	}
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	sm = nt.peers[3].(*raft)
	if sm.state != StatePreCandidate {
		t.Errorf("peer 3 state: %s, want %s", sm.state, StatePreCandidate)
	}
	nt.send(pb.Message{From: 2, To: 2, Type: pb.MsgHup})
	sm = nt.peers[1].(*raft)
	if sm.Term != 3 {
		t.Errorf("peer 1 term: %d, want %d", sm.Term, 3)
	}
	sm = nt.peers[2].(*raft)
	if sm.Term != 3 {
		t.Errorf("peer 2 term: %d, want %d", sm.Term, 3)
	}
	sm = nt.peers[3].(*raft)
	if sm.Term != 1 {
		t.Errorf("peer 3 term: %d, want %d", sm.Term, 1)
	}
	sm = nt.peers[1].(*raft)
	if sm.state != StateFollower {
		t.Errorf("peer 1 state: %s, want %s", sm.state, StateFollower)
	}
	sm = nt.peers[2].(*raft)
	if sm.state != StateLeader {
		t.Errorf("peer 2 state: %s, want %s", sm.state, StateLeader)
	}
	sm = nt.peers[3].(*raft)
	if sm.state != StatePreCandidate {
		t.Errorf("peer 3 state: %s, want %s", sm.state, StatePreCandidate)
	}
	sm.logger.Infof("going to bring back peer 3 and kill peer 2")
	nt.recover()
	nt.cut(2, 1)
	nt.cut(2, 3)
	nt.send(pb.Message{From: 3, To: 3, Type: pb.MsgHup})
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	sma := nt.peers[1].(*raft)
	smb := nt.peers[3].(*raft)
	if sma.state != StateLeader && smb.state != StateLeader {
		t.Errorf("no leader")
	}
}
func TestPreVoteWithSplitVote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	n1 := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	n2 := newTestRaft(2, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	n3 := newTestRaft(3, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	n1.becomeFollower(1, None)
	n2.becomeFollower(1, None)
	n3.becomeFollower(1, None)
	n1.preVote = true
	n2.preVote = true
	n3.preVote = true
	nt := newNetwork(n1, n2, n3)
	nt.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
	nt.isolate(1)
	nt.send([]pb.Message{{From: 2, To: 2, Type: pb.MsgHup}, {From: 3, To: 3, Type: pb.MsgHup}}...)
	sm := nt.peers[2].(*raft)
	if sm.Term != 3 {
		t.Errorf("peer 2 term: %d, want %d", sm.Term, 3)
	}
	sm = nt.peers[3].(*raft)
	if sm.Term != 3 {
		t.Errorf("peer 3 term: %d, want %d", sm.Term, 3)
	}
	sm = nt.peers[2].(*raft)
	if sm.state != StateCandidate {
		t.Errorf("peer 2 state: %s, want %s", sm.state, StateCandidate)
	}
	sm = nt.peers[3].(*raft)
	if sm.state != StateCandidate {
		t.Errorf("peer 3 state: %s, want %s", sm.state, StateCandidate)
	}
	nt.send(pb.Message{From: 2, To: 2, Type: pb.MsgHup})
	sm = nt.peers[2].(*raft)
	if sm.Term != 4 {
		t.Errorf("peer 2 term: %d, want %d", sm.Term, 4)
	}
	sm = nt.peers[3].(*raft)
	if sm.Term != 4 {
		t.Errorf("peer 3 term: %d, want %d", sm.Term, 4)
	}
	sm = nt.peers[2].(*raft)
	if sm.state != StateLeader {
		t.Errorf("peer 2 state: %s, want %s", sm.state, StateLeader)
	}
	sm = nt.peers[3].(*raft)
	if sm.state != StateFollower {
		t.Errorf("peer 3 state: %s, want %s", sm.state, StateFollower)
	}
}
func entsWithConfig(configFunc func(*Config), terms ...uint64) *raft {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := NewMemoryStorage()
	for i, term := range terms {
		storage.Append([]pb.Entry{{Index: uint64(i + 1), Term: term}})
	}
	cfg := newTestConfig(1, []uint64{}, 5, 1, storage)
	if configFunc != nil {
		configFunc(cfg)
	}
	sm := newRaft(cfg)
	sm.reset(terms[len(terms)-1])
	return sm
}
func votedWithConfig(configFunc func(*Config), vote, term uint64) *raft {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := NewMemoryStorage()
	storage.SetHardState(pb.HardState{Vote: vote, Term: term})
	cfg := newTestConfig(1, []uint64{}, 5, 1, storage)
	if configFunc != nil {
		configFunc(cfg)
	}
	sm := newRaft(cfg)
	sm.reset(term)
	return sm
}

type network struct {
	peers	map[uint64]stateMachine
	storage	map[uint64]*MemoryStorage
	dropm	map[connem]float64
	ignorem	map[pb.MessageType]bool
}

func newNetwork(peers ...stateMachine) *network {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newNetworkWithConfig(nil, peers...)
}
func newNetworkWithConfig(configFunc func(*Config), peers ...stateMachine) *network {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := len(peers)
	peerAddrs := idsBySize(size)
	npeers := make(map[uint64]stateMachine, size)
	nstorage := make(map[uint64]*MemoryStorage, size)
	for j, p := range peers {
		id := peerAddrs[j]
		switch v := p.(type) {
		case nil:
			nstorage[id] = NewMemoryStorage()
			cfg := newTestConfig(id, peerAddrs, 10, 1, nstorage[id])
			if configFunc != nil {
				configFunc(cfg)
			}
			sm := newRaft(cfg)
			npeers[id] = sm
		case *raft:
			learners := make(map[uint64]bool, len(v.learnerPrs))
			for i := range v.learnerPrs {
				learners[i] = true
			}
			v.id = id
			v.prs = make(map[uint64]*Progress)
			v.learnerPrs = make(map[uint64]*Progress)
			for i := 0; i < size; i++ {
				if _, ok := learners[peerAddrs[i]]; ok {
					v.learnerPrs[peerAddrs[i]] = &Progress{IsLearner: true}
				} else {
					v.prs[peerAddrs[i]] = &Progress{}
				}
			}
			v.reset(v.Term)
			npeers[id] = v
		case *blackHole:
			npeers[id] = v
		default:
			panic(fmt.Sprintf("unexpected state machine type: %T", p))
		}
	}
	return &network{peers: npeers, storage: nstorage, dropm: make(map[connem]float64), ignorem: make(map[pb.MessageType]bool)}
}
func preVoteConfig(c *Config) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.PreVote = true
}
func (nw *network) send(msgs ...pb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for len(msgs) > 0 {
		m := msgs[0]
		p := nw.peers[m.To]
		p.Step(m)
		msgs = append(msgs[1:], nw.filter(p.readMessages())...)
	}
}
func (nw *network) drop(from, to uint64, perc float64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nw.dropm[connem{from, to}] = perc
}
func (nw *network) cut(one, other uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nw.drop(one, other, 1)
	nw.drop(other, one, 1)
}
func (nw *network) isolate(id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < len(nw.peers); i++ {
		nid := uint64(i) + 1
		if nid != id {
			nw.drop(id, nid, 1.0)
			nw.drop(nid, id, 1.0)
		}
	}
}
func (nw *network) ignore(t pb.MessageType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nw.ignorem[t] = true
}
func (nw *network) recover() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nw.dropm = make(map[connem]float64)
	nw.ignorem = make(map[pb.MessageType]bool)
}
func (nw *network) filter(msgs []pb.Message) []pb.Message {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mm := []pb.Message{}
	for _, m := range msgs {
		if nw.ignorem[m.Type] {
			continue
		}
		switch m.Type {
		case pb.MsgHup:
			panic("unexpected msgHup")
		default:
			perc := nw.dropm[connem{m.From, m.To}]
			if n := rand.Float64(); n < perc {
				continue
			}
		}
		mm = append(mm, m)
	}
	return mm
}

type connem struct{ from, to uint64 }
type blackHole struct{}

func (blackHole) Step(pb.Message) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (blackHole) readMessages() []pb.Message {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

var nopStepper = &blackHole{}

func idsBySize(size int) []uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ids := make([]uint64, size)
	for i := 0; i < size; i++ {
		ids[i] = 1 + uint64(i)
	}
	return ids
}
func setRandomizedElectionTimeout(r *raft, v int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.randomizedElectionTimeout = v
}
func newTestConfig(id uint64, peers []uint64, election, heartbeat int, storage Storage) *Config {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Config{ID: id, peers: peers, ElectionTick: election, HeartbeatTick: heartbeat, Storage: storage, MaxSizePerMsg: noLimit, MaxInflightMsgs: 256}
}
func newTestRaft(id uint64, peers []uint64, election, heartbeat int, storage Storage) *raft {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newRaft(newTestConfig(id, peers, election, heartbeat, storage))
}
func newTestLearnerRaft(id uint64, peers []uint64, learners []uint64, election, heartbeat int, storage Storage) *raft {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := newTestConfig(id, peers, election, heartbeat, storage)
	cfg.learners = learners
	return newRaft(cfg)
}
