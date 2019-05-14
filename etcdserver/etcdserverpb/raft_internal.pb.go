package etcdserverpb

import (
	"fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
)

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RequestHeader struct {
	ID           uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	AuthRevision uint64 `protobuf:"varint,3,opt,name=auth_revision,json=authRevision,proto3" json:"auth_revision,omitempty"`
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
	Header                   *RequestHeader                   `protobuf:"bytes,100,opt,name=header" json:"header,omitempty"`
	ID                       uint64                           `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	V2                       *Request                         `protobuf:"bytes,2,opt,name=v2" json:"v2,omitempty"`
	Range                    *RangeRequest                    `protobuf:"bytes,3,opt,name=range" json:"range,omitempty"`
	Put                      *PutRequest                      `protobuf:"bytes,4,opt,name=put" json:"put,omitempty"`
	DeleteRange              *DeleteRangeRequest              `protobuf:"bytes,5,opt,name=delete_range,json=deleteRange" json:"delete_range,omitempty"`
	Txn                      *TxnRequest                      `protobuf:"bytes,6,opt,name=txn" json:"txn,omitempty"`
	Compaction               *CompactionRequest               `protobuf:"bytes,7,opt,name=compaction" json:"compaction,omitempty"`
	LeaseGrant               *LeaseGrantRequest               `protobuf:"bytes,8,opt,name=lease_grant,json=leaseGrant" json:"lease_grant,omitempty"`
	LeaseRevoke              *LeaseRevokeRequest              `protobuf:"bytes,9,opt,name=lease_revoke,json=leaseRevoke" json:"lease_revoke,omitempty"`
	Alarm                    *AlarmRequest                    `protobuf:"bytes,10,opt,name=alarm" json:"alarm,omitempty"`
	AuthEnable               *AuthEnableRequest               `protobuf:"bytes,1000,opt,name=auth_enable,json=authEnable" json:"auth_enable,omitempty"`
	AuthDisable              *AuthDisableRequest              `protobuf:"bytes,1011,opt,name=auth_disable,json=authDisable" json:"auth_disable,omitempty"`
	Authenticate             *InternalAuthenticateRequest     `protobuf:"bytes,1012,opt,name=authenticate" json:"authenticate,omitempty"`
	AuthUserAdd              *AuthUserAddRequest              `protobuf:"bytes,1100,opt,name=auth_user_add,json=authUserAdd" json:"auth_user_add,omitempty"`
	AuthUserDelete           *AuthUserDeleteRequest           `protobuf:"bytes,1101,opt,name=auth_user_delete,json=authUserDelete" json:"auth_user_delete,omitempty"`
	AuthUserGet              *AuthUserGetRequest              `protobuf:"bytes,1102,opt,name=auth_user_get,json=authUserGet" json:"auth_user_get,omitempty"`
	AuthUserChangePassword   *AuthUserChangePasswordRequest   `protobuf:"bytes,1103,opt,name=auth_user_change_password,json=authUserChangePassword" json:"auth_user_change_password,omitempty"`
	AuthUserGrantRole        *AuthUserGrantRoleRequest        `protobuf:"bytes,1104,opt,name=auth_user_grant_role,json=authUserGrantRole" json:"auth_user_grant_role,omitempty"`
	AuthUserRevokeRole       *AuthUserRevokeRoleRequest       `protobuf:"bytes,1105,opt,name=auth_user_revoke_role,json=authUserRevokeRole" json:"auth_user_revoke_role,omitempty"`
	AuthUserList             *AuthUserListRequest             `protobuf:"bytes,1106,opt,name=auth_user_list,json=authUserList" json:"auth_user_list,omitempty"`
	AuthRoleList             *AuthRoleListRequest             `protobuf:"bytes,1107,opt,name=auth_role_list,json=authRoleList" json:"auth_role_list,omitempty"`
	AuthRoleAdd              *AuthRoleAddRequest              `protobuf:"bytes,1200,opt,name=auth_role_add,json=authRoleAdd" json:"auth_role_add,omitempty"`
	AuthRoleDelete           *AuthRoleDeleteRequest           `protobuf:"bytes,1201,opt,name=auth_role_delete,json=authRoleDelete" json:"auth_role_delete,omitempty"`
	AuthRoleGet              *AuthRoleGetRequest              `protobuf:"bytes,1202,opt,name=auth_role_get,json=authRoleGet" json:"auth_role_get,omitempty"`
	AuthRoleGrantPermission  *AuthRoleGrantPermissionRequest  `protobuf:"bytes,1203,opt,name=auth_role_grant_permission,json=authRoleGrantPermission" json:"auth_role_grant_permission,omitempty"`
	AuthRoleRevokePermission *AuthRoleRevokePermissionRequest `protobuf:"bytes,1204,opt,name=auth_role_revoke_permission,json=authRoleRevokePermission" json:"auth_role_revoke_permission,omitempty"`
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
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	SimpleToken string `protobuf:"bytes,3,opt,name=simple_token,json=simpleToken,proto3" json:"simple_token,omitempty"`
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
	if m.Header != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.Header.Size()))
		n10, err := m.Header.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	if m.AuthEnable != nil {
		dAtA[i] = 0xc2
		i++
		dAtA[i] = 0x3e
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthEnable.Size()))
		n11, err := m.AuthEnable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	if m.AuthDisable != nil {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x3f
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthDisable.Size()))
		n12, err := m.AuthDisable.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	if m.Authenticate != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x3f
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.Authenticate.Size()))
		n13, err := m.Authenticate.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n13
	}
	if m.AuthUserAdd != nil {
		dAtA[i] = 0xe2
		i++
		dAtA[i] = 0x44
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserAdd.Size()))
		n14, err := m.AuthUserAdd.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n14
	}
	if m.AuthUserDelete != nil {
		dAtA[i] = 0xea
		i++
		dAtA[i] = 0x44
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserDelete.Size()))
		n15, err := m.AuthUserDelete.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n15
	}
	if m.AuthUserGet != nil {
		dAtA[i] = 0xf2
		i++
		dAtA[i] = 0x44
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserGet.Size()))
		n16, err := m.AuthUserGet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n16
	}
	if m.AuthUserChangePassword != nil {
		dAtA[i] = 0xfa
		i++
		dAtA[i] = 0x44
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserChangePassword.Size()))
		n17, err := m.AuthUserChangePassword.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n17
	}
	if m.AuthUserGrantRole != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x45
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserGrantRole.Size()))
		n18, err := m.AuthUserGrantRole.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n18
	}
	if m.AuthUserRevokeRole != nil {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x45
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserRevokeRole.Size()))
		n19, err := m.AuthUserRevokeRole.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n19
	}
	if m.AuthUserList != nil {
		dAtA[i] = 0x92
		i++
		dAtA[i] = 0x45
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthUserList.Size()))
		n20, err := m.AuthUserList.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n20
	}
	if m.AuthRoleList != nil {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x45
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleList.Size()))
		n21, err := m.AuthRoleList.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n21
	}
	if m.AuthRoleAdd != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleAdd.Size()))
		n22, err := m.AuthRoleAdd.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n22
	}
	if m.AuthRoleDelete != nil {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleDelete.Size()))
		n23, err := m.AuthRoleDelete.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n23
	}
	if m.AuthRoleGet != nil {
		dAtA[i] = 0x92
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleGet.Size()))
		n24, err := m.AuthRoleGet.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n24
	}
	if m.AuthRoleGrantPermission != nil {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleGrantPermission.Size()))
		n25, err := m.AuthRoleGrantPermission.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n25
	}
	if m.AuthRoleRevokePermission != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x4b
		i++
		i = encodeVarintRaftInternal(dAtA, i, uint64(m.AuthRoleRevokePermission.Size()))
		n26, err := m.AuthRoleRevokePermission.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n26
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
	ErrInvalidLengthRaftInternal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaftInternal   = fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("raft_internal.proto", fileDescriptorRaftInternal)
}

var fileDescriptorRaftInternal = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x96, 0xdd, 0x52, 0xd3, 0x40, 0x14, 0xc7, 0x69, 0xf9, 0xec, 0xb6, 0x54, 0x5c, 0x40, 0xd7, 0x32, 0x53, 0xa1, 0x8c, 0x8a, 0x5f, 0xe8, 0x94, 0x07, 0xd0, 0x4a, 0x19, 0x60, 0x86, 0x61, 0x98, 0x0c, 0xce, 0x38, 0xe3, 0x45, 0x5c, 0x9a, 0x43, 0x1b, 0x49, 0x93, 0xb8, 0xd9, 0x56, 0x7c, 0x13, 0x1f, 0xc3, 0xaf, 0x87, 0xe0, 0xc2, 0x0f, 0xd4, 0x17, 0x50, 0xbc, 0xf1, 0xca, 0x1b, 0x7d, 0x00, 0x67, 0x3f, 0x92, 0x34, 0x6d, 0xca, 0x5d, 0x72, 0xce, 0xff, 0xfc, 0xce, 0xd9, 0xec, 0x7f, 0xbb, 0x45, 0xb3, 0x8c, 0x1e, 0x72, 0xd3, 0x76, 0x39, 0x30, 0x97, 0x3a, 0xab, 0x3e, 0xf3, 0xb8, 0x87, 0x0b, 0xc0, 0x1b, 0x56, 0x00, 0xac, 0x0b, 0xcc, 0x3f, 0x28, 0xcd, 0x35, 0xbd, 0xa6, 0x27, 0x13, 0xf7, 0xc4, 0x93, 0xd2, 0x94, 0x66, 0x62, 0x8d, 0x8e, 0xe4, 0x98, 0xdf, 0x50, 0x8f, 0x95, 0x67, 0x68, 0xda, 0x80, 0x17, 0x1d, 0x08, 0xf8, 0x16, 0x50, 0x0b, 0x18, 0x2e, 0xa2, 0xec, 0x76, 0x9d, 0x64, 0x16, 0x33, 0x2b, 0x63, 0x46, 0x76, 0xbb, 0x8e, 0x4b, 0x68, 0xaa, 0x13, 0x88, 0x96, 0x6d, 0x20, 0xd9, 0xc5, 0xcc, 0x4a, 0xce, 0x88, 0xde, 0xf1, 0x32, 0x9a, 0xa6, 0x1d, 0xde, 0x32, 0x19, 0x74, 0xed, 0xc0, 0xf6, 0x5c, 0x32, 0x2a, 0xcb, 0x0a, 0x22, 0x68, 0xe8, 0x58, 0xe5, 0x4f, 0x11, 0xcd, 0x6e, 0xeb, 0xa9, 0x0d, 0x7a, 0xc8, 0x75, 0xbb, 0x81, 0x46, 0xd7, 0x50, 0xb6, 0x5b, 0x95, 0x2d, 0xf2, 0xd5, 0xf9, 0xd5, 0xde, 0x75, 0xad, 0xea, 0x12, 0x23, 0xdb, 0xad, 0xe2, 0xfb, 0x68, 0x9c, 0x51, 0xb7, 0x09, 0xb2, 0x57, 0xbe, 0x5a, 0xea, 0x53, 0x8a, 0x54, 0x28, 0x57, 0x42, 0x7c, 0x0b, 0x8d, 0xfa, 0x1d, 0x4e, 0xc6, 0xa4, 0x9e, 0x24, 0xf5, 0x7b, 0x9d, 0x70, 0x1e, 0x43, 0x88, 0xf0, 0x3a, 0x2a, 0x58, 0xe0, 0x00, 0x07, 0x53, 0x35, 0x19, 0x97, 0x45, 0x8b, 0xc9, 0xa2, 0xba, 0x54, 0x24, 0x5a, 0xe5, 0xad, 0x38, 0x26, 0x1a, 0xf2, 0x63, 0x97, 0x4c, 0xa4, 0x35, 0xdc, 0x3f, 0x76, 0xa3, 0x86, 0xfc, 0xd8, 0xc5, 0x0f, 0x10, 0x6a, 0x78, 0x6d, 0x9f, 0x36, 0xb8, 0xf8, 0x7e, 0x93, 0xb2, 0xe4, 0x6a, 0xb2, 0x64, 0x3d, 0xca, 0x87, 0x95, 0x3d, 0x25, 0xf8, 0x21, 0xca, 0x3b, 0x40, 0x03, 0x30, 0x9b, 0x8c, 0xba, 0x9c, 0x4c, 0xa5, 0x11, 0x76, 0x84, 0x60, 0x53, 0xe4, 0x23, 0x82, 0x13, 0x85, 0xc4, 0x9a, 0x15, 0x81, 0x41, 0xd7, 0x3b, 0x02, 0x92, 0x4b, 0x5b, 0xb3, 0x44, 0x18, 0x52, 0x10, 0xad, 0xd9, 0x89, 0x63, 0x62, 0x5b, 0xa8, 0x43, 0x59, 0x9b, 0xa0, 0xb4, 0x6d, 0xa9, 0x89, 0x54, 0xb4, 0x2d, 0x52, 0x88, 0xd7, 0xd0, 0x44, 0x4b, 0x5a, 0x8e, 0x58, 0xb2, 0x64, 0x21, 0x75, 0xcf, 0x95, 0x2b, 0x0d, 0x2d, 0xc5, 0x35, 0x94, 0x97, 0x8e, 0x03, 0x97, 0x1e, 0x38, 0x40, 0x7e, 0xa7, 0x7e, 0xb0, 0x5a, 0x87, 0xb7, 0x36, 0xa4, 0x20, 0x5a, 0x2e, 0x8d, 0x42, 0xb8, 0x8e, 0xa4, 0x3f, 0x4d, 0xcb, 0x0e, 0x24, 0xe3, 0xef, 0x64, 0xda, 0x7a, 0x05, 0xa3, 0xae, 0x14, 0xd1, 0x7a, 0x69, 0x1c, 0xc3, 0xbb, 0x8a, 0x02, 0x2e, 0xb7, 0x1b, 0x94, 0x03, 0xf9, 0xa7, 0x28, 0x37, 0x93, 0x94, 0xd0, 0xf7, 0xb5, 0x1e, 0x69, 0x88, 0x4b, 0xd4, 0xe3, 0x0d, 0x7d, 0x94, 0xc4, 0xd9, 0x32, 0xa9, 0x65, 0x91, 0x8f, 0x53, 0xc3, 0xc6, 0x7a, 0x1c, 0x00, 0xab, 0x59, 0x56, 0x62, 0x2c, 0x1d, 0xc3, 0xbb, 0x68, 0x26, 0xc6, 0x28, 0x4f, 0x92, 0x4f, 0x8a, 0xb4, 0x9c, 0x4e, 0xd2, 0x66, 0xd6, 0xb0, 0x22, 0x4d, 0x84, 0x93, 0x63, 0x35, 0x81, 0x93, 0xcf, 0xe7, 0x8e, 0xb5, 0x09, 0x7c, 0x60, 0xac, 0x4d, 0xe0, 0xb8, 0x89, 0xae, 0xc4, 0x98, 0x46, 0x4b, 0x9c, 0x12, 0xd3, 0xa7, 0x41, 0xf0, 0xd2, 0x63, 0x16, 0xf9, 0xa2, 0x90, 0xb7, 0xd3, 0x91, 0xeb, 0x52, 0xbd, 0xa7, 0xc5, 0x21, 0xfd, 0x12, 0x4d, 0x4d, 0xe3, 0x27, 0x68, 0xae, 0x67, 0x5e, 0x61, 0x6f, 0x93, 0x79, 0x0e, 0x90, 0x53, 0xd5, 0xe3, 0xfa, 0x90, 0xb1, 0xe5, 0xd1, 0xf0, 0xe2, 0xad, 0xbe, 0x48, 0xfb, 0x33, 0xf8, 0x29, 0x9a, 0x8f, 0xc9, 0xea, 0xa4, 0x28, 0xf4, 0x57, 0x85, 0xbe, 0x91, 0x8e, 0xd6, 0x47, 0xa6, 0x87, 0x8d, 0xe9, 0x40, 0x0a, 0x6f, 0xa1, 0x62, 0x0c, 0x77, 0xec, 0x80, 0x93, 0x6f, 0x8a, 0xba, 0x94, 0x4e, 0xdd, 0xb1, 0x03, 0x9e, 0xf0, 0x51, 0x18, 0x8c, 0x48, 0x62, 0x34, 0x45, 0xfa, 0x3e, 0x94, 0x24, 0x5a, 0x0f, 0x90, 0xc2, 0x60, 0xb4, 0xf5, 0x92, 0x24, 0x1c, 0xf9, 0x26, 0x37, 0x6c, 0xeb, 0x45, 0x4d, 0xbf, 0x23, 0x75, 0x2c, 0x72, 0xa4, 0xc4, 0x68, 0x47, 0xbe, 0xcd, 0x0d, 0x73, 0xa4, 0xa8, 0x4a, 0x71, 0x64, 0x1c, 0x4e, 0x8e, 0x25, 0x1c, 0xf9, 0xee, 0xdc, 0xb1, 0xfa, 0x1d, 0xa9, 0x63, 0xf8, 0x39, 0x2a, 0xf5, 0x60, 0xa4, 0x51, 0x7c, 0x60, 0x6d, 0x3b, 0x90, 0xf7, 0xd8, 0x7b, 0xc5, 0xbc, 0x33, 0x84, 0x29, 0xe4, 0x7b, 0x91, 0x3a, 0xe4, 0x5f, 0xa6, 0xe9, 0x79, 0xdc, 0x46, 0x0b, 0x71, 0x2f, 0x6d, 0x9d, 0x9e, 0x66, 0x1f, 0x54, 0xb3, 0xbb, 0xe9, 0xcd, 0x94, 0x4b, 0x06, 0xbb, 0x11, 0x3a, 0x44, 0x50, 0xb9, 0x80, 0xa6, 0x37, 0xda, 0x3e, 0x7f, 0x65, 0x40, 0xe0, 0x7b, 0x6e, 0x00, 0x15, 0x1f, 0x2d, 0x9c, 0xf3, 0x43, 0x84, 0x31, 0x1a, 0x93, 0xb7, 0x7b, 0x46, 0xde, 0xee, 0xf2, 0x59, 0xdc, 0xfa, 0xd1, 0xf9, 0xd4, 0xb7, 0x7e, 0xf8, 0x8e, 0x97, 0x50, 0x21, 0xb0, 0xdb, 0xbe, 0x03, 0x26, 0xf7, 0x8e, 0x40, 0x5d, 0xfa, 0x39, 0x23, 0xaf, 0x62, 0xfb, 0x22, 0xf4, 0x68, 0xee, 0xe4, 0x67, 0x79, 0xe4, 0xe4, 0xac, 0x9c, 0x39, 0x3d, 0x2b, 0x67, 0x7e, 0x9c, 0x95, 0x33, 0xaf, 0x7f, 0x95, 0x47, 0x0e, 0x26, 0xe4, 0x5f, 0x8e, 0xb5, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xff, 0xc9, 0xfc, 0x0e, 0xca, 0x08, 0x00, 0x00}
