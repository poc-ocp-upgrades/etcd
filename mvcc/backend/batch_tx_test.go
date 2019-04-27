package backend

import (
	"reflect"
	"testing"
	"time"
	bolt "github.com/coreos/bbolt"
)

func TestBatchTxPut(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := NewTmpBackend(time.Hour, 10000)
	defer cleanup(b, tmpPath)
	tx := b.batchTx
	tx.Lock()
	defer tx.Unlock()
	tx.UnsafeCreateBucket([]byte("test"))
	v := []byte("bar")
	tx.UnsafePut([]byte("test"), []byte("foo"), v)
	for k := 0; k < 2; k++ {
		_, gv := tx.UnsafeRange([]byte("test"), []byte("foo"), nil, 0)
		if !reflect.DeepEqual(gv[0], v) {
			t.Errorf("v = %s, want %s", string(gv[0]), string(v))
		}
		tx.commit(false)
	}
}
func TestBatchTxRange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := NewTmpBackend(time.Hour, 10000)
	defer cleanup(b, tmpPath)
	tx := b.batchTx
	tx.Lock()
	defer tx.Unlock()
	tx.UnsafeCreateBucket([]byte("test"))
	allKeys := [][]byte{[]byte("foo"), []byte("foo1"), []byte("foo2")}
	allVals := [][]byte{[]byte("bar"), []byte("bar1"), []byte("bar2")}
	for i := range allKeys {
		tx.UnsafePut([]byte("test"), allKeys[i], allVals[i])
	}
	tests := []struct {
		key	[]byte
		endKey	[]byte
		limit	int64
		wkeys	[][]byte
		wvals	[][]byte
	}{{[]byte("foo"), nil, 0, allKeys[:1], allVals[:1]}, {[]byte("doo"), nil, 0, nil, nil}, {[]byte("foo"), []byte("foo1"), 0, allKeys[:1], allVals[:1]}, {[]byte("foo"), []byte("foo3"), 0, allKeys, allVals}, {[]byte("goo"), []byte("goo3"), 0, nil, nil}, {[]byte("foo"), []byte("foo3"), 1, allKeys[:1], allVals[:1]}, {[]byte("foo"), []byte("foo3"), 4, allKeys, allVals}}
	for i, tt := range tests {
		keys, vals := tx.UnsafeRange([]byte("test"), tt.key, tt.endKey, tt.limit)
		if !reflect.DeepEqual(keys, tt.wkeys) {
			t.Errorf("#%d: keys = %+v, want %+v", i, keys, tt.wkeys)
		}
		if !reflect.DeepEqual(vals, tt.wvals) {
			t.Errorf("#%d: vals = %+v, want %+v", i, vals, tt.wvals)
		}
	}
}
func TestBatchTxDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := NewTmpBackend(time.Hour, 10000)
	defer cleanup(b, tmpPath)
	tx := b.batchTx
	tx.Lock()
	defer tx.Unlock()
	tx.UnsafeCreateBucket([]byte("test"))
	tx.UnsafePut([]byte("test"), []byte("foo"), []byte("bar"))
	tx.UnsafeDelete([]byte("test"), []byte("foo"))
	for k := 0; k < 2; k++ {
		ks, _ := tx.UnsafeRange([]byte("test"), []byte("foo"), nil, 0)
		if len(ks) != 0 {
			t.Errorf("keys on foo = %v, want nil", ks)
		}
		tx.commit(false)
	}
}
func TestBatchTxCommit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := NewTmpBackend(time.Hour, 10000)
	defer cleanup(b, tmpPath)
	tx := b.batchTx
	tx.Lock()
	tx.UnsafeCreateBucket([]byte("test"))
	tx.UnsafePut([]byte("test"), []byte("foo"), []byte("bar"))
	tx.Unlock()
	tx.Commit()
	b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("test"))
		if bucket == nil {
			t.Errorf("bucket test does not exit")
			return nil
		}
		v := bucket.Get([]byte("foo"))
		if v == nil {
			t.Errorf("foo key failed to written in backend")
		}
		return nil
	})
}
func TestBatchTxBatchLimitCommit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := NewTmpBackend(time.Hour, 1)
	defer cleanup(b, tmpPath)
	tx := b.batchTx
	tx.Lock()
	tx.UnsafeCreateBucket([]byte("test"))
	tx.UnsafePut([]byte("test"), []byte("foo"), []byte("bar"))
	tx.Unlock()
	b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("test"))
		if bucket == nil {
			t.Errorf("bucket test does not exit")
			return nil
		}
		v := bucket.Get([]byte("foo"))
		if v == nil {
			t.Errorf("foo key failed to written in backend")
		}
		return nil
	})
}
