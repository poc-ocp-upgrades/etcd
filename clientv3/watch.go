package clientv3

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	v3rpc "go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	mvccpb "go.etcd.io/etcd/mvcc/mvccpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	EventTypeDelete		= mvccpb.DELETE
	EventTypePut		= mvccpb.PUT
	closeSendErrTimeout	= 250 * time.Millisecond
)

type Event mvccpb.Event
type WatchChan <-chan WatchResponse
type Watcher interface {
	Watch(ctx context.Context, key string, opts ...OpOption) WatchChan
	RequestProgress(ctx context.Context) error
	Close() error
}
type WatchResponse struct {
	Header		pb.ResponseHeader
	Events		[]*Event
	CompactRevision	int64
	Canceled	bool
	Created		bool
	closeErr	error
	cancelReason	string
}

func (e *Event) IsCreate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.Type == EventTypePut && e.Kv.CreateRevision == e.Kv.ModRevision
}
func (e *Event) IsModify() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.Type == EventTypePut && e.Kv.CreateRevision != e.Kv.ModRevision
}
func (wr *WatchResponse) Err() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case wr.closeErr != nil:
		return v3rpc.Error(wr.closeErr)
	case wr.CompactRevision != 0:
		return v3rpc.ErrCompacted
	case wr.Canceled:
		if len(wr.cancelReason) != 0 {
			return v3rpc.Error(status.Error(codes.FailedPrecondition, wr.cancelReason))
		}
		return v3rpc.ErrFutureRev
	}
	return nil
}
func (wr *WatchResponse) IsProgressNotify() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(wr.Events) == 0 && !wr.Canceled && !wr.Created && wr.CompactRevision == 0 && wr.Header.Revision != 0
}

type watcher struct {
	remote		pb.WatchClient
	callOpts	[]grpc.CallOption
	mu		sync.RWMutex
	streams		map[string]*watchGrpcStream
}
type watchGrpcStream struct {
	owner		*watcher
	remote		pb.WatchClient
	callOpts	[]grpc.CallOption
	ctx		context.Context
	ctxKey		string
	cancel		context.CancelFunc
	substreams	map[int64]*watcherStream
	resuming	[]*watcherStream
	reqc		chan watchStreamRequest
	respc		chan *pb.WatchResponse
	donec		chan struct{}
	errc		chan error
	closingc	chan *watcherStream
	wg		sync.WaitGroup
	resumec		chan struct{}
	closeErr	error
}
type watchStreamRequest interface{ toPB() *pb.WatchRequest }
type watchRequest struct {
	ctx		context.Context
	key		string
	end		string
	rev		int64
	createdNotify	bool
	progressNotify	bool
	fragment	bool
	filters		[]pb.WatchCreateRequest_FilterType
	prevKV		bool
	retc		chan chan WatchResponse
}
type progressRequest struct{}
type watcherStream struct {
	initReq	watchRequest
	outc	chan WatchResponse
	recvc	chan *WatchResponse
	donec	chan struct{}
	closing	bool
	id	int64
	buf	[]*WatchResponse
}

func NewWatcher(c *Client) Watcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewWatchFromWatchClient(pb.NewWatchClient(c.conn), c)
}
func NewWatchFromWatchClient(wc pb.WatchClient, c *Client) Watcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w := &watcher{remote: wc, streams: make(map[string]*watchGrpcStream)}
	if c != nil {
		w.callOpts = c.callOpts
	}
	return w
}

var valCtxCh = make(chan struct{})
var zeroTime = time.Unix(0, 0)

type valCtx struct{ context.Context }

