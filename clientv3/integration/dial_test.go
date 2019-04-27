package integration

import (
	"context"
	"math/rand"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/pkg/transport"
)

var (
	testTLSInfo		= transport.TLSInfo{KeyFile: "../../integration/fixtures/server.key.insecure", CertFile: "../../integration/fixtures/server.crt", TrustedCAFile: "../../integration/fixtures/ca.crt", ClientCertAuth: true}
	testTLSInfoExpired	= transport.TLSInfo{KeyFile: "../../integration/fixtures-expired/server-key.pem", CertFile: "../../integration/fixtures-expired/server.pem", TrustedCAFile: "../../integration/fixtures-expired/etcd-root-ca.pem", ClientCertAuth: true}
)

func TestDialTLSExpired(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1, PeerTLS: &testTLSInfo, ClientTLS: &testTLSInfo, SkipCreatingClient: true})
	defer clus.Terminate(t)
	tls, err := testTLSInfoExpired.ClientConfig()
	if err != nil {
		t.Fatal(err)
	}
	_, err = clientv3.New(clientv3.Config{Endpoints: []string{clus.Members[0].GRPCAddr()}, DialTimeout: 3 * time.Second, TLS: tls})
	if err != context.DeadlineExceeded {
		t.Fatalf("expected %v, got %v", context.DeadlineExceeded, err)
	}
}
func TestDialTLSNoConfig(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1, ClientTLS: &testTLSInfo, SkipCreatingClient: true})
	defer clus.Terminate(t)
	_, err := clientv3.New(clientv3.Config{Endpoints: []string{clus.Members[0].GRPCAddr()}, DialTimeout: time.Second})
	if err != context.DeadlineExceeded {
		t.Fatalf("expected %v, got %v", context.DeadlineExceeded, err)
	}
}
func TestDialSetEndpointsBeforeFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testDialSetEndpoints(t, true)
}
func TestDialSetEndpointsAfterFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testDialSetEndpoints(t, false)
}
func testDialSetEndpoints(t *testing.T, setBefore bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3, SkipCreatingClient: true})
	defer clus.Terminate(t)
	eps := make([]string, 3)
	for i := range eps {
		eps[i] = clus.Members[i].GRPCAddr()
	}
	toKill := rand.Intn(len(eps))
	cfg := clientv3.Config{Endpoints: []string{eps[toKill]}, DialTimeout: 1 * time.Second}
	cli, err := clientv3.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close()
	if setBefore {
		cli.SetEndpoints(eps[toKill%3], eps[(toKill+1)%3])
	}
	clus.Members[toKill].Stop(t)
	clus.WaitLeader(t)
	if !setBefore {
		cli.SetEndpoints(eps[toKill%3], eps[(toKill+1)%3])
	}
	ctx, cancel := context.WithTimeout(context.Background(), integration.RequestWaitTimeout)
	if _, err = cli.Get(ctx, "foo", clientv3.WithSerializable()); err != nil {
		t.Fatal(err)
	}
	cancel()
}
func TestSwitchSetEndpoints(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	eps := []string{clus.Members[1].GRPCAddr(), clus.Members[2].GRPCAddr()}
	cli := clus.Client(0)
	clus.Members[0].InjectPartition(t, clus.Members[1:]...)
	cli.SetEndpoints(eps...)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := cli.Get(ctx, "foo"); err != nil {
		t.Fatal(err)
	}
}
func TestRejectOldCluster(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 2, SkipCreatingClient: true})
	defer clus.Terminate(t)
	cfg := clientv3.Config{Endpoints: []string{clus.Members[0].GRPCAddr(), clus.Members[1].GRPCAddr()}, DialTimeout: 5 * time.Second, RejectOldCluster: true}
	cli, err := clientv3.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	cli.Close()
}
func TestDialForeignEndpoint(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 2})
	defer clus.Terminate(t)
	conn, err := clus.Client(0).Dial(clus.Client(1).Endpoints()[0])
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	kvc := clientv3.NewKVFromKVClient(pb.NewKVClient(conn), clus.Client(0))
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	if _, gerr := kvc.Get(ctx, "abc"); gerr != nil {
		t.Fatal(err)
	}
}
func TestSetEndpointAndPut(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 2})
	defer clus.Terminate(t)
	clus.Client(1).SetEndpoints(clus.Members[0].GRPCAddr())
	_, err := clus.Client(1).Put(context.TODO(), "foo", "bar")
	if err != nil && !strings.Contains(err.Error(), "closing") {
		t.Fatal(err)
	}
}
