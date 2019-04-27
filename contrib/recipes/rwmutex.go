package recipe

import (
	"context"
	v3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type RWMutex struct {
	s	*concurrency.Session
	ctx	context.Context
	pfx	string
	myKey	*EphemeralKV
}

func NewRWMutex(s *concurrency.Session, prefix string) *RWMutex {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &RWMutex{s, context.TODO(), prefix + "/", nil}
}
func (rwm *RWMutex) RLock() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rk, err := newUniqueEphemeralKey(rwm.s, rwm.pfx+"read")
	if err != nil {
		return err
	}
	rwm.myKey = rk
	for {
		if done, werr := rwm.waitOnLastRev(rwm.pfx + "write"); done || werr != nil {
			return werr
		}
	}
}
func (rwm *RWMutex) Lock() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rk, err := newUniqueEphemeralKey(rwm.s, rwm.pfx+"write")
	if err != nil {
		return err
	}
	rwm.myKey = rk
	for {
		if done, werr := rwm.waitOnLastRev(rwm.pfx); done || werr != nil {
			return werr
		}
	}
}
func (rwm *RWMutex) waitOnLastRev(pfx string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := rwm.s.Client()
	opts := append(v3.WithLastRev(), v3.WithMaxModRev(rwm.myKey.Revision()-1))
	lastKey, err := client.Get(rwm.ctx, pfx, opts...)
	if err != nil {
		return false, err
	}
	if len(lastKey.Kvs) == 0 {
		return true, nil
	}
	_, err = WaitEvents(client, string(lastKey.Kvs[0].Key), rwm.myKey.Revision(), []mvccpb.Event_EventType{mvccpb.DELETE})
	return false, err
}
func (rwm *RWMutex) RUnlock() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rwm.myKey.Delete()
}
func (rwm *RWMutex) Unlock() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rwm.myKey.Delete()
}
