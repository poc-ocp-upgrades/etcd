package rafthttp

import (
	"context"
	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

const (
	ConnReadTimeout     = 5 * time.Second
	ConnWriteTimeout    = 5 * time.Second
	recvBufSize         = 4096
	maxPendingProposals = 4096
	streamAppV2         = "streamMsgAppV2"
	streamMsg           = "streamMsg"
	pipelineMsg         = "pipeline"
	sendSnap            = "sendMsgSnap"
)

type Peer interface {
	send(m raftpb.Message)
	sendSnap(m snap.Message)
	update(urls types.URLs)
	attachOutgoingConn(conn *outgoingConn)
	activeSince() time.Time
	stop()
}
type peer struct {
	id             types.ID
	r              Raft
	status         *peerStatus
	picker         *urlPicker
	msgAppV2Writer *streamWriter
	writer         *streamWriter
	pipeline       *pipeline
	snapSender     *snapshotSender
	msgAppV2Reader *streamReader
	msgAppReader   *streamReader
	recvc          chan raftpb.Message
	propc          chan raftpb.Message
	mu             sync.Mutex
	paused         bool
	cancel         context.CancelFunc
	stopc          chan struct{}
}

func startPeer(transport *Transport, urls types.URLs, peerID types.ID, fs *stats.FollowerStats) *peer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plog.Infof("starting peer %s...", peerID)
	defer plog.Infof("started peer %s", peerID)
	status := newPeerStatus(peerID)
	picker := newURLPicker(urls)
	errorc := transport.ErrorC
	r := transport.Raft
	pipeline := &pipeline{peerID: peerID, tr: transport, picker: picker, status: status, followerStats: fs, raft: r, errorc: errorc}
	pipeline.start()
	p := &peer{id: peerID, r: r, status: status, picker: picker, msgAppV2Writer: startStreamWriter(peerID, status, fs, r), writer: startStreamWriter(peerID, status, fs, r), pipeline: pipeline, snapSender: newSnapshotSender(transport, picker, peerID, status), recvc: make(chan raftpb.Message, recvBufSize), propc: make(chan raftpb.Message, maxPendingProposals), stopc: make(chan struct{})}
	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel
	go func() {
		for {
			select {
			case mm := <-p.recvc:
				if err := r.Process(ctx, mm); err != nil {
					plog.Warningf("failed to process raft message (%v)", err)
				}
			case <-p.stopc:
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case mm := <-p.propc:
				if err := r.Process(ctx, mm); err != nil {
					plog.Warningf("failed to process raft message (%v)", err)
				}
			case <-p.stopc:
				return
			}
		}
	}()
	p.msgAppV2Reader = &streamReader{peerID: peerID, typ: streamTypeMsgAppV2, tr: transport, picker: picker, status: status, recvc: p.recvc, propc: p.propc, rl: rate.NewLimiter(transport.DialRetryFrequency, 1)}
	p.msgAppReader = &streamReader{peerID: peerID, typ: streamTypeMessage, tr: transport, picker: picker, status: status, recvc: p.recvc, propc: p.propc, rl: rate.NewLimiter(transport.DialRetryFrequency, 1)}
	p.msgAppV2Reader.start()
	p.msgAppReader.start()
	return p
}
func (p *peer) send(m raftpb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.mu.Lock()
	paused := p.paused
	p.mu.Unlock()
	if paused {
		return
	}
	writec, name := p.pick(m)
	select {
	case writec <- m:
	default:
		p.r.ReportUnreachable(m.To)
		if isMsgSnap(m) {
			p.r.ReportSnapshot(m.To, raft.SnapshotFailure)
		}
		if p.status.isActive() {
			plog.MergeWarningf("dropped internal raft message to %s since %s's sending buffer is full (bad/overloaded network)", p.id, name)
		}
		plog.Debugf("dropped %s to %s since %s's sending buffer is full", m.Type, p.id, name)
		sentFailures.WithLabelValues(types.ID(m.To).String()).Inc()
	}
}
func (p *peer) sendSnap(m snap.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	go p.snapSender.send(m)
}
func (p *peer) update(urls types.URLs) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.picker.update(urls)
}
func (p *peer) attachOutgoingConn(conn *outgoingConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ok bool
	switch conn.t {
	case streamTypeMsgAppV2:
		ok = p.msgAppV2Writer.attach(conn)
	case streamTypeMessage:
		ok = p.writer.attach(conn)
	default:
		plog.Panicf("unhandled stream type %s", conn.t)
	}
	if !ok {
		conn.Close()
	}
}
func (p *peer) activeSince() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.status.activeSince()
}
func (p *peer) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.mu.Lock()
	defer p.mu.Unlock()
	p.paused = true
	p.msgAppReader.pause()
	p.msgAppV2Reader.pause()
}
func (p *peer) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.mu.Lock()
	defer p.mu.Unlock()
	p.paused = false
	p.msgAppReader.resume()
	p.msgAppV2Reader.resume()
}
func (p *peer) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plog.Infof("stopping peer %s...", p.id)
	defer plog.Infof("stopped peer %s", p.id)
	close(p.stopc)
	p.cancel()
	p.msgAppV2Writer.stop()
	p.writer.stop()
	p.pipeline.stop()
	p.snapSender.stop()
	p.msgAppV2Reader.stop()
	p.msgAppReader.stop()
}
func (p *peer) pick(m raftpb.Message) (writec chan<- raftpb.Message, picked string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ok bool
	if isMsgSnap(m) {
		return p.pipeline.msgc, pipelineMsg
	} else if writec, ok = p.msgAppV2Writer.writec(); ok && isMsgApp(m) {
		return writec, streamAppV2
	} else if writec, ok = p.writer.writec(); ok {
		return writec, streamMsg
	}
	return p.pipeline.msgc, pipelineMsg
}
func isMsgApp(m raftpb.Message) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Type == raftpb.MsgApp
}
func isMsgSnap(m raftpb.Message) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.Type == raftpb.MsgSnap
}
