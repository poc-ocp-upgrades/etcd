package rafthttp

import (
	"fmt"
	"github.com/coreos/etcd/pkg/types"
	"sync"
	"time"
)

type failureType struct {
	source string
	action string
}
type peerStatus struct {
	id     types.ID
	mu     sync.Mutex
	active bool
	since  time.Time
}

func newPeerStatus(id types.ID) *peerStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &peerStatus{id: id}
}
func (s *peerStatus) activate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.active {
		plog.Infof("peer %s became active", s.id)
		s.active = true
		s.since = time.Now()
	}
}
func (s *peerStatus) deactivate(failure failureType, reason string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	defer s.mu.Unlock()
	msg := fmt.Sprintf("failed to %s %s on %s (%s)", failure.action, s.id, failure.source, reason)
	if s.active {
		plog.Errorf(msg)
		plog.Infof("peer %s became inactive (message send to peer failed)", s.id)
		s.active = false
		s.since = time.Time{}
		return
	}
	plog.Debugf(msg)
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
