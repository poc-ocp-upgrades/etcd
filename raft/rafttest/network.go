package rafttest

import (
	"math/rand"
	"sync"
	"time"
	"github.com/coreos/etcd/raft/raftpb"
)

type iface interface {
	send(m raftpb.Message)
	recv() chan raftpb.Message
	disconnect()
	connect()
}
type network interface {
	drop(from, to uint64, rate float64)
	delay(from, to uint64, d time.Duration, rate float64)
	disconnect(id uint64)
	connect(id uint64)
	heal()
}
type raftNetwork struct {
	mu		sync.Mutex
	disconnected	map[uint64]bool
	dropmap		map[conn]float64
	delaymap	map[conn]delay
	recvQueues	map[uint64]chan raftpb.Message
}
type conn struct{ from, to uint64 }
type delay struct {
	d	time.Duration
	rate	float64
}

func newRaftNetwork(nodes ...uint64) *raftNetwork {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pn := &raftNetwork{recvQueues: make(map[uint64]chan raftpb.Message), dropmap: make(map[conn]float64), delaymap: make(map[conn]delay), disconnected: make(map[uint64]bool)}
	for _, n := range nodes {
		pn.recvQueues[n] = make(chan raftpb.Message, 1024)
	}
	return pn
}
func (rn *raftNetwork) nodeNetwork(id uint64) iface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &nodeNetwork{id: id, raftNetwork: rn}
}
func (rn *raftNetwork) send(m raftpb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.mu.Lock()
	to := rn.recvQueues[m.To]
	if rn.disconnected[m.To] {
		to = nil
	}
	drop := rn.dropmap[conn{m.From, m.To}]
	dl := rn.delaymap[conn{m.From, m.To}]
	rn.mu.Unlock()
	if to == nil {
		return
	}
	if drop != 0 && rand.Float64() < drop {
		return
	}
	if dl.d != 0 && rand.Float64() < dl.rate {
		rd := rand.Int63n(int64(dl.d))
		time.Sleep(time.Duration(rd))
	}
	b, err := m.Marshal()
	if err != nil {
		panic(err)
	}
	var cm raftpb.Message
	err = cm.Unmarshal(b)
	if err != nil {
		panic(err)
	}
	select {
	case to <- cm:
	default:
	}
}
func (rn *raftNetwork) recvFrom(from uint64) chan raftpb.Message {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.mu.Lock()
	fromc := rn.recvQueues[from]
	if rn.disconnected[from] {
		fromc = nil
	}
	rn.mu.Unlock()
	return fromc
}
func (rn *raftNetwork) drop(from, to uint64, rate float64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.dropmap[conn{from, to}] = rate
}
func (rn *raftNetwork) delay(from, to uint64, d time.Duration, rate float64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.delaymap[conn{from, to}] = delay{d, rate}
}
func (rn *raftNetwork) heal() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.dropmap = make(map[conn]float64)
	rn.delaymap = make(map[conn]delay)
}
func (rn *raftNetwork) disconnect(id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.disconnected[id] = true
}
func (rn *raftNetwork) connect(id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.mu.Lock()
	defer rn.mu.Unlock()
	rn.disconnected[id] = false
}

type nodeNetwork struct {
	id	uint64
	*raftNetwork
}

func (nt *nodeNetwork) connect() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt.raftNetwork.connect(nt.id)
}
func (nt *nodeNetwork) disconnect() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt.raftNetwork.disconnect(nt.id)
}
func (nt *nodeNetwork) send(m raftpb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nt.raftNetwork.send(m)
}
func (nt *nodeNetwork) recv() chan raftpb.Message {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nt.recvFrom(nt.id)
}
