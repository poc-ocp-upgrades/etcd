package embed

import (
	"fmt"
	"os"
	"go.etcd.io/etcd/pkg/logutil"
	"go.uber.org/zap/zapcore"
)

func getJournalWriteSyncer() (zapcore.WriteSyncer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	jw, err := logutil.NewJournalWriter(os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("can't find journal (%v)", err)
	}
	return zapcore.AddSync(jw), nil
}
