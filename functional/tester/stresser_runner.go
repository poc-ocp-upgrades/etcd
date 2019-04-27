package tester

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"syscall"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

type runnerStresser struct {
	stype			rpcpb.Stresser
	etcdClientEndpoint	string
	lg			*zap.Logger
	cmd			*exec.Cmd
	cmdStr			string
	args			[]string
	rl			*rate.Limiter
	reqRate			int
	errc			chan error
	donec			chan struct{}
}

func newRunnerStresser(stype rpcpb.Stresser, ep string, lg *zap.Logger, cmdStr string, args []string, rl *rate.Limiter, reqRate int) *runnerStresser {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rl.SetLimit(rl.Limit() - rate.Limit(reqRate))
	return &runnerStresser{stype: stype, etcdClientEndpoint: ep, cmdStr: cmdStr, args: args, rl: rl, reqRate: reqRate, errc: make(chan error, 1), donec: make(chan struct{})}
}
func (rs *runnerStresser) setupOnce() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rs.cmd != nil {
		return nil
	}
	rs.cmd = exec.Command(rs.cmdStr, rs.args...)
	stderr, err := rs.cmd.StderrPipe()
	if err != nil {
		return err
	}
	go func() {
		defer close(rs.donec)
		out, err := ioutil.ReadAll(stderr)
		if err != nil {
			rs.errc <- err
		} else {
			rs.errc <- fmt.Errorf("(%v %v) stderr %v", rs.cmdStr, rs.args, string(out))
		}
	}()
	return rs.cmd.Start()
}
func (rs *runnerStresser) Stress() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rs.lg.Info("stress START", zap.String("stress-type", rs.stype.String()))
	if err = rs.setupOnce(); err != nil {
		return err
	}
	return syscall.Kill(rs.cmd.Process.Pid, syscall.SIGCONT)
}
func (rs *runnerStresser) Pause() map[string]int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rs.lg.Info("stress STOP", zap.String("stress-type", rs.stype.String()))
	syscall.Kill(rs.cmd.Process.Pid, syscall.SIGSTOP)
	return nil
}
func (rs *runnerStresser) Close() map[string]int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	syscall.Kill(rs.cmd.Process.Pid, syscall.SIGINT)
	rs.cmd.Wait()
	<-rs.donec
	rs.rl.SetLimit(rs.rl.Limit() + rate.Limit(rs.reqRate))
	return nil
}
func (rs *runnerStresser) ModifiedKeys() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 1
}
