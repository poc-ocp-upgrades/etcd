package v3rpc

import (
	"context"
	"io"
	"math/rand"
	"sync"
	"time"
	"go.etcd.io/etcd/auth"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/mvcc"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"go.uber.org/zap"
)

type watchServer struct {
	lg		*zap.Logger
	clusterID	int64
	memberID	int64
	maxRequestBytes	int
	sg		etcdserver.RaftStatusGetter
	watchable	mvcc.WatchableKV
	ag		AuthGetter
}

func NewWatchServer(s *etcdserver.EtcdServer) pb.WatchServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &watchServer{lg: s.Cfg.Logger, clusterID: int64(s.Cluster().ID()), memberID: int64(s.ID()), maxRequestBytes: int(s.Cfg.MaxRequestBytes + grpcOverheadBytes), sg: s, watchable: s.Watchable(), ag: s}
}

var (
	progressReportInterval		= 10 * time.Minute
	progressReportIntervalMu	sync.RWMutex
)

func GetProgressReportInterval() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	progressReportIntervalMu.RLock()
	interval := progressReportInterval
	progressReportIntervalMu.RUnlock()
	jitter := time.Duration(rand.Int63n(int64(interval) / 10))
	return interval + jitter
}
func SetProgressReportInterval(newTimeout time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	progressReportIntervalMu.Lock()
	progressReportInterval = newTimeout
	progressReportIntervalMu.Unlock()
}

const ctrlStreamBufLen = 16

type serverWatchStream struct {
	lg		*zap.Logger
	clusterID	int64
	memberID	int64
	maxRequestBytes	int
	sg		etcdserver.RaftStatusGetter
	watchable	mvcc.WatchableKV
	ag		AuthGetter
	gRPCStream	pb.Watch_WatchServer
	watchStream	mvcc.WatchStream
	ctrlStream	chan *pb.WatchResponse
	mu		sync.RWMutex
	progress	map[mvcc.WatchID]bool
	prevKV		map[mvcc.WatchID]bool
	fragment	map[mvcc.WatchID]bool
	closec		chan struct{}
	wg		sync.WaitGroup
}

