package mockwait

import (
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/pkg/wait"
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
