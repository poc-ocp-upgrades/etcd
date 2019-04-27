package raft

import (
	"bytes"
	"reflect"
	"testing"
	"github.com/coreos/etcd/raft/raftpb"
)

func TestRawNodeStep(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, msgn := range raftpb.MessageType_name {
		s := NewMemoryStorage()
		rawNode, err := NewRawNode(newTestConfig(1, nil, 10, 1, s), []Peer{{ID: 1}})
		if err != nil {
			t.Fatal(err)
		}
		msgt := raftpb.MessageType(i)
		err = rawNode.Step(raftpb.Message{Type: msgt})
		if IsLocalMsg(msgt) {
			if err != ErrStepLocalMsg {
				t.Errorf("%d: step should ignore %s", msgt, msgn)
			}
		}
	}
}
func TestRawNodeProposeAndConfChange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := NewMemoryStorage()
	var err error
	rawNode, err := NewRawNode(newTestConfig(1, nil, 10, 1, s), []Peer{{ID: 1}})
	if err != nil {
		t.Fatal(err)
	}
	rd := rawNode.Ready()
	s.Append(rd.Entries)
	rawNode.Advance(rd)
	rawNode.Campaign()
	proposed := false
	var (
		lastIndex	uint64
		ccdata		[]byte
	)
	for {
		rd = rawNode.Ready()
		s.Append(rd.Entries)
		if !proposed && rd.SoftState.Lead == rawNode.raft.id {
			rawNode.Propose([]byte("somedata"))
			cc := raftpb.ConfChange{Type: raftpb.ConfChangeAddNode, NodeID: 1}
			ccdata, err = cc.Marshal()
			if err != nil {
				t.Fatal(err)
			}
			rawNode.ProposeConfChange(cc)
			proposed = true
		}
		rawNode.Advance(rd)
		lastIndex, err = s.LastIndex()
		if err != nil {
			t.Fatal(err)
		}
		if lastIndex >= 4 {
			break
		}
	}
	entries, err := s.Entries(lastIndex-1, lastIndex+1, noLimit)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 2 {
		t.Fatalf("len(entries) = %d, want %d", len(entries), 2)
	}
	if !bytes.Equal(entries[0].Data, []byte("somedata")) {
		t.Errorf("entries[0].Data = %v, want %v", entries[0].Data, []byte("somedata"))
	}
	if entries[1].Type != raftpb.EntryConfChange {
		t.Fatalf("type = %v, want %v", entries[1].Type, raftpb.EntryConfChange)
	}
	if !bytes.Equal(entries[1].Data, ccdata) {
		t.Errorf("data = %v, want %v", entries[1].Data, ccdata)
	}
}
func TestRawNodeProposeAddDuplicateNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := NewMemoryStorage()
	rawNode, err := NewRawNode(newTestConfig(1, nil, 10, 1, s), []Peer{{ID: 1}})
	if err != nil {
		t.Fatal(err)
	}
	rd := rawNode.Ready()
	s.Append(rd.Entries)
	rawNode.Advance(rd)
	rawNode.Campaign()
	for {
		rd = rawNode.Ready()
		s.Append(rd.Entries)
		if rd.SoftState.Lead == rawNode.raft.id {
			rawNode.Advance(rd)
			break
		}
		rawNode.Advance(rd)
	}
	proposeConfChangeAndApply := func(cc raftpb.ConfChange) {
		rawNode.ProposeConfChange(cc)
		rd = rawNode.Ready()
		s.Append(rd.Entries)
		for _, entry := range rd.CommittedEntries {
			if entry.Type == raftpb.EntryConfChange {
				var cc raftpb.ConfChange
				cc.Unmarshal(entry.Data)
				rawNode.ApplyConfChange(cc)
			}
		}
		rawNode.Advance(rd)
	}
	cc1 := raftpb.ConfChange{Type: raftpb.ConfChangeAddNode, NodeID: 1}
	ccdata1, err := cc1.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	proposeConfChangeAndApply(cc1)
	proposeConfChangeAndApply(cc1)
	cc2 := raftpb.ConfChange{Type: raftpb.ConfChangeAddNode, NodeID: 2}
	ccdata2, err := cc2.Marshal()
	if err != nil {
		t.Fatal(err)
	}
	proposeConfChangeAndApply(cc2)
	lastIndex, err := s.LastIndex()
	if err != nil {
		t.Fatal(err)
	}
	entries, err := s.Entries(lastIndex-2, lastIndex+1, noLimit)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 3 {
		t.Fatalf("len(entries) = %d, want %d", len(entries), 3)
	}
	if !bytes.Equal(entries[0].Data, ccdata1) {
		t.Errorf("entries[0].Data = %v, want %v", entries[0].Data, ccdata1)
	}
	if !bytes.Equal(entries[2].Data, ccdata2) {
		t.Errorf("entries[2].Data = %v, want %v", entries[2].Data, ccdata2)
	}
}
func TestRawNodeReadIndex(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	msgs := []raftpb.Message{}
	appendStep := func(r *raft, m raftpb.Message) {
		msgs = append(msgs, m)
	}
	wrs := []ReadState{{Index: uint64(1), RequestCtx: []byte("somedata")}}
	s := NewMemoryStorage()
	c := newTestConfig(1, nil, 10, 1, s)
	rawNode, err := NewRawNode(c, []Peer{{ID: 1}})
	if err != nil {
		t.Fatal(err)
	}
	rawNode.raft.readStates = wrs
	hasReady := rawNode.HasReady()
	if !hasReady {
		t.Errorf("HasReady() returns %t, want %t", hasReady, true)
	}
	rd := rawNode.Ready()
	if !reflect.DeepEqual(rd.ReadStates, wrs) {
		t.Errorf("ReadStates = %d, want %d", rd.ReadStates, wrs)
	}
	s.Append(rd.Entries)
	rawNode.Advance(rd)
	if rawNode.raft.readStates != nil {
		t.Errorf("readStates = %v, want %v", rawNode.raft.readStates, nil)
	}
	wrequestCtx := []byte("somedata2")
	rawNode.Campaign()
	for {
		rd = rawNode.Ready()
		s.Append(rd.Entries)
		if rd.SoftState.Lead == rawNode.raft.id {
			rawNode.Advance(rd)
			rawNode.raft.step = appendStep
			rawNode.ReadIndex(wrequestCtx)
			break
		}
		rawNode.Advance(rd)
	}
	if len(msgs) != 1 {
		t.Fatalf("len(msgs) = %d, want %d", len(msgs), 1)
	}
	if msgs[0].Type != raftpb.MsgReadIndex {
		t.Errorf("msg type = %d, want %d", msgs[0].Type, raftpb.MsgReadIndex)
	}
	if !bytes.Equal(msgs[0].Entries[0].Data, wrequestCtx) {
		t.Errorf("data = %v, want %v", msgs[0].Entries[0].Data, wrequestCtx)
	}
}
func TestRawNodeStart(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := raftpb.ConfChange{Type: raftpb.ConfChangeAddNode, NodeID: 1}
	ccdata, err := cc.Marshal()
	if err != nil {
		t.Fatalf("unexpected marshal error: %v", err)
	}
	wants := []Ready{{HardState: raftpb.HardState{Term: 1, Commit: 1, Vote: 0}, Entries: []raftpb.Entry{{Type: raftpb.EntryConfChange, Term: 1, Index: 1, Data: ccdata}}, CommittedEntries: []raftpb.Entry{{Type: raftpb.EntryConfChange, Term: 1, Index: 1, Data: ccdata}}, MustSync: true}, {HardState: raftpb.HardState{Term: 2, Commit: 3, Vote: 1}, Entries: []raftpb.Entry{{Term: 2, Index: 3, Data: []byte("foo")}}, CommittedEntries: []raftpb.Entry{{Term: 2, Index: 3, Data: []byte("foo")}}, MustSync: true}}
	storage := NewMemoryStorage()
	rawNode, err := NewRawNode(newTestConfig(1, nil, 10, 1, storage), []Peer{{ID: 1}})
	if err != nil {
		t.Fatal(err)
	}
	rd := rawNode.Ready()
	t.Logf("rd %v", rd)
	if !reflect.DeepEqual(rd, wants[0]) {
		t.Fatalf("#%d: g = %+v,\n             w   %+v", 1, rd, wants[0])
	} else {
		storage.Append(rd.Entries)
		rawNode.Advance(rd)
	}
	storage.Append(rd.Entries)
	rawNode.Advance(rd)
	rawNode.Campaign()
	rd = rawNode.Ready()
	storage.Append(rd.Entries)
	rawNode.Advance(rd)
	rawNode.Propose([]byte("foo"))
	if rd = rawNode.Ready(); !reflect.DeepEqual(rd, wants[1]) {
		t.Errorf("#%d: g = %+v,\n             w   %+v", 2, rd, wants[1])
	} else {
		storage.Append(rd.Entries)
		rawNode.Advance(rd)
	}
	if rawNode.HasReady() {
		t.Errorf("unexpected Ready: %+v", rawNode.Ready())
	}
}
func TestRawNodeRestart(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	entries := []raftpb.Entry{{Term: 1, Index: 1}, {Term: 1, Index: 2, Data: []byte("foo")}}
	st := raftpb.HardState{Term: 1, Commit: 1}
	want := Ready{HardState: emptyState, CommittedEntries: entries[:st.Commit], MustSync: true}
	storage := NewMemoryStorage()
	storage.SetHardState(st)
	storage.Append(entries)
	rawNode, err := NewRawNode(newTestConfig(1, nil, 10, 1, storage), nil)
	if err != nil {
		t.Fatal(err)
	}
	rd := rawNode.Ready()
	if !reflect.DeepEqual(rd, want) {
		t.Errorf("g = %+v,\n             w   %+v", rd, want)
	}
	rawNode.Advance(rd)
	if rawNode.HasReady() {
		t.Errorf("unexpected Ready: %+v", rawNode.Ready())
	}
}
func TestRawNodeRestartFromSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	snap := raftpb.Snapshot{Metadata: raftpb.SnapshotMetadata{ConfState: raftpb.ConfState{Nodes: []uint64{1, 2}}, Index: 2, Term: 1}}
	entries := []raftpb.Entry{{Term: 1, Index: 3, Data: []byte("foo")}}
	st := raftpb.HardState{Term: 1, Commit: 3}
	want := Ready{HardState: emptyState, CommittedEntries: entries, MustSync: true}
	s := NewMemoryStorage()
	s.SetHardState(st)
	s.ApplySnapshot(snap)
	s.Append(entries)
	rawNode, err := NewRawNode(newTestConfig(1, nil, 10, 1, s), nil)
	if err != nil {
		t.Fatal(err)
	}
	if rd := rawNode.Ready(); !reflect.DeepEqual(rd, want) {
		t.Errorf("g = %+v,\n             w   %+v", rd, want)
	} else {
		rawNode.Advance(rd)
	}
	if rawNode.HasReady() {
		t.Errorf("unexpected Ready: %+v", rawNode.HasReady())
	}
}
func TestRawNodeStatus(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storage := NewMemoryStorage()
	rawNode, err := NewRawNode(newTestConfig(1, nil, 10, 1, storage), []Peer{{ID: 1}})
	if err != nil {
		t.Fatal(err)
	}
	status := rawNode.Status()
	if status == nil {
		t.Errorf("expected status struct, got nil")
	}
}
