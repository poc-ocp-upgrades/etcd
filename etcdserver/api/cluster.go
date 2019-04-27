package api

import (
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/go-semver/semver"
)

type Cluster interface {
	ID() types.ID
	ClientURLs() []string
	Members() []*membership.Member
	Member(id types.ID) *membership.Member
	Version() *semver.Version
}
