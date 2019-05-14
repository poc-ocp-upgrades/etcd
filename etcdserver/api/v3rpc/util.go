package v3rpc

import (
	"context"
	"github.com/coreos/etcd/auth"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

var toGRPCErrorMap = map[error]error{membership.ErrIDRemoved: rpctypes.ErrGRPCMemberNotFound, membership.ErrIDNotFound: rpctypes.ErrGRPCMemberNotFound, membership.ErrIDExists: rpctypes.ErrGRPCMemberExist, membership.ErrPeerURLexists: rpctypes.ErrGRPCPeerURLExist, etcdserver.ErrNotEnoughStartedMembers: rpctypes.ErrMemberNotEnoughStarted, mvcc.ErrCompacted: rpctypes.ErrGRPCCompacted, mvcc.ErrFutureRev: rpctypes.ErrGRPCFutureRev, etcdserver.ErrRequestTooLarge: rpctypes.ErrGRPCRequestTooLarge, etcdserver.ErrNoSpace: rpctypes.ErrGRPCNoSpace, etcdserver.ErrTooManyRequests: rpctypes.ErrTooManyRequests, etcdserver.ErrNoLeader: rpctypes.ErrGRPCNoLeader, etcdserver.ErrNotLeader: rpctypes.ErrGRPCNotLeader, etcdserver.ErrStopped: rpctypes.ErrGRPCStopped, etcdserver.ErrTimeout: rpctypes.ErrGRPCTimeout, etcdserver.ErrTimeoutDueToLeaderFail: rpctypes.ErrGRPCTimeoutDueToLeaderFail, etcdserver.ErrTimeoutDueToConnectionLost: rpctypes.ErrGRPCTimeoutDueToConnectionLost, etcdserver.ErrUnhealthy: rpctypes.ErrGRPCUnhealthy, etcdserver.ErrKeyNotFound: rpctypes.ErrGRPCKeyNotFound, etcdserver.ErrCorrupt: rpctypes.ErrGRPCCorrupt, lease.ErrLeaseNotFound: rpctypes.ErrGRPCLeaseNotFound, lease.ErrLeaseExists: rpctypes.ErrGRPCLeaseExist, lease.ErrLeaseTTLTooLarge: rpctypes.ErrGRPCLeaseTTLTooLarge, auth.ErrRootUserNotExist: rpctypes.ErrGRPCRootUserNotExist, auth.ErrRootRoleNotExist: rpctypes.ErrGRPCRootRoleNotExist, auth.ErrUserAlreadyExist: rpctypes.ErrGRPCUserAlreadyExist, auth.ErrUserEmpty: rpctypes.ErrGRPCUserEmpty, auth.ErrUserNotFound: rpctypes.ErrGRPCUserNotFound, auth.ErrRoleAlreadyExist: rpctypes.ErrGRPCRoleAlreadyExist, auth.ErrRoleNotFound: rpctypes.ErrGRPCRoleNotFound, auth.ErrAuthFailed: rpctypes.ErrGRPCAuthFailed, auth.ErrPermissionDenied: rpctypes.ErrGRPCPermissionDenied, auth.ErrRoleNotGranted: rpctypes.ErrGRPCRoleNotGranted, auth.ErrPermissionNotGranted: rpctypes.ErrGRPCPermissionNotGranted, auth.ErrAuthNotEnabled: rpctypes.ErrGRPCAuthNotEnabled, auth.ErrInvalidAuthToken: rpctypes.ErrGRPCInvalidAuthToken, auth.ErrInvalidAuthMgmt: rpctypes.ErrGRPCInvalidAuthMgmt}

func togRPCError(err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == context.Canceled || err == context.DeadlineExceeded {
		return err
	}
	grpcErr, ok := toGRPCErrorMap[err]
	if !ok {
		return status.Error(codes.Unknown, err.Error())
	}
	return grpcErr
}
func isClientCtxErr(ctxErr error, err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ctxErr != nil {
		return true
	}
	ev, ok := status.FromError(err)
	if !ok {
		return false
	}
	switch ev.Code() {
	case codes.Canceled, codes.DeadlineExceeded:
		return true
	case codes.Unavailable:
		msg := ev.Message()
		if msg == "client disconnected" {
			return true
		}
		if strings.HasPrefix(msg, "stream error: ") && strings.HasSuffix(msg, "; CANCEL") {
			return true
		}
	}
	return false
}
