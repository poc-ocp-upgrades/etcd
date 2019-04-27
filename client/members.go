package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"github.com/coreos/etcd/pkg/types"
)

var (
	defaultV2MembersPrefix	= "/v2/members"
	defaultLeaderSuffix	= "/leader"
)

type Member struct {
	ID		string		`json:"id"`
	Name		string		`json:"name"`
	PeerURLs	[]string	`json:"peerURLs"`
	ClientURLs	[]string	`json:"clientURLs"`
}
type memberCollection []Member

func (c *memberCollection) UnmarshalJSON(data []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := struct{ Members []Member }{}
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	if d.Members == nil {
		*c = make([]Member, 0)
		return nil
	}
	*c = d.Members
	return nil
}

type memberCreateOrUpdateRequest struct{ PeerURLs types.URLs }

func (m *memberCreateOrUpdateRequest) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := struct {
		PeerURLs []string `json:"peerURLs"`
	}{PeerURLs: make([]string, len(m.PeerURLs))}
	for i, u := range m.PeerURLs {
		s.PeerURLs[i] = u.String()
	}
	return json.Marshal(&s)
}
func NewMembersAPI(c Client) MembersAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &httpMembersAPI{client: c}
}

type MembersAPI interface {
	List(ctx context.Context) ([]Member, error)
	Add(ctx context.Context, peerURL string) (*Member, error)
	Remove(ctx context.Context, mID string) error
	Update(ctx context.Context, mID string, peerURLs []string) error
	Leader(ctx context.Context) (*Member, error)
}
type httpMembersAPI struct{ client httpClient }

func (m *httpMembersAPI) List(ctx context.Context) ([]Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := &membersAPIActionList{}
	resp, body, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := assertStatusCode(resp.StatusCode, http.StatusOK); err != nil {
		return nil, err
	}
	var mCollection memberCollection
	if err := json.Unmarshal(body, &mCollection); err != nil {
		return nil, err
	}
	return []Member(mCollection), nil
}
func (m *httpMembersAPI) Add(ctx context.Context, peerURL string) (*Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	urls, err := types.NewURLs([]string{peerURL})
	if err != nil {
		return nil, err
	}
	req := &membersAPIActionAdd{peerURLs: urls}
	resp, body, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := assertStatusCode(resp.StatusCode, http.StatusCreated, http.StatusConflict); err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		var merr membersError
		if err := json.Unmarshal(body, &merr); err != nil {
			return nil, err
		}
		return nil, merr
	}
	var memb Member
	if err := json.Unmarshal(body, &memb); err != nil {
		return nil, err
	}
	return &memb, nil
}
func (m *httpMembersAPI) Update(ctx context.Context, memberID string, peerURLs []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	urls, err := types.NewURLs(peerURLs)
	if err != nil {
		return err
	}
	req := &membersAPIActionUpdate{peerURLs: urls, memberID: memberID}
	resp, body, err := m.client.Do(ctx, req)
	if err != nil {
		return err
	}
	if err := assertStatusCode(resp.StatusCode, http.StatusNoContent, http.StatusNotFound, http.StatusConflict); err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		var merr membersError
		if err := json.Unmarshal(body, &merr); err != nil {
			return err
		}
		return merr
	}
	return nil
}
func (m *httpMembersAPI) Remove(ctx context.Context, memberID string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := &membersAPIActionRemove{memberID: memberID}
	resp, _, err := m.client.Do(ctx, req)
	if err != nil {
		return err
	}
	return assertStatusCode(resp.StatusCode, http.StatusNoContent, http.StatusGone)
}
func (m *httpMembersAPI) Leader(ctx context.Context) (*Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := &membersAPIActionLeader{}
	resp, body, err := m.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := assertStatusCode(resp.StatusCode, http.StatusOK); err != nil {
		return nil, err
	}
	var leader Member
	if err := json.Unmarshal(body, &leader); err != nil {
		return nil, err
	}
	return &leader, nil
}

type membersAPIActionList struct{}

func (l *membersAPIActionList) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2MembersURL(ep)
	req, _ := http.NewRequest("GET", u.String(), nil)
	return req
}

type membersAPIActionRemove struct{ memberID string }

func (d *membersAPIActionRemove) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2MembersURL(ep)
	u.Path = path.Join(u.Path, d.memberID)
	req, _ := http.NewRequest("DELETE", u.String(), nil)
	return req
}

type membersAPIActionAdd struct{ peerURLs types.URLs }

func (a *membersAPIActionAdd) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2MembersURL(ep)
	m := memberCreateOrUpdateRequest{PeerURLs: a.peerURLs}
	b, _ := json.Marshal(&m)
	req, _ := http.NewRequest("POST", u.String(), bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	return req
}

type membersAPIActionUpdate struct {
	memberID	string
	peerURLs	types.URLs
}

func (a *membersAPIActionUpdate) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2MembersURL(ep)
	m := memberCreateOrUpdateRequest{PeerURLs: a.peerURLs}
	u.Path = path.Join(u.Path, a.memberID)
	b, _ := json.Marshal(&m)
	req, _ := http.NewRequest("PUT", u.String(), bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	return req
}
func assertStatusCode(got int, want ...int) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, w := range want {
		if w == got {
			return nil
		}
	}
	return fmt.Errorf("unexpected status code %d", got)
}

type membersAPIActionLeader struct{}

func (l *membersAPIActionLeader) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2MembersURL(ep)
	u.Path = path.Join(u.Path, defaultLeaderSuffix)
	req, _ := http.NewRequest("GET", u.String(), nil)
	return req
}
func v2MembersURL(ep url.URL) *url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ep.Path = path.Join(ep.Path, defaultV2MembersPrefix)
	return &ep
}

type membersError struct {
	Message	string	`json:"message"`
	Code	int	`json:"-"`
}

func (e membersError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.Message
}
