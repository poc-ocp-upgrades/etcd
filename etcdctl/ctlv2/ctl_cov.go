package ctlv2

import (
	"os"
	"strings"
	"github.com/urfave/cli"
)

func runCtlV2(app *cli.App) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return app.Run(strings.Split(os.Getenv("ETCDCTL_ARGS"), "\xe7\xcd"))
}
