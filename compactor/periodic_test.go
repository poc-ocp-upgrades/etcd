package compactor

import (
	"reflect"
	"testing"
	"time"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/jonboulle/clockwork"
)

func TestPeriodicHourly(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	retentionHours := 2
	retentionDuration := time.Duration(retentionHours) * time.Hour
	fc := clockwork.NewFakeClock()
	rg := &fakeRevGetter{testutil.NewRecorderStream(), 0}
	compactable := &fakeCompactable{testutil.NewRecorderStream()}
	tb := newPeriodic(fc, retentionDuration, rg, compactable)
	tb.Run()
	defer tb.Stop()
	initialIntervals, intervalsPerPeriod := tb.getRetentions(), 10
	for i := 0; i < initialIntervals; i++ {
		rg.Wait(1)
		fc.Advance(tb.getRetryInterval())
	}
	a, err := compactable.Wait(1)
	if err != nil {
		t.Fatal(err)
	}
	expectedRevision := int64(1)
	if !reflect.DeepEqual(a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision}) {
		t.Errorf("compact request = %v, want %v", a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision})
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < intervalsPerPeriod; j++ {
			rg.Wait(1)
			fc.Advance(tb.getRetryInterval())
		}
		a, err = compactable.Wait(1)
		if err != nil {
			t.Fatal(err)
		}
		expectedRevision = int64((i + 1) * 10)
		if !reflect.DeepEqual(a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision}) {
			t.Errorf("compact request = %v, want %v", a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision})
		}
	}
}
func TestPeriodicMinutes(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	retentionMinutes := 5
	retentionDuration := time.Duration(retentionMinutes) * time.Minute
	fc := clockwork.NewFakeClock()
	rg := &fakeRevGetter{testutil.NewRecorderStream(), 0}
	compactable := &fakeCompactable{testutil.NewRecorderStream()}
	tb := newPeriodic(fc, retentionDuration, rg, compactable)
	tb.Run()
	defer tb.Stop()
	initialIntervals, intervalsPerPeriod := tb.getRetentions(), 10
	for i := 0; i < initialIntervals; i++ {
		rg.Wait(1)
		fc.Advance(tb.getRetryInterval())
	}
	a, err := compactable.Wait(1)
	if err != nil {
		t.Fatal(err)
	}
	expectedRevision := int64(1)
	if !reflect.DeepEqual(a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision}) {
		t.Errorf("compact request = %v, want %v", a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision})
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < intervalsPerPeriod; j++ {
			rg.Wait(1)
			fc.Advance(tb.getRetryInterval())
		}
		a, err := compactable.Wait(1)
		if err != nil {
			t.Fatal(err)
		}
		expectedRevision = int64((i + 1) * 10)
		if !reflect.DeepEqual(a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision}) {
			t.Errorf("compact request = %v, want %v", a[0].Params[0], &pb.CompactionRequest{Revision: expectedRevision})
		}
	}
}
func TestPeriodicPause(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fc := clockwork.NewFakeClock()
	retentionDuration := time.Hour
	rg := &fakeRevGetter{testutil.NewRecorderStream(), 0}
	compactable := &fakeCompactable{testutil.NewRecorderStream()}
	tb := newPeriodic(fc, retentionDuration, rg, compactable)
	tb.Run()
	tb.Pause()
	n := tb.getRetentions()
	for i := 0; i < n*3; i++ {
		rg.Wait(1)
		fc.Advance(tb.getRetryInterval())
	}
	select {
	case a := <-compactable.Chan():
		t.Fatalf("unexpected action %v", a)
	case <-time.After(10 * time.Millisecond):
	}
	tb.Resume()
	rg.Wait(1)
	fc.Advance(tb.getRetryInterval())
	a, err := compactable.Wait(1)
	if err != nil {
		t.Fatal(err)
	}
	wreq := &pb.CompactionRequest{Revision: int64(1 + 2*n + 1)}
	if !reflect.DeepEqual(a[0].Params[0], wreq) {
		t.Errorf("compact request = %v, want %v", a[0].Params[0], wreq.Revision)
	}
}
