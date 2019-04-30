package logutil

import (
	"errors"
	"go.etcd.io/etcd/raft"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewRaftLogger(lcfg *zap.Config) (raft.Logger, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if lcfg == nil {
		return nil, errors.New("nil zap.Config")
	}
	lg, err := lcfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	return &zapRaftLogger{lg: lg, sugar: lg.Sugar()}, nil
}
func NewRaftLoggerFromZapCore(cr zapcore.Core, syncer zapcore.WriteSyncer) raft.Logger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg := zap.New(cr, zap.AddCaller(), zap.AddCallerSkip(1), zap.ErrorOutput(syncer))
	return &zapRaftLogger{lg: lg, sugar: lg.Sugar()}
}

type zapRaftLogger struct {
	lg	*zap.Logger
	sugar	*zap.SugaredLogger
}

func (zl *zapRaftLogger) Debug(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Debug(args...)
}
func (zl *zapRaftLogger) Debugf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Debugf(format, args...)
}
func (zl *zapRaftLogger) Error(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Error(args...)
}
func (zl *zapRaftLogger) Errorf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Errorf(format, args...)
}
func (zl *zapRaftLogger) Info(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Info(args...)
}
func (zl *zapRaftLogger) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Infof(format, args...)
}
func (zl *zapRaftLogger) Warning(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Warn(args...)
}
func (zl *zapRaftLogger) Warningf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Warnf(format, args...)
}
func (zl *zapRaftLogger) Fatal(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Fatal(args...)
}
func (zl *zapRaftLogger) Fatalf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Fatalf(format, args...)
}
func (zl *zapRaftLogger) Panic(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Panic(args...)
}
func (zl *zapRaftLogger) Panicf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Panicf(format, args...)
}
