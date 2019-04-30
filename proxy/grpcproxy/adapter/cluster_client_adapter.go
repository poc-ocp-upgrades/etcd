package adapter

import (
	"context"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
)

type cls2clc struct{ cls pb.ClusterServer }

func ClusterServerToClusterClient(cls pb.ClusterServer) pb.ClusterClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cls2clc{cls}
}
func (s *cls2clc) MemberList(ctx context.Context, r *pb.MemberListRequest, opts ...grpc.CallOption) (*pb.MemberListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.cls.MemberList(ctx, r)
}
func (s *cls2clc) MemberAdd(ctx context.Context, r *pb.MemberAddRequest, opts ...grpc.CallOption) (*pb.MemberAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.cls.MemberAdd(ctx, r)
}
func (s *cls2clc) MemberUpdate(ctx context.Context, r *pb.MemberUpdateRequest, opts ...grpc.CallOption) (*pb.MemberUpdateResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.cls.MemberUpdate(ctx, r)
}
func (s *cls2clc) MemberRemove(ctx context.Context, r *pb.MemberRemoveRequest, opts ...grpc.CallOption) (*pb.MemberRemoveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.cls.MemberRemove(ctx, r)
}
