package integration

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
	"google.golang.org/grpc"
)

func TestLeaseNotFoundError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	kv := clus.RandClient()
	_, err := kv.Put(context.TODO(), "foo", "bar", clientv3.WithLease(clientv3.LeaseID(500)))
	if err != rpctypes.ErrLeaseNotFound {
		t.Fatalf("expected %v, got %v", rpctypes.ErrLeaseNotFound, err)
	}
}
func TestLeaseGrant(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	lapi := clus.RandClient()
	kv := clus.RandClient()
	_, merr := lapi.Grant(context.Background(), clientv3.MaxLeaseTTL+1)
	if merr != rpctypes.ErrLeaseTTLTooLarge {
		t.Fatalf("err = %v, want %v", merr, rpctypes.ErrLeaseTTLTooLarge)
	}
	resp, err := lapi.Grant(context.Background(), 10)
	if err != nil {
		t.Errorf("failed to create lease %v", err)
	}
	_, err = kv.Put(context.TODO(), "foo", "bar", clientv3.WithLease(resp.ID))
	if err != nil {
		t.Fatalf("failed to create key with lease %v", err)
	}
}
func TestLeaseRevoke(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	lapi := clus.RandClient()
	kv := clus.RandClient()
	resp, err := lapi.Grant(context.Background(), 10)
	if err != nil {
		t.Errorf("failed to create lease %v", err)
	}
	_, err = lapi.Revoke(context.Background(), clientv3.LeaseID(resp.ID))
	if err != nil {
		t.Errorf("failed to revoke lease %v", err)
	}
	_, err = kv.Put(context.TODO(), "foo", "bar", clientv3.WithLease(resp.ID))
	if err != rpctypes.ErrLeaseNotFound {
		t.Fatalf("err = %v, want %v", err, rpctypes.ErrLeaseNotFound)
	}
}
func TestLeaseKeepAliveOnce(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	lapi := clus.RandClient()
	resp, err := lapi.Grant(context.Background(), 10)
	if err != nil {
		t.Errorf("failed to create lease %v", err)
	}
	_, err = lapi.KeepAliveOnce(context.Background(), resp.ID)
	if err != nil {
		t.Errorf("failed to keepalive lease %v", err)
	}
	_, err = lapi.KeepAliveOnce(context.Background(), clientv3.LeaseID(0))
	if err != rpctypes.ErrLeaseNotFound {
		t.Errorf("expected %v, got %v", rpctypes.ErrLeaseNotFound, err)
	}
}
func TestLeaseKeepAlive(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	lapi := clus.Client(0)
	clus.TakeClient(0)
	resp, err := lapi.Grant(context.Background(), 10)
	if err != nil {
		t.Errorf("failed to create lease %v", err)
	}
	rc, kerr := lapi.KeepAlive(context.Background(), resp.ID)
	if kerr != nil {
		t.Errorf("failed to keepalive lease %v", kerr)
	}
	kresp, ok := <-rc
	if !ok {
		t.Errorf("chan is closed, want not closed")
	}
	if kresp.ID != resp.ID {
		t.Errorf("ID = %x, want %x", kresp.ID, resp.ID)
	}
	lapi.Close()
	_, ok = <-rc
	if ok {
		t.Errorf("chan is not closed, want lease Close() closes chan")
	}
}
func TestLeaseKeepAliveOneSecond(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	resp, err := cli.Grant(context.Background(), 1)
	if err != nil {
		t.Errorf("failed to create lease %v", err)
	}
	rc, kerr := cli.KeepAlive(context.Background(), resp.ID)
	if kerr != nil {
		t.Errorf("failed to keepalive lease %v", kerr)
	}
	for i := 0; i < 3; i++ {
		if _, ok := <-rc; !ok {
			t.Errorf("chan is closed, want not closed")
		}
	}
}
func TestLeaseKeepAliveHandleFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Skip("test it when we have a cluster client")
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	lapi := clus.RandClient()
	resp, err := lapi.Grant(context.Background(), 10)
	if err != nil {
		t.Errorf("failed to create lease %v", err)
	}
	rc, kerr := lapi.KeepAlive(context.Background(), resp.ID)
	if kerr != nil {
		t.Errorf("failed to keepalive lease %v", kerr)
	}
	kresp := <-rc
	if kresp.ID != resp.ID {
		t.Errorf("ID = %x, want %x", kresp.ID, resp.ID)
	}
	clus.Members[0].Stop(t)
	select {
	case <-rc:
		t.Fatalf("unexpected keepalive")
	case <-time.After(10*time.Second/3 + 1):
	}
	clus.Members[0].Restart(t)
	kresp = <-rc
	if kresp.ID != resp.ID {
		t.Errorf("ID = %x, want %x", kresp.ID, resp.ID)
	}
	lapi.Close()
	_, ok := <-rc
	if ok {
		t.Errorf("chan is not closed, want lease Close() closes chan")
	}
}

