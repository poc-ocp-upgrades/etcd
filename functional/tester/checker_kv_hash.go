package tester

import (
	"fmt"
	"time"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
)

const retries = 7

type kvHashChecker struct {
	ctype	rpcpb.Checker
	clus	*Cluster
}

func newKVHashChecker(clus *Cluster) Checker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &kvHashChecker{ctype: rpcpb.Checker_KV_HASH, clus: clus}
}
func (hc *kvHashChecker) checkRevAndHashes() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		revs	map[string]int64
		hashes	map[string]int64
	)
	for i := 0; i < retries; i++ {
		revs, hashes, err = hc.clus.getRevisionHash()
		if err != nil {
			hc.clus.lg.Warn("failed to get revision and hash", zap.Int("retries", i), zap.Error(err))
		} else {
			sameRev := getSameValue(revs)
			sameHashes := getSameValue(hashes)
			if sameRev && sameHashes {
				return nil
			}
			hc.clus.lg.Warn("retrying; etcd cluster is not stable", zap.Int("retries", i), zap.Bool("same-revisions", sameRev), zap.Bool("same-hashes", sameHashes), zap.String("revisions", fmt.Sprintf("%+v", revs)), zap.String("hashes", fmt.Sprintf("%+v", hashes)))
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		return fmt.Errorf("failed revision and hash check (%v)", err)
	}
	return fmt.Errorf("etcd cluster is not stable: [revisions: %v] and [hashes: %v]", revs, hashes)
}
func (hc *kvHashChecker) Type() rpcpb.Checker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return hc.ctype
}
func (hc *kvHashChecker) EtcdClientEndpoints() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return hc.clus.EtcdClientEndpoints()
}
func (hc *kvHashChecker) Check() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return hc.checkRevAndHashes()
}
