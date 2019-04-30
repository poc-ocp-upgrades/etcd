package integration

import (
	"context"
	"testing"
	"time"
	"go.etcd.io/etcd/clientv3"
)

func mustWaitPinReady(t *testing.T, cli *clientv3.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		t.Fatal(err)
	}
}