type leaseCh struct {
	lid	clientv3.LeaseID
	ch	<-chan *clientv3.LeaseKeepAliveResponse
}

func TestLeaseKeepAliveNotFound(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.RandClient()
	lchs := []leaseCh{}
	for i := 0; i < 3; i++ {
		resp, rerr := cli.Grant(context.TODO(), 5)
		if rerr != nil {
			t.Fatal(rerr)
		}
		kach, kaerr := cli.KeepAlive(context.Background(), resp.ID)
		if kaerr != nil {
			t.Fatal(kaerr)
		}
		lchs = append(lchs, leaseCh{resp.ID, kach})
	}
	if _, err := cli.Revoke(context.TODO(), lchs[1].lid); err != nil {
		t.Fatal(err)
	}
	<-lchs[0].ch
	if _, ok := <-lchs[0].ch; !ok {
		t.Fatalf("closed keepalive on wrong lease")
	}
	timec := time.After(5 * time.Second)
	for range lchs[1].ch {
		select {
		case <-timec:
			t.Fatalf("revoke did not close keep alive")
		default:
		}
	}
}
func TestLeaseGrantErrConnClosed(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	clus.TakeClient(0)
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		_, err := cli.Grant(context.TODO(), 5)
		if err != nil && err != grpc.ErrClientConnClosing && err != context.Canceled {
			t.Fatalf("expected %v or %v, got %v", grpc.ErrClientConnClosing, context.Canceled, err)
		}
	}()
	if err := cli.Close(); err != nil {
		t.Fatal(err)
	}
	select {
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("le.Grant took too long")
	case <-donec:
	}
}
func TestLeaseKeepAliveFullResponseQueue(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	lapi := clus.Client(0)
	lresp, err := lapi.Grant(context.Background(), 30)
	if err != nil {
		t.Fatalf("failed to create lease %v", err)
	}
	id := lresp.ID
	old := clientv3.LeaseResponseChSize
	defer func() {
		clientv3.LeaseResponseChSize = old
	}()
	clientv3.LeaseResponseChSize = 0
	_, err = lapi.KeepAlive(context.Background(), id)
	if err != nil {
		t.Fatalf("failed to keepalive lease %v", err)
	}
	time.Sleep(3 * time.Second)
	tr, terr := lapi.TimeToLive(context.Background(), id)
	if terr != nil {
		t.Fatalf("failed to get lease information %v", terr)
	}
	if tr.TTL >= 29 {
		t.Errorf("unexpected kept-alive lease TTL %d", tr.TTL)
	}
}
func TestLeaseGrantNewAfterClose(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	clus.TakeClient(0)
	if err := cli.Close(); err != nil {
		t.Fatal(err)
	}
	donec := make(chan struct{})
	go func() {
		if _, err := cli.Grant(context.TODO(), 5); err != context.Canceled && err != grpc.ErrClientConnClosing {
			t.Fatalf("expected %v or %v, got %v", err != context.Canceled, grpc.ErrClientConnClosing, err)
		}
		close(donec)
	}()
	select {
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("le.Grant took too long")
	case <-donec:
	}
}
func TestLeaseRevokeNewAfterClose(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		t.Fatal(err)
	}
	leaseID := resp.ID
	clus.TakeClient(0)
	if err := cli.Close(); err != nil {
		t.Fatal(err)
	}
	donec := make(chan struct{})
	go func() {
		if _, err := cli.Revoke(context.TODO(), leaseID); err != context.Canceled && err != grpc.ErrClientConnClosing {
			t.Fatalf("expected %v or %v, got %v", err != context.Canceled, grpc.ErrClientConnClosing, err)
		}
		close(donec)
	}()
	select {
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("le.Revoke took too long")
	case <-donec:
	}
}
func TestLeaseKeepAliveCloseAfterDisconnectRevoke(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	resp, err := cli.Grant(context.Background(), 10)
	if err != nil {
		t.Fatal(err)
	}
	rc, kerr := cli.KeepAlive(context.Background(), resp.ID)
	if kerr != nil {
		t.Fatal(kerr)
	}
	kresp := <-rc
	if kresp.ID != resp.ID {
		t.Fatalf("ID = %x, want %x", kresp.ID, resp.ID)
	}
	clus.Members[0].Stop(t)
	time.Sleep(time.Second)
	clus.WaitLeader(t)
	if _, err := clus.Client(1).Revoke(context.TODO(), resp.ID); err != nil {
		t.Fatal(err)
	}
	clus.Members[0].Restart(t)
	timer := time.After(time.Duration(kresp.TTL) * time.Second)
	for kresp != nil {
		select {
		case kresp = <-rc:
		case <-timer:
			t.Fatalf("keepalive channel did not close")
		}
	}
}
func TestLeaseKeepAliveInitTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	resp, err := cli.Grant(context.Background(), 5)
	if err != nil {
		t.Fatal(err)
	}
	clus.Members[0].Stop(t)
	rc, kerr := cli.KeepAlive(context.Background(), resp.ID)
	if kerr != nil {
		t.Fatal(kerr)
	}
	select {
	case ka, ok := <-rc:
		if ok {
			t.Fatalf("unexpected keepalive %v, expected closed channel", ka)
		}
	case <-time.After(10 * time.Second):
		t.Fatalf("keepalive channel did not close")
	}
	clus.Members[0].Restart(t)
}
func TestLeaseKeepAliveTTLTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	resp, err := cli.Grant(context.Background(), 5)
	if err != nil {
		t.Fatal(err)
	}
	rc, kerr := cli.KeepAlive(context.Background(), resp.ID)
	if kerr != nil {
		t.Fatal(kerr)
	}
	if kresp := <-rc; kresp.ID != resp.ID {
		t.Fatalf("ID = %x, want %x", kresp.ID, resp.ID)
	}
	clus.Members[0].Stop(t)
	select {
	case ka, ok := <-rc:
		if ok {
			t.Fatalf("unexpected keepalive %v, expected closed channel", ka)
		}
	case <-time.After(10 * time.Second):
		t.Fatalf("keepalive channel did not close")
	}
	clus.Members[0].Restart(t)
}
func TestLeaseTimeToLive(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	c := clus.RandClient()
	lapi := c
	resp, err := lapi.Grant(context.Background(), 10)
	if err != nil {
		t.Errorf("failed to create lease %v", err)
	}
	kv := clus.RandClient()
	keys := []string{"foo1", "foo2"}
	for i := range keys {
		if _, err = kv.Put(context.TODO(), keys[i], "bar", clientv3.WithLease(resp.ID)); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := c.Get(context.TODO(), "abc"); err != nil {
		t.Fatal(err)
	}
	lresp, lerr := lapi.TimeToLive(context.Background(), resp.ID, clientv3.WithAttachedKeys())
	if lerr != nil {
		t.Fatal(lerr)
	}
	if lresp.ID != resp.ID {
		t.Fatalf("leaseID expected %d, got %d", resp.ID, lresp.ID)
	}
	if lresp.GrantedTTL != int64(10) {
		t.Fatalf("GrantedTTL expected %d, got %d", 10, lresp.GrantedTTL)
	}
	if lresp.TTL == 0 || lresp.TTL > lresp.GrantedTTL {
		t.Fatalf("unexpected TTL %d (granted %d)", lresp.TTL, lresp.GrantedTTL)
	}
	ks := make([]string, len(lresp.Keys))
	for i := range lresp.Keys {
		ks[i] = string(lresp.Keys[i])
	}
	sort.Strings(ks)
	if !reflect.DeepEqual(ks, keys) {
		t.Fatalf("keys expected %v, got %v", keys, ks)
	}
	lresp, lerr = lapi.TimeToLive(context.Background(), resp.ID)
	if lerr != nil {
		t.Fatal(lerr)
	}
	if len(lresp.Keys) != 0 {
		t.Fatalf("unexpected keys %+v", lresp.Keys)
	}
}
func TestLeaseTimeToLiveLeaseNotFound(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.RandClient()
	resp, err := cli.Grant(context.Background(), 10)
	if err != nil {
		t.Errorf("failed to create lease %v", err)
	}
	_, err = cli.Revoke(context.Background(), resp.ID)
	if err != nil {
		t.Errorf("failed to Revoke lease %v", err)
	}
	lresp, err := cli.TimeToLive(context.Background(), resp.ID)
	if err != nil {
		t.Fatalf("expected err to be nil")
	}
	if lresp == nil {
		t.Fatalf("expected lresp not to be nil")
	}
	if lresp.ResponseHeader == nil {
		t.Fatalf("expected ResponseHeader not to be nil")
	}
	if lresp.ID != resp.ID {
		t.Fatalf("expected Lease ID %v, but got %v", resp.ID, lresp.ID)
	}
	if lresp.TTL != -1 {
		t.Fatalf("expected TTL %v, but got %v", lresp.TTL, lresp.TTL)
	}
}
func TestLeaseLeases(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.RandClient()
	ids := []clientv3.LeaseID{}
	for i := 0; i < 5; i++ {
		resp, err := cli.Grant(context.Background(), 10)
		if err != nil {
			t.Errorf("failed to create lease %v", err)
		}
		ids = append(ids, resp.ID)
	}
	resp, err := cli.Leases(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Leases) != 5 {
		t.Fatalf("len(resp.Leases) expected 5, got %d", len(resp.Leases))
	}
	for i := range resp.Leases {
		if ids[i] != resp.Leases[i].ID {
			t.Fatalf("#%d: lease ID expected %d, got %d", i, ids[i], resp.Leases[i].ID)
		}
	}
}
func TestLeaseRenewLostQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	r, err := cli.Grant(context.TODO(), 4)
	if err != nil {
		t.Fatal(err)
	}
	kctx, kcancel := context.WithCancel(context.Background())
	defer kcancel()
	ka, err := cli.KeepAlive(kctx, r.ID)
	if err != nil {
		t.Fatal(err)
	}
	<-ka
	lastKa := time.Now()
	clus.Members[1].Stop(t)
	clus.Members[2].Stop(t)
	time.Sleep(time.Duration(r.TTL-2) * time.Second)
	clus.Members[1].Restart(t)
	clus.Members[2].Restart(t)
	if time.Since(lastKa) > time.Duration(r.TTL)*time.Second {
		t.Skip("waited too long for server stop and restart")
	}
	select {
	case _, ok := <-ka:
		if !ok {
			t.Fatalf("keepalive closed")
		}
	case <-time.After(time.Duration(r.TTL) * time.Second):
		t.Fatalf("timed out waiting for keepalive")
	}
}
func TestLeaseKeepAliveLoopExit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	ctx := context.Background()
	cli := clus.Client(0)
	clus.TakeClient(0)
	resp, err := cli.Grant(ctx, 5)
	if err != nil {
		t.Fatal(err)
	}
	cli.Close()
	_, err = cli.KeepAlive(ctx, resp.ID)
	if _, ok := err.(clientv3.ErrKeepAliveHalted); !ok {
		t.Fatalf("expected %T, got %v(%T)", clientv3.ErrKeepAliveHalted{}, err, err)
	}
}
func TestV3LeaseFailureOverlap(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 2})
	defer clus.Terminate(t)
	numReqs := 5
	cli := clus.Client(0)
	updown := func(i int) error {
		sess, err := concurrency.NewSession(cli)
		if err != nil {
			return err
		}
		ch := make(chan struct{})
		go func() {
			defer close(ch)
			sess.Close()
		}()
		select {
		case <-ch:
		case <-time.After(time.Minute / 4):
			t.Fatalf("timeout %d", i)
		}
		return nil
	}
	var wg sync.WaitGroup
	mkReqs := func(n int) {
		wg.Add(numReqs)
		for i := 0; i < numReqs; i++ {
			go func() {
				defer wg.Done()
				err := updown(n)
				if err == nil || err == rpctypes.ErrTimeoutDueToConnectionLost {
					return
				}
				t.Fatal(err)
			}()
		}
	}
	mkReqs(1)
	clus.Members[1].Stop(t)
	mkReqs(2)
	time.Sleep(time.Second)
	mkReqs(3)
	clus.Members[1].Restart(t)
	mkReqs(4)
	wg.Wait()
}
func TestLeaseWithRequireLeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 2})
	defer clus.Terminate(t)
	c := clus.Client(0)
	lid1, err1 := c.Grant(context.TODO(), 60)
	if err1 != nil {
		t.Fatal(err1)
	}
	lid2, err2 := c.Grant(context.TODO(), 60)
	if err2 != nil {
		t.Fatal(err2)
	}
	kaReqLeader, kerr1 := c.KeepAlive(clientv3.WithRequireLeader(context.TODO()), lid1.ID)
	if kerr1 != nil {
		t.Fatal(kerr1)
	}
	kaWait, kerr2 := c.KeepAlive(context.TODO(), lid2.ID)
	if kerr2 != nil {
		t.Fatal(kerr2)
	}
	select {
	case <-kaReqLeader:
	case <-time.After(5 * time.Second):
		t.Fatalf("require leader first keep-alive timed out")
	}
	select {
	case <-kaWait:
	case <-time.After(5 * time.Second):
		t.Fatalf("leader not required first keep-alive timed out")
	}
	clus.Members[1].Stop(t)
	time.Sleep(100 * time.Millisecond)
	for len(kaReqLeader) > 0 {
		<-kaReqLeader
	}
	select {
	case resp, ok := <-kaReqLeader:
		if ok {
			t.Fatalf("expected closed require leader, got response %+v", resp)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("keepalive with require leader took too long to close")
	}
	select {
	case _, ok := <-kaWait:
		if !ok {
			t.Fatalf("got closed channel with no require leader, expected non-closed")
		}
	case <-time.After(10 * time.Millisecond):
	}
}
