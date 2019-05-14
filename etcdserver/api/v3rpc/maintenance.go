package v3rpc

import (
	"context"
	"crypto/sha256"
	"github.com/coreos/etcd/auth"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/mvcc"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/version"
	"io"
)

type KVGetter interface {
	KV() mvcc.ConsistentWatchableKV
}
type BackendGetter interface{ Backend() backend.Backend }
type Alarmer interface {
	Alarm(ctx context.Context, ar *pb.AlarmRequest) (*pb.AlarmResponse, error)
}
type LeaderTransferrer interface {
	MoveLeader(ctx context.Context, lead, target uint64) error
}
type RaftStatusGetter interface {
	etcdserver.RaftTimer
	ID() types.ID
	Leader() types.ID
}
type AuthGetter interface {
	AuthInfoFromCtx(ctx context.Context) (*auth.AuthInfo, error)
	AuthStore() auth.AuthStore
}
type maintenanceServer struct {
	rg  RaftStatusGetter
	kg  KVGetter
	bg  BackendGetter
	a   Alarmer
	lt  LeaderTransferrer
	hdr header
}

func NewMaintenanceServer(s *etcdserver.EtcdServer) pb.MaintenanceServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	srv := &maintenanceServer{rg: s, kg: s, bg: s, a: s, lt: s, hdr: newHeader(s)}
	return &authMaintenanceServer{srv, s}
}
func (ms *maintenanceServer) Defragment(ctx context.Context, sr *pb.DefragmentRequest) (*pb.DefragmentResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plog.Noticef("starting to defragment the storage backend...")
	err := ms.bg.Backend().Defrag()
	if err != nil {
		plog.Errorf("failed to defragment the storage backend (%v)", err)
		return nil, err
	}
	plog.Noticef("finished defragmenting the storage backend")
	return &pb.DefragmentResponse{}, nil
}
func (ms *maintenanceServer) Snapshot(sr *pb.SnapshotRequest, srv pb.Maintenance_SnapshotServer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	snap := ms.bg.Backend().Snapshot()
	pr, pw := io.Pipe()
	defer pr.Close()
	go func() {
		snap.WriteTo(pw)
		if err := snap.Close(); err != nil {
			plog.Errorf("error closing snapshot (%v)", err)
		}
		pw.Close()
	}()
	h := sha256.New()
	br := int64(0)
	buf := make([]byte, 32*1024)
	sz := snap.Size()
	for br < sz {
		n, err := io.ReadFull(pr, buf)
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return togRPCError(err)
		}
		br += int64(n)
		resp := &pb.SnapshotResponse{RemainingBytes: uint64(sz - br), Blob: buf[:n]}
		if err = srv.Send(resp); err != nil {
			return togRPCError(err)
		}
		h.Write(buf[:n])
	}
	sha := h.Sum(nil)
	hresp := &pb.SnapshotResponse{RemainingBytes: 0, Blob: sha}
	if err := srv.Send(hresp); err != nil {
		return togRPCError(err)
	}
	return nil
}
func (ms *maintenanceServer) Hash(ctx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h, rev, err := ms.kg.KV().Hash()
	if err != nil {
		return nil, togRPCError(err)
	}
	resp := &pb.HashResponse{Header: &pb.ResponseHeader{Revision: rev}, Hash: h}
	ms.hdr.fill(resp.Header)
	return resp, nil
}
func (ms *maintenanceServer) HashKV(ctx context.Context, r *pb.HashKVRequest) (*pb.HashKVResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h, rev, compactRev, err := ms.kg.KV().HashByRev(r.Revision)
	if err != nil {
		return nil, togRPCError(err)
	}
	resp := &pb.HashKVResponse{Header: &pb.ResponseHeader{Revision: rev}, Hash: h, CompactRevision: compactRev}
	ms.hdr.fill(resp.Header)
	return resp, nil
}
func (ms *maintenanceServer) Alarm(ctx context.Context, ar *pb.AlarmRequest) (*pb.AlarmResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ms.a.Alarm(ctx, ar)
}
func (ms *maintenanceServer) Status(ctx context.Context, ar *pb.StatusRequest) (*pb.StatusResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp := &pb.StatusResponse{Header: &pb.ResponseHeader{Revision: ms.hdr.rev()}, Version: version.Version, DbSize: ms.bg.Backend().Size(), Leader: uint64(ms.rg.Leader()), RaftIndex: ms.rg.Index(), RaftTerm: ms.rg.Term()}
	ms.hdr.fill(resp.Header)
	return resp, nil
}
func (ms *maintenanceServer) MoveLeader(ctx context.Context, tr *pb.MoveLeaderRequest) (*pb.MoveLeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ms.rg.ID() != ms.rg.Leader() {
		return nil, rpctypes.ErrGRPCNotLeader
	}
	if err := ms.lt.MoveLeader(ctx, uint64(ms.rg.Leader()), tr.TargetID); err != nil {
		return nil, togRPCError(err)
	}
	return &pb.MoveLeaderResponse{}, nil
}

type authMaintenanceServer struct {
	*maintenanceServer
	ag AuthGetter
}

func (ams *authMaintenanceServer) isAuthenticated(ctx context.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	authInfo, err := ams.ag.AuthInfoFromCtx(ctx)
	if err != nil {
		return err
	}
	return ams.ag.AuthStore().IsAdminPermitted(authInfo)
}
func (ams *authMaintenanceServer) Defragment(ctx context.Context, sr *pb.DefragmentRequest) (*pb.DefragmentResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ams.isAuthenticated(ctx); err != nil {
		return nil, err
	}
	return ams.maintenanceServer.Defragment(ctx, sr)
}
func (ams *authMaintenanceServer) Snapshot(sr *pb.SnapshotRequest, srv pb.Maintenance_SnapshotServer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ams.isAuthenticated(srv.Context()); err != nil {
		return err
	}
	return ams.maintenanceServer.Snapshot(sr, srv)
}
func (ams *authMaintenanceServer) Hash(ctx context.Context, r *pb.HashRequest) (*pb.HashResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ams.isAuthenticated(ctx); err != nil {
		return nil, err
	}
	return ams.maintenanceServer.Hash(ctx, r)
}
func (ams *authMaintenanceServer) HashKV(ctx context.Context, r *pb.HashKVRequest) (*pb.HashKVResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ams.isAuthenticated(ctx); err != nil {
		return nil, err
	}
	return ams.maintenanceServer.HashKV(ctx, r)
}
func (ams *authMaintenanceServer) Status(ctx context.Context, ar *pb.StatusRequest) (*pb.StatusResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ams.maintenanceServer.Status(ctx, ar)
}
func (ams *authMaintenanceServer) MoveLeader(ctx context.Context, tr *pb.MoveLeaderRequest) (*pb.MoveLeaderResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ams.maintenanceServer.MoveLeader(ctx, tr)
}
