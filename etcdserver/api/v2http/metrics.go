package v2http

import (
	"strconv"
	"time"
	"net/http"
	"go.etcd.io/etcd/etcdserver/api/v2error"
	"go.etcd.io/etcd/etcdserver/api/v2http/httptypes"
	"go.etcd.io/etcd/etcdserver/etcdserverpb"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	incomingEvents			= prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "http", Name: "received_total", Help: "Counter of requests received into the system (successfully parsed and authd)."}, []string{"method"})
	failedEvents			= prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "http", Name: "failed_total", Help: "Counter of handle failures of requests (non-watches), by method (GET/PUT etc.) and code (400, 500 etc.)."}, []string{"method", "code"})
	successfulEventsHandlingSec	= prometheus.NewHistogramVec(prometheus.HistogramOpts{Namespace: "etcd", Subsystem: "http", Name: "successful_duration_seconds", Help: "Bucketed histogram of processing time (s) of successfully handled requests (non-watches), by method (GET/PUT etc.).", Buckets: prometheus.ExponentialBuckets(0.0005, 2, 13)}, []string{"method"})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prometheus.MustRegister(incomingEvents)
	prometheus.MustRegister(failedEvents)
	prometheus.MustRegister(successfulEventsHandlingSec)
}
func reportRequestReceived(request etcdserverpb.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	incomingEvents.WithLabelValues(methodFromRequest(request)).Inc()
}
func reportRequestCompleted(request etcdserverpb.Request, startTime time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	method := methodFromRequest(request)
	successfulEventsHandlingSec.WithLabelValues(method).Observe(time.Since(startTime).Seconds())
}
func reportRequestFailed(request etcdserverpb.Request, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	method := methodFromRequest(request)
	failedEvents.WithLabelValues(method, strconv.Itoa(codeFromError(err))).Inc()
}
func methodFromRequest(request etcdserverpb.Request) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if request.Method == "GET" && request.Quorum {
		return "QGET"
	}
	return request.Method
}
func codeFromError(err error) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return http.StatusInternalServerError
	}
	switch e := err.(type) {
	case *v2error.Error:
		return e.StatusCode()
	case *httptypes.HTTPError:
		return e.Code
	default:
		return http.StatusInternalServerError
	}
}
