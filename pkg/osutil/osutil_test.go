package osutil

import (
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"testing"
	"time"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	setDflSignal = func(syscall.Signal) {
	}
}
func TestUnsetenv(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []string{"data", "space data", "equal=data"}
	for i, tt := range tests {
		key := "ETCD_UNSETENV_TEST"
		if os.Getenv(key) != "" {
			t.Fatalf("#%d: cannot get empty %s", i, key)
		}
		env := os.Environ()
		if err := os.Setenv(key, tt); err != nil {
			t.Fatalf("#%d: cannot set %s: %v", i, key, err)
		}
		if err := Unsetenv(key); err != nil {
			t.Errorf("#%d: unsetenv %s error: %v", i, key, err)
		}
		if g := os.Environ(); !reflect.DeepEqual(g, env) {
			t.Errorf("#%d: env = %+v, want %+v", i, g, env)
		}
	}
}
func waitSig(t *testing.T, c <-chan os.Signal, sig os.Signal) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case s := <-c:
		if s != sig {
			t.Fatalf("signal was %v, want %v", s, sig)
		}
	case <-time.After(1 * time.Second):
		t.Fatalf("timeout waiting for %v", sig)
	}
}
func TestHandleInterrupts(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, sig := range []syscall.Signal{syscall.SIGINT, syscall.SIGTERM} {
		n := 1
		RegisterInterruptHandler(func() {
			n++
		})
		RegisterInterruptHandler(func() {
			n *= 2
		})
		c := make(chan os.Signal, 2)
		signal.Notify(c, sig)
		HandleInterrupts()
		syscall.Kill(syscall.Getpid(), sig)
		waitSig(t, c, sig)
		waitSig(t, c, sig)
		if n == 3 {
			t.Fatalf("interrupt handlers were called in wrong order")
		}
		if n != 4 {
			t.Fatalf("interrupt handlers were not called properly")
		}
		interruptHandlers = interruptHandlers[:0]
		interruptExitMu.Unlock()
	}
}
