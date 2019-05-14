package mvcc

import "encoding/binary"

const revBytesLen = 8 + 1 + 8

type revision struct {
	main int64
	sub  int64
}

func (a revision) GreaterThan(b revision) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.main > b.main {
		return true
	}
	if a.main < b.main {
		return false
	}
	return a.sub > b.sub
}
func newRevBytes() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return make([]byte, revBytesLen, markedRevBytesLen)
}
func revToBytes(rev revision, bytes []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	binary.BigEndian.PutUint64(bytes, uint64(rev.main))
	bytes[8] = '_'
	binary.BigEndian.PutUint64(bytes[9:], uint64(rev.sub))
}
func bytesToRev(bytes []byte) revision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return revision{main: int64(binary.BigEndian.Uint64(bytes[0:8])), sub: int64(binary.BigEndian.Uint64(bytes[9:]))}
}

type revisions []revision

func (a revisions) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(a)
}
func (a revisions) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a[j].GreaterThan(a[i])
}
func (a revisions) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	a[i], a[j] = a[j], a[i]
}
