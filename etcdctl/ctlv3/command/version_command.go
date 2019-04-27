package command

import (
	"fmt"
	"github.com/coreos/etcd/version"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cobra.Command{Use: "version", Short: "Prints the version of etcdctl", Run: versionCommandFunc}
}
func versionCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Println("etcdctl version:", version.Version)
	fmt.Println("API version:", version.APIVersion)
}
