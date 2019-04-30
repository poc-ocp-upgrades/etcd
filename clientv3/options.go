package clientv3

import (
	"math"
	"time"
	"google.golang.org/grpc"
)

var (
	defaultFailFast				= grpc.FailFast(false)
	defaultMaxCallSendMsgSize		= grpc.MaxCallSendMsgSize(2 * 1024 * 1024)
	defaultMaxCallRecvMsgSize		= grpc.MaxCallRecvMsgSize(math.MaxInt32)
	defaultUnaryMaxRetries		uint	= 100
	defaultStreamMaxRetries			= uint(^uint(0))
	defaultBackoffWaitBetween		= 25 * time.Millisecond
	defaultBackoffJitterFraction		= 0.10
)
var defaultCallOpts = []grpc.CallOption{defaultFailFast, defaultMaxCallSendMsgSize, defaultMaxCallRecvMsgSize}

const MaxLeaseTTL = 9000000000
