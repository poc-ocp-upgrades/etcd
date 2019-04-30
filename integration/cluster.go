package integration

import (
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"go.etcd.io/etcd/client"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/embed"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api/etcdhttp"
	"go.etcd.io/etcd/etcdserver/api/rafthttp"
	"go.etcd.io/etcd/etcdserver/api/v2http"
	"go.etcd.io/etcd/etcdserver/api/v3client"
	"go.etcd.io/etcd/etcdserver/api/v3election"
	epb "go.etcd.io/etcd/etcdserver/api/v3election/v3electionpb"
	"go.etcd.io/etcd/etcdserver/api/v3lock"
	lockpb "go.etcd.io/etcd/etcdserver/api/v3lock/v3lockpb"
	"go.etcd.io/etcd/etcdserver/api/v3rpc"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/pkg/logutil"
	"go.etcd.io/etcd/pkg/testutil"
	"go.etcd.io/etcd/pkg/tlsutil"
	"go.etcd.io/etcd/pkg/transport"
	"go.etcd.io/etcd/pkg/types"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/keepalive"
)

const (
	RequestWaitTimeout	= 3 * time.Second
	tickDuration		= 10 * time.Millisecond
	requestTimeout		= 20 * time.Second
	clusterName		= "etcd"
	basePort		= 21000
	URLScheme		= "unix"
	URLSchemeTLS		= "unixs"
)

var (
	electionTicks		= 10
	localListenCount	= int64(0)
	testTLSInfo		= transport.TLSInfo{KeyFile: "./fixtures/server.key.insecure", CertFile: "./fixtures/server.crt", TrustedCAFile: "./fixtures/ca.crt", ClientCertAuth: true}
	testTLSInfoIP		= transport.TLSInfo{KeyFile: "./fixtures/server-ip.key.insecure", CertFile: "./fixtures/server-ip.crt", TrustedCAFile: "./fixtures/ca.crt", ClientCertAuth: true}
	testTLSInfoExpired	= transport.TLSInfo{KeyFile: "./fixtures-expired/server.key.insecure", CertFile: "./fixtures-expired/server.crt", TrustedCAFile: "./fixtures-expired/ca.crt", ClientCertAuth: true}
	testTLSInfoExpiredIP	= transport.TLSInfo{KeyFile: "./fixtures-expired/server-ip.key.insecure", CertFile: "./fixtures-expired/server-ip.crt", TrustedCAFile: "./fixtures-expired/ca.crt", ClientCertAuth: true}
	defaultTokenJWT		= "jwt,pub-key=./fixtures/server.crt,priv-key=./fixtures/server.key.insecure,sign-method=RS256,ttl=1s"
	lg			= zap.NewNop()
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if os.Getenv("CLUSTER_DEBUG") != "" {
		lg, _ = zap.NewProduction()
	}
}

type ClusterConfig struct {
	Size				int
	PeerTLS				*transport.TLSInfo
	ClientTLS			*transport.TLSInfo
	DiscoveryURL			string
	AuthToken			string
	UseGRPC				bool
	QuotaBackendBytes		int64
	MaxTxnOps			uint
	MaxRequestBytes			uint
	SnapshotCount			uint64
	SnapshotCatchUpEntries		uint64
	GRPCKeepAliveMinTime		time.Duration
	GRPCKeepAliveInterval		time.Duration
	GRPCKeepAliveTimeout		time.Duration
	SkipCreatingClient		bool
	ClientMaxCallSendMsgSize	int
	ClientMaxCallRecvMsgSize	int
	UseIP				bool
	LeaseCheckpointInterval		time.Duration
}
type cluster struct {
	cfg	*ClusterConfig
	Members	[]*member
}

