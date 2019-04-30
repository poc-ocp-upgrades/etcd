package v2store

import (
	"fmt"
	"path"
	"strings"
	"sync"
	"go.etcd.io/etcd/etcdserver/api/v2error"
)

type EventHistory struct {
	Queue		eventQueue
	StartIndex	uint64
	LastIndex	uint64
	rwl		sync.RWMutex
}

func newEventHistory(capacity int) *EventHistory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &EventHistory{Queue: eventQueue{Capacity: capacity, Events: make([]*Event, capacity)}}
}
func (eh *EventHistory) addEvent(e *Event) *Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eh.rwl.Lock()
	defer eh.rwl.Unlock()
	eh.Queue.insert(e)
	eh.LastIndex = e.Index()
	eh.StartIndex = eh.Queue.Events[eh.Queue.Front].Index()
	return e
}
func (eh *EventHistory) scan(key string, recursive bool, index uint64) (*Event, *v2error.Error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eh.rwl.RLock()
	defer eh.rwl.RUnlock()
	if index < eh.StartIndex {
		return nil, v2error.NewError(v2error.EcodeEventIndexCleared, fmt.Sprintf("the requested history has been cleared [%v/%v]", eh.StartIndex, index), 0)
	}
	if index > eh.LastIndex {
		return nil, nil
	}
	offset := index - eh.StartIndex
	i := (eh.Queue.Front + int(offset)) % eh.Queue.Capacity
	for {
		e := eh.Queue.Events[i]
		if !e.Refresh {
			ok := e.Node.Key == key
			if recursive {
				nkey := path.Clean(key)
				if nkey[len(nkey)-1] != '/' {
					nkey = nkey + "/"
				}
				ok = ok || strings.HasPrefix(e.Node.Key, nkey)
			}
			if (e.Action == Delete || e.Action == Expire) && e.PrevNode != nil && e.PrevNode.Dir {
				ok = ok || strings.HasPrefix(key, e.PrevNode.Key)
			}
			if ok {
				return e, nil
			}
		}
		i = (i + 1) % eh.Queue.Capacity
		if i == eh.Queue.Back {
			return nil, nil
		}
	}
}
func (eh *EventHistory) clone() *EventHistory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clonedQueue := eventQueue{Capacity: eh.Queue.Capacity, Events: make([]*Event, eh.Queue.Capacity), Size: eh.Queue.Size, Front: eh.Queue.Front, Back: eh.Queue.Back}
	copy(clonedQueue.Events, eh.Queue.Events)
	return &EventHistory{StartIndex: eh.StartIndex, Queue: clonedQueue, LastIndex: eh.LastIndex}
}
