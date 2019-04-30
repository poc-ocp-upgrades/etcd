package logutil

import "google.golang.org/grpc/grpclog"

type Logger interface {
	grpclog.LoggerV2
	Lvl(lvl int) grpclog.LoggerV2
}

var _ Logger = &defaultLogger{}

func NewLogger(g grpclog.LoggerV2) Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &defaultLogger{g: g}
}

type defaultLogger struct{ g grpclog.LoggerV2 }

func (l *defaultLogger) Info(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Info(args...)
}
func (l *defaultLogger) Infoln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Info(args...)
}
func (l *defaultLogger) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Infof(format, args...)
}
func (l *defaultLogger) Warning(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Warning(args...)
}
func (l *defaultLogger) Warningln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Warning(args...)
}
func (l *defaultLogger) Warningf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Warningf(format, args...)
}
func (l *defaultLogger) Error(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Error(args...)
}
func (l *defaultLogger) Errorln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Error(args...)
}
func (l *defaultLogger) Errorf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Errorf(format, args...)
}
func (l *defaultLogger) Fatal(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Fatal(args...)
}
func (l *defaultLogger) Fatalln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Fatal(args...)
}
func (l *defaultLogger) Fatalf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.g.Fatalf(format, args...)
}
func (l *defaultLogger) V(lvl int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.g.V(lvl)
}
func (l *defaultLogger) Lvl(lvl int) grpclog.LoggerV2 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if l.g.V(lvl) {
		return l
	}
	return &discardLogger{}
}
