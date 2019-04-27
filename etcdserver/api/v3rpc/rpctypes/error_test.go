package rpctypes

import (
	"testing"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestConvert(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e1 := status.New(codes.InvalidArgument, "etcdserver: key is not provided").Err()
	e2 := ErrGRPCEmptyKey
	e3 := ErrEmptyKey
	if e1.Error() != e2.Error() {
		t.Fatalf("expected %q == %q", e1.Error(), e2.Error())
	}
	ev1, _ := status.FromError(e1)
	if ev1.Code() != e3.(EtcdError).Code() {
		t.Fatalf("expected them to be equal, got %v / %v", ev1.Code(), e3.(EtcdError).Code())
	}
	if e1.Error() == e3.Error() {
		t.Fatalf("expected %q != %q", e1.Error(), e3.Error())
	}
	ev2, _ := status.FromError(e2)
	if ev2.Code() != e3.(EtcdError).Code() {
		t.Fatalf("expected them to be equal, got %v / %v", ev2.Code(), e3.(EtcdError).Code())
	}
}
