package types

import (
	"strconv"
)

type ID uint64

func (i ID) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strconv.FormatUint(uint64(i), 16)
}
func IDFromString(s string) (ID, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i, err := strconv.ParseUint(s, 16, 64)
	return ID(i), err
}

type IDSlice []ID

func (p IDSlice) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(p)
}
func (p IDSlice) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return uint64(p[i]) < uint64(p[j])
}
func (p IDSlice) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p[i], p[j] = p[j], p[i]
}
