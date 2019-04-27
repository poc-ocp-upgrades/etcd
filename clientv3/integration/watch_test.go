package integration

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/integration"
	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/coreos/etcd/pkg/testutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type watcherTest func(*testing.T, *watchctx)
type watchctx struct {
	clus		*integration.ClusterV3
	w		clientv3.Watcher
	kv		clientv3.KV
	wclientMember	int
	kvMember	int
	ch		clientv3.WatchChan
}

func runWatchTest(t *testing.T, f watcherTest) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	wclientMember := rand.Intn(3)
	w := clus.Client(wclientMember).Watcher
	kvMember := rand.Intn(3)
	for kvMember == wclientMember {
		kvMember = rand.Intn(3)
	}
	kv := clus.Client(kvMember).KV
	wctx := &watchctx{clus, w, kv, wclientMember, kvMember, nil}
	f(t, wctx)
}
func TestWatchMultiWatcher(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	runWatchTest(t, testWatchMultiWatcher)
}
func testWatchMultiWatcher(t *testing.T, wctx *watchctx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	numKeyUpdates := 4
	keys := []string{"foo", "bar", "baz"}
	donec := make(chan struct{})
	readyc := make(chan struct{})
	for _, k := range keys {
		go func(key string) {
			ch := wctx.w.Watch(context.TODO(), key)
			if ch == nil {
				t.Fatalf("expected watcher channel, got nil")
			}
			readyc <- struct{}{}
			for i := 0; i < numKeyUpdates; i++ {
				resp, ok := <-ch
				if !ok {
					t.Fatalf("watcher unexpectedly closed")
				}
				v := fmt.Sprintf("%s-%d", key, i)
				gotv := string(resp.Events[0].Kv.Value)
				if gotv != v {
					t.Errorf("#%d: got %s, wanted %s", i, gotv, v)
				}
			}
			donec <- struct{}{}
		}(k)
	}
	go func() {
		prefixc := wctx.w.Watch(context.TODO(), "b", clientv3.WithPrefix())
		if prefixc == nil {
			t.Fatalf("expected watcher channel, got nil")
		}
		readyc <- struct{}{}
		evs := []*clientv3.Event{}
		for i := 0; i < numKeyUpdates*2; i++ {
			resp, ok := <-prefixc
			if !ok {
				t.Fatalf("watcher unexpectedly closed")
			}
			evs = append(evs, resp.Events...)
		}
		expected := []string{}
		bkeys := []string{"bar", "baz"}
		for _, k := range bkeys {
			for i := 0; i < numKeyUpdates; i++ {
				expected = append(expected, fmt.Sprintf("%s-%d", k, i))
			}
		}
		got := []string{}
		for _, ev := range evs {
			got = append(got, string(ev.Kv.Value))
		}
		sort.Strings(got)
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("got %v, expected %v", got, expected)
		}
		select {
		case resp, ok := <-prefixc:
			if !ok {
				t.Fatalf("watcher unexpectedly closed")
			}
			t.Fatalf("unexpected event %+v", resp)
		case <-time.After(time.Second):
		}
		donec <- struct{}{}
	}()
	for i := 0; i < len(keys)+1; i++ {
		<-readyc
	}
	ctx := context.TODO()
	for i := 0; i < numKeyUpdates; i++ {
		for _, k := range keys {
			v := fmt.Sprintf("%s-%d", k, i)
			if _, err := wctx.kv.Put(ctx, k, v); err != nil {
				t.Fatal(err)
			}
		}
	}
	for i := 0; i < len(keys)+1; i++ {
		<-donec
	}
}
func TestWatchRange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	runWatchTest(t, testWatchRange)
}
func testWatchRange(t *testing.T, wctx *watchctx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if wctx.ch = wctx.w.Watch(context.TODO(), "a", clientv3.WithRange("c")); wctx.ch == nil {
		t.Fatalf("expected non-nil channel")
	}
	putAndWatch(t, wctx, "a", "a")
	putAndWatch(t, wctx, "b", "b")
	putAndWatch(t, wctx, "bar", "bar")
}
func TestWatchReconnRequest(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	runWatchTest(t, testWatchReconnRequest)
}
func testWatchReconnRequest(t *testing.T, wctx *watchctx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	donec, stopc := make(chan struct{}), make(chan struct{}, 1)
	go func() {
		timer := time.After(2 * time.Second)
		defer close(donec)
		for {
			wctx.clus.Members[wctx.wclientMember].DropConnections()
			select {
			case <-timer:
				return
			case <-stopc:
				return
			default:
			}
		}
	}()
	if wctx.ch = wctx.w.Watch(context.TODO(), "a"); wctx.ch == nil {
		t.Fatalf("expected non-nil channel")
	}
	stopc <- struct{}{}
	<-donec
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	if _, err := wctx.kv.Get(ctx, "_"); err != nil {
		t.Fatal(err)
	}
	cancel()
	putAndWatch(t, wctx, "a", "a")
}
func TestWatchReconnInit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	runWatchTest(t, testWatchReconnInit)
}
func testWatchReconnInit(t *testing.T, wctx *watchctx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if wctx.ch = wctx.w.Watch(context.TODO(), "a"); wctx.ch == nil {
		t.Fatalf("expected non-nil channel")
	}
	wctx.clus.Members[wctx.wclientMember].DropConnections()
	putAndWatch(t, wctx, "a", "a")
}
func TestWatchReconnRunning(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	runWatchTest(t, testWatchReconnRunning)
}
func testWatchReconnRunning(t *testing.T, wctx *watchctx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if wctx.ch = wctx.w.Watch(context.TODO(), "a"); wctx.ch == nil {
		t.Fatalf("expected non-nil channel")
	}
	putAndWatch(t, wctx, "a", "a")
	wctx.clus.Members[wctx.wclientMember].DropConnections()
	putAndWatch(t, wctx, "a", "b")
}
func TestWatchCancelImmediate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	runWatchTest(t, testWatchCancelImmediate)
}
func testWatchCancelImmediate(t *testing.T, wctx *watchctx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	wch := wctx.w.Watch(ctx, "a")
	select {
	case wresp, ok := <-wch:
		if ok {
			t.Fatalf("read wch got %v; expected closed channel", wresp)
		}
	default:
		t.Fatalf("closed watcher channel should not block")
	}
}
func TestWatchCancelInit(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	runWatchTest(t, testWatchCancelInit)
}
func testWatchCancelInit(t *testing.T, wctx *watchctx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithCancel(context.Background())
	if wctx.ch = wctx.w.Watch(ctx, "a"); wctx.ch == nil {
		t.Fatalf("expected non-nil watcher channel")
	}
	cancel()
	select {
	case <-time.After(time.Second):
		t.Fatalf("took too long to cancel")
	case _, ok := <-wctx.ch:
		if ok {
			t.Fatalf("expected watcher channel to close")
		}
	}
}
func TestWatchCancelRunning(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	runWatchTest(t, testWatchCancelRunning)
}
func testWatchCancelRunning(t *testing.T, wctx *watchctx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithCancel(context.Background())
	if wctx.ch = wctx.w.Watch(ctx, "a"); wctx.ch == nil {
		t.Fatalf("expected non-nil watcher channel")
	}
	if _, err := wctx.kv.Put(ctx, "a", "a"); err != nil {
		t.Fatal(err)
	}
	cancel()
	select {
	case <-time.After(time.Second):
		t.Fatalf("took too long to cancel")
	case _, ok := <-wctx.ch:
		if !ok {
			break
		}
		select {
		case <-time.After(time.Second):
			t.Fatalf("took too long to close")
		case v, ok2 := <-wctx.ch:
			if ok2 {
				t.Fatalf("expected watcher channel to close, got %v", v)
			}
		}
	}
}
func putAndWatch(t *testing.T, wctx *watchctx, key, val string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, err := wctx.kv.Put(context.TODO(), key, val); err != nil {
		t.Fatal(err)
	}
	select {
	case <-time.After(5 * time.Second):
		t.Fatalf("watch timed out")
	case v, ok := <-wctx.ch:
		if !ok {
			t.Fatalf("unexpected watch close")
		}
		if string(v.Events[0].Kv.Value) != val {
			t.Fatalf("bad value got %v, wanted %v", v.Events[0].Kv.Value, val)
		}
	}
}
func TestWatchResumeInitRev(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	if _, err := cli.Put(context.TODO(), "b", "2"); err != nil {
		t.Fatal(err)
	}
	if _, err := cli.Put(context.TODO(), "a", "3"); err != nil {
		t.Fatal(err)
	}
	if _, err := cli.Put(context.TODO(), "a", "4"); err != nil {
		t.Fatal(err)
	}
	wch := clus.Client(0).Watch(context.Background(), "a", clientv3.WithRev(1), clientv3.WithCreatedNotify())
	if resp, ok := <-wch; !ok || resp.Header.Revision != 4 {
		t.Fatalf("got (%v, %v), expected create notification rev=4", resp, ok)
	}
	clus.Members[0].DropConnections()
	clus.Members[0].PauseConnections()
	select {
	case resp, ok := <-wch:
		t.Skipf("wch should block, got (%+v, %v); drop not fast enough", resp, ok)
	case <-time.After(100 * time.Millisecond):
	}
	clus.Members[0].UnpauseConnections()
	select {
	case resp, ok := <-wch:
		if !ok {
			t.Fatal("unexpected watch close")
		}
		if len(resp.Events) == 0 {
			t.Fatal("expected event on watch")
		}
		if string(resp.Events[0].Kv.Value) != "3" {
			t.Fatalf("expected value=3, got event %+v", resp.Events[0])
		}
	case <-time.After(5 * time.Second):
		t.Fatal("watch timed out")
	}
}
func TestWatchResumeCompacted(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	w := clus.Client(0)
	wch := w.Watch(context.Background(), "foo", clientv3.WithRev(1))
	select {
	case w := <-wch:
		t.Errorf("unexpected message from wch %v", w)
	default:
	}
	clus.Members[0].Stop(t)
	ticker := time.After(time.Second * 10)
	for clus.WaitLeader(t) <= 0 {
		select {
		case <-ticker:
			t.Fatalf("failed to wait for new leader")
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
	numPuts := 5
	kv := clus.Client(1)
	for i := 0; i < numPuts; i++ {
		if _, err := kv.Put(context.TODO(), "foo", "bar"); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := kv.Compact(context.TODO(), 3); err != nil {
		t.Fatal(err)
	}
	clus.Members[0].Restart(t)
	wRev := int64(2)
	for int(wRev) <= numPuts+1 {
		var wresp clientv3.WatchResponse
		var ok bool
		select {
		case wresp, ok = <-wch:
			if !ok {
				t.Fatalf("expected wresp, but got closed channel")
			}
		case <-time.After(5 * time.Second):
			t.Fatalf("compacted watch timed out")
		}
		for _, ev := range wresp.Events {
			if ev.Kv.ModRevision != wRev {
				t.Fatalf("expected modRev %v, got %+v", wRev, ev)
			}
			wRev++
		}
		if wresp.Err() == nil {
			continue
		}
		if wresp.Err() != rpctypes.ErrCompacted {
			t.Fatalf("wresp.Err() expected %v, got %+v", rpctypes.ErrCompacted, wresp.Err())
		}
		break
	}
	if int(wRev) > numPuts+1 {
		return
	}
	select {
	case wresp, ok := <-wch:
		if ok {
			t.Fatalf("expected closed channel, but got %v", wresp)
		}
	case <-time.After(5 * time.Second):
		t.Fatalf("timed out waiting for channel close")
	}
}
func TestWatchCompactRevision(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	kv := clus.RandClient()
	for i := 0; i < 5; i++ {
		if _, err := kv.Put(context.TODO(), "foo", "bar"); err != nil {
			t.Fatal(err)
		}
	}
	w := clus.RandClient()
	if _, err := kv.Compact(context.TODO(), 4); err != nil {
		t.Fatal(err)
	}
	wch := w.Watch(context.Background(), "foo", clientv3.WithRev(2))
	wresp, ok := <-wch
	if !ok {
		t.Fatalf("expected wresp, but got closed channel")
	}
	if wresp.Err() != rpctypes.ErrCompacted {
		t.Fatalf("wresp.Err() expected %v, but got %v", rpctypes.ErrCompacted, wresp.Err())
	}
	if !wresp.Canceled {
		t.Fatalf("wresp.Canceled expected true, got %+v", wresp)
	}
	if wresp, ok = <-wch; ok {
		t.Fatalf("expected closed channel, but got %v", wresp)
	}
}
func TestWatchWithProgressNotify(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testWatchWithProgressNotify(t, true)
}
func TestWatchWithProgressNotifyNoEvent(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testWatchWithProgressNotify(t, false)
}
func testWatchWithProgressNotify(t *testing.T, watchOnPut bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	oldpi := v3rpc.GetProgressReportInterval()
	v3rpc.SetProgressReportInterval(3 * time.Second)
	pi := 3 * time.Second
	defer func() {
		v3rpc.SetProgressReportInterval(oldpi)
	}()
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	wc := clus.RandClient()
	opts := []clientv3.OpOption{clientv3.WithProgressNotify()}
	if watchOnPut {
		opts = append(opts, clientv3.WithPrefix())
	}
	rch := wc.Watch(context.Background(), "foo", opts...)
	select {
	case resp := <-rch:
		if len(resp.Events) != 0 {
			t.Fatalf("resp.Events expected none, got %+v", resp.Events)
		}
	case <-time.After(2 * pi):
		t.Fatalf("watch response expected in %v, but timed out", pi)
	}
	kvc := clus.RandClient()
	if _, err := kvc.Put(context.TODO(), "foox", "bar"); err != nil {
		t.Fatal(err)
	}
	select {
	case resp := <-rch:
		if resp.Header.Revision != 2 {
			t.Fatalf("resp.Header.Revision expected 2, got %d", resp.Header.Revision)
		}
		if watchOnPut {
			ev := []*clientv3.Event{{Type: clientv3.EventTypePut, Kv: &mvccpb.KeyValue{Key: []byte("foox"), Value: []byte("bar"), CreateRevision: 2, ModRevision: 2, Version: 1}}}
			if !reflect.DeepEqual(ev, resp.Events) {
				t.Fatalf("expected %+v, got %+v", ev, resp.Events)
			}
		} else if len(resp.Events) != 0 {
			t.Fatalf("expected no events, but got %+v", resp.Events)
		}
	case <-time.After(time.Duration(1.5 * float64(pi))):
		t.Fatalf("watch response expected in %v, but timed out", pi)
	}
}
func TestWatchEventType(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer cluster.Terminate(t)
	client := cluster.RandClient()
	ctx := context.Background()
	watchChan := client.Watch(ctx, "/", clientv3.WithPrefix())
	if _, err := client.Put(ctx, "/toDelete", "foo"); err != nil {
		t.Fatalf("Put failed: %v", err)
	}
	if _, err := client.Put(ctx, "/toDelete", "bar"); err != nil {
		t.Fatalf("Put failed: %v", err)
	}
	if _, err := client.Delete(ctx, "/toDelete"); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	lcr, err := client.Lease.Grant(ctx, 1)
	if err != nil {
		t.Fatalf("lease create failed: %v", err)
	}
	if _, err := client.Put(ctx, "/toExpire", "foo", clientv3.WithLease(lcr.ID)); err != nil {
		t.Fatalf("Put failed: %v", err)
	}
	tests := []struct {
		et		mvccpb.Event_EventType
		isCreate	bool
		isModify	bool
	}{{et: clientv3.EventTypePut, isCreate: true}, {et: clientv3.EventTypePut, isModify: true}, {et: clientv3.EventTypeDelete}, {et: clientv3.EventTypePut, isCreate: true}, {et: clientv3.EventTypeDelete}}
	var res []*clientv3.Event
	for {
		select {
		case wres := <-watchChan:
			res = append(res, wres.Events...)
		case <-time.After(10 * time.Second):
			t.Fatalf("Should receive %d events and then break out loop", len(tests))
		}
		if len(res) == len(tests) {
			break
		}
	}
	for i, tt := range tests {
		ev := res[i]
		if tt.et != ev.Type {
			t.Errorf("#%d: event type want=%s, get=%s", i, tt.et, ev.Type)
		}
		if tt.isCreate && !ev.IsCreate() {
			t.Errorf("#%d: event should be CreateEvent", i)
		}
		if tt.isModify && !ev.IsModify() {
			t.Errorf("#%d: event should be ModifyEvent", i)
		}
	}
}
func TestWatchErrConnClosed(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		ch := cli.Watch(context.TODO(), "foo")
		if wr := <-ch; grpc.ErrorDesc(wr.Err()) != grpc.ErrClientConnClosing.Error() {
			t.Fatalf("expected %v, got %v", grpc.ErrClientConnClosing, grpc.ErrorDesc(wr.Err()))
		}
	}()
	if err := cli.ActiveConnection().Close(); err != nil {
		t.Fatal(err)
	}
	clus.TakeClient(0)
	select {
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("wc.Watch took too long")
	case <-donec:
	}
}
func TestWatchAfterClose(t *testing.T) {
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
		cli.Watch(context.TODO(), "foo")
		if err := cli.Close(); err != nil && err != grpc.ErrClientConnClosing {
			t.Fatalf("expected %v, got %v", grpc.ErrClientConnClosing, err)
		}
		close(donec)
	}()
	select {
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("wc.Watch took too long")
	case <-donec:
	}
}
func TestWatchWithRequireLeader(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	liveClient := clus.Client(0)
	if _, err := liveClient.Put(context.TODO(), "foo", "bar"); err != nil {
		t.Fatal(err)
	}
	clus.Members[1].Stop(t)
	clus.Members[2].Stop(t)
	clus.Client(1).Close()
	clus.Client(2).Close()
	clus.TakeClient(1)
	clus.TakeClient(2)
	tickDuration := 10 * time.Millisecond
	time.Sleep(time.Duration(5*clus.Members[0].ElectionTicks) * tickDuration)
	chLeader := liveClient.Watch(clientv3.WithRequireLeader(context.TODO()), "foo", clientv3.WithRev(1))
	chNoLeader := liveClient.Watch(context.TODO(), "foo", clientv3.WithRev(1))
	select {
	case resp, ok := <-chLeader:
		if !ok {
			t.Fatalf("expected %v watch channel, got closed channel", rpctypes.ErrNoLeader)
		}
		if resp.Err() != rpctypes.ErrNoLeader {
			t.Fatalf("expected %v watch response error, got %+v", rpctypes.ErrNoLeader, resp)
		}
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("watch without leader took too long to close")
	}
	select {
	case resp, ok := <-chLeader:
		if ok {
			t.Fatalf("expected closed channel, got response %v", resp)
		}
	case <-time.After(integration.RequestWaitTimeout):
		t.Fatal("waited too long for channel to close")
	}
	if _, ok := <-chNoLeader; !ok {
		t.Fatalf("expected response, got closed channel")
	}
}
func TestWatchWithFilter(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer cluster.Terminate(t)
	client := cluster.RandClient()
	ctx := context.Background()
	wcNoPut := client.Watch(ctx, "a", clientv3.WithFilterPut())
	wcNoDel := client.Watch(ctx, "a", clientv3.WithFilterDelete())
	if _, err := client.Put(ctx, "a", "abc"); err != nil {
		t.Fatal(err)
	}
	if _, err := client.Delete(ctx, "a"); err != nil {
		t.Fatal(err)
	}
	npResp := <-wcNoPut
	if len(npResp.Events) != 1 || npResp.Events[0].Type != clientv3.EventTypeDelete {
		t.Fatalf("expected delete event, got %+v", npResp.Events)
	}
	ndResp := <-wcNoDel
	if len(ndResp.Events) != 1 || ndResp.Events[0].Type != clientv3.EventTypePut {
		t.Fatalf("expected put event, got %+v", ndResp.Events)
	}
	select {
	case resp := <-wcNoPut:
		t.Fatalf("unexpected event on filtered put (%+v)", resp)
	case resp := <-wcNoDel:
		t.Fatalf("unexpected event on filtered delete (%+v)", resp)
	case <-time.After(100 * time.Millisecond):
	}
}
func TestWatchWithCreatedNotification(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer cluster.Terminate(t)
	client := cluster.RandClient()
	ctx := context.Background()
	createC := client.Watch(ctx, "a", clientv3.WithCreatedNotify())
	resp := <-createC
	if !resp.Created {
		t.Fatalf("expected created event, got %v", resp)
	}
}
func TestWatchWithCreatedNotificationDropConn(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer cluster.Terminate(t)
	client := cluster.RandClient()
	wch := client.Watch(context.Background(), "a", clientv3.WithCreatedNotify())
	resp := <-wch
	if !resp.Created {
		t.Fatalf("expected created event, got %v", resp)
	}
	cluster.Members[0].DropConnections()
	select {
	case wresp := <-wch:
		t.Fatalf("got unexpected watch response: %+v\n", wresp)
	case <-time.After(time.Second):
	}
}
func TestWatchCancelOnServer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer cluster.Terminate(t)
	client := cluster.RandClient()
	numWatches := 10
	for {
		ctx, cancel := context.WithCancel(clientv3.WithRequireLeader(context.TODO()))
		ww := client.Watch(ctx, "a", clientv3.WithCreatedNotify())
		wresp := <-ww
		cancel()
		if wresp.Err() == nil {
			break
		}
	}
	cancels := make([]context.CancelFunc, numWatches)
	for i := 0; i < numWatches; i++ {
		md := metadata.Pairs("some-key", fmt.Sprintf("%d", i))
		mctx := metadata.NewOutgoingContext(context.Background(), md)
		ctx, cancel := context.WithCancel(mctx)
		cancels[i] = cancel
		w := client.Watch(ctx, fmt.Sprintf("%d", i), clientv3.WithCreatedNotify())
		<-w
	}
	maxWatches, _ := cluster.Members[0].Metric("etcd_debugging_mvcc_watcher_total")
	for i := 0; i < numWatches; i++ {
		cancels[i]()
	}
	time.Sleep(time.Second)
	minWatches, err := cluster.Members[0].Metric("etcd_debugging_mvcc_watcher_total")
	if err != nil {
		t.Fatal(err)
	}
	maxWatchV, minWatchV := 0, 0
	n, serr := fmt.Sscanf(maxWatches+" "+minWatches, "%d %d", &maxWatchV, &minWatchV)
	if n != 2 || serr != nil {
		t.Fatalf("expected n=2 and err=nil, got n=%d and err=%v", n, serr)
	}
	if maxWatchV-minWatchV < numWatches {
		t.Fatalf("expected %d canceled watchers, got %d", numWatches, maxWatchV-minWatchV)
	}
}
func TestWatchOverlapContextCancel(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := func(clus *integration.ClusterV3) {
	}
	testWatchOverlapContextCancel(t, f)
}
func TestWatchOverlapDropConnContextCancel(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := func(clus *integration.ClusterV3) {
		clus.Members[0].DropConnections()
	}
	testWatchOverlapContextCancel(t, f)
}
func testWatchOverlapContextCancel(t *testing.T, f func(*integration.ClusterV3)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	n := 100
	ctxs, ctxc := make([]context.Context, 5), make([]chan struct{}, 5)
	for i := range ctxs {
		md := metadata.Pairs("some-key", fmt.Sprintf("%d", i))
		ctxs[i] = metadata.NewOutgoingContext(context.Background(), md)
		ctxc[i] = make(chan struct{}, 2)
	}
	cli := clus.RandClient()
	if _, err := cli.Put(context.TODO(), "abc", "def"); err != nil {
		t.Fatal(err)
	}
	ch := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				ch <- struct{}{}
			}()
			idx := rand.Intn(len(ctxs))
			ctx, cancel := context.WithCancel(ctxs[idx])
			ctxc[idx] <- struct{}{}
			wch := cli.Watch(ctx, "abc", clientv3.WithRev(1))
			f(clus)
			select {
			case _, ok := <-wch:
				if !ok {
					t.Fatalf("unexpected closed channel %p", wch)
				}
			case <-time.After(5 * time.Second):
				t.Errorf("timed out waiting for watch on %p", wch)
			}
			if rand.Intn(2) == 0 {
				<-ctxc[idx]
				cancel()
			} else {
				cancel()
				<-ctxc[idx]
			}
		}()
	}
	for i := 0; i < n; i++ {
		select {
		case <-ch:
		case <-time.After(5 * time.Second):
			t.Fatalf("timed out waiting for completed watch")
		}
	}
}
func TestWatchCancelAndCloseClient(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	ctx, cancel := context.WithCancel(context.Background())
	wch := cli.Watch(ctx, "abc")
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		select {
		case wr, ok := <-wch:
			if ok {
				t.Fatalf("expected closed watch after cancel(), got resp=%+v err=%v", wr, wr.Err())
			}
		case <-time.After(5 * time.Second):
			t.Fatal("timed out waiting for closed channel")
		}
	}()
	cancel()
	if err := cli.Close(); err != nil {
		t.Fatal(err)
	}
	<-donec
	clus.TakeClient(0)
}
func TestWatchStressResumeClose(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	ctx, cancel := context.WithCancel(context.Background())
	wchs := make([]clientv3.WatchChan, 2000)
	for i := range wchs {
		wchs[i] = cli.Watch(ctx, "abc")
	}
	clus.Members[0].DropConnections()
	cancel()
	if err := cli.Close(); err != nil {
		t.Fatal(err)
	}
	clus.TakeClient(0)
}
func TestWatchCancelDisconnected(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := integration.NewClusterV3(t, &integration.ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	ctx, cancel := context.WithCancel(context.Background())
	wch := cli.Watch(ctx, "abc")
	clus.Members[0].Stop(t)
	cancel()
	select {
	case <-wch:
	case <-time.After(time.Second):
		t.Fatal("took too long to cancel disconnected watcher")
	}
}
