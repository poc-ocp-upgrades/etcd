package rafthttp

import (
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
)

type remote struct {
	id       types.ID
	status   *peerStatus
	pipeline *pipeline
}

func startRemote(tr *Transport, urls types.URLs, id types.ID) *remote {
	_logClusterCodePath()
	defer _logClusterCodePath()
	picker := newURLPicker(urls)
	status := newPeerStatus(id)
	pipeline := &pipeline{peerID: id, tr: tr, picker: picker, status: status, raft: tr.Raft, errorc: tr.ErrorC}
	pipeline.start()
	return &remote{id: id, status: status, pipeline: pipeline}
}
func (g *remote) send(m raftpb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case g.pipeline.msgc <- m:
	default:
		if g.status.isActive() {
			plog.MergeWarningf("dropped internal raft message to %s since sending buffer is full (bad/overloaded network)", g.id)
		}
		plog.Debugf("dropped %s to %s since sending buffer is full", m.Type, g.id)
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
