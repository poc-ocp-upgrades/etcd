package tester

import "github.com/coreos/etcd/functional/rpcpb"

type Checker interface {
	Type() rpcpb.Checker
	EtcdClientEndpoints() []string
	Check() error
}
