package integration

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestBalancerUnderServerShutdownWatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3, SkipCreatingClient: true})
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr(), clus.Members[2].GRPCAddr()}
	lead := clus.WaitLeader(t)
	watchCli, err := clientv3.New(clientv3.Config{Endpoints: []string{eps[lead]}})
	if err != nil {
		t.Fatal(err)
	}
	defer watchCli.Close()
	mustWaitPinReady(t, watchCli)
	watchCli.SetEndpoints(eps...)
	key, val := "foo", "bar"
	wch := watchCli.Watch(context.Background(), key, clientv3.WithCreatedNotify())
	select {
	case <-wch:
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("took too long to create watch")
	}
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		select {
		case ev := <-wch:
			if werr := ev.Err(); werr != nil {
				t.Fatal(werr)
			}
			if len(ev.Events) != 1 {
				t.Fatalf("expected one event, got %+v", ev)
			}
			if !bytes.Equal(ev.Events[0].Kv.Value, []byte(val)) {
				t.Fatalf("expected %q, got %+v", val, ev.Events[0].Kv)
			}
		case <-time.After(7 * time.Second):
			t.Fatal("took too long to receive events")
		}
	}()
	clus.Members[lead].Terminate(t)
	putCli, err := clientv3.New(clientv3.Config{Endpoints: []string{eps[(lead+1)%3]}})
	if err != nil {
		t.Fatal(err)
	}
	defer putCli.Close()
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		_, err = putCli.Put(ctx, key, val)
		cancel()
		if err == nil {
			break
		}
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) || err == rpctypes.ErrTimeout || err == rpctypes.ErrTimeoutDueToLeaderFail {
			continue
		}
		t.Fatal(err)
	}
	select {
	case <-donec:
	case <-time.After(5 * time.Second):
		t.Fatal("took too long to receive events")
	}
}
func TestBalancerUnderServerShutdownPut(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderServerShutdownMutable(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Put(ctx, "foo", "bar")
		return err
	})
}
func TestBalancerUnderServerShutdownDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderServerShutdownMutable(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Delete(ctx, "foo")
		return err
	})
}
func TestBalancerUnderServerShutdownTxn(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderServerShutdownMutable(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Txn(ctx).If(clientv3.Compare(clientv3.Version("foo"), "=", 0)).Then(clientv3.OpPut("foo", "bar")).Else(clientv3.OpPut("foo", "baz")).Commit()
		return err
	})
}
func testBalancerUnderServerShutdownMutable(t *testing.T, op func(*clientv3.Client, context.Context) error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3, SkipCreatingClient: true})
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr(), clus.Members[2].GRPCAddr()}
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{eps[0]}})
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	mustWaitPinReady(t, cli)
	cli.SetEndpoints(eps...)
	clus.Members[0].Terminate(t)
	time.Sleep(time.Second)
	cctx, ccancel := context.WithTimeout(context.Background(), time.Second)
	err = op(cli, cctx)
	ccancel()
	if err != nil {
		t.Fatal(err)
	}
}
func TestBalancerUnderServerShutdownGetLinearizable(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderServerShutdownImmutable(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Get(ctx, "foo")
		return err
	}, 7*time.Second)
}
func TestBalancerUnderServerShutdownGetSerializable(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderServerShutdownImmutable(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Get(ctx, "foo", clientv3.WithSerializable())
		return err
	}, 2*time.Second)
}
func testBalancerUnderServerShutdownImmutable(t *testing.T, op func(*clientv3.Client, context.Context) error, timeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3, SkipCreatingClient: true})
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr(), clus.Members[2].GRPCAddr()}
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{eps[0]}})
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}
	defer cli.Close()
	mustWaitPinReady(t, cli)
	cli.SetEndpoints(eps...)
	clus.Members[0].Terminate(t)
	cctx, ccancel := context.WithTimeout(context.Background(), timeout)
	err = op(cli, cctx)
	ccancel()
	if err != nil {
		t.Errorf("failed to finish range request in time %v (timeout %v)", err, timeout)
	}
}
func TestBalancerUnderServerStopInflightLinearizableGetOnRestart(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := []pinTestOpt{{pinLeader: true, stopPinFirst: true}, {pinLeader: true, stopPinFirst: false}, {pinLeader: false, stopPinFirst: true}, {pinLeader: false, stopPinFirst: false}}
	for i := range tt {
		testBalancerUnderServerStopInflightRangeOnRestart(t, true, tt[i])
	}
}
func TestBalancerUnderServerStopInflightSerializableGetOnRestart(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := []pinTestOpt{{pinLeader: true, stopPinFirst: true}, {pinLeader: true, stopPinFirst: false}, {pinLeader: false, stopPinFirst: true}, {pinLeader: false, stopPinFirst: false}}
	for i := range tt {
		testBalancerUnderServerStopInflightRangeOnRestart(t, false, tt[i])
	}
}

type pinTestOpt struct {
	pinLeader	bool
	stopPinFirst	bool
}

func testBalancerUnderServerStopInflightRangeOnRestart(t *testing.T, linearizable bool, opt pinTestOpt) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	cfg := &integration.ClusterConfig{Size: 2, SkipCreatingClient: true}
	if linearizable {
		cfg.Size = 3
	}
	clus := integration.NewClusterV3(t, cfg)
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr()}
	if linearizable {
		eps = append(eps, clus.Members[2].GRPCAddr())
	}
	lead := clus.WaitLeader(t)
	target := lead
	if !opt.pinLeader {
		target = (target + 1) % 2
	}
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{eps[target]}})
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}
	defer cli.Close()
	mustWaitPinReady(t, cli)
	cli.SetEndpoints(eps...)
	if opt.stopPinFirst {
		clus.Members[target].Stop(t)
		time.Sleep(time.Second)
		clus.Members[(target+1)%2].Stop(t)
	} else {
		clus.Members[(target+1)%2].Stop(t)
		clus.Members[target].Stop(t)
	}
	clientTimeout := 7 * time.Second
	var gops []clientv3.OpOption
	if !linearizable {
		gops = append(gops, clientv3.WithSerializable())
	}
	donec, readyc := make(chan struct{}), make(chan struct{}, 1)
	go func() {
		defer close(donec)
		ctx, cancel := context.WithTimeout(context.TODO(), clientTimeout)
		readyc <- struct{}{}
		_, err := cli.Get(ctx, "abc", gops...)
		cancel()
		if err != nil {
			t.Fatal(err)
		}
	}()
	<-readyc
	clus.Members[target].Restart(t)
	select {
	case <-time.After(clientTimeout + integration.RequestWaitTimeout):
		t.Fatalf("timed out waiting for Get [linearizable: %v, opt: %+v]", linearizable, opt)
	case <-donec:
	}
}
func isServerCtxTimeout(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return false
	}
	ev, _ := status.FromError(err)
	code := ev.Code()
	return code == codes.DeadlineExceeded && strings.Contains(err.Error(), "context deadline exceeded")
}
