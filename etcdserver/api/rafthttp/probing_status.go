package rafthttp

import (
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/xiang90/probing"
	"go.uber.org/zap"
)

const (
	RoundTripperNameRaftMessage	= "ROUND_TRIPPER_RAFT_MESSAGE"
	RoundTripperNameSnapshot	= "ROUND_TRIPPER_SNAPSHOT"
)

var (
	proberInterval			= ConnReadTimeout - time.Second
	statusMonitoringInterval	= 30 * time.Second
	statusErrorInterval		= 5 * time.Second
)

func addPeerToProber(lg *zap.Logger, p probing.Prober, id string, us []string, roundTripperName string, rttSecProm *prometheus.HistogramVec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hus := make([]string, len(us))
	for i := range us {
		hus[i] = us[i] + ProbingPrefix
	}
	p.AddHTTP(id, proberInterval, hus)
	s, err := p.Status(id)
	if err != nil {
		if lg != nil {
			lg.Warn("failed to add peer into prober", zap.String("remote-peer-id", id))
		} else {
			plog.Errorf("failed to add peer %s into prober", id)
		}
		return
	}
	go monitorProbingStatus(lg, s, id, roundTripperName, rttSecProm)
}
func monitorProbingStatus(lg *zap.Logger, s probing.Status, id string, roundTripperName string, rttSecProm *prometheus.HistogramVec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	interval := statusErrorInterval
	for {
		select {
		case <-time.After(interval):
			if !s.Health() {
				if lg != nil {
					lg.Warn("prober detected unhealthy status", zap.String("round-tripper-name", roundTripperName), zap.String("remote-peer-id", id), zap.Duration("rtt", s.SRTT()), zap.Error(s.Err()))
				} else {
					plog.Warningf("health check for peer %s could not connect: %v", id, s.Err())
				}
				interval = statusErrorInterval
			} else {
				interval = statusMonitoringInterval
			}
			if s.ClockDiff() > time.Second {
				if lg != nil {
					lg.Warn("prober found high clock drift", zap.String("round-tripper-name", roundTripperName), zap.String("remote-peer-id", id), zap.Duration("clock-drift", s.SRTT()), zap.Duration("rtt", s.ClockDiff()), zap.Error(s.Err()))
				} else {
					plog.Warningf("the clock difference against peer %s is too high [%v > %v]", id, s.ClockDiff(), time.Second)
				}
			}
			rttSecProm.WithLabelValues(id).Observe(s.SRTT().Seconds())
		case <-s.StopNotify():
			return
		}
	}
}
