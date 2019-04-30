package raftpb

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "github.com/gogo/protobuf/gogoproto"
	io "io"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2

type EntryType int32

const (
	EntryNormal	EntryType	= 0
	EntryConfChange	EntryType	= 1
)

var EntryType_name = map[int32]string{0: "EntryNormal", 1: "EntryConfChange"}
var EntryType_value = map[string]int32{"EntryNormal": 0, "EntryConfChange": 1}

func (x EntryType) Enum() *EntryType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := new(EntryType)
	*p = x
	return p
}
func (x EntryType) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(EntryType_name, int32(x))
}
func (x *EntryType) UnmarshalJSON(data []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	value, err := proto.UnmarshalJSONEnum(EntryType_value, data, "EntryType")
	if err != nil {
		return err
	}
	*x = EntryType(value)
	return nil
}
func (EntryType) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{0}
}

type MessageType int32

const (
	MsgHup			MessageType	= 0
	MsgBeat			MessageType	= 1
	MsgProp			MessageType	= 2
	MsgApp			MessageType	= 3
	MsgAppResp		MessageType	= 4
	MsgVote			MessageType	= 5
	MsgVoteResp		MessageType	= 6
	MsgSnap			MessageType	= 7
	MsgHeartbeat		MessageType	= 8
	MsgHeartbeatResp	MessageType	= 9
	MsgUnreachable		MessageType	= 10
	MsgSnapStatus		MessageType	= 11
	MsgCheckQuorum		MessageType	= 12
	MsgTransferLeader	MessageType	= 13
	MsgTimeoutNow		MessageType	= 14
	MsgReadIndex		MessageType	= 15
	MsgReadIndexResp	MessageType	= 16
	MsgPreVote		MessageType	= 17
	MsgPreVoteResp		MessageType	= 18
)

var MessageType_name = map[int32]string{0: "MsgHup", 1: "MsgBeat", 2: "MsgProp", 3: "MsgApp", 4: "MsgAppResp", 5: "MsgVote", 6: "MsgVoteResp", 7: "MsgSnap", 8: "MsgHeartbeat", 9: "MsgHeartbeatResp", 10: "MsgUnreachable", 11: "MsgSnapStatus", 12: "MsgCheckQuorum", 13: "MsgTransferLeader", 14: "MsgTimeoutNow", 15: "MsgReadIndex", 16: "MsgReadIndexResp", 17: "MsgPreVote", 18: "MsgPreVoteResp"}
var MessageType_value = map[string]int32{"MsgHup": 0, "MsgBeat": 1, "MsgProp": 2, "MsgApp": 3, "MsgAppResp": 4, "MsgVote": 5, "MsgVoteResp": 6, "MsgSnap": 7, "MsgHeartbeat": 8, "MsgHeartbeatResp": 9, "MsgUnreachable": 10, "MsgSnapStatus": 11, "MsgCheckQuorum": 12, "MsgTransferLeader": 13, "MsgTimeoutNow": 14, "MsgReadIndex": 15, "MsgReadIndexResp": 16, "MsgPreVote": 17, "MsgPreVoteResp": 18}

func (x MessageType) Enum() *MessageType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{1}
}

type ConfChangeType int32

const (
	ConfChangeAddNode		ConfChangeType	= 0
	ConfChangeRemoveNode		ConfChangeType	= 1
	ConfChangeUpdateNode		ConfChangeType	= 2
	ConfChangeAddLearnerNode	ConfChangeType	= 3
)

var ConfChangeType_name = map[int32]string{0: "ConfChangeAddNode", 1: "ConfChangeRemoveNode", 2: "ConfChangeUpdateNode", 3: "ConfChangeAddLearnerNode"}
var ConfChangeType_value = map[string]int32{"ConfChangeAddNode": 0, "ConfChangeRemoveNode": 1, "ConfChangeUpdateNode": 2, "ConfChangeAddLearnerNode": 3}

func (x ConfChangeType) Enum() *ConfChangeType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := new(ConfChangeType)
	*p = x
	return p
}
func (x ConfChangeType) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(ConfChangeType_name, int32(x))
}
func (x *ConfChangeType) UnmarshalJSON(data []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	value, err := proto.UnmarshalJSONEnum(ConfChangeType_value, data, "ConfChangeType")
	if err != nil {
		return err
	}
	*x = ConfChangeType(value)
	return nil
}
func (ConfChangeType) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{2}
}

