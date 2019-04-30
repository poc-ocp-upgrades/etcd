package types

type Uint64Slice []uint64

func (p Uint64Slice) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(p)
}
func (p Uint64Slice) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p[i] < p[j]
}
func (p Uint64Slice) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p[i], p[j] = p[j], p[i]
}
