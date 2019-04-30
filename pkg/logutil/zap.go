package logutil

import (
	"sort"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DefaultZapLoggerConfig = zap.Config{Level: zap.NewAtomicLevelAt(zap.InfoLevel), Development: false, Sampling: &zap.SamplingConfig{Initial: 100, Thereafter: 100}, Encoding: "json", EncoderConfig: zapcore.EncoderConfig{TimeKey: "ts", LevelKey: "level", NameKey: "logger", CallerKey: "caller", MessageKey: "msg", StacktraceKey: "stacktrace", LineEnding: zapcore.DefaultLineEnding, EncodeLevel: zapcore.LowercaseLevelEncoder, EncodeTime: zapcore.ISO8601TimeEncoder, EncodeDuration: zapcore.StringDurationEncoder, EncodeCaller: zapcore.ShortCallerEncoder}, OutputPaths: []string{"stderr"}, ErrorOutputPaths: []string{"stderr"}}

func AddOutputPaths(cfg zap.Config, outputPaths, errorOutputPaths []string) zap.Config {
	_logClusterCodePath()
	defer _logClusterCodePath()
	outputs := make(map[string]struct{})
	for _, v := range cfg.OutputPaths {
		outputs[v] = struct{}{}
	}
	for _, v := range outputPaths {
		outputs[v] = struct{}{}
	}
	outputSlice := make([]string, 0)
	if _, ok := outputs["/dev/null"]; ok {
		outputSlice = []string{"/dev/null"}
	} else {
		for k := range outputs {
			outputSlice = append(outputSlice, k)
		}
	}
	cfg.OutputPaths = outputSlice
	sort.Strings(cfg.OutputPaths)
	errOutputs := make(map[string]struct{})
	for _, v := range cfg.ErrorOutputPaths {
		errOutputs[v] = struct{}{}
	}
	for _, v := range errorOutputPaths {
		errOutputs[v] = struct{}{}
	}
	errOutputSlice := make([]string, 0)
	if _, ok := errOutputs["/dev/null"]; ok {
		errOutputSlice = []string{"/dev/null"}
	} else {
		for k := range errOutputs {
			errOutputSlice = append(errOutputSlice, k)
		}
	}
	cfg.ErrorOutputPaths = errOutputSlice
	sort.Strings(cfg.ErrorOutputPaths)
	return cfg
}
