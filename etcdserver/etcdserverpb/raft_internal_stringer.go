package etcdserverpb

import (
	"fmt"
	"strings"
	proto "github.com/golang/protobuf/proto"
)

type InternalRaftStringer struct{ Request *InternalRaftRequest }

func (as *InternalRaftStringer) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case as.Request.LeaseGrant != nil:
		return fmt.Sprintf("header:<%s> lease_grant:<ttl:%d-second id:%016x>", as.Request.Header.String(), as.Request.LeaseGrant.TTL, as.Request.LeaseGrant.ID)
	case as.Request.LeaseRevoke != nil:
		return fmt.Sprintf("header:<%s> lease_revoke:<id:%016x>", as.Request.Header.String(), as.Request.LeaseRevoke.ID)
	case as.Request.Authenticate != nil:
		return fmt.Sprintf("header:<%s> authenticate:<name:%s simple_token:%s>", as.Request.Header.String(), as.Request.Authenticate.Name, as.Request.Authenticate.SimpleToken)
	case as.Request.AuthUserAdd != nil:
		return fmt.Sprintf("header:<%s> auth_user_add:<name:%s>", as.Request.Header.String(), as.Request.AuthUserAdd.Name)
	case as.Request.AuthUserChangePassword != nil:
		return fmt.Sprintf("header:<%s> auth_user_change_password:<name:%s>", as.Request.Header.String(), as.Request.AuthUserChangePassword.Name)
	case as.Request.Put != nil:
		return fmt.Sprintf("header:<%s> put:<%s>", as.Request.Header.String(), NewLoggablePutRequest(as.Request.Put).String())
	case as.Request.Txn != nil:
		return fmt.Sprintf("header:<%s> txn:<%s>", as.Request.Header.String(), NewLoggableTxnRequest(as.Request.Txn).String())
	default:
	}
	return as.Request.String()
}

type txnRequestStringer struct{ Request *TxnRequest }

func NewLoggableTxnRequest(request *TxnRequest) *txnRequestStringer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &txnRequestStringer{request}
}
func (as *txnRequestStringer) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var compare []string
	for _, c := range as.Request.Compare {
		switch cv := c.TargetUnion.(type) {
		case *Compare_Value:
			compare = append(compare, newLoggableValueCompare(c, cv).String())
		default:
			compare = append(compare, c.String())
		}
	}
	var success []string
	for _, s := range as.Request.Success {
		success = append(success, newLoggableRequestOp(s).String())
	}
	var failure []string
	for _, f := range as.Request.Failure {
		failure = append(failure, newLoggableRequestOp(f).String())
	}
	return fmt.Sprintf("compare:<%s> success:<%s> failure:<%s>", strings.Join(compare, " "), strings.Join(success, " "), strings.Join(failure, " "))
}

type requestOpStringer struct{ Op *RequestOp }

func newLoggableRequestOp(op *RequestOp) *requestOpStringer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &requestOpStringer{op}
}
func (as *requestOpStringer) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch op := as.Op.Request.(type) {
	case *RequestOp_RequestPut:
		return fmt.Sprintf("request_put:<%s>", NewLoggablePutRequest(op.RequestPut).String())
	case *RequestOp_RequestTxn:
		return fmt.Sprintf("request_txn:<%s>", NewLoggableTxnRequest(op.RequestTxn).String())
	default:
	}
	return as.Op.String()
}

type loggableValueCompare struct {
	Result		Compare_CompareResult	`protobuf:"varint,1,opt,name=result,proto3,enum=etcdserverpb.Compare_CompareResult"`
	Target		Compare_CompareTarget	`protobuf:"varint,2,opt,name=target,proto3,enum=etcdserverpb.Compare_CompareTarget"`
	Key		[]byte			`protobuf:"bytes,3,opt,name=key,proto3"`
	ValueSize	int			`protobuf:"bytes,7,opt,name=value_size,proto3"`
	RangeEnd	[]byte			`protobuf:"bytes,64,opt,name=range_end,proto3"`
}

func newLoggableValueCompare(c *Compare, cv *Compare_Value) *loggableValueCompare {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &loggableValueCompare{c.Result, c.Target, c.Key, len(cv.Value), c.RangeEnd}
}
func (m *loggableValueCompare) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = loggableValueCompare{}
}
func (m *loggableValueCompare) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*loggableValueCompare) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}

type loggablePutRequest struct {
	Key		[]byte	`protobuf:"bytes,1,opt,name=key,proto3"`
	ValueSize	int	`protobuf:"varint,2,opt,name=value_size,proto3"`
	Lease		int64	`protobuf:"varint,3,opt,name=lease,proto3"`
	PrevKv		bool	`protobuf:"varint,4,opt,name=prev_kv,proto3"`
	IgnoreValue	bool	`protobuf:"varint,5,opt,name=ignore_value,proto3"`
	IgnoreLease	bool	`protobuf:"varint,6,opt,name=ignore_lease,proto3"`
}

func NewLoggablePutRequest(request *PutRequest) *loggablePutRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &loggablePutRequest{request.Key, len(request.Value), request.Lease, request.PrevKv, request.IgnoreValue, request.IgnoreLease}
}
func (m *loggablePutRequest) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = loggablePutRequest{}
}
func (m *loggablePutRequest) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*loggablePutRequest) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
