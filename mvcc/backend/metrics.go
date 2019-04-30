package backend

import "github.com/prometheus/client_golang/prometheus"

var (
	commitSec		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "disk", Name: "backend_commit_duration_seconds", Help: "The latency distributions of commit called by backend.", Buckets: prometheus.ExponentialBuckets(0.001, 2, 14)})
	rebalanceSec		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd_debugging", Subsystem: "disk", Name: "backend_commit_rebalance_duration_seconds", Help: "The latency distributions of commit.rebalance called by bboltdb backend.", Buckets: prometheus.ExponentialBuckets(0.001, 2, 14)})
	spillSec		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd_debugging", Subsystem: "disk", Name: "backend_commit_spill_duration_seconds", Help: "The latency distributions of commit.spill called by bboltdb backend.", Buckets: prometheus.ExponentialBuckets(0.001, 2, 14)})
	writeSec		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd_debugging", Subsystem: "disk", Name: "backend_commit_write_duration_seconds", Help: "The latency distributions of commit.write called by bboltdb backend.", Buckets: prometheus.ExponentialBuckets(0.001, 2, 14)})
	defragSec		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "disk", Name: "backend_defrag_duration_seconds", Help: "The latency distribution of backend defragmentation.", Buckets: prometheus.ExponentialBuckets(.1, 2, 13)})
	snapshotTransferSec	= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "disk", Name: "backend_snapshot_duration_seconds", Help: "The latency distribution of backend snapshots.", Buckets: prometheus.ExponentialBuckets(.01, 2, 17)})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prometheus.MustRegister(commitSec)
	prometheus.MustRegister(rebalanceSec)
	prometheus.MustRegister(spillSec)
	prometheus.MustRegister(writeSec)
	prometheus.MustRegister(defragSec)
	prometheus.MustRegister(snapshotTransferSec)
}
