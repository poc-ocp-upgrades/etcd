package command

import (
	"errors"
	"os"
	"time"
	"github.com/coreos/etcd/client"
	"github.com/urfave/cli"
)

func NewUpdateCommand() cli.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cli.Command{Name: "update", Usage: "update an existing key with a given value", ArgsUsage: "<key> <value>", Flags: []cli.Flag{cli.IntFlag{Name: "ttl", Value: 0, Usage: "key time-to-live in seconds"}}, Action: func(c *cli.Context) error {
		updateCommandFunc(c, mustNewKeyAPI(c))
		return nil
	}}
}
func updateCommandFunc(c *cli.Context, ki client.KeysAPI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(c.Args()) == 0 {
		handleError(c, ExitBadArgs, errors.New("key required"))
	}
	key := c.Args()[0]
	value, err := argOrStdin(c.Args(), os.Stdin, 1)
	if err != nil {
		handleError(c, ExitBadArgs, errors.New("value required"))
	}
	ttl := c.Int("ttl")
	ctx, cancel := contextWithTotalTimeout(c)
	resp, err := ki.Set(ctx, key, value, &client.SetOptions{TTL: time.Duration(ttl) * time.Second, PrevExist: client.PrevExist})
	cancel()
	if err != nil {
		handleError(c, ExitServerError, err)
	}
	printResponseKey(resp, c.GlobalString("output"))
}
