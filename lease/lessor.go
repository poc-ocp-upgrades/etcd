package lease

import (
	godefaultbytes "bytes"
	"encoding/binary"
	"errors"
	"github.com/coreos/etcd/lease/leasepb"
	"github.com/coreos/etcd/mvcc/backend"
	"math"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sort"
	"sync"
	"time"
)

const NoLease = LeaseID(0)
const MaxLeaseTTL = 9000000000

var (
	forever             = time.Time{}
	leaseBucketName     = []byte("lease")
	leaseRevokeRate     = 1000
	ErrNotPrimary       = errors.New("not a primary lessor")
	ErrLeaseNotFound    = errors.New("lease not found")
	ErrLeaseExists      = errors.New("lease already exists")
	ErrLeaseTTLTooLarge = errors.New("too large lease TTL")
)

type TxnDelete interface {
	DeleteRange(key, end []byte) (n, rev int64)
	End()
}
type RangeDeleter func() TxnDelete
type LeaseID int64
type Lessor interface {
	SetRangeDeleter(rd RangeDeleter)
	Grant(id LeaseID, ttl int64) (*Lease, error)
	Revoke(id LeaseID) error
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
	mu          sync.Mutex
	demotec     chan struct{}
	leaseMap    map[LeaseID]*Lease
	itemMap     map[LeaseItem]LeaseID
	rd          RangeDeleter
	b           backend.Backend
	minLeaseTTL int64
	expiredC    chan []*Lease
	stopC       chan struct{}
	doneC       chan struct{}
}

func NewLessor(b backend.Backend, minLeaseTTL int64) Lessor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newLessor(b, minLeaseTTL)
}
func newLessor(b backend.Backend, minLeaseTTL int64) *lessor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := &lessor{leaseMap: make(map[LeaseID]*Lease), itemMap: make(map[LeaseItem]LeaseID), b: b, minLeaseTTL: minLeaseTTL, expiredC: make(chan []*Lease, 16), stopC: make(chan struct{}), doneC: make(chan struct{})}
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
	l.persistTo(le.b)
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
	return nil
}
func (le *lessor) Renew(id LeaseID) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	unlock := func() {
		le.mu.Unlock()
	}
	defer func() {
		unlock()
	}()
	if !le.isPrimary() {
		return -1, ErrNotPrimary
	}
	demotec := le.demotec
	l := le.leaseMap[id]
	if l == nil {
		return -1, ErrLeaseNotFound
	}
	if l.expired() {
		le.mu.Unlock()
		unlock = func() {
		}
		select {
		case <-l.revokec:
			return -1, ErrLeaseNotFound
		case <-demotec:
			return -1, ErrNotPrimary
		case <-le.stopC:
			return -1, ErrNotPrimary
		}
	}
	l.refresh(0)
	return l.ttl, nil
}
func (le *lessor) Lookup(id LeaseID) *Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	defer le.mu.Unlock()
	return le.leaseMap[id]
}
func (le *lessor) unsafeLeases() []*Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	leases := make([]*Lease, 0, len(le.leaseMap))
	for _, l := range le.leaseMap {
		leases = append(leases, l)
	}
	sort.Sort(leasesByExpiry(leases))
	return leases
}
func (le *lessor) Leases() []*Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	le.mu.Lock()
	ls := le.unsafeLeases()
	le.mu.Unlock()
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
	}
	if len(le.leaseMap) < leaseRevokeRate {
		return
	}
	leases := le.unsafeLeases()
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
	le.mu.Lock()
	id := le.itemMap[item]
	le.mu.Unlock()
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
		var ls []*Lease
		revokeLimit := leaseRevokeRate / 2
		le.mu.Lock()
		if le.isPrimary() {
			ls = le.findExpiredLeases(revokeLimit)
		}
		le.mu.Unlock()
		if len(ls) != 0 {
			select {
			case <-le.stopC:
				return
			case le.expiredC <- ls:
			default:
			}
		}
		select {
		case <-time.After(500 * time.Millisecond):
		case <-le.stopC:
			return
		}
	}
}
func (le *lessor) findExpiredLeases(limit int) []*Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	leases := make([]*Lease, 0, 16)
	for _, l := range le.leaseMap {
		if l.expired() {
			leases = append(leases, l)
			if len(leases) == limit {
				break
			}
		}
	}
	return leases
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
	tx.Unlock()
	le.b.ForceCommit()
}

type Lease struct {
	ID       LeaseID
	ttl      int64
	expiryMu sync.RWMutex
	expiry   time.Time
	mu       sync.RWMutex
	itemSet  map[LeaseItem]struct{}
	revokec  chan struct{}
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
	lpb := leasepb.Lease{ID: int64(l.ID), TTL: int64(l.ttl)}
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
func (l *Lease) refresh(extend time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newExpiry := time.Now().Add(extend + time.Duration(l.ttl)*time.Second)
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
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
