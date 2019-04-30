package adapter

import (
	"context"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
)

type ls2lc struct{ leaseServer pb.LeaseServer }

func LeaseServerToLeaseClient(ls pb.LeaseServer) pb.LeaseClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ls2lc{ls}
}
func (c *ls2lc) LeaseGrant(ctx context.Context, in *pb.LeaseGrantRequest, opts ...grpc.CallOption) (*pb.LeaseGrantResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.leaseServer.LeaseGrant(ctx, in)
}
func (c *ls2lc) LeaseRevoke(ctx context.Context, in *pb.LeaseRevokeRequest, opts ...grpc.CallOption) (*pb.LeaseRevokeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.leaseServer.LeaseRevoke(ctx, in)
}
func (c *ls2lc) LeaseKeepAlive(ctx context.Context, opts ...grpc.CallOption) (pb.Lease_LeaseKeepAliveClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cs := newPipeStream(ctx, func(ss chanServerStream) error {
		return c.leaseServer.LeaseKeepAlive(&ls2lcServerStream{ss})
	})
	return &ls2lcClientStream{cs}, nil
}
func (c *ls2lc) LeaseTimeToLive(ctx context.Context, in *pb.LeaseTimeToLiveRequest, opts ...grpc.CallOption) (*pb.LeaseTimeToLiveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.leaseServer.LeaseTimeToLive(ctx, in)
}
func (c *ls2lc) LeaseLeases(ctx context.Context, in *pb.LeaseLeasesRequest, opts ...grpc.CallOption) (*pb.LeaseLeasesResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.leaseServer.LeaseLeases(ctx, in)
}

type ls2lcClientStream struct{ chanClientStream }
type ls2lcServerStream struct{ chanServerStream }

func (s *ls2lcClientStream) Send(rr *pb.LeaseKeepAliveRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.SendMsg(rr)
}
func (s *ls2lcClientStream) Recv() (*pb.LeaseKeepAliveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var v interface{}
	if err := s.RecvMsg(&v); err != nil {
		return nil, err
	}
	return v.(*pb.LeaseKeepAliveResponse), nil
}
func (s *ls2lcServerStream) Send(rr *pb.LeaseKeepAliveResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.SendMsg(rr)
}
func (s *ls2lcServerStream) Recv() (*pb.LeaseKeepAliveRequest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var v interface{}
	if err := s.RecvMsg(&v); err != nil {
		return nil, err
	}
	return v.(*pb.LeaseKeepAliveRequest), nil
}
