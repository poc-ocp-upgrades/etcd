package mvcc

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func TestWatcherWatchID(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := WatchableKV(newWatchableStore(b, &lease.FakeLessor{}, nil))
	defer cleanup(s, b, tmpPath)
	w := s.NewWatchStream()
	defer w.Close()
	idm := make(map[WatchID]struct{})
	for i := 0; i < 10; i++ {
		id := w.Watch([]byte("foo"), nil, 0)
		if _, ok := idm[id]; ok {
			t.Errorf("#%d: id %d exists", i, id)
		}
		idm[id] = struct{}{}
		s.Put([]byte("foo"), []byte("bar"), lease.NoLease)
		resp := <-w.Chan()
		if resp.WatchID != id {
			t.Errorf("#%d: watch id in event = %d, want %d", i, resp.WatchID, id)
		}
		if err := w.Cancel(id); err != nil {
			t.Error(err)
		}
	}
	s.Put([]byte("foo2"), []byte("bar"), lease.NoLease)
	for i := 10; i < 20; i++ {
		id := w.Watch([]byte("foo2"), nil, 1)
		if _, ok := idm[id]; ok {
			t.Errorf("#%d: id %d exists", i, id)
		}
		idm[id] = struct{}{}
		resp := <-w.Chan()
		if resp.WatchID != id {
			t.Errorf("#%d: watch id in event = %d, want %d", i, resp.WatchID, id)
		}
		if err := w.Cancel(id); err != nil {
			t.Error(err)
		}
	}
}
func TestWatcherWatchPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := WatchableKV(newWatchableStore(b, &lease.FakeLessor{}, nil))
	defer cleanup(s, b, tmpPath)
	w := s.NewWatchStream()
	defer w.Close()
	idm := make(map[WatchID]struct{})
	val := []byte("bar")
	keyWatch, keyEnd, keyPut := []byte("foo"), []byte("fop"), []byte("foobar")
	for i := 0; i < 10; i++ {
		id := w.Watch(keyWatch, keyEnd, 0)
		if _, ok := idm[id]; ok {
			t.Errorf("#%d: unexpected duplicated id %x", i, id)
		}
		idm[id] = struct{}{}
		s.Put(keyPut, val, lease.NoLease)
		resp := <-w.Chan()
		if resp.WatchID != id {
			t.Errorf("#%d: watch id in event = %d, want %d", i, resp.WatchID, id)
		}
		if err := w.Cancel(id); err != nil {
			t.Errorf("#%d: unexpected cancel error %v", i, err)
		}
		if len(resp.Events) != 1 {
			t.Errorf("#%d: len(resp.Events) got = %d, want = 1", i, len(resp.Events))
		}
		if len(resp.Events) == 1 {
			if !bytes.Equal(resp.Events[0].Kv.Key, keyPut) {
				t.Errorf("#%d: resp.Events got = %s, want = %s", i, resp.Events[0].Kv.Key, keyPut)
			}
		}
	}
	keyWatch1, keyEnd1, keyPut1 := []byte("foo1"), []byte("foo2"), []byte("foo1bar")
	s.Put(keyPut1, val, lease.NoLease)
	for i := 10; i < 15; i++ {
		id := w.Watch(keyWatch1, keyEnd1, 1)
		if _, ok := idm[id]; ok {
			t.Errorf("#%d: id %d exists", i, id)
		}
		idm[id] = struct{}{}
		resp := <-w.Chan()
		if resp.WatchID != id {
			t.Errorf("#%d: watch id in event = %d, want %d", i, resp.WatchID, id)
		}
		if err := w.Cancel(id); err != nil {
			t.Error(err)
		}
		if len(resp.Events) != 1 {
			t.Errorf("#%d: len(resp.Events) got = %d, want = 1", i, len(resp.Events))
		}
		if len(resp.Events) == 1 {
			if !bytes.Equal(resp.Events[0].Kv.Key, keyPut1) {
				t.Errorf("#%d: resp.Events got = %s, want = %s", i, resp.Events[0].Kv.Key, keyPut1)
			}
		}
	}
}
func TestWatcherWatchWrongRange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := WatchableKV(newWatchableStore(b, &lease.FakeLessor{}, nil))
	defer cleanup(s, b, tmpPath)
	w := s.NewWatchStream()
	defer w.Close()
	if id := w.Watch([]byte("foa"), []byte("foa"), 1); id != -1 {
		t.Fatalf("key == end range given; id expected -1, got %d", id)
	}
	if id := w.Watch([]byte("fob"), []byte("foa"), 1); id != -1 {
		t.Fatalf("key > end range given; id expected -1, got %d", id)
	}
	if id := w.Watch([]byte("foo"), []byte{}, 1); id != 0 {
		t.Fatalf("\x00 is range given; id expected 0, got %d", id)
	}
}
func TestWatchDeleteRange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := newWatchableStore(b, &lease.FakeLessor{}, nil)
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKeyPrefix := []byte("foo")
	for i := 0; i < 3; i++ {
		s.Put([]byte(fmt.Sprintf("%s_%d", testKeyPrefix, i)), []byte("bar"), lease.NoLease)
	}
	w := s.NewWatchStream()
	from, to := []byte(testKeyPrefix), []byte(fmt.Sprintf("%s_%d", testKeyPrefix, 99))
	w.Watch(from, to, 0)
	s.DeleteRange(from, to)
	we := []mvccpb.Event{{Type: mvccpb.DELETE, Kv: &mvccpb.KeyValue{Key: []byte("foo_0"), ModRevision: 5}}, {Type: mvccpb.DELETE, Kv: &mvccpb.KeyValue{Key: []byte("foo_1"), ModRevision: 5}}, {Type: mvccpb.DELETE, Kv: &mvccpb.KeyValue{Key: []byte("foo_2"), ModRevision: 5}}}
	select {
	case r := <-w.Chan():
		if !reflect.DeepEqual(r.Events, we) {
			t.Errorf("event = %v, want %v", r.Events, we)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("failed to receive event after 10 seconds!")
	}
}
func TestWatchStreamCancelWatcherByID(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := WatchableKV(newWatchableStore(b, &lease.FakeLessor{}, nil))
	defer cleanup(s, b, tmpPath)
	w := s.NewWatchStream()
	defer w.Close()
	id := w.Watch([]byte("foo"), nil, 0)
	tests := []struct {
		cancelID	WatchID
		werr		error
	}{{id, nil}, {id, ErrWatcherNotExist}, {id + 1, ErrWatcherNotExist}}
	for i, tt := range tests {
		gerr := w.Cancel(tt.cancelID)
		if gerr != tt.werr {
			t.Errorf("#%d: err = %v, want %v", i, gerr, tt.werr)
		}
	}
	if l := len(w.(*watchStream).cancels); l != 0 {
		t.Errorf("cancels = %d, want 0", l)
	}
}
func TestWatcherRequestProgress(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := &watchableStore{store: NewStore(b, &lease.FakeLessor{}, nil), unsynced: newWatcherGroup(), synced: newWatcherGroup()}
	defer func() {
		s.store.Close()
		os.Remove(tmpPath)
	}()
	testKey := []byte("foo")
	notTestKey := []byte("bad")
	testValue := []byte("bar")
	s.Put(testKey, testValue, lease.NoLease)
	w := s.NewWatchStream()
	badID := WatchID(1000)
	w.RequestProgress(badID)
	select {
	case resp := <-w.Chan():
		t.Fatalf("unexpected %+v", resp)
	default:
	}
	id := w.Watch(notTestKey, nil, 1)
	w.RequestProgress(id)
	select {
	case resp := <-w.Chan():
		t.Fatalf("unexpected %+v", resp)
	default:
	}
	s.syncWatchers()
	w.RequestProgress(id)
	wrs := WatchResponse{WatchID: 0, Revision: 2}
	select {
	case resp := <-w.Chan():
		if !reflect.DeepEqual(resp, wrs) {
			t.Fatalf("got %+v, expect %+v", resp, wrs)
		}
	case <-time.After(time.Second):
		t.Fatal("failed to receive progress")
	}
}
func TestWatcherWatchWithFilter(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s := WatchableKV(newWatchableStore(b, &lease.FakeLessor{}, nil))
	defer cleanup(s, b, tmpPath)
	w := s.NewWatchStream()
	defer w.Close()
	filterPut := func(e mvccpb.Event) bool {
		return e.Type == mvccpb.PUT
	}
	w.Watch([]byte("foo"), nil, 0, filterPut)
	done := make(chan struct{})
	go func() {
		<-w.Chan()
		done <- struct{}{}
	}()
	s.Put([]byte("foo"), []byte("bar"), 0)
	select {
	case <-done:
		t.Fatal("failed to filter put request")
	case <-time.After(100 * time.Millisecond):
	}
	s.DeleteRange([]byte("foo"), nil)
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("failed to receive delete request")
	}
}
