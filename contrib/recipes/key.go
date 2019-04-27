package recipe

import (
	"context"
	"fmt"
	"strings"
	"time"
	v3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

type RemoteKV struct {
	kv	v3.KV
	key	string
	rev	int64
	val	string
}

func newKey(kv v3.KV, key string, leaseID v3.LeaseID) (*RemoteKV, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newKV(kv, key, "", leaseID)
}
func newKV(kv v3.KV, key, val string, leaseID v3.LeaseID) (*RemoteKV, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rev, err := putNewKV(kv, key, val, leaseID)
	if err != nil {
		return nil, err
	}
	return &RemoteKV{kv, key, rev, val}, nil
}
func newUniqueKV(kv v3.KV, prefix string, val string) (*RemoteKV, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		newKey := fmt.Sprintf("%s/%v", prefix, time.Now().UnixNano())
		rev, err := putNewKV(kv, newKey, val, 0)
		if err == nil {
			return &RemoteKV{kv, newKey, rev, val}, nil
		}
		if err != ErrKeyExists {
			return nil, err
		}
	}
}
func putNewKV(kv v3.KV, key, val string, leaseID v3.LeaseID) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmp := v3.Compare(v3.Version(key), "=", 0)
	req := v3.OpPut(key, val, v3.WithLease(leaseID))
	txnresp, err := kv.Txn(context.TODO()).If(cmp).Then(req).Commit()
	if err != nil {
		return 0, err
	}
	if !txnresp.Succeeded {
		return 0, ErrKeyExists
	}
	return txnresp.Header.Revision, nil
}
func newSequentialKV(kv v3.KV, prefix, val string) (*RemoteKV, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := kv.Get(context.TODO(), prefix, v3.WithLastKey()...)
	if err != nil {
		return nil, err
	}
	newSeqNum := 0
	if len(resp.Kvs) != 0 {
		fields := strings.Split(string(resp.Kvs[0].Key), "/")
		_, serr := fmt.Sscanf(fields[len(fields)-1], "%d", &newSeqNum)
		if serr != nil {
			return nil, serr
		}
		newSeqNum++
	}
	newKey := fmt.Sprintf("%s/%016d", prefix, newSeqNum)
	baseKey := "__" + prefix
	cmp := v3.Compare(v3.ModRevision(baseKey), "<", resp.Header.Revision+1)
	reqPrefix := v3.OpPut(baseKey, "")
	reqnewKey := v3.OpPut(newKey, val)
	txn := kv.Txn(context.TODO())
	txnresp, err := txn.If(cmp).Then(reqPrefix, reqnewKey).Commit()
	if err != nil {
		return nil, err
	}
	if !txnresp.Succeeded {
		return newSequentialKV(kv, prefix, val)
	}
	return &RemoteKV{kv, newKey, txnresp.Header.Revision, val}, nil
}
func (rk *RemoteKV) Key() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rk.key
}
func (rk *RemoteKV) Revision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rk.rev
}
func (rk *RemoteKV) Value() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rk.val
}
func (rk *RemoteKV) Delete() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rk.kv == nil {
		return nil
	}
	_, err := rk.kv.Delete(context.TODO(), rk.key)
	rk.kv = nil
	return err
}
func (rk *RemoteKV) Put(val string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := rk.kv.Put(context.TODO(), rk.key, val)
	return err
}

type EphemeralKV struct{ RemoteKV }

func newEphemeralKV(s *concurrency.Session, key, val string) (*EphemeralKV, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	k, err := newKV(s.Client(), key, val, s.Lease())
	if err != nil {
		return nil, err
	}
	return &EphemeralKV{*k}, nil
}
func newUniqueEphemeralKey(s *concurrency.Session, prefix string) (*EphemeralKV, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newUniqueEphemeralKV(s, prefix, "")
}
func newUniqueEphemeralKV(s *concurrency.Session, prefix, val string) (ek *EphemeralKV, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		newKey := fmt.Sprintf("%s/%v", prefix, time.Now().UnixNano())
		ek, err = newEphemeralKV(s, newKey, val)
		if err == nil || err != ErrKeyExists {
			break
		}
	}
	return ek, err
}
