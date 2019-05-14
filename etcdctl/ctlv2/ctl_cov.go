package ctlv2

import (
	"github.com/urfave/cli"
	"os"
	"strings"
)

func runCtlV2(app *cli.App) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return app.Run(strings.Split(os.Getenv("ETCDCTL_ARGS"), "\xe7\xcd"))
}
