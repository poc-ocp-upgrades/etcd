package leasing

import (
	"bytes"
	v3 "github.com/coreos/etcd/clientv3"
	v3pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

func compareInt64(a, b int64) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
func evalCmp(resp *v3.GetResponse, tcmp v3.Cmp) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var result int
	if len(resp.Kvs) != 0 {
		kv := resp.Kvs[0]
		switch tcmp.Target {
		case v3pb.Compare_VALUE:
			if tv, _ := tcmp.TargetUnion.(*v3pb.Compare_Value); tv != nil {
				result = bytes.Compare(kv.Value, tv.Value)
			}
		case v3pb.Compare_CREATE:
			if tv, _ := tcmp.TargetUnion.(*v3pb.Compare_CreateRevision); tv != nil {
				result = compareInt64(kv.CreateRevision, tv.CreateRevision)
			}
		case v3pb.Compare_MOD:
			if tv, _ := tcmp.TargetUnion.(*v3pb.Compare_ModRevision); tv != nil {
				result = compareInt64(kv.ModRevision, tv.ModRevision)
			}
		case v3pb.Compare_VERSION:
			if tv, _ := tcmp.TargetUnion.(*v3pb.Compare_Version); tv != nil {
				result = compareInt64(kv.Version, tv.Version)
			}
		}
	}
	switch tcmp.Result {
	case v3pb.Compare_EQUAL:
		return result == 0
	case v3pb.Compare_NOT_EQUAL:
		return result != 0
	case v3pb.Compare_GREATER:
		return result > 0
	case v3pb.Compare_LESS:
		return result < 0
	}
	return true
}
func gatherOps(ops []v3.Op) (ret []v3.Op) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, op := range ops {
		if !op.IsTxn() {
			ret = append(ret, op)
			continue
		}
		_, thenOps, elseOps := op.Txn()
		ret = append(ret, gatherOps(append(thenOps, elseOps...))...)
	}
	return ret
}
func gatherResponseOps(resp []*v3pb.ResponseOp, ops []v3.Op) (ret []v3.Op) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, op := range ops {
		if !op.IsTxn() {
			ret = append(ret, op)
			continue
		}
		_, thenOps, elseOps := op.Txn()
		if txnResp := resp[i].GetResponseTxn(); txnResp.Succeeded {
			ret = append(ret, gatherResponseOps(txnResp.Responses, thenOps)...)
		} else {
			ret = append(ret, gatherResponseOps(txnResp.Responses, elseOps)...)
		}
	}
	return ret
}
func copyHeader(hdr *v3pb.ResponseHeader) *v3pb.ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	h := *hdr
	return &h
}
func closeAll(chs []chan<- struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, ch := range chs {
		close(ch)
	}
}
