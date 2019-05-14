package tester

import (
	"go.uber.org/zap"
	"time"
)

type caseDelay struct {
	Case
	delayDuration time.Duration
}

func (c *caseDelay) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := c.Case.Inject(clus); err != nil {
		return err
	}
	if c.delayDuration > 0 {
		clus.lg.Info("wait after inject", zap.Duration("delay", c.delayDuration), zap.String("desc", c.Case.Desc()))
		time.Sleep(c.delayDuration)
	}
	return nil
}
