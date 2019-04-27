package mvcc

import (
	"github.com/coreos/etcd/lease"
)

type metricsTxnWrite struct {
	TxnWrite
	ranges	uint
	puts	uint
	deletes	uint
}

func newMetricsTxnRead(tr TxnRead) TxnRead {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &metricsTxnWrite{&txnReadWrite{tr}, 0, 0, 0}
}
func newMetricsTxnWrite(tw TxnWrite) TxnWrite {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &metricsTxnWrite{tw, 0, 0, 0}
}
func (tw *metricsTxnWrite) Range(key, end []byte, ro RangeOptions) (*RangeResult, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tw.ranges++
	return tw.TxnWrite.Range(key, end, ro)
}
func (tw *metricsTxnWrite) DeleteRange(key, end []byte) (n, rev int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tw.deletes++
	return tw.TxnWrite.DeleteRange(key, end)
}
func (tw *metricsTxnWrite) Put(key, value []byte, lease lease.LeaseID) (rev int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tw.puts++
	return tw.TxnWrite.Put(key, value, lease)
}
func (tw *metricsTxnWrite) End() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer tw.TxnWrite.End()
	if sum := tw.ranges + tw.puts + tw.deletes; sum > 1 {
		txnCounter.Inc()
	}
	rangeCounter.Add(float64(tw.ranges))
	putCounter.Add(float64(tw.puts))
	deleteCounter.Add(float64(tw.deletes))
}
