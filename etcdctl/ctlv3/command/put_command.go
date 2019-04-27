package command

import (
	"fmt"
	"os"
	"strconv"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/cobra"
)

var (
	leaseStr	string
	putPrevKV	bool
	putIgnoreVal	bool
	putIgnoreLease	bool
)

func NewPutCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "put [options] <key> <value> (<value> can also be given from stdin)", Short: "Puts the given key into the store", Long: `
Puts the given key into the store.

When <value> begins with '-', <value> is interpreted as a flag.
Insert '--' for workaround:

$ put <key> -- <value>
$ put -- <key> <value>

If <value> isn't given as a command line argument and '--ignore-value' is not specified,
this command tries to read the value from standard input.

If <lease> isn't given as a command line argument and '--ignore-lease' is not specified,
this command tries to read the value from standard input.

For example,
$ cat file | put <key>
will store the content of the file to <key>.
`, Run: putCommandFunc}
	cmd.Flags().StringVar(&leaseStr, "lease", "0", "lease ID (in hexadecimal) to attach to the key")
	cmd.Flags().BoolVar(&putPrevKV, "prev-kv", false, "return the previous key-value pair before modification")
	cmd.Flags().BoolVar(&putIgnoreVal, "ignore-value", false, "updates the key using its current value")
	cmd.Flags().BoolVar(&putIgnoreLease, "ignore-lease", false, "updates the key using its current lease")
	return cmd
}
func putCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, value, opts := getPutOp(cmd, args)
	ctx, cancel := commandCtx(cmd)
	resp, err := mustClientFromCmd(cmd).Put(ctx, key, value, opts...)
	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.Put(*resp)
}
func getPutOp(cmd *cobra.Command, args []string) (string, string, []clientv3.OpOption) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("put command needs 1 argument and input from stdin or 2 arguments."))
	}
	key := args[0]
	if putIgnoreVal && len(args) > 1 {
		ExitWithError(ExitBadArgs, fmt.Errorf("put command needs only 1 argument when 'ignore-value' is set."))
	}
	var value string
	var err error
	if !putIgnoreVal {
		value, err = argOrStdin(args, os.Stdin, 1)
		if err != nil {
			ExitWithError(ExitBadArgs, fmt.Errorf("put command needs 1 argument and input from stdin or 2 arguments."))
		}
	}
	id, err := strconv.ParseInt(leaseStr, 16, 64)
	if err != nil {
		ExitWithError(ExitBadArgs, fmt.Errorf("bad lease ID (%v), expecting ID in Hex", err))
	}
	opts := []clientv3.OpOption{}
	if id != 0 {
		opts = append(opts, clientv3.WithLease(clientv3.LeaseID(id)))
	}
	if putPrevKV {
		opts = append(opts, clientv3.WithPrevKV())
	}
	if putIgnoreVal {
		opts = append(opts, clientv3.WithIgnoreValue())
	}
	if putIgnoreLease {
		opts = append(opts, clientv3.WithIgnoreLease())
	}
	return key, value, opts
}
