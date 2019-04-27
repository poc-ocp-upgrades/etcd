package v2http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"reflect"
	"strings"
	"testing"
	"time"
	etcdErr "github.com/coreos/etcd/error"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api"
	"github.com/coreos/etcd/etcdserver/api/v2http/httptypes"
	"github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/store"
	"github.com/coreos/go-semver/semver"
	"github.com/jonboulle/clockwork"
)

func mustMarshalEvent(t *testing.T, ev *store.Event) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(ev); err != nil {
		t.Fatalf("error marshalling event %#v: %v", ev, err)
	}
	return b.String()
}
func mustNewForm(t *testing.T, p string, vals url.Values) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := testutil.MustNewURL(t, path.Join(keysPrefix, p))
	req, err := http.NewRequest("PUT", u.String(), strings.NewReader(vals.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatalf("error creating new request: %v", err)
	}
	return req
}
func mustNewPostForm(t *testing.T, p string, vals url.Values) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := testutil.MustNewURL(t, path.Join(keysPrefix, p))
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(vals.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatalf("error creating new request: %v", err)
	}
	return req
}
func mustNewRequest(t *testing.T, p string) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mustNewMethodRequest(t, "GET", p)
}
func mustNewMethodRequest(t *testing.T, m, p string) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &http.Request{Method: m, URL: testutil.MustNewURL(t, path.Join(keysPrefix, p))}
}

type fakeServer struct {
	dummyRaftTimer
	dummyStats
}

func (s *fakeServer) Leader() types.ID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.ID(1)
}
func (s *fakeServer) Alarms() []*etcdserverpb.AlarmMember {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *fakeServer) Cluster() api.Cluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *fakeServer) ClusterVersion() *semver.Version {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *fakeServer) RaftHandler() http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *fakeServer) Do(ctx context.Context, r etcdserverpb.Request) (rr etcdserver.Response, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return
}
func (s *fakeServer) ClientCertAuthEnabled() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}

type serverRecorder struct {
	fakeServer
	actions	[]action
}

func (s *serverRecorder) Do(_ context.Context, r etcdserverpb.Request) (etcdserver.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.actions = append(s.actions, action{name: "Do", params: []interface{}{r}})
	return etcdserver.Response{}, nil
}
func (s *serverRecorder) Process(_ context.Context, m raftpb.Message) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.actions = append(s.actions, action{name: "Process", params: []interface{}{m}})
	return nil
}
func (s *serverRecorder) AddMember(_ context.Context, m membership.Member) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.actions = append(s.actions, action{name: "AddMember", params: []interface{}{m}})
	return nil, nil
}
func (s *serverRecorder) RemoveMember(_ context.Context, id uint64) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.actions = append(s.actions, action{name: "RemoveMember", params: []interface{}{id}})
	return nil, nil
}
func (s *serverRecorder) UpdateMember(_ context.Context, m membership.Member) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.actions = append(s.actions, action{name: "UpdateMember", params: []interface{}{m}})
	return nil, nil
}

type action struct {
	name	string
	params	[]interface{}
}
type flushingRecorder struct {
	*httptest.ResponseRecorder
	ch	chan struct{}
}

func (fr *flushingRecorder) Flush() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fr.ResponseRecorder.Flush()
	fr.ch <- struct{}{}
}

type resServer struct {
	fakeServer
	res	etcdserver.Response
}

func (rs *resServer) Do(_ context.Context, _ etcdserverpb.Request) (etcdserver.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rs.res, nil
}
func (rs *resServer) Process(_ context.Context, _ raftpb.Message) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (rs *resServer) AddMember(_ context.Context, _ membership.Member) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func (rs *resServer) RemoveMember(_ context.Context, _ uint64) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func (rs *resServer) UpdateMember(_ context.Context, _ membership.Member) ([]*membership.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func boolp(b bool) *bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &b
}

type dummyRaftTimer struct{}

func (drt dummyRaftTimer) Index() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return uint64(100)
}
func (drt dummyRaftTimer) Term() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return uint64(5)
}

type dummyWatcher struct {
	echan	chan *store.Event
	sidx	uint64
}

