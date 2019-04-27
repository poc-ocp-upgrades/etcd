package command

import (
	"errors"
	"time"
	"github.com/coreos/etcd/client"
	"github.com/urfave/cli"
)

func NewUpdateDirCommand() cli.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cli.Command{Name: "updatedir", Usage: "update an existing directory", ArgsUsage: "<key> <value>", Flags: []cli.Flag{cli.IntFlag{Name: "ttl", Value: 0, Usage: "key time-to-live in seconds"}}, Action: func(c *cli.Context) error {
		updatedirCommandFunc(c, mustNewKeyAPI(c))
		return nil
	}}
}
func updatedirCommandFunc(c *cli.Context, ki client.KeysAPI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(c.Args()) == 0 {
		handleError(c, ExitBadArgs, errors.New("key required"))
	}
	key := c.Args()[0]
	ttl := c.Int("ttl")
	ctx, cancel := contextWithTotalTimeout(c)
	resp, err := ki.Set(ctx, key, "", &client.SetOptions{TTL: time.Duration(ttl) * time.Second, Dir: true, PrevExist: client.PrevExist})
	cancel()
	if err != nil {
		handleError(c, ExitServerError, err)
	}
	if c.GlobalString("output") != "simple" {
		printResponseKey(resp, c.GlobalString("output"))
	}
}
