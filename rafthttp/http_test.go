package rafthttp

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/pkg/pbutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/version"
)

func TestServeRaftPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCases := []struct {
		method		string
		body		io.Reader
		p		Raft
		clusterID	string
		wcode		int
	}{{"GET", bytes.NewReader(pbutil.MustMarshal(&raftpb.Message{})), &fakeRaft{}, "0", http.StatusMethodNotAllowed}, {"PUT", bytes.NewReader(pbutil.MustMarshal(&raftpb.Message{})), &fakeRaft{}, "0", http.StatusMethodNotAllowed}, {"DELETE", bytes.NewReader(pbutil.MustMarshal(&raftpb.Message{})), &fakeRaft{}, "0", http.StatusMethodNotAllowed}, {"POST", &errReader{}, &fakeRaft{}, "0", http.StatusBadRequest}, {"POST", strings.NewReader("malformed garbage"), &fakeRaft{}, "0", http.StatusBadRequest}, {"POST", bytes.NewReader(pbutil.MustMarshal(&raftpb.Message{})), &fakeRaft{}, "1", http.StatusPreconditionFailed}, {"POST", bytes.NewReader(pbutil.MustMarshal(&raftpb.Message{})), &fakeRaft{err: &resWriterToError{code: http.StatusForbidden}}, "0", http.StatusForbidden}, {"POST", bytes.NewReader(pbutil.MustMarshal(&raftpb.Message{})), &fakeRaft{err: &resWriterToError{code: http.StatusInternalServerError}}, "0", http.StatusInternalServerError}, {"POST", bytes.NewReader(pbutil.MustMarshal(&raftpb.Message{})), &fakeRaft{err: errors.New("blah")}, "0", http.StatusInternalServerError}, {"POST", bytes.NewReader(pbutil.MustMarshal(&raftpb.Message{})), &fakeRaft{}, "0", http.StatusNoContent}}
	for i, tt := range testCases {
		req, err := http.NewRequest(tt.method, "foo", tt.body)
		if err != nil {
			t.Fatalf("#%d: could not create request: %#v", i, err)
		}
		req.Header.Set("X-Etcd-Cluster-ID", tt.clusterID)
		req.Header.Set("X-Server-Version", version.Version)
		rw := httptest.NewRecorder()
		h := newPipelineHandler(NewNopTransporter(), tt.p, types.ID(0))
		donec := make(chan struct{})
		go func() {
			defer func() {
				recover()
				close(donec)
			}()
			h.ServeHTTP(rw, req)
		}()
		<-donec
		if rw.Code != tt.wcode {
			t.Errorf("#%d: got code=%d, want %d", i, rw.Code, tt.wcode)
		}
	}
}
func TestServeRaftStreamPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		path	string
		wtype	streamType
	}{{RaftStreamPrefix + "/message/1", streamTypeMessage}, {RaftStreamPrefix + "/msgapp/1", streamTypeMsgAppV2}}
	for i, tt := range tests {
		req, err := http.NewRequest("GET", "http://localhost:2380"+tt.path, nil)
		if err != nil {
			t.Fatalf("#%d: could not create request: %#v", i, err)
		}
		req.Header.Set("X-Etcd-Cluster-ID", "1")
		req.Header.Set("X-Server-Version", version.Version)
		req.Header.Set("X-Raft-To", "2")
		peer := newFakePeer()
		peerGetter := &fakePeerGetter{peers: map[types.ID]Peer{types.ID(1): peer}}
		tr := &Transport{}
		h := newStreamHandler(tr, peerGetter, &fakeRaft{}, types.ID(2), types.ID(1))
		rw := httptest.NewRecorder()
		go h.ServeHTTP(rw, req)
		var conn *outgoingConn
		select {
		case conn = <-peer.connc:
		case <-time.After(time.Second):
			t.Fatalf("#%d: failed to attach outgoingConn", i)
		}
		if g := rw.Header().Get("X-Server-Version"); g != version.Version {
			t.Errorf("#%d: X-Server-Version = %s, want %s", i, g, version.Version)
		}
		if conn.t != tt.wtype {
			t.Errorf("#%d: type = %s, want %s", i, conn.t, tt.wtype)
		}
		conn.Close()
	}
}
func TestServeRaftStreamPrefixBad(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	removedID := uint64(5)
	tests := []struct {
		method		string
		path		string
		clusterID	string
		remote		string
		wcode		int
	}{{"PUT", RaftStreamPrefix + "/message/1", "1", "1", http.StatusMethodNotAllowed}, {"POST", RaftStreamPrefix + "/message/1", "1", "1", http.StatusMethodNotAllowed}, {"DELETE", RaftStreamPrefix + "/message/1", "1", "1", http.StatusMethodNotAllowed}, {"GET", RaftStreamPrefix + "/strange/1", "1", "1", http.StatusNotFound}, {"GET", RaftStreamPrefix + "/strange", "1", "1", http.StatusNotFound}, {"GET", RaftStreamPrefix + "/message/2", "1", "1", http.StatusNotFound}, {"GET", RaftStreamPrefix + "/message/" + fmt.Sprint(removedID), "1", "1", http.StatusGone}, {"GET", RaftStreamPrefix + "/message/1", "2", "1", http.StatusPreconditionFailed}, {"GET", RaftStreamPrefix + "/message/1", "1", "2", http.StatusPreconditionFailed}}
	for i, tt := range tests {
		req, err := http.NewRequest(tt.method, "http://localhost:2380"+tt.path, nil)
		if err != nil {
			t.Fatalf("#%d: could not create request: %#v", i, err)
		}
		req.Header.Set("X-Etcd-Cluster-ID", tt.clusterID)
		req.Header.Set("X-Server-Version", version.Version)
		req.Header.Set("X-Raft-To", tt.remote)
		rw := httptest.NewRecorder()
		tr := &Transport{}
		peerGetter := &fakePeerGetter{peers: map[types.ID]Peer{types.ID(1): newFakePeer()}}
		r := &fakeRaft{removedID: removedID}
		h := newStreamHandler(tr, peerGetter, r, types.ID(1), types.ID(1))
		h.ServeHTTP(rw, req)
		if rw.Code != tt.wcode {
			t.Errorf("#%d: code = %d, want %d", i, rw.Code, tt.wcode)
		}
	}
}
func TestCloseNotifier(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := newCloseNotifier()
	select {
	case <-c.closeNotify():
		t.Fatalf("received unexpected close notification")
	default:
	}
	c.Close()
	select {
	case <-c.closeNotify():
	default:
		t.Fatalf("failed to get close notification")
	}
}

