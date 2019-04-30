package clientv3

import (
	"io/ioutil"
	"sync"
	"go.etcd.io/etcd/pkg/logutil"
	"google.golang.org/grpc/grpclog"
)

var (
	lgMu	sync.RWMutex
	lg	logutil.Logger
)

type settableLogger struct {
	l	grpclog.LoggerV2
	mu	sync.RWMutex
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg = &settableLogger{}
	SetLogger(grpclog.NewLoggerV2(ioutil.Discard, ioutil.Discard, ioutil.Discard))
}
func SetLogger(l grpclog.LoggerV2) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lgMu.Lock()
	lg = logutil.NewLogger(l)
	grpclog.SetLoggerV2(lg)
	lgMu.Unlock()
}
func GetLogger() logutil.Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lgMu.RLock()
	l := lg
	lgMu.RUnlock()
	return l
}
func NewLogger(gl grpclog.LoggerV2) logutil.Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &settableLogger{l: gl}
}
func (s *settableLogger) get() grpclog.LoggerV2 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.RLock()
	l := s.l
	s.mu.RUnlock()
	return l
}
func (s *settableLogger) Info(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Info(args...)
}
func (s *settableLogger) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Infof(format, args...)
}
func (s *settableLogger) Infoln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Infoln(args...)
}
func (s *settableLogger) Warning(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Warning(args...)
}
func (s *settableLogger) Warningf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Warningf(format, args...)
}
func (s *settableLogger) Warningln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Warningln(args...)
}
func (s *settableLogger) Error(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Error(args...)
}
func (s *settableLogger) Errorf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Errorf(format, args...)
}
func (s *settableLogger) Errorln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Errorln(args...)
}
func (s *settableLogger) Fatal(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Fatal(args...)
}
func (s *settableLogger) Fatalf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Fatalf(format, args...)
}
func (s *settableLogger) Fatalln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Fatalln(args...)
}
func (s *settableLogger) Print(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Info(args...)
}
func (s *settableLogger) Printf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Infof(format, args...)
}
func (s *settableLogger) Println(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.get().Infoln(args...)
}
func (s *settableLogger) V(l int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.get().V(l)
}
func (s *settableLogger) Lvl(lvl int) grpclog.LoggerV2 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.RLock()
	l := s.l
	s.mu.RUnlock()
	if l.V(lvl) {
		return s
	}
	return logutil.NewDiscardLogger()
}
