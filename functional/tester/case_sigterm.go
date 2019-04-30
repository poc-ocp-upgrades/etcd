package tester

import "go.etcd.io/etcd/functional/rpcpb"

func inject_SIGTERM_ETCD(clus *Cluster, idx int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clus.sendOp(idx, rpcpb.Operation_SIGTERM_ETCD)
}
func recover_SIGTERM_ETCD(clus *Cluster, idx int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clus.sendOp(idx, rpcpb.Operation_RESTART_ETCD)
}
func new_Case_SIGTERM_ONE_FOLLOWER(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_SIGTERM_ONE_FOLLOWER, injectMember: inject_SIGTERM_ETCD, recoverMember: recover_SIGTERM_ETCD}
	c := &caseFollower{cc, -1, -1}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
func new_Case_SIGTERM_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &caseUntilSnapshot{rpcpbCase: rpcpb.Case_SIGTERM_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT, Case: new_Case_SIGTERM_ONE_FOLLOWER(clus)}
}
func new_Case_SIGTERM_LEADER(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := caseByFunc{rpcpbCase: rpcpb.Case_SIGTERM_LEADER, injectMember: inject_SIGTERM_ETCD, recoverMember: recover_SIGTERM_ETCD}
	c := &caseLeader{cc, -1, -1}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
func new_Case_SIGTERM_LEADER_UNTIL_TRIGGER_SNAPSHOT(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &caseUntilSnapshot{rpcpbCase: rpcpb.Case_SIGTERM_LEADER_UNTIL_TRIGGER_SNAPSHOT, Case: new_Case_SIGTERM_LEADER(clus)}
}
func new_Case_SIGTERM_QUORUM(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &caseQuorum{caseByFunc: caseByFunc{rpcpbCase: rpcpb.Case_SIGTERM_QUORUM, injectMember: inject_SIGTERM_ETCD, recoverMember: recover_SIGTERM_ETCD}, injected: make(map[int]struct{})}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
func new_Case_SIGTERM_ALL(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &caseAll{rpcpbCase: rpcpb.Case_SIGTERM_ALL, injectMember: inject_SIGTERM_ETCD, recoverMember: recover_SIGTERM_ETCD}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
