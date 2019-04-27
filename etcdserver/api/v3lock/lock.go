package v3lock

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/etcdserver/api/v3lock/v3lockpb"
)

type lockServer struct{ c *clientv3.Client }

func NewLockServer(c *clientv3.Client) v3lockpb.LockServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &lockServer{c}
}
func (ls *lockServer) Lock(ctx context.Context, req *v3lockpb.LockRequest) (*v3lockpb.LockResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := concurrency.NewSession(ls.c, concurrency.WithLease(clientv3.LeaseID(req.Lease)), concurrency.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	s.Orphan()
	m := concurrency.NewMutex(s, string(req.Name))
	if err = m.Lock(ctx); err != nil {
		return nil, err
	}
	return &v3lockpb.LockResponse{Header: m.Header(), Key: []byte(m.Key())}, nil
}
func (ls *lockServer) Unlock(ctx context.Context, req *v3lockpb.UnlockRequest) (*v3lockpb.UnlockResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := ls.c.Delete(ctx, string(req.Key))
	if err != nil {
		return nil, err
	}
	return &v3lockpb.UnlockResponse{Header: resp.Header}, nil
}
