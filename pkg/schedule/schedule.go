package schedule

import (
	"context"
	"sync"
)

type Job func(context.Context)
type Scheduler interface {
	Schedule(j Job)
	Pending() int
	Scheduled() int
	Finished() int
	WaitFinish(n int)
	Stop()
}
type fifo struct {
	mu		sync.Mutex
	resume		chan struct{}
	scheduled	int
	finished	int
	pendings	[]Job
	ctx		context.Context
	cancel		context.CancelFunc
	finishCond	*sync.Cond
	donec		chan struct{}
}

func NewFIFOScheduler() Scheduler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := &fifo{resume: make(chan struct{}, 1), donec: make(chan struct{}, 1)}
	f.finishCond = sync.NewCond(&f.mu)
	f.ctx, f.cancel = context.WithCancel(context.Background())
	go f.run()
	return f
}
func (f *fifo) Schedule(j Job) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.cancel == nil {
		panic("schedule: schedule to stopped scheduler")
	}
	if len(f.pendings) == 0 {
		select {
		case f.resume <- struct{}{}:
		default:
		}
	}
	f.pendings = append(f.pendings, j)
}
func (f *fifo) Pending() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.pendings)
}
func (f *fifo) Scheduled() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.scheduled
}
func (f *fifo) Finished() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.finishCond.L.Lock()
	defer f.finishCond.L.Unlock()
	return f.finished
}
func (f *fifo) WaitFinish(n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.finishCond.L.Lock()
	for f.finished < n || len(f.pendings) != 0 {
		f.finishCond.Wait()
	}
	f.finishCond.L.Unlock()
}
func (f *fifo) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.mu.Lock()
	f.cancel()
	f.cancel = nil
	f.mu.Unlock()
	<-f.donec
}
func (f *fifo) run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer func() {
		close(f.donec)
		close(f.resume)
	}()
	for {
		var todo Job
		f.mu.Lock()
		if len(f.pendings) != 0 {
			f.scheduled++
			todo = f.pendings[0]
		}
		f.mu.Unlock()
		if todo == nil {
			select {
			case <-f.resume:
			case <-f.ctx.Done():
				f.mu.Lock()
				pendings := f.pendings
				f.pendings = nil
				f.mu.Unlock()
				for _, todo := range pendings {
					todo(f.ctx)
				}
				return
			}
		} else {
			todo(f.ctx)
			f.finishCond.L.Lock()
			f.finished++
			f.pendings = f.pendings[1:]
			f.finishCond.Broadcast()
			f.finishCond.L.Unlock()
		}
	}
}
