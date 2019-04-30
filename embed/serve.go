package embed

import (
	"context"
	"fmt"
	"io/ioutil"
	defaultLog "log"
	"net"
	"net/http"
	"strings"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api/v3client"
	"go.etcd.io/etcd/etcdserver/api/v3election"
	"go.etcd.io/etcd/etcdserver/api/v3election/v3electionpb"
	v3electiongw "go.etcd.io/etcd/etcdserver/api/v3election/v3electionpb/gw"
	"go.etcd.io/etcd/etcdserver/api/v3lock"
	"go.etcd.io/etcd/etcdserver/api/v3lock/v3lockpb"
	v3lockgw "go.etcd.io/etcd/etcdserver/api/v3lock/v3lockpb/gw"
	"go.etcd.io/etcd/etcdserver/api/v3rpc"
	etcdservergw "go.etcd.io/etcd/etcdserver/etcdserverpb/gw"
	"go.etcd.io/etcd/pkg/debugutil"
	"go.etcd.io/etcd/pkg/httputil"
	"go.etcd.io/etcd/pkg/transport"
	gw "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"go.uber.org/zap"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type serveCtx struct {
	lg		*zap.Logger
	l		net.Listener
	addr		string
	network		string
	secure		bool
	insecure	bool
	ctx		context.Context
	cancel		context.CancelFunc
	userHandlers	map[string]http.Handler
	serviceRegister	func(*grpc.Server)
	serversC	chan *servers
}
type servers struct {
	secure	bool
	grpc	*grpc.Server
	http	*http.Server
}

func newServeCtx(lg *zap.Logger) *serveCtx {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithCancel(context.Background())
	return &serveCtx{lg: lg, ctx: ctx, cancel: cancel, userHandlers: make(map[string]http.Handler), serversC: make(chan *servers, 2)}
}
func (sctx *serveCtx) serve(s *etcdserver.EtcdServer, tlsinfo *transport.TLSInfo, handler http.Handler, errHandler func(error), gopts ...grpc.ServerOption) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	logger := defaultLog.New(ioutil.Discard, "etcdhttp", 0)
	<-s.ReadyNotify()
	if sctx.lg == nil {
		plog.Info("ready to serve client requests")
	}
	m := cmux.New(sctx.l)
	v3c := v3client.New(s)
	servElection := v3election.NewElectionServer(v3c)
	servLock := v3lock.NewLockServer(v3c)
	var gs *grpc.Server
	defer func() {
		if err != nil && gs != nil {
			gs.Stop()
		}
	}()
	if sctx.insecure {
		gs = v3rpc.Server(s, nil, gopts...)
		v3electionpb.RegisterElectionServer(gs, servElection)
		v3lockpb.RegisterLockServer(gs, servLock)
		if sctx.serviceRegister != nil {
			sctx.serviceRegister(gs)
		}
		grpcl := m.Match(cmux.HTTP2())
		go func() {
			errHandler(gs.Serve(grpcl))
		}()
		var gwmux *gw.ServeMux
		if s.Cfg.EnableGRPCGateway {
			gwmux, err = sctx.registerGateway([]grpc.DialOption{grpc.WithInsecure()})
			if err != nil {
				return err
			}
		}
		httpmux := sctx.createMux(gwmux, handler)
		srvhttp := &http.Server{Handler: createAccessController(sctx.lg, s, httpmux), ErrorLog: logger}
		httpl := m.Match(cmux.HTTP1())
		go func() {
			errHandler(srvhttp.Serve(httpl))
		}()
		sctx.serversC <- &servers{grpc: gs, http: srvhttp}
		if sctx.lg != nil {
			sctx.lg.Info("serving client traffic insecurely; this is strongly discouraged!", zap.String("address", sctx.l.Addr().String()))
		} else {
			plog.Noticef("serving insecure client requests on %s, this is strongly discouraged!", sctx.l.Addr().String())
		}
	}
	if sctx.secure {
		tlscfg, tlsErr := tlsinfo.ServerConfig()
		if tlsErr != nil {
			return tlsErr
		}
		gs = v3rpc.Server(s, tlscfg, gopts...)
		v3electionpb.RegisterElectionServer(gs, servElection)
		v3lockpb.RegisterLockServer(gs, servLock)
		if sctx.serviceRegister != nil {
			sctx.serviceRegister(gs)
		}
		handler = grpcHandlerFunc(gs, handler)
		var gwmux *gw.ServeMux
		if s.Cfg.EnableGRPCGateway {
			dtls := tlscfg.Clone()
			dtls.InsecureSkipVerify = true
			creds := credentials.NewTLS(dtls)
			opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
			gwmux, err = sctx.registerGateway(opts)
			if err != nil {
				return err
			}
		}
		var tlsl net.Listener
		tlsl, err = transport.NewTLSListener(m.Match(cmux.Any()), tlsinfo)
		if err != nil {
			return err
		}
		httpmux := sctx.createMux(gwmux, handler)
		srv := &http.Server{Handler: createAccessController(sctx.lg, s, httpmux), TLSConfig: tlscfg, ErrorLog: logger}
		go func() {
			errHandler(srv.Serve(tlsl))
		}()
		sctx.serversC <- &servers{secure: true, grpc: gs, http: srv}
		if sctx.lg != nil {
			sctx.lg.Info("serving client traffic insecurely", zap.String("address", sctx.l.Addr().String()))
		} else {
			plog.Infof("serving client requests on %s", sctx.l.Addr().String())
		}
	}
	close(sctx.serversC)
	return m.Serve()
}
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if otherHandler == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			grpcServer.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

