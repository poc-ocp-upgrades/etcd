package command

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"go.etcd.io/etcd/clientv3"
)

func NewMoveLeaderCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "move-leader <transferee-member-id>", Short: "Transfers leadership to another etcd cluster member.", Run: transferLeadershipCommandFunc}
	return cmd
}
func transferLeadershipCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("move-leader command needs 1 argument"))
	}
	target, err := strconv.ParseUint(args[0], 16, 64)
	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}
	c := mustClientFromCmd(cmd)
	eps := c.Endpoints()
	c.Close()
	ctx, cancel := commandCtx(cmd)
	var leaderCli *clientv3.Client
	var leaderID uint64
	for _, ep := range eps {
		cfg := clientConfigFromCmd(cmd)
		cfg.endpoints = []string{ep}
		cli := cfg.mustClient()
		resp, serr := cli.Status(ctx, ep)
		if serr != nil {
			ExitWithError(ExitError, serr)
		}
		if resp.Header.GetMemberId() == resp.Leader {
			leaderCli = cli
			leaderID = resp.Leader
			break
		}
		cli.Close()
	}
	if leaderCli == nil {
		ExitWithError(ExitBadArgs, fmt.Errorf("no leader endpoint given at %v", eps))
	}
	var resp *clientv3.MoveLeaderResponse
	resp, err = leaderCli.MoveLeader(ctx, target)
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.MoveLeader(leaderID, target, *resp)
}
