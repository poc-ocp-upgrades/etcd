package clientv3

import (
	"context"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
	"io"
)

type (
	DefragmentResponse pb.DefragmentResponse
	AlarmResponse      pb.AlarmResponse
	AlarmMember        pb.AlarmMember
	StatusResponse     pb.StatusResponse
	HashKVResponse     pb.HashKVResponse
	MoveLeaderResponse pb.MoveLeaderResponse
)
type Maintenance interface {
	AlarmList(ctx context.Context) (*AlarmResponse, error)
	AlarmDisarm(ctx context.Context, m *AlarmMember) (*AlarmResponse, error)
	Defragment(ctx context.Context, endpoint string) (*DefragmentResponse, error)
	Status(ctx context.Context, endpoint string) (*StatusResponse, error)
	HashKV(ctx context.Context, endpoint string, rev int64) (*HashKVResponse, error)
	Snapshot(ctx context.Context) (io.ReadCloser, error)
	MoveLeader(ctx context.Context, transfereeID uint64) (*MoveLeaderResponse, error)
}
type maintenance struct {
	dial     func(endpoint string) (pb.MaintenanceClient, func(), error)
	remote   pb.MaintenanceClient
	callOpts []grpc.CallOption
}

func NewMaintenance(c *Client) Maintenance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api := &maintenance{dial: func(endpoint string) (pb.MaintenanceClient, func(), error) {
		conn, err := c.dial(endpoint)
		if err != nil {
			return nil, nil, err
		}
		cancel := func() {
			conn.Close()
		}
		return RetryMaintenanceClient(c, conn), cancel, nil
	}, remote: RetryMaintenanceClient(c, c.conn)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
func NewMaintenanceFromMaintenanceClient(remote pb.MaintenanceClient, c *Client) Maintenance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api := &maintenance{dial: func(string) (pb.MaintenanceClient, func(), error) {
		return remote, func() {
		}, nil
	}, remote: remote}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
func (m *maintenance) AlarmList(ctx context.Context) (*AlarmResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := &pb.AlarmRequest{Action: pb.AlarmRequest_GET, MemberID: 0, Alarm: pb.AlarmType_NONE}
	resp, err := m.remote.Alarm(ctx, req, m.callOpts...)
	if err == nil {
		return (*AlarmResponse)(resp), nil
	}
	return nil, toErr(ctx, err)
}
func (m *maintenance) AlarmDisarm(ctx context.Context, am *AlarmMember) (*AlarmResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := &pb.AlarmRequest{Action: pb.AlarmRequest_DEACTIVATE, MemberID: am.MemberID, Alarm: am.Alarm}
	if req.MemberID == 0 && req.Alarm == pb.AlarmType_NONE {
		ar, err := m.AlarmList(ctx)
		if err != nil {
			return nil, toErr(ctx, err)
		}
		ret := AlarmResponse{}
		for _, am := range ar.Alarms {
			dresp, derr := m.AlarmDisarm(ctx, (*AlarmMember)(am))
			if derr != nil {
				return nil, toErr(ctx, derr)
			}
			ret.Alarms = append(ret.Alarms, dresp.Alarms...)
		}
		return &ret, nil
	}
	resp, err := m.remote.Alarm(ctx, req, m.callOpts...)
	if err == nil {
		return (*AlarmResponse)(resp), nil
	}
	return nil, toErr(ctx, err)
}
func (m *maintenance) Defragment(ctx context.Context, endpoint string) (*DefragmentResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	remote, cancel, err := m.dial(endpoint)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	defer cancel()
	resp, err := remote.Defragment(ctx, &pb.DefragmentRequest{}, m.callOpts...)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	return (*DefragmentResponse)(resp), nil
}
func (m *maintenance) Status(ctx context.Context, endpoint string) (*StatusResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	remote, cancel, err := m.dial(endpoint)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	defer cancel()
	resp, err := remote.Status(ctx, &pb.StatusRequest{}, m.callOpts...)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	return (*StatusResponse)(resp), nil
}
func (m *maintenance) HashKV(ctx context.Context, endpoint string, rev int64) (*HashKVResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	remote, cancel, err := m.dial(endpoint)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	defer cancel()
	resp, err := remote.HashKV(ctx, &pb.HashKVRequest{Revision: rev}, m.callOpts...)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	return (*HashKVResponse)(resp), nil
}
func (m *maintenance) Snapshot(ctx context.Context) (io.ReadCloser, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss, err := m.remote.Snapshot(ctx, &pb.SnapshotRequest{}, m.callOpts...)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	pr, pw := io.Pipe()
	go func() {
		for {
			resp, err := ss.Recv()
			if err != nil {
				pw.CloseWithError(err)
				return
			}
			if resp == nil && err == nil {
				break
			}
			if _, werr := pw.Write(resp.Blob); werr != nil {
				pw.CloseWithError(werr)
				return
			}
		}
		pw.Close()
	}()
	return &snapshotReadCloser{ctx: ctx, ReadCloser: pr}, nil
}

type snapshotReadCloser struct {
	ctx context.Context
	io.ReadCloser
}

func (rc *snapshotReadCloser) Read(p []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n, err = rc.ReadCloser.Read(p)
	return n, toErr(rc.ctx, err)
}
func (m *maintenance) MoveLeader(ctx context.Context, transfereeID uint64) (*MoveLeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := m.remote.MoveLeader(ctx, &pb.MoveLeaderRequest{TargetID: transfereeID}, m.callOpts...)
	return (*MoveLeaderResponse)(resp), toErr(ctx, err)
}
