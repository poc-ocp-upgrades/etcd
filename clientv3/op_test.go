package clientv3

import (
	"reflect"
	"testing"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
)

func TestOpWithSort(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opReq := OpGet("foo", WithSort(SortByKey, SortAscend), WithLimit(10)).toRequestOp().Request
	q, ok := opReq.(*pb.RequestOp_RequestRange)
	if !ok {
		t.Fatalf("expected range request, got %v", reflect.TypeOf(opReq))
	}
	req := q.RequestRange
	wreq := &pb.RangeRequest{Key: []byte("foo"), SortOrder: pb.RangeRequest_NONE, Limit: 10}
	if !reflect.DeepEqual(req, wreq) {
		t.Fatalf("expected %+v, got %+v", wreq, req)
	}
}
