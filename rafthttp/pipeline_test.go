package rafthttp

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/version"
)

func TestPipelineSend(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := &roundTripperRecorder{rec: testutil.NewRecorderStream()}
	picker := mustNewURLPicker(t, []string{"http://localhost:2380"})
	tp := &Transport{pipelineRt: tr}
	p := startTestPipeline(tp, picker)
	p.msgc <- raftpb.Message{Type: raftpb.MsgApp}
	tr.rec.Wait(1)
	p.stop()
	if p.followerStats.Counts.Success != 1 {
		t.Errorf("success = %d, want 1", p.followerStats.Counts.Success)
	}
}
func TestPipelineKeepSendingWhenPostError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := &respRoundTripper{rec: testutil.NewRecorderStream(), err: fmt.Errorf("roundtrip error")}
	picker := mustNewURLPicker(t, []string{"http://localhost:2380"})
	tp := &Transport{pipelineRt: tr}
	p := startTestPipeline(tp, picker)
	defer p.stop()
	for i := 0; i < 50; i++ {
		p.msgc <- raftpb.Message{Type: raftpb.MsgApp}
	}
	_, err := tr.rec.Wait(50)
	if err != nil {
		t.Errorf("unexpected wait error %v", err)
	}
}
func TestPipelineExceedMaximumServing(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rt := newRoundTripperBlocker()
	picker := mustNewURLPicker(t, []string{"http://localhost:2380"})
	tp := &Transport{pipelineRt: rt}
	p := startTestPipeline(tp, picker)
	defer p.stop()
	for i := 0; i < connPerPipeline+pipelineBufSize; i++ {
		select {
		case p.msgc <- raftpb.Message{}:
		case <-time.After(time.Second):
			t.Errorf("failed to send out message")
		}
	}
	select {
	case p.msgc <- raftpb.Message{}:
		t.Errorf("unexpected message sendout")
	default:
	}
	rt.unblock()
	select {
	case p.msgc <- raftpb.Message{}:
	case <-time.After(time.Second):
		t.Errorf("failed to send out message")
	}
}
func TestPipelineSendFailed(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	picker := mustNewURLPicker(t, []string{"http://localhost:2380"})
	rt := newRespRoundTripper(0, errors.New("blah"))
	rt.rec = testutil.NewRecorderStream()
	tp := &Transport{pipelineRt: rt}
	p := startTestPipeline(tp, picker)
	p.msgc <- raftpb.Message{Type: raftpb.MsgApp}
	if _, err := rt.rec.Wait(1); err != nil {
		t.Fatal(err)
	}
	p.stop()
	if p.followerStats.Counts.Fail != 1 {
		t.Errorf("fail = %d, want 1", p.followerStats.Counts.Fail)
	}
}
func TestPipelinePost(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := &roundTripperRecorder{rec: &testutil.RecorderBuffered{}}
	picker := mustNewURLPicker(t, []string{"http://localhost:2380"})
	tp := &Transport{ClusterID: types.ID(1), pipelineRt: tr}
	p := startTestPipeline(tp, picker)
	if err := p.post([]byte("some data")); err != nil {
		t.Fatalf("unexpected post error: %v", err)
	}
	act, err := tr.rec.Wait(1)
	if err != nil {
		t.Fatal(err)
	}
	p.stop()
	req := act[0].Params[0].(*http.Request)
	if g := req.Method; g != "POST" {
		t.Errorf("method = %s, want %s", g, "POST")
	}
	if g := req.URL.String(); g != "http://localhost:2380/raft" {
		t.Errorf("url = %s, want %s", g, "http://localhost:2380/raft")
	}
	if g := req.Header.Get("Content-Type"); g != "application/protobuf" {
		t.Errorf("content type = %s, want %s", g, "application/protobuf")
	}
	if g := req.Header.Get("X-Server-Version"); g != version.Version {
		t.Errorf("version = %s, want %s", g, version.Version)
	}
	if g := req.Header.Get("X-Min-Cluster-Version"); g != version.MinClusterVersion {
		t.Errorf("min version = %s, want %s", g, version.MinClusterVersion)
	}
	if g := req.Header.Get("X-Etcd-Cluster-ID"); g != "1" {
		t.Errorf("cluster id = %s, want %s", g, "1")
	}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("unexpected ReadAll error: %v", err)
	}
	if string(b) != "some data" {
		t.Errorf("body = %s, want %s", b, "some data")
	}
}
func TestPipelinePostBad(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		u	string
		code	int
		err	error
	}{{"http://localhost:2380", 0, errors.New("blah")}, {"http://localhost:2380", http.StatusOK, nil}, {"http://localhost:2380", http.StatusCreated, nil}}
	for i, tt := range tests {
		picker := mustNewURLPicker(t, []string{tt.u})
		tp := &Transport{pipelineRt: newRespRoundTripper(tt.code, tt.err)}
		p := startTestPipeline(tp, picker)
		err := p.post([]byte("some data"))
		p.stop()
		if err == nil {
			t.Errorf("#%d: err = nil, want not nil", i)
		}
	}
}
func TestPipelinePostErrorc(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		u	string
		code	int
		err	error
	}{{"http://localhost:2380", http.StatusForbidden, nil}}
	for i, tt := range tests {
		picker := mustNewURLPicker(t, []string{tt.u})
		tp := &Transport{pipelineRt: newRespRoundTripper(tt.code, tt.err)}
		p := startTestPipeline(tp, picker)
		p.post([]byte("some data"))
		p.stop()
		select {
		case <-p.errorc:
		default:
			t.Fatalf("#%d: cannot receive from errorc", i)
		}
	}
}
func TestStopBlockedPipeline(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	picker := mustNewURLPicker(t, []string{"http://localhost:2380"})
	tp := &Transport{pipelineRt: newRoundTripperBlocker()}
	p := startTestPipeline(tp, picker)
	for i := 0; i < connPerPipeline*10; i++ {
		p.msgc <- raftpb.Message{}
	}
	done := make(chan struct{})
	go func() {
		p.stop()
		done <- struct{}{}
	}()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatalf("failed to stop pipeline in 1s")
	}
}

