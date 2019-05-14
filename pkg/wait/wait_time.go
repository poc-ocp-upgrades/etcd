package wait

import "sync"

type WaitTime interface {
	Wait(deadline uint64) <-chan struct{}
	Trigger(deadline uint64)
}

var closec chan struct{}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	closec = make(chan struct{})
	close(closec)
}

type timeList struct {
	l                   sync.Mutex
	lastTriggerDeadline uint64
	m                   map[uint64]chan struct{}
}

func NewTimeList() *timeList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &timeList{m: make(map[uint64]chan struct{})}
}
func (tl *timeList) Wait(deadline uint64) <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tl.l.Lock()
	defer tl.l.Unlock()
	if tl.lastTriggerDeadline >= deadline {
		return closec
	}
	ch := tl.m[deadline]
	if ch == nil {
		ch = make(chan struct{})
		tl.m[deadline] = ch
	}
	return ch
}
func (tl *timeList) Trigger(deadline uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tl.l.Lock()
	defer tl.l.Unlock()
	tl.lastTriggerDeadline = deadline
	for t, ch := range tl.m {
		if t <= deadline {
			delete(tl.m, t)
			close(ch)
		}
	}
}
