package embed

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"testing"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/ghodss/yaml"
)

func TestConfigFileOtherFields(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctls := securityConfig{CAFile: "cca", CertFile: "ccert", KeyFile: "ckey"}
	ptls := securityConfig{CAFile: "pca", CertFile: "pcert", KeyFile: "pkey"}
	yc := struct {
		ClientSecurityCfgFile	securityConfig	`json:"client-transport-security"`
		PeerSecurityCfgFile	securityConfig	`json:"peer-transport-security"`
		ForceNewCluster		bool		`json:"force-new-cluster"`
	}{ctls, ptls, true}
	b, err := yaml.Marshal(&yc)
	if err != nil {
		t.Fatal(err)
	}
	tmpfile := mustCreateCfgFile(t, b)
	defer os.Remove(tmpfile.Name())
	cfg, err := ConfigFromFile(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if !cfg.ForceNewCluster {
		t.Errorf("ForceNewCluster = %v, want %v", cfg.ForceNewCluster, true)
	}
	if !ctls.equals(&cfg.ClientTLSInfo) {
		t.Errorf("ClientTLS = %v, want %v", cfg.ClientTLSInfo, ctls)
	}
	if !ptls.equals(&cfg.PeerTLSInfo) {
		t.Errorf("PeerTLS = %v, want %v", cfg.PeerTLSInfo, ptls)
	}
}
func TestUpdateDefaultClusterFromName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := NewConfig()
	defaultInitialCluster := cfg.InitialCluster
	oldscheme := cfg.APUrls[0].Scheme
	origpeer := cfg.APUrls[0].String()
	origadvc := cfg.ACUrls[0].String()
	cfg.Name = "abc"
	lpport := cfg.LPUrls[0].Port()
	exp := fmt.Sprintf("%s=%s://localhost:%s", cfg.Name, oldscheme, lpport)
	cfg.UpdateDefaultClusterFromName(defaultInitialCluster)
	if exp != cfg.InitialCluster {
		t.Fatalf("initial-cluster expected %q, got %q", exp, cfg.InitialCluster)
	}
	if origpeer != cfg.APUrls[0].String() {
		t.Fatalf("advertise peer url expected %q, got %q", origadvc, cfg.APUrls[0].String())
	}
	if origadvc != cfg.ACUrls[0].String() {
		t.Fatalf("advertise client url expected %q, got %q", origadvc, cfg.ACUrls[0].String())
	}
}
func TestUpdateDefaultClusterFromNameOverwrite(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if defaultHostname == "" {
		t.Skip("machine's default host not found")
	}
	cfg := NewConfig()
	defaultInitialCluster := cfg.InitialCluster
	oldscheme := cfg.APUrls[0].Scheme
	origadvc := cfg.ACUrls[0].String()
	cfg.Name = "abc"
	lpport := cfg.LPUrls[0].Port()
	cfg.LPUrls[0] = url.URL{Scheme: cfg.LPUrls[0].Scheme, Host: fmt.Sprintf("0.0.0.0:%s", lpport)}
	dhost, _ := cfg.UpdateDefaultClusterFromName(defaultInitialCluster)
	if dhost != defaultHostname {
		t.Fatalf("expected default host %q, got %q", defaultHostname, dhost)
	}
	aphost, apport := cfg.APUrls[0].Hostname(), cfg.APUrls[0].Port()
	if apport != lpport {
		t.Fatalf("advertise peer url got different port %s, expected %s", apport, lpport)
	}
	if aphost != defaultHostname {
		t.Fatalf("advertise peer url expected machine default host %q, got %q", defaultHostname, aphost)
	}
	expected := fmt.Sprintf("%s=%s://%s:%s", cfg.Name, oldscheme, defaultHostname, lpport)
	if expected != cfg.InitialCluster {
		t.Fatalf("initial-cluster expected %q, got %q", expected, cfg.InitialCluster)
	}
	if origadvc != cfg.ACUrls[0].String() {
		t.Fatalf("advertise-client-url expected %q, got %q", origadvc, cfg.ACUrls[0].String())
	}
}
func (s *securityConfig) equals(t *transport.TLSInfo) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.CAFile == t.CAFile && s.CertFile == t.CertFile && s.CertAuth == t.ClientCertAuth && s.TrustedCAFile == t.TrustedCAFile
}
func mustCreateCfgFile(t *testing.T, b []byte) *os.File {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tmpfile, err := ioutil.TempFile("", "servercfg")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = tmpfile.Write(b); err != nil {
		t.Fatal(err)
	}
	if err = tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	return tmpfile
}
func TestAutoCompactionModeInvalid(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := NewConfig()
	cfg.AutoCompactionMode = "period"
	err := cfg.Validate()
	if err == nil {
		t.Errorf("expected non-nil error, got %v", err)
	}
}
func TestAutoCompactionModeParse(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dur, err := parseCompactionRetention("revision", "1")
	if err != nil {
		t.Error(err)
	}
	if dur != 1 {
		t.Fatalf("AutoCompactionRetention expected 1, got %d", dur)
	}
}
