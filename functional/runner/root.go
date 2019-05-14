package runner

import (
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"time"
)

const (
	cliName            = "etcd-runner"
	cliDescription     = "Stress tests using clientv3 functionality.."
	defaultDialTimeout = 2 * time.Second
)

var (
	rootCmd = &cobra.Command{Use: cliName, Short: cliDescription, SuggestFor: []string{"etcd-runner"}}
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cobra.EnablePrefixMatching = true
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.Lmicroseconds)
	rootCmd.PersistentFlags().StringSliceVar(&endpoints, "endpoints", []string{"127.0.0.1:2379"}, "gRPC endpoints")
	rootCmd.PersistentFlags().DurationVar(&dialTimeout, "dial-timeout", defaultDialTimeout, "dial timeout for client connections")
	rootCmd.PersistentFlags().IntVar(&reqRate, "req-rate", 30, "maximum number of requests per second")
	rootCmd.PersistentFlags().IntVar(&rounds, "rounds", 100, "number of rounds to run; 0 to run forever")
	rootCmd.AddCommand(NewElectionCommand(), NewLeaseRenewerCommand(), NewLockRacerCommand(), NewWatchCommand())
}
func Start() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd.SetUsageFunc(usageFunc)
	rootCmd.SetHelpTemplate(`{{.UsageString}}`)
	if err := rootCmd.Execute(); err != nil {
		ExitWithError(ExitError, err)
	}
}
