package v3rpc

import (
	"context"
	"sync"
	"time"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	maxNoLeaderCnt = 3
)

type streamsMap struct {
	mu	sync.Mutex
	streams	map[grpc.ServerStream]struct{}
}

func newUnaryInterceptor(s *etcdserver.EtcdServer) grpc.UnaryServerInterceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !api.IsCapabilityEnabled(api.V3rpcCapability) {
			return nil, rpctypes.ErrGRPCNotCapable
		}
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			if ks := md[rpctypes.MetadataRequireLeaderKey]; len(ks) > 0 && ks[0] == rpctypes.MetadataHasLeader {
				if s.Leader() == types.ID(raft.None) {
					return nil, rpctypes.ErrGRPCNoLeader
				}
			}
		}
		return handler(ctx, req)
	}
}
func newLogUnaryInterceptor(s *etcdserver.EtcdServer) grpc.UnaryServerInterceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()
		resp, err := handler(ctx, req)
		defer logUnaryRequestStats(ctx, nil, info, startTime, req, resp)
		return resp, err
	}
}
func logUnaryRequestStats(ctx context.Context, lg *zap.Logger, info *grpc.UnaryServerInfo, startTime time.Time, req interface{}, resp interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	duration := time.Since(startTime)
	remote := "No remote client info."
	peerInfo, ok := peer.FromContext(ctx)
	if ok {
		remote = peerInfo.Addr.String()
	}
	var responseType string = info.FullMethod
	var reqCount, respCount int64
	var reqSize, respSize int
	var reqContent string
	switch _resp := resp.(type) {
	case *pb.RangeResponse:
		_req, ok := req.(*pb.RangeRequest)
		if ok {
			reqCount = 0
			reqSize = _req.Size()
			reqContent = _req.String()
		}
		if _resp != nil {
			respCount = _resp.GetCount()
			respSize = _resp.Size()
		}
	case *pb.PutResponse:
		_req, ok := req.(*pb.PutRequest)
		if ok {
			reqCount = 1
			reqSize = _req.Size()
			reqContent = pb.NewLoggablePutRequest(_req).String()
		}
		if _resp != nil {
			respCount = 0
			respSize = _resp.Size()
		}
	case *pb.DeleteRangeResponse:
		_req, ok := req.(*pb.DeleteRangeRequest)
		if ok {
			reqCount = 0
			reqSize = _req.Size()
			reqContent = _req.String()
		}
		if _resp != nil {
			respCount = _resp.GetDeleted()
			respSize = _resp.Size()
		}
	case *pb.TxnResponse:
		_req, ok := req.(*pb.TxnRequest)
		if ok && _resp != nil {
			if _resp.GetSucceeded() {
				reqCount = int64(len(_req.GetSuccess()))
				reqSize = 0
				for _, r := range _req.GetSuccess() {
					reqSize += r.Size()
				}
			} else {
				reqCount = int64(len(_req.GetFailure()))
				reqSize = 0
				for _, r := range _req.GetFailure() {
					reqSize += r.Size()
				}
			}
			reqContent = pb.NewLoggableTxnRequest(_req).String()
		}
		if _resp != nil {
			respCount = 0
			respSize = _resp.Size()
		}
	default:
		reqCount = -1
		reqSize = -1
		respCount = -1
		respSize = -1
	}
	logGenericRequestStats(lg, startTime, duration, remote, responseType, reqCount, reqSize, respCount, respSize, reqContent)
}
func logGenericRequestStats(lg *zap.Logger, startTime time.Time, duration time.Duration, remote string, responseType string, reqCount int64, reqSize int, respCount int64, respSize int, reqContent string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if lg == nil {
		plog.Debugf("start time = %v, "+"time spent = %v, "+"remote = %s, "+"response type = %s, "+"request count = %d, "+"request size = %d, "+"response count = %d, "+"response size = %d, "+"request content = %s", startTime, duration, remote, responseType, reqCount, reqSize, respCount, respSize, reqContent)
	} else {
		lg.Debug("request stats", zap.Time("start time", startTime), zap.Duration("time spent", duration), zap.String("remote", remote), zap.String("response type", responseType), zap.Int64("request count", reqCount), zap.Int("request size", reqSize), zap.Int64("response count", respCount), zap.Int("response size", respSize), zap.String("request content", reqContent))
	}
}
func newStreamInterceptor(s *etcdserver.EtcdServer) grpc.StreamServerInterceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	smap := monitorLeader(s)
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if !api.IsCapabilityEnabled(api.V3rpcCapability) {
			return rpctypes.ErrGRPCNotCapable
		}
		md, ok := metadata.FromIncomingContext(ss.Context())
		if ok {
			if ks := md[rpctypes.MetadataRequireLeaderKey]; len(ks) > 0 && ks[0] == rpctypes.MetadataHasLeader {
				if s.Leader() == types.ID(raft.None) {
					return rpctypes.ErrGRPCNoLeader
				}
				cctx, cancel := context.WithCancel(ss.Context())
				ss = serverStreamWithCtx{ctx: cctx, cancel: &cancel, ServerStream: ss}
				smap.mu.Lock()
				smap.streams[ss] = struct{}{}
				smap.mu.Unlock()
				defer func() {
					smap.mu.Lock()
					delete(smap.streams, ss)
					smap.mu.Unlock()
					cancel()
				}()
			}
		}
		return handler(srv, ss)
	}
}

type serverStreamWithCtx struct {
	grpc.ServerStream
	ctx	context.Context
	cancel	*context.CancelFunc
}

func (ssc serverStreamWithCtx) Context() context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ssc.ctx
}
func monitorLeader(s *etcdserver.EtcdServer) *streamsMap {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	smap := &streamsMap{streams: make(map[grpc.ServerStream]struct{})}
	go func() {
		election := time.Duration(s.Cfg.TickMs) * time.Duration(s.Cfg.ElectionTicks) * time.Millisecond
		noLeaderCnt := 0
		for {
			select {
			case <-s.StopNotify():
				return
			case <-time.After(election):
				if s.Leader() == types.ID(raft.None) {
					noLeaderCnt++
				} else {
					noLeaderCnt = 0
				}
				if noLeaderCnt >= maxNoLeaderCnt {
					smap.mu.Lock()
					for ss := range smap.streams {
						if ssWithCtx, ok := ss.(serverStreamWithCtx); ok {
							(*ssWithCtx.cancel)()
							<-ss.Context().Done()
						}
					}
					smap.streams = make(map[grpc.ServerStream]struct{})
					smap.mu.Unlock()
				}
			}
		}
	}()
	return smap
}
