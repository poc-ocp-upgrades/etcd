package mvcc

import (
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func (tw *watchableStoreTxnWrite) End() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	changes := tw.Changes()
	if len(changes) == 0 {
		tw.TxnWrite.End()
		return
	}
	rev := tw.Rev() + 1
	evs := make([]mvccpb.Event, len(changes))
	for i, change := range changes {
		evs[i].Kv = &changes[i]
		if change.CreateRevision == 0 {
			evs[i].Type = mvccpb.DELETE
			evs[i].Kv.ModRevision = rev
		} else {
			evs[i].Type = mvccpb.PUT
		}
	}
	tw.s.mu.Lock()
	tw.s.notify(rev, evs)
	tw.TxnWrite.End()
	tw.s.mu.Unlock()
}

type watchableStoreTxnWrite struct {
	TxnWrite
	s	*watchableStore
}

func (s *watchableStore) Write() TxnWrite {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &watchableStoreTxnWrite{s.store.Write(), s}
}
