package mvcc

import (
	"context"
	"encoding/binary"
	"errors"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/coreos/etcd/pkg/schedule"
	"github.com/coreos/pkg/capnslog"
	"hash/crc32"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var (
	keyBucketName           = []byte("key")
	metaBucketName          = []byte("meta")
	consistentIndexKeyName  = []byte("consistent_index")
	scheduledCompactKeyName = []byte("scheduledCompactRev")
	finishedCompactKeyName  = []byte("finishedCompactRev")
	ErrCompacted            = errors.New("mvcc: required revision has been compacted")
	ErrFutureRev            = errors.New("mvcc: required revision is a future revision")
	ErrCanceled             = errors.New("mvcc: watcher is canceled")
	ErrClosed               = errors.New("mvcc: closed")
	plog                    = capnslog.NewPackageLogger("github.com/coreos/etcd", "mvcc")
)

const (
	markedRevBytesLen      = revBytesLen + 1
	markBytePosition       = markedRevBytesLen - 1
	markTombstone     byte = 't'
)

var restoreChunkKeys = 10000

type ConsistentIndexGetter interface{ ConsistentIndex() uint64 }
type store struct {
	ReadView
	WriteView
	consistentIndex uint64
	mu              sync.RWMutex
	ig              ConsistentIndexGetter
	b               backend.Backend
	kvindex         index
	le              lease.Lessor
	revMu           sync.RWMutex
	currentRev      int64
	compactMainRev  int64
	bytesBuf8       []byte
	fifoSched       schedule.Scheduler
	stopc           chan struct{}
}

func NewStore(b backend.Backend, le lease.Lessor, ig ConsistentIndexGetter) *store {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := &store{b: b, ig: ig, kvindex: newTreeIndex(), le: le, currentRev: 1, compactMainRev: -1, bytesBuf8: make([]byte, 8), fifoSched: schedule.NewFIFOScheduler(), stopc: make(chan struct{})}
	s.ReadView = &readView{s}
	s.WriteView = &writeView{s}
	if s.le != nil {
		s.le.SetRangeDeleter(func() lease.TxnDelete {
			return s.Write()
		})
	}
	tx := s.b.BatchTx()
	tx.Lock()
	tx.UnsafeCreateBucket(keyBucketName)
	tx.UnsafeCreateBucket(metaBucketName)
	tx.Unlock()
	s.b.ForceCommit()
	if err := s.restore(); err != nil {
		panic("failed to recover store from backend")
	}
	return s
}
func (s *store) compactBarrier(ctx context.Context, ch chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ctx == nil || ctx.Err() != nil {
		s.mu.Lock()
		select {
		case <-s.stopc:
		default:
			f := func(ctx context.Context) {
				s.compactBarrier(ctx, ch)
			}
			s.fifoSched.Schedule(f)
		}
		s.mu.Unlock()
		return
	}
	close(ch)
}
func (s *store) Hash() (hash uint32, revision int64, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	start := time.Now()
	s.b.ForceCommit()
	h, err := s.b.Hash(DefaultIgnores)
	hashDurations.Observe(time.Since(start).Seconds())
	return h, s.currentRev, err
}
func (s *store) HashByRev(rev int64) (hash uint32, currentRev int64, compactRev int64, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	start := time.Now()
	s.mu.RLock()
	s.revMu.RLock()
	compactRev, currentRev = s.compactMainRev, s.currentRev
	s.revMu.RUnlock()
	if rev > 0 && rev <= compactRev {
		s.mu.RUnlock()
		return 0, 0, compactRev, ErrCompacted
	} else if rev > 0 && rev > currentRev {
		s.mu.RUnlock()
		return 0, currentRev, 0, ErrFutureRev
	}
	if rev == 0 {
		rev = currentRev
	}
	keep := s.kvindex.Keep(rev)
	tx := s.b.ReadTx()
	tx.Lock()
	defer tx.Unlock()
	s.mu.RUnlock()
	upper := revision{main: rev + 1}
	lower := revision{main: compactRev + 1}
	h := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	h.Write(keyBucketName)
	err = tx.UnsafeForEach(keyBucketName, func(k, v []byte) error {
		kr := bytesToRev(k)
		if !upper.GreaterThan(kr) {
			return nil
		}
		if lower.GreaterThan(kr) && len(keep) > 0 {
			if _, ok := keep[kr]; !ok {
				return nil
			}
		}
		h.Write(k)
		h.Write(v)
		return nil
	})
	hash = h.Sum32()
	hashRevDurations.Observe(time.Since(start).Seconds())
	return hash, currentRev, compactRev, err
}
func (s *store) Compact(rev int64) (<-chan struct{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	s.revMu.Lock()
	defer s.revMu.Unlock()
	if rev <= s.compactMainRev {
		ch := make(chan struct{})
		f := func(ctx context.Context) {
			s.compactBarrier(ctx, ch)
		}
		s.fifoSched.Schedule(f)
		return ch, ErrCompacted
	}
	if rev > s.currentRev {
		return nil, ErrFutureRev
	}
	start := time.Now()
	s.compactMainRev = rev
	rbytes := newRevBytes()
	revToBytes(revision{main: rev}, rbytes)
	tx := s.b.BatchTx()
	tx.Lock()
	tx.UnsafePut(metaBucketName, scheduledCompactKeyName, rbytes)
	tx.Unlock()
	s.b.ForceCommit()
	keep := s.kvindex.Compact(rev)
	ch := make(chan struct{})
	var j = func(ctx context.Context) {
		if ctx.Err() != nil {
			s.compactBarrier(ctx, ch)
			return
		}
		if !s.scheduleCompaction(rev, keep) {
			s.compactBarrier(nil, ch)
			return
		}
		close(ch)
	}
	s.fifoSched.Schedule(j)
	indexCompactionPauseDurations.Observe(float64(time.Since(start) / time.Millisecond))
	return ch, nil
}

var DefaultIgnores map[backend.IgnoreKey]struct{}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	DefaultIgnores = map[backend.IgnoreKey]struct{}{{Bucket: string(metaBucketName), Key: string(consistentIndexKeyName)}: {}}
}
func (s *store) Commit() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	tx := s.b.BatchTx()
	tx.Lock()
	s.saveIndex(tx)
	tx.Unlock()
	s.b.ForceCommit()
}
func (s *store) Restore(b backend.Backend) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	close(s.stopc)
	s.fifoSched.Stop()
	atomic.StoreUint64(&s.consistentIndex, 0)
	s.b = b
	s.kvindex = newTreeIndex()
	s.currentRev = 1
	s.compactMainRev = -1
	s.fifoSched = schedule.NewFIFOScheduler()
	s.stopc = make(chan struct{})
	return s.restore()
}
func (s *store) restore() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := s.b
	reportDbTotalSizeInBytesMu.Lock()
	reportDbTotalSizeInBytes = func() float64 {
		return float64(b.Size())
	}
	reportDbTotalSizeInBytesMu.Unlock()
	reportDbTotalSizeInUseInBytesMu.Lock()
	reportDbTotalSizeInUseInBytes = func() float64 {
		return float64(b.SizeInUse())
	}
	reportDbTotalSizeInUseInBytesMu.Unlock()
	min, max := newRevBytes(), newRevBytes()
	revToBytes(revision{main: 1}, min)
	revToBytes(revision{main: math.MaxInt64, sub: math.MaxInt64}, max)
	keyToLease := make(map[string]lease.LeaseID)
	tx := s.b.BatchTx()
	tx.Lock()
	_, finishedCompactBytes := tx.UnsafeRange(metaBucketName, finishedCompactKeyName, nil, 0)
	if len(finishedCompactBytes) != 0 {
		s.compactMainRev = bytesToRev(finishedCompactBytes[0]).main
		plog.Printf("restore compact to %d", s.compactMainRev)
	}
	_, scheduledCompactBytes := tx.UnsafeRange(metaBucketName, scheduledCompactKeyName, nil, 0)
	scheduledCompact := int64(0)
	if len(scheduledCompactBytes) != 0 {
		scheduledCompact = bytesToRev(scheduledCompactBytes[0]).main
	}
	keysGauge.Set(0)
	rkvc, revc := restoreIntoIndex(s.kvindex)
	for {
		keys, vals := tx.UnsafeRange(keyBucketName, min, max, int64(restoreChunkKeys))
		if len(keys) == 0 {
			break
		}
		restoreChunk(rkvc, keys, vals, keyToLease)
		if len(keys) < restoreChunkKeys {
			break
		}
		newMin := bytesToRev(keys[len(keys)-1][:revBytesLen])
		newMin.sub++
		revToBytes(newMin, min)
	}
	close(rkvc)
	s.currentRev = <-revc
	if s.currentRev < s.compactMainRev {
		s.currentRev = s.compactMainRev
	}
	if scheduledCompact <= s.compactMainRev {
		scheduledCompact = 0
	}
	for key, lid := range keyToLease {
		if s.le == nil {
			panic("no lessor to attach lease")
		}
		err := s.le.Attach(lid, []lease.LeaseItem{{Key: key}})
		if err != nil {
			plog.Errorf("unexpected Attach error: %v", err)
		}
	}
	tx.Unlock()
	if scheduledCompact != 0 {
		s.Compact(scheduledCompact)
		plog.Printf("resume scheduled compaction at %d", scheduledCompact)
	}
	return nil
}

