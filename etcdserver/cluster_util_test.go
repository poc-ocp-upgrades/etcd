package etcdserver

import (
	"reflect"
	"testing"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/version"
	"github.com/coreos/go-semver/semver"
)

func TestDecideClusterVersion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		vers	map[string]*version.Versions
		wdver	*semver.Version
	}{{map[string]*version.Versions{"a": {Server: "2.0.0"}}, semver.Must(semver.NewVersion("2.0.0"))}, {map[string]*version.Versions{"a": nil}, nil}, {map[string]*version.Versions{"a": {Server: "2.0.0"}, "b": {Server: "2.1.0"}, "c": {Server: "2.1.0"}}, semver.Must(semver.NewVersion("2.0.0"))}, {map[string]*version.Versions{"a": {Server: "2.1.0"}, "b": {Server: "2.1.0"}, "c": {Server: "2.1.0"}}, semver.Must(semver.NewVersion("2.1.0"))}, {map[string]*version.Versions{"a": nil, "b": {Server: "2.1.0"}, "c": {Server: "2.1.0"}}, nil}}
	for i, tt := range tests {
		dver := decideClusterVersion(tt.vers)
		if !reflect.DeepEqual(dver, tt.wdver) {
			t.Errorf("#%d: ver = %+v, want %+v", i, dver, tt.wdver)
		}
	}
}
func TestIsCompatibleWithVers(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		vers		map[string]*version.Versions
		local		types.ID
		minV, maxV	*semver.Version
		wok		bool
	}{{map[string]*version.Versions{"a": {Server: "2.0.0", Cluster: "not_decided"}, "b": {Server: "2.1.0", Cluster: "2.1.0"}, "c": {Server: "2.1.0", Cluster: "2.1.0"}}, 0xa, semver.Must(semver.NewVersion("2.0.0")), semver.Must(semver.NewVersion("2.0.0")), false}, {map[string]*version.Versions{"a": {Server: "2.1.0", Cluster: "not_decided"}, "b": {Server: "2.1.0", Cluster: "2.1.0"}, "c": {Server: "2.1.0", Cluster: "2.1.0"}}, 0xa, semver.Must(semver.NewVersion("2.0.0")), semver.Must(semver.NewVersion("2.1.0")), true}, {map[string]*version.Versions{"a": {Server: "2.2.0", Cluster: "not_decided"}, "b": {Server: "2.0.0", Cluster: "2.0.0"}, "c": {Server: "2.0.0", Cluster: "2.0.0"}}, 0xa, semver.Must(semver.NewVersion("2.1.0")), semver.Must(semver.NewVersion("2.2.0")), false}, {map[string]*version.Versions{"a": {Server: "2.1.0", Cluster: "not_decided"}, "b": nil, "c": {Server: "2.1.0", Cluster: "2.1.0"}}, 0xa, semver.Must(semver.NewVersion("2.0.0")), semver.Must(semver.NewVersion("2.1.0")), true}, {map[string]*version.Versions{"a": {Server: "2.1.0", Cluster: "not_decided"}, "b": nil, "c": nil}, 0xa, semver.Must(semver.NewVersion("2.0.0")), semver.Must(semver.NewVersion("2.1.0")), false}}
	for i, tt := range tests {
		ok := isCompatibleWithVers(tt.vers, tt.local, tt.minV, tt.maxV)
		if ok != tt.wok {
			t.Errorf("#%d: ok = %+v, want %+v", i, ok, tt.wok)
		}
	}
}
