package raft

import (
	"testing"
	pb "github.com/coreos/etcd/raft/raftpb"
)

var (
	testingSnap = pb.Snapshot{Metadata: pb.SnapshotMetadata{Index: 11, Term: 11, ConfState: pb.ConfState{Nodes: []uint64{1, 2}}}}
)

func TestSendingSnapshotSetPendingSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := NewMemoryStorage()
	sm := newTestRaft(1, []uint64{1}, 10, 1, storage)
	sm.restore(testingSnap)
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.prs[2].Next = sm.raftLog.firstIndex()
	sm.Step(pb.Message{From: 2, To: 1, Type: pb.MsgAppResp, Index: sm.prs[2].Next - 1, Reject: true})
	if sm.prs[2].PendingSnapshot != 11 {
		t.Fatalf("PendingSnapshot = %d, want 11", sm.prs[2].PendingSnapshot)
	}
}
func TestPendingSnapshotPauseReplication(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := NewMemoryStorage()
	sm := newTestRaft(1, []uint64{1, 2}, 10, 1, storage)
	sm.restore(testingSnap)
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.prs[2].becomeSnapshot(11)
	sm.Step(pb.Message{From: 1, To: 1, Type: pb.MsgProp, Entries: []pb.Entry{{Data: []byte("somedata")}}})
	msgs := sm.readMessages()
	if len(msgs) != 0 {
		t.Fatalf("len(msgs) = %d, want 0", len(msgs))
	}
}
func TestSnapshotFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := NewMemoryStorage()
	sm := newTestRaft(1, []uint64{1, 2}, 10, 1, storage)
	sm.restore(testingSnap)
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.prs[2].Next = 1
	sm.prs[2].becomeSnapshot(11)
	sm.Step(pb.Message{From: 2, To: 1, Type: pb.MsgSnapStatus, Reject: true})
	if sm.prs[2].PendingSnapshot != 0 {
		t.Fatalf("PendingSnapshot = %d, want 0", sm.prs[2].PendingSnapshot)
	}
	if sm.prs[2].Next != 1 {
		t.Fatalf("Next = %d, want 1", sm.prs[2].Next)
	}
	if !sm.prs[2].Paused {
		t.Errorf("Paused = %v, want true", sm.prs[2].Paused)
	}
}
func TestSnapshotSucceed(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := NewMemoryStorage()
	sm := newTestRaft(1, []uint64{1, 2}, 10, 1, storage)
	sm.restore(testingSnap)
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.prs[2].Next = 1
	sm.prs[2].becomeSnapshot(11)
	sm.Step(pb.Message{From: 2, To: 1, Type: pb.MsgSnapStatus, Reject: false})
	if sm.prs[2].PendingSnapshot != 0 {
		t.Fatalf("PendingSnapshot = %d, want 0", sm.prs[2].PendingSnapshot)
	}
	if sm.prs[2].Next != 12 {
		t.Fatalf("Next = %d, want 12", sm.prs[2].Next)
	}
	if !sm.prs[2].Paused {
		t.Errorf("Paused = %v, want true", sm.prs[2].Paused)
	}
}
func TestSnapshotAbort(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := NewMemoryStorage()
	sm := newTestRaft(1, []uint64{1, 2}, 10, 1, storage)
	sm.restore(testingSnap)
	sm.becomeCandidate()
	sm.becomeLeader()
	sm.prs[2].Next = 1
	sm.prs[2].becomeSnapshot(11)
	sm.Step(pb.Message{From: 2, To: 1, Type: pb.MsgAppResp, Index: 11})
	if sm.prs[2].PendingSnapshot != 0 {
		t.Fatalf("PendingSnapshot = %d, want 0", sm.prs[2].PendingSnapshot)
	}
	if sm.prs[2].Next != 12 {
		t.Fatalf("Next = %d, want 12", sm.prs[2].Next)
	}
}
