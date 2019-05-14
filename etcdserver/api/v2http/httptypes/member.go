package httptypes

import (
	"encoding/json"
	"github.com/coreos/etcd/pkg/types"
)

type Member struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	PeerURLs   []string `json:"peerURLs"`
	ClientURLs []string `json:"clientURLs"`
}
type MemberCreateRequest struct{ PeerURLs types.URLs }
type MemberUpdateRequest struct{ MemberCreateRequest }

func (m *MemberCreateRequest) UnmarshalJSON(data []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := struct {
		PeerURLs []string `json:"peerURLs"`
	}{}
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	urls, err := types.NewURLs(s.PeerURLs)
	if err != nil {
		return err
	}
	m.PeerURLs = urls
	return nil
}

type MemberCollection []Member

func (c *MemberCollection) MarshalJSON() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := struct {
		Members []Member `json:"members"`
	}{Members: []Member(*c)}
	return json.Marshal(d)
}
