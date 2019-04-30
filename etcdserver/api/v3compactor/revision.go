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

type Revision struct {
	lg		*zap.Logger
	clock		clockwork.Clock
	retention	int64
	rg		RevGetter
	c		Compactable
	ctx		context.Context
	cancel		context.CancelFunc
	mu		sync.Mutex
	paused		bool
}

func newRevision(lg *zap.Logger, clock clockwork.Clock, retention int64, rg RevGetter, c Compactable) *Revision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rc := &Revision{lg: lg, clock: clock, retention: retention, rg: rg, c: c}
	rc.ctx, rc.cancel = context.WithCancel(context.Background())
	return rc
}

const revInterval = 5 * time.Minute

func (rc *Revision) Run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prev := int64(0)
	go func() {
		for {
			select {
			case <-rc.ctx.Done():
				return
			case <-rc.clock.After(revInterval):
				rc.mu.Lock()
				p := rc.paused
				rc.mu.Unlock()
				if p {
					continue
				}
			}
			rev := rc.rg.Rev() - rc.retention
			if rev <= 0 || rev == prev {
				continue
			}
			now := time.Now()
			if rc.lg != nil {
				rc.lg.Info("starting auto revision compaction", zap.Int64("revision", rev), zap.Int64("revision-compaction-retention", rc.retention))
			} else {
				plog.Noticef("Starting auto-compaction at revision %d (retention: %d revisions)", rev, rc.retention)
			}
			_, err := rc.c.Compact(rc.ctx, &pb.CompactionRequest{Revision: rev})
			if err == nil || err == mvcc.ErrCompacted {
				prev = rev
				if rc.lg != nil {
					rc.lg.Info("completed auto revision compaction", zap.Int64("revision", rev), zap.Int64("revision-compaction-retention", rc.retention), zap.Duration("took", time.Since(now)))
				} else {
					plog.Noticef("Finished auto-compaction at revision %d", rev)
				}
			} else {
				if rc.lg != nil {
					rc.lg.Warn("failed auto revision compaction", zap.Int64("revision", rev), zap.Int64("revision-compaction-retention", rc.retention), zap.Duration("retry-interval", revInterval), zap.Error(err))
				} else {
					plog.Noticef("Failed auto-compaction at revision %d (%v)", rev, err)
					plog.Noticef("Retry after %v", revInterval)
				}
			}
		}
	}()
}
func (rc *Revision) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rc.cancel()
}
func (rc *Revision) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rc.mu.Lock()
	rc.paused = true
	rc.mu.Unlock()
}
func (rc *Revision) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rc.mu.Lock()
	rc.paused = false
	rc.mu.Unlock()
}
