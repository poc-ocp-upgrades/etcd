package e2e

import (
	"github.com/coreos/etcd/pkg/expect"
	"os"
)

const noOutputLineCount = 0

func spawnCmd(args []string) (*expect.ExpectProcess, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if args[0] == ctlBinPath+"3" {
		env := append(os.Environ(), "ETCDCTL_API=3")
		return expect.NewExpectWithEnv(ctlBinPath, args[1:], env)
	}
	return expect.NewExpect(args[0], args[1:]...)
}
