package mvcc

import (
	"github.com/coreos/etcd/lease"
)

type readView struct{ kv KV }

func (rv *readView) FirstRev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := rv.kv.Read()
	defer tr.End()
	return tr.FirstRev()
}
func (rv *readView) Rev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := rv.kv.Read()
	defer tr.End()
	return tr.Rev()
}
func (rv *readView) Range(key, end []byte, ro RangeOptions) (r *RangeResult, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := rv.kv.Read()
	defer tr.End()
	return tr.Range(key, end, ro)
}

type writeView struct{ kv KV }

func (wv *writeView) DeleteRange(key, end []byte) (n, rev int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tw := wv.kv.Write()
	defer tw.End()
	return tw.DeleteRange(key, end)
}
func (wv *writeView) Put(key, value []byte, lease lease.LeaseID) (rev int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tw := wv.kv.Write()
	defer tw.End()
	return tw.Put(key, value, lease)
}
