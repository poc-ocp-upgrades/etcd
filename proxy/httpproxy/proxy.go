package httpproxy

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"golang.org/x/net/http2"
)

const (
	DefaultMaxIdleConnsPerHost = 128
)

type GetProxyURLs func() []string

func NewHandler(t *http.Transport, urlsFunc GetProxyURLs, failureWait time.Duration, refreshInterval time.Duration) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.TLSClientConfig != nil {
		err := http2.ConfigureTransport(t)
		if err != nil {
			plog.Infof("Error enabling Transport HTTP/2 support: %v", err)
		}
	}
	p := &reverseProxy{director: newDirector(urlsFunc, failureWait, refreshInterval), transport: t}
	mux := http.NewServeMux()
	mux.Handle("/", p)
	mux.HandleFunc("/v2/config/local/proxy", p.configHandler)
	return mux
}
func NewReadonlyHandler(hdlr http.Handler) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	readonly := readonlyHandlerFunc(hdlr)
	return http.HandlerFunc(readonly)
}
func readonlyHandlerFunc(next http.Handler) func(http.ResponseWriter, *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		next.ServeHTTP(w, req)
	}
}
func (p *reverseProxy) configHandler(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !allowMethod(w, r.Method, "GET") {
		return
	}
	eps := p.director.endpoints()
	epstr := make([]string, len(eps))
	for i, e := range eps {
		epstr[i] = e.URL.String()
	}
	proxyConfig := struct {
		Endpoints []string `json:"endpoints"`
	}{Endpoints: epstr}
	json.NewEncoder(w).Encode(proxyConfig)
}
func allowMethod(w http.ResponseWriter, m string, ms ...string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, meth := range ms {
		if m == meth {
			return true
		}
	}
	w.Header().Set("Allow", strings.Join(ms, ","))
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return false
}
