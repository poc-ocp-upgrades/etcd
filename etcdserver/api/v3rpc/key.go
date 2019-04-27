package v3rpc

import (
	"context"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/adt"
	"github.com/coreos/pkg/capnslog"
)

var (
	plog = capnslog.NewPackageLogger("github.com/coreos/etcd", "etcdserver/api/v3rpc")
)

type kvServer struct {
	hdr		header
	kv		etcdserver.RaftKV
	maxTxnOps	uint
}

func NewKVServer(s *etcdserver.EtcdServer) pb.KVServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &kvServer{hdr: newHeader(s), kv: s, maxTxnOps: s.Cfg.MaxTxnOps}
}
func (s *kvServer) Range(ctx context.Context, r *pb.RangeRequest) (*pb.RangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := checkRangeRequest(r); err != nil {
		return nil, err
	}
	resp, err := s.kv.Range(ctx, r)
	if err != nil {
		return nil, togRPCError(err)
	}
	s.hdr.fill(resp.Header)
	return resp, nil
}
func (s *kvServer) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := checkPutRequest(r); err != nil {
		return nil, err
	}
	resp, err := s.kv.Put(ctx, r)
	if err != nil {
		return nil, togRPCError(err)
	}
	s.hdr.fill(resp.Header)
	return resp, nil
}
func (s *kvServer) DeleteRange(ctx context.Context, r *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := checkDeleteRequest(r); err != nil {
		return nil, err
	}
	resp, err := s.kv.DeleteRange(ctx, r)
	if err != nil {
		return nil, togRPCError(err)
	}
	s.hdr.fill(resp.Header)
	return resp, nil
}
func (s *kvServer) Txn(ctx context.Context, r *pb.TxnRequest) (*pb.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := checkTxnRequest(r, int(s.maxTxnOps)); err != nil {
		return nil, err
	}
	if _, _, err := checkIntervals(r.Success); err != nil {
		return nil, err
	}
	if _, _, err := checkIntervals(r.Failure); err != nil {
		return nil, err
	}
	resp, err := s.kv.Txn(ctx, r)
	if err != nil {
		return nil, togRPCError(err)
	}
	s.hdr.fill(resp.Header)
	return resp, nil
}
func (s *kvServer) Compact(ctx context.Context, r *pb.CompactionRequest) (*pb.CompactionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := s.kv.Compact(ctx, r)
	if err != nil {
		return nil, togRPCError(err)
	}
	s.hdr.fill(resp.Header)
	return resp, nil
}
func checkRangeRequest(r *pb.RangeRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(r.Key) == 0 {
		return rpctypes.ErrGRPCEmptyKey
	}
	return nil
}
func checkPutRequest(r *pb.PutRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(r.Key) == 0 {
		return rpctypes.ErrGRPCEmptyKey
	}
	if r.IgnoreValue && len(r.Value) != 0 {
		return rpctypes.ErrGRPCValueProvided
	}
	if r.IgnoreLease && r.Lease != 0 {
		return rpctypes.ErrGRPCLeaseProvided
	}
	return nil
}
func checkDeleteRequest(r *pb.DeleteRangeRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(r.Key) == 0 {
		return rpctypes.ErrGRPCEmptyKey
	}
	return nil
}
func checkTxnRequest(r *pb.TxnRequest, maxTxnOps int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	opc := len(r.Compare)
	if opc < len(r.Success) {
		opc = len(r.Success)
	}
	if opc < len(r.Failure) {
		opc = len(r.Failure)
	}
	if opc > maxTxnOps {
		return rpctypes.ErrGRPCTooManyOps
	}
	for _, c := range r.Compare {
		if len(c.Key) == 0 {
			return rpctypes.ErrGRPCEmptyKey
		}
	}
	for _, u := range r.Success {
		if err := checkRequestOp(u, maxTxnOps-opc); err != nil {
			return err
		}
	}
	for _, u := range r.Failure {
		if err := checkRequestOp(u, maxTxnOps-opc); err != nil {
			return err
		}
	}
	return nil
}
func checkIntervals(reqs []*pb.RequestOp) (map[string]struct{}, adt.IntervalTree, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var dels adt.IntervalTree
	for _, req := range reqs {
		tv, ok := req.Request.(*pb.RequestOp_RequestDeleteRange)
		if !ok {
			continue
		}
		dreq := tv.RequestDeleteRange
		if dreq == nil {
			continue
		}
		var iv adt.Interval
		if len(dreq.RangeEnd) != 0 {
			iv = adt.NewStringAffineInterval(string(dreq.Key), string(dreq.RangeEnd))
		} else {
			iv = adt.NewStringAffinePoint(string(dreq.Key))
		}
		dels.Insert(iv, struct{}{})
	}
	puts := make(map[string]struct{})
	for _, req := range reqs {
		tv, ok := req.Request.(*pb.RequestOp_RequestTxn)
		if !ok {
			continue
		}
		putsThen, delsThen, err := checkIntervals(tv.RequestTxn.Success)
		if err != nil {
			return nil, dels, err
		}
		putsElse, delsElse, err := checkIntervals(tv.RequestTxn.Failure)
		if err != nil {
			return nil, dels, err
		}
		for k := range putsThen {
			if _, ok := puts[k]; ok {
				return nil, dels, rpctypes.ErrGRPCDuplicateKey
			}
			if dels.Intersects(adt.NewStringAffinePoint(k)) {
				return nil, dels, rpctypes.ErrGRPCDuplicateKey
			}
			puts[k] = struct{}{}
		}
		for k := range putsElse {
			if _, ok := puts[k]; ok {
				if _, isSafe := putsThen[k]; !isSafe {
					return nil, dels, rpctypes.ErrGRPCDuplicateKey
				}
			}
			if dels.Intersects(adt.NewStringAffinePoint(k)) {
				return nil, dels, rpctypes.ErrGRPCDuplicateKey
			}
			puts[k] = struct{}{}
		}
		dels.Union(delsThen, adt.NewStringAffineInterval("\x00", ""))
		dels.Union(delsElse, adt.NewStringAffineInterval("\x00", ""))
	}
	for _, req := range reqs {
		tv, ok := req.Request.(*pb.RequestOp_RequestPut)
		if !ok || tv.RequestPut == nil {
			continue
		}
		k := string(tv.RequestPut.Key)
		if _, ok := puts[k]; ok {
			return nil, dels, rpctypes.ErrGRPCDuplicateKey
		}
		if dels.Intersects(adt.NewStringAffinePoint(k)) {
			return nil, dels, rpctypes.ErrGRPCDuplicateKey
		}
		puts[k] = struct{}{}
	}
	return puts, dels, nil
}
func checkRequestOp(u *pb.RequestOp, maxTxnOps int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch uv := u.Request.(type) {
	case *pb.RequestOp_RequestRange:
		return checkRangeRequest(uv.RequestRange)
	case *pb.RequestOp_RequestPut:
		return checkPutRequest(uv.RequestPut)
	case *pb.RequestOp_RequestDeleteRange:
		return checkDeleteRequest(uv.RequestDeleteRange)
	case *pb.RequestOp_RequestTxn:
		return checkTxnRequest(uv.RequestTxn, maxTxnOps)
	default:
		return rpctypes.ErrGRPCKeyNotFound
	}
}
