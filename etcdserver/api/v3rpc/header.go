package v3rpc

import (
	"github.com/coreos/etcd/etcdserver"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

type header struct {
	clusterID	int64
	memberID	int64
	raftTimer	etcdserver.RaftTimer
	rev		func() int64
}

func newHeader(s *etcdserver.EtcdServer) header {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return header{clusterID: int64(s.Cluster().ID()), memberID: int64(s.ID()), raftTimer: s, rev: func() int64 {
		return s.KV().Rev()
	}}
}
func (h *header) fill(rh *pb.ResponseHeader) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rh == nil {
		plog.Panic("unexpected nil resp.Header")
	}
	rh.ClusterId = uint64(h.clusterID)
	rh.MemberId = uint64(h.memberID)
	rh.RaftTerm = h.raftTimer.Term()
	if rh.Revision == 0 {
		rh.Revision = h.rev()
	}
}
