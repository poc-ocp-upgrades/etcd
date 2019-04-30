package command

import (
	"errors"
	"os"
	"time"
	"github.com/urfave/cli"
	"go.etcd.io/etcd/client"
)

func NewSetCommand() cli.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cli.Command{Name: "set", Usage: "set the value of a key", ArgsUsage: "<key> <value>", Description: `Set sets the value of a key.

   When <value> begins with '-', <value> is interpreted as a flag.
   Insert '--' for workaround:

   $ set -- <key> <value>`, Flags: []cli.Flag{cli.IntFlag{Name: "ttl", Value: 0, Usage: "key time-to-live in seconds"}, cli.StringFlag{Name: "swap-with-value", Value: "", Usage: "previous value"}, cli.IntFlag{Name: "swap-with-index", Value: 0, Usage: "previous index"}}, Action: func(c *cli.Context) error {
		setCommandFunc(c, mustNewKeyAPI(c))
		return nil
	}}
}
func setCommandFunc(c *cli.Context, ki client.KeysAPI) {
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
	prevValue := c.String("swap-with-value")
	prevIndex := c.Int("swap-with-index")
	ctx, cancel := contextWithTotalTimeout(c)
	resp, err := ki.Set(ctx, key, value, &client.SetOptions{TTL: time.Duration(ttl) * time.Second, PrevIndex: uint64(prevIndex), PrevValue: prevValue})
	cancel()
	if err != nil {
		handleError(c, ExitServerError, err)
	}
	printResponseKey(resp, c.GlobalString("output"))
}
