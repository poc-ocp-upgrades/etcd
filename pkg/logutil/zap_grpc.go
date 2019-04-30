package logutil

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
)

func NewGRPCLoggerV2(lcfg zap.Config) (grpclog.LoggerV2, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg, err := lcfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}
	return &zapGRPCLogger{lg: lg, sugar: lg.Sugar()}, nil
}
func NewGRPCLoggerV2FromZapCore(cr zapcore.Core, syncer zapcore.WriteSyncer) grpclog.LoggerV2 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg := zap.New(cr, zap.AddCaller(), zap.AddCallerSkip(1), zap.ErrorOutput(syncer))
	return &zapGRPCLogger{lg: lg, sugar: lg.Sugar()}
}

type zapGRPCLogger struct {
	lg	*zap.Logger
	sugar	*zap.SugaredLogger
}

func (zl *zapGRPCLogger) Info(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !zl.lg.Core().Enabled(zapcore.DebugLevel) {
		return
	}
	zl.sugar.Info(args...)
}
func (zl *zapGRPCLogger) Infoln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !zl.lg.Core().Enabled(zapcore.DebugLevel) {
		return
	}
	zl.sugar.Info(args...)
}
func (zl *zapGRPCLogger) Infof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !zl.lg.Core().Enabled(zapcore.DebugLevel) {
		return
	}
	zl.sugar.Infof(format, args...)
}
func (zl *zapGRPCLogger) Warning(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Warn(args...)
}
func (zl *zapGRPCLogger) Warningln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Warn(args...)
}
func (zl *zapGRPCLogger) Warningf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Warnf(format, args...)
}
func (zl *zapGRPCLogger) Error(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Error(args...)
}
func (zl *zapGRPCLogger) Errorln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Error(args...)
}
func (zl *zapGRPCLogger) Errorf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Errorf(format, args...)
}
func (zl *zapGRPCLogger) Fatal(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Fatal(args...)
}
func (zl *zapGRPCLogger) Fatalln(args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Fatal(args...)
}
func (zl *zapGRPCLogger) Fatalf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	zl.sugar.Fatalf(format, args...)
}
func (zl *zapGRPCLogger) V(l int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if l <= 0 {
		return !zl.lg.Core().Enabled(zapcore.DebugLevel)
	}
	return true
}
