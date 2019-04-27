package snap

import "github.com/prometheus/client_golang/prometheus"

var (
	saveDurations		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd_debugging", Subsystem: "snap", Name: "save_total_duration_seconds", Help: "The total latency distributions of save called by snapshot.", Buckets: prometheus.ExponentialBuckets(0.001, 2, 14)})
	marshallingDurations	= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd_debugging", Subsystem: "snap", Name: "save_marshalling_duration_seconds", Help: "The marshalling cost distributions of save called by snapshot.", Buckets: prometheus.ExponentialBuckets(0.001, 2, 14)})
	snapDBSaveSec		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "snap_db", Name: "save_total_duration_seconds", Help: "The total latency distributions of v3 snapshot save", Buckets: prometheus.ExponentialBuckets(0.1, 2, 10)})
	snapDBFsyncSec		= prometheus.NewHistogram(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "snap_db", Name: "fsync_duration_seconds", Help: "The latency distributions of fsyncing .snap.db file", Buckets: prometheus.ExponentialBuckets(0.001, 2, 14)})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prometheus.MustRegister(saveDurations)
	prometheus.MustRegister(marshallingDurations)
	prometheus.MustRegister(snapDBSaveSec)
	prometheus.MustRegister(snapDBFsyncSec)
}
