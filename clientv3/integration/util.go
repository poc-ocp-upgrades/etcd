package integration

import (
	"context"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
)

func mustWaitPinReady(t *testing.T, cli *clientv3.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		t.Fatal(err)
	}
}
