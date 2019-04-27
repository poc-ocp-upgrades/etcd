package wait

import (
	"testing"
	"time"
)

func TestWaitTime(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wt := NewTimeList()
	ch1 := wt.Wait(1)
	wt.Trigger(2)
	select {
	case <-ch1:
	default:
		t.Fatalf("cannot receive from ch as expected")
	}
	ch2 := wt.Wait(4)
	wt.Trigger(3)
	select {
	case <-ch2:
		t.Fatalf("unexpected to receive from ch2")
	default:
	}
	wt.Trigger(4)
	select {
	case <-ch2:
	default:
		t.Fatalf("cannot receive from ch2 as expected")
	}
	select {
	case <-wt.Wait(4):
	default:
		t.Fatalf("unexpected blocking when wait on triggered deadline")
	}
}
func TestWaitTestStress(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	chs := make([]<-chan struct{}, 0)
	wt := NewTimeList()
	for i := 0; i < 10000; i++ {
		chs = append(chs, wt.Wait(uint64(i)))
	}
	wt.Trigger(10000 + 1)
	for _, ch := range chs {
		select {
		case <-ch:
		case <-time.After(time.Second):
			t.Fatalf("cannot receive from ch as expected")
		}
	}
}
func BenchmarkWaitTime(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	wt := NewTimeList()
	for i := 0; i < b.N; i++ {
		wt.Wait(1)
	}
}
func BenchmarkTriggerAnd10KWaitTime(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < b.N; i++ {
		wt := NewTimeList()
		for j := 0; j < 10000; j++ {
			wt.Wait(uint64(j))
		}
		wt.Trigger(10000 + 1)
	}
}
