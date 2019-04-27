package etcdserver

import pb "github.com/coreos/etcd/etcdserver/etcdserverpb"

const (
	DefaultQuotaBytes	= int64(2 * 1024 * 1024 * 1024)
	MaxQuotaBytes		= int64(8 * 1024 * 1024 * 1024)
)

type Quota interface {
	Available(req interface{}) bool
	Cost(req interface{}) int
	Remaining() int64
}
type passthroughQuota struct{}

func (*passthroughQuota) Available(interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (*passthroughQuota) Cost(interface{}) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (*passthroughQuota) Remaining() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 1
}

type backendQuota struct {
	s		*EtcdServer
	maxBackendBytes	int64
}

const (
	leaseOverhead	= 64
	kvOverhead	= 256
)

func NewBackendQuota(s *EtcdServer) Quota {
	_logClusterCodePath()
	defer _logClusterCodePath()
	quotaBackendBytes.Set(float64(s.Cfg.QuotaBackendBytes))
	if s.Cfg.QuotaBackendBytes < 0 {
		plog.Warningf("disabling backend quota")
		return &passthroughQuota{}
	}
	if s.Cfg.QuotaBackendBytes == 0 {
		quotaBackendBytes.Set(float64(DefaultQuotaBytes))
		return &backendQuota{s, DefaultQuotaBytes}
	}
	if s.Cfg.QuotaBackendBytes > MaxQuotaBytes {
		plog.Warningf("backend quota %v exceeds maximum recommended quota %v", s.Cfg.QuotaBackendBytes, MaxQuotaBytes)
	}
	return &backendQuota{s, s.Cfg.QuotaBackendBytes}
}
func (b *backendQuota) Available(v interface{}) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.s.Backend().Size()+int64(b.Cost(v)) < b.maxBackendBytes
}
func (b *backendQuota) Cost(v interface{}) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch r := v.(type) {
	case *pb.PutRequest:
		return costPut(r)
	case *pb.TxnRequest:
		return costTxn(r)
	case *pb.LeaseGrantRequest:
		return leaseOverhead
	default:
		panic("unexpected cost")
	}
}
func costPut(r *pb.PutRequest) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kvOverhead + len(r.Key) + len(r.Value)
}
func costTxnReq(u *pb.RequestOp) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := u.GetRequestPut()
	if r == nil {
		return 0
	}
	return costPut(r)
}
func costTxn(r *pb.TxnRequest) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sizeSuccess := 0
	for _, u := range r.Success {
		sizeSuccess += costTxnReq(u)
	}
	sizeFailure := 0
	for _, u := range r.Failure {
		sizeFailure += costTxnReq(u)
	}
	if sizeFailure > sizeSuccess {
		return sizeFailure
	}
	return sizeSuccess
}
func (b *backendQuota) Remaining() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return b.maxBackendBytes - b.s.Backend().Size()
}