type roundTripperBlocker struct {
	unblockc	chan struct{}
	mu		sync.Mutex
	cancel		map[*http.Request]chan struct{}
}

func newRoundTripperBlocker() *roundTripperBlocker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &roundTripperBlocker{unblockc: make(chan struct{}), cancel: make(map[*http.Request]chan struct{})}
}
func (t *roundTripperBlocker) unblock() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(t.unblockc)
}
func (t *roundTripperBlocker) CancelRequest(req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	if c, ok := t.cancel[req]; ok {
		c <- struct{}{}
		delete(t.cancel, req)
	}
}

type respRoundTripper struct {
	mu	sync.Mutex
	rec	testutil.Recorder
	code	int
	header	http.Header
	err	error
}

func newRespRoundTripper(code int, err error) *respRoundTripper {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &respRoundTripper{code: code, err: err}
}
func (t *respRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.rec != nil {
		t.rec.Record(testutil.Action{Name: "req", Params: []interface{}{req}})
	}
	return &http.Response{StatusCode: t.code, Header: t.header, Body: &nopReadCloser{}}, t.err
}

type roundTripperRecorder struct{ rec testutil.Recorder }

func (t *roundTripperRecorder) RoundTrip(req *http.Request) (*http.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.rec != nil {
		t.rec.Record(testutil.Action{Name: "req", Params: []interface{}{req}})
	}
	return &http.Response{StatusCode: http.StatusNoContent, Body: &nopReadCloser{}}, nil
}

type nopReadCloser struct{}

func (n *nopReadCloser) Read(p []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0, io.EOF
}
func (n *nopReadCloser) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func startTestPipeline(tr *Transport, picker *urlPicker) *pipeline {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &pipeline{peerID: types.ID(1), tr: tr, picker: picker, status: newPeerStatus(types.ID(1)), raft: &fakeRaft{}, followerStats: &stats.FollowerStats{}, errorc: make(chan error, 1)}
	p.start()
	return p
}
