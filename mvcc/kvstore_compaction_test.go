package mvcc

import (
	"os"
	"reflect"
	"testing"
	"time"
	"github.com/coreos/etcd/lease"
	"github.com/coreos/etcd/mvcc/backend"
)

func TestScheduleCompaction(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	revs := []revision{{1, 0}, {2, 0}, {3, 0}}
	tests := []struct {
		rev	int64
		keep	map[revision]struct{}
		wrevs	[]revision
	}{{1, nil, revs[1:]}, {3, nil, nil}, {1, map[revision]struct{}{{main: 1}: {}}, revs}, {3, map[revision]struct{}{{main: 2}: {}, {main: 3}: {}}, revs[1:]}}
	for i, tt := range tests {
		b, tmpPath := backend.NewDefaultTmpBackend()
		s := NewStore(b, &lease.FakeLessor{}, nil)
		tx := s.b.BatchTx()
		tx.Lock()
		ibytes := newRevBytes()
		for _, rev := range revs {
			revToBytes(rev, ibytes)
			tx.UnsafePut(keyBucketName, ibytes, []byte("bar"))
		}
		tx.Unlock()
		s.scheduleCompaction(tt.rev, tt.keep)
		tx.Lock()
		for _, rev := range tt.wrevs {
			revToBytes(rev, ibytes)
			keys, _ := tx.UnsafeRange(keyBucketName, ibytes, nil, 0)
			if len(keys) != 1 {
				t.Errorf("#%d: range on %v = %d, want 1", i, rev, len(keys))
			}
		}
		_, vals := tx.UnsafeRange(metaBucketName, finishedCompactKeyName, nil, 0)
		revToBytes(revision{main: tt.rev}, ibytes)
		if w := [][]byte{ibytes}; !reflect.DeepEqual(vals, w) {
			t.Errorf("#%d: vals on %v = %+v, want %+v", i, finishedCompactKeyName, vals, w)
		}
		tx.Unlock()
		cleanup(s, b, tmpPath)
	}
}
func TestCompactAllAndRestore(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, tmpPath := backend.NewDefaultTmpBackend()
	s0 := NewStore(b, &lease.FakeLessor{}, nil)
	defer os.Remove(tmpPath)
	s0.Put([]byte("foo"), []byte("bar"), lease.NoLease)
	s0.Put([]byte("foo"), []byte("bar1"), lease.NoLease)
	s0.Put([]byte("foo"), []byte("bar2"), lease.NoLease)
	s0.DeleteRange([]byte("foo"), nil)
	rev := s0.Rev()
	done, err := s0.Compact(rev)
	if err != nil {
		t.Fatal(err)
	}
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("timeout waiting for compaction to finish")
	}
	err = s0.Close()
	if err != nil {
		t.Fatal(err)
	}
	s1 := NewStore(b, &lease.FakeLessor{}, nil)
	if s1.Rev() != rev {
		t.Errorf("rev = %v, want %v", s1.Rev(), rev)
	}
	_, err = s1.Range([]byte("foo"), nil, RangeOptions{})
	if err != nil {
		t.Errorf("unexpect range error %v", err)
	}
}
