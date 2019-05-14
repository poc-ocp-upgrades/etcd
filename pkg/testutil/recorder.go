package testutil

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Action struct {
	Name   string
	Params []interface{}
}
type Recorder interface {
	Record(a Action)
	Wait(n int) ([]Action, error)
	Action() []Action
	Chan() <-chan Action
}
type RecorderBuffered struct {
	sync.Mutex
	actions []Action
}

func (r *RecorderBuffered) Record(a Action) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.Lock()
	r.actions = append(r.actions, a)
	r.Unlock()
}
func (r *RecorderBuffered) Action() []Action {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.Lock()
	cpy := make([]Action, len(r.actions))
	copy(cpy, r.actions)
	r.Unlock()
	return cpy
}
func (r *RecorderBuffered) Wait(n int) (acts []Action, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	WaitSchedule()
	acts = r.Action()
	if len(acts) < n {
		err = newLenErr(n, len(acts))
	}
	return acts, err
}
func (r *RecorderBuffered) Chan() <-chan Action {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ch := make(chan Action)
	go func() {
		acts := r.Action()
		for i := range acts {
			ch <- acts[i]
		}
		close(ch)
	}()
	return ch
}

type recorderStream struct{ ch chan Action }

func NewRecorderStream() Recorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &recorderStream{ch: make(chan Action)}
}
func (r *recorderStream) Record(a Action) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.ch <- a
}
func (r *recorderStream) Action() (acts []Action) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		select {
		case act := <-r.ch:
			acts = append(acts, act)
		default:
			return acts
		}
	}
}
func (r *recorderStream) Chan() <-chan Action {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.ch
}
func (r *recorderStream) Wait(n int) ([]Action, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	acts := make([]Action, n)
	timeoutC := time.After(5 * time.Second)
	for i := 0; i < n; i++ {
		select {
		case acts[i] = <-r.ch:
		case <-timeoutC:
			acts = acts[:i]
			return acts, newLenErr(n, i)
		}
	}
	select {
	case act := <-r.ch:
		acts = append(acts, act)
	case <-time.After(10 * time.Millisecond):
	}
	return acts, nil
}
func newLenErr(expected int, actual int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := fmt.Sprintf("len(actions) = %d, expected >= %d", actual, expected)
	return errors.New(s)
}