func schemeFromTLSInfo(tls *transport.TLSInfo) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if tls == nil {
		return URLScheme
	}
	return URLSchemeTLS
}
func (c *cluster) fillClusterForMembers() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.cfg.DiscoveryURL != "" {
		return nil
	}
	addrs := make([]string, 0)
	for _, m := range c.Members {
		scheme := schemeFromTLSInfo(m.PeerTLSInfo)
		for _, l := range m.PeerListeners {
			addrs = append(addrs, fmt.Sprintf("%s=%s://%s", m.Name, scheme, l.Addr().String()))
		}
	}
	clusterStr := strings.Join(addrs, ",")
	var err error
	for _, m := range c.Members {
		m.InitialPeerURLsMap, err = types.NewURLsMap(clusterStr)
		if err != nil {
			return err
		}
	}
	return nil
}
func newCluster(t testing.TB, cfg *ClusterConfig) *cluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &cluster{cfg: cfg}
	ms := make([]*member, cfg.Size)
	for i := 0; i < cfg.Size; i++ {
		ms[i] = c.mustNewMember(t)
	}
	c.Members = ms
	if err := c.fillClusterForMembers(); err != nil {
		t.Fatal(err)
	}
	return c
}
func NewCluster(t testing.TB, size int) *cluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newCluster(t, &ClusterConfig{Size: size})
}
func NewClusterByConfig(t testing.TB, cfg *ClusterConfig) *cluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newCluster(t, cfg)
}
func (c *cluster) Launch(t testing.TB) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errc := make(chan error)
	for _, m := range c.Members {
		go func(m *member) {
			errc <- m.Launch()
		}(m)
	}
	for range c.Members {
		if err := <-errc; err != nil {
			t.Fatalf("error setting up member: %v", err)
		}
	}
	c.waitMembersMatch(t, c.HTTPMembers())
	c.waitVersion()
}
func (c *cluster) URL(i int) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Members[i].ClientURLs[0].String()
}
func (c *cluster) URLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return getMembersURLs(c.Members)
}
func getMembersURLs(members []*member) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	urls := make([]string, 0)
	for _, m := range members {
		select {
		case <-m.s.StopNotify():
			continue
		default:
		}
		for _, u := range m.ClientURLs {
			urls = append(urls, u.String())
		}
	}
	return urls
}
func (c *cluster) HTTPMembers() []client.Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ms := []client.Member{}
	for _, m := range c.Members {
		pScheme := schemeFromTLSInfo(m.PeerTLSInfo)
		cScheme := schemeFromTLSInfo(m.ClientTLSInfo)
		cm := client.Member{Name: m.Name}
		for _, ln := range m.PeerListeners {
			cm.PeerURLs = append(cm.PeerURLs, pScheme+"://"+ln.Addr().String())
		}
		for _, ln := range m.ClientListeners {
			cm.ClientURLs = append(cm.ClientURLs, cScheme+"://"+ln.Addr().String())
		}
		ms = append(ms, cm)
	}
	return ms
}
func (c *cluster) mustNewMember(t testing.TB) *member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := mustNewMember(t, memberConfig{name: c.name(rand.Int()), authToken: c.cfg.AuthToken, peerTLS: c.cfg.PeerTLS, clientTLS: c.cfg.ClientTLS, quotaBackendBytes: c.cfg.QuotaBackendBytes, maxTxnOps: c.cfg.MaxTxnOps, maxRequestBytes: c.cfg.MaxRequestBytes, snapshotCount: c.cfg.SnapshotCount, snapshotCatchUpEntries: c.cfg.SnapshotCatchUpEntries, grpcKeepAliveMinTime: c.cfg.GRPCKeepAliveMinTime, grpcKeepAliveInterval: c.cfg.GRPCKeepAliveInterval, grpcKeepAliveTimeout: c.cfg.GRPCKeepAliveTimeout, clientMaxCallSendMsgSize: c.cfg.ClientMaxCallSendMsgSize, clientMaxCallRecvMsgSize: c.cfg.ClientMaxCallRecvMsgSize, useIP: c.cfg.UseIP, leaseCheckpointInterval: c.cfg.LeaseCheckpointInterval})
	m.DiscoveryURL = c.cfg.DiscoveryURL
	if c.cfg.UseGRPC {
		if err := m.listenGRPC(); err != nil {
			t.Fatal(err)
		}
	}
	return m
}
func (c *cluster) addMember(t testing.TB) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := c.mustNewMember(t)
	scheme := schemeFromTLSInfo(c.cfg.PeerTLS)
	var err error
	for i := 0; i < len(c.Members); i++ {
		clientURL := c.URL(i)
		peerURL := scheme + "://" + m.PeerListeners[0].Addr().String()
		if err = c.addMemberByURL(t, clientURL, peerURL); err == nil {
			break
		}
	}
	if err != nil {
		t.Fatalf("add member failed on all members error: %v", err)
	}
	m.InitialPeerURLsMap = types.URLsMap{}
	for _, mm := range c.Members {
		m.InitialPeerURLsMap[mm.Name] = mm.PeerURLs
	}
	m.InitialPeerURLsMap[m.Name] = m.PeerURLs
	m.NewCluster = false
	if err := m.Launch(); err != nil {
		t.Fatal(err)
	}
	c.Members = append(c.Members, m)
	c.waitMembersMatch(t, c.HTTPMembers())
}
func (c *cluster) addMemberByURL(t testing.TB, clientURL, peerURL string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := MustNewHTTPClient(t, []string{clientURL}, c.cfg.ClientTLS)
	ma := client.NewMembersAPI(cc)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err := ma.Add(ctx, peerURL)
	cancel()
	if err != nil {
		return err
	}
	members := append(c.HTTPMembers(), client.Member{PeerURLs: []string{peerURL}, ClientURLs: []string{}})
	c.waitMembersMatch(t, members)
	return nil
}
func (c *cluster) AddMember(t testing.TB) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.addMember(t)
}
func (c *cluster) RemoveMember(t testing.TB, id uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.removeMember(t, id); err != nil {
		t.Fatal(err)
	}
}
func (c *cluster) removeMember(t testing.TB, id uint64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := MustNewHTTPClient(t, c.URLs(), c.cfg.ClientTLS)
	ma := client.NewMembersAPI(cc)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	err := ma.Remove(ctx, types.ID(id).String())
	cancel()
	if err != nil {
		return err
	}
	newMembers := make([]*member, 0)
	for _, m := range c.Members {
		if uint64(m.s.ID()) != id {
			newMembers = append(newMembers, m)
		} else {
			select {
			case <-m.s.StopNotify():
				m.Terminate(t)
			case <-time.After(time.Second + time.Duration(electionTicks)*tickDuration + time.Second + rafthttp.ConnWriteTimeout):
				t.Fatalf("failed to remove member %s in time", m.s.ID())
			}
		}
	}
	c.Members = newMembers
	c.waitMembersMatch(t, c.HTTPMembers())
	return nil
}
func (c *cluster) Terminate(t testing.TB) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wg sync.WaitGroup
	wg.Add(len(c.Members))
	for _, m := range c.Members {
		go func(mm *member) {
			defer wg.Done()
			mm.Terminate(t)
		}(m)
	}
	wg.Wait()
}
func (c *cluster) waitMembersMatch(t testing.TB, membs []client.Member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, u := range c.URLs() {
		cc := MustNewHTTPClient(t, []string{u}, c.cfg.ClientTLS)
		ma := client.NewMembersAPI(cc)
		for {
			ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
			ms, err := ma.List(ctx)
			cancel()
			if err == nil && isMembersEqual(ms, membs) {
				break
			}
			time.Sleep(tickDuration)
		}
	}
}
func (c *cluster) WaitLeader(t testing.TB) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.waitLeader(t, c.Members)
}
func (c *cluster) waitLeader(t testing.TB, membs []*member) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	possibleLead := make(map[uint64]bool)
	var lead uint64
	for _, m := range membs {
		possibleLead[uint64(m.s.ID())] = true
	}
	cc := MustNewHTTPClient(t, getMembersURLs(membs), nil)
	kapi := client.NewKeysAPI(cc)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*tickDuration+time.Second)
		_, err := kapi.Get(ctx, "0", &client.GetOptions{Quorum: true})
		cancel()
		if err == nil || strings.Contains(err.Error(), "Key not found") {
			break
		}
	}
	for lead == 0 || !possibleLead[lead] {
		lead = 0
		for _, m := range membs {
			select {
			case <-m.s.StopNotify():
				continue
			default:
			}
			if lead != 0 && lead != m.s.Lead() {
				lead = 0
				time.Sleep(10 * tickDuration)
				break
			}
			lead = m.s.Lead()
		}
	}
	for i, m := range membs {
		if uint64(m.s.ID()) == lead {
			return i
		}
	}
	return -1
}
func (c *cluster) WaitNoLeader() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.waitNoLeader(c.Members)
}
func (c *cluster) waitNoLeader(membs []*member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	noLeader := false
	for !noLeader {
		noLeader = true
		for _, m := range membs {
			select {
			case <-m.s.StopNotify():
				continue
			default:
			}
			if m.s.Lead() != 0 {
				noLeader = false
				time.Sleep(10 * tickDuration)
				break
			}
		}
	}
}
func (c *cluster) waitVersion() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range c.Members {
		for {
			if m.s.ClusterVersion() != nil {
				break
			}
			time.Sleep(tickDuration)
		}
	}
}
func (c *cluster) name(i int) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprint(i)
}
func isMembersEqual(membs []client.Member, wmembs []client.Member) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sort.Sort(SortableMemberSliceByPeerURLs(membs))
	sort.Sort(SortableMemberSliceByPeerURLs(wmembs))
	for i := range membs {
		membs[i].ID = ""
	}
	return reflect.DeepEqual(membs, wmembs)
}
func newLocalListener(t testing.TB) net.Listener {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := atomic.AddInt64(&localListenCount, 1)
	addr := fmt.Sprintf("127.0.0.1:%05d%05d", c+basePort, os.Getpid())
	return NewListenerWithAddr(t, addr)
}
func NewListenerWithAddr(t testing.TB, addr string) net.Listener {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l, err := transport.NewUnixListener(addr)
	if err != nil {
		t.Fatal(err)
	}
	return l
}

