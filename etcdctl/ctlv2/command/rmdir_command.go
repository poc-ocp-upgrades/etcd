package command

import (
	"errors"
	"github.com/coreos/etcd/client"
	"github.com/urfave/cli"
)

func NewRemoveDirCommand() cli.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cli.Command{Name: "rmdir", Usage: "removes the key if it is an empty directory or a key-value pair", ArgsUsage: "<key>", Action: func(c *cli.Context) error {
		rmdirCommandFunc(c, mustNewKeyAPI(c))
		return nil
	}}
}
func rmdirCommandFunc(c *cli.Context, ki client.KeysAPI) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(c.Args()) == 0 {
		handleError(c, ExitBadArgs, errors.New("key required"))
	}
	key := c.Args()[0]
	ctx, cancel := contextWithTotalTimeout(c)
	resp, err := ki.Delete(ctx, key, &client.DeleteOptions{Dir: true})
	cancel()
	if err != nil {
		handleError(c, ExitServerError, err)
	}
	if !resp.Node.Dir || c.GlobalString("output") != "simple" {
		printResponseKey(resp, c.GlobalString("output"))
	}
}
