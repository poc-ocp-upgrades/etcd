package snapshot

import "encoding/binary"

type revision struct {
	main	int64
	sub	int64
}

func bytesToRev(bytes []byte) revision {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return revision{main: int64(binary.BigEndian.Uint64(bytes[0:8])), sub: int64(binary.BigEndian.Uint64(bytes[9:]))}
}

type initIndex int

func (i *initIndex) ConsistentIndex() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return uint64(*i)
}
