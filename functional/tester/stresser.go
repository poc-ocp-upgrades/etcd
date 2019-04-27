package tester

import (
	"fmt"
	"time"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
)

type Stresser interface {
	Stress() error
	Pause() map[string]int
	Close() map[string]int
	ModifiedKeys() int64
}

func newStresser(clus *Cluster, m *rpcpb.Member) (stressers []Stresser) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	stressers = make([]Stresser, len(clus.Tester.Stressers))
	for i, stype := range clus.Tester.Stressers {
		clus.lg.Info("creating stresser", zap.String("type", stype), zap.String("endpoint", m.EtcdClientEndpoint))
		switch stype {
		case "KV":
			stressers[i] = &keyStresser{stype: rpcpb.Stresser_KV, lg: clus.lg, m: m, keySize: int(clus.Tester.StressKeySize), keyLargeSize: int(clus.Tester.StressKeySizeLarge), keySuffixRange: int(clus.Tester.StressKeySuffixRange), keyTxnSuffixRange: int(clus.Tester.StressKeySuffixRangeTxn), keyTxnOps: int(clus.Tester.StressKeyTxnOps), clientsN: int(clus.Tester.StressClients), rateLimiter: clus.rateLimiter}
		case "LEASE":
			stressers[i] = &leaseStresser{stype: rpcpb.Stresser_LEASE, lg: clus.lg, m: m, numLeases: 10, keysPerLease: 10, rateLimiter: clus.rateLimiter}
		case "ELECTION_RUNNER":
			reqRate := 100
			args := []string{"election", fmt.Sprintf("%v", time.Now().UnixNano()), "--dial-timeout=10s", "--endpoints", m.EtcdClientEndpoint, "--total-client-connections=10", "--rounds=0", "--req-rate", fmt.Sprintf("%v", reqRate)}
			stressers[i] = newRunnerStresser(rpcpb.Stresser_ELECTION_RUNNER, m.EtcdClientEndpoint, clus.lg, clus.Tester.RunnerExecPath, args, clus.rateLimiter, reqRate)
		case "WATCH_RUNNER":
			reqRate := 100
			args := []string{"watcher", "--prefix", fmt.Sprintf("%v", time.Now().UnixNano()), "--total-keys=1", "--total-prefixes=1", "--watch-per-prefix=1", "--endpoints", m.EtcdClientEndpoint, "--rounds=0", "--req-rate", fmt.Sprintf("%v", reqRate)}
			stressers[i] = newRunnerStresser(rpcpb.Stresser_WATCH_RUNNER, m.EtcdClientEndpoint, clus.lg, clus.Tester.RunnerExecPath, args, clus.rateLimiter, reqRate)
		case "LOCK_RACER_RUNNER":
			reqRate := 100
			args := []string{"lock-racer", fmt.Sprintf("%v", time.Now().UnixNano()), "--endpoints", m.EtcdClientEndpoint, "--total-client-connections=10", "--rounds=0", "--req-rate", fmt.Sprintf("%v", reqRate)}
			stressers[i] = newRunnerStresser(rpcpb.Stresser_LOCK_RACER_RUNNER, m.EtcdClientEndpoint, clus.lg, clus.Tester.RunnerExecPath, args, clus.rateLimiter, reqRate)
		case "LEASE_RUNNER":
			args := []string{"lease-renewer", "--ttl=30", "--endpoints", m.EtcdClientEndpoint}
			stressers[i] = newRunnerStresser(rpcpb.Stresser_LEASE_RUNNER, m.EtcdClientEndpoint, clus.lg, clus.Tester.RunnerExecPath, args, clus.rateLimiter, 0)
		}
	}
	return stressers
}
