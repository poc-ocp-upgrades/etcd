package rafthttp

import (
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/xiang90/probing"
)

var (
	proberInterval			= ConnReadTimeout - time.Second
	statusMonitoringInterval	= 30 * time.Second
	statusErrorInterval		= 5 * time.Second
)

const (
	RoundTripperNameRaftMessage	= "ROUND_TRIPPER_RAFT_MESSAGE"
	RoundTripperNameSnapshot	= "ROUND_TRIPPER_SNAPSHOT"
)

func addPeerToProber(p probing.Prober, id string, us []string, roundTripperName string, rttSecProm *prometheus.HistogramVec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hus := make([]string, len(us))
	for i := range us {
		hus[i] = us[i] + ProbingPrefix
	}
	p.AddHTTP(id, proberInterval, hus)
	s, err := p.Status(id)
	if err != nil {
		plog.Errorf("failed to add peer %s into prober", id)
	} else {
		go monitorProbingStatus(s, id, roundTripperName, rttSecProm)
	}
}
func monitorProbingStatus(s probing.Status, id string, roundTripperName string, rttSecProm *prometheus.HistogramVec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	interval := statusErrorInterval
	for {
		select {
		case <-time.After(interval):
			if !s.Health() {
				plog.Warningf("health check for peer %s could not connect: %v (prober %q)", id, s.Err(), roundTripperName)
				interval = statusErrorInterval
			} else {
				interval = statusMonitoringInterval
			}
			if s.ClockDiff() > time.Second {
				plog.Warningf("the clock difference against peer %s is too high [%v > %v] (prober %q)", id, s.ClockDiff(), time.Second, roundTripperName)
			}
			rttSecProm.WithLabelValues(id).Observe(s.SRTT().Seconds())
		case <-s.StopNotify():
			return
		}
	}
}
