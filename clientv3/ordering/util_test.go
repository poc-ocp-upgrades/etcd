package ordering

import (
	"context"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestEndpointSwitchResolvesViolation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr(), clus.Members[2].GRPCAddr()}
	cfg := clientv3.Config{Endpoints: []string{clus.Members[0].GRPCAddr()}}
	cli, err := clientv3.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.TODO()
	if _, err = clus.Client(0).Put(ctx, "foo", "bar"); err != nil {
		t.Fatal(err)
	}
	if _, err = clus.Client(1).Get(ctx, "foo"); err != nil {
		t.Fatal(err)
	}
	clus.Members[2].InjectPartition(t, clus.Members[:2]...)
	time.Sleep(1 * time.Second)
	if _, err = clus.Client(1).Put(ctx, "foo", "buzz"); err != nil {
		t.Fatal(err)
	}
	cli.SetEndpoints(eps...)
	OrderingKv := NewKV(cli.KV, NewOrderViolationSwitchEndpointClosure(*cli))
	_, err = OrderingKv.Get(ctx, "foo")
	if err != nil {
		t.Fatal(err)
	}
	cli.SetEndpoints(clus.Members[2].GRPCAddr())
	time.Sleep(1 * time.Second)
	_, err = OrderingKv.Get(ctx, "foo", clientv3.WithSerializable())
	if err != nil {
		t.Fatalf("failed to resolve order violation %v", err)
	}
}
func TestUnresolvableOrderViolation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 5, SkipCreatingClient: true})
	defer clus.Terminate(t)
	cfg := clientv3.Config{Endpoints: []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr(), clus.Members[2].GRPCAddr(), clus.Members[3].GRPCAddr(), clus.Members[4].GRPCAddr()}}
	cli, err := clientv3.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	eps := cli.Endpoints()
	ctx := context.TODO()
	cli.SetEndpoints(clus.Members[0].GRPCAddr())
	time.Sleep(1 * time.Second)
	_, err = cli.Put(ctx, "foo", "bar")
	if err != nil {
		t.Fatal(err)
	}
	clus.Members[3].Stop(t)
	time.Sleep(1 * time.Second)
	clus.Members[4].Stop(t)
	time.Sleep(1 * time.Second)
	_, err = cli.Put(ctx, "foo", "buzz")
	if err != nil {
		t.Fatal(err)
	}
	cli.SetEndpoints(eps...)
	OrderingKv := NewKV(cli.KV, NewOrderViolationSwitchEndpointClosure(*cli))
	_, err = OrderingKv.Get(ctx, "foo")
	if err != nil {
		t.Fatal(err)
	}
	clus.Members[0].Stop(t)
	clus.Members[1].Stop(t)
	clus.Members[2].Stop(t)
	clus.Members[3].Restart(t)
	clus.Members[4].Restart(t)
	cli.SetEndpoints(clus.Members[3].GRPCAddr())
	time.Sleep(1 * time.Second)
	_, err = OrderingKv.Get(ctx, "foo", clientv3.WithSerializable())
	if err != ErrNoGreaterRev {
		t.Fatalf("expected %v, got %v", ErrNoGreaterRev, err)
	}
}
