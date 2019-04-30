package v2http

import (
	"math"
	"net/http"
	"strings"
	"time"
	"go.etcd.io/etcd/etcdserver/api/etcdhttp"
	"go.etcd.io/etcd/etcdserver/api/v2auth"
	"go.etcd.io/etcd/etcdserver/api/v2http/httptypes"
	"go.etcd.io/etcd/pkg/logutil"
	"github.com/coreos/pkg/capnslog"
	"go.uber.org/zap"
)

const (
	defaultWatchTimeout = time.Duration(math.MaxInt64)
)

var (
	plog	= capnslog.NewPackageLogger("go.etcd.io/etcd", "etcdserver/api/v2http")
	mlog	= logutil.NewMergeLogger(plog)
)

func writeError(lg *zap.Logger, w http.ResponseWriter, r *http.Request, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return
	}
	if e, ok := err.(v2auth.Error); ok {
		herr := httptypes.NewHTTPError(e.HTTPStatus(), e.Error())
		if et := herr.WriteTo(w); et != nil {
			if lg != nil {
				lg.Debug("failed to write v2 HTTP error", zap.String("remote-addr", r.RemoteAddr), zap.String("v2auth-error", e.Error()), zap.Error(et))
			} else {
				plog.Debugf("error writing HTTPError (%v) to %s", et, r.RemoteAddr)
			}
		}
		return
	}
	etcdhttp.WriteError(lg, w, r, err)
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
func requestLogger(lg *zap.Logger, handler http.Handler) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if lg != nil {
			lg.Debug("handling HTTP request", zap.String("method", r.Method), zap.String("request-uri", r.RequestURI), zap.String("remote-addr", r.RemoteAddr))
		} else {
			plog.Debugf("[%s] %s remote:%s", r.Method, r.RequestURI, r.RemoteAddr)
		}
		handler.ServeHTTP(w, r)
	})
}
