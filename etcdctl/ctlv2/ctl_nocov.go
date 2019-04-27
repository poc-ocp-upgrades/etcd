package ctlv2

import (
	"os"
	"github.com/urfave/cli"
)

func runCtlV2(app *cli.App) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return app.Run(os.Args)
}
