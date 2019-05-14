package clientv3

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	minHealthRetryDuration = 3 * time.Second
	unknownService         = "unknown service grpc.health.v1.Health"
)

var ErrNoAddrAvilable = status.Error(codes.Unavailable, "there is no address available")

type healthCheckFunc func(ep string) (bool, error)
type notifyMsg int

const (
	notifyReset notifyMsg = iota
	notifyNext
)

type healthBalancer struct {
	addrs              []grpc.Address
	eps                []string
	notifyCh           chan []grpc.Address
	readyc             chan struct{}
	readyOnce          sync.Once
	healthCheck        healthCheckFunc
	healthCheckTimeout time.Duration
	unhealthyMu        sync.RWMutex
	unhealthyHostPorts map[string]time.Time
	mu                 sync.RWMutex
	upc                chan struct{}
	downc              chan struct{}
	stopc              chan struct{}
	stopOnce           sync.Once
	wg                 sync.WaitGroup
	donec              chan struct{}
	updateAddrsC       chan notifyMsg
	hostPort2ep        map[string]string
	pinAddr            string
	closed             bool
}

func newHealthBalancer(eps []string, timeout time.Duration, hc healthCheckFunc) *healthBalancer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	notifyCh := make(chan []grpc.Address)
	addrs := eps2addrs(eps)
	hb := &healthBalancer{addrs: addrs, eps: eps, notifyCh: notifyCh, readyc: make(chan struct{}), healthCheck: hc, unhealthyHostPorts: make(map[string]time.Time), upc: make(chan struct{}), stopc: make(chan struct{}), downc: make(chan struct{}), donec: make(chan struct{}), updateAddrsC: make(chan notifyMsg), hostPort2ep: getHostPort2ep(eps)}
	if timeout < minHealthRetryDuration {
		timeout = minHealthRetryDuration
	}
	hb.healthCheckTimeout = timeout
	close(hb.downc)
	go hb.updateNotifyLoop()
	hb.wg.Add(1)
	go func() {
		defer hb.wg.Done()
		hb.updateUnhealthy()
	}()
	return hb
}
func (b *healthBalancer) Start(target string, config grpc.BalancerConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (b *healthBalancer) ConnectNotify() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.upc
}
func (b *healthBalancer) ready() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.readyc
}
func (b *healthBalancer) endpoint(hostPort string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.hostPort2ep[hostPort]
}
func (b *healthBalancer) pinned() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.pinAddr
}
func (b *healthBalancer) hostPortError(hostPort string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b.endpoint(hostPort) == "" {
		logger.Lvl(4).Infof("clientv3/balancer: %q is stale (skip marking as unhealthy on %q)", hostPort, err.Error())
		return
	}
	b.unhealthyMu.Lock()
	b.unhealthyHostPorts[hostPort] = time.Now()
	b.unhealthyMu.Unlock()
	logger.Lvl(4).Infof("clientv3/balancer: %q is marked unhealthy (%q)", hostPort, err.Error())
}
func (b *healthBalancer) removeUnhealthy(hostPort, msg string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b.endpoint(hostPort) == "" {
		logger.Lvl(4).Infof("clientv3/balancer: %q was not in unhealthy (%q)", hostPort, msg)
		return
	}
	b.unhealthyMu.Lock()
	delete(b.unhealthyHostPorts, hostPort)
	b.unhealthyMu.Unlock()
	logger.Lvl(4).Infof("clientv3/balancer: %q is removed from unhealthy (%q)", hostPort, msg)
}
func (b *healthBalancer) countUnhealthy() (count int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.unhealthyMu.RLock()
	count = len(b.unhealthyHostPorts)
	b.unhealthyMu.RUnlock()
	return count
}
func (b *healthBalancer) isUnhealthy(hostPort string) (unhealthy bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.unhealthyMu.RLock()
	_, unhealthy = b.unhealthyHostPorts[hostPort]
	b.unhealthyMu.RUnlock()
	return unhealthy
}
func (b *healthBalancer) cleanupUnhealthy() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.unhealthyMu.Lock()
	for k, v := range b.unhealthyHostPorts {
		if time.Since(v) > b.healthCheckTimeout {
			delete(b.unhealthyHostPorts, k)
			logger.Lvl(4).Infof("clientv3/balancer: removed %q from unhealthy after %v", k, b.healthCheckTimeout)
		}
	}
	b.unhealthyMu.Unlock()
}
func (b *healthBalancer) liveAddrs() ([]grpc.Address, map[string]struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	unhealthyCnt := b.countUnhealthy()
	b.mu.RLock()
	defer b.mu.RUnlock()
	hbAddrs := b.addrs
	if len(b.addrs) == 1 || unhealthyCnt == 0 || unhealthyCnt == len(b.addrs) {
		liveHostPorts := make(map[string]struct{}, len(b.hostPort2ep))
		for k := range b.hostPort2ep {
			liveHostPorts[k] = struct{}{}
		}
		return hbAddrs, liveHostPorts
	}
	addrs := make([]grpc.Address, 0, len(b.addrs)-unhealthyCnt)
	liveHostPorts := make(map[string]struct{}, len(addrs))
	for _, addr := range b.addrs {
		if !b.isUnhealthy(addr.Addr) {
			addrs = append(addrs, addr)
			liveHostPorts[addr.Addr] = struct{}{}
		}
	}
	return addrs, liveHostPorts
}
func (b *healthBalancer) updateUnhealthy() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		select {
		case <-time.After(b.healthCheckTimeout):
			b.cleanupUnhealthy()
			pinned := b.pinned()
			if pinned == "" || b.isUnhealthy(pinned) {
				select {
				case b.updateAddrsC <- notifyNext:
				case <-b.stopc:
					return
				}
			}
		case <-b.stopc:
			return
		}
	}
}
func (b *healthBalancer) updateAddrs(eps ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	np := getHostPort2ep(eps)
	b.mu.Lock()
	defer b.mu.Unlock()
	match := len(np) == len(b.hostPort2ep)
	if match {
		for k, v := range np {
			if b.hostPort2ep[k] != v {
				match = false
				break
			}
		}
	}
	if match {
		return
	}
	b.hostPort2ep = np
	b.addrs, b.eps = eps2addrs(eps), eps
	b.unhealthyMu.Lock()
	b.unhealthyHostPorts = make(map[string]time.Time)
	b.unhealthyMu.Unlock()
}
func (b *healthBalancer) next() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.RLock()
	downc := b.downc
	b.mu.RUnlock()
	select {
	case b.updateAddrsC <- notifyNext:
	case <-b.stopc:
	}
	select {
	case <-downc:
	case <-b.stopc:
	}
}
func (b *healthBalancer) updateNotifyLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(b.donec)
	for {
		b.mu.RLock()
		upc, downc, addr := b.upc, b.downc, b.pinAddr
		b.mu.RUnlock()
		select {
		case <-downc:
			downc = nil
		default:
		}
		select {
		case <-upc:
			upc = nil
		default:
		}
		switch {
		case downc == nil && upc == nil:
			select {
			case <-b.stopc:
				return
			default:
			}
		case downc == nil:
			b.notifyAddrs(notifyReset)
			select {
			case <-upc:
			case msg := <-b.updateAddrsC:
				b.notifyAddrs(msg)
			case <-b.stopc:
				return
			}
		case upc == nil:
			select {
			case b.notifyCh <- []grpc.Address{{Addr: addr}}:
			case <-downc:
			case <-b.stopc:
				return
			}
			select {
			case <-downc:
				b.notifyAddrs(notifyReset)
			case msg := <-b.updateAddrsC:
				b.notifyAddrs(msg)
			case <-b.stopc:
				return
			}
		}
	}
}
func (b *healthBalancer) notifyAddrs(msg notifyMsg) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if msg == notifyNext {
		select {
		case b.notifyCh <- []grpc.Address{}:
		case <-b.stopc:
			return
		}
	}
	b.mu.RLock()
	pinAddr := b.pinAddr
	downc := b.downc
	b.mu.RUnlock()
	addrs, hostPorts := b.liveAddrs()
	var waitDown bool
	if pinAddr != "" {
		_, ok := hostPorts[pinAddr]
		waitDown = !ok
	}
	select {
	case b.notifyCh <- addrs:
		if waitDown {
			select {
			case <-downc:
			case <-b.stopc:
			}
		}
	case <-b.stopc:
	}
}
func (b *healthBalancer) Up(addr grpc.Address) func(error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !b.mayPin(addr) {
		return func(err error) {
		}
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return func(err error) {
		}
	}
	if !hasAddr(b.addrs, addr.Addr) {
		return func(err error) {
		}
	}
	if b.pinAddr != "" {
		logger.Lvl(4).Infof("clientv3/balancer: %q is up but not pinned (already pinned %q)", addr.Addr, b.pinAddr)
		return func(err error) {
		}
	}
	close(b.upc)
	b.downc = make(chan struct{})
	b.pinAddr = addr.Addr
	logger.Lvl(4).Infof("clientv3/balancer: pin %q", addr.Addr)
	b.readyOnce.Do(func() {
		close(b.readyc)
	})
	return func(err error) {
		b.hostPortError(addr.Addr, err)
		b.mu.Lock()
		b.upc = make(chan struct{})
		close(b.downc)
		b.pinAddr = ""
		b.mu.Unlock()
		logger.Lvl(4).Infof("clientv3/balancer: unpin %q (%q)", addr.Addr, err.Error())
	}
}
func (b *healthBalancer) mayPin(addr grpc.Address) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b.endpoint(addr.Addr) == "" {
		return false
	}
	b.unhealthyMu.RLock()
	unhealthyCnt := len(b.unhealthyHostPorts)
	failedTime, bad := b.unhealthyHostPorts[addr.Addr]
	b.unhealthyMu.RUnlock()
	b.mu.RLock()
	skip := len(b.addrs) == 1 || unhealthyCnt == 0 || len(b.addrs) == unhealthyCnt
	b.mu.RUnlock()
	if skip || !bad {
		return true
	}
	if elapsed := time.Since(failedTime); elapsed < b.healthCheckTimeout {
		logger.Lvl(4).Infof("clientv3/balancer: %q is up but not pinned (failed %v ago, require minimum %v after failure)", addr.Addr, elapsed, b.healthCheckTimeout)
		return false
	}
	if ok, _ := b.healthCheck(addr.Addr); ok {
		b.removeUnhealthy(addr.Addr, "health check success")
		return true
	}
	b.hostPortError(addr.Addr, errors.New("health check failed"))
	return false
}
func (b *healthBalancer) Get(ctx context.Context, opts grpc.BalancerGetOptions) (grpc.Address, func(), error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		addr   string
		closed bool
	)
	if !opts.BlockingWait {
		b.mu.RLock()
		closed = b.closed
		addr = b.pinAddr
		b.mu.RUnlock()
		if closed {
			return grpc.Address{Addr: ""}, nil, grpc.ErrClientConnClosing
		}
		if addr == "" {
			return grpc.Address{Addr: ""}, nil, ErrNoAddrAvilable
		}
		return grpc.Address{Addr: addr}, func() {
		}, nil
	}
	for {
		b.mu.RLock()
		ch := b.upc
		b.mu.RUnlock()
		select {
		case <-ch:
		case <-b.donec:
			return grpc.Address{Addr: ""}, nil, grpc.ErrClientConnClosing
		case <-ctx.Done():
			return grpc.Address{Addr: ""}, nil, ctx.Err()
		}
		b.mu.RLock()
		closed = b.closed
		addr = b.pinAddr
		b.mu.RUnlock()
		if closed {
			return grpc.Address{Addr: ""}, nil, grpc.ErrClientConnClosing
		}
		if addr != "" {
			break
		}
	}
	return grpc.Address{Addr: addr}, func() {
	}, nil
}
func (b *healthBalancer) Notify() <-chan []grpc.Address {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.notifyCh
}
func (b *healthBalancer) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.Lock()
	if b.closed {
		b.mu.Unlock()
		<-b.donec
		return nil
	}
	b.closed = true
	b.stopOnce.Do(func() {
		close(b.stopc)
	})
	b.pinAddr = ""
	select {
	case <-b.upc:
	default:
		close(b.upc)
	}
	b.mu.Unlock()
	b.wg.Wait()
	<-b.donec
	close(b.notifyCh)
	return nil
}
func grpcHealthCheck(client *Client, ep string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := client.dial(ep)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	cli := healthpb.NewHealthClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Check(ctx, &healthpb.HealthCheckRequest{})
	cancel()
	if err != nil {
		if s, ok := status.FromError(err); ok && s.Code() == codes.Unavailable {
			if s.Message() == unknownService {
				return true, nil
			}
		}
		return false, err
	}
	return resp.Status == healthpb.HealthCheckResponse_SERVING, nil
}
func hasAddr(addrs []grpc.Address, targetAddr string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, addr := range addrs {
		if targetAddr == addr.Addr {
			return true
		}
	}
	return false
}
func getHost(ep string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	url, uerr := url.Parse(ep)
	if uerr != nil || !strings.Contains(ep, "://") {
		return ep
	}
	return url.Host
}
func eps2addrs(eps []string) []grpc.Address {
	_logClusterCodePath()
	defer _logClusterCodePath()
	addrs := make([]grpc.Address, len(eps))
	for i := range eps {
		addrs[i].Addr = getHost(eps[i])
	}
	return addrs
}
func getHostPort2ep(eps []string) map[string]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hm := make(map[string]string, len(eps))
	for i := range eps {
		_, host, _ := parseEndpoint(eps[i])
		hm[host] = eps[i]
	}
	return hm
}
