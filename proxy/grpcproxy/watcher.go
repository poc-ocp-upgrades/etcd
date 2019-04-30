package grpcproxy

import (
	"time"
	"go.etcd.io/etcd/clientv3"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/mvcc"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

type watchRange struct{ key, end string }

func (wr *watchRange) valid() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(wr.end) == 0 || wr.end > wr.key || (wr.end[0] == 0 && len(wr.end) == 1)
}

type watcher struct {
	wr		watchRange
	filters		[]mvcc.FilterFunc
	progress	bool
	prevKV		bool
	id		int64
	nextrev		int64
	lastHeader	pb.ResponseHeader
	wps		*watchProxyStream
}

func (w *watcher) send(wr clientv3.WatchResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if wr.IsProgressNotify() && !w.progress {
		return
	}
	if w.nextrev > wr.Header.Revision && len(wr.Events) > 0 {
		return
	}
	if w.nextrev == 0 {
		w.nextrev = wr.Header.Revision + 1
	}
	events := make([]*mvccpb.Event, 0, len(wr.Events))
	var lastRev int64
	for i := range wr.Events {
		ev := (*mvccpb.Event)(wr.Events[i])
		if ev.Kv.ModRevision < w.nextrev {
			continue
		} else {
			lastRev = ev.Kv.ModRevision
		}
		filtered := false
		for _, filter := range w.filters {
			if filter(*ev) {
				filtered = true
				break
			}
		}
		if filtered {
			continue
		}
		if !w.prevKV {
			evCopy := *ev
			evCopy.PrevKv = nil
			ev = &evCopy
		}
		events = append(events, ev)
	}
	if lastRev >= w.nextrev {
		w.nextrev = lastRev + 1
	}
	if !wr.IsProgressNotify() && !wr.Created && len(events) == 0 && wr.CompactRevision == 0 {
		return
	}
	w.lastHeader = wr.Header
	w.post(&pb.WatchResponse{Header: &wr.Header, Created: wr.Created, CompactRevision: wr.CompactRevision, Canceled: wr.Canceled, WatchId: w.id, Events: events})
}
func (w *watcher) post(wr *pb.WatchResponse) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case w.wps.watchCh <- wr:
	case <-time.After(50 * time.Millisecond):
		w.wps.cancel()
		return false
	}
	return true
}
