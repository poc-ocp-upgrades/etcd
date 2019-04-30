package tester

import "go.etcd.io/etcd/functional/rpcpb"

func inject_BLACKHOLE_PEER_PORT_TX_RX(clus *Cluster, idx int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clus.sendOp(idx, rpcpb.Operation_BLACKHOLE_PEER_PORT_TX_RX)
}
func recover_BLACKHOLE_PEER_PORT_TX_RX(clus *Cluster, idx int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clus.sendOp(idx, rpcpb.Operation_UNBLACKHOLE_PEER_PORT_TX_RX)
}
func new_Case_BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER, injectMember: inject_BLACKHOLE_PEER_PORT_TX_RX, recoverMember: recover_BLACKHOLE_PEER_PORT_TX_RX}
	c := &caseFollower{cc, -1, -1}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
func new_Case_BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT() Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT, injectMember: inject_BLACKHOLE_PEER_PORT_TX_RX, recoverMember: recover_BLACKHOLE_PEER_PORT_TX_RX}
	c := &caseFollower{cc, -1, -1}
	return &caseUntilSnapshot{rpcpbCase: rpcpb.Case_BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT, Case: c}
}
func new_Case_BLACKHOLE_PEER_PORT_TX_RX_LEADER(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_BLACKHOLE_PEER_PORT_TX_RX_LEADER, injectMember: inject_BLACKHOLE_PEER_PORT_TX_RX, recoverMember: recover_BLACKHOLE_PEER_PORT_TX_RX}
	c := &caseLeader{cc, -1, -1}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
func new_Case_BLACKHOLE_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT() Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_BLACKHOLE_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT, injectMember: inject_BLACKHOLE_PEER_PORT_TX_RX, recoverMember: recover_BLACKHOLE_PEER_PORT_TX_RX}
	c := &caseLeader{cc, -1, -1}
	return &caseUntilSnapshot{rpcpbCase: rpcpb.Case_BLACKHOLE_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT, Case: c}
}
func new_Case_BLACKHOLE_PEER_PORT_TX_RX_QUORUM(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &caseQuorum{caseByFunc: caseByFunc{rpcpbCase: rpcpb.Case_BLACKHOLE_PEER_PORT_TX_RX_QUORUM, injectMember: inject_BLACKHOLE_PEER_PORT_TX_RX, recoverMember: recover_BLACKHOLE_PEER_PORT_TX_RX}, injected: make(map[int]struct{})}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
func new_Case_BLACKHOLE_PEER_PORT_TX_RX_ALL(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &caseAll{rpcpbCase: rpcpb.Case_BLACKHOLE_PEER_PORT_TX_RX_ALL, injectMember: inject_BLACKHOLE_PEER_PORT_TX_RX, recoverMember: recover_BLACKHOLE_PEER_PORT_TX_RX}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
