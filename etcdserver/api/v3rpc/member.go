package v3rpc

import (
	"context"
	"time"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api"
	"go.etcd.io/etcd/etcdserver/api/membership"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/pkg/types"
)

type ClusterServer struct {
	cluster	api.Cluster
	server	etcdserver.ServerV3
}

func NewClusterServer(s etcdserver.ServerV3) *ClusterServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ClusterServer{cluster: s.Cluster(), server: s}
}
func (cs *ClusterServer) MemberAdd(ctx context.Context, r *pb.MemberAddRequest) (*pb.MemberAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	urls, err := types.NewURLs(r.PeerURLs)
	if err != nil {
		return nil, rpctypes.ErrGRPCMemberBadURLs
	}
	now := time.Now()
	m := membership.NewMember("", urls, "", &now)
	membs, merr := cs.server.AddMember(ctx, *m)
	if merr != nil {
		return nil, togRPCError(merr)
	}
	return &pb.MemberAddResponse{Header: cs.header(), Member: &pb.Member{ID: uint64(m.ID), PeerURLs: m.PeerURLs}, Members: membersToProtoMembers(membs)}, nil
}
func (cs *ClusterServer) MemberRemove(ctx context.Context, r *pb.MemberRemoveRequest) (*pb.MemberRemoveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	membs, err := cs.server.RemoveMember(ctx, r.ID)
	if err != nil {
		return nil, togRPCError(err)
	}
	return &pb.MemberRemoveResponse{Header: cs.header(), Members: membersToProtoMembers(membs)}, nil
}
func (cs *ClusterServer) MemberUpdate(ctx context.Context, r *pb.MemberUpdateRequest) (*pb.MemberUpdateResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := membership.Member{ID: types.ID(r.ID), RaftAttributes: membership.RaftAttributes{PeerURLs: r.PeerURLs}}
	membs, err := cs.server.UpdateMember(ctx, m)
	if err != nil {
		return nil, togRPCError(err)
	}
	return &pb.MemberUpdateResponse{Header: cs.header(), Members: membersToProtoMembers(membs)}, nil
}
func (cs *ClusterServer) MemberList(ctx context.Context, r *pb.MemberListRequest) (*pb.MemberListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	membs := membersToProtoMembers(cs.cluster.Members())
	return &pb.MemberListResponse{Header: cs.header(), Members: membs}, nil
}
func (cs *ClusterServer) header() *pb.ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pb.ResponseHeader{ClusterId: uint64(cs.cluster.ID()), MemberId: uint64(cs.server.ID()), RaftTerm: cs.server.Term()}
}
func membersToProtoMembers(membs []*membership.Member) []*pb.Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	protoMembs := make([]*pb.Member, len(membs))
	for i := range membs {
		protoMembs[i] = &pb.Member{Name: membs[i].Name, ID: uint64(membs[i].ID), PeerURLs: membs[i].PeerURLs, ClientURLs: membs[i].ClientURLs}
	}
	return protoMembs
}
