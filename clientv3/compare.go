package clientv3

import (
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
)

type CompareTarget int
type CompareResult int

const (
	CompareVersion	CompareTarget	= iota
	CompareCreated
	CompareModified
	CompareValue
)

type Cmp pb.Compare

func Compare(cmp Cmp, result string, v interface{}) Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var r pb.Compare_CompareResult
	switch result {
	case "=":
		r = pb.Compare_EQUAL
	case "!=":
		r = pb.Compare_NOT_EQUAL
	case ">":
		r = pb.Compare_GREATER
	case "<":
		r = pb.Compare_LESS
	default:
		panic("Unknown result op")
	}
	cmp.Result = r
	switch cmp.Target {
	case pb.Compare_VALUE:
		val, ok := v.(string)
		if !ok {
			panic("bad compare value")
		}
		cmp.TargetUnion = &pb.Compare_Value{Value: []byte(val)}
	case pb.Compare_VERSION:
		cmp.TargetUnion = &pb.Compare_Version{Version: mustInt64(v)}
	case pb.Compare_CREATE:
		cmp.TargetUnion = &pb.Compare_CreateRevision{CreateRevision: mustInt64(v)}
	case pb.Compare_MOD:
		cmp.TargetUnion = &pb.Compare_ModRevision{ModRevision: mustInt64(v)}
	case pb.Compare_LEASE:
		cmp.TargetUnion = &pb.Compare_Lease{Lease: mustInt64orLeaseID(v)}
	default:
		panic("Unknown compare type")
	}
	return cmp
}
func Value(key string) Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Cmp{Key: []byte(key), Target: pb.Compare_VALUE}
}
func Version(key string) Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Cmp{Key: []byte(key), Target: pb.Compare_VERSION}
}
func CreateRevision(key string) Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Cmp{Key: []byte(key), Target: pb.Compare_CREATE}
}
func ModRevision(key string) Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Cmp{Key: []byte(key), Target: pb.Compare_MOD}
}
func LeaseValue(key string) Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Cmp{Key: []byte(key), Target: pb.Compare_LEASE}
}
func (cmp *Cmp) KeyBytes() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cmp.Key
}
func (cmp *Cmp) WithKeyBytes(key []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmp.Key = key
}
func (cmp *Cmp) ValueBytes() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if tu, ok := cmp.TargetUnion.(*pb.Compare_Value); ok {
		return tu.Value
	}
	return nil
}
func (cmp *Cmp) WithValueBytes(v []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmp.TargetUnion.(*pb.Compare_Value).Value = v
}
func (cmp Cmp) WithRange(end string) Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmp.RangeEnd = []byte(end)
	return cmp
}
func (cmp Cmp) WithPrefix() Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmp.RangeEnd = getPrefix(cmp.Key)
	return cmp
}
func mustInt64(val interface{}) int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if v, ok := val.(int64); ok {
		return v
	}
	if v, ok := val.(int); ok {
		return int64(v)
	}
	panic("bad value")
}
func mustInt64orLeaseID(val interface{}) int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if v, ok := val.(LeaseID); ok {
		return int64(v)
	}
	return mustInt64(val)
}
