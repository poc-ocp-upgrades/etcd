package integration

import (
	"io/ioutil"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/pkg/capnslog"
	"google.golang.org/grpc/grpclog"
)

const defaultLogLevel = capnslog.CRITICAL

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	capnslog.SetGlobalLogLevel(defaultLogLevel)
	clientv3.SetLogger(grpclog.NewLoggerV2(ioutil.Discard, ioutil.Discard, ioutil.Discard))
}
