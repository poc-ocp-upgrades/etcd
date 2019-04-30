package logutil

import (
	"github.com/coreos/pkg/capnslog"
	"google.golang.org/grpc/grpclog"
)

var _ Logger = &packageLogger{}

func NewPackageLogger(repo, pkg string) Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &packageLogger{p: capnslog.NewPackageLogger(repo, pkg)}
}

type packageLogger struct{ p *capnslog.PackageLogger }

func (l *packageLogger) Info(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Info(args...)
}
func (l *packageLogger) Infoln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Info(args...)
}
func (l *packageLogger) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Infof(format, args...)
}
func (l *packageLogger) Warning(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Warning(args...)
}
func (l *packageLogger) Warningln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Warning(args...)
}
func (l *packageLogger) Warningf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Warningf(format, args...)
}
func (l *packageLogger) Error(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Error(args...)
}
func (l *packageLogger) Errorln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Error(args...)
}
func (l *packageLogger) Errorf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Errorf(format, args...)
}
func (l *packageLogger) Fatal(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Fatal(args...)
}
func (l *packageLogger) Fatalln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Fatal(args...)
}
func (l *packageLogger) Fatalf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.p.Fatalf(format, args...)
}
func (l *packageLogger) V(lvl int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return l.p.LevelAt(capnslog.LogLevel(lvl))
}
func (l *packageLogger) Lvl(lvl int) grpclog.LoggerV2 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if l.p.LevelAt(capnslog.LogLevel(lvl)) {
		return l
	}
	return &discardLogger{}
}
