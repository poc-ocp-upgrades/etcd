package command

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/urfave/cli"
	"go.etcd.io/etcd/client"
)

const (
	ExitSuccess	= iota
	ExitBadArgs
	ExitBadConnection
	ExitBadAuth
	ExitServerError
	ExitClusterNotHealthy
)

func handleError(c *cli.Context, code int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.GlobalString("output") == "json" {
		if err, ok := err.(*client.Error); ok {
			b, err := json.Marshal(err)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(os.Stderr, string(b))
			os.Exit(code)
		}
	}
	fmt.Fprintln(os.Stderr, "Error: ", err)
	if cerr, ok := err.(*client.ClusterError); ok {
		fmt.Fprintln(os.Stderr, cerr.Detail())
	}
	os.Exit(code)
}
