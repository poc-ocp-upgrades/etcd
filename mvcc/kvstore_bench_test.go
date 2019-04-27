package mvcc

import (
	"sync/atomic"
	"testing"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc/backend"
)

type fakeConsistentIndex uint64

func (i *fakeConsistentIndex) ConsistentIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return atomic.LoadUint64((*uint64)(i))
}
func BenchmarkStorePut(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i fakeConsistentIndex
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := NewStore(be, &lease.FakeLessor{}, &i)
	defer cleanup(s, be, tmpPath)
	bytesN := 64
	keys := createBytesSlice(bytesN, b.N)
	vals := createBytesSlice(bytesN, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Put(keys[i], vals[i], lease.NoLease)
	}
}
func BenchmarkStoreRangeKey1(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	benchmarkStoreRange(b, 1)
}
func BenchmarkStoreRangeKey100(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	benchmarkStoreRange(b, 100)
}
func benchmarkStoreRange(b *testing.B, n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i fakeConsistentIndex
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := NewStore(be, &lease.FakeLessor{}, &i)
	defer cleanup(s, be, tmpPath)
	keys, val := createBytesSlice(64, n), createBytesSlice(64, 1)
	for i := range keys {
		s.Put(keys[i], val[0], lease.NoLease)
	}
	s.Commit()
	var begin, end []byte
	if n == 1 {
		begin, end = keys[0], nil
	} else {
		begin, end = []byte{}, []byte{}
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Range(begin, end, RangeOptions{})
	}
}
func BenchmarkConsistentIndex(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fci := fakeConsistentIndex(10)
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := NewStore(be, &lease.FakeLessor{}, &fci)
	defer cleanup(s, be, tmpPath)
	tx := s.b.BatchTx()
	tx.Lock()
	s.saveIndex(tx)
	tx.Unlock()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ConsistentIndex()
	}
}
func BenchmarkStorePutUpdate(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i fakeConsistentIndex
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := NewStore(be, &lease.FakeLessor{}, &i)
	defer cleanup(s, be, tmpPath)
	keys := createBytesSlice(64, 1)
	vals := createBytesSlice(1024, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Put(keys[0], vals[0], lease.NoLease)
	}
}
func BenchmarkStoreTxnPut(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i fakeConsistentIndex
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := NewStore(be, &lease.FakeLessor{}, &i)
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
func benchmarkStoreRestore(revsPerKey int, b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i fakeConsistentIndex
	be, tmpPath := backend.NewDefaultTmpBackend()
	s := NewStore(be, &lease.FakeLessor{}, &i)
	defer func() {
		cleanup(s, be, tmpPath)
	}()
	bytesN := 64
	keys := createBytesSlice(bytesN, b.N)
	vals := createBytesSlice(bytesN, b.N)
	for i := 0; i < b.N; i++ {
		for j := 0; j < revsPerKey; j++ {
			txn := s.Write()
			txn.Put(keys[i], vals[i], lease.NoLease)
			txn.End()
		}
	}
	s.Close()
	b.ReportAllocs()
	b.ResetTimer()
	s = NewStore(be, &lease.FakeLessor{}, &i)
}
func BenchmarkStoreRestoreRevs1(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	benchmarkStoreRestore(1, b)
}
func BenchmarkStoreRestoreRevs10(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	benchmarkStoreRestore(10, b)
}
func BenchmarkStoreRestoreRevs20(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	benchmarkStoreRestore(20, b)
}
