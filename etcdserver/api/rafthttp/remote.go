package rafthttp

import (
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/raft/raftpb"
	"go.uber.org/zap"
)

type remote struct {
	lg		*zap.Logger
	localID		types.ID
	id		types.ID
	status		*peerStatus
	pipeline	*pipeline
}

func startRemote(tr *Transport, urls types.URLs, id types.ID) *remote {
	_logClusterCodePath()
	defer _logClusterCodePath()
	picker := newURLPicker(urls)
	status := newPeerStatus(tr.Logger, tr.ID, id)
	pipeline := &pipeline{peerID: id, tr: tr, picker: picker, status: status, raft: tr.Raft, errorc: tr.ErrorC}
	pipeline.start()
	return &remote{lg: tr.Logger, localID: tr.ID, id: id, status: status, pipeline: pipeline}
}
func (g *remote) send(m raftpb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case g.pipeline.msgc <- m:
	default:
		if g.status.isActive() {
			if g.lg != nil {
				g.lg.Warn("dropped internal Raft message since sending buffer is full (overloaded network)", zap.String("message-type", m.Type.String()), zap.String("local-member-id", g.localID.String()), zap.String("from", types.ID(m.From).String()), zap.String("remote-peer-id", g.id.String()), zap.Bool("remote-peer-active", g.status.isActive()))
			} else {
				plog.MergeWarningf("dropped internal raft message to %s since sending buffer is full (bad/overloaded network)", g.id)
			}
		} else {
			if g.lg != nil {
				g.lg.Warn("dropped Raft message since sending buffer is full (overloaded network)", zap.String("message-type", m.Type.String()), zap.String("local-member-id", g.localID.String()), zap.String("from", types.ID(m.From).String()), zap.String("remote-peer-id", g.id.String()), zap.Bool("remote-peer-active", g.status.isActive()))
			} else {
				plog.Debugf("dropped %s to %s since sending buffer is full", m.Type, g.id)
			}
		}
		sentFailures.WithLabelValues(types.ID(m.To).String()).Inc()
	}
}
func (g *remote) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g.pipeline.stop()
}
func (g *remote) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g.stop()
}
func (g *remote) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	g.pipeline.start()
}
