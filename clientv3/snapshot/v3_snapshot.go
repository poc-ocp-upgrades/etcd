package snapshot

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api/membership"
	"go.etcd.io/etcd/etcdserver/api/snap"
	"go.etcd.io/etcd/etcdserver/api/v2store"
	"go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/lease"
	"go.etcd.io/etcd/mvcc"
	"go.etcd.io/etcd/mvcc/backend"
	"go.etcd.io/etcd/pkg/fileutil"
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/raft"
	"go.etcd.io/etcd/raft/raftpb"
	"go.etcd.io/etcd/wal"
	"go.etcd.io/etcd/wal/walpb"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
)

type Manager interface {
	Save(ctx context.Context, cfg clientv3.Config, dbPath string) error
	Status(dbPath string) (Status, error)
	Restore(cfg RestoreConfig) error
}

func NewV3(lg *zap.Logger) Manager {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if lg == nil {
		lg = zap.NewExample()
	}
	return &v3Manager{lg: lg}
}

type v3Manager struct {
	lg		*zap.Logger
	name		string
	dbPath		string
	walDir		string
	snapDir		string
	cl		*membership.RaftCluster
	skipHashCheck	bool
}

func (s *v3Manager) Save(ctx context.Context, cfg clientv3.Config, dbPath string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(cfg.Endpoints) != 1 {
		return fmt.Errorf("snapshot must be requested to one selected node, not multiple %v", cfg.Endpoints)
	}
	cli, err := clientv3.New(cfg)
	if err != nil {
		return err
	}
	defer cli.Close()
	partpath := dbPath + ".part"
	defer os.RemoveAll(partpath)
	var f *os.File
	f, err = os.OpenFile(partpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileutil.PrivateFileMode)
	if err != nil {
		return fmt.Errorf("could not open %s (%v)", partpath, err)
	}
	s.lg.Info("created temporary db file", zap.String("path", partpath))
	now := time.Now()
	var rd io.ReadCloser
	rd, err = cli.Snapshot(ctx)
	if err != nil {
		return err
	}
	s.lg.Info("fetching snapshot", zap.String("endpoint", cfg.Endpoints[0]))
	if _, err = io.Copy(f, rd); err != nil {
		return err
	}
	if err = fileutil.Fsync(f); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	s.lg.Info("fetched snapshot", zap.String("endpoint", cfg.Endpoints[0]), zap.Duration("took", time.Since(now)))
	if err = os.Rename(partpath, dbPath); err != nil {
		return fmt.Errorf("could not rename %s to %s (%v)", partpath, dbPath, err)
	}
	s.lg.Info("saved", zap.String("path", dbPath))
	return nil
}

type Status struct {
	Hash		uint32	`json:"hash"`
	Revision	int64	`json:"revision"`
	TotalKey	int	`json:"totalKey"`
	TotalSize	int64	`json:"totalSize"`
}

func (s *v3Manager) Status(dbPath string) (ds Status, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, err = os.Stat(dbPath); err != nil {
		return ds, err
	}
	db, err := bolt.Open(dbPath, 0400, &bolt.Options{ReadOnly: true})
	if err != nil {
		return ds, err
	}
	defer db.Close()
	h := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	if err = db.View(func(tx *bolt.Tx) error {
		var dbErrStrings []string
		for dbErr := range tx.Check() {
			dbErrStrings = append(dbErrStrings, dbErr.Error())
		}
		if len(dbErrStrings) > 0 {
			return fmt.Errorf("snapshot file integrity check failed. %d errors found.\n"+strings.Join(dbErrStrings, "\n"), len(dbErrStrings))
		}
		ds.TotalSize = tx.Size()
		c := tx.Cursor()
		for next, _ := c.First(); next != nil; next, _ = c.Next() {
			b := tx.Bucket(next)
			if b == nil {
				return fmt.Errorf("cannot get hash of bucket %s", string(next))
			}
			h.Write(next)
			iskeyb := (string(next) == "key")
			b.ForEach(func(k, v []byte) error {
				h.Write(k)
				h.Write(v)
				if iskeyb {
					rev := bytesToRev(k)
					ds.Revision = rev.main
				}
				ds.TotalKey++
				return nil
			})
		}
		return nil
	}); err != nil {
		return ds, err
	}
	ds.Hash = h.Sum32()
	return ds, nil
}

type RestoreConfig struct {
	SnapshotPath		string
	Name			string
	OutputDataDir		string
	OutputWALDir		string
	PeerURLs		[]string
	InitialCluster		string
	InitialClusterToken	string
	SkipHashCheck		bool
}

