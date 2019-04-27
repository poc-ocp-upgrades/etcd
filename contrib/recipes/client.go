package recipe

import (
	"context"
	"errors"
	v3 "github.com/coreos/etcd/clientv3"
	spb "github.com/coreos/etcd/mvcc/mvccpb"
)

var (
	ErrKeyExists		= errors.New("key already exists")
	ErrWaitMismatch		= errors.New("unexpected wait result")
	ErrTooManyClients	= errors.New("too many clients")
	ErrNoWatcher		= errors.New("no watcher channel")
)

func deleteRevKey(kv v3.KV, key string, rev int64) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmp := v3.Compare(v3.ModRevision(key), "=", rev)
	req := v3.OpDelete(key)
	txnresp, err := kv.Txn(context.TODO()).If(cmp).Then(req).Commit()
	if err != nil {
		return false, err
	} else if !txnresp.Succeeded {
		return false, nil
	}
	return true, nil
}
func claimFirstKey(kv v3.KV, kvs []*spb.KeyValue) (*spb.KeyValue, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, k := range kvs {
		ok, err := deleteRevKey(kv, string(k.Key), k.ModRevision)
		if err != nil {
			return nil, err
		} else if ok {
			return k, nil
		}
	}
	return nil, nil
}
