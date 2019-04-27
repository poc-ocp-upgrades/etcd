package etcdserver

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/mock/mockstorage"
	"github.com/coreos/etcd/pkg/pbutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/rafthttp"
)

func TestGetIDs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	addcc := &raftpb.ConfChange{Type: raftpb.ConfChangeAddNode, NodeID: 2}
	addEntry := raftpb.Entry{Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(addcc)}
	removecc := &raftpb.ConfChange{Type: raftpb.ConfChangeRemoveNode, NodeID: 2}
	removeEntry := raftpb.Entry{Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(removecc)}
	normalEntry := raftpb.Entry{Type: raftpb.EntryNormal}
	updatecc := &raftpb.ConfChange{Type: raftpb.ConfChangeUpdateNode, NodeID: 2}
	updateEntry := raftpb.Entry{Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(updatecc)}
	tests := []struct {
		confState	*raftpb.ConfState
		ents		[]raftpb.Entry
		widSet		[]uint64
	}{{nil, []raftpb.Entry{}, []uint64{}}, {&raftpb.ConfState{Nodes: []uint64{1}}, []raftpb.Entry{}, []uint64{1}}, {&raftpb.ConfState{Nodes: []uint64{1}}, []raftpb.Entry{addEntry}, []uint64{1, 2}}, {&raftpb.ConfState{Nodes: []uint64{1}}, []raftpb.Entry{addEntry, removeEntry}, []uint64{1}}, {&raftpb.ConfState{Nodes: []uint64{1}}, []raftpb.Entry{addEntry, normalEntry}, []uint64{1, 2}}, {&raftpb.ConfState{Nodes: []uint64{1}}, []raftpb.Entry{addEntry, normalEntry, updateEntry}, []uint64{1, 2}}, {&raftpb.ConfState{Nodes: []uint64{1}}, []raftpb.Entry{addEntry, removeEntry, normalEntry}, []uint64{1}}}
	for i, tt := range tests {
		var snap raftpb.Snapshot
		if tt.confState != nil {
			snap.Metadata.ConfState = *tt.confState
		}
		idSet := getIDs(&snap, tt.ents)
		if !reflect.DeepEqual(idSet, tt.widSet) {
			t.Errorf("#%d: idset = %#v, want %#v", i, idSet, tt.widSet)
		}
	}
}
func TestCreateConfigChangeEnts(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := membership.Member{ID: types.ID(1), RaftAttributes: membership.RaftAttributes{PeerURLs: []string{"http://localhost:2380"}}}
	ctx, err := json.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}
	addcc1 := &raftpb.ConfChange{Type: raftpb.ConfChangeAddNode, NodeID: 1, Context: ctx}
	removecc2 := &raftpb.ConfChange{Type: raftpb.ConfChangeRemoveNode, NodeID: 2}
	removecc3 := &raftpb.ConfChange{Type: raftpb.ConfChangeRemoveNode, NodeID: 3}
	tests := []struct {
		ids		[]uint64
		self		uint64
		term, index	uint64
		wents		[]raftpb.Entry
	}{{[]uint64{1}, 1, 1, 1, []raftpb.Entry{}}, {[]uint64{1, 2}, 1, 1, 1, []raftpb.Entry{{Term: 1, Index: 2, Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(removecc2)}}}, {[]uint64{1, 2}, 1, 2, 2, []raftpb.Entry{{Term: 2, Index: 3, Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(removecc2)}}}, {[]uint64{1, 2, 3}, 1, 2, 2, []raftpb.Entry{{Term: 2, Index: 3, Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(removecc2)}, {Term: 2, Index: 4, Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(removecc3)}}}, {[]uint64{2, 3}, 2, 2, 2, []raftpb.Entry{{Term: 2, Index: 3, Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(removecc3)}}}, {[]uint64{2, 3}, 1, 2, 2, []raftpb.Entry{{Term: 2, Index: 3, Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(removecc2)}, {Term: 2, Index: 4, Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(removecc3)}, {Term: 2, Index: 5, Type: raftpb.EntryConfChange, Data: pbutil.MustMarshal(addcc1)}}}}
	for i, tt := range tests {
		gents := createConfigChangeEnts(tt.ids, tt.self, tt.term, tt.index)
		if !reflect.DeepEqual(gents, tt.wents) {
			t.Errorf("#%d: ents = %v, want %v", i, gents, tt.wents)
		}
	}
}
func TestStopRaftWhenWaitingForApplyDone(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := newNopReadyNode()
	r := newRaftNode(raftNodeConfig{Node: n, storage: mockstorage.NewStorageRecorder(""), raftStorage: raft.NewMemoryStorage(), transport: rafthttp.NewNopTransporter()})
	srv := &EtcdServer{r: *r}
	srv.r.start(nil)
	n.readyc <- raft.Ready{}
	select {
	case <-srv.r.applyc:
	case <-time.After(time.Second):
		t.Fatalf("failed to receive apply struct")
	}
	srv.r.stopped <- struct{}{}
	select {
	case <-srv.r.done:
	case <-time.After(time.Second):
		t.Fatalf("failed to stop raft loop")
	}
}
func TestConfgChangeBlocksApply(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := newNopReadyNode()
	r := newRaftNode(raftNodeConfig{Node: n, storage: mockstorage.NewStorageRecorder(""), raftStorage: raft.NewMemoryStorage(), transport: rafthttp.NewNopTransporter()})
	srv := &EtcdServer{r: *r}
	srv.r.start(&raftReadyHandler{updateLeadership: func(bool) {
	}})
	defer srv.r.Stop()
	n.readyc <- raft.Ready{SoftState: &raft.SoftState{RaftState: raft.StateFollower}, CommittedEntries: []raftpb.Entry{{Type: raftpb.EntryConfChange}}}
	ap := <-srv.r.applyc
	continueC := make(chan struct{})
	go func() {
		n.readyc <- raft.Ready{}
		<-srv.r.applyc
		close(continueC)
	}()
	select {
	case <-continueC:
		t.Fatalf("unexpected execution: raft routine should block waiting for apply")
	case <-time.After(time.Second):
	}
	<-ap.notifyc
	select {
	case <-continueC:
	case <-time.After(time.Second):
		t.Fatalf("unexpected blocking on execution")
	}
}
