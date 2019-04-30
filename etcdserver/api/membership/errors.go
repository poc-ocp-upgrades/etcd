package membership

import (
	"errors"
	"go.etcd.io/etcd/etcdserver/api/v2error"
)

var (
	ErrIDRemoved		= errors.New("membership: ID removed")
	ErrIDExists		= errors.New("membership: ID exists")
	ErrIDNotFound		= errors.New("membership: ID not found")
	ErrPeerURLexists	= errors.New("membership: peerURL exists")
)

func isKeyNotFound(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e, ok := err.(*v2error.Error)
	return ok && e.ErrorCode == v2error.EcodeKeyNotFound
}
