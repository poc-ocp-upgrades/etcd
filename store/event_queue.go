package store

type eventQueue struct {
	Events   []*Event
	Size     int
	Front    int
	Back     int
	Capacity int
}

func (eq *eventQueue) insert(e *Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eq.Events[eq.Back] = e
	eq.Back = (eq.Back + 1) % eq.Capacity
	if eq.Size == eq.Capacity {
		eq.Front = (eq.Front + 1) % eq.Capacity
	} else {
		eq.Size++
	}
}
