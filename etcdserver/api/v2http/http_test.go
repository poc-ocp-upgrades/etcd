package v2http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	etcdErr "github.com/coreos/etcd/error"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
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

type errServer struct {
	err	error
	fakeServer
}

func (fs *errServer) Do(ctx context.Context, r etcdserverpb.Request) (etcdserver.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return etcdserver.Response{}, fs.err
}
func (fs *errServer) Process(ctx context.Context, m raftpb.Message) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fs.err
}
func (fs *errServer) AddMember(ctx context.Context, m membership.Member) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, fs.err
}
func (fs *errServer) RemoveMember(ctx context.Context, id uint64) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, fs.err
}
func (fs *errServer) UpdateMember(ctx context.Context, m membership.Member) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, fs.err
}
func TestWriteError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rec := httptest.NewRecorder()
	r := new(http.Request)
	writeError(rec, r, nil)
	h := rec.Header()
	if len(h) > 0 {
		t.Fatalf("unexpected non-empty headers: %#v", h)
	}
	b := rec.Body.String()
	if len(b) > 0 {
		t.Fatalf("unexpected non-empty body: %q", b)
	}
	tests := []struct {
		err	error
		wcode	int
		wi	string
	}{{etcdErr.NewError(etcdErr.EcodeKeyNotFound, "/foo/bar", 123), http.StatusNotFound, "123"}, {etcdErr.NewError(etcdErr.EcodeTestFailed, "/foo/bar", 456), http.StatusPreconditionFailed, "456"}, {err: errors.New("something went wrong"), wcode: http.StatusInternalServerError}}
	for i, tt := range tests {
		rw := httptest.NewRecorder()
		writeError(rw, r, tt.err)
		if code := rw.Code; code != tt.wcode {
			t.Errorf("#%d: code=%d, want %d", i, code, tt.wcode)
		}
		if idx := rw.Header().Get("X-Etcd-Index"); idx != tt.wi {
			t.Errorf("#%d: X-Etcd-Index=%q, want %q", i, idx, tt.wi)
		}
	}
}
func TestAllowMethod(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		m	string
		ms	[]string
		w	bool
		wh	string
	}{{m: "GET", ms: []string{"GET", "POST", "PUT"}, w: true}, {m: "POST", ms: []string{"POST"}, w: true}, {m: "FAKE", ms: []string{"GET", "POST", "PUT"}, w: false, wh: "GET,POST,PUT"}, {m: "", ms: []string{"GET", "POST"}, w: false, wh: "GET,POST"}, {m: "GET", ms: []string{""}, w: false, wh: ""}, {m: "GET", ms: []string{}, w: false, wh: ""}}
	for i, tt := range tests {
		rw := httptest.NewRecorder()
		g := allowMethod(rw, tt.m, tt.ms...)
		if g != tt.w {
			t.Errorf("#%d: got allowMethod()=%t, want %t", i, g, tt.w)
		}
		if !tt.w {
			if rw.Code != http.StatusMethodNotAllowed {
				t.Errorf("#%d: code=%d, want %d", i, rw.Code, http.StatusMethodNotAllowed)
			}
			gh := rw.Header().Get("Allow")
			if gh != tt.wh {
				t.Errorf("#%d: Allow header=%q, want %q", i, gh, tt.wh)
			}
		}
	}
}
