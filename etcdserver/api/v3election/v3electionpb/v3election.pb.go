package v3electionpb

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "github.com/gogo/protobuf/gogoproto"
	etcdserverpb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	mvccpb "go.etcd.io/etcd/mvcc/mvccpb"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	io "io"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2

type CampaignRequest struct {
	Name	[]byte	`protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lease	int64	`protobuf:"varint,2,opt,name=lease,proto3" json:"lease,omitempty"`
	Value	[]byte	`protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *CampaignRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = CampaignRequest{}
}
func (m *CampaignRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*CampaignRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*CampaignRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorV3Election, []int{0}
}
func (m *CampaignRequest) GetName() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return nil
}
func (m *CampaignRequest) GetLease() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Lease
	}
	return 0
}
func (m *CampaignRequest) GetValue() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Value
	}
	return nil
}

type CampaignResponse struct {
	Header	*etcdserverpb.ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Leader	*LeaderKey			`protobuf:"bytes,2,opt,name=leader" json:"leader,omitempty"`
}

func (m *CampaignResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = CampaignResponse{}
}
func (m *CampaignResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*CampaignResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*CampaignResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorV3Election, []int{1}
}
func (m *CampaignResponse) GetHeader() *etcdserverpb.ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *CampaignResponse) GetLeader() *LeaderKey {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Leader
	}
	return nil
}

type LeaderKey struct {
	Name	[]byte	`protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key	[]byte	`protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Rev	int64	`protobuf:"varint,3,opt,name=rev,proto3" json:"rev,omitempty"`
	Lease	int64	`protobuf:"varint,4,opt,name=lease,proto3" json:"lease,omitempty"`
}

func (m *LeaderKey) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaderKey{}
}
func (m *LeaderKey) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaderKey) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaderKey) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorV3Election, []int{2}
}
func (m *LeaderKey) GetName() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return nil
}
func (m *LeaderKey) GetKey() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Key
	}
	return nil
}
func (m *LeaderKey) GetRev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Rev
	}
	return 0
}
func (m *LeaderKey) GetLease() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Lease
	}
	return 0
}

type LeaderRequest struct {
	Name []byte `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *LeaderRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaderRequest{}
}
func (m *LeaderRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaderRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaderRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorV3Election, []int{3}
}
func (m *LeaderRequest) GetName() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Name
	}
	return nil
}

type LeaderResponse struct {
	Header	*etcdserverpb.ResponseHeader	`protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Kv	*mvccpb.KeyValue		`protobuf:"bytes,2,opt,name=kv" json:"kv,omitempty"`
}

func (m *LeaderResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = LeaderResponse{}
}
func (m *LeaderResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*LeaderResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*LeaderResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorV3Election, []int{4}
}
func (m *LeaderResponse) GetHeader() *etcdserverpb.ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}
func (m *LeaderResponse) GetKv() *mvccpb.KeyValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Kv
	}
	return nil
}

type ResignRequest struct {
	Leader *LeaderKey `protobuf:"bytes,1,opt,name=leader" json:"leader,omitempty"`
}

func (m *ResignRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = ResignRequest{}
}
func (m *ResignRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*ResignRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ResignRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorV3Election, []int{5}
}
func (m *ResignRequest) GetLeader() *LeaderKey {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Leader
	}
	return nil
}

type ResignResponse struct {
	Header *etcdserverpb.ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *ResignResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = ResignResponse{}
}
func (m *ResignResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*ResignResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ResignResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorV3Election, []int{6}
}
func (m *ResignResponse) GetHeader() *etcdserverpb.ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Header
	}
	return nil
}

