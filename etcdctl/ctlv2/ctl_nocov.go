package ctlv2

import (
	"github.com/urfave/cli"
	"os"
)

func runCtlV2(app *cli.App) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return app.Run(os.Args)
}
