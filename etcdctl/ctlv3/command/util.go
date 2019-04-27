package command

import (
	"context"
	"encoding/hex"
	"fmt"
	"regexp"
	pb "github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/spf13/cobra"
)

func printKV(isHex bool, valueOnly bool, kv *pb.KeyValue) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	k, v := string(kv.Key), string(kv.Value)
	if isHex {
		k = addHexPrefix(hex.EncodeToString(kv.Key))
		v = addHexPrefix(hex.EncodeToString(kv.Value))
	}
	if !valueOnly {
		fmt.Println(k)
	}
	fmt.Println(v)
}
func addHexPrefix(s string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := make([]byte, len(s)*2)
	for i := 0; i < len(s); i += 2 {
		ns[i*2] = '\\'
		ns[i*2+1] = 'x'
		ns[i*2+2] = s[i]
		ns[i*2+3] = s[i+1]
	}
	return string(ns)
}
func argify(s string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := regexp.MustCompile(`"(?:[^"\\]|\\.)*"|'[^']*'|[^'"\s]\S*[^'"\s]?`)
	args := r.FindAllString(s, -1)
	for i := range args {
		if len(args[i]) == 0 {
			continue
		}
		if args[i][0] == '\'' {
			args[i] = args[i][1 : len(args)-1]
		} else if args[i][0] == '"' {
			if _, err := fmt.Sscanf(args[i], "%q", &args[i]); err != nil {
				ExitWithError(ExitInvalidInput, err)
			}
		}
	}
	return args
}
func commandCtx(cmd *cobra.Command) (context.Context, context.CancelFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	timeOut, err := cmd.Flags().GetDuration("command-timeout")
	if err != nil {
		ExitWithError(ExitError, err)
	}
	return context.WithTimeout(context.Background(), timeOut)
}
