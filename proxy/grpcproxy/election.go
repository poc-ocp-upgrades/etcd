package grpcproxy

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3election/v3electionpb"
)

type electionProxy struct{ client *clientv3.Client }

func NewElectionProxy(client *clientv3.Client) v3electionpb.ElectionServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &electionProxy{client: client}
}
func (ep *electionProxy) Campaign(ctx context.Context, req *v3electionpb.CampaignRequest) (*v3electionpb.CampaignResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v3electionpb.NewElectionClient(ep.client.ActiveConnection()).Campaign(ctx, req)
}
func (ep *electionProxy) Proclaim(ctx context.Context, req *v3electionpb.ProclaimRequest) (*v3electionpb.ProclaimResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v3electionpb.NewElectionClient(ep.client.ActiveConnection()).Proclaim(ctx, req)
}
func (ep *electionProxy) Leader(ctx context.Context, req *v3electionpb.LeaderRequest) (*v3electionpb.LeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v3electionpb.NewElectionClient(ep.client.ActiveConnection()).Leader(ctx, req)
}
func (ep *electionProxy) Observe(req *v3electionpb.LeaderRequest, s v3electionpb.Election_ObserveServer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ep.client.ActiveConnection()
	ctx, cancel := context.WithCancel(s.Context())
	defer cancel()
	sc, err := v3electionpb.NewElectionClient(conn).Observe(ctx, req)
	if err != nil {
		return err
	}
	for {
		rr, err := sc.Recv()
		if err != nil {
			return err
		}
		if err = s.Send(rr); err != nil {
			return err
		}
	}
}
func (ep *electionProxy) Resign(ctx context.Context, req *v3electionpb.ResignRequest) (*v3electionpb.ResignResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return v3electionpb.NewElectionClient(ep.client.ActiveConnection()).Resign(ctx, req)
}
