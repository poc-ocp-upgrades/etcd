package compactor

import (
	"reflect"
	"testing"
	"time"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/jonboulle/clockwork"
)

func TestRevision(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fc := clockwork.NewFakeClock()
	rg := &fakeRevGetter{testutil.NewRecorderStream(), 0}
	compactable := &fakeCompactable{testutil.NewRecorderStream()}
	tb := newRevision(fc, 10, rg, compactable)
	tb.Run()
	defer tb.Stop()
	fc.Advance(revInterval)
	rg.Wait(1)
	rg.SetRev(99)
	expectedRevision := int64(90)
	fc.Advance(revInterval)
	rg.Wait(1)
	a, err := compactable.Wait(1)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision}) {
		t.Errorf("compact request = %v, want %v", a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision})
	}
	rg.SetRev(99)
	rg.Wait(1)
	rg.SetRev(199)
	expectedRevision = int64(190)
	fc.Advance(revInterval)
	rg.Wait(1)
	a, err = compactable.Wait(1)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision}) {
		t.Errorf("compact request = %v, want %v", a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision})
	}
}
func TestRevisionPause(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fc := clockwork.NewFakeClock()
	rg := &fakeRevGetter{testutil.NewRecorderStream(), 99}
	compactable := &fakeCompactable{testutil.NewRecorderStream()}
	tb := newRevision(fc, 10, rg, compactable)
	tb.Run()
	tb.Pause()
	n := int(time.Hour / revInterval)
	for i := 0; i < 3*n; i++ {
		fc.Advance(revInterval)
	}
	select {
	case a := <-compactable.Chan():
		t.Fatalf("unexpected action %v", a)
	case <-time.After(10 * time.Millisecond):
	}
	tb.Resume()
	fc.Advance(revInterval)
	rg.Wait(1)
	a, err := compactable.Wait(1)
	if err != nil {
		t.Fatal(err)
	}
	wreq := &pb.CompactionRequest{Revision: int64(90)}
	if !reflect.DeepEqual(a[0].Params[0], wreq) {
		t.Errorf("compact request = %v, want %v", a[0].Params[0], wreq.Revision)
	}
}
