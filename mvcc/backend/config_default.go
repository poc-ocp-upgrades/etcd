package backend

import bolt "github.com/coreos/bbolt"

var boltOpenOptions *bolt.Options = nil

func (bcfg *BackendConfig) mmapSize() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int(bcfg.MmapSize)
}
