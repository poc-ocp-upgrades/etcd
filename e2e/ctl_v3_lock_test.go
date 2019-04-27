package e2e

import (
	"os"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/pkg/expect"
)

func TestCtlV3Lock(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldenv := os.Getenv("EXPECT_DEBUG")
	defer os.Setenv("EXPECT_DEBUG", oldenv)
	os.Setenv("EXPECT_DEBUG", "1")
	testCtl(t, testLock)
}
func testLock(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	name := "a"
	holder, ch, err := ctlV3Lock(cx, name)
	if err != nil {
		cx.t.Fatal(err)
	}
	l1 := ""
	select {
	case <-time.After(2 * time.Second):
		cx.t.Fatalf("timed out locking")
	case l1 = <-ch:
		if !strings.HasPrefix(l1, name) {
			cx.t.Errorf("got %q, expected %q prefix", l1, name)
		}
	}
	blocked, ch, err := ctlV3Lock(cx, name)
	if err != nil {
		cx.t.Fatal(err)
	}
	select {
	case <-time.After(100 * time.Millisecond):
	case <-ch:
		cx.t.Fatalf("should block")
	}
	blockAcquire, ch, err := ctlV3Lock(cx, name)
	if err != nil {
		cx.t.Fatal(err)
	}
	defer blockAcquire.Stop()
	select {
	case <-time.After(100 * time.Millisecond):
	case <-ch:
		cx.t.Fatalf("should block")
	}
	if err = blocked.Signal(os.Interrupt); err != nil {
		cx.t.Fatal(err)
	}
	if err = closeWithTimeout(blocked, time.Second); err != nil {
		cx.t.Fatal(err)
	}
	if err = holder.Signal(os.Interrupt); err != nil {
		cx.t.Fatal(err)
	}
	if err = closeWithTimeout(holder, time.Second); err != nil {
		cx.t.Fatal(err)
	}
	select {
	case <-time.After(time.Second):
		cx.t.Fatalf("timed out from waiting to holding")
	case l2 := <-ch:
		if l1 == l2 || !strings.HasPrefix(l2, name) {
			cx.t.Fatalf("expected different lock name, got l1=%q, l2=%q", l1, l2)
		}
	}
}
func ctlV3Lock(cx ctlCtx, name string) (*expect.ExpectProcess, <-chan string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "lock", name)
	proc, err := spawnCmd(cmdArgs)
	outc := make(chan string, 1)
	if err != nil {
		close(outc)
		return proc, outc, err
	}
	go func() {
		s, xerr := proc.ExpectFunc(func(string) bool {
			return true
		})
		if xerr != nil {
			cx.t.Errorf("expect failed (%v)", xerr)
		}
		outc <- s
	}()
	return proc, outc, err
}
