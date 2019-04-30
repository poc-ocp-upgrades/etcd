package clientv3

import (
	"context"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type retryPolicy uint8

const (
	repeatable	retryPolicy	= iota
	nonRepeatable
)

func (rp retryPolicy) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch rp {
	case repeatable:
		return "repeatable"
	case nonRepeatable:
		return "nonRepeatable"
	default:
		return "UNKNOWN"
	}
}

type rpcFunc func(ctx context.Context) error
type retryRPCFunc func(context.Context, rpcFunc, retryPolicy) error
type retryStopErrFunc func(error) bool

func isSafeRetryImmutableRPC(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eErr := rpctypes.Error(err)
	if serverErr, ok := eErr.(rpctypes.EtcdError); ok && serverErr.Code() != codes.Unavailable {
		return false
	}
	ev, ok := status.FromError(err)
	if !ok {
		return false
	}
	return ev.Code() == codes.Unavailable
}
func isSafeRetryMutableRPC(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ev, ok := status.FromError(err); ok && ev.Code() != codes.Unavailable {
		return false
	}
	desc := rpctypes.ErrorDesc(err)
	return desc == "there is no address available" || desc == "there is no connection available"
}

type retryKVClient struct{ kc pb.KVClient }

func RetryKVClient(c *Client) pb.KVClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryKVClient{kc: pb.NewKVClient(c.conn)}
}
func (rkv *retryKVClient) Range(ctx context.Context, in *pb.RangeRequest, opts ...grpc.CallOption) (resp *pb.RangeResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rkv.kc.Range(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rkv *retryKVClient) Put(ctx context.Context, in *pb.PutRequest, opts ...grpc.CallOption) (resp *pb.PutResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rkv.kc.Put(ctx, in, opts...)
}
func (rkv *retryKVClient) DeleteRange(ctx context.Context, in *pb.DeleteRangeRequest, opts ...grpc.CallOption) (resp *pb.DeleteRangeResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rkv.kc.DeleteRange(ctx, in, opts...)
}
func (rkv *retryKVClient) Txn(ctx context.Context, in *pb.TxnRequest, opts ...grpc.CallOption) (resp *pb.TxnResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rkv.kc.Txn(ctx, in, opts...)
}
func (rkv *retryKVClient) Compact(ctx context.Context, in *pb.CompactionRequest, opts ...grpc.CallOption) (resp *pb.CompactionResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rkv.kc.Compact(ctx, in, opts...)
}

type retryLeaseClient struct{ lc pb.LeaseClient }

func RetryLeaseClient(c *Client) pb.LeaseClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryLeaseClient{lc: pb.NewLeaseClient(c.conn)}
}
func (rlc *retryLeaseClient) LeaseTimeToLive(ctx context.Context, in *pb.LeaseTimeToLiveRequest, opts ...grpc.CallOption) (resp *pb.LeaseTimeToLiveResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rlc.lc.LeaseTimeToLive(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rlc *retryLeaseClient) LeaseLeases(ctx context.Context, in *pb.LeaseLeasesRequest, opts ...grpc.CallOption) (resp *pb.LeaseLeasesResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rlc.lc.LeaseLeases(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rlc *retryLeaseClient) LeaseGrant(ctx context.Context, in *pb.LeaseGrantRequest, opts ...grpc.CallOption) (resp *pb.LeaseGrantResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rlc.lc.LeaseGrant(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rlc *retryLeaseClient) LeaseRevoke(ctx context.Context, in *pb.LeaseRevokeRequest, opts ...grpc.CallOption) (resp *pb.LeaseRevokeResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rlc.lc.LeaseRevoke(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rlc *retryLeaseClient) LeaseKeepAlive(ctx context.Context, opts ...grpc.CallOption) (stream pb.Lease_LeaseKeepAliveClient, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rlc.lc.LeaseKeepAlive(ctx, append(opts, withRetryPolicy(repeatable))...)
}

type retryClusterClient struct{ cc pb.ClusterClient }

func RetryClusterClient(c *Client) pb.ClusterClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryClusterClient{cc: pb.NewClusterClient(c.conn)}
}
func (rcc *retryClusterClient) MemberList(ctx context.Context, in *pb.MemberListRequest, opts ...grpc.CallOption) (resp *pb.MemberListResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rcc.cc.MemberList(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rcc *retryClusterClient) MemberAdd(ctx context.Context, in *pb.MemberAddRequest, opts ...grpc.CallOption) (resp *pb.MemberAddResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rcc.cc.MemberAdd(ctx, in, opts...)
}
func (rcc *retryClusterClient) MemberRemove(ctx context.Context, in *pb.MemberRemoveRequest, opts ...grpc.CallOption) (resp *pb.MemberRemoveResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rcc.cc.MemberRemove(ctx, in, opts...)
}
func (rcc *retryClusterClient) MemberUpdate(ctx context.Context, in *pb.MemberUpdateRequest, opts ...grpc.CallOption) (resp *pb.MemberUpdateResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rcc.cc.MemberUpdate(ctx, in, opts...)
}

type retryMaintenanceClient struct{ mc pb.MaintenanceClient }

func RetryMaintenanceClient(c *Client, conn *grpc.ClientConn) pb.MaintenanceClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryMaintenanceClient{mc: pb.NewMaintenanceClient(conn)}
}
func (rmc *retryMaintenanceClient) Alarm(ctx context.Context, in *pb.AlarmRequest, opts ...grpc.CallOption) (resp *pb.AlarmResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rmc.mc.Alarm(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rmc *retryMaintenanceClient) Status(ctx context.Context, in *pb.StatusRequest, opts ...grpc.CallOption) (resp *pb.StatusResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rmc.mc.Status(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rmc *retryMaintenanceClient) Hash(ctx context.Context, in *pb.HashRequest, opts ...grpc.CallOption) (resp *pb.HashResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rmc.mc.Hash(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rmc *retryMaintenanceClient) HashKV(ctx context.Context, in *pb.HashKVRequest, opts ...grpc.CallOption) (resp *pb.HashKVResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rmc.mc.HashKV(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rmc *retryMaintenanceClient) Snapshot(ctx context.Context, in *pb.SnapshotRequest, opts ...grpc.CallOption) (stream pb.Maintenance_SnapshotClient, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rmc.mc.Snapshot(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rmc *retryMaintenanceClient) MoveLeader(ctx context.Context, in *pb.MoveLeaderRequest, opts ...grpc.CallOption) (resp *pb.MoveLeaderResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rmc.mc.MoveLeader(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rmc *retryMaintenanceClient) Defragment(ctx context.Context, in *pb.DefragmentRequest, opts ...grpc.CallOption) (resp *pb.DefragmentResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rmc.mc.Defragment(ctx, in, opts...)
}

type retryAuthClient struct{ ac pb.AuthClient }

func RetryAuthClient(c *Client) pb.AuthClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &retryAuthClient{ac: pb.NewAuthClient(c.conn)}
}
func (rac *retryAuthClient) UserList(ctx context.Context, in *pb.AuthUserListRequest, opts ...grpc.CallOption) (resp *pb.AuthUserListResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.UserList(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rac *retryAuthClient) UserGet(ctx context.Context, in *pb.AuthUserGetRequest, opts ...grpc.CallOption) (resp *pb.AuthUserGetResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.UserGet(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rac *retryAuthClient) RoleGet(ctx context.Context, in *pb.AuthRoleGetRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleGetResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.RoleGet(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rac *retryAuthClient) RoleList(ctx context.Context, in *pb.AuthRoleListRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleListResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.RoleList(ctx, in, append(opts, withRetryPolicy(repeatable))...)
}
func (rac *retryAuthClient) AuthEnable(ctx context.Context, in *pb.AuthEnableRequest, opts ...grpc.CallOption) (resp *pb.AuthEnableResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.AuthEnable(ctx, in, opts...)
}
func (rac *retryAuthClient) AuthDisable(ctx context.Context, in *pb.AuthDisableRequest, opts ...grpc.CallOption) (resp *pb.AuthDisableResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.AuthDisable(ctx, in, opts...)
}
func (rac *retryAuthClient) UserAdd(ctx context.Context, in *pb.AuthUserAddRequest, opts ...grpc.CallOption) (resp *pb.AuthUserAddResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.UserAdd(ctx, in, opts...)
}
func (rac *retryAuthClient) UserDelete(ctx context.Context, in *pb.AuthUserDeleteRequest, opts ...grpc.CallOption) (resp *pb.AuthUserDeleteResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.UserDelete(ctx, in, opts...)
}
func (rac *retryAuthClient) UserChangePassword(ctx context.Context, in *pb.AuthUserChangePasswordRequest, opts ...grpc.CallOption) (resp *pb.AuthUserChangePasswordResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.UserChangePassword(ctx, in, opts...)
}
func (rac *retryAuthClient) UserGrantRole(ctx context.Context, in *pb.AuthUserGrantRoleRequest, opts ...grpc.CallOption) (resp *pb.AuthUserGrantRoleResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.UserGrantRole(ctx, in, opts...)
}
func (rac *retryAuthClient) UserRevokeRole(ctx context.Context, in *pb.AuthUserRevokeRoleRequest, opts ...grpc.CallOption) (resp *pb.AuthUserRevokeRoleResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.UserRevokeRole(ctx, in, opts...)
}
func (rac *retryAuthClient) RoleAdd(ctx context.Context, in *pb.AuthRoleAddRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleAddResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.RoleAdd(ctx, in, opts...)
}
func (rac *retryAuthClient) RoleDelete(ctx context.Context, in *pb.AuthRoleDeleteRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleDeleteResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.RoleDelete(ctx, in, opts...)
}
func (rac *retryAuthClient) RoleGrantPermission(ctx context.Context, in *pb.AuthRoleGrantPermissionRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleGrantPermissionResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.RoleGrantPermission(ctx, in, opts...)
}
func (rac *retryAuthClient) RoleRevokePermission(ctx context.Context, in *pb.AuthRoleRevokePermissionRequest, opts ...grpc.CallOption) (resp *pb.AuthRoleRevokePermissionResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.RoleRevokePermission(ctx, in, opts...)
}
func (rac *retryAuthClient) Authenticate(ctx context.Context, in *pb.AuthenticateRequest, opts ...grpc.CallOption) (resp *pb.AuthenticateResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rac.ac.Authenticate(ctx, in, opts...)
}
