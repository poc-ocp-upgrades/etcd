package ordering

import (
	"context"
	"sync"
	"github.com/coreos/etcd/clientv3"
)

type kvOrdering struct {
	clientv3.KV
	orderViolationFunc	OrderViolationFunc
	prevRev			int64
	revMu			sync.RWMutex
}

func NewKV(kv clientv3.KV, orderViolationFunc OrderViolationFunc) *kvOrdering {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &kvOrdering{kv, orderViolationFunc, 0, sync.RWMutex{}}
}
func (kv *kvOrdering) getPrevRev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kv.revMu.RLock()
	defer kv.revMu.RUnlock()
	return kv.prevRev
}
func (kv *kvOrdering) setPrevRev(currRev int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kv.revMu.Lock()
	defer kv.revMu.Unlock()
	if currRev > kv.prevRev {
		kv.prevRev = currRev
	}
}
func (kv *kvOrdering) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prevRev := kv.getPrevRev()
	op := clientv3.OpGet(key, opts...)
	for {
		r, err := kv.KV.Do(ctx, op)
		if err != nil {
			return nil, err
		}
		resp := r.Get()
		if resp.Header.Revision == prevRev {
			return resp, nil
		} else if resp.Header.Revision > prevRev {
			kv.setPrevRev(resp.Header.Revision)
			return resp, nil
		}
		err = kv.orderViolationFunc(op, r, prevRev)
		if err != nil {
			return nil, err
		}
	}
}
func (kv *kvOrdering) Txn(ctx context.Context) clientv3.Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &txnOrdering{kv.KV.Txn(ctx), kv, ctx, sync.Mutex{}, []clientv3.Cmp{}, []clientv3.Op{}, []clientv3.Op{}}
}

type txnOrdering struct {
	clientv3.Txn
	*kvOrdering
	ctx	context.Context
	mu	sync.Mutex
	cmps	[]clientv3.Cmp
	thenOps	[]clientv3.Op
	elseOps	[]clientv3.Op
}

func (txn *txnOrdering) If(cs ...clientv3.Cmp) clientv3.Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.mu.Lock()
	defer txn.mu.Unlock()
	txn.cmps = cs
	txn.Txn.If(cs...)
	return txn
}
func (txn *txnOrdering) Then(ops ...clientv3.Op) clientv3.Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.mu.Lock()
	defer txn.mu.Unlock()
	txn.thenOps = ops
	txn.Txn.Then(ops...)
	return txn
}
func (txn *txnOrdering) Else(ops ...clientv3.Op) clientv3.Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.mu.Lock()
	defer txn.mu.Unlock()
	txn.elseOps = ops
	txn.Txn.Else(ops...)
	return txn
}
func (txn *txnOrdering) Commit() (*clientv3.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prevRev := txn.getPrevRev()
	opTxn := clientv3.OpTxn(txn.cmps, txn.thenOps, txn.elseOps)
	for {
		opResp, err := txn.KV.Do(txn.ctx, opTxn)
		if err != nil {
			return nil, err
		}
		txnResp := opResp.Txn()
		if txnResp.Header.Revision >= prevRev {
			txn.setPrevRev(txnResp.Header.Revision)
			return txnResp, nil
		}
		err = txn.orderViolationFunc(opTxn, opResp, prevRev)
		if err != nil {
			return nil, err
		}
	}
}
