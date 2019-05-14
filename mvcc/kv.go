package mvcc

import (
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type RangeOptions struct {
	Limit int64
	Rev   int64
	Count bool
}
type RangeResult struct {
	KVs   []mvccpb.KeyValue
	Rev   int64
	Count int
}
type ReadView interface {
	FirstRev() int64
	Rev() int64
	Range(key, end []byte, ro RangeOptions) (r *RangeResult, err error)
}
type TxnRead interface {
	ReadView
	End()
}
type WriteView interface {
	DeleteRange(key, end []byte) (n, rev int64)
	Put(key, value []byte, lease lease.LeaseID) (rev int64)
}
type TxnWrite interface {
	TxnRead
	WriteView
	Changes() []mvccpb.KeyValue
}
type txnReadWrite struct{ TxnRead }

func (trw *txnReadWrite) DeleteRange(key, end []byte) (n, rev int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("unexpected DeleteRange")
}
func (trw *txnReadWrite) Put(key, value []byte, lease lease.LeaseID) (rev int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("unexpected Put")
}
func (trw *txnReadWrite) Changes() []mvccpb.KeyValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func NewReadOnlyTxnWrite(txn TxnRead) TxnWrite {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &txnReadWrite{txn}
}

type KV interface {
	ReadView
	WriteView
	Read() TxnRead
	Write() TxnWrite
	Hash() (hash uint32, revision int64, err error)
	HashByRev(rev int64) (hash uint32, revision int64, compactRev int64, err error)
	Compact(rev int64) (<-chan struct{}, error)
	Commit()
	Restore(b backend.Backend) error
	Close() error
}
type WatchableKV interface {
	KV
	Watchable
}
type Watchable interface{ NewWatchStream() WatchStream }
type ConsistentWatchableKV interface {
	WatchableKV
	ConsistentIndex() uint64
}
