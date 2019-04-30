package osutil

import (
	"os"
	"go.uber.org/zap"
)

type InterruptHandler func()

func RegisterInterruptHandler(h InterruptHandler) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func HandleInterrupts(*zap.Logger) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func Exit(code int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	os.Exit(code)
}
