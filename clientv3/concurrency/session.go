package concurrency

import (
	"context"
	"time"
	v3 "go.etcd.io/etcd/clientv3"
)

const defaultSessionTTL = 60

type Session struct {
	client	*v3.Client
	opts	*sessionOptions
	id	v3.LeaseID
	cancel	context.CancelFunc
	donec	<-chan struct{}
}

func NewSession(client *v3.Client, opts ...SessionOption) (*Session, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ops := &sessionOptions{ttl: defaultSessionTTL, ctx: client.Ctx()}
	for _, opt := range opts {
		opt(ops)
	}
	id := ops.leaseID
	if id == v3.NoLease {
		resp, err := client.Grant(ops.ctx, int64(ops.ttl))
		if err != nil {
			return nil, err
		}
		id = v3.LeaseID(resp.ID)
	}
	ctx, cancel := context.WithCancel(ops.ctx)
	keepAlive, err := client.KeepAlive(ctx, id)
	if err != nil || keepAlive == nil {
		cancel()
		return nil, err
	}
	donec := make(chan struct{})
	s := &Session{client: client, opts: ops, id: id, cancel: cancel, donec: donec}
	go func() {
		defer close(donec)
		for range keepAlive {
		}
	}()
	return s, nil
}
func (s *Session) Client() *v3.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.client
}
func (s *Session) Lease() v3.LeaseID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.id
}
func (s *Session) Done() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.donec
}
func (s *Session) Orphan() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.cancel()
	<-s.donec
}
func (s *Session) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Orphan()
	ctx, cancel := context.WithTimeout(s.opts.ctx, time.Duration(s.opts.ttl)*time.Second)
	_, err := s.client.Revoke(ctx, s.id)
	cancel()
	return err
}

type sessionOptions struct {
	ttl	int
	leaseID	v3.LeaseID
	ctx	context.Context
}
type SessionOption func(*sessionOptions)

func WithTTL(ttl int) SessionOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(so *sessionOptions) {
		if ttl > 0 {
			so.ttl = ttl
		}
	}
}
func WithLease(leaseID v3.LeaseID) SessionOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(so *sessionOptions) {
		so.leaseID = leaseID
	}
}
func WithContext(ctx context.Context) SessionOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(so *sessionOptions) {
		so.ctx = ctx
	}
}
