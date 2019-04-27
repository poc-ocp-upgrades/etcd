package store_test

import (
	"testing"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/store"
)

type v2TestStore struct{ store.Store }

func (s *v2TestStore) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func newTestStore(t *testing.T, ns ...string) StoreCloser {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &v2TestStore{store.New(ns...)}
}
func TestStoreRecover(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newTestStore(t)
	defer s.Close()
	var eidx uint64 = 4
	s.Create("/foo", true, "", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	s.Create("/foo/x", false, "bar", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	s.Update("/foo/x", "barbar", store.TTLOptionSet{ExpireTime: store.Permanent})
	s.Create("/foo/y", false, "baz", false, store.TTLOptionSet{ExpireTime: store.Permanent})
	b, err := s.Save()
	testutil.AssertNil(t, err)
	s2 := newTestStore(t)
	s2.Recovery(b)
	e, err := s.Get("/foo/x", false, false)
	testutil.AssertEqual(t, e.Node.CreatedIndex, uint64(2))
	testutil.AssertEqual(t, e.Node.ModifiedIndex, uint64(3))
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, *e.Node.Value, "barbar")
	e, err = s.Get("/foo/y", false, false)
	testutil.AssertEqual(t, e.EtcdIndex, eidx)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, *e.Node.Value, "baz")
}
