package clientv3

import (
	"context"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type retryPolicy uint8

const (
	repeatable	retryPolicy	= iota
	nonRepeatable
)

type rpcFunc func(ctx context.Context) error
type retryRPCFunc func(context.Context, rpcFunc, retryPolicy) error
type retryStopErrFunc func(error) bool

func isRepeatableStopError(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	eErr := rpctypes.Error(err)
	if serverErr, ok := eErr.(rpctypes.EtcdError); ok && serverErr.Code() != codes.Unavailable {
		return true
	}
	ev, _ := status.FromError(err)
	return ev.Code() != codes.Unavailable
}
func isNonRepeatableStopError(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ev, _ := status.FromError(err)
	if ev.Code() != codes.Unavailable {
		return true
	}
	desc := rpctypes.ErrorDesc(err)
	return desc != "there is no address available" && desc != "there is no connection available"
}
func (c *Client) newRetryWrapper() retryRPCFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(rpcCtx context.Context, f rpcFunc, rp retryPolicy) error {
		var isStop retryStopErrFunc
		switch rp {
		case repeatable:
			isStop = isRepeatableStopError
		case nonRepeatable:
			isStop = isNonRepeatableStopError
		}
		for {
			if err := readyWait(rpcCtx, c.ctx, c.balancer.ConnectNotify()); err != nil {
				return err
			}
			pinned := c.balancer.pinned()
			err := f(rpcCtx)
			if err == nil {
				return nil
			}
			logger.Lvl(4).Infof("clientv3/retry: error %q on pinned endpoint %q", err.Error(), pinned)
			if s, ok := status.FromError(err); ok && (s.Code() == codes.Unavailable || s.Code() == codes.DeadlineExceeded || s.Code() == codes.Internal) {
				c.balancer.hostPortError(pinned, err)
				c.balancer.next()
				logger.Lvl(4).Infof("clientv3/retry: switching from %q due to error %q", pinned, err.Error())
			}
			if isStop(err) {
				return err
			}
		}
	}
}
func (c *Client) newAuthRetryWrapper(retryf retryRPCFunc) retryRPCFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(rpcCtx context.Context, f rpcFunc, rp retryPolicy) error {
		for {
			pinned := c.balancer.pinned()
			err := retryf(rpcCtx, f, rp)
			if err == nil {
				return nil
			}
			logger.Lvl(4).Infof("clientv3/auth-retry: error %q on pinned endpoint %q", err.Error(), pinned)
			if rpctypes.Error(err) == rpctypes.ErrInvalidAuthToken {
				gterr := c.getToken(rpcCtx)
				if gterr != nil {
					logger.Lvl(4).Infof("clientv3/auth-retry: cannot retry due to error %q(%q) on pinned endpoint %q", err.Error(), gterr.Error(), pinned)
					return err
				}
				continue
			}
			return err
		}
	}
}

type retryKVClient struct {
	kc	pb.KVClient
	retryf	retryRPCFunc
}

