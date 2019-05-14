package tester

import "github.com/coreos/etcd/functional/rpcpb"

type noCheck struct{}

func newNoChecker() Checker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &noCheck{}
}
func (nc *noCheck) Type() rpcpb.Checker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rpcpb.Checker_NO_CHECK
}
func (nc *noCheck) EtcdClientEndpoints() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (nc *noCheck) Check() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
