package store_test

import (
	"testing"
	"time"
	etcdErr "github.com/coreos/etcd/error"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/store"
)

type StoreCloser interface {
	store.Store
	Close()
}

func TestNewStoreWithNamespaces(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t, "/0", "/1")
	defer s.Close()
	_, err := s.Get("/0", false, false)
	testutil.AssertNil(t, err)
	_, err = s.Get("/1", false, false)
	testutil.AssertNil(t, err)
}
func TestStoreGetValue(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	var eidx uint64 = 1
	e, err := s.Get("/foo", false, false)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "get")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertEqual(t, *e.Node.Value, "bar")
}
func TestStoreGetSorted(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/foo", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	s.Create("/foo/x", false, "0", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	s.Create("/foo/z", false, "0", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	s.Create("/foo/y", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	s.Create("/foo/y/a", false, "0", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	s.Create("/foo/y/b", false, "0", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	var eidx uint64 = 6
	e, err := s.Get("/foo", true, true)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	var yNodes store.NodeExterns
	sortedStrings := []string{"/foo/x", "/foo/y", "/foo/z"}
	for i := range e.Node.Nodes {
		node := e.Node.Nodes[i]
		if node.Key != sortedStrings[i] {
			t.Errorf("expect key = %s, got key = %s", sortedStrings[i], node.Key)
		}
		if node.Key == "/foo/y" {
			yNodes = node.Nodes
		}
	}
	sortedStrings = []string{"/foo/y/a", "/foo/y/b"}
	for i := range yNodes {
		node := yNodes[i]
		if node.Key != sortedStrings[i] {
			t.Errorf("expect key = %s, got key = %s", sortedStrings[i], node.Key)
		}
	}
}
func TestSet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	e, err := s.Set("/foo", false, "", store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "set")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertFalse(t, e.Node.Dir)
	testutil.AssertEqual(t, *e.Node.Value, "")
	testutil.AssertNil(t, e.Node.Nodes)
	testutil.AssertNil(t, e.Node.Expiration)
	testutil.AssertEqual(t, e.Node.TTL, int64(0))
	testutil.AssertEqual(t, e.Node.ModifiedIndex, uint64(1))
	eidx = 2
	e, err = s.Set("/foo", false, "bar", store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "set")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertFalse(t, e.Node.Dir)
	testutil.AssertEqual(t, *e.Node.Value, "bar")
	testutil.AssertNil(t, e.Node.Nodes)
	testutil.AssertNil(t, e.Node.Expiration)
	testutil.AssertEqual(t, e.Node.TTL, int64(0))
	testutil.AssertEqual(t, e.Node.ModifiedIndex, uint64(2))
	testutil.AssertNotNil(t, e.PrevNode)
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, *e.PrevNode.Value, "")
	testutil.AssertEqual(t, e.PrevNode.ModifiedIndex, uint64(1))
	eidx = 3
	e, err = s.Set("/foo", false, "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "set")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertFalse(t, e.Node.Dir)
	testutil.AssertEqual(t, *e.Node.Value, "baz")
	testutil.AssertNil(t, e.Node.Nodes)
	testutil.AssertNil(t, e.Node.Expiration)
	testutil.AssertEqual(t, e.Node.TTL, int64(0))
	testutil.AssertEqual(t, e.Node.ModifiedIndex, uint64(3))
	testutil.AssertNotNil(t, e.PrevNode)
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, *e.PrevNode.Value, "bar")
	testutil.AssertEqual(t, e.PrevNode.ModifiedIndex, uint64(2))
	eidx = 4
	e, err = s.Set("/dir", true, "", store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "set")
	testutil.AssertEqual(t, e.Node.Key, "/dir")
	testutil.AssertTrue(t, e.Node.Dir)
	testutil.AssertNil(t, e.Node.Value)
	testutil.AssertNil(t, e.Node.Nodes)
	testutil.AssertNil(t, e.Node.Expiration)
	testutil.AssertEqual(t, e.Node.TTL, int64(0))
	testutil.AssertEqual(t, e.Node.ModifiedIndex, uint64(4))
}
func TestStoreCreateValue(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	e, err := s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "create")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertFalse(t, e.Node.Dir)
	testutil.AssertEqual(t, *e.Node.Value, "bar")
	testutil.AssertNil(t, e.Node.Nodes)
	testutil.AssertNil(t, e.Node.Expiration)
	testutil.AssertEqual(t, e.Node.TTL, int64(0))
	testutil.AssertEqual(t, e.Node.ModifiedIndex, uint64(1))
	eidx = 2
	e, err = s.Create("/empty", false, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "create")
	testutil.AssertEqual(t, e.Node.Key, "/empty")
	testutil.AssertFalse(t, e.Node.Dir)
	testutil.AssertEqual(t, *e.Node.Value, "")
	testutil.AssertNil(t, e.Node.Nodes)
	testutil.AssertNil(t, e.Node.Expiration)
	testutil.AssertEqual(t, e.Node.TTL, int64(0))
	testutil.AssertEqual(t, e.Node.ModifiedIndex, uint64(2))
}
func TestStoreCreateDirectory(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	e, err := s.Create("/foo", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "create")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertTrue(t, e.Node.Dir)
}
func TestStoreCreateFailsIfExists(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/foo", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, _err := s.Create("/foo", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	err := _err.(*etcdErr.Error)
	testutil.AssertEqual(t, err.ErrorCode, etcdErr.EcodeNodeExist)
	testutil.AssertEqual(t, err.Message, "Key already exists")
	testutil.AssertEqual(t, err.Cause, "/foo")
	testutil.AssertEqual(t, err.Index, uint64(1))
	testutil.AssertNil(t, e)
}
func TestStoreUpdateValue(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	var eidx uint64 = 2
	e, err := s.Update("/foo", "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "update")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertFalse(t, e.Node.Dir)
	testutil.AssertEqual(t, *e.Node.Value, "baz")
	testutil.AssertEqual(t, e.Node.TTL, int64(0))
	testutil.AssertEqual(t, e.Node.ModifiedIndex, uint64(2))
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, *e.PrevNode.Value, "bar")
	testutil.AssertEqual(t, e.PrevNode.TTL, int64(0))
	testutil.AssertEqual(t, e.PrevNode.ModifiedIndex, uint64(1))
	e, _ = s.Get("/foo", false, false)
	testutil.AssertEqual(t, *e.Node.Value, "baz")
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	eidx = 3
	e, err = s.Update("/foo", "", store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "update")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertFalse(t, e.Node.Dir)
	testutil.AssertEqual(t, *e.Node.Value, "")
	testutil.AssertEqual(t, e.Node.TTL, int64(0))
	testutil.AssertEqual(t, e.Node.ModifiedIndex, uint64(3))
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, *e.PrevNode.Value, "baz")
	testutil.AssertEqual(t, e.PrevNode.TTL, int64(0))
	testutil.AssertEqual(t, e.PrevNode.ModifiedIndex, uint64(2))
	e, _ = s.Get("/foo", false, false)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, *e.Node.Value, "")
}
func TestStoreUpdateFailsIfDirectory(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/foo", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, _err := s.Update("/foo", "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	err := _err.(*etcdErr.Error)
	testutil.AssertEqual(t, err.ErrorCode, etcdErr.EcodeNotFile)
	testutil.AssertEqual(t, err.Message, "Not a file")
	testutil.AssertEqual(t, err.Cause, "/foo")
	testutil.AssertNil(t, e)
}
func TestStoreDeleteValue(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 2
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, err := s.Delete("/foo", false, false)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "delete")
	testutil.AssertNotNil(t, e.PrevNode)
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, *e.PrevNode.Value, "bar")
}
func TestStoreDeleteDirectory(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 2
	s.Create("/foo", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, err := s.Delete("/foo", true, false)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "delete")
	testutil.AssertNotNil(t, e.PrevNode)
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, e.PrevNode.Dir, true)
	_, err = s.Create("/foo/bar", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	_, err = s.Delete("/foo", true, false)
	testutil.AssertNotNil(t, err)
	e, err = s.Delete("/foo", false, true)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.Action, "delete")
}
func TestStoreDeleteDirectoryFailsIfNonRecursiveAndDir(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/foo", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, _err := s.Delete("/foo", false, false)
	err := _err.(*etcdErr.Error)
	testutil.AssertEqual(t, err.ErrorCode, etcdErr.EcodeNotFile)
	testutil.AssertEqual(t, err.Message, "Not a file")
	testutil.AssertNil(t, e)
}
func TestRootRdOnly(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t, "/0")
	defer s.Close()
	for _, tt := range []string{"/", "/0"} {
		_, err := s.Set(tt, true, "", store.TTLOptionSet{ExpireTime: store.Permanent})
		testutil.AssertNotNil(t, err)
		_, err = s.Delete(tt, true, true)
		testutil.AssertNotNil(t, err)
		_, err = s.Create(tt, true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
		testutil.AssertNotNil(t, err)
		_, err = s.Update(tt, "", store.TTLOptionSet{ExpireTime: store.Permanent})
		testutil.AssertNotNil(t, err)
		_, err = s.CompareAndSwap(tt, "", 0, "", store.TTLOptionSet{ExpireTime: store.Permanent})
		testutil.AssertNotNil(t, err)
	}
}
func TestStoreCompareAndDeletePrevValue(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 2
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, err := s.CompareAndDelete("/foo", "bar", 0)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "compareAndDelete")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertNotNil(t, e.PrevNode)
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, *e.PrevNode.Value, "bar")
	testutil.AssertEqual(t, e.PrevNode.ModifiedIndex, uint64(1))
	testutil.AssertEqual(t, e.PrevNode.CreatedIndex, uint64(1))
}
func TestStoreCompareAndDeletePrevValueFailsIfNotMatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, _err := s.CompareAndDelete("/foo", "baz", 0)
	err := _err.(*etcdErr.Error)
	testutil.AssertEqual(t, err.ErrorCode, etcdErr.EcodeTestFailed)
	testutil.AssertEqual(t, err.Message, "Compare failed")
	testutil.AssertNil(t, e)
	e, _ = s.Get("/foo", false, false)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, *e.Node.Value, "bar")
}
func TestStoreCompareAndDeletePrevIndex(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 2
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, err := s.CompareAndDelete("/foo", "", 1)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "compareAndDelete")
	testutil.AssertNotNil(t, e.PrevNode)
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, *e.PrevNode.Value, "bar")
	testutil.AssertEqual(t, e.PrevNode.ModifiedIndex, uint64(1))
	testutil.AssertEqual(t, e.PrevNode.CreatedIndex, uint64(1))
}
func TestStoreCompareAndDeletePrevIndexFailsIfNotMatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, _err := s.CompareAndDelete("/foo", "", 100)
	testutil.AssertNotNil(t, _err)
	err := _err.(*etcdErr.Error)
	testutil.AssertEqual(t, err.ErrorCode, etcdErr.EcodeTestFailed)
	testutil.AssertEqual(t, err.Message, "Compare failed")
	testutil.AssertNil(t, e)
	e, _ = s.Get("/foo", false, false)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, *e.Node.Value, "bar")
}
func TestStoreCompareAndDeleteDirectoryFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/foo", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	_, _err := s.CompareAndDelete("/foo", "", 0)
	testutil.AssertNotNil(t, _err)
	err := _err.(*etcdErr.Error)
	testutil.AssertEqual(t, err.ErrorCode, etcdErr.EcodeNotFile)
}
func TestStoreCompareAndSwapPrevValue(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 2
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, err := s.CompareAndSwap("/foo", "bar", 0, "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "compareAndSwap")
	testutil.AssertEqual(t, *e.Node.Value, "baz")
	testutil.AssertNotNil(t, e.PrevNode)
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, *e.PrevNode.Value, "bar")
	testutil.AssertEqual(t, e.PrevNode.ModifiedIndex, uint64(1))
	testutil.AssertEqual(t, e.PrevNode.CreatedIndex, uint64(1))
	e, _ = s.Get("/foo", false, false)
	testutil.AssertEqual(t, *e.Node.Value, "baz")
}
func TestStoreCompareAndSwapPrevValueFailsIfNotMatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, _err := s.CompareAndSwap("/foo", "wrong_value", 0, "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	err := _err.(*etcdErr.Error)
	testutil.AssertEqual(t, err.ErrorCode, etcdErr.EcodeTestFailed)
	testutil.AssertEqual(t, err.Message, "Compare failed")
	testutil.AssertNil(t, e)
	e, _ = s.Get("/foo", false, false)
	testutil.AssertEqual(t, *e.Node.Value, "bar")
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
}
func TestStoreCompareAndSwapPrevIndex(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 2
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, err := s.CompareAndSwap("/foo", "", 1, "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "compareAndSwap")
	testutil.AssertEqual(t, *e.Node.Value, "baz")
	testutil.AssertNotNil(t, e.PrevNode)
	testutil.AssertEqual(t, e.PrevNode.Key, "/foo")
	testutil.AssertEqual(t, *e.PrevNode.Value, "bar")
	testutil.AssertEqual(t, e.PrevNode.ModifiedIndex, uint64(1))
	testutil.AssertEqual(t, e.PrevNode.CreatedIndex, uint64(1))
	e, _ = s.Get("/foo", false, false)
	testutil.AssertEqual(t, *e.Node.Value, "baz")
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
}
func TestStoreCompareAndSwapPrevIndexFailsIfNotMatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e, _err := s.CompareAndSwap("/foo", "", 100, "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	err := _err.(*etcdErr.Error)
	testutil.AssertEqual(t, err.ErrorCode, etcdErr.EcodeTestFailed)
	testutil.AssertEqual(t, err.Message, "Compare failed")
	testutil.AssertNil(t, e)
	e, _ = s.Get("/foo", false, false)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, *e.Node.Value, "bar")
}
func TestStoreWatchCreate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 0
	w, _ := s.Watch("/foo", false, false, 0)
	c := w.EventChan()
	testutil.AssertEqual(t, w.StartIndex(), eidx)
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	eidx = 1
	e := timeoutSelect(t, c)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "create")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	select {
	case e = <-w.EventChan():
		testutil.AssertNil(t, e)
	case <-time.After(100 * time.Millisecond):
	}
}
func TestStoreWatchRecursiveCreate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 0
	w, err := s.Watch("/foo", true, false, 0)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, w.StartIndex(), eidx)
	eidx = 1
	s.Create("/foo/bar", false, "baz", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "create")
	testutil.AssertEqual(t, e.Node.Key, "/foo/bar")
}
func TestStoreWatchUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, _ := s.Watch("/foo", false, false, 0)
	testutil.AssertEqual(t, w.StartIndex(), eidx)
	eidx = 2
	s.Update("/foo", "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "update")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
}
func TestStoreWatchRecursiveUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo/bar", false, "baz", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, err := s.Watch("/foo", true, false, 0)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, w.StartIndex(), eidx)
	eidx = 2
	s.Update("/foo/bar", "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "update")
	testutil.AssertEqual(t, e.Node.Key, "/foo/bar")
}
func TestStoreWatchDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, _ := s.Watch("/foo", false, false, 0)
	testutil.AssertEqual(t, w.StartIndex(), eidx)
	eidx = 2
	s.Delete("/foo", false, false)
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "delete")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
}
func TestStoreWatchRecursiveDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo/bar", false, "baz", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, err := s.Watch("/foo", true, false, 0)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, w.StartIndex(), eidx)
	eidx = 2
	s.Delete("/foo/bar", false, false)
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "delete")
	testutil.AssertEqual(t, e.Node.Key, "/foo/bar")
}
func TestStoreWatchCompareAndSwap(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, _ := s.Watch("/foo", false, false, 0)
	testutil.AssertEqual(t, w.StartIndex(), eidx)
	eidx = 2
	s.CompareAndSwap("/foo", "bar", 0, "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "compareAndSwap")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
}
func TestStoreWatchRecursiveCompareAndSwap(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	s.Create("/foo/bar", false, "baz", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, _ := s.Watch("/foo", true, false, 0)
	testutil.AssertEqual(t, w.StartIndex(), eidx)
	eidx = 2
	s.CompareAndSwap("/foo/bar", "baz", 0, "bat", store.TTLOptionSet{ExpireTime: store.Permanent})
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "compareAndSwap")
	testutil.AssertEqual(t, e.Node.Key, "/foo/bar")
}
func TestStoreWatchStream(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	w, _ := s.Watch("/foo", false, true, 0)
	s.Create("/foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "create")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertEqual(t, *e.Node.Value, "bar")
	select {
	case e = <-w.EventChan():
		testutil.AssertNil(t, e)
	case <-time.After(100 * time.Millisecond):
	}
	eidx = 2
	s.Update("/foo", "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	e = timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "update")
	testutil.AssertEqual(t, e.Node.Key, "/foo")
	testutil.AssertEqual(t, *e.Node.Value, "baz")
	select {
	case e = <-w.EventChan():
		testutil.AssertNil(t, e)
	case <-time.After(100 * time.Millisecond):
	}
}
func TestStoreWatchCreateWithHiddenKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	w, _ := s.Watch("/_foo", false, false, 0)
	s.Create("/_foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "create")
	testutil.AssertEqual(t, e.Node.Key, "/_foo")
	select {
	case e = <-w.EventChan():
		testutil.AssertNil(t, e)
	case <-time.After(100 * time.Millisecond):
	}
}
func TestStoreWatchRecursiveCreateWithHiddenKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	w, _ := s.Watch("/foo", true, false, 0)
	s.Create("/foo/_bar", false, "baz", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e := nbselect(w.EventChan())
	testutil.AssertNil(t, e)
	w, _ = s.Watch("/foo", true, false, 0)
	s.Create("/foo/_baz", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	select {
	case e = <-w.EventChan():
		testutil.AssertNil(t, e)
	case <-time.After(100 * time.Millisecond):
	}
	s.Create("/foo/_baz/quux", false, "quux", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	select {
	case e = <-w.EventChan():
		testutil.AssertNil(t, e)
	case <-time.After(100 * time.Millisecond):
	}
}
func TestStoreWatchUpdateWithHiddenKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/_foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, _ := s.Watch("/_foo", false, false, 0)
	s.Update("/_foo", "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.Action, "update")
	testutil.AssertEqual(t, e.Node.Key, "/_foo")
	e = nbselect(w.EventChan())
	testutil.AssertNil(t, e)
}
func TestStoreWatchRecursiveUpdateWithHiddenKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/foo/_bar", false, "baz", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, _ := s.Watch("/foo", true, false, 0)
	s.Update("/foo/_bar", "baz", store.TTLOptionSet{ExpireTime: store.Permanent})
	e := nbselect(w.EventChan())
	testutil.AssertNil(t, e)
}
func TestStoreWatchDeleteWithHiddenKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 2
	s.Create("/_foo", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, _ := s.Watch("/_foo", false, false, 0)
	s.Delete("/_foo", false, false)
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "delete")
	testutil.AssertEqual(t, e.Node.Key, "/_foo")
	e = nbselect(w.EventChan())
	testutil.AssertNil(t, e)
}
func TestStoreWatchRecursiveDeleteWithHiddenKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Create("/foo/_bar", false, "baz", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	w, _ := s.Watch("/foo", true, false, 0)
	s.Delete("/foo/_bar", false, false)
	e := nbselect(w.EventChan())
	testutil.AssertNil(t, e)
}
func TestStoreWatchRecursiveCreateDeeperThanHiddenKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 1
	w, _ := s.Watch("/_foo/bar", true, false, 0)
	s.Create("/_foo/bar/baz", false, "baz", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	e := timeoutSelect(t, w.EventChan())
	testutil.AssertNotNil(t, e)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertEqual(t, e.Action, "create")
	testutil.AssertEqual(t, e.Node.Key, "/_foo/bar/baz")
}
func TestStoreWatchSlowConsumer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	s.Watch("/foo", true, true, 0)
	for i := 1; i <= 100; i++ {
		s.Set("/foo", false, string(i), store.TTLOptionSet{ExpireTime: store.Permanent})
	}
	s.Set("/foo", false, "101", store.TTLOptionSet{ExpireTime: store.Permanent})
	s.Set("/foo", false, "102", store.TTLOptionSet{ExpireTime: store.Permanent})
}
func nbselect(c <-chan *store.Event) *store.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case e := <-c:
		return e
	default:
		return nil
	}
}
func timeoutSelect(t *testing.T, c <-chan *store.Event) *store.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case e := <-c:
		return e
	case <-time.After(time.Second):
		t.Errorf("timed out waiting on event")
		return nil
	}
}
