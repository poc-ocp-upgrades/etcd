package clientv3

import (
	"google.golang.org/grpc"
	"math"
)

var (
	defaultFailFast           = grpc.FailFast(true)
	defaultMaxCallSendMsgSize = grpc.MaxCallSendMsgSize(2 * 1024 * 1024)
	defaultMaxCallRecvMsgSize = grpc.MaxCallRecvMsgSize(math.MaxInt32)
)
var defaultCallOpts = []grpc.CallOption{defaultFailFast, defaultMaxCallSendMsgSize, defaultMaxCallRecvMsgSize}

const MaxLeaseTTL = 9000000000
