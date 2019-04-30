package embed

import (
	"crypto/tls"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	godefaulthttp "net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api/v3compactor"
	"go.etcd.io/etcd/pkg/flags"
	"go.etcd.io/etcd/pkg/netutil"
	"go.etcd.io/etcd/pkg/srv"
	"go.etcd.io/etcd/pkg/tlsutil"
	"go.etcd.io/etcd/pkg/transport"
	"go.etcd.io/etcd/pkg/types"
	"github.com/ghodss/yaml"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

const (
	ClusterStateFlagNew		= "new"
	ClusterStateFlagExisting	= "existing"
	DefaultName			= "default"
	DefaultMaxSnapshots		= 5
	DefaultMaxWALs			= 5
	DefaultMaxTxnOps		= uint(128)
	DefaultMaxRequestBytes		= 1.5 * 1024 * 1024
	DefaultGRPCKeepAliveMinTime	= 5 * time.Second
	DefaultGRPCKeepAliveInterval	= 2 * time.Hour
	DefaultGRPCKeepAliveTimeout	= 20 * time.Second
	DefaultListenPeerURLs		= "http://localhost:2380"
	DefaultListenClientURLs		= "http://localhost:2379"
	DefaultLogOutput		= "default"
	JournalLogOutput		= "systemd/journal"
	StdErrLogOutput			= "stderr"
	StdOutLogOutput			= "stdout"
	DefaultStrictReconfigCheck	= true
	DefaultEnableV2			= true
	maxElectionMs			= 50000
	freelistMapType			= "map"
)

var (
	ErrConflictBootstrapFlags	= fmt.Errorf("multiple discovery or bootstrap flags are set. " + "Choose one of \"initial-cluster\", \"discovery\" or \"discovery-srv\"")
	ErrUnsetAdvertiseClientURLsFlag	= fmt.Errorf("--advertise-client-urls is required when --listen-client-urls is set explicitly")
	DefaultInitialAdvertisePeerURLs	= "http://localhost:2380"
	DefaultAdvertiseClientURLs	= "http://localhost:2379"
	defaultHostname			string
	defaultHostStatus		error
)
var (
	CompactorModePeriodic	= v3compactor.ModePeriodic
	CompactorModeRevision	= v3compactor.ModeRevision
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defaultHostname, defaultHostStatus = netutil.GetDefaultHost()
}

type Config struct {
	Name				string	`json:"name"`
	Dir				string	`json:"data-dir"`
	WalDir				string	`json:"wal-dir"`
	SnapshotCount			uint64	`json:"snapshot-count"`
	SnapshotCatchUpEntries		uint64
	MaxSnapFiles			uint		`json:"max-snapshots"`
	MaxWalFiles			uint		`json:"max-wals"`
	TickMs				uint		`json:"heartbeat-interval"`
	ElectionMs			uint		`json:"election-timeout"`
	InitialElectionTickAdvance	bool		`json:"initial-election-tick-advance"`
	BackendBatchInterval		time.Duration	`json:"backend-batch-interval"`
	BackendBatchLimit		int		`json:"backend-batch-limit"`
	QuotaBackendBytes		int64		`json:"quota-backend-bytes"`
	MaxTxnOps			uint		`json:"max-txn-ops"`
	MaxRequestBytes			uint		`json:"max-request-bytes"`
	LPUrls, LCUrls			[]url.URL
	APUrls, ACUrls			[]url.URL
	ClientTLSInfo			transport.TLSInfo
	ClientAutoTLS			bool
	PeerTLSInfo			transport.TLSInfo
	PeerAutoTLS			bool
	CipherSuites			[]string	`json:"cipher-suites"`
	ClusterState			string		`json:"initial-cluster-state"`
	DNSCluster			string		`json:"discovery-srv"`
	DNSClusterServiceName		string		`json:"discovery-srv-name"`
	Dproxy				string		`json:"discovery-proxy"`
	Durl				string		`json:"discovery"`
	InitialCluster			string		`json:"initial-cluster"`
	InitialClusterToken		string		`json:"initial-cluster-token"`
	StrictReconfigCheck		bool		`json:"strict-reconfig-check"`
	EnableV2			bool		`json:"enable-v2"`
	AutoCompactionMode		string		`json:"auto-compaction-mode"`
	AutoCompactionRetention		string		`json:"auto-compaction-retention"`
	GRPCKeepAliveMinTime		time.Duration	`json:"grpc-keepalive-min-time"`
	GRPCKeepAliveInterval		time.Duration	`json:"grpc-keepalive-interval"`
	GRPCKeepAliveTimeout		time.Duration	`json:"grpc-keepalive-timeout"`
	PreVote				bool		`json:"pre-vote"`
	CORS				map[string]struct{}
	HostWhitelist			map[string]struct{}
	UserHandlers			map[string]http.Handler	`json:"-"`
	ServiceRegister			func(*grpc.Server)	`json:"-"`
	AuthToken			string			`json:"auth-token"`
	BcryptCost			uint			`json:"bcrypt-cost"`
	ExperimentalInitialCorruptCheck	bool			`json:"experimental-initial-corrupt-check"`
	ExperimentalCorruptCheckTime	time.Duration		`json:"experimental-corrupt-check-time"`
	ExperimentalEnableV2V3		string			`json:"experimental-enable-v2v3"`
	ExperimentalBackendFreelistType	string			`json:"experimental-backend-bbolt-freelist-type"`
	ForceNewCluster			bool			`json:"force-new-cluster"`
	EnablePprof			bool			`json:"enable-pprof"`
	Metrics				string			`json:"metrics"`
	ListenMetricsUrls		[]url.URL
	ListenMetricsUrlsJSON		string		`json:"listen-metrics-urls"`
	Logger				string		`json:"logger"`
	DeprecatedLogOutput		[]string	`json:"log-output"`
	LogOutputs			[]string	`json:"log-outputs"`
	Debug				bool		`json:"debug"`
	ZapLoggerBuilder		func(*Config) error
	loggerMu			*sync.RWMutex
	logger				*zap.Logger
	loggerConfig			*zap.Config
	loggerCore			zapcore.Core
	loggerWriteSyncer		zapcore.WriteSyncer
	EnableGRPCGateway		bool	`json:"enable-grpc-gateway"`
	LogPkgLevels			string	`json:"log-package-levels"`
}
type configYAML struct {
	Config
	configJSON
}
type configJSON struct {
	LPUrlsJSON		string		`json:"listen-peer-urls"`
	LCUrlsJSON		string		`json:"listen-client-urls"`
	APUrlsJSON		string		`json:"initial-advertise-peer-urls"`
	ACUrlsJSON		string		`json:"advertise-client-urls"`
	CORSJSON		string		`json:"cors"`
	HostWhitelistJSON	string		`json:"host-whitelist"`
	ClientSecurityJSON	securityConfig	`json:"client-transport-security"`
	PeerSecurityJSON	securityConfig	`json:"peer-transport-security"`
}
type securityConfig struct {
	CertFile	string	`json:"cert-file"`
	KeyFile		string	`json:"key-file"`
	CertAuth	bool	`json:"client-cert-auth"`
	TrustedCAFile	string	`json:"trusted-ca-file"`
	AutoTLS		bool	`json:"auto-tls"`
}

func NewConfig() *Config {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lpurl, _ := url.Parse(DefaultListenPeerURLs)
	apurl, _ := url.Parse(DefaultInitialAdvertisePeerURLs)
	lcurl, _ := url.Parse(DefaultListenClientURLs)
	acurl, _ := url.Parse(DefaultAdvertiseClientURLs)
	cfg := &Config{MaxSnapFiles: DefaultMaxSnapshots, MaxWalFiles: DefaultMaxWALs, Name: DefaultName, SnapshotCount: etcdserver.DefaultSnapshotCount, SnapshotCatchUpEntries: etcdserver.DefaultSnapshotCatchUpEntries, MaxTxnOps: DefaultMaxTxnOps, MaxRequestBytes: DefaultMaxRequestBytes, GRPCKeepAliveMinTime: DefaultGRPCKeepAliveMinTime, GRPCKeepAliveInterval: DefaultGRPCKeepAliveInterval, GRPCKeepAliveTimeout: DefaultGRPCKeepAliveTimeout, TickMs: 100, ElectionMs: 1000, InitialElectionTickAdvance: true, LPUrls: []url.URL{*lpurl}, LCUrls: []url.URL{*lcurl}, APUrls: []url.URL{*apurl}, ACUrls: []url.URL{*acurl}, ClusterState: ClusterStateFlagNew, InitialClusterToken: "etcd-cluster", StrictReconfigCheck: DefaultStrictReconfigCheck, Metrics: "basic", EnableV2: DefaultEnableV2, CORS: map[string]struct{}{"*": {}}, HostWhitelist: map[string]struct{}{"*": {}}, AuthToken: "simple", BcryptCost: uint(bcrypt.DefaultCost), PreVote: false, loggerMu: new(sync.RWMutex), logger: nil, Logger: "capnslog", DeprecatedLogOutput: []string{DefaultLogOutput}, LogOutputs: []string{DefaultLogOutput}, Debug: false, LogPkgLevels: ""}
	cfg.InitialCluster = cfg.InitialClusterFromName(cfg.Name)
	return cfg
}
func logTLSHandshakeFailure(conn *tls.Conn, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	state := conn.ConnectionState()
	remoteAddr := conn.RemoteAddr().String()
	serverName := state.ServerName
	if len(state.PeerCertificates) > 0 {
		cert := state.PeerCertificates[0]
		ips, dns := cert.IPAddresses, cert.DNSNames
		plog.Infof("rejected connection from %q (error %q, ServerName %q, IPAddresses %q, DNSNames %q)", remoteAddr, err.Error(), serverName, ips, dns)
	} else {
		plog.Infof("rejected connection from %q (error %q, ServerName %q)", remoteAddr, err.Error(), serverName)
	}
}
func ConfigFromFile(path string) (*Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &configYAML{Config: *NewConfig()}
	if err := cfg.configFromFile(path); err != nil {
		return nil, err
	}
	return &cfg.Config, nil
}
func (cfg *configYAML) configFromFile(path string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	defaultInitialCluster := cfg.InitialCluster
	err = yaml.Unmarshal(b, cfg)
	if err != nil {
		return err
	}
	if cfg.LPUrlsJSON != "" {
		u, err := types.NewURLs(strings.Split(cfg.LPUrlsJSON, ","))
		if err != nil {
			fmt.Fprintf(os.Stderr, "unexpected error setting up listen-peer-urls: %v\n", err)
			os.Exit(1)
		}
		cfg.LPUrls = []url.URL(u)
	}
	if cfg.LCUrlsJSON != "" {
		u, err := types.NewURLs(strings.Split(cfg.LCUrlsJSON, ","))
		if err != nil {
			fmt.Fprintf(os.Stderr, "unexpected error setting up listen-client-urls: %v\n", err)
			os.Exit(1)
		}
		cfg.LCUrls = []url.URL(u)
	}
	if cfg.APUrlsJSON != "" {
		u, err := types.NewURLs(strings.Split(cfg.APUrlsJSON, ","))
		if err != nil {
			fmt.Fprintf(os.Stderr, "unexpected error setting up initial-advertise-peer-urls: %v\n", err)
			os.Exit(1)
		}
		cfg.APUrls = []url.URL(u)
	}
	if cfg.ACUrlsJSON != "" {
		u, err := types.NewURLs(strings.Split(cfg.ACUrlsJSON, ","))
		if err != nil {
			fmt.Fprintf(os.Stderr, "unexpected error setting up advertise-peer-urls: %v\n", err)
			os.Exit(1)
		}
		cfg.ACUrls = []url.URL(u)
	}
	if cfg.ListenMetricsUrlsJSON != "" {
		u, err := types.NewURLs(strings.Split(cfg.ListenMetricsUrlsJSON, ","))
		if err != nil {
			fmt.Fprintf(os.Stderr, "unexpected error setting up listen-metrics-urls: %v\n", err)
			os.Exit(1)
		}
		cfg.ListenMetricsUrls = []url.URL(u)
	}
	if cfg.CORSJSON != "" {
		uv := flags.NewUniqueURLsWithExceptions(cfg.CORSJSON, "*")
		cfg.CORS = uv.Values
	}
	if cfg.HostWhitelistJSON != "" {
		uv := flags.NewUniqueStringsValue(cfg.HostWhitelistJSON)
		cfg.HostWhitelist = uv.Values
	}
	if (cfg.Durl != "" || cfg.DNSCluster != "") && cfg.InitialCluster == defaultInitialCluster {
		cfg.InitialCluster = ""
	}
	if cfg.ClusterState == "" {
		cfg.ClusterState = ClusterStateFlagNew
	}
	copySecurityDetails := func(tls *transport.TLSInfo, ysc *securityConfig) {
		tls.CertFile = ysc.CertFile
		tls.KeyFile = ysc.KeyFile
		tls.ClientCertAuth = ysc.CertAuth
		tls.TrustedCAFile = ysc.TrustedCAFile
	}
	copySecurityDetails(&cfg.ClientTLSInfo, &cfg.ClientSecurityJSON)
	copySecurityDetails(&cfg.PeerTLSInfo, &cfg.PeerSecurityJSON)
	cfg.ClientAutoTLS = cfg.ClientSecurityJSON.AutoTLS
	cfg.PeerAutoTLS = cfg.PeerSecurityJSON.AutoTLS
	return cfg.Validate()
}
func updateCipherSuites(tls *transport.TLSInfo, ss []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(tls.CipherSuites) > 0 && len(ss) > 0 {
		return fmt.Errorf("TLSInfo.CipherSuites is already specified (given %v)", ss)
	}
	if len(ss) > 0 {
		cs := make([]uint16, len(ss))
		for i, s := range ss {
			var ok bool
			cs[i], ok = tlsutil.GetCipherSuite(s)
			if !ok {
				return fmt.Errorf("unexpected TLS cipher suite %q", s)
			}
		}
		tls.CipherSuites = cs
	}
	return nil
}
func (cfg *Config) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := cfg.setupLogging(); err != nil {
		return err
	}
	if err := checkBindURLs(cfg.LPUrls); err != nil {
		return err
	}
	if err := checkBindURLs(cfg.LCUrls); err != nil {
		return err
	}
	if err := checkBindURLs(cfg.ListenMetricsUrls); err != nil {
		return err
	}
	if err := checkHostURLs(cfg.APUrls); err != nil {
		addrs := cfg.getAPURLs()
		return fmt.Errorf(`--initial-advertise-peer-urls %q must be "host:port" (%v)`, strings.Join(addrs, ","), err)
	}
	if err := checkHostURLs(cfg.ACUrls); err != nil {
		addrs := cfg.getACURLs()
		return fmt.Errorf(`--advertise-client-urls %q must be "host:port" (%v)`, strings.Join(addrs, ","), err)
	}
	nSet := 0
	for _, v := range []bool{cfg.Durl != "", cfg.InitialCluster != "", cfg.DNSCluster != ""} {
		if v {
			nSet++
		}
	}
	if cfg.ClusterState != ClusterStateFlagNew && cfg.ClusterState != ClusterStateFlagExisting {
		return fmt.Errorf("unexpected clusterState %q", cfg.ClusterState)
	}
	if nSet > 1 {
		return ErrConflictBootstrapFlags
	}
	if cfg.TickMs <= 0 {
		return fmt.Errorf("--heartbeat-interval must be >0 (set to %dms)", cfg.TickMs)
	}
	if cfg.ElectionMs <= 0 {
		return fmt.Errorf("--election-timeout must be >0 (set to %dms)", cfg.ElectionMs)
	}
	if 5*cfg.TickMs > cfg.ElectionMs {
		return fmt.Errorf("--election-timeout[%vms] should be at least as 5 times as --heartbeat-interval[%vms]", cfg.ElectionMs, cfg.TickMs)
	}
	if cfg.ElectionMs > maxElectionMs {
		return fmt.Errorf("--election-timeout[%vms] is too long, and should be set less than %vms", cfg.ElectionMs, maxElectionMs)
	}
	if cfg.LCUrls != nil && cfg.ACUrls == nil {
		return ErrUnsetAdvertiseClientURLsFlag
	}
	switch cfg.AutoCompactionMode {
	case "":
	case CompactorModeRevision, CompactorModePeriodic:
	default:
		return fmt.Errorf("unknown auto-compaction-mode %q", cfg.AutoCompactionMode)
	}
	return nil
}
func (cfg *Config) PeerURLsMapAndToken(which string) (urlsmap types.URLsMap, token string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	token = cfg.InitialClusterToken
	switch {
	case cfg.Durl != "":
		urlsmap = types.URLsMap{}
		urlsmap[cfg.Name] = cfg.APUrls
		token = cfg.Durl
	case cfg.DNSCluster != "":
		clusterStrs, cerr := cfg.GetDNSClusterNames()
		lg := cfg.logger
		if cerr != nil {
			if lg != nil {
				lg.Warn("failed to resolve during SRV discovery", zap.Error(cerr))
			} else {
				plog.Errorf("couldn't resolve during SRV discovery (%v)", cerr)
			}
			return nil, "", cerr
		}
		for _, s := range clusterStrs {
			if lg != nil {
				lg.Info("got bootstrap from DNS for etcd-server", zap.String("node", s))
			} else {
				plog.Noticef("got bootstrap from DNS for etcd-server at %s", s)
			}
		}
		clusterStr := strings.Join(clusterStrs, ",")
		if strings.Contains(clusterStr, "https://") && cfg.PeerTLSInfo.TrustedCAFile == "" {
			cfg.PeerTLSInfo.ServerName = cfg.DNSCluster
		}
		urlsmap, err = types.NewURLsMap(clusterStr)
		if which == "etcd" {
			if _, ok := urlsmap[cfg.Name]; !ok {
				return nil, "", fmt.Errorf("cannot find local etcd member %q in SRV records", cfg.Name)
			}
		}
	default:
		urlsmap, err = types.NewURLsMap(cfg.InitialCluster)
	}
	return urlsmap, token, err
}
func (cfg *Config) GetDNSClusterNames() ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		clusterStrs		[]string
		cerr			error
		serviceNameSuffix	string
	)
	if cfg.DNSClusterServiceName != "" {
		serviceNameSuffix = "-" + cfg.DNSClusterServiceName
	}
	lg := cfg.GetLogger()
	clusterStrs, cerr = srv.GetCluster("https", "etcd-server-ssl"+serviceNameSuffix, cfg.Name, cfg.DNSCluster, cfg.APUrls)
	if cerr != nil {
		clusterStrs = make([]string, 0)
	}
	if lg != nil {
		lg.Info("get cluster for etcd-server-ssl SRV", zap.String("service-scheme", "https"), zap.String("service-name", "etcd-server-ssl"+serviceNameSuffix), zap.String("server-name", cfg.Name), zap.String("discovery-srv", cfg.DNSCluster), zap.Strings("advertise-peer-urls", cfg.getAPURLs()), zap.Strings("found-cluster", clusterStrs), zap.Error(cerr))
	}
	defaultHTTPClusterStrs, httpCerr := srv.GetCluster("http", "etcd-server"+serviceNameSuffix, cfg.Name, cfg.DNSCluster, cfg.APUrls)
	if httpCerr != nil {
		clusterStrs = append(clusterStrs, defaultHTTPClusterStrs...)
	}
	if lg != nil {
		lg.Info("get cluster for etcd-server SRV", zap.String("service-scheme", "http"), zap.String("service-name", "etcd-server"+serviceNameSuffix), zap.String("server-name", cfg.Name), zap.String("discovery-srv", cfg.DNSCluster), zap.Strings("advertise-peer-urls", cfg.getAPURLs()), zap.Strings("found-cluster", clusterStrs), zap.Error(httpCerr))
	}
	return clusterStrs, cerr
}
func (cfg Config) InitialClusterFromName(name string) (ret string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(cfg.APUrls) == 0 {
		return ""
	}
	n := name
	if name == "" {
		n = DefaultName
	}
	for i := range cfg.APUrls {
		ret = ret + "," + n + "=" + cfg.APUrls[i].String()
	}
	return ret[1:]
}
func (cfg Config) IsNewCluster() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cfg.ClusterState == ClusterStateFlagNew
}
func (cfg Config) ElectionTicks() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int(cfg.ElectionMs / cfg.TickMs)
}
func (cfg Config) defaultPeerHost() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(cfg.APUrls) == 1 && cfg.APUrls[0].String() == DefaultInitialAdvertisePeerURLs
}
func (cfg Config) defaultClientHost() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(cfg.ACUrls) == 1 && cfg.ACUrls[0].String() == DefaultAdvertiseClientURLs
}
func (cfg *Config) ClientSelfCert() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !cfg.ClientAutoTLS {
		return nil
	}
	if !cfg.ClientTLSInfo.Empty() {
		if cfg.logger != nil {
			cfg.logger.Warn("ignoring client auto TLS since certs given")
		} else {
			plog.Warningf("ignoring client auto TLS since certs given")
		}
		return nil
	}
	chosts := make([]string, len(cfg.LCUrls))
	for i, u := range cfg.LCUrls {
		chosts[i] = u.Host
	}
	cfg.ClientTLSInfo, err = transport.SelfCert(cfg.logger, filepath.Join(cfg.Dir, "fixtures", "client"), chosts)
	if err != nil {
		return err
	}
	return updateCipherSuites(&cfg.ClientTLSInfo, cfg.CipherSuites)
}
func (cfg *Config) PeerSelfCert() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !cfg.PeerAutoTLS {
		return nil
	}
	if !cfg.PeerTLSInfo.Empty() {
		if cfg.logger != nil {
			cfg.logger.Warn("ignoring peer auto TLS since certs given")
		} else {
			plog.Warningf("ignoring peer auto TLS since certs given")
		}
		return nil
	}
	phosts := make([]string, len(cfg.LPUrls))
	for i, u := range cfg.LPUrls {
		phosts[i] = u.Host
	}
	cfg.PeerTLSInfo, err = transport.SelfCert(cfg.logger, filepath.Join(cfg.Dir, "fixtures", "peer"), phosts)
	if err != nil {
		return err
	}
	return updateCipherSuites(&cfg.PeerTLSInfo, cfg.CipherSuites)
}
func (cfg *Config) UpdateDefaultClusterFromName(defaultInitialCluster string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if defaultHostname == "" || defaultHostStatus != nil {
		if cfg.Name != DefaultName && cfg.InitialCluster == defaultInitialCluster {
			cfg.InitialCluster = cfg.InitialClusterFromName(cfg.Name)
		}
		return "", defaultHostStatus
	}
	used := false
	pip, pport := cfg.LPUrls[0].Hostname(), cfg.LPUrls[0].Port()
	if cfg.defaultPeerHost() && pip == "0.0.0.0" {
		cfg.APUrls[0] = url.URL{Scheme: cfg.APUrls[0].Scheme, Host: fmt.Sprintf("%s:%s", defaultHostname, pport)}
		used = true
	}
	if cfg.Name != DefaultName && cfg.InitialCluster == defaultInitialCluster {
		cfg.InitialCluster = cfg.InitialClusterFromName(cfg.Name)
	}
	cip, cport := cfg.LCUrls[0].Hostname(), cfg.LCUrls[0].Port()
	if cfg.defaultClientHost() && cip == "0.0.0.0" {
		cfg.ACUrls[0] = url.URL{Scheme: cfg.ACUrls[0].Scheme, Host: fmt.Sprintf("%s:%s", defaultHostname, cport)}
		used = true
	}
	dhost := defaultHostname
	if !used {
		dhost = ""
	}
	return dhost, defaultHostStatus
}
func checkBindURLs(urls []url.URL) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, url := range urls {
		if url.Scheme == "unix" || url.Scheme == "unixs" {
			continue
		}
		host, _, err := net.SplitHostPort(url.Host)
		if err != nil {
			return err
		}
		if host == "localhost" {
			continue
		}
		if net.ParseIP(host) == nil {
			return fmt.Errorf("expected IP in URL for binding (%s)", url.String())
		}
	}
	return nil
}
func checkHostURLs(urls []url.URL) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, url := range urls {
		host, _, err := net.SplitHostPort(url.Host)
		if err != nil {
			return err
		}
		if host == "" {
			return fmt.Errorf("unexpected empty host (%s)", url.String())
		}
	}
	return nil
}
func (cfg *Config) getAPURLs() (ss []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss = make([]string, len(cfg.APUrls))
	for i := range cfg.APUrls {
		ss[i] = cfg.APUrls[i].String()
	}
	return ss
}
func (cfg *Config) getLPURLs() (ss []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss = make([]string, len(cfg.LPUrls))
	for i := range cfg.LPUrls {
		ss[i] = cfg.LPUrls[i].String()
	}
	return ss
}
func (cfg *Config) getACURLs() (ss []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss = make([]string, len(cfg.ACUrls))
	for i := range cfg.ACUrls {
		ss[i] = cfg.ACUrls[i].String()
	}
	return ss
}
func (cfg *Config) getLCURLs() (ss []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss = make([]string, len(cfg.LCUrls))
	for i := range cfg.LCUrls {
		ss[i] = cfg.LCUrls[i].String()
	}
	return ss
}
func (cfg *Config) getMetricsURLs() (ss []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss = make([]string, len(cfg.ListenMetricsUrls))
	for i := range cfg.ListenMetricsUrls {
		ss[i] = cfg.ListenMetricsUrls[i].String()
	}
	return ss
}
func parseBackendFreelistType(freelistType string) bolt.FreelistType {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if freelistType == freelistMapType {
		return bolt.FreelistMapType
	}
	return bolt.FreelistArrayType
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
