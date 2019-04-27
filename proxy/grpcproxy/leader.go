package grpcproxy

import (
	"context"
	"math"
	"sync"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

const (
	lostLeaderKey	= "__lostleader"
	retryPerSecond	= 10
)

type leader struct {
	ctx		context.Context
	w		clientv3.Watcher
	mu		sync.RWMutex
	leaderc		chan struct{}
	disconnc	chan struct{}
	donec		chan struct{}
}

func newLeader(ctx context.Context, w clientv3.Watcher) *leader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := &leader{ctx: clientv3.WithRequireLeader(ctx), w: w, leaderc: make(chan struct{}), disconnc: make(chan struct{}), donec: make(chan struct{})}
	close(l.leaderc)
	go l.recvLoop()
	return l
}
func (l *leader) recvLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(l.donec)
	limiter := rate.NewLimiter(rate.Limit(retryPerSecond), retryPerSecond)
	rev := int64(math.MaxInt64 - 2)
	for limiter.Wait(l.ctx) == nil {
		wch := l.w.Watch(l.ctx, lostLeaderKey, clientv3.WithRev(rev), clientv3.WithCreatedNotify())
		cresp, ok := <-wch
		if !ok {
			l.loseLeader()
			continue
		}
		if cresp.Err() != nil {
			l.loseLeader()
			if rpctypes.ErrorDesc(cresp.Err()) == grpc.ErrClientConnClosing.Error() {
				close(l.disconnc)
				return
			}
			continue
		}
		l.gotLeader()
		<-wch
		l.loseLeader()
	}
}
func (l *leader) loseLeader() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.mu.RLock()
	defer l.mu.RUnlock()
	select {
	case <-l.leaderc:
	default:
		close(l.leaderc)
	}
}
func (l *leader) gotLeader() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.mu.Lock()
	defer l.mu.Unlock()
	select {
	case <-l.leaderc:
		l.leaderc = make(chan struct{})
	default:
	}
}
func (l *leader) disconnectNotify() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.disconnc
}
func (l *leader) stopNotify() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.donec
}
func (l *leader) lostNotify() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.leaderc
}
