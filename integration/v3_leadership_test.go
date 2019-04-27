package integration

import (
	"context"
	"testing"
	"time"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestMoveLeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testMoveLeader(t, true)
}
func TestMoveLeaderService(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testMoveLeader(t, false)
}
func testMoveLeader(t *testing.T, auto bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	oldLeadIdx := clus.WaitLeader(t)
	oldLeadID := uint64(clus.Members[oldLeadIdx].s.ID())
	idc := make(chan uint64)
	for i := range clus.Members {
		if oldLeadIdx != i {
			go func(m *member) {
				idc <- checkLeaderTransition(t, m, oldLeadID)
			}(clus.Members[i])
		}
	}
	target := uint64(clus.Members[(oldLeadIdx+1)%3].s.ID())
	if auto {
		err := clus.Members[oldLeadIdx].s.TransferLeadership()
		if err != nil {
			t.Fatal(err)
		}
	} else {
		mvc := toGRPC(clus.Client(oldLeadIdx)).Maintenance
		_, err := mvc.MoveLeader(context.TODO(), &pb.MoveLeaderRequest{TargetID: target})
		if err != nil {
			t.Fatal(err)
		}
	}
	var newLeadIDs [2]uint64
	for i := range newLeadIDs {
		select {
		case newLeadIDs[i] = <-idc:
		case <-time.After(time.Second):
			t.Fatal("timed out waiting for leader transition")
		}
	}
	if newLeadIDs[0] != newLeadIDs[1] {
		t.Fatalf("expected same new leader %d == %d", newLeadIDs[0], newLeadIDs[1])
	}
	if oldLeadID == newLeadIDs[0] {
		t.Fatalf("expected old leader %d != new leader %d", oldLeadID, newLeadIDs[0])
	}
	if !auto {
		if newLeadIDs[0] != target {
			t.Fatalf("expected new leader %d != target %d", newLeadIDs[0], target)
		}
	}
}
func TestMoveLeaderError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	oldLeadIdx := clus.WaitLeader(t)
	followerIdx := (oldLeadIdx + 1) % 3
	target := uint64(clus.Members[(oldLeadIdx+2)%3].s.ID())
	mvc := toGRPC(clus.Client(followerIdx)).Maintenance
	_, err := mvc.MoveLeader(context.TODO(), &pb.MoveLeaderRequest{TargetID: target})
	if !eqErrGRPC(err, rpctypes.ErrGRPCNotLeader) {
		t.Errorf("err = %v, want %v", err, rpctypes.ErrGRPCNotLeader)
	}
}
