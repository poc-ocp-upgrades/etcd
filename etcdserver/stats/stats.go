package stats

import "github.com/coreos/pkg/capnslog"

var (
	plog = capnslog.NewPackageLogger("github.com/coreos/etcd", "etcdserver/stats")
)

type Stats interface {
	SelfStats() []byte
	LeaderStats() []byte
	StoreStats() []byte
}
