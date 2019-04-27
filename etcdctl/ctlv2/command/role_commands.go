package command

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/pkg/pathutil"
	"github.com/urfave/cli"
)

func NewRoleCommands() cli.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cli.Command{Name: "role", Usage: "role add, grant and revoke subcommands", Subcommands: []cli.Command{{Name: "add", Usage: "add a new role for the etcd cluster", ArgsUsage: "<role> ", Action: actionRoleAdd}, {Name: "get", Usage: "get details for a role", ArgsUsage: "<role>", Action: actionRoleGet}, {Name: "list", Usage: "list all roles", ArgsUsage: " ", Action: actionRoleList}, {Name: "remove", Usage: "remove a role from the etcd cluster", ArgsUsage: "<role>", Action: actionRoleRemove}, {Name: "grant", Usage: "grant path matches to an etcd role", ArgsUsage: "<role>", Flags: []cli.Flag{cli.StringFlag{Name: "path", Value: "", Usage: "Path granted for the role to access"}, cli.BoolFlag{Name: "read", Usage: "Grant read-only access"}, cli.BoolFlag{Name: "write", Usage: "Grant write-only access"}, cli.BoolFlag{Name: "readwrite, rw", Usage: "Grant read-write access"}}, Action: actionRoleGrant}, {Name: "revoke", Usage: "revoke path matches for an etcd role", ArgsUsage: "<role>", Flags: []cli.Flag{cli.StringFlag{Name: "path", Value: "", Usage: "Path revoked for the role to access"}, cli.BoolFlag{Name: "read", Usage: "Revoke read access"}, cli.BoolFlag{Name: "write", Usage: "Revoke write access"}, cli.BoolFlag{Name: "readwrite, rw", Usage: "Revoke read-write access"}}, Action: actionRoleRevoke}}}
}
func mustNewAuthRoleAPI(c *cli.Context) client.AuthRoleAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hc := mustNewClient(c)
	if c.GlobalBool("debug") {
		fmt.Fprintf(os.Stderr, "Cluster-Endpoints: %s\n", strings.Join(hc.Endpoints(), ", "))
	}
	return client.NewAuthRoleAPI(hc)
}
func actionRoleList(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(c.Args()) != 0 {
		fmt.Fprintln(os.Stderr, "No arguments accepted")
		os.Exit(1)
	}
	r := mustNewAuthRoleAPI(c)
	ctx, cancel := contextWithTotalTimeout(c)
	roles, err := r.ListRoles(ctx)
	cancel()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	for _, role := range roles {
		fmt.Printf("%s\n", role)
	}
	return nil
}
func actionRoleAdd(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api, role := mustRoleAPIAndName(c)
	ctx, cancel := contextWithTotalTimeout(c)
	defer cancel()
	currentRole, _ := api.GetRole(ctx, role)
	if currentRole != nil {
		fmt.Fprintf(os.Stderr, "Role %s already exists\n", role)
		os.Exit(1)
	}
	err := api.AddRole(ctx, role)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("Role %s created\n", role)
	return nil
}
func actionRoleRemove(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api, role := mustRoleAPIAndName(c)
	ctx, cancel := contextWithTotalTimeout(c)
	err := api.RemoveRole(ctx, role)
	cancel()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("Role %s removed\n", role)
	return nil
}
func actionRoleGrant(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	roleGrantRevoke(c, true)
	return nil
}
func actionRoleRevoke(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	roleGrantRevoke(c, false)
	return nil
}
func roleGrantRevoke(c *cli.Context, grant bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	path := c.String("path")
	if path == "" {
		fmt.Fprintln(os.Stderr, "No path specified; please use `--path`")
		os.Exit(1)
	}
	if pathutil.CanonicalURLPath(path) != path {
		fmt.Fprintf(os.Stderr, "Not canonical path; please use `--path=%s`\n", pathutil.CanonicalURLPath(path))
		os.Exit(1)
	}
	read := c.Bool("read")
	write := c.Bool("write")
	rw := c.Bool("readwrite")
	permcount := 0
	for _, v := range []bool{read, write, rw} {
		if v {
			permcount++
		}
	}
	if permcount != 1 {
		fmt.Fprintln(os.Stderr, "Please specify exactly one of --read, --write or --readwrite")
		os.Exit(1)
	}
	var permType client.PermissionType
	switch {
	case read:
		permType = client.ReadPermission
	case write:
		permType = client.WritePermission
	case rw:
		permType = client.ReadWritePermission
	}
	api, role := mustRoleAPIAndName(c)
	ctx, cancel := contextWithTotalTimeout(c)
	defer cancel()
	currentRole, err := api.GetRole(ctx, role)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	var newRole *client.Role
	if grant {
		newRole, err = api.GrantRoleKV(ctx, role, []string{path}, permType)
	} else {
		newRole, err = api.RevokeRoleKV(ctx, role, []string{path}, permType)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if reflect.DeepEqual(newRole, currentRole) {
		if grant {
			fmt.Printf("Role unchanged; already granted")
		} else {
			fmt.Printf("Role unchanged; already revoked")
		}
	}
	fmt.Printf("Role %s updated\n", role)
}
func actionRoleGet(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api, rolename := mustRoleAPIAndName(c)
	ctx, cancel := contextWithTotalTimeout(c)
	role, err := api.GetRole(ctx, rolename)
	cancel()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("Role: %s\n", role.Role)
	fmt.Printf("KV Read:\n")
	for _, v := range role.Permissions.KV.Read {
		fmt.Printf("\t%s\n", v)
	}
	fmt.Printf("KV Write:\n")
	for _, v := range role.Permissions.KV.Write {
		fmt.Printf("\t%s\n", v)
	}
	return nil
}
func mustRoleAPIAndName(c *cli.Context) (client.AuthRoleAPI, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := c.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Please provide a role name")
		os.Exit(1)
	}
	name := args[0]
	api := mustNewAuthRoleAPI(c)
	return api, name
}
