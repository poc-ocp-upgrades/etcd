package etcdserver

import (
	"sync/atomic"
)

type consistentIndex uint64

func (i *consistentIndex) setConsistentIndex(v uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	atomic.StoreUint64((*uint64)(i), v)
}
func (i *consistentIndex) ConsistentIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return atomic.LoadUint64((*uint64)(i))
}
