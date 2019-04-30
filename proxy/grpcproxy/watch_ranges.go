package grpcproxy

import (
	"sync"
)

type watchRanges struct {
	wp	*watchProxy
	mu	sync.Mutex
	bcasts	map[watchRange]*watchBroadcasts
}

func newWatchRanges(wp *watchProxy) *watchRanges {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &watchRanges{wp: wp, bcasts: make(map[watchRange]*watchBroadcasts)}
}
func (wrs *watchRanges) add(w *watcher) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wrs.mu.Lock()
	defer wrs.mu.Unlock()
	if wbs := wrs.bcasts[w.wr]; wbs != nil {
		wbs.add(w)
		return
	}
	wbs := newWatchBroadcasts(wrs.wp)
	wrs.bcasts[w.wr] = wbs
	wbs.add(w)
}
func (wrs *watchRanges) delete(w *watcher) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wrs.mu.Lock()
	defer wrs.mu.Unlock()
	wbs, ok := wrs.bcasts[w.wr]
	if !ok {
		panic("deleting missing range")
	}
	if wbs.delete(w) == 0 {
		wbs.stop()
		delete(wrs.bcasts, w.wr)
	}
}
func (wrs *watchRanges) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wrs.mu.Lock()
	defer wrs.mu.Unlock()
	for _, wb := range wrs.bcasts {
		wb.stop()
	}
	wrs.bcasts = nil
}