func (w *dummyWatcher) EventChan() chan *store.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return w.echan
}
func (w *dummyWatcher) StartIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return w.sidx
}
func (w *dummyWatcher) Remove() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func TestBadRefreshRequest(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		in	*http.Request
		wcode	int
	}{{mustNewRequest(t, "foo?refresh=true&value=test"), etcdErr.EcodeRefreshValue}, {mustNewRequest(t, "foo?refresh=true&value=10"), etcdErr.EcodeRefreshValue}, {mustNewRequest(t, "foo?refresh=true"), etcdErr.EcodeRefreshTTLRequired}, {mustNewRequest(t, "foo?refresh=true&ttl="), etcdErr.EcodeRefreshTTLRequired}}
	for i, tt := range tests {
		got, _, err := parseKeyRequest(tt.in, clockwork.NewFakeClock())
		if err == nil {
			t.Errorf("#%d: unexpected nil error!", i)
			continue
		}
		ee, ok := err.(*etcdErr.Error)
		if !ok {
			t.Errorf("#%d: err is not etcd.Error!", i)
			continue
		}
		if ee.ErrorCode != tt.wcode {
			t.Errorf("#%d: code=%d, want %v", i, ee.ErrorCode, tt.wcode)
			t.Logf("cause: %#v", ee.Cause)
		}
		if !reflect.DeepEqual(got, etcdserverpb.Request{}) {
			t.Errorf("#%d: unexpected non-empty Request: %#v", i, got)
		}
	}
}
func TestBadParseRequest(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		in	*http.Request
		wcode	int
	}{{&http.Request{Body: nil, Method: "PUT"}, etcdErr.EcodeInvalidForm}, {&http.Request{URL: testutil.MustNewURL(t, "/badprefix/")}, etcdErr.EcodeInvalidForm}, {mustNewForm(t, "foo", url.Values{"prevIndex": []string{"garbage"}}), etcdErr.EcodeIndexNaN}, {mustNewForm(t, "foo", url.Values{"prevIndex": []string{"1.5"}}), etcdErr.EcodeIndexNaN}, {mustNewForm(t, "foo", url.Values{"prevIndex": []string{"-1"}}), etcdErr.EcodeIndexNaN}, {mustNewForm(t, "foo", url.Values{"waitIndex": []string{"garbage"}}), etcdErr.EcodeIndexNaN}, {mustNewForm(t, "foo", url.Values{"waitIndex": []string{"??"}}), etcdErr.EcodeIndexNaN}, {mustNewForm(t, "foo", url.Values{"ttl": []string{"-1"}}), etcdErr.EcodeTTLNaN}, {mustNewForm(t, "foo", url.Values{"recursive": []string{"hahaha"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"recursive": []string{"1234"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"recursive": []string{"?"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"sorted": []string{"?"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"sorted": []string{"x"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"wait": []string{"?!"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"wait": []string{"yes"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"prevExist": []string{"yes"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"prevExist": []string{"#2"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"dir": []string{"no"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"dir": []string{"file"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"quorum": []string{"no"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"quorum": []string{"file"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"stream": []string{"zzz"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"stream": []string{"something"}}), etcdErr.EcodeInvalidField}, {mustNewForm(t, "foo", url.Values{"prevValue": []string{""}}), etcdErr.EcodePrevValueRequired}, {mustNewMethodRequest(t, "HEAD", "foo?wait=true"), etcdErr.EcodeInvalidField}, {mustNewRequest(t, "foo?prevExist=wrong"), etcdErr.EcodeInvalidField}, {mustNewRequest(t, "foo?ttl=wrong"), etcdErr.EcodeTTLNaN}, {mustNewForm(t, "foo?ttl=12", url.Values{"ttl": []string{"garbage"}}), etcdErr.EcodeTTLNaN}, {mustNewForm(t, "foo?prevExist=false", url.Values{"prevExist": []string{"yes"}}), etcdErr.EcodeInvalidField}}
	for i, tt := range tests {
		got, _, err := parseKeyRequest(tt.in, clockwork.NewFakeClock())
		if err == nil {
			t.Errorf("#%d: unexpected nil error!", i)
			continue
		}
		ee, ok := err.(*etcdErr.Error)
		if !ok {
			t.Errorf("#%d: err is not etcd.Error!", i)
			continue
		}
		if ee.ErrorCode != tt.wcode {
			t.Errorf("#%d: code=%d, want %v", i, ee.ErrorCode, tt.wcode)
			t.Logf("cause: %#v", ee.Cause)
		}
		if !reflect.DeepEqual(got, etcdserverpb.Request{}) {
			t.Errorf("#%d: unexpected non-empty Request: %#v", i, got)
		}
	}
}
func TestGoodParseRequest(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fc := clockwork.NewFakeClock()
	fc.Advance(1111)
	tests := []struct {
		in	*http.Request
		w	etcdserverpb.Request
		noValue	bool
	}{{mustNewRequest(t, "foo"), etcdserverpb.Request{Method: "GET", Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo", url.Values{"value": []string{"some_value"}}), etcdserverpb.Request{Method: "PUT", Val: "some_value", Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo", url.Values{"prevIndex": []string{"98765"}}), etcdserverpb.Request{Method: "PUT", PrevIndex: 98765, Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo", url.Values{"recursive": []string{"true"}}), etcdserverpb.Request{Method: "PUT", Recursive: true, Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo", url.Values{"sorted": []string{"true"}}), etcdserverpb.Request{Method: "PUT", Sorted: true, Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo", url.Values{"quorum": []string{"true"}}), etcdserverpb.Request{Method: "PUT", Quorum: true, Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewRequest(t, "foo?wait=true"), etcdserverpb.Request{Method: "GET", Wait: true, Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewRequest(t, "foo?ttl="), etcdserverpb.Request{Method: "GET", Path: path.Join(etcdserver.StoreKeysPrefix, "/foo"), Expiration: 0}, false}, {mustNewRequest(t, "foo?ttl=5678"), etcdserverpb.Request{Method: "GET", Path: path.Join(etcdserver.StoreKeysPrefix, "/foo"), Expiration: fc.Now().Add(5678 * time.Second).UnixNano()}, false}, {mustNewRequest(t, "foo?ttl=0"), etcdserverpb.Request{Method: "GET", Path: path.Join(etcdserver.StoreKeysPrefix, "/foo"), Expiration: fc.Now().UnixNano()}, false}, {mustNewRequest(t, "foo?dir=true"), etcdserverpb.Request{Method: "GET", Dir: true, Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewRequest(t, "foo?dir=false"), etcdserverpb.Request{Method: "GET", Dir: false, Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo", url.Values{"prevExist": []string{"true"}}), etcdserverpb.Request{Method: "PUT", PrevExist: boolp(true), Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo", url.Values{"prevExist": []string{"false"}}), etcdserverpb.Request{Method: "PUT", PrevExist: boolp(false), Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo", url.Values{"value": []string{"some value"}, "prevExist": []string{"true"}, "prevValue": []string{"previous value"}}), etcdserverpb.Request{Method: "PUT", PrevExist: boolp(true), PrevValue: "previous value", Val: "some value", Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo?prevValue=woof", url.Values{}), etcdserverpb.Request{Method: "PUT", PrevValue: "woof", Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo?prevValue=woof", url.Values{"prevValue": []string{"miaow"}}), etcdserverpb.Request{Method: "PUT", PrevValue: "miaow", Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, false}, {mustNewForm(t, "foo", url.Values{"noValueOnSuccess": []string{"true"}}), etcdserverpb.Request{Method: "PUT", Path: path.Join(etcdserver.StoreKeysPrefix, "/foo")}, true}}
	for i, tt := range tests {
		got, noValueOnSuccess, err := parseKeyRequest(tt.in, fc)
		if err != nil {
			t.Errorf("#%d: err = %v, want %v", i, err, nil)
		}
		if noValueOnSuccess != tt.noValue {
			t.Errorf("#%d: noValue=%t, want %t", i, noValueOnSuccess, tt.noValue)
		}
		if !reflect.DeepEqual(got, tt.w) {
			t.Errorf("#%d: request=%#v, want %#v", i, got, tt.w)
		}
	}
}
func TestServeMembers(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	memb1 := membership.Member{ID: 12, Attributes: membership.Attributes{ClientURLs: []string{"http://localhost:8080"}}}
	memb2 := membership.Member{ID: 13, Attributes: membership.Attributes{ClientURLs: []string{"http://localhost:8081"}}}
	cluster := &fakeCluster{id: 1, members: map[uint64]*membership.Member{1: &memb1, 2: &memb2}}
	h := &membersHandler{server: &serverRecorder{}, clock: clockwork.NewFakeClock(), cluster: cluster}
	wmc := string(`{"members":[{"id":"c","name":"","peerURLs":[],"clientURLs":["http://localhost:8080"]},{"id":"d","name":"","peerURLs":[],"clientURLs":["http://localhost:8081"]}]}`)
	tests := []struct {
		path	string
		wcode	int
		wct	string
		wbody	string
	}{{membersPrefix, http.StatusOK, "application/json", wmc + "\n"}, {membersPrefix + "/", http.StatusOK, "application/json", wmc + "\n"}, {path.Join(membersPrefix, "100"), http.StatusNotFound, "application/json", `{"message":"Not found"}`}, {path.Join(membersPrefix, "foobar"), http.StatusNotFound, "application/json", `{"message":"Not found"}`}}
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
		gcid := rw.Header().Get("X-Etcd-Cluster-ID")
		wcid := cluster.ID().String()
		if gcid != wcid {
			t.Errorf("#%d: cid = %s, want %s", i, gcid, wcid)
		}
		if rw.Body.String() != tt.wbody {
			t.Errorf("#%d: body = %q, want %q", i, rw.Body.String(), tt.wbody)
		}
	}
}
func TestServeLeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	memb1 := membership.Member{ID: 1, Attributes: membership.Attributes{ClientURLs: []string{"http://localhost:8080"}}}
	memb2 := membership.Member{ID: 2, Attributes: membership.Attributes{ClientURLs: []string{"http://localhost:8081"}}}
	cluster := &fakeCluster{id: 1, members: map[uint64]*membership.Member{1: &memb1, 2: &memb2}}
	h := &membersHandler{server: &serverRecorder{}, clock: clockwork.NewFakeClock(), cluster: cluster}
	wmc := string(`{"id":"1","name":"","peerURLs":[],"clientURLs":["http://localhost:8080"]}`)
	tests := []struct {
		path	string
		wcode	int
		wct	string
		wbody	string
	}{{membersPrefix + "leader", http.StatusOK, "application/json", wmc + "\n"}}
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
		gcid := rw.Header().Get("X-Etcd-Cluster-ID")
		wcid := cluster.ID().String()
		if gcid != wcid {
			t.Errorf("#%d: cid = %s, want %s", i, gcid, wcid)
		}
		if rw.Body.String() != tt.wbody {
			t.Errorf("#%d: body = %q, want %q", i, rw.Body.String(), tt.wbody)
		}
	}
}
func TestServeMembersCreate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := testutil.MustNewURL(t, membersPrefix)
	b := []byte(`{"peerURLs":["http://127.0.0.1:1"]}`)
	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	s := &serverRecorder{}
	h := &membersHandler{server: s, clock: clockwork.NewFakeClock(), cluster: &fakeCluster{id: 1}}
	rw := httptest.NewRecorder()
	h.ServeHTTP(rw, req)
	wcode := http.StatusCreated
	if rw.Code != wcode {
		t.Errorf("code=%d, want %d", rw.Code, wcode)
	}
	wct := "application/json"
	if gct := rw.Header().Get("Content-Type"); gct != wct {
		t.Errorf("content-type = %s, want %s", gct, wct)
	}
	gcid := rw.Header().Get("X-Etcd-Cluster-ID")
	wcid := h.cluster.ID().String()
	if gcid != wcid {
		t.Errorf("cid = %s, want %s", gcid, wcid)
	}
	wb := `{"id":"c29b431f04be0bc7","name":"","peerURLs":["http://127.0.0.1:1"],"clientURLs":[]}` + "\n"
	g := rw.Body.String()
	if g != wb {
		t.Errorf("got body=%q, want %q", g, wb)
	}
	wm := membership.Member{ID: 14022875665250782151, RaftAttributes: membership.RaftAttributes{PeerURLs: []string{"http://127.0.0.1:1"}}}
	wactions := []action{{name: "AddMember", params: []interface{}{wm}}}
	if !reflect.DeepEqual(s.actions, wactions) {
		t.Errorf("actions = %+v, want %+v", s.actions, wactions)
	}
}
func TestServeMembersDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := &http.Request{Method: "DELETE", URL: testutil.MustNewURL(t, path.Join(membersPrefix, "BEEF"))}
	s := &serverRecorder{}
	h := &membersHandler{server: s, cluster: &fakeCluster{id: 1}}
	rw := httptest.NewRecorder()
	h.ServeHTTP(rw, req)
	wcode := http.StatusNoContent
	if rw.Code != wcode {
		t.Errorf("code=%d, want %d", rw.Code, wcode)
	}
	gcid := rw.Header().Get("X-Etcd-Cluster-ID")
	wcid := h.cluster.ID().String()
	if gcid != wcid {
		t.Errorf("cid = %s, want %s", gcid, wcid)
	}
	g := rw.Body.String()
	if g != "" {
		t.Errorf("got body=%q, want %q", g, "")
	}
	wactions := []action{{name: "RemoveMember", params: []interface{}{uint64(0xBEEF)}}}
	if !reflect.DeepEqual(s.actions, wactions) {
		t.Errorf("actions = %+v, want %+v", s.actions, wactions)
	}
}
func TestServeMembersUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := testutil.MustNewURL(t, path.Join(membersPrefix, "1"))
	b := []byte(`{"peerURLs":["http://127.0.0.1:1"]}`)
	req, err := http.NewRequest("PUT", u.String(), bytes.NewReader(b))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	s := &serverRecorder{}
	h := &membersHandler{server: s, clock: clockwork.NewFakeClock(), cluster: &fakeCluster{id: 1}}
	rw := httptest.NewRecorder()
	h.ServeHTTP(rw, req)
	wcode := http.StatusNoContent
	if rw.Code != wcode {
		t.Errorf("code=%d, want %d", rw.Code, wcode)
	}
	gcid := rw.Header().Get("X-Etcd-Cluster-ID")
	wcid := h.cluster.ID().String()
	if gcid != wcid {
		t.Errorf("cid = %s, want %s", gcid, wcid)
	}
	wm := membership.Member{ID: 1, RaftAttributes: membership.RaftAttributes{PeerURLs: []string{"http://127.0.0.1:1"}}}
	wactions := []action{{name: "UpdateMember", params: []interface{}{wm}}}
	if !reflect.DeepEqual(s.actions, wactions) {
		t.Errorf("actions = %+v, want %+v", s.actions, wactions)
	}
}
func TestServeMembersFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		req	*http.Request
		server	etcdserver.ServerV2
		wcode	int
	}{{&http.Request{Method: "CONNECT"}, &resServer{}, http.StatusMethodNotAllowed}, {&http.Request{Method: "TRACE"}, &resServer{}, http.StatusMethodNotAllowed}, {&http.Request{URL: testutil.MustNewURL(t, membersPrefix), Method: "POST", Body: ioutil.NopCloser(strings.NewReader("bad json")), Header: map[string][]string{"Content-Type": {"application/json"}}}, &resServer{}, http.StatusBadRequest}, {&http.Request{URL: testutil.MustNewURL(t, membersPrefix), Method: "POST", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://127.0.0.1:1"]}`)), Header: map[string][]string{"Content-Type": {"application/bad"}}}, &errServer{}, http.StatusUnsupportedMediaType}, {&http.Request{URL: testutil.MustNewURL(t, membersPrefix), Method: "POST", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://a"]}`)), Header: map[string][]string{"Content-Type": {"application/json"}}}, &errServer{}, http.StatusBadRequest}, {&http.Request{URL: testutil.MustNewURL(t, membersPrefix), Method: "POST", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://127.0.0.1:1"]}`)), Header: map[string][]string{"Content-Type": {"application/json"}}}, &errServer{err: errors.New("Error while adding a member")}, http.StatusInternalServerError}, {&http.Request{URL: testutil.MustNewURL(t, membersPrefix), Method: "POST", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://127.0.0.1:1"]}`)), Header: map[string][]string{"Content-Type": {"application/json"}}}, &errServer{err: membership.ErrIDExists}, http.StatusConflict}, {&http.Request{URL: testutil.MustNewURL(t, membersPrefix), Method: "POST", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://127.0.0.1:1"]}`)), Header: map[string][]string{"Content-Type": {"application/json"}}}, &errServer{err: membership.ErrPeerURLexists}, http.StatusConflict}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "1")), Method: "DELETE"}, &errServer{err: errors.New("Error while removing member")}, http.StatusInternalServerError}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "0")), Method: "DELETE"}, &errServer{err: membership.ErrIDRemoved}, http.StatusGone}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "0")), Method: "DELETE"}, &errServer{err: membership.ErrIDNotFound}, http.StatusNotFound}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "bad_id")), Method: "DELETE"}, nil, http.StatusNotFound}, {&http.Request{URL: testutil.MustNewURL(t, membersPrefix), Method: "DELETE"}, nil, http.StatusMethodNotAllowed}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "0")), Method: "PUT", Body: ioutil.NopCloser(strings.NewReader("bad json")), Header: map[string][]string{"Content-Type": {"application/json"}}}, &resServer{}, http.StatusBadRequest}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "0")), Method: "PUT", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://127.0.0.1:1"]}`)), Header: map[string][]string{"Content-Type": {"application/bad"}}}, &errServer{}, http.StatusUnsupportedMediaType}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "0")), Method: "PUT", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://a"]}`)), Header: map[string][]string{"Content-Type": {"application/json"}}}, &errServer{}, http.StatusBadRequest}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "0")), Method: "PUT", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://127.0.0.1:1"]}`)), Header: map[string][]string{"Content-Type": {"application/json"}}}, &errServer{err: errors.New("blah")}, http.StatusInternalServerError}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "0")), Method: "PUT", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://127.0.0.1:1"]}`)), Header: map[string][]string{"Content-Type": {"application/json"}}}, &errServer{err: membership.ErrPeerURLexists}, http.StatusConflict}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "0")), Method: "PUT", Body: ioutil.NopCloser(strings.NewReader(`{"PeerURLs": ["http://127.0.0.1:1"]}`)), Header: map[string][]string{"Content-Type": {"application/json"}}}, &errServer{err: membership.ErrIDNotFound}, http.StatusNotFound}, {&http.Request{URL: testutil.MustNewURL(t, path.Join(membersPrefix, "bad_id")), Method: "PUT"}, nil, http.StatusNotFound}, {&http.Request{URL: testutil.MustNewURL(t, membersPrefix), Method: "PUT"}, nil, http.StatusMethodNotAllowed}}
	for i, tt := range tests {
		h := &membersHandler{server: tt.server, cluster: &fakeCluster{id: 1}, clock: clockwork.NewFakeClock()}
		rw := httptest.NewRecorder()
		h.ServeHTTP(rw, tt.req)
		if rw.Code != tt.wcode {
			t.Errorf("#%d: code=%d, want %d", i, rw.Code, tt.wcode)
		}
		if rw.Code != http.StatusMethodNotAllowed {
			gcid := rw.Header().Get("X-Etcd-Cluster-ID")
			wcid := h.cluster.ID().String()
			if gcid != wcid {
				t.Errorf("#%d: cid = %s, want %s", i, gcid, wcid)
			}
		}
	}
}
func TestWriteEvent(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rec := httptest.NewRecorder()
	writeKeyEvent(rec, etcdserver.Response{}, false)
	h := rec.Header()
	if len(h) > 0 {
		t.Fatalf("unexpected non-empty headers: %#v", h)
	}
	b := rec.Body.String()
	if len(b) > 0 {
		t.Fatalf("unexpected non-empty body: %q", b)
	}
	tests := []struct {
		ev	*store.Event
		noValue	bool
		idx	string
		code	int
		err	error
	}{{&store.Event{Action: store.Get, Node: &store.NodeExtern{}, PrevNode: &store.NodeExtern{}}, false, "0", http.StatusOK, nil}, {&store.Event{Action: store.Create, Node: &store.NodeExtern{}, PrevNode: &store.NodeExtern{}}, false, "0", http.StatusCreated, nil}}
	for i, tt := range tests {
		rw := httptest.NewRecorder()
		resp := etcdserver.Response{Event: tt.ev, Term: 5, Index: 100}
		writeKeyEvent(rw, resp, tt.noValue)
		if gct := rw.Header().Get("Content-Type"); gct != "application/json" {
			t.Errorf("case %d: bad Content-Type: got %q, want application/json", i, gct)
		}
		if gri := rw.Header().Get("X-Raft-Index"); gri != "100" {
			t.Errorf("case %d: bad X-Raft-Index header: got %s, want %s", i, gri, "100")
		}
		if grt := rw.Header().Get("X-Raft-Term"); grt != "5" {
			t.Errorf("case %d: bad X-Raft-Term header: got %s, want %s", i, grt, "5")
		}
		if gei := rw.Header().Get("X-Etcd-Index"); gei != tt.idx {
			t.Errorf("case %d: bad X-Etcd-Index header: got %s, want %s", i, gei, tt.idx)
		}
		if rw.Code != tt.code {
			t.Errorf("case %d: bad response code: got %d, want %v", i, rw.Code, tt.code)
		}
	}
}
func TestV2DMachinesEndpoint(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		method	string
		wcode	int
	}{{"GET", http.StatusOK}, {"HEAD", http.StatusOK}, {"POST", http.StatusMethodNotAllowed}}
	m := &machinesHandler{cluster: &fakeCluster{}}
	s := httptest.NewServer(m)
	defer s.Close()
	for _, tt := range tests {
		req, err := http.NewRequest(tt.method, s.URL+machinesPrefix, nil)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != tt.wcode {
			t.Errorf("StatusCode = %d, expected %d", resp.StatusCode, tt.wcode)
		}
	}
}
func TestServeMachines(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := &fakeCluster{clientURLs: []string{"http://localhost:8080", "http://localhost:8081", "http://localhost:8082"}}
	writer := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	h := &machinesHandler{cluster: cluster}
	h.ServeHTTP(writer, req)
	w := "http://localhost:8080, http://localhost:8081, http://localhost:8082"
	if g := writer.Body.String(); g != w {
		t.Errorf("body = %s, want %s", g, w)
	}
	if writer.Code != http.StatusOK {
		t.Errorf("code = %d, want %d", writer.Code, http.StatusOK)
	}
}
func TestGetID(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		path	string
		wok	bool
		wid	types.ID
		wcode	int
	}{{"123", true, 0x123, http.StatusOK}, {"bad_id", false, 0, http.StatusNotFound}, {"", false, 0, http.StatusMethodNotAllowed}}
	for i, tt := range tests {
		w := httptest.NewRecorder()
		id, ok := getID(tt.path, w)
		if id != tt.wid {
			t.Errorf("#%d: id = %d, want %d", i, id, tt.wid)
		}
		if ok != tt.wok {
			t.Errorf("#%d: ok = %t, want %t", i, ok, tt.wok)
		}
		if w.Code != tt.wcode {
			t.Errorf("#%d code = %d, want %d", i, w.Code, tt.wcode)
		}
	}
}

type dummyStats struct{ data []byte }

func (ds *dummyStats) SelfStats() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ds.data
}
func (ds *dummyStats) LeaderStats() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ds.data
}
func (ds *dummyStats) StoreStats() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ds.data
}
func (ds *dummyStats) UpdateRecvApp(_ types.ID, _ int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func TestServeSelfStats(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wb := []byte("some statistics")
	w := string(wb)
	sh := &statsHandler{stats: &dummyStats{data: wb}}
	rw := httptest.NewRecorder()
	sh.serveSelf(rw, &http.Request{Method: "GET"})
	if rw.Code != http.StatusOK {
		t.Errorf("code = %d, want %d", rw.Code, http.StatusOK)
	}
	wct := "application/json"
	if gct := rw.Header().Get("Content-Type"); gct != wct {
		t.Errorf("Content-Type = %q, want %q", gct, wct)
	}
	if g := rw.Body.String(); g != w {
		t.Errorf("body = %s, want %s", g, w)
	}
}
func TestSelfServeStatsBad(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range []string{"PUT", "POST", "DELETE"} {
		sh := &statsHandler{}
		rw := httptest.NewRecorder()
		sh.serveSelf(rw, &http.Request{Method: m})
		if rw.Code != http.StatusMethodNotAllowed {
			t.Errorf("method %s: code=%d, want %d", m, rw.Code, http.StatusMethodNotAllowed)
		}
	}
}
func TestLeaderServeStatsBad(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range []string{"PUT", "POST", "DELETE"} {
		sh := &statsHandler{}
		rw := httptest.NewRecorder()
		sh.serveLeader(rw, &http.Request{Method: m})
		if rw.Code != http.StatusMethodNotAllowed {
			t.Errorf("method %s: code=%d, want %d", m, rw.Code, http.StatusMethodNotAllowed)
		}
	}
}
func TestServeLeaderStats(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wb := []byte("some statistics")
	w := string(wb)
	sh := &statsHandler{stats: &dummyStats{data: wb}}
	rw := httptest.NewRecorder()
	sh.serveLeader(rw, &http.Request{Method: "GET"})
	if rw.Code != http.StatusOK {
		t.Errorf("code = %d, want %d", rw.Code, http.StatusOK)
	}
	wct := "application/json"
	if gct := rw.Header().Get("Content-Type"); gct != wct {
		t.Errorf("Content-Type = %q, want %q", gct, wct)
	}
	if g := rw.Body.String(); g != w {
		t.Errorf("body = %s, want %s", g, w)
	}
}
func TestServeStoreStats(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wb := []byte("some statistics")
	w := string(wb)
	sh := &statsHandler{stats: &dummyStats{data: wb}}
	rw := httptest.NewRecorder()
	sh.serveStore(rw, &http.Request{Method: "GET"})
	if rw.Code != http.StatusOK {
		t.Errorf("code = %d, want %d", rw.Code, http.StatusOK)
	}
	wct := "application/json"
	if gct := rw.Header().Get("Content-Type"); gct != wct {
		t.Errorf("Content-Type = %q, want %q", gct, wct)
	}
	if g := rw.Body.String(); g != w {
		t.Errorf("body = %s, want %s", g, w)
	}
}
func TestBadServeKeys(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testBadCases := []struct {
		req	*http.Request
		server	etcdserver.ServerV2
		wcode	int
		wbody	string
	}{{&http.Request{Method: "CONNECT"}, &resServer{}, http.StatusMethodNotAllowed, "Method Not Allowed"}, {&http.Request{Method: "TRACE"}, &resServer{}, http.StatusMethodNotAllowed, "Method Not Allowed"}, {&http.Request{Body: nil, Method: "PUT"}, &resServer{}, http.StatusBadRequest, `{"errorCode":210,"message":"Invalid POST form","cause":"missing form body","index":0}`}, {mustNewRequest(t, "foo"), &errServer{err: errors.New("Internal Server Error")}, http.StatusInternalServerError, `{"errorCode":300,"message":"Raft Internal Error","cause":"Internal Server Error","index":0}`}, {mustNewRequest(t, "foo"), &errServer{err: etcdErr.NewError(etcdErr.EcodeKeyNotFound, "/1/pant", 0)}, http.StatusNotFound, `{"errorCode":100,"message":"Key not found","cause":"/pant","index":0}`}, {mustNewRequest(t, "foo"), &resServer{res: etcdserver.Response{}}, http.StatusInternalServerError, `{"errorCode":300,"message":"Raft Internal Error","cause":"received response with no Event/Watcher!","index":0}`}}
	for i, tt := range testBadCases {
		h := &keysHandler{timeout: 0, server: tt.server, cluster: &fakeCluster{id: 1}}
		rw := httptest.NewRecorder()
		h.ServeHTTP(rw, tt.req)
		if rw.Code != tt.wcode {
			t.Errorf("#%d: got code=%d, want %d", i, rw.Code, tt.wcode)
		}
		if rw.Code != http.StatusMethodNotAllowed {
			gcid := rw.Header().Get("X-Etcd-Cluster-ID")
			wcid := h.cluster.ID().String()
			if gcid != wcid {
				t.Errorf("#%d: cid = %s, want %s", i, gcid, wcid)
			}
		}
		if g := strings.TrimSuffix(rw.Body.String(), "\n"); g != tt.wbody {
			t.Errorf("#%d: body = %s, want %s", i, g, tt.wbody)
		}
	}
}
func TestServeKeysGood(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		req	*http.Request
		wcode	int
	}{{mustNewMethodRequest(t, "HEAD", "foo"), http.StatusOK}, {mustNewMethodRequest(t, "GET", "foo"), http.StatusOK}, {mustNewForm(t, "foo", url.Values{"value": []string{"bar"}}), http.StatusOK}, {mustNewMethodRequest(t, "DELETE", "foo"), http.StatusOK}, {mustNewPostForm(t, "foo", url.Values{"value": []string{"bar"}}), http.StatusOK}}
	server := &resServer{res: etcdserver.Response{Event: &store.Event{Action: store.Get, Node: &store.NodeExtern{}}}}
	for i, tt := range tests {
		h := &keysHandler{timeout: time.Hour, server: server, cluster: &fakeCluster{id: 1}}
		rw := httptest.NewRecorder()
		h.ServeHTTP(rw, tt.req)
		if rw.Code != tt.wcode {
			t.Errorf("#%d: got code=%d, want %d", i, rw.Code, tt.wcode)
		}
	}
}
func TestServeKeysEvent(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		req	*http.Request
		rsp	etcdserver.Response
		wcode	int
		event	*store.Event
	}{{mustNewRequest(t, "foo"), etcdserver.Response{Event: &store.Event{Action: store.Get, Node: &store.NodeExtern{}}}, http.StatusOK, &store.Event{Action: store.Get, Node: &store.NodeExtern{}}}, {mustNewForm(t, "foo", url.Values{"noValueOnSuccess": []string{"true"}}), etcdserver.Response{Event: &store.Event{Action: store.CompareAndSwap, Node: &store.NodeExtern{}}}, http.StatusOK, &store.Event{Action: store.CompareAndSwap, Node: nil}}}
	server := &resServer{}
	h := &keysHandler{timeout: time.Hour, server: server, cluster: &fakeCluster{id: 1}}
	for _, tt := range tests {
		server.res = tt.rsp
		rw := httptest.NewRecorder()
		h.ServeHTTP(rw, tt.req)
		wbody := mustMarshalEvent(t, tt.event)
		if rw.Code != tt.wcode {
			t.Errorf("got code=%d, want %d", rw.Code, tt.wcode)
		}
		gcid := rw.Header().Get("X-Etcd-Cluster-ID")
		wcid := h.cluster.ID().String()
		if gcid != wcid {
			t.Errorf("cid = %s, want %s", gcid, wcid)
		}
		g := rw.Body.String()
		if g != wbody {
			t.Errorf("got body=%#v, want %#v", g, wbody)
		}
	}
}
func TestServeKeysWatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := mustNewRequest(t, "/foo/bar")
	ec := make(chan *store.Event)
	dw := &dummyWatcher{echan: ec}
	server := &resServer{res: etcdserver.Response{Watcher: dw}}
	h := &keysHandler{timeout: time.Hour, server: server, cluster: &fakeCluster{id: 1}}
	go func() {
		ec <- &store.Event{Action: store.Get, Node: &store.NodeExtern{}}
	}()
	rw := httptest.NewRecorder()
	h.ServeHTTP(rw, req)
	wcode := http.StatusOK
	wbody := mustMarshalEvent(t, &store.Event{Action: store.Get, Node: &store.NodeExtern{}})
	if rw.Code != wcode {
		t.Errorf("got code=%d, want %d", rw.Code, wcode)
	}
	gcid := rw.Header().Get("X-Etcd-Cluster-ID")
	wcid := h.cluster.ID().String()
	if gcid != wcid {
		t.Errorf("cid = %s, want %s", gcid, wcid)
	}
	g := rw.Body.String()
	if g != wbody {
		t.Errorf("got body=%#v, want %#v", g, wbody)
	}
}

type recordingCloseNotifier struct {
	*httptest.ResponseRecorder
	cn	chan bool
}

func (rcn *recordingCloseNotifier) CloseNotify() <-chan bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rcn.cn
}
func TestHandleWatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defaultRwRr := func() (http.ResponseWriter, *httptest.ResponseRecorder) {
		r := httptest.NewRecorder()
		return r, r
	}
	noopEv := func(chan *store.Event) {
	}
	tests := []struct {
		getCtx		func() context.Context
		getRwRr		func() (http.ResponseWriter, *httptest.ResponseRecorder)
		doToChan	func(chan *store.Event)
		wbody		string
	}{{context.Background, defaultRwRr, func(ch chan *store.Event) {
		ch <- &store.Event{Action: store.Get, Node: &store.NodeExtern{}}
	}, mustMarshalEvent(t, &store.Event{Action: store.Get, Node: &store.NodeExtern{}})}, {context.Background, defaultRwRr, func(ch chan *store.Event) {
		close(ch)
	}, ""}, {func() context.Context {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		return ctx
	}, defaultRwRr, noopEv, ""}, {context.Background, func() (http.ResponseWriter, *httptest.ResponseRecorder) {
		rw := &recordingCloseNotifier{ResponseRecorder: httptest.NewRecorder(), cn: make(chan bool, 1)}
		rw.cn <- true
		return rw, rw.ResponseRecorder
	}, noopEv, ""}}
	for i, tt := range tests {
		rw, rr := tt.getRwRr()
		wa := &dummyWatcher{echan: make(chan *store.Event, 1), sidx: 10}
		tt.doToChan(wa.echan)
		resp := etcdserver.Response{Term: 5, Index: 100, Watcher: wa}
		handleKeyWatch(tt.getCtx(), rw, resp, false)
		wcode := http.StatusOK
		wct := "application/json"
		wei := "10"
		wri := "100"
		wrt := "5"
		if rr.Code != wcode {
			t.Errorf("#%d: got code=%d, want %d", i, rr.Code, wcode)
		}
		h := rr.Header()
		if ct := h.Get("Content-Type"); ct != wct {
			t.Errorf("#%d: Content-Type=%q, want %q", i, ct, wct)
		}
		if ei := h.Get("X-Etcd-Index"); ei != wei {
			t.Errorf("#%d: X-Etcd-Index=%q, want %q", i, ei, wei)
		}
		if ri := h.Get("X-Raft-Index"); ri != wri {
			t.Errorf("#%d: X-Raft-Index=%q, want %q", i, ri, wri)
		}
		if rt := h.Get("X-Raft-Term"); rt != wrt {
			t.Errorf("#%d: X-Raft-Term=%q, want %q", i, rt, wrt)
		}
		g := rr.Body.String()
		if g != tt.wbody {
			t.Errorf("#%d: got body=%#v, want %#v", i, g, tt.wbody)
		}
	}
}
func TestHandleWatchStreaming(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rw := &flushingRecorder{httptest.NewRecorder(), make(chan struct{}, 1)}
	wa := &dummyWatcher{echan: make(chan *store.Event)}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func() {
		resp := etcdserver.Response{Watcher: wa}
		handleKeyWatch(ctx, rw, resp, true)
		close(done)
	}()
	select {
	case <-rw.ch:
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for flush")
	}
	wcode := http.StatusOK
	wct := "application/json"
	wbody := ""
	if rw.Code != wcode {
		t.Errorf("got code=%d, want %d", rw.Code, wcode)
	}
	h := rw.Header()
	if ct := h.Get("Content-Type"); ct != wct {
		t.Errorf("Content-Type=%q, want %q", ct, wct)
	}
	g := rw.Body.String()
	if g != wbody {
		t.Errorf("got body=%#v, want %#v", g, wbody)
	}
	select {
	case wa.echan <- &store.Event{Action: store.Get, Node: &store.NodeExtern{}}:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for send")
	}
	select {
	case <-rw.ch:
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for flush")
	}
	wbody = mustMarshalEvent(t, &store.Event{Action: store.Get, Node: &store.NodeExtern{}})
	g = rw.Body.String()
	if g != wbody {
		t.Errorf("got body=%#v, want %#v", g, wbody)
	}
	select {
	case wa.echan <- &store.Event{Action: store.Get, Node: &store.NodeExtern{}}:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for send")
	}
	select {
	case <-rw.ch:
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for flush")
	}
	wbody = wbody + wbody
	g = rw.Body.String()
	if g != wbody {
		t.Errorf("got body=%#v, want %#v", g, wbody)
	}
	cancel()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for done")
	}
}
func TestTrimEventPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pre := "/abc"
	tests := []struct {
		ev	*store.Event
		wev	*store.Event
	}{{nil, nil}, {&store.Event{}, &store.Event{}}, {&store.Event{Node: &store.NodeExtern{Key: "/abc/def"}}, &store.Event{Node: &store.NodeExtern{Key: "/def"}}}, {&store.Event{PrevNode: &store.NodeExtern{Key: "/abc/ghi"}}, &store.Event{PrevNode: &store.NodeExtern{Key: "/ghi"}}}, {&store.Event{Node: &store.NodeExtern{Key: "/abc/def"}, PrevNode: &store.NodeExtern{Key: "/abc/ghi"}}, &store.Event{Node: &store.NodeExtern{Key: "/def"}, PrevNode: &store.NodeExtern{Key: "/ghi"}}}}
	for i, tt := range tests {
		ev := trimEventPrefix(tt.ev, pre)
		if !reflect.DeepEqual(ev, tt.wev) {
			t.Errorf("#%d: event = %+v, want %+v", i, ev, tt.wev)
		}
	}
}
func TestTrimNodeExternPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pre := "/abc"
	tests := []struct {
		n	*store.NodeExtern
		wn	*store.NodeExtern
	}{{nil, nil}, {&store.NodeExtern{Key: "/abc/def"}, &store.NodeExtern{Key: "/def"}}, {&store.NodeExtern{Key: "/abc/def", Nodes: []*store.NodeExtern{{Key: "/abc/def/1"}, {Key: "/abc/def/2"}}}, &store.NodeExtern{Key: "/def", Nodes: []*store.NodeExtern{{Key: "/def/1"}, {Key: "/def/2"}}}}}
	for i, tt := range tests {
		trimNodeExternPrefix(tt.n, pre)
		if !reflect.DeepEqual(tt.n, tt.wn) {
			t.Errorf("#%d: node = %+v, want %+v", i, tt.n, tt.wn)
		}
	}
}
func TestTrimPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		in	string
		prefix	string
		w	string
	}{{"/v2/members", "/v2/members", ""}, {"/v2/members/", "/v2/members", ""}, {"/v2/members/foo", "/v2/members", "foo"}}
	for i, tt := range tests {
		if g := trimPrefix(tt.in, tt.prefix); g != tt.w {
			t.Errorf("#%d: trimPrefix = %q, want %q", i, g, tt.w)
		}
	}
}
func TestNewMemberCollection(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fixture := []*membership.Member{{ID: 12, Attributes: membership.Attributes{ClientURLs: []string{"http://localhost:8080", "http://localhost:8081"}}, RaftAttributes: membership.RaftAttributes{PeerURLs: []string{"http://localhost:8082", "http://localhost:8083"}}}, {ID: 13, Attributes: membership.Attributes{ClientURLs: []string{"http://localhost:9090", "http://localhost:9091"}}, RaftAttributes: membership.RaftAttributes{PeerURLs: []string{"http://localhost:9092", "http://localhost:9093"}}}}
	got := newMemberCollection(fixture)
	want := httptypes.MemberCollection([]httptypes.Member{{ID: "c", ClientURLs: []string{"http://localhost:8080", "http://localhost:8081"}, PeerURLs: []string{"http://localhost:8082", "http://localhost:8083"}}, {ID: "d", ClientURLs: []string{"http://localhost:9090", "http://localhost:9091"}, PeerURLs: []string{"http://localhost:9092", "http://localhost:9093"}}})
	if !reflect.DeepEqual(&want, got) {
		t.Fatalf("newMemberCollection failure: want=%#v, got=%#v", &want, got)
	}
}
func TestNewMember(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fixture := &membership.Member{ID: 12, Attributes: membership.Attributes{ClientURLs: []string{"http://localhost:8080", "http://localhost:8081"}}, RaftAttributes: membership.RaftAttributes{PeerURLs: []string{"http://localhost:8082", "http://localhost:8083"}}}
	got := newMember(fixture)
	want := httptypes.Member{ID: "c", ClientURLs: []string{"http://localhost:8080", "http://localhost:8081"}, PeerURLs: []string{"http://localhost:8082", "http://localhost:8083"}}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("newMember failure: want=%#v, got=%#v", want, got)
	}
}
