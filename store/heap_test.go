package store

import (
	"fmt"
	"testing"
	"time"
)

func TestHeapPushPop(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	h := newTtlKeyHeap()
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("%v", 10-i)
		m := time.Duration(10 - i)
		n := newKV(nil, path, path, 0, nil, time.Now().Add(time.Second*m))
		h.push(n)
	}
	min := time.Now()
	for i := 0; i < 10; i++ {
		node := h.pop()
		if node.ExpireTime.Before(min) {
			t.Fatal("heap sort wrong!")
		}
		min = node.ExpireTime
	}
}
func TestHeapUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	h := newTtlKeyHeap()
	kvs := make([]*node, 10)
	for i := range kvs {
		path := fmt.Sprintf("%v", 10-i)
		m := time.Duration(10 - i)
		n := newKV(nil, path, path, 0, nil, time.Now().Add(time.Second*m))
		kvs[i] = n
		h.push(n)
	}
	kvs[3].ExpireTime = time.Now().Add(time.Second * 11)
	kvs[5].ExpireTime = time.Now().Add(time.Second * 12)
	h.update(kvs[3])
	h.update(kvs[5])
	min := time.Now()
	for i := 0; i < 10; i++ {
		node := h.pop()
		if node.ExpireTime.Before(min) {
			t.Fatal("heap sort wrong!")
		}
		min = node.ExpireTime
		if i == 8 {
			if node.Path != "7" {
				t.Fatal("heap sort wrong!", node.Path)
			}
		}
		if i == 9 {
			if node.Path != "5" {
				t.Fatal("heap sort wrong!")
			}
		}
	}
}
