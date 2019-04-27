package store_test

import (
	"io/ioutil"
	"testing"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v2v3"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/store"
	"github.com/coreos/pkg/capnslog"
	"google.golang.org/grpc/grpclog"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	capnslog.SetGlobalLogLevel(capnslog.CRITICAL)
	clientv3.SetLogger(grpclog.NewLoggerV2(ioutil.Discard, ioutil.Discard, ioutil.Discard))
}

type v2v3TestStore struct {
	store.Store
	clus	*integration.ClusterV3
	t	*testing.T
}

func (s *v2v3TestStore) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.clus.Terminate(s.t)
}
func newTestStore(t *testing.T, ns ...string) StoreCloser {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	return &v2v3TestStore{v2v3.NewStore(clus.Client(0), "/v2/"), clus, t}
}
