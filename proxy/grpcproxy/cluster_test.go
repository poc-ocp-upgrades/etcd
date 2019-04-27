package grpcproxy

import (
	"context"
	"net"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
	"google.golang.org/grpc"
)

func TestClusterProxyMemberList(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cts := newClusterProxyServer([]string{clus.Members[0].GRPCAddr()}, t)
	defer cts.close(t)
	cfg := clientv3.Config{Endpoints: []string{cts.caddr}, DialTimeout: 5 * time.Second}
	client, err := clientv3.New(cfg)
	if err != nil {
		t.Fatalf("err %v, want nil", err)
	}
	defer client.Close()
	time.Sleep(time.Second)
	var mresp *clientv3.MemberListResponse
	mresp, err = client.Cluster.MemberList(context.Background())
	if err != nil {
		t.Fatalf("err %v, want nil", err)
	}
	if len(mresp.Members) != 1 {
		t.Fatalf("len(mresp.Members) expected 1, got %d (%+v)", len(mresp.Members), mresp.Members)
	}
	if len(mresp.Members[0].ClientURLs) != 1 {
		t.Fatalf("len(mresp.Members[0].ClientURLs) expected 1, got %d (%+v)", len(mresp.Members[0].ClientURLs), mresp.Members[0].ClientURLs[0])
	}
	if mresp.Members[0].ClientURLs[0] != cts.caddr {
		t.Fatalf("mresp.Members[0].ClientURLs[0] expected %q, got %q", cts.caddr, mresp.Members[0].ClientURLs[0])
	}
}

type clusterproxyTestServer struct {
	cp	pb.ClusterServer
	c	*clientv3.Client
	server	*grpc.Server
	l	net.Listener
	donec	<-chan struct{}
	caddr	string
}

func (cts *clusterproxyTestServer) close(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cts.server.Stop()
	cts.l.Close()
	cts.c.Close()
	select {
	case <-cts.donec:
		return
	case <-time.After(5 * time.Second):
		t.Fatalf("register-loop took too long to return")
	}
}
func newClusterProxyServer(endpoints []string, t *testing.T) *clusterproxyTestServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := clientv3.Config{Endpoints: endpoints, DialTimeout: 5 * time.Second}
	client, err := clientv3.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	cts := &clusterproxyTestServer{c: client}
	cts.l, err = net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	var opts []grpc.ServerOption
	cts.server = grpc.NewServer(opts...)
	servec := make(chan struct{})
	go func() {
		<-servec
		cts.server.Serve(cts.l)
	}()
	Register(client, "test-prefix", cts.l.Addr().String(), 7)
	cts.cp, cts.donec = NewClusterProxy(client, cts.l.Addr().String(), "test-prefix")
	cts.caddr = cts.l.Addr().String()
	pb.RegisterClusterServer(cts.server, cts.cp)
	close(servec)
	time.Sleep(500 * time.Millisecond)
	return cts
}
