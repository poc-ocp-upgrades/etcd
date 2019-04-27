package store

import (
	"container/list"
	"path"
	"strings"
	"sync"
	"sync/atomic"
	etcdErr "github.com/coreos/etcd/error"
)

type watcherHub struct {
	count		int64
	mutex		sync.Mutex
	watchers	map[string]*list.List
	EventHistory	*EventHistory
}

func newWatchHub(capacity int) *watcherHub {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &watcherHub{watchers: make(map[string]*list.List), EventHistory: newEventHistory(capacity)}
}
func (wh *watcherHub) watch(key string, recursive, stream bool, index, storeIndex uint64) (Watcher, *etcdErr.Error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	reportWatchRequest()
	event, err := wh.EventHistory.scan(key, recursive, index)
	if err != nil {
		err.Index = storeIndex
		return nil, err
	}
	w := &watcher{eventChan: make(chan *Event, 100), recursive: recursive, stream: stream, sinceIndex: index, startIndex: storeIndex, hub: wh}
	wh.mutex.Lock()
	defer wh.mutex.Unlock()
	if event != nil {
		ne := event.Clone()
		ne.EtcdIndex = storeIndex
		w.eventChan <- ne
		return w, nil
	}
	l, ok := wh.watchers[key]
	var elem *list.Element
	if ok {
		elem = l.PushBack(w)
	} else {
		l = list.New()
		elem = l.PushBack(w)
		wh.watchers[key] = l
	}
	w.remove = func() {
		if w.removed {
			return
		}
		w.removed = true
		l.Remove(elem)
		atomic.AddInt64(&wh.count, -1)
		reportWatcherRemoved()
		if l.Len() == 0 {
			delete(wh.watchers, key)
		}
	}
	atomic.AddInt64(&wh.count, 1)
	reportWatcherAdded()
	return w, nil
}
func (wh *watcherHub) add(e *Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wh.EventHistory.addEvent(e)
}
func (wh *watcherHub) notify(e *Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e = wh.EventHistory.addEvent(e)
	segments := strings.Split(e.Node.Key, "/")
	currPath := "/"
	for _, segment := range segments {
		currPath = path.Join(currPath, segment)
		wh.notifyWatchers(e, currPath, false)
	}
}
func (wh *watcherHub) notifyWatchers(e *Event, nodePath string, deleted bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wh.mutex.Lock()
	defer wh.mutex.Unlock()
	l, ok := wh.watchers[nodePath]
	if ok {
		curr := l.Front()
		for curr != nil {
			next := curr.Next()
			w, _ := curr.Value.(*watcher)
			originalPath := (e.Node.Key == nodePath)
			if (originalPath || !isHidden(nodePath, e.Node.Key)) && w.notify(e, originalPath, deleted) {
				if !w.stream {
					w.removed = true
					l.Remove(curr)
					atomic.AddInt64(&wh.count, -1)
					reportWatcherRemoved()
				}
			}
			curr = next
		}
		if l.Len() == 0 {
			delete(wh.watchers, nodePath)
		}
	}
}
func (wh *watcherHub) clone() *watcherHub {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clonedHistory := wh.EventHistory.clone()
	return &watcherHub{EventHistory: clonedHistory}
}
func isHidden(watchPath, keyPath string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(watchPath) > len(keyPath) {
		return false
	}
	afterPath := path.Clean("/" + keyPath[len(watchPath):])
	return strings.Contains(afterPath, "/_")
}
