package e2e

import "strings"

type kvExec struct {
	key, val	string
	execOutput	string
}

func setupWatchArgs(cx ctlCtx, args []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "watch")
	if cx.interactive {
		cmdArgs = append(cmdArgs, "--interactive")
	} else {
		cmdArgs = append(cmdArgs, args...)
	}
	return cmdArgs
}
func ctlV3Watch(cx ctlCtx, args []string, kvs ...kvExec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := setupWatchArgs(cx, args)
	proc, err := spawnCmd(cmdArgs)
	if err != nil {
		return err
	}
	if cx.interactive {
		wl := strings.Join(append([]string{"watch"}, args...), " ") + "\r"
		if err = proc.Send(wl); err != nil {
			return err
		}
	}
	for _, elem := range kvs {
		if _, err = proc.Expect(elem.key); err != nil {
			return err
		}
		if _, err = proc.Expect(elem.val); err != nil {
			return err
		}
		if elem.execOutput != "" {
			if _, err = proc.Expect(elem.execOutput); err != nil {
				return err
			}
		}
	}
	return proc.Stop()
}
func ctlV3WatchFailPerm(cx ctlCtx, args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := setupWatchArgs(cx, args)
	proc, err := spawnCmd(cmdArgs)
	if err != nil {
		return err
	}
	if cx.interactive {
		wl := strings.Join(append([]string{"watch"}, args...), " ") + "\r"
		if err = proc.Send(wl); err != nil {
			return err
		}
	}
	_, err = proc.Expect("watch is canceled by the server")
	if err != nil {
		return err
	}
	return proc.Close()
}
