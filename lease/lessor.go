package lease

import (
	"container/heap"
	"context"
	"encoding/binary"
	"errors"
	"math"
	"sort"
	"sync"
	"time"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/lease/leasepb"
	"go.etcd.io/etcd/mvcc/backend"
	"go.uber.org/zap"
)

const NoLease = LeaseID(0)
const MaxLeaseTTL = 9000000000

var (
	forever				= time.Time{}
	leaseBucketName			= []byte("lease")
	leaseRevokeRate			= 1000
	leaseCheckpointRate		= 1000
	maxLeaseCheckpointBatchSize	= 1000
	ErrNotPrimary			= errors.New("not a primary lessor")
	ErrLeaseNotFound		= errors.New("lease not found")
	ErrLeaseExists			= errors.New("lease already exists")
	ErrLeaseTTLTooLarge		= errors.New("too large lease TTL")
)

type TxnDelete interface {
	DeleteRange(key, end []byte) (n, rev int64)
	End()
}
type RangeDeleter func() TxnDelete
type Checkpointer func(ctx context.Context, lc *pb.LeaseCheckpointRequest)
type LeaseID int64
type Lessor interface {
	SetRangeDeleter(rd RangeDeleter)
	SetCheckpointer(cp Checkpointer)
	Grant(id LeaseID, ttl int64) (*Lease, error)
	Revoke(id LeaseID) error
	Checkpoint(id LeaseID, remainingTTL int64) error
	Attach(id LeaseID, items []LeaseItem) error
	GetLease(item LeaseItem) LeaseID
	Detach(id LeaseID, items []LeaseItem) error
	Promote(extend time.Duration)
	Demote()
	Renew(id LeaseID) (int64, error)
	Lookup(id LeaseID) *Lease
	Leases() []*Lease
	ExpiredLeasesC() <-chan []*Lease
	Recover(b backend.Backend, rd RangeDeleter)
	Stop()
}
type lessor struct {
	mu			sync.RWMutex
	demotec			chan struct{}
	leaseMap		map[LeaseID]*Lease
	leaseHeap		LeaseQueue
	leaseCheckpointHeap	LeaseQueue
	itemMap			map[LeaseItem]LeaseID
	rd			RangeDeleter
	cp			Checkpointer
	b			backend.Backend
	minLeaseTTL		int64
	expiredC		chan []*Lease
	stopC			chan struct{}
	doneC			chan struct{}
	lg			*zap.Logger
	checkpointInterval	time.Duration
}
type LessorConfig struct {
	MinLeaseTTL		int64
	CheckpointInterval	time.Duration
}

