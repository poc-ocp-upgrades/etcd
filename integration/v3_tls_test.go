package integration

import (
	"context"
	"crypto/tls"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestTLSClientCipherSuitesValid(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testTLSCipherSuites(t, true)
}
func TestTLSClientCipherSuitesMismatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testTLSCipherSuites(t, false)
}
func testTLSCipherSuites(t *testing.T, valid bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	cipherSuites := []uint16{tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305, tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305}
	srvTLS, cliTLS := testTLSInfo, testTLSInfo
	if valid {
		srvTLS.CipherSuites, cliTLS.CipherSuites = cipherSuites, cipherSuites
	} else {
		srvTLS.CipherSuites, cliTLS.CipherSuites = cipherSuites[:2], cipherSuites[2:]
	}
	clus := NewClusterV3(t, &ClusterConfig{Size: 1, ClientTLS: &srvTLS})
	defer clus.Terminate(t)
	cc, err := cliTLS.ClientConfig()
	if err != nil {
		t.Fatal(err)
	}
	cli, cerr := clientv3.New(clientv3.Config{Endpoints: []string{clus.Members[0].GRPCAddr()}, DialTimeout: time.Second, TLS: cc})
	if cli != nil {
		cli.Close()
	}
	if !valid && cerr != context.DeadlineExceeded {
		t.Fatalf("expected %v with TLS handshake failure, got %v", context.DeadlineExceeded, cerr)
	}
	if valid && cerr != nil {
		t.Fatalf("expected TLS handshake success, got %v", cerr)
	}
}
