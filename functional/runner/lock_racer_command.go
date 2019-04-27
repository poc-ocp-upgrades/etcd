package runner

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/spf13/cobra"
)

func NewLockRacerCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "lock-racer [name of lock (defaults to 'racers')]", Short: "Performs lock race operation", Run: runRacerFunc}
	cmd.Flags().IntVar(&totalClientConnections, "total-client-connections", 10, "total number of client connections")
	return cmd
}
func runRacerFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	racers := "racers"
	if len(args) == 1 {
		racers = args[0]
	}
	if len(args) > 1 {
		ExitWithError(ExitBadArgs, errors.New("lock-racer takes at most one argument"))
	}
	rcs := make([]roundClient, totalClientConnections)
	ctx := context.Background()
	var mu sync.Mutex
	cnt := 0
	eps := endpointsFromFlag(cmd)
	for i := range rcs {
		var (
			s	*concurrency.Session
			err	error
		)
		rcs[i].c = newClient(eps, dialTimeout)
		for {
			s, err = concurrency.NewSession(rcs[i].c)
			if err == nil {
				break
			}
		}
		m := concurrency.NewMutex(s, racers)
		rcs[i].acquire = func() error {
			return m.Lock(ctx)
		}
		rcs[i].validate = func() error {
			mu.Lock()
			defer mu.Unlock()
			if cnt++; cnt != 1 {
				return fmt.Errorf("bad lock; count: %d", cnt)
			}
			return nil
		}
		rcs[i].release = func() error {
			mu.Lock()
			defer mu.Unlock()
			if err := m.Unlock(ctx); err != nil {
				return err
			}
			cnt = 0
			return nil
		}
	}
	doRounds(rcs, rounds, 2*len(rcs))
}
