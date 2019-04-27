package stats

import (
	"encoding/json"
	"log"
	"sync"
	"time"
	"github.com/coreos/etcd/raft"
)

type ServerStats struct {
	serverStats
	sync.Mutex
}

func NewServerStats(name, id string) *ServerStats {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss := &ServerStats{serverStats: serverStats{Name: name, ID: id}}
	now := time.Now()
	ss.StartTime = now
	ss.LeaderInfo.StartTime = now
	ss.sendRateQueue = &statsQueue{back: -1}
	ss.recvRateQueue = &statsQueue{back: -1}
	return ss
}

type serverStats struct {
	Name		string		`json:"name"`
	ID		string		`json:"id"`
	State		raft.StateType	`json:"state"`
	StartTime	time.Time	`json:"startTime"`
	LeaderInfo	struct {
		Name		string		`json:"leader"`
		Uptime		string		`json:"uptime"`
		StartTime	time.Time	`json:"startTime"`
	}	`json:"leaderInfo"`
	RecvAppendRequestCnt	uint64	`json:"recvAppendRequestCnt,"`
	RecvingPkgRate		float64	`json:"recvPkgRate,omitempty"`
	RecvingBandwidthRate	float64	`json:"recvBandwidthRate,omitempty"`
	SendAppendRequestCnt	uint64	`json:"sendAppendRequestCnt"`
	SendingPkgRate		float64	`json:"sendPkgRate,omitempty"`
	SendingBandwidthRate	float64	`json:"sendBandwidthRate,omitempty"`
	sendRateQueue		*statsQueue
	recvRateQueue		*statsQueue
}

func (ss *ServerStats) JSON() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss.Lock()
	stats := ss.serverStats
	stats.SendingPkgRate, stats.SendingBandwidthRate = stats.sendRateQueue.Rate()
	stats.RecvingPkgRate, stats.RecvingBandwidthRate = stats.recvRateQueue.Rate()
	stats.LeaderInfo.Uptime = time.Since(stats.LeaderInfo.StartTime).String()
	ss.Unlock()
	b, err := json.Marshal(stats)
	if err != nil {
		log.Printf("stats: error marshalling server stats: %v", err)
	}
	return b
}
func (ss *ServerStats) RecvAppendReq(leader string, reqSize int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss.Lock()
	defer ss.Unlock()
	now := time.Now()
	ss.State = raft.StateFollower
	if leader != ss.LeaderInfo.Name {
		ss.LeaderInfo.Name = leader
		ss.LeaderInfo.StartTime = now
	}
	ss.recvRateQueue.Insert(&RequestStats{SendingTime: now, Size: reqSize})
	ss.RecvAppendRequestCnt++
}
func (ss *ServerStats) SendAppendReq(reqSize int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss.Lock()
	defer ss.Unlock()
	ss.becomeLeader()
	ss.sendRateQueue.Insert(&RequestStats{SendingTime: time.Now(), Size: reqSize})
	ss.SendAppendRequestCnt++
}
func (ss *ServerStats) BecomeLeader() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss.Lock()
	defer ss.Unlock()
	ss.becomeLeader()
}
func (ss *ServerStats) becomeLeader() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ss.State != raft.StateLeader {
		ss.State = raft.StateLeader
		ss.LeaderInfo.Name = ss.ID
		ss.LeaderInfo.StartTime = time.Now()
	}
}
