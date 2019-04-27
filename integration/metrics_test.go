package integration

import (
	"context"
	"strconv"
	"testing"
	"time"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestMetricDbSizeBoot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	v, err := clus.Members[0].Metric("etcd_debugging_mvcc_db_total_size_in_bytes")
	if err != nil {
		t.Fatal(err)
	}
	if v == "0" {
		t.Fatalf("expected non-zero, got %q", v)
	}
}
func TestMetricDbSizeDefrag(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	kvc := toGRPC(clus.Client(0)).KV
	mc := toGRPC(clus.Client(0)).Maintenance
	numPuts := 25
	putreq := &pb.PutRequest{Key: []byte("k"), Value: make([]byte, 4096)}
	for i := 0; i < numPuts; i++ {
		if _, err := kvc.Put(context.TODO(), putreq); err != nil {
			t.Fatal(err)
		}
	}
	time.Sleep(500 * time.Millisecond)
	beforeDefrag, err := clus.Members[0].Metric("etcd_debugging_mvcc_db_total_size_in_bytes")
	if err != nil {
		t.Fatal(err)
	}
	bv, err := strconv.Atoi(beforeDefrag)
	if err != nil {
		t.Fatal(err)
	}
	if expected := numPuts * len(putreq.Value); bv < expected {
		t.Fatalf("expected db size greater than %d, got %d", expected, bv)
	}
	creq := &pb.CompactionRequest{Revision: int64(numPuts), Physical: true}
	if _, err := kvc.Compact(context.TODO(), creq); err != nil {
		t.Fatal(err)
	}
	mc.Defragment(context.TODO(), &pb.DefragmentRequest{})
	afterDefrag, err := clus.Members[0].Metric("etcd_debugging_mvcc_db_total_size_in_bytes")
	if err != nil {
		t.Fatal(err)
	}
	av, err := strconv.Atoi(afterDefrag)
	if err != nil {
		t.Fatal(err)
	}
	if bv <= av {
		t.Fatalf("expected less than %d, got %d after defrag", bv, av)
	}
}
