package backend

import (
	"bytes"
	"math"
	"sync"
	"sync/atomic"
	"time"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
)

type BatchTx interface {
	ReadTx
	UnsafeCreateBucket(name []byte)
	UnsafePut(bucketName []byte, key []byte, value []byte)
	UnsafeSeqPut(bucketName []byte, key []byte, value []byte)
	UnsafeDelete(bucketName []byte, key []byte)
	Commit()
	CommitAndStop()
}
type batchTx struct {
	sync.Mutex
	tx	*bolt.Tx
	backend	*backend
	pending	int
}

func (t *batchTx) Lock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Mutex.Lock()
}
func (t *batchTx) Unlock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.pending >= t.backend.batchLimit {
		t.commit(false)
	}
	t.Mutex.Unlock()
}
func (t *batchTx) RLock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("unexpected RLock")
}
func (t *batchTx) RUnlock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("unexpected RUnlock")
}
func (t *batchTx) UnsafeCreateBucket(name []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := t.tx.CreateBucket(name)
	if err != nil && err != bolt.ErrBucketExists {
		if t.backend.lg != nil {
			t.backend.lg.Fatal("failed to create a bucket", zap.String("bucket-name", string(name)), zap.Error(err))
		} else {
			plog.Fatalf("cannot create bucket %s (%v)", name, err)
		}
	}
	t.pending++
}
func (t *batchTx) UnsafePut(bucketName []byte, key []byte, value []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.unsafePut(bucketName, key, value, false)
}
func (t *batchTx) UnsafeSeqPut(bucketName []byte, key []byte, value []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.unsafePut(bucketName, key, value, true)
}
func (t *batchTx) unsafePut(bucketName []byte, key []byte, value []byte, seq bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bucket := t.tx.Bucket(bucketName)
	if bucket == nil {
		if t.backend.lg != nil {
			t.backend.lg.Fatal("failed to find a bucket", zap.String("bucket-name", string(bucketName)))
		} else {
			plog.Fatalf("bucket %s does not exist", bucketName)
		}
	}
	if seq {
		bucket.FillPercent = 0.9
	}
	if err := bucket.Put(key, value); err != nil {
		if t.backend.lg != nil {
			t.backend.lg.Fatal("failed to write to a bucket", zap.String("bucket-name", string(bucketName)), zap.Error(err))
		} else {
			plog.Fatalf("cannot put key into bucket (%v)", err)
		}
	}
	t.pending++
}
func (t *batchTx) UnsafeRange(bucketName, key, endKey []byte, limit int64) ([][]byte, [][]byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bucket := t.tx.Bucket(bucketName)
	if bucket == nil {
		if t.backend.lg != nil {
			t.backend.lg.Fatal("failed to find a bucket", zap.String("bucket-name", string(bucketName)))
		} else {
			plog.Fatalf("bucket %s does not exist", bucketName)
		}
	}
	return unsafeRange(bucket.Cursor(), key, endKey, limit)
}
func unsafeRange(c *bolt.Cursor, key, endKey []byte, limit int64) (keys [][]byte, vs [][]byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if limit <= 0 {
		limit = math.MaxInt64
	}
	var isMatch func(b []byte) bool
	if len(endKey) > 0 {
		isMatch = func(b []byte) bool {
			return bytes.Compare(b, endKey) < 0
		}
	} else {
		isMatch = func(b []byte) bool {
			return bytes.Equal(b, key)
		}
		limit = 1
	}
	for ck, cv := c.Seek(key); ck != nil && isMatch(ck); ck, cv = c.Next() {
		vs = append(vs, cv)
		keys = append(keys, ck)
		if limit == int64(len(keys)) {
			break
		}
	}
	return keys, vs
}
func (t *batchTx) UnsafeDelete(bucketName []byte, key []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bucket := t.tx.Bucket(bucketName)
	if bucket == nil {
		if t.backend.lg != nil {
			t.backend.lg.Fatal("failed to find a bucket", zap.String("bucket-name", string(bucketName)))
		} else {
			plog.Fatalf("bucket %s does not exist", bucketName)
		}
	}
	err := bucket.Delete(key)
	if err != nil {
		if t.backend.lg != nil {
			t.backend.lg.Fatal("failed to delete a key", zap.String("bucket-name", string(bucketName)), zap.Error(err))
		} else {
			plog.Fatalf("cannot delete key from bucket (%v)", err)
		}
	}
	t.pending++
}
func (t *batchTx) UnsafeForEach(bucketName []byte, visitor func(k, v []byte) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return unsafeForEach(t.tx, bucketName, visitor)
}
func unsafeForEach(tx *bolt.Tx, bucket []byte, visitor func(k, v []byte) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b := tx.Bucket(bucket); b != nil {
		return b.ForEach(visitor)
	}
	return nil
}
func (t *batchTx) Commit() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Lock()
	t.commit(false)
	t.Unlock()
}
func (t *batchTx) CommitAndStop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Lock()
	t.commit(true)
	t.Unlock()
}
func (t *batchTx) safePending() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	return t.pending
}
func (t *batchTx) commit(stop bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.tx != nil {
		if t.pending == 0 && !stop {
			return
		}
		start := time.Now()
		err := t.tx.Commit()
		rebalanceSec.Observe(t.tx.Stats().RebalanceTime.Seconds())
		spillSec.Observe(t.tx.Stats().SpillTime.Seconds())
		writeSec.Observe(t.tx.Stats().WriteTime.Seconds())
		commitSec.Observe(time.Since(start).Seconds())
		atomic.AddInt64(&t.backend.commits, 1)
		t.pending = 0
		if err != nil {
			if t.backend.lg != nil {
				t.backend.lg.Fatal("failed to commit tx", zap.Error(err))
			} else {
				plog.Fatalf("cannot commit tx (%s)", err)
			}
		}
	}
	if !stop {
		t.tx = t.backend.begin(true)
	}
}

