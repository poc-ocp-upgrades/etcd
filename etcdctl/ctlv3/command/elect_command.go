package command

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/spf13/cobra"
)

var (
	electListen bool
)

func NewElectCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "elect <election-name> [proposal]", Short: "Observes and participates in leader election", Run: electCommandFunc}
	cmd.Flags().BoolVarP(&electListen, "listen", "l", false, "observation mode")
	return cmd
}
func electCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 && len(args) != 2 {
		ExitWithError(ExitBadArgs, errors.New("elect takes one election name argument and an optional proposal argument."))
	}
	c := mustClientFromCmd(cmd)
	var err error
	if len(args) == 1 {
		if !electListen {
			ExitWithError(ExitBadArgs, errors.New("no proposal argument but -l not set"))
		}
		err = observe(c, args[0])
	} else {
		if electListen {
			ExitWithError(ExitBadArgs, errors.New("proposal given but -l is set"))
		}
		err = campaign(c, args[0], args[1])
	}
	if err != nil {
		ExitWithError(ExitError, err)
	}
}
func observe(c *clientv3.Client, election string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := concurrency.NewSession(c)
	if err != nil {
		return err
	}
	e := concurrency.NewElection(s, election)
	ctx, cancel := context.WithCancel(context.TODO())
	donec := make(chan struct{})
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigc
		cancel()
	}()
	go func() {
		for resp := range e.Observe(ctx) {
			display.Get(resp)
		}
		close(donec)
	}()
	<-donec
	select {
	case <-ctx.Done():
	default:
		return errors.New("elect: observer lost")
	}
	return nil
}
func campaign(c *clientv3.Client, election string, prop string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := concurrency.NewSession(c)
	if err != nil {
		return err
	}
	e := concurrency.NewElection(s, election)
	ctx, cancel := context.WithCancel(context.TODO())
	donec := make(chan struct{})
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigc
		cancel()
		close(donec)
	}()
	if err = e.Campaign(ctx, prop); err != nil {
		return err
	}
	resp, err := c.Get(ctx, e.Key())
	if err != nil {
		return err
	}
	display.Get(*resp)
	select {
	case <-donec:
	case <-s.Done():
		return errors.New("elect: session expired")
	}
	return e.Resign(context.TODO())
}
