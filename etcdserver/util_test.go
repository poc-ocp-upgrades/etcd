package etcdserver

import (
	"net/http"
	"testing"
	"time"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/rafthttp"
	"github.com/coreos/etcd/snap"
)

func TestLongestConnected(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	umap, err := types.NewURLsMap("mem1=http://10.1:2379,mem2=http://10.2:2379,mem3=http://10.3:2379")
	if err != nil {
		t.Fatal(err)
	}
	clus, err := membership.NewClusterFromURLsMap("test", umap)
	if err != nil {
		t.Fatal(err)
	}
	memberIDs := clus.MemberIDs()
	tr := newNopTransporterWithActiveTime(memberIDs)
	transferee, ok := longestConnected(tr, memberIDs)
	if !ok {
		t.Fatalf("unexpected ok %v", ok)
	}
	if memberIDs[0] != transferee {
		t.Fatalf("expected first member %s to be transferee, got %s", memberIDs[0], transferee)
	}
	amap := make(map[types.ID]time.Time)
	for _, id := range memberIDs {
		amap[id] = time.Time{}
	}
	tr.(*nopTransporterWithActiveTime).reset(amap)
	_, ok2 := longestConnected(tr, memberIDs)
	if ok2 {
		t.Fatalf("unexpected ok %v", ok)
	}
}

type nopTransporterWithActiveTime struct{ activeMap map[types.ID]time.Time }

func newNopTransporterWithActiveTime(memberIDs []types.ID) rafthttp.Transporter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	am := make(map[types.ID]time.Time)
	for i, id := range memberIDs {
		am[id] = time.Now().Add(time.Duration(i) * time.Second)
	}
	return &nopTransporterWithActiveTime{activeMap: am}
}
func (s *nopTransporterWithActiveTime) Start() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *nopTransporterWithActiveTime) Handler() http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *nopTransporterWithActiveTime) Send(m []raftpb.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) SendSnapshot(m snap.Message) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) AddRemote(id types.ID, us []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) AddPeer(id types.ID, us []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) RemovePeer(id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) RemoveAllPeers() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) UpdatePeer(id types.ID, us []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) ActiveSince(id types.ID) time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.activeMap[id]
}
func (s *nopTransporterWithActiveTime) ActivePeers() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (s *nopTransporterWithActiveTime) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) Resume() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (s *nopTransporterWithActiveTime) reset(am map[types.ID]time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.activeMap = am
}
