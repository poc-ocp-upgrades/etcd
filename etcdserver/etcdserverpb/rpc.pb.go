package etcdserverpb

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "github.com/gogo/protobuf/gogoproto"
	mvccpb "go.etcd.io/etcd/mvcc/mvccpb"
	authpb "go.etcd.io/etcd/auth/authpb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	io "io"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AlarmType int32

const (
	AlarmType_NONE		AlarmType	= 0
	AlarmType_NOSPACE	AlarmType	= 1
	AlarmType_CORRUPT	AlarmType	= 2
)

var AlarmType_name = map[int32]string{0: "NONE", 1: "NOSPACE", 2: "CORRUPT"}
var AlarmType_value = map[string]int32{"NONE": 0, "NOSPACE": 1, "CORRUPT": 2}

func (x AlarmType) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(AlarmType_name, int32(x))
}
func (AlarmType) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{0}
}

type RangeRequest_SortOrder int32

const (
	RangeRequest_NONE	RangeRequest_SortOrder	= 0
	RangeRequest_ASCEND	RangeRequest_SortOrder	= 1
	RangeRequest_DESCEND	RangeRequest_SortOrder	= 2
)

var RangeRequest_SortOrder_name = map[int32]string{0: "NONE", 1: "ASCEND", 2: "DESCEND"}
var RangeRequest_SortOrder_value = map[string]int32{"NONE": 0, "ASCEND": 1, "DESCEND": 2}

func (x RangeRequest_SortOrder) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(RangeRequest_SortOrder_name, int32(x))
}
func (RangeRequest_SortOrder) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{1, 0}
}

type RangeRequest_SortTarget int32

const (
	RangeRequest_KEY	RangeRequest_SortTarget	= 0
	RangeRequest_VERSION	RangeRequest_SortTarget	= 1
	RangeRequest_CREATE	RangeRequest_SortTarget	= 2
	RangeRequest_MOD	RangeRequest_SortTarget	= 3
	RangeRequest_VALUE	RangeRequest_SortTarget	= 4
)

var RangeRequest_SortTarget_name = map[int32]string{0: "KEY", 1: "VERSION", 2: "CREATE", 3: "MOD", 4: "VALUE"}
var RangeRequest_SortTarget_value = map[string]int32{"KEY": 0, "VERSION": 1, "CREATE": 2, "MOD": 3, "VALUE": 4}

func (x RangeRequest_SortTarget) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(RangeRequest_SortTarget_name, int32(x))
}
func (RangeRequest_SortTarget) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{1, 1}
}

type Compare_CompareResult int32

const (
	Compare_EQUAL		Compare_CompareResult	= 0
	Compare_GREATER		Compare_CompareResult	= 1
	Compare_LESS		Compare_CompareResult	= 2
	Compare_NOT_EQUAL	Compare_CompareResult	= 3
)

var Compare_CompareResult_name = map[int32]string{0: "EQUAL", 1: "GREATER", 2: "LESS", 3: "NOT_EQUAL"}
var Compare_CompareResult_value = map[string]int32{"EQUAL": 0, "GREATER": 1, "LESS": 2, "NOT_EQUAL": 3}

func (x Compare_CompareResult) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(Compare_CompareResult_name, int32(x))
}
func (Compare_CompareResult) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{9, 0}
}

type Compare_CompareTarget int32

const (
	Compare_VERSION	Compare_CompareTarget	= 0
	Compare_CREATE	Compare_CompareTarget	= 1
	Compare_MOD	Compare_CompareTarget	= 2
	Compare_VALUE	Compare_CompareTarget	= 3
	Compare_LEASE	Compare_CompareTarget	= 4
)

var Compare_CompareTarget_name = map[int32]string{0: "VERSION", 1: "CREATE", 2: "MOD", 3: "VALUE", 4: "LEASE"}
var Compare_CompareTarget_value = map[string]int32{"VERSION": 0, "CREATE": 1, "MOD": 2, "VALUE": 3, "LEASE": 4}

func (x Compare_CompareTarget) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(Compare_CompareTarget_name, int32(x))
}
func (Compare_CompareTarget) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{9, 1}
}

type WatchCreateRequest_FilterType int32

const (
	WatchCreateRequest_NOPUT	WatchCreateRequest_FilterType	= 0
	WatchCreateRequest_NODELETE	WatchCreateRequest_FilterType	= 1
)

var WatchCreateRequest_FilterType_name = map[int32]string{0: "NOPUT", 1: "NODELETE"}
var WatchCreateRequest_FilterType_value = map[string]int32{"NOPUT": 0, "NODELETE": 1}

func (x WatchCreateRequest_FilterType) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(WatchCreateRequest_FilterType_name, int32(x))
}
func (WatchCreateRequest_FilterType) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{21, 0}
}

type AlarmRequest_AlarmAction int32

const (
	AlarmRequest_GET	AlarmRequest_AlarmAction	= 0
	AlarmRequest_ACTIVATE	AlarmRequest_AlarmAction	= 1
	AlarmRequest_DEACTIVATE	AlarmRequest_AlarmAction	= 2
)

var AlarmRequest_AlarmAction_name = map[int32]string{0: "GET", 1: "ACTIVATE", 2: "DEACTIVATE"}
var AlarmRequest_AlarmAction_value = map[string]int32{"GET": 0, "ACTIVATE": 1, "DEACTIVATE": 2}

func (x AlarmRequest_AlarmAction) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(AlarmRequest_AlarmAction_name, int32(x))
}
func (AlarmRequest_AlarmAction) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{52, 0}
}

type ResponseHeader struct {
	ClusterId	uint64	`protobuf:"varint,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	MemberId	uint64	`protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	Revision	int64	`protobuf:"varint,3,opt,name=revision,proto3" json:"revision,omitempty"`
	RaftTerm	uint64	`protobuf:"varint,4,opt,name=raft_term,json=raftTerm,proto3" json:"raft_term,omitempty"`
}

func (m *ResponseHeader) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = ResponseHeader{}
}
func (m *ResponseHeader) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*ResponseHeader) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ResponseHeader) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{0}
}
func (m *ResponseHeader) GetClusterId() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ClusterId
	}
	return 0
}
func (m *ResponseHeader) GetMemberId() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.MemberId
	}
	return 0
}
func (m *ResponseHeader) GetRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Revision
	}
	return 0
}
func (m *ResponseHeader) GetRaftTerm() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RaftTerm
	}
	return 0
}

type RangeRequest struct {
	Key			[]byte			`protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	RangeEnd		[]byte			`protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	Limit			int64			`protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Revision		int64			`protobuf:"varint,4,opt,name=revision,proto3" json:"revision,omitempty"`
	SortOrder		RangeRequest_SortOrder	`protobuf:"varint,5,opt,name=sort_order,json=sortOrder,proto3,enum=etcdserverpb.RangeRequest_SortOrder" json:"sort_order,omitempty"`
	SortTarget		RangeRequest_SortTarget	`protobuf:"varint,6,opt,name=sort_target,json=sortTarget,proto3,enum=etcdserverpb.RangeRequest_SortTarget" json:"sort_target,omitempty"`
	Serializable		bool			`protobuf:"varint,7,opt,name=serializable,proto3" json:"serializable,omitempty"`
	KeysOnly		bool			`protobuf:"varint,8,opt,name=keys_only,json=keysOnly,proto3" json:"keys_only,omitempty"`
	CountOnly		bool			`protobuf:"varint,9,opt,name=count_only,json=countOnly,proto3" json:"count_only,omitempty"`
	MinModRevision		int64			`protobuf:"varint,10,opt,name=min_mod_revision,json=minModRevision,proto3" json:"min_mod_revision,omitempty"`
	MaxModRevision		int64			`protobuf:"varint,11,opt,name=max_mod_revision,json=maxModRevision,proto3" json:"max_mod_revision,omitempty"`
	MinCreateRevision	int64			`protobuf:"varint,12,opt,name=min_create_revision,json=minCreateRevision,proto3" json:"min_create_revision,omitempty"`
	MaxCreateRevision	int64			`protobuf:"varint,13,opt,name=max_create_revision,json=maxCreateRevision,proto3" json:"max_create_revision,omitempty"`
}

func (m *RangeRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = RangeRequest{}
}
func (m *RangeRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*RangeRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*RangeRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{1}
}
func (m *RangeRequest) GetKey() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Key
	}
	return nil
}
func (m *RangeRequest) GetRangeEnd() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RangeEnd
	}
	return nil
}
func (m *RangeRequest) GetLimit() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Limit
	}
	return 0
}
func (m *RangeRequest) GetRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Revision
	}
	return 0
}
func (m *RangeRequest) GetSortOrder() RangeRequest_SortOrder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.SortOrder
	}
	return RangeRequest_NONE
}
func (m *RangeRequest) GetSortTarget() RangeRequest_SortTarget {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.SortTarget
	}
	return RangeRequest_KEY
}
func (m *RangeRequest) GetSerializable() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Serializable
	}
	return false
}
func (m *RangeRequest) GetKeysOnly() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.KeysOnly
	}
	return false
}
func (m *RangeRequest) GetCountOnly() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.CountOnly
	}
	return false
}
func (m *RangeRequest) GetMinModRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.MinModRevision
	}
	return 0
}
func (m *RangeRequest) GetMaxModRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.MaxModRevision
	}
	return 0
}
func (m *RangeRequest) GetMinCreateRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.MinCreateRevision
	}
	return 0
}
func (m *RangeRequest) GetMaxCreateRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.MaxCreateRevision
	}
	return 0
}

type RangeResponse struct {
	Header	*ResponseHeader		`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Kvs	[]*mvccpb.KeyValue	`protobuf:"bytes,2,rep,name=kvs" json:"kvs,omitempty"`
	More	bool			`protobuf:"varint,3,opt,name=more,proto3" json:"more,omitempty"`
	Count	int64			`protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *RangeResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = RangeResponse{}
}
func (m *RangeResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*RangeResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*RangeResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{2}
}
func (m *RangeResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *RangeResponse) GetKvs() []*mvccpb.KeyValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Kvs
	}
	return nil
}
func (m *RangeResponse) GetMore() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.More
	}
	return false
}
func (m *RangeResponse) GetCount() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Count
	}
	return 0
}

type PutRequest struct {
	Key		[]byte	`protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value		[]byte	`protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Lease		int64	`protobuf:"varint,3,opt,name=lease,proto3" json:"lease,omitempty"`
	PrevKv		bool	`protobuf:"varint,4,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	IgnoreValue	bool	`protobuf:"varint,5,opt,name=ignore_value,json=ignoreValue,proto3" json:"ignore_value,omitempty"`
	IgnoreLease	bool	`protobuf:"varint,6,opt,name=ignore_lease,json=ignoreLease,proto3" json:"ignore_lease,omitempty"`
}

func (m *PutRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = PutRequest{}
}
func (m *PutRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*PutRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*PutRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{3}
}
func (m *PutRequest) GetKey() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Key
	}
	return nil
}
func (m *PutRequest) GetValue() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Value
	}
	return nil
}
func (m *PutRequest) GetLease() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Lease
	}
	return 0
}
func (m *PutRequest) GetPrevKv() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.PrevKv
	}
	return false
}
func (m *PutRequest) GetIgnoreValue() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.IgnoreValue
	}
	return false
}
func (m *PutRequest) GetIgnoreLease() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.IgnoreLease
	}
	return false
}

type PutResponse struct {
	Header	*ResponseHeader		`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	PrevKv	*mvccpb.KeyValue	`protobuf:"bytes,2,opt,name=prev_kv,json=prevKv" json:"prev_kv,omitempty"`
}

func (m *PutResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = PutResponse{}
}
func (m *PutResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*PutResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*PutResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{4}
}
func (m *PutResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *PutResponse) GetPrevKv() *mvccpb.KeyValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.PrevKv
	}
	return nil
}

type DeleteRangeRequest struct {
	Key		[]byte	`protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	RangeEnd	[]byte	`protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	PrevKv		bool	`protobuf:"varint,3,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
}

func (m *DeleteRangeRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = DeleteRangeRequest{}
}
func (m *DeleteRangeRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*DeleteRangeRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*DeleteRangeRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{5}
}
func (m *DeleteRangeRequest) GetKey() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Key
	}
	return nil
}
func (m *DeleteRangeRequest) GetRangeEnd() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RangeEnd
	}
	return nil
}
func (m *DeleteRangeRequest) GetPrevKv() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.PrevKv
	}
	return false
}

type DeleteRangeResponse struct {
	Header	*ResponseHeader		`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Deleted	int64			`protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	PrevKvs	[]*mvccpb.KeyValue	`protobuf:"bytes,3,rep,name=prev_kvs,json=prevKvs" json:"prev_kvs,omitempty"`
}

func (m *DeleteRangeResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = DeleteRangeResponse{}
}
func (m *DeleteRangeResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*DeleteRangeResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*DeleteRangeResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{6}
}
func (m *DeleteRangeResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *DeleteRangeResponse) GetDeleted() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Deleted
	}
	return 0
}
func (m *DeleteRangeResponse) GetPrevKvs() []*mvccpb.KeyValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.PrevKvs
	}
	return nil
}

type RequestOp struct {
	Request isRequestOp_Request `protobuf_oneof:"request"`
}

func (m *RequestOp) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = RequestOp{}
}
func (m *RequestOp) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*RequestOp) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*RequestOp) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{7}
}

type isRequestOp_Request interface {
	isRequestOp_Request()
	MarshalTo([]byte) (int, error)
	Size() int
}
type RequestOp_RequestRange struct {
	RequestRange *RangeRequest `protobuf:"bytes,1,opt,name=request_range,json=requestRange,oneof"`
}
type RequestOp_RequestPut struct {
	RequestPut *PutRequest `protobuf:"bytes,2,opt,name=request_put,json=requestPut,oneof"`
}
type RequestOp_RequestDeleteRange struct {
	RequestDeleteRange *DeleteRangeRequest `protobuf:"bytes,3,opt,name=request_delete_range,json=requestDeleteRange,oneof"`
}
type RequestOp_RequestTxn struct {
	RequestTxn *TxnRequest `protobuf:"bytes,4,opt,name=request_txn,json=requestTxn,oneof"`
}

func (*RequestOp_RequestRange) isRequestOp_Request() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*RequestOp_RequestPut) isRequestOp_Request() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*RequestOp_RequestDeleteRange) isRequestOp_Request() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*RequestOp_RequestTxn) isRequestOp_Request() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (m *RequestOp) GetRequest() isRequestOp_Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Request
	}
	return nil
}
func (m *RequestOp) GetRequestRange() *RangeRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetRequest().(*RequestOp_RequestRange); ok {
		return x.RequestRange
	}
	return nil
}
func (m *RequestOp) GetRequestPut() *PutRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetRequest().(*RequestOp_RequestPut); ok {
		return x.RequestPut
	}
	return nil
}
func (m *RequestOp) GetRequestDeleteRange() *DeleteRangeRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetRequest().(*RequestOp_RequestDeleteRange); ok {
		return x.RequestDeleteRange
	}
	return nil
}
func (m *RequestOp) GetRequestTxn() *TxnRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetRequest().(*RequestOp_RequestTxn); ok {
		return x.RequestTxn
	}
	return nil
}
func (*RequestOp) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _RequestOp_OneofMarshaler, _RequestOp_OneofUnmarshaler, _RequestOp_OneofSizer, []interface{}{(*RequestOp_RequestRange)(nil), (*RequestOp_RequestPut)(nil), (*RequestOp_RequestDeleteRange)(nil), (*RequestOp_RequestTxn)(nil)}
}
func _RequestOp_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*RequestOp)
	switch x := m.Request.(type) {
	case *RequestOp_RequestRange:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestRange); err != nil {
			return err
		}
	case *RequestOp_RequestPut:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestPut); err != nil {
			return err
		}
	case *RequestOp_RequestDeleteRange:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestDeleteRange); err != nil {
			return err
		}
	case *RequestOp_RequestTxn:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestTxn); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RequestOp.Request has unexpected type %T", x)
	}
	return nil
}
func _RequestOp_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*RequestOp)
	switch tag {
	case 1:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RangeRequest)
		err := b.DecodeMessage(msg)
		m.Request = &RequestOp_RequestRange{msg}
		return true, err
	case 2:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PutRequest)
		err := b.DecodeMessage(msg)
		m.Request = &RequestOp_RequestPut{msg}
		return true, err
	case 3:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeleteRangeRequest)
		err := b.DecodeMessage(msg)
		m.Request = &RequestOp_RequestDeleteRange{msg}
		return true, err
	case 4:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TxnRequest)
		err := b.DecodeMessage(msg)
		m.Request = &RequestOp_RequestTxn{msg}
		return true, err
	default:
		return false, nil
	}
}
func _RequestOp_OneofSizer(msg proto.Message) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*RequestOp)
	switch x := m.Request.(type) {
	case *RequestOp_RequestRange:
		s := proto.Size(x.RequestRange)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RequestOp_RequestPut:
		s := proto.Size(x.RequestPut)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RequestOp_RequestDeleteRange:
		s := proto.Size(x.RequestDeleteRange)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *RequestOp_RequestTxn:
		s := proto.Size(x.RequestTxn)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ResponseOp struct {
	Response isResponseOp_Response `protobuf_oneof:"response"`
}

func (m *ResponseOp) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = ResponseOp{}
}
func (m *ResponseOp) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*ResponseOp) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ResponseOp) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{8}
}

type isResponseOp_Response interface {
	isResponseOp_Response()
	MarshalTo([]byte) (int, error)
	Size() int
}
type ResponseOp_ResponseRange struct {
	ResponseRange *RangeResponse `protobuf:"bytes,1,opt,name=response_range,json=responseRange,oneof"`
}
type ResponseOp_ResponsePut struct {
	ResponsePut *PutResponse `protobuf:"bytes,2,opt,name=response_put,json=responsePut,oneof"`
}
type ResponseOp_ResponseDeleteRange struct {
	ResponseDeleteRange *DeleteRangeResponse `protobuf:"bytes,3,opt,name=response_delete_range,json=responseDeleteRange,oneof"`
}
type ResponseOp_ResponseTxn struct {
	ResponseTxn *TxnResponse `protobuf:"bytes,4,opt,name=response_txn,json=responseTxn,oneof"`
}

func (*ResponseOp_ResponseRange) isResponseOp_Response() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ResponseOp_ResponsePut) isResponseOp_Response() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ResponseOp_ResponseDeleteRange) isResponseOp_Response() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ResponseOp_ResponseTxn) isResponseOp_Response() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (m *ResponseOp) GetResponse() isResponseOp_Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Response
	}
	return nil
}
func (m *ResponseOp) GetResponseRange() *RangeResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetResponse().(*ResponseOp_ResponseRange); ok {
		return x.ResponseRange
	}
	return nil
}
func (m *ResponseOp) GetResponsePut() *PutResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetResponse().(*ResponseOp_ResponsePut); ok {
		return x.ResponsePut
	}
	return nil
}
func (m *ResponseOp) GetResponseDeleteRange() *DeleteRangeResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetResponse().(*ResponseOp_ResponseDeleteRange); ok {
		return x.ResponseDeleteRange
	}
	return nil
}
func (m *ResponseOp) GetResponseTxn() *TxnResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetResponse().(*ResponseOp_ResponseTxn); ok {
		return x.ResponseTxn
	}
	return nil
}
func (*ResponseOp) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _ResponseOp_OneofMarshaler, _ResponseOp_OneofUnmarshaler, _ResponseOp_OneofSizer, []interface{}{(*ResponseOp_ResponseRange)(nil), (*ResponseOp_ResponsePut)(nil), (*ResponseOp_ResponseDeleteRange)(nil), (*ResponseOp_ResponseTxn)(nil)}
}
func _ResponseOp_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*ResponseOp)
	switch x := m.Response.(type) {
	case *ResponseOp_ResponseRange:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResponseRange); err != nil {
			return err
		}
	case *ResponseOp_ResponsePut:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResponsePut); err != nil {
			return err
		}
	case *ResponseOp_ResponseDeleteRange:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResponseDeleteRange); err != nil {
			return err
		}
	case *ResponseOp_ResponseTxn:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResponseTxn); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ResponseOp.Response has unexpected type %T", x)
	}
	return nil
}
func _ResponseOp_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*ResponseOp)
	switch tag {
	case 1:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RangeResponse)
		err := b.DecodeMessage(msg)
		m.Response = &ResponseOp_ResponseRange{msg}
		return true, err
	case 2:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PutResponse)
		err := b.DecodeMessage(msg)
		m.Response = &ResponseOp_ResponsePut{msg}
		return true, err
	case 3:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeleteRangeResponse)
		err := b.DecodeMessage(msg)
		m.Response = &ResponseOp_ResponseDeleteRange{msg}
		return true, err
	case 4:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TxnResponse)
		err := b.DecodeMessage(msg)
		m.Response = &ResponseOp_ResponseTxn{msg}
		return true, err
	default:
		return false, nil
	}
}
func _ResponseOp_OneofSizer(msg proto.Message) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*ResponseOp)
	switch x := m.Response.(type) {
	case *ResponseOp_ResponseRange:
		s := proto.Size(x.ResponseRange)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResponseOp_ResponsePut:
		s := proto.Size(x.ResponsePut)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResponseOp_ResponseDeleteRange:
		s := proto.Size(x.ResponseDeleteRange)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResponseOp_ResponseTxn:
		s := proto.Size(x.ResponseTxn)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Compare struct {
	Result		Compare_CompareResult	`protobuf:"varint,1,opt,name=result,proto3,enum=etcdserverpb.Compare_CompareResult" json:"result,omitempty"`
	Target		Compare_CompareTarget	`protobuf:"varint,2,opt,name=target,proto3,enum=etcdserverpb.Compare_CompareTarget" json:"target,omitempty"`
	Key		[]byte			`protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	TargetUnion	isCompare_TargetUnion	`protobuf_oneof:"target_union"`
	RangeEnd	[]byte			`protobuf:"bytes,64,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
}

func (m *Compare) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Compare{}
}
func (m *Compare) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Compare) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Compare) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{9}
}

type isCompare_TargetUnion interface {
	isCompare_TargetUnion()
	MarshalTo([]byte) (int, error)
	Size() int
}
type Compare_Version struct {
	Version int64 `protobuf:"varint,4,opt,name=version,proto3,oneof"`
}
type Compare_CreateRevision struct {
	CreateRevision int64 `protobuf:"varint,5,opt,name=create_revision,json=createRevision,proto3,oneof"`
}
type Compare_ModRevision struct {
	ModRevision int64 `protobuf:"varint,6,opt,name=mod_revision,json=modRevision,proto3,oneof"`
}
type Compare_Value struct {
	Value []byte `protobuf:"bytes,7,opt,name=value,proto3,oneof"`
}
type Compare_Lease struct {
	Lease int64 `protobuf:"varint,8,opt,name=lease,proto3,oneof"`
}

