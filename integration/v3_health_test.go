package integration

import (
	"context"
	"testing"
	"github.com/coreos/etcd/pkg/testutil"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func TestHealthCheck(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := healthpb.NewHealthClient(clus.RandClient().ActiveConnection())
	resp, err := cli.Check(context.TODO(), &healthpb.HealthCheckRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Status != healthpb.HealthCheckResponse_SERVING {
		t.Fatalf("status expected %s, got %s", healthpb.HealthCheckResponse_SERVING, resp.Status)
	}
}
