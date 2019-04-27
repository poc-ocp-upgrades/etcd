package etcdserver

import (
	"io"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/pbutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/snap"
	"github.com/coreos/etcd/wal"
	"github.com/coreos/etcd/wal/walpb"
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
func readWAL(waldir string, snap walpb.Snapshot) (w *wal.WAL, id, cid types.ID, st raftpb.HardState, ents []raftpb.Entry) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		err		error
		wmetadata	[]byte
	)
	repaired := false
	for {
		if w, err = wal.Open(waldir, snap); err != nil {
			plog.Fatalf("open wal error: %v", err)
		}
		if wmetadata, st, ents, err = w.ReadAll(); err != nil {
			w.Close()
			if repaired || err != io.ErrUnexpectedEOF {
				plog.Fatalf("read wal error (%v) and cannot be repaired", err)
			}
			if !wal.Repair(waldir) {
				plog.Fatalf("WAL error (%v) cannot be repaired", err)
			} else {
				plog.Infof("repaired WAL error (%v)", err)
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