func (*Compare_Version) isCompare_TargetUnion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Compare_CreateRevision) isCompare_TargetUnion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Compare_ModRevision) isCompare_TargetUnion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Compare_Value) isCompare_TargetUnion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Compare_Lease) isCompare_TargetUnion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (m *Compare) GetTargetUnion() isCompare_TargetUnion {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.TargetUnion
	}
	return nil
}
func (m *Compare) GetResult() Compare_CompareResult {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Result
	}
	return Compare_EQUAL
}
func (m *Compare) GetTarget() Compare_CompareTarget {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Target
	}
	return Compare_VERSION
}
func (m *Compare) GetKey() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Key
	}
	return nil
}
func (m *Compare) GetVersion() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetTargetUnion().(*Compare_Version); ok {
		return x.Version
	}
	return 0
}
func (m *Compare) GetCreateRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetTargetUnion().(*Compare_CreateRevision); ok {
		return x.CreateRevision
	}
	return 0
}
func (m *Compare) GetModRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetTargetUnion().(*Compare_ModRevision); ok {
		return x.ModRevision
	}
	return 0
}
func (m *Compare) GetValue() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetTargetUnion().(*Compare_Value); ok {
		return x.Value
	}
	return nil
}
func (m *Compare) GetLease() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetTargetUnion().(*Compare_Lease); ok {
		return x.Lease
	}
	return 0
}
func (m *Compare) GetRangeEnd() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RangeEnd
	}
	return nil
}
func (*Compare) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _Compare_OneofMarshaler, _Compare_OneofUnmarshaler, _Compare_OneofSizer, []interface{}{(*Compare_Version)(nil), (*Compare_CreateRevision)(nil), (*Compare_ModRevision)(nil), (*Compare_Value)(nil), (*Compare_Lease)(nil)}
}
func _Compare_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*Compare)
	switch x := m.TargetUnion.(type) {
	case *Compare_Version:
		_ = b.EncodeVarint(4<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Version))
	case *Compare_CreateRevision:
		_ = b.EncodeVarint(5<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.CreateRevision))
	case *Compare_ModRevision:
		_ = b.EncodeVarint(6<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.ModRevision))
	case *Compare_Value:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		_ = b.EncodeRawBytes(x.Value)
	case *Compare_Lease:
		_ = b.EncodeVarint(8<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.Lease))
	case nil:
	default:
		return fmt.Errorf("Compare.TargetUnion has unexpected type %T", x)
	}
	return nil
}
func _Compare_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*Compare)
	switch tag {
	case 4:
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TargetUnion = &Compare_Version{int64(x)}
		return true, err
	case 5:
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TargetUnion = &Compare_CreateRevision{int64(x)}
		return true, err
	case 6:
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TargetUnion = &Compare_ModRevision{int64(x)}
		return true, err
	case 7:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.TargetUnion = &Compare_Value{x}
		return true, err
	case 8:
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TargetUnion = &Compare_Lease{int64(x)}
		return true, err
	default:
		return false, nil
	}
}
func _Compare_OneofSizer(msg proto.Message) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*Compare)
	switch x := m.TargetUnion.(type) {
	case *Compare_Version:
		n += proto.SizeVarint(4<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Version))
	case *Compare_CreateRevision:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.CreateRevision))
	case *Compare_ModRevision:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.ModRevision))
	case *Compare_Value:
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Value)))
		n += len(x.Value)
	case *Compare_Lease:
		n += proto.SizeVarint(8<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Lease))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TxnRequest struct {
	Compare	[]*Compare	`protobuf:"bytes,1,rep,name=compare" json:"compare,omitempty"`
	Success	[]*RequestOp	`protobuf:"bytes,2,rep,name=success" json:"success,omitempty"`
	Failure	[]*RequestOp	`protobuf:"bytes,3,rep,name=failure" json:"failure,omitempty"`
}

func (m *TxnRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = TxnRequest{}
}
func (m *TxnRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*TxnRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*TxnRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{10}
}
func (m *TxnRequest) GetCompare() []*Compare {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Compare
	}
	return nil
}
func (m *TxnRequest) GetSuccess() []*RequestOp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Success
	}
	return nil
}
func (m *TxnRequest) GetFailure() []*RequestOp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Failure
	}
	return nil
}

type TxnResponse struct {
	Header		*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Succeeded	bool		`protobuf:"varint,2,opt,name=succeeded,proto3" json:"succeeded,omitempty"`
	Responses	[]*ResponseOp	`protobuf:"bytes,3,rep,name=responses" json:"responses,omitempty"`
}

func (m *TxnResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = TxnResponse{}
}
func (m *TxnResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*TxnResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*TxnResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{11}
}
func (m *TxnResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *TxnResponse) GetSucceeded() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Succeeded
	}
	return false
}
func (m *TxnResponse) GetResponses() []*ResponseOp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Responses
	}
	return nil
}

type CompactionRequest struct {
	Revision	int64	`protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	Physical	bool	`protobuf:"varint,2,opt,name=physical,proto3" json:"physical,omitempty"`
}

func (m *CompactionRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = CompactionRequest{}
}
func (m *CompactionRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*CompactionRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*CompactionRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{12}
}
func (m *CompactionRequest) GetRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Revision
	}
	return 0
}
func (m *CompactionRequest) GetPhysical() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Physical
	}
	return false
}

type CompactionResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *CompactionResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = CompactionResponse{}
}
func (m *CompactionResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*CompactionResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*CompactionResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{13}
}
func (m *CompactionResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type HashRequest struct{}

func (m *HashRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = HashRequest{}
}
func (m *HashRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*HashRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*HashRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{14}
}

type HashKVRequest struct {
	Revision int64 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (m *HashKVRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = HashKVRequest{}
}
func (m *HashKVRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*HashKVRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*HashKVRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{15}
}
func (m *HashKVRequest) GetRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Revision
	}
	return 0
}

type HashKVResponse struct {
	Header		*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Hash		uint32		`protobuf:"varint,2,opt,name=hash,proto3" json:"hash,omitempty"`
	CompactRevision	int64		`protobuf:"varint,3,opt,name=compact_revision,json=compactRevision,proto3" json:"compact_revision,omitempty"`
}

func (m *HashKVResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = HashKVResponse{}
}
func (m *HashKVResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*HashKVResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*HashKVResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{16}
}
func (m *HashKVResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *HashKVResponse) GetHash() uint32 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Hash
	}
	return 0
}
func (m *HashKVResponse) GetCompactRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.CompactRevision
	}
	return 0
}

type HashResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Hash	uint32		`protobuf:"varint,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *HashResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = HashResponse{}
}
func (m *HashResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*HashResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*HashResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{17}
}
func (m *HashResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *HashResponse) GetHash() uint32 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Hash
	}
	return 0
}

type SnapshotRequest struct{}

func (m *SnapshotRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = SnapshotRequest{}
}
func (m *SnapshotRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*SnapshotRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*SnapshotRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{18}
}

type SnapshotResponse struct {
	Header		*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	RemainingBytes	uint64		`protobuf:"varint,2,opt,name=remaining_bytes,json=remainingBytes,proto3" json:"remaining_bytes,omitempty"`
	Blob		[]byte		`protobuf:"bytes,3,opt,name=blob,proto3" json:"blob,omitempty"`
}

func (m *SnapshotResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = SnapshotResponse{}
}
func (m *SnapshotResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*SnapshotResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*SnapshotResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{19}
}
func (m *SnapshotResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *SnapshotResponse) GetRemainingBytes() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RemainingBytes
	}
	return 0
}
func (m *SnapshotResponse) GetBlob() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Blob
	}
	return nil
}

type WatchRequest struct {
	RequestUnion isWatchRequest_RequestUnion `protobuf_oneof:"request_union"`
}

func (m *WatchRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = WatchRequest{}
}
func (m *WatchRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*WatchRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{20}
}

type isWatchRequest_RequestUnion interface {
	isWatchRequest_RequestUnion()
	MarshalTo([]byte) (int, error)
	Size() int
}
type WatchRequest_CreateRequest struct {
	CreateRequest *WatchCreateRequest `protobuf:"bytes,1,opt,name=create_request,json=createRequest,oneof"`
}
type WatchRequest_CancelRequest struct {
	CancelRequest *WatchCancelRequest `protobuf:"bytes,2,opt,name=cancel_request,json=cancelRequest,oneof"`
}
type WatchRequest_ProgressRequest struct {
	ProgressRequest *WatchProgressRequest `protobuf:"bytes,3,opt,name=progress_request,json=progressRequest,oneof"`
}

func (*WatchRequest_CreateRequest) isWatchRequest_RequestUnion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*WatchRequest_CancelRequest) isWatchRequest_RequestUnion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*WatchRequest_ProgressRequest) isWatchRequest_RequestUnion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (m *WatchRequest) GetRequestUnion() isWatchRequest_RequestUnion {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RequestUnion
	}
	return nil
}
func (m *WatchRequest) GetCreateRequest() *WatchCreateRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetRequestUnion().(*WatchRequest_CreateRequest); ok {
		return x.CreateRequest
	}
	return nil
}
func (m *WatchRequest) GetCancelRequest() *WatchCancelRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetRequestUnion().(*WatchRequest_CancelRequest); ok {
		return x.CancelRequest
	}
	return nil
}
func (m *WatchRequest) GetProgressRequest() *WatchProgressRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if x, ok := m.GetRequestUnion().(*WatchRequest_ProgressRequest); ok {
		return x.ProgressRequest
	}
	return nil
}
func (*WatchRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return _WatchRequest_OneofMarshaler, _WatchRequest_OneofUnmarshaler, _WatchRequest_OneofSizer, []interface{}{(*WatchRequest_CreateRequest)(nil), (*WatchRequest_CancelRequest)(nil), (*WatchRequest_ProgressRequest)(nil)}
}
func _WatchRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*WatchRequest)
	switch x := m.RequestUnion.(type) {
	case *WatchRequest_CreateRequest:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateRequest); err != nil {
			return err
		}
	case *WatchRequest_CancelRequest:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CancelRequest); err != nil {
			return err
		}
	case *WatchRequest_ProgressRequest:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ProgressRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("WatchRequest.RequestUnion has unexpected type %T", x)
	}
	return nil
}
func _WatchRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*WatchRequest)
	switch tag {
	case 1:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WatchCreateRequest)
		err := b.DecodeMessage(msg)
		m.RequestUnion = &WatchRequest_CreateRequest{msg}
		return true, err
	case 2:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WatchCancelRequest)
		err := b.DecodeMessage(msg)
		m.RequestUnion = &WatchRequest_CancelRequest{msg}
		return true, err
	case 3:
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WatchProgressRequest)
		err := b.DecodeMessage(msg)
		m.RequestUnion = &WatchRequest_ProgressRequest{msg}
		return true, err
	default:
		return false, nil
	}
}
func _WatchRequest_OneofSizer(msg proto.Message) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := msg.(*WatchRequest)
	switch x := m.RequestUnion.(type) {
	case *WatchRequest_CreateRequest:
		s := proto.Size(x.CreateRequest)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WatchRequest_CancelRequest:
		s := proto.Size(x.CancelRequest)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *WatchRequest_ProgressRequest:
		s := proto.Size(x.ProgressRequest)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type WatchCreateRequest struct {
	Key		[]byte				`protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	RangeEnd	[]byte				`protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	StartRevision	int64				`protobuf:"varint,3,opt,name=start_revision,json=startRevision,proto3" json:"start_revision,omitempty"`
	ProgressNotify	bool				`protobuf:"varint,4,opt,name=progress_notify,json=progressNotify,proto3" json:"progress_notify,omitempty"`
	Filters		[]WatchCreateRequest_FilterType	`protobuf:"varint,5,rep,packed,name=filters,enum=etcdserverpb.WatchCreateRequest_FilterType" json:"filters,omitempty"`
	PrevKv		bool				`protobuf:"varint,6,opt,name=prev_kv,json=prevKv,proto3" json:"prev_kv,omitempty"`
	WatchId		int64				`protobuf:"varint,7,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
	Fragment	bool				`protobuf:"varint,8,opt,name=fragment,proto3" json:"fragment,omitempty"`
}

func (m *WatchCreateRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = WatchCreateRequest{}
}
func (m *WatchCreateRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*WatchCreateRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*WatchCreateRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{21}
}
func (m *WatchCreateRequest) GetKey() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Key
	}
	return nil
}
func (m *WatchCreateRequest) GetRangeEnd() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RangeEnd
	}
	return nil
}
func (m *WatchCreateRequest) GetStartRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.StartRevision
	}
	return 0
}
func (m *WatchCreateRequest) GetProgressNotify() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ProgressNotify
	}
	return false
}
func (m *WatchCreateRequest) GetFilters() []WatchCreateRequest_FilterType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Filters
	}
	return nil
}
func (m *WatchCreateRequest) GetPrevKv() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.PrevKv
	}
	return false
}
func (m *WatchCreateRequest) GetWatchId() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.WatchId
	}
	return 0
}
func (m *WatchCreateRequest) GetFragment() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Fragment
	}
	return false
}

type WatchCancelRequest struct {
	WatchId int64 `protobuf:"varint,1,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
}

func (m *WatchCancelRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = WatchCancelRequest{}
}
func (m *WatchCancelRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*WatchCancelRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*WatchCancelRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{22}
}
func (m *WatchCancelRequest) GetWatchId() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.WatchId
	}
	return 0
}

type WatchProgressRequest struct{}

func (m *WatchProgressRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = WatchProgressRequest{}
}
func (m *WatchProgressRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*WatchProgressRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*WatchProgressRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{23}
}

type WatchResponse struct {
	Header		*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	WatchId		int64		`protobuf:"varint,2,opt,name=watch_id,json=watchId,proto3" json:"watch_id,omitempty"`
	Created		bool		`protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Canceled	bool		`protobuf:"varint,4,opt,name=canceled,proto3" json:"canceled,omitempty"`
	CompactRevision	int64		`protobuf:"varint,5,opt,name=compact_revision,json=compactRevision,proto3" json:"compact_revision,omitempty"`
	CancelReason	string		`protobuf:"bytes,6,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason,omitempty"`
	Fragment	bool		`protobuf:"varint,7,opt,name=fragment,proto3" json:"fragment,omitempty"`
	Events		[]*mvccpb.Event	`protobuf:"bytes,11,rep,name=events" json:"events,omitempty"`
}

func (m *WatchResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = WatchResponse{}
}
func (m *WatchResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*WatchResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{24}
}
func (m *WatchResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *WatchResponse) GetWatchId() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.WatchId
	}
	return 0
}
func (m *WatchResponse) GetCreated() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Created
	}
	return false
}
func (m *WatchResponse) GetCanceled() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Canceled
	}
	return false
}
func (m *WatchResponse) GetCompactRevision() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.CompactRevision
	}
	return 0
}
func (m *WatchResponse) GetCancelReason() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.CancelReason
	}
	return ""
}
func (m *WatchResponse) GetFragment() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Fragment
	}
	return false
}
func (m *WatchResponse) GetEvents() []*mvccpb.Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Events
	}
	return nil
}

type LeaseGrantRequest struct {
	TTL	int64	`protobuf:"varint,1,opt,name=TTL,proto3" json:"TTL,omitempty"`
	ID	int64	`protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *LeaseGrantRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseGrantRequest{}
}
func (m *LeaseGrantRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseGrantRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseGrantRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{25}
}
func (m *LeaseGrantRequest) GetTTL() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.TTL
	}
	return 0
}
func (m *LeaseGrantRequest) GetID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}

type LeaseGrantResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ID	int64		`protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	TTL	int64		`protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	Error	string		`protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *LeaseGrantResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseGrantResponse{}
}
func (m *LeaseGrantResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseGrantResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseGrantResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{26}
}
func (m *LeaseGrantResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *LeaseGrantResponse) GetID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}
func (m *LeaseGrantResponse) GetTTL() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.TTL
	}
	return 0
}
func (m *LeaseGrantResponse) GetError() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Error
	}
	return ""
}

type LeaseRevokeRequest struct {
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *LeaseRevokeRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseRevokeRequest{}
}
func (m *LeaseRevokeRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseRevokeRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseRevokeRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{27}
}
func (m *LeaseRevokeRequest) GetID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}

type LeaseRevokeResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *LeaseRevokeResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseRevokeResponse{}
}
func (m *LeaseRevokeResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseRevokeResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseRevokeResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{28}
}
func (m *LeaseRevokeResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type LeaseCheckpoint struct {
	ID		int64	`protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Remaining_TTL	int64	`protobuf:"varint,2,opt,name=remaining_TTL,json=remainingTTL,proto3" json:"remaining_TTL,omitempty"`
}

func (m *LeaseCheckpoint) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseCheckpoint{}
}
func (m *LeaseCheckpoint) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseCheckpoint) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseCheckpoint) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{29}
}
func (m *LeaseCheckpoint) GetID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}
func (m *LeaseCheckpoint) GetRemaining_TTL() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Remaining_TTL
	}
	return 0
}

type LeaseCheckpointRequest struct {
	Checkpoints []*LeaseCheckpoint `protobuf:"bytes,1,rep,name=checkpoints" json:"checkpoints,omitempty"`
}

func (m *LeaseCheckpointRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseCheckpointRequest{}
}
func (m *LeaseCheckpointRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseCheckpointRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseCheckpointRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{30}
}
func (m *LeaseCheckpointRequest) GetCheckpoints() []*LeaseCheckpoint {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Checkpoints
	}
	return nil
}

type LeaseCheckpointResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *LeaseCheckpointResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseCheckpointResponse{}
}
func (m *LeaseCheckpointResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseCheckpointResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseCheckpointResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{31}
}
func (m *LeaseCheckpointResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type LeaseKeepAliveRequest struct {
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *LeaseKeepAliveRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseKeepAliveRequest{}
}
func (m *LeaseKeepAliveRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseKeepAliveRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseKeepAliveRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{32}
}
func (m *LeaseKeepAliveRequest) GetID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}

type LeaseKeepAliveResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ID	int64		`protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	TTL	int64		`protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
}

func (m *LeaseKeepAliveResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseKeepAliveResponse{}
}
func (m *LeaseKeepAliveResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseKeepAliveResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseKeepAliveResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{33}
}
func (m *LeaseKeepAliveResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *LeaseKeepAliveResponse) GetID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}
func (m *LeaseKeepAliveResponse) GetTTL() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.TTL
	}
	return 0
}

type LeaseTimeToLiveRequest struct {
	ID	int64	`protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Keys	bool	`protobuf:"varint,2,opt,name=keys,proto3" json:"keys,omitempty"`
}

func (m *LeaseTimeToLiveRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseTimeToLiveRequest{}
}
func (m *LeaseTimeToLiveRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseTimeToLiveRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseTimeToLiveRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{34}
}
func (m *LeaseTimeToLiveRequest) GetID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}
func (m *LeaseTimeToLiveRequest) GetKeys() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Keys
	}
	return false
}

type LeaseTimeToLiveResponse struct {
	Header		*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	ID		int64		`protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	TTL		int64		`protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	GrantedTTL	int64		`protobuf:"varint,4,opt,name=grantedTTL,proto3" json:"grantedTTL,omitempty"`
	Keys		[][]byte	`protobuf:"bytes,5,rep,name=keys" json:"keys,omitempty"`
}

func (m *LeaseTimeToLiveResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseTimeToLiveResponse{}
}
func (m *LeaseTimeToLiveResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseTimeToLiveResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseTimeToLiveResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{35}
}
func (m *LeaseTimeToLiveResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *LeaseTimeToLiveResponse) GetID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}
func (m *LeaseTimeToLiveResponse) GetTTL() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.TTL
	}
	return 0
}
func (m *LeaseTimeToLiveResponse) GetGrantedTTL() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.GrantedTTL
	}
	return 0
}
func (m *LeaseTimeToLiveResponse) GetKeys() [][]byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Keys
	}
	return nil
}

type LeaseLeasesRequest struct{}

func (m *LeaseLeasesRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseLeasesRequest{}
}
func (m *LeaseLeasesRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseLeasesRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseLeasesRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{36}
}

type LeaseStatus struct {
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *LeaseStatus) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseStatus{}
}
func (m *LeaseStatus) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseStatus) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseStatus) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{37}
}
func (m *LeaseStatus) GetID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}

type LeaseLeasesResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Leases	[]*LeaseStatus	`protobuf:"bytes,2,rep,name=leases" json:"leases,omitempty"`
}

func (m *LeaseLeasesResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaseLeasesResponse{}
}
func (m *LeaseLeasesResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaseLeasesResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaseLeasesResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{38}
}
func (m *LeaseLeasesResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *LeaseLeasesResponse) GetLeases() []*LeaseStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Leases
	}
	return nil
}

type Member struct {
	ID		uint64		`protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name		string		`protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PeerURLs	[]string	`protobuf:"bytes,3,rep,name=peerURLs" json:"peerURLs,omitempty"`
	ClientURLs	[]string	`protobuf:"bytes,4,rep,name=clientURLs" json:"clientURLs,omitempty"`
}

func (m *Member) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Member{}
}
func (m *Member) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Member) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Member) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{39}
}
func (m *Member) GetID() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}
func (m *Member) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *Member) GetPeerURLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.PeerURLs
	}
	return nil
}
func (m *Member) GetClientURLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ClientURLs
	}
	return nil
}

type MemberAddRequest struct {
	PeerURLs []string `protobuf:"bytes,1,rep,name=peerURLs" json:"peerURLs,omitempty"`
}

func (m *MemberAddRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MemberAddRequest{}
}
func (m *MemberAddRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MemberAddRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MemberAddRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{40}
}
func (m *MemberAddRequest) GetPeerURLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.PeerURLs
	}
	return nil
}

type MemberAddResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Member	*Member		`protobuf:"bytes,2,opt,name=member" json:"member,omitempty"`
	Members	[]*Member	`protobuf:"bytes,3,rep,name=members" json:"members,omitempty"`
}

func (m *MemberAddResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MemberAddResponse{}
}
func (m *MemberAddResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MemberAddResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MemberAddResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{41}
}
func (m *MemberAddResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *MemberAddResponse) GetMember() *Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Member
	}
	return nil
}
func (m *MemberAddResponse) GetMembers() []*Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Members
	}
	return nil
}

type MemberRemoveRequest struct {
	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (m *MemberRemoveRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MemberRemoveRequest{}
}
func (m *MemberRemoveRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MemberRemoveRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MemberRemoveRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{42}
}
func (m *MemberRemoveRequest) GetID() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}

type MemberRemoveResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Members	[]*Member	`protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *MemberRemoveResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MemberRemoveResponse{}
}
func (m *MemberRemoveResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MemberRemoveResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MemberRemoveResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{43}
}
func (m *MemberRemoveResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *MemberRemoveResponse) GetMembers() []*Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Members
	}
	return nil
}

type MemberUpdateRequest struct {
	ID		uint64		`protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PeerURLs	[]string	`protobuf:"bytes,2,rep,name=peerURLs" json:"peerURLs,omitempty"`
}

func (m *MemberUpdateRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MemberUpdateRequest{}
}
func (m *MemberUpdateRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MemberUpdateRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MemberUpdateRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{44}
}
func (m *MemberUpdateRequest) GetID() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.ID
	}
	return 0
}
func (m *MemberUpdateRequest) GetPeerURLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.PeerURLs
	}
	return nil
}

type MemberUpdateResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Members	[]*Member	`protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *MemberUpdateResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MemberUpdateResponse{}
}
func (m *MemberUpdateResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MemberUpdateResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MemberUpdateResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{45}
}
func (m *MemberUpdateResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *MemberUpdateResponse) GetMembers() []*Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Members
	}
	return nil
}

type MemberListRequest struct{}

func (m *MemberListRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MemberListRequest{}
}
func (m *MemberListRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MemberListRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MemberListRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{46}
}

type MemberListResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Members	[]*Member	`protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *MemberListResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MemberListResponse{}
}
func (m *MemberListResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MemberListResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MemberListResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{47}
}
func (m *MemberListResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *MemberListResponse) GetMembers() []*Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Members
	}
	return nil
}

type DefragmentRequest struct{}

func (m *DefragmentRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = DefragmentRequest{}
}
func (m *DefragmentRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*DefragmentRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*DefragmentRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{48}
}

type DefragmentResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *DefragmentResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = DefragmentResponse{}
}
func (m *DefragmentResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*DefragmentResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*DefragmentResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{49}
}
func (m *DefragmentResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type MoveLeaderRequest struct {
	TargetID uint64 `protobuf:"varint,1,opt,name=targetID,proto3" json:"targetID,omitempty"`
}

func (m *MoveLeaderRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MoveLeaderRequest{}
}
func (m *MoveLeaderRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MoveLeaderRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MoveLeaderRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{50}
}
func (m *MoveLeaderRequest) GetTargetID() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.TargetID
	}
	return 0
}

type MoveLeaderResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *MoveLeaderResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = MoveLeaderResponse{}
}
func (m *MoveLeaderResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*MoveLeaderResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*MoveLeaderResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{51}
}
func (m *MoveLeaderResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AlarmRequest struct {
	Action		AlarmRequest_AlarmAction	`protobuf:"varint,1,opt,name=action,proto3,enum=etcdserverpb.AlarmRequest_AlarmAction" json:"action,omitempty"`
	MemberID	uint64				`protobuf:"varint,2,opt,name=memberID,proto3" json:"memberID,omitempty"`
	Alarm		AlarmType			`protobuf:"varint,3,opt,name=alarm,proto3,enum=etcdserverpb.AlarmType" json:"alarm,omitempty"`
}

func (m *AlarmRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AlarmRequest{}
}
func (m *AlarmRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AlarmRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AlarmRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{52}
}
func (m *AlarmRequest) GetAction() AlarmRequest_AlarmAction {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Action
	}
	return AlarmRequest_GET
}
func (m *AlarmRequest) GetMemberID() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.MemberID
	}
	return 0
}
func (m *AlarmRequest) GetAlarm() AlarmType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Alarm
	}
	return AlarmType_NONE
}

type AlarmMember struct {
	MemberID	uint64		`protobuf:"varint,1,opt,name=memberID,proto3" json:"memberID,omitempty"`
	Alarm		AlarmType	`protobuf:"varint,2,opt,name=alarm,proto3,enum=etcdserverpb.AlarmType" json:"alarm,omitempty"`
}

func (m *AlarmMember) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AlarmMember{}
}
func (m *AlarmMember) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AlarmMember) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AlarmMember) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{53}
}
func (m *AlarmMember) GetMemberID() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.MemberID
	}
	return 0
}
func (m *AlarmMember) GetAlarm() AlarmType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Alarm
	}
	return AlarmType_NONE
}

type AlarmResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Alarms	[]*AlarmMember	`protobuf:"bytes,2,rep,name=alarms" json:"alarms,omitempty"`
}

func (m *AlarmResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AlarmResponse{}
}
func (m *AlarmResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AlarmResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AlarmResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{54}
}
func (m *AlarmResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *AlarmResponse) GetAlarms() []*AlarmMember {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Alarms
	}
	return nil
}

type StatusRequest struct{}

func (m *StatusRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = StatusRequest{}
}
func (m *StatusRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*StatusRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{55}
}

type StatusResponse struct {
	Header			*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Version			string		`protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	DbSize			int64		`protobuf:"varint,3,opt,name=dbSize,proto3" json:"dbSize,omitempty"`
	Leader			uint64		`protobuf:"varint,4,opt,name=leader,proto3" json:"leader,omitempty"`
	RaftIndex		uint64		`protobuf:"varint,5,opt,name=raftIndex,proto3" json:"raftIndex,omitempty"`
	RaftTerm		uint64		`protobuf:"varint,6,opt,name=raftTerm,proto3" json:"raftTerm,omitempty"`
	RaftAppliedIndex	uint64		`protobuf:"varint,7,opt,name=raftAppliedIndex,proto3" json:"raftAppliedIndex,omitempty"`
	Errors			[]string	`protobuf:"bytes,8,rep,name=errors" json:"errors,omitempty"`
	DbSizeInUse		int64		`protobuf:"varint,9,opt,name=dbSizeInUse,proto3" json:"dbSizeInUse,omitempty"`
}

func (m *StatusResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = StatusResponse{}
}
func (m *StatusResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*StatusResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{56}
}
func (m *StatusResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *StatusResponse) GetVersion() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Version
	}
	return ""
}
func (m *StatusResponse) GetDbSize() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.DbSize
	}
	return 0
}
func (m *StatusResponse) GetLeader() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Leader
	}
	return 0
}
func (m *StatusResponse) GetRaftIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RaftIndex
	}
	return 0
}
func (m *StatusResponse) GetRaftTerm() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RaftTerm
	}
	return 0
}
func (m *StatusResponse) GetRaftAppliedIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RaftAppliedIndex
	}
	return 0
}
func (m *StatusResponse) GetErrors() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Errors
	}
	return nil
}
func (m *StatusResponse) GetDbSizeInUse() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.DbSizeInUse
	}
	return 0
}

type AuthEnableRequest struct{}

func (m *AuthEnableRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthEnableRequest{}
}
func (m *AuthEnableRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthEnableRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthEnableRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{57}
}

type AuthDisableRequest struct{}

func (m *AuthDisableRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthDisableRequest{}
}
func (m *AuthDisableRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthDisableRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthDisableRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{58}
}

type AuthenticateRequest struct {
	Name		string	`protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password	string	`protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *AuthenticateRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthenticateRequest{}
}
func (m *AuthenticateRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthenticateRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{59}
}
func (m *AuthenticateRequest) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *AuthenticateRequest) GetPassword() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthUserAddRequest struct {
	Name		string	`protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password	string	`protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *AuthUserAddRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserAddRequest{}
}
func (m *AuthUserAddRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserAddRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserAddRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{60}
}
func (m *AuthUserAddRequest) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *AuthUserAddRequest) GetPassword() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthUserGetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *AuthUserGetRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserGetRequest{}
}
func (m *AuthUserGetRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserGetRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserGetRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{61}
}
func (m *AuthUserGetRequest) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return ""
}

type AuthUserDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *AuthUserDeleteRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserDeleteRequest{}
}
func (m *AuthUserDeleteRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserDeleteRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserDeleteRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{62}
}
func (m *AuthUserDeleteRequest) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return ""
}

type AuthUserChangePasswordRequest struct {
	Name		string	`protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password	string	`protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (m *AuthUserChangePasswordRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserChangePasswordRequest{}
}
func (m *AuthUserChangePasswordRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserChangePasswordRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserChangePasswordRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{63}
}
func (m *AuthUserChangePasswordRequest) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *AuthUserChangePasswordRequest) GetPassword() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthUserGrantRoleRequest struct {
	User	string	`protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Role	string	`protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *AuthUserGrantRoleRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserGrantRoleRequest{}
}
func (m *AuthUserGrantRoleRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserGrantRoleRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserGrantRoleRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{64}
}
func (m *AuthUserGrantRoleRequest) GetUser() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.User
	}
	return ""
}
func (m *AuthUserGrantRoleRequest) GetRole() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Role
	}
	return ""
}

type AuthUserRevokeRoleRequest struct {
	Name	string	`protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Role	string	`protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *AuthUserRevokeRoleRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserRevokeRoleRequest{}
}
func (m *AuthUserRevokeRoleRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserRevokeRoleRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserRevokeRoleRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{65}
}
func (m *AuthUserRevokeRoleRequest) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *AuthUserRevokeRoleRequest) GetRole() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Role
	}
	return ""
}

type AuthRoleAddRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *AuthRoleAddRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleAddRequest{}
}
func (m *AuthRoleAddRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleAddRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleAddRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{66}
}
func (m *AuthRoleAddRequest) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return ""
}

type AuthRoleGetRequest struct {
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *AuthRoleGetRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleGetRequest{}
}
func (m *AuthRoleGetRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleGetRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleGetRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{67}
}
func (m *AuthRoleGetRequest) GetRole() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Role
	}
	return ""
}

type AuthUserListRequest struct{}

func (m *AuthUserListRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserListRequest{}
}
func (m *AuthUserListRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserListRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserListRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{68}
}

type AuthRoleListRequest struct{}

func (m *AuthRoleListRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleListRequest{}
}
func (m *AuthRoleListRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleListRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleListRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{69}
}

type AuthRoleDeleteRequest struct {
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *AuthRoleDeleteRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleDeleteRequest{}
}
func (m *AuthRoleDeleteRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleDeleteRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleDeleteRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{70}
}
func (m *AuthRoleDeleteRequest) GetRole() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Role
	}
	return ""
}

type AuthRoleGrantPermissionRequest struct {
	Name	string			`protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Perm	*authpb.Permission	`protobuf:"bytes,2,opt,name=perm" json:"perm,omitempty"`
}

func (m *AuthRoleGrantPermissionRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleGrantPermissionRequest{}
}
func (m *AuthRoleGrantPermissionRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleGrantPermissionRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleGrantPermissionRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{71}
}
func (m *AuthRoleGrantPermissionRequest) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return ""
}
func (m *AuthRoleGrantPermissionRequest) GetPerm() *authpb.Permission {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Perm
	}
	return nil
}

type AuthRoleRevokePermissionRequest struct {
	Role		string	`protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Key		[]byte	`protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	RangeEnd	[]byte	`protobuf:"bytes,3,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
}

func (m *AuthRoleRevokePermissionRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleRevokePermissionRequest{}
}
func (m *AuthRoleRevokePermissionRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleRevokePermissionRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleRevokePermissionRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{72}
}
func (m *AuthRoleRevokePermissionRequest) GetRole() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Role
	}
	return ""
}
func (m *AuthRoleRevokePermissionRequest) GetKey() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Key
	}
	return nil
}
func (m *AuthRoleRevokePermissionRequest) GetRangeEnd() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.RangeEnd
	}
	return nil
}

type AuthEnableResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthEnableResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthEnableResponse{}
}
func (m *AuthEnableResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthEnableResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthEnableResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{73}
}
func (m *AuthEnableResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthDisableResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthDisableResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthDisableResponse{}
}
func (m *AuthDisableResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthDisableResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthDisableResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{74}
}
func (m *AuthDisableResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthenticateResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Token	string		`protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *AuthenticateResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthenticateResponse{}
}
func (m *AuthenticateResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthenticateResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{75}
}
func (m *AuthenticateResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *AuthenticateResponse) GetToken() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthUserAddResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthUserAddResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserAddResponse{}
}
func (m *AuthUserAddResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserAddResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserAddResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{76}
}
func (m *AuthUserAddResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthUserGetResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Roles	[]string	`protobuf:"bytes,2,rep,name=roles" json:"roles,omitempty"`
}

func (m *AuthUserGetResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserGetResponse{}
}
func (m *AuthUserGetResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserGetResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserGetResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{77}
}
func (m *AuthUserGetResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *AuthUserGetResponse) GetRoles() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Roles
	}
	return nil
}

type AuthUserDeleteResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthUserDeleteResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserDeleteResponse{}
}
func (m *AuthUserDeleteResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserDeleteResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserDeleteResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{78}
}
func (m *AuthUserDeleteResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthUserChangePasswordResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthUserChangePasswordResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserChangePasswordResponse{}
}
func (m *AuthUserChangePasswordResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserChangePasswordResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserChangePasswordResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{79}
}
func (m *AuthUserChangePasswordResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthUserGrantRoleResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthUserGrantRoleResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserGrantRoleResponse{}
}
func (m *AuthUserGrantRoleResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserGrantRoleResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserGrantRoleResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{80}
}
func (m *AuthUserGrantRoleResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthUserRevokeRoleResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthUserRevokeRoleResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserRevokeRoleResponse{}
}
func (m *AuthUserRevokeRoleResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserRevokeRoleResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserRevokeRoleResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{81}
}
func (m *AuthUserRevokeRoleResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthRoleAddResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthRoleAddResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleAddResponse{}
}
func (m *AuthRoleAddResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleAddResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleAddResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{82}
}
func (m *AuthRoleAddResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthRoleGetResponse struct {
	Header	*ResponseHeader		`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Perm	[]*authpb.Permission	`protobuf:"bytes,2,rep,name=perm" json:"perm,omitempty"`
}

func (m *AuthRoleGetResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleGetResponse{}
}
func (m *AuthRoleGetResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleGetResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleGetResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{83}
}
func (m *AuthRoleGetResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *AuthRoleGetResponse) GetPerm() []*authpb.Permission {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Perm
	}
	return nil
}

type AuthRoleListResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Roles	[]string	`protobuf:"bytes,2,rep,name=roles" json:"roles,omitempty"`
}

func (m *AuthRoleListResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleListResponse{}
}
func (m *AuthRoleListResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleListResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleListResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{84}
}
func (m *AuthRoleListResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *AuthRoleListResponse) GetRoles() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Roles
	}
	return nil
}

