package runtime

import (
	"fmt"
	"runtime"
)

func FDLimit() (uint64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0, fmt.Errorf("cannot get FDLimit on %s", runtime.GOOS)
}
func FDUsage() (uint64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0, fmt.Errorf("cannot get FDUsage on %s", runtime.GOOS)
}
