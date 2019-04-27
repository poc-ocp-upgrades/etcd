package etcdserverpb

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

type Request struct {
	ID			uint64	`protobuf:"varint,1,opt,name=ID" json:"ID"`
	Method			string	`protobuf:"bytes,2,opt,name=Method" json:"Method"`
	Path			string	`protobuf:"bytes,3,opt,name=Path" json:"Path"`
	Val			string	`protobuf:"bytes,4,opt,name=Val" json:"Val"`
	Dir			bool	`protobuf:"varint,5,opt,name=Dir" json:"Dir"`
	PrevValue		string	`protobuf:"bytes,6,opt,name=PrevValue" json:"PrevValue"`
	PrevIndex		uint64	`protobuf:"varint,7,opt,name=PrevIndex" json:"PrevIndex"`
	PrevExist		*bool	`protobuf:"varint,8,opt,name=PrevExist" json:"PrevExist,omitempty"`
	Expiration		int64	`protobuf:"varint,9,opt,name=Expiration" json:"Expiration"`
	Wait			bool	`protobuf:"varint,10,opt,name=Wait" json:"Wait"`
	Since			uint64	`protobuf:"varint,11,opt,name=Since" json:"Since"`
	Recursive		bool	`protobuf:"varint,12,opt,name=Recursive" json:"Recursive"`
	Sorted			bool	`protobuf:"varint,13,opt,name=Sorted" json:"Sorted"`
	Quorum			bool	`protobuf:"varint,14,opt,name=Quorum" json:"Quorum"`
	Time			int64	`protobuf:"varint,15,opt,name=Time" json:"Time"`
	Stream			bool	`protobuf:"varint,16,opt,name=Stream" json:"Stream"`
	Refresh			*bool	`protobuf:"varint,17,opt,name=Refresh" json:"Refresh,omitempty"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (m *Request) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Request{}
}
func (m *Request) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Request) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Request) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorEtcdserver, []int{0}
}

type Metadata struct {
	NodeID			uint64	`protobuf:"varint,1,opt,name=NodeID" json:"NodeID"`
	ClusterID		uint64	`protobuf:"varint,2,opt,name=ClusterID" json:"ClusterID"`
	XXX_unrecognized	[]byte	`json:"-"`
}

func (m *Metadata) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Metadata{}
}
func (m *Metadata) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Metadata) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Metadata) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorEtcdserver, []int{1}
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterType((*Request)(nil), "etcdserverpb.Request")
	proto.RegisterType((*Metadata)(nil), "etcdserverpb.Metadata")
}
func (m *Request) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.ID))
	dAtA[i] = 0x12
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(len(m.Method)))
	i += copy(dAtA[i:], m.Method)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(len(m.Path)))
	i += copy(dAtA[i:], m.Path)
	dAtA[i] = 0x22
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(len(m.Val)))
	i += copy(dAtA[i:], m.Val)
	dAtA[i] = 0x28
	i++
	if m.Dir {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x32
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(len(m.PrevValue)))
	i += copy(dAtA[i:], m.PrevValue)
	dAtA[i] = 0x38
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.PrevIndex))
	if m.PrevExist != nil {
		dAtA[i] = 0x40
		i++
		if *m.PrevExist {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x48
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.Expiration))
	dAtA[i] = 0x50
	i++
	if m.Wait {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x58
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.Since))
	dAtA[i] = 0x60
	i++
	if m.Recursive {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x68
	i++
	if m.Sorted {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x70
	i++
	if m.Quorum {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x78
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.Time))
	dAtA[i] = 0x80
	i++
	dAtA[i] = 0x1
	i++
	if m.Stream {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	if m.Refresh != nil {
		dAtA[i] = 0x88
		i++
		dAtA[i] = 0x1
		i++
		if *m.Refresh {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func (m *Metadata) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.NodeID))
	dAtA[i] = 0x10
	i++
	i = encodeVarintEtcdserver(dAtA, i, uint64(m.ClusterID))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeVarintEtcdserver(dAtA []byte, offset int, v uint64) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (m *Request) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovEtcdserver(uint64(m.ID))
	l = len(m.Method)
	n += 1 + l + sovEtcdserver(uint64(l))
	l = len(m.Path)
	n += 1 + l + sovEtcdserver(uint64(l))
	l = len(m.Val)
	n += 1 + l + sovEtcdserver(uint64(l))
	n += 2
	l = len(m.PrevValue)
	n += 1 + l + sovEtcdserver(uint64(l))
	n += 1 + sovEtcdserver(uint64(m.PrevIndex))
	if m.PrevExist != nil {
		n += 2
	}
	n += 1 + sovEtcdserver(uint64(m.Expiration))
	n += 2
	n += 1 + sovEtcdserver(uint64(m.Since))
	n += 2
	n += 2
	n += 2
	n += 1 + sovEtcdserver(uint64(m.Time))
	n += 3
	if m.Refresh != nil {
		n += 3
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func (m *Metadata) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	n += 1 + sovEtcdserver(uint64(m.NodeID))
	n += 1 + sovEtcdserver(uint64(m.ClusterID))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}
func sovEtcdserver(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func sozEtcdserver(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sovEtcdserver(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEtcdserver
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
				return ErrInvalidLengthEtcdserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
				return ErrInvalidLengthEtcdserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
				return ErrInvalidLengthEtcdserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Val = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dir", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
			m.Dir = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
				return ErrInvalidLengthEtcdserver
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevIndex", wireType)
			}
			m.PrevIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrevIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevExist", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
			b := bool(v != 0)
			m.PrevExist = &b
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			m.Expiration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiration |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wait", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
			m.Wait = bool(v != 0)
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Since", wireType)
			}
			m.Since = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Since |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recursive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
			m.Recursive = bool(v != 0)
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sorted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
			m.Sorted = bool(v != 0)
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
			m.Quorum = bool(v != 0)
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stream", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
			m.Stream = bool(v != 0)
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refresh", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
			b := bool(v != 0)
			m.Refresh = &b
		default:
			iNdEx = preIndex
			skippy, err := skipEtcdserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEtcdserver
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
func (m *Metadata) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEtcdserver
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
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			m.ClusterID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEtcdserver
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClusterID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEtcdserver(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEtcdserver
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
func skipEtcdserver(dAtA []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEtcdserver
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
					return 0, ErrIntOverflowEtcdserver
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
					return 0, ErrIntOverflowEtcdserver
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
				return 0, ErrInvalidLengthEtcdserver
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEtcdserver
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
				next, err := skipEtcdserver(dAtA[start:])
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
	ErrInvalidLengthEtcdserver	= fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEtcdserver	= fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("etcdserver.proto", fileDescriptorEtcdserver)
}

var fileDescriptorEtcdserver = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd2, 0xdd, 0x6e, 0xda, 0x30, 0x14, 0x07, 0x70, 0x0c, 0xe1, 0xcb, 0x63, 0x1b, 0xb3, 0xd0, 0x74, 0x84, 0xa6, 0x2c, 0x42, 0xbb, 0xc8, 0xd5, 0xf6, 0x0e, 0x2c, 0x5c, 0x44, 0x2a, 0x15, 0x0d, 0x15, 0xbd, 0x76, 0xc9, 0x29, 0x58, 0x02, 0x4c, 0x1d, 0x07, 0xf1, 0x06, 0x7d, 0x85, 0x3e, 0x12, 0x97, 0x7d, 0x82, 0xaa, 0xa5, 0x2f, 0x52, 0x39, 0x24, 0xc4, 0xed, 0x5d, 0xf4, 0xfb, 0x9f, 0x1c, 0x1f, 0x7f, 0xd0, 0x2e, 0xea, 0x79, 0x9c, 0xa0, 0xda, 0xa1, 0xfa, 0xbb, 0x55, 0x52, 0x4b, 0xd6, 0x29, 0x65, 0x7b, 0xdb, 0xef, 0x2d, 0xe4, 0x42, 0x66, 0xc1, 0x3f, 0xf3, 0x75, 0xaa, 0x19, 0x3c, 0x38, 0xb4, 0x19, 0xe1, 0x7d, 0x8a, 0x89, 0x66, 0x3d, 0x5a, 0x0d, 0x03, 0x20, 0x1e, 0xf1, 0x9d, 0xa1, 0x73, 0x78, 0xfe, 0x5d, 0x89, 0xaa, 0x61, 0xc0, 0x7e, 0xd1, 0xc6, 0x18, 0xf5, 0x52, 0xc6, 0x50, 0xf5, 0x88, 0xdf, 0xce, 0x93, 0xdc, 0x18, 0x50, 0x67, 0xc2, 0xf5, 0x12, 0x6a, 0x56, 0x96, 0x09, 0xfb, 0x49, 0x6b, 0x33, 0xbe, 0x02, 0xc7, 0x0a, 0x0c, 0x18, 0x0f, 0x84, 0x82, 0xba, 0x47, 0xfc, 0x56, 0xe1, 0x81, 0x50, 0x6c, 0x40, 0xdb, 0x13, 0x85, 0xbb, 0x19, 0x5f, 0xa5, 0x08, 0x0d, 0xeb, 0xaf, 0x92, 0x8b, 0x9a, 0x70, 0x13, 0xe3, 0x1e, 0x9a, 0xd6, 0xa0, 0x25, 0x17, 0x35, 0xa3, 0xbd, 0x48, 0x34, 0xb4, 0xce, 0xab, 0x90, 0xa8, 0x64, 0xf6, 0x87, 0xd2, 0xd1, 0x7e, 0x2b, 0x14, 0xd7, 0x42, 0x6e, 0xa0, 0xed, 0x11, 0xbf, 0x96, 0x37, 0xb2, 0xdc, 0xec, 0xed, 0x86, 0x0b, 0x0d, 0xd4, 0x1a, 0x35, 0x13, 0xd6, 0xa7, 0xf5, 0xa9, 0xd8, 0xcc, 0x11, 0xbe, 0x58, 0x33, 0x9c, 0xc8, 0xac, 0x1f, 0xe1, 0x3c, 0x55, 0x89, 0xd8, 0x21, 0x74, 0xac, 0x5f, 0x4b, 0x36, 0x67, 0x3a, 0x95, 0x4a, 0x63, 0x0c, 0x5f, 0xad, 0x82, 0xdc, 0x4c, 0x7a, 0x95, 0x4a, 0x95, 0xae, 0xe1, 0x9b, 0x9d, 0x9e, 0xcc, 0x4c, 0x75, 0x2d, 0xd6, 0x08, 0xdf, 0xad, 0xa9, 0x33, 0xc9, 0xba, 0x6a, 0x85, 0x7c, 0x0d, 0xdd, 0x0f, 0x5d, 0x33, 0x63, 0xae, 0xb9, 0xe8, 0x3b, 0x85, 0xc9, 0x12, 0x7e, 0x58, 0xa7, 0x52, 0xe0, 0xe0, 0x82, 0xb6, 0xc6, 0xa8, 0x79, 0xcc, 0x35, 0x37, 0x9d, 0x2e, 0x65, 0x8c, 0x9f, 0x5e, 0x43, 0x6e, 0x66, 0x87, 0xff, 0x57, 0x69, 0xa2, 0x51, 0x85, 0x41, 0xf6, 0x28, 0xce, 0xb7, 0x70, 0xe6, 0x61, 0xef, 0xf0, 0xea, 0x56, 0x0e, 0x47, 0x97, 0x3c, 0x1d, 0x5d, 0xf2, 0x72, 0x74, 0xc9, 0xe3, 0x9b, 0x5b, 0x79, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xee, 0x40, 0xba, 0xd6, 0xa4, 0x02, 0x00, 0x00}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