func RetryKVClient(c *Client) pb.KVClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryKVClient{kc: pb.NewKVClient(c.conn), retryf: c.newAuthRetryWrapper(c.newRetryWrapper())}
}
func (rkv *retryKVClient) Range(ctx context.Context, in *pb.RangeRequest, opts ...grpc.CallOption) (resp *pb.RangeResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rkv.retryf(ctx, func(rctx context.Context) error {
		resp, err = rkv.kc.Range(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rkv *retryKVClient) Put(ctx context.Context, in *pb.PutRequest, opts ...grpc.CallOption) (resp *pb.PutResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rkv.retryf(ctx, func(rctx context.Context) error {
		resp, err = rkv.kc.Put(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rkv *retryKVClient) DeleteRange(ctx context.Context, in *pb.DeleteRangeRequest, opts ...grpc.CallOption) (resp *pb.DeleteRangeResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rkv.retryf(ctx, func(rctx context.Context) error {
		resp, err = rkv.kc.DeleteRange(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rkv *retryKVClient) Txn(ctx context.Context, in *pb.TxnRequest, opts ...grpc.CallOption) (resp *pb.TxnResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rkv.retryf(ctx, func(rctx context.Context) error {
		resp, err = rkv.kc.Txn(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rkv *retryKVClient) Compact(ctx context.Context, in *pb.CompactionRequest, opts ...grpc.CallOption) (resp *pb.CompactionResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rkv.retryf(ctx, func(rctx context.Context) error {
		resp, err = rkv.kc.Compact(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}

type retryLeaseClient struct {
	lc	pb.LeaseClient
	retryf	retryRPCFunc
}

func RetryLeaseClient(c *Client) pb.LeaseClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryLeaseClient{lc: pb.NewLeaseClient(c.conn), retryf: c.newAuthRetryWrapper(c.newRetryWrapper())}
}
func (rlc *retryLeaseClient) LeaseTimeToLive(ctx context.Context, in *pb.LeaseTimeToLiveRequest, opts ...grpc.CallOption) (resp *pb.LeaseTimeToLiveResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rlc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rlc.lc.LeaseTimeToLive(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rlc *retryLeaseClient) LeaseLeases(ctx context.Context, in *pb.LeaseLeasesRequest, opts ...grpc.CallOption) (resp *pb.LeaseLeasesResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rlc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rlc.lc.LeaseLeases(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rlc *retryLeaseClient) LeaseGrant(ctx context.Context, in *pb.LeaseGrantRequest, opts ...grpc.CallOption) (resp *pb.LeaseGrantResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rlc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rlc.lc.LeaseGrant(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rlc *retryLeaseClient) LeaseRevoke(ctx context.Context, in *pb.LeaseRevokeRequest, opts ...grpc.CallOption) (resp *pb.LeaseRevokeResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rlc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rlc.lc.LeaseRevoke(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rlc *retryLeaseClient) LeaseKeepAlive(ctx context.Context, opts ...grpc.CallOption) (stream pb.Lease_LeaseKeepAliveClient, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rlc.retryf(ctx, func(rctx context.Context) error {
		stream, err = rlc.lc.LeaseKeepAlive(rctx, opts...)
		return err
	}, repeatable)
	return stream, err
}

type retryClusterClient struct {
	cc	pb.ClusterClient
	retryf	retryRPCFunc
}

func RetryClusterClient(c *Client) pb.ClusterClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryClusterClient{cc: pb.NewClusterClient(c.conn), retryf: c.newRetryWrapper()}
}
func (rcc *retryClusterClient) MemberList(ctx context.Context, in *pb.MemberListRequest, opts ...grpc.CallOption) (resp *pb.MemberListResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rcc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rcc.cc.MemberList(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rcc *retryClusterClient) MemberAdd(ctx context.Context, in *pb.MemberAddRequest, opts ...grpc.CallOption) (resp *pb.MemberAddResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rcc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rcc.cc.MemberAdd(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rcc *retryClusterClient) MemberRemove(ctx context.Context, in *pb.MemberRemoveRequest, opts ...grpc.CallOption) (resp *pb.MemberRemoveResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rcc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rcc.cc.MemberRemove(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rcc *retryClusterClient) MemberUpdate(ctx context.Context, in *pb.MemberUpdateRequest, opts ...grpc.CallOption) (resp *pb.MemberUpdateResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rcc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rcc.cc.MemberUpdate(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}

type retryMaintenanceClient struct {
	mc	pb.MaintenanceClient
	retryf	retryRPCFunc
}

func RetryMaintenanceClient(c *Client, conn *grpc.ClientConn) pb.MaintenanceClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryMaintenanceClient{mc: pb.NewMaintenanceClient(conn), retryf: c.newRetryWrapper()}
}
func (rmc *retryMaintenanceClient) Alarm(ctx context.Context, in *pb.AlarmRequest, opts ...grpc.CallOption) (resp *pb.AlarmResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rmc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rmc.mc.Alarm(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rmc *retryMaintenanceClient) Status(ctx context.Context, in *pb.StatusRequest, opts ...grpc.CallOption) (resp *pb.StatusResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rmc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rmc.mc.Status(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rmc *retryMaintenanceClient) Hash(ctx context.Context, in *pb.HashRequest, opts ...grpc.CallOption) (resp *pb.HashResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rmc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rmc.mc.Hash(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rmc *retryMaintenanceClient) HashKV(ctx context.Context, in *pb.HashKVRequest, opts ...grpc.CallOption) (resp *pb.HashKVResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rmc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rmc.mc.HashKV(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rmc *retryMaintenanceClient) Snapshot(ctx context.Context, in *pb.SnapshotRequest, opts ...grpc.CallOption) (stream pb.Maintenance_SnapshotClient, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rmc.retryf(ctx, func(rctx context.Context) error {
		stream, err = rmc.mc.Snapshot(rctx, in, opts...)
		return err
	}, repeatable)
	return stream, err
}
func (rmc *retryMaintenanceClient) MoveLeader(ctx context.Context, in *pb.MoveLeaderRequest, opts ...grpc.CallOption) (resp *pb.MoveLeaderResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rmc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rmc.mc.MoveLeader(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rmc *retryMaintenanceClient) Defragment(ctx context.Context, in *pb.DefragmentRequest, opts ...grpc.CallOption) (resp *pb.DefragmentResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rmc.retryf(ctx, func(rctx context.Context) error {
		resp, err = rmc.mc.Defragment(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}

type retryAuthClient struct {
	ac	pb.AuthClient
	retryf	retryRPCFunc
}

func RetryAuthClient(c *Client) pb.AuthClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryAuthClient{ac: pb.NewAuthClient(c.conn), retryf: c.newRetryWrapper()}
}
func (rac *retryAuthClient) UserList(ctx context.Context, in *pb.AuthUserListRequest, opts ...grpc.CallOption) (resp *pb.AuthUserListResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.UserList(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rac *retryAuthClient) UserGet(ctx context.Context, in *pb.AuthUserGetRequest, opts ...grpc.CallOption) (resp *pb.AuthUserGetResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.UserGet(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rac *retryAuthClient) RoleGet(ctx context.Context, in *pb.AuthRoleGetRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleGetResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.RoleGet(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rac *retryAuthClient) RoleList(ctx context.Context, in *pb.AuthRoleListRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleListResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.RoleList(rctx, in, opts...)
		return err
	}, repeatable)
	return resp, err
}
func (rac *retryAuthClient) AuthEnable(ctx context.Context, in *pb.AuthEnableRequest, opts ...grpc.CallOption) (resp *pb.AuthEnableResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.AuthEnable(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) AuthDisable(ctx context.Context, in *pb.AuthDisableRequest, opts ...grpc.CallOption) (resp *pb.AuthDisableResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.AuthDisable(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) UserAdd(ctx context.Context, in *pb.AuthUserAddRequest, opts ...grpc.CallOption) (resp *pb.AuthUserAddResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.UserAdd(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) UserDelete(ctx context.Context, in *pb.AuthUserDeleteRequest, opts ...grpc.CallOption) (resp *pb.AuthUserDeleteResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.UserDelete(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) UserChangePassword(ctx context.Context, in *pb.AuthUserChangePasswordRequest, opts ...grpc.CallOption) (resp *pb.AuthUserChangePasswordResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.UserChangePassword(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) UserGrantRole(ctx context.Context, in *pb.AuthUserGrantRoleRequest, opts ...grpc.CallOption) (resp *pb.AuthUserGrantRoleResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.UserGrantRole(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) UserRevokeRole(ctx context.Context, in *pb.AuthUserRevokeRoleRequest, opts ...grpc.CallOption) (resp *pb.AuthUserRevokeRoleResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.UserRevokeRole(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) RoleAdd(ctx context.Context, in *pb.AuthRoleAddRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleAddResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.RoleAdd(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) RoleDelete(ctx context.Context, in *pb.AuthRoleDeleteRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleDeleteResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.RoleDelete(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) RoleGrantPermission(ctx context.Context, in *pb.AuthRoleGrantPermissionRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleGrantPermissionResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.RoleGrantPermission(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) RoleRevokePermission(ctx context.Context, in *pb.AuthRoleRevokePermissionRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleRevokePermissionResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.RoleRevokePermission(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
func (rac *retryAuthClient) Authenticate(ctx context.Context, in *pb.AuthenticateRequest, opts ...grpc.CallOption) (resp *pb.AuthenticateResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = rac.retryf(ctx, func(rctx context.Context) error {
		resp, err = rac.ac.Authenticate(rctx, in, opts...)
		return err
	}, nonRepeatable)
	return resp, err
}
