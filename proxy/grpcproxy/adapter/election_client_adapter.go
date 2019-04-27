package adapter

import (
	"context"
	"github.com/coreos/etcd/etcdserver/api/v3election/v3electionpb"
	"google.golang.org/grpc"
)

type es2ec struct{ es v3electionpb.ElectionServer }

func ElectionServerToElectionClient(es v3electionpb.ElectionServer) v3electionpb.ElectionClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &es2ec{es}
}
func (s *es2ec) Campaign(ctx context.Context, r *v3electionpb.CampaignRequest, opts ...grpc.CallOption) (*v3electionpb.CampaignResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.es.Campaign(ctx, r)
}
func (s *es2ec) Proclaim(ctx context.Context, r *v3electionpb.ProclaimRequest, opts ...grpc.CallOption) (*v3electionpb.ProclaimResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.es.Proclaim(ctx, r)
}
func (s *es2ec) Leader(ctx context.Context, r *v3electionpb.LeaderRequest, opts ...grpc.CallOption) (*v3electionpb.LeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.es.Leader(ctx, r)
}
func (s *es2ec) Resign(ctx context.Context, r *v3electionpb.ResignRequest, opts ...grpc.CallOption) (*v3electionpb.ResignResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.es.Resign(ctx, r)
}
func (s *es2ec) Observe(ctx context.Context, in *v3electionpb.LeaderRequest, opts ...grpc.CallOption) (v3electionpb.Election_ObserveClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cs := newPipeStream(ctx, func(ss chanServerStream) error {
		return s.es.Observe(in, &es2ecServerStream{ss})
	})
	return &es2ecClientStream{cs}, nil
}

type es2ecClientStream struct{ chanClientStream }
type es2ecServerStream struct{ chanServerStream }

func (s *es2ecClientStream) Send(rr *v3electionpb.LeaderRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.SendMsg(rr)
}
func (s *es2ecClientStream) Recv() (*v3electionpb.LeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var v interface{}
	if err := s.RecvMsg(&v); err != nil {
		return nil, err
	}
	return v.(*v3electionpb.LeaderResponse), nil
}
func (s *es2ecServerStream) Send(rr *v3electionpb.LeaderResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.SendMsg(rr)
}
func (s *es2ecServerStream) Recv() (*v3electionpb.LeaderRequest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var v interface{}
	if err := s.RecvMsg(&v); err != nil {
		return nil, err
	}
	return v.(*v3electionpb.LeaderRequest), nil
}
