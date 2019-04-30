package picker

import (
	"context"
	"sync"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/resolver"
)

func NewRoundrobinBalanced(lg *zap.Logger, scs []balancer.SubConn, addrToSc map[resolver.Address]balancer.SubConn, scToAddr map[balancer.SubConn]resolver.Address) Picker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &rrBalanced{lg: lg, scs: scs, addrToSc: addrToSc, scToAddr: scToAddr}
}

type rrBalanced struct {
	lg		*zap.Logger
	mu		sync.RWMutex
	next		int
	scs		[]balancer.SubConn
	addrToSc	map[resolver.Address]balancer.SubConn
	scToAddr	map[balancer.SubConn]resolver.Address
}

func (rb *rrBalanced) Pick(ctx context.Context, opts balancer.PickOptions) (balancer.SubConn, func(balancer.DoneInfo), error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rb.mu.RLock()
	n := len(rb.scs)
	rb.mu.RUnlock()
	if n == 0 {
		return nil, nil, balancer.ErrNoSubConnAvailable
	}
	rb.mu.Lock()
	cur := rb.next
	sc := rb.scs[cur]
	picked := rb.scToAddr[sc].Addr
	rb.next = (rb.next + 1) % len(rb.scs)
	rb.mu.Unlock()
	rb.lg.Debug("picked", zap.String("address", picked), zap.Int("subconn-index", cur), zap.Int("subconn-size", n))
	doneFunc := func(info balancer.DoneInfo) {
		fss := []zapcore.Field{zap.Error(info.Err), zap.String("address", picked), zap.Bool("success", info.Err == nil), zap.Bool("bytes-sent", info.BytesSent), zap.Bool("bytes-received", info.BytesReceived)}
		if info.Err == nil {
			rb.lg.Debug("balancer done", fss...)
		} else {
			rb.lg.Warn("balancer failed", fss...)
		}
	}
	return sc, doneFunc, nil
}
