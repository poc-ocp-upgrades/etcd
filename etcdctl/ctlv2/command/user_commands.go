package command

import (
	"fmt"
	"os"
	"strings"
	"github.com/bgentry/speakeasy"
	"github.com/urfave/cli"
	"go.etcd.io/etcd/client"
)

func NewUserCommands() cli.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cli.Command{Name: "user", Usage: "user add, grant and revoke subcommands", Subcommands: []cli.Command{{Name: "add", Usage: "add a new user for the etcd cluster", ArgsUsage: "<user>", Action: actionUserAdd}, {Name: "get", Usage: "get details for a user", ArgsUsage: "<user>", Action: actionUserGet}, {Name: "list", Usage: "list all current users", ArgsUsage: "<user>", Action: actionUserList}, {Name: "remove", Usage: "remove a user for the etcd cluster", ArgsUsage: "<user>", Action: actionUserRemove}, {Name: "grant", Usage: "grant roles to an etcd user", ArgsUsage: "<user>", Flags: []cli.Flag{cli.StringSliceFlag{Name: "roles", Value: new(cli.StringSlice), Usage: "List of roles to grant or revoke"}}, Action: actionUserGrant}, {Name: "revoke", Usage: "revoke roles for an etcd user", ArgsUsage: "<user>", Flags: []cli.Flag{cli.StringSliceFlag{Name: "roles", Value: new(cli.StringSlice), Usage: "List of roles to grant or revoke"}}, Action: actionUserRevoke}, {Name: "passwd", Usage: "change password for a user", ArgsUsage: "<user>", Action: actionUserPasswd}}}
}
func mustNewAuthUserAPI(c *cli.Context) client.AuthUserAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hc := mustNewClient(c)
	if c.GlobalBool("debug") {
		fmt.Fprintf(os.Stderr, "Cluster-Endpoints: %s\n", strings.Join(hc.Endpoints(), ", "))
	}
	return client.NewAuthUserAPI(hc)
}
func actionUserList(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(c.Args()) != 0 {
		fmt.Fprintln(os.Stderr, "No arguments accepted")
		os.Exit(1)
	}
	u := mustNewAuthUserAPI(c)
	ctx, cancel := contextWithTotalTimeout(c)
	users, err := u.ListUsers(ctx)
	cancel()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	for _, user := range users {
		fmt.Printf("%s\n", user)
	}
	return nil
}
func actionUserAdd(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api, userarg := mustUserAPIAndName(c)
	ctx, cancel := contextWithTotalTimeout(c)
	defer cancel()
	user, _, _ := getUsernamePassword("", userarg+":")
	_, pass, err := getUsernamePassword("New password: ", userarg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading password:", err)
		os.Exit(1)
	}
	err = api.AddUser(ctx, user, pass)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("User %s created\n", user)
	return nil
}
func actionUserRemove(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api, user := mustUserAPIAndName(c)
	ctx, cancel := contextWithTotalTimeout(c)
	err := api.RemoveUser(ctx, user)
	cancel()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("User %s removed\n", user)
	return nil
}
func actionUserPasswd(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api, user := mustUserAPIAndName(c)
	ctx, cancel := contextWithTotalTimeout(c)
	defer cancel()
	pass, err := speakeasy.Ask("New password: ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading password:", err)
		os.Exit(1)
	}
	_, err = api.ChangePassword(ctx, user, pass)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("Password updated\n")
	return nil
}
func actionUserGrant(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	userGrantRevoke(c, true)
	return nil
}
func actionUserRevoke(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	userGrantRevoke(c, false)
	return nil
}
func userGrantRevoke(c *cli.Context, grant bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	roles := c.StringSlice("roles")
	if len(roles) == 0 {
		fmt.Fprintln(os.Stderr, "No roles specified; please use `--roles`")
		os.Exit(1)
	}
	ctx, cancel := contextWithTotalTimeout(c)
	defer cancel()
	api, user := mustUserAPIAndName(c)
	var err error
	if grant {
		_, err = api.GrantUser(ctx, user, roles)
	} else {
		_, err = api.RevokeUser(ctx, user, roles)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("User %s updated\n", user)
}
func actionUserGet(c *cli.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api, username := mustUserAPIAndName(c)
	ctx, cancel := contextWithTotalTimeout(c)
	user, err := api.GetUser(ctx, username)
	cancel()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("User: %s\n", user.User)
	fmt.Printf("Roles: %s\n", strings.Join(user.Roles, " "))
	return nil
}
func mustUserAPIAndName(c *cli.Context) (client.AuthUserAPI, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := c.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Please provide a username")
		os.Exit(1)
	}
	api := mustNewAuthUserAPI(c)
	username := args[0]
	return api, username
}
