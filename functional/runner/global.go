package runner

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/cobra"
	"golang.org/x/time/rate"
	"log"
	"sync"
	"time"
)

var (
	totalClientConnections int
	endpoints              []string
	dialTimeout            time.Duration
	rounds                 int
	reqRate                int
)

type roundClient struct {
	c        *clientv3.Client
	progress int
	acquire  func() error
	validate func() error
	release  func() error
}

func newClient(eps []string, timeout time.Duration) *clientv3.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, err := clientv3.New(clientv3.Config{Endpoints: eps, DialTimeout: time.Duration(timeout) * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return c
}
func doRounds(rcs []roundClient, rounds int, requests int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wg sync.WaitGroup
	wg.Add(len(rcs))
	finished := make(chan struct{})
	limiter := rate.NewLimiter(rate.Limit(reqRate), reqRate)
	for i := range rcs {
		go func(rc *roundClient) {
			defer wg.Done()
			for rc.progress < rounds || rounds <= 0 {
				if err := limiter.WaitN(context.Background(), requests/len(rcs)); err != nil {
					log.Panicf("rate limiter error %v", err)
				}
				for rc.acquire() != nil {
				}
				if err := rc.validate(); err != nil {
					log.Fatal(err)
				}
				time.Sleep(10 * time.Millisecond)
				rc.progress++
				finished <- struct{}{}
				for rc.release() != nil {
				}
			}
		}(&rcs[i])
	}
	start := time.Now()
	for i := 1; i < len(rcs)*rounds+1 || rounds <= 0; i++ {
		select {
		case <-finished:
			if i%100 == 0 {
				fmt.Printf("finished %d, took %v\n", i, time.Since(start))
				start = time.Now()
			}
		case <-time.After(time.Minute):
			log.Panic("no progress after 1 minute!")
		}
	}
	wg.Wait()
	for _, rc := range rcs {
		rc.c.Close()
	}
}
func endpointsFromFlag(cmd *cobra.Command) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eps, err := cmd.Flags().GetStringSlice("endpoints")
	if err != nil {
		ExitWithError(ExitError, err)
	}
	return eps
}
