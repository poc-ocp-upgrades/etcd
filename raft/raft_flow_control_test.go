package raft

import (
	"testing"
	pb "github.com/coreos/etcd/raft/raftpb"
)

func TestMsgAppFlowControlFull(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 5, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	pr2 := r.prs[2]
	pr2.becomeReplicate()
	for i := 0; i < r.maxInflight; i++ {
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
		ms := r.readMessages()
		if len(ms) != 1 {
			t.Fatalf("#%d: len(ms) = %d, want 1", i, len(ms))
		}
	}
	if !pr2.ins.full() {
		t.Fatalf("inflights.full = %t, want %t", pr2.ins.full(), true)
	}
	for i := 0; i < 10; i++ {
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
		ms := r.readMessages()
		if len(ms) != 0 {
			t.Fatalf("#%d: len(ms) = %d, want 0", i, len(ms))
		}
	}
}
func TestMsgAppFlowControlMoveForward(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 5, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	pr2 := r.prs[2]
	pr2.becomeReplicate()
	for i := 0; i < r.maxInflight; i++ {
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
		r.readMessages()
	}
	for tt := 2; tt < r.maxInflight; tt++ {
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgAppResp, Index: uint64(tt)})
		r.readMessages()
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
		ms := r.readMessages()
		if len(ms) != 1 {
			t.Fatalf("#%d: len(ms) = %d, want 1", tt, len(ms))
		}
		if !pr2.ins.full() {
			t.Fatalf("inflights.full = %t, want %t", pr2.ins.full(), true)
		}
		for i := 0; i < tt; i++ {
			r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgAppResp, Index: uint64(i)})
			if !pr2.ins.full() {
				t.Fatalf("#%d: inflights.full = %t, want %t", tt, pr2.ins.full(), true)
			}
		}
	}
}
func TestMsgAppFlowControlRecvHeartbeat(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := newTestRaft(1, []uint64{1, 2}, 5, 1, NewMemoryStorage())
	r.becomeCandidate()
	r.becomeLeader()
	pr2 := r.prs[2]
	pr2.becomeReplicate()
	for i := 0; i < r.maxInflight; i++ {
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
		r.readMessages()
	}
	for tt := 1; tt < 5; tt++ {
		if !pr2.ins.full() {
			t.Fatalf("#%d: inflights.full = %t, want %t", tt, pr2.ins.full(), true)
		}
		for i := 0; i < tt; i++ {
			r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgHeartbeatResp})
			r.readMessages()
			if pr2.ins.full() {
				t.Fatalf("#%d.%d: inflights.full = %t, want %t", tt, i, pr2.ins.full(), false)
			}
		}
		r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
		ms := r.readMessages()
		if len(ms) != 1 {
			t.Fatalf("#%d: free slot = 0, want 1", tt)
		}
		for i := 0; i < 10; i++ {
			r.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
			ms1 := r.readMessages()
			if len(ms1) != 0 {
				t.Fatalf("#%d.%d: len(ms) = %d, want 0", tt, i, len(ms1))
			}
		}
		r.Step(pb.Message{From: 2, To: 1, Type: pb.MsgHeartbeatResp})
		r.readMessages()
	}
}
