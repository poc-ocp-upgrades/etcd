package v3rpc

import (
	"context"
	"io"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/lease"
)

type LeaseServer struct {
	hdr	header
	le	etcdserver.Lessor
}

func NewLeaseServer(s *etcdserver.EtcdServer) pb.LeaseServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &LeaseServer{le: s, hdr: newHeader(s)}
}
func (ls *LeaseServer) LeaseGrant(ctx context.Context, cr *pb.LeaseGrantRequest) (*pb.LeaseGrantResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := ls.le.LeaseGrant(ctx, cr)
	if err != nil {
		return nil, togRPCError(err)
	}
	ls.hdr.fill(resp.Header)
	return resp, nil
}
func (ls *LeaseServer) LeaseRevoke(ctx context.Context, rr *pb.LeaseRevokeRequest) (*pb.LeaseRevokeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := ls.le.LeaseRevoke(ctx, rr)
	if err != nil {
		return nil, togRPCError(err)
	}
	ls.hdr.fill(resp.Header)
	return resp, nil
}
func (ls *LeaseServer) LeaseTimeToLive(ctx context.Context, rr *pb.LeaseTimeToLiveRequest) (*pb.LeaseTimeToLiveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := ls.le.LeaseTimeToLive(ctx, rr)
	if err != nil && err != lease.ErrLeaseNotFound {
		return nil, togRPCError(err)
	}
	if err == lease.ErrLeaseNotFound {
		resp = &pb.LeaseTimeToLiveResponse{Header: &pb.ResponseHeader{}, ID: rr.ID, TTL: -1}
	}
	ls.hdr.fill(resp.Header)
	return resp, nil
}
func (ls *LeaseServer) LeaseLeases(ctx context.Context, rr *pb.LeaseLeasesRequest) (*pb.LeaseLeasesResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := ls.le.LeaseLeases(ctx, rr)
	if err != nil && err != lease.ErrLeaseNotFound {
		return nil, togRPCError(err)
	}
	if err == lease.ErrLeaseNotFound {
		resp = &pb.LeaseLeasesResponse{Header: &pb.ResponseHeader{}, Leases: []*pb.LeaseStatus{}}
	}
	ls.hdr.fill(resp.Header)
	return resp, nil
}
func (ls *LeaseServer) LeaseKeepAlive(stream pb.Lease_LeaseKeepAliveServer) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errc := make(chan error, 1)
	go func() {
		errc <- ls.leaseKeepAlive(stream)
	}()
	select {
	case err = <-errc:
	case <-stream.Context().Done():
		err = stream.Context().Err()
		if err == context.Canceled {
			err = rpctypes.ErrGRPCNoLeader
		}
	}
	return err
}
func (ls *LeaseServer) leaseKeepAlive(stream pb.Lease_LeaseKeepAliveServer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			if isClientCtxErr(stream.Context().Err(), err) {
				plog.Debugf("failed to receive lease keepalive request from gRPC stream (%q)", err.Error())
			} else {
				plog.Warningf("failed to receive lease keepalive request from gRPC stream (%q)", err.Error())
			}
			return err
		}
		resp := &pb.LeaseKeepAliveResponse{ID: req.ID, Header: &pb.ResponseHeader{}}
		ls.hdr.fill(resp.Header)
		ttl, err := ls.le.LeaseRenew(stream.Context(), lease.LeaseID(req.ID))
		if err == lease.ErrLeaseNotFound {
			err = nil
			ttl = 0
		}
		if err != nil {
			return togRPCError(err)
		}
		resp.TTL = ttl
		err = stream.Send(resp)
		if err != nil {
			if isClientCtxErr(stream.Context().Err(), err) {
				plog.Debugf("failed to send lease keepalive response to gRPC stream (%q)", err.Error())
			} else {
				plog.Warningf("failed to send lease keepalive response to gRPC stream (%q)", err.Error())
			}
			return err
		}
	}
}
