package netutil

import (
	"fmt"
	"runtime"
)

func GetDefaultHost() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "", fmt.Errorf("default host not supported on %s_%s", runtime.GOOS, runtime.GOARCH)
}
func GetDefaultInterfaces() (map[string]uint8, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, fmt.Errorf("default host not supported on %s_%s", runtime.GOOS, runtime.GOARCH)
}
