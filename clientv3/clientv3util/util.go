package clientv3util

import (
	"github.com/coreos/etcd/clientv3"
)

func KeyExists(key string) clientv3.Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clientv3.Compare(clientv3.Version(key), ">", 0)
}
func KeyMissing(key string) clientv3.Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clientv3.Compare(clientv3.Version(key), "=", 0)
}
