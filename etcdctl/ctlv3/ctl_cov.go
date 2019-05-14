package ctlv3

import (
	"github.com/coreos/etcd/etcdctl/ctlv3/command"
	"os"
	"strings"
)

func Start() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd.SetArgs(strings.Split(os.Getenv("ETCDCTL_ARGS"), "\xe7\xcd")[1:])
	os.Unsetenv("ETCDCTL_ARGS")
	if err := rootCmd.Execute(); err != nil {
		command.ExitWithError(command.ExitError, err)
	}
}
