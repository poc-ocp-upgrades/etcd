package mvcc

import (
	"fmt"
	"testing"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc/backend"
)

func BenchmarkKVWatcherMemoryUsage(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	be, tmpPath := backend.NewDefaultTmpBackend()
	watchable := newWatchableStore(be, &lease.FakeLessor{}, nil)
	defer cleanup(watchable, be, tmpPath)
	w := watchable.NewWatchStream()
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w.Watch([]byte(fmt.Sprint("foo", i)), nil, 0)
	}
}
