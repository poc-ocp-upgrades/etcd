package rafthttp

import (
	"context"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
)

func BenchmarkSendingMsgApp(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := &Transport{ID: types.ID(1), ClusterID: types.ID(1), Raft: &fakeRaft{}, ServerStats: newServerStats(), LeaderStats: stats.NewLeaderStats("1")}
	tr.Start()
	srv := httptest.NewServer(tr.Handler())
	defer srv.Close()
	r := &countRaft{}
	tr2 := &Transport{ID: types.ID(2), ClusterID: types.ID(1), Raft: r, ServerStats: newServerStats(), LeaderStats: stats.NewLeaderStats("2")}
	tr2.Start()
	srv2 := httptest.NewServer(tr2.Handler())
	defer srv2.Close()
	tr.AddPeer(types.ID(2), []string{srv2.URL})
	defer tr.Stop()
	tr2.AddPeer(types.ID(1), []string{srv.URL})
	defer tr2.Stop()
	if !waitStreamWorking(tr.Get(types.ID(2)).(*peer)) {
		b.Fatalf("stream from 1 to 2 is not in work as expected")
	}
	b.ReportAllocs()
	b.SetBytes(64)
	b.ResetTimer()
	data := make([]byte, 64)
	for i := 0; i < b.N; i++ {
		tr.Send([]raftpb.Message{{Type: raftpb.MsgApp, From: 1, To: 2, Index: uint64(i), Entries: []raftpb.Entry{{Index: uint64(i + 1), Data: data}}}})
	}
	for r.count() != b.N {
		time.Sleep(time.Millisecond)
	}
	b.StopTimer()
}

type countRaft struct {
	mu	sync.Mutex
	cnt	int
}

func (r *countRaft) Process(ctx context.Context, m raftpb.Message) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cnt++
	return nil
}
func (r *countRaft) IsIDRemoved(id uint64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (r *countRaft) ReportUnreachable(id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *countRaft) ReportSnapshot(id uint64, status raft.SnapshotStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *countRaft) count() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.cnt
}
