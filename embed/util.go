package embed

import (
	"path/filepath"
	"github.com/coreos/etcd/wal"
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
