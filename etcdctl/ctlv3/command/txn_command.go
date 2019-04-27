package command

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/coreos/etcd/clientv3"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/spf13/cobra"
)

var (
	txnInteractive bool
)

func NewTxnCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "txn [options]", Short: "Txn processes all the requests in one transaction", Run: txnCommandFunc}
	cmd.Flags().BoolVarP(&txnInteractive, "interactive", "i", false, "Input transaction in interactive mode")
	return cmd
}
func txnCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("txn command does not accept argument."))
	}
	reader := bufio.NewReader(os.Stdin)
	txn := mustClientFromCmd(cmd).Txn(context.Background())
	promptInteractive("compares:")
	txn.If(readCompares(reader)...)
	promptInteractive("success requests (get, put, del):")
	txn.Then(readOps(reader)...)
	promptInteractive("failure requests (get, put, del):")
	txn.Else(readOps(reader)...)
	resp, err := txn.Commit()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	display.Txn(*resp)
}
func promptInteractive(s string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if txnInteractive {
		fmt.Println(s)
	}
}
func readCompares(r *bufio.Reader) (cmps []clientv3.Cmp) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			ExitWithError(ExitInvalidInput, err)
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}
		cmp, err := parseCompare(line)
		if err != nil {
			ExitWithError(ExitInvalidInput, err)
		}
		cmps = append(cmps, *cmp)
	}
	return cmps
}
func readOps(r *bufio.Reader) (ops []clientv3.Op) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			ExitWithError(ExitInvalidInput, err)
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}
		op, err := parseRequestUnion(line)
		if err != nil {
			ExitWithError(ExitInvalidInput, err)
		}
		ops = append(ops, *op)
	}
	return ops
}
func parseRequestUnion(line string) (*clientv3.Op, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := argify(line)
	if len(args) < 2 {
		return nil, fmt.Errorf("invalid txn compare request: %s", line)
	}
	opc := make(chan clientv3.Op, 1)
	put := NewPutCommand()
	put.Run = func(cmd *cobra.Command, args []string) {
		key, value, opts := getPutOp(cmd, args)
		opc <- clientv3.OpPut(key, value, opts...)
	}
	get := NewGetCommand()
	get.Run = func(cmd *cobra.Command, args []string) {
		key, opts := getGetOp(cmd, args)
		opc <- clientv3.OpGet(key, opts...)
	}
	del := NewDelCommand()
	del.Run = func(cmd *cobra.Command, args []string) {
		key, opts := getDelOp(cmd, args)
		opc <- clientv3.OpDelete(key, opts...)
	}
	cmds := &cobra.Command{SilenceErrors: true}
	cmds.AddCommand(put, get, del)
	cmds.SetArgs(args)
	if err := cmds.Execute(); err != nil {
		return nil, fmt.Errorf("invalid txn request: %s", line)
	}
	op := <-opc
	return &op, nil
}
func parseCompare(line string) (*clientv3.Cmp, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		key	string
		op	string
		val	string
	)
	lparenSplit := strings.SplitN(line, "(", 2)
	if len(lparenSplit) != 2 {
		return nil, fmt.Errorf("malformed comparison: %s", line)
	}
	target := lparenSplit[0]
	n, serr := fmt.Sscanf(lparenSplit[1], "%q) %s %q", &key, &op, &val)
	if n != 3 {
		return nil, fmt.Errorf("malformed comparison: %s; got %s(%q) %s %q", line, target, key, op, val)
	}
	if serr != nil {
		return nil, fmt.Errorf("malformed comparison: %s (%v)", line, serr)
	}
	var (
		v	int64
		err	error
		cmp	clientv3.Cmp
	)
	switch target {
	case "ver", "version":
		if v, err = strconv.ParseInt(val, 10, 64); err == nil {
			cmp = clientv3.Compare(clientv3.Version(key), op, v)
		}
	case "c", "create":
		if v, err = strconv.ParseInt(val, 10, 64); err == nil {
			cmp = clientv3.Compare(clientv3.CreateRevision(key), op, v)
		}
	case "m", "mod":
		if v, err = strconv.ParseInt(val, 10, 64); err == nil {
			cmp = clientv3.Compare(clientv3.ModRevision(key), op, v)
		}
	case "val", "value":
		cmp = clientv3.Compare(clientv3.Value(key), op, val)
	case "lease":
		cmp = clientv3.Compare(clientv3.Cmp{Target: pb.Compare_LEASE}, op, val)
	default:
		return nil, fmt.Errorf("malformed comparison: %s (unknown target %s)", line, target)
	}
	if err != nil {
		return nil, fmt.Errorf("invalid txn compare request: %s", line)
	}
	return &cmp, nil
}
