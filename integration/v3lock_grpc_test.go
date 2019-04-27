package integration

import (
	"context"
	"testing"
	"time"
	lockpb "github.com/coreos/etcd/etcdserver/api/v3lock/v3lockpb"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestV3LockLockWaiter(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	lease1, err1 := toGRPC(clus.RandClient()).Lease.LeaseGrant(context.TODO(), &pb.LeaseGrantRequest{TTL: 30})
	if err1 != nil {
		t.Fatal(err1)
	}
	lease2, err2 := toGRPC(clus.RandClient()).Lease.LeaseGrant(context.TODO(), &pb.LeaseGrantRequest{TTL: 30})
	if err2 != nil {
		t.Fatal(err2)
	}
	lc := toGRPC(clus.Client(0)).Lock
	l1, lerr1 := lc.Lock(context.TODO(), &lockpb.LockRequest{Name: []byte("foo"), Lease: lease1.ID})
	if lerr1 != nil {
		t.Fatal(lerr1)
	}
	lockc := make(chan struct{})
	go func() {
		l2, lerr2 := lc.Lock(context.TODO(), &lockpb.LockRequest{Name: []byte("foo"), Lease: lease2.ID})
		if lerr2 != nil {
			t.Fatal(lerr2)
		}
		if l1.Header.Revision >= l2.Header.Revision {
			t.Fatalf("expected l1 revision < l2 revision, got %d >= %d", l1.Header.Revision, l2.Header.Revision)
		}
		close(lockc)
	}()
	select {
	case <-time.After(200 * time.Millisecond):
	case <-lockc:
		t.Fatalf("locked before unlock")
	}
	if _, uerr := lc.Unlock(context.TODO(), &lockpb.UnlockRequest{Key: l1.Key}); uerr != nil {
		t.Fatal(uerr)
	}
	select {
	case <-time.After(200 * time.Millisecond):
		t.Fatalf("waiter did not lock after unlock")
	case <-lockc:
	}
}