type member struct {
	etcdserver.ServerConfig
	PeerListeners, ClientListeners	[]net.Listener
	grpcListener			net.Listener
	PeerTLSInfo			*transport.TLSInfo
	ClientTLSInfo			*transport.TLSInfo
	DialOptions			[]grpc.DialOption
	raftHandler			*testutil.PauseableHandler
	s				*etcdserver.EtcdServer
	serverClosers			[]func()
	grpcServerOpts			[]grpc.ServerOption
	grpcServer			*grpc.Server
	grpcServerPeer			*grpc.Server
	grpcAddr			string
	grpcBridge			*bridge
	serverClient			*clientv3.Client
	keepDataDirTerminate		bool
	clientMaxCallSendMsgSize	int
	clientMaxCallRecvMsgSize	int
	useIP				bool
}

func (m *member) GRPCAddr() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.grpcAddr
}

type memberConfig struct {
	name				string
	peerTLS				*transport.TLSInfo
	clientTLS			*transport.TLSInfo
	authToken			string
	quotaBackendBytes		int64
	maxTxnOps			uint
	maxRequestBytes			uint
	snapshotCount			uint64
	snapshotCatchUpEntries		uint64
	grpcKeepAliveMinTime		time.Duration
	grpcKeepAliveInterval		time.Duration
	grpcKeepAliveTimeout		time.Duration
	clientMaxCallSendMsgSize	int
	clientMaxCallRecvMsgSize	int
	useIP				bool
	leaseCheckpointInterval		time.Duration
}

