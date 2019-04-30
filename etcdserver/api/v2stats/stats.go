package v2stats

import "github.com/coreos/pkg/capnslog"

var plog = capnslog.NewPackageLogger("go.etcd.io/etcd", "etcdserver/stats")

type Stats interface {
	SelfStats() []byte
	LeaderStats() []byte
	StoreStats() []byte
}
