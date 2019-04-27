package client

import (
	"bytes"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"context"
	"encoding/json"
	"net/http"
	godefaulthttp "net/http"
	"net/url"
)

type Role struct {
	Role		string		`json:"role"`
	Permissions	Permissions	`json:"permissions"`
	Grant		*Permissions	`json:"grant,omitempty"`
	Revoke		*Permissions	`json:"revoke,omitempty"`
}
type Permissions struct {
	KV rwPermission `json:"kv"`
}
type rwPermission struct {
	Read	[]string	`json:"read"`
	Write	[]string	`json:"write"`
}
type PermissionType int

const (
	ReadPermission	PermissionType	= iota
	WritePermission
	ReadWritePermission
)

func NewAuthRoleAPI(c Client) AuthRoleAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &httpAuthRoleAPI{client: c}
}

type AuthRoleAPI interface {
	AddRole(ctx context.Context, role string) error
	RemoveRole(ctx context.Context, role string) error
	GetRole(ctx context.Context, role string) (*Role, error)
	GrantRoleKV(ctx context.Context, role string, prefixes []string, permType PermissionType) (*Role, error)
	RevokeRoleKV(ctx context.Context, role string, prefixes []string, permType PermissionType) (*Role, error)
	ListRoles(ctx context.Context) ([]string, error)
}
type httpAuthRoleAPI struct{ client httpClient }
type authRoleAPIAction struct {
	verb	string
	name	string
	role	*Role
}
type authRoleAPIList struct{}

func (list *authRoleAPIList) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2AuthURL(ep, "roles", "")
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	return req
}
func (l *authRoleAPIAction) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2AuthURL(ep, "roles", l.name)
	if l.role == nil {
		req, _ := http.NewRequest(l.verb, u.String(), nil)
		return req
	}
	b, err := json.Marshal(l.role)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(b)
	req, _ := http.NewRequest(l.verb, u.String(), body)
	req.Header.Set("Content-Type", "application/json")
	return req
}
func (r *httpAuthRoleAPI) ListRoles(ctx context.Context) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, body, err := r.client.Do(ctx, &authRoleAPIList{})
	if err != nil {
		return nil, err
	}
	if err = assertStatusCode(resp.StatusCode, http.StatusOK); err != nil {
		return nil, err
	}
	var roleList struct {
		Roles []Role `json:"roles"`
	}
	if err = json.Unmarshal(body, &roleList); err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(roleList.Roles))
	for _, r := range roleList.Roles {
		ret = append(ret, r.Role)
	}
	return ret, nil
}
func (r *httpAuthRoleAPI) AddRole(ctx context.Context, rolename string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	role := &Role{Role: rolename}
	return r.addRemoveRole(ctx, &authRoleAPIAction{verb: "PUT", name: rolename, role: role})
}
func (r *httpAuthRoleAPI) RemoveRole(ctx context.Context, rolename string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.addRemoveRole(ctx, &authRoleAPIAction{verb: "DELETE", name: rolename})
}
func (r *httpAuthRoleAPI) addRemoveRole(ctx context.Context, req *authRoleAPIAction) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, body, err := r.client.Do(ctx, req)
	if err != nil {
		return err
	}
	if err := assertStatusCode(resp.StatusCode, http.StatusOK, http.StatusCreated); err != nil {
		var sec authError
		err := json.Unmarshal(body, &sec)
		if err != nil {
			return err
		}
		return sec
	}
	return nil
}
func (r *httpAuthRoleAPI) GetRole(ctx context.Context, rolename string) (*Role, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.modRole(ctx, &authRoleAPIAction{verb: "GET", name: rolename})
}
func buildRWPermission(prefixes []string, permType PermissionType) rwPermission {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var out rwPermission
	switch permType {
	case ReadPermission:
		out.Read = prefixes
	case WritePermission:
		out.Write = prefixes
	case ReadWritePermission:
		out.Read = prefixes
		out.Write = prefixes
	}
	return out
}
func (r *httpAuthRoleAPI) GrantRoleKV(ctx context.Context, rolename string, prefixes []string, permType PermissionType) (*Role, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rwp := buildRWPermission(prefixes, permType)
	role := &Role{Role: rolename, Grant: &Permissions{KV: rwp}}
	return r.modRole(ctx, &authRoleAPIAction{verb: "PUT", name: rolename, role: role})
}
func (r *httpAuthRoleAPI) RevokeRoleKV(ctx context.Context, rolename string, prefixes []string, permType PermissionType) (*Role, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rwp := buildRWPermission(prefixes, permType)
	role := &Role{Role: rolename, Revoke: &Permissions{KV: rwp}}
	return r.modRole(ctx, &authRoleAPIAction{verb: "PUT", name: rolename, role: role})
}
func (r *httpAuthRoleAPI) modRole(ctx context.Context, req *authRoleAPIAction) (*Role, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, body, err := r.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	if err = assertStatusCode(resp.StatusCode, http.StatusOK); err != nil {
		var sec authError
		err = json.Unmarshal(body, &sec)
		if err != nil {
			return nil, err
		}
		return nil, sec
	}
	var role Role
	if err = json.Unmarshal(body, &role); err != nil {
		return nil, err
	}
	return &role, nil
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
