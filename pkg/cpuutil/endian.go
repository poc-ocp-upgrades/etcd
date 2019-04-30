package cpuutil

import (
	"encoding/binary"
	"unsafe"
)

const intWidth int = int(unsafe.Sizeof(0))

var byteOrder binary.ByteOrder

func ByteOrder() binary.ByteOrder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return byteOrder
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	i := int(0x1)
	if v := (*[intWidth]byte)(unsafe.Pointer(&i)); v[0] == 0 {
		byteOrder = binary.BigEndian
	} else {
		byteOrder = binary.LittleEndian
	}
}
