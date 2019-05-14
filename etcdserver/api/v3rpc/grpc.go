package v3rpc

import (
	"crypto/tls"
	"github.com/coreos/etcd/etcdserver"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"math"
)

const (
	grpcOverheadBytes = 512 * 1024
	maxStreams        = math.MaxUint32
	maxSendBytes      = math.MaxInt32
)

func Server(s *etcdserver.EtcdServer, tls *tls.Config, gopts ...grpc.ServerOption) *grpc.Server {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var opts []grpc.ServerOption
	opts = append(opts, grpc.CustomCodec(&codec{}))
	if tls != nil {
		opts = append(opts, grpc.Creds(credentials.NewTLS(tls)))
	}
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(newLogUnaryInterceptor(s), newUnaryInterceptor(s), grpc_prometheus.UnaryServerInterceptor)))
	opts = append(opts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(newStreamInterceptor(s), grpc_prometheus.StreamServerInterceptor)))
	opts = append(opts, grpc.MaxRecvMsgSize(int(s.Cfg.MaxRequestBytes+grpcOverheadBytes)))
	opts = append(opts, grpc.MaxSendMsgSize(maxSendBytes))
	opts = append(opts, grpc.MaxConcurrentStreams(maxStreams))
	grpcServer := grpc.NewServer(append(opts, gopts...)...)
	pb.RegisterKVServer(grpcServer, NewQuotaKVServer(s))
	pb.RegisterWatchServer(grpcServer, NewWatchServer(s))
	pb.RegisterLeaseServer(grpcServer, NewQuotaLeaseServer(s))
	pb.RegisterClusterServer(grpcServer, NewClusterServer(s))
	pb.RegisterAuthServer(grpcServer, NewAuthServer(s))
	pb.RegisterMaintenanceServer(grpcServer, NewMaintenanceServer(s))
	hsrv := health.NewServer()
	hsrv.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(grpcServer, hsrv)
	grpc_prometheus.Register(grpcServer)
	return grpcServer
}
