package membership

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math/rand"
	"sort"
	"time"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/pkg/capnslog"
)

var (
	plog = capnslog.NewPackageLogger("github.com/coreos/etcd", "etcdserver/membership")
)

type RaftAttributes struct {
	PeerURLs []string `json:"peerURLs"`
}
type Attributes struct {
	Name		string		`json:"name,omitempty"`
	ClientURLs	[]string	`json:"clientURLs,omitempty"`
}
type Member struct {
	ID	types.ID	`json:"id"`
	RaftAttributes
	Attributes
}

func NewMember(name string, peerURLs types.URLs, clusterName string, now *time.Time) *Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := &Member{RaftAttributes: RaftAttributes{PeerURLs: peerURLs.StringSlice()}, Attributes: Attributes{Name: name}}
	var b []byte
	sort.Strings(m.PeerURLs)
	for _, p := range m.PeerURLs {
		b = append(b, []byte(p)...)
	}
	b = append(b, []byte(clusterName)...)
	if now != nil {
		b = append(b, []byte(fmt.Sprintf("%d", now.Unix()))...)
	}
	hash := sha1.Sum(b)
	m.ID = types.ID(binary.BigEndian.Uint64(hash[:8]))
	return m
}
func (m *Member) PickPeerURL() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(m.PeerURLs) == 0 {
		plog.Panicf("member should always have some peer url")
	}
	return m.PeerURLs[rand.Intn(len(m.PeerURLs))]
}
func (m *Member) Clone() *Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m == nil {
		return nil
	}
	mm := &Member{ID: m.ID, Attributes: Attributes{Name: m.Name}}
	if m.PeerURLs != nil {
		mm.PeerURLs = make([]string, len(m.PeerURLs))
		copy(mm.PeerURLs, m.PeerURLs)
	}
	if m.ClientURLs != nil {
		mm.ClientURLs = make([]string, len(m.ClientURLs))
		copy(mm.ClientURLs, m.ClientURLs)
	}
	return mm
}
func (m *Member) IsStarted() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(m.Name) != 0
}

type MembersByID []*Member

func (ms MembersByID) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(ms)
}
func (ms MembersByID) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ms[i].ID < ms[j].ID
}
func (ms MembersByID) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ms[i], ms[j] = ms[j], ms[i]
}

type MembersByPeerURLs []*Member

func (ms MembersByPeerURLs) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(ms)
}
func (ms MembersByPeerURLs) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ms[i].PeerURLs[0] < ms[j].PeerURLs[0]
}
func (ms MembersByPeerURLs) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ms[i], ms[j] = ms[j], ms[i]
}
