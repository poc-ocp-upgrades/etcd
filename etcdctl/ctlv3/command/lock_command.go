package command

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/spf13/cobra"
)

var lockTTL = 10

func NewLockCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &cobra.Command{Use: "lock <lockname> [exec-command arg1 arg2 ...]", Short: "Acquires a named lock", Run: lockCommandFunc}
	c.Flags().IntVarP(&lockTTL, "ttl", "", lockTTL, "timeout for session")
	return c
}
func lockCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 0 {
		ExitWithError(ExitBadArgs, errors.New("lock takes a lock name argument and an optional command to execute."))
	}
	c := mustClientFromCmd(cmd)
	if err := lockUntilSignal(c, args[0], args[1:]); err != nil {
		ExitWithError(ExitError, err)
	}
}
func lockUntilSignal(c *clientv3.Client, lockname string, cmdArgs []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := concurrency.NewSession(c, concurrency.WithTTL(lockTTL))
	if err != nil {
		return err
	}
	m := concurrency.NewMutex(s, lockname)
	ctx, cancel := context.WithCancel(context.TODO())
	donec := make(chan struct{})
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigc
		cancel()
		close(donec)
	}()
	if err := m.Lock(ctx); err != nil {
		return err
	}
	if len(cmdArgs) > 0 {
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		cmd.Env = append(environLockResponse(m), os.Environ()...)
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		err := cmd.Run()
		unlockErr := m.Unlock(context.TODO())
		if err != nil {
			return err
		}
		return unlockErr
	}
	k, kerr := c.Get(ctx, m.Key())
	if kerr != nil {
		return kerr
	}
	if len(k.Kvs) == 0 {
		return errors.New("lock lost on init")
	}
	display.Get(*k)
	select {
	case <-donec:
		return m.Unlock(context.TODO())
	case <-s.Done():
	}
	return errors.New("session expired")
}
func environLockResponse(m *concurrency.Mutex) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{"ETCD_LOCK_KEY=" + m.Key(), fmt.Sprintf("ETCD_LOCK_REV=%d", m.Header().Revision)}
}
