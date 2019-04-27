package grpcproxy

import (
	"context"
	"sync"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type watchProxy struct {
	cw	clientv3.Watcher
	ctx	context.Context
	leader	*leader
	ranges	*watchRanges
	mu	sync.Mutex
	wg	sync.WaitGroup
	kv	clientv3.KV
}

func NewWatchProxy(c *clientv3.Client) (pb.WatchServer, <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cctx, cancel := context.WithCancel(c.Ctx())
	wp := &watchProxy{cw: c.Watcher, ctx: cctx, leader: newLeader(c.Ctx(), c.Watcher), kv: c.KV}
	wp.ranges = newWatchRanges(wp)
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		<-wp.leader.stopNotify()
		wp.mu.Lock()
		select {
		case <-wp.ctx.Done():
		case <-wp.leader.disconnectNotify():
			cancel()
		}
		<-wp.ctx.Done()
		wp.mu.Unlock()
		wp.wg.Wait()
		wp.ranges.stop()
	}()
	return wp, ch
}
func (wp *watchProxy) Watch(stream pb.Watch_WatchServer) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wp.mu.Lock()
	select {
	case <-wp.ctx.Done():
		wp.mu.Unlock()
		select {
		case <-wp.leader.disconnectNotify():
			return grpc.ErrClientConnClosing
		default:
			return wp.ctx.Err()
		}
	default:
		wp.wg.Add(1)
	}
	wp.mu.Unlock()
	ctx, cancel := context.WithCancel(stream.Context())
	wps := &watchProxyStream{ranges: wp.ranges, watchers: make(map[int64]*watcher), stream: stream, watchCh: make(chan *pb.WatchResponse, 1024), ctx: ctx, cancel: cancel, kv: wp.kv}
	var lostLeaderC <-chan struct{}
	if md, ok := metadata.FromOutgoingContext(stream.Context()); ok {
		v := md[rpctypes.MetadataRequireLeaderKey]
		if len(v) > 0 && v[0] == rpctypes.MetadataHasLeader {
			lostLeaderC = wp.leader.lostNotify()
			select {
			case <-lostLeaderC:
				wp.wg.Done()
				return rpctypes.ErrNoLeader
			default:
			}
		}
	}
	stopc := make(chan struct{}, 3)
	go func() {
		defer func() {
			stopc <- struct{}{}
		}()
		wps.recvLoop()
	}()
	go func() {
		defer func() {
			stopc <- struct{}{}
		}()
		wps.sendLoop()
	}()
	go func() {
		defer func() {
			stopc <- struct{}{}
		}()
		select {
		case <-lostLeaderC:
		case <-ctx.Done():
		case <-wp.ctx.Done():
		}
	}()
	<-stopc
	cancel()
	go func() {
		<-stopc
		<-stopc
		wps.close()
		wp.wg.Done()
	}()
	select {
	case <-lostLeaderC:
		return rpctypes.ErrNoLeader
	case <-wp.leader.disconnectNotify():
		return grpc.ErrClientConnClosing
	default:
		return wps.ctx.Err()
	}
}

type watchProxyStream struct {
	ranges		*watchRanges
	mu		sync.Mutex
	watchers	map[int64]*watcher
	nextWatcherID	int64
	stream		pb.Watch_WatchServer
	watchCh		chan *pb.WatchResponse
	ctx		context.Context
	cancel		context.CancelFunc
	kv		clientv3.KV
}

func (wps *watchProxyStream) close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wg sync.WaitGroup
	wps.cancel()
	wps.mu.Lock()
	wg.Add(len(wps.watchers))
	for _, wpsw := range wps.watchers {
		go func(w *watcher) {
			wps.ranges.delete(w)
			wg.Done()
		}(wpsw)
	}
	wps.watchers = nil
	wps.mu.Unlock()
	wg.Wait()
	close(wps.watchCh)
}
func (wps *watchProxyStream) checkPermissionForWatch(key, rangeEnd []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(key) == 0 {
		key = []byte{0}
		rangeEnd = []byte{0}
	}
	req := &pb.RangeRequest{Serializable: true, Key: key, RangeEnd: rangeEnd, CountOnly: true, Limit: 1}
	_, err := wps.kv.Do(wps.ctx, RangeRequestToOp(req))
	return err
}
func (wps *watchProxyStream) recvLoop() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		req, err := wps.stream.Recv()
		if err != nil {
			return err
		}
		switch uv := req.RequestUnion.(type) {
		case *pb.WatchRequest_CreateRequest:
			cr := uv.CreateRequest
			if err = wps.checkPermissionForWatch(cr.Key, cr.RangeEnd); err != nil && err == rpctypes.ErrPermissionDenied {
				wps.watchCh <- &pb.WatchResponse{Header: &pb.ResponseHeader{}, WatchId: -1, Created: true, Canceled: true}
				continue
			}
			w := &watcher{wr: watchRange{string(cr.Key), string(cr.RangeEnd)}, id: wps.nextWatcherID, wps: wps, nextrev: cr.StartRevision, progress: cr.ProgressNotify, prevKV: cr.PrevKv, filters: v3rpc.FiltersFromRequest(cr)}
			if !w.wr.valid() {
				w.post(&pb.WatchResponse{WatchId: -1, Created: true, Canceled: true})
				continue
			}
			wps.nextWatcherID++
			w.nextrev = cr.StartRevision
			wps.watchers[w.id] = w
			wps.ranges.add(w)
		case *pb.WatchRequest_CancelRequest:
			wps.delete(uv.CancelRequest.WatchId)
		default:
			panic("not implemented")
		}
	}
}
func (wps *watchProxyStream) sendLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		select {
		case wresp, ok := <-wps.watchCh:
			if !ok {
				return
			}
			if err := wps.stream.Send(wresp); err != nil {
				return
			}
		case <-wps.ctx.Done():
			return
		}
	}
}
func (wps *watchProxyStream) delete(id int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wps.mu.Lock()
	defer wps.mu.Unlock()
	w, ok := wps.watchers[id]
	if !ok {
		return
	}
	wps.ranges.delete(w)
	delete(wps.watchers, id)
	resp := &pb.WatchResponse{Header: &w.lastHeader, WatchId: id, Canceled: true}
	wps.watchCh <- resp
}
