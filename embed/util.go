package embed

import (
	"github.com/coreos/etcd/wal"
	"path/filepath"
)

func isMemberInitialized(cfg *Config) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	waldir := cfg.WalDir
	if waldir == "" {
		waldir = filepath.Join(cfg.Dir, "member", "wal")
	}
	return wal.Exist(waldir)
}