func (s *v3Manager) Restore(cfg RestoreConfig) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pURLs, err := types.NewURLs(cfg.PeerURLs)
	if err != nil {
		return err
	}
	var ics types.URLsMap
	ics, err = types.NewURLsMap(cfg.InitialCluster)
	if err != nil {
		return err
	}
	srv := etcdserver.ServerConfig{Logger: s.lg, Name: cfg.Name, PeerURLs: pURLs, InitialPeerURLsMap: ics, InitialClusterToken: cfg.InitialClusterToken}
	if err = srv.VerifyBootstrap(); err != nil {
		return err
	}
	s.cl, err = membership.NewClusterFromURLsMap(s.lg, cfg.InitialClusterToken, ics)
	if err != nil {
		return err
	}
	dataDir := cfg.OutputDataDir
	if dataDir == "" {
		dataDir = cfg.Name + ".etcd"
	}
	if fileutil.Exist(dataDir) {
		return fmt.Errorf("data-dir %q exists", dataDir)
	}
	walDir := cfg.OutputWALDir
	if walDir == "" {
		walDir = filepath.Join(dataDir, "member", "wal")
	} else if fileutil.Exist(walDir) {
		return fmt.Errorf("wal-dir %q exists", walDir)
	}
	s.name = cfg.Name
	s.dbPath = cfg.SnapshotPath
	s.walDir = walDir
	s.snapDir = filepath.Join(dataDir, "member", "snap")
	s.skipHashCheck = cfg.SkipHashCheck
	s.lg.Info("restoring snapshot", zap.String("path", s.dbPath), zap.String("wal-dir", s.walDir), zap.String("data-dir", dataDir), zap.String("snap-dir", s.snapDir))
	if err = s.saveDB(); err != nil {
		return err
	}
	if err = s.saveWALAndSnap(); err != nil {
		return err
	}
	s.lg.Info("restored snapshot", zap.String("path", s.dbPath), zap.String("wal-dir", s.walDir), zap.String("data-dir", dataDir), zap.String("snap-dir", s.snapDir))
	return nil
}
func (s *v3Manager) saveDB() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f, ferr := os.OpenFile(s.dbPath, os.O_RDONLY, 0600)
	if ferr != nil {
		return ferr
	}
	defer f.Close()
	if _, err := f.Seek(-sha256.Size, io.SeekEnd); err != nil {
		return err
	}
	sha := make([]byte, sha256.Size)
	if _, err := f.Read(sha); err != nil {
		return err
	}
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return err
	}
	if err := fileutil.CreateDirAll(s.snapDir); err != nil {
		return err
	}
	dbpath := filepath.Join(s.snapDir, "db")
	db, dberr := os.OpenFile(dbpath, os.O_RDWR|os.O_CREATE, 0600)
	if dberr != nil {
		return dberr
	}
	if _, err := io.Copy(db, f); err != nil {
		return err
	}
	off, serr := db.Seek(0, io.SeekEnd)
	if serr != nil {
		return serr
	}
	hasHash := (off % 512) == sha256.Size
	if hasHash {
		if err := db.Truncate(off - sha256.Size); err != nil {
			return err
		}
	}
	if !hasHash && !s.skipHashCheck {
		return fmt.Errorf("snapshot missing hash but --skip-hash-check=false")
	}
	if hasHash && !s.skipHashCheck {
		if _, err := db.Seek(0, io.SeekStart); err != nil {
			return err
		}
		h := sha256.New()
		if _, err := io.Copy(h, db); err != nil {
			return err
		}
		dbsha := h.Sum(nil)
		if !reflect.DeepEqual(sha, dbsha) {
			return fmt.Errorf("expected sha256 %v, got %v", sha, dbsha)
		}
	}
	db.Close()
	commit := len(s.cl.Members())
	be := backend.NewDefaultBackend(dbpath)
	lessor := lease.NewLessor(s.lg, be, lease.LessorConfig{MinLeaseTTL: math.MaxInt64})
	mvs := mvcc.NewStore(s.lg, be, lessor, (*initIndex)(&commit))
	txn := mvs.Write()
	btx := be.BatchTx()
	del := func(k, v []byte) error {
		txn.DeleteRange(k, nil)
		return nil
	}
	btx.UnsafeForEach([]byte("members"), del)
	btx.UnsafeForEach([]byte("members_removed"), del)
	txn.End()
	mvs.Commit()
	mvs.Close()
	be.Close()
	return nil
}
func (s *v3Manager) saveWALAndSnap() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := fileutil.CreateDirAll(s.walDir); err != nil {
		return err
	}
	st := v2store.New(etcdserver.StoreClusterPrefix, etcdserver.StoreKeysPrefix)
	s.cl.SetStore(st)
	for _, m := range s.cl.Members() {
		s.cl.AddMember(m)
	}
	m := s.cl.MemberByName(s.name)
	md := &etcdserverpb.Metadata{NodeID: uint64(m.ID), ClusterID: uint64(s.cl.ID())}
	metadata, merr := md.Marshal()
	if merr != nil {
		return merr
	}
	w, walerr := wal.Create(s.lg, s.walDir, metadata)
	if walerr != nil {
		return walerr
	}
	defer w.Close()
	peers := make([]raft.Peer, len(s.cl.MemberIDs()))
	for i, id := range s.cl.MemberIDs() {
		ctx, err := json.Marshal((*s.cl).Member(id))
		if err != nil {
			return err
		}
		peers[i] = raft.Peer{ID: uint64(id), Context: ctx}
	}
	ents := make([]raftpb.Entry, len(peers))
	nodeIDs := make([]uint64, len(peers))
	for i, p := range peers {
		nodeIDs[i] = p.ID
		cc := raftpb.ConfChange{Type: raftpb.ConfChangeAddNode, NodeID: p.ID, Context: p.Context}
		d, err := cc.Marshal()
		if err != nil {
			return err
		}
		ents[i] = raftpb.Entry{Type: raftpb.EntryConfChange, Term: 1, Index: uint64(i + 1), Data: d}
	}
	commit, term := uint64(len(ents)), uint64(1)
	if err := w.Save(raftpb.HardState{Term: term, Vote: peers[0].ID, Commit: commit}, ents); err != nil {
		return err
	}
	b, berr := st.Save()
	if berr != nil {
		return berr
	}
	raftSnap := raftpb.Snapshot{Data: b, Metadata: raftpb.SnapshotMetadata{Index: commit, Term: term, ConfState: raftpb.ConfState{Nodes: nodeIDs}}}
	sn := snap.New(s.lg, s.snapDir)
	if err := sn.SaveSnap(raftSnap); err != nil {
		return err
	}
	return w.SaveSnapshot(walpb.Snapshot{Index: commit, Term: term})
}
