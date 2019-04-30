package rafthttp

import (
	"context"
	"sync"
	"time"
	"go.etcd.io/etcd/etcdserver/api/snap"
	stats "go.etcd.io/etcd/etcdserver/api/v2stats"
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/raft"
	"go.etcd.io/etcd/raft/raftpb"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

const (
	ConnReadTimeout		= 5 * time.Second
	ConnWriteTimeout	= 5 * time.Second
	recvBufSize		= 4096
	maxPendingProposals	= 4096
	streamAppV2		= "streamMsgAppV2"
	streamMsg		= "streamMsg"
	pipelineMsg		= "pipeline"
	sendSnap		= "sendMsgSnap"
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
	lg		*zap.Logger
	localID		types.ID
	id		types.ID
	r		Raft
	status		*peerStatus
	picker		*urlPicker
	msgAppV2Writer	*streamWriter
	writer		*streamWriter
	pipeline	*pipeline
	snapSender	*snapshotSender
	msgAppV2Reader	*streamReader
	msgAppReader	*streamReader
	recvc		chan raftpb.Message
	propc		chan raftpb.Message
	mu		sync.Mutex
	paused		bool
	cancel		context.CancelFunc
	stopc		chan struct{}
}

func startPeer(t *Transport, urls types.URLs, peerID types.ID, fs *stats.FollowerStats) *peer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.Logger != nil {
		t.Logger.Info("starting remote peer", zap.String("remote-peer-id", peerID.String()))
	} else {
		plog.Infof("starting peer %s...", peerID)
	}
	defer func() {
		if t.Logger != nil {
			t.Logger.Info("started remote peer", zap.String("remote-peer-id", peerID.String()))
		} else {
			plog.Infof("started peer %s", peerID)
		}
	}()
	status := newPeerStatus(t.Logger, t.ID, peerID)
	picker := newURLPicker(urls)
	errorc := t.ErrorC
	r := t.Raft
	pipeline := &pipeline{peerID: peerID, tr: t, picker: picker, status: status, followerStats: fs, raft: r, errorc: errorc}
	pipeline.start()
	p := &peer{lg: t.Logger, localID: t.ID, id: peerID, r: r, status: status, picker: picker, msgAppV2Writer: startStreamWriter(t.Logger, t.ID, peerID, status, fs, r), writer: startStreamWriter(t.Logger, t.ID, peerID, status, fs, r), pipeline: pipeline, snapSender: newSnapshotSender(t, picker, peerID, status), recvc: make(chan raftpb.Message, recvBufSize), propc: make(chan raftpb.Message, maxPendingProposals), stopc: make(chan struct{})}
	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel
	go func() {
		for {
			select {
			case mm := <-p.recvc:
				if err := r.Process(ctx, mm); err != nil {
					if t.Logger != nil {
						t.Logger.Warn("failed to process Raft message", zap.Error(err))
					} else {
						plog.Warningf("failed to process raft message (%v)", err)
					}
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
	p.msgAppV2Reader = &streamReader{lg: t.Logger, peerID: peerID, typ: streamTypeMsgAppV2, tr: t, picker: picker, status: status, recvc: p.recvc, propc: p.propc, rl: rate.NewLimiter(t.DialRetryFrequency, 1)}
	p.msgAppReader = &streamReader{lg: t.Logger, peerID: peerID, typ: streamTypeMessage, tr: t, picker: picker, status: status, recvc: p.recvc, propc: p.propc, rl: rate.NewLimiter(t.DialRetryFrequency, 1)}
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
			if p.lg != nil {
				p.lg.Warn("dropped internal Raft message since sending buffer is full (overloaded network)", zap.String("message-type", m.Type.String()), zap.String("local-member-id", p.localID.String()), zap.String("from", types.ID(m.From).String()), zap.String("remote-peer-id", p.id.String()), zap.Bool("remote-peer-active", p.status.isActive()))
			} else {
				plog.MergeWarningf("dropped internal raft message to %s since %s's sending buffer is full (bad/overloaded network)", p.id, name)
			}
		} else {
			if p.lg != nil {
				p.lg.Warn("dropped internal Raft message since sending buffer is full (overloaded network)", zap.String("message-type", m.Type.String()), zap.String("local-member-id", p.localID.String()), zap.String("from", types.ID(m.From).String()), zap.String("remote-peer-id", p.id.String()), zap.Bool("remote-peer-active", p.status.isActive()))
			} else {
				plog.Debugf("dropped %s to %s since %s's sending buffer is full", m.Type, p.id, name)
			}
		}
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
		if p.lg != nil {
			p.lg.Panic("unknown stream type", zap.String("type", conn.t.String()))
		} else {
			plog.Panicf("unhandled stream type %s", conn.t)
		}
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
	if p.lg != nil {
		p.lg.Info("stopping remote peer", zap.String("remote-peer-id", p.id.String()))
	} else {
		plog.Infof("stopping peer %s...", p.id)
	}
	defer func() {
		if p.lg != nil {
			p.lg.Info("stopped remote peer", zap.String("remote-peer-id", p.id.String()))
		} else {
			plog.Infof("stopped peer %s", p.id)
		}
	}()
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
