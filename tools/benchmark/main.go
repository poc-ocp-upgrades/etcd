package main

import (
	"fmt"
	"os"
	"go.etcd.io/etcd/tools/benchmark/cmd"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