func (vc *valCtx) Deadline() (time.Time, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return zeroTime, false
}
func (vc *valCtx) Done() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return valCtxCh
}
func (vc *valCtx) Err() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (w *watcher) newWatcherGrpcStream(inctx context.Context) *watchGrpcStream {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithCancel(&valCtx{inctx})
	wgs := &watchGrpcStream{owner: w, remote: w.remote, callOpts: w.callOpts, ctx: ctx, ctxKey: streamKeyFromCtx(inctx), cancel: cancel, substreams: make(map[int64]*watcherStream), respc: make(chan *pb.WatchResponse), reqc: make(chan watchStreamRequest), donec: make(chan struct{}), errc: make(chan error, 1), closingc: make(chan *watcherStream), resumec: make(chan struct{})}
	go wgs.run()
	return wgs
}
func (w *watcher) Watch(ctx context.Context, key string, opts ...OpOption) WatchChan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ow := opWatch(key, opts...)
	var filters []pb.WatchCreateRequest_FilterType
	if ow.filterPut {
		filters = append(filters, pb.WatchCreateRequest_NOPUT)
	}
	if ow.filterDelete {
		filters = append(filters, pb.WatchCreateRequest_NODELETE)
	}
	wr := &watchRequest{ctx: ctx, createdNotify: ow.createdNotify, key: string(ow.key), end: string(ow.end), rev: ow.rev, progressNotify: ow.progressNotify, fragment: ow.fragment, filters: filters, prevKV: ow.prevKV, retc: make(chan chan WatchResponse, 1)}
	ok := false
	ctxKey := streamKeyFromCtx(ctx)
	w.mu.Lock()
	if w.streams == nil {
		w.mu.Unlock()
		ch := make(chan WatchResponse)
		close(ch)
		return ch
	}
	wgs := w.streams[ctxKey]
	if wgs == nil {
		wgs = w.newWatcherGrpcStream(ctx)
		w.streams[ctxKey] = wgs
	}
	donec := wgs.donec
	reqc := wgs.reqc
	w.mu.Unlock()
	closeCh := make(chan WatchResponse, 1)
	select {
	case reqc <- wr:
		ok = true
	case <-wr.ctx.Done():
	case <-donec:
		if wgs.closeErr != nil {
			closeCh <- WatchResponse{Canceled: true, closeErr: wgs.closeErr}
			break
		}
		return w.Watch(ctx, key, opts...)
	}
	if ok {
		select {
		case ret := <-wr.retc:
			return ret
		case <-ctx.Done():
		case <-donec:
			if wgs.closeErr != nil {
				closeCh <- WatchResponse{Canceled: true, closeErr: wgs.closeErr}
				break
			}
			return w.Watch(ctx, key, opts...)
		}
	}
	close(closeCh)
	return closeCh
}
func (w *watcher) Close() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.mu.Lock()
	streams := w.streams
	w.streams = nil
	w.mu.Unlock()
	for _, wgs := range streams {
		if werr := wgs.close(); werr != nil {
			err = werr
		}
	}
	if err == context.Canceled {
		err = nil
	}
	return err
}
func (w *watcher) RequestProgress(ctx context.Context) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctxKey := streamKeyFromCtx(ctx)
	w.mu.Lock()
	if w.streams == nil {
		return fmt.Errorf("no stream found for context")
	}
	wgs := w.streams[ctxKey]
	if wgs == nil {
		wgs = w.newWatcherGrpcStream(ctx)
		w.streams[ctxKey] = wgs
	}
	donec := wgs.donec
	reqc := wgs.reqc
	w.mu.Unlock()
	pr := &progressRequest{}
	select {
	case reqc <- pr:
		return nil
	case <-ctx.Done():
		if err == nil {
			return ctx.Err()
		}
		return err
	case <-donec:
		if wgs.closeErr != nil {
			return wgs.closeErr
		}
		return w.RequestProgress(ctx)
	}
}
func (w *watchGrpcStream) close() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.cancel()
	<-w.donec
	select {
	case err = <-w.errc:
	default:
	}
	return toErr(w.ctx, err)
}
func (w *watcher) closeStream(wgs *watchGrpcStream) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.mu.Lock()
	close(wgs.donec)
	wgs.cancel()
	if w.streams != nil {
		delete(w.streams, wgs.ctxKey)
	}
	w.mu.Unlock()
}
func (w *watchGrpcStream) addSubstream(resp *pb.WatchResponse, ws *watcherStream) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if resp.WatchId == -1 || (resp.Canceled && resp.CancelReason != "") {
		w.closeErr = v3rpc.Error(errors.New(resp.CancelReason))
		close(ws.recvc)
		return
	}
	ws.id = resp.WatchId
	w.substreams[ws.id] = ws
}
func (w *watchGrpcStream) sendCloseSubstream(ws *watcherStream, resp *WatchResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case ws.outc <- *resp:
	case <-ws.initReq.ctx.Done():
	case <-time.After(closeSendErrTimeout):
	}
	close(ws.outc)
}
func (w *watchGrpcStream) closeSubstream(ws *watcherStream) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case ws.initReq.retc <- ws.outc:
	default:
	}
	if closeErr := w.closeErr; closeErr != nil && ws.initReq.ctx.Err() == nil {
		go w.sendCloseSubstream(ws, &WatchResponse{Canceled: true, closeErr: w.closeErr})
	} else if ws.outc != nil {
		close(ws.outc)
	}
	if ws.id != -1 {
		delete(w.substreams, ws.id)
		return
	}
	for i := range w.resuming {
		if w.resuming[i] == ws {
			w.resuming[i] = nil
			return
		}
	}
}
func (w *watchGrpcStream) run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wc pb.Watch_WatchClient
	var closeErr error
	closing := make(map[*watcherStream]struct{})
	defer func() {
		w.closeErr = closeErr
		for _, ws := range w.substreams {
			if _, ok := closing[ws]; !ok {
				close(ws.recvc)
				closing[ws] = struct{}{}
			}
		}
		for _, ws := range w.resuming {
			if _, ok := closing[ws]; ws != nil && !ok {
				close(ws.recvc)
				closing[ws] = struct{}{}
			}
		}
		w.joinSubstreams()
		for range closing {
			w.closeSubstream(<-w.closingc)
		}
		w.wg.Wait()
		w.owner.closeStream(w)
	}()
	if wc, closeErr = w.newWatchClient(); closeErr != nil {
		return
	}
	cancelSet := make(map[int64]struct{})
	var cur *pb.WatchResponse
	for {
		select {
		case req := <-w.reqc:
			switch wreq := req.(type) {
			case *watchRequest:
				outc := make(chan WatchResponse, 1)
				ws := &watcherStream{initReq: *wreq, id: -1, outc: outc, recvc: make(chan *WatchResponse)}
				ws.donec = make(chan struct{})
				w.wg.Add(1)
				go w.serveSubstream(ws, w.resumec)
				w.resuming = append(w.resuming, ws)
				if len(w.resuming) == 1 {
					wc.Send(ws.initReq.toPB())
				}
			case *progressRequest:
				wc.Send(wreq.toPB())
			}
		case pbresp := <-w.respc:
			if cur == nil || pbresp.Created || pbresp.Canceled {
				cur = pbresp
			} else if cur != nil && cur.WatchId == pbresp.WatchId {
				cur.Events = append(cur.Events, pbresp.Events...)
				cur.Fragment = pbresp.Fragment
			}
			switch {
			case pbresp.Created:
				if ws := w.resuming[0]; ws != nil {
					w.addSubstream(pbresp, ws)
					w.dispatchEvent(pbresp)
					w.resuming[0] = nil
				}
				if ws := w.nextResume(); ws != nil {
					wc.Send(ws.initReq.toPB())
				}
				cur = nil
			case pbresp.Canceled && pbresp.CompactRevision == 0:
				delete(cancelSet, pbresp.WatchId)
				if ws, ok := w.substreams[pbresp.WatchId]; ok {
					close(ws.recvc)
					closing[ws] = struct{}{}
				}
				cur = nil
			case cur.Fragment:
				continue
			default:
				ok := w.dispatchEvent(cur)
				cur = nil
				if ok {
					break
				}
				if _, ok := cancelSet[pbresp.WatchId]; ok {
					break
				}
				cancelSet[pbresp.WatchId] = struct{}{}
				cr := &pb.WatchRequest_CancelRequest{CancelRequest: &pb.WatchCancelRequest{WatchId: pbresp.WatchId}}
				req := &pb.WatchRequest{RequestUnion: cr}
				wc.Send(req)
			}
		case err := <-w.errc:
			if isHaltErr(w.ctx, err) || toErr(w.ctx, err) == v3rpc.ErrNoLeader {
				closeErr = err
				return
			}
			if wc, closeErr = w.newWatchClient(); closeErr != nil {
				return
			}
			if ws := w.nextResume(); ws != nil {
				wc.Send(ws.initReq.toPB())
			}
			cancelSet = make(map[int64]struct{})
		case <-w.ctx.Done():
			return
		case ws := <-w.closingc:
			w.closeSubstream(ws)
			delete(closing, ws)
			if len(w.substreams)+len(w.resuming) == 0 {
				return
			}
		}
	}
}
func (w *watchGrpcStream) nextResume() *watcherStream {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for len(w.resuming) != 0 {
		if w.resuming[0] != nil {
			return w.resuming[0]
		}
		w.resuming = w.resuming[1:len(w.resuming)]
	}
	return nil
}
func (w *watchGrpcStream) dispatchEvent(pbresp *pb.WatchResponse) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	events := make([]*Event, len(pbresp.Events))
	for i, ev := range pbresp.Events {
		events[i] = (*Event)(ev)
	}
	wr := &WatchResponse{Header: *pbresp.Header, Events: events, CompactRevision: pbresp.CompactRevision, Created: pbresp.Created, Canceled: pbresp.Canceled, cancelReason: pbresp.CancelReason}
	if wr.IsProgressNotify() && pbresp.WatchId == -1 {
		return w.broadcastResponse(wr)
	}
	return w.unicastResponse(wr, pbresp.WatchId)
}
func (w *watchGrpcStream) broadcastResponse(wr *WatchResponse) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, ws := range w.substreams {
		select {
		case ws.recvc <- wr:
		case <-ws.donec:
		}
	}
	return true
}
func (w *watchGrpcStream) unicastResponse(wr *WatchResponse, watchId int64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ws, ok := w.substreams[watchId]
	if !ok {
		return false
	}
	select {
	case ws.recvc <- wr:
	case <-ws.donec:
		return false
	}
	return true
}
func (w *watchGrpcStream) serveWatchClient(wc pb.Watch_WatchClient) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		resp, err := wc.Recv()
		if err != nil {
			select {
			case w.errc <- err:
			case <-w.donec:
			}
			return
		}
		select {
		case w.respc <- resp:
		case <-w.donec:
			return
		}
	}
}
func (w *watchGrpcStream) serveSubstream(ws *watcherStream, resumec chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ws.closing {
		panic("created substream goroutine but substream is closing")
	}
	nextRev := ws.initReq.rev
	resuming := false
	defer func() {
		if !resuming {
			ws.closing = true
		}
		close(ws.donec)
		if !resuming {
			w.closingc <- ws
		}
		w.wg.Done()
	}()
	emptyWr := &WatchResponse{}
	for {
		curWr := emptyWr
		outc := ws.outc
		if len(ws.buf) > 0 {
			curWr = ws.buf[0]
		} else {
			outc = nil
		}
		select {
		case outc <- *curWr:
			if ws.buf[0].Err() != nil {
				return
			}
			ws.buf[0] = nil
			ws.buf = ws.buf[1:]
		case wr, ok := <-ws.recvc:
			if !ok {
				return
			}
			if wr.Created {
				if ws.initReq.retc != nil {
					ws.initReq.retc <- ws.outc
					ws.initReq.retc = nil
					if ws.initReq.createdNotify {
						ws.outc <- *wr
					}
					if ws.initReq.rev == 0 {
						nextRev = wr.Header.Revision
					}
				}
			} else {
				nextRev = wr.Header.Revision
			}
			if len(wr.Events) > 0 {
				nextRev = wr.Events[len(wr.Events)-1].Kv.ModRevision + 1
			}
			ws.initReq.rev = nextRev
			if wr.Created {
				continue
			}
			ws.buf = append(ws.buf, wr)
		case <-w.ctx.Done():
			return
		case <-ws.initReq.ctx.Done():
			return
		case <-resumec:
			resuming = true
			return
		}
	}
}
func (w *watchGrpcStream) newWatchClient() (pb.Watch_WatchClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(w.resumec)
	w.resumec = make(chan struct{})
	w.joinSubstreams()
	for _, ws := range w.substreams {
		ws.id = -1
		w.resuming = append(w.resuming, ws)
	}
	var resuming []*watcherStream
	for _, ws := range w.resuming {
		if ws != nil {
			resuming = append(resuming, ws)
		}
	}
	w.resuming = resuming
	w.substreams = make(map[int64]*watcherStream)
	stopc := make(chan struct{})
	donec := w.waitCancelSubstreams(stopc)
	wc, err := w.openWatchClient()
	close(stopc)
	<-donec
	for _, ws := range w.resuming {
		if ws.closing {
			continue
		}
		ws.donec = make(chan struct{})
		w.wg.Add(1)
		go w.serveSubstream(ws, w.resumec)
	}
	if err != nil {
		return nil, v3rpc.Error(err)
	}
	go w.serveWatchClient(wc)
	return wc, nil
}
func (w *watchGrpcStream) waitCancelSubstreams(stopc <-chan struct{}) <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wg sync.WaitGroup
	wg.Add(len(w.resuming))
	donec := make(chan struct{})
	for i := range w.resuming {
		go func(ws *watcherStream) {
			defer wg.Done()
			if ws.closing {
				if ws.initReq.ctx.Err() != nil && ws.outc != nil {
					close(ws.outc)
					ws.outc = nil
				}
				return
			}
			select {
			case <-ws.initReq.ctx.Done():
				ws.closing = true
				close(ws.outc)
				ws.outc = nil
				w.wg.Add(1)
				go func() {
					defer w.wg.Done()
					w.closingc <- ws
				}()
			case <-stopc:
			}
		}(w.resuming[i])
	}
	go func() {
		defer close(donec)
		wg.Wait()
	}()
	return donec
}
func (w *watchGrpcStream) joinSubstreams() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, ws := range w.substreams {
		<-ws.donec
	}
	for _, ws := range w.resuming {
		if ws != nil {
			<-ws.donec
		}
	}
}

