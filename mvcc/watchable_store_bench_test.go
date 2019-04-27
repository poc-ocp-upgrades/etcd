package mvcc

import (
	"math/rand"
	"os"
	"testing"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc/backend"
)

func BenchmarkWatchableStorePut(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := New(be, &lease.FakeLessor{}, nil)
	defer cleanup(s, be, tmpPath)
	bytesN := 64
	keys := createBytesSlice(bytesN, b.N)
	vals := createBytesSlice(bytesN, b.N)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Put(keys[i], vals[i], lease.NoLease)
	}
}
func BenchmarkWatchableStoreTxnPut(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i fakeConsistentIndex
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := New(be, &lease.FakeLessor{}, &i)
	defer cleanup(s, be, tmpPath)
	bytesN := 64
	keys := createBytesSlice(bytesN, b.N)
	vals := createBytesSlice(bytesN, b.N)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		txn := s.Write()
		txn.Put(keys[i], vals[i], lease.NoLease)
		txn.End()
	}
}
func BenchmarkWatchableStoreWatchSyncPut(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(be, &lease.FakeLessor{}, nil)
	defer cleanup(s, be, tmpPath)
	k := []byte("testkey")
	v := []byte("testval")
	w := s.NewWatchStream()
	defer w.Close()
	watchIDs := make([]WatchID, b.N)
	for i := range watchIDs {
		watchIDs[i] = w.Watch(k, nil, 1)
	}
	b.ResetTimer()
	b.ReportAllocs()
	s.Put(k, v, lease.NoLease)
	for range watchIDs {
		<-w.Chan()
	}
	select {
	case wc := <-w.Chan():
		b.Fatalf("unexpected data %v", wc)
	default:
	}
}
func BenchmarkWatchableStoreUnsyncedCancel(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := NewStore(be, &lease.FakeLessor{}, nil)
	ws := &watchableStore{store: s, unsynced: newWatcherGroup(), synced: newWatcherGroup()}
	defer func() {
		ws.store.Close()
		os.Remove(tmpPath)
	}()
	testKey := []byte("foo")
	testValue := []byte("bar")
	s.Put(testKey, testValue, lease.NoLease)
	w := ws.NewWatchStream()
	const k int = 2
	benchSampleN := b.N
	watcherN := k * benchSampleN
	watchIDs := make([]WatchID, watcherN)
	for i := 0; i < watcherN; i++ {
		watchIDs[i] = w.Watch(testKey, nil, 1)
	}
	ix := rand.Perm(watcherN)
	b.ResetTimer()
	b.ReportAllocs()
	for _, idx := range ix[:benchSampleN] {
		if err := w.Cancel(watchIDs[idx]); err != nil {
			b.Error(err)
		}
	}
}
func BenchmarkWatchableStoreSyncedCancel(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(be, &lease.FakeLessor{}, nil)
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKey := []byte("foo")
	testValue := []byte("bar")
	s.Put(testKey, testValue, lease.NoLease)
	w := s.NewWatchStream()
	const watcherN = 1000000
	watchIDs := make([]WatchID, watcherN)
	for i := 0; i < watcherN; i++ {
		watchIDs[i] = w.Watch(testKey, nil, 0)
	}
	ix := rand.Perm(watcherN)
	b.ResetTimer()
	b.ReportAllocs()
	for _, idx := range ix {
		if err := w.Cancel(watchIDs[idx]); err != nil {
			b.Error(err)
		}
	}
}
