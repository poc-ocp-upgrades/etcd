package runner

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/stringutil"
	"github.com/spf13/cobra"
	"golang.org/x/time/rate"
)

var (
	runningTime	time.Duration
	noOfPrefixes	int
	watchPerPrefix	int
	watchPrefix	string
	totalKeys	int
)

func NewWatchCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "watcher", Short: "Performs watch operation", Run: runWatcherFunc}
	cmd.Flags().DurationVar(&runningTime, "running-time", 60, "number of seconds to run")
	cmd.Flags().StringVar(&watchPrefix, "prefix", "", "the prefix to append on all keys")
	cmd.Flags().IntVar(&noOfPrefixes, "total-prefixes", 10, "total no of prefixes to use")
	cmd.Flags().IntVar(&watchPerPrefix, "watch-per-prefix", 10, "number of watchers per prefix")
	cmd.Flags().IntVar(&totalKeys, "total-keys", 1000, "total number of keys to watch")
	return cmd
}
func runWatcherFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) > 0 {
		ExitWithError(ExitBadArgs, errors.New("watcher does not take any argument"))
	}
	ctx := context.Background()
	for round := 0; round < rounds || rounds <= 0; round++ {
		fmt.Println("round", round)
		performWatchOnPrefixes(ctx, cmd, round)
	}
}
func performWatchOnPrefixes(ctx context.Context, cmd *cobra.Command, round int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keyPerPrefix := totalKeys / noOfPrefixes
	prefixes := stringutil.UniqueStrings(5, noOfPrefixes)
	keys := stringutil.RandomStrings(10, keyPerPrefix)
	roundPrefix := fmt.Sprintf("%16x", round)
	eps := endpointsFromFlag(cmd)
	var (
		revision	int64
		wg		sync.WaitGroup
		gr		*clientv3.GetResponse
		err		error
	)
	client := newClient(eps, dialTimeout)
	defer client.Close()
	gr, err = getKey(ctx, client, "non-existent")
	if err != nil {
		log.Fatalf("failed to get the initial revision: %v", err)
	}
	revision = gr.Header.Revision
	ctxt, cancel := context.WithDeadline(ctx, time.Now().Add(runningTime*time.Second))
	defer cancel()
	limiter := rate.NewLimiter(rate.Limit(reqRate), reqRate)
	go func() {
		for _, key := range keys {
			for _, prefix := range prefixes {
				if err = limiter.Wait(ctxt); err != nil {
					return
				}
				if err = putKeyAtMostOnce(ctxt, client, watchPrefix+"-"+roundPrefix+"-"+prefix+"-"+key); err != nil {
					log.Fatalf("failed to put key: %v", err)
					return
				}
			}
		}
	}()
	ctxc, cancelc := context.WithCancel(ctx)
	wcs := make([]clientv3.WatchChan, 0)
	rcs := make([]*clientv3.Client, 0)
	for _, prefix := range prefixes {
		for j := 0; j < watchPerPrefix; j++ {
			rc := newClient(eps, dialTimeout)
			rcs = append(rcs, rc)
			wprefix := watchPrefix + "-" + roundPrefix + "-" + prefix
			wc := rc.Watch(ctxc, wprefix, clientv3.WithPrefix(), clientv3.WithRev(revision))
			wcs = append(wcs, wc)
			wg.Add(1)
			go func() {
				defer wg.Done()
				checkWatchResponse(wc, wprefix, keys)
			}()
		}
	}
	wg.Wait()
	cancelc()
	for e, wc := range wcs {
		if _, ok := <-wc; ok {
			log.Fatalf("expected wc to be closed, but received %v", e)
		}
	}
	for _, rc := range rcs {
		rc.Close()
	}
	if err = deletePrefix(ctx, client, watchPrefix); err != nil {
		log.Fatalf("failed to clean up keys after test: %v", err)
	}
}
func checkWatchResponse(wc clientv3.WatchChan, prefix string, keys []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for n := 0; n < len(keys); {
		wr, more := <-wc
		if !more {
			log.Fatalf("expect more keys (received %d/%d) for %s", n, len(keys), prefix)
		}
		for _, event := range wr.Events {
			expectedKey := prefix + "-" + keys[n]
			receivedKey := string(event.Kv.Key)
			if expectedKey != receivedKey {
				log.Fatalf("expected key %q, got %q for prefix : %q\n", expectedKey, receivedKey, prefix)
			}
			n++
		}
	}
}
func putKeyAtMostOnce(ctx context.Context, client *clientv3.Client, key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	gr, err := getKey(ctx, client, key)
	if err != nil {
		return err
	}
	var modrev int64
	if len(gr.Kvs) > 0 {
		modrev = gr.Kvs[0].ModRevision
	}
	for ctx.Err() == nil {
		_, err := client.Txn(ctx).If(clientv3.Compare(clientv3.ModRevision(key), "=", modrev)).Then(clientv3.OpPut(key, key)).Commit()
		if err == nil {
			return nil
		}
	}
	return ctx.Err()
}
func deletePrefix(ctx context.Context, client *clientv3.Client, key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for ctx.Err() == nil {
		if _, err := client.Delete(ctx, key, clientv3.WithPrefix()); err == nil {
			return nil
		}
	}
	return ctx.Err()
}
func getKey(ctx context.Context, client *clientv3.Client, key string) (*clientv3.GetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for ctx.Err() == nil {
		if gr, err := client.Get(ctx, key); err == nil {
			return gr, nil
		}
	}
	return nil, ctx.Err()
}