type revKeyValue struct {
	key  []byte
	kv   mvccpb.KeyValue
	kstr string
}

func restoreIntoIndex(idx index) (chan<- revKeyValue, <-chan int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rkvc, revc := make(chan revKeyValue, restoreChunkKeys), make(chan int64, 1)
	go func() {
		currentRev := int64(1)
		defer func() {
			revc <- currentRev
		}()
		kiCache := make(map[string]*keyIndex, restoreChunkKeys)
		for rkv := range rkvc {
			ki, ok := kiCache[rkv.kstr]
			if !ok && len(kiCache) >= restoreChunkKeys {
				i := 10
				for k := range kiCache {
					delete(kiCache, k)
					if i--; i == 0 {
						break
					}
				}
			}
			if !ok {
				ki = &keyIndex{key: rkv.kv.Key}
				if idxKey := idx.KeyIndex(ki); idxKey != nil {
					kiCache[rkv.kstr], ki = idxKey, idxKey
					ok = true
				}
			}
			rev := bytesToRev(rkv.key)
			currentRev = rev.main
			if ok {
				if isTombstone(rkv.key) {
					ki.tombstone(rev.main, rev.sub)
					continue
				}
				ki.put(rev.main, rev.sub)
			} else if !isTombstone(rkv.key) {
				ki.restore(revision{rkv.kv.CreateRevision, 0}, rev, rkv.kv.Version)
				idx.Insert(ki)
				kiCache[rkv.kstr] = ki
			}
		}
	}()
	return rkvc, revc
}
func restoreChunk(kvc chan<- revKeyValue, keys, vals [][]byte, keyToLease map[string]lease.LeaseID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, key := range keys {
		rkv := revKeyValue{key: key}
		if err := rkv.kv.Unmarshal(vals[i]); err != nil {
			plog.Fatalf("cannot unmarshal event: %v", err)
		}
		rkv.kstr = string(rkv.kv.Key)
		if isTombstone(key) {
			delete(keyToLease, rkv.kstr)
		} else if lid := lease.LeaseID(rkv.kv.Lease); lid != lease.NoLease {
			keyToLease[rkv.kstr] = lid
		} else {
			delete(keyToLease, rkv.kstr)
		}
		kvc <- rkv
	}
}
func (s *store) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(s.stopc)
	s.fifoSched.Stop()
	return nil
}
func (s *store) saveIndex(tx backend.BatchTx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.ig == nil {
		return
	}
	bs := s.bytesBuf8
	ci := s.ig.ConsistentIndex()
	binary.BigEndian.PutUint64(bs, ci)
	tx.UnsafePut(metaBucketName, consistentIndexKeyName, bs)
	atomic.StoreUint64(&s.consistentIndex, ci)
}
func (s *store) ConsistentIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ci := atomic.LoadUint64(&s.consistentIndex); ci > 0 {
		return ci
	}
	tx := s.b.BatchTx()
	tx.Lock()
	defer tx.Unlock()
	_, vs := tx.UnsafeRange(metaBucketName, consistentIndexKeyName, nil, 0)
	if len(vs) == 0 {
		return 0
	}
	v := binary.BigEndian.Uint64(vs[0])
	atomic.StoreUint64(&s.consistentIndex, v)
	return v
}
func appendMarkTombstone(b []byte) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(b) != revBytesLen {
		plog.Panicf("cannot append mark to non normal revision bytes")
	}
	return append(b, markTombstone)
}
func isTombstone(b []byte) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(b) == markedRevBytesLen && b[markBytePosition] == markTombstone
}
