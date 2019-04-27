package fileutil

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLockAndUnlock(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f, err := ioutil.TempFile("", "lock")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	defer func() {
		err = os.Remove(f.Name())
		if err != nil {
			t.Fatal(err)
		}
	}()
	l, err := LockFile(f.Name(), os.O_WRONLY, PrivateFileMode)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = TryLockFile(f.Name(), os.O_WRONLY, PrivateFileMode); err != ErrLocked {
		t.Fatal(err)
	}
	if err = l.Close(); err != nil {
		t.Fatal(err)
	}
	dupl, err := TryLockFile(f.Name(), os.O_WRONLY, PrivateFileMode)
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}
	locked := make(chan struct{}, 1)
	go func() {
		bl, blerr := LockFile(f.Name(), os.O_WRONLY, PrivateFileMode)
		if blerr != nil {
			t.Fatal(blerr)
		}
		locked <- struct{}{}
		if blerr = bl.Close(); blerr != nil {
			t.Fatal(blerr)
		}
	}()
	select {
	case <-locked:
		t.Error("unexpected unblocking")
	case <-time.After(100 * time.Millisecond):
	}
	if err = dupl.Close(); err != nil {
		t.Fatal(err)
	}
	select {
	case <-locked:
	case <-time.After(1 * time.Second):
		t.Error("unexpected blocking")
	}
}
