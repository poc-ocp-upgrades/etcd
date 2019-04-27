package compactor

import (
	"context"
	"sync"
	"time"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/mvcc"
	"github.com/jonboulle/clockwork"
)

type Revision struct {
	clock		clockwork.Clock
	retention	int64
	rg		RevGetter
	c		Compactable
	ctx		context.Context
	cancel		context.CancelFunc
	mu		sync.Mutex
	paused		bool
}

func NewRevision(retention int64, rg RevGetter, c Compactable) *Revision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newRevision(clockwork.NewRealClock(), retention, rg, c)
}
func newRevision(clock clockwork.Clock, retention int64, rg RevGetter, c Compactable) *Revision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := &Revision{clock: clock, retention: retention, rg: rg, c: c}
	t.ctx, t.cancel = context.WithCancel(context.Background())
	return t
}

const revInterval = 5 * time.Minute

func (t *Revision) Run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prev := int64(0)
	go func() {
		for {
			select {
			case <-t.ctx.Done():
				return
			case <-t.clock.After(revInterval):
				t.mu.Lock()
				p := t.paused
				t.mu.Unlock()
				if p {
					continue
				}
			}
			rev := t.rg.Rev() - t.retention
			if rev <= 0 || rev == prev {
				continue
			}
			plog.Noticef("Starting auto-compaction at revision %d (retention: %d revisions)", rev, t.retention)
			_, err := t.c.Compact(t.ctx, &pb.CompactionRequest{Revision: rev})
			if err == nil || err == mvcc.ErrCompacted {
				prev = rev
				plog.Noticef("Finished auto-compaction at revision %d", rev)
			} else {
				plog.Noticef("Failed auto-compaction at revision %d (%v)", rev, err)
				plog.Noticef("Retry after %v", revInterval)
			}
		}
	}()
}
func (t *Revision) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.cancel()
}
func (t *Revision) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.paused = true
}
func (t *Revision) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.mu.Lock()
	defer t.mu.Unlock()
	t.paused = false
}
