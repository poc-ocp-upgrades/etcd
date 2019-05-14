package recipe

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func WaitEvents(c *clientv3.Client, key string, rev int64, evs []mvccpb.Event_EventType) (*clientv3.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wc := c.Watch(context.Background(), key, clientv3.WithRev(rev))
	if wc == nil {
		return nil, ErrNoWatcher
	}
	return waitEvents(wc, evs), nil
}
func WaitPrefixEvents(c *clientv3.Client, prefix string, rev int64, evs []mvccpb.Event_EventType) (*clientv3.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wc := c.Watch(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithRev(rev))
	if wc == nil {
		return nil, ErrNoWatcher
	}
	return waitEvents(wc, evs), nil
}
func waitEvents(wc clientv3.WatchChan, evs []mvccpb.Event_EventType) *clientv3.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	for wresp := range wc {
		for _, ev := range wresp.Events {
			if ev.Type == evs[i] {
				i++
				if i == len(evs) {
					return ev
				}
			}
		}
	}
	return nil
}
