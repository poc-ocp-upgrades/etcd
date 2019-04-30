package clientv3

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/google/uuid"
	"go.etcd.io/etcd/clientv3/balancer"
	"go.etcd.io/etcd/clientv3/balancer/picker"
	"go.etcd.io/etcd/clientv3/balancer/resolver/endpoint"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"go.etcd.io/etcd/pkg/logutil"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	ErrNoAvailableEndpoints	= errors.New("etcdclient: no available endpoints")
	ErrOldCluster		= errors.New("etcdclient: old cluster version")
	roundRobinBalancerName	= fmt.Sprintf("etcd-%s", picker.RoundrobinBalanced.String())
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg := zap.NewNop()
	if os.Getenv("ETCD_CLIENT_DEBUG") != "" {
		var err error
		lg, err = zap.NewProductionConfig().Build()
		if err != nil {
			panic(err)
		}
	}
	balancer.RegisterBuilder(balancer.Config{Policy: picker.RoundrobinBalanced, Name: roundRobinBalancerName, Logger: lg})
}

type Client struct {
	Cluster
	KV
	Lease
	Watcher
	Auth
	Maintenance
	conn		*grpc.ClientConn
	cfg		Config
	creds		*credentials.TransportCredentials
	balancer	balancer.Balancer
	resolverGroup	*endpoint.ResolverGroup
	mu		*sync.Mutex
	ctx		context.Context
	cancel		context.CancelFunc
	Username	string
	Password	string
	tokenCred	*authTokenCredential
	callOpts	[]grpc.CallOption
	lg		*zap.Logger
}

func New(cfg Config) (*Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(cfg.Endpoints) == 0 {
		return nil, ErrNoAvailableEndpoints
	}
	return newClient(&cfg)
}
func NewCtxClient(ctx context.Context) *Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cctx, cancel := context.WithCancel(ctx)
	return &Client{ctx: cctx, cancel: cancel}
}
func NewFromURL(url string) (*Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return New(Config{Endpoints: []string{url}})
}
func NewFromURLs(urls []string) (*Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return New(Config{Endpoints: urls})
}
func (c *Client) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.cancel()
	c.Watcher.Close()
	c.Lease.Close()
	if c.resolverGroup != nil {
		c.resolverGroup.Close()
	}
	if c.conn != nil {
		return toErr(c.ctx, c.conn.Close())
	}
	return c.ctx.Err()
}
func (c *Client) Ctx() context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ctx
}
func (c *Client) Endpoints() (eps []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eps = make([]string, len(c.cfg.Endpoints))
	copy(eps, c.cfg.Endpoints)
	return
}
func (c *Client) SetEndpoints(eps ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cfg.Endpoints = eps
	c.resolverGroup.SetEndpoints(eps)
}
func (c *Client) Sync(ctx context.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mresp, err := c.MemberList(ctx)
	if err != nil {
		return err
	}
	var eps []string
	for _, m := range mresp.Members {
		eps = append(eps, m.ClientURLs...)
	}
	c.SetEndpoints(eps...)
	return nil
}
func (c *Client) autoSync() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.cfg.AutoSyncInterval == time.Duration(0) {
		return
	}
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(c.cfg.AutoSyncInterval):
			ctx, cancel := context.WithTimeout(c.ctx, 5*time.Second)
			err := c.Sync(ctx)
			cancel()
			if err != nil && err != c.ctx.Err() {
				lg.Lvl(4).Infof("Auto sync endpoints failed: %v", err)
			}
		}
	}
}

type authTokenCredential struct {
	token	string
	tokenMu	*sync.RWMutex
}