type errReader struct{}

func (er *errReader) Read(_ []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0, errors.New("some error")
}

type resWriterToError struct{ code int }

func (e *resWriterToError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ""
}
func (e *resWriterToError) WriteTo(w http.ResponseWriter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.WriteHeader(e.code)
}

type fakePeerGetter struct{ peers map[types.ID]Peer }

func (pg *fakePeerGetter) Get(id types.ID) Peer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pg.peers[id]
}

type fakePeer struct {
	msgs		[]raftpb.Message
	snapMsgs	[]snap.Message
	peerURLs	types.URLs
	connc		chan *outgoingConn
	paused		bool
}

func newFakePeer() *fakePeer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeURL, _ := url.Parse("http://localhost")
	return &fakePeer{connc: make(chan *outgoingConn, 1), peerURLs: types.URLs{*fakeURL}}
}
func (pr *fakePeer) send(m raftpb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.paused {
		return
	}
	pr.msgs = append(pr.msgs, m)
}
func (pr *fakePeer) sendSnap(m snap.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.paused {
		return
	}
	pr.snapMsgs = append(pr.snapMsgs, m)
}
func (pr *fakePeer) update(urls types.URLs) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.peerURLs = urls
}
func (pr *fakePeer) attachOutgoingConn(conn *outgoingConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.connc <- conn
}
func (pr *fakePeer) activeSince() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Time{}
}
func (pr *fakePeer) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (pr *fakePeer) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.paused = true
}
func (pr *fakePeer) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.paused = false
}