func mustNewMember(t testing.TB, mcfg memberConfig) *member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	m := &member{}
	peerScheme := schemeFromTLSInfo(mcfg.peerTLS)
	clientScheme := schemeFromTLSInfo(mcfg.clientTLS)
	pln := newLocalListener(t)
	m.PeerListeners = []net.Listener{pln}
	m.PeerURLs, err = types.NewURLs([]string{peerScheme + "://" + pln.Addr().String()})
	if err != nil {
		t.Fatal(err)
	}
	m.PeerTLSInfo = mcfg.peerTLS
	cln := newLocalListener(t)
	m.ClientListeners = []net.Listener{cln}
	m.ClientURLs, err = types.NewURLs([]string{clientScheme + "://" + cln.Addr().String()})
	if err != nil {
		t.Fatal(err)
	}
	m.ClientTLSInfo = mcfg.clientTLS
	m.Name = mcfg.name
	m.DataDir, err = ioutil.TempDir(os.TempDir(), "etcd")
	if err != nil {
		t.Fatal(err)
	}
	clusterStr := fmt.Sprintf("%s=%s://%s", mcfg.name, peerScheme, pln.Addr().String())
	m.InitialPeerURLsMap, err = types.NewURLsMap(clusterStr)
	if err != nil {
		t.Fatal(err)
	}
	m.InitialClusterToken = clusterName
	m.NewCluster = true
	m.BootstrapTimeout = 10 * time.Millisecond
	if m.PeerTLSInfo != nil {
		m.ServerConfig.PeerTLSInfo = *m.PeerTLSInfo
	}
	m.ElectionTicks = electionTicks
	m.InitialElectionTickAdvance = true
	m.TickMs = uint(tickDuration / time.Millisecond)
	m.QuotaBackendBytes = mcfg.quotaBackendBytes
	m.MaxTxnOps = mcfg.maxTxnOps
	if m.MaxTxnOps == 0 {
		m.MaxTxnOps = embed.DefaultMaxTxnOps
	}
	m.MaxRequestBytes = mcfg.maxRequestBytes
	if m.MaxRequestBytes == 0 {
		m.MaxRequestBytes = embed.DefaultMaxRequestBytes
	}
	m.SnapshotCount = etcdserver.DefaultSnapshotCount
	if mcfg.snapshotCount != 0 {
		m.SnapshotCount = mcfg.snapshotCount
	}
	m.SnapshotCatchUpEntries = etcdserver.DefaultSnapshotCatchUpEntries
	if mcfg.snapshotCatchUpEntries != 0 {
		m.SnapshotCatchUpEntries = mcfg.snapshotCatchUpEntries
	}
	m.AuthToken = "simple"
	if mcfg.authToken != "" {
		m.AuthToken = mcfg.authToken
	}
	m.BcryptCost = uint(bcrypt.MinCost)
	m.grpcServerOpts = []grpc.ServerOption{}
	if mcfg.grpcKeepAliveMinTime > time.Duration(0) {
		m.grpcServerOpts = append(m.grpcServerOpts, grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{MinTime: mcfg.grpcKeepAliveMinTime, PermitWithoutStream: false}))
	}
	if mcfg.grpcKeepAliveInterval > time.Duration(0) && mcfg.grpcKeepAliveTimeout > time.Duration(0) {
		m.grpcServerOpts = append(m.grpcServerOpts, grpc.KeepaliveParams(keepalive.ServerParameters{Time: mcfg.grpcKeepAliveInterval, Timeout: mcfg.grpcKeepAliveTimeout}))
	}
	m.clientMaxCallSendMsgSize = mcfg.clientMaxCallSendMsgSize
	m.clientMaxCallRecvMsgSize = mcfg.clientMaxCallRecvMsgSize
	m.useIP = mcfg.useIP
	m.LeaseCheckpointInterval = mcfg.leaseCheckpointInterval
	m.InitialCorruptCheck = true
	lcfg := logutil.DefaultZapLoggerConfig
	m.LoggerConfig = &lcfg
	m.LoggerConfig.OutputPaths = []string{"/dev/null"}
	m.LoggerConfig.ErrorOutputPaths = []string{"/dev/null"}
	if os.Getenv("CLUSTER_DEBUG") != "" {
		m.LoggerConfig.OutputPaths = []string{"stderr"}
		m.LoggerConfig.ErrorOutputPaths = []string{"stderr"}
	}
	m.Logger, err = m.LoggerConfig.Build()
	if err != nil {
		t.Fatal(err)
	}
	return m
}
func (m *member) listenGRPC() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.grpcAddr = "localhost:" + m.Name
	if m.useIP {
		m.grpcAddr = "127.0.0.1:" + m.Name
	}
	l, err := transport.NewUnixListener(m.grpcAddr)
	if err != nil {
		return fmt.Errorf("listen failed on grpc socket %s (%v)", m.grpcAddr, err)
	}
	m.grpcBridge, err = newBridge(m.grpcAddr)
	if err != nil {
		l.Close()
		return err
	}
	m.grpcAddr = schemeFromTLSInfo(m.ClientTLSInfo) + "://" + m.grpcBridge.inaddr
	m.grpcListener = l
	return nil
}
func (m *member) ElectionTimeout() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Duration(m.s.Cfg.ElectionTicks*int(m.s.Cfg.TickMs)) * time.Millisecond
}
func (m *member) ID() types.ID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.s.ID()
}
func (m *member) DropConnections() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.grpcBridge.Reset()
}
func (m *member) PauseConnections() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.grpcBridge.Pause()
}
func (m *member) UnpauseConnections() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.grpcBridge.Unpause()
}
func (m *member) Blackhole() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.grpcBridge.Blackhole()
}
func (m *member) Unblackhole() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.grpcBridge.Unblackhole()
}
func NewClientV3(m *member) (*clientv3.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.grpcAddr == "" {
		return nil, fmt.Errorf("member not configured for grpc")
	}
	cfg := clientv3.Config{Endpoints: []string{m.grpcAddr}, DialTimeout: 5 * time.Second, DialOptions: []grpc.DialOption{grpc.WithBlock()}, MaxCallSendMsgSize: m.clientMaxCallSendMsgSize, MaxCallRecvMsgSize: m.clientMaxCallRecvMsgSize}
	if m.ClientTLSInfo != nil {
		tls, err := m.ClientTLSInfo.ClientConfig()
		if err != nil {
			return nil, err
		}
		cfg.TLS = tls
	}
	if m.DialOptions != nil {
		cfg.DialOptions = append(cfg.DialOptions, m.DialOptions...)
	}
	return newClientV3(cfg)
}
func (m *member) Clone(t testing.TB) *member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mm := &member{}
	mm.ServerConfig = m.ServerConfig
	var err error
	clientURLStrs := m.ClientURLs.StringSlice()
	mm.ClientURLs, err = types.NewURLs(clientURLStrs)
	if err != nil {
		panic(err)
	}
	peerURLStrs := m.PeerURLs.StringSlice()
	mm.PeerURLs, err = types.NewURLs(peerURLStrs)
	if err != nil {
		panic(err)
	}
	clusterStr := m.InitialPeerURLsMap.String()
	mm.InitialPeerURLsMap, err = types.NewURLsMap(clusterStr)
	if err != nil {
		panic(err)
	}
	mm.InitialClusterToken = m.InitialClusterToken
	mm.ElectionTicks = m.ElectionTicks
	mm.PeerTLSInfo = m.PeerTLSInfo
	mm.ClientTLSInfo = m.ClientTLSInfo
	return mm
}
func (m *member) Launch() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg.Info("launching a member", zap.String("name", m.Name), zap.Strings("advertise-peer-urls", m.PeerURLs.StringSlice()), zap.Strings("listen-client-urls", m.ClientURLs.StringSlice()), zap.String("grpc-address", m.grpcAddr))
	var err error
	if m.s, err = etcdserver.NewServer(m.ServerConfig); err != nil {
		return fmt.Errorf("failed to initialize the etcd server: %v", err)
	}
	m.s.SyncTicker = time.NewTicker(500 * time.Millisecond)
	m.s.Start()
	var peerTLScfg *tls.Config
	if m.PeerTLSInfo != nil && !m.PeerTLSInfo.Empty() {
		if peerTLScfg, err = m.PeerTLSInfo.ServerConfig(); err != nil {
			return err
		}
	}
	if m.grpcListener != nil {
		var (
			tlscfg *tls.Config
		)
		if m.ClientTLSInfo != nil && !m.ClientTLSInfo.Empty() {
			tlscfg, err = m.ClientTLSInfo.ServerConfig()
			if err != nil {
				return err
			}
		}
		m.grpcServer = v3rpc.Server(m.s, tlscfg, m.grpcServerOpts...)
		m.grpcServerPeer = v3rpc.Server(m.s, peerTLScfg)
		m.serverClient = v3client.New(m.s)
		lockpb.RegisterLockServer(m.grpcServer, v3lock.NewLockServer(m.serverClient))
		epb.RegisterElectionServer(m.grpcServer, v3election.NewElectionServer(m.serverClient))
		go m.grpcServer.Serve(m.grpcListener)
	}
	m.raftHandler = &testutil.PauseableHandler{Next: etcdhttp.NewPeerHandler(m.Logger, m.s)}
	h := (http.Handler)(m.raftHandler)
	if m.grpcListener != nil {
		h = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				m.grpcServerPeer.ServeHTTP(w, r)
			} else {
				m.raftHandler.ServeHTTP(w, r)
			}
		})
	}
	for _, ln := range m.PeerListeners {
		cm := cmux.New(ln)
		cm.SetReadTimeout(time.Second)
		if m.grpcServer != nil {
			grpcl := cm.Match(cmux.HTTP2())
			go m.grpcServerPeer.Serve(grpcl)
		}
		ll := cm.Match(cmux.Any())
		if peerTLScfg != nil {
			if ll, err = transport.NewTLSListener(ll, m.PeerTLSInfo); err != nil {
				return err
			}
		}
		hs := &httptest.Server{Listener: ll, Config: &http.Server{Handler: h, TLSConfig: peerTLScfg, ErrorLog: log.New(ioutil.Discard, "net/http", 0)}, TLS: peerTLScfg}
		hs.Start()
		donec := make(chan struct{})
		go func() {
			defer close(donec)
			cm.Serve()
		}()
		closer := func() {
			ll.Close()
			hs.CloseClientConnections()
			hs.Close()
			<-donec
		}
		m.serverClosers = append(m.serverClosers, closer)
	}
	for _, ln := range m.ClientListeners {
		hs := &httptest.Server{Listener: ln, Config: &http.Server{Handler: v2http.NewClientHandler(m.Logger, m.s, m.ServerConfig.ReqTimeout()), ErrorLog: log.New(ioutil.Discard, "net/http", 0)}}
		if m.ClientTLSInfo == nil {
			hs.Start()
		} else {
			info := m.ClientTLSInfo
			hs.TLS, err = info.ServerConfig()
			if err != nil {
				return err
			}
			tlsCert, err := tlsutil.NewCert(info.CertFile, info.KeyFile, nil)
			if err != nil {
				return err
			}
			hs.TLS.Certificates = []tls.Certificate{*tlsCert}
			hs.StartTLS()
		}
		closer := func() {
			ln.Close()
			hs.CloseClientConnections()
			hs.Close()
		}
		m.serverClosers = append(m.serverClosers, closer)
	}
	lg.Info("launched a member", zap.String("name", m.Name), zap.Strings("advertise-peer-urls", m.PeerURLs.StringSlice()), zap.Strings("listen-client-urls", m.ClientURLs.StringSlice()), zap.String("grpc-address", m.grpcAddr))
	return nil
}
func (m *member) WaitOK(t testing.TB) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.WaitStarted(t)
	for m.s.Leader() == 0 {
		time.Sleep(tickDuration)
	}
}
func (m *member) WaitStarted(t testing.TB) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := MustNewHTTPClient(t, []string{m.URL()}, m.ClientTLSInfo)
	kapi := client.NewKeysAPI(cc)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		_, err := kapi.Get(ctx, "/", nil)
		if err != nil {
			time.Sleep(tickDuration)
			continue
		}
		cancel()
		break
	}
}
func WaitClientV3(t testing.TB, kv clientv3.KV) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	timeout := time.Now().Add(requestTimeout)
	var err error
	for time.Now().Before(timeout) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		_, err = kv.Get(ctx, "/")
		cancel()
		if err == nil {
			return
		}
		time.Sleep(tickDuration)
	}
	if err != nil {
		t.Fatalf("timed out waiting for client: %v", err)
	}
}
func (m *member) URL() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.ClientURLs[0].String()
}
func (m *member) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.raftHandler.Pause()
	m.s.PauseSending()
}
func (m *member) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.raftHandler.Resume()
	m.s.ResumeSending()
}
func (m *member) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.grpcBridge != nil {
		m.grpcBridge.Close()
		m.grpcBridge = nil
	}
	if m.serverClient != nil {
		m.serverClient.Close()
		m.serverClient = nil
	}
	if m.grpcServer != nil {
		m.grpcServer.Stop()
		m.grpcServer.GracefulStop()
		m.grpcServer = nil
		m.grpcServerPeer.Stop()
		m.grpcServerPeer.GracefulStop()
		m.grpcServerPeer = nil
	}
	m.s.HardStop()
	for _, f := range m.serverClosers {
		f()
	}
}
func (m *member) Stop(t testing.TB) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg.Info("stopping a member", zap.String("name", m.Name), zap.Strings("advertise-peer-urls", m.PeerURLs.StringSlice()), zap.Strings("listen-client-urls", m.ClientURLs.StringSlice()), zap.String("grpc-address", m.grpcAddr))
	m.Close()
	m.serverClosers = nil
	lg.Info("stopped a member", zap.String("name", m.Name), zap.Strings("advertise-peer-urls", m.PeerURLs.StringSlice()), zap.Strings("listen-client-urls", m.ClientURLs.StringSlice()), zap.String("grpc-address", m.grpcAddr))
}
func checkLeaderTransition(m *member, oldLead uint64) uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	interval := time.Duration(m.s.Cfg.TickMs) * time.Millisecond
	for m.s.Lead() == 0 || (m.s.Lead() == oldLead) {
		time.Sleep(interval)
	}
	return m.s.Lead()
}
func (m *member) StopNotify() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return m.s.StopNotify()
}
func (m *member) Restart(t testing.TB) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg.Info("restarting a member", zap.String("name", m.Name), zap.Strings("advertise-peer-urls", m.PeerURLs.StringSlice()), zap.Strings("listen-client-urls", m.ClientURLs.StringSlice()), zap.String("grpc-address", m.grpcAddr))
	newPeerListeners := make([]net.Listener, 0)
	for _, ln := range m.PeerListeners {
		newPeerListeners = append(newPeerListeners, NewListenerWithAddr(t, ln.Addr().String()))
	}
	m.PeerListeners = newPeerListeners
	newClientListeners := make([]net.Listener, 0)
	for _, ln := range m.ClientListeners {
		newClientListeners = append(newClientListeners, NewListenerWithAddr(t, ln.Addr().String()))
	}
	m.ClientListeners = newClientListeners
	if m.grpcListener != nil {
		if err := m.listenGRPC(); err != nil {
			t.Fatal(err)
		}
	}
	err := m.Launch()
	lg.Info("restarted a member", zap.String("name", m.Name), zap.Strings("advertise-peer-urls", m.PeerURLs.StringSlice()), zap.Strings("listen-client-urls", m.ClientURLs.StringSlice()), zap.String("grpc-address", m.grpcAddr), zap.Error(err))
	return err
}
func (m *member) Terminate(t testing.TB) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lg.Info("terminating a member", zap.String("name", m.Name), zap.Strings("advertise-peer-urls", m.PeerURLs.StringSlice()), zap.Strings("listen-client-urls", m.ClientURLs.StringSlice()), zap.String("grpc-address", m.grpcAddr))
	m.Close()
	if !m.keepDataDirTerminate {
		if err := os.RemoveAll(m.ServerConfig.DataDir); err != nil {
			t.Fatal(err)
		}
	}
	lg.Info("terminated a member", zap.String("name", m.Name), zap.Strings("advertise-peer-urls", m.PeerURLs.StringSlice()), zap.Strings("listen-client-urls", m.ClientURLs.StringSlice()), zap.String("grpc-address", m.grpcAddr))
}
func (m *member) Metric(metricName string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfgtls := transport.TLSInfo{}
	tr, err := transport.NewTimeoutTransport(cfgtls, time.Second, time.Second, time.Second)
	if err != nil {
		return "", err
	}
	cli := &http.Client{Transport: tr}
	resp, err := cli.Get(m.ClientURLs[0].String() + "/metrics")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		return "", rerr
	}
	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		if strings.HasPrefix(l, metricName) {
			return strings.Split(l, " ")[1], nil
		}
	}
	return "", nil
}
func (m *member) InjectPartition(t testing.TB, others ...*member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, other := range others {
		m.s.CutPeer(other.s.ID())
		other.s.CutPeer(m.s.ID())
	}
}
func (m *member) RecoverPartition(t testing.TB, others ...*member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, other := range others {
		m.s.MendPeer(other.s.ID())
		other.s.MendPeer(m.s.ID())
	}
}
func MustNewHTTPClient(t testing.TB, eps []string, tls *transport.TLSInfo) client.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfgtls := transport.TLSInfo{}
	if tls != nil {
		cfgtls = *tls
	}
	cfg := client.Config{Transport: mustNewTransport(t, cfgtls), Endpoints: eps}
	c, err := client.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	return c
}
func mustNewTransport(t testing.TB, tlsInfo transport.TLSInfo) *http.Transport {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr, err := transport.NewTimeoutTransport(tlsInfo, time.Second, rafthttp.ConnReadTimeout, rafthttp.ConnWriteTimeout)
	if err != nil {
		t.Fatal(err)
	}
	return tr
}

