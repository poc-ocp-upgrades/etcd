package backend

import (
	"bytes"
	"math"
	"sync"
	bolt "github.com/coreos/bbolt"
)

var safeRangeBucket = []byte("key")

type ReadTx interface {
	Lock()
	Unlock()
	UnsafeRange(bucketName []byte, key, endKey []byte, limit int64) (keys [][]byte, vals [][]byte)
	UnsafeForEach(bucketName []byte, visitor func(k, v []byte) error) error
}
type readTx struct {
	mu	sync.RWMutex
	buf	txReadBuffer
	txmu	sync.RWMutex
	tx	*bolt.Tx
	buckets	map[string]*bolt.Bucket
}

func (rt *readTx) Lock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rt.mu.RLock()
}
func (rt *readTx) Unlock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rt.mu.RUnlock()
}
func (rt *readTx) UnsafeRange(bucketName, key, endKey []byte, limit int64) ([][]byte, [][]byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if endKey == nil {
		limit = 1
	}
	if limit <= 0 {
		limit = math.MaxInt64
	}
	if limit > 1 && !bytes.Equal(bucketName, safeRangeBucket) {
		panic("do not use unsafeRange on non-keys bucket")
	}
	keys, vals := rt.buf.Range(bucketName, key, endKey, limit)
	if int64(len(keys)) == limit {
		return keys, vals
	}
	bn := string(bucketName)
	rt.txmu.RLock()
	bucket, ok := rt.buckets[bn]
	rt.txmu.RUnlock()
	if !ok {
		rt.txmu.Lock()
		bucket = rt.tx.Bucket(bucketName)
		rt.buckets[bn] = bucket
		rt.txmu.Unlock()
	}
	if bucket == nil {
		return keys, vals
	}
	rt.txmu.Lock()
	c := bucket.Cursor()
	rt.txmu.Unlock()
	k2, v2 := unsafeRange(c, key, endKey, limit-int64(len(keys)))
	return append(k2, keys...), append(v2, vals...)
}
func (rt *readTx) UnsafeForEach(bucketName []byte, visitor func(k, v []byte) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dups := make(map[string]struct{})
	getDups := func(k, v []byte) error {
		dups[string(k)] = struct{}{}
		return nil
	}
	visitNoDup := func(k, v []byte) error {
		if _, ok := dups[string(k)]; ok {
			return nil
		}
		return visitor(k, v)
	}
	if err := rt.buf.ForEach(bucketName, getDups); err != nil {
		return err
	}
	rt.txmu.Lock()
	err := unsafeForEach(rt.tx, bucketName, visitNoDup)
	rt.txmu.Unlock()
	if err != nil {
		return err
	}
	return rt.buf.ForEach(bucketName, visitor)
}
func (rt *readTx) reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rt.buf.reset()
	rt.buckets = make(map[string]*bolt.Bucket)
	rt.tx = nil
}
