package rafthttp

import "github.com/prometheus/client_golang/prometheus"

var (
	sentBytes               = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "network", Name: "peer_sent_bytes_total", Help: "The total number of bytes sent to peers."}, []string{"To"})
	receivedBytes           = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "network", Name: "peer_received_bytes_total", Help: "The total number of bytes received from peers."}, []string{"From"})
	sentFailures            = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "network", Name: "peer_sent_failures_total", Help: "The total number of send failures from peers."}, []string{"To"})
	recvFailures            = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "network", Name: "peer_received_failures_total", Help: "The total number of receive failures from peers."}, []string{"From"})
	snapshotSend            = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "network", Name: "snapshot_send_success", Help: "Total number of successful snapshot sends"}, []string{"To"})
	snapshotSendFailures    = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "network", Name: "snapshot_send_failures", Help: "Total number of snapshot send failures"}, []string{"To"})
	snapshotSendSeconds     = prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "network", Name: "snapshot_send_total_duration_seconds", Help: "Total latency distributions of v3 snapshot sends", Buckets: prometheus.ExponentialBuckets(0.1, 2, 10)}, []string{"To"})
	snapshotReceive         = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "network", Name: "snapshot_receive_success", Help: "Total number of successful snapshot receives"}, []string{"From"})
	snapshotReceiveFailures = prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "network", Name: "snapshot_receive_failures", Help: "Total number of snapshot receive failures"}, []string{"From"})
	snapshotReceiveSeconds  = prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "network", Name: "snapshot_receive_total_duration_seconds", Help: "Total latency distributions of v3 snapshot receives", Buckets: prometheus.ExponentialBuckets(0.1, 2, 10)}, []string{"From"})
	rtts                    = prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "network", Name: "peer_round_trip_time_seconds", Help: "Round-Trip-Time histogram between peers.", Buckets: prometheus.ExponentialBuckets(0.0001, 2, 14)}, []string{"To"})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prometheus.MustRegister(sentBytes)
	prometheus.MustRegister(receivedBytes)
	prometheus.MustRegister(sentFailures)
	prometheus.MustRegister(recvFailures)
	prometheus.MustRegister(snapshotSend)
	prometheus.MustRegister(snapshotSendFailures)
	prometheus.MustRegister(snapshotSendSeconds)
	prometheus.MustRegister(snapshotReceive)
	prometheus.MustRegister(snapshotReceiveFailures)
	prometheus.MustRegister(snapshotReceiveSeconds)
	prometheus.MustRegister(rtts)
}
