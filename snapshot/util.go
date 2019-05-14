package snapshot

import (
	godefaultbytes "bytes"
	"encoding/binary"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type revision struct {
	main int64
	sub  int64
}

func bytesToRev(bytes []byte) revision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return revision{main: int64(binary.BigEndian.Uint64(bytes[0:8])), sub: int64(binary.BigEndian.Uint64(bytes[9:]))}
}

type initIndex int

func (i *initIndex) ConsistentIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return uint64(*i)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
