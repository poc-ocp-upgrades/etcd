package v3rpc

import (
	"context"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/types"
)

type quotaKVServer struct {
	pb.KVServer
	qa quotaAlarmer
}
type quotaAlarmer struct {
	q  etcdserver.Quota
	a  Alarmer
	id types.ID
}

func (qa *quotaAlarmer) check(ctx context.Context, r interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if qa.q.Available(r) {
		return nil
	}
	req := &pb.AlarmRequest{MemberID: uint64(qa.id), Action: pb.AlarmRequest_ACTIVATE, Alarm: pb.AlarmType_NOSPACE}
	qa.a.Alarm(ctx, req)
	return rpctypes.ErrGRPCNoSpace
}
func NewQuotaKVServer(s *etcdserver.EtcdServer) pb.KVServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &quotaKVServer{NewKVServer(s), quotaAlarmer{etcdserver.NewBackendQuota(s), s, s.ID()}}
}
func (s *quotaKVServer) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.qa.check(ctx, r); err != nil {
		return nil, err
	}
	return s.KVServer.Put(ctx, r)
}
func (s *quotaKVServer) Txn(ctx context.Context, r *pb.TxnRequest) (*pb.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.qa.check(ctx, r); err != nil {
		return nil, err
	}
	return s.KVServer.Txn(ctx, r)
}

type quotaLeaseServer struct {
	pb.LeaseServer
	qa quotaAlarmer
}

func (s *quotaLeaseServer) LeaseGrant(ctx context.Context, cr *pb.LeaseGrantRequest) (*pb.LeaseGrantResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.qa.check(ctx, cr); err != nil {
		return nil, err
	}
	return s.LeaseServer.LeaseGrant(ctx, cr)
}
func NewQuotaLeaseServer(s *etcdserver.EtcdServer) pb.LeaseServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &quotaLeaseServer{NewLeaseServer(s), quotaAlarmer{etcdserver.NewBackendQuota(s), s, s.ID()}}
}
