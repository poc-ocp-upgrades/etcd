package fileutil

import "os"

func preallocExtend(f *os.File, sizeInBytes int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return preallocExtendTrunc(f, sizeInBytes)
}
func preallocFixed(f *os.File, sizeInBytes int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
