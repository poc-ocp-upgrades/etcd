package tester

import "github.com/coreos/etcd/functional/rpcpb"

type runnerChecker struct {
	ctype			rpcpb.Checker
	etcdClientEndpoint	string
	errc			chan error
}

func newRunnerChecker(ep string, errc chan error) Checker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &runnerChecker{ctype: rpcpb.Checker_RUNNER, etcdClientEndpoint: ep, errc: errc}
}
func (rc *runnerChecker) Type() rpcpb.Checker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rc.ctype
}
func (rc *runnerChecker) EtcdClientEndpoints() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{rc.etcdClientEndpoint}
}
func (rc *runnerChecker) Check() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case err := <-rc.errc:
		return err
	default:
		return nil
	}
}
