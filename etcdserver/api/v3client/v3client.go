package v3client

import (
	"context"
	"time"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api/v3rpc"
	"go.etcd.io/etcd/proxy/grpcproxy/adapter"
)

func New(s *etcdserver.EtcdServer) *clientv3.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := clientv3.NewCtxClient(context.Background())
	kvc := adapter.KvServerToKvClient(v3rpc.NewQuotaKVServer(s))
	c.KV = clientv3.NewKVFromKVClient(kvc, c)
	lc := adapter.LeaseServerToLeaseClient(v3rpc.NewQuotaLeaseServer(s))
	c.Lease = clientv3.NewLeaseFromLeaseClient(lc, c, time.Second)
	wc := adapter.WatchServerToWatchClient(v3rpc.NewWatchServer(s))
	c.Watcher = &watchWrapper{clientv3.NewWatchFromWatchClient(wc, c)}
	mc := adapter.MaintenanceServerToMaintenanceClient(v3rpc.NewMaintenanceServer(s))
	c.Maintenance = clientv3.NewMaintenanceFromMaintenanceClient(mc, c)
	clc := adapter.ClusterServerToClusterClient(v3rpc.NewClusterServer(s))
	c.Cluster = clientv3.NewClusterFromClusterClient(clc, c)
	return c
}

type blankContext struct{ context.Context }

func (*blankContext) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "(blankCtx)"
}

type watchWrapper struct{ clientv3.Watcher }

func (ww *watchWrapper) Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ww.Watcher.Watch(&blankContext{ctx}, key, opts...)
}
