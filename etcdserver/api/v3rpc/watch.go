package v3rpc

import (
	"context"
	"github.com/coreos/etcd/auth"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/mvcc"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"io"
	"sync"
	"time"
)

type watchServer struct {
	clusterID int64
	memberID  int64
	raftTimer etcdserver.RaftTimer
	watchable mvcc.WatchableKV
	ag        AuthGetter
}

func NewWatchServer(s *etcdserver.EtcdServer) pb.WatchServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &watchServer{clusterID: int64(s.Cluster().ID()), memberID: int64(s.ID()), raftTimer: s, watchable: s.Watchable(), ag: s}
}

var (
	progressReportInterval   = 10 * time.Minute
	progressReportIntervalMu sync.RWMutex
)

func GetProgressReportInterval() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	progressReportIntervalMu.RLock()
	defer progressReportIntervalMu.RUnlock()
	return progressReportInterval
}
func SetProgressReportInterval(newTimeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	progressReportIntervalMu.Lock()
	defer progressReportIntervalMu.Unlock()
	progressReportInterval = newTimeout
}

const (
	ctrlStreamBufLen = 16
)

type serverWatchStream struct {
	clusterID   int64
	memberID    int64
	raftTimer   etcdserver.RaftTimer
	watchable   mvcc.WatchableKV
	gRPCStream  pb.Watch_WatchServer
	watchStream mvcc.WatchStream
	ctrlStream  chan *pb.WatchResponse
	mu          sync.Mutex
	progress    map[mvcc.WatchID]bool
	prevKV      map[mvcc.WatchID]bool
	closec      chan struct{}
	wg          sync.WaitGroup
	ag          AuthGetter
}

