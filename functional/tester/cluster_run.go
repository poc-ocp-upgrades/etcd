package tester

import (
	"fmt"
	"os"
	"time"
	"github.com/coreos/etcd/functional/rpcpb"
	"github.com/coreos/etcd/pkg/fileutil"
	"go.uber.org/zap"
)

const compactQPS = 50000

func (clus *Cluster) Run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer printReport()
	if err := fileutil.TouchDirAll(clus.Tester.DataDir); err != nil {
		clus.lg.Panic("failed to create test data directory", zap.String("dir", clus.Tester.DataDir), zap.Error(err))
	}
	var preModifiedKey int64
	for round := 0; round < int(clus.Tester.RoundLimit) || clus.Tester.RoundLimit == -1; round++ {
		roundTotalCounter.Inc()
		clus.rd = round
		if err := clus.doRound(); err != nil {
			clus.lg.Warn("round FAIL", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.Error(err))
			if clus.cleanup() != nil {
				return
			}
			preModifiedKey = 0
			continue
		}
		clus.cs = -1
		revToCompact := max(0, clus.currentRevision-10000)
		currentModifiedKey := clus.stresser.ModifiedKeys()
		modifiedKey := currentModifiedKey - preModifiedKey
		preModifiedKey = currentModifiedKey
		timeout := 10 * time.Second
		timeout += time.Duration(modifiedKey/compactQPS) * time.Second
		clus.lg.Info("compact START", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.Duration("timeout", timeout))
		if err := clus.compact(revToCompact, timeout); err != nil {
			clus.lg.Warn("compact FAIL", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.Error(err))
			if err = clus.cleanup(); err != nil {
				clus.lg.Warn("cleanup FAIL", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.Error(err))
				return
			}
			preModifiedKey = 0
		}
		if round > 0 && round%500 == 0 {
			if err := clus.defrag(); err != nil {
				clus.failed()
				return
			}
		}
	}
	clus.lg.Info("functional-tester PASS", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)))
}
func (clus *Cluster) doRound() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if clus.Tester.CaseShuffle {
		clus.shuffleCases()
	}
	roundNow := time.Now()
	clus.lg.Info("round START", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.Strings("cases", clus.listCases()))
	for i, fa := range clus.cases {
		clus.cs = i
		caseTotal[fa.Desc()]++
		caseTotalCounter.WithLabelValues(fa.Desc()).Inc()
		caseNow := time.Now()
		clus.lg.Info("case START", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.String("desc", fa.Desc()))
		clus.lg.Info("wait health before injecting failures")
		if err := clus.WaitHealth(); err != nil {
			return fmt.Errorf("wait full health error: %v", err)
		}
		stressStarted := false
		fcase := fa.TestCase()
		if fcase != rpcpb.Case_NO_FAIL_WITH_NO_STRESS_FOR_LIVENESS {
			clus.lg.Info("stress START", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.String("desc", fa.Desc()))
			if err := clus.stresser.Stress(); err != nil {
				return fmt.Errorf("start stresser error: %v", err)
			}
			stressStarted = true
		}
		clus.lg.Info("inject START", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.String("desc", fa.Desc()))
		if err := fa.Inject(clus); err != nil {
			return fmt.Errorf("injection error: %v", err)
		}
		clus.lg.Info("recover START", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.String("desc", fa.Desc()))
		if err := fa.Recover(clus); err != nil {
			return fmt.Errorf("recovery error: %v", err)
		}
		if stressStarted {
			clus.lg.Info("stress PAUSE", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.String("desc", fa.Desc()))
			ems := clus.stresser.Pause()
			if fcase == rpcpb.Case_NO_FAIL_WITH_STRESS && len(ems) > 0 {
				ess := make([]string, 0, len(ems))
				cnt := 0
				for k, v := range ems {
					ess = append(ess, fmt.Sprintf("%s (count: %d)", k, v))
					cnt += v
				}
				clus.lg.Warn("expected no errors", zap.String("desc", fa.Desc()), zap.Strings("errors", ess))
				if cnt > int(clus.Tester.StressQPS)/10 {
					return fmt.Errorf("expected no error in %q, got %q", fcase.String(), ess)
				}
			}
		}
		clus.lg.Info("health check START", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.String("desc", fa.Desc()))
		if err := clus.WaitHealth(); err != nil {
			return fmt.Errorf("wait full health error: %v", err)
		}
		checkerFailExceptions := []rpcpb.Checker{}
		switch fcase {
		case rpcpb.Case_SIGQUIT_AND_REMOVE_QUORUM_AND_RESTORE_LEADER_SNAPSHOT_FROM_SCRATCH:
			checkerFailExceptions = append(checkerFailExceptions, rpcpb.Checker_LEASE_EXPIRE)
		}
		clus.lg.Info("consistency check START", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.String("desc", fa.Desc()))
		if err := clus.runCheckers(checkerFailExceptions...); err != nil {
			return fmt.Errorf("consistency check error (%v)", err)
		}
		clus.lg.Info("consistency check PASS", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.String("desc", fa.Desc()), zap.Duration("took", time.Since(caseNow)))
	}
	clus.lg.Info("round ALL PASS", zap.Int("round", clus.rd), zap.Strings("cases", clus.listCases()), zap.Int("case-total", len(clus.cases)), zap.Duration("took", time.Since(roundNow)))
	return nil
}
func (clus *Cluster) updateRevision() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	revs, _, err := clus.getRevisionHash()
	for _, rev := range revs {
		clus.currentRevision = rev
		break
	}
	clus.lg.Info("updated current revision", zap.Int64("current-revision", clus.currentRevision))
	return err
}
func (clus *Cluster) compact(rev int64, timeout time.Duration) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err = clus.compactKV(rev, timeout); err != nil {
		clus.lg.Warn("compact FAIL", zap.Int64("current-revision", clus.currentRevision), zap.Int64("compact-revision", rev), zap.Error(err))
		return err
	}
	clus.lg.Info("compact DONE", zap.Int64("current-revision", clus.currentRevision), zap.Int64("compact-revision", rev))
	if err = clus.checkCompact(rev); err != nil {
		clus.lg.Warn("check compact FAIL", zap.Int64("current-revision", clus.currentRevision), zap.Int64("compact-revision", rev), zap.Error(err))
		return err
	}
	clus.lg.Info("check compact DONE", zap.Int64("current-revision", clus.currentRevision), zap.Int64("compact-revision", rev))
	return nil
}
func (clus *Cluster) failed() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus.lg.Info("functional-tester FAIL", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)))
	clus.Send_SIGQUIT_ETCD_AND_REMOVE_DATA_AND_STOP_AGENT()
	os.Exit(2)
}
func (clus *Cluster) cleanup() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if clus.Tester.ExitOnCaseFail {
		defer clus.failed()
	}
	roundFailedTotalCounter.Inc()
	desc := "compact/defrag"
	if clus.cs != -1 {
		desc = clus.cases[clus.cs].Desc()
	}
	caseFailedTotalCounter.WithLabelValues(desc).Inc()
	clus.lg.Info("closing stressers before archiving failure data", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)))
	clus.stresser.Close()
	if err := clus.send_SIGQUIT_ETCD_AND_ARCHIVE_DATA(); err != nil {
		clus.lg.Warn("cleanup FAIL", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.Error(err))
		return err
	}
	if err := clus.send_RESTART_ETCD(); err != nil {
		clus.lg.Warn("restart FAIL", zap.Int("round", clus.rd), zap.Int("case", clus.cs), zap.Int("case-total", len(clus.cases)), zap.Error(err))
		return err
	}
	clus.setStresserChecker()
	return nil
}
