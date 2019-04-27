package osutil

import (
	"os"
	"strings"
	"github.com/coreos/pkg/capnslog"
)

var (
	plog		= capnslog.NewPackageLogger("github.com/coreos/etcd", "pkg/osutil")
	setDflSignal	= dflSignal
)

func Unsetenv(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	envs := os.Environ()
	os.Clearenv()
	for _, e := range envs {
		strs := strings.SplitN(e, "=", 2)
		if strs[0] == key {
			continue
		}
		if err := os.Setenv(strs[0], strs[1]); err != nil {
			return err
		}
	}
	return nil
}
