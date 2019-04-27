package clientv3

import (
	"math"
	"google.golang.org/grpc"
)

var (
	defaultFailFast			= grpc.FailFast(true)
	defaultMaxCallSendMsgSize	= grpc.MaxCallSendMsgSize(2 * 1024 * 1024)
	defaultMaxCallRecvMsgSize	= grpc.MaxCallRecvMsgSize(math.MaxInt32)
)
var defaultCallOpts = []grpc.CallOption{defaultFailFast, defaultMaxCallSendMsgSize, defaultMaxCallRecvMsgSize}

const MaxLeaseTTL = 9000000000
