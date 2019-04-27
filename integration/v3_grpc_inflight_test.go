package integration

import (
	"context"
	"sync"
	"testing"
	"time"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/testutil"
	"google.golang.org/grpc"
)

func TestV3MaintenanceDefragmentInflightRange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.RandClient()
	kvc := toGRPC(cli).KV
	if _, err := kvc.Put(context.Background(), &pb.PutRequest{Key: []byte("foo"), Value: []byte("bar")}); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		kvc.Range(ctx, &pb.RangeRequest{Key: []byte("foo")})
	}()
	mvc := toGRPC(cli).Maintenance
	mvc.Defragment(context.Background(), &pb.DefragmentRequest{})
	cancel()
	<-donec
}
func TestV3KVInflightRangeRequests(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.RandClient()
	kvc := toGRPC(cli).KV
	if _, err := kvc.Put(context.Background(), &pb.PutRequest{Key: []byte("foo"), Value: []byte("bar")}); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	reqN := 10
	var wg sync.WaitGroup
	wg.Add(reqN)
	for i := 0; i < reqN; i++ {
		go func() {
			defer wg.Done()
			_, err := kvc.Range(ctx, &pb.RangeRequest{Key: []byte("foo"), Serializable: true}, grpc.FailFast(false))
			if err != nil {
				if err != nil && rpctypes.ErrorDesc(err) != context.Canceled.Error() {
					t.Fatalf("inflight request should be canceld with %v, got %v", context.Canceled, err)
				}
			}
		}()
	}
	clus.Members[0].Stop(t)
	cancel()
	wg.Wait()
}
