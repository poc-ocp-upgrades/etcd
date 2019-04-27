package store

import (
	"container/heap"
)

type ttlKeyHeap struct {
	array	[]*node
	keyMap	map[*node]int
}

func newTtlKeyHeap() *ttlKeyHeap {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h := &ttlKeyHeap{keyMap: make(map[*node]int)}
	heap.Init(h)
	return h
}
func (h ttlKeyHeap) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(h.array)
}
func (h ttlKeyHeap) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return h.array[i].ExpireTime.Before(h.array[j].ExpireTime)
}
func (h ttlKeyHeap) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h.array[i], h.array[j] = h.array[j], h.array[i]
	h.keyMap[h.array[i]] = i
	h.keyMap[h.array[j]] = j
}
func (h *ttlKeyHeap) Push(x interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n, _ := x.(*node)
	h.keyMap[n] = len(h.array)
	h.array = append(h.array, n)
}
func (h *ttlKeyHeap) Pop() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	old := h.array
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	h.array = old[0 : n-1]
	delete(h.keyMap, x)
	return x
}
func (h *ttlKeyHeap) top() *node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.Len() != 0 {
		return h.array[0]
	}
	return nil
}
func (h *ttlKeyHeap) pop() *node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	x := heap.Pop(h)
	n, _ := x.(*node)
	return n
}
func (h *ttlKeyHeap) push(x interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	heap.Push(h, x)
}
func (h *ttlKeyHeap) update(n *node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	index, ok := h.keyMap[n]
	if ok {
		heap.Remove(h, index)
		heap.Push(h, n)
	}
}
func (h *ttlKeyHeap) remove(n *node) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	index, ok := h.keyMap[n]
	if ok {
		heap.Remove(h, index)
	}
}
