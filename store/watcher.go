package store

type Watcher interface {
	EventChan() chan *Event
	StartIndex() uint64
	Remove()
}
type watcher struct {
	eventChan	chan *Event
	stream		bool
	recursive	bool
	sinceIndex	uint64
	startIndex	uint64
	hub		*watcherHub
	removed		bool
	remove		func()
}

func (w *watcher) EventChan() chan *Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return w.eventChan
}
func (w *watcher) StartIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return w.startIndex
}
func (w *watcher) notify(e *Event, originalPath bool, deleted bool) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if (w.recursive || originalPath || deleted) && e.Index() >= w.sinceIndex {
		select {
		case w.eventChan <- e:
		default:
			w.remove()
		}
		return true
	}
	return false
}
func (w *watcher) Remove() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.hub.mutex.Lock()
	defer w.hub.mutex.Unlock()
	close(w.eventChan)
	if w.remove != nil {
		w.remove()
	}
}

type nopWatcher struct{}

func NewNopWatcher() Watcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &nopWatcher{}
}
func (w *nopWatcher) EventChan() chan *Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (w *nopWatcher) StartIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (w *nopWatcher) Remove() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
