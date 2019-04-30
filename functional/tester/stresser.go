package tester

import (
	"fmt"
	"time"
	"go.etcd.io/etcd/functional/rpcpb"
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
	ks := &keyStresser{lg: clus.lg, m: m, keySize: int(clus.Tester.StressKeySize), keyLargeSize: int(clus.Tester.StressKeySizeLarge), keySuffixRange: int(clus.Tester.StressKeySuffixRange), keyTxnSuffixRange: int(clus.Tester.StressKeySuffixRangeTxn), keyTxnOps: int(clus.Tester.StressKeyTxnOps), clientsN: int(clus.Tester.StressClients), rateLimiter: clus.rateLimiter}
	ksExist := false
	for _, s := range clus.Tester.Stressers {
		clus.lg.Info("creating stresser", zap.String("type", s.Type), zap.Float64("weight", s.Weight), zap.String("endpoint", m.EtcdClientEndpoint))
		switch s.Type {
		case "KV_WRITE_SMALL":
			ksExist = true
			ks.weightKVWriteSmall = s.Weight
		case "KV_WRITE_LARGE":
			ksExist = true
			ks.weightKVWriteLarge = s.Weight
		case "KV_READ_ONE_KEY":
			ksExist = true
			ks.weightKVReadOneKey = s.Weight
		case "KV_READ_RANGE":
			ksExist = true
			ks.weightKVReadRange = s.Weight
		case "KV_DELETE_ONE_KEY":
			ksExist = true
			ks.weightKVDeleteOneKey = s.Weight
		case "KV_DELETE_RANGE":
			ksExist = true
			ks.weightKVDeleteRange = s.Weight
		case "KV_TXN_WRITE_DELETE":
			ksExist = true
			ks.weightKVTxnWriteDelete = s.Weight
		case "LEASE":
			stressers = append(stressers, &leaseStresser{stype: rpcpb.StresserType_LEASE, lg: clus.lg, m: m, numLeases: 10, keysPerLease: 10, rateLimiter: clus.rateLimiter})
		case "ELECTION_RUNNER":
			reqRate := 100
			args := []string{"election", fmt.Sprintf("%v", time.Now().UnixNano()), "--dial-timeout=10s", "--endpoints", m.EtcdClientEndpoint, "--total-client-connections=10", "--rounds=0", "--req-rate", fmt.Sprintf("%v", reqRate)}
			stressers = append(stressers, newRunnerStresser(rpcpb.StresserType_ELECTION_RUNNER, m.EtcdClientEndpoint, clus.lg, clus.Tester.RunnerExecPath, args, clus.rateLimiter, reqRate))
		case "WATCH_RUNNER":
			reqRate := 100
			args := []string{"watcher", "--prefix", fmt.Sprintf("%v", time.Now().UnixNano()), "--total-keys=1", "--total-prefixes=1", "--watch-per-prefix=1", "--endpoints", m.EtcdClientEndpoint, "--rounds=0", "--req-rate", fmt.Sprintf("%v", reqRate)}
			stressers = append(stressers, newRunnerStresser(rpcpb.StresserType_WATCH_RUNNER, m.EtcdClientEndpoint, clus.lg, clus.Tester.RunnerExecPath, args, clus.rateLimiter, reqRate))
		case "LOCK_RACER_RUNNER":
			reqRate := 100
			args := []string{"lock-racer", fmt.Sprintf("%v", time.Now().UnixNano()), "--endpoints", m.EtcdClientEndpoint, "--total-client-connections=10", "--rounds=0", "--req-rate", fmt.Sprintf("%v", reqRate)}
			stressers = append(stressers, newRunnerStresser(rpcpb.StresserType_LOCK_RACER_RUNNER, m.EtcdClientEndpoint, clus.lg, clus.Tester.RunnerExecPath, args, clus.rateLimiter, reqRate))
		case "LEASE_RUNNER":
			args := []string{"lease-renewer", "--ttl=30", "--endpoints", m.EtcdClientEndpoint}
			stressers = append(stressers, newRunnerStresser(rpcpb.StresserType_LEASE_RUNNER, m.EtcdClientEndpoint, clus.lg, clus.Tester.RunnerExecPath, args, clus.rateLimiter, 0))
		}
	}
	if ksExist {
		return append(stressers, ks)
	}
	return stressers
}
