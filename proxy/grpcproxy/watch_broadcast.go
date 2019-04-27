package grpcproxy

import (
	"context"
	"sync"
	"github.com/coreos/etcd/clientv3"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

type watchBroadcast struct {
	cancel		context.CancelFunc
	donec		chan struct{}
	mu		sync.RWMutex
	nextrev		int64
	receivers	map[*watcher]struct{}
	responses	int
}

func newWatchBroadcast(wp *watchProxy, w *watcher, update func(*watchBroadcast)) *watchBroadcast {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cctx, cancel := context.WithCancel(wp.ctx)
	wb := &watchBroadcast{cancel: cancel, nextrev: w.nextrev, receivers: make(map[*watcher]struct{}), donec: make(chan struct{})}
	wb.add(w)
	go func() {
		defer close(wb.donec)
		opts := []clientv3.OpOption{clientv3.WithRange(w.wr.end), clientv3.WithProgressNotify(), clientv3.WithRev(wb.nextrev), clientv3.WithPrevKV(), clientv3.WithCreatedNotify()}
		cctx = withClientAuthToken(cctx, w.wps.stream.Context())
		wch := wp.cw.Watch(cctx, w.wr.key, opts...)
		for wr := range wch {
			wb.bcast(wr)
			update(wb)
		}
	}()
	return wb
}
func (wb *watchBroadcast) bcast(wr clientv3.WatchResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wb.mu.Lock()
	defer wb.mu.Unlock()
	if wb.responses > 0 || wb.nextrev == 0 {
		wb.nextrev = wr.Header.Revision + 1
	}
	wb.responses++
	for r := range wb.receivers {
		r.send(wr)
	}
	if len(wb.receivers) > 0 {
		eventsCoalescing.Add(float64(len(wb.receivers) - 1))
	}
}
func (wb *watchBroadcast) add(w *watcher) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wb.mu.Lock()
	defer wb.mu.Unlock()
	if wb.nextrev > w.nextrev || (wb.nextrev == 0 && w.nextrev != 0) {
		return false
	}
	if wb.responses == 0 {
		wb.receivers[w] = struct{}{}
		return true
	}
	ok := w.post(&pb.WatchResponse{Header: &pb.ResponseHeader{Revision: w.nextrev}, WatchId: w.id, Created: true})
	if !ok {
		return false
	}
	wb.receivers[w] = struct{}{}
	watchersCoalescing.Inc()
	return true
}
func (wb *watchBroadcast) delete(w *watcher) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wb.mu.Lock()
	defer wb.mu.Unlock()
	if _, ok := wb.receivers[w]; !ok {
		panic("deleting missing watcher from broadcast")
	}
	delete(wb.receivers, w)
	if len(wb.receivers) > 0 {
		watchersCoalescing.Dec()
	}
}
func (wb *watchBroadcast) size() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wb.mu.RLock()
	defer wb.mu.RUnlock()
	return len(wb.receivers)
}
func (wb *watchBroadcast) empty() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wb.size() == 0
}
func (wb *watchBroadcast) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !wb.empty() {
		watchersCoalescing.Sub(float64(wb.size() - 1))
	}
	wb.cancel()
	<-wb.donec
}
