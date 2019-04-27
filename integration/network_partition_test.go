package integration

import (
	"fmt"
	"testing"
	"time"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestNetworkPartition5MembersLeaderInMinority(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 5})
	defer clus.Terminate(t)
	leadIndex := clus.WaitLeader(t)
	minority := []int{leadIndex, (leadIndex + 1) % 5}
	majority := []int{(leadIndex + 2) % 5, (leadIndex + 3) % 5, (leadIndex + 4) % 5}
	minorityMembers := getMembersByIndexSlice(clus.cluster, minority)
	majorityMembers := getMembersByIndexSlice(clus.cluster, majority)
	injectPartition(t, minorityMembers, majorityMembers)
	clus.waitNoLeader(t, minorityMembers)
	time.Sleep(2 * majorityMembers[0].ElectionTimeout())
	clus.waitLeader(t, majorityMembers)
	recoverPartition(t, minorityMembers, majorityMembers)
	clusterMustProgress(t, append(majorityMembers, minorityMembers...))
}
func TestNetworkPartition5MembersLeaderInMajority(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	for i := 0; i < 3; i++ {
		if err = testNetworkPartition5MembersLeaderInMajority(t); err == nil {
			break
		}
		t.Logf("[%d] got %v", i, err)
	}
	if err != nil {
		t.Fatalf("failed after 3 tries (%v)", err)
	}
}
func testNetworkPartition5MembersLeaderInMajority(t *testing.T) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 5})
	defer clus.Terminate(t)
	leadIndex := clus.WaitLeader(t)
	majority := []int{leadIndex, (leadIndex + 1) % 5, (leadIndex + 2) % 5}
	minority := []int{(leadIndex + 3) % 5, (leadIndex + 4) % 5}
	majorityMembers := getMembersByIndexSlice(clus.cluster, majority)
	minorityMembers := getMembersByIndexSlice(clus.cluster, minority)
	injectPartition(t, majorityMembers, minorityMembers)
	clus.waitNoLeader(t, minorityMembers)
	time.Sleep(2 * majorityMembers[0].ElectionTimeout())
	leadIndex2 := clus.waitLeader(t, majorityMembers)
	leadID, leadID2 := clus.Members[leadIndex].s.ID(), majorityMembers[leadIndex2].s.ID()
	if leadID != leadID2 {
		return fmt.Errorf("unexpected leader change from %s, got %s", leadID, leadID2)
	}
	recoverPartition(t, majorityMembers, minorityMembers)
	clusterMustProgress(t, append(majorityMembers, minorityMembers...))
	return nil
}
func TestNetworkPartition4Members(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 4})
	defer clus.Terminate(t)
	leadIndex := clus.WaitLeader(t)
	groupA := []int{leadIndex, (leadIndex + 1) % 4}
	groupB := []int{(leadIndex + 2) % 4, (leadIndex + 3) % 4}
	leaderPartition := getMembersByIndexSlice(clus.cluster, groupA)
	followerPartition := getMembersByIndexSlice(clus.cluster, groupB)
	injectPartition(t, leaderPartition, followerPartition)
	clus.WaitNoLeader(t)
	recoverPartition(t, leaderPartition, followerPartition)
	clus.WaitLeader(t)
	clusterMustProgress(t, clus.Members)
}
func getMembersByIndexSlice(clus *cluster, idxs []int) []*member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ms := make([]*member, len(idxs))
	for i, idx := range idxs {
		ms[i] = clus.Members[idx]
	}
	return ms
}
func injectPartition(t *testing.T, src, others []*member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range src {
		m.InjectPartition(t, others...)
	}
}
func recoverPartition(t *testing.T, src, others []*member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range src {
		m.RecoverPartition(t, others...)
	}
}
