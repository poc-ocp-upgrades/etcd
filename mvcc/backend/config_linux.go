package backend

import (
	bolt "github.com/coreos/bbolt"
	"syscall"
)

var boltOpenOptions = &bolt.Options{MmapFlags: syscall.MAP_POPULATE, NoFreelistSync: true}

func (bcfg *BackendConfig) mmapSize() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int(bcfg.MmapSize)
}
