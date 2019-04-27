package rafthttp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"
	"time"
	"golang.org/x/time/rate"
	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/version"
	"github.com/coreos/go-semver/semver"
)

func TestStreamWriterAttachOutgoingConn(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sw := startStreamWriter(types.ID(1), newPeerStatus(types.ID(1)), &stats.FollowerStats{}, &fakeRaft{})
	if _, ok := sw.writec(); ok {
		t.Errorf("initial working status = %v, want false", ok)
	}
	var wfc *fakeWriteFlushCloser
	for i := 0; i < 3; i++ {
		prevwfc := wfc
		wfc = newFakeWriteFlushCloser(nil)
		sw.attach(&outgoingConn{t: streamTypeMessage, Writer: wfc, Flusher: wfc, Closer: wfc})
		if prevwfc != nil {
			select {
			case <-prevwfc.closed:
			case <-time.After(time.Second):
				t.Errorf("#%d: close of previous connection timed out", i)
			}
		}
		msgc, _ := sw.writec()
		msgc <- raftpb.Message{}
		select {
		case <-wfc.writec:
		case <-time.After(time.Second):
			t.Errorf("#%d: failed to write to the underlying connection", i)
		}
		if _, ok := sw.writec(); !ok {
			t.Errorf("#%d: working status = %v, want true", i, ok)
		}
	}
	sw.stop()
	if _, ok := sw.writec(); ok {
		t.Errorf("working status after stop = %v, want false", ok)
	}
	if !wfc.Closed() {
		t.Errorf("failed to close the underlying connection")
	}
}
func TestStreamWriterAttachBadOutgoingConn(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sw := startStreamWriter(types.ID(1), newPeerStatus(types.ID(1)), &stats.FollowerStats{}, &fakeRaft{})
	defer sw.stop()
	wfc := newFakeWriteFlushCloser(errors.New("blah"))
	sw.attach(&outgoingConn{t: streamTypeMessage, Writer: wfc, Flusher: wfc, Closer: wfc})
	sw.msgc <- raftpb.Message{}
	select {
	case <-wfc.closed:
	case <-time.After(time.Second):
		t.Errorf("failed to close the underlying connection in time")
	}
	if _, ok := sw.writec(); ok {
		t.Errorf("working = %v, want false", ok)
	}
}
func TestStreamReaderDialRequest(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, tt := range []streamType{streamTypeMessage, streamTypeMsgAppV2} {
		tr := &roundTripperRecorder{rec: &testutil.RecorderBuffered{}}
		sr := &streamReader{peerID: types.ID(2), tr: &Transport{streamRt: tr, ClusterID: types.ID(1), ID: types.ID(1)}, picker: mustNewURLPicker(t, []string{"http://localhost:2380"}), ctx: context.Background()}
		sr.dial(tt)
		act, err := tr.rec.Wait(1)
		if err != nil {
			t.Fatal(err)
		}
		req := act[0].Params[0].(*http.Request)
		wurl := fmt.Sprintf("http://localhost:2380" + tt.endpoint() + "/1")
		if req.URL.String() != wurl {
			t.Errorf("#%d: url = %s, want %s", i, req.URL.String(), wurl)
		}
		if w := "GET"; req.Method != w {
			t.Errorf("#%d: method = %s, want %s", i, req.Method, w)
		}
		if g := req.Header.Get("X-Etcd-Cluster-ID"); g != "1" {
			t.Errorf("#%d: header X-Etcd-Cluster-ID = %s, want 1", i, g)
		}
		if g := req.Header.Get("X-Raft-To"); g != "2" {
			t.Errorf("#%d: header X-Raft-To = %s, want 2", i, g)
		}
	}
}
func TestStreamReaderDialResult(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		code	int
		err	error
		wok	bool
		whalt	bool
	}{{0, errors.New("blah"), false, false}, {http.StatusOK, nil, true, false}, {http.StatusMethodNotAllowed, nil, false, false}, {http.StatusNotFound, nil, false, false}, {http.StatusPreconditionFailed, nil, false, false}, {http.StatusGone, nil, false, true}}
	for i, tt := range tests {
		h := http.Header{}
		h.Add("X-Server-Version", version.Version)
		tr := &respRoundTripper{code: tt.code, header: h, err: tt.err}
		sr := &streamReader{peerID: types.ID(2), tr: &Transport{streamRt: tr, ClusterID: types.ID(1)}, picker: mustNewURLPicker(t, []string{"http://localhost:2380"}), errorc: make(chan error, 1), ctx: context.Background()}
		_, err := sr.dial(streamTypeMessage)
		if ok := err == nil; ok != tt.wok {
			t.Errorf("#%d: ok = %v, want %v", i, ok, tt.wok)
		}
		if halt := len(sr.errorc) > 0; halt != tt.whalt {
			t.Errorf("#%d: halt = %v, want %v", i, halt, tt.whalt)
		}
	}
}
func TestStreamReaderStopOnDial(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	h := http.Header{}
	h.Add("X-Server-Version", version.Version)
	tr := &respWaitRoundTripper{rrt: &respRoundTripper{code: http.StatusOK, header: h}}
	sr := &streamReader{peerID: types.ID(2), tr: &Transport{streamRt: tr, ClusterID: types.ID(1)}, picker: mustNewURLPicker(t, []string{"http://localhost:2380"}), errorc: make(chan error, 1), typ: streamTypeMessage, status: newPeerStatus(types.ID(2)), rl: rate.NewLimiter(rate.Every(100*time.Millisecond), 1)}
	tr.onResp = func() {
		go sr.stop()
		time.Sleep(10 * time.Millisecond)
	}
	sr.start()
	select {
	case <-sr.done:
	case <-time.After(time.Second):
		t.Fatal("streamReader did not stop in time")
	}
}

