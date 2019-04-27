package command

import (
	"github.com/coreos/etcd/client"
	"github.com/urfave/cli"
)

func NewSetDirCommand() cli.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cli.Command{Name: "setdir", Usage: "create a new directory or update an existing directory TTL", ArgsUsage: "<key>", Flags: []cli.Flag{cli.IntFlag{Name: "ttl", Value: 0, Usage: "key time-to-live in seconds"}}, Action: func(c *cli.Context) error {
		mkdirCommandFunc(c, mustNewKeyAPI(c), client.PrevIgnore)
		return nil
	}}
}
