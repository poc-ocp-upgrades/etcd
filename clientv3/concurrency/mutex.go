package concurrency

import (
	"context"
	"fmt"
	"sync"
	v3 "github.com/coreos/etcd/clientv3"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

type Mutex struct {
	s	*Session
	pfx	string
	myKey	string
	myRev	int64
	hdr	*pb.ResponseHeader
}

func NewMutex(s *Session, pfx string) *Mutex {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Mutex{s, pfx + "/", "", -1, nil}
}
func (m *Mutex) Lock(ctx context.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := m.s
	client := m.s.Client()
	m.myKey = fmt.Sprintf("%s%x", m.pfx, s.Lease())
	cmp := v3.Compare(v3.CreateRevision(m.myKey), "=", 0)
	put := v3.OpPut(m.myKey, "", v3.WithLease(s.Lease()))
	get := v3.OpGet(m.myKey)
	getOwner := v3.OpGet(m.pfx, v3.WithFirstCreate()...)
	resp, err := client.Txn(ctx).If(cmp).Then(put, getOwner).Else(get, getOwner).Commit()
	if err != nil {
		return err
	}
	m.myRev = resp.Header.Revision
	if !resp.Succeeded {
		m.myRev = resp.Responses[0].GetResponseRange().Kvs[0].CreateRevision
	}
	ownerKey := resp.Responses[1].GetResponseRange().Kvs
	if len(ownerKey) == 0 || ownerKey[0].CreateRevision == m.myRev {
		m.hdr = resp.Header
		return nil
	}
	hdr, werr := waitDeletes(ctx, client, m.pfx, m.myRev-1)
	if werr != nil {
		m.Unlock(client.Ctx())
	} else {
		m.hdr = hdr
	}
	return werr
}
func (m *Mutex) Unlock(ctx context.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := m.s.Client()
	if _, err := client.Delete(ctx, m.myKey); err != nil {
		return err
	}
	m.myKey = "\x00"
	m.myRev = -1
	return nil
}
func (m *Mutex) IsOwner() v3.Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v3.Compare(v3.CreateRevision(m.myKey), "=", m.myRev)
}
func (m *Mutex) Key() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.myKey
}
func (m *Mutex) Header() *pb.ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.hdr
}

type lockerMutex struct{ *Mutex }

func (lm *lockerMutex) Lock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := lm.s.Client()
	if err := lm.Mutex.Lock(client.Ctx()); err != nil {
		panic(err)
	}
}
func (lm *lockerMutex) Unlock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := lm.s.Client()
	if err := lm.Mutex.Unlock(client.Ctx()); err != nil {
		panic(err)
	}
}
func NewLocker(s *Session, pfx string) sync.Locker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &lockerMutex{NewMutex(s, pfx)}
}
