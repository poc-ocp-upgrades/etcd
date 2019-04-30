package tester

import (
	"fmt"
	"sort"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	caseTotal		= make(map[string]int)
	caseTotalCounter	= prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "funcational_tester", Name: "case_total", Help: "Total number of finished test cases"}, []string{"desc"})
	caseFailedTotalCounter	= prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "funcational_tester", Name: "case_failed_total", Help: "Total number of failed test cases"}, []string{"desc"})
	roundTotalCounter	= prometheus.NewCounter(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "funcational_tester", Name: "round_total", Help: "Total number of finished test rounds."})
	roundFailedTotalCounter	= prometheus.NewCounter(prometheus.CounterOpts{Namespace: "etcd", Subsystem: "funcational_tester", Name: "round_failed_total", Help: "Total number of failed test rounds."})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prometheus.MustRegister(caseTotalCounter)
	prometheus.MustRegister(caseFailedTotalCounter)
	prometheus.MustRegister(roundTotalCounter)
	prometheus.MustRegister(roundFailedTotalCounter)
}
func printReport() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rows := make([]string, 0, len(caseTotal))
	for k, v := range caseTotal {
		rows = append(rows, fmt.Sprintf("%s: %d", k, v))
	}
	sort.Strings(rows)
	println()
	for _, row := range rows {
		fmt.Println(row)
	}
	println()
}
