package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

var (
	defaultV2AuthPrefix = "/v2/auth"
)

type User struct {
	User     string   `json:"user"`
	Password string   `json:"password,omitempty"`
	Roles    []string `json:"roles"`
	Grant    []string `json:"grant,omitempty"`
	Revoke   []string `json:"revoke,omitempty"`
}
type userListEntry struct {
	User  string `json:"user"`
	Roles []Role `json:"roles"`
}
type UserRoles struct {
	User  string `json:"user"`
	Roles []Role `json:"roles"`
}

func v2AuthURL(ep url.URL, action string, name string) *url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if name != "" {
		ep.Path = path.Join(ep.Path, defaultV2AuthPrefix, action, name)
		return &ep
	}
	ep.Path = path.Join(ep.Path, defaultV2AuthPrefix, action)
	return &ep
}
func NewAuthAPI(c Client) AuthAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &httpAuthAPI{client: c}
}

type AuthAPI interface {
	Enable(ctx context.Context) error
	Disable(ctx context.Context) error
}
type httpAuthAPI struct{ client httpClient }

func (s *httpAuthAPI) Enable(ctx context.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.enableDisable(ctx, &authAPIAction{"PUT"})
}
func (s *httpAuthAPI) Disable(ctx context.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.enableDisable(ctx, &authAPIAction{"DELETE"})
}
func (s *httpAuthAPI) enableDisable(ctx context.Context, req httpAction) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, body, err := s.client.Do(ctx, req)
	if err != nil {
		return err
	}
	if err = assertStatusCode(resp.StatusCode, http.StatusOK, http.StatusCreated); err != nil {
		var sec authError
		err = json.Unmarshal(body, &sec)
		if err != nil {
			return err
		}
		return sec
	}
	return nil
}

type authAPIAction struct{ verb string }

func (l *authAPIAction) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2AuthURL(ep, "enable", "")
	req, _ := http.NewRequest(l.verb, u.String(), nil)
	return req
}

type authError struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func (e authError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.Message
}
func NewAuthUserAPI(c Client) AuthUserAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &httpAuthUserAPI{client: c}
}

type AuthUserAPI interface {
	AddUser(ctx context.Context, username string, password string) error
	RemoveUser(ctx context.Context, username string) error
	GetUser(ctx context.Context, username string) (*User, error)
	GrantUser(ctx context.Context, username string, roles []string) (*User, error)
	RevokeUser(ctx context.Context, username string, roles []string) (*User, error)
	ChangePassword(ctx context.Context, username string, password string) (*User, error)
	ListUsers(ctx context.Context) ([]string, error)
}
type httpAuthUserAPI struct{ client httpClient }
type authUserAPIAction struct {
	verb     string
	username string
	user     *User
}
type authUserAPIList struct{}

func (list *authUserAPIList) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2AuthURL(ep, "users", "")
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	return req
}
func (l *authUserAPIAction) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2AuthURL(ep, "users", l.username)
	if l.user == nil {
		req, _ := http.NewRequest(l.verb, u.String(), nil)
		return req
	}
	b, err := json.Marshal(l.user)
	if err != nil {
		panic(err)
	}
	body := bytes.NewReader(b)
	req, _ := http.NewRequest(l.verb, u.String(), body)
	req.Header.Set("Content-Type", "application/json")
	return req
}
func (u *httpAuthUserAPI) ListUsers(ctx context.Context) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, body, err := u.client.Do(ctx, &authUserAPIList{})
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
	var userList struct {
		Users []userListEntry `json:"users"`
	}
	if err = json.Unmarshal(body, &userList); err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(userList.Users))
	for _, u := range userList.Users {
		ret = append(ret, u.User)
	}
	return ret, nil
}
func (u *httpAuthUserAPI) AddUser(ctx context.Context, username string, password string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user := &User{User: username, Password: password}
	return u.addRemoveUser(ctx, &authUserAPIAction{verb: "PUT", username: username, user: user})
}
func (u *httpAuthUserAPI) RemoveUser(ctx context.Context, username string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return u.addRemoveUser(ctx, &authUserAPIAction{verb: "DELETE", username: username})
}
func (u *httpAuthUserAPI) addRemoveUser(ctx context.Context, req *authUserAPIAction) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, body, err := u.client.Do(ctx, req)
	if err != nil {
		return err
	}
	if err = assertStatusCode(resp.StatusCode, http.StatusOK, http.StatusCreated); err != nil {
		var sec authError
		err = json.Unmarshal(body, &sec)
		if err != nil {
			return err
		}
		return sec
	}
	return nil
}
func (u *httpAuthUserAPI) GetUser(ctx context.Context, username string) (*User, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return u.modUser(ctx, &authUserAPIAction{verb: "GET", username: username})
}
func (u *httpAuthUserAPI) GrantUser(ctx context.Context, username string, roles []string) (*User, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user := &User{User: username, Grant: roles}
	return u.modUser(ctx, &authUserAPIAction{verb: "PUT", username: username, user: user})
}
func (u *httpAuthUserAPI) RevokeUser(ctx context.Context, username string, roles []string) (*User, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user := &User{User: username, Revoke: roles}
	return u.modUser(ctx, &authUserAPIAction{verb: "PUT", username: username, user: user})
}
func (u *httpAuthUserAPI) ChangePassword(ctx context.Context, username string, password string) (*User, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	user := &User{User: username, Password: password}
	return u.modUser(ctx, &authUserAPIAction{verb: "PUT", username: username, user: user})
}
func (u *httpAuthUserAPI) modUser(ctx context.Context, req *authUserAPIAction) (*User, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, body, err := u.client.Do(ctx, req)
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
	var user User
	if err = json.Unmarshal(body, &user); err != nil {
		var userR UserRoles
		if urerr := json.Unmarshal(body, &userR); urerr != nil {
			return nil, err
		}
		user.User = userR.User
		for _, r := range userR.Roles {
			user.Roles = append(user.Roles, r.Role)
		}
	}
	return &user, nil
}
