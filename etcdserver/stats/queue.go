package stats

import (
	"sync"
	"time"
)

const (
	queueCapacity = 200
)

type RequestStats struct {
	SendingTime time.Time
	Size        int
}
type statsQueue struct {
	items        [queueCapacity]*RequestStats
	size         int
	front        int
	back         int
	totalReqSize int
	rwl          sync.RWMutex
}

func (q *statsQueue) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return q.size
}
func (q *statsQueue) ReqSize() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return q.totalReqSize
}
func (q *statsQueue) frontAndBack() (*RequestStats, *RequestStats) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	q.rwl.RLock()
	defer q.rwl.RUnlock()
	if q.size != 0 {
		return q.items[q.front], q.items[q.back]
	}
	return nil, nil
}
func (q *statsQueue) Insert(p *RequestStats) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	q.rwl.Lock()
	defer q.rwl.Unlock()
	q.back = (q.back + 1) % queueCapacity
	if q.size == queueCapacity {
		q.totalReqSize -= q.items[q.front].Size
		q.front = (q.back + 1) % queueCapacity
	} else {
		q.size++
	}
	q.items[q.back] = p
	q.totalReqSize += q.items[q.back].Size
}
func (q *statsQueue) Rate() (float64, float64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	front, back := q.frontAndBack()
	if front == nil || back == nil {
		return 0, 0
	}
	if time.Since(back.SendingTime) > time.Second {
		q.Clear()
		return 0, 0
	}
	sampleDuration := back.SendingTime.Sub(front.SendingTime)
	pr := float64(q.Len()) / float64(sampleDuration) * float64(time.Second)
	br := float64(q.ReqSize()) / float64(sampleDuration) * float64(time.Second)
	return pr, br
}
func (q *statsQueue) Clear() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	q.rwl.Lock()
	defer q.rwl.Unlock()
	q.back = -1
	q.front = 0
	q.size = 0
	q.totalReqSize = 0
}
