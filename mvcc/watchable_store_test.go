package mvcc

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func TestWatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(b, &lease.FakeLessor{}, nil)
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKey := []byte("foo")
	testValue := []byte("bar")
	s.Put(testKey, testValue, lease.NoLease)
	w := s.NewWatchStream()
	w.Watch(testKey, nil, 0)
	if !s.synced.contains(string(testKey)) {
		t.Errorf("existence = false, want true")
	}
}
func TestNewWatcherCancel(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(b, &lease.FakeLessor{}, nil)
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKey := []byte("foo")
	testValue := []byte("bar")
	s.Put(testKey, testValue, lease.NoLease)
	w := s.NewWatchStream()
	wt := w.Watch(testKey, nil, 0)
	if err := w.Cancel(wt); err != nil {
		t.Error(err)
	}
	if s.synced.contains(string(testKey)) {
		t.Errorf("existence = true, want false")
	}
}
func TestCancelUnsynced(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := &watchableStore{store: NewStore(b, &lease.FakeLessor{}, nil), unsynced: newWatcherGroup(), synced: newWatcherGroup()}
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKey := []byte("foo")
	testValue := []byte("bar")
	s.Put(testKey, testValue, lease.NoLease)
	w := s.NewWatchStream()
	watcherN := 100
	watchIDs := make([]WatchID, watcherN)
	for i := 0; i < watcherN; i++ {
		watchIDs[i] = w.Watch(testKey, nil, 1)
	}
	for _, idx := range watchIDs {
		if err := w.Cancel(idx); err != nil {
			t.Error(err)
		}
	}
	if size := s.unsynced.size(); size != 0 {
		t.Errorf("unsynced size = %d, want 0", size)
	}
}
func TestSyncWatchers(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := &watchableStore{store: NewStore(b, &lease.FakeLessor{}, nil), unsynced: newWatcherGroup(), synced: newWatcherGroup()}
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKey := []byte("foo")
	testValue := []byte("bar")
	s.Put(testKey, testValue, lease.NoLease)
	w := s.NewWatchStream()
	watcherN := 100
	for i := 0; i < watcherN; i++ {
		w.Watch(testKey, nil, 1)
	}
	sws := s.synced.watcherSetByKey(string(testKey))
	uws := s.unsynced.watcherSetByKey(string(testKey))
	if len(sws) != 0 {
		t.Fatalf("synced[string(testKey)] size = %d, want 0", len(sws))
	}
	if len(uws) != watcherN {
		t.Errorf("unsynced size = %d, want %d", len(uws), watcherN)
	}
	s.syncWatchers()
	sws = s.synced.watcherSetByKey(string(testKey))
	uws = s.unsynced.watcherSetByKey(string(testKey))
	if len(sws) != watcherN {
		t.Errorf("synced[string(testKey)] size = %d, want %d", len(sws), watcherN)
	}
	if len(uws) != 0 {
		t.Errorf("unsynced size = %d, want 0", len(uws))
	}
	for w := range sws {
		if w.minRev != s.Rev()+1 {
			t.Errorf("w.minRev = %d, want %d", w.minRev, s.Rev()+1)
		}
	}
	if len(w.(*watchStream).ch) != watcherN {
		t.Errorf("watched event size = %d, want %d", len(w.(*watchStream).ch), watcherN)
	}
	evs := (<-w.(*watchStream).ch).Events
	if len(evs) != 1 {
		t.Errorf("len(evs) got = %d, want = 1", len(evs))
	}
	if evs[0].Type != mvccpb.PUT {
		t.Errorf("got = %v, want = %v", evs[0].Type, mvccpb.PUT)
	}
	if !bytes.Equal(evs[0].Kv.Key, testKey) {
		t.Errorf("got = %s, want = %s", evs[0].Kv.Key, testKey)
	}
	if !bytes.Equal(evs[0].Kv.Value, testValue) {
		t.Errorf("got = %s, want = %s", evs[0].Kv.Value, testValue)
	}
}
func TestWatchCompacted(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(b, &lease.FakeLessor{}, nil)
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKey := []byte("foo")
	testValue := []byte("bar")
	maxRev := 10
	compactRev := int64(5)
	for i := 0; i < maxRev; i++ {
		s.Put(testKey, testValue, lease.NoLease)
	}
	_, err := s.Compact(compactRev)
	if err != nil {
		t.Fatalf("failed to compact kv (%v)", err)
	}
	w := s.NewWatchStream()
	wt := w.Watch(testKey, nil, compactRev-1)
	select {
	case resp := <-w.Chan():
		if resp.WatchID != wt {
			t.Errorf("resp.WatchID = %x, want %x", resp.WatchID, wt)
		}
		if resp.CompactRevision == 0 {
			t.Errorf("resp.Compacted = %v, want %v", resp.CompactRevision, compactRev)
		}
	case <-time.After(1 * time.Second):
		t.Fatalf("failed to receive response (timeout)")
	}
}
func TestWatchFutureRev(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(b, &lease.FakeLessor{}, nil)
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKey := []byte("foo")
	testValue := []byte("bar")
	w := s.NewWatchStream()
	wrev := int64(10)
	w.Watch(testKey, nil, wrev)
	for i := 0; i < 10; i++ {
		rev := s.Put(testKey, testValue, lease.NoLease)
		if rev >= wrev {
			break
		}
	}
	select {
	case resp := <-w.Chan():
		if resp.Revision != wrev {
			t.Fatalf("rev = %d, want %d", resp.Revision, wrev)
		}
		if len(resp.Events) != 1 {
			t.Fatalf("failed to get events from the response")
		}
		if resp.Events[0].Kv.ModRevision != wrev {
			t.Fatalf("kv.rev = %d, want %d", resp.Events[0].Kv.ModRevision, wrev)
		}
	case <-time.After(time.Second):
		t.Fatal("failed to receive event in 1 second.")
	}
}
func TestWatchRestore(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	test := func(delay time.Duration) func(t *testing.T) {
		return func(t *testing.T) {
			b, tmpPath := backend.NewDefaultTmpBackend()
			s := newWatchableStore(b, &lease.FakeLessor{}, nil)
			defer cleanup(s, b, tmpPath)
			testKey := []byte("foo")
			testValue := []byte("bar")
			rev := s.Put(testKey, testValue, lease.NoLease)
			newBackend, newPath := backend.NewDefaultTmpBackend()
			newStore := newWatchableStore(newBackend, &lease.FakeLessor{}, nil)
			defer cleanup(newStore, newBackend, newPath)
			w := newStore.NewWatchStream()
			w.Watch(testKey, nil, rev-1)
			time.Sleep(delay)
			newStore.Restore(b)
			select {
			case resp := <-w.Chan():
				if resp.Revision != rev {
					t.Fatalf("rev = %d, want %d", resp.Revision, rev)
				}
				if len(resp.Events) != 1 {
					t.Fatalf("failed to get events from the response")
				}
				if resp.Events[0].Kv.ModRevision != rev {
					t.Fatalf("kv.rev = %d, want %d", resp.Events[0].Kv.ModRevision, rev)
				}
			case <-time.After(time.Second):
				t.Fatal("failed to receive event in 1 second.")
			}
		}
	}
	t.Run("Normal", test(0))
	t.Run("RunSyncWatchLoopBeforeRestore", test(time.Millisecond*120))
}
func TestWatchRestoreSyncedWatcher(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b1, b1Path := backend.NewDefaultTmpBackend()
	s1 := newWatchableStore(b1, &lease.FakeLessor{}, nil)
	defer cleanup(s1, b1, b1Path)
	b2, b2Path := backend.NewDefaultTmpBackend()
	s2 := newWatchableStore(b2, &lease.FakeLessor{}, nil)
	defer cleanup(s2, b2, b2Path)
	testKey, testValue := []byte("foo"), []byte("bar")
	rev := s1.Put(testKey, testValue, lease.NoLease)
	startRev := rev + 2
	w1 := s1.NewWatchStream()
	w1.Watch(testKey, nil, startRev)
	s2.Put(testKey, testValue, lease.NoLease)
	s2.Put(testKey, testValue, lease.NoLease)
	if err := s1.Restore(b2); err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	s1.Put(testKey, testValue, lease.NoLease)
	select {
	case resp := <-w1.Chan():
		if resp.Revision != startRev {
			t.Fatalf("resp.Revision expect %d, got %d", startRev, resp.Revision)
		}
		if len(resp.Events) != 1 {
			t.Fatalf("len(resp.Events) expect 1, got %d", len(resp.Events))
		}
		if resp.Events[0].Kv.ModRevision != startRev {
			t.Fatalf("resp.Events[0].Kv.ModRevision expect %d, got %d", startRev, resp.Events[0].Kv.ModRevision)
		}
	case <-time.After(time.Second):
		t.Fatal("failed to receive event in 1 second")
	}
}
func TestWatchBatchUnsynced(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(b, &lease.FakeLessor{}, nil)
	oldMaxRevs := watchBatchMaxRevs
	defer func() {
		watchBatchMaxRevs = oldMaxRevs
		s.store.Close()
		os.Remove(tmpPath)
	}()
	batches := 3
	watchBatchMaxRevs = 4
	v := []byte("foo")
	for i := 0; i < watchBatchMaxRevs*batches; i++ {
		s.Put(v, v, lease.NoLease)
	}
	w := s.NewWatchStream()
	w.Watch(v, nil, 1)
	for i := 0; i < batches; i++ {
		if resp := <-w.Chan(); len(resp.Events) != watchBatchMaxRevs {
			t.Fatalf("len(events) = %d, want %d", len(resp.Events), watchBatchMaxRevs)
		}
	}
	s.store.revMu.Lock()
	defer s.store.revMu.Unlock()
	if size := s.synced.size(); size != 1 {
		t.Errorf("synced size = %d, want 1", size)
	}
}
func TestNewMapwatcherToEventMap(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	k0, k1, k2 := []byte("foo0"), []byte("foo1"), []byte("foo2")
	v0, v1, v2 := []byte("bar0"), []byte("bar1"), []byte("bar2")
	ws := []*watcher{{key: k0}, {key: k1}, {key: k2}}
	evs := []mvccpb.Event{{Type: mvccpb.PUT, Kv: &mvccpb.KeyValue{Key: k0, Value: v0}}, {Type: mvccpb.PUT, Kv: &mvccpb.KeyValue{Key: k1, Value: v1}}, {Type: mvccpb.PUT, Kv: &mvccpb.KeyValue{Key: k2, Value: v2}}}
	tests := []struct {
		sync	[]*watcher
		evs	[]mvccpb.Event
		wwe	map[*watcher][]mvccpb.Event
	}{{nil, evs, map[*watcher][]mvccpb.Event{}}, {[]*watcher{ws[2]}, evs[:1], map[*watcher][]mvccpb.Event{}}, {[]*watcher{ws[1]}, evs[1:2], map[*watcher][]mvccpb.Event{ws[1]: evs[1:2]}}, {[]*watcher{ws[0], ws[2]}, evs[2:], map[*watcher][]mvccpb.Event{ws[2]: evs[2:]}}, {[]*watcher{ws[0], ws[1]}, evs[:2], map[*watcher][]mvccpb.Event{ws[0]: evs[:1], ws[1]: evs[1:2]}}}
	for i, tt := range tests {
		wg := newWatcherGroup()
		for _, w := range tt.sync {
			wg.add(w)
		}
		gwe := newWatcherBatch(&wg, tt.evs)
		if len(gwe) != len(tt.wwe) {
			t.Errorf("#%d: len(gwe) got = %d, want = %d", i, len(gwe), len(tt.wwe))
		}
		for w, eb := range gwe {
			if len(eb.evs) != len(tt.wwe[w]) {
				t.Errorf("#%d: len(eb.evs) got = %d, want = %d", i, len(eb.evs), len(tt.wwe[w]))
			}
			if !reflect.DeepEqual(eb.evs, tt.wwe[w]) {
				t.Errorf("#%d: reflect.DeepEqual events got = %v, want = true", i, false)
			}
		}
	}
}
func TestWatchVictims(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldChanBufLen, oldMaxWatchersPerSync := chanBufLen, maxWatchersPerSync
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(b, &lease.FakeLessor{}, nil)
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
		chanBufLen, maxWatchersPerSync = oldChanBufLen, oldMaxWatchersPerSync
	}()
	chanBufLen, maxWatchersPerSync = 1, 2
	numPuts := chanBufLen * 64
	testKey, testValue := []byte("foo"), []byte("bar")
	var wg sync.WaitGroup
	numWatches := maxWatchersPerSync * 128
	errc := make(chan error, numWatches)
	wg.Add(numWatches)
	for i := 0; i < numWatches; i++ {
		go func() {
			w := s.NewWatchStream()
			w.Watch(testKey, nil, 1)
			defer func() {
				w.Close()
				wg.Done()
			}()
			tc := time.After(10 * time.Second)
			evs, nextRev := 0, int64(2)
			for evs < numPuts {
				select {
				case <-tc:
					errc <- fmt.Errorf("time out")
					return
				case wr := <-w.Chan():
					evs += len(wr.Events)
					for _, ev := range wr.Events {
						if ev.Kv.ModRevision != nextRev {
							errc <- fmt.Errorf("expected rev=%d, got %d", nextRev, ev.Kv.ModRevision)
							return
						}
						nextRev++
					}
					time.Sleep(time.Millisecond)
				}
			}
			if evs != numPuts {
				errc <- fmt.Errorf("expected %d events, got %d", numPuts, evs)
				return
			}
			select {
			case <-w.Chan():
				errc <- fmt.Errorf("unexpected response")
			default:
			}
		}()
		time.Sleep(time.Millisecond)
	}
	var wgPut sync.WaitGroup
	wgPut.Add(numPuts)
	for i := 0; i < numPuts; i++ {
		go func() {
			defer wgPut.Done()
			s.Put(testKey, testValue, lease.NoLease)
		}()
	}
	wgPut.Wait()
	wg.Wait()
	select {
	case err := <-errc:
		t.Fatal(err)
	default:
	}
}
func TestStressWatchCancelClose(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(b, &lease.FakeLessor{}, nil)
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKey, testValue := []byte("foo"), []byte("bar")
	var wg sync.WaitGroup
	readyc := make(chan struct{})
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			w := s.NewWatchStream()
			ids := make([]WatchID, 10)
			for i := range ids {
				ids[i] = w.Watch(testKey, nil, 0)
			}
			<-readyc
			wg.Add(1 + len(ids)/2)
			for i := range ids[:len(ids)/2] {
				go func(n int) {
					defer wg.Done()
					w.Cancel(ids[n])
				}(i)
			}
			go func() {
				defer wg.Done()
				w.Close()
			}()
		}()
	}
	close(readyc)
	for i := 0; i < 100; i++ {
		s.Put(testKey, testValue, lease.NoLease)
	}
	wg.Wait()
}
