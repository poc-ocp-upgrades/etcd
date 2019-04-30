package v3compactor

import (
	"context"
	"sync"
	"time"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/mvcc"
	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"
)

type Periodic struct {
	lg	*zap.Logger
	clock	clockwork.Clock
	period	time.Duration
	rg	RevGetter
	c	Compactable
	revs	[]int64
	ctx	context.Context
	cancel	context.CancelFunc
	mu	sync.RWMutex
	paused	bool
}

func newPeriodic(lg *zap.Logger, clock clockwork.Clock, h time.Duration, rg RevGetter, c Compactable) *Periodic {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc := &Periodic{lg: lg, clock: clock, period: h, rg: rg, c: c, revs: make([]int64, 0)}
	pc.ctx, pc.cancel = context.WithCancel(context.Background())
	return pc
}
func (pc *Periodic) Run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	compactInterval := pc.getCompactInterval()
	retryInterval := pc.getRetryInterval()
	retentions := pc.getRetentions()
	go func() {
		lastSuccess := pc.clock.Now()
		baseInterval := pc.period
		for {
			pc.revs = append(pc.revs, pc.rg.Rev())
			if len(pc.revs) > retentions {
				pc.revs = pc.revs[1:]
			}
			select {
			case <-pc.ctx.Done():
				return
			case <-pc.clock.After(retryInterval):
				pc.mu.Lock()
				p := pc.paused
				pc.mu.Unlock()
				if p {
					continue
				}
			}
			if pc.clock.Now().Sub(lastSuccess) < baseInterval {
				continue
			}
			if baseInterval == pc.period {
				baseInterval = compactInterval
			}
			rev := pc.revs[0]
			if pc.lg != nil {
				pc.lg.Info("starting auto periodic compaction", zap.Int64("revision", rev), zap.Duration("compact-period", pc.period))
			} else {
				plog.Noticef("Starting auto-compaction at revision %d (retention: %v)", rev, pc.period)
			}
			_, err := pc.c.Compact(pc.ctx, &pb.CompactionRequest{Revision: rev})
			if err == nil || err == mvcc.ErrCompacted {
				if pc.lg != nil {
					pc.lg.Info("completed auto periodic compaction", zap.Int64("revision", rev), zap.Duration("compact-period", pc.period), zap.Duration("took", time.Since(lastSuccess)))
				} else {
					plog.Noticef("Finished auto-compaction at revision %d", rev)
				}
				lastSuccess = pc.clock.Now()
			} else {
				if pc.lg != nil {
					pc.lg.Warn("failed auto periodic compaction", zap.Int64("revision", rev), zap.Duration("compact-period", pc.period), zap.Duration("retry-interval", retryInterval), zap.Error(err))
				} else {
					plog.Noticef("Failed auto-compaction at revision %d (%v)", rev, err)
					plog.Noticef("Retry after %v", retryInterval)
				}
			}
		}
	}()
}
func (pc *Periodic) getCompactInterval() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	itv := pc.period
	if itv > time.Hour {
		itv = time.Hour
	}
	return itv
}
func (pc *Periodic) getRetentions() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int(pc.period/pc.getRetryInterval()) + 1
}

const retryDivisor = 10

func (pc *Periodic) getRetryInterval() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	itv := pc.period
	if itv > time.Hour {
		itv = time.Hour
	}
	return itv / retryDivisor
}
func (pc *Periodic) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc.cancel()
}
func (pc *Periodic) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc.mu.Lock()
	pc.paused = true
	pc.mu.Unlock()
}
func (pc *Periodic) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc.mu.Lock()
	pc.paused = false
	pc.mu.Unlock()
}
