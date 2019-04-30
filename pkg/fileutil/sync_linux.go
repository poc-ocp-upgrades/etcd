package fileutil

import (
	"os"
	"syscall"
)

func Fsync(f *os.File) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.Sync()
}
func Fdatasync(f *os.File) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return syscall.Fdatasync(int(f.Fd()))
}
