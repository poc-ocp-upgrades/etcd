package e2e

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/coreos/etcd/pkg/types"
)

func TestCtlV3MoveLeaderSecure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV3MoveLeader(t, configTLS)
}
func TestCtlV3MoveLeaderInsecure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV3MoveLeader(t, configNoTLS)
}
func testCtlV3MoveLeader(t *testing.T, cfg etcdProcessClusterConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, &cfg, true)
	defer func() {
		if errC := epc.Close(); errC != nil {
			t.Fatalf("error closing etcd processes (%v)", errC)
		}
	}()
	var tcfg *tls.Config
	if cfg.clientTLS == clientTLS {
		tinfo := transport.TLSInfo{CertFile: certPath, KeyFile: privateKeyPath, TrustedCAFile: caPath}
		var err error
		tcfg, err = tinfo.ClientConfig()
		if err != nil {
			t.Fatal(err)
		}
	}
	var leadIdx int
	var leaderID uint64
	var transferee uint64
	for i, ep := range epc.EndpointsV3() {
		cli, err := clientv3.New(clientv3.Config{Endpoints: []string{ep}, DialTimeout: 3 * time.Second, TLS: tcfg})
		if err != nil {
			t.Fatal(err)
		}
		resp, err := cli.Status(context.Background(), ep)
		if err != nil {
			t.Fatal(err)
		}
		cli.Close()
		if resp.Header.GetMemberId() == resp.Leader {
			leadIdx = i
			leaderID = resp.Leader
		} else {
			transferee = resp.Header.GetMemberId()
		}
	}
	os.Setenv("ETCDCTL_API", "3")
	defer os.Unsetenv("ETCDCTL_API")
	cx := ctlCtx{t: t, cfg: configNoTLS, dialTimeout: 7 * time.Second, epc: epc}
	tests := []struct {
		prefixes	[]string
		expect		string
	}{{cx.prefixArgs([]string{cx.epc.EndpointsV3()[(leadIdx+1)%3]}), "no leader endpoint given at "}, {cx.prefixArgs([]string{cx.epc.EndpointsV3()[leadIdx]}), fmt.Sprintf("Leadership transferred from %s to %s", types.ID(leaderID), types.ID(transferee))}}
	for i, tc := range tests {
		cmdArgs := append(tc.prefixes, "move-leader", types.ID(transferee).String())
		if err := spawnWithExpect(cmdArgs, tc.expect); err != nil {
			t.Fatalf("#%d: %v", i, err)
		}
	}
}
