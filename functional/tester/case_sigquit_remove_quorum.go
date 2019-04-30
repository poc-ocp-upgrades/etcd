package tester

import (
	"context"
	"fmt"
	"strings"
	"time"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/functional/rpcpb"
	"go.uber.org/zap"
)

type fetchSnapshotCaseQuorum struct {
	desc		string
	rpcpbCase	rpcpb.Case
	injected	map[int]struct{}
	snapshotted	int
}

func (c *fetchSnapshotCaseQuorum) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lead, err := clus.GetLeader()
	if err != nil {
		return err
	}
	c.snapshotted = lead
	clus.lg.Info("save snapshot on leader node START", zap.String("target-endpoint", clus.Members[lead].EtcdClientEndpoint))
	var resp *rpcpb.Response
	resp, err = clus.sendOpWithResp(lead, rpcpb.Operation_SAVE_SNAPSHOT)
	if resp == nil || (resp != nil && !resp.Success) || err != nil {
		clus.lg.Info("save snapshot on leader node FAIL", zap.String("target-endpoint", clus.Members[lead].EtcdClientEndpoint), zap.Error(err))
		return err
	}
	clus.lg.Info("save snapshot on leader node SUCCESS", zap.String("target-endpoint", clus.Members[lead].EtcdClientEndpoint), zap.String("member-name", resp.SnapshotInfo.MemberName), zap.Strings("member-client-urls", resp.SnapshotInfo.MemberClientURLs), zap.String("snapshot-path", resp.SnapshotInfo.SnapshotPath), zap.String("snapshot-file-size", resp.SnapshotInfo.SnapshotFileSize), zap.String("snapshot-total-size", resp.SnapshotInfo.SnapshotTotalSize), zap.Int64("snapshot-total-key", resp.SnapshotInfo.SnapshotTotalKey), zap.Int64("snapshot-hash", resp.SnapshotInfo.SnapshotHash), zap.Int64("snapshot-revision", resp.SnapshotInfo.SnapshotRevision), zap.String("took", resp.SnapshotInfo.Took), zap.Error(err))
	if err != nil {
		return err
	}
	clus.Members[lead].SnapshotInfo = resp.SnapshotInfo
	leaderc, err := clus.Members[lead].CreateEtcdClient()
	if err != nil {
		return err
	}
	defer leaderc.Close()
	var mresp *clientv3.MemberListResponse
	mresp, err = leaderc.MemberList(context.Background())
	mss := []string{}
	if err == nil && mresp != nil {
		mss = describeMembers(mresp)
	}
	clus.lg.Info("member list before disastrous machine failure", zap.String("request-to", clus.Members[lead].EtcdClientEndpoint), zap.Strings("members", mss), zap.Error(err))
	if err != nil {
		return err
	}
	time.Sleep(time.Second)
	for {
		c.injected = pickQuorum(len(clus.Members))
		if _, ok := c.injected[lead]; !ok {
			break
		}
	}
	for idx := range c.injected {
		clus.lg.Info("disastrous machine failure to quorum START", zap.String("target-endpoint", clus.Members[idx].EtcdClientEndpoint))
		err = clus.sendOp(idx, rpcpb.Operation_SIGQUIT_ETCD_AND_REMOVE_DATA)
		clus.lg.Info("disastrous machine failure to quorum END", zap.String("target-endpoint", clus.Members[idx].EtcdClientEndpoint), zap.Error(err))
		if err != nil {
			return err
		}
	}
	clus.lg.Info("disastrous machine failure to old leader START", zap.String("target-endpoint", clus.Members[lead].EtcdClientEndpoint))
	err = clus.sendOp(lead, rpcpb.Operation_SIGQUIT_ETCD_AND_REMOVE_DATA)
	clus.lg.Info("disastrous machine failure to old leader END", zap.String("target-endpoint", clus.Members[lead].EtcdClientEndpoint), zap.Error(err))
	return err
}
func (c *fetchSnapshotCaseQuorum) Recover(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	oldlead := c.snapshotted
	clus.Members[oldlead].EtcdOnSnapshotRestore = clus.Members[oldlead].Etcd
	clus.Members[oldlead].EtcdOnSnapshotRestore.InitialClusterState = "existing"
	name := clus.Members[oldlead].Etcd.Name
	initClus := []string{}
	for _, u := range clus.Members[oldlead].Etcd.AdvertisePeerURLs {
		initClus = append(initClus, fmt.Sprintf("%s=%s", name, u))
	}
	clus.Members[oldlead].EtcdOnSnapshotRestore.InitialCluster = strings.Join(initClus, ",")
	clus.lg.Info("restore snapshot and restart from snapshot request START", zap.String("target-endpoint", clus.Members[oldlead].EtcdClientEndpoint), zap.Strings("initial-cluster", initClus))
	err := clus.sendOp(oldlead, rpcpb.Operation_RESTORE_RESTART_FROM_SNAPSHOT)
	clus.lg.Info("restore snapshot and restart from snapshot request END", zap.String("target-endpoint", clus.Members[oldlead].EtcdClientEndpoint), zap.Strings("initial-cluster", initClus), zap.Error(err))
	if err != nil {
		return err
	}
	leaderc, err := clus.Members[oldlead].CreateEtcdClient()
	if err != nil {
		return err
	}
	defer leaderc.Close()
	idxs := make([]int, 0, len(c.injected))
	for idx := range c.injected {
		idxs = append(idxs, idx)
	}
	clus.lg.Info("member add START", zap.Int("members-to-add", len(idxs)))
	for i, idx := range idxs {
		clus.lg.Info("member add request SENT", zap.String("target-endpoint", clus.Members[idx].EtcdClientEndpoint), zap.Strings("peer-urls", clus.Members[idx].Etcd.AdvertisePeerURLs))
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		_, err := leaderc.MemberAdd(ctx, clus.Members[idx].Etcd.AdvertisePeerURLs)
		cancel()
		clus.lg.Info("member add request DONE", zap.String("target-endpoint", clus.Members[idx].EtcdClientEndpoint), zap.Strings("peer-urls", clus.Members[idx].Etcd.AdvertisePeerURLs), zap.Error(err))
		if err != nil {
			return err
		}
		clus.Members[idx].EtcdOnSnapshotRestore = clus.Members[idx].Etcd
		clus.Members[idx].EtcdOnSnapshotRestore.InitialClusterState = "existing"
		name := clus.Members[idx].Etcd.Name
		for _, u := range clus.Members[idx].Etcd.AdvertisePeerURLs {
			initClus = append(initClus, fmt.Sprintf("%s=%s", name, u))
		}
		clus.Members[idx].EtcdOnSnapshotRestore.InitialCluster = strings.Join(initClus, ",")
		clus.lg.Info("restart from snapshot request SENT", zap.String("target-endpoint", clus.Members[idx].EtcdClientEndpoint), zap.Strings("initial-cluster", initClus))
		err = clus.sendOp(idx, rpcpb.Operation_RESTART_FROM_SNAPSHOT)
		clus.lg.Info("restart from snapshot request DONE", zap.String("target-endpoint", clus.Members[idx].EtcdClientEndpoint), zap.Strings("initial-cluster", initClus), zap.Error(err))
		if err != nil {
			return err
		}
		if i != len(c.injected)-1 {
			dur := 5 * clus.Members[idx].ElectionTimeout()
			clus.lg.Info("waiting after restart from snapshot request", zap.Int("i", i), zap.Int("idx", idx), zap.Duration("sleep", dur))
			time.Sleep(dur)
		} else {
			clus.lg.Info("restart from snapshot request ALL END", zap.Int("i", i), zap.Int("idx", idx))
		}
	}
	return nil
}
func (c *fetchSnapshotCaseQuorum) Desc() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.desc != "" {
		return c.desc
	}
	return c.rpcpbCase.String()
}
func (c *fetchSnapshotCaseQuorum) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}
func new_Case_SIGQUIT_AND_REMOVE_QUORUM_AND_RESTORE_LEADER_SNAPSHOT_FROM_SCRATCH(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &fetchSnapshotCaseQuorum{rpcpbCase: rpcpb.Case_SIGQUIT_AND_REMOVE_QUORUM_AND_RESTORE_LEADER_SNAPSHOT_FROM_SCRATCH, injected: make(map[int]struct{}), snapshotted: -1}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
