package clientv3

import (
	"io/ioutil"
	"sync"
	"google.golang.org/grpc/grpclog"
)

type Logger interface {
	grpclog.LoggerV2
	Lvl(lvl int) Logger
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

var (
	loggerMu	sync.RWMutex
	logger		Logger
)

type settableLogger struct {
	l	grpclog.LoggerV2
	mu	sync.RWMutex
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger = &settableLogger{}
	SetLogger(grpclog.NewLoggerV2(ioutil.Discard, ioutil.Discard, ioutil.Discard))
}
func SetLogger(l grpclog.LoggerV2) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	loggerMu.Lock()
	logger = NewLogger(l)
	grpclog.SetLoggerV2(logger)
	loggerMu.Unlock()
}
func GetLogger() Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	loggerMu.RLock()
	l := logger
	loggerMu.RUnlock()
	return l
}
func NewLogger(gl grpclog.LoggerV2) Logger {
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
func (s *settableLogger) Lvl(lvl int) Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.RLock()
	l := s.l
	s.mu.RUnlock()
	if l.V(lvl) {
		return s
	}
	return &noLogger{}
}

type noLogger struct{}

func (*noLogger) Info(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Infoln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Warning(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Warningf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Warningln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Error(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Errorf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Errorln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Fatal(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Fatalf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Fatalln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Print(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Printf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) Println(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*noLogger) V(l int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (ng *noLogger) Lvl(lvl int) Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ng
}