func (ws *watchServer) Watch(stream pb.Watch_WatchServer) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sws := serverWatchStream{clusterID: ws.clusterID, memberID: ws.memberID, raftTimer: ws.raftTimer, watchable: ws.watchable, gRPCStream: stream, watchStream: ws.watchable.NewWatchStream(), ctrlStream: make(chan *pb.WatchResponse, ctrlStreamBufLen), progress: make(map[mvcc.WatchID]bool), prevKV: make(map[mvcc.WatchID]bool), closec: make(chan struct{}), ag: ws.ag}
	sws.wg.Add(1)
	go func() {
		sws.sendLoop()
		sws.wg.Done()
	}()
	errc := make(chan error, 1)
	go func() {
		if rerr := sws.recvLoop(); rerr != nil {
			if isClientCtxErr(stream.Context().Err(), rerr) {
				plog.Debugf("failed to receive watch request from gRPC stream (%q)", rerr.Error())
			} else {
				plog.Warningf("failed to receive watch request from gRPC stream (%q)", rerr.Error())
			}
			errc <- rerr
		}
	}()
	select {
	case err = <-errc:
		close(sws.ctrlStream)
	case <-stream.Context().Done():
		err = stream.Context().Err()
		if err == context.Canceled {
			err = rpctypes.ErrGRPCNoLeader
		}
	}
	sws.close()
	return err
}
func (sws *serverWatchStream) isWatchPermitted(wcr *pb.WatchCreateRequest) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	authInfo, err := sws.ag.AuthInfoFromCtx(sws.gRPCStream.Context())
	if err != nil {
		return false
	}
	if authInfo == nil {
		authInfo = &auth.AuthInfo{}
	}
	return sws.ag.AuthStore().IsRangePermitted(authInfo, wcr.Key, wcr.RangeEnd) == nil
}
func (sws *serverWatchStream) recvLoop() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		req, err := sws.gRPCStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		switch uv := req.RequestUnion.(type) {
		case *pb.WatchRequest_CreateRequest:
			if uv.CreateRequest == nil {
				break
			}
			creq := uv.CreateRequest
			if len(creq.Key) == 0 {
				creq.Key = []byte{0}
			}
			if len(creq.RangeEnd) == 0 {
				creq.RangeEnd = nil
			}
			if len(creq.RangeEnd) == 1 && creq.RangeEnd[0] == 0 {
				creq.RangeEnd = []byte{}
			}
			if !sws.isWatchPermitted(creq) {
				wr := &pb.WatchResponse{Header: sws.newResponseHeader(sws.watchStream.Rev()), WatchId: -1, Canceled: true, Created: true, CancelReason: rpctypes.ErrGRPCPermissionDenied.Error()}
				select {
				case sws.ctrlStream <- wr:
				case <-sws.closec:
				}
				return nil
			}
			filters := FiltersFromRequest(creq)
			wsrev := sws.watchStream.Rev()
			rev := creq.StartRevision
			if rev == 0 {
				rev = wsrev + 1
			}
			id := sws.watchStream.Watch(creq.Key, creq.RangeEnd, rev, filters...)
			if id != -1 {
				sws.mu.Lock()
				if creq.ProgressNotify {
					sws.progress[id] = true
				}
				if creq.PrevKv {
					sws.prevKV[id] = true
				}
				sws.mu.Unlock()
			}
			wr := &pb.WatchResponse{Header: sws.newResponseHeader(wsrev), WatchId: int64(id), Created: true, Canceled: id == -1}
			select {
			case sws.ctrlStream <- wr:
			case <-sws.closec:
				return nil
			}
		case *pb.WatchRequest_CancelRequest:
			if uv.CancelRequest != nil {
				id := uv.CancelRequest.WatchId
				err := sws.watchStream.Cancel(mvcc.WatchID(id))
				if err == nil {
					sws.ctrlStream <- &pb.WatchResponse{Header: sws.newResponseHeader(sws.watchStream.Rev()), WatchId: id, Canceled: true}
					sws.mu.Lock()
					delete(sws.progress, mvcc.WatchID(id))
					delete(sws.prevKV, mvcc.WatchID(id))
					sws.mu.Unlock()
				}
			}
		default:
			continue
		}
	}
}
func (sws *serverWatchStream) sendLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ids := make(map[mvcc.WatchID]struct{})
	pending := make(map[mvcc.WatchID][]*pb.WatchResponse)
	interval := GetProgressReportInterval()
	progressTicker := time.NewTicker(interval)
	defer func() {
		progressTicker.Stop()
		for ws := range sws.watchStream.Chan() {
			mvcc.ReportEventReceived(len(ws.Events))
		}
		for _, wrs := range pending {
			for _, ws := range wrs {
				mvcc.ReportEventReceived(len(ws.Events))
			}
		}
	}()
	for {
		select {
		case wresp, ok := <-sws.watchStream.Chan():
			if !ok {
				return
			}
			evs := wresp.Events
			events := make([]*mvccpb.Event, len(evs))
			sws.mu.Lock()
			needPrevKV := sws.prevKV[wresp.WatchID]
			sws.mu.Unlock()
			for i := range evs {
				events[i] = &evs[i]
				if needPrevKV {
					opt := mvcc.RangeOptions{Rev: evs[i].Kv.ModRevision - 1}
					r, err := sws.watchable.Range(evs[i].Kv.Key, nil, opt)
					if err == nil && len(r.KVs) != 0 {
						events[i].PrevKv = &(r.KVs[0])
					}
				}
			}
			canceled := wresp.CompactRevision != 0
			wr := &pb.WatchResponse{Header: sws.newResponseHeader(wresp.Revision), WatchId: int64(wresp.WatchID), Events: events, CompactRevision: wresp.CompactRevision, Canceled: canceled}
			if _, hasId := ids[wresp.WatchID]; !hasId {
				wrs := append(pending[wresp.WatchID], wr)
				pending[wresp.WatchID] = wrs
				continue
			}
			mvcc.ReportEventReceived(len(evs))
			if err := sws.gRPCStream.Send(wr); err != nil {
				if isClientCtxErr(sws.gRPCStream.Context().Err(), err) {
					plog.Debugf("failed to send watch response to gRPC stream (%q)", err.Error())
				} else {
					plog.Warningf("failed to send watch response to gRPC stream (%q)", err.Error())
				}
				return
			}
			sws.mu.Lock()
			if len(evs) > 0 && sws.progress[wresp.WatchID] {
				sws.progress[wresp.WatchID] = false
			}
			sws.mu.Unlock()
		case c, ok := <-sws.ctrlStream:
			if !ok {
				return
			}
			if err := sws.gRPCStream.Send(c); err != nil {
				if isClientCtxErr(sws.gRPCStream.Context().Err(), err) {
					plog.Debugf("failed to send watch control response to gRPC stream (%q)", err.Error())
				} else {
					plog.Warningf("failed to send watch control response to gRPC stream (%q)", err.Error())
				}
				return
			}
			wid := mvcc.WatchID(c.WatchId)
			if c.Canceled {
				delete(ids, wid)
				continue
			}
			if c.Created {
				ids[wid] = struct{}{}
				for _, v := range pending[wid] {
					mvcc.ReportEventReceived(len(v.Events))
					if err := sws.gRPCStream.Send(v); err != nil {
						if isClientCtxErr(sws.gRPCStream.Context().Err(), err) {
							plog.Debugf("failed to send pending watch response to gRPC stream (%q)", err.Error())
						} else {
							plog.Warningf("failed to send pending watch response to gRPC stream (%q)", err.Error())
						}
						return
					}
				}
				delete(pending, wid)
			}
		case <-progressTicker.C:
			sws.mu.Lock()
			for id, ok := range sws.progress {
				if ok {
					sws.watchStream.RequestProgress(id)
				}
				sws.progress[id] = true
			}
			sws.mu.Unlock()
		case <-sws.closec:
			return
		}
	}
}
func (sws *serverWatchStream) close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sws.watchStream.Close()
	close(sws.closec)
	sws.wg.Wait()
}
func (sws *serverWatchStream) newResponseHeader(rev int64) *pb.ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pb.ResponseHeader{ClusterId: uint64(sws.clusterID), MemberId: uint64(sws.memberID), Revision: rev, RaftTerm: sws.raftTimer.Term()}
}
func filterNoDelete(e mvccpb.Event) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.Type == mvccpb.DELETE
}
func filterNoPut(e mvccpb.Event) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.Type == mvccpb.PUT
}
func FiltersFromRequest(creq *pb.WatchCreateRequest) []mvcc.FilterFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	filters := make([]mvcc.FilterFunc, 0, len(creq.Filters))
	for _, ft := range creq.Filters {
		switch ft {
		case pb.WatchCreateRequest_NOPUT:
			filters = append(filters, filterNoPut)
		case pb.WatchCreateRequest_NODELETE:
			filters = append(filters, filterNoDelete)
		default:
		}
	}
	return filters
}
