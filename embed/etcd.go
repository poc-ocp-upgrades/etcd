package embed

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	defaultLog "log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
	"github.com/coreos/etcd/compactor"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api/etcdhttp"
	"github.com/coreos/etcd/etcdserver/api/v2http"
	"github.com/coreos/etcd/etcdserver/api/v2v3"
	"github.com/coreos/etcd/etcdserver/api/v3client"
	"github.com/coreos/etcd/etcdserver/api/v3rpc"
	"github.com/coreos/etcd/pkg/cors"
	"github.com/coreos/etcd/pkg/debugutil"
	runtimeutil "github.com/coreos/etcd/pkg/runtime"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/rafthttp"
	"github.com/coreos/pkg/capnslog"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var plog = capnslog.NewPackageLogger("github.com/coreos/etcd", "embed")

const (
	reservedInternalFDNum = 150
)

type Etcd struct {
	Peers			[]*peerListener
	Clients			[]net.Listener
	sctxs			map[string]*serveCtx
	metricsListeners	[]net.Listener
	Server			*etcdserver.EtcdServer
	cfg			Config
	stopc			chan struct{}
	errc			chan error
	closeOnce		sync.Once
}
type peerListener struct {
	net.Listener
	serve	func() error
	close	func(context.Context) error
}

func StartEtcd(inCfg *Config) (e *Etcd, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err = inCfg.Validate(); err != nil {
		return nil, err
	}
	serving := false
	e = &Etcd{cfg: *inCfg, stopc: make(chan struct{})}
	cfg := &e.cfg
	defer func() {
		if e == nil || err == nil {
			return
		}
		if !serving {
			for _, sctx := range e.sctxs {
				close(sctx.serversC)
			}
		}
		e.Close()
		e = nil
	}()
	if e.Peers, err = startPeerListeners(cfg); err != nil {
		return e, err
	}
	if e.sctxs, err = startClientListeners(cfg); err != nil {
		return e, err
	}
	for _, sctx := range e.sctxs {
		e.Clients = append(e.Clients, sctx.l)
	}
	var (
		urlsmap	types.URLsMap
		token	string
	)
	memberInitialized := true
	if !isMemberInitialized(cfg) {
		memberInitialized = false
		urlsmap, token, err = cfg.PeerURLsMapAndToken("etcd")
		if err != nil {
			return e, fmt.Errorf("error setting up initial cluster: %v", err)
		}
	}
	if len(cfg.AutoCompactionRetention) == 0 {
		cfg.AutoCompactionRetention = "0"
	}
	autoCompactionRetention, err := parseCompactionRetention(cfg.AutoCompactionMode, cfg.AutoCompactionRetention)
	if err != nil {
		return e, err
	}
	srvcfg := etcdserver.ServerConfig{Name: cfg.Name, ClientURLs: cfg.ACUrls, PeerURLs: cfg.APUrls, DataDir: cfg.Dir, DedicatedWALDir: cfg.WalDir, SnapCount: cfg.SnapCount, MaxSnapFiles: cfg.MaxSnapFiles, MaxWALFiles: cfg.MaxWalFiles, InitialPeerURLsMap: urlsmap, InitialClusterToken: token, DiscoveryURL: cfg.Durl, DiscoveryProxy: cfg.Dproxy, NewCluster: cfg.IsNewCluster(), ForceNewCluster: cfg.ForceNewCluster, PeerTLSInfo: cfg.PeerTLSInfo, TickMs: cfg.TickMs, ElectionTicks: cfg.ElectionTicks(), InitialElectionTickAdvance: cfg.InitialElectionTickAdvance, AutoCompactionRetention: autoCompactionRetention, AutoCompactionMode: cfg.AutoCompactionMode, QuotaBackendBytes: cfg.QuotaBackendBytes, MaxTxnOps: cfg.MaxTxnOps, MaxRequestBytes: cfg.MaxRequestBytes, StrictReconfigCheck: cfg.StrictReconfigCheck, ClientCertAuthEnabled: cfg.ClientTLSInfo.ClientCertAuth, AuthToken: cfg.AuthToken, InitialCorruptCheck: cfg.ExperimentalInitialCorruptCheck, CorruptCheckTime: cfg.ExperimentalCorruptCheckTime, Debug: cfg.Debug}
	if e.Server, err = etcdserver.NewServer(srvcfg); err != nil {
		return e, err
	}
	e.errc = make(chan error, len(e.Peers)+len(e.Clients)+2*len(e.sctxs))
	if memberInitialized {
		if err = e.Server.CheckInitialHashKV(); err != nil {
			e.Server = nil
			return e, err
		}
	}
	e.Server.Start()
	if err = e.servePeers(); err != nil {
		return e, err
	}
	if err = e.serveClients(); err != nil {
		return e, err
	}
	if err = e.serveMetrics(); err != nil {
		return e, err
	}
	serving = true
	return e, nil
}
func (e *Etcd) Config() Config {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.cfg
}
func (e *Etcd) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.closeOnce.Do(func() {
		close(e.stopc)
	})
	timeout := 2 * time.Second
	if e.Server != nil {
		timeout = e.Server.Cfg.ReqTimeout()
	}
	for _, sctx := range e.sctxs {
		for ss := range sctx.serversC {
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			stopServers(ctx, ss)
			cancel()
		}
	}
	for _, sctx := range e.sctxs {
		sctx.cancel()
	}
	for i := range e.Clients {
		if e.Clients[i] != nil {
			e.Clients[i].Close()
		}
	}
	for i := range e.metricsListeners {
		e.metricsListeners[i].Close()
	}
	if e.Server != nil {
		e.Server.Stop()
	}
	for i := range e.Peers {
		if e.Peers[i] != nil && e.Peers[i].close != nil {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			e.Peers[i].close(ctx)
			cancel()
		}
	}
}
func stopServers(ctx context.Context, ss *servers) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	shutdownNow := func() {
		ss.http.Shutdown(ctx)
		ss.grpc.Stop()
	}
	if ss.secure {
		shutdownNow()
		return
	}
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		ss.grpc.GracefulStop()
	}()
	select {
	case <-ch:
	case <-ctx.Done():
		shutdownNow()
		<-ch
	}
}
func (e *Etcd) Err() <-chan error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.errc
}
func startPeerListeners(cfg *Config) (peers []*peerListener, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err = updateCipherSuites(&cfg.PeerTLSInfo, cfg.CipherSuites); err != nil {
		return nil, err
	}
	if err = cfg.PeerSelfCert(); err != nil {
		plog.Fatalf("could not get certs (%v)", err)
	}
	if !cfg.PeerTLSInfo.Empty() {
		plog.Infof("peerTLS: %s", cfg.PeerTLSInfo)
	}
	peers = make([]*peerListener, len(cfg.LPUrls))
	defer func() {
		if err == nil {
			return
		}
		for i := range peers {
			if peers[i] != nil && peers[i].close != nil {
				plog.Info("stopping listening for peers on ", cfg.LPUrls[i].String())
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				peers[i].close(ctx)
				cancel()
			}
		}
	}()
	for i, u := range cfg.LPUrls {
		if u.Scheme == "http" {
			if !cfg.PeerTLSInfo.Empty() {
				plog.Warningf("The scheme of peer url %s is HTTP while peer key/cert files are presented. Ignored peer key/cert files.", u.String())
			}
			if cfg.PeerTLSInfo.ClientCertAuth {
				plog.Warningf("The scheme of peer url %s is HTTP while client cert auth (--peer-client-cert-auth) is enabled. Ignored client cert auth for this url.", u.String())
			}
		}
		peers[i] = &peerListener{close: func(context.Context) error {
			return nil
		}}
		peers[i].Listener, err = rafthttp.NewListener(u, &cfg.PeerTLSInfo)
		if err != nil {
			return nil, err
		}
		peers[i].close = func(context.Context) error {
			return peers[i].Listener.Close()
		}
		plog.Info("listening for peers on ", u.String())
	}
	return peers, nil
}
func (e *Etcd) servePeers() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ph := etcdhttp.NewPeerHandler(e.Server)
	var peerTLScfg *tls.Config
	if !e.cfg.PeerTLSInfo.Empty() {
		if peerTLScfg, err = e.cfg.PeerTLSInfo.ServerConfig(); err != nil {
			return err
		}
	}
	for _, p := range e.Peers {
		gs := v3rpc.Server(e.Server, peerTLScfg)
		m := cmux.New(p.Listener)
		go gs.Serve(m.Match(cmux.HTTP2()))
		srv := &http.Server{Handler: grpcHandlerFunc(gs, ph), ReadTimeout: 5 * time.Minute, ErrorLog: defaultLog.New(ioutil.Discard, "", 0)}
		go srv.Serve(m.Match(cmux.Any()))
		p.serve = func() error {
			return m.Serve()
		}
		p.close = func(ctx context.Context) error {
			stopServers(ctx, &servers{secure: peerTLScfg != nil, grpc: gs, http: srv})
			return nil
		}
	}
	for _, pl := range e.Peers {
		go func(l *peerListener) {
			e.errHandler(l.serve())
		}(pl)
	}
	return nil
}
func startClientListeners(cfg *Config) (sctxs map[string]*serveCtx, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err = updateCipherSuites(&cfg.ClientTLSInfo, cfg.CipherSuites); err != nil {
		return nil, err
	}
	if err = cfg.ClientSelfCert(); err != nil {
		plog.Fatalf("could not get certs (%v)", err)
	}
	if cfg.EnablePprof {
		plog.Infof("pprof is enabled under %s", debugutil.HTTPPrefixPProf)
	}
	sctxs = make(map[string]*serveCtx)
	for _, u := range cfg.LCUrls {
		sctx := newServeCtx()
		if u.Scheme == "http" || u.Scheme == "unix" {
			if !cfg.ClientTLSInfo.Empty() {
				plog.Warningf("The scheme of client url %s is HTTP while peer key/cert files are presented. Ignored key/cert files.", u.String())
			}
			if cfg.ClientTLSInfo.ClientCertAuth {
				plog.Warningf("The scheme of client url %s is HTTP while client cert auth (--client-cert-auth) is enabled. Ignored client cert auth for this url.", u.String())
			}
		}
		if (u.Scheme == "https" || u.Scheme == "unixs") && cfg.ClientTLSInfo.Empty() {
			return nil, fmt.Errorf("TLS key/cert (--cert-file, --key-file) must be provided for client url %s with HTTPs scheme", u.String())
		}
		proto := "tcp"
		addr := u.Host
		if u.Scheme == "unix" || u.Scheme == "unixs" {
			proto = "unix"
			addr = u.Host + u.Path
		}
		sctx.secure = u.Scheme == "https" || u.Scheme == "unixs"
		sctx.insecure = !sctx.secure
		if oldctx := sctxs[addr]; oldctx != nil {
			oldctx.secure = oldctx.secure || sctx.secure
			oldctx.insecure = oldctx.insecure || sctx.insecure
			continue
		}
		if sctx.l, err = net.Listen(proto, addr); err != nil {
			return nil, err
		}
		sctx.addr = addr
		if fdLimit, fderr := runtimeutil.FDLimit(); fderr == nil {
			if fdLimit <= reservedInternalFDNum {
				plog.Fatalf("file descriptor limit[%d] of etcd process is too low, and should be set higher than %d to ensure internal usage", fdLimit, reservedInternalFDNum)
			}
			sctx.l = transport.LimitListener(sctx.l, int(fdLimit-reservedInternalFDNum))
		}
		if proto == "tcp" {
			if sctx.l, err = transport.NewKeepAliveListener(sctx.l, "tcp", nil); err != nil {
				return nil, err
			}
		}
		plog.Info("listening for client requests on ", u.Host)
		defer func() {
			if err != nil {
				sctx.l.Close()
				plog.Info("stopping listening for client requests on ", u.Host)
			}
		}()
		for k := range cfg.UserHandlers {
			sctx.userHandlers[k] = cfg.UserHandlers[k]
		}
		sctx.serviceRegister = cfg.ServiceRegister
		if cfg.EnablePprof || cfg.Debug {
			sctx.registerPprof()
		}
		if cfg.Debug {
			sctx.registerTrace()
		}
		sctxs[addr] = sctx
	}
	return sctxs, nil
}
func (e *Etcd) serveClients() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !e.cfg.ClientTLSInfo.Empty() {
		plog.Infof("ClientTLS: %s", e.cfg.ClientTLSInfo)
	}
	if e.cfg.CorsInfo.String() != "" {
		plog.Infof("cors = %s", e.cfg.CorsInfo)
	}
	var h http.Handler
	if e.Config().EnableV2 {
		if len(e.Config().ExperimentalEnableV2V3) > 0 {
			srv := v2v3.NewServer(v3client.New(e.Server), e.cfg.ExperimentalEnableV2V3)
			h = v2http.NewClientHandler(srv, e.Server.Cfg.ReqTimeout())
		} else {
			h = v2http.NewClientHandler(e.Server, e.Server.Cfg.ReqTimeout())
		}
	} else {
		mux := http.NewServeMux()
		etcdhttp.HandleBasic(mux, e.Server)
		h = mux
	}
	h = http.Handler(&cors.CORSHandler{Handler: h, Info: e.cfg.CorsInfo})
	gopts := []grpc.ServerOption{}
	if e.cfg.GRPCKeepAliveMinTime > time.Duration(0) {
		gopts = append(gopts, grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{MinTime: e.cfg.GRPCKeepAliveMinTime, PermitWithoutStream: false}))
	}
	if e.cfg.GRPCKeepAliveInterval > time.Duration(0) && e.cfg.GRPCKeepAliveTimeout > time.Duration(0) {
		gopts = append(gopts, grpc.KeepaliveParams(keepalive.ServerParameters{Time: e.cfg.GRPCKeepAliveInterval, Timeout: e.cfg.GRPCKeepAliveTimeout}))
	}
	for _, sctx := range e.sctxs {
		go func(s *serveCtx) {
			e.errHandler(s.serve(e.Server, &e.cfg.ClientTLSInfo, h, e.errHandler, gopts...))
		}(sctx)
	}
	return nil
}
func (e *Etcd) serveMetrics() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if e.cfg.Metrics == "extensive" {
		grpc_prometheus.EnableHandlingTimeHistogram()
	}
	if len(e.cfg.ListenMetricsUrls) > 0 {
		metricsMux := http.NewServeMux()
		etcdhttp.HandleMetricsHealth(metricsMux, e.Server)
		for _, murl := range e.cfg.ListenMetricsUrls {
			tlsInfo := &e.cfg.ClientTLSInfo
			if murl.Scheme == "http" {
				tlsInfo = nil
			}
			ml, err := transport.NewListener(murl.Host, murl.Scheme, tlsInfo)
			if err != nil {
				return err
			}
			e.metricsListeners = append(e.metricsListeners, ml)
			go func(u url.URL, ln net.Listener) {
				plog.Info("listening for metrics on ", u.String())
				e.errHandler(http.Serve(ln, metricsMux))
			}(murl, ml)
		}
	}
	return nil
}
func (e *Etcd) errHandler(err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case <-e.stopc:
		return
	default:
	}
	select {
	case <-e.stopc:
	case e.errc <- err:
	}
}
func parseCompactionRetention(mode, retention string) (ret time.Duration, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	h, err := strconv.Atoi(retention)
	if err == nil {
		switch mode {
		case compactor.ModeRevision:
			ret = time.Duration(int64(h))
		case compactor.ModePeriodic:
			ret = time.Duration(int64(h)) * time.Hour
		}
	} else {
		ret, err = time.ParseDuration(retention)
		if err != nil {
			return 0, fmt.Errorf("error parsing CompactionRetention: %v", err)
		}
	}
	return ret, nil
}
