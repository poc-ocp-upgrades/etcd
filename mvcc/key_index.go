package mvcc

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/btree"
)

var (
	ErrRevisionNotFound = errors.New("mvcc: revision not found")
)

type keyIndex struct {
	key         []byte
	modified    revision
	generations []generation
}

func (ki *keyIndex) put(main int64, sub int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rev := revision{main: main, sub: sub}
	if !rev.GreaterThan(ki.modified) {
		plog.Panicf("store.keyindex: put with unexpected smaller revision [%v / %v]", rev, ki.modified)
	}
	if len(ki.generations) == 0 {
		ki.generations = append(ki.generations, generation{})
	}
	g := &ki.generations[len(ki.generations)-1]
	if len(g.revs) == 0 {
		keysGauge.Inc()
		g.created = rev
	}
	g.revs = append(g.revs, rev)
	g.ver++
	ki.modified = rev
}
func (ki *keyIndex) restore(created, modified revision, ver int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(ki.generations) != 0 {
		plog.Panicf("store.keyindex: cannot restore non-empty keyIndex")
	}
	ki.modified = modified
	g := generation{created: created, ver: ver, revs: []revision{modified}}
	ki.generations = append(ki.generations, g)
	keysGauge.Inc()
}
func (ki *keyIndex) tombstone(main int64, sub int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ki.isEmpty() {
		plog.Panicf("store.keyindex: unexpected tombstone on empty keyIndex %s", string(ki.key))
	}
	if ki.generations[len(ki.generations)-1].isEmpty() {
		return ErrRevisionNotFound
	}
	ki.put(main, sub)
	ki.generations = append(ki.generations, generation{})
	keysGauge.Dec()
	return nil
}
func (ki *keyIndex) get(atRev int64) (modified, created revision, ver int64, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ki.isEmpty() {
		plog.Panicf("store.keyindex: unexpected get on empty keyIndex %s", string(ki.key))
	}
	g := ki.findGeneration(atRev)
	if g.isEmpty() {
		return revision{}, revision{}, 0, ErrRevisionNotFound
	}
	n := g.walk(func(rev revision) bool {
		return rev.main > atRev
	})
	if n != -1 {
		return g.revs[n], g.created, g.ver - int64(len(g.revs)-n-1), nil
	}
	return revision{}, revision{}, 0, ErrRevisionNotFound
}
func (ki *keyIndex) since(rev int64) []revision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ki.isEmpty() {
		plog.Panicf("store.keyindex: unexpected get on empty keyIndex %s", string(ki.key))
	}
	since := revision{rev, 0}
	var gi int
	for gi = len(ki.generations) - 1; gi > 0; gi-- {
		g := ki.generations[gi]
		if g.isEmpty() {
			continue
		}
		if since.GreaterThan(g.created) {
			break
		}
	}
	var revs []revision
	var last int64
	for ; gi < len(ki.generations); gi++ {
		for _, r := range ki.generations[gi].revs {
			if since.GreaterThan(r) {
				continue
			}
			if r.main == last {
				revs[len(revs)-1] = r
				continue
			}
			revs = append(revs, r)
			last = r.main
		}
	}
	return revs
}
func (ki *keyIndex) compact(atRev int64, available map[revision]struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ki.isEmpty() {
		plog.Panicf("store.keyindex: unexpected compact on empty keyIndex %s", string(ki.key))
	}
	genIdx, revIndex := ki.doCompact(atRev, available)
	g := &ki.generations[genIdx]
	if !g.isEmpty() {
		if revIndex != -1 {
			g.revs = g.revs[revIndex:]
		}
		if len(g.revs) == 1 && genIdx != len(ki.generations)-1 {
			delete(available, g.revs[0])
			genIdx++
		}
	}
	ki.generations = ki.generations[genIdx:]
}
func (ki *keyIndex) keep(atRev int64, available map[revision]struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ki.isEmpty() {
		return
	}
	genIdx, revIndex := ki.doCompact(atRev, available)
	g := &ki.generations[genIdx]
	if !g.isEmpty() {
		if revIndex == len(g.revs)-1 && genIdx != len(ki.generations)-1 {
			delete(available, g.revs[revIndex])
		}
	}
}
func (ki *keyIndex) doCompact(atRev int64, available map[revision]struct{}) (genIdx int, revIndex int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := func(rev revision) bool {
		if rev.main <= atRev {
			available[rev] = struct{}{}
			return false
		}
		return true
	}
	genIdx, g := 0, &ki.generations[0]
	for genIdx < len(ki.generations)-1 {
		if tomb := g.revs[len(g.revs)-1].main; tomb > atRev {
			break
		}
		genIdx++
		g = &ki.generations[genIdx]
	}
	revIndex = g.walk(f)
	return genIdx, revIndex
}
func (ki *keyIndex) isEmpty() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(ki.generations) == 1 && ki.generations[0].isEmpty()
}
func (ki *keyIndex) findGeneration(rev int64) *generation {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lastg := len(ki.generations) - 1
	cg := lastg
	for cg >= 0 {
		if len(ki.generations[cg].revs) == 0 {
			cg--
			continue
		}
		g := ki.generations[cg]
		if cg != lastg {
			if tomb := g.revs[len(g.revs)-1].main; tomb <= rev {
				return nil
			}
		}
		if g.revs[0].main <= rev {
			return &ki.generations[cg]
		}
		cg--
	}
	return nil
}
func (a *keyIndex) Less(b btree.Item) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bytes.Compare(a.key, b.(*keyIndex).key) == -1
}
func (a *keyIndex) equal(b *keyIndex) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !bytes.Equal(a.key, b.key) {
		return false
	}
	if a.modified != b.modified {
		return false
	}
	if len(a.generations) != len(b.generations) {
		return false
	}
	for i := range a.generations {
		ag, bg := a.generations[i], b.generations[i]
		if !ag.equal(bg) {
			return false
		}
	}
	return true
}
func (ki *keyIndex) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var s string
	for _, g := range ki.generations {
		s += g.String()
	}
	return s
}

type generation struct {
	ver     int64
	created revision
	revs    []revision
}

func (g *generation) isEmpty() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return g == nil || len(g.revs) == 0
}
func (g *generation) walk(f func(rev revision) bool) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(g.revs)
	for i := range g.revs {
		ok := f(g.revs[l-i-1])
		if !ok {
			return l - i - 1
		}
	}
	return -1
}
func (g *generation) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("g: created[%d] ver[%d], revs %#v\n", g.created, g.ver, g.revs)
}
func (a generation) equal(b generation) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.ver != b.ver {
		return false
	}
	if len(a.revs) != len(b.revs) {
		return false
	}
	for i := range a.revs {
		ar, br := a.revs[i], b.revs[i]
		if ar != br {
			return false
		}
	}
	return true
}
