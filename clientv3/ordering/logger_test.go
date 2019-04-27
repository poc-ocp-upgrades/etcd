package ordering

import (
	"io/ioutil"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/pkg/capnslog"
	"google.golang.org/grpc/grpclog"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	capnslog.SetGlobalLogLevel(capnslog.CRITICAL)
	clientv3.SetLogger(grpclog.NewLoggerV2(ioutil.Discard, ioutil.Discard, ioutil.Discard))
}
