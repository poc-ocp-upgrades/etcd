package clientv3

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestDialCancel(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	ln, err := net.Listen("unix", "dialcancel:12345")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()
	ep := "unix://dialcancel:12345"
	cfg := Config{Endpoints: []string{ep}, DialTimeout: 30 * time.Second}
	c, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	c.SetEndpoints("http://254.0.0.1:12345")
	getc := make(chan struct{})
	go func() {
		defer close(getc)
		c.Get(c.Ctx(), "abc")
	}()
	time.Sleep(100 * time.Millisecond)
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		c.Close()
	}()
	select {
	case <-time.After(5 * time.Second):
		t.Fatalf("failed to close")
	case <-donec:
	}
	select {
	case <-time.After(5 * time.Second):
		t.Fatalf("get failed to exit")
	case <-getc:
	}
}
func TestDialTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	testCfgs := []Config{{Endpoints: []string{"http://254.0.0.1:12345"}, DialTimeout: 2 * time.Second}, {Endpoints: []string{"http://254.0.0.1:12345"}, DialTimeout: time.Second, Username: "abc", Password: "def"}}
	for i, cfg := range testCfgs {
		donec := make(chan error)
		go func() {
			c, err := New(cfg)
			if c != nil || err == nil {
				t.Errorf("#%d: new client should fail", i)
			}
			donec <- err
		}()
		time.Sleep(10 * time.Millisecond)
		select {
		case err := <-donec:
			t.Errorf("#%d: dial didn't wait (%v)", i, err)
		default:
		}
		select {
		case <-time.After(5 * time.Second):
			t.Errorf("#%d: failed to timeout dial on time", i)
		case err := <-donec:
			if err != context.DeadlineExceeded {
				t.Errorf("#%d: unexpected error %v, want %v", i, err, context.DeadlineExceeded)
			}
		}
	}
}
func TestDialNoTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := Config{Endpoints: []string{"127.0.0.1:12345"}}
	c, err := New(cfg)
	if c == nil || err != nil {
		t.Fatalf("new client with DialNoWait should succeed, got %v", err)
	}
	c.Close()
}
func TestIsHaltErr(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !isHaltErr(nil, fmt.Errorf("etcdserver: some etcdserver error")) {
		t.Errorf(`error prefixed with "etcdserver: " should be Halted by default`)
	}
	if isHaltErr(nil, rpctypes.ErrGRPCStopped) {
		t.Errorf("error %v should not halt", rpctypes.ErrGRPCStopped)
	}
	if isHaltErr(nil, rpctypes.ErrGRPCNoLeader) {
		t.Errorf("error %v should not halt", rpctypes.ErrGRPCNoLeader)
	}
	ctx, cancel := context.WithCancel(context.TODO())
	if isHaltErr(ctx, nil) {
		t.Errorf("no error and active context should not be Halted")
	}
	cancel()
	if !isHaltErr(ctx, nil) {
		t.Errorf("cancel on context should be Halted")
	}
}
