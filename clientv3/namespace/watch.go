package namespace

import (
	"context"
	"sync"
	"go.etcd.io/etcd/clientv3"
)

type watcherPrefix struct {
	clientv3.Watcher
	pfx		string
	wg		sync.WaitGroup
	stopc		chan struct{}
	stopOnce	sync.Once
}

func NewWatcher(w clientv3.Watcher, prefix string) clientv3.Watcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &watcherPrefix{Watcher: w, pfx: prefix, stopc: make(chan struct{})}
}
func (w *watcherPrefix) Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	op := clientv3.OpGet(key, opts...)
	end := op.RangeBytes()
	pfxBegin, pfxEnd := prefixInterval(w.pfx, []byte(key), end)
	if pfxEnd != nil {
		opts = append(opts, clientv3.WithRange(string(pfxEnd)))
	}
	wch := w.Watcher.Watch(ctx, string(pfxBegin), opts...)
	pfxWch := make(chan clientv3.WatchResponse)
	w.wg.Add(1)
	go func() {
		defer func() {
			close(pfxWch)
			w.wg.Done()
		}()
		for wr := range wch {
			for i := range wr.Events {
				wr.Events[i].Kv.Key = wr.Events[i].Kv.Key[len(w.pfx):]
				if wr.Events[i].PrevKv != nil {
					wr.Events[i].PrevKv.Key = wr.Events[i].Kv.Key
				}
			}
			select {
			case pfxWch <- wr:
			case <-ctx.Done():
				return
			case <-w.stopc:
				return
			}
		}
	}()
	return pfxWch
}
func (w *watcherPrefix) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := w.Watcher.Close()
	w.stopOnce.Do(func() {
		close(w.stopc)
	})
	w.wg.Wait()
	return err
}
