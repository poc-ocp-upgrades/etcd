package clientv3

import (
	"context"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
)

type (
	CompactResponse	pb.CompactionResponse
	PutResponse	pb.PutResponse
	GetResponse	pb.RangeResponse
	DeleteResponse	pb.DeleteRangeResponse
	TxnResponse	pb.TxnResponse
)
type KV interface {
	Put(ctx context.Context, key, val string, opts ...OpOption) (*PutResponse, error)
	Get(ctx context.Context, key string, opts ...OpOption) (*GetResponse, error)
	Delete(ctx context.Context, key string, opts ...OpOption) (*DeleteResponse, error)
	Compact(ctx context.Context, rev int64, opts ...CompactOption) (*CompactResponse, error)
	Do(ctx context.Context, op Op) (OpResponse, error)
	Txn(ctx context.Context) Txn
}
type OpResponse struct {
	put	*PutResponse
	get	*GetResponse
	del	*DeleteResponse
	txn	*TxnResponse
}

func (op OpResponse) Put() *PutResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.put
}
func (op OpResponse) Get() *GetResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.get
}
func (op OpResponse) Del() *DeleteResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.del
}
func (op OpResponse) Txn() *TxnResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.txn
}
func (resp *PutResponse) OpResponse() OpResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return OpResponse{put: resp}
}
func (resp *GetResponse) OpResponse() OpResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return OpResponse{get: resp}
}
func (resp *DeleteResponse) OpResponse() OpResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return OpResponse{del: resp}
}
func (resp *TxnResponse) OpResponse() OpResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return OpResponse{txn: resp}
}

type kv struct {
	remote		pb.KVClient
	callOpts	[]grpc.CallOption
}

func NewKV(c *Client) KV {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api := &kv{remote: RetryKVClient(c)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
func NewKVFromKVClient(remote pb.KVClient, c *Client) KV {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api := &kv{remote: remote}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
func (kv *kv) Put(ctx context.Context, key, val string, opts ...OpOption) (*PutResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r, err := kv.Do(ctx, OpPut(key, val, opts...))
	return r.put, toErr(ctx, err)
}
func (kv *kv) Get(ctx context.Context, key string, opts ...OpOption) (*GetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r, err := kv.Do(ctx, OpGet(key, opts...))
	return r.get, toErr(ctx, err)
}
func (kv *kv) Delete(ctx context.Context, key string, opts ...OpOption) (*DeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r, err := kv.Do(ctx, OpDelete(key, opts...))
	return r.del, toErr(ctx, err)
}
func (kv *kv) Compact(ctx context.Context, rev int64, opts ...CompactOption) (*CompactResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := kv.remote.Compact(ctx, OpCompact(rev, opts...).toRequest(), kv.callOpts...)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	return (*CompactResponse)(resp), err
}
func (kv *kv) Txn(ctx context.Context) Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &txn{kv: kv, ctx: ctx, callOpts: kv.callOpts}
}
func (kv *kv) Do(ctx context.Context, op Op) (OpResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	switch op.t {
	case tRange:
		var resp *pb.RangeResponse
		resp, err = kv.remote.Range(ctx, op.toRangeRequest(), kv.callOpts...)
		if err == nil {
			return OpResponse{get: (*GetResponse)(resp)}, nil
		}
	case tPut:
		var resp *pb.PutResponse
		r := &pb.PutRequest{Key: op.key, Value: op.val, Lease: int64(op.leaseID), PrevKv: op.prevKV, IgnoreValue: op.ignoreValue, IgnoreLease: op.ignoreLease}
		resp, err = kv.remote.Put(ctx, r, kv.callOpts...)
		if err == nil {
			return OpResponse{put: (*PutResponse)(resp)}, nil
		}
	case tDeleteRange:
		var resp *pb.DeleteRangeResponse
		r := &pb.DeleteRangeRequest{Key: op.key, RangeEnd: op.end, PrevKv: op.prevKV}
		resp, err = kv.remote.DeleteRange(ctx, r, kv.callOpts...)
		if err == nil {
			return OpResponse{del: (*DeleteResponse)(resp)}, nil
		}
	case tTxn:
		var resp *pb.TxnResponse
		resp, err = kv.remote.Txn(ctx, op.toTxnRequest(), kv.callOpts...)
		if err == nil {
			return OpResponse{txn: (*TxnResponse)(resp)}, nil
		}
	default:
		panic("Unknown op")
	}
	return OpResponse{}, toErr(ctx, err)
}