func NewLessor(lg *zap.Logger, b backend.Backend, cfg LessorConfig) Lessor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newLessor(lg, b, cfg)
}
func newLessor(lg *zap.Logger, b backend.Backend, cfg LessorConfig) *lessor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	checkpointInterval := cfg.CheckpointInterval
	if checkpointInterval == 0 {
		checkpointInterval = 5 * time.Minute
	}
	l := &lessor{leaseMap: make(map[LeaseID]*Lease), itemMap: make(map[LeaseItem]LeaseID), leaseHeap: make(LeaseQueue, 0), leaseCheckpointHeap: make(LeaseQueue, 0), b: b, minLeaseTTL: cfg.MinLeaseTTL, checkpointInterval: checkpointInterval, expiredC: make(chan []*Lease, 16), stopC: make(chan struct{}), doneC: make(chan struct{}), lg: lg}
	l.initAndRecover()
	go l.runLoop()
	return l
}
func (le *lessor) isPrimary() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return le.demotec != nil
}
func (le *lessor) SetRangeDeleter(rd RangeDeleter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	defer le.mu.Unlock()
	le.rd = rd
}
func (le *lessor) SetCheckpointer(cp Checkpointer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	defer le.mu.Unlock()
	le.cp = cp
}
func (le *lessor) Grant(id LeaseID, ttl int64) (*Lease, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if id == NoLease {
		return nil, ErrLeaseNotFound
	}
	if ttl > MaxLeaseTTL {
		return nil, ErrLeaseTTLTooLarge
	}
	l := &Lease{ID: id, ttl: ttl, itemSet: make(map[LeaseItem]struct{}), revokec: make(chan struct{})}
	le.mu.Lock()
	defer le.mu.Unlock()
	if _, ok := le.leaseMap[id]; ok {
		return nil, ErrLeaseExists
	}
	if l.ttl < le.minLeaseTTL {
		l.ttl = le.minLeaseTTL
	}
	if le.isPrimary() {
		l.refresh(0)
	} else {
		l.forever()
	}
	le.leaseMap[id] = l
	item := &LeaseWithTime{id: l.ID, time: l.expiry.UnixNano()}
	heap.Push(&le.leaseHeap, item)
	l.persistTo(le.b)
	leaseTotalTTLs.Observe(float64(l.ttl))
	leaseGranted.Inc()
	if le.isPrimary() {
		le.scheduleCheckpointIfNeeded(l)
	}
	return l, nil
}
func (le *lessor) Revoke(id LeaseID) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	l := le.leaseMap[id]
	if l == nil {
		le.mu.Unlock()
		return ErrLeaseNotFound
	}
	defer close(l.revokec)
	le.mu.Unlock()
	if le.rd == nil {
		return nil
	}
	txn := le.rd()
	keys := l.Keys()
	sort.StringSlice(keys).Sort()
	for _, key := range keys {
		txn.DeleteRange([]byte(key), nil)
	}
	le.mu.Lock()
	defer le.mu.Unlock()
	delete(le.leaseMap, l.ID)
	le.b.BatchTx().UnsafeDelete(leaseBucketName, int64ToBytes(int64(l.ID)))
	txn.End()
	leaseRevoked.Inc()
	return nil
}
func (le *lessor) Checkpoint(id LeaseID, remainingTTL int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	defer le.mu.Unlock()
	if l, ok := le.leaseMap[id]; ok {
		l.remainingTTL = remainingTTL
		if le.isPrimary() {
			le.scheduleCheckpointIfNeeded(l)
		}
	}
	return nil
}
func (le *lessor) Renew(id LeaseID) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.RLock()
	if !le.isPrimary() {
		le.mu.RUnlock()
		return -1, ErrNotPrimary
	}
	demotec := le.demotec
	l := le.leaseMap[id]
	if l == nil {
		le.mu.RUnlock()
		return -1, ErrLeaseNotFound
	}
	clearRemainingTTL := le.cp != nil && l.remainingTTL > 0
	le.mu.RUnlock()
	if l.expired() {
		select {
		case <-l.revokec:
			return -1, ErrLeaseNotFound
		case <-demotec:
			return -1, ErrNotPrimary
		case <-le.stopC:
			return -1, ErrNotPrimary
		}
	}
	if clearRemainingTTL {
		le.cp(context.Background(), &pb.LeaseCheckpointRequest{Checkpoints: []*pb.LeaseCheckpoint{{ID: int64(l.ID), Remaining_TTL: 0}}})
	}
	le.mu.Lock()
	l.refresh(0)
	item := &LeaseWithTime{id: l.ID, time: l.expiry.UnixNano()}
	heap.Push(&le.leaseHeap, item)
	le.mu.Unlock()
	leaseRenewed.Inc()
	return l.ttl, nil
}
func (le *lessor) Lookup(id LeaseID) *Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.RLock()
	defer le.mu.RUnlock()
	return le.leaseMap[id]
}
func (le *lessor) unsafeLeases() []*Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	leases := make([]*Lease, 0, len(le.leaseMap))
	for _, l := range le.leaseMap {
		leases = append(leases, l)
	}
	return leases
}
func (le *lessor) Leases() []*Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.RLock()
	ls := le.unsafeLeases()
	le.mu.RUnlock()
	sort.Sort(leasesByExpiry(ls))
	return ls
}
func (le *lessor) Promote(extend time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	defer le.mu.Unlock()
	le.demotec = make(chan struct{})
	for _, l := range le.leaseMap {
		l.refresh(extend)
		item := &LeaseWithTime{id: l.ID, time: l.expiry.UnixNano()}
		heap.Push(&le.leaseHeap, item)
	}
	if len(le.leaseMap) < leaseRevokeRate {
		return
	}
	leases := le.unsafeLeases()
	sort.Sort(leasesByExpiry(leases))
	baseWindow := leases[0].Remaining()
	nextWindow := baseWindow + time.Second
	expires := 0
	targetExpiresPerSecond := (3 * leaseRevokeRate) / 4
	for _, l := range leases {
		remaining := l.Remaining()
		if remaining > nextWindow {
			baseWindow = remaining
			nextWindow = baseWindow + time.Second
			expires = 1
			continue
		}
		expires++
		if expires <= targetExpiresPerSecond {
			continue
		}
		rateDelay := float64(time.Second) * (float64(expires) / float64(targetExpiresPerSecond))
		rateDelay -= float64(remaining - baseWindow)
		delay := time.Duration(rateDelay)
		nextWindow = baseWindow + delay
		l.refresh(delay + extend)
		item := &LeaseWithTime{id: l.ID, time: l.expiry.UnixNano()}
		heap.Push(&le.leaseHeap, item)
		le.scheduleCheckpointIfNeeded(l)
	}
}

