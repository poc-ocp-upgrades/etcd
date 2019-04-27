package backend

import (
	"syscall"
	bolt "github.com/coreos/bbolt"
)

var boltOpenOptions = &bolt.Options{MmapFlags: syscall.MAP_POPULATE, NoFreelistSync: true}

func (bcfg *BackendConfig) mmapSize() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int(bcfg.MmapSize)
}