type AuthUserListResponse struct {
	Header	*ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Users	[]string	`protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
}

func (m *AuthUserListResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthUserListResponse{}
}
func (m *AuthUserListResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthUserListResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthUserListResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{85}
}
func (m *AuthUserListResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *AuthUserListResponse) GetUsers() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Users
	}
	return nil
}

type AuthRoleDeleteResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthRoleDeleteResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleDeleteResponse{}
}
func (m *AuthRoleDeleteResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleDeleteResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleDeleteResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{86}
}
func (m *AuthRoleDeleteResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthRoleGrantPermissionResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthRoleGrantPermissionResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleGrantPermissionResponse{}
}
func (m *AuthRoleGrantPermissionResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleGrantPermissionResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleGrantPermissionResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{87}
}
func (m *AuthRoleGrantPermissionResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type AuthRoleRevokePermissionResponse struct {
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *AuthRoleRevokePermissionResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = AuthRoleRevokePermissionResponse{}
}
func (m *AuthRoleRevokePermissionResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*AuthRoleRevokePermissionResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*AuthRoleRevokePermissionResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{88}
}
func (m *AuthRoleRevokePermissionResponse) GetHeader() *ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterType((*ResponseHeader)(nil), "etcdserverpb.ResponseHeader")
	proto.RegisterType((*RangeRequest)(nil), "etcdserverpb.RangeRequest")
	proto.RegisterType((*RangeResponse)(nil), "etcdserverpb.RangeResponse")
	proto.RegisterType((*PutRequest)(nil), "etcdserverpb.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "etcdserverpb.PutResponse")
	proto.RegisterType((*DeleteRangeRequest)(nil), "etcdserverpb.DeleteRangeRequest")
	proto.RegisterType((*DeleteRangeResponse)(nil), "etcdserverpb.DeleteRangeResponse")
	proto.RegisterType((*RequestOp)(nil), "etcdserverpb.RequestOp")
	proto.RegisterType((*ResponseOp)(nil), "etcdserverpb.ResponseOp")
	proto.RegisterType((*Compare)(nil), "etcdserverpb.Compare")
	proto.RegisterType((*TxnRequest)(nil), "etcdserverpb.TxnRequest")
	proto.RegisterType((*TxnResponse)(nil), "etcdserverpb.TxnResponse")
	proto.RegisterType((*CompactionRequest)(nil), "etcdserverpb.CompactionRequest")
	proto.RegisterType((*CompactionResponse)(nil), "etcdserverpb.CompactionResponse")
	proto.RegisterType((*HashRequest)(nil), "etcdserverpb.HashRequest")
	proto.RegisterType((*HashKVRequest)(nil), "etcdserverpb.HashKVRequest")
	proto.RegisterType((*HashKVResponse)(nil), "etcdserverpb.HashKVResponse")
	proto.RegisterType((*HashResponse)(nil), "etcdserverpb.HashResponse")
	proto.RegisterType((*SnapshotRequest)(nil), "etcdserverpb.SnapshotRequest")
	proto.RegisterType((*SnapshotResponse)(nil), "etcdserverpb.SnapshotResponse")
	proto.RegisterType((*WatchRequest)(nil), "etcdserverpb.WatchRequest")
	proto.RegisterType((*WatchCreateRequest)(nil), "etcdserverpb.WatchCreateRequest")
	proto.RegisterType((*WatchCancelRequest)(nil), "etcdserverpb.WatchCancelRequest")
	proto.RegisterType((*WatchProgressRequest)(nil), "etcdserverpb.WatchProgressRequest")
	proto.RegisterType((*WatchResponse)(nil), "etcdserverpb.WatchResponse")
	proto.RegisterType((*LeaseGrantRequest)(nil), "etcdserverpb.LeaseGrantRequest")
	proto.RegisterType((*LeaseGrantResponse)(nil), "etcdserverpb.LeaseGrantResponse")
	proto.RegisterType((*LeaseRevokeRequest)(nil), "etcdserverpb.LeaseRevokeRequest")
	proto.RegisterType((*LeaseRevokeResponse)(nil), "etcdserverpb.LeaseRevokeResponse")
	proto.RegisterType((*LeaseCheckpoint)(nil), "etcdserverpb.LeaseCheckpoint")
	proto.RegisterType((*LeaseCheckpointRequest)(nil), "etcdserverpb.LeaseCheckpointRequest")
	proto.RegisterType((*LeaseCheckpointResponse)(nil), "etcdserverpb.LeaseCheckpointResponse")
	proto.RegisterType((*LeaseKeepAliveRequest)(nil), "etcdserverpb.LeaseKeepAliveRequest")
	proto.RegisterType((*LeaseKeepAliveResponse)(nil), "etcdserverpb.LeaseKeepAliveResponse")
	proto.RegisterType((*LeaseTimeToLiveRequest)(nil), "etcdserverpb.LeaseTimeToLiveRequest")
	proto.RegisterType((*LeaseTimeToLiveResponse)(nil), "etcdserverpb.LeaseTimeToLiveResponse")
	proto.RegisterType((*LeaseLeasesRequest)(nil), "etcdserverpb.LeaseLeasesRequest")
	proto.RegisterType((*LeaseStatus)(nil), "etcdserverpb.LeaseStatus")
	proto.RegisterType((*LeaseLeasesResponse)(nil), "etcdserverpb.LeaseLeasesResponse")
	proto.RegisterType((*Member)(nil), "etcdserverpb.Member")
	proto.RegisterType((*MemberAddRequest)(nil), "etcdserverpb.MemberAddRequest")
	proto.RegisterType((*MemberAddResponse)(nil), "etcdserverpb.MemberAddResponse")
	proto.RegisterType((*MemberRemoveRequest)(nil), "etcdserverpb.MemberRemoveRequest")
	proto.RegisterType((*MemberRemoveResponse)(nil), "etcdserverpb.MemberRemoveResponse")
	proto.RegisterType((*MemberUpdateRequest)(nil), "etcdserverpb.MemberUpdateRequest")
	proto.RegisterType((*MemberUpdateResponse)(nil), "etcdserverpb.MemberUpdateResponse")
	proto.RegisterType((*MemberListRequest)(nil), "etcdserverpb.MemberListRequest")
	proto.RegisterType((*MemberListResponse)(nil), "etcdserverpb.MemberListResponse")
	proto.RegisterType((*DefragmentRequest)(nil), "etcdserverpb.DefragmentRequest")
	proto.RegisterType((*DefragmentResponse)(nil), "etcdserverpb.DefragmentResponse")
	proto.RegisterType((*MoveLeaderRequest)(nil), "etcdserverpb.MoveLeaderRequest")
	proto.RegisterType((*MoveLeaderResponse)(nil), "etcdserverpb.MoveLeaderResponse")
	proto.RegisterType((*AlarmRequest)(nil), "etcdserverpb.AlarmRequest")
	proto.RegisterType((*AlarmMember)(nil), "etcdserverpb.AlarmMember")
	proto.RegisterType((*AlarmResponse)(nil), "etcdserverpb.AlarmResponse")
	proto.RegisterType((*StatusRequest)(nil), "etcdserverpb.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "etcdserverpb.StatusResponse")
	proto.RegisterType((*AuthEnableRequest)(nil), "etcdserverpb.AuthEnableRequest")
	proto.RegisterType((*AuthDisableRequest)(nil), "etcdserverpb.AuthDisableRequest")
	proto.RegisterType((*AuthenticateRequest)(nil), "etcdserverpb.AuthenticateRequest")
	proto.RegisterType((*AuthUserAddRequest)(nil), "etcdserverpb.AuthUserAddRequest")
	proto.RegisterType((*AuthUserGetRequest)(nil), "etcdserverpb.AuthUserGetRequest")
	proto.RegisterType((*AuthUserDeleteRequest)(nil), "etcdserverpb.AuthUserDeleteRequest")
	proto.RegisterType((*AuthUserChangePasswordRequest)(nil), "etcdserverpb.AuthUserChangePasswordRequest")
	proto.RegisterType((*AuthUserGrantRoleRequest)(nil), "etcdserverpb.AuthUserGrantRoleRequest")
	proto.RegisterType((*AuthUserRevokeRoleRequest)(nil), "etcdserverpb.AuthUserRevokeRoleRequest")
	proto.RegisterType((*AuthRoleAddRequest)(nil), "etcdserverpb.AuthRoleAddRequest")
	proto.RegisterType((*AuthRoleGetRequest)(nil), "etcdserverpb.AuthRoleGetRequest")
	proto.RegisterType((*AuthUserListRequest)(nil), "etcdserverpb.AuthUserListRequest")
	proto.RegisterType((*AuthRoleListRequest)(nil), "etcdserverpb.AuthRoleListRequest")
	proto.RegisterType((*AuthRoleDeleteRequest)(nil), "etcdserverpb.AuthRoleDeleteRequest")
	proto.RegisterType((*AuthRoleGrantPermissionRequest)(nil), "etcdserverpb.AuthRoleGrantPermissionRequest")
	proto.RegisterType((*AuthRoleRevokePermissionRequest)(nil), "etcdserverpb.AuthRoleRevokePermissionRequest")
	proto.RegisterType((*AuthEnableResponse)(nil), "etcdserverpb.AuthEnableResponse")
	proto.RegisterType((*AuthDisableResponse)(nil), "etcdserverpb.AuthDisableResponse")
	proto.RegisterType((*AuthenticateResponse)(nil), "etcdserverpb.AuthenticateResponse")
	proto.RegisterType((*AuthUserAddResponse)(nil), "etcdserverpb.AuthUserAddResponse")
	proto.RegisterType((*AuthUserGetResponse)(nil), "etcdserverpb.AuthUserGetResponse")
	proto.RegisterType((*AuthUserDeleteResponse)(nil), "etcdserverpb.AuthUserDeleteResponse")
	proto.RegisterType((*AuthUserChangePasswordResponse)(nil), "etcdserverpb.AuthUserChangePasswordResponse")
	proto.RegisterType((*AuthUserGrantRoleResponse)(nil), "etcdserverpb.AuthUserGrantRoleResponse")
	proto.RegisterType((*AuthUserRevokeRoleResponse)(nil), "etcdserverpb.AuthUserRevokeRoleResponse")
	proto.RegisterType((*AuthRoleAddResponse)(nil), "etcdserverpb.AuthRoleAddResponse")
	proto.RegisterType((*AuthRoleGetResponse)(nil), "etcdserverpb.AuthRoleGetResponse")
	proto.RegisterType((*AuthRoleListResponse)(nil), "etcdserverpb.AuthRoleListResponse")
	proto.RegisterType((*AuthUserListResponse)(nil), "etcdserverpb.AuthUserListResponse")
	proto.RegisterType((*AuthRoleDeleteResponse)(nil), "etcdserverpb.AuthRoleDeleteResponse")
	proto.RegisterType((*AuthRoleGrantPermissionResponse)(nil), "etcdserverpb.AuthRoleGrantPermissionResponse")
	proto.RegisterType((*AuthRoleRevokePermissionResponse)(nil), "etcdserverpb.AuthRoleRevokePermissionResponse")
	proto.RegisterEnum("etcdserverpb.AlarmType", AlarmType_name, AlarmType_value)
	proto.RegisterEnum("etcdserverpb.RangeRequest_SortOrder", RangeRequest_SortOrder_name, RangeRequest_SortOrder_value)
	proto.RegisterEnum("etcdserverpb.RangeRequest_SortTarget", RangeRequest_SortTarget_name, RangeRequest_SortTarget_value)
	proto.RegisterEnum("etcdserverpb.Compare_CompareResult", Compare_CompareResult_name, Compare_CompareResult_value)
	proto.RegisterEnum("etcdserverpb.Compare_CompareTarget", Compare_CompareTarget_name, Compare_CompareTarget_value)
	proto.RegisterEnum("etcdserverpb.WatchCreateRequest_FilterType", WatchCreateRequest_FilterType_name, WatchCreateRequest_FilterType_value)
	proto.RegisterEnum("etcdserverpb.AlarmRequest_AlarmAction", AlarmRequest_AlarmAction_name, AlarmRequest_AlarmAction_value)
}

var _ context.Context
var _ grpc.ClientConn

const _ = grpc.SupportPackageIsVersion4

type KVClient interface {
	Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*RangeResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	DeleteRange(ctx context.Context, in *DeleteRangeRequest, opts ...grpc.CallOption) (*DeleteRangeResponse, error)
	Txn(ctx context.Context, in *TxnRequest, opts ...grpc.CallOption) (*TxnResponse, error)
	Compact(ctx context.Context, in *CompactionRequest, opts ...grpc.CallOption) (*CompactionResponse, error)
}
type kVClient struct{ cc *grpc.ClientConn }

func NewKVClient(cc *grpc.ClientConn) KVClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &kVClient{cc}
}
func (c *kVClient) Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*RangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(RangeResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.KV/Range", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *kVClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.KV/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *kVClient) DeleteRange(ctx context.Context, in *DeleteRangeRequest, opts ...grpc.CallOption) (*DeleteRangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(DeleteRangeResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.KV/DeleteRange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *kVClient) Txn(ctx context.Context, in *TxnRequest, opts ...grpc.CallOption) (*TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(TxnResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.KV/Txn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *kVClient) Compact(ctx context.Context, in *CompactionRequest, opts ...grpc.CallOption) (*CompactionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(CompactionResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.KV/Compact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type KVServer interface {
	Range(context.Context, *RangeRequest) (*RangeResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	DeleteRange(context.Context, *DeleteRangeRequest) (*DeleteRangeResponse, error)
	Txn(context.Context, *TxnRequest) (*TxnResponse, error)
	Compact(context.Context, *CompactionRequest) (*CompactionResponse, error)
}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.RegisterService(&_KV_serviceDesc, srv)
}
func _KV_Range_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(RangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Range(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.KV/Range"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Range(ctx, req.(*RangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _KV_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.KV/Put"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _KV_DeleteRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(DeleteRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).DeleteRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.KV/DeleteRange"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).DeleteRange(ctx, req.(*DeleteRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _KV_Txn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(TxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Txn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.KV/Txn"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Txn(ctx, req.(*TxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _KV_Compact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(CompactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Compact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.KV/Compact"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Compact(ctx, req.(*CompactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KV_serviceDesc = grpc.ServiceDesc{ServiceName: "etcdserverpb.KV", HandlerType: (*KVServer)(nil), Methods: []grpc.MethodDesc{{MethodName: "Range", Handler: _KV_Range_Handler}, {MethodName: "Put", Handler: _KV_Put_Handler}, {MethodName: "DeleteRange", Handler: _KV_DeleteRange_Handler}, {MethodName: "Txn", Handler: _KV_Txn_Handler}, {MethodName: "Compact", Handler: _KV_Compact_Handler}}, Streams: []grpc.StreamDesc{}, Metadata: "rpc.proto"}

type WatchClient interface {
	Watch(ctx context.Context, opts ...grpc.CallOption) (Watch_WatchClient, error)
}
type watchClient struct{ cc *grpc.ClientConn }

func NewWatchClient(cc *grpc.ClientConn) WatchClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &watchClient{cc}
}
func (c *watchClient) Watch(ctx context.Context, opts ...grpc.CallOption) (Watch_WatchClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stream, err := grpc.NewClientStream(ctx, &_Watch_serviceDesc.Streams[0], c.cc, "/etcdserverpb.Watch/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &watchWatchClient{stream}
	return x, nil
}

type Watch_WatchClient interface {
	Send(*WatchRequest) error
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}
type watchWatchClient struct{ grpc.ClientStream }

func (x *watchWatchClient) Send(m *WatchRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return x.ClientStream.SendMsg(m)
}
func (x *watchWatchClient) Recv() (*WatchResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type WatchServer interface{ Watch(Watch_WatchServer) error }

func RegisterWatchServer(s *grpc.Server, srv WatchServer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.RegisterService(&_Watch_serviceDesc, srv)
}
func _Watch_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return srv.(WatchServer).Watch(&watchWatchServer{stream})
}

type Watch_WatchServer interface {
	Send(*WatchResponse) error
	Recv() (*WatchRequest, error)
	grpc.ServerStream
}
type watchWatchServer struct{ grpc.ServerStream }

func (x *watchWatchServer) Send(m *WatchResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return x.ServerStream.SendMsg(m)
}
func (x *watchWatchServer) Recv() (*WatchRequest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(WatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Watch_serviceDesc = grpc.ServiceDesc{ServiceName: "etcdserverpb.Watch", HandlerType: (*WatchServer)(nil), Methods: []grpc.MethodDesc{}, Streams: []grpc.StreamDesc{{StreamName: "Watch", Handler: _Watch_Watch_Handler, ServerStreams: true, ClientStreams: true}}, Metadata: "rpc.proto"}

type LeaseClient interface {
	LeaseGrant(ctx context.Context, in *LeaseGrantRequest, opts ...grpc.CallOption) (*LeaseGrantResponse, error)
	LeaseRevoke(ctx context.Context, in *LeaseRevokeRequest, opts ...grpc.CallOption) (*LeaseRevokeResponse, error)
	LeaseKeepAlive(ctx context.Context, opts ...grpc.CallOption) (Lease_LeaseKeepAliveClient, error)
	LeaseTimeToLive(ctx context.Context, in *LeaseTimeToLiveRequest, opts ...grpc.CallOption) (*LeaseTimeToLiveResponse, error)
	LeaseLeases(ctx context.Context, in *LeaseLeasesRequest, opts ...grpc.CallOption) (*LeaseLeasesResponse, error)
}
type leaseClient struct{ cc *grpc.ClientConn }

func NewLeaseClient(cc *grpc.ClientConn) LeaseClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &leaseClient{cc}
}
func (c *leaseClient) LeaseGrant(ctx context.Context, in *LeaseGrantRequest, opts ...grpc.CallOption) (*LeaseGrantResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(LeaseGrantResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Lease/LeaseGrant", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *leaseClient) LeaseRevoke(ctx context.Context, in *LeaseRevokeRequest, opts ...grpc.CallOption) (*LeaseRevokeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(LeaseRevokeResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Lease/LeaseRevoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *leaseClient) LeaseKeepAlive(ctx context.Context, opts ...grpc.CallOption) (Lease_LeaseKeepAliveClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stream, err := grpc.NewClientStream(ctx, &_Lease_serviceDesc.Streams[0], c.cc, "/etcdserverpb.Lease/LeaseKeepAlive", opts...)
	if err != nil {
		return nil, err
	}
	x := &leaseLeaseKeepAliveClient{stream}
	return x, nil
}

type Lease_LeaseKeepAliveClient interface {
	Send(*LeaseKeepAliveRequest) error
	Recv() (*LeaseKeepAliveResponse, error)
	grpc.ClientStream
}
type leaseLeaseKeepAliveClient struct{ grpc.ClientStream }

func (x *leaseLeaseKeepAliveClient) Send(m *LeaseKeepAliveRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return x.ClientStream.SendMsg(m)
}
func (x *leaseLeaseKeepAliveClient) Recv() (*LeaseKeepAliveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(LeaseKeepAliveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
func (c *leaseClient) LeaseTimeToLive(ctx context.Context, in *LeaseTimeToLiveRequest, opts ...grpc.CallOption) (*LeaseTimeToLiveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(LeaseTimeToLiveResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Lease/LeaseTimeToLive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *leaseClient) LeaseLeases(ctx context.Context, in *LeaseLeasesRequest, opts ...grpc.CallOption) (*LeaseLeasesResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(LeaseLeasesResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Lease/LeaseLeases", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type LeaseServer interface {
	LeaseGrant(context.Context, *LeaseGrantRequest) (*LeaseGrantResponse, error)
	LeaseRevoke(context.Context, *LeaseRevokeRequest) (*LeaseRevokeResponse, error)
	LeaseKeepAlive(Lease_LeaseKeepAliveServer) error
	LeaseTimeToLive(context.Context, *LeaseTimeToLiveRequest) (*LeaseTimeToLiveResponse, error)
	LeaseLeases(context.Context, *LeaseLeasesRequest) (*LeaseLeasesResponse, error)
}

func RegisterLeaseServer(s *grpc.Server, srv LeaseServer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.RegisterService(&_Lease_serviceDesc, srv)
}
func _Lease_LeaseGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(LeaseGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServer).LeaseGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Lease/LeaseGrant"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServer).LeaseGrant(ctx, req.(*LeaseGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Lease_LeaseRevoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(LeaseRevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServer).LeaseRevoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Lease/LeaseRevoke"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServer).LeaseRevoke(ctx, req.(*LeaseRevokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Lease_LeaseKeepAlive_Handler(srv interface{}, stream grpc.ServerStream) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return srv.(LeaseServer).LeaseKeepAlive(&leaseLeaseKeepAliveServer{stream})
}

type Lease_LeaseKeepAliveServer interface {
	Send(*LeaseKeepAliveResponse) error
	Recv() (*LeaseKeepAliveRequest, error)
	grpc.ServerStream
}
type leaseLeaseKeepAliveServer struct{ grpc.ServerStream }

func (x *leaseLeaseKeepAliveServer) Send(m *LeaseKeepAliveResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return x.ServerStream.SendMsg(m)
}
func (x *leaseLeaseKeepAliveServer) Recv() (*LeaseKeepAliveRequest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(LeaseKeepAliveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
func _Lease_LeaseTimeToLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(LeaseTimeToLiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServer).LeaseTimeToLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Lease/LeaseTimeToLive"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServer).LeaseTimeToLive(ctx, req.(*LeaseTimeToLiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Lease_LeaseLeases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(LeaseLeasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServer).LeaseLeases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Lease/LeaseLeases"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServer).LeaseLeases(ctx, req.(*LeaseLeasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lease_serviceDesc = grpc.ServiceDesc{ServiceName: "etcdserverpb.Lease", HandlerType: (*LeaseServer)(nil), Methods: []grpc.MethodDesc{{MethodName: "LeaseGrant", Handler: _Lease_LeaseGrant_Handler}, {MethodName: "LeaseRevoke", Handler: _Lease_LeaseRevoke_Handler}, {MethodName: "LeaseTimeToLive", Handler: _Lease_LeaseTimeToLive_Handler}, {MethodName: "LeaseLeases", Handler: _Lease_LeaseLeases_Handler}}, Streams: []grpc.StreamDesc{{StreamName: "LeaseKeepAlive", Handler: _Lease_LeaseKeepAlive_Handler, ServerStreams: true, ClientStreams: true}}, Metadata: "rpc.proto"}

type ClusterClient interface {
	MemberAdd(ctx context.Context, in *MemberAddRequest, opts ...grpc.CallOption) (*MemberAddResponse, error)
	MemberRemove(ctx context.Context, in *MemberRemoveRequest, opts ...grpc.CallOption) (*MemberRemoveResponse, error)
	MemberUpdate(ctx context.Context, in *MemberUpdateRequest, opts ...grpc.CallOption) (*MemberUpdateResponse, error)
	MemberList(ctx context.Context, in *MemberListRequest, opts ...grpc.CallOption) (*MemberListResponse, error)
}
type clusterClient struct{ cc *grpc.ClientConn }

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterClient{cc}
}
func (c *clusterClient) MemberAdd(ctx context.Context, in *MemberAddRequest, opts ...grpc.CallOption) (*MemberAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(MemberAddResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Cluster/MemberAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *clusterClient) MemberRemove(ctx context.Context, in *MemberRemoveRequest, opts ...grpc.CallOption) (*MemberRemoveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(MemberRemoveResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Cluster/MemberRemove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *clusterClient) MemberUpdate(ctx context.Context, in *MemberUpdateRequest, opts ...grpc.CallOption) (*MemberUpdateResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(MemberUpdateResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Cluster/MemberUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *clusterClient) MemberList(ctx context.Context, in *MemberListRequest, opts ...grpc.CallOption) (*MemberListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(MemberListResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Cluster/MemberList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ClusterServer interface {
	MemberAdd(context.Context, *MemberAddRequest) (*MemberAddResponse, error)
	MemberRemove(context.Context, *MemberRemoveRequest) (*MemberRemoveResponse, error)
	MemberUpdate(context.Context, *MemberUpdateRequest) (*MemberUpdateResponse, error)
	MemberList(context.Context, *MemberListRequest) (*MemberListResponse, error)
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.RegisterService(&_Cluster_serviceDesc, srv)
}
func _Cluster_MemberAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(MemberAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Cluster/MemberAdd"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberAdd(ctx, req.(*MemberAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Cluster_MemberRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(MemberRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Cluster/MemberRemove"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberRemove(ctx, req.(*MemberRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Cluster_MemberUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(MemberUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Cluster/MemberUpdate"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberUpdate(ctx, req.(*MemberUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Cluster_MemberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(MemberListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).MemberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Cluster/MemberList"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).MemberList(ctx, req.(*MemberListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{ServiceName: "etcdserverpb.Cluster", HandlerType: (*ClusterServer)(nil), Methods: []grpc.MethodDesc{{MethodName: "MemberAdd", Handler: _Cluster_MemberAdd_Handler}, {MethodName: "MemberRemove", Handler: _Cluster_MemberRemove_Handler}, {MethodName: "MemberUpdate", Handler: _Cluster_MemberUpdate_Handler}, {MethodName: "MemberList", Handler: _Cluster_MemberList_Handler}}, Streams: []grpc.StreamDesc{}, Metadata: "rpc.proto"}

type MaintenanceClient interface {
	Alarm(ctx context.Context, in *AlarmRequest, opts ...grpc.CallOption) (*AlarmResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Defragment(ctx context.Context, in *DefragmentRequest, opts ...grpc.CallOption) (*DefragmentResponse, error)
	Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error)
	HashKV(ctx context.Context, in *HashKVRequest, opts ...grpc.CallOption) (*HashKVResponse, error)
	Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (Maintenance_SnapshotClient, error)
	MoveLeader(ctx context.Context, in *MoveLeaderRequest, opts ...grpc.CallOption) (*MoveLeaderResponse, error)
}
type maintenanceClient struct{ cc *grpc.ClientConn }

func NewMaintenanceClient(cc *grpc.ClientConn) MaintenanceClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &maintenanceClient{cc}
}
func (c *maintenanceClient) Alarm(ctx context.Context, in *AlarmRequest, opts ...grpc.CallOption) (*AlarmResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AlarmResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Maintenance/Alarm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *maintenanceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Maintenance/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *maintenanceClient) Defragment(ctx context.Context, in *DefragmentRequest, opts ...grpc.CallOption) (*DefragmentResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(DefragmentResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Maintenance/Defragment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *maintenanceClient) Hash(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*HashResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(HashResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Maintenance/Hash", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *maintenanceClient) HashKV(ctx context.Context, in *HashKVRequest, opts ...grpc.CallOption) (*HashKVResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(HashKVResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Maintenance/HashKV", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *maintenanceClient) Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (Maintenance_SnapshotClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stream, err := grpc.NewClientStream(ctx, &_Maintenance_serviceDesc.Streams[0], c.cc, "/etcdserverpb.Maintenance/Snapshot", opts...)
	if err != nil {
		return nil, err
	}
	x := &maintenanceSnapshotClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Maintenance_SnapshotClient interface {
	Recv() (*SnapshotResponse, error)
	grpc.ClientStream
}
type maintenanceSnapshotClient struct{ grpc.ClientStream }

func (x *maintenanceSnapshotClient) Recv() (*SnapshotResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(SnapshotResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
func (c *maintenanceClient) MoveLeader(ctx context.Context, in *MoveLeaderRequest, opts ...grpc.CallOption) (*MoveLeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(MoveLeaderResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Maintenance/MoveLeader", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type MaintenanceServer interface {
	Alarm(context.Context, *AlarmRequest) (*AlarmResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	Defragment(context.Context, *DefragmentRequest) (*DefragmentResponse, error)
	Hash(context.Context, *HashRequest) (*HashResponse, error)
	HashKV(context.Context, *HashKVRequest) (*HashKVResponse, error)
	Snapshot(*SnapshotRequest, Maintenance_SnapshotServer) error
	MoveLeader(context.Context, *MoveLeaderRequest) (*MoveLeaderResponse, error)
}

func RegisterMaintenanceServer(s *grpc.Server, srv MaintenanceServer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.RegisterService(&_Maintenance_serviceDesc, srv)
}
func _Maintenance_Alarm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AlarmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Alarm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Maintenance/Alarm"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Alarm(ctx, req.(*AlarmRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Maintenance_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Maintenance/Status"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Maintenance_Defragment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(DefragmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Defragment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Maintenance/Defragment"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Defragment(ctx, req.(*DefragmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Maintenance_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Maintenance/Hash"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).Hash(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Maintenance_HashKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(HashKVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).HashKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Maintenance/HashKV"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).HashKV(ctx, req.(*HashKVRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Maintenance_Snapshot_Handler(srv interface{}, stream grpc.ServerStream) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(SnapshotRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MaintenanceServer).Snapshot(m, &maintenanceSnapshotServer{stream})
}

type Maintenance_SnapshotServer interface {
	Send(*SnapshotResponse) error
	grpc.ServerStream
}
type maintenanceSnapshotServer struct{ grpc.ServerStream }

func (x *maintenanceSnapshotServer) Send(m *SnapshotResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return x.ServerStream.SendMsg(m)
}
func _Maintenance_MoveLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(MoveLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServer).MoveLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Maintenance/MoveLeader"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServer).MoveLeader(ctx, req.(*MoveLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Maintenance_serviceDesc = grpc.ServiceDesc{ServiceName: "etcdserverpb.Maintenance", HandlerType: (*MaintenanceServer)(nil), Methods: []grpc.MethodDesc{{MethodName: "Alarm", Handler: _Maintenance_Alarm_Handler}, {MethodName: "Status", Handler: _Maintenance_Status_Handler}, {MethodName: "Defragment", Handler: _Maintenance_Defragment_Handler}, {MethodName: "Hash", Handler: _Maintenance_Hash_Handler}, {MethodName: "HashKV", Handler: _Maintenance_HashKV_Handler}, {MethodName: "MoveLeader", Handler: _Maintenance_MoveLeader_Handler}}, Streams: []grpc.StreamDesc{{StreamName: "Snapshot", Handler: _Maintenance_Snapshot_Handler, ServerStreams: true}}, Metadata: "rpc.proto"}

type AuthClient interface {
	AuthEnable(ctx context.Context, in *AuthEnableRequest, opts ...grpc.CallOption) (*AuthEnableResponse, error)
	AuthDisable(ctx context.Context, in *AuthDisableRequest, opts ...grpc.CallOption) (*AuthDisableResponse, error)
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	UserAdd(ctx context.Context, in *AuthUserAddRequest, opts ...grpc.CallOption) (*AuthUserAddResponse, error)
	UserGet(ctx context.Context, in *AuthUserGetRequest, opts ...grpc.CallOption) (*AuthUserGetResponse, error)
	UserList(ctx context.Context, in *AuthUserListRequest, opts ...grpc.CallOption) (*AuthUserListResponse, error)
	UserDelete(ctx context.Context, in *AuthUserDeleteRequest, opts ...grpc.CallOption) (*AuthUserDeleteResponse, error)
	UserChangePassword(ctx context.Context, in *AuthUserChangePasswordRequest, opts ...grpc.CallOption) (*AuthUserChangePasswordResponse, error)
	UserGrantRole(ctx context.Context, in *AuthUserGrantRoleRequest, opts ...grpc.CallOption) (*AuthUserGrantRoleResponse, error)
	UserRevokeRole(ctx context.Context, in *AuthUserRevokeRoleRequest, opts ...grpc.CallOption) (*AuthUserRevokeRoleResponse, error)
	RoleAdd(ctx context.Context, in *AuthRoleAddRequest, opts ...grpc.CallOption) (*AuthRoleAddResponse, error)
	RoleGet(ctx context.Context, in *AuthRoleGetRequest, opts ...grpc.CallOption) (*AuthRoleGetResponse, error)
	RoleList(ctx context.Context, in *AuthRoleListRequest, opts ...grpc.CallOption) (*AuthRoleListResponse, error)
	RoleDelete(ctx context.Context, in *AuthRoleDeleteRequest, opts ...grpc.CallOption) (*AuthRoleDeleteResponse, error)
	RoleGrantPermission(ctx context.Context, in *AuthRoleGrantPermissionRequest, opts ...grpc.CallOption) (*AuthRoleGrantPermissionResponse, error)
	RoleRevokePermission(ctx context.Context, in *AuthRoleRevokePermissionRequest, opts ...grpc.CallOption) (*AuthRoleRevokePermissionResponse, error)
}
type authClient struct{ cc *grpc.ClientConn }

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &authClient{cc}
}
func (c *authClient) AuthEnable(ctx context.Context, in *AuthEnableRequest, opts ...grpc.CallOption) (*AuthEnableResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthEnableResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/AuthEnable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) AuthDisable(ctx context.Context, in *AuthDisableRequest, opts ...grpc.CallOption) (*AuthDisableResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthDisableResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/AuthDisable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthenticateResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) UserAdd(ctx context.Context, in *AuthUserAddRequest, opts ...grpc.CallOption) (*AuthUserAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthUserAddResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/UserAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) UserGet(ctx context.Context, in *AuthUserGetRequest, opts ...grpc.CallOption) (*AuthUserGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthUserGetResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/UserGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) UserList(ctx context.Context, in *AuthUserListRequest, opts ...grpc.CallOption) (*AuthUserListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthUserListResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/UserList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) UserDelete(ctx context.Context, in *AuthUserDeleteRequest, opts ...grpc.CallOption) (*AuthUserDeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthUserDeleteResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/UserDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) UserChangePassword(ctx context.Context, in *AuthUserChangePasswordRequest, opts ...grpc.CallOption) (*AuthUserChangePasswordResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthUserChangePasswordResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/UserChangePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) UserGrantRole(ctx context.Context, in *AuthUserGrantRoleRequest, opts ...grpc.CallOption) (*AuthUserGrantRoleResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthUserGrantRoleResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/UserGrantRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) UserRevokeRole(ctx context.Context, in *AuthUserRevokeRoleRequest, opts ...grpc.CallOption) (*AuthUserRevokeRoleResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthUserRevokeRoleResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/UserRevokeRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) RoleAdd(ctx context.Context, in *AuthRoleAddRequest, opts ...grpc.CallOption) (*AuthRoleAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthRoleAddResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/RoleAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) RoleGet(ctx context.Context, in *AuthRoleGetRequest, opts ...grpc.CallOption) (*AuthRoleGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthRoleGetResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/RoleGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) RoleList(ctx context.Context, in *AuthRoleListRequest, opts ...grpc.CallOption) (*AuthRoleListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthRoleListResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/RoleList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) RoleDelete(ctx context.Context, in *AuthRoleDeleteRequest, opts ...grpc.CallOption) (*AuthRoleDeleteResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthRoleDeleteResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/RoleDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) RoleGrantPermission(ctx context.Context, in *AuthRoleGrantPermissionRequest, opts ...grpc.CallOption) (*AuthRoleGrantPermissionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthRoleGrantPermissionResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/RoleGrantPermission", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *authClient) RoleRevokePermission(ctx context.Context, in *AuthRoleRevokePermissionRequest, opts ...grpc.CallOption) (*AuthRoleRevokePermissionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(AuthRoleRevokePermissionResponse)
	err := grpc.Invoke(ctx, "/etcdserverpb.Auth/RoleRevokePermission", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type AuthServer interface {
	AuthEnable(context.Context, *AuthEnableRequest) (*AuthEnableResponse, error)
	AuthDisable(context.Context, *AuthDisableRequest) (*AuthDisableResponse, error)
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	UserAdd(context.Context, *AuthUserAddRequest) (*AuthUserAddResponse, error)
	UserGet(context.Context, *AuthUserGetRequest) (*AuthUserGetResponse, error)
	UserList(context.Context, *AuthUserListRequest) (*AuthUserListResponse, error)
	UserDelete(context.Context, *AuthUserDeleteRequest) (*AuthUserDeleteResponse, error)
	UserChangePassword(context.Context, *AuthUserChangePasswordRequest) (*AuthUserChangePasswordResponse, error)
	UserGrantRole(context.Context, *AuthUserGrantRoleRequest) (*AuthUserGrantRoleResponse, error)
	UserRevokeRole(context.Context, *AuthUserRevokeRoleRequest) (*AuthUserRevokeRoleResponse, error)
	RoleAdd(context.Context, *AuthRoleAddRequest) (*AuthRoleAddResponse, error)
	RoleGet(context.Context, *AuthRoleGetRequest) (*AuthRoleGetResponse, error)
	RoleList(context.Context, *AuthRoleListRequest) (*AuthRoleListResponse, error)
	RoleDelete(context.Context, *AuthRoleDeleteRequest) (*AuthRoleDeleteResponse, error)
	RoleGrantPermission(context.Context, *AuthRoleGrantPermissionRequest) (*AuthRoleGrantPermissionResponse, error)
	RoleRevokePermission(context.Context, *AuthRoleRevokePermissionRequest) (*AuthRoleRevokePermissionResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.RegisterService(&_Auth_serviceDesc, srv)
}
func _Auth_AuthEnable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthEnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthEnable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/AuthEnable"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthEnable(ctx, req.(*AuthEnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_AuthDisable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthDisableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthDisable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/AuthDisable"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthDisable(ctx, req.(*AuthDisableRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/Authenticate"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthUserAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/UserAdd"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserAdd(ctx, req.(*AuthUserAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_UserGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthUserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/UserGet"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserGet(ctx, req.(*AuthUserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/UserList"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserList(ctx, req.(*AuthUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthUserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/UserDelete"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserDelete(ctx, req.(*AuthUserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_UserChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthUserChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/UserChangePassword"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserChangePassword(ctx, req.(*AuthUserChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_UserGrantRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthUserGrantRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserGrantRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/UserGrantRole"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserGrantRole(ctx, req.(*AuthUserGrantRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_UserRevokeRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthUserRevokeRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UserRevokeRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/UserRevokeRole"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UserRevokeRole(ctx, req.(*AuthUserRevokeRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_RoleAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthRoleAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/RoleAdd"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleAdd(ctx, req.(*AuthRoleAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_RoleGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthRoleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/RoleGet"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleGet(ctx, req.(*AuthRoleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_RoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthRoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/RoleList"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleList(ctx, req.(*AuthRoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_RoleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthRoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/RoleDelete"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleDelete(ctx, req.(*AuthRoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_RoleGrantPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthRoleGrantPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleGrantPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/RoleGrantPermission"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleGrantPermission(ctx, req.(*AuthRoleGrantPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Auth_RoleRevokePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(AuthRoleRevokePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RoleRevokePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/etcdserverpb.Auth/RoleRevokePermission"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RoleRevokePermission(ctx, req.(*AuthRoleRevokePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{ServiceName: "etcdserverpb.Auth", HandlerType: (*AuthServer)(nil), Methods: []grpc.MethodDesc{{MethodName: "AuthEnable", Handler: _Auth_AuthEnable_Handler}, {MethodName: "AuthDisable", Handler: _Auth_AuthDisable_Handler}, {MethodName: "Authenticate", Handler: _Auth_Authenticate_Handler}, {MethodName: "UserAdd", Handler: _Auth_UserAdd_Handler}, {MethodName: "UserGet", Handler: _Auth_UserGet_Handler}, {MethodName: "UserList", Handler: _Auth_UserList_Handler}, {MethodName: "UserDelete", Handler: _Auth_UserDelete_Handler}, {MethodName: "UserChangePassword", Handler: _Auth_UserChangePassword_Handler}, {MethodName: "UserGrantRole", Handler: _Auth_UserGrantRole_Handler}, {MethodName: "UserRevokeRole", Handler: _Auth_UserRevokeRole_Handler}, {MethodName: "RoleAdd", Handler: _Auth_RoleAdd_Handler}, {MethodName: "RoleGet", Handler: _Auth_RoleGet_Handler}, {MethodName: "RoleList", Handler: _Auth_RoleList_Handler}, {MethodName: "RoleDelete", Handler: _Auth_RoleDelete_Handler}, {MethodName: "RoleGrantPermission", Handler: _Auth_RoleGrantPermission_Handler}, {MethodName: "RoleRevokePermission", Handler: _Auth_RoleRevokePermission_Handler}}, Streams: []grpc.StreamDesc{}, Metadata: "rpc.proto"}

func (m *ResponseHeader) Marshal() (dAtA []byte, err error) {
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
func (m *ResponseHeader) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ClusterId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ClusterId))
	}
	if m.MemberId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.MemberId))
	}
	if m.Revision != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Revision))
	}
	if m.RaftTerm != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RaftTerm))
	}
	return i, nil
}
func (m *RangeRequest) Marshal() (dAtA []byte, err error) {
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
func (m *RangeRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.RangeEnd) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i += copy(dAtA[i:], m.RangeEnd)
	}
	if m.Limit != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Limit))
	}
	if m.Revision != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Revision))
	}
	if m.SortOrder != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SortOrder))
	}
	if m.SortTarget != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SortTarget))
	}
	if m.Serializable {
		dAtA[i] = 0x38
		i++
		if m.Serializable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.KeysOnly {
		dAtA[i] = 0x40
		i++
		if m.KeysOnly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CountOnly {
		dAtA[i] = 0x48
		i++
		if m.CountOnly {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.MinModRevision != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.MinModRevision))
	}
	if m.MaxModRevision != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.MaxModRevision))
	}
	if m.MinCreateRevision != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.MinCreateRevision))
	}
	if m.MaxCreateRevision != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.MaxCreateRevision))
	}
	return i, nil
}
func (m *RangeResponse) Marshal() (dAtA []byte, err error) {
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
func (m *RangeResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n1, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Kvs) > 0 {
		for _, msg := range m.Kvs {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.More {
		dAtA[i] = 0x18
		i++
		if m.More {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Count != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Count))
	}
	return i, nil
}
func (m *PutRequest) Marshal() (dAtA []byte, err error) {
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
func (m *PutRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.Lease != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Lease))
	}
	if m.PrevKv {
		dAtA[i] = 0x20
		i++
		if m.PrevKv {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.IgnoreValue {
		dAtA[i] = 0x28
		i++
		if m.IgnoreValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.IgnoreLease {
		dAtA[i] = 0x30
		i++
		if m.IgnoreLease {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}
func (m *PutResponse) Marshal() (dAtA []byte, err error) {
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
func (m *PutResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n2, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.PrevKv != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.PrevKv.Size()))
		n3, err := m.PrevKv.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *DeleteRangeRequest) Marshal() (dAtA []byte, err error) {
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
func (m *DeleteRangeRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.RangeEnd) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i += copy(dAtA[i:], m.RangeEnd)
	}
	if m.PrevKv {
		dAtA[i] = 0x18
		i++
		if m.PrevKv {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}
func (m *DeleteRangeResponse) Marshal() (dAtA []byte, err error) {
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
func (m *DeleteRangeResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n4, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Deleted != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Deleted))
	}
	if len(m.PrevKvs) > 0 {
		for _, msg := range m.PrevKvs {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *RequestOp) Marshal() (dAtA []byte, err error) {
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
func (m *RequestOp) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Request != nil {
		nn5, err := m.Request.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn5
	}
	return i, nil
}
func (m *RequestOp_RequestRange) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.RequestRange != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RequestRange.Size()))
		n6, err := m.RequestRange.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *RequestOp_RequestPut) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.RequestPut != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RequestPut.Size()))
		n7, err := m.RequestPut.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func (m *RequestOp_RequestDeleteRange) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.RequestDeleteRange != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RequestDeleteRange.Size()))
		n8, err := m.RequestDeleteRange.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}
func (m *RequestOp_RequestTxn) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.RequestTxn != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RequestTxn.Size()))
		n9, err := m.RequestTxn.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	return i, nil
}
func (m *ResponseOp) Marshal() (dAtA []byte, err error) {
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
func (m *ResponseOp) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Response != nil {
		nn10, err := m.Response.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn10
	}
	return i, nil
}
func (m *ResponseOp_ResponseRange) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.ResponseRange != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ResponseRange.Size()))
		n11, err := m.ResponseRange.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	return i, nil
}
func (m *ResponseOp_ResponsePut) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.ResponsePut != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ResponsePut.Size()))
		n12, err := m.ResponsePut.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	return i, nil
}
func (m *ResponseOp_ResponseDeleteRange) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.ResponseDeleteRange != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ResponseDeleteRange.Size()))
		n13, err := m.ResponseDeleteRange.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	return i, nil
}
func (m *ResponseOp_ResponseTxn) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.ResponseTxn != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ResponseTxn.Size()))
		n14, err := m.ResponseTxn.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n14
	}
	return i, nil
}
func (m *Compare) Marshal() (dAtA []byte, err error) {
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
func (m *Compare) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Result))
	}
	if m.Target != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Target))
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.TargetUnion != nil {
		nn15, err := m.TargetUnion.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn15
	}
	if len(m.RangeEnd) > 0 {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x4
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i += copy(dAtA[i:], m.RangeEnd)
	}
	return i, nil
}
func (m *Compare_Version) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	dAtA[i] = 0x20
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.Version))
	return i, nil
}
func (m *Compare_CreateRevision) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	dAtA[i] = 0x28
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.CreateRevision))
	return i, nil
}
func (m *Compare_ModRevision) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	dAtA[i] = 0x30
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.ModRevision))
	return i, nil
}
func (m *Compare_Value) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.Value != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}
func (m *Compare_Lease) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	dAtA[i] = 0x40
	i++
	i = encodeVarintRpc(dAtA, i, uint64(m.Lease))
	return i, nil
}
func (m *TxnRequest) Marshal() (dAtA []byte, err error) {
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
func (m *TxnRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Compare) > 0 {
		for _, msg := range m.Compare {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Success) > 0 {
		for _, msg := range m.Success {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Failure) > 0 {
		for _, msg := range m.Failure {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *TxnResponse) Marshal() (dAtA []byte, err error) {
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
func (m *TxnResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n16, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n16
	}
	if m.Succeeded {
		dAtA[i] = 0x10
		i++
		if m.Succeeded {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Responses) > 0 {
		for _, msg := range m.Responses {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *CompactionRequest) Marshal() (dAtA []byte, err error) {
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
func (m *CompactionRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Revision != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Revision))
	}
	if m.Physical {
		dAtA[i] = 0x10
		i++
		if m.Physical {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}
func (m *CompactionResponse) Marshal() (dAtA []byte, err error) {
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
func (m *CompactionResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n17, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n17
	}
	return i, nil
}
func (m *HashRequest) Marshal() (dAtA []byte, err error) {
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
func (m *HashRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *HashKVRequest) Marshal() (dAtA []byte, err error) {
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
func (m *HashKVRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Revision != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Revision))
	}
	return i, nil
}
func (m *HashKVResponse) Marshal() (dAtA []byte, err error) {
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
func (m *HashKVResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n18, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n18
	}
	if m.Hash != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Hash))
	}
	if m.CompactRevision != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.CompactRevision))
	}
	return i, nil
}
func (m *HashResponse) Marshal() (dAtA []byte, err error) {
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
func (m *HashResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n19, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n19
	}
	if m.Hash != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Hash))
	}
	return i, nil
}
func (m *SnapshotRequest) Marshal() (dAtA []byte, err error) {
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
func (m *SnapshotRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *SnapshotResponse) Marshal() (dAtA []byte, err error) {
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
func (m *SnapshotResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n20, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n20
	}
	if m.RemainingBytes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RemainingBytes))
	}
	if len(m.Blob) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Blob)))
		i += copy(dAtA[i:], m.Blob)
	}
	return i, nil
}
func (m *WatchRequest) Marshal() (dAtA []byte, err error) {
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
func (m *WatchRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.RequestUnion != nil {
		nn21, err := m.RequestUnion.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn21
	}
	return i, nil
}
func (m *WatchRequest_CreateRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.CreateRequest != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.CreateRequest.Size()))
		n22, err := m.CreateRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n22
	}
	return i, nil
}
func (m *WatchRequest_CancelRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.CancelRequest != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.CancelRequest.Size()))
		n23, err := m.CancelRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n23
	}
	return i, nil
}
func (m *WatchRequest_ProgressRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := 0
	if m.ProgressRequest != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ProgressRequest.Size()))
		n24, err := m.ProgressRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n24
	}
	return i, nil
}
func (m *WatchCreateRequest) Marshal() (dAtA []byte, err error) {
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
func (m *WatchCreateRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.RangeEnd) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i += copy(dAtA[i:], m.RangeEnd)
	}
	if m.StartRevision != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.StartRevision))
	}
	if m.ProgressNotify {
		dAtA[i] = 0x20
		i++
		if m.ProgressNotify {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Filters) > 0 {
		dAtA26 := make([]byte, len(m.Filters)*10)
		var j25 int
		for _, num := range m.Filters {
			for num >= 1<<7 {
				dAtA26[j25] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j25++
			}
			dAtA26[j25] = uint8(num)
			j25++
		}
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(j25))
		i += copy(dAtA[i:], dAtA26[:j25])
	}
	if m.PrevKv {
		dAtA[i] = 0x30
		i++
		if m.PrevKv {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.WatchId != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.WatchId))
	}
	if m.Fragment {
		dAtA[i] = 0x40
		i++
		if m.Fragment {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}
func (m *WatchCancelRequest) Marshal() (dAtA []byte, err error) {
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
func (m *WatchCancelRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.WatchId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.WatchId))
	}
	return i, nil
}
func (m *WatchProgressRequest) Marshal() (dAtA []byte, err error) {
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
func (m *WatchProgressRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *WatchResponse) Marshal() (dAtA []byte, err error) {
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
func (m *WatchResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n27, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n27
	}
	if m.WatchId != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.WatchId))
	}
	if m.Created {
		dAtA[i] = 0x18
		i++
		if m.Created {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Canceled {
		dAtA[i] = 0x20
		i++
		if m.Canceled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CompactRevision != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.CompactRevision))
	}
	if len(m.CancelReason) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.CancelReason)))
		i += copy(dAtA[i:], m.CancelReason)
	}
	if m.Fragment {
		dAtA[i] = 0x38
		i++
		if m.Fragment {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Events) > 0 {
		for _, msg := range m.Events {
			dAtA[i] = 0x5a
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *LeaseGrantRequest) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseGrantRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.TTL != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.TTL))
	}
	if m.ID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	return i, nil
}
func (m *LeaseGrantResponse) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseGrantResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n28, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n28
	}
	if m.ID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	if m.TTL != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.TTL))
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	return i, nil
}
func (m *LeaseRevokeRequest) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseRevokeRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	return i, nil
}
func (m *LeaseRevokeResponse) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseRevokeResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n29, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n29
	}
	return i, nil
}
func (m *LeaseCheckpoint) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseCheckpoint) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	if m.Remaining_TTL != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Remaining_TTL))
	}
	return i, nil
}
func (m *LeaseCheckpointRequest) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseCheckpointRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Checkpoints) > 0 {
		for _, msg := range m.Checkpoints {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *LeaseCheckpointResponse) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseCheckpointResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n30, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n30
	}
	return i, nil
}
func (m *LeaseKeepAliveRequest) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseKeepAliveRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	return i, nil
}
func (m *LeaseKeepAliveResponse) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseKeepAliveResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n31, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n31
	}
	if m.ID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	if m.TTL != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.TTL))
	}
	return i, nil
}
func (m *LeaseTimeToLiveRequest) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseTimeToLiveRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	if m.Keys {
		dAtA[i] = 0x10
		i++
		if m.Keys {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}
func (m *LeaseTimeToLiveResponse) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseTimeToLiveResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n32, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n32
	}
	if m.ID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	if m.TTL != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.TTL))
	}
	if m.GrantedTTL != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.GrantedTTL))
	}
	if len(m.Keys) > 0 {
		for _, b := range m.Keys {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintRpc(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}
func (m *LeaseLeasesRequest) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseLeasesRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *LeaseStatus) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseStatus) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	return i, nil
}
func (m *LeaseLeasesResponse) Marshal() (dAtA []byte, err error) {
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
func (m *LeaseLeasesResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n33, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n33
	}
	if len(m.Leases) > 0 {
		for _, msg := range m.Leases {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *Member) Marshal() (dAtA []byte, err error) {
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
func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.PeerURLs) > 0 {
		for _, s := range m.PeerURLs {
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.ClientURLs) > 0 {
		for _, s := range m.ClientURLs {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}
func (m *MemberAddRequest) Marshal() (dAtA []byte, err error) {
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
func (m *MemberAddRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PeerURLs) > 0 {
		for _, s := range m.PeerURLs {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}
func (m *MemberAddResponse) Marshal() (dAtA []byte, err error) {
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
func (m *MemberAddResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n34, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n34
	}
	if m.Member != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Member.Size()))
		n35, err := m.Member.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n35
	}
	if len(m.Members) > 0 {
		for _, msg := range m.Members {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *MemberRemoveRequest) Marshal() (dAtA []byte, err error) {
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
func (m *MemberRemoveRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	return i, nil
}
func (m *MemberRemoveResponse) Marshal() (dAtA []byte, err error) {
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
func (m *MemberRemoveResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n36, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n36
	}
	if len(m.Members) > 0 {
		for _, msg := range m.Members {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *MemberUpdateRequest) Marshal() (dAtA []byte, err error) {
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
func (m *MemberUpdateRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ID))
	}
	if len(m.PeerURLs) > 0 {
		for _, s := range m.PeerURLs {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}
func (m *MemberUpdateResponse) Marshal() (dAtA []byte, err error) {
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
func (m *MemberUpdateResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n37, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n37
	}
	if len(m.Members) > 0 {
		for _, msg := range m.Members {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *MemberListRequest) Marshal() (dAtA []byte, err error) {
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
func (m *MemberListRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *MemberListResponse) Marshal() (dAtA []byte, err error) {
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
func (m *MemberListResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n38, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n38
	}
	if len(m.Members) > 0 {
		for _, msg := range m.Members {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *DefragmentRequest) Marshal() (dAtA []byte, err error) {
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
func (m *DefragmentRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *DefragmentResponse) Marshal() (dAtA []byte, err error) {
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
func (m *DefragmentResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n39, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n39
	}
	return i, nil
}
func (m *MoveLeaderRequest) Marshal() (dAtA []byte, err error) {
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
func (m *MoveLeaderRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.TargetID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.TargetID))
	}
	return i, nil
}
func (m *MoveLeaderResponse) Marshal() (dAtA []byte, err error) {
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
func (m *MoveLeaderResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n40, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n40
	}
	return i, nil
}
func (m *AlarmRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AlarmRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Action != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Action))
	}
	if m.MemberID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.MemberID))
	}
	if m.Alarm != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Alarm))
	}
	return i, nil
}
func (m *AlarmMember) Marshal() (dAtA []byte, err error) {
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
func (m *AlarmMember) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.MemberID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.MemberID))
	}
	if m.Alarm != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Alarm))
	}
	return i, nil
}
func (m *AlarmResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AlarmResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n41, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n41
	}
	if len(m.Alarms) > 0 {
		for _, msg := range m.Alarms {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *StatusRequest) Marshal() (dAtA []byte, err error) {
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
func (m *StatusRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *StatusResponse) Marshal() (dAtA []byte, err error) {
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
func (m *StatusResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n42, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n42
	}
	if len(m.Version) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	if m.DbSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.DbSize))
	}
	if m.Leader != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Leader))
	}
	if m.RaftIndex != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RaftIndex))
	}
	if m.RaftTerm != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RaftTerm))
	}
	if m.RaftAppliedIndex != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RaftAppliedIndex))
	}
	if len(m.Errors) > 0 {
		for _, s := range m.Errors {
			dAtA[i] = 0x42
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.DbSizeInUse != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.DbSizeInUse))
	}
	return i, nil
}
func (m *AuthEnableRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthEnableRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *AuthDisableRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthDisableRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *AuthenticateRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthenticateRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	return i, nil
}
func (m *AuthUserAddRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserAddRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	return i, nil
}
func (m *AuthUserGetRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserGetRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}
func (m *AuthUserDeleteRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserDeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}
func (m *AuthUserChangePasswordRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserChangePasswordRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	return i, nil
}
func (m *AuthUserGrantRoleRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserGrantRoleRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.User)))
		i += copy(dAtA[i:], m.User)
	}
	if len(m.Role) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	return i, nil
}
func (m *AuthUserRevokeRoleRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserRevokeRoleRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Role) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	return i, nil
}
func (m *AuthRoleAddRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleAddRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}
func (m *AuthRoleGetRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleGetRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Role) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	return i, nil
}
func (m *AuthUserListRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserListRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *AuthRoleListRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleListRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *AuthRoleDeleteRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleDeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Role) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	return i, nil
}
func (m *AuthRoleGrantPermissionRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleGrantPermissionRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Perm != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Perm.Size()))
		n43, err := m.Perm.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n43
	}
	return i, nil
}
func (m *AuthRoleRevokePermissionRequest) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleRevokePermissionRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Role) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Role)))
		i += copy(dAtA[i:], m.Role)
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.RangeEnd) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RangeEnd)))
		i += copy(dAtA[i:], m.RangeEnd)
	}
	return i, nil
}
func (m *AuthEnableResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthEnableResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n44, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n44
	}
	return i, nil
}
func (m *AuthDisableResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthDisableResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n45, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n45
	}
	return i, nil
}
func (m *AuthenticateResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthenticateResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n46, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n46
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}
func (m *AuthUserAddResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserAddResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n47, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n47
	}
	return i, nil
}
func (m *AuthUserGetResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserGetResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n48, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n48
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}
func (m *AuthUserDeleteResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserDeleteResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n49, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n49
	}
	return i, nil
}
func (m *AuthUserChangePasswordResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserChangePasswordResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n50, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n50
	}
	return i, nil
}
func (m *AuthUserGrantRoleResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserGrantRoleResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n51, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n51
	}
	return i, nil
}
func (m *AuthUserRevokeRoleResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserRevokeRoleResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n52, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n52
	}
	return i, nil
}
func (m *AuthRoleAddResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleAddResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n53, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n53
	}
	return i, nil
}
func (m *AuthRoleGetResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleGetResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n54, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n54
	}
	if len(m.Perm) > 0 {
		for _, msg := range m.Perm {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}
func (m *AuthRoleListResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleListResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n55, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n55
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}
func (m *AuthUserListResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthUserListResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n56, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n56
	}
	if len(m.Users) > 0 {
		for _, s := range m.Users {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}
func (m *AuthRoleDeleteResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleDeleteResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n57, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n57
	}
	return i, nil
}
func (m *AuthRoleGrantPermissionResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleGrantPermissionResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n58, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n58
	}
	return i, nil
}
func (m *AuthRoleRevokePermissionResponse) Marshal() (dAtA []byte, err error) {
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
func (m *AuthRoleRevokePermissionResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Header.Size()))
		n59, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n59
	}
	return i, nil
}
func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
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
func (m *ResponseHeader) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ClusterId != 0 {
		n += 1 + sovRpc(uint64(m.ClusterId))
	}
	if m.MemberId != 0 {
		n += 1 + sovRpc(uint64(m.MemberId))
	}
	if m.Revision != 0 {
		n += 1 + sovRpc(uint64(m.Revision))
	}
	if m.RaftTerm != 0 {
		n += 1 + sovRpc(uint64(m.RaftTerm))
	}
	return n
}
func (m *RangeRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Limit != 0 {
		n += 1 + sovRpc(uint64(m.Limit))
	}
	if m.Revision != 0 {
		n += 1 + sovRpc(uint64(m.Revision))
	}
	if m.SortOrder != 0 {
		n += 1 + sovRpc(uint64(m.SortOrder))
	}
	if m.SortTarget != 0 {
		n += 1 + sovRpc(uint64(m.SortTarget))
	}
	if m.Serializable {
		n += 2
	}
	if m.KeysOnly {
		n += 2
	}
	if m.CountOnly {
		n += 2
	}
	if m.MinModRevision != 0 {
		n += 1 + sovRpc(uint64(m.MinModRevision))
	}
	if m.MaxModRevision != 0 {
		n += 1 + sovRpc(uint64(m.MaxModRevision))
	}
	if m.MinCreateRevision != 0 {
		n += 1 + sovRpc(uint64(m.MinCreateRevision))
	}
	if m.MaxCreateRevision != 0 {
		n += 1 + sovRpc(uint64(m.MaxCreateRevision))
	}
	return n
}
func (m *RangeResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Kvs) > 0 {
		for _, e := range m.Kvs {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.More {
		n += 2
	}
	if m.Count != 0 {
		n += 1 + sovRpc(uint64(m.Count))
	}
	return n
}
func (m *PutRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Lease != 0 {
		n += 1 + sovRpc(uint64(m.Lease))
	}
	if m.PrevKv {
		n += 2
	}
	if m.IgnoreValue {
		n += 2
	}
	if m.IgnoreLease {
		n += 2
	}
	return n
}
func (m *PutResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.PrevKv != nil {
		l = m.PrevKv.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *DeleteRangeRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.PrevKv {
		n += 2
	}
	return n
}
func (m *DeleteRangeResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Deleted != 0 {
		n += 1 + sovRpc(uint64(m.Deleted))
	}
	if len(m.PrevKvs) > 0 {
		for _, e := range m.PrevKvs {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *RequestOp) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Request != nil {
		n += m.Request.Size()
	}
	return n
}
func (m *RequestOp_RequestRange) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.RequestRange != nil {
		l = m.RequestRange.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *RequestOp_RequestPut) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.RequestPut != nil {
		l = m.RequestPut.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *RequestOp_RequestDeleteRange) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.RequestDeleteRange != nil {
		l = m.RequestDeleteRange.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *RequestOp_RequestTxn) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.RequestTxn != nil {
		l = m.RequestTxn.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *ResponseOp) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Response != nil {
		n += m.Response.Size()
	}
	return n
}
func (m *ResponseOp_ResponseRange) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ResponseRange != nil {
		l = m.ResponseRange.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *ResponseOp_ResponsePut) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ResponsePut != nil {
		l = m.ResponsePut.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *ResponseOp_ResponseDeleteRange) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ResponseDeleteRange != nil {
		l = m.ResponseDeleteRange.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *ResponseOp_ResponseTxn) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ResponseTxn != nil {
		l = m.ResponseTxn.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *Compare) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovRpc(uint64(m.Result))
	}
	if m.Target != 0 {
		n += 1 + sovRpc(uint64(m.Target))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.TargetUnion != nil {
		n += m.TargetUnion.Size()
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *Compare_Version) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovRpc(uint64(m.Version))
	return n
}
func (m *Compare_CreateRevision) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovRpc(uint64(m.CreateRevision))
	return n
}
func (m *Compare_ModRevision) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovRpc(uint64(m.ModRevision))
	return n
}
func (m *Compare_Value) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Value != nil {
		l = len(m.Value)
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *Compare_Lease) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovRpc(uint64(m.Lease))
	return n
}
func (m *TxnRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if len(m.Compare) > 0 {
		for _, e := range m.Compare {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if len(m.Success) > 0 {
		for _, e := range m.Success {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if len(m.Failure) > 0 {
		for _, e := range m.Failure {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *TxnResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Succeeded {
		n += 2
	}
	if len(m.Responses) > 0 {
		for _, e := range m.Responses {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *CompactionRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Revision != 0 {
		n += 1 + sovRpc(uint64(m.Revision))
	}
	if m.Physical {
		n += 2
	}
	return n
}
func (m *CompactionResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *HashRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *HashKVRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Revision != 0 {
		n += 1 + sovRpc(uint64(m.Revision))
	}
	return n
}
func (m *HashKVResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Hash != 0 {
		n += 1 + sovRpc(uint64(m.Hash))
	}
	if m.CompactRevision != 0 {
		n += 1 + sovRpc(uint64(m.CompactRevision))
	}
	return n
}
func (m *HashResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Hash != 0 {
		n += 1 + sovRpc(uint64(m.Hash))
	}
	return n
}
func (m *SnapshotRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *SnapshotResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.RemainingBytes != 0 {
		n += 1 + sovRpc(uint64(m.RemainingBytes))
	}
	l = len(m.Blob)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *WatchRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.RequestUnion != nil {
		n += m.RequestUnion.Size()
	}
	return n
}
func (m *WatchRequest_CreateRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.CreateRequest != nil {
		l = m.CreateRequest.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *WatchRequest_CancelRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.CancelRequest != nil {
		l = m.CancelRequest.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *WatchRequest_ProgressRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ProgressRequest != nil {
		l = m.ProgressRequest.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *WatchCreateRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.StartRevision != 0 {
		n += 1 + sovRpc(uint64(m.StartRevision))
	}
	if m.ProgressNotify {
		n += 2
	}
	if len(m.Filters) > 0 {
		l = 0
		for _, e := range m.Filters {
			l += sovRpc(uint64(e))
		}
		n += 1 + sovRpc(uint64(l)) + l
	}
	if m.PrevKv {
		n += 2
	}
	if m.WatchId != 0 {
		n += 1 + sovRpc(uint64(m.WatchId))
	}
	if m.Fragment {
		n += 2
	}
	return n
}
func (m *WatchCancelRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.WatchId != 0 {
		n += 1 + sovRpc(uint64(m.WatchId))
	}
	return n
}
func (m *WatchProgressRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *WatchResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.WatchId != 0 {
		n += 1 + sovRpc(uint64(m.WatchId))
	}
	if m.Created {
		n += 2
	}
	if m.Canceled {
		n += 2
	}
	if m.CompactRevision != 0 {
		n += 1 + sovRpc(uint64(m.CompactRevision))
	}
	l = len(m.CancelReason)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Fragment {
		n += 2
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *LeaseGrantRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.TTL != 0 {
		n += 1 + sovRpc(uint64(m.TTL))
	}
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	return n
}
func (m *LeaseGrantResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovRpc(uint64(m.TTL))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *LeaseRevokeRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	return n
}
func (m *LeaseRevokeResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *LeaseCheckpoint) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.Remaining_TTL != 0 {
		n += 1 + sovRpc(uint64(m.Remaining_TTL))
	}
	return n
}
func (m *LeaseCheckpointRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if len(m.Checkpoints) > 0 {
		for _, e := range m.Checkpoints {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *LeaseCheckpointResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *LeaseKeepAliveRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	return n
}
func (m *LeaseKeepAliveResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovRpc(uint64(m.TTL))
	}
	return n
}
func (m *LeaseTimeToLiveRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.Keys {
		n += 2
	}
	return n
}
func (m *LeaseTimeToLiveResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if m.TTL != 0 {
		n += 1 + sovRpc(uint64(m.TTL))
	}
	if m.GrantedTTL != 0 {
		n += 1 + sovRpc(uint64(m.GrantedTTL))
	}
	if len(m.Keys) > 0 {
		for _, b := range m.Keys {
			l = len(b)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *LeaseLeasesRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *LeaseStatus) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	return n
}
func (m *LeaseLeasesResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Leases) > 0 {
		for _, e := range m.Leases {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *Member) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.PeerURLs) > 0 {
		for _, s := range m.PeerURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if len(m.ClientURLs) > 0 {
		for _, s := range m.ClientURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *MemberAddRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if len(m.PeerURLs) > 0 {
		for _, s := range m.PeerURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *MemberAddResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Member != nil {
		l = m.Member.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *MemberRemoveRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	return n
}
func (m *MemberRemoveResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *MemberUpdateRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRpc(uint64(m.ID))
	}
	if len(m.PeerURLs) > 0 {
		for _, s := range m.PeerURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *MemberUpdateResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *MemberListRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *MemberListResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *DefragmentRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *DefragmentResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *MoveLeaderRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.TargetID != 0 {
		n += 1 + sovRpc(uint64(m.TargetID))
	}
	return n
}
func (m *MoveLeaderResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AlarmRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Action != 0 {
		n += 1 + sovRpc(uint64(m.Action))
	}
	if m.MemberID != 0 {
		n += 1 + sovRpc(uint64(m.MemberID))
	}
	if m.Alarm != 0 {
		n += 1 + sovRpc(uint64(m.Alarm))
	}
	return n
}
func (m *AlarmMember) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.MemberID != 0 {
		n += 1 + sovRpc(uint64(m.MemberID))
	}
	if m.Alarm != 0 {
		n += 1 + sovRpc(uint64(m.Alarm))
	}
	return n
}
func (m *AlarmResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Alarms) > 0 {
		for _, e := range m.Alarms {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *StatusRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *StatusResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.DbSize != 0 {
		n += 1 + sovRpc(uint64(m.DbSize))
	}
	if m.Leader != 0 {
		n += 1 + sovRpc(uint64(m.Leader))
	}
	if m.RaftIndex != 0 {
		n += 1 + sovRpc(uint64(m.RaftIndex))
	}
	if m.RaftTerm != 0 {
		n += 1 + sovRpc(uint64(m.RaftTerm))
	}
	if m.RaftAppliedIndex != 0 {
		n += 1 + sovRpc(uint64(m.RaftAppliedIndex))
	}
	if len(m.Errors) > 0 {
		for _, s := range m.Errors {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.DbSizeInUse != 0 {
		n += 1 + sovRpc(uint64(m.DbSizeInUse))
	}
	return n
}
func (m *AuthEnableRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *AuthDisableRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *AuthenticateRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserAddRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserGetRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserDeleteRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserChangePasswordRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserGrantRoleRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserRevokeRoleRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthRoleAddRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthRoleGetRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserListRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *AuthRoleListRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *AuthRoleDeleteRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthRoleGrantPermissionRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Perm != nil {
		l = m.Perm.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthRoleRevokePermissionRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Role)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.RangeEnd)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthEnableResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthDisableResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthenticateResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserAddResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserGetResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *AuthUserDeleteResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserChangePasswordResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserGrantRoleResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthUserRevokeRoleResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthRoleAddResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthRoleGetResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Perm) > 0 {
		for _, e := range m.Perm {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *AuthRoleListResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *AuthUserListResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Users) > 0 {
		for _, s := range m.Users {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	return n
}
func (m *AuthRoleDeleteResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthRoleGrantPermissionResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *AuthRoleRevokePermissionResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func sovRpc(x uint64) (n int) {
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
func sozRpc(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResponseHeader) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: ResponseHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterId", wireType)
			}
			m.ClusterId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClusterId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberId", wireType)
			}
			m.MemberId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemberId |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftTerm", wireType)
			}
			m.RaftTerm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftTerm |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RangeRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: RangeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortOrder", wireType)
			}
			m.SortOrder = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SortOrder |= (RangeRequest_SortOrder(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SortTarget", wireType)
			}
			m.SortTarget = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SortTarget |= (RangeRequest_SortTarget(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Serializable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.Serializable = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeysOnly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.KeysOnly = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountOnly", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.CountOnly = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinModRevision", wireType)
			}
			m.MinModRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinModRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxModRevision", wireType)
			}
			m.MaxModRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxModRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCreateRevision", wireType)
			}
			m.MinCreateRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinCreateRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCreateRevision", wireType)
			}
			m.MaxCreateRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCreateRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RangeResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: RangeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kvs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kvs = append(m.Kvs, &mvccpb.KeyValue{})
			if err := m.Kvs[len(m.Kvs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field More", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.More = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PutRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: PutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			m.Lease = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lease |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.PrevKv = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoreValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.IgnoreValue = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IgnoreLease", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.IgnoreLease = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PutResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: PutResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PutResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrevKv == nil {
				m.PrevKv = &mvccpb.KeyValue{}
			}
			if err := m.PrevKv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeleteRangeRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DeleteRangeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteRangeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.PrevKv = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeleteRangeResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DeleteRangeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteRangeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			m.Deleted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deleted |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKvs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevKvs = append(m.PrevKvs, &mvccpb.KeyValue{})
			if err := m.PrevKvs[len(m.PrevKvs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestOp) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: RequestOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RangeRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Request = &RequestOp_RequestRange{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestPut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PutRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Request = &RequestOp_RequestPut{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestDeleteRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DeleteRangeRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Request = &RequestOp_RequestDeleteRange{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTxn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TxnRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Request = &RequestOp_RequestTxn{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseOp) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: ResponseOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &RangeResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &ResponseOp_ResponseRange{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponsePut", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PutResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &ResponseOp_ResponsePut{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseDeleteRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &DeleteRangeResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &ResponseOp_ResponseDeleteRange{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseTxn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TxnResponse{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Response = &ResponseOp_ResponseTxn{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Compare) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: Compare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Compare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= (Compare_CompareResult(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			m.Target = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Target |= (Compare_CompareTarget(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TargetUnion = &Compare_Version{v}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateRevision", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TargetUnion = &Compare_CreateRevision{v}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModRevision", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TargetUnion = &Compare_ModRevision{v}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.TargetUnion = &Compare_Value{v}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TargetUnion = &Compare_Lease{v}
		case 64:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TxnRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: TxnRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxnRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compare", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Compare = append(m.Compare, &Compare{})
			if err := m.Compare[len(m.Compare)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Success = append(m.Success, &RequestOp{})
			if err := m.Success[len(m.Success)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failure", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Failure = append(m.Failure, &RequestOp{})
			if err := m.Failure[len(m.Failure)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TxnResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: TxnResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxnResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Succeeded", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.Succeeded = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Responses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Responses = append(m.Responses, &ResponseOp{})
			if err := m.Responses[len(m.Responses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompactionRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: CompactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Physical", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.Physical = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CompactionResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: CompactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HashRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: HashRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HashKVRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: HashKVRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashKVRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Revision", wireType)
			}
			m.Revision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Revision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HashKVResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: HashKVResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashKVResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hash |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompactRevision", wireType)
			}
			m.CompactRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompactRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HashResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: HashResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			m.Hash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hash |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SnapshotRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: SnapshotRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SnapshotResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: SnapshotResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingBytes", wireType)
			}
			m.RemainingBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemainingBytes |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blob", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blob = append(m.Blob[:0], dAtA[iNdEx:postIndex]...)
			if m.Blob == nil {
				m.Blob = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &WatchCreateRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.RequestUnion = &WatchRequest_CreateRequest{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CancelRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &WatchCancelRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.RequestUnion = &WatchRequest_CancelRequest{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgressRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &WatchProgressRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.RequestUnion = &WatchRequest_ProgressRequest{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchCreateRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchCreateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchCreateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartRevision", wireType)
			}
			m.StartRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgressNotify", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.ProgressNotify = bool(v != 0)
		case 5:
			if wireType == 0 {
				var v WatchCreateRequest_FilterType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRpc
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (WatchCreateRequest_FilterType(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Filters = append(m.Filters, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRpc
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
					return ErrInvalidLengthRpc
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v WatchCreateRequest_FilterType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRpc
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (WatchCreateRequest_FilterType(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Filters = append(m.Filters, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevKv", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.PrevKv = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchId", wireType)
			}
			m.WatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WatchId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fragment", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.Fragment = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchCancelRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchCancelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchCancelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchId", wireType)
			}
			m.WatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WatchId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchProgressRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchProgressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchProgressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WatchResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: WatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WatchId", wireType)
			}
			m.WatchId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WatchId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.Created = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Canceled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.Canceled = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompactRevision", wireType)
			}
			m.CompactRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CompactRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CancelReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CancelReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fragment", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.Fragment = bool(v != 0)
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, &mvccpb.Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseGrantRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseGrantRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseGrantRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseGrantResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseGrantResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseGrantResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseRevokeRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseRevokeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseRevokeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseRevokeResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseRevokeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseRevokeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseCheckpoint) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseCheckpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseCheckpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remaining_TTL", wireType)
			}
			m.Remaining_TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Remaining_TTL |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseCheckpointRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseCheckpointRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseCheckpointRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checkpoints = append(m.Checkpoints, &LeaseCheckpoint{})
			if err := m.Checkpoints[len(m.Checkpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseCheckpointResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseCheckpointResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseCheckpointResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseKeepAliveRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseKeepAliveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseKeepAliveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseKeepAliveResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseKeepAliveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseKeepAliveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseTimeToLiveRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseTimeToLiveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseTimeToLiveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			m.Keys = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseTimeToLiveResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseTimeToLiveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseTimeToLiveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GrantedTTL", wireType)
			}
			m.GrantedTTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GrantedTTL |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Keys", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Keys = append(m.Keys, make([]byte, postIndex-iNdEx))
			copy(m.Keys[len(m.Keys)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseLeasesRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseLeasesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseLeasesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseStatus) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LeaseLeasesResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: LeaseLeasesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaseLeasesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leases", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Leases = append(m.Leases, &LeaseStatus{})
			if err := m.Leases[len(m.Leases)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Member) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerURLs = append(m.PeerURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientURLs = append(m.ClientURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MemberAddRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerURLs = append(m.PeerURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MemberAddResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberAddResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberAddResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Member == nil {
				m.Member = &Member{}
			}
			if err := m.Member.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MemberRemoveRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberRemoveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberRemoveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MemberRemoveResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberRemoveResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberRemoveResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MemberUpdateRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberUpdateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberUpdateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerURLs = append(m.PeerURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MemberUpdateResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberUpdateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberUpdateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MemberListRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MemberListResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MemberListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MemberListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &Member{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DefragmentRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DefragmentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefragmentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DefragmentResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: DefragmentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefragmentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MoveLeaderRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MoveLeaderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MoveLeaderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetID", wireType)
			}
			m.TargetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MoveLeaderResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: MoveLeaderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MoveLeaderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AlarmRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AlarmRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlarmRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= (AlarmRequest_AlarmAction(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberID", wireType)
			}
			m.MemberID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemberID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alarm", wireType)
			}
			m.Alarm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Alarm |= (AlarmType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AlarmMember) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AlarmMember: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlarmMember: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberID", wireType)
			}
			m.MemberID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemberID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alarm", wireType)
			}
			m.Alarm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Alarm |= (AlarmType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AlarmResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AlarmResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AlarmResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alarms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alarms = append(m.Alarms, &AlarmMember{})
			if err := m.Alarms[len(m.Alarms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: StatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: StatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbSize", wireType)
			}
			m.DbSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DbSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leader", wireType)
			}
			m.Leader = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Leader |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftIndex", wireType)
			}
			m.RaftIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftTerm", wireType)
			}
			m.RaftTerm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftTerm |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAppliedIndex", wireType)
			}
			m.RaftAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errors = append(m.Errors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbSizeInUse", wireType)
			}
			m.DbSizeInUse = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DbSizeInUse |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthEnableRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthEnableRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthEnableRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthDisableRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthDisableRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthDisableRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthenticateRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthenticateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserAddRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserGetRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserGetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserGetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserDeleteRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserDeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserDeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserChangePasswordRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserChangePasswordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserChangePasswordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserGrantRoleRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserGrantRoleRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserGrantRoleRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserRevokeRoleRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserRevokeRoleRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserRevokeRoleRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleAddRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleAddRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleAddRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleGetRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleGetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleGetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserListRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleListRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleDeleteRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleDeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleDeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleGrantPermissionRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleGrantPermissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleGrantPermissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Perm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Perm == nil {
				m.Perm = &authpb.Permission{}
			}
			if err := m.Perm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleRevokePermissionRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleRevokePermissionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleRevokePermissionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Role = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeEnd", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RangeEnd = append(m.RangeEnd[:0], dAtA[iNdEx:postIndex]...)
			if m.RangeEnd == nil {
				m.RangeEnd = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthEnableResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthEnableResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthEnableResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthDisableResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthDisableResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthDisableResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthenticateResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthenticateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserAddResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserAddResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserAddResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserGetResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserDeleteResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserDeleteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserDeleteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserChangePasswordResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserChangePasswordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserChangePasswordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserGrantRoleResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserGrantRoleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserGrantRoleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserRevokeRoleResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserRevokeRoleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserRevokeRoleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleAddResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleAddResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleAddResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleGetResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Perm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Perm = append(m.Perm, &authpb.Permission{})
			if err := m.Perm[len(m.Perm)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleListResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthUserListResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthUserListResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthUserListResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleDeleteResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleDeleteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleDeleteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleGrantPermissionResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleGrantPermissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleGrantPermissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AuthRoleRevokePermissionResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: AuthRoleRevokePermissionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthRoleRevokePermissionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRpc(dAtA []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
				return 0, ErrInvalidLengthRpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRpc
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
				next, err := skipRpc(dAtA[start:])
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
	ErrInvalidLengthRpc	= fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc	= fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("rpc.proto", fileDescriptorRpc)
}

var fileDescriptorRpc = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x5b, 0xdd, 0x6f, 0x23, 0xc9, 0x71, 0xd7, 0x90, 0xe2, 0x57, 0xf1, 0x43, 0x54, 0xeb, 0x63, 0x29, 0xee, 0xae, 0x56, 0xd7, 0xbb, 0x7b, 0xab, 0xdb, 0xbd, 0x13, 0x6d, 0xd9, 0x4e, 0x80, 0x4d, 0xe2, 0x58, 0x2b, 0xf1, 0x56, 0x3a, 0x69, 0x45, 0xdd, 0x88, 0xda, 0xfb, 0x80, 0x11, 0x61, 0x44, 0xf6, 0x4a, 0x13, 0x91, 0x33, 0xf4, 0xcc, 0x90, 0x2b, 0x5d, 0x82, 0x38, 0x30, 0x9c, 0x00, 0xc9, 0xa3, 0x0d, 0x04, 0xc9, 0x43, 0x9e, 0x82, 0x20, 0xf0, 0x43, 0x80, 0xbc, 0x05, 0xc8, 0x5f, 0x90, 0xb7, 0x24, 0xc8, 0x3f, 0x10, 0x5c, 0xfc, 0x92, 0xff, 0x22, 0xe8, 0xaf, 0x99, 0x9e, 0x2f, 0x69, 0x6d, 0xfa, 0xfc, 0x22, 0x4d, 0x57, 0x57, 0x57, 0x55, 0x57, 0x77, 0x57, 0x55, 0xff, 0x66, 0x08, 0x25, 0x67, 0xd4, 0xdb, 0x18, 0x39, 0xb6, 0x67, 0xa3, 0x0a, 0xf1, 0x7a, 0x7d, 0x97, 0x38, 0x13, 0xe2, 0x8c, 0xce, 0x9a, 0x8b, 0xe7, 0xf6, 0xb9, 0xcd, 0x3a, 0x5a, 0xf4, 0x89, 0xf3, 0x34, 0x57, 0x28, 0x4f, 0x6b, 0x38, 0xe9, 0xf5, 0xd8, 0x9f, 0xd1, 0x59, 0xeb, 0x72, 0x22, 0xba, 0xee, 0xb2, 0x2e, 0x63, 0xec, 0x5d, 0xb0, 0x3f, 0xa3, 0x33, 0xf6, 0x4f, 0x74, 0xde, 0x3b, 0xb7, 0xed, 0xf3, 0x01, 0x69, 0x19, 0x23, 0xb3, 0x65, 0x58, 0x96, 0xed, 0x19, 0x9e, 0x69, 0x5b, 0x2e, 0xef, 0xc5, 0x7f, 0xa1, 0x41, 0x4d, 0x27, 0xee, 0xc8, 0xb6, 0x5c, 0xb2, 0x4b, 0x8c, 0x3e, 0x71, 0xd0, 0x7d, 0x80, 0xde, 0x60, 0xec, 0x7a, 0xc4, 0x39, 0x35, 0xfb, 0x0d, 0x6d, 0x4d, 0x5b, 0x9f, 0xd5, 0x4b, 0x82, 0xb2, 0xd7, 0x47, 0x77, 0xa1, 0x34, 0x24, 0xc3, 0x33, 0xde, 0x9b, 0x61, 0xbd, 0x45, 0x4e, 0xd8, 0xeb, 0xa3, 0x26, 0x14, 0x1d, 0x32, 0x31, 0x5d, 0xd3, 0xb6, 0x1a, 0xd9, 0x35, 0x6d, 0x3d, 0xab, 0xfb, 0x6d, 0x3a, 0xd0, 0x31, 0xde, 0x78, 0xa7, 0x1e, 0x71, 0x86, 0x8d, 0x59, 0x3e, 0x90, 0x12, 0xba, 0xc4, 0x19, 0xe2, 0x9f, 0xe6, 0xa0, 0xa2, 0x1b, 0xd6, 0x39, 0xd1, 0xc9, 0x8f, 0xc6, 0xc4, 0xf5, 0x50, 0x1d, 0xb2, 0x97, 0xe4, 0x9a, 0xa9, 0xaf, 0xe8, 0xf4, 0x91, 0x8f, 0xb7, 0xce, 0xc9, 0x29, 0xb1, 0xb8, 0xe2, 0x0a, 0x1d, 0x6f, 0x9d, 0x93, 0xb6, 0xd5, 0x47, 0x8b, 0x90, 0x1b, 0x98, 0x43, 0xd3, 0x13, 0x5a, 0x79, 0x23, 0x64, 0xce, 0x6c, 0xc4, 0x9c, 0x6d, 0x00, 0xd7, 0x76, 0xbc, 0x53, 0xdb, 0xe9, 0x13, 0xa7, 0x91, 0x5b, 0xd3, 0xd6, 0x6b, 0x9b, 0x8f, 0x36, 0xd4, 0x85, 0xd8, 0x50, 0x0d, 0xda, 0x38, 0xb6, 0x1d, 0xaf, 0x43, 0x79, 0xf5, 0x92, 0x2b, 0x1f, 0xd1, 0xc7, 0x50, 0x66, 0x42, 0x3c, 0xc3, 0x39, 0x27, 0x5e, 0x23, 0xcf, 0xa4, 0x3c, 0xbe, 0x45, 0x4a, 0x97, 0x31, 0xeb, 0x4c, 0x3d, 0x7f, 0x46, 0x18, 0x2a, 0x2e, 0x71, 0x4c, 0x63, 0x60, 0x7e, 0x65, 0x9c, 0x0d, 0x48, 0xa3, 0xb0, 0xa6, 0xad, 0x17, 0xf5, 0x10, 0x8d, 0xce, 0xff, 0x92, 0x5c, 0xbb, 0xa7, 0xb6, 0x35, 0xb8, 0x6e, 0x14, 0x19, 0x43, 0x91, 0x12, 0x3a, 0xd6, 0xe0, 0x9a, 0x2d, 0x9a, 0x3d, 0xb6, 0x3c, 0xde, 0x5b, 0x62, 0xbd, 0x25, 0x46, 0x61, 0xdd, 0xeb, 0x50, 0x1f, 0x9a, 0xd6, 0xe9, 0xd0, 0xee, 0x9f, 0xfa, 0x0e, 0x01, 0xe6, 0x90, 0xda, 0xd0, 0xb4, 0x5e, 0xd9, 0x7d, 0x5d, 0xba, 0x85, 0x72, 0x1a, 0x57, 0x61, 0xce, 0xb2, 0xe0, 0x34, 0xae, 0x54, 0xce, 0x0d, 0x58, 0xa0, 0x32, 0x7b, 0x0e, 0x31, 0x3c, 0x12, 0x30, 0x57, 0x18, 0xf3, 0xfc, 0xd0, 0xb4, 0xb6, 0x59, 0x4f, 0x88, 0xdf, 0xb8, 0x8a, 0xf1, 0x57, 0x05, 0xbf, 0x71, 0x15, 0xe6, 0xc7, 0x1b, 0x50, 0xf2, 0x7d, 0x8e, 0x8a, 0x30, 0x7b, 0xd8, 0x39, 0x6c, 0xd7, 0x67, 0x10, 0x40, 0x7e, 0xeb, 0x78, 0xbb, 0x7d, 0xb8, 0x53, 0xd7, 0x50, 0x19, 0x0a, 0x3b, 0x6d, 0xde, 0xc8, 0xe0, 0x17, 0x00, 0x81, 0x77, 0x51, 0x01, 0xb2, 0xfb, 0xed, 0x2f, 0xea, 0x33, 0x94, 0xe7, 0x75, 0x5b, 0x3f, 0xde, 0xeb, 0x1c, 0xd6, 0x35, 0x3a, 0x78, 0x5b, 0x6f, 0x6f, 0x75, 0xdb, 0xf5, 0x0c, 0xe5, 0x78, 0xd5, 0xd9, 0xa9, 0x67, 0x51, 0x09, 0x72, 0xaf, 0xb7, 0x0e, 0x4e, 0xda, 0xf5, 0x59, 0xfc, 0x73, 0x0d, 0xaa, 0x62, 0xbd, 0xf8, 0x99, 0x40, 0xdf, 0x85, 0xfc, 0x05, 0x3b, 0x17, 0x6c, 0x2b, 0x96, 0x37, 0xef, 0x45, 0x16, 0x37, 0x74, 0x76, 0x74, 0xc1, 0x8b, 0x30, 0x64, 0x2f, 0x27, 0x6e, 0x23, 0xb3, 0x96, 0x5d, 0x2f, 0x6f, 0xd6, 0x37, 0xf8, 0x81, 0xdd, 0xd8, 0x27, 0xd7, 0xaf, 0x8d, 0xc1, 0x98, 0xe8, 0xb4, 0x13, 0x21, 0x98, 0x1d, 0xda, 0x0e, 0x61, 0x3b, 0xb6, 0xa8, 0xb3, 0x67, 0xba, 0x8d, 0xd9, 0xa2, 0x89, 0xdd, 0xca, 0x1b, 0xf8, 0x17, 0x1a, 0xc0, 0xd1, 0xd8, 0x4b, 0x3f, 0x1a, 0x8b, 0x90, 0x9b, 0x50, 0xc1, 0xe2, 0x58, 0xf0, 0x06, 0x3b, 0x13, 0xc4, 0x70, 0x89, 0x7f, 0x26, 0x68, 0x03, 0xdd, 0x81, 0xc2, 0xc8, 0x21, 0x93, 0xd3, 0xcb, 0x09, 0x53, 0x52, 0xd4, 0xf3, 0xb4, 0xb9, 0x3f, 0x41, 0xef, 0x41, 0xc5, 0x3c, 0xb7, 0x6c, 0x87, 0x9c, 0x72, 0x59, 0x39, 0xd6, 0x5b, 0xe6, 0x34, 0x66, 0xb7, 0xc2, 0xc2, 0x05, 0xe7, 0x55, 0x96, 0x03, 0x4a, 0xc2, 0x16, 0x94, 0x99, 0xa9, 0x53, 0xb9, 0xef, 0x83, 0xc0, 0xc6, 0x0c, 0x1b, 0x16, 0x77, 0xa1, 0xb0, 0x1a, 0xff, 0x10, 0xd0, 0x0e, 0x19, 0x10, 0x8f, 0x4c, 0x13, 0x3d, 0x14, 0x9f, 0x64, 0x55, 0x9f, 0xe0, 0x9f, 0x69, 0xb0, 0x10, 0x12, 0x3f, 0xd5, 0xb4, 0x1a, 0x50, 0xe8, 0x33, 0x61, 0xdc, 0x82, 0xac, 0x2e, 0x9b, 0xe8, 0x19, 0x14, 0x85, 0x01, 0x6e, 0x23, 0x9b, 0xb2, 0x69, 0x0a, 0xdc, 0x26, 0x17, 0xff, 0x22, 0x03, 0x25, 0x31, 0xd1, 0xce, 0x08, 0x6d, 0x41, 0xd5, 0xe1, 0x8d, 0x53, 0x36, 0x1f, 0x61, 0x51, 0x33, 0x3d, 0x08, 0xed, 0xce, 0xe8, 0x15, 0x31, 0x84, 0x91, 0xd1, 0xef, 0x41, 0x59, 0x8a, 0x18, 0x8d, 0x3d, 0xe1, 0xf2, 0x46, 0x58, 0x40, 0xb0, 0xff, 0x76, 0x67, 0x74, 0x10, 0xec, 0x47, 0x63, 0x0f, 0x75, 0x61, 0x51, 0x0e, 0xe6, 0xb3, 0x11, 0x66, 0x64, 0x99, 0x94, 0xb5, 0xb0, 0x94, 0xf8, 0x52, 0xed, 0xce, 0xe8, 0x48, 0x8c, 0x57, 0x3a, 0x55, 0x93, 0xbc, 0x2b, 0x1e, 0xbc, 0x63, 0x26, 0x75, 0xaf, 0xac, 0xb8, 0x49, 0xdd, 0x2b, 0xeb, 0x45, 0x09, 0x0a, 0xa2, 0x85, 0xff, 0x35, 0x03, 0x20, 0x57, 0xa3, 0x33, 0x42, 0x3b, 0x50, 0x73, 0x44, 0x2b, 0xe4, 0xad, 0xbb, 0x89, 0xde, 0x12, 0x8b, 0x38, 0xa3, 0x57, 0xe5, 0x20, 0x6e, 0xdc, 0xf7, 0xa1, 0xe2, 0x4b, 0x09, 0x1c, 0xb6, 0x92, 0xe0, 0x30, 0x5f, 0x42, 0x59, 0x0e, 0xa0, 0x2e, 0xfb, 0x0c, 0x96, 0xfc, 0xf1, 0x09, 0x3e, 0x7b, 0xef, 0x06, 0x9f, 0xf9, 0x02, 0x17, 0xa4, 0x04, 0xd5, 0x6b, 0xaa, 0x61, 0x81, 0xdb, 0x56, 0x12, 0xdc, 0x16, 0x37, 0x8c, 0x3a, 0x0e, 0x68, 0xbe, 0xe4, 0x4d, 0xfc, 0x7f, 0x59, 0x28, 0x6c, 0xdb, 0xc3, 0x91, 0xe1, 0xd0, 0xd5, 0xc8, 0x3b, 0xc4, 0x1d, 0x0f, 0x3c, 0xe6, 0xae, 0xda, 0xe6, 0xc3, 0xb0, 0x44, 0xc1, 0x26, 0xff, 0xeb, 0x8c, 0x55, 0x17, 0x43, 0xe8, 0x60, 0x91, 0x1e, 0x33, 0xef, 0x30, 0x58, 0x24, 0x47, 0x31, 0x44, 0x1e, 0xe4, 0x6c, 0x70, 0x90, 0x9b, 0x50, 0x98, 0x10, 0x27, 0x48, 0xe9, 0xbb, 0x33, 0xba, 0x24, 0xa0, 0x0f, 0x60, 0x2e, 0x9a, 0x5e, 0x72, 0x82, 0xa7, 0xd6, 0x0b, 0x67, 0xa3, 0x87, 0x50, 0x09, 0xe5, 0xb8, 0xbc, 0xe0, 0x2b, 0x0f, 0x95, 0x14, 0xb7, 0x2c, 0xe3, 0x2a, 0xcd, 0xc7, 0x95, 0xdd, 0x19, 0x19, 0x59, 0x97, 0x65, 0x64, 0x2d, 0x8a, 0x51, 0x22, 0xb6, 0x86, 0x82, 0xcc, 0x0f, 0xc2, 0x41, 0x06, 0xff, 0x00, 0xaa, 0x21, 0x07, 0xd1, 0xbc, 0xd3, 0xfe, 0xf4, 0x64, 0xeb, 0x80, 0x27, 0xa9, 0x97, 0x2c, 0x2f, 0xe9, 0x75, 0x8d, 0xe6, 0xba, 0x83, 0xf6, 0xf1, 0x71, 0x3d, 0x83, 0xaa, 0x50, 0x3a, 0xec, 0x74, 0x4f, 0x39, 0x57, 0x16, 0xbf, 0xf4, 0x25, 0x88, 0x24, 0xa7, 0xe4, 0xb6, 0x19, 0x25, 0xb7, 0x69, 0x32, 0xb7, 0x65, 0x82, 0xdc, 0xc6, 0xd2, 0xdc, 0x41, 0x7b, 0xeb, 0xb8, 0x5d, 0x9f, 0x7d, 0x51, 0x83, 0x0a, 0xf7, 0xef, 0xe9, 0xd8, 0xa2, 0xa9, 0xf6, 0x1f, 0x34, 0x80, 0xe0, 0x34, 0xa1, 0x16, 0x14, 0x7a, 0x5c, 0x4f, 0x43, 0x63, 0xc1, 0x68, 0x29, 0x71, 0xc9, 0x74, 0xc9, 0x85, 0xbe, 0x0d, 0x05, 0x77, 0xdc, 0xeb, 0x11, 0x57, 0xa6, 0xbc, 0x3b, 0xd1, 0x78, 0x28, 0xa2, 0x95, 0x2e, 0xf9, 0xe8, 0x90, 0x37, 0x86, 0x39, 0x18, 0xb3, 0x04, 0x78, 0xf3, 0x10, 0xc1, 0x87, 0xff, 0x4e, 0x83, 0xb2, 0xb2, 0x79, 0x7f, 0xcd, 0x20, 0x7c, 0x0f, 0x4a, 0xcc, 0x06, 0xd2, 0x17, 0x61, 0xb8, 0xa8, 0x07, 0x04, 0xf4, 0x3b, 0x50, 0x92, 0x27, 0x40, 0x46, 0xe2, 0x46, 0xb2, 0xd8, 0xce, 0x48, 0x0f, 0x58, 0xf1, 0x3e, 0xcc, 0x33, 0xaf, 0xf4, 0x68, 0x71, 0x2d, 0xfd, 0xa8, 0x96, 0x9f, 0x5a, 0xa4, 0xfc, 0x6c, 0x42, 0x71, 0x74, 0x71, 0xed, 0x9a, 0x3d, 0x63, 0x20, 0xac, 0xf0, 0xdb, 0xf8, 0x13, 0x40, 0xaa, 0xb0, 0x69, 0xa6, 0x8b, 0xab, 0x50, 0xde, 0x35, 0xdc, 0x0b, 0x61, 0x12, 0x7e, 0x06, 0x55, 0xda, 0xdc, 0x7f, 0xfd, 0x0e, 0x36, 0xb2, 0xcb, 0x81, 0xe4, 0x9e, 0xca, 0xe7, 0x08, 0x66, 0x2f, 0x0c, 0xf7, 0x82, 0x4d, 0xb4, 0xaa, 0xb3, 0x67, 0xf4, 0x01, 0xd4, 0x7b, 0x7c, 0x92, 0xa7, 0x91, 0x2b, 0xc3, 0x9c, 0xa0, 0xfb, 0x95, 0xe0, 0xe7, 0x50, 0xe1, 0x73, 0xf8, 0x4d, 0x1b, 0x81, 0xe7, 0x61, 0xee, 0xd8, 0x32, 0x46, 0xee, 0x85, 0x2d, 0xb3, 0x1b, 0x9d, 0x74, 0x3d, 0xa0, 0x4d, 0xa5, 0xf1, 0x09, 0xcc, 0x39, 0x64, 0x68, 0x98, 0x96, 0x69, 0x9d, 0x9f, 0x9e, 0x5d, 0x7b, 0xc4, 0x15, 0x17, 0xa6, 0x9a, 0x4f, 0x7e, 0x41, 0xa9, 0xd4, 0xb4, 0xb3, 0x81, 0x7d, 0x26, 0xc2, 0x1c, 0x7b, 0xc6, 0x7f, 0x99, 0x81, 0xca, 0x67, 0x86, 0xd7, 0x93, 0x4b, 0x87, 0xf6, 0xa0, 0xe6, 0x07, 0x37, 0x46, 0x11, 0xb6, 0x44, 0x52, 0x2c, 0x1b, 0x23, 0x4b, 0x69, 0x99, 0x1d, 0xab, 0x3d, 0x95, 0xc0, 0x44, 0x19, 0x56, 0x8f, 0x0c, 0x7c, 0x51, 0x99, 0x74, 0x51, 0x8c, 0x51, 0x15, 0xa5, 0x12, 0x50, 0x07, 0xea, 0x23, 0xc7, 0x3e, 0x77, 0x88, 0xeb, 0xfa, 0xc2, 0x78, 0x1a, 0xc3, 0x09, 0xc2, 0x8e, 0x04, 0x6b, 0x20, 0x6e, 0x6e, 0x14, 0x26, 0xbd, 0x98, 0x0b, 0xea, 0x19, 0x1e, 0x9c, 0xfe, 0x2b, 0x03, 0x28, 0x3e, 0xa9, 0x5f, 0xb5, 0xc4, 0x7b, 0x0c, 0x35, 0xd7, 0x33, 0x9c, 0xd8, 0x66, 0xab, 0x32, 0xaa, 0x1f, 0xf1, 0x9f, 0x80, 0x6f, 0xd0, 0xa9, 0x65, 0x7b, 0xe6, 0x9b, 0x6b, 0x51, 0x25, 0xd7, 0x24, 0xf9, 0x90, 0x51, 0x51, 0x1b, 0x0a, 0x6f, 0xcc, 0x81, 0x47, 0x1c, 0xb7, 0x91, 0x5b, 0xcb, 0xae, 0xd7, 0x36, 0x9f, 0xdd, 0xb6, 0x0c, 0x1b, 0x1f, 0x33, 0xfe, 0xee, 0xf5, 0x88, 0xe8, 0x72, 0xac, 0x5a, 0x79, 0xe6, 0x43, 0xd5, 0xf8, 0x0a, 0x14, 0xdf, 0x52, 0x11, 0xf4, 0x96, 0x5d, 0xe0, 0xc5, 0x22, 0x6b, 0xf3, 0x4b, 0xf6, 0x1b, 0xc7, 0x38, 0x1f, 0x12, 0xcb, 0x93, 0xf7, 0x40, 0xd9, 0xc6, 0x8f, 0x01, 0x02, 0x35, 0x34, 0xe4, 0x1f, 0x76, 0x8e, 0x4e, 0xba, 0xf5, 0x19, 0x54, 0x81, 0xe2, 0x61, 0x67, 0xa7, 0x7d, 0xd0, 0xa6, 0xf9, 0x01, 0xb7, 0xa4, 0x4b, 0x43, 0x6b, 0xa9, 0xea, 0xd4, 0x42, 0x3a, 0xf1, 0x32, 0x2c, 0x26, 0x2d, 0x20, 0xad, 0x45, 0xab, 0x62, 0x97, 0x4e, 0x75, 0x54, 0x54, 0xd5, 0x99, 0xf0, 0x74, 0x1b, 0x50, 0xe0, 0xbb, 0xb7, 0x2f, 0x8a, 0x73, 0xd9, 0xa4, 0x8e, 0xe0, 0x9b, 0x91, 0xf4, 0xc5, 0x2a, 0xf9, 0xed, 0xc4, 0xf0, 0x92, 0x4b, 0x0c, 0x2f, 0xe8, 0x21, 0x54, 0xfd, 0xd3, 0x60, 0xb8, 0xa2, 0x16, 0x28, 0xe9, 0x15, 0xb9, 0xd1, 0x29, 0x2d, 0xe4, 0xf4, 0x42, 0xd8, 0xe9, 0xe8, 0x31, 0xe4, 0xc9, 0x84, 0x58, 0x9e, 0xdb, 0x28, 0xb3, 0x8c, 0x51, 0x95, 0xb5, 0x7b, 0x9b, 0x52, 0x75, 0xd1, 0x89, 0xbf, 0x07, 0xf3, 0xec, 0x8e, 0xf4, 0xd2, 0x31, 0x2c, 0xf5, 0x32, 0xd7, 0xed, 0x1e, 0x08, 0x77, 0xd3, 0x47, 0x54, 0x83, 0xcc, 0xde, 0x8e, 0x70, 0x42, 0x66, 0x6f, 0x07, 0xff, 0x44, 0x03, 0xa4, 0x8e, 0x9b, 0xca, 0xcf, 0x11, 0xe1, 0x52, 0x7d, 0x36, 0x50, 0xbf, 0x08, 0x39, 0xe2, 0x38, 0xb6, 0xc3, 0x3c, 0x5a, 0xd2, 0x79, 0x03, 0x3f, 0x12, 0x36, 0xe8, 0x64, 0x62, 0x5f, 0xfa, 0x67, 0x90, 0x4b, 0xd3, 0x7c, 0x53, 0xf7, 0x61, 0x21, 0xc4, 0x35, 0x55, 0xe6, 0xfa, 0x18, 0xe6, 0x98, 0xb0, 0xed, 0x0b, 0xd2, 0xbb, 0x1c, 0xd9, 0xa6, 0x15, 0xd3, 0x47, 0x57, 0x2e, 0x08, 0xb0, 0x74, 0x1e, 0x7c, 0x62, 0x15, 0x9f, 0xd8, 0xed, 0x1e, 0xe0, 0x2f, 0x60, 0x39, 0x22, 0x47, 0x9a, 0xff, 0x87, 0x50, 0xee, 0xf9, 0x44, 0x57, 0xd4, 0x3a, 0xf7, 0xc3, 0xc6, 0x45, 0x87, 0xaa, 0x23, 0x70, 0x07, 0xee, 0xc4, 0x44, 0x4f, 0x35, 0xe7, 0x27, 0xb0, 0xc4, 0x04, 0xee, 0x13, 0x32, 0xda, 0x1a, 0x98, 0x93, 0x54, 0x4f, 0x8f, 0xc4, 0xa4, 0x14, 0xc6, 0x6f, 0x76, 0x5f, 0xe0, 0xdf, 0x17, 0x1a, 0xbb, 0xe6, 0x90, 0x74, 0xed, 0x83, 0x74, 0xdb, 0x68, 0x36, 0xbb, 0x24, 0xd7, 0xae, 0x28, 0x6b, 0xd8, 0x33, 0xfe, 0x47, 0x4d, 0xb8, 0x4a, 0x1d, 0xfe, 0x0d, 0xef, 0xe4, 0x55, 0x80, 0x73, 0x7a, 0x64, 0x48, 0x9f, 0x76, 0x70, 0x44, 0x45, 0xa1, 0xf8, 0x76, 0xd2, 0xf8, 0x5d, 0x11, 0x76, 0x2e, 0x8a, 0x7d, 0xce, 0xfe, 0xf8, 0x51, 0xee, 0x3e, 0x94, 0x19, 0xe1, 0xd8, 0x33, 0xbc, 0xb1, 0x1b, 0x5b, 0x8c, 0x3f, 0x13, 0xdb, 0x5e, 0x0e, 0x9a, 0x6a, 0x5e, 0xdf, 0x86, 0x3c, 0xbb, 0x4c, 0xc8, 0x52, 0x7a, 0x25, 0x61, 0x3f, 0x72, 0x3b, 0x74, 0xc1, 0x88, 0x2f, 0x20, 0xff, 0x8a, 0x21, 0xb0, 0x8a, 0x65, 0xb3, 0x72, 0x29, 0x2c, 0x63, 0xc8, 0x71, 0xa1, 0x92, 0xce, 0x9e, 0x59, 0xe5, 0x49, 0x88, 0x73, 0xa2, 0x1f, 0xf0, 0x0a, 0xb7, 0xa4, 0xfb, 0x6d, 0xea, 0xb2, 0xde, 0xc0, 0x24, 0x96, 0xc7, 0x7a, 0x67, 0x59, 0xaf, 0x42, 0xc1, 0x1b, 0x50, 0xe7, 0x9a, 0xb6, 0xfa, 0x7d, 0xa5, 0x82, 0xf4, 0xe5, 0x69, 0x61, 0x79, 0xf8, 0x9f, 0x34, 0x98, 0x57, 0x06, 0x4c, 0xe5, 0x98, 0x0f, 0x21, 0xcf, 0x71, 0x66, 0x51, 0xac, 0x2c, 0x86, 0x47, 0x71, 0x35, 0xba, 0xe0, 0x41, 0x1b, 0x50, 0xe0, 0x4f, 0xb2, 0x8c, 0x4f, 0x66, 0x97, 0x4c, 0xf8, 0x31, 0x2c, 0x08, 0x12, 0x19, 0xda, 0x49, 0x7b, 0x9b, 0x39, 0x14, 0xff, 0x29, 0x2c, 0x86, 0xd9, 0xa6, 0x9a, 0x92, 0x62, 0x64, 0xe6, 0x5d, 0x8c, 0xdc, 0x92, 0x46, 0x9e, 0x8c, 0xfa, 0x4a, 0x29, 0x14, 0x5d, 0x75, 0x75, 0x45, 0x32, 0x91, 0x15, 0xf1, 0x27, 0x20, 0x45, 0xfc, 0x56, 0x27, 0xb0, 0x20, 0xb7, 0xc3, 0x81, 0xe9, 0xfa, 0x15, 0xf7, 0x57, 0x80, 0x54, 0xe2, 0x6f, 0xdb, 0xa0, 0x1d, 0x22, 0x13, 0xb9, 0x34, 0xe8, 0x13, 0x40, 0x2a, 0x71, 0xaa, 0x88, 0xde, 0x82, 0xf9, 0x57, 0xf6, 0x84, 0x86, 0x06, 0x4a, 0x0d, 0x8e, 0x0c, 0xbf, 0x7f, 0xfb, 0xcb, 0xe6, 0xb7, 0xa9, 0x72, 0x75, 0xc0, 0x54, 0xca, 0xff, 0x43, 0x83, 0xca, 0xd6, 0xc0, 0x70, 0x86, 0x52, 0xf1, 0xf7, 0x21, 0xcf, 0x6f, 0x95, 0x02, 0xc8, 0x79, 0x3f, 0x2c, 0x46, 0xe5, 0xe5, 0x8d, 0x2d, 0x7e, 0x07, 0x15, 0xa3, 0xa8, 0xe1, 0xe2, 0x5d, 0xcf, 0x4e, 0xe4, 0xdd, 0xcf, 0x0e, 0xfa, 0x08, 0x72, 0x06, 0x1d, 0xc2, 0x42, 0x70, 0x2d, 0x7a, 0x9f, 0x67, 0xd2, 0x58, 0xed, 0xcb, 0xb9, 0xf0, 0x77, 0xa1, 0xac, 0x68, 0x40, 0x05, 0xc8, 0xbe, 0x6c, 0x8b, 0x42, 0x75, 0x6b, 0xbb, 0xbb, 0xf7, 0x9a, 0x03, 0x19, 0x35, 0x80, 0x9d, 0xb6, 0xdf, 0xce, 0xe0, 0xcf, 0xc5, 0x28, 0x11, 0xef, 0x54, 0x7b, 0xb4, 0x34, 0x7b, 0x32, 0xef, 0x64, 0xcf, 0x15, 0x54, 0xc5, 0xf4, 0xa7, 0x0d, 0xdf, 0x4c, 0x5e, 0x4a, 0xf8, 0x56, 0x8c, 0xd7, 0x05, 0x23, 0x9e, 0x83, 0xaa, 0x08, 0xe8, 0x62, 0xff, 0xfd, 0x4b, 0x06, 0x6a, 0x92, 0x32, 0x2d, 0xe0, 0x2c, 0xb1, 0x32, 0x9e, 0x01, 0x7c, 0xa4, 0x6c, 0x19, 0xf2, 0xfd, 0xb3, 0x63, 0xf3, 0x2b, 0xf9, 0x72, 0x40, 0xb4, 0x28, 0x7d, 0xc0, 0xf5, 0xf0, 0x37, 0x74, 0xa2, 0x85, 0xee, 0xf1, 0x97, 0x77, 0x7b, 0x56, 0x9f, 0x5c, 0xb1, 0x3a, 0x7a, 0x56, 0x0f, 0x08, 0x0c, 0x44, 0x10, 0x6f, 0xf2, 0x58, 0xf1, 0xac, 0xbc, 0xd9, 0x43, 0x4f, 0xa1, 0x4e, 0x9f, 0xb7, 0x46, 0xa3, 0x81, 0x49, 0xfa, 0x5c, 0x40, 0x81, 0xf1, 0xc4, 0xe8, 0x54, 0x3b, 0x2b, 0x37, 0xdd, 0x46, 0x91, 0x85, 0x2d, 0xd1, 0x42, 0x6b, 0x50, 0xe6, 0xf6, 0xed, 0x59, 0x27, 0x2e, 0x61, 0xaf, 0xb7, 0xb2, 0xba, 0x4a, 0xa2, 0xe7, 0x78, 0x6b, 0xec, 0x5d, 0xb4, 0x2d, 0xe3, 0x6c, 0x20, 0xe3, 0x22, 0x4d, 0xe6, 0x94, 0xb8, 0x63, 0xba, 0x2a, 0xb5, 0x0d, 0x0b, 0x94, 0x4a, 0x2c, 0xcf, 0xec, 0x29, 0x41, 0x54, 0xa6, 0x4a, 0x2d, 0x92, 0x2a, 0x0d, 0xd7, 0x7d, 0x6b, 0x3b, 0x7d, 0xe1, 0x40, 0xbf, 0x8d, 0x77, 0xb8, 0xf0, 0x13, 0x37, 0x94, 0x0c, 0x7f, 0x55, 0x29, 0xeb, 0x81, 0x94, 0x97, 0xc4, 0xbb, 0x41, 0x0a, 0x7e, 0x06, 0x4b, 0x92, 0x53, 0x40, 0xbe, 0x37, 0x30, 0x77, 0xe0, 0xbe, 0x64, 0xde, 0xbe, 0xa0, 0x57, 0xe0, 0x23, 0xa1, 0xf0, 0xd7, 0xb5, 0xf3, 0x05, 0x34, 0x7c, 0x3b, 0xd9, 0x35, 0xc4, 0x1e, 0xa8, 0x06, 0x8c, 0x5d, 0xb1, 0x33, 0x4b, 0x3a, 0x7b, 0xa6, 0x34, 0xc7, 0x1e, 0xf8, 0x85, 0x07, 0x7d, 0xc6, 0xdb, 0xb0, 0x22, 0x65, 0x88, 0x0b, 0x42, 0x58, 0x48, 0xcc, 0xa0, 0x24, 0x21, 0xc2, 0x61, 0x74, 0xe8, 0xcd, 0x6e, 0x57, 0x39, 0xc3, 0xae, 0x65, 0x32, 0x35, 0x45, 0xe6, 0x12, 0xdf, 0x11, 0xd4, 0x30, 0x35, 0x2f, 0x09, 0x32, 0x15, 0xa0, 0x92, 0xc5, 0x42, 0x50, 0x72, 0x6c, 0x21, 0x62, 0xa2, 0x7f, 0x08, 0xab, 0xbe, 0x11, 0xd4, 0x6f, 0x47, 0xc4, 0x19, 0x9a, 0xae, 0xab, 0x80, 0x84, 0x49, 0x13, 0x7f, 0x1f, 0x66, 0x47, 0x44, 0x44, 0xae, 0xf2, 0x26, 0xda, 0xe0, 0x6f, 0xf5, 0x37, 0x94, 0xc1, 0xac, 0x1f, 0xf7, 0xe1, 0x81, 0x94, 0xce, 0x3d, 0x9a, 0x28, 0x3e, 0x6a, 0x94, 0x84, 0x4e, 0x32, 0x29, 0xd0, 0x49, 0x36, 0x02, 0x5c, 0x7f, 0xc2, 0x1d, 0x29, 0xcf, 0xd6, 0x54, 0x19, 0x69, 0x9f, 0xfb, 0xd4, 0x3f, 0x92, 0x53, 0x09, 0x3b, 0x83, 0xc5, 0xf0, 0x49, 0x9e, 0x2a, 0x58, 0x2e, 0x42, 0xce, 0xb3, 0x2f, 0x89, 0x0c, 0x95, 0xbc, 0x21, 0x0d, 0xf6, 0x8f, 0xf9, 0x54, 0x06, 0x1b, 0x81, 0x30, 0xb6, 0x25, 0xa7, 0xb5, 0x97, 0xae, 0xa6, 0x2c, 0xf1, 0x78, 0x03, 0x1f, 0xc2, 0x72, 0x34, 0x4c, 0x4c, 0x65, 0xf2, 0x6b, 0xbe, 0x81, 0x93, 0x22, 0xc9, 0x54, 0x72, 0x3f, 0x0d, 0x82, 0x81, 0x12, 0x50, 0xa6, 0x12, 0xa9, 0x43, 0x33, 0x29, 0xbe, 0xfc, 0x26, 0xf6, 0xab, 0x1f, 0x6e, 0xa6, 0x12, 0xe6, 0x06, 0xc2, 0xa6, 0x5f, 0xfe, 0x20, 0x46, 0x64, 0x6f, 0x8c, 0x11, 0xe2, 0x90, 0x04, 0x51, 0xec, 0x1b, 0xd8, 0x74, 0x42, 0x47, 0x10, 0x40, 0xa7, 0xd5, 0x41, 0x73, 0x88, 0xaf, 0x83, 0x35, 0xe4, 0xc6, 0x56, 0xc3, 0xee, 0x54, 0x8b, 0xf1, 0x59, 0x10, 0x3b, 0x63, 0x91, 0x79, 0x2a, 0xc1, 0x9f, 0xc3, 0x5a, 0x7a, 0x50, 0x9e, 0x46, 0xf2, 0xd3, 0x16, 0x94, 0xfc, 0xb2, 0x55, 0xf9, 0x22, 0xa6, 0x0c, 0x85, 0xc3, 0xce, 0xf1, 0xd1, 0xd6, 0x76, 0x9b, 0x7f, 0x12, 0xb3, 0xdd, 0xd1, 0xf5, 0x93, 0xa3, 0x6e, 0x3d, 0xb3, 0xf9, 0xcb, 0x2c, 0x64, 0xf6, 0x5f, 0xa3, 0x2f, 0x20, 0xc7, 0xdf, 0x0f, 0xdf, 0xf0, 0x51, 0x40, 0xf3, 0xa6, 0x57, 0xe0, 0xf8, 0xce, 0x4f, 0xfe, 0xfb, 0x97, 0x3f, 0xcf, 0xcc, 0xe3, 0x4a, 0x6b, 0xf2, 0x9d, 0xd6, 0xe5, 0xa4, 0xc5, 0x72, 0xc3, 0x73, 0xed, 0x29, 0xfa, 0x14, 0xb2, 0x47, 0x63, 0x0f, 0xa5, 0x7e, 0x2c, 0xd0, 0x4c, 0x7f, 0x2b, 0x8e, 0x97, 0x98, 0xd0, 0x39, 0x0c, 0x42, 0xe8, 0x68, 0xec, 0x51, 0x91, 0x3f, 0x82, 0xb2, 0xfa, 0x4e, 0xfb, 0xd6, 0x2f, 0x08, 0x9a, 0xb7, 0xbf, 0x2f, 0xc7, 0xf7, 0x99, 0xaa, 0x3b, 0x18, 0x09, 0x55, 0xfc, 0xad, 0xbb, 0x3a, 0x8b, 0xee, 0x95, 0x85, 0x52, 0xbf, 0x2f, 0x68, 0xa6, 0xbf, 0x42, 0x8f, 0xcd, 0xc2, 0xbb, 0xb2, 0xa8, 0xc8, 0x3f, 0x16, 0x6f, 0xcf, 0x7b, 0x1e, 0x7a, 0x90, 0xf0, 0xf6, 0x54, 0x7d, 0x4f, 0xd8, 0x5c, 0x4b, 0x67, 0x10, 0x4a, 0xee, 0x31, 0x25, 0xcb, 0x78, 0x5e, 0x28, 0xe9, 0xf9, 0x2c, 0xcf, 0xb5, 0xa7, 0x9b, 0x3d, 0xc8, 0x31, 0x0c, 0x1e, 0x7d, 0x29, 0x1f, 0x9a, 0x09, 0x2f, 0x23, 0x52, 0x16, 0x3a, 0x84, 0xde, 0xe3, 0x45, 0xa6, 0xa8, 0x86, 0x4b, 0x54, 0x11, 0x43, 0xe0, 0x9f, 0x6b, 0x4f, 0xd7, 0xb5, 0x6f, 0x69, 0x9b, 0xff, 0x9c, 0x83, 0x1c, 0x03, 0x9f, 0xd0, 0x25, 0x40, 0x80, 0x47, 0x47, 0x67, 0x17, 0x43, 0xb8, 0xa3, 0xb3, 0x8b, 0x43, 0xd9, 0xb8, 0xc9, 0x94, 0x2e, 0xe2, 0x39, 0xaa, 0x94, 0x61, 0x5a, 0x2d, 0x06, 0xd3, 0x51, 0x3f, 0xfe, 0x95, 0x26, 0xb0, 0x37, 0x7e, 0x96, 0x50, 0x92, 0xb4, 0x10, 0x28, 0x1d, 0xdd, 0x0e, 0x09, 0x80, 0x34, 0xfe, 0x1e, 0x53, 0xd8, 0xc2, 0xf5, 0x40, 0xa1, 0xc3, 0x38, 0x9e, 0x6b, 0x4f, 0xbf, 0x6c, 0xe0, 0x05, 0xe1, 0xe5, 0x48, 0x0f, 0xfa, 0x31, 0xd4, 0xc2, 0xa0, 0x2b, 0x7a, 0x98, 0xa0, 0x2b, 0x8a, 0xdd, 0x36, 0x1f, 0xdd, 0xcc, 0x24, 0x6c, 0x5a, 0x65, 0x36, 0x09, 0xe5, 0x5c, 0xf3, 0x25, 0x21, 0x23, 0x83, 0x32, 0x89, 0x35, 0x40, 0x7f, 0xaf, 0x09, 0x4c, 0x3c, 0x40, 0x51, 0x51, 0x92, 0xf4, 0x18, 0x46, 0xdb, 0x7c, 0x7c, 0x0b, 0x97, 0x30, 0xe2, 0x0f, 0x98, 0x11, 0xbf, 0x8b, 0x17, 0x03, 0x23, 0x3c, 0x73, 0x48, 0x3c, 0x5b, 0x58, 0xf1, 0xe5, 0x3d, 0x7c, 0x27, 0xe4, 0x9c, 0x50, 0x6f, 0xb0, 0x58, 0x1c, 0x09, 0x4d, 0x5c, 0xac, 0x10, 0xb2, 0x9a, 0xb8, 0x58, 0x61, 0x18, 0x35, 0x69, 0xb1, 0x38, 0xee, 0x99, 0xb4, 0x58, 0x7e, 0xcf, 0x26, 0xfb, 0x7e, 0x85, 0x7f, 0xb5, 0x8a, 0x6c, 0x28, 0xf9, 0x28, 0x24, 0x5a, 0x4d, 0x42, 0x84, 0x82, 0xbb, 0x44, 0xf3, 0x41, 0x6a, 0xbf, 0x30, 0xe8, 0x3d, 0x66, 0xd0, 0x5d, 0xbc, 0x4c, 0x35, 0x8b, 0x0f, 0x63, 0x5b, 0x1c, 0x76, 0x68, 0x19, 0xfd, 0x3e, 0x75, 0xc4, 0x9f, 0x40, 0x45, 0x85, 0x09, 0xd1, 0x7b, 0x89, 0x28, 0x94, 0x8a, 0x34, 0x36, 0xf1, 0x4d, 0x2c, 0x42, 0xf3, 0x23, 0xa6, 0x79, 0x15, 0xaf, 0x24, 0x68, 0x76, 0x18, 0x6b, 0x48, 0x39, 0x87, 0xf8, 0x92, 0x95, 0x87, 0x10, 0xc4, 0x64, 0xe5, 0x61, 0x84, 0xf0, 0x46, 0xe5, 0x63, 0xc6, 0x4a, 0x95, 0xbb, 0x00, 0x01, 0x98, 0x87, 0x12, 0x7d, 0xa9, 0x5c, 0xa6, 0xa2, 0xc1, 0x21, 0x8e, 0x03, 0x62, 0xcc, 0xd4, 0x8a, 0x7d, 0x17, 0x51, 0x3b, 0x30, 0x5d, 0x1a, 0x24, 0x36, 0xff, 0x3a, 0x0f, 0xe5, 0x57, 0x86, 0x69, 0x79, 0xc4, 0x32, 0xac, 0x1e, 0x41, 0x67, 0x90, 0x63, 0x89, 0x32, 0x1a, 0x07, 0x55, 0x7c, 0x2b, 0x1a, 0x07, 0x43, 0xe0, 0x0f, 0x5e, 0x63, 0x5a, 0x9b, 0x78, 0x89, 0x6a, 0x1d, 0x06, 0xa2, 0x5b, 0x0c, 0xb3, 0xa1, 0x13, 0x7d, 0x03, 0x79, 0xf1, 0x3a, 0x20, 0x22, 0x28, 0x84, 0xe5, 0x34, 0xef, 0x25, 0x77, 0x26, 0x6d, 0x25, 0x55, 0x8d, 0xcb, 0xf8, 0xa8, 0x9e, 0x09, 0x40, 0x00, 0x46, 0x46, 0x1d, 0x1a, 0xc3, 0x2e, 0x9b, 0x6b, 0xe9, 0x0c, 0x42, 0xe7, 0x63, 0xa6, 0xf3, 0x01, 0x6e, 0x46, 0x75, 0xf6, 0x7d, 0x5e, 0xaa, 0xf7, 0x8f, 0x60, 0x76, 0xd7, 0x70, 0x2f, 0x50, 0x24, 0xf5, 0x29, 0x1f, 0x93, 0x34, 0x9b, 0x49, 0x5d, 0x42, 0xcb, 0x03, 0xa6, 0x65, 0x85, 0x47, 0x12, 0x55, 0xcb, 0x85, 0xe1, 0xd2, 0x9c, 0x82, 0xfa, 0x90, 0xe7, 0xdf, 0x96, 0x44, 0xfd, 0x17, 0xfa, 0x3e, 0x25, 0xea, 0xbf, 0xf0, 0xe7, 0x28, 0xb7, 0x6b, 0x19, 0x41, 0x51, 0x7e, 0xcc, 0x81, 0x22, 0x6f, 0xf6, 0x22, 0x1f, 0x7e, 0x34, 0x57, 0xd3, 0xba, 0x85, 0xae, 0x87, 0x4c, 0xd7, 0x7d, 0xdc, 0x88, 0xad, 0x95, 0xe0, 0x7c, 0xae, 0x3d, 0xfd, 0x96, 0x86, 0x7e, 0x0c, 0x10, 0xe0, 0xb7, 0xb1, 0x03, 0x10, 0x85, 0x82, 0x63, 0x07, 0x20, 0x06, 0xfd, 0xe2, 0x0d, 0xa6, 0x77, 0x1d, 0x3f, 0x8c, 0xea, 0xf5, 0x1c, 0xc3, 0x72, 0xdf, 0x10, 0xe7, 0x23, 0x8e, 0xd1, 0xb9, 0x17, 0xe6, 0x88, 0x1e, 0x86, 0x7f, 0x9b, 0x83, 0x59, 0x5a, 0x80, 0xd2, 0x3c, 0x1d, 0xdc, 0xdb, 0xa3, 0x96, 0xc4, 0xd0, 0xb2, 0xa8, 0x25, 0xf1, 0x2b, 0x7f, 0x38, 0x4f, 0xb3, 0x9f, 0x1b, 0x10, 0xc6, 0x40, 0x1d, 0x6d, 0x43, 0x59, 0xb9, 0xd8, 0xa3, 0x04, 0x61, 0x61, 0x18, 0x2e, 0x1a, 0xf9, 0x13, 0x50, 0x01, 0x7c, 0x97, 0xe9, 0x5b, 0xe2, 0x91, 0x9f, 0xe9, 0xeb, 0x73, 0x0e, 0xaa, 0xf0, 0x2d, 0x54, 0xd4, 0xcb, 0x3f, 0x4a, 0x90, 0x17, 0x81, 0xf8, 0xa2, 0x51, 0x2e, 0x09, 0x3b, 0x08, 0x1f, 0x7c, 0xff, 0x27, 0x15, 0x92, 0x8d, 0x2a, 0x1e, 0x40, 0x41, 0xa0, 0x01, 0x49, 0xb3, 0x0c, 0xe3, 0x81, 0x49, 0xb3, 0x8c, 0x40, 0x09, 0xe1, 0xda, 0x8e, 0x69, 0xa4, 0x17, 0x1e, 0x99, 0x49, 0x84, 0xb6, 0x97, 0xc4, 0x4b, 0xd3, 0x16, 0x80, 0x5b, 0x69, 0xda, 0x94, 0xcb, 0x66, 0x9a, 0xb6, 0x73, 0xe2, 0x89, 0xe3, 0x22, 0x2f, 0x71, 0x28, 0x45, 0x98, 0x1a, 0xbd, 0xf1, 0x4d, 0x2c, 0x49, 0xa5, 0x77, 0xa0, 0x50, 0x84, 0x6e, 0x74, 0x05, 0x10, 0x60, 0x15, 0xd1, 0x7a, 0x2a, 0x11, 0xf0, 0x8c, 0xd6, 0x53, 0xc9, 0x70, 0x47, 0x38, 0x34, 0x04, 0x7a, 0x79, 0xe5, 0x4f, 0x35, 0xff, 0x4c, 0x03, 0x14, 0x87, 0x35, 0xd0, 0xb3, 0x64, 0xe9, 0x89, 0x30, 0x6a, 0xf3, 0xc3, 0x77, 0x63, 0x4e, 0x8a, 0xf6, 0x81, 0x49, 0x3d, 0xc6, 0x3d, 0x7a, 0x4b, 0x8d, 0xfa, 0x73, 0x0d, 0xaa, 0x21, 0x4c, 0x04, 0xbd, 0x9f, 0xb2, 0xa6, 0x11, 0x14, 0xb6, 0xf9, 0xe4, 0x56, 0xbe, 0xa4, 0x42, 0x53, 0xd9, 0x01, 0xb2, 0xe2, 0xfe, 0xa9, 0x06, 0xb5, 0x30, 0x86, 0x82, 0x52, 0x64, 0xc7, 0x50, 0xdc, 0xe6, 0xfa, 0xed, 0x8c, 0x37, 0x2f, 0x4f, 0x50, 0x6c, 0x0f, 0xa0, 0x20, 0x50, 0x97, 0xa4, 0x8d, 0x1f, 0xc6, 0x7f, 0x93, 0x36, 0x7e, 0x04, 0xb2, 0x49, 0xd8, 0xf8, 0x8e, 0x3d, 0x20, 0xca, 0x31, 0x13, 0xb0, 0x4c, 0x9a, 0xb6, 0x9b, 0x8f, 0x59, 0x04, 0xd3, 0x49, 0xd3, 0x16, 0x1c, 0x33, 0x89, 0xc7, 0xa0, 0x14, 0x61, 0xb7, 0x1c, 0xb3, 0x28, 0x9c, 0x93, 0x70, 0xcc, 0x98, 0x42, 0xe5, 0x98, 0x05, 0xc8, 0x49, 0xd2, 0x31, 0x8b, 0xc1, 0xd9, 0x49, 0xc7, 0x2c, 0x0e, 0xbe, 0x24, 0xac, 0x23, 0xd3, 0x1b, 0x3a, 0x66, 0x0b, 0x09, 0x20, 0x0b, 0xfa, 0x30, 0xc5, 0x89, 0x89, 0x28, 0x79, 0xf3, 0xa3, 0x77, 0xe4, 0x4e, 0xdd, 0xe3, 0xdc, 0xfd, 0x72, 0x8f, 0xff, 0x8d, 0x06, 0x8b, 0x49, 0x00, 0x0d, 0x4a, 0xd1, 0x93, 0x82, 0xae, 0x37, 0x37, 0xde, 0x95, 0xfd, 0x66, 0x6f, 0xf9, 0xbb, 0xfe, 0x45, 0xfd, 0xdf, 0xbf, 0x5e, 0xd5, 0xfe, 0xf3, 0xeb, 0x55, 0xed, 0x7f, 0xbe, 0x5e, 0xd5, 0xfe, 0xf6, 0x7f, 0x57, 0x67, 0xce, 0xf2, 0xec, 0x87, 0x7a, 0xdf, 0xf9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xc3, 0xa2, 0xb2, 0x2f, 0x38, 0x00, 0x00}