type leasesByExpiry []*Lease

func (le leasesByExpiry) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(le)
}
func (le leasesByExpiry) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return le[i].Remaining() < le[j].Remaining()
}
func (le leasesByExpiry) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le[i], le[j] = le[j], le[i]
}
func (le *lessor) Demote() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	defer le.mu.Unlock()
	for _, l := range le.leaseMap {
		l.forever()
	}
	le.clearScheduledLeasesCheckpoints()
	if le.demotec != nil {
		close(le.demotec)
		le.demotec = nil
	}
}
func (le *lessor) Attach(id LeaseID, items []LeaseItem) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	defer le.mu.Unlock()
	l := le.leaseMap[id]
	if l == nil {
		return ErrLeaseNotFound
	}
	l.mu.Lock()
	for _, it := range items {
		l.itemSet[it] = struct{}{}
		le.itemMap[it] = id
	}
	l.mu.Unlock()
	return nil
}
func (le *lessor) GetLease(item LeaseItem) LeaseID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.RLock()
	id := le.itemMap[item]
	le.mu.RUnlock()
	return id
}
func (le *lessor) Detach(id LeaseID, items []LeaseItem) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	defer le.mu.Unlock()
	l := le.leaseMap[id]
	if l == nil {
		return ErrLeaseNotFound
	}
	l.mu.Lock()
	for _, it := range items {
		delete(l.itemSet, it)
		delete(le.itemMap, it)
	}
	l.mu.Unlock()
	return nil
}
func (le *lessor) Recover(b backend.Backend, rd RangeDeleter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	defer le.mu.Unlock()
	le.b = b
	le.rd = rd
	le.leaseMap = make(map[LeaseID]*Lease)
	le.itemMap = make(map[LeaseItem]LeaseID)
	le.initAndRecover()
}
func (le *lessor) ExpiredLeasesC() <-chan []*Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return le.expiredC
}
func (le *lessor) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(le.stopC)
	<-le.doneC
}
func (le *lessor) runLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(le.doneC)
	for {
		le.revokeExpiredLeases()
		le.checkpointScheduledLeases()
		select {
		case <-time.After(500 * time.Millisecond):
		case <-le.stopC:
			return
		}
	}
}
func (le *lessor) revokeExpiredLeases() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ls []*Lease
	revokeLimit := leaseRevokeRate / 2
	le.mu.RLock()
	if le.isPrimary() {
		ls = le.findExpiredLeases(revokeLimit)
	}
	le.mu.RUnlock()
	if len(ls) != 0 {
		select {
		case <-le.stopC:
			return
		case le.expiredC <- ls:
		default:
		}
	}
}
func (le *lessor) checkpointScheduledLeases() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cps []*pb.LeaseCheckpoint
	for i := 0; i < leaseCheckpointRate/2; i++ {
		le.mu.Lock()
		if le.isPrimary() {
			cps = le.findDueScheduledCheckpoints(maxLeaseCheckpointBatchSize)
		}
		le.mu.Unlock()
		if len(cps) != 0 {
			le.cp(context.Background(), &pb.LeaseCheckpointRequest{Checkpoints: cps})
		}
		if len(cps) < maxLeaseCheckpointBatchSize {
			return
		}
	}
}
func (le *lessor) clearScheduledLeasesCheckpoints() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.leaseCheckpointHeap = make(LeaseQueue, 0)
}
func (le *lessor) expireExists() (l *Lease, ok bool, next bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if le.leaseHeap.Len() == 0 {
		return nil, false, false
	}
	item := le.leaseHeap[0]
	l = le.leaseMap[item.id]
	if l == nil {
		heap.Pop(&le.leaseHeap)
		return nil, false, true
	}
	if time.Now().UnixNano() < item.time {
		return l, false, false
	}
	heap.Pop(&le.leaseHeap)
	return l, true, false
}
func (le *lessor) findExpiredLeases(limit int) []*Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	leases := make([]*Lease, 0, 16)
	for {
		l, ok, next := le.expireExists()
		if !ok && !next {
			break
		}
		if !ok {
			continue
		}
		if next {
			continue
		}
		if l.expired() {
			leases = append(leases, l)
			if len(leases) == limit {
				break
			}
		}
	}
	return leases
}
func (le *lessor) scheduleCheckpointIfNeeded(lease *Lease) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if le.cp == nil {
		return
	}
	if lease.RemainingTTL() > int64(le.checkpointInterval.Seconds()) {
		if le.lg != nil {
			le.lg.Debug("Scheduling lease checkpoint", zap.Int64("leaseID", int64(lease.ID)), zap.Duration("intervalSeconds", le.checkpointInterval))
		}
		heap.Push(&le.leaseCheckpointHeap, &LeaseWithTime{id: lease.ID, time: time.Now().Add(le.checkpointInterval).UnixNano()})
	}
}
func (le *lessor) findDueScheduledCheckpoints(checkpointLimit int) []*pb.LeaseCheckpoint {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if le.cp == nil {
		return nil
	}
	now := time.Now()
	cps := []*pb.LeaseCheckpoint{}
	for le.leaseCheckpointHeap.Len() > 0 && len(cps) < checkpointLimit {
		lt := le.leaseCheckpointHeap[0]
		if lt.time > now.UnixNano() {
			return cps
		}
		heap.Pop(&le.leaseCheckpointHeap)
		var l *Lease
		var ok bool
		if l, ok = le.leaseMap[lt.id]; !ok {
			continue
		}
		if !now.Before(l.expiry) {
			continue
		}
		remainingTTL := int64(math.Ceil(l.expiry.Sub(now).Seconds()))
		if remainingTTL >= l.ttl {
			continue
		}
		if le.lg != nil {
			le.lg.Debug("Checkpointing lease", zap.Int64("leaseID", int64(lt.id)), zap.Int64("remainingTTL", remainingTTL))
		}
		cps = append(cps, &pb.LeaseCheckpoint{ID: int64(lt.id), Remaining_TTL: remainingTTL})
	}
	return cps
}
func (le *lessor) initAndRecover() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tx := le.b.BatchTx()
	tx.Lock()
	tx.UnsafeCreateBucket(leaseBucketName)
	_, vs := tx.UnsafeRange(leaseBucketName, int64ToBytes(0), int64ToBytes(math.MaxInt64), 0)
	for i := range vs {
		var lpb leasepb.Lease
		err := lpb.Unmarshal(vs[i])
		if err != nil {
			tx.Unlock()
			panic("failed to unmarshal lease proto item")
		}
		ID := LeaseID(lpb.ID)
		if lpb.TTL < le.minLeaseTTL {
			lpb.TTL = le.minLeaseTTL
		}
		le.leaseMap[ID] = &Lease{ID: ID, ttl: lpb.TTL, itemSet: make(map[LeaseItem]struct{}), expiry: forever, revokec: make(chan struct{})}
	}
	heap.Init(&le.leaseHeap)
	heap.Init(&le.leaseCheckpointHeap)
	tx.Unlock()
	le.b.ForceCommit()
}

