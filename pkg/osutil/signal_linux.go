package osutil

import (
	"syscall"
	"unsafe"
)

func dflSignal(sig syscall.Signal) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var sigactBuf [32]uint64
	ptr := unsafe.Pointer(&sigactBuf)
	syscall.Syscall6(uintptr(syscall.SYS_RT_SIGACTION), uintptr(sig), uintptr(ptr), 0, 8, 0, 0)
}
