package clientv3

import (
	"context"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/types"
	"google.golang.org/grpc"
)

type (
	Member               pb.Member
	MemberListResponse   pb.MemberListResponse
	MemberAddResponse    pb.MemberAddResponse
	MemberRemoveResponse pb.MemberRemoveResponse
	MemberUpdateResponse pb.MemberUpdateResponse
)
type Cluster interface {
	MemberList(ctx context.Context) (*MemberListResponse, error)
	MemberAdd(ctx context.Context, peerAddrs []string) (*MemberAddResponse, error)
	MemberRemove(ctx context.Context, id uint64) (*MemberRemoveResponse, error)
	MemberUpdate(ctx context.Context, id uint64, peerAddrs []string) (*MemberUpdateResponse, error)
}
type cluster struct {
	remote   pb.ClusterClient
	callOpts []grpc.CallOption
}

func NewCluster(c *Client) Cluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api := &cluster{remote: RetryClusterClient(c)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
func NewClusterFromClusterClient(remote pb.ClusterClient, c *Client) Cluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api := &cluster{remote: remote}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
func (c *cluster) MemberAdd(ctx context.Context, peerAddrs []string) (*MemberAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, err := types.NewURLs(peerAddrs); err != nil {
		return nil, err
	}
	r := &pb.MemberAddRequest{PeerURLs: peerAddrs}
	resp, err := c.remote.MemberAdd(ctx, r, c.callOpts...)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	return (*MemberAddResponse)(resp), nil
}
func (c *cluster) MemberRemove(ctx context.Context, id uint64) (*MemberRemoveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := &pb.MemberRemoveRequest{ID: id}
	resp, err := c.remote.MemberRemove(ctx, r, c.callOpts...)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	return (*MemberRemoveResponse)(resp), nil
}
func (c *cluster) MemberUpdate(ctx context.Context, id uint64, peerAddrs []string) (*MemberUpdateResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, err := types.NewURLs(peerAddrs); err != nil {
		return nil, err
	}
	r := &pb.MemberUpdateRequest{ID: id, PeerURLs: peerAddrs}
	resp, err := c.remote.MemberUpdate(ctx, r, c.callOpts...)
	if err == nil {
		return (*MemberUpdateResponse)(resp), nil
	}
	return nil, toErr(ctx, err)
}
func (c *cluster) MemberList(ctx context.Context) (*MemberListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := c.remote.MemberList(ctx, &pb.MemberListRequest{}, c.callOpts...)
	if err == nil {
		return (*MemberListResponse)(resp), nil
	}
	return nil, toErr(ctx, err)
}
