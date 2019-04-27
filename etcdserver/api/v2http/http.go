package v2http

import (
	"math"
	"net/http"
	"strings"
	"time"
	"github.com/coreos/etcd/etcdserver/api/etcdhttp"
	"github.com/coreos/etcd/etcdserver/api/v2http/httptypes"
	"github.com/coreos/etcd/etcdserver/auth"
	"github.com/coreos/etcd/pkg/logutil"
	"github.com/coreos/pkg/capnslog"
)

const (
	defaultWatchTimeout = time.Duration(math.MaxInt64)
)

var (
	plog	= capnslog.NewPackageLogger("github.com/coreos/etcd", "etcdserver/api/v2http")
	mlog	= logutil.NewMergeLogger(plog)
)

func writeError(w http.ResponseWriter, r *http.Request, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return
	}
	if e, ok := err.(auth.Error); ok {
		herr := httptypes.NewHTTPError(e.HTTPStatus(), e.Error())
		if et := herr.WriteTo(w); et != nil {
			plog.Debugf("error writing HTTPError (%v) to %s", et, r.RemoteAddr)
		}
		return
	}
	etcdhttp.WriteError(w, r, err)
}
func allowMethod(w http.ResponseWriter, m string, ms ...string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func requestLogger(handler http.Handler) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		plog.Debugf("[%s] %s remote:%s", r.Method, r.RequestURI, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}