type batchTxBuffered struct {
	batchTx
	buf	txWriteBuffer
}

func newBatchTxBuffered(backend *backend) *batchTxBuffered {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tx := &batchTxBuffered{batchTx: batchTx{backend: backend}, buf: txWriteBuffer{txBuffer: txBuffer{make(map[string]*bucketBuffer)}, seq: true}}
	tx.Commit()
	return tx
}
func (t *batchTxBuffered) Unlock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.pending != 0 {
		t.backend.readTx.Lock()
		t.buf.writeback(&t.backend.readTx.buf)
		t.backend.readTx.Unlock()
		if t.pending >= t.backend.batchLimit {
			t.commit(false)
		}
	}
	t.batchTx.Unlock()
}
func (t *batchTxBuffered) Commit() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Lock()
	t.commit(false)
	t.Unlock()
}
func (t *batchTxBuffered) CommitAndStop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Lock()
	t.commit(true)
	t.Unlock()
}
func (t *batchTxBuffered) commit(stop bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.backend.readTx.Lock()
	t.unsafeCommit(stop)
	t.backend.readTx.Unlock()
}
func (t *batchTxBuffered) unsafeCommit(stop bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.backend.readTx.tx != nil {
		if err := t.backend.readTx.tx.Rollback(); err != nil {
			if t.backend.lg != nil {
				t.backend.lg.Fatal("failed to rollback tx", zap.Error(err))
			} else {
				plog.Fatalf("cannot rollback tx (%s)", err)
			}
		}
		t.backend.readTx.reset()
	}
	t.batchTx.commit(stop)
	if !stop {
		t.backend.readTx.tx = t.backend.begin(false)
	}
}
func (t *batchTxBuffered) UnsafePut(bucketName []byte, key []byte, value []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.batchTx.UnsafePut(bucketName, key, value)
	t.buf.put(bucketName, key, value)
}
func (t *batchTxBuffered) UnsafeSeqPut(bucketName []byte, key []byte, value []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.batchTx.UnsafeSeqPut(bucketName, key, value)
	t.buf.putSeq(bucketName, key, value)
}