type respWaitRoundTripper struct {
	rrt	*respRoundTripper
	onResp	func()
}

func (t *respWaitRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := t.rrt.RoundTrip(req)
	resp.Body = newWaitReadCloser()
	t.onResp()
	return resp, err
}

type waitReadCloser struct{ closec chan struct{} }

func newWaitReadCloser() *waitReadCloser {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &waitReadCloser{make(chan struct{})}
}
func (wrc *waitReadCloser) Read(p []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	<-wrc.closec
	return 0, io.EOF
}
func (wrc *waitReadCloser) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(wrc.closec)
	return nil
}
func TestStreamReaderDialDetectUnsupport(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, typ := range []streamType{streamTypeMsgAppV2, streamTypeMessage} {
		tr := &respRoundTripper{code: http.StatusNotFound, header: http.Header{}}
		sr := &streamReader{peerID: types.ID(2), tr: &Transport{streamRt: tr, ClusterID: types.ID(1)}, picker: mustNewURLPicker(t, []string{"http://localhost:2380"}), ctx: context.Background()}
		_, err := sr.dial(typ)
		if err != errUnsupportedStreamType {
			t.Errorf("#%d: error = %v, want %v", i, err, errUnsupportedStreamType)
		}
	}
}
func TestStream(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	recvc := make(chan raftpb.Message, streamBufSize)
	propc := make(chan raftpb.Message, streamBufSize)
	msgapp := raftpb.Message{Type: raftpb.MsgApp, From: 2, To: 1, Term: 1, LogTerm: 1, Index: 3, Entries: []raftpb.Entry{{Term: 1, Index: 4}}}
	tests := []struct {
		t	streamType
		m	raftpb.Message
		wc	chan raftpb.Message
	}{{streamTypeMessage, raftpb.Message{Type: raftpb.MsgProp, To: 2}, propc}, {streamTypeMessage, msgapp, recvc}, {streamTypeMsgAppV2, msgapp, recvc}}
	for i, tt := range tests {
		h := &fakeStreamHandler{t: tt.t}
		srv := httptest.NewServer(h)
		defer srv.Close()
		sw := startStreamWriter(types.ID(1), newPeerStatus(types.ID(1)), &stats.FollowerStats{}, &fakeRaft{})
		defer sw.stop()
		h.sw = sw
		picker := mustNewURLPicker(t, []string{srv.URL})
		tr := &Transport{streamRt: &http.Transport{}, ClusterID: types.ID(1)}
		sr := &streamReader{peerID: types.ID(2), typ: tt.t, tr: tr, picker: picker, status: newPeerStatus(types.ID(2)), recvc: recvc, propc: propc, rl: rate.NewLimiter(rate.Every(100*time.Millisecond), 1)}
		sr.start()
		var writec chan<- raftpb.Message
		for {
			var ok bool
			if writec, ok = sw.writec(); ok {
				break
			}
			time.Sleep(time.Millisecond)
		}
		writec <- tt.m
		var m raftpb.Message
		select {
		case m = <-tt.wc:
		case <-time.After(time.Second):
			t.Fatalf("#%d: failed to receive message from the channel", i)
		}
		if !reflect.DeepEqual(m, tt.m) {
			t.Fatalf("#%d: message = %+v, want %+v", i, m, tt.m)
		}
		sr.stop()
	}
}
func TestCheckStreamSupport(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		v	*semver.Version
		t	streamType
		w	bool
	}{{semver.Must(semver.NewVersion("2.1.0")), streamTypeMsgAppV2, true}, {semver.Must(semver.NewVersion("2.1.9")), streamTypeMsgAppV2, true}, {semver.Must(semver.NewVersion("2.1.0-alpha")), streamTypeMsgAppV2, true}}
	for i, tt := range tests {
		if g := checkStreamSupport(tt.v, tt.t); g != tt.w {
			t.Errorf("#%d: check = %v, want %v", i, g, tt.w)
		}
	}
}

type fakeWriteFlushCloser struct {
	mu	sync.Mutex
	err	error
	written	int
	closed	chan struct{}
	writec	chan struct{}
}

func newFakeWriteFlushCloser(err error) *fakeWriteFlushCloser {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &fakeWriteFlushCloser{err: err, closed: make(chan struct{}), writec: make(chan struct{}, 1)}
}
func (wfc *fakeWriteFlushCloser) Write(p []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wfc.mu.Lock()
	defer wfc.mu.Unlock()
	select {
	case wfc.writec <- struct{}{}:
	default:
	}
	wfc.written += len(p)
	return len(p), wfc.err
}
func (wfc *fakeWriteFlushCloser) Flush() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (wfc *fakeWriteFlushCloser) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(wfc.closed)
	return wfc.err
}
func (wfc *fakeWriteFlushCloser) Written() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wfc.mu.Lock()
	defer wfc.mu.Unlock()
	return wfc.written
}
func (wfc *fakeWriteFlushCloser) Closed() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case <-wfc.closed:
		return true
	default:
		return false
	}
}

type fakeStreamHandler struct {
	t	streamType
	sw	*streamWriter
}

func (h *fakeStreamHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Header().Add("X-Server-Version", version.Version)
	w.(http.Flusher).Flush()
	c := newCloseNotifier()
	h.sw.attach(&outgoingConn{t: h.t, Writer: w, Flusher: w.(http.Flusher), Closer: c})
	<-c.closeNotify()
}
