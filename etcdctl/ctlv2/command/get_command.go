package command

import (
	"errors"
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/urfave/cli"
	"os"
)

func NewGetCommand() cli.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cli.Command{Name: "get", Usage: "retrieve the value of a key", ArgsUsage: "<key>", Flags: []cli.Flag{cli.BoolFlag{Name: "sort", Usage: "returns result in sorted order"}, cli.BoolFlag{Name: "quorum, q", Usage: "require quorum for get request"}}, Action: func(c *cli.Context) error {
		getCommandFunc(c, mustNewKeyAPI(c))
		return nil
	}}
}
func getCommandFunc(c *cli.Context, ki client.KeysAPI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(c.Args()) == 0 {
		handleError(c, ExitBadArgs, errors.New("key required"))
	}
	key := c.Args()[0]
	sorted := c.Bool("sort")
	quorum := c.Bool("quorum")
	ctx, cancel := contextWithTotalTimeout(c)
	resp, err := ki.Get(ctx, key, &client.GetOptions{Sort: sorted, Quorum: quorum})
	cancel()
	if err != nil {
		handleError(c, ExitServerError, err)
	}
	if resp.Node.Dir {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("%s: is a directory", resp.Node.Key))
		os.Exit(1)
	}
	printResponseKey(resp, c.GlobalString("output"))
}
