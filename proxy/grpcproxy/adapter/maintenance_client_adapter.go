package adapter

import (
	"context"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
)

type mts2mtc struct{ mts pb.MaintenanceServer }

func MaintenanceServerToMaintenanceClient(mts pb.MaintenanceServer) pb.MaintenanceClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &mts2mtc{mts}
}
func (s *mts2mtc) Alarm(ctx context.Context, r *pb.AlarmRequest, opts ...grpc.CallOption) (*pb.AlarmResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.mts.Alarm(ctx, r)
}
func (s *mts2mtc) Status(ctx context.Context, r *pb.StatusRequest, opts ...grpc.CallOption) (*pb.StatusResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.mts.Status(ctx, r)
}
func (s *mts2mtc) Defragment(ctx context.Context, dr *pb.DefragmentRequest, opts ...grpc.CallOption) (*pb.DefragmentResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.mts.Defragment(ctx, dr)
}
func (s *mts2mtc) Hash(ctx context.Context, r *pb.HashRequest, opts ...grpc.CallOption) (*pb.HashResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.mts.Hash(ctx, r)
}
func (s *mts2mtc) HashKV(ctx context.Context, r *pb.HashKVRequest, opts ...grpc.CallOption) (*pb.HashKVResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.mts.HashKV(ctx, r)
}
func (s *mts2mtc) MoveLeader(ctx context.Context, r *pb.MoveLeaderRequest, opts ...grpc.CallOption) (*pb.MoveLeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.mts.MoveLeader(ctx, r)
}
func (s *mts2mtc) Snapshot(ctx context.Context, in *pb.SnapshotRequest, opts ...grpc.CallOption) (pb.Maintenance_SnapshotClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cs := newPipeStream(ctx, func(ss chanServerStream) error {
		return s.mts.Snapshot(in, &ss2scServerStream{ss})
	})
	return &ss2scClientStream{cs}, nil
}

type ss2scClientStream struct{ chanClientStream }
type ss2scServerStream struct{ chanServerStream }

func (s *ss2scClientStream) Send(rr *pb.SnapshotRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.SendMsg(rr)
}
func (s *ss2scClientStream) Recv() (*pb.SnapshotResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var v interface{}
	if err := s.RecvMsg(&v); err != nil {
		return nil, err
	}
	return v.(*pb.SnapshotResponse), nil
}
func (s *ss2scServerStream) Send(rr *pb.SnapshotResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.SendMsg(rr)
}
func (s *ss2scServerStream) Recv() (*pb.SnapshotRequest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var v interface{}
	if err := s.RecvMsg(&v); err != nil {
		return nil, err
	}
	return v.(*pb.SnapshotRequest), nil
}
