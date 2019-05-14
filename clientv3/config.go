package clientv3

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"time"
)

type Config struct {
	Endpoints            []string      `json:"endpoints"`
	AutoSyncInterval     time.Duration `json:"auto-sync-interval"`
	DialTimeout          time.Duration `json:"dial-timeout"`
	DialKeepAliveTime    time.Duration `json:"dial-keep-alive-time"`
	DialKeepAliveTimeout time.Duration `json:"dial-keep-alive-timeout"`
	MaxCallSendMsgSize   int
	MaxCallRecvMsgSize   int
	TLS                  *tls.Config
	Username             string `json:"username"`
	Password             string `json:"password"`
	RejectOldCluster     bool   `json:"reject-old-cluster"`
	DialOptions          []grpc.DialOption
	Context              context.Context
}
