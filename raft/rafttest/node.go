package rafttest

import (
	"context"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"log"
	"sync"
	"time"
)

type node struct {
	raft.Node
	id      uint64
	iface   iface
	stopc   chan struct{}
	pausec  chan bool
	storage *raft.MemoryStorage
	mu      sync.Mutex
	state   raftpb.HardState
}

func startNode(id uint64, peers []raft.Peer, iface iface) *node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	st := raft.NewMemoryStorage()
	c := &raft.Config{ID: id, ElectionTick: 10, HeartbeatTick: 1, Storage: st, MaxSizePerMsg: 1024 * 1024, MaxInflightMsgs: 256}
	rn := raft.StartNode(c, peers)
	n := &node{Node: rn, id: id, storage: st, iface: iface, pausec: make(chan bool)}
	n.start()
	return n
}
func (n *node) start() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.stopc = make(chan struct{})
	ticker := time.Tick(5 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker:
				n.Tick()
			case rd := <-n.Ready():
				if !raft.IsEmptyHardState(rd.HardState) {
					n.mu.Lock()
					n.state = rd.HardState
					n.mu.Unlock()
					n.storage.SetHardState(n.state)
				}
				n.storage.Append(rd.Entries)
				time.Sleep(time.Millisecond)
				for _, m := range rd.Messages {
					n.iface.send(m)
				}
				n.Advance()
			case m := <-n.iface.recv():
				go n.Step(context.TODO(), m)
			case <-n.stopc:
				n.Stop()
				log.Printf("raft.%d: stop", n.id)
				n.Node = nil
				close(n.stopc)
				return
			case p := <-n.pausec:
				recvms := make([]raftpb.Message, 0)
				for p {
					select {
					case m := <-n.iface.recv():
						recvms = append(recvms, m)
					case p = <-n.pausec:
					}
				}
				for _, m := range recvms {
					n.Step(context.TODO(), m)
				}
			}
		}
	}()
}
func (n *node) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.iface.disconnect()
	n.stopc <- struct{}{}
	<-n.stopc
}
func (n *node) restart() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	<-n.stopc
	c := &raft.Config{ID: n.id, ElectionTick: 10, HeartbeatTick: 1, Storage: n.storage, MaxSizePerMsg: 1024 * 1024, MaxInflightMsgs: 256}
	n.Node = raft.RestartNode(c)
	n.start()
	n.iface.connect()
}
func (n *node) pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.pausec <- true
}
func (n *node) resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n.pausec <- false
}