type SortableMemberSliceByPeerURLs []client.Member

func (p SortableMemberSliceByPeerURLs) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(p)
}
func (p SortableMemberSliceByPeerURLs) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p[i].PeerURLs[0] < p[j].PeerURLs[0]
}
func (p SortableMemberSliceByPeerURLs) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p[i], p[j] = p[j], p[i]
}

type ClusterV3 struct {
	*cluster
	mu	sync.Mutex
	clients	[]*clientv3.Client
}

func NewClusterV3(t testing.TB, cfg *ClusterConfig) *ClusterV3 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg.UseGRPC = true
	if os.Getenv("CLIENT_DEBUG") != "" {
		clientv3.SetLogger(grpclog.NewLoggerV2WithVerbosity(os.Stderr, os.Stderr, os.Stderr, 4))
	}
	clus := &ClusterV3{cluster: NewClusterByConfig(t, cfg)}
	clus.Launch(t)
	if !cfg.SkipCreatingClient {
		for _, m := range clus.Members {
			client, err := NewClientV3(m)
			if err != nil {
				t.Fatalf("cannot create client: %v", err)
			}
			clus.clients = append(clus.clients, client)
		}
	}
	return clus
}
func (c *ClusterV3) TakeClient(idx int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mu.Lock()
	c.clients[idx] = nil
	c.mu.Unlock()
}
func (c *ClusterV3) Terminate(t testing.TB) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mu.Lock()
	for _, client := range c.clients {
		if client == nil {
			continue
		}
		if err := client.Close(); err != nil {
			t.Error(err)
		}
	}
	c.mu.Unlock()
	c.cluster.Terminate(t)
}
func (c *ClusterV3) RandClient() *clientv3.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.clients[rand.Intn(len(c.clients))]
}
func (c *ClusterV3) Client(i int) *clientv3.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.clients[i]
}

type grpcAPI struct {
	Cluster		pb.ClusterClient
	KV		pb.KVClient
	Lease		pb.LeaseClient
	Watch		pb.WatchClient
	Maintenance	pb.MaintenanceClient
	Auth		pb.AuthClient
	Lock		lockpb.LockClient
	Election	epb.ElectionClient
}
