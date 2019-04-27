package command

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/spf13/cobra"
)

var memberPeerURLs string

func NewMemberCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mc := &cobra.Command{Use: "member <subcommand>", Short: "Membership related commands"}
	mc.AddCommand(NewMemberAddCommand())
	mc.AddCommand(NewMemberRemoveCommand())
	mc.AddCommand(NewMemberUpdateCommand())
	mc.AddCommand(NewMemberListCommand())
	return mc
}
func NewMemberAddCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := &cobra.Command{Use: "add <memberName> [options]", Short: "Adds a member into the cluster", Run: memberAddCommandFunc}
	cc.Flags().StringVar(&memberPeerURLs, "peer-urls", "", "comma separated peer URLs for the new member.")
	return cc
}
func NewMemberRemoveCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := &cobra.Command{Use: "remove <memberID>", Short: "Removes a member from the cluster", Run: memberRemoveCommandFunc}
	return cc
}
func NewMemberUpdateCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := &cobra.Command{Use: "update <memberID> [options]", Short: "Updates a member in the cluster", Run: memberUpdateCommandFunc}
	cc.Flags().StringVar(&memberPeerURLs, "peer-urls", "", "comma separated peer URLs for the updated member.")
	return cc
}
func NewMemberListCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := &cobra.Command{Use: "list", Short: "Lists all members in the cluster", Long: `When --write-out is set to simple, this command prints out comma-separated member lists for each endpoint.
The items in the lists are ID, Status, Name, Peer Addrs, Client Addrs.
`, Run: memberListCommandFunc}
	return cc
}
func memberAddCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("member name not provided."))
	}
	newMemberName := args[0]
	if len(memberPeerURLs) == 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("member peer urls not provided."))
	}
	urls := strings.Split(memberPeerURLs, ",")
	ctx, cancel := commandCtx(cmd)
	cli := mustClientFromCmd(cmd)
	resp, err := cli.MemberAdd(ctx, urls)
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	newID := resp.Member.ID
	display.MemberAdd(*resp)
	if _, ok := (display).(*simplePrinter); ok {
		ctx, cancel = commandCtx(cmd)
		listResp, err := cli.MemberList(ctx)
		for {
			if err != nil {
				ExitWithError(ExitError, err)
			}
			if listResp.Header.MemberId == resp.Header.MemberId {
				break
			}
			gresp, gerr := cli.Get(ctx, "_")
			if gerr != nil {
				ExitWithError(ExitError, err)
			}
			resp.Header.MemberId = gresp.Header.MemberId
			listResp, err = cli.MemberList(ctx)
		}
		cancel()
		conf := []string{}
		for _, memb := range listResp.Members {
			for _, u := range memb.PeerURLs {
				n := memb.Name
				if memb.ID == newID {
					n = newMemberName
				}
				conf = append(conf, fmt.Sprintf("%s=%s", n, u))
			}
		}
		fmt.Print("\n")
		fmt.Printf("ETCD_NAME=%q\n", newMemberName)
		fmt.Printf("ETCD_INITIAL_CLUSTER=%q\n", strings.Join(conf, ","))
		fmt.Printf("ETCD_INITIAL_ADVERTISE_PEER_URLS=%q\n", memberPeerURLs)
		fmt.Printf("ETCD_INITIAL_CLUSTER_STATE=\"existing\"\n")
	}
}
func memberRemoveCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("member ID is not provided"))
	}
	id, err := strconv.ParseUint(args[0], 16, 64)
	if err != nil {
		ExitWithError(ExitBadArgs, fmt.Errorf("bad member ID arg (%v), expecting ID in Hex", err))
	}
	ctx, cancel := commandCtx(cmd)
	resp, err := mustClientFromCmd(cmd).MemberRemove(ctx, id)
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.MemberRemove(id, *resp)
}
func memberUpdateCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("member ID is not provided"))
	}
	id, err := strconv.ParseUint(args[0], 16, 64)
	if err != nil {
		ExitWithError(ExitBadArgs, fmt.Errorf("bad member ID arg (%v), expecting ID in Hex", err))
	}
	if len(memberPeerURLs) == 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("member peer urls not provided."))
	}
	urls := strings.Split(memberPeerURLs, ",")
	ctx, cancel := commandCtx(cmd)
	resp, err := mustClientFromCmd(cmd).MemberUpdate(ctx, id, urls)
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.MemberUpdate(id, *resp)
}
func memberListCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := commandCtx(cmd)
	resp, err := mustClientFromCmd(cmd).MemberList(ctx)
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.MemberList(*resp)
}
