package etcdhttp

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/raft"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	PathMetrics	= "/metrics"
	PathHealth	= "/health"
)

func HandleMetricsHealth(mux *http.ServeMux, srv etcdserver.ServerV2) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle(PathMetrics, promhttp.Handler())
	mux.Handle(PathHealth, NewHealthHandler(func() Health {
		return checkHealth(srv)
	}))
}
func HandlePrometheus(mux *http.ServeMux) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle(PathMetrics, promhttp.Handler())
}
func NewHealthHandler(hfunc func() Health) http.HandlerFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		h := hfunc()
		d, _ := json.Marshal(h)
		if h.Health != "true" {
			http.Error(w, string(d), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(d)
	}
}

var (
	healthSuccess	= prometheus.NewCounter(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "server", Name: "health_success", Help: "The total number of successful health checks"})
	healthFailed	= prometheus.NewCounter(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "server", Name: "health_failures", Help: "The total number of failed health checks"})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prometheus.MustRegister(healthSuccess)
	prometheus.MustRegister(healthFailed)
}

type Health struct {
	Health string `json:"health"`
}

func checkHealth(srv etcdserver.ServerV2) Health {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	h := Health{Health: "true"}
	as := srv.Alarms()
	if len(as) > 0 {
		h.Health = "false"
	}
	if h.Health == "true" {
		if uint64(srv.Leader()) == raft.None {
			h.Health = "false"
		}
	}
	if h.Health == "true" {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		_, err := srv.Do(ctx, etcdserverpb.Request{Method: "QGET"})
		cancel()
		if err != nil {
			h.Health = "false"
		}
	}
	if h.Health == "true" {
		healthSuccess.Inc()
	} else {
		healthFailed.Inc()
	}
	return h
}
