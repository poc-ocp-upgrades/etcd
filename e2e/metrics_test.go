package e2e

import (
	"fmt"
	"testing"
	"github.com/coreos/etcd/version"
)

func TestV3MetricsSecure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := configTLS
	cfg.clusterSize = 1
	cfg.metricsURLScheme = "https"
	testCtl(t, metricsTest)
}
func TestV3MetricsInsecure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := configTLS
	cfg.clusterSize = 1
	cfg.metricsURLScheme = "http"
	testCtl(t, metricsTest)
}
func metricsTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3Put(cx, "k", "v", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := cURLGet(cx.epc, cURLReq{endpoint: "/metrics", expected: `etcd_debugging_mvcc_keys_total 1`, metricsURLScheme: cx.cfg.metricsURLScheme}); err != nil {
		cx.t.Fatalf("failed get with curl (%v)", err)
	}
	if err := cURLGet(cx.epc, cURLReq{endpoint: "/metrics", expected: fmt.Sprintf(`etcd_server_version{server_version="%s"} 1`, version.Version), metricsURLScheme: cx.cfg.metricsURLScheme}); err != nil {
		cx.t.Fatalf("failed get with curl (%v)", err)
	}
	if err := cURLGet(cx.epc, cURLReq{endpoint: "/health", expected: `{"health":"true"}`, metricsURLScheme: cx.cfg.metricsURLScheme}); err != nil {
		cx.t.Fatalf("failed get with curl (%v)", err)
	}
}
