package etcdserver

import "testing"

func TestConsistentIndex(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i consistentIndex
	i.setConsistentIndex(10)
	if g := i.ConsistentIndex(); g != 10 {
		t.Errorf("value = %d, want 10", g)
	}
}
