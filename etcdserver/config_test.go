package etcdserver

import (
	"net/url"
	"testing"
	"github.com/coreos/etcd/pkg/types"
)

func mustNewURLs(t *testing.T, urls []string) []url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(urls) == 0 {
		return nil
	}
	u, err := types.NewURLs(urls)
	if err != nil {
		t.Fatalf("error creating new URLs from %q: %v", urls, err)
	}
	return u
}
func TestConfigVerifyBootstrapWithoutClusterAndDiscoveryURLFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &ServerConfig{Name: "node1", DiscoveryURL: "", InitialPeerURLsMap: types.URLsMap{}}
	if err := c.VerifyBootstrap(); err == nil {
		t.Errorf("err = nil, want not nil")
	}
}
func TestConfigVerifyExistingWithDiscoveryURLFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster, err := types.NewURLsMap("node1=http://127.0.0.1:2380")
	if err != nil {
		t.Fatalf("NewCluster error: %v", err)
	}
	c := &ServerConfig{Name: "node1", DiscoveryURL: "http://127.0.0.1:2379/abcdefg", PeerURLs: mustNewURLs(t, []string{"http://127.0.0.1:2380"}), InitialPeerURLsMap: cluster, NewCluster: false}
	if err := c.VerifyJoinExisting(); err == nil {
		t.Errorf("err = nil, want not nil")
	}
}
func TestConfigVerifyLocalMember(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		clusterSetting	string
		apurls		[]string
		strict		bool
		shouldError	bool
	}{{"", nil, true, true}, {"node1=http://localhost:7001,node2=http://localhost:7002", []string{"http://localhost:7001"}, true, false}, {"node1=http://localhost:2380,node1=http://localhost:7001", []string{"http://localhost:2380", "http://localhost:7001"}, true, false}, {"node1=http://localhost:7001", []string{"http://localhost:12345"}, true, true}, {"node1=http://localhost:2380,node1=http://localhost:12345", []string{"http://localhost:12345"}, true, true}, {"node1=http://localhost:12345", []string{"http://localhost:2380", "http://localhost:12345"}, true, true}, {"node1=http://localhost:2380", []string{}, true, true}, {"node1=http://localhost:2380", []string{}, false, false}}
	for i, tt := range tests {
		cluster, err := types.NewURLsMap(tt.clusterSetting)
		if err != nil {
			t.Fatalf("#%d: Got unexpected error: %v", i, err)
		}
		cfg := ServerConfig{Name: "node1", InitialPeerURLsMap: cluster}
		if tt.apurls != nil {
			cfg.PeerURLs = mustNewURLs(t, tt.apurls)
		}
		if err = cfg.hasLocalMember(); err == nil && tt.strict {
			err = cfg.advertiseMatchesCluster()
		}
		if (err == nil) && tt.shouldError {
			t.Errorf("#%d: Got no error where one was expected", i)
		}
		if (err != nil) && !tt.shouldError {
			t.Errorf("#%d: Got unexpected error: %v", i, err)
		}
	}
}
func TestSnapDir(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := map[string]string{"/": "/member/snap", "/var/lib/etc": "/var/lib/etc/member/snap"}
	for dd, w := range tests {
		cfg := ServerConfig{DataDir: dd}
		if g := cfg.SnapDir(); g != w {
			t.Errorf("DataDir=%q: SnapDir()=%q, want=%q", dd, g, w)
		}
	}
}
func TestWALDir(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := map[string]string{"/": "/member/wal", "/var/lib/etc": "/var/lib/etc/member/wal"}
	for dd, w := range tests {
		cfg := ServerConfig{DataDir: dd}
		if g := cfg.WALDir(); g != w {
			t.Errorf("DataDir=%q: WALDir()=%q, want=%q", dd, g, w)
		}
	}
}
func TestShouldDiscover(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := map[string]bool{"": false, "foo": true, "http://discovery.etcd.io/asdf": true}
	for durl, w := range tests {
		cfg := ServerConfig{DiscoveryURL: durl}
		if g := cfg.ShouldDiscover(); g != w {
			t.Errorf("durl=%q: ShouldDiscover()=%t, want=%t", durl, g, w)
		}
	}
}
