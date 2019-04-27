package grpcproxy

import (
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/naming"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
	gnaming "google.golang.org/grpc/naming"
)

func TestRegister(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	paddr := clus.Members[0].GRPCAddr()
	testPrefix := "test-name"
	wa := createWatcher(t, cli, testPrefix)
	ups, err := wa.Next()
	if err != nil {
		t.Fatal(err)
	}
	if len(ups) != 0 {
		t.Fatalf("len(ups) expected 0, got %d (%v)", len(ups), ups)
	}
	donec := Register(cli, testPrefix, paddr, 5)
	ups, err = wa.Next()
	if err != nil {
		t.Fatal(err)
	}
	if len(ups) != 1 {
		t.Fatalf("len(ups) expected 1, got %d (%v)", len(ups), ups)
	}
	if ups[0].Addr != paddr {
		t.Fatalf("ups[0].Addr expected %q, got %q", paddr, ups[0].Addr)
	}
	cli.Close()
	clus.TakeClient(0)
	select {
	case <-donec:
	case <-time.After(5 * time.Second):
		t.Fatal("donec 'register' did not return in time")
	}
}
func createWatcher(t *testing.T, c *clientv3.Client, prefix string) gnaming.Watcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	gr := &naming.GRPCResolver{Client: c}
	watcher, err := gr.Resolve(prefix)
	if err != nil {
		t.Fatalf("failed to resolve %q (%v)", prefix, err)
	}
	return watcher
}
