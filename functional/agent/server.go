package agent

import (
	"math"
	"net"
	"os"
	"os/exec"
	"strings"
	"github.com/coreos/etcd/functional/rpcpb"
	"github.com/coreos/etcd/pkg/proxy"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer	*grpc.Server
	lg		*zap.Logger
	network		string
	address		string
	ln		net.Listener
	rpcpb.TransportServer
	last	rpcpb.Operation
	*rpcpb.Member
	*rpcpb.Tester
	etcdCmd				*exec.Cmd
	etcdLogFile			*os.File
	advertiseClientPortToProxy	map[int]proxy.Server
	advertisePeerPortToProxy	map[int]proxy.Server
}

func NewServer(lg *zap.Logger, network string, address string) *Server {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Server{lg: lg, network: network, address: address, last: rpcpb.Operation_NOT_STARTED, advertiseClientPortToProxy: make(map[int]proxy.Server), advertisePeerPortToProxy: make(map[int]proxy.Server)}
}

const (
	maxRequestBytes		= 1.5 * 1024 * 1024
	grpcOverheadBytes	= 512 * 1024
	maxStreams		= math.MaxUint32
	maxSendBytes		= math.MaxInt32
)

func (srv *Server) StartServe() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	srv.ln, err = net.Listen(srv.network, srv.address)
	if err != nil {
		return err
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.MaxRecvMsgSize(int(maxRequestBytes+grpcOverheadBytes)))
	opts = append(opts, grpc.MaxSendMsgSize(maxSendBytes))
	opts = append(opts, grpc.MaxConcurrentStreams(maxStreams))
	srv.grpcServer = grpc.NewServer(opts...)
	rpcpb.RegisterTransportServer(srv.grpcServer, srv)
	srv.lg.Info("gRPC server started", zap.String("address", srv.address), zap.String("listener-address", srv.ln.Addr().String()))
	err = srv.grpcServer.Serve(srv.ln)
	if err != nil && strings.Contains(err.Error(), "use of closed network connection") {
		srv.lg.Info("gRPC server is shut down", zap.String("address", srv.address), zap.Error(err))
	} else {
		srv.lg.Warn("gRPC server returned with error", zap.String("address", srv.address), zap.Error(err))
	}
	return err
}
func (srv *Server) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	srv.lg.Info("gRPC server stopping", zap.String("address", srv.address))
	srv.grpcServer.Stop()
	srv.lg.Info("gRPC server stopped", zap.String("address", srv.address))
}
func (srv *Server) Transport(stream rpcpb.Transport_TransportServer) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	errc := make(chan error)
	go func() {
		for {
			var req *rpcpb.Request
			req, err = stream.Recv()
			if err != nil {
				errc <- err
				return
			}
			if req.Member != nil {
				srv.Member = req.Member
			}
			if req.Tester != nil {
				srv.Tester = req.Tester
			}
			var resp *rpcpb.Response
			resp, err = srv.handleTesterRequest(req)
			if err != nil {
				errc <- err
				return
			}
			if err = stream.Send(resp); err != nil {
				errc <- err
				return
			}
		}
	}()
	select {
	case err = <-errc:
	case <-stream.Context().Done():
		err = stream.Context().Err()
	}
	return err
}
