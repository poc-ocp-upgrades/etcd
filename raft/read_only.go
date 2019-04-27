package raft

import pb "github.com/coreos/etcd/raft/raftpb"

type ReadState struct {
	Index		uint64
	RequestCtx	[]byte
}
type readIndexStatus struct {
	req	pb.Message
	index	uint64
	acks	map[uint64]struct{}
}
type readOnly struct {
	option			ReadOnlyOption
	pendingReadIndex	map[string]*readIndexStatus
	readIndexQueue		[]string
}

func newReadOnly(option ReadOnlyOption) *readOnly {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &readOnly{option: option, pendingReadIndex: make(map[string]*readIndexStatus)}
}
func (ro *readOnly) addRequest(index uint64, m pb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx := string(m.Entries[0].Data)
	if _, ok := ro.pendingReadIndex[ctx]; ok {
		return
	}
	ro.pendingReadIndex[ctx] = &readIndexStatus{index: index, req: m, acks: make(map[uint64]struct{})}
	ro.readIndexQueue = append(ro.readIndexQueue, ctx)
}
func (ro *readOnly) recvAck(m pb.Message) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rs, ok := ro.pendingReadIndex[string(m.Context)]
	if !ok {
		return 0
	}
	rs.acks[m.From] = struct{}{}
	return len(rs.acks) + 1
}
func (ro *readOnly) advance(m pb.Message) []*readIndexStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		i	int
		found	bool
	)
	ctx := string(m.Context)
	rss := []*readIndexStatus{}
	for _, okctx := range ro.readIndexQueue {
		i++
		rs, ok := ro.pendingReadIndex[okctx]
		if !ok {
			panic("cannot find corresponding read state from pending map")
		}
		rss = append(rss, rs)
		if okctx == ctx {
			found = true
			break
		}
	}
	if found {
		ro.readIndexQueue = ro.readIndexQueue[i:]
		for _, rs := range rss {
			delete(ro.pendingReadIndex, string(rs.req.Entries[0].Data))
		}
		return rss
	}
	return nil
}
func (ro *readOnly) lastPendingRequestCtx() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(ro.readIndexQueue) == 0 {
		return ""
	}
	return ro.readIndexQueue[len(ro.readIndexQueue)-1]
}
