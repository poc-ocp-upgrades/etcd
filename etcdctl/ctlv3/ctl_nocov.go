package ctlv3

import "github.com/coreos/etcd/etcdctl/ctlv3/command"

func Start() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd.SetUsageFunc(usageFunc)
	rootCmd.SetHelpTemplate(`{{.UsageString}}`)
	if err := rootCmd.Execute(); err != nil {
		command.ExitWithError(command.ExitError, err)
	}
}
