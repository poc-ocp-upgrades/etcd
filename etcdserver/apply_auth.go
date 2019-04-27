package etcdserver

import (
	"sync"
	"github.com/coreos/etcd/auth"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc"
)

type authApplierV3 struct {
	applierV3
	as		auth.AuthStore
	lessor		lease.Lessor
	mu		sync.Mutex
	authInfo	auth.AuthInfo
}

func newAuthApplierV3(as auth.AuthStore, base applierV3, lessor lease.Lessor) *authApplierV3 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &authApplierV3{applierV3: base, as: as, lessor: lessor}
}
func (aa *authApplierV3) Apply(r *pb.InternalRaftRequest) *applyResult {
	_logClusterCodePath()
	defer _logClusterCodePath()
	aa.mu.Lock()
	defer aa.mu.Unlock()
	if r.Header != nil {
		aa.authInfo.Username = r.Header.Username
		aa.authInfo.Revision = r.Header.AuthRevision
	}
	if needAdminPermission(r) {
		if err := aa.as.IsAdminPermitted(&aa.authInfo); err != nil {
			aa.authInfo.Username = ""
			aa.authInfo.Revision = 0
			return &applyResult{err: err}
		}
	}
	ret := aa.applierV3.Apply(r)
	aa.authInfo.Username = ""
	aa.authInfo.Revision = 0
	return ret
}
func (aa *authApplierV3) Put(txn mvcc.TxnWrite, r *pb.PutRequest) (*pb.PutResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := aa.as.IsPutPermitted(&aa.authInfo, r.Key); err != nil {
		return nil, err
	}
	if err := aa.checkLeasePuts(lease.LeaseID(r.Lease)); err != nil {
		return nil, err
	}
	if r.PrevKv {
		err := aa.as.IsRangePermitted(&aa.authInfo, r.Key, nil)
		if err != nil {
			return nil, err
		}
	}
	return aa.applierV3.Put(txn, r)
}
func (aa *authApplierV3) Range(txn mvcc.TxnRead, r *pb.RangeRequest) (*pb.RangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := aa.as.IsRangePermitted(&aa.authInfo, r.Key, r.RangeEnd); err != nil {
		return nil, err
	}
	return aa.applierV3.Range(txn, r)
}
func (aa *authApplierV3) DeleteRange(txn mvcc.TxnWrite, r *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := aa.as.IsDeleteRangePermitted(&aa.authInfo, r.Key, r.RangeEnd); err != nil {
		return nil, err
	}
	if r.PrevKv {
		err := aa.as.IsRangePermitted(&aa.authInfo, r.Key, r.RangeEnd)
		if err != nil {
			return nil, err
		}
	}
	return aa.applierV3.DeleteRange(txn, r)
}
func checkTxnReqsPermission(as auth.AuthStore, ai *auth.AuthInfo, reqs []*pb.RequestOp) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, requ := range reqs {
		switch tv := requ.Request.(type) {
		case *pb.RequestOp_RequestRange:
			if tv.RequestRange == nil {
				continue
			}
			if err := as.IsRangePermitted(ai, tv.RequestRange.Key, tv.RequestRange.RangeEnd); err != nil {
				return err
			}
		case *pb.RequestOp_RequestPut:
			if tv.RequestPut == nil {
				continue
			}
			if err := as.IsPutPermitted(ai, tv.RequestPut.Key); err != nil {
				return err
			}
		case *pb.RequestOp_RequestDeleteRange:
			if tv.RequestDeleteRange == nil {
				continue
			}
			if tv.RequestDeleteRange.PrevKv {
				err := as.IsRangePermitted(ai, tv.RequestDeleteRange.Key, tv.RequestDeleteRange.RangeEnd)
				if err != nil {
					return err
				}
			}
			err := as.IsDeleteRangePermitted(ai, tv.RequestDeleteRange.Key, tv.RequestDeleteRange.RangeEnd)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func checkTxnAuth(as auth.AuthStore, ai *auth.AuthInfo, rt *pb.TxnRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, c := range rt.Compare {
		if err := as.IsRangePermitted(ai, c.Key, c.RangeEnd); err != nil {
			return err
		}
	}
	if err := checkTxnReqsPermission(as, ai, rt.Success); err != nil {
		return err
	}
	if err := checkTxnReqsPermission(as, ai, rt.Failure); err != nil {
		return err
	}
	return nil
}
func (aa *authApplierV3) Txn(rt *pb.TxnRequest) (*pb.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := checkTxnAuth(aa.as, &aa.authInfo, rt); err != nil {
		return nil, err
	}
	return aa.applierV3.Txn(rt)
}
func (aa *authApplierV3) LeaseRevoke(lc *pb.LeaseRevokeRequest) (*pb.LeaseRevokeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := aa.checkLeasePuts(lease.LeaseID(lc.ID)); err != nil {
		return nil, err
	}
	return aa.applierV3.LeaseRevoke(lc)
}
func (aa *authApplierV3) checkLeasePuts(leaseID lease.LeaseID) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lease := aa.lessor.Lookup(leaseID)
	if lease != nil {
		for _, key := range lease.Keys() {
			if err := aa.as.IsPutPermitted(&aa.authInfo, []byte(key)); err != nil {
				return err
			}
		}
	}
	return nil
}
func (aa *authApplierV3) UserGet(r *pb.AuthUserGetRequest) (*pb.AuthUserGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := aa.as.IsAdminPermitted(&aa.authInfo)
	if err != nil && r.Name != aa.authInfo.Username {
		aa.authInfo.Username = ""
		aa.authInfo.Revision = 0
		return &pb.AuthUserGetResponse{}, err
	}
	return aa.applierV3.UserGet(r)
}
func (aa *authApplierV3) RoleGet(r *pb.AuthRoleGetRequest) (*pb.AuthRoleGetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := aa.as.IsAdminPermitted(&aa.authInfo)
	if err != nil && !aa.as.HasRole(aa.authInfo.Username, r.Role) {
		aa.authInfo.Username = ""
		aa.authInfo.Revision = 0
		return &pb.AuthRoleGetResponse{}, err
	}
	return aa.applierV3.RoleGet(r)
}
func needAdminPermission(r *pb.InternalRaftRequest) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case r.AuthEnable != nil:
		return true
	case r.AuthDisable != nil:
		return true
	case r.AuthUserAdd != nil:
		return true
	case r.AuthUserDelete != nil:
		return true
	case r.AuthUserChangePassword != nil:
		return true
	case r.AuthUserGrantRole != nil:
		return true
	case r.AuthUserRevokeRole != nil:
		return true
	case r.AuthRoleAdd != nil:
		return true
	case r.AuthRoleGrantPermission != nil:
		return true
	case r.AuthRoleRevokePermission != nil:
		return true
	case r.AuthRoleDelete != nil:
		return true
	case r.AuthUserList != nil:
		return true
	case r.AuthRoleList != nil:
		return true
	default:
		return false
	}
}