type registerHandlerFunc func(context.Context, *gw.ServeMux, *grpc.ClientConn) error

func (sctx *serveCtx) registerGateway(opts []grpc.DialOption) (*gw.ServeMux, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx := sctx.ctx
	addr := sctx.addr
	if network := sctx.network; network == "unix" {
		addr = fmt.Sprintf("%s://%s", network, addr)
	}
	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, err
	}
	gwmux := gw.NewServeMux()
	handlers := []registerHandlerFunc{etcdservergw.RegisterKVHandler, etcdservergw.RegisterWatchHandler, etcdservergw.RegisterLeaseHandler, etcdservergw.RegisterClusterHandler, etcdservergw.RegisterMaintenanceHandler, etcdservergw.RegisterAuthHandler, v3lockgw.RegisterLockHandler, v3electiongw.RegisterElectionHandler}
	for _, h := range handlers {
		if err := h(ctx, gwmux, conn); err != nil {
			return nil, err
		}
	}
	go func() {
		<-ctx.Done()
		if cerr := conn.Close(); cerr != nil {
			if sctx.lg != nil {
				sctx.lg.Warn("failed to close connection", zap.String("address", sctx.l.Addr().String()), zap.Error(cerr))
			} else {
				plog.Warningf("failed to close conn to %s: %v", sctx.l.Addr().String(), cerr)
			}
		}
	}()
	return gwmux, nil
}
func (sctx *serveCtx) createMux(gwmux *gw.ServeMux, handler http.Handler) *http.ServeMux {
	_logClusterCodePath()
	defer _logClusterCodePath()
	httpmux := http.NewServeMux()
	for path, h := range sctx.userHandlers {
		httpmux.Handle(path, h)
	}
	if gwmux != nil {
		httpmux.Handle("/v3/", wsproxy.WebsocketProxy(gwmux, wsproxy.WithRequestMutator(func(_ *http.Request, outgoing *http.Request) *http.Request {
			outgoing.Method = "POST"
			return outgoing
		})))
	}
	if handler != nil {
		httpmux.Handle("/", handler)
	}
	return httpmux
}
func createAccessController(lg *zap.Logger, s *etcdserver.EtcdServer, mux *http.ServeMux) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &accessController{lg: lg, s: s, mux: mux}
}

