package integration

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/namespace"
	"github.com/coreos/etcd/proxy/grpcproxy"
	"github.com/coreos/etcd/proxy/grpcproxy/adapter"
	"sync"
)

var (
	pmu     sync.Mutex
	proxies map[*clientv3.Client]grpcClientProxy = make(map[*clientv3.Client]grpcClientProxy)
)

const proxyNamespace = "proxy-namespace"

type grpcClientProxy struct {
	grpc    grpcAPI
	wdonec  <-chan struct{}
	kvdonec <-chan struct{}
	lpdonec <-chan struct{}
}

func toGRPC(c *clientv3.Client) grpcAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pmu.Lock()
	defer pmu.Unlock()
	if v, ok := proxies[c]; ok {
		return v.grpc
	}
	c.KV = namespace.NewKV(c.KV, proxyNamespace)
	c.Watcher = namespace.NewWatcher(c.Watcher, proxyNamespace)
	c.Lease = namespace.NewLease(c.Lease, proxyNamespace)
	kvp, kvpch := grpcproxy.NewKvProxy(c)
	wp, wpch := grpcproxy.NewWatchProxy(c)
	lp, lpch := grpcproxy.NewLeaseProxy(c)
	mp := grpcproxy.NewMaintenanceProxy(c)
	clp, _ := grpcproxy.NewClusterProxy(c, "", "")
	authp := grpcproxy.NewAuthProxy(c)
	lockp := grpcproxy.NewLockProxy(c)
	electp := grpcproxy.NewElectionProxy(c)
	grpc := grpcAPI{adapter.ClusterServerToClusterClient(clp), adapter.KvServerToKvClient(kvp), adapter.LeaseServerToLeaseClient(lp), adapter.WatchServerToWatchClient(wp), adapter.MaintenanceServerToMaintenanceClient(mp), adapter.AuthServerToAuthClient(authp), adapter.LockServerToLockClient(lockp), adapter.ElectionServerToElectionClient(electp)}
	proxies[c] = grpcClientProxy{grpc: grpc, wdonec: wpch, kvdonec: kvpch, lpdonec: lpch}
	return grpc
}

type proxyCloser struct {
	clientv3.Watcher
	wdonec  <-chan struct{}
	kvdonec <-chan struct{}
	lclose  func()
	lpdonec <-chan struct{}
}

func (pc *proxyCloser) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	<-pc.kvdonec
	err := pc.Watcher.Close()
	<-pc.wdonec
	pc.lclose()
	<-pc.lpdonec
	return err
}
func newClientV3(cfg clientv3.Config) (*clientv3.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	rpc := toGRPC(c)
	c.KV = clientv3.NewKVFromKVClient(rpc.KV, c)
	pmu.Lock()
	lc := c.Lease
	c.Lease = clientv3.NewLeaseFromLeaseClient(rpc.Lease, c, cfg.DialTimeout)
	c.Watcher = &proxyCloser{Watcher: clientv3.NewWatchFromWatchClient(rpc.Watch, c), wdonec: proxies[c].wdonec, kvdonec: proxies[c].kvdonec, lclose: func() {
		lc.Close()
	}, lpdonec: proxies[c].lpdonec}
	pmu.Unlock()
	return c, nil
}
