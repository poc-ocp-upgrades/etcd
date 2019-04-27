package crc

import (
	"hash/crc32"
	"reflect"
	"testing"
)

func TestHash32(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stdhash := crc32.New(crc32.IEEETable)
	if _, err := stdhash.Write([]byte("test data")); err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}
	hash := New(stdhash.Sum32(), crc32.IEEETable)
	wsize := stdhash.Size()
	if g := hash.Size(); g != wsize {
		t.Errorf("size = %d, want %d", g, wsize)
	}
	wbsize := stdhash.BlockSize()
	if g := hash.BlockSize(); g != wbsize {
		t.Errorf("block size = %d, want %d", g, wbsize)
	}
	wsum32 := stdhash.Sum32()
	if g := hash.Sum32(); g != wsum32 {
		t.Errorf("Sum32 = %d, want %d", g, wsum32)
	}
	wsum := stdhash.Sum(make([]byte, 32))
	if g := hash.Sum(make([]byte, 32)); !reflect.DeepEqual(g, wsum) {
		t.Errorf("sum = %v, want %v", g, wsum)
	}
	if _, err := stdhash.Write([]byte("test data")); err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}
	if _, err := hash.Write([]byte("test data")); err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}
	wsum32 = stdhash.Sum32()
	if g := hash.Sum32(); g != wsum32 {
		t.Errorf("Sum32 after write = %d, want %d", g, wsum32)
	}
	stdhash.Reset()
	hash.Reset()
	wsum32 = stdhash.Sum32()
	if g := hash.Sum32(); g != wsum32 {
		t.Errorf("Sum32 after reset = %d, want %d", g, wsum32)
	}
}
