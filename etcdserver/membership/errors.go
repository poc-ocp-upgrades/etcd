package membership

import (
	"errors"
	etcdErr "github.com/coreos/etcd/error"
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	e, ok := err.(*etcdErr.Error)
	return ok && e.ErrorCode == etcdErr.EcodeKeyNotFound
}
