package adapter

import (
	"context"
	"github.com/coreos/etcd/etcdserver/api/v3lock/v3lockpb"
	"google.golang.org/grpc"
)

type ls2lsc struct{ ls v3lockpb.LockServer }

func LockServerToLockClient(ls v3lockpb.LockServer) v3lockpb.LockClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ls2lsc{ls}
}
func (s *ls2lsc) Lock(ctx context.Context, r *v3lockpb.LockRequest, opts ...grpc.CallOption) (*v3lockpb.LockResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.ls.Lock(ctx, r)
}
func (s *ls2lsc) Unlock(ctx context.Context, r *v3lockpb.UnlockRequest, opts ...grpc.CallOption) (*v3lockpb.UnlockResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.ls.Unlock(ctx, r)
}
