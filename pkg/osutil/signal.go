package osutil

import "syscall"

func dflSignal(sig syscall.Signal) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
