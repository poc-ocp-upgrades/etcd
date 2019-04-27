package command

import (
	"fmt"
	"strconv"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/cobra"
)

var compactPhysical bool

func NewCompactionCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "compaction [options] <revision>", Short: "Compacts the event history in etcd", Run: compactionCommandFunc}
	cmd.Flags().BoolVar(&compactPhysical, "physical", false, "'true' to wait for compaction to physically remove all old revisions")
	return cmd
}
func compactionCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("compaction command needs 1 argument."))
	}
	rev, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		ExitWithError(ExitError, err)
	}
	var opts []clientv3.CompactOption
	if compactPhysical {
		opts = append(opts, clientv3.WithCompactPhysical())
	}
	c := mustClientFromCmd(cmd)
	ctx, cancel := commandCtx(cmd)
	_, cerr := c.Compact(ctx, rev, opts...)
	cancel()
	if cerr != nil {
		ExitWithError(ExitError, cerr)
	}
	fmt.Println("compacted revision", rev)
}
