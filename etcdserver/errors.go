package etcdserver

import (
	"errors"
	"fmt"
)

var (
	ErrUnknownMethod		= errors.New("etcdserver: unknown method")
	ErrStopped			= errors.New("etcdserver: server stopped")
	ErrCanceled			= errors.New("etcdserver: request cancelled")
	ErrTimeout			= errors.New("etcdserver: request timed out")
	ErrTimeoutDueToLeaderFail	= errors.New("etcdserver: request timed out, possibly due to previous leader failure")
	ErrTimeoutDueToConnectionLost	= errors.New("etcdserver: request timed out, possibly due to connection lost")
	ErrTimeoutLeaderTransfer	= errors.New("etcdserver: request timed out, leader transfer took too long")
	ErrNotEnoughStartedMembers	= errors.New("etcdserver: re-configuration failed due to not enough started members")
	ErrNoLeader			= errors.New("etcdserver: no leader")
	ErrNotLeader			= errors.New("etcdserver: not leader")
	ErrRequestTooLarge		= errors.New("etcdserver: request is too large")
	ErrNoSpace			= errors.New("etcdserver: no space")
	ErrTooManyRequests		= errors.New("etcdserver: too many requests")
	ErrUnhealthy			= errors.New("etcdserver: unhealthy cluster")
	ErrKeyNotFound			= errors.New("etcdserver: key not found")
	ErrCorrupt			= errors.New("etcdserver: corrupt cluster")
)

type DiscoveryError struct {
	Op	string
	Err	error
}

func (e DiscoveryError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("failed to %s discovery cluster (%v)", e.Op, e.Err)
}
