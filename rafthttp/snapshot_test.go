package rafthttp

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
)

type strReaderCloser struct{ *strings.Reader }

func (s strReaderCloser) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func TestSnapshotSend(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		m	raftpb.Message
		rc	io.ReadCloser
		size	int64
		wsent	bool
		wfiles	int
	}{{m: raftpb.Message{Type: raftpb.MsgSnap, To: 1}, rc: strReaderCloser{strings.NewReader("hello")}, size: 5, wsent: true, wfiles: 1}, {m: raftpb.Message{Type: raftpb.MsgSnap, To: 1}, rc: &errReadCloser{fmt.Errorf("snapshot error")}, size: 1, wsent: false, wfiles: 0}, {m: raftpb.Message{Type: raftpb.MsgSnap, To: 1}, rc: strReaderCloser{strings.NewReader("hello")}, size: 10000, wsent: false, wfiles: 0}, {m: raftpb.Message{Type: raftpb.MsgSnap, To: 1}, rc: strReaderCloser{strings.NewReader("hello")}, size: 1, wsent: false, wfiles: 0}}
	for i, tt := range tests {
		sent, files := testSnapshotSend(t, snap.NewMessage(tt.m, tt.rc, tt.size))
		if tt.wsent != sent {
			t.Errorf("#%d: snapshot expected %v, got %v", i, tt.wsent, sent)
		}
		if tt.wfiles != len(files) {
			t.Fatalf("#%d: expected %d files, got %d files", i, tt.wfiles, len(files))
		}
	}
}
func testSnapshotSend(t *testing.T, sm *snap.Message) (bool, []os.FileInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d, err := ioutil.TempDir(os.TempDir(), "snapdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(d)
	r := &fakeRaft{}
	tr := &Transport{pipelineRt: &http.Transport{}, ClusterID: types.ID(1), Raft: r}
	ch := make(chan struct{}, 1)
	h := &syncHandler{newSnapshotHandler(tr, r, snap.New(d), types.ID(1)), ch}
	srv := httptest.NewServer(h)
	defer srv.Close()
	picker := mustNewURLPicker(t, []string{srv.URL})
	snapsend := newSnapshotSender(tr, picker, types.ID(1), newPeerStatus(types.ID(1)))
	defer snapsend.stop()
	snapsend.send(*sm)
	sent := false
	select {
	case <-time.After(time.Second):
		t.Fatalf("timed out sending snapshot")
	case sent = <-sm.CloseNotify():
	}
	<-ch
	files, rerr := ioutil.ReadDir(d)
	if rerr != nil {
		t.Fatal(rerr)
	}
	return sent, files
}

type errReadCloser struct{ err error }

func (s *errReadCloser) Read(p []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0, s.err
}
func (s *errReadCloser) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.err
}

type syncHandler struct {
	h	http.Handler
	ch	chan<- struct{}
}

func (sh *syncHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sh.h.ServeHTTP(w, r)
	sh.ch <- struct{}{}
}
