package command

import (
	"context"
	"fmt"
	"strconv"
	v3 "go.etcd.io/etcd/clientv3"
	"github.com/spf13/cobra"
)

func NewLeaseCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lc := &cobra.Command{Use: "lease <subcommand>", Short: "Lease related commands"}
	lc.AddCommand(NewLeaseGrantCommand())
	lc.AddCommand(NewLeaseRevokeCommand())
	lc.AddCommand(NewLeaseTimeToLiveCommand())
	lc.AddCommand(NewLeaseListCommand())
	lc.AddCommand(NewLeaseKeepAliveCommand())
	return lc
}
func NewLeaseGrantCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lc := &cobra.Command{Use: "grant <ttl>", Short: "Creates leases", Run: leaseGrantCommandFunc}
	return lc
}
func leaseGrantCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("lease grant command needs TTL argument"))
	}
	ttl, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		ExitWithError(ExitBadArgs, fmt.Errorf("bad TTL (%v)", err))
	}
	ctx, cancel := commandCtx(cmd)
	resp, err := mustClientFromCmd(cmd).Grant(ctx, ttl)
	cancel()
	if err != nil {
		ExitWithError(ExitError, fmt.Errorf("failed to grant lease (%v)", err))
	}
	display.Grant(*resp)
}
func NewLeaseRevokeCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lc := &cobra.Command{Use: "revoke <leaseID>", Short: "Revokes leases", Run: leaseRevokeCommandFunc}
	return lc
}
func leaseRevokeCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("lease revoke command needs 1 argument"))
	}
	id := leaseFromArgs(args[0])
	ctx, cancel := commandCtx(cmd)
	resp, err := mustClientFromCmd(cmd).Revoke(ctx, id)
	cancel()
	if err != nil {
		ExitWithError(ExitError, fmt.Errorf("failed to revoke lease (%v)", err))
	}
	display.Revoke(id, *resp)
}

var timeToLiveKeys bool

func NewLeaseTimeToLiveCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lc := &cobra.Command{Use: "timetolive <leaseID> [options]", Short: "Get lease information", Run: leaseTimeToLiveCommandFunc}
	lc.Flags().BoolVar(&timeToLiveKeys, "keys", false, "Get keys attached to this lease")
	return lc
}
func leaseTimeToLiveCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("lease timetolive command needs lease ID as argument"))
	}
	var opts []v3.LeaseOption
	if timeToLiveKeys {
		opts = append(opts, v3.WithAttachedKeys())
	}
	resp, rerr := mustClientFromCmd(cmd).TimeToLive(context.TODO(), leaseFromArgs(args[0]), opts...)
	if rerr != nil {
		ExitWithError(ExitBadConnection, rerr)
	}
	display.TimeToLive(*resp, timeToLiveKeys)
}
func NewLeaseListCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lc := &cobra.Command{Use: "list", Short: "List all active leases", Run: leaseListCommandFunc}
	return lc
}
func leaseListCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, rerr := mustClientFromCmd(cmd).Leases(context.TODO())
	if rerr != nil {
		ExitWithError(ExitBadConnection, rerr)
	}
	display.Leases(*resp)
}

var (
	leaseKeepAliveOnce bool
)

func NewLeaseKeepAliveCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lc := &cobra.Command{Use: "keep-alive [options] <leaseID>", Short: "Keeps leases alive (renew)", Run: leaseKeepAliveCommandFunc}
	lc.Flags().BoolVar(&leaseKeepAliveOnce, "once", false, "Resets the keep-alive time to its original value and exits immediately")
	return lc
}
func leaseKeepAliveCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("lease keep-alive command needs lease ID as argument"))
	}
	id := leaseFromArgs(args[0])
	if leaseKeepAliveOnce {
		respc, kerr := mustClientFromCmd(cmd).KeepAliveOnce(context.TODO(), id)
		if kerr != nil {
			ExitWithError(ExitBadConnection, kerr)
		}
		display.KeepAlive(*respc)
		return
	}
	respc, kerr := mustClientFromCmd(cmd).KeepAlive(context.TODO(), id)
	if kerr != nil {
		ExitWithError(ExitBadConnection, kerr)
	}
	for resp := range respc {
		display.KeepAlive(*resp)
	}
	if _, ok := (display).(*simplePrinter); ok {
		fmt.Printf("lease %016x expired or revoked.\n", id)
	}
}
func leaseFromArgs(arg string) v3.LeaseID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	id, err := strconv.ParseInt(arg, 16, 64)
	if err != nil {
		ExitWithError(ExitBadArgs, fmt.Errorf("bad lease ID arg (%v), expecting ID in Hex", err))
	}
	return v3.LeaseID(id)
}
