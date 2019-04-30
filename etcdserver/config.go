package etcdserver

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"go.etcd.io/etcd/pkg/netutil"
	"go.etcd.io/etcd/pkg/transport"
	"go.etcd.io/etcd/pkg/types"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ServerConfig struct {
	Name				string
	DiscoveryURL			string
	DiscoveryProxy			string
	ClientURLs			types.URLs
	PeerURLs			types.URLs
	DataDir				string
	DedicatedWALDir			string
	SnapshotCount			uint64
	SnapshotCatchUpEntries		uint64
	MaxSnapFiles			uint
	MaxWALFiles			uint
	BackendBatchInterval		time.Duration
	BackendBatchLimit		int
	BackendFreelistType		bolt.FreelistType
	InitialPeerURLsMap		types.URLsMap
	InitialClusterToken		string
	NewCluster			bool
	PeerTLSInfo			transport.TLSInfo
	CORS				map[string]struct{}
	HostWhitelist			map[string]struct{}
	TickMs				uint
	ElectionTicks			int
	InitialElectionTickAdvance	bool
	BootstrapTimeout		time.Duration
	AutoCompactionRetention		time.Duration
	AutoCompactionMode		string
	QuotaBackendBytes		int64
	MaxTxnOps			uint
	MaxRequestBytes			uint
	StrictReconfigCheck		bool
	ClientCertAuthEnabled		bool
	AuthToken			string
	BcryptCost			uint
	InitialCorruptCheck		bool
	CorruptCheckTime		time.Duration
	PreVote				bool
	Logger				*zap.Logger
	LoggerConfig			*zap.Config
	LoggerCore			zapcore.Core
	LoggerWriteSyncer		zapcore.WriteSyncer
	Debug				bool
	ForceNewCluster			bool
	LeaseCheckpointInterval		time.Duration
	EnableGRPCGateway		bool
}

func (c *ServerConfig) VerifyBootstrap() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.hasLocalMember(); err != nil {
		return err
	}
	if err := c.advertiseMatchesCluster(); err != nil {
		return err
	}
	if checkDuplicateURL(c.InitialPeerURLsMap) {
		return fmt.Errorf("initial cluster %s has duplicate url", c.InitialPeerURLsMap)
	}
	if c.InitialPeerURLsMap.String() == "" && c.DiscoveryURL == "" {
		return fmt.Errorf("initial cluster unset and no discovery URL found")
	}
	return nil
}
func (c *ServerConfig) VerifyJoinExisting() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.hasLocalMember(); err != nil {
		return err
	}
	if checkDuplicateURL(c.InitialPeerURLsMap) {
		return fmt.Errorf("initial cluster %s has duplicate url", c.InitialPeerURLsMap)
	}
	if c.DiscoveryURL != "" {
		return fmt.Errorf("discovery URL should not be set when joining existing initial cluster")
	}
	return nil
}
func (c *ServerConfig) hasLocalMember() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if urls := c.InitialPeerURLsMap[c.Name]; urls == nil {
		return fmt.Errorf("couldn't find local name %q in the initial cluster configuration", c.Name)
	}
	return nil
}
func (c *ServerConfig) advertiseMatchesCluster() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	urls, apurls := c.InitialPeerURLsMap[c.Name], c.PeerURLs.StringSlice()
	urls.Sort()
	sort.Strings(apurls)
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	ok, err := netutil.URLStringsEqual(ctx, c.Logger, apurls, urls.StringSlice())
	if ok {
		return nil
	}
	initMap, apMap := make(map[string]struct{}), make(map[string]struct{})
	for _, url := range c.PeerURLs {
		apMap[url.String()] = struct{}{}
	}
	for _, url := range c.InitialPeerURLsMap[c.Name] {
		initMap[url.String()] = struct{}{}
	}
	missing := []string{}
	for url := range initMap {
		if _, ok := apMap[url]; !ok {
			missing = append(missing, url)
		}
	}
	if len(missing) > 0 {
		for i := range missing {
			missing[i] = c.Name + "=" + missing[i]
		}
		mstr := strings.Join(missing, ",")
		apStr := strings.Join(apurls, ",")
		return fmt.Errorf("--initial-cluster has %s but missing from --initial-advertise-peer-urls=%s (%v)", mstr, apStr, err)
	}
	for url := range apMap {
		if _, ok := initMap[url]; !ok {
			missing = append(missing, url)
		}
	}
	if len(missing) > 0 {
		mstr := strings.Join(missing, ",")
		umap := types.URLsMap(map[string]types.URLs{c.Name: c.PeerURLs})
		return fmt.Errorf("--initial-advertise-peer-urls has %s but missing from --initial-cluster=%s", mstr, umap.String())
	}
	apStr := strings.Join(apurls, ",")
	umap := types.URLsMap(map[string]types.URLs{c.Name: c.PeerURLs})
	return fmt.Errorf("failed to resolve %s to match --initial-cluster=%s (%v)", apStr, umap.String(), err)
}
func (c *ServerConfig) MemberDir() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return filepath.Join(c.DataDir, "member")
}
func (c *ServerConfig) WALDir() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.DedicatedWALDir != "" {
		return c.DedicatedWALDir
	}
	return filepath.Join(c.MemberDir(), "wal")
}
func (c *ServerConfig) SnapDir() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return filepath.Join(c.MemberDir(), "snap")
}
func (c *ServerConfig) ShouldDiscover() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.DiscoveryURL != ""
}
func (c *ServerConfig) ReqTimeout() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 5*time.Second + 2*time.Duration(c.ElectionTicks*int(c.TickMs))*time.Millisecond
}
func (c *ServerConfig) electionTimeout() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Duration(c.ElectionTicks*int(c.TickMs)) * time.Millisecond
}
func (c *ServerConfig) peerDialTimeout() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Second + time.Duration(c.ElectionTicks*int(c.TickMs))*time.Millisecond
}
func checkDuplicateURL(urlsmap types.URLsMap) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	um := make(map[string]bool)
	for _, urls := range urlsmap {
		for _, url := range urls {
			u := url.String()
			if um[u] {
				return true
			}
			um[u] = true
		}
	}
	return false
}
func (c *ServerConfig) bootstrapTimeout() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.BootstrapTimeout != 0 {
		return c.BootstrapTimeout
	}
	return time.Second
}
func (c *ServerConfig) backendPath() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return filepath.Join(c.SnapDir(), "db")
}