func (cred authTokenCredential) RequireTransportSecurity() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (cred authTokenCredential) GetRequestMetadata(ctx context.Context, s ...string) (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cred.tokenMu.RLock()
	defer cred.tokenMu.RUnlock()
	return map[string]string{rpctypes.TokenFieldNameGRPC: cred.token}, nil
}
func (c *Client) processCreds(scheme string) (creds *credentials.TransportCredentials) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	creds = c.creds
	switch scheme {
	case "unix":
	case "http":
		creds = nil
	case "https", "unixs":
		if creds != nil {
			break
		}
		tlsconfig := &tls.Config{}
		emptyCreds := credentials.NewTLS(tlsconfig)
		creds = &emptyCreds
	default:
		creds = nil
	}
	return creds
}
func (c *Client) dialSetupOpts(creds *credentials.TransportCredentials, dopts ...grpc.DialOption) (opts []grpc.DialOption, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.cfg.DialKeepAliveTime > 0 {
		params := keepalive.ClientParameters{Time: c.cfg.DialKeepAliveTime, Timeout: c.cfg.DialKeepAliveTimeout, PermitWithoutStream: c.cfg.PermitWithoutStream}
		opts = append(opts, grpc.WithKeepaliveParams(params))
	}
	opts = append(opts, dopts...)
	f := func(dialEp string, t time.Duration) (net.Conn, error) {
		proto, host, _ := endpoint.ParseEndpoint(dialEp)
		select {
		case <-c.ctx.Done():
			return nil, c.ctx.Err()
		default:
		}
		dialer := &net.Dialer{Timeout: t}
		return dialer.DialContext(c.ctx, proto, host)
	}
	opts = append(opts, grpc.WithDialer(f))
	if creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(*creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	rrBackoff := withBackoff(c.roundRobinQuorumBackoff(defaultBackoffWaitBetween, defaultBackoffJitterFraction))
	opts = append(opts, grpc.WithStreamInterceptor(c.streamClientInterceptor(c.lg, withMax(0), rrBackoff)), grpc.WithUnaryInterceptor(c.unaryClientInterceptor(c.lg, withMax(defaultUnaryMaxRetries), rrBackoff)))
	return opts, nil
}
func (c *Client) Dial(ep string) (*grpc.ClientConn, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	creds := c.directDialCreds(ep)
	return c.dial(fmt.Sprintf("passthrough:///%s", ep), creds)
}
func (c *Client) getToken(ctx context.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	var auth *authenticator
	for i := 0; i < len(c.cfg.Endpoints); i++ {
		ep := c.cfg.Endpoints[i]
		var dOpts []grpc.DialOption
		_, host, _ := endpoint.ParseEndpoint(ep)
		target := c.resolverGroup.Target(host)
		creds := c.dialWithBalancerCreds(ep)
		dOpts, err = c.dialSetupOpts(creds, c.cfg.DialOptions...)
		if err != nil {
			err = fmt.Errorf("failed to configure auth dialer: %v", err)
			continue
		}
		dOpts = append(dOpts, grpc.WithBalancerName(roundRobinBalancerName))
		auth, err = newAuthenticator(ctx, target, dOpts, c)
		if err != nil {
			continue
		}
		defer auth.close()
		var resp *AuthenticateResponse
		resp, err = auth.authenticate(ctx, c.Username, c.Password)
		if err != nil {
			if err == rpctypes.ErrAuthNotEnabled {
				return err
			}
			continue
		}
		c.tokenCred.tokenMu.Lock()
		c.tokenCred.token = resp.Token
		c.tokenCred.tokenMu.Unlock()
		return nil
	}
	return err
}
func (c *Client) dialWithBalancer(ep string, dopts ...grpc.DialOption) (*grpc.ClientConn, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, host, _ := endpoint.ParseEndpoint(ep)
	target := c.resolverGroup.Target(host)
	creds := c.dialWithBalancerCreds(ep)
	return c.dial(target, creds, dopts...)
}
func (c *Client) dial(target string, creds *credentials.TransportCredentials, dopts ...grpc.DialOption) (*grpc.ClientConn, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts, err := c.dialSetupOpts(creds, dopts...)
	if err != nil {
		return nil, fmt.Errorf("failed to configure dialer: %v", err)
	}
	if c.Username != "" && c.Password != "" {
		c.tokenCred = &authTokenCredential{tokenMu: &sync.RWMutex{}}
		ctx, cancel := c.ctx, func() {
		}
		if c.cfg.DialTimeout > 0 {
			ctx, cancel = context.WithTimeout(ctx, c.cfg.DialTimeout)
		}
		err = c.getToken(ctx)
		if err != nil {
			if toErr(ctx, err) != rpctypes.ErrAuthNotEnabled {
				if err == ctx.Err() && ctx.Err() != c.ctx.Err() {
					err = context.DeadlineExceeded
				}
				cancel()
				return nil, err
			}
		} else {
			opts = append(opts, grpc.WithPerRPCCredentials(c.tokenCred))
		}
		cancel()
	}
	opts = append(opts, c.cfg.DialOptions...)
	dctx := c.ctx
	if c.cfg.DialTimeout > 0 {
		var cancel context.CancelFunc
		dctx, cancel = context.WithTimeout(c.ctx, c.cfg.DialTimeout)
		defer cancel()
	}
	conn, err := grpc.DialContext(dctx, target, opts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func (c *Client) directDialCreds(ep string) *credentials.TransportCredentials {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, hostPort, scheme := endpoint.ParseEndpoint(ep)
	creds := c.creds
	if len(scheme) != 0 {
		creds = c.processCreds(scheme)
		if creds != nil {
			c := *creds
			clone := c.Clone()
			host, _ := endpoint.ParseHostPort(hostPort)
			clone.OverrideServerName(host)
			creds = &clone
		}
	}
	return creds
}
func (c *Client) dialWithBalancerCreds(ep string) *credentials.TransportCredentials {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, _, scheme := endpoint.ParseEndpoint(ep)
	creds := c.creds
	if len(scheme) != 0 {
		creds = c.processCreds(scheme)
	}
	return creds
}
func WithRequireLeader(ctx context.Context) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	md := metadata.Pairs(rpctypes.MetadataRequireLeaderKey, rpctypes.MetadataHasLeader)
	return metadata.NewOutgoingContext(ctx, md)
}
func newClient(cfg *Config) (*Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if cfg == nil {
		cfg = &Config{}
	}
	var creds *credentials.TransportCredentials
	if cfg.TLS != nil {
		c := credentials.NewTLS(cfg.TLS)
		creds = &c
	}
	baseCtx := context.TODO()
	if cfg.Context != nil {
		baseCtx = cfg.Context
	}
	ctx, cancel := context.WithCancel(baseCtx)
	client := &Client{conn: nil, cfg: *cfg, creds: creds, ctx: ctx, cancel: cancel, mu: new(sync.Mutex), callOpts: defaultCallOpts}
	lcfg := logutil.DefaultZapLoggerConfig
	if cfg.LogConfig != nil {
		lcfg = *cfg.LogConfig
	}
	var err error
	client.lg, err = lcfg.Build()
	if err != nil {
		return nil, err
	}
	if cfg.Username != "" && cfg.Password != "" {
		client.Username = cfg.Username
		client.Password = cfg.Password
	}
	if cfg.MaxCallSendMsgSize > 0 || cfg.MaxCallRecvMsgSize > 0 {
		if cfg.MaxCallRecvMsgSize > 0 && cfg.MaxCallSendMsgSize > cfg.MaxCallRecvMsgSize {
			return nil, fmt.Errorf("gRPC message recv limit (%d bytes) must be greater than send limit (%d bytes)", cfg.MaxCallRecvMsgSize, cfg.MaxCallSendMsgSize)
		}
		callOpts := []grpc.CallOption{defaultFailFast, defaultMaxCallSendMsgSize, defaultMaxCallRecvMsgSize}
		if cfg.MaxCallSendMsgSize > 0 {
			callOpts[1] = grpc.MaxCallSendMsgSize(cfg.MaxCallSendMsgSize)
		}
		if cfg.MaxCallRecvMsgSize > 0 {
			callOpts[2] = grpc.MaxCallRecvMsgSize(cfg.MaxCallRecvMsgSize)
		}
		client.callOpts = callOpts
	}
	client.resolverGroup, err = endpoint.NewResolverGroup(fmt.Sprintf("client-%s", uuid.New().String()))
	if err != nil {
		client.cancel()
		return nil, err
	}
	client.resolverGroup.SetEndpoints(cfg.Endpoints)
	if len(cfg.Endpoints) < 1 {
		return nil, fmt.Errorf("at least one Endpoint must is required in client config")
	}
	dialEndpoint := cfg.Endpoints[0]
	conn, err := client.dialWithBalancer(dialEndpoint, grpc.WithBalancerName(roundRobinBalancerName))
	if err != nil {
		client.cancel()
		client.resolverGroup.Close()
		return nil, err
	}
	client.conn = conn
	client.Cluster = NewCluster(client)
	client.KV = NewKV(client)
	client.Lease = NewLease(client)
	client.Watcher = NewWatcher(client)
	client.Auth = NewAuth(client)
	client.Maintenance = NewMaintenance(client)
	if cfg.RejectOldCluster {
		if err := client.checkVersion(); err != nil {
			client.Close()
			return nil, err
		}
	}
	go client.autoSync()
	return client, nil
}
func (c *Client) roundRobinQuorumBackoff(waitBetween time.Duration, jitterFraction float64) backoffFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(attempt uint) time.Duration {
		n := uint(len(c.Endpoints()))
		quorum := (n/2 + 1)
		if attempt%quorum == 0 {
			c.lg.Debug("backoff", zap.Uint("attempt", attempt), zap.Uint("quorum", quorum), zap.Duration("waitBetween", waitBetween), zap.Float64("jitterFraction", jitterFraction))
			return jitterUp(waitBetween, jitterFraction)
		}
		c.lg.Debug("backoff skipped", zap.Uint("attempt", attempt), zap.Uint("quorum", quorum))
		return 0
	}
}
func (c *Client) checkVersion() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wg sync.WaitGroup
	errc := make(chan error, len(c.cfg.Endpoints))
	ctx, cancel := context.WithCancel(c.ctx)
	if c.cfg.DialTimeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, c.cfg.DialTimeout)
	}
	wg.Add(len(c.cfg.Endpoints))
	for _, ep := range c.cfg.Endpoints {
		go func(e string) {
			defer wg.Done()
			resp, rerr := c.Status(ctx, e)
			if rerr != nil {
				errc <- rerr
				return
			}
			vs := strings.Split(resp.Version, ".")
			maj, min := 0, 0
			if len(vs) >= 2 {
				maj, _ = strconv.Atoi(vs[0])
				min, rerr = strconv.Atoi(vs[1])
			}
			if maj < 3 || (maj == 3 && min < 2) {
				rerr = ErrOldCluster
			}
			errc <- rerr
		}(ep)
	}
	for i := 0; i < len(c.cfg.Endpoints); i++ {
		if err = <-errc; err == nil {
			break
		}
	}
	cancel()
	wg.Wait()
	return err
}
func (c *Client) ActiveConnection() *grpc.ClientConn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.conn
}
func isHaltErr(ctx context.Context, err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ctx != nil && ctx.Err() != nil {
		return true
	}
	if err == nil {
		return false
	}
	ev, _ := status.FromError(err)
	return ev.Code() != codes.Unavailable && ev.Code() != codes.Internal
}
func isUnavailableErr(ctx context.Context, err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ctx != nil && ctx.Err() != nil {
		return false
	}
	if err == nil {
		return false
	}
	ev, _ := status.FromError(err)
	return ev.Code() == codes.Unavailable
}
func toErr(ctx context.Context, err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return nil
	}
	err = rpctypes.Error(err)
	if _, ok := err.(rpctypes.EtcdError); ok {
		return err
	}
	if ev, ok := status.FromError(err); ok {
		code := ev.Code()
		switch code {
		case codes.DeadlineExceeded:
			fallthrough
		case codes.Canceled:
			if ctx.Err() != nil {
				err = ctx.Err()
			}
		case codes.Unavailable:
		case codes.FailedPrecondition:
			err = grpc.ErrClientConnClosing
		}
	}
	return err
}
func canceledByCaller(stopCtx context.Context, err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if stopCtx.Err() == nil || err == nil {
		return false
	}
	return err == context.Canceled || err == context.DeadlineExceeded
}
func IsConnCanceled(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err == nil {
		return false
	}
	s, ok := status.FromError(err)
	if ok {
		return s.Code() == codes.Canceled || s.Message() == "transport is closing"
	}
	if err == context.Canceled {
		return true
	}
	return strings.Contains(err.Error(), "grpc: the client connection is closing")
}
