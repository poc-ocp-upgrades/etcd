package raft

import (
	"errors"
	pb "github.com/coreos/etcd/raft/raftpb"
)

var ErrStepLocalMsg = errors.New("raft: cannot step raft local message")
var ErrStepPeerNotFound = errors.New("raft: cannot step as peer not found")

type RawNode struct {
	raft       *raft
	prevSoftSt *SoftState
	prevHardSt pb.HardState
}

func (rn *RawNode) newReady() Ready {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newReady(rn.raft, rn.prevSoftSt, rn.prevHardSt)
}
func (rn *RawNode) commitReady(rd Ready) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rd.SoftState != nil {
		rn.prevSoftSt = rd.SoftState
	}
	if !IsEmptyHardState(rd.HardState) {
		rn.prevHardSt = rd.HardState
	}
	if rn.prevHardSt.Commit != 0 {
		rn.raft.raftLog.appliedTo(rn.prevHardSt.Commit)
	}
	if len(rd.Entries) > 0 {
		e := rd.Entries[len(rd.Entries)-1]
		rn.raft.raftLog.stableTo(e.Index, e.Term)
	}
	if !IsEmptySnap(rd.Snapshot) {
		rn.raft.raftLog.stableSnapTo(rd.Snapshot.Metadata.Index)
	}
	if len(rd.ReadStates) != 0 {
		rn.raft.readStates = nil
	}
}
func NewRawNode(config *Config, peers []Peer) (*RawNode, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if config.ID == 0 {
		panic("config.ID must not be zero")
	}
	r := newRaft(config)
	rn := &RawNode{raft: r}
	lastIndex, err := config.Storage.LastIndex()
	if err != nil {
		panic(err)
	}
	if lastIndex == 0 {
		r.becomeFollower(1, None)
		ents := make([]pb.Entry, len(peers))
		for i, peer := range peers {
			cc := pb.ConfChange{Type: pb.ConfChangeAddNode, NodeID: peer.ID, Context: peer.Context}
			data, err := cc.Marshal()
			if err != nil {
				panic("unexpected marshal error")
			}
			ents[i] = pb.Entry{Type: pb.EntryConfChange, Term: 1, Index: uint64(i + 1), Data: data}
		}
		r.raftLog.append(ents...)
		r.raftLog.committed = uint64(len(ents))
		for _, peer := range peers {
			r.addNode(peer.ID)
		}
	}
	rn.prevSoftSt = r.softState()
	if lastIndex == 0 {
		rn.prevHardSt = emptyState
	} else {
		rn.prevHardSt = r.hardState()
	}
	return rn, nil
}
func (rn *RawNode) Tick() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.raft.tick()
}
func (rn *RawNode) TickQuiesced() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.raft.electionElapsed++
}
func (rn *RawNode) Campaign() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rn.raft.Step(pb.Message{Type: pb.MsgHup})
}
func (rn *RawNode) Propose(data []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rn.raft.Step(pb.Message{Type: pb.MsgProp, From: rn.raft.id, Entries: []pb.Entry{{Data: data}}})
}
func (rn *RawNode) ProposeConfChange(cc pb.ConfChange) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := cc.Marshal()
	if err != nil {
		return err
	}
	return rn.raft.Step(pb.Message{Type: pb.MsgProp, Entries: []pb.Entry{{Type: pb.EntryConfChange, Data: data}}})
}
func (rn *RawNode) ApplyConfChange(cc pb.ConfChange) *pb.ConfState {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if cc.NodeID == None {
		rn.raft.resetPendingConf()
		return &pb.ConfState{Nodes: rn.raft.nodes()}
	}
	switch cc.Type {
	case pb.ConfChangeAddNode:
		rn.raft.addNode(cc.NodeID)
	case pb.ConfChangeAddLearnerNode:
		rn.raft.addLearner(cc.NodeID)
	case pb.ConfChangeRemoveNode:
		rn.raft.removeNode(cc.NodeID)
	case pb.ConfChangeUpdateNode:
		rn.raft.resetPendingConf()
	default:
		panic("unexpected conf type")
	}
	return &pb.ConfState{Nodes: rn.raft.nodes()}
}
func (rn *RawNode) Step(m pb.Message) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if IsLocalMsg(m.Type) {
		return ErrStepLocalMsg
	}
	if pr := rn.raft.getProgress(m.From); pr != nil || !IsResponseMsg(m.Type) {
		return rn.raft.Step(m)
	}
	return ErrStepPeerNotFound
}
func (rn *RawNode) Ready() Ready {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rd := rn.newReady()
	rn.raft.msgs = nil
	return rd
}
func (rn *RawNode) HasReady() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := rn.raft
	if !r.softState().equal(rn.prevSoftSt) {
		return true
	}
	if hardSt := r.hardState(); !IsEmptyHardState(hardSt) && !isHardStateEqual(hardSt, rn.prevHardSt) {
		return true
	}
	if r.raftLog.unstable.snapshot != nil && !IsEmptySnap(*r.raftLog.unstable.snapshot) {
		return true
	}
	if len(r.msgs) > 0 || len(r.raftLog.unstableEntries()) > 0 || r.raftLog.hasNextEnts() {
		return true
	}
	if len(r.readStates) != 0 {
		return true
	}
	return false
}
func (rn *RawNode) Advance(rd Ready) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rn.commitReady(rd)
}
func (rn *RawNode) Status() *Status {
	_logClusterCodePath()
	defer _logClusterCodePath()
	status := getStatus(rn.raft)
	return &status
}
func (rn *RawNode) ReportUnreachable(id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = rn.raft.Step(pb.Message{Type: pb.MsgUnreachable, From: id})
}
func (rn *RawNode) ReportSnapshot(id uint64, status SnapshotStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rej := status == SnapshotFailure
	_ = rn.raft.Step(pb.Message{Type: pb.MsgSnapStatus, From: id, Reject: rej})
}
func (rn *RawNode) TransferLeader(transferee uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = rn.raft.Step(pb.Message{Type: pb.MsgTransferLeader, From: transferee})
}
func (rn *RawNode) ReadIndex(rctx []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = rn.raft.Step(pb.Message{Type: pb.MsgReadIndex, Entries: []pb.Entry{{Data: rctx}}})
}
