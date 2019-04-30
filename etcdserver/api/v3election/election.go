package v3election

import (
	"context"
	"errors"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	epb "go.etcd.io/etcd/etcdserver/api/v3election/v3electionpb"
)

var ErrMissingLeaderKey = errors.New(`"leader" field must be provided`)

type electionServer struct{ c *clientv3.Client }

func NewElectionServer(c *clientv3.Client) epb.ElectionServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &electionServer{c}
}
func (es *electionServer) Campaign(ctx context.Context, req *epb.CampaignRequest) (*epb.CampaignResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := es.session(ctx, req.Lease)
	if err != nil {
		return nil, err
	}
	e := concurrency.NewElection(s, string(req.Name))
	if err = e.Campaign(ctx, string(req.Value)); err != nil {
		return nil, err
	}
	return &epb.CampaignResponse{Header: e.Header(), Leader: &epb.LeaderKey{Name: req.Name, Key: []byte(e.Key()), Rev: e.Rev(), Lease: int64(s.Lease())}}, nil
}
func (es *electionServer) Proclaim(ctx context.Context, req *epb.ProclaimRequest) (*epb.ProclaimResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if req.Leader == nil {
		return nil, ErrMissingLeaderKey
	}
	s, err := es.session(ctx, req.Leader.Lease)
	if err != nil {
		return nil, err
	}
	e := concurrency.ResumeElection(s, string(req.Leader.Name), string(req.Leader.Key), req.Leader.Rev)
	if err := e.Proclaim(ctx, string(req.Value)); err != nil {
		return nil, err
	}
	return &epb.ProclaimResponse{Header: e.Header()}, nil
}
func (es *electionServer) Observe(req *epb.LeaderRequest, stream epb.Election_ObserveServer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := es.session(stream.Context(), -1)
	if err != nil {
		return err
	}
	e := concurrency.NewElection(s, string(req.Name))
	ch := e.Observe(stream.Context())
	for stream.Context().Err() == nil {
		select {
		case <-stream.Context().Done():
		case resp, ok := <-ch:
			if !ok {
				return nil
			}
			lresp := &epb.LeaderResponse{Header: resp.Header, Kv: resp.Kvs[0]}
			if err := stream.Send(lresp); err != nil {
				return err
			}
		}
	}
	return stream.Context().Err()
}
func (es *electionServer) Leader(ctx context.Context, req *epb.LeaderRequest) (*epb.LeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := es.session(ctx, -1)
	if err != nil {
		return nil, err
	}
	l, lerr := concurrency.NewElection(s, string(req.Name)).Leader(ctx)
	if lerr != nil {
		return nil, lerr
	}
	return &epb.LeaderResponse{Header: l.Header, Kv: l.Kvs[0]}, nil
}
func (es *electionServer) Resign(ctx context.Context, req *epb.ResignRequest) (*epb.ResignResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if req.Leader == nil {
		return nil, ErrMissingLeaderKey
	}
	s, err := es.session(ctx, req.Leader.Lease)
	if err != nil {
		return nil, err
	}
	e := concurrency.ResumeElection(s, string(req.Leader.Name), string(req.Leader.Key), req.Leader.Rev)
	if err := e.Resign(ctx); err != nil {
		return nil, err
	}
	return &epb.ResignResponse{Header: e.Header()}, nil
}
func (es *electionServer) session(ctx context.Context, lease int64) (*concurrency.Session, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := concurrency.NewSession(es.c, concurrency.WithLease(clientv3.LeaseID(lease)), concurrency.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	s.Orphan()
	return s, nil
}