type accessController struct {
	lg	*zap.Logger
	s	*etcdserver.EtcdServer
	mux	*http.ServeMux
}

func (ac *accessController) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if req != nil && req.URL != nil && strings.HasPrefix(req.URL.Path, "/v3beta/") {
		req.URL.Path = strings.Replace(req.URL.Path, "/v3beta/", "/v3/", 1)
	}
	if req.TLS == nil {
		host := httputil.GetHostname(req)
		if !ac.s.AccessController.IsHostWhitelisted(host) {
			if ac.lg != nil {
				ac.lg.Warn("rejecting HTTP request to prevent DNS rebinding attacks", zap.String("host", host))
			} else {
				plog.Warningf("rejecting HTTP request from %q to prevent DNS rebinding attacks", host)
			}
			http.Error(rw, errCVE20185702(host), 421)
			return
		}
	} else if ac.s.Cfg.ClientCertAuthEnabled && ac.s.Cfg.EnableGRPCGateway && ac.s.AuthStore().IsAuthEnabled() && strings.HasPrefix(req.URL.Path, "/v3/") {
		for _, chains := range req.TLS.VerifiedChains {
			if len(chains) < 1 {
				continue
			}
			if len(chains[0].Subject.CommonName) != 0 {
				http.Error(rw, "CommonName of client sending a request against gateway will be ignored and not used as expected", 400)
				return
			}
		}
	}
	if ac.s.AccessController.OriginAllowed("*") {
		addCORSHeader(rw, "*")
	} else if origin := req.Header.Get("Origin"); ac.s.OriginAllowed(origin) {
		addCORSHeader(rw, origin)
	}
	if req.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusOK)
		return
	}
	ac.mux.ServeHTTP(rw, req)
}
func addCORSHeader(w http.ResponseWriter, origin string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Add("Access-Control-Allow-Origin", origin)
	w.Header().Add("Access-Control-Allow-Headers", "accept, content-type, authorization")
}
func errCVE20185702(host string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf(`
etcd received your request, but the Host header was unrecognized.

To fix this, choose one of the following options:
- Enable TLS, then any HTTPS request will be allowed.
- Add the hostname you want to use to the whitelist in settings.
  - e.g. etcd --host-whitelist %q

This requirement has been added to help prevent "DNS Rebinding" attacks (CVE-2018-5702).
`, host)
}
func WrapCORS(cors map[string]struct{}, h http.Handler) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corsHandler{ac: &etcdserver.AccessController{CORS: cors}, h: h}
}

type corsHandler struct {
	ac	*etcdserver.AccessController
	h	http.Handler
}

func (ch *corsHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ch.ac.OriginAllowed("*") {
		addCORSHeader(rw, "*")
	} else if origin := req.Header.Get("Origin"); ch.ac.OriginAllowed(origin) {
		addCORSHeader(rw, origin)
	}
	if req.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusOK)
		return
	}
	ch.h.ServeHTTP(rw, req)
}
func (sctx *serveCtx) registerUserHandler(s string, h http.Handler) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if sctx.userHandlers[s] != nil {
		if sctx.lg != nil {
			sctx.lg.Warn("path is already registered by user handler", zap.String("path", s))
		} else {
			plog.Warningf("path %s already registered by user handler", s)
		}
		return
	}
	sctx.userHandlers[s] = h
}
func (sctx *serveCtx) registerPprof() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for p, h := range debugutil.PProfHandlers() {
		sctx.registerUserHandler(p, h)
	}
}
func (sctx *serveCtx) registerTrace() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	reqf := func(w http.ResponseWriter, r *http.Request) {
		trace.Render(w, r, true)
	}
	sctx.registerUserHandler("/debug/requests", http.HandlerFunc(reqf))
	evf := func(w http.ResponseWriter, r *http.Request) {
		trace.RenderEvents(w, r, true)
	}
	sctx.registerUserHandler("/debug/events", http.HandlerFunc(evf))
}
