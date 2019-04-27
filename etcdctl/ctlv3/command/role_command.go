package command

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/cobra"
)

var (
	rolePermPrefix	bool
	rolePermFromKey	bool
)

func NewRoleCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ac := &cobra.Command{Use: "role <subcommand>", Short: "Role related commands"}
	ac.AddCommand(newRoleAddCommand())
	ac.AddCommand(newRoleDeleteCommand())
	ac.AddCommand(newRoleGetCommand())
	ac.AddCommand(newRoleListCommand())
	ac.AddCommand(newRoleGrantPermissionCommand())
	ac.AddCommand(newRoleRevokePermissionCommand())
	return ac
}
func newRoleAddCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cobra.Command{Use: "add <role name>", Short: "Adds a new role", Run: roleAddCommandFunc}
}
func newRoleDeleteCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cobra.Command{Use: "delete <role name>", Short: "Deletes a role", Run: roleDeleteCommandFunc}
}
func newRoleGetCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cobra.Command{Use: "get <role name>", Short: "Gets detailed information of a role", Run: roleGetCommandFunc}
}
func newRoleListCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cobra.Command{Use: "list", Short: "Lists all roles", Run: roleListCommandFunc}
}
func newRoleGrantPermissionCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "grant-permission [options] <role name> <permission type> <key> [endkey]", Short: "Grants a key to a role", Run: roleGrantPermissionCommandFunc}
	cmd.Flags().BoolVar(&rolePermPrefix, "prefix", false, "grant a prefix permission")
	cmd.Flags().BoolVar(&rolePermFromKey, "from-key", false, "grant a permission of keys that are greater than or equal to the given key using byte compare")
	return cmd
}
func newRoleRevokePermissionCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "revoke-permission <role name> <key> [endkey]", Short: "Revokes a key from a role", Run: roleRevokePermissionCommandFunc}
	cmd.Flags().BoolVar(&rolePermPrefix, "prefix", false, "revoke a prefix permission")
	cmd.Flags().BoolVar(&rolePermFromKey, "from-key", false, "revoke a permission of keys that are greater than or equal to the given key using byte compare")
	return cmd
}
func roleAddCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("role add command requires role name as its argument."))
	}
	resp, err := mustClientFromCmd(cmd).Auth.RoleAdd(context.TODO(), args[0])
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.RoleAdd(args[0], *resp)
}
func roleDeleteCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("role delete command requires role name as its argument."))
	}
	resp, err := mustClientFromCmd(cmd).Auth.RoleDelete(context.TODO(), args[0])
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.RoleDelete(args[0], *resp)
}
func roleGetCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("role get command requires role name as its argument."))
	}
	name := args[0]
	resp, err := mustClientFromCmd(cmd).Auth.RoleGet(context.TODO(), name)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.RoleGet(name, *resp)
}
func roleListCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("role list command requires no arguments."))
	}
	resp, err := mustClientFromCmd(cmd).Auth.RoleList(context.TODO())
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.RoleList(*resp)
}
func roleGrantPermissionCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) < 3 {
		ExitWithError(ExitBadArgs, fmt.Errorf("role grant command requires role name, permission type, and key [endkey] as its argument."))
	}
	perm, err := clientv3.StrToPermissionType(args[1])
	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}
	key, rangeEnd := permRange(args[2:])
	resp, err := mustClientFromCmd(cmd).Auth.RoleGrantPermission(context.TODO(), args[0], key, rangeEnd, perm)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.RoleGrantPermission(args[0], *resp)
}
func roleRevokePermissionCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) < 2 {
		ExitWithError(ExitBadArgs, fmt.Errorf("role revoke-permission command requires role name and key [endkey] as its argument."))
	}
	key, rangeEnd := permRange(args[1:])
	resp, err := mustClientFromCmd(cmd).Auth.RoleRevokePermission(context.TODO(), args[0], key, rangeEnd)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.RoleRevokePermission(args[0], args[1], rangeEnd, *resp)
}
func permRange(args []string) (string, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := args[0]
	var rangeEnd string
	if len(key) == 0 {
		if rolePermPrefix && rolePermFromKey {
			ExitWithError(ExitBadArgs, fmt.Errorf("--from-key and --prefix flags are mutually exclusive"))
		}
		key = "\x00"
		if rolePermPrefix || rolePermFromKey {
			rangeEnd = "\x00"
		}
	} else {
		var err error
		rangeEnd, err = rangeEndFromPermFlags(args[0:])
		if err != nil {
			ExitWithError(ExitBadArgs, err)
		}
	}
	return key, rangeEnd
}
func rangeEndFromPermFlags(args []string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 1 {
		if rolePermPrefix {
			if rolePermFromKey {
				return "", fmt.Errorf("--from-key and --prefix flags are mutually exclusive")
			}
			return clientv3.GetPrefixRangeEnd(args[0]), nil
		}
		if rolePermFromKey {
			return "\x00", nil
		}
		return "", nil
	}
	if rolePermPrefix {
		return "", fmt.Errorf("unexpected endkey argument with --prefix flag")
	}
	if rolePermFromKey {
		return "", fmt.Errorf("unexpected endkey argument with --from-key flag")
	}
	return args[1], nil
}
