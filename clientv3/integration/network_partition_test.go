package integration

import (
	"context"
	"errors"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
)

var errExpected = errors.New("expected error")

func TestBalancerUnderNetworkPartitionPut(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderNetworkPartition(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Put(ctx, "a", "b")
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) || err == rpctypes.ErrTimeout {
			return errExpected
		}
		return err
	}, time.Second)
}
func TestBalancerUnderNetworkPartitionDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderNetworkPartition(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Delete(ctx, "a")
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) || err == rpctypes.ErrTimeout {
			return errExpected
		}
		return err
	}, time.Second)
}
func TestBalancerUnderNetworkPartitionTxn(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderNetworkPartition(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Txn(ctx).If(clientv3.Compare(clientv3.Version("foo"), "=", 0)).Then(clientv3.OpPut("foo", "bar")).Else(clientv3.OpPut("foo", "baz")).Commit()
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) || err == rpctypes.ErrTimeout {
			return errExpected
		}
		return err
	}, time.Second)
}
func TestBalancerUnderNetworkPartitionLinearizableGetWithLongTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderNetworkPartition(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Get(ctx, "a")
		return err
	}, 7*time.Second)
}
func TestBalancerUnderNetworkPartitionLinearizableGetWithShortTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderNetworkPartition(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Get(ctx, "a")
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) {
			return errExpected
		}
		return err
	}, time.Second)
}
func TestBalancerUnderNetworkPartitionSerializableGet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderNetworkPartition(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Get(ctx, "a", clientv3.WithSerializable())
		return err
	}, time.Second)
}
func testBalancerUnderNetworkPartition(t *testing.T, op func(*clientv3.Client, context.Context) error, timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3, SkipCreatingClient: true})
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr(), clus.Members[2].GRPCAddr()}
	ccfg := clientv3.Config{Endpoints: []string{eps[0]}, DialTimeout: 3 * time.Second}
	cli, err := clientv3.New(ccfg)
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	mustWaitPinReady(t, cli)
	cli.SetEndpoints(eps...)
	clus.Members[0].InjectPartition(t, clus.Members[1:]...)
	for i := 0; i < 2; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		err = op(cli, ctx)
		cancel()
		if err == nil {
			break
		}
		if err != errExpected {
			t.Errorf("#%d: expected %v, got %v", i, errExpected, err)
		}
		if i == 0 {
			time.Sleep(5 * time.Second)
		}
	}
	if err != nil {
		t.Errorf("balancer did not switch in time (%v)", err)
	}
}
func TestBalancerUnderNetworkPartitionLinearizableGetLeaderElection(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3, SkipCreatingClient: true})
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr(), clus.Members[2].GRPCAddr()}
	lead := clus.WaitLeader(t)
	timeout := 3 * clus.Members[(lead+1)%2].ServerConfig.ReqTimeout()
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{eps[(lead+1)%2]}, DialTimeout: 1 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	mustWaitPinReady(t, cli)
	cli.SetEndpoints(eps[lead], eps[(lead+1)%2])
	clus.Members[lead].InjectPartition(t, clus.Members[(lead+1)%3], clus.Members[(lead+2)%3])
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	_, err = cli.Get(ctx, "a")
	cancel()
	if err != nil {
		t.Fatal(err)
	}
}
func TestBalancerUnderNetworkPartitionWatchLeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderNetworkPartitionWatch(t, true)
}
func TestBalancerUnderNetworkPartitionWatchFollower(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderNetworkPartitionWatch(t, false)
}
func testBalancerUnderNetworkPartitionWatch(t *testing.T, isolateLeader bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3, SkipCreatingClient: true})
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr(), clus.Members[2].GRPCAddr()}
	target := clus.WaitLeader(t)
	if !isolateLeader {
		target = (target + 1) % 3
	}
	watchCli, err := clientv3.New(clientv3.Config{Endpoints: []string{eps[target]}})
	if err != nil {
		t.Fatal(err)
	}
	defer watchCli.Close()
	mustWaitPinReady(t, watchCli)
	watchCli.SetEndpoints(eps...)
	wch := watchCli.Watch(clientv3.WithRequireLeader(context.Background()), "foo", clientv3.WithCreatedNotify())
	select {
	case <-wch:
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("took too long to create watch")
	}
	clus.Members[target].InjectPartition(t, clus.Members[(target+1)%3], clus.Members[(target+2)%3])
	select {
	case ev := <-wch:
		if len(ev.Events) != 0 {
			t.Fatal("expected no event")
		}
		if err = ev.Err(); err != rpctypes.ErrNoLeader {
			t.Fatalf("expected %v, got %v", rpctypes.ErrNoLeader, err)
		}
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("took too long to detect leader lost")
	}
}
