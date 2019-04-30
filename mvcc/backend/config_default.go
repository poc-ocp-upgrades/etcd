package backend

import bolt "go.etcd.io/bbolt"

var boltOpenOptions *bolt.Options

func (bcfg *BackendConfig) mmapSize() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int(bcfg.MmapSize)
}
