package raft

import "fmt"

const (
	ProgressStateProbe	ProgressStateType	= iota
	ProgressStateReplicate
	ProgressStateSnapshot
)

type ProgressStateType uint64

var prstmap = [...]string{"ProgressStateProbe", "ProgressStateReplicate", "ProgressStateSnapshot"}

func (st ProgressStateType) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return prstmap[uint64(st)]
}

type Progress struct {
	Match, Next	uint64
	State		ProgressStateType
	Paused		bool
	PendingSnapshot	uint64
	RecentActive	bool
	ins		*inflights
	IsLearner	bool
}

func (pr *Progress) resetState(state ProgressStateType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.Paused = false
	pr.PendingSnapshot = 0
	pr.State = state
	pr.ins.reset()
}
func (pr *Progress) becomeProbe() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.State == ProgressStateSnapshot {
		pendingSnapshot := pr.PendingSnapshot
		pr.resetState(ProgressStateProbe)
		pr.Next = max(pr.Match+1, pendingSnapshot+1)
	} else {
		pr.resetState(ProgressStateProbe)
		pr.Next = pr.Match + 1
	}
}
func (pr *Progress) becomeReplicate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.resetState(ProgressStateReplicate)
	pr.Next = pr.Match + 1
}
func (pr *Progress) becomeSnapshot(snapshoti uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.resetState(ProgressStateSnapshot)
	pr.PendingSnapshot = snapshoti
}
func (pr *Progress) maybeUpdate(n uint64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var updated bool
	if pr.Match < n {
		pr.Match = n
		updated = true
		pr.resume()
	}
	if pr.Next < n+1 {
		pr.Next = n + 1
	}
	return updated
}
func (pr *Progress) optimisticUpdate(n uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.Next = n + 1
}
func (pr *Progress) maybeDecrTo(rejected, last uint64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.State == ProgressStateReplicate {
		if rejected <= pr.Match {
			return false
		}
		pr.Next = pr.Match + 1
		return true
	}
	if pr.Next-1 != rejected {
		return false
	}
	if pr.Next = min(rejected, last+1); pr.Next < 1 {
		pr.Next = 1
	}
	pr.resume()
	return true
}
func (pr *Progress) pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.Paused = true
}
func (pr *Progress) resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.Paused = false
}
func (pr *Progress) IsPaused() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch pr.State {
	case ProgressStateProbe:
		return pr.Paused
	case ProgressStateReplicate:
		return pr.ins.full()
	case ProgressStateSnapshot:
		return true
	default:
		panic("unexpected state")
	}
}
func (pr *Progress) snapshotFailure() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pr.PendingSnapshot = 0
}
func (pr *Progress) needSnapshotAbort() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pr.State == ProgressStateSnapshot && pr.Match >= pr.PendingSnapshot
}
func (pr *Progress) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("next = %d, match = %d, state = %s, waiting = %v, pendingSnapshot = %d", pr.Next, pr.Match, pr.State, pr.IsPaused(), pr.PendingSnapshot)
}

type inflights struct {
	start	int
	count	int
	size	int
	buffer	[]uint64
}

func newInflights(size int) *inflights {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &inflights{size: size}
}
func (in *inflights) add(inflight uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.full() {
		panic("cannot add into a full inflights")
	}
	next := in.start + in.count
	size := in.size
	if next >= size {
		next -= size
	}
	if next >= len(in.buffer) {
		in.growBuf()
	}
	in.buffer[next] = inflight
	in.count++
}
func (in *inflights) growBuf() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newSize := len(in.buffer) * 2
	if newSize == 0 {
		newSize = 1
	} else if newSize > in.size {
		newSize = in.size
	}
	newBuffer := make([]uint64, newSize)
	copy(newBuffer, in.buffer)
	in.buffer = newBuffer
}
func (in *inflights) freeTo(to uint64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in.count == 0 || to < in.buffer[in.start] {
		return
	}
	idx := in.start
	var i int
	for i = 0; i < in.count; i++ {
		if to < in.buffer[idx] {
			break
		}
		size := in.size
		if idx++; idx >= size {
			idx -= size
		}
	}
	in.count -= i
	in.start = idx
	if in.count == 0 {
		in.start = 0
	}
}
func (in *inflights) freeFirstOne() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in.freeTo(in.buffer[in.start])
}
func (in *inflights) full() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return in.count == in.size
}
func (in *inflights) reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	in.count = 0
	in.start = 0
}
