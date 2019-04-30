package v2v3

import (
	"context"
	"net/http"
	"time"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api"
	"go.etcd.io/etcd/etcdserver/api/membership"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/pkg/types"
	"github.com/coreos/go-semver/semver"
	"go.uber.org/zap"
)

type fakeStats struct{}

func (s *fakeStats) SelfStats() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *fakeStats) LeaderStats() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *fakeStats) StoreStats() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

type v2v3Server struct {
	lg	*zap.Logger
	c	*clientv3.Client
	store	*v2v3Store
	fakeStats
}

func NewServer(lg *zap.Logger, c *clientv3.Client, pfx string) etcdserver.ServerPeer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &v2v3Server{lg: lg, c: c, store: newStore(c, pfx)}
}
func (s *v2v3Server) ClientCertAuthEnabled() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (s *v2v3Server) LeaseHandler() http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("STUB: lease handler")
}
func (s *v2v3Server) RaftHandler() http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("STUB: raft handler")
}
func (s *v2v3Server) Leader() types.ID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	resp, err := s.c.Status(ctx, s.c.Endpoints()[0])
	if err != nil {
		return 0
	}
	return types.ID(resp.Leader)
}
func (s *v2v3Server) AddMember(ctx context.Context, memb membership.Member) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := s.c.MemberAdd(ctx, memb.PeerURLs)
	if err != nil {
		return nil, err
	}
	return v3MembersToMembership(resp.Members), nil
}
func (s *v2v3Server) RemoveMember(ctx context.Context, id uint64) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := s.c.MemberRemove(ctx, id)
	if err != nil {
		return nil, err
	}
	return v3MembersToMembership(resp.Members), nil
}
func (s *v2v3Server) UpdateMember(ctx context.Context, m membership.Member) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := s.c.MemberUpdate(ctx, uint64(m.ID), m.PeerURLs)
	if err != nil {
		return nil, err
	}
	return v3MembersToMembership(resp.Members), nil
}
func v3MembersToMembership(v3membs []*pb.Member) []*membership.Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	membs := make([]*membership.Member, len(v3membs))
	for i, m := range v3membs {
		membs[i] = &membership.Member{ID: types.ID(m.ID), RaftAttributes: membership.RaftAttributes{PeerURLs: m.PeerURLs}, Attributes: membership.Attributes{Name: m.Name, ClientURLs: m.ClientURLs}}
	}
	return membs
}
func (s *v2v3Server) ClusterVersion() *semver.Version {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.Version()
}
func (s *v2v3Server) Cluster() api.Cluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s
}
func (s *v2v3Server) Alarms() []*pb.AlarmMember {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *v2v3Server) Do(ctx context.Context, r pb.Request) (etcdserver.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	applier := etcdserver.NewApplierV2(s.lg, s.store, nil)
	reqHandler := etcdserver.NewStoreRequestV2Handler(s.store, applier)
	req := (*etcdserver.RequestV2)(&r)
	resp, err := req.Handle(ctx, reqHandler)
	if resp.Err != nil {
		return resp, resp.Err
	}
	return resp, err
}
