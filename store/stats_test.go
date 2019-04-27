package store

import (
	"testing"
	"time"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestStoreStatsGetSuccess(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Create("/foo", false, "bar", false, TTLOptionSet{ExpireTime: Permanent})
	s.Get("/foo", false, false)
	testutil.AssertEqual(t, uint64(1), s.Stats.GetSuccess, "")
}
func TestStoreStatsGetFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Create("/foo", false, "bar", false, TTLOptionSet{ExpireTime: Permanent})
	s.Get("/no_such_key", false, false)
	testutil.AssertEqual(t, uint64(1), s.Stats.GetFail, "")
}
func TestStoreStatsCreateSuccess(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Create("/foo", false, "bar", false, TTLOptionSet{ExpireTime: Permanent})
	testutil.AssertEqual(t, uint64(1), s.Stats.CreateSuccess, "")
}
func TestStoreStatsCreateFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Create("/foo", true, "", false, TTLOptionSet{ExpireTime: Permanent})
	s.Create("/foo", false, "bar", false, TTLOptionSet{ExpireTime: Permanent})
	testutil.AssertEqual(t, uint64(1), s.Stats.CreateFail, "")
}
func TestStoreStatsUpdateSuccess(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Create("/foo", false, "bar", false, TTLOptionSet{ExpireTime: Permanent})
	s.Update("/foo", "baz", TTLOptionSet{ExpireTime: Permanent})
	testutil.AssertEqual(t, uint64(1), s.Stats.UpdateSuccess, "")
}
func TestStoreStatsUpdateFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Update("/foo", "bar", TTLOptionSet{ExpireTime: Permanent})
	testutil.AssertEqual(t, uint64(1), s.Stats.UpdateFail, "")
}
func TestStoreStatsCompareAndSwapSuccess(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Create("/foo", false, "bar", false, TTLOptionSet{ExpireTime: Permanent})
	s.CompareAndSwap("/foo", "bar", 0, "baz", TTLOptionSet{ExpireTime: Permanent})
	testutil.AssertEqual(t, uint64(1), s.Stats.CompareAndSwapSuccess, "")
}
func TestStoreStatsCompareAndSwapFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Create("/foo", false, "bar", false, TTLOptionSet{ExpireTime: Permanent})
	s.CompareAndSwap("/foo", "wrong_value", 0, "baz", TTLOptionSet{ExpireTime: Permanent})
	testutil.AssertEqual(t, uint64(1), s.Stats.CompareAndSwapFail, "")
}
func TestStoreStatsDeleteSuccess(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Create("/foo", false, "bar", false, TTLOptionSet{ExpireTime: Permanent})
	s.Delete("/foo", false, false)
	testutil.AssertEqual(t, uint64(1), s.Stats.DeleteSuccess, "")
}
func TestStoreStatsDeleteFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	s.Delete("/foo", false, false)
	testutil.AssertEqual(t, uint64(1), s.Stats.DeleteFail, "")
}
func TestStoreStatsExpireCount(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore()
	fc := newFakeClock()
	s.clock = fc
	s.Create("/foo", false, "bar", false, TTLOptionSet{ExpireTime: fc.Now().Add(500 * time.Millisecond)})
	testutil.AssertEqual(t, uint64(0), s.Stats.ExpireCount, "")
	fc.Advance(600 * time.Millisecond)
	s.DeleteExpiredKeys(fc.Now())
	testutil.AssertEqual(t, uint64(1), s.Stats.ExpireCount, "")
}
