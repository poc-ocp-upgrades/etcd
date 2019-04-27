package store

import (
	"testing"
	etcdErr "github.com/coreos/etcd/error"
)

func TestEventQueue(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eh := newEventHistory(100)
	for i := 0; i < 200; i++ {
		e := newEvent(Create, "/foo", uint64(i), uint64(i))
		eh.addEvent(e)
	}
	j := 100
	i := eh.Queue.Front
	n := eh.Queue.Size
	for ; n > 0; n-- {
		e := eh.Queue.Events[i]
		if e.Index() != uint64(j) {
			t.Fatalf("queue error!")
		}
		j++
		i = (i + 1) % eh.Queue.Capacity
	}
}
func TestScanHistory(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eh := newEventHistory(100)
	eh.addEvent(newEvent(Create, "/foo", 1, 1))
	eh.addEvent(newEvent(Create, "/foo/bar", 2, 2))
	eh.addEvent(newEvent(Create, "/foo/foo", 3, 3))
	eh.addEvent(newEvent(Create, "/foo/bar/bar", 4, 4))
	eh.addEvent(newEvent(Create, "/foo/foo/foo", 5, 5))
	de := newEvent(Delete, "/foo", 6, 6)
	de.PrevNode = newDir(nil, "/foo", 1, nil, Permanent).Repr(false, false, nil)
	eh.addEvent(de)
	e, err := eh.scan("/foo", false, 1)
	if err != nil || e.Index() != 1 {
		t.Fatalf("scan error [/foo] [1] %d (%v)", e.Index(), err)
	}
	e, err = eh.scan("/foo/bar", false, 1)
	if err != nil || e.Index() != 2 {
		t.Fatalf("scan error [/foo/bar] [2] %d (%v)", e.Index(), err)
	}
	e, err = eh.scan("/foo/bar", true, 3)
	if err != nil || e.Index() != 4 {
		t.Fatalf("scan error [/foo/bar/bar] [4] %d (%v)", e.Index(), err)
	}
	e, err = eh.scan("/foo/foo/foo", false, 6)
	if err != nil || e.Index() != 6 {
		t.Fatalf("scan error [/foo/foo/foo] [6] %d (%v)", e.Index(), err)
	}
	e, _ = eh.scan("/foo/bar", true, 7)
	if e != nil {
		t.Fatalf("bad index shoud reuturn nil")
	}
}
func TestEventIndexHistoryCleared(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eh := newEventHistory(5)
	eh.addEvent(newEvent(Create, "/foo", 1, 1))
	eh.addEvent(newEvent(Create, "/foo/bar", 2, 2))
	eh.addEvent(newEvent(Create, "/foo/foo", 3, 3))
	eh.addEvent(newEvent(Create, "/foo/bar/bar", 4, 4))
	eh.addEvent(newEvent(Create, "/foo/foo/foo", 5, 5))
	eh.addEvent(newEvent(Create, "/foo/bar/bar/bar", 6, 6))
	_, err := eh.scan("/foo", false, 1)
	if err == nil || err.ErrorCode != etcdErr.EcodeEventIndexCleared {
		t.Fatalf("scan error cleared index should return err with %d got (%v)", etcdErr.EcodeEventIndexCleared, err)
	}
}
func TestFullEventQueue(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eh := newEventHistory(10)
	for i := 0; i < 1000; i++ {
		ce := newEvent(Create, "/foo", uint64(i), uint64(i))
		eh.addEvent(ce)
		e, err := eh.scan("/foo", true, uint64(i-1))
		if i > 0 {
			if e == nil || err != nil {
				t.Fatalf("scan error [/foo] [%v] %v", i-1, i)
			}
		}
	}
}
func TestCloneEvent(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e1 := &Event{Action: Create, EtcdIndex: 1, Node: nil, PrevNode: nil}
	e2 := e1.Clone()
	if e2.Action != Create {
		t.Fatalf("Action=%q, want %q", e2.Action, Create)
	}
	if e2.EtcdIndex != e1.EtcdIndex {
		t.Fatalf("EtcdIndex=%d, want %d", e2.EtcdIndex, e1.EtcdIndex)
	}
	e2.Action = Delete
	e2.EtcdIndex = uint64(5)
	if e1.Action != Create {
		t.Fatalf("Action=%q, want %q", e1.Action, Create)
	}
	if e1.EtcdIndex != uint64(1) {
		t.Fatalf("EtcdIndex=%d, want %d", e1.EtcdIndex, uint64(1))
	}
	if e2.Action != Delete {
		t.Fatalf("Action=%q, want %q", e2.Action, Delete)
	}
	if e2.EtcdIndex != uint64(5) {
		t.Fatalf("EtcdIndex=%d, want %d", e2.EtcdIndex, uint64(5))
	}
}
