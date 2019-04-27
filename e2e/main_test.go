package e2e

import (
	"flag"
	"os"
	"runtime"
	"testing"
	"github.com/coreos/etcd/pkg/testutil"
)

var (
	binDir			string
	certDir			string
	certPath		string
	privateKeyPath		string
	caPath			string
	certPath2		string
	privateKeyPath2		string
	crlPath			string
	revokedCertPath		string
	revokedPrivateKeyPath	string
)

func TestMain(m *testing.M) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	os.Setenv("ETCD_UNSUPPORTED_ARCH", runtime.GOARCH)
	os.Unsetenv("ETCDCTL_API")
	flag.StringVar(&binDir, "bin-dir", "../bin", "The directory for store etcd and etcdctl binaries.")
	flag.StringVar(&certDir, "cert-dir", "../integration/fixtures", "The directory for store certificate files.")
	flag.Parse()
	binPath = binDir + "/etcd"
	ctlBinPath = binDir + "/etcdctl"
	certPath = certDir + "/server.crt"
	privateKeyPath = certDir + "/server.key.insecure"
	caPath = certDir + "/ca.crt"
	revokedCertPath = certDir + "/server-revoked.crt"
	revokedPrivateKeyPath = certDir + "/server-revoked.key.insecure"
	crlPath = certDir + "/revoke.crl"
	certPath2 = certDir + "/server2.crt"
	privateKeyPath2 = certDir + "/server2.key.insecure"
	v := m.Run()
	if v == 0 && testutil.CheckLeakedGoroutine() {
		os.Exit(1)
	}
	os.Exit(v)
}
