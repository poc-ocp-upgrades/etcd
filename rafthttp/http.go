package rafthttp

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"time"
	pioutil "github.com/coreos/etcd/pkg/ioutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/version"
)

const (
	connReadLimitByte = 64 * 1024
)

var (
	RaftPrefix		= "/raft"
	ProbingPrefix		= path.Join(RaftPrefix, "probing")
	RaftStreamPrefix	= path.Join(RaftPrefix, "stream")
	RaftSnapshotPrefix	= path.Join(RaftPrefix, "snapshot")
	errIncompatibleVersion	= errors.New("incompatible version")
	errClusterIDMismatch	= errors.New("cluster ID mismatch")
)

type peerGetter interface{ Get(id types.ID) Peer }
type writerToResponse interface{ WriteTo(w http.ResponseWriter) }
type pipelineHandler struct {
	tr	Transporter
	r	Raft
	cid	types.ID
}

func newPipelineHandler(tr Transporter, r Raft, cid types.ID) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pipelineHandler{tr: tr, r: r, cid: cid}
}
func (h *pipelineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("X-Etcd-Cluster-ID", h.cid.String())
	if err := checkClusterCompatibilityFromHeader(r.Header, h.cid); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}
	addRemoteFromRequest(h.tr, r)
	limitedr := pioutil.NewLimitedBufferReader(r.Body, connReadLimitByte)
	b, err := ioutil.ReadAll(limitedr)
	if err != nil {
		plog.Errorf("failed to read raft message (%v)", err)
		http.Error(w, "error reading raft message", http.StatusBadRequest)
		recvFailures.WithLabelValues(r.RemoteAddr).Inc()
		return
	}
	var m raftpb.Message
	if err := m.Unmarshal(b); err != nil {
		plog.Errorf("failed to unmarshal raft message (%v)", err)
		http.Error(w, "error unmarshaling raft message", http.StatusBadRequest)
		recvFailures.WithLabelValues(r.RemoteAddr).Inc()
		return
	}
	receivedBytes.WithLabelValues(types.ID(m.From).String()).Add(float64(len(b)))
	if err := h.r.Process(context.TODO(), m); err != nil {
		switch v := err.(type) {
		case writerToResponse:
			v.WriteTo(w)
		default:
			plog.Warningf("failed to process raft message (%v)", err)
			http.Error(w, "error processing raft message", http.StatusInternalServerError)
			w.(http.Flusher).Flush()
			panic(err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type snapshotHandler struct {
	tr		Transporter
	r		Raft
	snapshotter	*snap.Snapshotter
	cid		types.ID
}

func newSnapshotHandler(tr Transporter, r Raft, snapshotter *snap.Snapshotter, cid types.ID) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &snapshotHandler{tr: tr, r: r, snapshotter: snapshotter, cid: cid}
}

const unknownSnapshotSender = "UNKNOWN_SNAPSHOT_SENDER"

func (h *snapshotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	start := time.Now()
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		snapshotReceiveFailures.WithLabelValues(unknownSnapshotSender).Inc()
		return
	}
	w.Header().Set("X-Etcd-Cluster-ID", h.cid.String())
	if err := checkClusterCompatibilityFromHeader(r.Header, h.cid); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		snapshotReceiveFailures.WithLabelValues(unknownSnapshotSender).Inc()
		return
	}
	addRemoteFromRequest(h.tr, r)
	dec := &messageDecoder{r: r.Body}
	m, err := dec.decodeLimit(uint64(1 << 63))
	from := types.ID(m.From).String()
	if err != nil {
		msg := fmt.Sprintf("failed to decode raft message (%v)", err)
		plog.Errorf(msg)
		http.Error(w, msg, http.StatusBadRequest)
		recvFailures.WithLabelValues(r.RemoteAddr).Inc()
		snapshotReceiveFailures.WithLabelValues(from).Inc()
		return
	}
	receivedBytes.WithLabelValues(from).Add(float64(m.Size()))
	if m.Type != raftpb.MsgSnap {
		plog.Errorf("unexpected raft message type %s on snapshot path", m.Type)
		http.Error(w, "wrong raft message type", http.StatusBadRequest)
		snapshotReceiveFailures.WithLabelValues(from).Inc()
		return
	}
	plog.Infof("receiving database snapshot [index:%d, from %s] ...", m.Snapshot.Metadata.Index, types.ID(m.From))
	n, err := h.snapshotter.SaveDBFrom(r.Body, m.Snapshot.Metadata.Index)
	if err != nil {
		msg := fmt.Sprintf("failed to save KV snapshot (%v)", err)
		plog.Error(msg)
		http.Error(w, msg, http.StatusInternalServerError)
		snapshotReceiveFailures.WithLabelValues(from).Inc()
		return
	}
	receivedBytes.WithLabelValues(from).Add(float64(n))
	plog.Infof("received and saved database snapshot [index: %d, from: %s] successfully", m.Snapshot.Metadata.Index, types.ID(m.From))
	if err := h.r.Process(context.TODO(), m); err != nil {
		switch v := err.(type) {
		case writerToResponse:
			v.WriteTo(w)
		default:
			msg := fmt.Sprintf("failed to process raft message (%v)", err)
			plog.Warningf(msg)
			http.Error(w, msg, http.StatusInternalServerError)
			snapshotReceiveFailures.WithLabelValues(from).Inc()
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
	snapshotReceive.WithLabelValues(from).Inc()
	snapshotReceiveSeconds.WithLabelValues(from).Observe(time.Since(start).Seconds())
}

type streamHandler struct {
	tr		*Transport
	peerGetter	peerGetter
	r		Raft
	id		types.ID
	cid		types.ID
}

func newStreamHandler(tr *Transport, pg peerGetter, r Raft, id, cid types.ID) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &streamHandler{tr: tr, peerGetter: pg, r: r, id: id, cid: cid}
}
func (h *streamHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("X-Server-Version", version.Version)
	w.Header().Set("X-Etcd-Cluster-ID", h.cid.String())
	if err := checkClusterCompatibilityFromHeader(r.Header, h.cid); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}
	var t streamType
	switch path.Dir(r.URL.Path) {
	case streamTypeMsgAppV2.endpoint():
		t = streamTypeMsgAppV2
	case streamTypeMessage.endpoint():
		t = streamTypeMessage
	default:
		plog.Debugf("ignored unexpected streaming request path %s", r.URL.Path)
		http.Error(w, "invalid path", http.StatusNotFound)
		return
	}
	fromStr := path.Base(r.URL.Path)
	from, err := types.IDFromString(fromStr)
	if err != nil {
		plog.Errorf("failed to parse from %s into ID (%v)", fromStr, err)
		http.Error(w, "invalid from", http.StatusNotFound)
		return
	}
	if h.r.IsIDRemoved(uint64(from)) {
		plog.Warningf("rejected the stream from peer %s since it was removed", from)
		http.Error(w, "removed member", http.StatusGone)
		return
	}
	p := h.peerGetter.Get(from)
	if p == nil {
		if urls := r.Header.Get("X-PeerURLs"); urls != "" {
			h.tr.AddRemote(from, strings.Split(urls, ","))
		}
		plog.Errorf("failed to find member %s in cluster %s", from, h.cid)
		http.Error(w, "error sender not found", http.StatusNotFound)
		return
	}
	wto := h.id.String()
	if gto := r.Header.Get("X-Raft-To"); gto != wto {
		plog.Errorf("streaming request ignored (ID mismatch got %s want %s)", gto, wto)
		http.Error(w, "to field mismatch", http.StatusPreconditionFailed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.(http.Flusher).Flush()
	c := newCloseNotifier()
	conn := &outgoingConn{t: t, Writer: w, Flusher: w.(http.Flusher), Closer: c}
	p.attachOutgoingConn(conn)
	<-c.closeNotify()
}
func checkClusterCompatibilityFromHeader(header http.Header, cid types.ID) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := checkVersionCompability(header.Get("X-Server-From"), serverVersion(header), minClusterVersion(header)); err != nil {
		plog.Errorf("request version incompatibility (%v)", err)
		return errIncompatibleVersion
	}
	if gcid := header.Get("X-Etcd-Cluster-ID"); gcid != cid.String() {
		plog.Errorf("request cluster ID mismatch (got %s want %s)", gcid, cid)
		return errClusterIDMismatch
	}
	return nil
}

type closeNotifier struct{ done chan struct{} }

func newCloseNotifier() *closeNotifier {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &closeNotifier{done: make(chan struct{})}
}
func (n *closeNotifier) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(n.done)
	return nil
}
func (n *closeNotifier) closeNotify() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.done
}
