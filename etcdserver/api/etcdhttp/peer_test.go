package etcdhttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"sort"
	"testing"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/rafthttp"
	"github.com/coreos/go-semver/semver"
)

type fakeCluster struct {
	id		uint64
	clientURLs	[]string
	members		map[uint64]*membership.Member
}

func (c *fakeCluster) ID() types.ID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.ID(c.id)
}
func (c *fakeCluster) ClientURLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.clientURLs
}
func (c *fakeCluster) Members() []*membership.Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ms membership.MembersByID
	for _, m := range c.members {
		ms = append(ms, m)
	}
	sort.Sort(ms)
	return []*membership.Member(ms)
}
func (c *fakeCluster) Member(id types.ID) *membership.Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.members[uint64(id)]
}
func (c *fakeCluster) Version() *semver.Version {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func TestNewPeerHandlerOnRaftPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test data"))
	})
	ph := newPeerHandler(&fakeCluster{}, h, nil)
	srv := httptest.NewServer(ph)
	defer srv.Close()
	tests := []string{rafthttp.RaftPrefix, rafthttp.RaftPrefix + "/hello"}
	for i, tt := range tests {
		resp, err := http.Get(srv.URL + tt)
		if err != nil {
			t.Fatalf("unexpected http.Get error: %v", err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("unexpected ioutil.ReadAll error: %v", err)
		}
		if w := "test data"; string(body) != w {
			t.Errorf("#%d: body = %s, want %s", i, body, w)
		}
	}
}
func TestServeMembersFails(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		method	string
		wcode	int
	}{{"POST", http.StatusMethodNotAllowed}, {"DELETE", http.StatusMethodNotAllowed}, {"BAD", http.StatusMethodNotAllowed}}
	for i, tt := range tests {
		rw := httptest.NewRecorder()
		h := &peerMembersHandler{cluster: nil}
		h.ServeHTTP(rw, &http.Request{Method: tt.method})
		if rw.Code != tt.wcode {
			t.Errorf("#%d: code=%d, want %d", i, rw.Code, tt.wcode)
		}
	}
}
func TestServeMembersGet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	memb1 := membership.Member{ID: 1, Attributes: membership.Attributes{ClientURLs: []string{"http://localhost:8080"}}}
	memb2 := membership.Member{ID: 2, Attributes: membership.Attributes{ClientURLs: []string{"http://localhost:8081"}}}
	cluster := &fakeCluster{id: 1, members: map[uint64]*membership.Member{1: &memb1, 2: &memb2}}
	h := &peerMembersHandler{cluster: cluster}
	msb, err := json.Marshal([]membership.Member{memb1, memb2})
	if err != nil {
		t.Fatal(err)
	}
	wms := string(msb) + "\n"
	tests := []struct {
		path	string
		wcode	int
		wct	string
		wbody	string
	}{{peerMembersPrefix, http.StatusOK, "application/json", wms}, {path.Join(peerMembersPrefix, "bad"), http.StatusBadRequest, "text/plain; charset=utf-8", "bad path\n"}}
	for i, tt := range tests {
		req, err := http.NewRequest("GET", testutil.MustNewURL(t, tt.path).String(), nil)
		if err != nil {
			t.Fatal(err)
		}
		rw := httptest.NewRecorder()
		h.ServeHTTP(rw, req)
		if rw.Code != tt.wcode {
			t.Errorf("#%d: code=%d, want %d", i, rw.Code, tt.wcode)
		}
		if gct := rw.Header().Get("Content-Type"); gct != tt.wct {
			t.Errorf("#%d: content-type = %s, want %s", i, gct, tt.wct)
		}
		if rw.Body.String() != tt.wbody {
			t.Errorf("#%d: body = %s, want %s", i, rw.Body.String(), tt.wbody)
		}
		gcid := rw.Header().Get("X-Etcd-Cluster-ID")
		wcid := cluster.ID().String()
		if gcid != wcid {
			t.Errorf("#%d: cid = %s, want %s", i, gcid, wcid)
		}
	}
}
