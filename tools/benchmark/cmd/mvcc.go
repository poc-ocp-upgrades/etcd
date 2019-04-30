package cmd

import (
	"os"
	"time"
	"go.uber.org/zap"
	"go.etcd.io/etcd/lease"
	"go.etcd.io/etcd/mvcc"
	"go.etcd.io/etcd/mvcc/backend"
	"github.com/spf13/cobra"
)

var (
	batchInterval	int
	batchLimit	int
	s		mvcc.KV
)

func initMVCC() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bcfg := backend.DefaultBackendConfig()
	bcfg.Path, bcfg.BatchInterval, bcfg.BatchLimit = "mvcc-bench", time.Duration(batchInterval)*time.Millisecond, batchLimit
	be := backend.New(bcfg)
	s = mvcc.NewStore(zap.NewExample(), be, &lease.FakeLessor{}, nil)
	os.Remove("mvcc-bench")
}

var mvccCmd = &cobra.Command{Use: "mvcc", Short: "Benchmark mvcc", Long: `storage subcommand is a set of various benchmark tools for MVCC storage subsystem of etcd.
Actual benchmarks are implemented as its subcommands.`, PersistentPreRun: mvccPreRun}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	RootCmd.AddCommand(mvccCmd)
	mvccCmd.PersistentFlags().IntVar(&batchInterval, "batch-interval", 100, "Interval of batching (milliseconds)")
	mvccCmd.PersistentFlags().IntVar(&batchLimit, "batch-limit", 10000, "A limit of batched transaction")
}
func mvccPreRun(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	initMVCC()
}
