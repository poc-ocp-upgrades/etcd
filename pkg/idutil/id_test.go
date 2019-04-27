package idutil

import (
	"testing"
	"time"
)

func TestNewGenerator(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	g := NewGenerator(0x12, time.Unix(0, 0).Add(0x3456*time.Millisecond))
	id := g.Next()
	wid := uint64(0x12000000345601)
	if id != wid {
		t.Errorf("id = %x, want %x", id, wid)
	}
}
func TestNewGeneratorUnique(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	g := NewGenerator(0, time.Time{})
	id := g.Next()
	g1 := NewGenerator(1, time.Time{})
	if gid := g1.Next(); id == gid {
		t.Errorf("generate the same id %x using different server ID", id)
	}
	g2 := NewGenerator(0, time.Now())
	if gid := g2.Next(); id == gid {
		t.Errorf("generate the same id %x after restart", id)
	}
}
func TestNext(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	g := NewGenerator(0x12, time.Unix(0, 0).Add(0x3456*time.Millisecond))
	wid := uint64(0x12000000345601)
	for i := 0; i < 1000; i++ {
		id := g.Next()
		if id != wid+uint64(i) {
			t.Errorf("id = %x, want %x", id, wid+uint64(i))
		}
	}
}
