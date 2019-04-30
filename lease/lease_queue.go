package lease

type LeaseWithTime struct {
	id	LeaseID
	time	int64
	index	int
}
type LeaseQueue []*LeaseWithTime

func (pq LeaseQueue) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(pq)
}
func (pq LeaseQueue) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pq[i].time < pq[j].time
}
func (pq LeaseQueue) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *LeaseQueue) Push(x interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := len(*pq)
	item := x.(*LeaseWithTime)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *LeaseQueue) Pop() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
