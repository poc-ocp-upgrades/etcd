package clientv3

import (
	"context"
	"io"
	"sync"
	"time"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func (c *Client) unaryClientInterceptor(logger *zap.Logger, optFuncs ...retryOption) grpc.UnaryClientInterceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	intOpts := reuseOrNewWithCallOptions(defaultOptions, optFuncs)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := reuseOrNewWithCallOptions(intOpts, retryOpts)
		if callOpts.max == 0 {
			return invoker(ctx, method, req, reply, cc, grpcOpts...)
		}
		var lastErr error
		for attempt := uint(0); attempt < callOpts.max; attempt++ {
			if err := waitRetryBackoff(ctx, attempt, callOpts); err != nil {
				return err
			}
			logger.Debug("retrying of unary invoker", zap.String("target", cc.Target()), zap.Uint("attempt", attempt))
			lastErr = invoker(ctx, method, req, reply, cc, grpcOpts...)
			if lastErr == nil {
				return nil
			}
			logger.Warn("retrying of unary invoker failed", zap.String("target", cc.Target()), zap.Uint("attempt", attempt), zap.Error(lastErr))
			if isContextError(lastErr) {
				if ctx.Err() != nil {
					return lastErr
				}
				continue
			}
			if callOpts.retryAuth && rpctypes.Error(lastErr) == rpctypes.ErrInvalidAuthToken {
				gterr := c.getToken(ctx)
				if gterr != nil {
					logger.Warn("retrying of unary invoker failed to fetch new auth token", zap.String("target", cc.Target()), zap.Error(gterr))
					return lastErr
				}
				continue
			}
			if !isSafeRetry(c.lg, lastErr, callOpts) {
				return lastErr
			}
		}
		return lastErr
	}
}
func (c *Client) streamClientInterceptor(logger *zap.Logger, optFuncs ...retryOption) grpc.StreamClientInterceptor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	intOpts := reuseOrNewWithCallOptions(defaultOptions, optFuncs)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		grpcOpts, retryOpts := filterCallOptions(opts)
		callOpts := reuseOrNewWithCallOptions(intOpts, retryOpts)
		if callOpts.max == 0 {
			return streamer(ctx, desc, cc, method, grpcOpts...)
		}
		if desc.ClientStreams {
			return nil, grpc.Errorf(codes.Unimplemented, "clientv3/retry_interceptor: cannot retry on ClientStreams, set Disable()")
		}
		newStreamer, err := streamer(ctx, desc, cc, method, grpcOpts...)
		logger.Warn("retry stream intercept", zap.Error(err))
		if err != nil {
			return nil, err
		}
		retryingStreamer := &serverStreamingRetryingStream{client: c, ClientStream: newStreamer, callOpts: callOpts, ctx: ctx, streamerCall: func(ctx context.Context) (grpc.ClientStream, error) {
			return streamer(ctx, desc, cc, method, grpcOpts...)
		}}
		return retryingStreamer, nil
	}
}

type serverStreamingRetryingStream struct {
	grpc.ClientStream
	client		*Client
	bufferedSends	[]interface{}
	receivedGood	bool
	wasClosedSend	bool
	ctx		context.Context
	callOpts	*options
	streamerCall	func(ctx context.Context) (grpc.ClientStream, error)
	mu		sync.RWMutex
}

