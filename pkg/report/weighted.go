package report

import (
	"time"
)

type weightedReport struct {
	baseReport	Report
	report		*report
	results		chan Result
	weightTotal	float64
}

func NewWeightedReport(r Report, precision string) Report {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &weightedReport{baseReport: r, report: newReport(precision), results: make(chan Result, 16)}
}
func (wr *weightedReport) Results() chan<- Result {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wr.results
}
func (wr *weightedReport) Run() <-chan string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	donec := make(chan string, 2)
	go func() {
		defer close(donec)
		basec, rc := make(chan string, 1), make(chan Stats, 1)
		go func() {
			basec <- (<-wr.baseReport.Run())
		}()
		go func() {
			rc <- (<-wr.report.Stats())
		}()
		go wr.processResults()
		wr.report.stats = wr.reweighStat(<-rc)
		donec <- wr.report.String()
		donec <- (<-basec)
	}()
	return donec
}
func (wr *weightedReport) Stats() <-chan Stats {
	_logClusterCodePath()
	defer _logClusterCodePath()
	donec := make(chan Stats, 2)
	go func() {
		defer close(donec)
		basec, rc := make(chan Stats, 1), make(chan Stats, 1)
		go func() {
			basec <- (<-wr.baseReport.Stats())
		}()
		go func() {
			rc <- (<-wr.report.Stats())
		}()
		go wr.processResults()
		donec <- wr.reweighStat(<-rc)
		donec <- (<-basec)
	}()
	return donec
}
func (wr *weightedReport) processResults() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(wr.report.results)
	defer close(wr.baseReport.Results())
	for res := range wr.results {
		wr.processResult(res)
		wr.baseReport.Results() <- res
	}
}
func (wr *weightedReport) processResult(res Result) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if res.Err != nil {
		wr.report.results <- res
		return
	}
	if res.Weight == 0 {
		res.Weight = 1.0
	}
	wr.weightTotal += res.Weight
	res.End = res.Start.Add(time.Duration(float64(res.End.Sub(res.Start)) / res.Weight))
	res.Weight = 1.0
	wr.report.results <- res
}
func (wr *weightedReport) reweighStat(s Stats) Stats {
	_logClusterCodePath()
	defer _logClusterCodePath()
	weightCoef := wr.weightTotal / float64(len(s.Lats))
	s.RPS *= weightCoef
	s.AvgTotal *= weightCoef * weightCoef
	return s
}
