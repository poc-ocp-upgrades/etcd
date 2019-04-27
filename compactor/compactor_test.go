package compactor

import (
	"context"
	"sync/atomic"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/testutil"
)

type fakeCompactable struct{ testutil.Recorder }

func (fc *fakeCompactable) Compact(ctx context.Context, r *pb.CompactionRequest) (*pb.CompactionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fc.Record(testutil.Action{Name: "c", Params: []interface{}{r}})
	return &pb.CompactionResponse{}, nil
}

type fakeRevGetter struct {
	testutil.Recorder
	rev	int64
}

func (fr *fakeRevGetter) Rev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fr.Record(testutil.Action{Name: "g"})
	rev := atomic.AddInt64(&fr.rev, 1)
	return rev
}
func (fr *fakeRevGetter) SetRev(rev int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	atomic.StoreInt64(&fr.rev, rev)
}
