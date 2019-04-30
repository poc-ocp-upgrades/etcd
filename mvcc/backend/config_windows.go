package backend

import bolt "go.etcd.io/bbolt"

var boltOpenOptions *bolt.Options = nil

func (bcfg *BackendConfig) mmapSize() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
