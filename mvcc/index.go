package mvcc

import (
	"sort"
	"sync"
	"github.com/google/btree"
	"go.uber.org/zap"
)

type index interface {
	Get(key []byte, atRev int64) (rev, created revision, ver int64, err error)
	Range(key, end []byte, atRev int64) ([][]byte, []revision)
	Revisions(key, end []byte, atRev int64) []revision
	Put(key []byte, rev revision)
	Tombstone(key []byte, rev revision) error
	RangeSince(key, end []byte, rev int64) []revision
	Compact(rev int64) map[revision]struct{}
	Keep(rev int64) map[revision]struct{}
	Equal(b index) bool
	Insert(ki *keyIndex)
	KeyIndex(ki *keyIndex) *keyIndex
}
type treeIndex struct {
	sync.RWMutex
	tree	*btree.BTree
	lg	*zap.Logger
}

func newTreeIndex(lg *zap.Logger) index {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &treeIndex{tree: btree.New(32), lg: lg}
}
func (ti *treeIndex) Put(key []byte, rev revision) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keyi := &keyIndex{key: key}
	ti.Lock()
	defer ti.Unlock()
	item := ti.tree.Get(keyi)
	if item == nil {
		keyi.put(ti.lg, rev.main, rev.sub)
		ti.tree.ReplaceOrInsert(keyi)
		return
	}
	okeyi := item.(*keyIndex)
	okeyi.put(ti.lg, rev.main, rev.sub)
}
func (ti *treeIndex) Get(key []byte, atRev int64) (modified, created revision, ver int64, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keyi := &keyIndex{key: key}
	ti.RLock()
	defer ti.RUnlock()
	if keyi = ti.keyIndex(keyi); keyi == nil {
		return revision{}, revision{}, 0, ErrRevisionNotFound
	}
	return keyi.get(ti.lg, atRev)
}
func (ti *treeIndex) KeyIndex(keyi *keyIndex) *keyIndex {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ti.RLock()
	defer ti.RUnlock()
	return ti.keyIndex(keyi)
}
func (ti *treeIndex) keyIndex(keyi *keyIndex) *keyIndex {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if item := ti.tree.Get(keyi); item != nil {
		return item.(*keyIndex)
	}
	return nil
}
func (ti *treeIndex) visit(key, end []byte, f func(ki *keyIndex)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keyi, endi := &keyIndex{key: key}, &keyIndex{key: end}
	ti.RLock()
	clone := ti.tree.Clone()
	ti.RUnlock()
	clone.AscendGreaterOrEqual(keyi, func(item btree.Item) bool {
		if len(endi.key) > 0 && !item.Less(endi) {
			return false
		}
		f(item.(*keyIndex))
		return true
	})
}
func (ti *treeIndex) Revisions(key, end []byte, atRev int64) (revs []revision) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if end == nil {
		rev, _, _, err := ti.Get(key, atRev)
		if err != nil {
			return nil
		}
		return []revision{rev}
	}
	ti.visit(key, end, func(ki *keyIndex) {
		if rev, _, _, err := ki.get(ti.lg, atRev); err == nil {
			revs = append(revs, rev)
		}
	})
	return revs
}
func (ti *treeIndex) Range(key, end []byte, atRev int64) (keys [][]byte, revs []revision) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if end == nil {
		rev, _, _, err := ti.Get(key, atRev)
		if err != nil {
			return nil, nil
		}
		return [][]byte{key}, []revision{rev}
	}
	ti.visit(key, end, func(ki *keyIndex) {
		if rev, _, _, err := ki.get(ti.lg, atRev); err == nil {
			revs = append(revs, rev)
			keys = append(keys, ki.key)
		}
	})
	return keys, revs
}
func (ti *treeIndex) Tombstone(key []byte, rev revision) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keyi := &keyIndex{key: key}
	ti.Lock()
	defer ti.Unlock()
	item := ti.tree.Get(keyi)
	if item == nil {
		return ErrRevisionNotFound
	}
	ki := item.(*keyIndex)
	return ki.tombstone(ti.lg, rev.main, rev.sub)
}
func (ti *treeIndex) RangeSince(key, end []byte, rev int64) []revision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keyi := &keyIndex{key: key}
	ti.RLock()
	defer ti.RUnlock()
	if end == nil {
		item := ti.tree.Get(keyi)
		if item == nil {
			return nil
		}
		keyi = item.(*keyIndex)
		return keyi.since(ti.lg, rev)
	}
	endi := &keyIndex{key: end}
	var revs []revision
	ti.tree.AscendGreaterOrEqual(keyi, func(item btree.Item) bool {
		if len(endi.key) > 0 && !item.Less(endi) {
			return false
		}
		curKeyi := item.(*keyIndex)
		revs = append(revs, curKeyi.since(ti.lg, rev)...)
		return true
	})
	sort.Sort(revisions(revs))
	return revs
}
func (ti *treeIndex) Compact(rev int64) map[revision]struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	available := make(map[revision]struct{})
	if ti.lg != nil {
		ti.lg.Info("compact tree index", zap.Int64("revision", rev))
	} else {
		plog.Printf("store.index: compact %d", rev)
	}
	ti.Lock()
	clone := ti.tree.Clone()
	ti.Unlock()
	clone.Ascend(func(item btree.Item) bool {
		keyi := item.(*keyIndex)
		ti.Lock()
		keyi.compact(ti.lg, rev, available)
		if keyi.isEmpty() {
			item := ti.tree.Delete(keyi)
			if item == nil {
				if ti.lg != nil {
					ti.lg.Panic("failed to delete during compaction")
				} else {
					plog.Panic("store.index: unexpected delete failure during compaction")
				}
			}
		}
		ti.Unlock()
		return true
	})
	return available
}
func (ti *treeIndex) Keep(rev int64) map[revision]struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	available := make(map[revision]struct{})
	ti.RLock()
	defer ti.RUnlock()
	ti.tree.Ascend(func(i btree.Item) bool {
		keyi := i.(*keyIndex)
		keyi.keep(rev, available)
		return true
	})
	return available
}
func (ti *treeIndex) Equal(bi index) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := bi.(*treeIndex)
	if ti.tree.Len() != b.tree.Len() {
		return false
	}
	equal := true
	ti.tree.Ascend(func(item btree.Item) bool {
		aki := item.(*keyIndex)
		bki := b.tree.Get(item).(*keyIndex)
		if !aki.equal(bki) {
			equal = false
			return false
		}
		return true
	})
	return equal
}
func (ti *treeIndex) Insert(ki *keyIndex) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ti.Lock()
	defer ti.Unlock()
	ti.tree.ReplaceOrInsert(ki)
}
