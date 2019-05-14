package etcdserver

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/mvcc"
	"github.com/coreos/etcd/pkg/types"
	"time"
)

func (s *EtcdServer) CheckInitialHashKV() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !s.Cfg.InitialCorruptCheck {
		return nil
	}
	plog.Infof("%s starting initial corruption check with timeout %v...", s.ID(), s.Cfg.ReqTimeout())
	h, rev, crev, err := s.kv.HashByRev(0)
	if err != nil {
		return fmt.Errorf("%s failed to fetch hash (%v)", s.ID(), err)
	}
	peers := s.getPeerHashKVs(rev)
	mismatch := 0
	for _, p := range peers {
		if p.resp != nil {
			peerID := types.ID(p.resp.Header.MemberId)
			if h != p.resp.Hash {
				if crev == p.resp.CompactRevision {
					plog.Errorf("%s's hash %d != %s's hash %d (revision %d, peer revision %d, compact revision %d)", s.ID(), h, peerID, p.resp.Hash, rev, p.resp.Header.Revision, crev)
					mismatch++
				} else {
					plog.Warningf("%s cannot check hash of peer(%s): peer has a different compact revision %d (revision:%d)", s.ID(), peerID, p.resp.CompactRevision, rev)
				}
			}
			continue
		}
		if p.err != nil {
			switch p.err {
			case rpctypes.ErrFutureRev:
				plog.Warningf("%s cannot check the hash of peer(%q) at revision %d: peer is lagging behind(%q)", s.ID(), p.eps, rev, p.err.Error())
			case rpctypes.ErrCompacted:
				plog.Warningf("%s cannot check the hash of peer(%q) at revision %d: local node is lagging behind(%q)", s.ID(), p.eps, rev, p.err.Error())
			}
		}
	}
	if mismatch > 0 {
		return fmt.Errorf("%s found data inconsistency with peers", s.ID())
	}
	plog.Infof("%s succeeded on initial corruption checking: no corruption", s.ID())
	return nil
}
func (s *EtcdServer) monitorKVHash() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := s.Cfg.CorruptCheckTime
	if t == 0 {
		return
	}
	plog.Infof("enabled corruption checking with %s interval", t)
	for {
		select {
		case <-s.stopping:
			return
		case <-time.After(t):
		}
		if !s.isLeader() {
			continue
		}
		if err := s.checkHashKV(); err != nil {
			plog.Debugf("check hash kv failed %v", err)
		}
	}
}
func (s *EtcdServer) checkHashKV() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h, rev, crev, err := s.kv.HashByRev(0)
	if err != nil {
		plog.Fatalf("failed to hash kv store (%v)", err)
	}
	peers := s.getPeerHashKVs(rev)
	ctx, cancel := context.WithTimeout(context.Background(), s.Cfg.ReqTimeout())
	err = s.linearizableReadNotify(ctx)
	cancel()
	if err != nil {
		return err
	}
	h2, rev2, crev2, err := s.kv.HashByRev(0)
	if err != nil {
		plog.Warningf("failed to hash kv store (%v)", err)
		return err
	}
	alarmed := false
	mismatch := func(id uint64) {
		if alarmed {
			return
		}
		alarmed = true
		a := &pb.AlarmRequest{MemberID: uint64(id), Action: pb.AlarmRequest_ACTIVATE, Alarm: pb.AlarmType_CORRUPT}
		s.goAttach(func() {
			s.raftRequest(s.ctx, pb.InternalRaftRequest{Alarm: a})
		})
	}
	if h2 != h && rev2 == rev && crev == crev2 {
		plog.Warningf("mismatched hashes %d and %d for revision %d", h, h2, rev)
		mismatch(uint64(s.ID()))
	}
	for _, p := range peers {
		if p.resp == nil {
			continue
		}
		id := p.resp.Header.MemberId
		if p.resp.Header.Revision > rev2 {
			plog.Warningf("revision %d from member %v, expected at most %d", p.resp.Header.Revision, types.ID(id), rev2)
			mismatch(id)
		}
		if p.resp.CompactRevision > crev2 {
			plog.Warningf("compact revision %d from member %v, expected at most %d", p.resp.CompactRevision, types.ID(id), crev2)
			mismatch(id)
		}
		if p.resp.CompactRevision == crev && p.resp.Hash != h {
			plog.Warningf("hash %d at revision %d from member %v, expected hash %d", p.resp.Hash, rev, types.ID(id), h)
			mismatch(id)
		}
	}
	return nil
}

type peerHashKVResp struct {
	resp *clientv3.HashKVResponse
	err  error
	eps  []string
}

func (s *EtcdServer) getPeerHashKVs(rev int64) (resps []*peerHashKVResp) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mbs := s.cluster.Members()
	pURLs := make([][]string, len(mbs))
	for _, m := range mbs {
		if m.ID == s.ID() {
			continue
		}
		pURLs = append(pURLs, m.PeerURLs)
	}
	for _, purls := range pURLs {
		if len(purls) == 0 {
			continue
		}
		cli, cerr := clientv3.New(clientv3.Config{DialTimeout: s.Cfg.ReqTimeout(), Endpoints: purls})
		if cerr != nil {
			plog.Warningf("%s failed to create client to peer %q for hash checking (%q)", s.ID(), purls, cerr.Error())
			continue
		}
		respsLen := len(resps)
		for _, c := range cli.Endpoints() {
			ctx, cancel := context.WithTimeout(context.Background(), s.Cfg.ReqTimeout())
			var resp *clientv3.HashKVResponse
			resp, cerr = cli.HashKV(ctx, c, rev)
			cancel()
			if cerr == nil {
				resps = append(resps, &peerHashKVResp{resp: resp})
				break
			}
			plog.Warningf("%s hash-kv error %q on peer %q with revision %d", s.ID(), cerr.Error(), c, rev)
		}
		cli.Close()
		if respsLen == len(resps) {
			resps = append(resps, &peerHashKVResp{err: cerr, eps: purls})
		}
	}
	return resps
}

type applierV3Corrupt struct{ applierV3 }

func newApplierV3Corrupt(a applierV3) *applierV3Corrupt {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &applierV3Corrupt{a}
}
func (a *applierV3Corrupt) Put(txn mvcc.TxnWrite, p *pb.PutRequest) (*pb.PutResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, ErrCorrupt
}
func (a *applierV3Corrupt) Range(txn mvcc.TxnRead, p *pb.RangeRequest) (*pb.RangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, ErrCorrupt
}
func (a *applierV3Corrupt) DeleteRange(txn mvcc.TxnWrite, p *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, ErrCorrupt
}
func (a *applierV3Corrupt) Txn(rt *pb.TxnRequest) (*pb.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, ErrCorrupt
}
func (a *applierV3Corrupt) Compaction(compaction *pb.CompactionRequest) (*pb.CompactionResponse, <-chan struct{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil, ErrCorrupt
}
func (a *applierV3Corrupt) LeaseGrant(lc *pb.LeaseGrantRequest) (*pb.LeaseGrantResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, ErrCorrupt
}
func (a *applierV3Corrupt) LeaseRevoke(lc *pb.LeaseRevokeRequest) (*pb.LeaseRevokeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, ErrCorrupt
}