type Lease struct {
	ID		LeaseID
	ttl		int64
	remainingTTL	int64
	expiryMu	sync.RWMutex
	expiry		time.Time
	mu		sync.RWMutex
	itemSet		map[LeaseItem]struct{}
	revokec		chan struct{}
}

func (l *Lease) expired() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.Remaining() <= 0
}
func (l *Lease) persistTo(b backend.Backend) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := int64ToBytes(int64(l.ID))
	lpb := leasepb.Lease{ID: int64(l.ID), TTL: l.ttl, RemainingTTL: l.remainingTTL}
	val, err := lpb.Marshal()
	if err != nil {
		panic("failed to marshal lease proto item")
	}
	b.BatchTx().Lock()
	b.BatchTx().UnsafePut(leaseBucketName, key, val)
	b.BatchTx().Unlock()
}
func (l *Lease) TTL() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.ttl
}
func (l *Lease) RemainingTTL() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if l.remainingTTL > 0 {
		return l.remainingTTL
	}
	return l.ttl
}
func (l *Lease) refresh(extend time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newExpiry := time.Now().Add(extend + time.Duration(l.RemainingTTL())*time.Second)
	l.expiryMu.Lock()
	defer l.expiryMu.Unlock()
	l.expiry = newExpiry
}
func (l *Lease) forever() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.expiryMu.Lock()
	defer l.expiryMu.Unlock()
	l.expiry = forever
}
func (l *Lease) Keys() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.mu.RLock()
	keys := make([]string, 0, len(l.itemSet))
	for k := range l.itemSet {
		keys = append(keys, k.Key)
	}
	l.mu.RUnlock()
	return keys
}
func (l *Lease) Remaining() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.expiryMu.RLock()
	defer l.expiryMu.RUnlock()
	if l.expiry.IsZero() {
		return time.Duration(math.MaxInt64)
	}
	return time.Until(l.expiry)
}

type LeaseItem struct{ Key string }

func int64ToBytes(n int64) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(n))
	return bytes
}

type FakeLessor struct{}

func (fl *FakeLessor) SetRangeDeleter(dr RangeDeleter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (fl *FakeLessor) SetCheckpointer(cp Checkpointer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (fl *FakeLessor) Grant(id LeaseID, ttl int64) (*Lease, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func (fl *FakeLessor) Revoke(id LeaseID) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (fl *FakeLessor) Checkpoint(id LeaseID, remainingTTL int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (fl *FakeLessor) Attach(id LeaseID, items []LeaseItem) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (fl *FakeLessor) GetLease(item LeaseItem) LeaseID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (fl *FakeLessor) Detach(id LeaseID, items []LeaseItem) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (fl *FakeLessor) Promote(extend time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (fl *FakeLessor) Demote() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (fl *FakeLessor) Renew(id LeaseID) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 10, nil
}
func (fl *FakeLessor) Lookup(id LeaseID) *Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (fl *FakeLessor) Leases() []*Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (fl *FakeLessor) ExpiredLeasesC() <-chan []*Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (fl *FakeLessor) Recover(b backend.Backend, rd RangeDeleter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (fl *FakeLessor) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
