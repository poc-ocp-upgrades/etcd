package etcdserver

import (
	"io"
	"go.etcd.io/etcd/etcdserver/api/snap"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/pkg/pbutil"
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/raft/raftpb"
	"go.etcd.io/etcd/wal"
	"go.etcd.io/etcd/wal/walpb"
	"go.uber.org/zap"
)

type Storage interface {
	Save(st raftpb.HardState, ents []raftpb.Entry) error
	SaveSnap(snap raftpb.Snapshot) error
	Close() error
}
type storage struct {
	*wal.WAL
	*snap.Snapshotter
}

func NewStorage(w *wal.WAL, s *snap.Snapshotter) Storage {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &storage{w, s}
}
func (st *storage) SaveSnap(snap raftpb.Snapshot) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	walsnap := walpb.Snapshot{Index: snap.Metadata.Index, Term: snap.Metadata.Term}
	err := st.WAL.SaveSnapshot(walsnap)
	if err != nil {
		return err
	}
	err = st.Snapshotter.SaveSnap(snap)
	if err != nil {
		return err
	}
	return st.WAL.ReleaseLockTo(snap.Metadata.Index)
}
func readWAL(lg *zap.Logger, waldir string, snap walpb.Snapshot) (w *wal.WAL, id, cid types.ID, st raftpb.HardState, ents []raftpb.Entry) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		err		error
		wmetadata	[]byte
	)
	repaired := false
	for {
		if w, err = wal.Open(lg, waldir, snap); err != nil {
			if lg != nil {
				lg.Fatal("failed to open WAL", zap.Error(err))
			} else {
				plog.Fatalf("open wal error: %v", err)
			}
		}
		if wmetadata, st, ents, err = w.ReadAll(); err != nil {
			w.Close()
			if repaired || err != io.ErrUnexpectedEOF {
				if lg != nil {
					lg.Fatal("failed to read WAL, cannot be repaired", zap.Error(err))
				} else {
					plog.Fatalf("read wal error (%v) and cannot be repaired", err)
				}
			}
			if !wal.Repair(lg, waldir) {
				if lg != nil {
					lg.Fatal("failed to repair WAL", zap.Error(err))
				} else {
					plog.Fatalf("WAL error (%v) cannot be repaired", err)
				}
			} else {
				if lg != nil {
					lg.Info("repaired WAL", zap.Error(err))
				} else {
					plog.Infof("repaired WAL error (%v)", err)
				}
				repaired = true
			}
			continue
		}
		break
	}
	var metadata pb.Metadata
	pbutil.MustUnmarshal(&metadata, wmetadata)
	id = types.ID(metadata.NodeID)
	cid = types.ID(metadata.ClusterID)
	return w, id, cid, st, ents
}
