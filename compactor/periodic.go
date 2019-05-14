package compactor

import (
	"context"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/mvcc"
	"github.com/jonboulle/clockwork"
	"sync"
	"time"
)

type Periodic struct {
	clock  clockwork.Clock
	period time.Duration
	rg     RevGetter
	c      Compactable
	revs   []int64
	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.RWMutex
	paused bool
}

func NewPeriodic(h time.Duration, rg RevGetter, c Compactable) *Periodic {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newPeriodic(clockwork.NewRealClock(), h, rg, c)
}
func newPeriodic(clock clockwork.Clock, h time.Duration, rg RevGetter, c Compactable) *Periodic {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := &Periodic{clock: clock, period: h, rg: rg, c: c, revs: make([]int64, 0)}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	return t
}
func (t *Periodic) Run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	compactInterval := t.getCompactInterval()
	retryInterval := t.getRetryInterval()
	retentions := t.getRetentions()
	go func() {
		lastSuccess := t.clock.Now()
		baseInterval := t.period
		for {
			t.revs = append(t.revs, t.rg.Rev())
			if len(t.revs) > retentions {
				t.revs = t.revs[1:]
			}
			select {
			case <-t.ctx.Done():
				return
			case <-t.clock.After(retryInterval):
				t.mu.Lock()
				p := t.paused
				t.mu.Unlock()
				if p {
					continue
				}
			}
			if t.clock.Now().Sub(lastSuccess) < baseInterval {
				continue
			}
			if baseInterval == t.period {
				baseInterval = compactInterval
			}
			rev := t.revs[0]
			plog.Noticef("Starting auto-compaction at revision %d (retention: %v)", rev, t.period)
			_, err := t.c.Compact(t.ctx, &pb.CompactionRequest{Revision: rev})
			if err == nil || err == mvcc.ErrCompacted {
				lastSuccess = t.clock.Now()
				plog.Noticef("Finished auto-compaction at revision %d", rev)
			} else {
				plog.Noticef("Failed auto-compaction at revision %d (%v)", rev, err)
				plog.Noticef("Retry after %v", retryInterval)
			}
		}
	}()
}
func (t *Periodic) getCompactInterval() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	itv := t.period
	if itv > time.Hour {
		itv = time.Hour
	}
	return itv
}
func (t *Periodic) getRetentions() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int(t.period/t.getRetryInterval()) + 1
}

const retryDivisor = 10

func (t *Periodic) getRetryInterval() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	itv := t.period
	if itv > time.Hour {
		itv = time.Hour
	}
	return itv / retryDivisor
}
func (t *Periodic) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.cancel()
}
func (t *Periodic) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.paused = true
}
func (t *Periodic) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.paused = false
}
