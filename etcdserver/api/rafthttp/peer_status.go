package rafthttp

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"go.etcd.io/etcd/pkg/types"
	"go.uber.org/zap"
)

type failureType struct {
	source	string
	action	string
}
type peerStatus struct {
	lg	*zap.Logger
	local	types.ID
	id	types.ID
	mu	sync.Mutex
	active	bool
	since	time.Time
}

func newPeerStatus(lg *zap.Logger, local, id types.ID) *peerStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &peerStatus{lg: lg, local: local, id: id}
}
func (s *peerStatus) activate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.active {
		if s.lg != nil {
			s.lg.Info("peer became active", zap.String("peer-id", s.id.String()))
		} else {
			plog.Infof("peer %s became active", s.id)
		}
		s.active = true
		s.since = time.Now()
		activePeers.WithLabelValues(s.local.String(), s.id.String()).Inc()
	}
}
func (s *peerStatus) deactivate(failure failureType, reason string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	msg := fmt.Sprintf("failed to %s %s on %s (%s)", failure.action, s.id, failure.source, reason)
	if s.active {
		if s.lg != nil {
			s.lg.Warn("peer became inactive (message send to peer failed)", zap.String("peer-id", s.id.String()), zap.Error(errors.New(msg)))
		} else {
			plog.Errorf(msg)
			plog.Infof("peer %s became inactive (message send to peer failed)", s.id)
		}
		s.active = false
		s.since = time.Time{}
		activePeers.WithLabelValues(s.local.String(), s.id.String()).Dec()
		disconnectedPeers.WithLabelValues(s.local.String(), s.id.String()).Inc()
		return
	}
	if s.lg != nil {
		s.lg.Debug("peer deactivated again", zap.String("peer-id", s.id.String()), zap.Error(errors.New(msg)))
	}
}
func (s *peerStatus) isActive() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.active
}
func (s *peerStatus) activeSince() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.since
}
