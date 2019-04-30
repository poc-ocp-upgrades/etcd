package backend

import (
	"syscall"
	bolt "go.etcd.io/bbolt"
)

var boltOpenOptions = &bolt.Options{MmapFlags: syscall.MAP_POPULATE, NoFreelistSync: true}

func (bcfg *BackendConfig) mmapSize() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int(bcfg.MmapSize)
}
