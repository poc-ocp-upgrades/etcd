package namespace

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

type kvPrefix struct {
	clientv3.KV
	pfx	string
}

func NewKV(kv clientv3.KV, prefix string) clientv3.KV {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &kvPrefix{kv, prefix}
}
func (kv *kvPrefix) Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(key) == 0 {
		return nil, rpctypes.ErrEmptyKey
	}
	op := kv.prefixOp(clientv3.OpPut(key, val, opts...))
	r, err := kv.KV.Do(ctx, op)
	if err != nil {
		return nil, err
	}
	put := r.Put()
	kv.unprefixPutResponse(put)
	return put, nil
}
func (kv *kvPrefix) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(key) == 0 {
		return nil, rpctypes.ErrEmptyKey
	}
	r, err := kv.KV.Do(ctx, kv.prefixOp(clientv3.OpGet(key, opts...)))
	if err != nil {
		return nil, err
	}
	get := r.Get()
	kv.unprefixGetResponse(get)
	return get, nil
}
func (kv *kvPrefix) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(key) == 0 {
		return nil, rpctypes.ErrEmptyKey
	}
	r, err := kv.KV.Do(ctx, kv.prefixOp(clientv3.OpDelete(key, opts...)))
	if err != nil {
		return nil, err
	}
	del := r.Del()
	kv.unprefixDeleteResponse(del)
	return del, nil
}
func (kv *kvPrefix) Do(ctx context.Context, op clientv3.Op) (clientv3.OpResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(op.KeyBytes()) == 0 && !op.IsTxn() {
		return clientv3.OpResponse{}, rpctypes.ErrEmptyKey
	}
	r, err := kv.KV.Do(ctx, kv.prefixOp(op))
	if err != nil {
		return r, err
	}
	switch {
	case r.Get() != nil:
		kv.unprefixGetResponse(r.Get())
	case r.Put() != nil:
		kv.unprefixPutResponse(r.Put())
	case r.Del() != nil:
		kv.unprefixDeleteResponse(r.Del())
	case r.Txn() != nil:
		kv.unprefixTxnResponse(r.Txn())
	}
	return r, nil
}

type txnPrefix struct {
	clientv3.Txn
	kv	*kvPrefix
}

func (kv *kvPrefix) Txn(ctx context.Context) clientv3.Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &txnPrefix{kv.KV.Txn(ctx), kv}
}
func (txn *txnPrefix) If(cs ...clientv3.Cmp) clientv3.Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.Txn = txn.Txn.If(txn.kv.prefixCmps(cs)...)
	return txn
}
func (txn *txnPrefix) Then(ops ...clientv3.Op) clientv3.Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.Txn = txn.Txn.Then(txn.kv.prefixOps(ops)...)
	return txn
}
func (txn *txnPrefix) Else(ops ...clientv3.Op) clientv3.Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.Txn = txn.Txn.Else(txn.kv.prefixOps(ops)...)
	return txn
}
func (txn *txnPrefix) Commit() (*clientv3.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := txn.Txn.Commit()
	if err != nil {
		return nil, err
	}
	txn.kv.unprefixTxnResponse(resp)
	return resp, nil
}
func (kv *kvPrefix) prefixOp(op clientv3.Op) clientv3.Op {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !op.IsTxn() {
		begin, end := kv.prefixInterval(op.KeyBytes(), op.RangeBytes())
		op.WithKeyBytes(begin)
		op.WithRangeBytes(end)
		return op
	}
	cmps, thenOps, elseOps := op.Txn()
	return clientv3.OpTxn(kv.prefixCmps(cmps), kv.prefixOps(thenOps), kv.prefixOps(elseOps))
}
func (kv *kvPrefix) unprefixGetResponse(resp *clientv3.GetResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range resp.Kvs {
		resp.Kvs[i].Key = resp.Kvs[i].Key[len(kv.pfx):]
	}
}
func (kv *kvPrefix) unprefixPutResponse(resp *clientv3.PutResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if resp.PrevKv != nil {
		resp.PrevKv.Key = resp.PrevKv.Key[len(kv.pfx):]
	}
}
func (kv *kvPrefix) unprefixDeleteResponse(resp *clientv3.DeleteResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range resp.PrevKvs {
		resp.PrevKvs[i].Key = resp.PrevKvs[i].Key[len(kv.pfx):]
	}
}
func (kv *kvPrefix) unprefixTxnResponse(resp *clientv3.TxnResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, r := range resp.Responses {
		switch tv := r.Response.(type) {
		case *pb.ResponseOp_ResponseRange:
			if tv.ResponseRange != nil {
				kv.unprefixGetResponse((*clientv3.GetResponse)(tv.ResponseRange))
			}
		case *pb.ResponseOp_ResponsePut:
			if tv.ResponsePut != nil {
				kv.unprefixPutResponse((*clientv3.PutResponse)(tv.ResponsePut))
			}
		case *pb.ResponseOp_ResponseDeleteRange:
			if tv.ResponseDeleteRange != nil {
				kv.unprefixDeleteResponse((*clientv3.DeleteResponse)(tv.ResponseDeleteRange))
			}
		case *pb.ResponseOp_ResponseTxn:
			if tv.ResponseTxn != nil {
				kv.unprefixTxnResponse((*clientv3.TxnResponse)(tv.ResponseTxn))
			}
		default:
		}
	}
}
func (kv *kvPrefix) prefixInterval(key, end []byte) (pfxKey []byte, pfxEnd []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return prefixInterval(kv.pfx, key, end)
}
func (kv *kvPrefix) prefixCmps(cs []clientv3.Cmp) []clientv3.Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newCmps := make([]clientv3.Cmp, len(cs))
	for i := range cs {
		newCmps[i] = cs[i]
		pfxKey, endKey := kv.prefixInterval(cs[i].KeyBytes(), cs[i].RangeEnd)
		newCmps[i].WithKeyBytes(pfxKey)
		if len(cs[i].RangeEnd) != 0 {
			newCmps[i].RangeEnd = endKey
		}
	}
	return newCmps
}
func (kv *kvPrefix) prefixOps(ops []clientv3.Op) []clientv3.Op {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newOps := make([]clientv3.Op, len(ops))
	for i := range ops {
		newOps[i] = kv.prefixOp(ops[i])
	}
	return newOps
}
