package command

import (
	"fmt"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/spf13/cobra"
)

func NewAuthCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ac := &cobra.Command{Use: "auth <enable or disable>", Short: "Enable or disable authentication"}
	ac.AddCommand(newAuthEnableCommand())
	ac.AddCommand(newAuthDisableCommand())
	return ac
}
func newAuthEnableCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cobra.Command{Use: "enable", Short: "Enables authentication", Run: authEnableCommandFunc}
}
func authEnableCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("auth enable command does not accept any arguments."))
	}
	ctx, cancel := commandCtx(cmd)
	cli := mustClientFromCmd(cmd)
	var err error
	for err == nil {
		if _, err = cli.AuthEnable(ctx); err == nil {
			break
		}
		if err == rpctypes.ErrRootRoleNotExist {
			if _, err = cli.RoleAdd(ctx, "root"); err != nil {
				break
			}
			if _, err = cli.UserGrantRole(ctx, "root", "root"); err != nil {
				break
			}
		}
	}
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	fmt.Println("Authentication Enabled")
}
func newAuthDisableCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cobra.Command{Use: "disable", Short: "Disables authentication", Run: authDisableCommandFunc}
}
func authDisableCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("auth disable command does not accept any arguments."))
	}
	ctx, cancel := commandCtx(cmd)
	_, err := mustClientFromCmd(cmd).Auth.AuthDisable(ctx)
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	fmt.Println("Authentication Disabled")
}
