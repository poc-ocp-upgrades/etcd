package adapter

import (
	"context"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	grpc "google.golang.org/grpc"
)

type kvs2kvc struct{ kvs pb.KVServer }

func KvServerToKvClient(kvs pb.KVServer) pb.KVClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &kvs2kvc{kvs}
}
func (s *kvs2kvc) Range(ctx context.Context, in *pb.RangeRequest, opts ...grpc.CallOption) (*pb.RangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.kvs.Range(ctx, in)
}
func (s *kvs2kvc) Put(ctx context.Context, in *pb.PutRequest, opts ...grpc.CallOption) (*pb.PutResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.kvs.Put(ctx, in)
}
func (s *kvs2kvc) DeleteRange(ctx context.Context, in *pb.DeleteRangeRequest, opts ...grpc.CallOption) (*pb.DeleteRangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.kvs.DeleteRange(ctx, in)
}
func (s *kvs2kvc) Txn(ctx context.Context, in *pb.TxnRequest, opts ...grpc.CallOption) (*pb.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.kvs.Txn(ctx, in)
}
func (s *kvs2kvc) Compact(ctx context.Context, in *pb.CompactionRequest, opts ...grpc.CallOption) (*pb.CompactionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.kvs.Compact(ctx, in)
}