func (s *serverStreamingRetryingStream) setStream(clientStream grpc.ClientStream) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	s.ClientStream = clientStream
	s.mu.Unlock()
}
func (s *serverStreamingRetryingStream) getStream() grpc.ClientStream {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ClientStream
}
func (s *serverStreamingRetryingStream) SendMsg(m interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	s.bufferedSends = append(s.bufferedSends, m)
	s.mu.Unlock()
	return s.getStream().SendMsg(m)
}
func (s *serverStreamingRetryingStream) CloseSend() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.Lock()
	s.wasClosedSend = true
	s.mu.Unlock()
	return s.getStream().CloseSend()
}
func (s *serverStreamingRetryingStream) Header() (metadata.MD, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.getStream().Header()
}
func (s *serverStreamingRetryingStream) Trailer() metadata.MD {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.getStream().Trailer()
}
func (s *serverStreamingRetryingStream) RecvMsg(m interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	attemptRetry, lastErr := s.receiveMsgAndIndicateRetry(m)
	if !attemptRetry {
		return lastErr
	}
	for attempt := uint(1); attempt < s.callOpts.max; attempt++ {
		if err := waitRetryBackoff(s.ctx, attempt, s.callOpts); err != nil {
			return err
		}
		newStream, err := s.reestablishStreamAndResendBuffer(s.ctx)
		if err != nil {
			return err
		}
		s.setStream(newStream)
		attemptRetry, lastErr = s.receiveMsgAndIndicateRetry(m)
		if !attemptRetry {
			return lastErr
		}
	}
	return lastErr
}
func (s *serverStreamingRetryingStream) receiveMsgAndIndicateRetry(m interface{}) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.RLock()
	wasGood := s.receivedGood
	s.mu.RUnlock()
	err := s.getStream().RecvMsg(m)
	if err == nil || err == io.EOF {
		s.mu.Lock()
		s.receivedGood = true
		s.mu.Unlock()
		return false, err
	} else if wasGood {
		return false, err
	}
	if isContextError(err) {
		if s.ctx.Err() != nil {
			return false, err
		}
		return true, err
	}
	if s.callOpts.retryAuth && rpctypes.Error(err) == rpctypes.ErrInvalidAuthToken {
		gterr := s.client.getToken(s.ctx)
		if gterr != nil {
			s.client.lg.Warn("retry failed to fetch new auth token", zap.Error(gterr))
			return false, err
		}
		return true, err
	}
	return isSafeRetry(s.client.lg, err, s.callOpts), err
}
func (s *serverStreamingRetryingStream) reestablishStreamAndResendBuffer(callCtx context.Context) (grpc.ClientStream, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.mu.RLock()
	bufferedSends := s.bufferedSends
	s.mu.RUnlock()
	newStream, err := s.streamerCall(callCtx)
	if err != nil {
		return nil, err
	}
	for _, msg := range bufferedSends {
		if err := newStream.SendMsg(msg); err != nil {
			return nil, err
		}
	}
	if err := newStream.CloseSend(); err != nil {
		return nil, err
	}
	return newStream, nil
}
func waitRetryBackoff(ctx context.Context, attempt uint, callOpts *options) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	waitTime := time.Duration(0)
	if attempt > 0 {
		waitTime = callOpts.backoffFunc(attempt)
	}
	if waitTime > 0 {
		timer := time.NewTimer(waitTime)
		select {
		case <-ctx.Done():
			timer.Stop()
			return contextErrToGrpcErr(ctx.Err())
		case <-timer.C:
		}
	}
	return nil
}
func isSafeRetry(lg *zap.Logger, err error, callOpts *options) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isContextError(err) {
		return false
	}
	switch callOpts.retryPolicy {
	case repeatable:
		return isSafeRetryImmutableRPC(err)
	case nonRepeatable:
		return isSafeRetryMutableRPC(err)
	default:
		lg.Warn("unrecognized retry policy", zap.String("retryPolicy", callOpts.retryPolicy.String()))
		return false
	}
}
func isContextError(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return grpc.Code(err) == codes.DeadlineExceeded || grpc.Code(err) == codes.Canceled
}
func contextErrToGrpcErr(err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch err {
	case context.DeadlineExceeded:
		return grpc.Errorf(codes.DeadlineExceeded, err.Error())
	case context.Canceled:
		return grpc.Errorf(codes.Canceled, err.Error())
	default:
		return grpc.Errorf(codes.Unknown, err.Error())
	}
}

var (
	defaultOptions = &options{retryPolicy: nonRepeatable, max: 0, backoffFunc: backoffLinearWithJitter(50*time.Millisecond, 0.10), retryAuth: true}
)

type backoffFunc func(attempt uint) time.Duration

func withRetryPolicy(rp retryPolicy) retryOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return retryOption{applyFunc: func(o *options) {
		o.retryPolicy = rp
	}}
}
func withAuthRetry(retryAuth bool) retryOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return retryOption{applyFunc: func(o *options) {
		o.retryAuth = retryAuth
	}}
}
func withMax(maxRetries uint) retryOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return retryOption{applyFunc: func(o *options) {
		o.max = maxRetries
	}}
}
func withBackoff(bf backoffFunc) retryOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return retryOption{applyFunc: func(o *options) {
		o.backoffFunc = bf
	}}
}

type options struct {
	retryPolicy	retryPolicy
	max		uint
	backoffFunc	backoffFunc
	retryAuth	bool
}
type retryOption struct {
	grpc.EmptyCallOption
	applyFunc	func(opt *options)
}

func reuseOrNewWithCallOptions(opt *options, retryOptions []retryOption) *options {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(retryOptions) == 0 {
		return opt
	}
	optCopy := &options{}
	*optCopy = *opt
	for _, f := range retryOptions {
		f.applyFunc(optCopy)
	}
	return optCopy
}
func filterCallOptions(callOptions []grpc.CallOption) (grpcOptions []grpc.CallOption, retryOptions []retryOption) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, opt := range callOptions {
		if co, ok := opt.(retryOption); ok {
			retryOptions = append(retryOptions, co)
		} else {
			grpcOptions = append(grpcOptions, opt)
		}
	}
	return grpcOptions, retryOptions
}
func backoffLinearWithJitter(waitBetween time.Duration, jitterFraction float64) backoffFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(attempt uint) time.Duration {
		return jitterUp(waitBetween, jitterFraction)
	}
}
