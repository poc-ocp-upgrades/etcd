package mvcc

import (
	"bytes"
	"errors"
	"sync"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

const AutoWatchID WatchID = 0

var (
	ErrWatcherNotExist	= errors.New("mvcc: watcher does not exist")
	ErrEmptyWatcherRange	= errors.New("mvcc: watcher range is empty")
	ErrWatcherDuplicateID	= errors.New("mvcc: duplicate watch ID provided on the WatchStream")
)

type WatchID int64
type FilterFunc func(e mvccpb.Event) bool
type WatchStream interface {
	Watch(id WatchID, key, end []byte, startRev int64, fcs ...FilterFunc) (WatchID, error)
	Chan() <-chan WatchResponse
	RequestProgress(id WatchID)
	Cancel(id WatchID) error
	Close()
	Rev() int64
}
type WatchResponse struct {
	WatchID		WatchID
	Events		[]mvccpb.Event
	Revision	int64
	CompactRevision	int64
}
type watchStream struct {
	watchable	watchable
	ch		chan WatchResponse
	mu		sync.Mutex
	nextID		WatchID
	closed		bool
	cancels		map[WatchID]cancelFunc
	watchers	map[WatchID]*watcher
}

func (ws *watchStream) Watch(id WatchID, key, end []byte, startRev int64, fcs ...FilterFunc) (WatchID, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(end) != 0 && bytes.Compare(key, end) != -1 {
		return -1, ErrEmptyWatcherRange
	}
	ws.mu.Lock()
	defer ws.mu.Unlock()
	if ws.closed {
		return -1, ErrEmptyWatcherRange
	}
	if id == AutoWatchID {
		for ws.watchers[ws.nextID] != nil {
			ws.nextID++
		}
		id = ws.nextID
		ws.nextID++
	} else if _, ok := ws.watchers[id]; ok {
		return -1, ErrWatcherDuplicateID
	}
	w, c := ws.watchable.watch(key, end, startRev, id, ws.ch, fcs...)
	ws.cancels[id] = c
	ws.watchers[id] = w
	return id, nil
}
func (ws *watchStream) Chan() <-chan WatchResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ws.ch
}
func (ws *watchStream) Cancel(id WatchID) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ws.mu.Lock()
	cancel, ok := ws.cancels[id]
	w := ws.watchers[id]
	ok = ok && !ws.closed
	ws.mu.Unlock()
	if !ok {
		return ErrWatcherNotExist
	}
	cancel()
	ws.mu.Lock()
	if ww := ws.watchers[id]; ww == w {
		delete(ws.cancels, id)
		delete(ws.watchers, id)
	}
	ws.mu.Unlock()
	return nil
}
func (ws *watchStream) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ws.mu.Lock()
	defer ws.mu.Unlock()
	for _, cancel := range ws.cancels {
		cancel()
	}
	ws.closed = true
	close(ws.ch)
	watchStreamGauge.Dec()
}
func (ws *watchStream) Rev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ws.mu.Lock()
	defer ws.mu.Unlock()
	return ws.watchable.rev()
}
func (ws *watchStream) RequestProgress(id WatchID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ws.mu.Lock()
	w, ok := ws.watchers[id]
	ws.mu.Unlock()
	if !ok {
		return
	}
	ws.watchable.progress(w)
}
