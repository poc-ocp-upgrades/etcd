package integration

import (
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3election/v3electionpb"
	"go.etcd.io/etcd/etcdserver/api/v3lock/v3lockpb"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
)

func toGRPC(c *clientv3.Client) grpcAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return grpcAPI{pb.NewClusterClient(c.ActiveConnection()), pb.NewKVClient(c.ActiveConnection()), pb.NewLeaseClient(c.ActiveConnection()), pb.NewWatchClient(c.ActiveConnection()), pb.NewMaintenanceClient(c.ActiveConnection()), pb.NewAuthClient(c.ActiveConnection()), v3lockpb.NewLockClient(c.ActiveConnection()), v3electionpb.NewElectionClient(c.ActiveConnection())}
}
func newClientV3(cfg clientv3.Config) (*clientv3.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clientv3.New(cfg)
}
