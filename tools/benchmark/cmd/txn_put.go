package cmd

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"time"
	v3 "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/report"
	"github.com/spf13/cobra"
	"golang.org/x/time/rate"
	"gopkg.in/cheggaaa/pb.v1"
)

var txnPutCmd = &cobra.Command{Use: "txn-put", Short: "Benchmark txn-put", Run: txnPutFunc}
var (
	txnPutTotal	int
	txnPutRate	int
	txnPutOpsPerTxn	int
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	RootCmd.AddCommand(txnPutCmd)
	txnPutCmd.Flags().IntVar(&keySize, "key-size", 8, "Key size of txn put")
	txnPutCmd.Flags().IntVar(&valSize, "val-size", 8, "Value size of txn put")
	txnPutCmd.Flags().IntVar(&txnPutOpsPerTxn, "txn-ops", 1, "Number of puts per txn")
	txnPutCmd.Flags().IntVar(&txnPutRate, "rate", 0, "Maximum txns per second (0 is no limit)")
	txnPutCmd.Flags().IntVar(&txnPutTotal, "total", 10000, "Total number of txn requests")
	txnPutCmd.Flags().IntVar(&keySpaceSize, "key-space-size", 1, "Maximum possible keys")
}
func txnPutFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if keySpaceSize <= 0 {
		fmt.Fprintf(os.Stderr, "expected positive --key-space-size, got (%v)", keySpaceSize)
		os.Exit(1)
	}
	requests := make(chan []v3.Op, totalClients)
	if txnPutRate == 0 {
		txnPutRate = math.MaxInt32
	}
	limit := rate.NewLimiter(rate.Limit(txnPutRate), 1)
	clients := mustCreateClients(totalClients, totalConns)
	k, v := make([]byte, keySize), string(mustRandBytes(valSize))
	bar = pb.New(txnPutTotal)
	bar.Format("Bom !")
	bar.Start()
	r := newReport()
	for i := range clients {
		wg.Add(1)
		go func(c *v3.Client) {
			defer wg.Done()
			for ops := range requests {
				limit.Wait(context.Background())
				st := time.Now()
				_, err := c.Txn(context.TODO()).Then(ops...).Commit()
				r.Results() <- report.Result{Err: err, Start: st, End: time.Now()}
				bar.Increment()
			}
		}(clients[i])
	}
	go func() {
		for i := 0; i < txnPutTotal; i++ {
			ops := make([]v3.Op, txnPutOpsPerTxn)
			for j := 0; j < txnPutOpsPerTxn; j++ {
				binary.PutVarint(k, int64(((i*txnPutOpsPerTxn)+j)%keySpaceSize))
				ops[j] = v3.OpPut(string(k), v)
			}
			requests <- ops
		}
		close(requests)
	}()
	rc := r.Run()
	wg.Wait()
	close(r.Results())
	bar.Finish()
	fmt.Println(<-rc)
}
