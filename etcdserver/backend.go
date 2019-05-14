package etcdserver

import (
	"fmt"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"os"
	"time"
)

func newBackend(cfg ServerConfig) backend.Backend {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bcfg := backend.DefaultBackendConfig()
	bcfg.Path = cfg.backendPath()
	if cfg.QuotaBackendBytes > 0 && cfg.QuotaBackendBytes != DefaultQuotaBytes {
		bcfg.MmapSize = uint64(cfg.QuotaBackendBytes + cfg.QuotaBackendBytes/10)
	}
	return backend.New(bcfg)
}
func openSnapshotBackend(cfg ServerConfig, ss *snap.Snapshotter, snapshot raftpb.Snapshot) (backend.Backend, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	snapPath, err := ss.DBFilePath(snapshot.Metadata.Index)
	if err != nil {
		return nil, fmt.Errorf("database snapshot file path error: %v", err)
	}
	if err := os.Rename(snapPath, cfg.backendPath()); err != nil {
		return nil, fmt.Errorf("rename snapshot file error: %v", err)
	}
	return openBackend(cfg), nil
}
func openBackend(cfg ServerConfig) backend.Backend {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fn := cfg.backendPath()
	beOpened := make(chan backend.Backend)
	go func() {
		beOpened <- newBackend(cfg)
	}()
	select {
	case be := <-beOpened:
		return be
	case <-time.After(10 * time.Second):
		plog.Warningf("another etcd process is using %q and holds the file lock, or loading backend file is taking >10 seconds", fn)
		plog.Warningf("waiting for it to exit before starting...")
	}
	return <-beOpened
}
func recoverSnapshotBackend(cfg ServerConfig, oldbe backend.Backend, snapshot raftpb.Snapshot) (backend.Backend, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cIndex consistentIndex
	kv := mvcc.New(oldbe, &lease.FakeLessor{}, &cIndex)
	defer kv.Close()
	if snapshot.Metadata.Index <= kv.ConsistentIndex() {
		return oldbe, nil
	}
	oldbe.Close()
	return openSnapshotBackend(cfg, snap.New(cfg.SnapDir()), snapshot)
}
