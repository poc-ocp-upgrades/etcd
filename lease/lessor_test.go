package lease

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"
	"github.com/coreos/etcd/mvcc/backend"
)

const (
	minLeaseTTL		= int64(5)
	minLeaseTTLDuration	= time.Duration(minLeaseTTL) * time.Second
)

func TestLessorGrant(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, be := NewTestBackend(t)
	defer os.RemoveAll(dir)
	defer be.Close()
	le := newLessor(be, minLeaseTTL)
	defer le.Stop()
	le.Promote(0)
	l, err := le.Grant(1, 1)
	if err != nil {
		t.Fatalf("could not grant lease 1 (%v)", err)
	}
	if l.ttl != minLeaseTTL {
		t.Fatalf("ttl = %v, expect minLeaseTTL %v", l.ttl, minLeaseTTL)
	}
	gl := le.Lookup(l.ID)
	if !reflect.DeepEqual(gl, l) {
		t.Errorf("lease = %v, want %v", gl, l)
	}
	if l.Remaining() < minLeaseTTLDuration-time.Second {
		t.Errorf("term = %v, want at least %v", l.Remaining(), minLeaseTTLDuration-time.Second)
	}
	_, err = le.Grant(1, 1)
	if err == nil {
		t.Errorf("allocated the same lease")
	}
	var nl *Lease
	nl, err = le.Grant(2, 1)
	if err != nil {
		t.Errorf("could not grant lease 2 (%v)", err)
	}
	if nl.ID == l.ID {
		t.Errorf("new lease.id = %x, want != %x", nl.ID, l.ID)
	}
	lss := []*Lease{gl, nl}
	leases := le.Leases()
	for i := range lss {
		if lss[i].ID != leases[i].ID {
			t.Fatalf("lease ID expected %d, got %d", lss[i].ID, leases[i].ID)
		}
		if lss[i].ttl != leases[i].ttl {
			t.Fatalf("ttl expected %d, got %d", lss[i].ttl, leases[i].ttl)
		}
	}
	be.BatchTx().Lock()
	_, vs := be.BatchTx().UnsafeRange(leaseBucketName, int64ToBytes(int64(l.ID)), nil, 0)
	if len(vs) != 1 {
		t.Errorf("len(vs) = %d, want 1", len(vs))
	}
	be.BatchTx().Unlock()
}
func TestLeaseConcurrentKeys(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, be := NewTestBackend(t)
	defer os.RemoveAll(dir)
	defer be.Close()
	le := newLessor(be, minLeaseTTL)
	defer le.Stop()
	le.SetRangeDeleter(func() TxnDelete {
		return newFakeDeleter(be)
	})
	l, err := le.Grant(1, 100)
	if err != nil {
		t.Fatalf("could not grant lease for 100s ttl (%v)", err)
	}
	itemn := 10
	items := make([]LeaseItem, itemn)
	for i := 0; i < itemn; i++ {
		items[i] = LeaseItem{Key: fmt.Sprintf("foo%d", i)}
	}
	if err = le.Attach(l.ID, items); err != nil {
		t.Fatalf("failed to attach items to the lease: %v", err)
	}
	donec := make(chan struct{})
	go func() {
		le.Detach(l.ID, items)
		close(donec)
	}()
	var wg sync.WaitGroup
	wg.Add(itemn)
	for i := 0; i < itemn; i++ {
		go func() {
			defer wg.Done()
			l.Keys()
		}()
	}
	<-donec
	wg.Wait()
}
func TestLessorRevoke(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, be := NewTestBackend(t)
	defer os.RemoveAll(dir)
	defer be.Close()
	le := newLessor(be, minLeaseTTL)
	defer le.Stop()
	var fd *fakeDeleter
	le.SetRangeDeleter(func() TxnDelete {
		fd = newFakeDeleter(be)
		return fd
	})
	l, err := le.Grant(1, 100)
	if err != nil {
		t.Fatalf("could not grant lease for 100s ttl (%v)", err)
	}
	items := []LeaseItem{{"foo"}, {"bar"}}
	if err = le.Attach(l.ID, items); err != nil {
		t.Fatalf("failed to attach items to the lease: %v", err)
	}
	if err = le.Revoke(l.ID); err != nil {
		t.Fatal("failed to revoke lease:", err)
	}
	if le.Lookup(l.ID) != nil {
		t.Errorf("got revoked lease %x", l.ID)
	}
	wdeleted := []string{"bar_", "foo_"}
	sort.Strings(fd.deleted)
	if !reflect.DeepEqual(fd.deleted, wdeleted) {
		t.Errorf("deleted= %v, want %v", fd.deleted, wdeleted)
	}
	be.BatchTx().Lock()
	_, vs := be.BatchTx().UnsafeRange(leaseBucketName, int64ToBytes(int64(l.ID)), nil, 0)
	if len(vs) != 0 {
		t.Errorf("len(vs) = %d, want 0", len(vs))
	}
	be.BatchTx().Unlock()
}
func TestLessorRenew(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, be := NewTestBackend(t)
	defer be.Close()
	defer os.RemoveAll(dir)
	le := newLessor(be, minLeaseTTL)
	defer le.Stop()
	le.Promote(0)
	l, err := le.Grant(1, minLeaseTTL)
	if err != nil {
		t.Fatalf("failed to grant lease (%v)", err)
	}
	le.mu.Lock()
	l.ttl = 10
	le.mu.Unlock()
	ttl, err := le.Renew(l.ID)
	if err != nil {
		t.Fatalf("failed to renew lease (%v)", err)
	}
	if ttl != l.ttl {
		t.Errorf("ttl = %d, want %d", ttl, l.ttl)
	}
	l = le.Lookup(l.ID)
	if l.Remaining() < 9*time.Second {
		t.Errorf("failed to renew the lease")
	}
}
func TestLessorRenewExtendPileup(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldRevokeRate := leaseRevokeRate
	defer func() {
		leaseRevokeRate = oldRevokeRate
	}()
	leaseRevokeRate = 10
	dir, be := NewTestBackend(t)
	defer os.RemoveAll(dir)
	le := newLessor(be, minLeaseTTL)
	ttl := int64(10)
	for i := 1; i <= leaseRevokeRate*10; i++ {
		if _, err := le.Grant(LeaseID(2*i), ttl); err != nil {
			t.Fatal(err)
		}
		if _, err := le.Grant(LeaseID(2*i+1), ttl+1); err != nil {
			t.Fatal(err)
		}
	}
	le.Stop()
	be.Close()
	bcfg := backend.DefaultBackendConfig()
	bcfg.Path = filepath.Join(dir, "be")
	be = backend.New(bcfg)
	defer be.Close()
	le = newLessor(be, minLeaseTTL)
	defer le.Stop()
	le.Promote(0)
	windowCounts := make(map[int64]int)
	for _, l := range le.leaseMap {
		s := int64(l.Remaining().Seconds() + 0.1)
		windowCounts[s]++
	}
	for i := ttl; i < ttl+20; i++ {
		c := windowCounts[i]
		if c > leaseRevokeRate {
			t.Errorf("expected at most %d expiring at %ds, got %d", leaseRevokeRate, i, c)
		}
		if c < leaseRevokeRate/2 {
			t.Errorf("expected at least %d expiring at %ds, got %d", leaseRevokeRate/2, i, c)
		}
	}
}
func TestLessorDetach(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, be := NewTestBackend(t)
	defer os.RemoveAll(dir)
	defer be.Close()
	le := newLessor(be, minLeaseTTL)
	defer le.Stop()
	le.SetRangeDeleter(func() TxnDelete {
		return newFakeDeleter(be)
	})
	l, err := le.Grant(1, 100)
	if err != nil {
		t.Fatalf("could not grant lease for 100s ttl (%v)", err)
	}
	items := []LeaseItem{{"foo"}, {"bar"}}
	if err := le.Attach(l.ID, items); err != nil {
		t.Fatalf("failed to attach items to the lease: %v", err)
	}
	if err := le.Detach(l.ID, items[0:1]); err != nil {
		t.Fatalf("failed to de-attach items to the lease: %v", err)
	}
	l = le.Lookup(l.ID)
	if len(l.itemSet) != 1 {
		t.Fatalf("len(l.itemSet) = %d, failed to de-attach items", len(l.itemSet))
	}
	if _, ok := l.itemSet[LeaseItem{"bar"}]; !ok {
		t.Fatalf("de-attached wrong item, want %q exists", "bar")
	}
}
func TestLessorRecover(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, be := NewTestBackend(t)
	defer os.RemoveAll(dir)
	defer be.Close()
	le := newLessor(be, minLeaseTTL)
	defer le.Stop()
	l1, err1 := le.Grant(1, 10)
	l2, err2 := le.Grant(2, 20)
	if err1 != nil || err2 != nil {
		t.Fatalf("could not grant initial leases (%v, %v)", err1, err2)
	}
	nle := newLessor(be, minLeaseTTL)
	defer nle.Stop()
	nl1 := nle.Lookup(l1.ID)
	if nl1 == nil || nl1.ttl != l1.ttl {
		t.Errorf("nl1 = %v, want nl1.ttl= %d", nl1.ttl, l1.ttl)
	}
	nl2 := nle.Lookup(l2.ID)
	if nl2 == nil || nl2.ttl != l2.ttl {
		t.Errorf("nl2 = %v, want nl2.ttl= %d", nl2.ttl, l2.ttl)
	}
}
func TestLessorExpire(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, be := NewTestBackend(t)
	defer os.RemoveAll(dir)
	defer be.Close()
	testMinTTL := int64(1)
	le := newLessor(be, testMinTTL)
	defer le.Stop()
	le.Promote(1 * time.Second)
	l, err := le.Grant(1, testMinTTL)
	if err != nil {
		t.Fatalf("failed to create lease: %v", err)
	}
	select {
	case el := <-le.ExpiredLeasesC():
		if el[0].ID != l.ID {
			t.Fatalf("expired id = %x, want %x", el[0].ID, l.ID)
		}
	case <-time.After(10 * time.Second):
		t.Fatalf("failed to receive expired lease")
	}
	donec := make(chan struct{})
	go func() {
		if _, err := le.Renew(l.ID); err != ErrLeaseNotFound {
			t.Fatalf("unexpected renew")
		}
		donec <- struct{}{}
	}()
	select {
	case <-donec:
		t.Fatalf("renew finished before lease revocation")
	case <-time.After(50 * time.Millisecond):
	}
	if err := le.Revoke(l.ID); err != nil {
		t.Fatalf("failed to revoke expired lease: %v", err)
	}
	select {
	case <-donec:
	case <-time.After(10 * time.Second):
		t.Fatalf("renew has not returned after lease revocation")
	}
}
func TestLessorExpireAndDemote(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, be := NewTestBackend(t)
	defer os.RemoveAll(dir)
	defer be.Close()
	testMinTTL := int64(1)
	le := newLessor(be, testMinTTL)
	defer le.Stop()
	le.Promote(1 * time.Second)
	l, err := le.Grant(1, testMinTTL)
	if err != nil {
		t.Fatalf("failed to create lease: %v", err)
	}
	select {
	case el := <-le.ExpiredLeasesC():
		if el[0].ID != l.ID {
			t.Fatalf("expired id = %x, want %x", el[0].ID, l.ID)
		}
	case <-time.After(10 * time.Second):
		t.Fatalf("failed to receive expired lease")
	}
	donec := make(chan struct{})
	go func() {
		if _, err := le.Renew(l.ID); err != ErrNotPrimary {
			t.Fatalf("unexpected renew: %v", err)
		}
		donec <- struct{}{}
	}()
	select {
	case <-donec:
		t.Fatalf("renew finished before demotion")
	case <-time.After(50 * time.Millisecond):
	}
	le.Demote()
	select {
	case <-donec:
	case <-time.After(10 * time.Second):
		t.Fatalf("renew has not returned after lessor demotion")
	}
}
func TestLessorMaxTTL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir, be := NewTestBackend(t)
	defer os.RemoveAll(dir)
	defer be.Close()
	le := newLessor(be, minLeaseTTL)
	defer le.Stop()
	_, err := le.Grant(1, MaxLeaseTTL+1)
	if err != ErrLeaseTTLTooLarge {
		t.Fatalf("grant unexpectedly succeeded")
	}
}

type fakeDeleter struct {
	deleted	[]string
	tx	backend.BatchTx
}

func newFakeDeleter(be backend.Backend) *fakeDeleter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fd := &fakeDeleter{nil, be.BatchTx()}
	fd.tx.Lock()
	return fd
}
func (fd *fakeDeleter) End() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fd.tx.Unlock()
}
func (fd *fakeDeleter) DeleteRange(key, end []byte) (int64, int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fd.deleted = append(fd.deleted, string(key)+"_"+string(end))
	return 0, 0
}
func NewTestBackend(t *testing.T) (string, backend.Backend) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tmpPath, err := ioutil.TempDir("", "lease")
	if err != nil {
		t.Fatalf("failed to create tmpdir (%v)", err)
	}
	bcfg := backend.DefaultBackendConfig()
	bcfg.Path = filepath.Join(tmpPath, "be")
	return tmpPath, backend.New(bcfg)
}
