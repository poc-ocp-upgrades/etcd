package clientv3

import (
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

type CompactOp struct {
	revision	int64
	physical	bool
}
type CompactOption func(*CompactOp)

func (op *CompactOp) applyCompactOpts(opts []CompactOption) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, opt := range opts {
		opt(op)
	}
}
func OpCompact(rev int64, opts ...CompactOption) CompactOp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := CompactOp{revision: rev}
	ret.applyCompactOpts(opts)
	return ret
}
func (op CompactOp) toRequest() *pb.CompactionRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pb.CompactionRequest{Revision: op.revision, Physical: op.physical}
}
func WithCompactPhysical() CompactOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *CompactOp) {
		op.physical = true
	}
}
