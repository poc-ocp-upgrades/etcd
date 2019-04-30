package etcdserver

import (
	"sync"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	humanize "github.com/dustin/go-humanize"
	"go.uber.org/zap"
)

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

var (
	quotaLogOnce		sync.Once
	DefaultQuotaSize	= humanize.Bytes(uint64(DefaultQuotaBytes))
	maxQuotaSize		= humanize.Bytes(uint64(MaxQuotaBytes))
)

func NewBackendQuota(s *EtcdServer, name string) Quota {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg := s.getLogger()
	quotaBackendBytes.Set(float64(s.Cfg.QuotaBackendBytes))
	if s.Cfg.QuotaBackendBytes < 0 {
		quotaLogOnce.Do(func() {
			if lg != nil {
				lg.Info("disabled backend quota", zap.String("quota-name", name), zap.Int64("quota-size-bytes", s.Cfg.QuotaBackendBytes))
			} else {
				plog.Warningf("disabling backend quota")
			}
		})
		return &passthroughQuota{}
	}
	if s.Cfg.QuotaBackendBytes == 0 {
		quotaLogOnce.Do(func() {
			if lg != nil {
				lg.Info("enabled backend quota with default value", zap.String("quota-name", name), zap.Int64("quota-size-bytes", DefaultQuotaBytes), zap.String("quota-size", DefaultQuotaSize))
			}
		})
		quotaBackendBytes.Set(float64(DefaultQuotaBytes))
		return &backendQuota{s, DefaultQuotaBytes}
	}
	quotaLogOnce.Do(func() {
		if s.Cfg.QuotaBackendBytes > MaxQuotaBytes {
			if lg != nil {
				lg.Warn("quota exceeds the maximum value", zap.String("quota-name", name), zap.Int64("quota-size-bytes", s.Cfg.QuotaBackendBytes), zap.String("quota-size", humanize.Bytes(uint64(s.Cfg.QuotaBackendBytes))), zap.Int64("quota-maximum-size-bytes", MaxQuotaBytes), zap.String("quota-maximum-size", maxQuotaSize))
			} else {
				plog.Warningf("backend quota %v exceeds maximum recommended quota %v", s.Cfg.QuotaBackendBytes, MaxQuotaBytes)
			}
		}
		if lg != nil {
			lg.Info("enabled backend quota", zap.String("quota-name", name), zap.Int64("quota-size-bytes", s.Cfg.QuotaBackendBytes), zap.String("quota-size", humanize.Bytes(uint64(s.Cfg.QuotaBackendBytes))))
		}
	})
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