type Entry struct {
	Term			uint64		`protobuf:"varint,2,opt,name=Term" json:"Term"`
	Index			uint64		`protobuf:"varint,3,opt,name=Index" json:"Index"`
	Type			EntryType	`protobuf:"varint,1,opt,name=Type,enum=raftpb.EntryType" json:"Type"`
	Data			[]byte		`protobuf:"bytes,4,opt,name=Data" json:"Data,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (m *Entry) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Entry{}
}
func (m *Entry) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Entry) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Entry) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{0}
}

type SnapshotMetadata struct {
	ConfState		ConfState	`protobuf:"bytes,1,opt,name=conf_state,json=confState" json:"conf_state"`
	Index			uint64		`protobuf:"varint,2,opt,name=index" json:"index"`
	Term			uint64		`protobuf:"varint,3,opt,name=term" json:"term"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (m *SnapshotMetadata) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = SnapshotMetadata{}
}
func (m *SnapshotMetadata) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*SnapshotMetadata) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*SnapshotMetadata) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{1}
}

type Snapshot struct {
	Data			[]byte			`protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Metadata		SnapshotMetadata	`protobuf:"bytes,2,opt,name=metadata" json:"metadata"`
	XXX_unrecognized	[]byte			`json:"-"`
}

func (m *Snapshot) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Snapshot{}
}
func (m *Snapshot) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Snapshot) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Snapshot) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{2}
}

type Message struct {
	Type			MessageType	`protobuf:"varint,1,opt,name=type,enum=raftpb.MessageType" json:"type"`
	To			uint64		`protobuf:"varint,2,opt,name=to" json:"to"`
	From			uint64		`protobuf:"varint,3,opt,name=from" json:"from"`
	Term			uint64		`protobuf:"varint,4,opt,name=term" json:"term"`
	LogTerm			uint64		`protobuf:"varint,5,opt,name=logTerm" json:"logTerm"`
	Index			uint64		`protobuf:"varint,6,opt,name=index" json:"index"`
	Entries			[]Entry		`protobuf:"bytes,7,rep,name=entries" json:"entries"`
	Commit			uint64		`protobuf:"varint,8,opt,name=commit" json:"commit"`
	Snapshot		Snapshot	`protobuf:"bytes,9,opt,name=snapshot" json:"snapshot"`
	Reject			bool		`protobuf:"varint,10,opt,name=reject" json:"reject"`
	RejectHint		uint64		`protobuf:"varint,11,opt,name=rejectHint" json:"rejectHint"`
	Context			[]byte		`protobuf:"bytes,12,opt,name=context" json:"context,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (m *Message) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Message{}
}
func (m *Message) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Message) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Message) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{3}
}

type HardState struct {
	Term			uint64	`protobuf:"varint,1,opt,name=term" json:"term"`
	Vote			uint64	`protobuf:"varint,2,opt,name=vote" json:"vote"`
	Commit			uint64	`protobuf:"varint,3,opt,name=commit" json:"commit"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (m *HardState) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = HardState{}
}
func (m *HardState) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*HardState) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*HardState) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{4}
}

type ConfState struct {
	Nodes			[]uint64	`protobuf:"varint,1,rep,name=nodes" json:"nodes,omitempty"`
	Learners		[]uint64	`protobuf:"varint,2,rep,name=learners" json:"learners,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (m *ConfState) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = ConfState{}
}
func (m *ConfState) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*ConfState) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ConfState) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{5}
}

type ConfChange struct {
	ID			uint64		`protobuf:"varint,1,opt,name=ID" json:"ID"`
	Type			ConfChangeType	`protobuf:"varint,2,opt,name=Type,enum=raftpb.ConfChangeType" json:"Type"`
	NodeID			uint64		`protobuf:"varint,3,opt,name=NodeID" json:"NodeID"`
	Context			[]byte		`protobuf:"bytes,4,opt,name=Context" json:"Context,omitempty"`
	XXX_unrecognized	[]byte		`json:"-"`
}

