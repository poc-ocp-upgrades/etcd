package integration

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/mvcc"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestV3StorageQuotaApply(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testutil.AfterTest(t)
	quotasize := int64(16 * os.Getpagesize())
	clus := NewClusterV3(t, &ClusterConfig{Size: 2})
	defer clus.Terminate(t)
	kvc0 := toGRPC(clus.Client(0)).KV
	kvc1 := toGRPC(clus.Client(1)).KV
	clus.Members[0].QuotaBackendBytes = quotasize
	clus.Members[0].Stop(t)
	clus.Members[0].Restart(t)
	clus.waitLeader(t, clus.Members)
	waitForRestart(t, kvc0)
	key := []byte("abc")
	smallbuf := make([]byte, 1024)
	_, serr := kvc0.Put(context.TODO(), &pb.PutRequest{Key: key, Value: smallbuf})
	if serr != nil {
		t.Fatal(serr)
	}
	bigbuf := make([]byte, quotasize)
	_, err := kvc1.Put(context.TODO(), &pb.PutRequest{Key: key, Value: bigbuf})
	if err != nil {
		t.Fatal(err)
	}
	_, err = kvc0.Range(context.TODO(), &pb.RangeRequest{Key: []byte("foo")})
	if err != nil {
		t.Fatal(err)
	}
	stopc := time.After(5 * time.Second)
	for {
		req := &pb.AlarmRequest{Action: pb.AlarmRequest_GET}
		resp, aerr := clus.Members[0].s.Alarm(context.TODO(), req)
		if aerr != nil {
			t.Fatal(aerr)
		}
		if len(resp.Alarms) != 0 {
			break
		}
		select {
		case <-stopc:
			t.Fatalf("timed out waiting for alarm")
		case <-time.After(10 * time.Millisecond):
		}
	}
	if _, err := kvc0.Put(context.TODO(), &pb.PutRequest{Key: key, Value: smallbuf}); err == nil {
		t.Fatalf("past-quota instance should reject put")
	}
	if _, err := kvc1.Put(context.TODO(), &pb.PutRequest{Key: key, Value: smallbuf}); err == nil {
		t.Fatalf("past-quota instance should reject put")
	}
	clus.Members[1].Stop(t)
	clus.Members[1].Restart(t)
	clus.waitLeader(t, clus.Members)
	if _, err := kvc1.Put(context.TODO(), &pb.PutRequest{Key: key, Value: smallbuf}); err == nil {
		t.Fatalf("alarmed instance should reject put after reset")
	}
}
func TestV3AlarmDeactivate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := NewClusterV3(t, &ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	kvc := toGRPC(clus.RandClient()).KV
	mt := toGRPC(clus.RandClient()).Maintenance
	alarmReq := &pb.AlarmRequest{MemberID: 123, Action: pb.AlarmRequest_ACTIVATE, Alarm: pb.AlarmType_NOSPACE}
	if _, err := mt.Alarm(context.TODO(), alarmReq); err != nil {
		t.Fatal(err)
	}
	key := []byte("abc")
	smallbuf := make([]byte, 512)
	_, err := kvc.Put(context.TODO(), &pb.PutRequest{Key: key, Value: smallbuf})
	if err == nil && !eqErrGRPC(err, rpctypes.ErrGRPCNoSpace) {
		t.Fatalf("put got %v, expected %v", err, rpctypes.ErrGRPCNoSpace)
	}
	alarmReq.Action = pb.AlarmRequest_DEACTIVATE
	if _, err = mt.Alarm(context.TODO(), alarmReq); err != nil {
		t.Fatal(err)
	}
	if _, err = kvc.Put(context.TODO(), &pb.PutRequest{Key: key, Value: smallbuf}); err != nil {
		t.Fatal(err)
	}
}

type fakeConsistentIndex struct{ rev uint64 }

func (f *fakeConsistentIndex) ConsistentIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return f.rev
}
func TestV3CorruptAlarm(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if _, err := clus.Client(0).Put(context.TODO(), "k", "v"); err != nil {
				t.Fatal(err)
			}
		}()
	}
	wg.Wait()
	clus.Members[0].Stop(t)
	fp := filepath.Join(clus.Members[0].DataDir, "member", "snap", "db")
	be := backend.NewDefaultBackend(fp)
	s := mvcc.NewStore(be, nil, &fakeConsistentIndex{13})
	s.Put([]byte("abc"), []byte("def"), 0)
	s.Put([]byte("xyz"), []byte("123"), 0)
	s.Compact(5)
	s.Commit()
	s.Close()
	be.Close()
	if _, err := clus.Client(1).Get(context.TODO(), "k"); err != nil {
		t.Fatal(err)
	}
	clus.Client(1).Put(context.TODO(), "xyz", "321")
	clus.Client(1).Put(context.TODO(), "abc", "fed")
	clus.Members[1].Stop(t)
	clus.Members[2].Stop(t)
	for _, m := range clus.Members {
		m.CorruptCheckTime = time.Second
		m.Restart(t)
	}
	resp0, err0 := clus.Client(0).Get(context.TODO(), "abc")
	if err0 != nil {
		t.Fatal(err0)
	}
	resp1, err1 := clus.Client(1).Get(context.TODO(), "abc")
	if err1 != nil {
		t.Fatal(err1)
	}
	if resp0.Kvs[0].ModRevision == resp1.Kvs[0].ModRevision {
		t.Fatalf("matching ModRevision values")
	}
	for i := 0; i < 5; i++ {
		presp, perr := clus.Client(0).Put(context.TODO(), "abc", "aaa")
		if perr != nil {
			if !eqErrGRPC(perr, rpctypes.ErrCorrupt) {
				t.Fatalf("expected %v, got %+v (%v)", rpctypes.ErrCorrupt, presp, perr)
			} else {
				return
			}
		}
		time.Sleep(time.Second)
	}
	t.Fatalf("expected error %v after %s", rpctypes.ErrCorrupt, 5*time.Second)
}
