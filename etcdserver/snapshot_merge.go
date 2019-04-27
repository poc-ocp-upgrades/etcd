package etcdserver

import (
	"io"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
)

func (s *EtcdServer) createMergedSnapshotMessage(m raftpb.Message, snapt, snapi uint64, confState raftpb.ConfState) snap.Message {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clone := s.store.Clone()
	d, err := clone.SaveNoCopy()
	if err != nil {
		plog.Panicf("store save should never fail: %v", err)
	}
	s.KV().Commit()
	dbsnap := s.be.Snapshot()
	rc := newSnapshotReaderCloser(dbsnap)
	snapshot := raftpb.Snapshot{Metadata: raftpb.SnapshotMetadata{Index: snapi, Term: snapt, ConfState: confState}, Data: d}
	m.Snapshot = snapshot
	return *snap.NewMessage(m, rc, dbsnap.Size())
}
func newSnapshotReaderCloser(snapshot backend.Snapshot) io.ReadCloser {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr, pw := io.Pipe()
	go func() {
		n, err := snapshot.WriteTo(pw)
		if err == nil {
			plog.Infof("wrote database snapshot out [total bytes: %d]", n)
		} else {
			plog.Warningf("failed to write database snapshot out [written bytes: %d]: %v", n, err)
		}
		pw.CloseWithError(err)
		err = snapshot.Close()
		if err != nil {
			plog.Panicf("failed to close database snapshot: %v", err)
		}
	}()
	return pr
}
