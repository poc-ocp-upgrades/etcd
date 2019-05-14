package grpcproxy

import (
	"fmt"
	"github.com/coreos/etcd/etcdserver/api/etcdhttp"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	watchersCoalescing = prometheus.NewGauge(prometheus.GaugeOpts{Namespace: "etcd", Subsystem: "grpc_proxy", Name: "watchers_coalescing_total", Help: "Total number of current watchers coalescing"})
	eventsCoalescing   = prometheus.NewCounter(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "grpc_proxy", Name: "events_coalescing_total", Help: "Total number of events coalescing"})
	cacheKeys          = prometheus.NewGauge(prometheus.GaugeOpts{Namespace: "etcd", Subsystem: "grpc_proxy", Name: "cache_keys_total", Help: "Total number of keys/ranges cached"})
	cacheHits          = prometheus.NewGauge(prometheus.GaugeOpts{Namespace: "etcd", Subsystem: "grpc_proxy", Name: "cache_hits_total", Help: "Total number of cache hits"})
	cachedMisses       = prometheus.NewGauge(prometheus.GaugeOpts{Namespace: "etcd", Subsystem: "grpc_proxy", Name: "cache_misses_total", Help: "Total number of cache misses"})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prometheus.MustRegister(watchersCoalescing)
	prometheus.MustRegister(eventsCoalescing)
	prometheus.MustRegister(cacheKeys)
	prometheus.MustRegister(cacheHits)
	prometheus.MustRegister(cachedMisses)
}
func HandleMetrics(mux *http.ServeMux, c *http.Client, eps []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	if len(eps) > 1 {
		eps = shuffleEndpoints(r, eps)
	}
	pathMetrics := etcdhttp.PathMetrics
	mux.HandleFunc(pathMetrics, func(w http.ResponseWriter, r *http.Request) {
		target := fmt.Sprintf("%s%s", eps[0], pathMetrics)
		if !strings.HasPrefix(target, "http") {
			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}
			target = fmt.Sprintf("%s://%s", scheme, target)
		}
		resp, err := c.Get(target)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		defer resp.Body.Close()
		w.Header().Set("Content-Type", "text/plain; version=0.0.4")
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(w, "%s", body)
	})
}
func shuffleEndpoints(r *rand.Rand, eps []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := len(eps)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		j := r.Intn(i + 1)
		p[i] = p[j]
		p[j] = i
	}
	neps := make([]string, n)
	for i, k := range p {
		neps[i] = eps[k]
	}
	return neps
}
