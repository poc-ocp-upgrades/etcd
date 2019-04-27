package report

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestPercentiles(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nums := make([]float64, 100)
	nums[99] = 1
	data := percentiles(nums)
	if data[len(pctls)-2] != 1 {
		t.Fatalf("99-percentile expected 1, got %f", data[len(pctls)-2])
	}
	nums = make([]float64, 1000)
	nums[999] = 1
	data = percentiles(nums)
	if data[len(pctls)-1] != 1 {
		t.Fatalf("99.9-percentile expected 1, got %f", data[len(pctls)-1])
	}
}
func TestReport(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := NewReportSample("%f")
	go func() {
		start := time.Now()
		for i := 0; i < 5; i++ {
			end := start.Add(time.Second)
			r.Results() <- Result{Start: start, End: end}
			start = end
		}
		r.Results() <- Result{Start: start, End: start.Add(time.Second), Err: fmt.Errorf("oops")}
		close(r.Results())
	}()
	stats := <-r.Stats()
	stats.TimeSeries = nil
	wStats := Stats{AvgTotal: 5.0, Fastest: 1.0, Slowest: 1.0, Average: 1.0, Stddev: 0.0, Total: stats.Total, RPS: 5.0 / stats.Total.Seconds(), ErrorDist: map[string]int{"oops": 1}, Lats: []float64{1.0, 1.0, 1.0, 1.0, 1.0}}
	if !reflect.DeepEqual(stats, wStats) {
		t.Fatalf("got %+v, want %+v", stats, wStats)
	}
	wstrs := []string{"Stddev:\t0", "Average:\t1.0", "Slowest:\t1.0", "Fastest:\t1.0"}
	ss := <-r.Run()
	for i, ws := range wstrs {
		if !strings.Contains(ss, ws) {
			t.Errorf("#%d: stats string missing %s", i, ws)
		}
	}
}
func TestWeightedReport(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := NewWeightedReport(NewReport("%f"), "%f")
	go func() {
		start := time.Now()
		for i := 0; i < 5; i++ {
			end := start.Add(time.Second)
			r.Results() <- Result{Start: start, End: end, Weight: 2.0}
			start = end
		}
		r.Results() <- Result{Start: start, End: start.Add(time.Second), Err: fmt.Errorf("oops")}
		close(r.Results())
	}()
	stats := <-r.Stats()
	stats.TimeSeries = nil
	wStats := Stats{AvgTotal: 10.0, Fastest: 0.5, Slowest: 0.5, Average: 0.5, Stddev: 0.0, Total: stats.Total, RPS: 10.0 / stats.Total.Seconds(), ErrorDist: map[string]int{"oops": 1}, Lats: []float64{0.5, 0.5, 0.5, 0.5, 0.5}}
	if !reflect.DeepEqual(stats, wStats) {
		t.Fatalf("got %+v, want %+v", stats, wStats)
	}
}
