package e2e

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/pkg/flags"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/version"
)

func TestCtlV3Version(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, versionTest)
}
func versionTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3Version(cx); err != nil {
		cx.t.Fatalf("versionTest ctlV3Version error (%v)", err)
	}
}
func ctlV3Version(cx ctlCtx) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "version")
	return spawnWithExpect(cmdArgs, version.Version)
}
func TestCtlV3DialWithHTTPScheme(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, dialWithSchemeTest, withCfg(configClientTLS))
}
func dialWithSchemeTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.prefixArgs(cx.epc.EndpointsV3()), "put", "foo", "bar")
	if err := spawnWithExpect(cmdArgs, "OK"); err != nil {
		cx.t.Fatal(err)
	}
}

type ctlCtx struct {
	t			*testing.T
	cfg			etcdProcessClusterConfig
	quotaBackendBytes	int64
	corruptFunc		func(string) error
	noStrictReconfig	bool
	epc			*etcdProcessCluster
	envMap			map[string]struct{}
	dialTimeout		time.Duration
	quorum			bool
	interactive		bool
	user			string
	pass			string
	initialCorruptCheck	bool
	compactPhysical		bool
}
type ctlOption func(*ctlCtx)

func (cx *ctlCtx) applyOpts(opts []ctlOption) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, opt := range opts {
		opt(cx)
	}
	cx.initialCorruptCheck = true
}
func withCfg(cfg etcdProcessClusterConfig) ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.cfg = cfg
	}
}
func withDialTimeout(timeout time.Duration) ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.dialTimeout = timeout
	}
}
func withQuorum() ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.quorum = true
	}
}
func withInteractive() ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.interactive = true
	}
}
func withQuota(b int64) ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.quotaBackendBytes = b
	}
}
func withCompactPhysical() ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.compactPhysical = true
	}
}
func withInitialCorruptCheck() ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.initialCorruptCheck = true
	}
}
func withCorruptFunc(f func(string) error) ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.corruptFunc = f
	}
}
func withNoStrictReconfig() ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.noStrictReconfig = true
	}
}
func withFlagByEnv() ctlOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(cx *ctlCtx) {
		cx.envMap = make(map[string]struct{})
	}
}
func testCtl(t *testing.T, testFunc func(ctlCtx), opts ...ctlOption) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	ret := ctlCtx{t: t, cfg: configAutoTLS, dialTimeout: 7 * time.Second}
	ret.applyOpts(opts)
	mustEtcdctl(t)
	if !ret.quorum {
		ret.cfg = *configStandalone(ret.cfg)
	}
	if ret.quotaBackendBytes > 0 {
		ret.cfg.quotaBackendBytes = ret.quotaBackendBytes
	}
	ret.cfg.noStrictReconfig = ret.noStrictReconfig
	if ret.initialCorruptCheck {
		ret.cfg.initialCorruptCheck = ret.initialCorruptCheck
	}
	epc, err := newEtcdProcessCluster(&ret.cfg)
	if err != nil {
		t.Fatalf("could not start etcd process cluster (%v)", err)
	}
	ret.epc = epc
	defer func() {
		if ret.envMap != nil {
			for k := range ret.envMap {
				os.Unsetenv(k)
			}
		}
		if errC := ret.epc.Close(); errC != nil {
			t.Fatalf("error closing etcd processes (%v)", errC)
		}
	}()
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		testFunc(ret)
	}()
	timeout := 2*ret.dialTimeout + time.Second
	if ret.dialTimeout == 0 {
		timeout = 30 * time.Second
	}
	select {
	case <-time.After(timeout):
		testutil.FatalStack(t, fmt.Sprintf("test timed out after %v", timeout))
	case <-donec:
	}
}
func (cx *ctlCtx) prefixArgs(eps []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmap := make(map[string]string)
	fmap["endpoints"] = strings.Join(eps, ",")
	fmap["dial-timeout"] = cx.dialTimeout.String()
	if cx.epc.cfg.clientTLS == clientTLS {
		if cx.epc.cfg.isClientAutoTLS {
			fmap["insecure-transport"] = "false"
			fmap["insecure-skip-tls-verify"] = "true"
		} else if cx.epc.cfg.isClientCRL {
			fmap["cacert"] = caPath
			fmap["cert"] = revokedCertPath
			fmap["key"] = revokedPrivateKeyPath
		} else {
			fmap["cacert"] = caPath
			fmap["cert"] = certPath
			fmap["key"] = privateKeyPath
		}
	}
	if cx.user != "" {
		fmap["user"] = cx.user + ":" + cx.pass
	}
	useEnv := cx.envMap != nil
	cmdArgs := []string{ctlBinPath + "3"}
	for k, v := range fmap {
		if useEnv {
			ek := flags.FlagToEnv("ETCDCTL", k)
			os.Setenv(ek, v)
			cx.envMap[ek] = struct{}{}
		} else {
			cmdArgs = append(cmdArgs, fmt.Sprintf("--%s=%s", k, v))
		}
	}
	return cmdArgs
}
func (cx *ctlCtx) PrefixArgs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return cx.prefixArgs(cx.epc.EndpointsV3())
}
func isGRPCTimedout(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Contains(err.Error(), "grpc: timed out trying to connect")
}
func (cx *ctlCtx) memberToRemove() (ep string, memberID string, clusterID string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	n1 := cx.cfg.clusterSize
	if n1 < 2 {
		cx.t.Fatalf("%d-node is too small to test 'member remove'", n1)
	}
	resp, err := getMemberList(*cx)
	if err != nil {
		cx.t.Fatal(err)
	}
	if n1 != len(resp.Members) {
		cx.t.Fatalf("expected %d, got %d", n1, len(resp.Members))
	}
	ep = resp.Members[0].ClientURLs[0]
	clusterID = fmt.Sprintf("%x", resp.Header.ClusterId)
	memberID = fmt.Sprintf("%x", resp.Members[1].ID)
	return ep, memberID, clusterID
}
