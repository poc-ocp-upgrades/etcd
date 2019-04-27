package tester

import (
	"time"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
)

const (
	waitRecover = 5 * time.Second
)

func inject_DELAY_PEER_PORT_TX_RX(clus *Cluster, idx int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus.lg.Info("injecting delay latency", zap.Duration("latency", time.Duration(clus.Tester.UpdatedDelayLatencyMs)*time.Millisecond), zap.Duration("latency-rv", time.Duration(clus.Tester.DelayLatencyMsRv)*time.Millisecond), zap.String("endpoint", clus.Members[idx].EtcdClientEndpoint))
	return clus.sendOp(idx, rpcpb.Operation_DELAY_PEER_PORT_TX_RX)
}
func recover_DELAY_PEER_PORT_TX_RX(clus *Cluster, idx int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := clus.sendOp(idx, rpcpb.Operation_UNDELAY_PEER_PORT_TX_RX)
	time.Sleep(waitRecover)
	return err
}
func new_Case_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER(clus *Cluster, random bool) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER, injectMember: inject_DELAY_PEER_PORT_TX_RX, recoverMember: recover_DELAY_PEER_PORT_TX_RX}
	clus.Tester.UpdatedDelayLatencyMs = clus.Tester.DelayLatencyMs
	if random {
		clus.UpdateDelayLatencyMs()
		cc.rpcpbCase = rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER
	}
	c := &caseFollower{cc, -1, -1}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
func new_Case_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT(clus *Cluster, random bool) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT, injectMember: inject_DELAY_PEER_PORT_TX_RX, recoverMember: recover_DELAY_PEER_PORT_TX_RX}
	clus.Tester.UpdatedDelayLatencyMs = clus.Tester.DelayLatencyMs
	if random {
		clus.UpdateDelayLatencyMs()
		cc.rpcpbCase = rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT
	}
	c := &caseFollower{cc, -1, -1}
	return &caseUntilSnapshot{rpcpbCase: cc.rpcpbCase, Case: c}
}
func new_Case_DELAY_PEER_PORT_TX_RX_LEADER(clus *Cluster, random bool) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_DELAY_PEER_PORT_TX_RX_LEADER, injectMember: inject_DELAY_PEER_PORT_TX_RX, recoverMember: recover_DELAY_PEER_PORT_TX_RX}
	clus.Tester.UpdatedDelayLatencyMs = clus.Tester.DelayLatencyMs
	if random {
		clus.UpdateDelayLatencyMs()
		cc.rpcpbCase = rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_LEADER
	}
	c := &caseLeader{cc, -1, -1}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
func new_Case_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT(clus *Cluster, random bool) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT, injectMember: inject_DELAY_PEER_PORT_TX_RX, recoverMember: recover_DELAY_PEER_PORT_TX_RX}
	clus.Tester.UpdatedDelayLatencyMs = clus.Tester.DelayLatencyMs
	if random {
		clus.UpdateDelayLatencyMs()
		cc.rpcpbCase = rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT
	}
	c := &caseLeader{cc, -1, -1}
	return &caseUntilSnapshot{rpcpbCase: cc.rpcpbCase, Case: c}
}
func new_Case_DELAY_PEER_PORT_TX_RX_QUORUM(clus *Cluster, random bool) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &caseQuorum{caseByFunc: caseByFunc{rpcpbCase: rpcpb.Case_DELAY_PEER_PORT_TX_RX_QUORUM, injectMember: inject_DELAY_PEER_PORT_TX_RX, recoverMember: recover_DELAY_PEER_PORT_TX_RX}, injected: make(map[int]struct{})}
	clus.Tester.UpdatedDelayLatencyMs = clus.Tester.DelayLatencyMs
	if random {
		clus.UpdateDelayLatencyMs()
		c.rpcpbCase = rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_QUORUM
	}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
func new_Case_DELAY_PEER_PORT_TX_RX_ALL(clus *Cluster, random bool) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &caseAll{rpcpbCase: rpcpb.Case_DELAY_PEER_PORT_TX_RX_ALL, injectMember: inject_DELAY_PEER_PORT_TX_RX, recoverMember: recover_DELAY_PEER_PORT_TX_RX}
	clus.Tester.UpdatedDelayLatencyMs = clus.Tester.DelayLatencyMs
	if random {
		clus.UpdateDelayLatencyMs()
		c.rpcpbCase = rpcpb.Case_RANDOM_DELAY_PEER_PORT_TX_RX_ALL
	}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
