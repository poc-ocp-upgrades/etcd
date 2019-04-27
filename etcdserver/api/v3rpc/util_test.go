package v3rpc

import (
	"context"
	"errors"
	"testing"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/mvcc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGRPCError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := []struct {
		err	error
		exp	error
	}{{err: mvcc.ErrCompacted, exp: rpctypes.ErrGRPCCompacted}, {err: mvcc.ErrFutureRev, exp: rpctypes.ErrGRPCFutureRev}, {err: context.Canceled, exp: context.Canceled}, {err: context.DeadlineExceeded, exp: context.DeadlineExceeded}, {err: errors.New("foo"), exp: status.Error(codes.Unknown, "foo")}}
	for i := range tt {
		if err := togRPCError(tt[i].err); err != tt[i].exp {
			if _, ok := status.FromError(err); ok {
				if err.Error() == tt[i].exp.Error() {
					continue
				}
			}
			t.Errorf("#%d: got %v, expected %v", i, err, tt[i].exp)
		}
	}
}
