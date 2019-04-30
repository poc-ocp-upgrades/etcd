package tester

import "go.etcd.io/etcd/functional/rpcpb"

type Checker interface {
	Type() rpcpb.Checker
	EtcdClientEndpoints() []string
	Check() error
}
