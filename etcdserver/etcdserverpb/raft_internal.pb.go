package etcdserverpb

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "github.com/gogo/protobuf/gogoproto"
	io "io"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RequestHeader struct {
	ID		uint64	`protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username	string	`protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	AuthRevision	uint64	`protobuf:"varint,3,opt,name=auth_revision,json=authRevision,proto3" json:"auth_revision,omitempty"`
}

func (m *RequestHeader) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = RequestHeader{}
}
func (m *RequestHeader) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*RequestHeader) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*RequestHeader) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaftInternal, []int{0}
}

type InternalRaftRequest struct {
	Header				*RequestHeader				`protobuf:"bytes,100,opt,name=header" json:"header,omitempty"`
	ID				uint64					`protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	V2				*Request				`protobuf:"bytes,2,opt,name=v2" json:"v2,omitempty"`
	Range				*RangeRequest				`protobuf:"bytes,3,opt,name=range" json:"range,omitempty"`
	Put				*PutRequest				`protobuf:"bytes,4,opt,name=put" json:"put,omitempty"`
	DeleteRange			*DeleteRangeRequest			`protobuf:"bytes,5,opt,name=delete_range,json=deleteRange" json:"delete_range,omitempty"`
	Txn				*TxnRequest				`protobuf:"bytes,6,opt,name=txn" json:"txn,omitempty"`
	Compaction			*CompactionRequest			`protobuf:"bytes,7,opt,name=compaction" json:"compaction,omitempty"`
	LeaseGrant			*LeaseGrantRequest			`protobuf:"bytes,8,opt,name=lease_grant,json=leaseGrant" json:"lease_grant,omitempty"`
	LeaseRevoke			*LeaseRevokeRequest			`protobuf:"bytes,9,opt,name=lease_revoke,json=leaseRevoke" json:"lease_revoke,omitempty"`
	Alarm				*AlarmRequest				`protobuf:"bytes,10,opt,name=alarm" json:"alarm,omitempty"`
	LeaseCheckpoint			*LeaseCheckpointRequest			`protobuf:"bytes,11,opt,name=lease_checkpoint,json=leaseCheckpoint" json:"lease_checkpoint,omitempty"`
	AuthEnable			*AuthEnableRequest			`protobuf:"bytes,1000,opt,name=auth_enable,json=authEnable" json:"auth_enable,omitempty"`
	AuthDisable			*AuthDisableRequest			`protobuf:"bytes,1011,opt,name=auth_disable,json=authDisable" json:"auth_disable,omitempty"`
	Authenticate			*InternalAuthenticateRequest		`protobuf:"bytes,1012,opt,name=authenticate" json:"authenticate,omitempty"`
	AuthUserAdd			*AuthUserAddRequest			`protobuf:"bytes,1100,opt,name=auth_user_add,json=authUserAdd" json:"auth_user_add,omitempty"`
	AuthUserDelete			*AuthUserDeleteRequest			`protobuf:"bytes,1101,opt,name=auth_user_delete,json=authUserDelete" json:"auth_user_delete,omitempty"`
	AuthUserGet			*AuthUserGetRequest			`protobuf:"bytes,1102,opt,name=auth_user_get,json=authUserGet" json:"auth_user_get,omitempty"`
	AuthUserChangePassword		*AuthUserChangePasswordRequest		`protobuf:"bytes,1103,opt,name=auth_user_change_password,json=authUserChangePassword" json:"auth_user_change_password,omitempty"`
	AuthUserGrantRole		*AuthUserGrantRoleRequest		`protobuf:"bytes,1104,opt,name=auth_user_grant_role,json=authUserGrantRole" json:"auth_user_grant_role,omitempty"`
	AuthUserRevokeRole		*AuthUserRevokeRoleRequest		`protobuf:"bytes,1105,opt,name=auth_user_revoke_role,json=authUserRevokeRole" json:"auth_user_revoke_role,omitempty"`
	AuthUserList			*AuthUserListRequest			`protobuf:"bytes,1106,opt,name=auth_user_list,json=authUserList" json:"auth_user_list,omitempty"`
	AuthRoleList			*AuthRoleListRequest			`protobuf:"bytes,1107,opt,name=auth_role_list,json=authRoleList" json:"auth_role_list,omitempty"`
	AuthRoleAdd			*AuthRoleAddRequest			`protobuf:"bytes,1200,opt,name=auth_role_add,json=authRoleAdd" json:"auth_role_add,omitempty"`
	AuthRoleDelete			*AuthRoleDeleteRequest			`protobuf:"bytes,1201,opt,name=auth_role_delete,json=authRoleDelete" json:"auth_role_delete,omitempty"`
	AuthRoleGet			*AuthRoleGetRequest			`protobuf:"bytes,1202,opt,name=auth_role_get,json=authRoleGet" json:"auth_role_get,omitempty"`
	AuthRoleGrantPermission		*AuthRoleGrantPermissionRequest		`protobuf:"bytes,1203,opt,name=auth_role_grant_permission,json=authRoleGrantPermission" json:"auth_role_grant_permission,omitempty"`
	AuthRoleRevokePermission	*AuthRoleRevokePermissionRequest	`protobuf:"bytes,1204,opt,name=auth_role_revoke_permission,json=authRoleRevokePermission" json:"auth_role_revoke_permission,omitempty"`
}

func (m *InternalRaftRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = InternalRaftRequest{}
}
func (m *InternalRaftRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*InternalRaftRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*InternalRaftRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaftInternal, []int{1}
}

type EmptyResponse struct{}

func (m *EmptyResponse) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = EmptyResponse{}
}
func (m *EmptyResponse) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*EmptyResponse) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaftInternal, []int{2}
}

type InternalAuthenticateRequest struct {
	Name		string	`protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password	string	`protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	SimpleToken	string	`protobuf:"bytes,3,opt,name=simple_token,json=simpleToken,proto3" json:"simple_token,omitempty"`
}

func (m *InternalAuthenticateRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = InternalAuthenticateRequest{}
}
func (m *InternalAuthenticateRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*InternalAuthenticateRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*InternalAuthenticateRequest) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRaftInternal, []int{3}
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterType((*RequestHeader)(nil), "etcdserverpb.RequestHeader")
	proto.RegisterType((*InternalRaftRequest)(nil), "etcdserverpb.InternalRaftRequest")
	proto.RegisterType((*EmptyResponse)(nil), "etcdserverpb.EmptyResponse")
	proto.RegisterType((*InternalAuthenticateRequest)(nil), "etcdserverpb.InternalAuthenticateRequest")
}
func (m *RequestHeader) Marshal() (dAtA []byte, err error) {
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
func (m *RequestHeader) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.ID))
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if m.AuthRevision != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRevision))
	}
	return i, nil
}
func (m *InternalRaftRequest) Marshal() (dAtA []byte, err error) {
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
func (m *InternalRaftRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.ID))
	}
	if m.V2 != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.V2.Size()))
		n1, err := m.V2.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Range != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.Range.Size()))
		n2, err := m.Range.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Put != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.Put.Size()))
		n3, err := m.Put.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.DeleteRange != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.DeleteRange.Size()))
		n4, err := m.DeleteRange.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Txn != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.Txn.Size()))
		n5, err := m.Txn.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Compaction != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.Compaction.Size()))
		n6, err := m.Compaction.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.LeaseGrant != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.LeaseGrant.Size()))
		n7, err := m.LeaseGrant.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.LeaseRevoke != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.LeaseRevoke.Size()))
		n8, err := m.LeaseRevoke.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.Alarm != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.Alarm.Size()))
		n9, err := m.Alarm.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	if m.LeaseCheckpoint != nil {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.LeaseCheckpoint.Size()))
		n10, err := m.LeaseCheckpoint.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.Header != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.Header.Size()))
		n11, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	if m.AuthEnable != nil {
		dAtA[i] = 0xc2
		i++
		dAtA[i] = 0x3e
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthEnable.Size()))
		n12, err := m.AuthEnable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	if m.AuthDisable != nil {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x3f
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthDisable.Size()))
		n13, err := m.AuthDisable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	if m.Authenticate != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x3f
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.Authenticate.Size()))
		n14, err := m.Authenticate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n14
	}
	if m.AuthUserAdd != nil {
		dAtA[i] = 0xe2
		i++
		dAtA[i] = 0x44
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserAdd.Size()))
		n15, err := m.AuthUserAdd.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n15
	}
	if m.AuthUserDelete != nil {
		dAtA[i] = 0xea
		i++
		dAtA[i] = 0x44
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserDelete.Size()))
		n16, err := m.AuthUserDelete.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n16
	}
	if m.AuthUserGet != nil {
		dAtA[i] = 0xf2
		i++
		dAtA[i] = 0x44
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserGet.Size()))
		n17, err := m.AuthUserGet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n17
	}
	if m.AuthUserChangePassword != nil {
		dAtA[i] = 0xfa
		i++
		dAtA[i] = 0x44
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserChangePassword.Size()))
		n18, err := m.AuthUserChangePassword.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n18
	}
	if m.AuthUserGrantRole != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x45
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserGrantRole.Size()))
		n19, err := m.AuthUserGrantRole.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n19
	}
	if m.AuthUserRevokeRole != nil {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x45
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserRevokeRole.Size()))
		n20, err := m.AuthUserRevokeRole.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n20
	}
	if m.AuthUserList != nil {
		dAtA[i] = 0x92
		i++
		dAtA[i] = 0x45
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserList.Size()))
		n21, err := m.AuthUserList.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n21
	}
	if m.AuthRoleList != nil {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x45
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleList.Size()))
		n22, err := m.AuthRoleList.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n22
	}
	if m.AuthRoleAdd != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleAdd.Size()))
		n23, err := m.AuthRoleAdd.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n23
	}
	if m.AuthRoleDelete != nil {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleDelete.Size()))
		n24, err := m.AuthRoleDelete.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n24
	}
	if m.AuthRoleGet != nil {
		dAtA[i] = 0x92
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleGet.Size()))
		n25, err := m.AuthRoleGet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n25
	}
	if m.AuthRoleGrantPermission != nil {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleGrantPermission.Size()))
		n26, err := m.AuthRoleGrantPermission.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n26
	}
	if m.AuthRoleRevokePermission != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleRevokePermission.Size()))
		n27, err := m.AuthRoleRevokePermission.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n27
	}
	return i, nil
}
func (m *EmptyResponse) Marshal() (dAtA []byte, err error) {
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
func (m *EmptyResponse) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}
func (m *InternalAuthenticateRequest) Marshal() (dAtA []byte, err error) {
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
func (m *InternalAuthenticateRequest) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if len(m.SimpleToken) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(len(m.SimpleToken)))
		i += copy(dAtA[i:], m.SimpleToken)
	}
	return i, nil
}
func encodeVarintRaftInternal(dAtA []byte, offset int, v uint64) int {
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
func (m *RequestHeader) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRaftInternal(uint64(m.ID))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRevision != 0 {
		n += 1 + sovRaftInternal(uint64(m.AuthRevision))
	}
	return n
}
func (m *InternalRaftRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRaftInternal(uint64(m.ID))
	}
	if m.V2 != nil {
		l = m.V2.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Put != nil {
		l = m.Put.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.DeleteRange != nil {
		l = m.DeleteRange.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Txn != nil {
		l = m.Txn.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Compaction != nil {
		l = m.Compaction.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseGrant != nil {
		l = m.LeaseGrant.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseRevoke != nil {
		l = m.LeaseRevoke.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Alarm != nil {
		l = m.Alarm.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.LeaseCheckpoint != nil {
		l = m.LeaseCheckpoint.Size()
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	if m.Header != nil {
		l = m.Header.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthEnable != nil {
		l = m.AuthEnable.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthDisable != nil {
		l = m.AuthDisable.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.Authenticate != nil {
		l = m.Authenticate.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserAdd != nil {
		l = m.AuthUserAdd.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserDelete != nil {
		l = m.AuthUserDelete.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserGet != nil {
		l = m.AuthUserGet.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserChangePassword != nil {
		l = m.AuthUserChangePassword.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserGrantRole != nil {
		l = m.AuthUserGrantRole.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserRevokeRole != nil {
		l = m.AuthUserRevokeRole.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthUserList != nil {
		l = m.AuthUserList.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleList != nil {
		l = m.AuthRoleList.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleAdd != nil {
		l = m.AuthRoleAdd.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleDelete != nil {
		l = m.AuthRoleDelete.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleGet != nil {
		l = m.AuthRoleGet.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleGrantPermission != nil {
		l = m.AuthRoleGrantPermission.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	if m.AuthRoleRevokePermission != nil {
		l = m.AuthRoleRevokePermission.Size()
		n += 2 + l + sovRaftInternal(uint64(l))
	}
	return n
}
func (m *EmptyResponse) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	return n
}
func (m *InternalAuthenticateRequest) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	l = len(m.SimpleToken)
	if l > 0 {
		n += 1 + l + sovRaftInternal(uint64(l))
	}
	return n
}
func sovRaftInternal(x uint64) (n int) {
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
func sozRaftInternal(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sovRaftInternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestHeader) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
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
			return fmt.Errorf("proto: RequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRevision", wireType)
			}
			m.AuthRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthRevision |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftInternal
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
func (m *InternalRaftRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
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
			return fmt.Errorf("proto: InternalRaftRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalRaftRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return fmt.Errorf("proto: wrong wireType = %d for field V2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.V2 == nil {
				m.V2 = &Request{}
			}
			if err := m.V2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Range == nil {
				m.Range = &RangeRequest{}
			}
			if err := m.Range.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Put", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Put == nil {
				m.Put = &PutRequest{}
			}
			if err := m.Put.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeleteRange == nil {
				m.DeleteRange = &DeleteRangeRequest{}
			}
			if err := m.DeleteRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Txn == nil {
				m.Txn = &TxnRequest{}
			}
			if err := m.Txn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Compaction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Compaction == nil {
				m.Compaction = &CompactionRequest{}
			}
			if err := m.Compaction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseGrant", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseGrant == nil {
				m.LeaseGrant = &LeaseGrantRequest{}
			}
			if err := m.LeaseGrant.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseRevoke", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseRevoke == nil {
				m.LeaseRevoke = &LeaseRevokeRequest{}
			}
			if err := m.LeaseRevoke.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alarm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Alarm == nil {
				m.Alarm = &AlarmRequest{}
			}
			if err := m.Alarm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseCheckpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LeaseCheckpoint == nil {
				m.LeaseCheckpoint = &LeaseCheckpointRequest{}
			}
			if err := m.LeaseCheckpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &RequestHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1000:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthEnable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthEnable == nil {
				m.AuthEnable = &AuthEnableRequest{}
			}
			if err := m.AuthEnable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1011:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthDisable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthDisable == nil {
				m.AuthDisable = &AuthDisableRequest{}
			}
			if err := m.AuthDisable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1012:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authenticate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authenticate == nil {
				m.Authenticate = &InternalAuthenticateRequest{}
			}
			if err := m.Authenticate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserAdd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserAdd == nil {
				m.AuthUserAdd = &AuthUserAddRequest{}
			}
			if err := m.AuthUserAdd.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1101:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserDelete", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserDelete == nil {
				m.AuthUserDelete = &AuthUserDeleteRequest{}
			}
			if err := m.AuthUserDelete.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1102:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserGet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserGet == nil {
				m.AuthUserGet = &AuthUserGetRequest{}
			}
			if err := m.AuthUserGet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1103:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserChangePassword", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserChangePassword == nil {
				m.AuthUserChangePassword = &AuthUserChangePasswordRequest{}
			}
			if err := m.AuthUserChangePassword.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1104:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserGrantRole", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserGrantRole == nil {
				m.AuthUserGrantRole = &AuthUserGrantRoleRequest{}
			}
			if err := m.AuthUserGrantRole.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1105:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserRevokeRole", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserRevokeRole == nil {
				m.AuthUserRevokeRole = &AuthUserRevokeRoleRequest{}
			}
			if err := m.AuthUserRevokeRole.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1106:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthUserList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthUserList == nil {
				m.AuthUserList = &AuthUserListRequest{}
			}
			if err := m.AuthUserList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1107:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthRoleList == nil {
				m.AuthRoleList = &AuthRoleListRequest{}
			}
			if err := m.AuthRoleList.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1200:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleAdd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthRoleAdd == nil {
				m.AuthRoleAdd = &AuthRoleAddRequest{}
			}
			if err := m.AuthRoleAdd.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1201:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleDelete", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthRoleDelete == nil {
				m.AuthRoleDelete = &AuthRoleDeleteRequest{}
			}
			if err := m.AuthRoleDelete.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1202:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleGet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthRoleGet == nil {
				m.AuthRoleGet = &AuthRoleGetRequest{}
			}
			if err := m.AuthRoleGet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1203:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleGrantPermission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthRoleGrantPermission == nil {
				m.AuthRoleGrantPermission = &AuthRoleGrantPermissionRequest{}
			}
			if err := m.AuthRoleGrantPermission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 1204:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthRoleRevokePermission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AuthRoleRevokePermission == nil {
				m.AuthRoleRevokePermission = &AuthRoleRevokePermissionRequest{}
			}
			if err := m.AuthRoleRevokePermission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftInternal
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
func (m *EmptyResponse) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
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
			return fmt.Errorf("proto: EmptyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EmptyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftInternal
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
func (m *InternalAuthenticateRequest) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaftInternal
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
			return fmt.Errorf("proto: InternalAuthenticateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalAuthenticateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
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
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SimpleToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaftInternal
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
				return ErrInvalidLengthRaftInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SimpleToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaftInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaftInternal
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
func skipRaftInternal(dAtA []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaftInternal
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
					return 0, ErrIntOverflowRaftInternal
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
					return 0, ErrIntOverflowRaftInternal
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
				return 0, ErrInvalidLengthRaftInternal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaftInternal
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
				next, err := skipRaftInternal(dAtA[start:])
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
	ErrInvalidLengthRaftInternal	= fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaftInternal	= fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("raft_internal.proto", fileDescriptorRaftInternal)
}

var fileDescriptorRaftInternal = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0xcb, 0x52, 0x33, 0x45, 0x14, 0xc7, 0x49, 0xb8, 0xa6, 0x13, 0x2e, 0x36, 0xa0, 0x6d, 0xa8, 0x8a, 0x10, 0xbc, 0xe0, 0x0d, 0xad, 0xf0, 0x00, 0x1a, 0x13, 0x0a, 0xa8, 0xa2, 0x90, 0x9a, 0xc2, 0x2a, 0xab, 0x5c, 0x8c, 0xcd, 0xcc, 0x21, 0x19, 0x99, 0xcc, 0x8c, 0x3d, 0x9d, 0x88, 0x6f, 0xe2, 0x63, 0x78, 0xdb, 0xbb, 0x65, 0xe1, 0x05, 0xf5, 0x05, 0x14, 0x37, 0xee, 0xbf, 0xef, 0x01, 0xbe, 0xea, 0xcb, 0xf4, 0x64, 0x92, 0x0e, 0xbb, 0xc9, 0x39, 0xff, 0xf3, 0xfb, 0x9f, 0x99, 0x3e, 0x07, 0x1a, 0x6d, 0x32, 0x7a, 0xc3, 0xdd, 0x20, 0xe2, 0xc0, 0x22, 0x1a, 0x1e, 0x26, 0x2c, 0xe6, 0x31, 0xae, 0x01, 0xf7, 0xfc, 0x14, 0xd8, 0x08, 0x58, 0x72, 0x5d, 0xdf, 0xea, 0xc5, 0xbd, 0x58, 0x26, 0x3e, 0x10, 0x4f, 0x4a, 0x53, 0xdf, 0xc8, 0x35, 0x3a, 0x52, 0x61, 0x89, 0xa7, 0x1e, 0x9b, 0x5f, 0xa2, 0x55, 0x07, 0xbe, 0x1e, 0x42, 0xca, 0x4f, 0x81, 0xfa, 0xc0, 0xf0, 0x1a, 0x2a, 0x9f, 0x75, 0x49, 0x69, 0xb7, 0x74, 0xb0, 0xe0, 0x94, 0xcf, 0xba, 0xb8, 0x8e, 0x56, 0x86, 0xa9, 0xb0, 0x1c, 0x00, 0x29, 0xef, 0x96, 0x0e, 0x2a, 0x8e, 0xf9, 0x8d, 0xf7, 0xd1, 0x2a, 0x1d, 0xf2, 0xbe, 0xcb, 0x60, 0x14, 0xa4, 0x41, 0x1c, 0x91, 0x79, 0x59, 0x56, 0x13, 0x41, 0x47, 0xc7, 0x9a, 0xbf, 0xac, 0xa3, 0xcd, 0x33, 0xdd, 0xb5, 0x43, 0x6f, 0xb8, 0xb6, 0x9b, 0x32, 0x7a, 0x03, 0x95, 0x47, 0x2d, 0x69, 0x51, 0x6d, 0x6d, 0x1f, 0x8e, 0xbf, 0xd7, 0xa1, 0x2e, 0x71, 0xca, 0xa3, 0x16, 0xfe, 0x10, 0x2d, 0x32, 0x1a, 0xf5, 0x40, 0x7a, 0x55, 0x5b, 0xf5, 0x09, 0xa5, 0x48, 0x65, 0x72, 0x25, 0xc4, 0xef, 0xa0, 0xf9, 0x64, 0xc8, 0xc9, 0x82, 0xd4, 0x93, 0xa2, 0xfe, 0x72, 0x98, 0xf5, 0xe3, 0x08, 0x11, 0xee, 0xa0, 0x9a, 0x0f, 0x21, 0x70, 0x70, 0x95, 0xc9, 0xa2, 0x2c, 0xda, 0x2d, 0x16, 0x75, 0xa5, 0xa2, 0x60, 0x55, 0xf5, 0xf3, 0x98, 0x30, 0xe4, 0x77, 0x11, 0x59, 0xb2, 0x19, 0x5e, 0xdd, 0x45, 0xc6, 0x90, 0xdf, 0x45, 0xf8, 0x23, 0x84, 0xbc, 0x78, 0x90, 0x50, 0x8f, 0x8b, 0xef, 0xb7, 0x2c, 0x4b, 0x5e, 0x2b, 0x96, 0x74, 0x4c, 0x3e, 0xab, 0x1c, 0x2b, 0xc1, 0x1f, 0xa3, 0x6a, 0x08, 0x34, 0x05, 0xb7, 0xc7, 0x68, 0xc4, 0xc9, 0x8a, 0x8d, 0x70, 0x2e, 0x04, 0x27, 0x22, 0x6f, 0x08, 0xa1, 0x09, 0x89, 0x77, 0x56, 0x04, 0x06, 0xa3, 0xf8, 0x16, 0x48, 0xc5, 0xf6, 0xce, 0x12, 0xe1, 0x48, 0x81, 0x79, 0xe7, 0x30, 0x8f, 0x89, 0x63, 0xa1, 0x21, 0x65, 0x03, 0x82, 0x6c, 0xc7, 0xd2, 0x16, 0x29, 0x73, 0x2c, 0x52, 0x88, 0x3f, 0x45, 0x1b, 0xca, 0xd6, 0xeb, 0x83, 0x77, 0x9b, 0xc4, 0x41, 0xc4, 0x49, 0x55, 0x16, 0xbf, 0x6e, 0xb1, 0xee, 0x18, 0x51, 0x86, 0x59, 0x0f, 0x8b, 0x71, 0x7c, 0x84, 0x96, 0xfa, 0x72, 0x86, 0x89, 0x2f, 0x31, 0x3b, 0xd6, 0x21, 0x52, 0x63, 0xee, 0x68, 0x29, 0x6e, 0xa3, 0xaa, 0x1c, 0x61, 0x88, 0xe8, 0x75, 0x08, 0xe4, 0x7f, 0xeb, 0x09, 0xb4, 0x87, 0xbc, 0x7f, 0x2c, 0x05, 0xe6, 0xfb, 0x51, 0x13, 0xc2, 0x5d, 0x24, 0x07, 0xde, 0xf5, 0x83, 0x54, 0x32, 0x9e, 0x2d, 0xdb, 0x3e, 0xa0, 0x60, 0x74, 0x95, 0xc2, 0x7c, 0x40, 0x9a, 0xc7, 0xf0, 0x85, 0xa2, 0x40, 0xc4, 0x03, 0x8f, 0x72, 0x20, 0xcf, 0x15, 0xe5, 0xed, 0x22, 0x25, 0x5b, 0xa4, 0xf6, 0x98, 0x34, 0xc3, 0x15, 0xea, 0xf1, 0xb1, 0xde, 0x4d, 0xb1, 0xac, 0x2e, 0xf5, 0x7d, 0xf2, 0xeb, 0xca, 0xac, 0xb6, 0x3e, 0x4b, 0x81, 0xb5, 0x7d, 0xbf, 0xd0, 0x96, 0x8e, 0xe1, 0x0b, 0xb4, 0x91, 0x63, 0xd4, 0x90, 0x93, 0xdf, 0x14, 0x69, 0xdf, 0x4e, 0xd2, 0xdb, 0xa1, 0x61, 0x6b, 0xb4, 0x10, 0x2e, 0xb6, 0xd5, 0x03, 0x4e, 0x7e, 0x7f, 0xb2, 0xad, 0x13, 0xe0, 0x53, 0x6d, 0x9d, 0x00, 0xc7, 0x3d, 0xf4, 0x6a, 0x8e, 0xf1, 0xfa, 0x62, 0xed, 0xdc, 0x84, 0xa6, 0xe9, 0x37, 0x31, 0xf3, 0xc9, 0x1f, 0x0a, 0xf9, 0xae, 0x1d, 0xd9, 0x91, 0xea, 0x4b, 0x2d, 0xce, 0xe8, 0x2f, 0x53, 0x6b, 0x1a, 0x7f, 0x8e, 0xb6, 0xc6, 0xfa, 0x15, 0xfb, 0xe2, 0xb2, 0x38, 0x04, 0xf2, 0xa0, 0x3c, 0xde, 0x9c, 0xd1, 0xb6, 0xdc, 0xb5, 0x38, 0x3f, 0xea, 0x97, 0xe8, 0x64, 0x06, 0x7f, 0x81, 0xb6, 0x73, 0xb2, 0x5a, 0x3d, 0x85, 0xfe, 0x53, 0xa1, 0xdf, 0xb2, 0xa3, 0xf5, 0x0e, 0x8e, 0xb1, 0x31, 0x9d, 0x4a, 0xe1, 0x53, 0xb4, 0x96, 0xc3, 0xc3, 0x20, 0xe5, 0xe4, 0x2f, 0x45, 0xdd, 0xb3, 0x53, 0xcf, 0x83, 0x94, 0x17, 0xe6, 0x28, 0x0b, 0x1a, 0x92, 0x68, 0x4d, 0x91, 0xfe, 0x9e, 0x49, 0x12, 0xd6, 0x53, 0xa4, 0x2c, 0x68, 0x8e, 0x5e, 0x92, 0xc4, 0x44, 0x7e, 0x5f, 0x99, 0x75, 0xf4, 0xa2, 0x66, 0x72, 0x22, 0x75, 0xcc, 0x4c, 0xa4, 0xc4, 0xe8, 0x89, 0xfc, 0xa1, 0x32, 0x6b, 0x22, 0x45, 0x95, 0x65, 0x22, 0xf3, 0x70, 0xb1, 0x2d, 0x31, 0x91, 0x3f, 0x3e, 0xd9, 0xd6, 0xe4, 0x44, 0xea, 0x18, 0xfe, 0x0a, 0xd5, 0xc7, 0x30, 0x72, 0x50, 0x12, 0x60, 0x83, 0x20, 0x95, 0xff, 0x18, 0x7f, 0x52, 0xcc, 0xf7, 0x66, 0x30, 0x85, 0xfc, 0xd2, 0xa8, 0x33, 0xfe, 0x2b, 0xd4, 0x9e, 0xc7, 0x03, 0xb4, 0x93, 0x7b, 0xe9, 0xd1, 0x19, 0x33, 0xfb, 0x59, 0x99, 0xbd, 0x6f, 0x37, 0x53, 0x53, 0x32, 0xed, 0x46, 0xe8, 0x0c, 0x41, 0x73, 0x1d, 0xad, 0x1e, 0x0f, 0x12, 0xfe, 0xad, 0x03, 0x69, 0x12, 0x47, 0x29, 0x34, 0x13, 0xb4, 0xf3, 0xc4, 0x1f, 0x22, 0x8c, 0xd1, 0x82, 0xbc, 0x2e, 0x94, 0xe4, 0x75, 0x41, 0x3e, 0x8b, 0x6b, 0x84, 0xd9, 0x4f, 0x7d, 0x8d, 0xc8, 0x7e, 0xe3, 0x3d, 0x54, 0x4b, 0x83, 0x41, 0x12, 0x82, 0xcb, 0xe3, 0x5b, 0x50, 0xb7, 0x88, 0x8a, 0x53, 0x55, 0xb1, 0x2b, 0x11, 0xfa, 0x64, 0xeb, 0xfe, 0xdf, 0xc6, 0xdc, 0xfd, 0x63, 0xa3, 0xf4, 0xf0, 0xd8, 0x28, 0xfd, 0xf3, 0xd8, 0x28, 0x7d, 0xf7, 0x5f, 0x63, 0xee, 0x7a, 0x49, 0xde, 0x61, 0x8e, 0x5e, 0x04, 0x00, 0x00, 0xff, 0xff, 0xed, 0x36, 0xf0, 0x6f, 0x1b, 0x09, 0x00, 0x00}
