package v2v3

import (
	"context"
	"strings"
	"github.com/coreos/etcd/clientv3"
	etcdErr "github.com/coreos/etcd/error"
	"github.com/coreos/etcd/store"
)

func (s *v2v3Store) Watch(prefix string, recursive, stream bool, sinceIndex uint64) (store.Watcher, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithCancel(s.ctx)
	wch := s.c.Watch(ctx, s.pfx, clientv3.WithPrefix(), clientv3.WithRev(int64(sinceIndex)), clientv3.WithCreatedNotify(), clientv3.WithPrevKV())
	resp, ok := <-wch
	if err := resp.Err(); err != nil || !ok {
		cancel()
		return nil, etcdErr.NewError(etcdErr.EcodeRaftInternal, prefix, 0)
	}
	evc, donec := make(chan *store.Event), make(chan struct{})
	go func() {
		defer func() {
			close(evc)
			close(donec)
		}()
		for resp := range wch {
			for _, ev := range s.mkV2Events(resp) {
				k := ev.Node.Key
				if recursive {
					if !strings.HasPrefix(k, prefix) {
						continue
					}
					k = strings.Replace(k, prefix, "/", 1)
					if strings.Contains(k, "/_") {
						continue
					}
				}
				if !recursive && k != prefix {
					continue
				}
				select {
				case evc <- ev:
				case <-ctx.Done():
					return
				}
				if !stream {
					return
				}
			}
		}
	}()
	return &v2v3Watcher{startRev: resp.Header.Revision, evc: evc, donec: donec, cancel: cancel}, nil
}
func (s *v2v3Store) mkV2Events(wr clientv3.WatchResponse) (evs []*store.Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ak := s.mkActionKey()
	for _, rev := range mkRevs(wr) {
		var act, key *clientv3.Event
		for _, ev := range rev {
			if string(ev.Kv.Key) == ak {
				act = ev
			} else if key != nil && len(key.Kv.Key) < len(ev.Kv.Key) {
				key = ev
			} else if key == nil {
				key = ev
			}
		}
		v2ev := &store.Event{Action: string(act.Kv.Value), Node: s.mkV2Node(key.Kv), PrevNode: s.mkV2Node(key.PrevKv), EtcdIndex: mkV2Rev(wr.Header.Revision)}
		evs = append(evs, v2ev)
	}
	return evs
}
func mkRevs(wr clientv3.WatchResponse) (revs [][]*clientv3.Event) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var curRev []*clientv3.Event
	for _, ev := range wr.Events {
		if curRev != nil && ev.Kv.ModRevision != curRev[0].Kv.ModRevision {
			revs = append(revs, curRev)
			curRev = nil
		}
		curRev = append(curRev, ev)
	}
	if curRev != nil {
		revs = append(revs, curRev)
	}
	return revs
}

type v2v3Watcher struct {
	startRev	int64
	evc		chan *store.Event
	donec		chan struct{}
	cancel		context.CancelFunc
}

func (w *v2v3Watcher) StartIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return mkV2Rev(w.startRev)
}
func (w *v2v3Watcher) Remove() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.cancel()
	<-w.donec
}
func (w *v2v3Watcher) EventChan() chan *store.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return w.evc
}