var maxBackoff = 100 * time.Millisecond

func (w *watchGrpcStream) openWatchClient() (ws pb.Watch_WatchClient, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	backoff := time.Millisecond
	for {
		select {
		case <-w.ctx.Done():
			if err == nil {
				return nil, w.ctx.Err()
			}
			return nil, err
		default:
		}
		if ws, err = w.remote.Watch(w.ctx, w.callOpts...); ws != nil && err == nil {
			break
		}
		if isHaltErr(w.ctx, err) {
			return nil, v3rpc.Error(err)
		}
		if isUnavailableErr(w.ctx, err) {
			if backoff < maxBackoff {
				backoff = backoff + backoff/4
				if backoff > maxBackoff {
					backoff = maxBackoff
				}
			}
			time.Sleep(backoff)
		}
	}
	return ws, nil
}
func (wr *watchRequest) toPB() *pb.WatchRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := &pb.WatchCreateRequest{StartRevision: wr.rev, Key: []byte(wr.key), RangeEnd: []byte(wr.end), ProgressNotify: wr.progressNotify, Filters: wr.filters, PrevKv: wr.prevKV, Fragment: wr.fragment}
	cr := &pb.WatchRequest_CreateRequest{CreateRequest: req}
	return &pb.WatchRequest{RequestUnion: cr}
}
func (pr *progressRequest) toPB() *pb.WatchRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := &pb.WatchProgressRequest{}
	cr := &pb.WatchRequest_ProgressRequest{ProgressRequest: req}
	return &pb.WatchRequest{RequestUnion: cr}
}
func streamKeyFromCtx(ctx context.Context) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		return fmt.Sprintf("%+v", md)
	}
	return ""
}
