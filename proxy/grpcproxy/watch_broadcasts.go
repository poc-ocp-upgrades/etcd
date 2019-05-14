package grpcproxy

import (
	"sync"
)

type watchBroadcasts struct {
	wp       *watchProxy
	mu       sync.Mutex
	bcasts   map[*watchBroadcast]struct{}
	watchers map[*watcher]*watchBroadcast
	updatec  chan *watchBroadcast
	donec    chan struct{}
}

const maxCoalesceReceivers = 5

func newWatchBroadcasts(wp *watchProxy) *watchBroadcasts {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wbs := &watchBroadcasts{wp: wp, bcasts: make(map[*watchBroadcast]struct{}), watchers: make(map[*watcher]*watchBroadcast), updatec: make(chan *watchBroadcast, 1), donec: make(chan struct{})}
	go func() {
		defer close(wbs.donec)
		for wb := range wbs.updatec {
			wbs.coalesce(wb)
		}
	}()
	return wbs
}
func (wbs *watchBroadcasts) coalesce(wb *watchBroadcast) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if wb.size() >= maxCoalesceReceivers {
		return
	}
	wbs.mu.Lock()
	for wbswb := range wbs.bcasts {
		if wbswb == wb {
			continue
		}
		wb.mu.Lock()
		wbswb.mu.Lock()
		if wb.nextrev >= wbswb.nextrev && wbswb.responses > 0 {
			for w := range wb.receivers {
				wbswb.receivers[w] = struct{}{}
				wbs.watchers[w] = wbswb
			}
			wb.receivers = nil
		}
		wbswb.mu.Unlock()
		wb.mu.Unlock()
		if wb.empty() {
			delete(wbs.bcasts, wb)
			wb.stop()
			break
		}
	}
	wbs.mu.Unlock()
}
func (wbs *watchBroadcasts) add(w *watcher) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wbs.mu.Lock()
	defer wbs.mu.Unlock()
	for wb := range wbs.bcasts {
		if wb.add(w) {
			wbs.watchers[w] = wb
			return
		}
	}
	wb := newWatchBroadcast(wbs.wp, w, wbs.update)
	wbs.watchers[w] = wb
	wbs.bcasts[wb] = struct{}{}
}
func (wbs *watchBroadcasts) delete(w *watcher) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wbs.mu.Lock()
	defer wbs.mu.Unlock()
	wb, ok := wbs.watchers[w]
	if !ok {
		panic("deleting missing watcher from broadcasts")
	}
	delete(wbs.watchers, w)
	wb.delete(w)
	if wb.empty() {
		delete(wbs.bcasts, wb)
		wb.stop()
	}
	return len(wbs.bcasts)
}
func (wbs *watchBroadcasts) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wbs.mu.Lock()
	for wb := range wbs.bcasts {
		wb.stop()
	}
	wbs.bcasts = nil
	close(wbs.updatec)
	wbs.mu.Unlock()
	<-wbs.donec
}
func (wbs *watchBroadcasts) update(wb *watchBroadcast) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case wbs.updatec <- wb:
	default:
	}
}
