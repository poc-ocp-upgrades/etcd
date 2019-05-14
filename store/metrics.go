package store

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	readCounter        = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd_debugging", Subsystem: "store", Name: "reads_total", Help: "Total number of reads action by (get/getRecursive), local to this member."}, []string{"action"})
	writeCounter       = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd_debugging", Subsystem: "store", Name: "writes_total", Help: "Total number of writes (e.g. set/compareAndDelete) seen by this member."}, []string{"action"})
	readFailedCounter  = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd_debugging", Subsystem: "store", Name: "reads_failed_total", Help: "Failed read actions by (get/getRecursive), local to this member."}, []string{"action"})
	writeFailedCounter = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd_debugging", Subsystem: "store", Name: "writes_failed_total", Help: "Failed write actions (e.g. set/compareAndDelete), seen by this member."}, []string{"action"})
	expireCounter      = prometheus.NewCounter(prometheus.CounterOpts{Namespace: "etcd_debugging", Subsystem: "store", Name: "expires_total", Help: "Total number of expired keys."})
	watchRequests      = prometheus.NewCounter(prometheus.CounterOpts{Namespace: "etcd_debugging", Subsystem: "store", Name: "watch_requests_total", Help: "Total number of incoming watch requests (new or reestablished)."})
	watcherCount       = prometheus.NewGauge(prometheus.GaugeOpts{Namespace: "etcd_debugging", Subsystem: "store", Name: "watchers", Help: "Count of currently active watchers."})
)

const (
	GetRecursive = "getRecursive"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if prometheus.Register(readCounter) != nil {
		return
	}
	prometheus.MustRegister(writeCounter)
	prometheus.MustRegister(expireCounter)
	prometheus.MustRegister(watchRequests)
	prometheus.MustRegister(watcherCount)
}
func reportReadSuccess(read_action string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	readCounter.WithLabelValues(read_action).Inc()
}
func reportReadFailure(read_action string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	readCounter.WithLabelValues(read_action).Inc()
	readFailedCounter.WithLabelValues(read_action).Inc()
}
func reportWriteSuccess(write_action string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	writeCounter.WithLabelValues(write_action).Inc()
}
func reportWriteFailure(write_action string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	writeCounter.WithLabelValues(write_action).Inc()
	writeFailedCounter.WithLabelValues(write_action).Inc()
}
func reportExpiredKey() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	expireCounter.Inc()
}
func reportWatchRequest() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	watchRequests.Inc()
}
func reportWatcherAdded() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	watcherCount.Inc()
}
func reportWatcherRemoved() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	watcherCount.Dec()
}
