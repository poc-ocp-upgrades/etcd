package etcdserver

import (
	"context"
	"time"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/store"
)

type RequestV2 pb.Request
type RequestV2Handler interface {
	Post(ctx context.Context, r *RequestV2) (Response, error)
	Put(ctx context.Context, r *RequestV2) (Response, error)
	Delete(ctx context.Context, r *RequestV2) (Response, error)
	QGet(ctx context.Context, r *RequestV2) (Response, error)
	Get(ctx context.Context, r *RequestV2) (Response, error)
	Head(ctx context.Context, r *RequestV2) (Response, error)
}
type reqV2HandlerEtcdServer struct {
	reqV2HandlerStore
	s	*EtcdServer
}
type reqV2HandlerStore struct {
	store	store.Store
	applier	ApplierV2
}

func NewStoreRequestV2Handler(s store.Store, applier ApplierV2) RequestV2Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &reqV2HandlerStore{s, applier}
}
func (a *reqV2HandlerStore) Post(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.applier.Post(r), nil
}
func (a *reqV2HandlerStore) Put(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.applier.Put(r), nil
}
func (a *reqV2HandlerStore) Delete(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.applier.Delete(r), nil
}
func (a *reqV2HandlerStore) QGet(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.applier.QGet(r), nil
}
func (a *reqV2HandlerStore) Get(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.Wait {
		wc, err := a.store.Watch(r.Path, r.Recursive, r.Stream, r.Since)
		return Response{Watcher: wc}, err
	}
	ev, err := a.store.Get(r.Path, r.Recursive, r.Sorted)
	return Response{Event: ev}, err
}
func (a *reqV2HandlerStore) Head(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ev, err := a.store.Get(r.Path, r.Recursive, r.Sorted)
	return Response{Event: ev}, err
}
func (a *reqV2HandlerEtcdServer) Post(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.processRaftRequest(ctx, r)
}
func (a *reqV2HandlerEtcdServer) Put(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.processRaftRequest(ctx, r)
}
func (a *reqV2HandlerEtcdServer) Delete(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.processRaftRequest(ctx, r)
}
func (a *reqV2HandlerEtcdServer) QGet(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a.processRaftRequest(ctx, r)
}
func (a *reqV2HandlerEtcdServer) processRaftRequest(ctx context.Context, r *RequestV2) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := ((*pb.Request)(r)).Marshal()
	if err != nil {
		return Response{}, err
	}
	ch := a.s.w.Register(r.ID)
	start := time.Now()
	a.s.r.Propose(ctx, data)
	proposalsPending.Inc()
	defer proposalsPending.Dec()
	select {
	case x := <-ch:
		resp := x.(Response)
		return resp, resp.Err
	case <-ctx.Done():
		proposalsFailed.Inc()
		a.s.w.Trigger(r.ID, nil)
		return Response{}, a.s.parseProposeCtxErr(ctx.Err(), start)
	case <-a.s.stopping:
	}
	return Response{}, ErrStopped
}
func (s *EtcdServer) Do(ctx context.Context, r pb.Request) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.ID = s.reqIDGen.Next()
	h := &reqV2HandlerEtcdServer{reqV2HandlerStore: reqV2HandlerStore{store: s.store, applier: s.applyV2}, s: s}
	rp := &r
	resp, err := ((*RequestV2)(rp)).Handle(ctx, h)
	resp.Term, resp.Index = s.Term(), s.Index()
	return resp, err
}
func (r *RequestV2) Handle(ctx context.Context, v2api RequestV2Handler) (Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if r.Method == "GET" && r.Quorum {
		r.Method = "QGET"
	}
	switch r.Method {
	case "POST":
		return v2api.Post(ctx, r)
	case "PUT":
		return v2api.Put(ctx, r)
	case "DELETE":
		return v2api.Delete(ctx, r)
	case "QGET":
		return v2api.QGet(ctx, r)
	case "GET":
		return v2api.Get(ctx, r)
	case "HEAD":
		return v2api.Head(ctx, r)
	}
	return Response{}, ErrUnknownMethod
}
func (r *RequestV2) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rpb := pb.Request(*r)
	return rpb.String()
}
