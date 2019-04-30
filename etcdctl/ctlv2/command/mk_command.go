package command

import (
	"errors"
	"os"
	"time"
	"github.com/urfave/cli"
	"go.etcd.io/etcd/client"
)

func NewMakeCommand() cli.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cli.Command{Name: "mk", Usage: "make a new key with a given value", ArgsUsage: "<key> <value>", Flags: []cli.Flag{cli.BoolFlag{Name: "in-order", Usage: "create in-order key under directory <key>"}, cli.IntFlag{Name: "ttl", Value: 0, Usage: "key time-to-live in seconds"}}, Action: func(c *cli.Context) error {
		mkCommandFunc(c, mustNewKeyAPI(c))
		return nil
	}}
}
func mkCommandFunc(c *cli.Context, ki client.KeysAPI) {
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
	inorder := c.Bool("in-order")
	var resp *client.Response
	ctx, cancel := contextWithTotalTimeout(c)
	if !inorder {
		resp, err = ki.Set(ctx, key, value, &client.SetOptions{TTL: time.Duration(ttl) * time.Second, PrevExist: client.PrevNoExist})
	} else {
		resp, err = ki.CreateInOrder(ctx, key, value, &client.CreateInOrderOptions{TTL: time.Duration(ttl) * time.Second})
	}
	cancel()
	if err != nil {
		handleError(c, ExitServerError, err)
	}
	printResponseKey(resp, c.GlobalString("output"))
}
