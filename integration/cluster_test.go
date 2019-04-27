package integration

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/pkg/capnslog"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)
	if t := os.Getenv("ETCD_ELECTION_TIMEOUT_TICKS"); t != "" {
		if i, err := strconv.ParseInt(t, 10, 64); err == nil {
			electionTicks = int(i)
		}
	}
}
func TestClusterOf1(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCluster(t, 1)
}
func TestClusterOf3(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCluster(t, 3)
}
func testCluster(t *testing.T, size int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, size)
	c.Launch(t)
	defer c.Terminate(t)
	clusterMustProgress(t, c.Members)
}
func TestTLSClusterOf3(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewClusterByConfig(t, &ClusterConfig{Size: 3, PeerTLS: &testTLSInfo})
	c.Launch(t)
	defer c.Terminate(t)
	clusterMustProgress(t, c.Members)
}
func TestClusterOf1UsingDiscovery(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testClusterUsingDiscovery(t, 1)
}
func TestClusterOf3UsingDiscovery(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testClusterUsingDiscovery(t, 3)
}
func testClusterUsingDiscovery(t *testing.T, size int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	dc := NewCluster(t, 1)
	dc.Launch(t)
	defer dc.Terminate(t)
	dcc := MustNewHTTPClient(t, dc.URLs(), nil)
	dkapi := client.NewKeysAPI(dcc)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	if _, err := dkapi.Create(ctx, "/_config/size", fmt.Sprintf("%d", size)); err != nil {
		t.Fatal(err)
	}
	cancel()
	c := NewClusterByConfig(t, &ClusterConfig{Size: size, DiscoveryURL: dc.URL(0) + "/v2/keys"})
	c.Launch(t)
	defer c.Terminate(t)
	clusterMustProgress(t, c.Members)
}
func TestTLSClusterOf3UsingDiscovery(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	dc := NewCluster(t, 1)
	dc.Launch(t)
	defer dc.Terminate(t)
	dcc := MustNewHTTPClient(t, dc.URLs(), nil)
	dkapi := client.NewKeysAPI(dcc)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	if _, err := dkapi.Create(ctx, "/_config/size", fmt.Sprintf("%d", 3)); err != nil {
		t.Fatal(err)
	}
	cancel()
	c := NewClusterByConfig(t, &ClusterConfig{Size: 3, PeerTLS: &testTLSInfo, DiscoveryURL: dc.URL(0) + "/v2/keys"})
	c.Launch(t)
	defer c.Terminate(t)
	clusterMustProgress(t, c.Members)
}
func TestDoubleClusterSizeOf1(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testDoubleClusterSize(t, 1)
}
func TestDoubleClusterSizeOf3(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testDoubleClusterSize(t, 3)
}
func testDoubleClusterSize(t *testing.T, size int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, size)
	c.Launch(t)
	defer c.Terminate(t)
	for i := 0; i < size; i++ {
		c.AddMember(t)
	}
	clusterMustProgress(t, c.Members)
}
func TestDoubleTLSClusterSizeOf3(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewClusterByConfig(t, &ClusterConfig{Size: 3, PeerTLS: &testTLSInfo})
	c.Launch(t)
	defer c.Terminate(t)
	for i := 0; i < 3; i++ {
		c.AddMember(t)
	}
	clusterMustProgress(t, c.Members)
}
func TestDecreaseClusterSizeOf3(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testDecreaseClusterSize(t, 3)
}
func TestDecreaseClusterSizeOf5(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testDecreaseClusterSize(t, 5)
}
func testDecreaseClusterSize(t *testing.T, size int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, size)
	c.Launch(t)
	defer c.Terminate(t)
	for i := 0; i < size-1; i++ {
		id := c.Members[len(c.Members)-1].s.ID()
		if err := c.removeMember(t, uint64(id)); err != nil {
			if strings.Contains(err.Error(), "no leader") {
				t.Logf("got leader error (%v)", err)
				i--
				continue
			}
			t.Fatal(err)
		}
		c.waitLeader(t, c.Members)
	}
	clusterMustProgress(t, c.Members)
}
func TestForceNewCluster(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := NewCluster(t, 3)
	c.Launch(t)
	cc := MustNewHTTPClient(t, []string{c.Members[0].URL()}, nil)
	kapi := client.NewKeysAPI(cc)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := kapi.Create(ctx, "/foo", "bar")
	if err != nil {
		t.Fatalf("unexpected create error: %v", err)
	}
	cancel()
	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	if _, err = kapi.Watcher("/foo", &client.WatcherOptions{AfterIndex: resp.Node.ModifiedIndex - 1}).Next(ctx); err != nil {
		t.Fatalf("unexpected watch error: %v", err)
	}
	cancel()
	c.Members[0].Stop(t)
	c.Members[1].Terminate(t)
	c.Members[2].Terminate(t)
	c.Members[0].ForceNewCluster = true
	err = c.Members[0].Restart(t)
	if err != nil {
		t.Fatalf("unexpected ForceRestart error: %v", err)
	}
	defer c.Members[0].Terminate(t)
	c.waitLeader(t, c.Members[:1])
	cc = MustNewHTTPClient(t, []string{c.Members[0].URL()}, nil)
	kapi = client.NewKeysAPI(cc)
	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	if _, err := kapi.Watcher("/foo", &client.WatcherOptions{AfterIndex: resp.Node.ModifiedIndex - 1}).Next(ctx); err != nil {
		t.Fatalf("unexpected watch error: %v", err)
	}
	cancel()
	clusterMustProgress(t, c.Members[:1])
}
func TestAddMemberAfterClusterFullRotation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, 3)
	c.Launch(t)
	defer c.Terminate(t)
	for i := 0; i < 3; i++ {
		c.RemoveMember(t, uint64(c.Members[0].s.ID()))
		c.waitLeader(t, c.Members)
		c.AddMember(t)
		c.waitLeader(t, c.Members)
	}
	c.AddMember(t)
	c.waitLeader(t, c.Members)
	clusterMustProgress(t, c.Members)
}
func TestIssue2681(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, 5)
	c.Launch(t)
	defer c.Terminate(t)
	c.RemoveMember(t, uint64(c.Members[4].s.ID()))
	c.waitLeader(t, c.Members)
	c.AddMember(t)
	c.waitLeader(t, c.Members)
	clusterMustProgress(t, c.Members)
}
func TestIssue2746(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testIssue2746(t, 5)
}
func TestIssue2746WithThree(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testIssue2746(t, 3)
}
func testIssue2746(t *testing.T, members int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, members)
	for _, m := range c.Members {
		m.SnapCount = 10
	}
	c.Launch(t)
	defer c.Terminate(t)
	for i := 0; i < 20; i++ {
		clusterMustProgress(t, c.Members)
	}
	c.RemoveMember(t, uint64(c.Members[members-1].s.ID()))
	c.waitLeader(t, c.Members)
	c.AddMember(t)
	c.waitLeader(t, c.Members)
	clusterMustProgress(t, c.Members)
}
func TestIssue2904(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, 1)
	c.Launch(t)
	defer c.Terminate(t)
	c.AddMember(t)
	c.Members[1].Stop(t)
	cc := MustNewHTTPClient(t, c.URLs(), nil)
	ma := client.NewMembersAPI(cc)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	ma.Remove(ctx, c.Members[1].s.ID().String())
	cancel()
	c.Members[1].Restart(t)
	<-c.Members[1].s.StopNotify()
	c.Members[1].Terminate(t)
	c.Members = c.Members[:1]
	c.waitMembersMatch(t, c.HTTPMembers())
}
func TestIssue3699(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, 3)
	c.Launch(t)
	defer c.Terminate(t)
	c.Members[0].Stop(t)
	c.AddMember(t)
	leaderID := c.waitLeader(t, c.Members)
	for leaderID != 3 {
		c.Members[leaderID].Stop(t)
		<-c.Members[leaderID].s.StopNotify()
		time.Sleep(time.Duration(electionTicks * int(tickDuration)))
		c.Members[leaderID].Restart(t)
		leaderID = c.waitLeader(t, c.Members)
	}
	if err := c.Members[0].Restart(t); err != nil {
		t.Fatal(err)
	}
	select {
	case <-time.After(10 * time.Second):
		t.Fatalf("waited too long for ready notification")
	case <-c.Members[0].s.StopNotify():
		t.Fatalf("should not be stopped")
	case <-c.Members[0].s.ReadyNotify():
	}
	c.waitLeader(t, c.Members)
	cc := MustNewHTTPClient(t, []string{c.URL(0)}, c.cfg.ClientTLS)
	kapi := client.NewKeysAPI(cc)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	if _, err := kapi.Set(ctx, "/foo", "bar", nil); err != nil {
		t.Fatalf("unexpected error on Set (%v)", err)
	}
	cancel()
}
func TestRejectUnhealthyAdd(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, 3)
	for _, m := range c.Members {
		m.ServerConfig.StrictReconfigCheck = true
	}
	c.Launch(t)
	defer c.Terminate(t)
	c.Members[0].Stop(t)
	c.WaitLeader(t)
	for i := 1; i < len(c.Members); i++ {
		err := c.addMemberByURL(t, c.URL(i), "unix://foo:12345")
		if err == nil {
			t.Fatalf("should have failed adding peer")
		}
		if !strings.Contains(err.Error(), "has no leader") {
			t.Errorf("unexpected error (%v)", err)
		}
	}
	c.Members[0].Restart(t)
	c.WaitLeader(t)
	time.Sleep(2 * etcdserver.HealthInterval)
	var err error
	for i := 1; i < len(c.Members); i++ {
		if err = c.addMemberByURL(t, c.URL(i), "unix://foo:12345"); err == nil {
			break
		}
	}
	if err != nil {
		t.Fatalf("should have added peer to healthy cluster (%v)", err)
	}
}
func TestRejectUnhealthyRemove(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	c := NewCluster(t, 5)
	for _, m := range c.Members {
		m.ServerConfig.StrictReconfigCheck = true
	}
	c.Launch(t)
	defer c.Terminate(t)
	c.Members[0].Stop(t)
	c.Members[1].Stop(t)
	c.WaitLeader(t)
	err := c.removeMember(t, uint64(c.Members[2].s.ID()))
	if err == nil {
		t.Fatalf("should reject quorum breaking remove")
	}
	if !strings.Contains(err.Error(), "has no leader") {
		t.Errorf("unexpected error (%v)", err)
	}
	time.Sleep(time.Duration(electionTicks * int(tickDuration)))
	if err = c.removeMember(t, uint64(c.Members[0].s.ID())); err != nil {
		t.Fatalf("should accept removing down member")
	}
	c.Members[0].Restart(t)
	time.Sleep((3 * etcdserver.HealthInterval) / 2)
	if err = c.removeMember(t, uint64(c.Members[0].s.ID())); err != nil {
		t.Fatalf("expected to remove member, got error %v", err)
	}
}
func TestRestartRemoved(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	capnslog.SetGlobalLogLevel(capnslog.INFO)
	defer capnslog.SetGlobalLogLevel(defaultLogLevel)
	c := NewCluster(t, 1)
	for _, m := range c.Members {
		m.ServerConfig.StrictReconfigCheck = true
	}
	c.Launch(t)
	defer c.Terminate(t)
	c.AddMember(t)
	c.WaitLeader(t)
	oldm := c.Members[0]
	oldm.keepDataDirTerminate = true
	if err := c.removeMember(t, uint64(c.Members[0].s.ID())); err != nil {
		t.Fatalf("expected to remove member, got error %v", err)
	}
	c.WaitLeader(t)
	oldm.ServerConfig.NewCluster = false
	if err := oldm.Restart(t); err != nil {
		t.Fatalf("unexpected ForceRestart error: %v", err)
	}
	defer func() {
		oldm.Close()
		os.RemoveAll(oldm.ServerConfig.DataDir)
	}()
	select {
	case <-oldm.s.StopNotify():
	case <-time.After(time.Minute):
		t.Fatalf("removed member didn't exit within %v", time.Minute)
	}
}
func clusterMustProgress(t *testing.T, membs []*member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := MustNewHTTPClient(t, []string{membs[0].URL()}, nil)
	kapi := client.NewKeysAPI(cc)
	key := fmt.Sprintf("foo%d", rand.Int())
	var (
		err	error
		resp	*client.Response
	)
	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		resp, err = kapi.Create(ctx, "/"+key, "bar")
		cancel()
		if err == nil {
			break
		}
		t.Logf("failed to create key on %q (%v)", membs[0].URL(), err)
	}
	if err != nil {
		t.Fatalf("create on %s error: %v", membs[0].URL(), err)
	}
	for i, m := range membs {
		u := m.URL()
		mcc := MustNewHTTPClient(t, []string{u}, nil)
		mkapi := client.NewKeysAPI(mcc)
		mctx, mcancel := context.WithTimeout(context.Background(), requestTimeout)
		if _, err := mkapi.Watcher(key, &client.WatcherOptions{AfterIndex: resp.Node.ModifiedIndex - 1}).Next(mctx); err != nil {
			t.Fatalf("#%d: watch on %s error: %v", i, u, err)
		}
		mcancel()
	}
}
func TestSpeedyTerminate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	clus := NewClusterV3(t, &ClusterConfig{Size: 3})
	for i := 0; i < 3; i++ {
		clus.Members[i].Stop(t)
		clus.Members[i].Restart(t)
	}
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		clus.Terminate(t)
	}()
	select {
	case <-time.After(10 * time.Second):
		t.Fatalf("cluster took too long to terminate")
	case <-donec:
	}
}
