package tester

import (
	"time"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
)

type caseNoFailWithStress caseByFunc

func (c *caseNoFailWithStress) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (c *caseNoFailWithStress) Recover(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (c *caseNoFailWithStress) Desc() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.desc != "" {
		return c.desc
	}
	return c.rpcpbCase.String()
}
func (c *caseNoFailWithStress) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}
func new_Case_NO_FAIL_WITH_STRESS(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &caseNoFailWithStress{rpcpbCase: rpcpb.Case_NO_FAIL_WITH_STRESS}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}

type caseNoFailWithNoStressForLiveness caseByFunc

func (c *caseNoFailWithNoStressForLiveness) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus.lg.Info("extra delay for liveness mode with no stresser", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.String("desc", c.Desc()))
	time.Sleep(clus.GetCaseDelayDuration())
	clus.lg.Info("wait health in liveness mode", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.String("desc", c.Desc()))
	return clus.WaitHealth()
}
func (c *caseNoFailWithNoStressForLiveness) Recover(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (c *caseNoFailWithNoStressForLiveness) Desc() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.desc != "" {
		return c.desc
	}
	return c.rpcpbCase.String()
}
func (c *caseNoFailWithNoStressForLiveness) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}
func new_Case_NO_FAIL_WITH_NO_STRESS_FOR_LIVENESS(clus *Cluster) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &caseNoFailWithNoStressForLiveness{rpcpbCase: rpcpb.Case_NO_FAIL_WITH_NO_STRESS_FOR_LIVENESS}
	return &caseDelay{Case: c, delayDuration: clus.GetCaseDelayDuration()}
}
