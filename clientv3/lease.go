package clientv3

import (
	"context"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"sync"
	"time"
)

type (
	LeaseRevokeResponse pb.LeaseRevokeResponse
	LeaseID             int64
)
type LeaseGrantResponse struct {
	*pb.ResponseHeader
	ID    LeaseID
	TTL   int64
	Error string
}
type LeaseKeepAliveResponse struct {
	*pb.ResponseHeader
	ID  LeaseID
	TTL int64
}
type LeaseTimeToLiveResponse struct {
	*pb.ResponseHeader
	ID         LeaseID  `json:"id"`
	TTL        int64    `json:"ttl"`
	GrantedTTL int64    `json:"granted-ttl"`
	Keys       [][]byte `json:"keys"`
}
type LeaseStatus struct {
	ID LeaseID `json:"id"`
}
type LeaseLeasesResponse struct {
	*pb.ResponseHeader
	Leases []LeaseStatus `json:"leases"`
}

const (
	defaultTTL            = 5 * time.Second
	NoLease       LeaseID = 0
	retryConnWait         = 500 * time.Millisecond
)

var LeaseResponseChSize = 16

type ErrKeepAliveHalted struct{ Reason error }

func (e ErrKeepAliveHalted) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := "etcdclient: leases keep alive halted"
	if e.Reason != nil {
		s += ": " + e.Reason.Error()
	}
	return s
}

type Lease interface {
	Grant(ctx context.Context, ttl int64) (*LeaseGrantResponse, error)
	Revoke(ctx context.Context, id LeaseID) (*LeaseRevokeResponse, error)
	TimeToLive(ctx context.Context, id LeaseID, opts ...LeaseOption) (*LeaseTimeToLiveResponse, error)
	Leases(ctx context.Context) (*LeaseLeasesResponse, error)
	KeepAlive(ctx context.Context, id LeaseID) (<-chan *LeaseKeepAliveResponse, error)
	KeepAliveOnce(ctx context.Context, id LeaseID) (*LeaseKeepAliveResponse, error)
	Close() error
}
type lessor struct {
	mu                    sync.Mutex
	donec                 chan struct{}
	loopErr               error
	remote                pb.LeaseClient
	stream                pb.Lease_LeaseKeepAliveClient
	streamCancel          context.CancelFunc
	stopCtx               context.Context
	stopCancel            context.CancelFunc
	keepAlives            map[LeaseID]*keepAlive
	firstKeepAliveTimeout time.Duration
	firstKeepAliveOnce    sync.Once
	callOpts              []grpc.CallOption
}
type keepAlive struct {
	chs           []chan<- *LeaseKeepAliveResponse
	ctxs          []context.Context
	deadline      time.Time
	nextKeepAlive time.Time
	donec         chan struct{}
}

