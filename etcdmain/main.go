package etcdmain

import (
	"fmt"
	"github.com/coreos/go-systemd/daemon"
	systemdutil "github.com/coreos/go-systemd/util"
	"os"
	"strings"
)

func Main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	checkSupportArch()
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		if covArgs := os.Getenv("ETCDCOV_ARGS"); len(covArgs) > 0 {
			args := strings.Split(os.Getenv("ETCDCOV_ARGS"), "\xe7\xcd")[1:]
			rootCmd.SetArgs(args)
			cmd = "grpc-proxy"
		}
		switch cmd {
		case "gateway", "grpc-proxy":
			if err := rootCmd.Execute(); err != nil {
				fmt.Fprint(os.Stderr, err)
				os.Exit(1)
			}
			return
		}
	}
	startEtcdOrProxyV2()
}
func notifySystemd() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !systemdutil.IsRunningSystemd() {
		return
	}
	sent, err := daemon.SdNotify(false, "READY=1")
	if err != nil {
		plog.Errorf("failed to notify systemd for readiness: %v", err)
	}
	if !sent {
		plog.Errorf("forgot to set Type=notify in systemd service file?")
	}
}
