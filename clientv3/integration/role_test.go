package integration

import (
	"context"
	"testing"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestRoleError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	authapi := clus.RandClient()
	_, err := authapi.RoleAdd(context.TODO(), "test-role")
	if err != nil {
		t.Fatal(err)
	}
	_, err = authapi.RoleAdd(context.TODO(), "test-role")
	if err != rpctypes.ErrRoleAlreadyExist {
		t.Fatalf("expected %v, got %v", rpctypes.ErrRoleAlreadyExist, err)
	}
}
