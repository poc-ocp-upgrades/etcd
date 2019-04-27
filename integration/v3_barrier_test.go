package integration

import (
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/contrib/recipes"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestBarrierSingleNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	testBarrier(t, 5, func() *clientv3.Client {
		return clus.clients[0]
	})
}
func TestBarrierMultiNode(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	testBarrier(t, 5, func() *clientv3.Client {
		return clus.RandClient()
	})
}
func testBarrier(t *testing.T, waiters int, chooseClient func() *clientv3.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := recipe.NewBarrier(chooseClient(), "test-barrier")
	if err := b.Hold(); err != nil {
		t.Fatalf("could not hold barrier (%v)", err)
	}
	if err := b.Hold(); err == nil {
		t.Fatalf("able to double-hold barrier")
	}
	donec := make(chan struct{})
	for i := 0; i < waiters; i++ {
		go func() {
			br := recipe.NewBarrier(chooseClient(), "test-barrier")
			if err := br.Wait(); err != nil {
				t.Fatalf("could not wait on barrier (%v)", err)
			}
			donec <- struct{}{}
		}()
	}
	select {
	case <-donec:
		t.Fatalf("barrier did not wait")
	default:
	}
	if err := b.Release(); err != nil {
		t.Fatalf("could not release barrier (%v)", err)
	}
	timerC := time.After(time.Duration(waiters*100) * time.Millisecond)
	for i := 0; i < waiters; i++ {
		select {
		case <-timerC:
			t.Fatalf("barrier timed out")
		case <-donec:
		}
	}
}
