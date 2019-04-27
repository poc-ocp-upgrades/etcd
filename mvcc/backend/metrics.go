package backend

import "github.com/prometheus/client_golang/prometheus"

var (
	commitDurations		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "disk", Name: "backend_commit_duration_seconds", Help: "The latency distributions of commit called by backend.", Buckets: prometheus.ExponentialBuckets(0.001, 2, 14)})
	defragDurations		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "disk", Name: "backend_defrag_duration_seconds", Help: "The latency distribution of backend defragmentation.", Buckets: prometheus.ExponentialBuckets(.1, 2, 13)})
	snapshotDurations	= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "disk", Name: "backend_snapshot_duration_seconds", Help: "The latency distribution of backend snapshots.", Buckets: prometheus.ExponentialBuckets(.01, 2, 17)})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prometheus.MustRegister(commitDurations)
	prometheus.MustRegister(defragDurations)
	prometheus.MustRegister(snapshotDurations)
}
