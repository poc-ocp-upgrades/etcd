package embed

import (
	"os"
	"go.uber.org/zap/zapcore"
)

func getJournalWriteSyncer() (zapcore.WriteSyncer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return zapcore.AddSync(os.Stderr), nil
}
