package etcdserver

import (
	"io"
	"go.etcd.io/etcd/etcdserver/api/snap"
	"go.etcd.io/etcd/mvcc/backend"
	"go.etcd.io/etcd/raft/raftpb"
	humanize "github.com/dustin/go-humanize"
	"go.uber.org/zap"
)

func (s *EtcdServer) createMergedSnapshotMessage(m raftpb.Message, snapt, snapi uint64, confState raftpb.ConfState) snap.Message {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clone := s.v2store.Clone()
	d, err := clone.SaveNoCopy()
	if err != nil {
		if lg := s.getLogger(); lg != nil {
			lg.Panic("failed to save v2 store data", zap.Error(err))
		} else {
			plog.Panicf("store save should never fail: %v", err)
		}
	}
	s.KV().Commit()
	dbsnap := s.be.Snapshot()
	rc := newSnapshotReaderCloser(s.getLogger(), dbsnap)
	snapshot := raftpb.Snapshot{Metadata: raftpb.SnapshotMetadata{Index: snapi, Term: snapt, ConfState: confState}, Data: d}
	m.Snapshot = snapshot
	return *snap.NewMessage(m, rc, dbsnap.Size())
}
func newSnapshotReaderCloser(lg *zap.Logger, snapshot backend.Snapshot) io.ReadCloser {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr, pw := io.Pipe()
	go func() {
		n, err := snapshot.WriteTo(pw)
		if err == nil {
			if lg != nil {
				lg.Info("sent database snapshot to writer", zap.Int64("bytes", n), zap.String("size", humanize.Bytes(uint64(n))))
			} else {
				plog.Infof("wrote database snapshot out [total bytes: %d]", n)
			}
		} else {
			if lg != nil {
				lg.Warn("failed to send database snapshot to writer", zap.String("size", humanize.Bytes(uint64(n))), zap.Error(err))
			} else {
				plog.Warningf("failed to write database snapshot out [written bytes: %d]: %v", n, err)
			}
		}
		pw.CloseWithError(err)
		err = snapshot.Close()
		if err != nil {
			if lg != nil {
				lg.Panic("failed to close database snapshot", zap.Error(err))
			} else {
				plog.Panicf("failed to close database snapshot: %v", err)
			}
		}
	}()
	return pr
}
