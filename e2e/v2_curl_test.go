package e2e

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestV2CurlNoTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCurlPutGet(t, &configNoTLS)
}
func TestV2CurlAutoTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCurlPutGet(t, &configAutoTLS)
}
func TestV2CurlAllTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCurlPutGet(t, &configTLS)
}
func TestV2CurlPeerTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCurlPutGet(t, &configPeerTLS)
}
func TestV2CurlClientTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCurlPutGet(t, &configClientTLS)
}
func TestV2CurlClientBoth(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCurlPutGet(t, &configClientBoth)
}
func testCurlPutGet(t *testing.T, cfg *etcdProcessClusterConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	cfg = configStandalone(*cfg)
	epc, err := newEtcdProcessCluster(cfg)
	if err != nil {
		t.Fatalf("could not start etcd process cluster (%v)", err)
	}
	defer func() {
		if err := epc.Close(); err != nil {
			t.Fatalf("error closing etcd processes (%v)", err)
		}
	}()
	var (
		expectPut	= `{"action":"set","node":{"key":"/foo","value":"bar","`
		expectGet	= `{"action":"get","node":{"key":"/foo","value":"bar","`
	)
	if err := cURLPut(epc, cURLReq{endpoint: "/v2/keys/foo", value: "bar", expected: expectPut}); err != nil {
		t.Fatalf("failed put with curl (%v)", err)
	}
	if err := cURLGet(epc, cURLReq{endpoint: "/v2/keys/foo", expected: expectGet}); err != nil {
		t.Fatalf("failed get with curl (%v)", err)
	}
	if cfg.clientTLS == clientTLSAndNonTLS {
		if err := cURLGet(epc, cURLReq{endpoint: "/v2/keys/foo", expected: expectGet, isTLS: true}); err != nil {
			t.Fatalf("failed get with curl (%v)", err)
		}
	}
}
func TestV2CurlIssue5182(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, &configNoTLS, false)
	defer func() {
		if err := epc.Close(); err != nil {
			t.Fatalf("error closing etcd processes (%v)", err)
		}
	}()
	expectPut := `{"action":"set","node":{"key":"/foo","value":"bar","`
	if err := cURLPut(epc, cURLReq{endpoint: "/v2/keys/foo", value: "bar", expected: expectPut}); err != nil {
		t.Fatal(err)
	}
	expectUserAdd := `{"user":"foo","roles":null}`
	if err := cURLPut(epc, cURLReq{endpoint: "/v2/auth/users/foo", value: `{"user":"foo", "password":"pass"}`, expected: expectUserAdd}); err != nil {
		t.Fatal(err)
	}
	expectRoleAdd := `{"role":"foo","permissions":{"kv":{"read":["/foo/*"],"write":null}}`
	if err := cURLPut(epc, cURLReq{endpoint: "/v2/auth/roles/foo", value: `{"role":"foo", "permissions": {"kv": {"read": ["/foo/*"]}}}`, expected: expectRoleAdd}); err != nil {
		t.Fatal(err)
	}
	expectUserUpdate := `{"user":"foo","roles":["foo"]}`
	if err := cURLPut(epc, cURLReq{endpoint: "/v2/auth/users/foo", value: `{"user": "foo", "grant": ["foo"]}`, expected: expectUserUpdate}); err != nil {
		t.Fatal(err)
	}
	if err := etcdctlUserAdd(epc, "root", "a"); err != nil {
		t.Fatal(err)
	}
	if err := etcdctlAuthEnable(epc); err != nil {
		t.Fatal(err)
	}
	if err := cURLGet(epc, cURLReq{endpoint: "/v2/keys/foo/", username: "root", password: "a", expected: "bar"}); err != nil {
		t.Fatal(err)
	}
	if err := cURLGet(epc, cURLReq{endpoint: "/v2/keys/foo/", username: "foo", password: "pass", expected: "bar"}); err != nil {
		t.Fatal(err)
	}
	if err := cURLGet(epc, cURLReq{endpoint: "/v2/keys/foo/", username: "foo", password: "", expected: "bar"}); err != nil {
		if !strings.Contains(err.Error(), `The request requires user authentication`) {
			t.Fatalf("expected 'The request requires user authentication' error, got %v", err)
		}
	} else {
		t.Fatalf("expected 'The request requires user authentication' error")
	}
}

type cURLReq struct {
	username		string
	password		string
	isTLS			bool
	timeout			int
	endpoint		string
	value			string
	expected		string
	header			string
	metricsURLScheme	string
	ciphers			string
}

func cURLPrefixArgs(clus *etcdProcessCluster, method string, req cURLReq) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		cmdArgs	= []string{"curl"}
		acurl	= clus.procs[rand.Intn(clus.cfg.clusterSize)].Config().acurl
	)
	if req.metricsURLScheme != "https" {
		if req.isTLS {
			if clus.cfg.clientTLS != clientTLSAndNonTLS {
				panic("should not use cURLPrefixArgsUseTLS when serving only TLS or non-TLS")
			}
			cmdArgs = append(cmdArgs, "--cacert", caPath, "--cert", certPath, "--key", privateKeyPath)
			acurl = toTLS(clus.procs[rand.Intn(clus.cfg.clusterSize)].Config().acurl)
		} else if clus.cfg.clientTLS == clientTLS {
			cmdArgs = append(cmdArgs, "--cacert", caPath, "--cert", certPath, "--key", privateKeyPath)
		}
	}
	if req.metricsURLScheme != "" {
		acurl = clus.procs[rand.Intn(clus.cfg.clusterSize)].EndpointsMetrics()[0]
	}
	ep := acurl + req.endpoint
	if req.username != "" || req.password != "" {
		cmdArgs = append(cmdArgs, "-L", "-u", fmt.Sprintf("%s:%s", req.username, req.password), ep)
	} else {
		cmdArgs = append(cmdArgs, "-L", ep)
	}
	if req.timeout != 0 {
		cmdArgs = append(cmdArgs, "-m", fmt.Sprintf("%d", req.timeout))
	}
	if req.header != "" {
		cmdArgs = append(cmdArgs, "-H", req.header)
	}
	if req.ciphers != "" {
		cmdArgs = append(cmdArgs, "--ciphers", req.ciphers)
	}
	switch method {
	case "POST", "PUT":
		dt := req.value
		if !strings.HasPrefix(dt, "{") {
			dt = "value=" + dt
		}
		cmdArgs = append(cmdArgs, "-X", method, "-d", dt)
	}
	return cmdArgs
}
func cURLPost(clus *etcdProcessCluster, req cURLReq) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return spawnWithExpect(cURLPrefixArgs(clus, "POST", req), req.expected)
}
func cURLPut(clus *etcdProcessCluster, req cURLReq) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return spawnWithExpect(cURLPrefixArgs(clus, "PUT", req), req.expected)
}
func cURLGet(clus *etcdProcessCluster, req cURLReq) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return spawnWithExpect(cURLPrefixArgs(clus, "GET", req), req.expected)
}
