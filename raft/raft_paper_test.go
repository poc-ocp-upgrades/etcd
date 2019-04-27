package raft

import (
	"fmt"
	"testing"
	"reflect"
	"sort"
	pb "github.com/coreos/etcd/raft/raftpb"
)

func TestFollowerUpdateTermFromMessage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testUpdateTermFromMessage(t, StateFollower)
}
func TestCandidateUpdateTermFromMessage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testUpdateTermFromMessage(t, StateCandidate)
}
func TestLeaderUpdateTermFromMessage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testUpdateTermFromMessage(t, StateLeader)
}
func testUpdateTermFromMessage(t *testing.T, state StateType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	switch state {
	case StateFollower:
		r.becomeFollower(1, 2)
	case StateCandidate:
		r.becomeCandidate()
	case StateLeader:
		r.becomeCandidate()
		r.becomeLeader()
	}
	r.Step(pb.Message{Type: pb.MsgApp, Term: 2})
	if r.Term != 2 {
		t.Errorf("term = %d, want %d", r.Term, 2)
	}
	if r.state != StateFollower {
		t.Errorf("state = %v, want %v", r.state, StateFollower)
	}
}
func TestRejectStaleTermMessage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	called := false
	fakeStep := func(r *raft, m pb.Message) {
		called = true
	}
	r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	r.step = fakeStep
	r.loadState(pb.HardState{Term: 2})
	r.Step(pb.Message{Type: pb.MsgApp, Term: r.Term - 1})
	if called {
		t.Errorf("stepFunc called = %v, want %v", called, false)
	}
}
func TestStartAsFollower(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
	if r.state != StateFollower {
		t.Errorf("state = %s, want %s", r.state, StateFollower)
	}
}
func TestLeaderBcastBeat(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hi := 1
	r := newTestRaft(1, []uint64{1, 2, 3}, 10, hi, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	for i := 0; i < 10; i++ {
		r.appendEntry(pb.Entry{Index: uint64(i) + 1})
	}
	for i := 0; i < hi; i++ {
		r.tick()
	}
	msgs := r.readMessages()
	sort.Sort(messageSlice(msgs))
	wmsgs := []pb.Message{{From: 1, To: 2, Term: 1, Type: pb.MsgHeartbeat}, {From: 1, To: 3, Term: 1, Type: pb.MsgHeartbeat}}
	if !reflect.DeepEqual(msgs, wmsgs) {
		t.Errorf("msgs = %v, want %v", msgs, wmsgs)
	}
}
func TestFollowerStartElection(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testNonleaderStartElection(t, StateFollower)
}
func TestCandidateStartNewElection(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testNonleaderStartElection(t, StateCandidate)
}
func testNonleaderStartElection(t *testing.T, state StateType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	et := 10
	r := newTestRaft(1, []uint64{1, 2, 3}, et, 1, NewMemoryStorage())
	switch state {
	case StateFollower:
		r.becomeFollower(1, 2)
	case StateCandidate:
		r.becomeCandidate()
	}
	for i := 1; i < 2*et; i++ {
		r.tick()
	}
	if r.Term != 2 {
		t.Errorf("term = %d, want 2", r.Term)
	}
	if r.state != StateCandidate {
		t.Errorf("state = %s, want %s", r.state, StateCandidate)
	}
	if !r.votes[r.id] {
		t.Errorf("vote for self = false, want true")
	}
	msgs := r.readMessages()
	sort.Sort(messageSlice(msgs))
	wmsgs := []pb.Message{{From: 1, To: 2, Term: 2, Type: pb.MsgVote}, {From: 1, To: 3, Term: 2, Type: pb.MsgVote}}
	if !reflect.DeepEqual(msgs, wmsgs) {
		t.Errorf("msgs = %v, want %v", msgs, wmsgs)
	}
}
func TestLeaderElectionInOneRoundRPC(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		size	int
		votes	map[uint64]bool
		state	StateType
	}{{1, map[uint64]bool{}, StateLeader}, {3, map[uint64]bool{2: true, 3: true}, StateLeader}, {3, map[uint64]bool{2: true}, StateLeader}, {5, map[uint64]bool{2: true, 3: true, 4: true, 5: true}, StateLeader}, {5, map[uint64]bool{2: true, 3: true, 4: true}, StateLeader}, {5, map[uint64]bool{2: true, 3: true}, StateLeader}, {3, map[uint64]bool{2: false, 3: false}, StateFollower}, {5, map[uint64]bool{2: false, 3: false, 4: false, 5: false}, StateFollower}, {5, map[uint64]bool{2: true, 3: false, 4: false, 5: false}, StateFollower}, {3, map[uint64]bool{}, StateCandidate}, {5, map[uint64]bool{2: true}, StateCandidate}, {5, map[uint64]bool{2: false, 3: false}, StateCandidate}, {5, map[uint64]bool{}, StateCandidate}}
	for i, tt := range tests {
		r := newTestRaft(1, idsBySize(tt.size), 10, 1, NewMemoryStorage())
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
		for id, vote := range tt.votes {
			r.Step(pb.Message{From: id, To: 1, Type: pb.MsgVoteResp, Reject: !vote})
		}
		if r.state != tt.state {
			t.Errorf("#%d: state = %s, want %s", i, r.state, tt.state)
		}
		if g := r.Term; g != 1 {
			t.Errorf("#%d: term = %d, want %d", i, g, 1)
		}
	}
}
func TestFollowerVote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		vote	uint64
		nvote	uint64
		wreject	bool
	}{{None, 1, false}, {None, 2, false}, {1, 1, false}, {2, 2, false}, {1, 2, true}, {2, 1, true}}
	for i, tt := range tests {
		r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
		r.loadState(pb.HardState{Term: 1, Vote: tt.vote})
		r.Step(pb.Message{From: tt.nvote, To: 1, Term: 1, Type: pb.MsgVote})
		msgs := r.readMessages()
		wmsgs := []pb.Message{{From: 1, To: tt.nvote, Term: 1, Type: pb.MsgVoteResp, Reject: tt.wreject}}
		if !reflect.DeepEqual(msgs, wmsgs) {
			t.Errorf("#%d: msgs = %v, want %v", i, msgs, wmsgs)
		}
	}
}
func TestCandidateFallback(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []pb.Message{{From: 2, To: 1, Term: 1, Type: pb.MsgApp}, {From: 2, To: 1, Term: 2, Type: pb.MsgApp}}
	for i, tt := range tests {
		r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
		if r.state != StateCandidate {
			t.Fatalf("unexpected state = %s, want %s", r.state, StateCandidate)
		}
		r.Step(tt)
		if g := r.state; g != StateFollower {
			t.Errorf("#%d: state = %s, want %s", i, g, StateFollower)
		}
		if g := r.Term; g != tt.Term {
			t.Errorf("#%d: term = %d, want %d", i, g, tt.Term)
		}
	}
}
func TestFollowerElectionTimeoutRandomized(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetLogger(discardLogger)
	defer SetLogger(defaultLogger)
	testNonleaderElectionTimeoutRandomized(t, StateFollower)
}
func TestCandidateElectionTimeoutRandomized(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetLogger(discardLogger)
	defer SetLogger(defaultLogger)
	testNonleaderElectionTimeoutRandomized(t, StateCandidate)
}
func testNonleaderElectionTimeoutRandomized(t *testing.T, state StateType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	et := 10
	r := newTestRaft(1, []uint64{1, 2, 3}, et, 1, NewMemoryStorage())
	timeouts := make(map[int]bool)
	for round := 0; round < 50*et; round++ {
		switch state {
		case StateFollower:
			r.becomeFollower(r.Term+1, 2)
		case StateCandidate:
			r.becomeCandidate()
		}
		time := 0
		for len(r.readMessages()) == 0 {
			r.tick()
			time++
		}
		timeouts[time] = true
	}
	for d := et + 1; d < 2*et; d++ {
		if !timeouts[d] {
			t.Errorf("timeout in %d ticks should happen", d)
		}
	}
}
func TestFollowersElectioinTimeoutNonconflict(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetLogger(discardLogger)
	defer SetLogger(defaultLogger)
	testNonleadersElectionTimeoutNonconflict(t, StateFollower)
}
func TestCandidatesElectionTimeoutNonconflict(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetLogger(discardLogger)
	defer SetLogger(defaultLogger)
	testNonleadersElectionTimeoutNonconflict(t, StateCandidate)
}
func testNonleadersElectionTimeoutNonconflict(t *testing.T, state StateType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	et := 10
	size := 5
	rs := make([]*raft, size)
	ids := idsBySize(size)
	for k := range rs {
		rs[k] = newTestRaft(ids[k], ids, et, 1, NewMemoryStorage())
	}
	conflicts := 0
	for round := 0; round < 1000; round++ {
		for _, r := range rs {
			switch state {
			case StateFollower:
				r.becomeFollower(r.Term+1, None)
			case StateCandidate:
				r.becomeCandidate()
			}
		}
		timeoutNum := 0
		for timeoutNum == 0 {
			for _, r := range rs {
				r.tick()
				if len(r.readMessages()) > 0 {
					timeoutNum++
				}
			}
		}
		if timeoutNum > 1 {
			conflicts++
		}
	}
	if g := float64(conflicts) / 1000; g > 0.3 {
		t.Errorf("probability of conflicts = %v, want <= 0.3", g)
	}
}
func TestLeaderStartReplication(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := NewMemoryStorage()
	r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, s)
	r.becomeCandidate()
	r.becomeLeader()
	commitNoopEntry(r, s)
	li := r.raftLog.lastIndex()
	ents := []pb.Entry{{Data: []byte("some data")}}
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: ents})
	if g := r.raftLog.lastIndex(); g != li+1 {
		t.Errorf("lastIndex = %d, want %d", g, li+1)
	}
	if g := r.raftLog.committed; g != li {
		t.Errorf("committed = %d, want %d", g, li)
	}
	msgs := r.readMessages()
	sort.Sort(messageSlice(msgs))
	wents := []pb.Entry{{Index: li + 1, Term: 1, Data: []byte("some data")}}
	wmsgs := []pb.Message{{From: 1, To: 2, Term: 1, Type: pb.MsgApp, Index: li, LogTerm: 1, Entries: wents, Commit: li}, {From: 1, To: 3, Term: 1, Type: pb.MsgApp, Index: li, LogTerm: 1, Entries: wents, Commit: li}}
	if !reflect.DeepEqual(msgs, wmsgs) {
		t.Errorf("msgs = %+v, want %+v", msgs, wmsgs)
	}
	if g := r.raftLog.unstableEntries(); !reflect.DeepEqual(g, wents) {
		t.Errorf("ents = %+v, want %+v", g, wents)
	}
}
func TestLeaderCommitEntry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := NewMemoryStorage()
	r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, s)
	r.becomeCandidate()
	r.becomeLeader()
	commitNoopEntry(r, s)
	li := r.raftLog.lastIndex()
	r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
	for _, m := range r.readMessages() {
		r.Step(acceptAndReply(m))
	}
	if g := r.raftLog.committed; g != li+1 {
		t.Errorf("committed = %d, want %d", g, li+1)
	}
	wents := []pb.Entry{{Index: li + 1, Term: 1, Data: []byte("some data")}}
	if g := r.raftLog.nextEnts(); !reflect.DeepEqual(g, wents) {
		t.Errorf("nextEnts = %+v, want %+v", g, wents)
	}
	msgs := r.readMessages()
	sort.Sort(messageSlice(msgs))
	for i, m := range msgs {
		if w := uint64(i + 2); m.To != w {
			t.Errorf("to = %x, want %x", m.To, w)
		}
		if m.Type != pb.MsgApp {
			t.Errorf("type = %v, want %v", m.Type, pb.MsgApp)
		}
		if m.Commit != li+1 {
			t.Errorf("commit = %d, want %d", m.Commit, li+1)
		}
	}
}
func TestLeaderAcknowledgeCommit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		size		int
		acceptors	map[uint64]bool
		wack		bool
	}{{1, nil, true}, {3, nil, false}, {3, map[uint64]bool{2: true}, true}, {3, map[uint64]bool{2: true, 3: true}, true}, {5, nil, false}, {5, map[uint64]bool{2: true}, false}, {5, map[uint64]bool{2: true, 3: true}, true}, {5, map[uint64]bool{2: true, 3: true, 4: true}, true}, {5, map[uint64]bool{2: true, 3: true, 4: true, 5: true}, true}}
	for i, tt := range tests {
		s := NewMemoryStorage()
		r := newTestRaft(1, idsBySize(tt.size), 10, 1, s)
		r.becomeCandidate()
		r.becomeLeader()
		commitNoopEntry(r, s)
		li := r.raftLog.lastIndex()
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
		for _, m := range r.readMessages() {
			if tt.acceptors[m.To] {
				r.Step(acceptAndReply(m))
			}
		}
		if g := r.raftLog.committed > li; g != tt.wack {
			t.Errorf("#%d: ack commit = %v, want %v", i, g, tt.wack)
		}
	}
}
func TestLeaderCommitPrecedingEntries(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := [][]pb.Entry{{}, {{Term: 2, Index: 1}}, {{Term: 1, Index: 1}, {Term: 2, Index: 2}}, {{Term: 1, Index: 1}}}
	for i, tt := range tests {
		storage := NewMemoryStorage()
		storage.Append(tt)
		r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, storage)
		r.loadState(pb.HardState{Term: 2})
		r.becomeCandidate()
		r.becomeLeader()
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("some data")}}})
		for _, m := range r.readMessages() {
			r.Step(acceptAndReply(m))
		}
		li := uint64(len(tt))
		wents := append(tt, pb.Entry{Term: 3, Index: li + 1}, pb.Entry{Term: 3, Index: li + 2, Data: []byte("some data")})
		if g := r.raftLog.nextEnts(); !reflect.DeepEqual(g, wents) {
			t.Errorf("#%d: ents = %+v, want %+v", i, g, wents)
		}
	}
}
func TestFollowerCommitEntry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		ents	[]pb.Entry
		commit	uint64
	}{{[]pb.Entry{{Term: 1, Index: 1, Data: []byte("some data")}}, 1}, {[]pb.Entry{{Term: 1, Index: 1, Data: []byte("some data")}, {Term: 1, Index: 2, Data: []byte("some data2")}}, 2}, {[]pb.Entry{{Term: 1, Index: 1, Data: []byte("some data2")}, {Term: 1, Index: 2, Data: []byte("some data")}}, 2}, {[]pb.Entry{{Term: 1, Index: 1, Data: []byte("some data")}, {Term: 1, Index: 2, Data: []byte("some data2")}}, 1}}
	for i, tt := range tests {
		r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
		r.becomeFollower(1, 2)
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgApp, Term: 1, Entries: tt.ents, Commit: tt.commit})
		if g := r.raftLog.committed; g != tt.commit {
			t.Errorf("#%d: committed = %d, want %d", i, g, tt.commit)
		}
		wents := tt.ents[:int(tt.commit)]
		if g := r.raftLog.nextEnts(); !reflect.DeepEqual(g, wents) {
			t.Errorf("#%d: nextEnts = %v, want %v", i, g, wents)
		}
	}
}
func TestFollowerCheckMsgApp(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ents := []pb.Entry{{Term: 1, Index: 1}, {Term: 2, Index: 2}}
	tests := []struct {
		term		uint64
		index		uint64
		windex		uint64
		wreject		bool
		wrejectHint	uint64
	}{{0, 0, 1, false, 0}, {ents[0].Term, ents[0].Index, 1, false, 0}, {ents[1].Term, ents[1].Index, 2, false, 0}, {ents[0].Term, ents[1].Index, ents[1].Index, true, 2}, {ents[1].Term + 1, ents[1].Index + 1, ents[1].Index + 1, true, 2}}
	for i, tt := range tests {
		storage := NewMemoryStorage()
		storage.Append(ents)
		r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, storage)
		r.loadState(pb.HardState{Commit: 1})
		r.becomeFollower(2, 2)
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgApp, Term: 2, LogTerm: tt.term, Index: tt.index})
		msgs := r.readMessages()
		wmsgs := []pb.Message{{From: 1, To: 2, Type: pb.MsgAppResp, Term: 2, Index: tt.windex, Reject: tt.wreject, RejectHint: tt.wrejectHint}}
		if !reflect.DeepEqual(msgs, wmsgs) {
			t.Errorf("#%d: msgs = %+v, want %+v", i, msgs, wmsgs)
		}
	}
}
func TestFollowerAppendEntries(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		index, term	uint64
		ents		[]pb.Entry
		wents		[]pb.Entry
		wunstable	[]pb.Entry
	}{{2, 2, []pb.Entry{{Term: 3, Index: 3}}, []pb.Entry{{Term: 1, Index: 1}, {Term: 2, Index: 2}, {Term: 3, Index: 3}}, []pb.Entry{{Term: 3, Index: 3}}}, {1, 1, []pb.Entry{{Term: 3, Index: 2}, {Term: 4, Index: 3}}, []pb.Entry{{Term: 1, Index: 1}, {Term: 3, Index: 2}, {Term: 4, Index: 3}}, []pb.Entry{{Term: 3, Index: 2}, {Term: 4, Index: 3}}}, {0, 0, []pb.Entry{{Term: 1, Index: 1}}, []pb.Entry{{Term: 1, Index: 1}, {Term: 2, Index: 2}}, nil}, {0, 0, []pb.Entry{{Term: 3, Index: 1}}, []pb.Entry{{Term: 3, Index: 1}}, []pb.Entry{{Term: 3, Index: 1}}}}
	for i, tt := range tests {
		storage := NewMemoryStorage()
		storage.Append([]pb.Entry{{Term: 1, Index: 1}, {Term: 2, Index: 2}})
		r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, storage)
		r.becomeFollower(2, 2)
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgApp, Term: 2, LogTerm: tt.term, Index: tt.index, Entries: tt.ents})
		if g := r.raftLog.allEntries(); !reflect.DeepEqual(g, tt.wents) {
			t.Errorf("#%d: ents = %+v, want %+v", i, g, tt.wents)
		}
		if g := r.raftLog.unstableEntries(); !reflect.DeepEqual(g, tt.wunstable) {
			t.Errorf("#%d: unstableEnts = %+v, want %+v", i, g, tt.wunstable)
		}
	}
}
func TestLeaderSyncFollowerLog(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ents := []pb.Entry{{}, {Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}, {Term: 4, Index: 4}, {Term: 4, Index: 5}, {Term: 5, Index: 6}, {Term: 5, Index: 7}, {Term: 6, Index: 8}, {Term: 6, Index: 9}, {Term: 6, Index: 10}}
	term := uint64(8)
	tests := [][]pb.Entry{{{}, {Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}, {Term: 4, Index: 4}, {Term: 4, Index: 5}, {Term: 5, Index: 6}, {Term: 5, Index: 7}, {Term: 6, Index: 8}, {Term: 6, Index: 9}}, {{}, {Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}, {Term: 4, Index: 4}}, {{}, {Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}, {Term: 4, Index: 4}, {Term: 4, Index: 5}, {Term: 5, Index: 6}, {Term: 5, Index: 7}, {Term: 6, Index: 8}, {Term: 6, Index: 9}, {Term: 6, Index: 10}, {Term: 6, Index: 11}}, {{}, {Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}, {Term: 4, Index: 4}, {Term: 4, Index: 5}, {Term: 5, Index: 6}, {Term: 5, Index: 7}, {Term: 6, Index: 8}, {Term: 6, Index: 9}, {Term: 6, Index: 10}, {Term: 7, Index: 11}, {Term: 7, Index: 12}}, {{}, {Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}, {Term: 4, Index: 4}, {Term: 4, Index: 5}, {Term: 4, Index: 6}, {Term: 4, Index: 7}}, {{}, {Term: 1, Index: 1}, {Term: 1, Index: 2}, {Term: 1, Index: 3}, {Term: 2, Index: 4}, {Term: 2, Index: 5}, {Term: 2, Index: 6}, {Term: 3, Index: 7}, {Term: 3, Index: 8}, {Term: 3, Index: 9}, {Term: 3, Index: 10}, {Term: 3, Index: 11}}}
	for i, tt := range tests {
		leadStorage := NewMemoryStorage()
		leadStorage.Append(ents)
		lead := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, leadStorage)
		lead.loadState(pb.HardState{Commit: lead.raftLog.lastIndex(), Term: term})
		followerStorage := NewMemoryStorage()
		followerStorage.Append(tt)
		follower := newTestRaft(2, []uint64{1, 2, 3}, 10, 1, followerStorage)
		follower.loadState(pb.HardState{Term: term - 1})
		n := newNetwork(lead, follower, nopStepper)
		n.send(pb.Message{From: 1, To: 1, Type: pb.MsgHup})
		n.send(pb.Message{From: 3, To: 1, Type: pb.MsgVoteResp, Term: term + 1})
		n.send(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
		if g := diffu(ltoa(lead.raftLog), ltoa(follower.raftLog)); g != "" {
			t.Errorf("#%d: log diff:\n%s", i, g)
		}
	}
}
func TestVoteRequest(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		ents	[]pb.Entry
		wterm	uint64
	}{{[]pb.Entry{{Term: 1, Index: 1}}, 2}, {[]pb.Entry{{Term: 1, Index: 1}, {Term: 2, Index: 2}}, 3}}
	for j, tt := range tests {
		r := newTestRaft(1, []uint64{1, 2, 3}, 10, 1, NewMemoryStorage())
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgApp, Term: tt.wterm - 1, LogTerm: 0, Index: 0, Entries: tt.ents})
		r.readMessages()
		for i := 1; i < r.electionTimeout*2; i++ {
			r.tickElection()
		}
		msgs := r.readMessages()
		sort.Sort(messageSlice(msgs))
		if len(msgs) != 2 {
			t.Fatalf("#%d: len(msg) = %d, want %d", j, len(msgs), 2)
		}
		for i, m := range msgs {
			if m.Type != pb.MsgVote {
				t.Errorf("#%d: msgType = %d, want %d", i, m.Type, pb.MsgVote)
			}
			if m.To != uint64(i+2) {
				t.Errorf("#%d: to = %d, want %d", i, m.To, i+2)
			}
			if m.Term != tt.wterm {
				t.Errorf("#%d: term = %d, want %d", i, m.Term, tt.wterm)
			}
			windex, wlogterm := tt.ents[len(tt.ents)-1].Index, tt.ents[len(tt.ents)-1].Term
			if m.Index != windex {
				t.Errorf("#%d: index = %d, want %d", i, m.Index, windex)
			}
			if m.LogTerm != wlogterm {
				t.Errorf("#%d: logterm = %d, want %d", i, m.LogTerm, wlogterm)
			}
		}
	}
}
func TestVoter(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		ents	[]pb.Entry
		logterm	uint64
		index	uint64
		wreject	bool
	}{{[]pb.Entry{{Term: 1, Index: 1}}, 1, 1, false}, {[]pb.Entry{{Term: 1, Index: 1}}, 1, 2, false}, {[]pb.Entry{{Term: 1, Index: 1}, {Term: 1, Index: 2}}, 1, 1, true}, {[]pb.Entry{{Term: 1, Index: 1}}, 2, 1, false}, {[]pb.Entry{{Term: 1, Index: 1}}, 2, 2, false}, {[]pb.Entry{{Term: 1, Index: 1}, {Term: 1, Index: 2}}, 2, 1, false}, {[]pb.Entry{{Term: 2, Index: 1}}, 1, 1, true}, {[]pb.Entry{{Term: 2, Index: 1}}, 1, 2, true}, {[]pb.Entry{{Term: 2, Index: 1}, {Term: 1, Index: 2}}, 1, 1, true}}
	for i, tt := range tests {
		storage := NewMemoryStorage()
		storage.Append(tt.ents)
		r := newTestRaft(1, []uint64{1, 2}, 10, 1, storage)
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgVote, Term: 3, LogTerm: tt.logterm, Index: tt.index})
		msgs := r.readMessages()
		if len(msgs) != 1 {
			t.Fatalf("#%d: len(msg) = %d, want %d", i, len(msgs), 1)
		}
		m := msgs[0]
		if m.Type != pb.MsgVoteResp {
			t.Errorf("#%d: msgType = %d, want %d", i, m.Type, pb.MsgVoteResp)
		}
		if m.Reject != tt.wreject {
			t.Errorf("#%d: reject = %t, want %t", i, m.Reject, tt.wreject)
		}
	}
}
func TestLeaderOnlyCommitsLogFromCurrentTerm(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ents := []pb.Entry{{Term: 1, Index: 1}, {Term: 2, Index: 2}}
	tests := []struct {
		index	uint64
		wcommit	uint64
	}{{1, 0}, {2, 0}, {3, 3}}
	for i, tt := range tests {
		storage := NewMemoryStorage()
		storage.Append(ents)
		r := newTestRaft(1, []uint64{1, 2}, 10, 1, storage)
		r.loadState(pb.HardState{Term: 2})
		r.becomeCandidate()
		r.becomeLeader()
		r.readMessages()
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{}}})
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgAppResp, Term: r.Term, Index: tt.index})
		if r.raftLog.committed != tt.wcommit {
			t.Errorf("#%d: commit = %d, want %d", i, r.raftLog.committed, tt.wcommit)
		}
	}
}

type messageSlice []pb.Message

func (s messageSlice) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(s)
}
func (s messageSlice) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprint(s[i]) < fmt.Sprint(s[j])
}
func (s messageSlice) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s[i], s[j] = s[j], s[i]
}
func commitNoopEntry(r *raft, s *MemoryStorage) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.state != StateLeader {
		panic("it should only be used when it is the leader")
	}
	r.bcastAppend()
	msgs := r.readMessages()
	for _, m := range msgs {
		if m.Type != pb.MsgApp || len(m.Entries) != 1 || m.Entries[0].Data != nil {
			panic("not a message to append noop entry")
		}
		r.Step(acceptAndReply(m))
	}
	r.readMessages()
	s.Append(r.raftLog.unstableEntries())
	r.raftLog.appliedTo(r.raftLog.committed)
	r.raftLog.stableTo(r.raftLog.lastIndex(), r.raftLog.lastTerm())
}
func acceptAndReply(m pb.Message) pb.Message {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.Type != pb.MsgApp {
		panic("type should be MsgApp")
	}
	return pb.Message{From: m.To, To: m.From, Term: m.Term, Type: pb.MsgAppResp, Index: m.Index + uint64(len(m.Entries))}
}