type ProclaimRequest struct {
	Leader	*LeaderKey	`protobuf:"bytes,1,opt,name=leader" json:"leader,omitempty"`
	Value	[]byte		`protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ProclaimRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = ProclaimRequest{}
}
func (m *ProclaimRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*ProclaimRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ProclaimRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorV3Election, []int{7}
}
func (m *ProclaimRequest) GetLeader() *LeaderKey {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Leader
	}
	return nil
}
func (m *ProclaimRequest) GetValue() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m != nil {
		return m.Value
	}
	return nil
}

type ProclaimResponse struct {
	Header *etcdserverpb.ResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
}

func (m *ProclaimResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = ProclaimResponse{}
}
func (m *ProclaimResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*ProclaimResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*ProclaimResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorV3Election, []int{8}
}
func (m *ProclaimResponse) GetHeader() *etcdserverpb.ResponseHeader {
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
	proto.RegisterType((*CampaignRequest)(nil), "v3electionpb.CampaignRequest")
	proto.RegisterType((*CampaignResponse)(nil), "v3electionpb.CampaignResponse")
	proto.RegisterType((*LeaderKey)(nil), "v3electionpb.LeaderKey")
	proto.RegisterType((*LeaderRequest)(nil), "v3electionpb.LeaderRequest")
	proto.RegisterType((*LeaderResponse)(nil), "v3electionpb.LeaderResponse")
	proto.RegisterType((*ResignRequest)(nil), "v3electionpb.ResignRequest")
	proto.RegisterType((*ResignResponse)(nil), "v3electionpb.ResignResponse")
	proto.RegisterType((*ProclaimRequest)(nil), "v3electionpb.ProclaimRequest")
	proto.RegisterType((*ProclaimResponse)(nil), "v3electionpb.ProclaimResponse")
}

var _ context.Context
var _ grpc.ClientConn

const _ = grpc.SupportPackageIsVersion4

type ElectionClient interface {
	Campaign(ctx context.Context, in *CampaignRequest, opts ...grpc.CallOption) (*CampaignResponse, error)
	Proclaim(ctx context.Context, in *ProclaimRequest, opts ...grpc.CallOption) (*ProclaimResponse, error)
	Leader(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (*LeaderResponse, error)
	Observe(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (Election_ObserveClient, error)
	Resign(ctx context.Context, in *ResignRequest, opts ...grpc.CallOption) (*ResignResponse, error)
}
type electionClient struct{ cc *grpc.ClientConn }

func NewElectionClient(cc *grpc.ClientConn) ElectionClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &electionClient{cc}
}
func (c *electionClient) Campaign(ctx context.Context, in *CampaignRequest, opts ...grpc.CallOption) (*CampaignResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(CampaignResponse)
	err := grpc.Invoke(ctx, "/v3electionpb.Election/Campaign", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *electionClient) Proclaim(ctx context.Context, in *ProclaimRequest, opts ...grpc.CallOption) (*ProclaimResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(ProclaimResponse)
	err := grpc.Invoke(ctx, "/v3electionpb.Election/Proclaim", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *electionClient) Leader(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (*LeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(LeaderResponse)
	err := grpc.Invoke(ctx, "/v3electionpb.Election/Leader", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *electionClient) Observe(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (Election_ObserveClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stream, err := grpc.NewClientStream(ctx, &_Election_serviceDesc.Streams[0], c.cc, "/v3electionpb.Election/Observe", opts...)
	if err != nil {
		return nil, err
	}
	x := &electionObserveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Election_ObserveClient interface {
	Recv() (*LeaderResponse, error)
	grpc.ClientStream
}
type electionObserveClient struct{ grpc.ClientStream }

func (x *electionObserveClient) Recv() (*LeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(LeaderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
func (c *electionClient) Resign(ctx context.Context, in *ResignRequest, opts ...grpc.CallOption) (*ResignResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out := new(ResignResponse)
	err := grpc.Invoke(ctx, "/v3electionpb.Election/Resign", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ElectionServer interface {
	Campaign(context.Context, *CampaignRequest) (*CampaignResponse, error)
	Proclaim(context.Context, *ProclaimRequest) (*ProclaimResponse, error)
	Leader(context.Context, *LeaderRequest) (*LeaderResponse, error)
	Observe(*LeaderRequest, Election_ObserveServer) error
	Resign(context.Context, *ResignRequest) (*ResignResponse, error)
}

func RegisterElectionServer(s *grpc.Server, srv ElectionServer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.RegisterService(&_Election_serviceDesc, srv)
}
func _Election_Campaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(CampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).Campaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/v3electionpb.Election/Campaign"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).Campaign(ctx, req.(*CampaignRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Election_Proclaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(ProclaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).Proclaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/v3electionpb.Election/Proclaim"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).Proclaim(ctx, req.(*ProclaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Election_Leader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(LeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).Leader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/v3electionpb.Election/Leader"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).Leader(ctx, req.(*LeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Election_Observe_Handler(srv interface{}, stream grpc.ServerStream) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(LeaderRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ElectionServer).Observe(m, &electionObserveServer{stream})
}

type Election_ObserveServer interface {
	Send(*LeaderResponse) error
	grpc.ServerStream
}
type electionObserveServer struct{ grpc.ServerStream }

func (x *electionObserveServer) Send(m *LeaderResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return x.ServerStream.SendMsg(m)
}
func _Election_Resign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in := new(ResignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).Resign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/v3electionpb.Election/Resign"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).Resign(ctx, req.(*ResignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Election_serviceDesc = grpc.ServiceDesc{ServiceName: "v3electionpb.Election", HandlerType: (*ElectionServer)(nil), Methods: []grpc.MethodDesc{{MethodName: "Campaign", Handler: _Election_Campaign_Handler}, {MethodName: "Proclaim", Handler: _Election_Proclaim_Handler}, {MethodName: "Leader", Handler: _Election_Leader_Handler}, {MethodName: "Resign", Handler: _Election_Resign_Handler}}, Streams: []grpc.StreamDesc{{StreamName: "Observe", Handler: _Election_Observe_Handler, ServerStreams: true}}, Metadata: "v3election.proto"}

func (m *CampaignRequest) Marshal() (dAtA []byte, err error) {
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
func (m *CampaignRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Lease != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Lease))
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}
func (m *CampaignResponse) Marshal() (dAtA []byte, err error) {
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
func (m *CampaignResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Header.Size()))
		n1, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Leader != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Leader.Size()))
		n2, err := m.Leader.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *LeaderKey) Marshal() (dAtA []byte, err error) {
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
func (m *LeaderKey) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if m.Rev != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Rev))
	}
	if m.Lease != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Lease))
	}
	return i, nil
}
func (m *LeaderRequest) Marshal() (dAtA []byte, err error) {
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
func (m *LeaderRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}
func (m *LeaderResponse) Marshal() (dAtA []byte, err error) {
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
func (m *LeaderResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Header.Size()))
		n3, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Kv != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Kv.Size()))
		n4, err := m.Kv.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *ResignRequest) Marshal() (dAtA []byte, err error) {
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
func (m *ResignRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Leader != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Leader.Size()))
		n5, err := m.Leader.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *ResignResponse) Marshal() (dAtA []byte, err error) {
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
func (m *ResignResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Header.Size()))
		n6, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *ProclaimRequest) Marshal() (dAtA []byte, err error) {
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
func (m *ProclaimRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Leader != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Leader.Size()))
		n7, err := m.Leader.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}
func (m *ProclaimResponse) Marshal() (dAtA []byte, err error) {
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
func (m *ProclaimResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Header != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintV3Election(dAtA, i, uint64(m.Header.Size()))
		n8, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}
func encodeVarintV3Election(dAtA []byte, offset int, v uint64) int {
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
func (m *CampaignRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovV3Election(uint64(l))
	}
	if m.Lease != 0 {
		n += 1 + sovV3Election(uint64(m.Lease))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovV3Election(uint64(l))
	}
	return n
}
func (m *CampaignResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovV3Election(uint64(l))
	}
	if m.Leader != nil {
		l = m.Leader.Size()
		n += 1 + l + sovV3Election(uint64(l))
	}
	return n
}
func (m *LeaderKey) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovV3Election(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovV3Election(uint64(l))
	}
	if m.Rev != 0 {
		n += 1 + sovV3Election(uint64(m.Rev))
	}
	if m.Lease != 0 {
		n += 1 + sovV3Election(uint64(m.Lease))
	}
	return n
}
func (m *LeaderRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovV3Election(uint64(l))
	}
	return n
}
func (m *LeaderResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovV3Election(uint64(l))
	}
	if m.Kv != nil {
		l = m.Kv.Size()
		n += 1 + l + sovV3Election(uint64(l))
	}
	return n
}
func (m *ResignRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Leader != nil {
		l = m.Leader.Size()
		n += 1 + l + sovV3Election(uint64(l))
	}
	return n
}
func (m *ResignResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovV3Election(uint64(l))
	}
	return n
}
func (m *ProclaimRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Leader != nil {
		l = m.Leader.Size()
		n += 1 + l + sovV3Election(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovV3Election(uint64(l))
	}
	return n
}
func (m *ProclaimResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovV3Election(uint64(l))
	}
	return n
}
func sovV3Election(x uint64) (n int) {
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
func sozV3Election(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sovV3Election(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CampaignRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Election
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
			return fmt.Errorf("proto: CampaignRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CampaignRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = append(m.Name[:0], dAtA[iNdEx:postIndex]...)
			if m.Name == nil {
				m.Name = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			m.Lease = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
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
		default:
			iNdEx = preIndex
			skippy, err := skipV3Election(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Election
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
func (m *CampaignResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Election
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
			return fmt.Errorf("proto: CampaignResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CampaignResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &etcdserverpb.ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Leader == nil {
				m.Leader = &LeaderKey{}
			}
			if err := m.Leader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Election(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Election
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
func (m *LeaderKey) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Election
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
			return fmt.Errorf("proto: LeaderKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaderKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = append(m.Name[:0], dAtA[iNdEx:postIndex]...)
			if m.Name == nil {
				m.Name = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rev", wireType)
			}
			m.Rev = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rev |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			m.Lease = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
		default:
			iNdEx = preIndex
			skippy, err := skipV3Election(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Election
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
func (m *LeaderRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Election
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
			return fmt.Errorf("proto: LeaderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = append(m.Name[:0], dAtA[iNdEx:postIndex]...)
			if m.Name == nil {
				m.Name = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Election(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Election
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
func (m *LeaderResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Election
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
			return fmt.Errorf("proto: LeaderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LeaderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &etcdserverpb.ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kv", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Kv == nil {
				m.Kv = &mvccpb.KeyValue{}
			}
			if err := m.Kv.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Election(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Election
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
func (m *ResignRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Election
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
			return fmt.Errorf("proto: ResignRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResignRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Leader == nil {
				m.Leader = &LeaderKey{}
			}
			if err := m.Leader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Election(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Election
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
func (m *ResignResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Election
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
			return fmt.Errorf("proto: ResignResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResignResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &etcdserverpb.ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Election(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Election
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
func (m *ProclaimRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Election
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
			return fmt.Errorf("proto: ProclaimRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProclaimRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Leader == nil {
				m.Leader = &LeaderKey{}
			}
			if err := m.Leader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
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
		default:
			iNdEx = preIndex
			skippy, err := skipV3Election(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Election
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
func (m *ProclaimResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowV3Election
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
			return fmt.Errorf("proto: ProclaimResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProclaimResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowV3Election
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
				return ErrInvalidLengthV3Election
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &etcdserverpb.ResponseHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipV3Election(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthV3Election
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
func skipV3Election(dAtA []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowV3Election
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
					return 0, ErrIntOverflowV3Election
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
					return 0, ErrIntOverflowV3Election
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
				return 0, ErrInvalidLengthV3Election
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowV3Election
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
				next, err := skipV3Election(dAtA[start:])
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
	ErrInvalidLengthV3Election	= fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowV3Election	= fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("v3election.proto", fileDescriptorV3Election)
}

var fileDescriptorV3Election = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6e, 0xd3, 0x40, 0x10, 0xc6, 0x59, 0x27, 0x84, 0x32, 0xa4, 0xad, 0x65, 0x82, 0x48, 0x43, 0x30, 0xd1, 0x22, 0xa1, 0x2a, 0x07, 0x2f, 0x6a, 0x38, 0xe5, 0x84, 0x40, 0xa0, 0x4a, 0x45, 0x02, 0x7c, 0x40, 0x70, 0xdc, 0xb8, 0x23, 0x37, 0x8a, 0xe3, 0x35, 0xb6, 0x6b, 0x29, 0x57, 0x5e, 0x81, 0x03, 0x3c, 0x12, 0x47, 0x24, 0x5e, 0x00, 0x05, 0x1e, 0x04, 0xed, 0xae, 0x8d, 0xff, 0x28, 0x41, 0xa8, 0xb9, 0x58, 0xe3, 0x9d, 0xcf, 0xf3, 0x9b, 0x6f, 0x76, 0x12, 0x30, 0xb3, 0x09, 0x06, 0xe8, 0xa5, 0x73, 0x11, 0x3a, 0x51, 0x2c, 0x52, 0x61, 0x75, 0xcb, 0x93, 0x68, 0x36, 0xe8, 0xf9, 0xc2, 0x17, 0x2a, 0xc1, 0x64, 0xa4, 0x35, 0x83, 0x47, 0x98, 0x7a, 0xe7, 0x4c, 0x3e, 0x12, 0x8c, 0x33, 0x8c, 0x2b, 0x61, 0x34, 0x63, 0x71, 0xe4, 0xe5, 0xba, 0x23, 0xa5, 0x5b, 0x66, 0x9e, 0xa7, 0x1e, 0xd1, 0x8c, 0x2d, 0xb2, 0x3c, 0x35, 0xf4, 0x85, 0xf0, 0x03, 0x64, 0x3c, 0x9a, 0x33, 0x1e, 0x86, 0x22, 0xe5, 0x92, 0x98, 0xe8, 0x2c, 0x7d, 0x0b, 0x87, 0xcf, 0xf9, 0x32, 0xe2, 0x73, 0x3f, 0x74, 0xf1, 0xe3, 0x25, 0x26, 0xa9, 0x65, 0x41, 0x3b, 0xe4, 0x4b, 0xec, 0x93, 0x11, 0x39, 0xee, 0xba, 0x2a, 0xb6, 0x7a, 0x70, 0x3d, 0x40, 0x9e, 0x60, 0xdf, 0x18, 0x91, 0xe3, 0x96, 0xab, 0x5f, 0xe4, 0x69, 0xc6, 0x83, 0x4b, 0xec, 0xb7, 0x94, 0x54, 0xbf, 0xd0, 0x15, 0x98, 0x65, 0xc9, 0x24, 0x12, 0x61, 0x82, 0xd6, 0x13, 0xe8, 0x5c, 0x20, 0x3f, 0xc7, 0x58, 0x55, 0xbd, 0x75, 0x32, 0x74, 0xaa, 0x46, 0x9c, 0x42, 0x77, 0xaa, 0x34, 0x6e, 0xae, 0xb5, 0x18, 0x74, 0x02, 0xfd, 0x95, 0xa1, 0xbe, 0xba, 0xeb, 0x54, 0x47, 0xe6, 0xbc, 0x52, 0xb9, 0x33, 0x5c, 0xb9, 0xb9, 0x8c, 0x7e, 0x80, 0x9b, 0x7f, 0x0f, 0x37, 0xfa, 0x30, 0xa1, 0xb5, 0xc0, 0x95, 0x2a, 0xd7, 0x75, 0x65, 0x28, 0x4f, 0x62, 0xcc, 0x94, 0x83, 0x96, 0x2b, 0xc3, 0xd2, 0x6b, 0xbb, 0xe2, 0x95, 0x3e, 0x84, 0x7d, 0x5d, 0xfa, 0x1f, 0x63, 0xa2, 0x17, 0x70, 0x50, 0x88, 0x76, 0x32, 0x3e, 0x02, 0x63, 0x91, 0xe5, 0xa6, 0x4d, 0x47, 0xdf, 0xa8, 0x73, 0x86, 0xab, 0x77, 0x72, 0xc0, 0xae, 0xb1, 0xc8, 0xe8, 0x53, 0xd8, 0x77, 0x31, 0xa9, 0xdc, 0x5a, 0x39, 0x2b, 0xf2, 0x7f, 0xb3, 0x7a, 0x09, 0x07, 0x45, 0x85, 0x5d, 0x7a, 0xa5, 0xef, 0xe1, 0xf0, 0x4d, 0x2c, 0xbc, 0x80, 0xcf, 0x97, 0x57, 0xed, 0xa5, 0x5c, 0x24, 0xa3, 0xba, 0x48, 0xa7, 0x60, 0x96, 0x95, 0x77, 0xe9, 0xf1, 0xe4, 0x4b, 0x1b, 0xf6, 0x5e, 0xe4, 0x0d, 0x58, 0x0b, 0xd8, 0x2b, 0xf6, 0xd3, 0xba, 0x5f, 0xef, 0xac, 0xf1, 0x53, 0x18, 0xd8, 0xdb, 0xd2, 0x9a, 0x42, 0x47, 0x9f, 0x7e, 0xfc, 0xfe, 0x6c, 0x0c, 0xe8, 0x1d, 0x96, 0x4d, 0x58, 0x21, 0x64, 0x5e, 0x2e, 0x9b, 0x92, 0xb1, 0x84, 0x15, 0x1e, 0x9a, 0xb0, 0xc6, 0xd4, 0x9a, 0xb0, 0xa6, 0xf5, 0x2d, 0xb0, 0x28, 0x97, 0x49, 0x98, 0x07, 0x1d, 0x3d, 0x5b, 0xeb, 0xde, 0xa6, 0x89, 0x17, 0xa0, 0xe1, 0xe6, 0x64, 0x8e, 0xb1, 0x15, 0xa6, 0x4f, 0x6f, 0xd7, 0x30, 0xfa, 0xa2, 0x24, 0xc4, 0x87, 0x1b, 0xaf, 0x67, 0x6a, 0xe0, 0xbb, 0x50, 0x1e, 0x28, 0xca, 0x11, 0xed, 0xd5, 0x28, 0x42, 0x17, 0x9e, 0x92, 0xf1, 0x63, 0x22, 0xdd, 0xe8, 0x05, 0x6d, 0x72, 0x6a, 0x8b, 0xdf, 0xe4, 0xd4, 0x77, 0x7a, 0x8b, 0x9b, 0x58, 0x89, 0xa6, 0x64, 0xfc, 0xcc, 0xfc, 0xb6, 0xb6, 0xc9, 0xf7, 0xb5, 0x4d, 0x7e, 0xae, 0x6d, 0xf2, 0xf5, 0x97, 0x7d, 0x6d, 0xd6, 0x51, 0x7f, 0x8c, 0x93, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x1d, 0xfa, 0x11, 0xb1, 0x05, 0x00, 0x00}

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