func (ws *watchServer) Watch(stream pb.Watch_WatchServer) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sws := serverWatchStream{lg: ws.lg, clusterID: ws.clusterID, memberID: ws.memberID, maxRequestBytes: ws.maxRequestBytes, sg: ws.sg, watchable: ws.watchable, ag: ws.ag, gRPCStream: stream, watchStream: ws.watchable.NewWatchStream(), ctrlStream: make(chan *pb.WatchResponse, ctrlStreamBufLen), progress: make(map[mvcc.WatchID]bool), prevKV: make(map[mvcc.WatchID]bool), fragment: make(map[mvcc.WatchID]bool), closec: make(chan struct{})}
	sws.wg.Add(1)
	go func() {
		sws.sendLoop()
		sws.wg.Done()
	}()
	errc := make(chan error, 1)
	go func() {
		if rerr := sws.recvLoop(); rerr != nil {
			if isClientCtxErr(stream.Context().Err(), rerr) {
				if sws.lg != nil {
					sws.lg.Debug("failed to receive watch request from gRPC stream", zap.Error(rerr))
				} else {
					plog.Debugf("failed to receive watch request from gRPC stream (%q)", rerr.Error())
				}
			} else {
				if sws.lg != nil {
					sws.lg.Warn("failed to receive watch request from gRPC stream", zap.Error(err))
				} else {
					plog.Warningf("failed to receive watch request from gRPC stream (%q)", rerr.Error())
				}
				streamFailures.WithLabelValues("receive", "watch").Inc()
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
				wr := &pb.WatchResponse{Header: sws.newResponseHeader(sws.watchStream.Rev()), WatchId: creq.WatchId, Canceled: true, Created: true, CancelReason: rpctypes.ErrGRPCPermissionDenied.Error()}
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
			id, err := sws.watchStream.Watch(mvcc.WatchID(creq.WatchId), creq.Key, creq.RangeEnd, rev, filters...)
			if err == nil {
				sws.mu.Lock()
				if creq.ProgressNotify {
					sws.progress[id] = true
				}
				if creq.PrevKv {
					sws.prevKV[id] = true
				}
				if creq.Fragment {
					sws.fragment[id] = true
				}
				sws.mu.Unlock()
			}
			wr := &pb.WatchResponse{Header: sws.newResponseHeader(wsrev), WatchId: int64(id), Created: true, Canceled: err != nil}
			if err != nil {
				wr.CancelReason = err.Error()
			}
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
					delete(sws.fragment, mvcc.WatchID(id))
					sws.mu.Unlock()
				}
			}
		case *pb.WatchRequest_ProgressRequest:
			if uv.ProgressRequest != nil {
				sws.ctrlStream <- &pb.WatchResponse{Header: sws.newResponseHeader(sws.watchStream.Rev()), WatchId: -1}
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
			sws.mu.RLock()
			needPrevKV := sws.prevKV[wresp.WatchID]
			sws.mu.RUnlock()
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
			if _, okID := ids[wresp.WatchID]; !okID {
				wrs := append(pending[wresp.WatchID], wr)
				pending[wresp.WatchID] = wrs
				continue
			}
			mvcc.ReportEventReceived(len(evs))
			sws.mu.RLock()
			fragmented, ok := sws.fragment[wresp.WatchID]
			sws.mu.RUnlock()
			var serr error
			if !fragmented && !ok {
				serr = sws.gRPCStream.Send(wr)
			} else {
				serr = sendFragments(wr, sws.maxRequestBytes, sws.gRPCStream.Send)
			}
			if serr != nil {
				if isClientCtxErr(sws.gRPCStream.Context().Err(), serr) {
					if sws.lg != nil {
						sws.lg.Debug("failed to send watch response to gRPC stream", zap.Error(serr))
					} else {
						plog.Debugf("failed to send watch response to gRPC stream (%q)", serr.Error())
					}
				} else {
					if sws.lg != nil {
						sws.lg.Warn("failed to send watch response to gRPC stream", zap.Error(serr))
					} else {
						plog.Warningf("failed to send watch response to gRPC stream (%q)", serr.Error())
					}
					streamFailures.WithLabelValues("send", "watch").Inc()
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
					if sws.lg != nil {
						sws.lg.Debug("failed to send watch control response to gRPC stream", zap.Error(err))
					} else {
						plog.Debugf("failed to send watch control response to gRPC stream (%q)", err.Error())
					}
				} else {
					if sws.lg != nil {
						sws.lg.Warn("failed to send watch control response to gRPC stream", zap.Error(err))
					} else {
						plog.Warningf("failed to send watch control response to gRPC stream (%q)", err.Error())
					}
					streamFailures.WithLabelValues("send", "watch").Inc()
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
							if sws.lg != nil {
								sws.lg.Debug("failed to send pending watch response to gRPC stream", zap.Error(err))
							} else {
								plog.Debugf("failed to send pending watch response to gRPC stream (%q)", err.Error())
							}
						} else {
							if sws.lg != nil {
								sws.lg.Warn("failed to send pending watch response to gRPC stream", zap.Error(err))
							} else {
								plog.Warningf("failed to send pending watch response to gRPC stream (%q)", err.Error())
							}
							streamFailures.WithLabelValues("send", "watch").Inc()
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
func sendFragments(wr *pb.WatchResponse, maxRequestBytes int, sendFunc func(*pb.WatchResponse) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if wr.Size() < maxRequestBytes || len(wr.Events) < 2 {
		return sendFunc(wr)
	}
	ow := *wr
	ow.Events = make([]*mvccpb.Event, 0)
	ow.Fragment = true
	var idx int
	for {
		cur := ow
		for _, ev := range wr.Events[idx:] {
			cur.Events = append(cur.Events, ev)
			if len(cur.Events) > 1 && cur.Size() >= maxRequestBytes {
				cur.Events = cur.Events[:len(cur.Events)-1]
				break
			}
			idx++
		}
		if idx == len(wr.Events) {
			cur.Fragment = false
		}
		if err := sendFunc(&cur); err != nil {
			return err
		}
		if !cur.Fragment {
			break
		}
	}
	return nil
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
	return &pb.ResponseHeader{ClusterId: uint64(sws.clusterID), MemberId: uint64(sws.memberID), Revision: rev, RaftTerm: sws.sg.Term()}
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