func NewLease(c *Client) Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewLeaseFromLeaseClient(RetryLeaseClient(c), c, c.cfg.DialTimeout+time.Second)
}
func NewLeaseFromLeaseClient(remote pb.LeaseClient, c *Client, keepAliveTimeout time.Duration) Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := &lessor{donec: make(chan struct{}), keepAlives: make(map[LeaseID]*keepAlive), remote: remote, firstKeepAliveTimeout: keepAliveTimeout}
	if l.firstKeepAliveTimeout == time.Second {
		l.firstKeepAliveTimeout = defaultTTL
	}
	if c != nil {
		l.callOpts = c.callOpts
	}
	reqLeaderCtx := WithRequireLeader(context.Background())
	l.stopCtx, l.stopCancel = context.WithCancel(reqLeaderCtx)
	return l
}
func (l *lessor) Grant(ctx context.Context, ttl int64) (*LeaseGrantResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := &pb.LeaseGrantRequest{TTL: ttl}
	resp, err := l.remote.LeaseGrant(ctx, r, l.callOpts...)
	if err == nil {
		gresp := &LeaseGrantResponse{ResponseHeader: resp.GetHeader(), ID: LeaseID(resp.ID), TTL: resp.TTL, Error: resp.Error}
		return gresp, nil
	}
	return nil, toErr(ctx, err)
}
func (l *lessor) Revoke(ctx context.Context, id LeaseID) (*LeaseRevokeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := &pb.LeaseRevokeRequest{ID: int64(id)}
	resp, err := l.remote.LeaseRevoke(ctx, r, l.callOpts...)
	if err == nil {
		return (*LeaseRevokeResponse)(resp), nil
	}
	return nil, toErr(ctx, err)
}
func (l *lessor) TimeToLive(ctx context.Context, id LeaseID, opts ...LeaseOption) (*LeaseTimeToLiveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := toLeaseTimeToLiveRequest(id, opts...)
	resp, err := l.remote.LeaseTimeToLive(ctx, r, l.callOpts...)
	if err == nil {
		gresp := &LeaseTimeToLiveResponse{ResponseHeader: resp.GetHeader(), ID: LeaseID(resp.ID), TTL: resp.TTL, GrantedTTL: resp.GrantedTTL, Keys: resp.Keys}
		return gresp, nil
	}
	return nil, toErr(ctx, err)
}
func (l *lessor) Leases(ctx context.Context) (*LeaseLeasesResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := l.remote.LeaseLeases(ctx, &pb.LeaseLeasesRequest{}, l.callOpts...)
	if err == nil {
		leases := make([]LeaseStatus, len(resp.Leases))
		for i := range resp.Leases {
			leases[i] = LeaseStatus{ID: LeaseID(resp.Leases[i].ID)}
		}
		return &LeaseLeasesResponse{ResponseHeader: resp.GetHeader(), Leases: leases}, nil
	}
	return nil, toErr(ctx, err)
}
func (l *lessor) KeepAlive(ctx context.Context, id LeaseID) (<-chan *LeaseKeepAliveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ch := make(chan *LeaseKeepAliveResponse, LeaseResponseChSize)
	l.mu.Lock()
	select {
	case <-l.donec:
		err := l.loopErr
		l.mu.Unlock()
		close(ch)
		return ch, ErrKeepAliveHalted{Reason: err}
	default:
	}
	ka, ok := l.keepAlives[id]
	if !ok {
		ka = &keepAlive{chs: []chan<- *LeaseKeepAliveResponse{ch}, ctxs: []context.Context{ctx}, deadline: time.Now().Add(l.firstKeepAliveTimeout), nextKeepAlive: time.Now(), donec: make(chan struct{})}
		l.keepAlives[id] = ka
	} else {
		ka.ctxs = append(ka.ctxs, ctx)
		ka.chs = append(ka.chs, ch)
	}
	l.mu.Unlock()
	go l.keepAliveCtxCloser(id, ctx, ka.donec)
	l.firstKeepAliveOnce.Do(func() {
		go l.recvKeepAliveLoop()
		go l.deadlineLoop()
	})
	return ch, nil
}
func (l *lessor) KeepAliveOnce(ctx context.Context, id LeaseID) (*LeaseKeepAliveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		resp, err := l.keepAliveOnce(ctx, id)
		if err == nil {
			if resp.TTL <= 0 {
				err = rpctypes.ErrLeaseNotFound
			}
			return resp, err
		}
		if isHaltErr(ctx, err) {
			return nil, toErr(ctx, err)
		}
	}
}
func (l *lessor) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.stopCancel()
	l.firstKeepAliveOnce.Do(func() {
		close(l.donec)
	})
	<-l.donec
	return nil
}
func (l *lessor) keepAliveCtxCloser(id LeaseID, ctx context.Context, donec <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case <-donec:
		return
	case <-l.donec:
		return
	case <-ctx.Done():
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	ka, ok := l.keepAlives[id]
	if !ok {
		return
	}
	for i, c := range ka.ctxs {
		if c == ctx {
			close(ka.chs[i])
			ka.ctxs = append(ka.ctxs[:i], ka.ctxs[i+1:]...)
			ka.chs = append(ka.chs[:i], ka.chs[i+1:]...)
			break
		}
	}
	if len(ka.chs) == 0 {
		delete(l.keepAlives, id)
	}
}
func (l *lessor) closeRequireLeader() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, ka := range l.keepAlives {
		reqIdxs := 0
		for i, ctx := range ka.ctxs {
			md, ok := metadata.FromOutgoingContext(ctx)
			if !ok {
				continue
			}
			ks := md[rpctypes.MetadataRequireLeaderKey]
			if len(ks) < 1 || ks[0] != rpctypes.MetadataHasLeader {
				continue
			}
			close(ka.chs[i])
			ka.chs[i] = nil
			reqIdxs++
		}
		if reqIdxs == 0 {
			continue
		}
		newChs := make([]chan<- *LeaseKeepAliveResponse, len(ka.chs)-reqIdxs)
		newCtxs := make([]context.Context, len(newChs))
		newIdx := 0
		for i := range ka.chs {
			if ka.chs[i] == nil {
				continue
			}
			newChs[newIdx], newCtxs[newIdx] = ka.chs[i], ka.ctxs[newIdx]
			newIdx++
		}
		ka.chs, ka.ctxs = newChs, newCtxs
	}
}
func (l *lessor) keepAliveOnce(ctx context.Context, id LeaseID) (*LeaseKeepAliveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	stream, err := l.remote.LeaseKeepAlive(cctx, l.callOpts...)
	if err != nil {
		return nil, toErr(ctx, err)
	}
	err = stream.Send(&pb.LeaseKeepAliveRequest{ID: int64(id)})
	if err != nil {
		return nil, toErr(ctx, err)
	}
	resp, rerr := stream.Recv()
	if rerr != nil {
		return nil, toErr(ctx, rerr)
	}
	karesp := &LeaseKeepAliveResponse{ResponseHeader: resp.GetHeader(), ID: LeaseID(resp.ID), TTL: resp.TTL}
	return karesp, nil
}
func (l *lessor) recvKeepAliveLoop() (gerr error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer func() {
		l.mu.Lock()
		close(l.donec)
		l.loopErr = gerr
		for _, ka := range l.keepAlives {
			ka.close()
		}
		l.keepAlives = make(map[LeaseID]*keepAlive)
		l.mu.Unlock()
	}()
	for {
		stream, err := l.resetRecv()
		if err != nil {
			if canceledByCaller(l.stopCtx, err) {
				return err
			}
		} else {
			for {
				resp, err := stream.Recv()
				if err != nil {
					if canceledByCaller(l.stopCtx, err) {
						return err
					}
					if toErr(l.stopCtx, err) == rpctypes.ErrNoLeader {
						l.closeRequireLeader()
					}
					break
				}
				l.recvKeepAlive(resp)
			}
		}
		select {
		case <-time.After(retryConnWait):
			continue
		case <-l.stopCtx.Done():
			return l.stopCtx.Err()
		}
	}
}
func (l *lessor) resetRecv() (pb.Lease_LeaseKeepAliveClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sctx, cancel := context.WithCancel(l.stopCtx)
	stream, err := l.remote.LeaseKeepAlive(sctx, l.callOpts...)
	if err != nil {
		cancel()
		return nil, err
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.stream != nil && l.streamCancel != nil {
		l.streamCancel()
	}
	l.streamCancel = cancel
	l.stream = stream
	go l.sendKeepAliveLoop(stream)
	return stream, nil
}
func (l *lessor) recvKeepAlive(resp *pb.LeaseKeepAliveResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	karesp := &LeaseKeepAliveResponse{ResponseHeader: resp.GetHeader(), ID: LeaseID(resp.ID), TTL: resp.TTL}
	l.mu.Lock()
	defer l.mu.Unlock()
	ka, ok := l.keepAlives[karesp.ID]
	if !ok {
		return
	}
	if karesp.TTL <= 0 {
		delete(l.keepAlives, karesp.ID)
		ka.close()
		return
	}
	nextKeepAlive := time.Now().Add((time.Duration(karesp.TTL) * time.Second) / 3.0)
	ka.deadline = time.Now().Add(time.Duration(karesp.TTL) * time.Second)
	for _, ch := range ka.chs {
		select {
		case ch <- karesp:
		default:
		}
		ka.nextKeepAlive = nextKeepAlive
	}
}
func (l *lessor) deadlineLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		select {
		case <-time.After(time.Second):
		case <-l.donec:
			return
		}
		now := time.Now()
		l.mu.Lock()
		for id, ka := range l.keepAlives {
			if ka.deadline.Before(now) {
				ka.close()
				delete(l.keepAlives, id)
			}
		}
		l.mu.Unlock()
	}
}
func (l *lessor) sendKeepAliveLoop(stream pb.Lease_LeaseKeepAliveClient) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		var tosend []LeaseID
		now := time.Now()
		l.mu.Lock()
		for id, ka := range l.keepAlives {
			if ka.nextKeepAlive.Before(now) {
				tosend = append(tosend, id)
			}
		}
		l.mu.Unlock()
		for _, id := range tosend {
			r := &pb.LeaseKeepAliveRequest{ID: int64(id)}
			if err := stream.Send(r); err != nil {
				return
			}
		}
		select {
		case <-time.After(500 * time.Millisecond):
		case <-stream.Context().Done():
			return
		case <-l.donec:
			return
		case <-l.stopCtx.Done():
			return
		}
	}
}
func (ka *keepAlive) close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(ka.donec)
	for _, ch := range ka.chs {
		close(ch)
	}
}
