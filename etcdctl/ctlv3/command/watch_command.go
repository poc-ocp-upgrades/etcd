package command

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"go.etcd.io/etcd/clientv3"
	"github.com/spf13/cobra"
)

var (
	errBadArgsNum			= errors.New("bad number of arguments")
	errBadArgsNumConflictEnv	= errors.New("bad number of arguments (found conflicting environment key)")
	errBadArgsNumSeparator		= errors.New("bad number of arguments (found separator --, but no commands)")
	errBadArgsInteractiveWatch	= errors.New("args[0] must be 'watch' for interactive calls")
)
var (
	watchRev		int64
	watchPrefix		bool
	watchInteractive	bool
	watchPrevKey		bool
)

func NewWatchCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "watch [options] [key or prefix] [range_end] [--] [exec-command arg1 arg2 ...]", Short: "Watches events stream on keys or prefixes", Run: watchCommandFunc}
	cmd.Flags().BoolVarP(&watchInteractive, "interactive", "i", false, "Interactive mode")
	cmd.Flags().BoolVar(&watchPrefix, "prefix", false, "Watch on a prefix if prefix is set")
	cmd.Flags().Int64Var(&watchRev, "rev", 0, "Revision to start watching")
	cmd.Flags().BoolVar(&watchPrevKey, "prev-kv", false, "get the previous key-value pair before the event happens")
	return cmd
}
func watchCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	envKey, envRange := os.Getenv("ETCDCTL_WATCH_KEY"), os.Getenv("ETCDCTL_WATCH_RANGE_END")
	if envKey == "" && envRange != "" {
		ExitWithError(ExitBadArgs, fmt.Errorf("ETCDCTL_WATCH_KEY is empty but got ETCDCTL_WATCH_RANGE_END=%q", envRange))
	}
	if watchInteractive {
		watchInteractiveFunc(cmd, os.Args, envKey, envRange)
		return
	}
	watchArgs, execArgs, err := parseWatchArgs(os.Args, args, envKey, envRange, false)
	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}
	c := mustClientFromCmd(cmd)
	wc, err := getWatchChan(c, watchArgs)
	if err != nil {
		ExitWithError(ExitBadArgs, err)
	}
	printWatchCh(c, wc, execArgs)
	if err = c.Close(); err != nil {
		ExitWithError(ExitBadConnection, err)
	}
	ExitWithError(ExitInterrupted, fmt.Errorf("watch is canceled by the server"))
}
func watchInteractiveFunc(cmd *cobra.Command, osArgs []string, envKey, envRange string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := mustClientFromCmd(cmd)
	reader := bufio.NewReader(os.Stdin)
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			ExitWithError(ExitInvalidInput, fmt.Errorf("Error reading watch request line: %v", err))
		}
		l = strings.TrimSuffix(l, "\n")
		args := argify(l)
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "Invalid command: %s (watch and progress supported)\n", l)
			continue
		}
		switch args[0] {
		case "watch":
			if len(args) < 2 && envKey == "" {
				fmt.Fprintf(os.Stderr, "Invalid command %s (command type or key is not provided)\n", l)
				continue
			}
			watchArgs, execArgs, perr := parseWatchArgs(osArgs, args, envKey, envRange, true)
			if perr != nil {
				ExitWithError(ExitBadArgs, perr)
			}
			ch, err := getWatchChan(c, watchArgs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid command %s (%v)\n", l, err)
				continue
			}
			go printWatchCh(c, ch, execArgs)
		case "progress":
			err := c.RequestProgress(clientv3.WithRequireLeader(context.Background()))
			if err != nil {
				ExitWithError(ExitError, err)
			}
		default:
			fmt.Fprintf(os.Stderr, "Invalid command %s (only support watch)\n", l)
			continue
		}
	}
}
func getWatchChan(c *clientv3.Client, args []string) (clientv3.WatchChan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) < 1 {
		return nil, errBadArgsNum
	}
	key := args[0]
	opts := []clientv3.OpOption{clientv3.WithRev(watchRev)}
	if len(args) == 2 {
		if watchPrefix {
			return nil, fmt.Errorf("`range_end` and `--prefix` are mutually exclusive")
		}
		opts = append(opts, clientv3.WithRange(args[1]))
	}
	if watchPrefix {
		opts = append(opts, clientv3.WithPrefix())
	}
	if watchPrevKey {
		opts = append(opts, clientv3.WithPrevKV())
	}
	return c.Watch(clientv3.WithRequireLeader(context.Background()), key, opts...), nil
}
func printWatchCh(c *clientv3.Client, ch clientv3.WatchChan, execArgs []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for resp := range ch {
		if resp.Canceled {
			fmt.Fprintf(os.Stderr, "watch was canceled (%v)\n", resp.Err())
		}
		if resp.IsProgressNotify() {
			fmt.Fprintf(os.Stdout, "progress notify: %d\n", resp.Header.Revision)
		}
		display.Watch(resp)
		if len(execArgs) > 0 {
			for _, ev := range resp.Events {
				cmd := exec.CommandContext(c.Ctx(), execArgs[0], execArgs[1:]...)
				cmd.Env = os.Environ()
				cmd.Env = append(cmd.Env, fmt.Sprintf("ETCD_WATCH_REVISION=%d", resp.Header.Revision))
				cmd.Env = append(cmd.Env, fmt.Sprintf("ETCD_WATCH_EVENT_TYPE=%q", ev.Type))
				cmd.Env = append(cmd.Env, fmt.Sprintf("ETCD_WATCH_KEY=%q", ev.Kv.Key))
				cmd.Env = append(cmd.Env, fmt.Sprintf("ETCD_WATCH_VALUE=%q", ev.Kv.Value))
				cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Fprintf(os.Stderr, "command %q error (%v)\n", execArgs, err)
					os.Exit(1)
				}
			}
		}
	}
}
func parseWatchArgs(osArgs, commandArgs []string, envKey, envRange string, interactive bool) (watchArgs []string, execArgs []string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rawArgs := make([]string, len(osArgs))
	copy(rawArgs, osArgs)
	watchArgs = make([]string, len(commandArgs))
	copy(watchArgs, commandArgs)
	for idx := range rawArgs {
		if rawArgs[idx] == "watch" {
			rawArgs = rawArgs[idx+1:]
			break
		}
	}
	if interactive {
		if watchArgs[0] != "watch" {
			watchPrefix, watchRev, watchPrevKey = false, 0, false
			return nil, nil, errBadArgsInteractiveWatch
		}
		watchArgs = watchArgs[1:]
	}
	execIdx, execExist := 0, false
	if !interactive {
		for execIdx = range rawArgs {
			if rawArgs[execIdx] == "--" {
				execExist = true
				break
			}
		}
		if execExist && execIdx == len(rawArgs)-1 {
			return nil, nil, errBadArgsNumSeparator
		}
		if !execExist && len(rawArgs) < 1 && envKey == "" {
			return nil, nil, errBadArgsNum
		}
		if execExist && envKey != "" {
			widx, ridx := len(watchArgs)-1, len(rawArgs)-1
			for ; widx >= 0; widx-- {
				if watchArgs[widx] == rawArgs[ridx] {
					ridx--
					continue
				}
				if ridx == execIdx {
					return nil, nil, errBadArgsNumConflictEnv
				}
			}
		}
		if !execExist && len(watchArgs) > 0 && envKey != "" {
			return nil, nil, errBadArgsNumConflictEnv
		}
	} else {
		for execIdx = range watchArgs {
			if watchArgs[execIdx] == "--" {
				execExist = true
				break
			}
		}
		if execExist && execIdx == len(watchArgs)-1 {
			watchPrefix, watchRev, watchPrevKey = false, 0, false
			return nil, nil, errBadArgsNumSeparator
		}
		flagset := NewWatchCommand().Flags()
		if perr := flagset.Parse(watchArgs); perr != nil {
			watchPrefix, watchRev, watchPrevKey = false, 0, false
			return nil, nil, perr
		}
		pArgs := flagset.Args()
		if !execExist && envKey == "" && len(pArgs) < 1 {
			watchPrefix, watchRev, watchPrevKey = false, 0, false
			return nil, nil, errBadArgsNum
		}
		if !execExist && len(pArgs) > 0 && envKey != "" {
			watchPrefix, watchRev, watchPrevKey = false, 0, false
			return nil, nil, errBadArgsNumConflictEnv
		}
	}
	argsWithSep := rawArgs
	if interactive {
		argsWithSep = watchArgs
	}
	idx, foundSep := 0, false
	for idx = range argsWithSep {
		if argsWithSep[idx] == "--" {
			foundSep = true
			break
		}
	}
	if foundSep {
		execArgs = argsWithSep[idx+1:]
	}
	if interactive {
		flagset := NewWatchCommand().Flags()
		if perr := flagset.Parse(argsWithSep); perr != nil {
			return nil, nil, perr
		}
		watchArgs = flagset.Args()
		watchPrefix, err = flagset.GetBool("prefix")
		if err != nil {
			return nil, nil, err
		}
		watchRev, err = flagset.GetInt64("rev")
		if err != nil {
			return nil, nil, err
		}
		watchPrevKey, err = flagset.GetBool("prev-kv")
		if err != nil {
			return nil, nil, err
		}
	}
	if envKey != "" {
		ranges := []string{envKey}
		if envRange != "" {
			ranges = append(ranges, envRange)
		}
		watchArgs = append(ranges, watchArgs...)
	}
	if !foundSep {
		return watchArgs, nil, nil
	}
	endIdx := 0
	for endIdx = len(watchArgs) - 1; endIdx >= 0; endIdx-- {
		if watchArgs[endIdx] == argsWithSep[idx+1] {
			break
		}
	}
	watchArgs = watchArgs[:endIdx]
	return watchArgs, execArgs, nil
}