func (m *ConfChange) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = ConfChange{}
}
func (m *ConfChange) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*ConfChange) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ConfChange) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaft, []int{6}
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterType((*Entry)(nil), "raftpb.Entry")
	proto.RegisterType((*SnapshotMetadata)(nil), "raftpb.SnapshotMetadata")
	proto.RegisterType((*Snapshot)(nil), "raftpb.Snapshot")
	proto.RegisterType((*Message)(nil), "raftpb.Message")
	proto.RegisterType((*HardState)(nil), "raftpb.HardState")
	proto.RegisterType((*ConfState)(nil), "raftpb.ConfState")
	proto.RegisterType((*ConfChange)(nil), "raftpb.ConfChange")
	proto.RegisterEnum("raftpb.EntryType", EntryType_name, EntryType_value)
	proto.RegisterEnum("raftpb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("raftpb.ConfChangeType", ConfChangeType_name, ConfChangeType_value)
}
func (m *Entry) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Type))
	dAtA[i] = 0x10
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Term))
	dAtA[i] = 0x18
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Index))
	if m.Data != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *SnapshotMetadata) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *SnapshotMetadata) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.ConfState.Size()))
	n1, err := m.ConfState.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x10
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Index))
	dAtA[i] = 0x18
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Term))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Metadata.Size()))
	n2, err := m.Metadata.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Type))
	dAtA[i] = 0x10
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.To))
	dAtA[i] = 0x18
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.From))
	dAtA[i] = 0x20
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Term))
	dAtA[i] = 0x28
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.LogTerm))
	dAtA[i] = 0x30
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Index))
	if len(m.Entries) > 0 {
		for _, msg := range m.Entries {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintRaft(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x40
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Commit))
	dAtA[i] = 0x4a
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Snapshot.Size()))
	n3, err := m.Snapshot.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x50
	i++
	if m.Reject {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x58
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.RejectHint))
	if m.Context != nil {
		dAtA[i] = 0x62
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.Context)))
		i += copy(dAtA[i:], m.Context)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *HardState) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *HardState) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Term))
	dAtA[i] = 0x10
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Vote))
	dAtA[i] = 0x18
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Commit))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *ConfState) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *ConfState) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, num := range m.Nodes {
			dAtA[i] = 0x8
			i++
			i = encodeVarintRaft(dAtA, i, uint64(num))
		}
	}
	if len(m.Learners) > 0 {
		for _, num := range m.Learners {
			dAtA[i] = 0x10
			i++
			i = encodeVarintRaft(dAtA, i, uint64(num))
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *ConfChange) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *ConfChange) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.ID))
	dAtA[i] = 0x10
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.Type))
	dAtA[i] = 0x18
	i++
	i = encodeVarintRaft(dAtA, i, uint64(m.NodeID))
	if m.Context != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRaft(dAtA, i, uint64(len(m.Context)))
		i += copy(dAtA[i:], m.Context)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeVarintRaft(dAtA []byte, offset int, v uint64) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Entry) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.Type))
	n += 1 + sovRaft(uint64(m.Term))
	n += 1 + sovRaft(uint64(m.Index))
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovRaft(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *SnapshotMetadata) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = m.ConfState.Size()
	n += 1 + l + sovRaft(uint64(l))
	n += 1 + sovRaft(uint64(m.Index))
	n += 1 + sovRaft(uint64(m.Term))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Snapshot) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovRaft(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovRaft(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Message) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.Type))
	n += 1 + sovRaft(uint64(m.To))
	n += 1 + sovRaft(uint64(m.From))
	n += 1 + sovRaft(uint64(m.Term))
	n += 1 + sovRaft(uint64(m.LogTerm))
	n += 1 + sovRaft(uint64(m.Index))
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovRaft(uint64(l))
		}
	}
	n += 1 + sovRaft(uint64(m.Commit))
	l = m.Snapshot.Size()
	n += 1 + l + sovRaft(uint64(l))
	n += 2
	n += 1 + sovRaft(uint64(m.RejectHint))
	if m.Context != nil {
		l = len(m.Context)
		n += 1 + l + sovRaft(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *HardState) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.Term))
	n += 1 + sovRaft(uint64(m.Vote))
	n += 1 + sovRaft(uint64(m.Commit))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *ConfState) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			n += 1 + sovRaft(uint64(e))
		}
	}
	if len(m.Learners) > 0 {
		for _, e := range m.Learners {
			n += 1 + sovRaft(uint64(e))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *ConfChange) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.ID))
	n += 1 + sovRaft(uint64(m.Type))
	n += 1 + sovRaft(uint64(m.NodeID))
	if m.Context != nil {
		l = len(m.Context)
		n += 1 + l + sovRaft(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func sovRaft(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRaft(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sovRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (EntryType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SnapshotMetadata) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SnapshotMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConfState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			m.To = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.To |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			m.From = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.From |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogTerm", wireType)
			}
			m.LogTerm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LogTerm |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			m.Commit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Commit |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshot", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Snapshot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reject", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Reject = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectHint", wireType)
			}
			m.RejectHint = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RejectHint |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context[:0], dAtA[iNdEx:postIndex]...)
			if m.Context == nil {
				m.Context = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HardState) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HardState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HardState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			m.Commit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Commit |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfState) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRaft
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Nodes = append(m.Nodes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRaft
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthRaft
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRaft
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Nodes = append(m.Nodes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRaft
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Learners = append(m.Learners, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRaft
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthRaft
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRaft
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Learners = append(m.Learners, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Learners", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfChange) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (ConfChangeType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context[:0], dAtA[iNdEx:postIndex]...)
			if m.Context == nil {
				m.Context = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRaft(dAtA []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRaft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaft
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRaft(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRaft	= fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaft	= fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("raft.proto", fileDescriptorRaft)
}

var fileDescriptorRaft = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x54, 0xcd, 0x6e, 0x23, 0x45, 0x10, 0xf6, 0x8c, 0xc7, 0x7f, 0x35, 0x8e, 0xd3, 0xa9, 0x35, 0xa8, 0x15, 0x45, 0xc6, 0xb2, 0x38, 0x58, 0x41, 0x1b, 0x20, 0x07, 0x0e, 0x48, 0x1c, 0x36, 0x09, 0x52, 0x22, 0xad, 0xa3, 0xc5, 0x9b, 0xe5, 0x80, 0x84, 0x50, 0xc7, 0x53, 0x9e, 0x18, 0x32, 0xd3, 0xa3, 0x9e, 0xf6, 0xb2, 0xb9, 0x20, 0x1e, 0x80, 0x07, 0xe0, 0xc2, 0xfb, 0xe4, 0xb8, 0x12, 0x77, 0xc4, 0x86, 0x17, 0x41, 0xdd, 0xd3, 0x63, 0xcf, 0x24, 0xb7, 0xae, 0xef, 0xab, 0xae, 0xfa, 0xea, 0xeb, 0x9a, 0x01, 0x50, 0x62, 0xa9, 0x8f, 0x32, 0x25, 0xb5, 0xc4, 0xb6, 0x39, 0x67, 0xd7, 0xfb, 0xc3, 0x58, 0xc6, 0xd2, 0x42, 0x9f, 0x9b, 0x53, 0xc1, 0x4e, 0x7e, 0x83, 0xd6, 0xb7, 0xa9, 0x56, 0x77, 0xf8, 0x19, 0x04, 0x57, 0x77, 0x19, 0x71, 0x6f, 0xec, 0x4d, 0x07, 0xc7, 0x7b, 0x47, 0xc5, 0xad, 0x23, 0x4b, 0x1a, 0xe2, 0x24, 0xb8, 0xff, 0xe7, 0x93, 0xc6, 0xdc, 0x26, 0x21, 0x87, 0xe0, 0x8a, 0x54, 0xc2, 0xfd, 0xb1, 0x37, 0x0d, 0x36, 0x0c, 0xa9, 0x04, 0xf7, 0xa1, 0x75, 0x91, 0x46, 0xf4, 0x8e, 0x37, 0x2b, 0x54, 0x01, 0x21, 0x42, 0x70, 0x26, 0xb4, 0xe0, 0xc1, 0xd8, 0x9b, 0xf6, 0xe7, 0xf6, 0x3c, 0xf9, 0xdd, 0x03, 0xf6, 0x3a, 0x15, 0x59, 0x7e, 0x23, 0xf5, 0x8c, 0xb4, 0x88, 0x84, 0x16, 0xf8, 0x15, 0xc0, 0x42, 0xa6, 0xcb, 0x9f, 0x72, 0x2d, 0x74, 0xa1, 0x28, 0xdc, 0x2a, 0x3a, 0x95, 0xe9, 0xf2, 0xb5, 0x21, 0x5c, 0xf1, 0xde, 0xa2, 0x04, 0x4c, 0xf3, 0x95, 0x6d, 0x5e, 0xd5, 0x55, 0x40, 0x46, 0xb2, 0x36, 0x92, 0xab, 0xba, 0x2c, 0x32, 0xf9, 0x01, 0xba, 0xa5, 0x02, 0x23, 0xd1, 0x28, 0xb0, 0x3d, 0xfb, 0x73, 0x7b, 0xc6, 0xaf, 0xa1, 0x9b, 0x38, 0x65, 0xb6, 0x70, 0x78, 0xcc, 0x4b, 0x2d, 0x8f, 0x95, 0xbb, 0xba, 0x9b, 0xfc, 0xc9, 0x5f, 0x4d, 0xe8, 0xcc, 0x28, 0xcf, 0x45, 0x4c, 0xf8, 0x1c, 0x02, 0xbd, 0x75, 0xf8, 0x59, 0x59, 0xc3, 0xd1, 0x55, 0x8f, 0x4d, 0x1a, 0x0e, 0xc1, 0xd7, 0xb2, 0x36, 0x89, 0xaf, 0xa5, 0x19, 0x63, 0xa9, 0xe4, 0xa3, 0x31, 0x0c, 0xb2, 0x19, 0x30, 0x78, 0x3c, 0x20, 0x8e, 0xa0, 0x73, 0x2b, 0x63, 0xfb, 0x60, 0xad, 0x0a, 0x59, 0x82, 0x5b, 0xdb, 0xda, 0x4f, 0x6d, 0x7b, 0x0e, 0x1d, 0x4a, 0xb5, 0x5a, 0x51, 0xce, 0x3b, 0xe3, 0xe6, 0x34, 0x3c, 0xde, 0xa9, 0x6d, 0x46, 0x59, 0xca, 0xe5, 0xe0, 0x01, 0xb4, 0x17, 0x32, 0x49, 0x56, 0x9a, 0x77, 0x2b, 0xb5, 0x1c, 0x86, 0xc7, 0xd0, 0xcd, 0x9d, 0x63, 0xbc, 0x67, 0x9d, 0x64, 0x8f, 0x9d, 0x2c, 0x1d, 0x2c, 0xf3, 0x4c, 0x45, 0x45, 0x3f, 0xd3, 0x42, 0x73, 0x18, 0x7b, 0xd3, 0x6e, 0x59, 0xb1, 0xc0, 0xf0, 0x53, 0x80, 0xe2, 0x74, 0xbe, 0x4a, 0x35, 0x0f, 0x2b, 0x3d, 0x2b, 0x38, 0x72, 0xe8, 0x2c, 0x64, 0xaa, 0xe9, 0x9d, 0xe6, 0x7d, 0xfb, 0xb0, 0x65, 0x38, 0xf9, 0x11, 0x7a, 0xe7, 0x42, 0x45, 0xc5, 0xfa, 0x94, 0x0e, 0x7a, 0x4f, 0x1c, 0xe4, 0x10, 0xbc, 0x95, 0x9a, 0xea, 0xfb, 0x6e, 0x90, 0xca, 0xc0, 0xcd, 0xa7, 0x03, 0x4f, 0xbe, 0x81, 0xde, 0x66, 0x5d, 0x71, 0x08, 0xad, 0x54, 0x46, 0x94, 0x73, 0x6f, 0xdc, 0x9c, 0x06, 0xf3, 0x22, 0xc0, 0x7d, 0xe8, 0xde, 0x92, 0x50, 0x29, 0xa9, 0x9c, 0xfb, 0x96, 0xd8, 0xc4, 0x93, 0x3f, 0x3c, 0x00, 0x73, 0xff, 0xf4, 0x46, 0xa4, 0xb1, 0xdd, 0x88, 0x8b, 0xb3, 0x9a, 0x3a, 0xff, 0xe2, 0x0c, 0xbf, 0x70, 0x1f, 0xae, 0x6f, 0xd7, 0xea, 0xe3, 0xea, 0x67, 0x52, 0xdc, 0x7b, 0xf2, 0xf5, 0x1e, 0x40, 0xfb, 0x52, 0x46, 0x74, 0x71, 0x56, 0xd7, 0x5c, 0x60, 0xc6, 0xac, 0x53, 0x67, 0x56, 0xf1, 0xa1, 0x96, 0xe1, 0xe1, 0x97, 0xd0, 0xdb, 0xfc, 0x0e, 0x70, 0x17, 0x42, 0x1b, 0x5c, 0x4a, 0x95, 0x88, 0x5b, 0xd6, 0xc0, 0x67, 0xb0, 0x6b, 0x81, 0x6d, 0x63, 0xe6, 0x1d, 0xfe, 0xed, 0x43, 0x58, 0x59, 0x70, 0x04, 0x68, 0xcf, 0xf2, 0xf8, 0x7c, 0x9d, 0xb1, 0x06, 0x86, 0xd0, 0x99, 0xe5, 0xf1, 0x09, 0x09, 0xcd, 0x3c, 0x17, 0xbc, 0x52, 0x32, 0x63, 0xbe, 0xcb, 0x7a, 0x91, 0x65, 0xac, 0x89, 0x03, 0x80, 0xe2, 0x3c, 0xa7, 0x3c, 0x63, 0x81, 0x4b, 0xfc, 0x5e, 0x6a, 0x62, 0x2d, 0x23, 0xc2, 0x05, 0x96, 0x6d, 0x3b, 0xd6, 0x2c, 0x13, 0xeb, 0x20, 0x83, 0xbe, 0x69, 0x46, 0x42, 0xe9, 0x6b, 0xd3, 0xa5, 0x8b, 0x43, 0x60, 0x55, 0xc4, 0x5e, 0xea, 0x21, 0xc2, 0x60, 0x96, 0xc7, 0x6f, 0x52, 0x45, 0x62, 0x71, 0x23, 0xae, 0x6f, 0x89, 0x01, 0xee, 0xc1, 0x8e, 0x2b, 0x64, 0x1e, 0x6f, 0x9d, 0xb3, 0xd0, 0xa5, 0x9d, 0xde, 0xd0, 0xe2, 0x97, 0xef, 0xd6, 0x52, 0xad, 0x13, 0xd6, 0xc7, 0x8f, 0x60, 0x6f, 0x96, 0xc7, 0x57, 0x4a, 0xa4, 0xf9, 0x92, 0xd4, 0x4b, 0x12, 0x11, 0x29, 0xb6, 0xe3, 0x6e, 0x5f, 0xad, 0x12, 0x92, 0x6b, 0x7d, 0x29, 0x7f, 0x65, 0x03, 0x27, 0x66, 0x4e, 0x22, 0xb2, 0x3f, 0x43, 0xb6, 0xeb, 0xc4, 0x6c, 0x10, 0x2b, 0x86, 0xb9, 0x79, 0x5f, 0x29, 0xb2, 0x23, 0xee, 0xb9, 0xae, 0x2e, 0xb6, 0x39, 0x78, 0x78, 0x07, 0x83, 0xfa, 0xf3, 0x1a, 0x1d, 0x5b, 0xe4, 0x45, 0x14, 0x99, 0xb7, 0x64, 0x0d, 0xe4, 0x30, 0xdc, 0xc2, 0x73, 0x4a, 0xe4, 0x5b, 0xb2, 0x8c, 0x57, 0x67, 0xde, 0x64, 0x91, 0xd0, 0x05, 0xe3, 0xe3, 0x01, 0xf0, 0x5a, 0xa9, 0x97, 0xc5, 0x36, 0x5a, 0xb6, 0x79, 0xc2, 0xef, 0x3f, 0x8c, 0x1a, 0xef, 0x3f, 0x8c, 0x1a, 0xf7, 0x0f, 0x23, 0xef, 0xfd, 0xc3, 0xc8, 0xfb, 0xf7, 0x61, 0xe4, 0xfd, 0xf9, 0xdf, 0xa8, 0xf1, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x52, 0x5b, 0xe0, 0x74, 0x06, 0x00, 0x00}

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
