package api

import (
	"go.etcd.io/etcd/etcdserver/api/membership"
	"go.etcd.io/etcd/pkg/types"
	"github.com/coreos/go-semver/semver"
)

type Cluster interface {
	ID() types.ID
	ClientURLs() []string
	Members() []*membership.Member
	Member(id types.ID) *membership.Member
	Version() *semver.Version
}
