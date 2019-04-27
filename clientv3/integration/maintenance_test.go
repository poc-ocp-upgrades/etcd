package integration

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestMaintenanceHashKV(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	for i := 0; i < 3; i++ {
		if _, err := clus.RandClient().Put(context.Background(), "foo", "bar"); err != nil {
			t.Fatal(err)
		}
	}
	var hv uint32
	for i := 0; i < 3; i++ {
		cli := clus.Client(i)
		if _, err := cli.Get(context.TODO(), "foo"); err != nil {
			t.Fatal(err)
		}
		hresp, err := cli.HashKV(context.Background(), clus.Members[i].GRPCAddr(), 0)
		if err != nil {
			t.Fatal(err)
		}
		if hv == 0 {
			hv = hresp.Hash
			continue
		}
		if hv != hresp.Hash {
			t.Fatalf("#%d: hash expected %d, got %d", i, hv, hresp.Hash)
		}
	}
}
func TestMaintenanceMoveLeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	oldLeadIdx := clus.WaitLeader(t)
	targetIdx := (oldLeadIdx + 1) % 3
	target := uint64(clus.Members[targetIdx].ID())
	cli := clus.Client(targetIdx)
	_, err := cli.MoveLeader(context.Background(), target)
	if err != rpctypes.ErrNotLeader {
		t.Fatalf("error expected %v, got %v", rpctypes.ErrNotLeader, err)
	}
	cli = clus.Client(oldLeadIdx)
	_, err = cli.MoveLeader(context.Background(), target)
	if err != nil {
		t.Fatal(err)
	}
	leadIdx := clus.WaitLeader(t)
	lead := uint64(clus.Members[leadIdx].ID())
	if target != lead {
		t.Fatalf("new leader expected %d, got %d", target, lead)
	}
}
func TestMaintenanceSnapshotError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	ctx, cancel := context.WithCancel(context.Background())
	rc1, err := clus.RandClient().Snapshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rc1.Close()
	cancel()
	_, err = io.Copy(ioutil.Discard, rc1)
	if err != context.Canceled {
		t.Errorf("expected %v, got %v", context.Canceled, err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rc2, err := clus.RandClient().Snapshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rc2.Close()
	time.Sleep(2 * time.Second)
	_, err = io.Copy(ioutil.Discard, rc2)
	if err != nil && err != context.DeadlineExceeded {
		t.Errorf("expected %v, got %v", context.DeadlineExceeded, err)
	}
}
func TestMaintenanceSnapshotErrorInflight(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	clus.Members[0].Stop(t)
	dpath := filepath.Join(clus.Members[0].DataDir, "member", "snap", "db")
	b := backend.NewDefaultBackend(dpath)
	s := mvcc.NewStore(b, &lease.FakeLessor{}, nil)
	rev := 100000
	for i := 2; i <= rev; i++ {
		s.Put([]byte(fmt.Sprintf("%10d", i)), bytes.Repeat([]byte("a"), 1024), lease.NoLease)
	}
	s.Close()
	b.Close()
	clus.Members[0].Restart(t)
	ctx, cancel := context.WithCancel(context.Background())
	rc1, err := clus.RandClient().Snapshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rc1.Close()
	donec := make(chan struct{})
	go func() {
		time.Sleep(300 * time.Millisecond)
		cancel()
		close(donec)
	}()
	_, err = io.Copy(ioutil.Discard, rc1)
	if err != nil && err != context.Canceled {
		t.Errorf("expected %v, got %v", context.Canceled, err)
	}
	<-donec
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rc2, err := clus.RandClient().Snapshot(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer rc2.Close()
	time.Sleep(700 * time.Millisecond)
	_, err = io.Copy(ioutil.Discard, rc2)
	if err != nil && err != context.DeadlineExceeded {
		t.Errorf("expected %v, got %v", context.DeadlineExceeded, err)
	}
}
