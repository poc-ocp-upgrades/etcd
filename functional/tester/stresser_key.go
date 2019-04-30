package tester

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"go.etcd.io/etcd/functional/rpcpb"
	"go.etcd.io/etcd/raft"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type keyStresser struct {
	lg			*zap.Logger
	m			*rpcpb.Member
	weightKVWriteSmall	float64
	weightKVWriteLarge	float64
	weightKVReadOneKey	float64
	weightKVReadRange	float64
	weightKVDeleteOneKey	float64
	weightKVDeleteRange	float64
	weightKVTxnWriteDelete	float64
	keySize			int
	keyLargeSize		int
	keySuffixRange		int
	keyTxnSuffixRange	int
	keyTxnOps		int
	rateLimiter		*rate.Limiter
	wg			sync.WaitGroup
	clientsN		int
	ctx			context.Context
	cancel			func()
	cli			*clientv3.Client
	emu			sync.RWMutex
	ems			map[string]int
	paused			bool
	atomicModifiedKeys	int64
	stressTable		*stressTable
}

func (s *keyStresser) Stress() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	s.cli, err = s.m.CreateEtcdClient(grpc.WithBackoffMaxDelay(1 * time.Second))
	if err != nil {
		return fmt.Errorf("%v (%q)", err, s.m.EtcdClientEndpoint)
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.wg.Add(s.clientsN)
	s.stressTable = createStressTable([]stressEntry{{weight: s.weightKVWriteSmall, f: newStressPut(s.cli, s.keySuffixRange, s.keySize)}, {weight: s.weightKVWriteLarge, f: newStressPut(s.cli, s.keySuffixRange, s.keyLargeSize)}, {weight: s.weightKVReadOneKey, f: newStressRange(s.cli, s.keySuffixRange)}, {weight: s.weightKVReadRange, f: newStressRangeInterval(s.cli, s.keySuffixRange)}, {weight: s.weightKVDeleteOneKey, f: newStressDelete(s.cli, s.keySuffixRange)}, {weight: s.weightKVDeleteRange, f: newStressDeleteInterval(s.cli, s.keySuffixRange)}, {weight: s.weightKVTxnWriteDelete, f: newStressTxn(s.cli, s.keyTxnSuffixRange, s.keyTxnOps)}})
	s.emu.Lock()
	s.paused = false
	s.ems = make(map[string]int, 100)
	s.emu.Unlock()
	for i := 0; i < s.clientsN; i++ {
		go s.run()
	}
	s.lg.Info("stress START", zap.String("stress-type", "KV"), zap.String("endpoint", s.m.EtcdClientEndpoint))
	return nil
}
func (s *keyStresser) run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer s.wg.Done()
	for {
		if err := s.rateLimiter.Wait(s.ctx); err == context.Canceled {
			return
		}
		sctx, scancel := context.WithTimeout(s.ctx, 10*time.Second)
		err, modifiedKeys := s.stressTable.choose()(sctx)
		scancel()
		if err == nil {
			atomic.AddInt64(&s.atomicModifiedKeys, modifiedKeys)
			continue
		}
		if !s.isRetryableError(err) {
			return
		}
		s.emu.Lock()
		if !s.paused {
			s.ems[err.Error()]++
		}
		s.emu.Unlock()
	}
}
func (s *keyStresser) isRetryableError(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch rpctypes.ErrorDesc(err) {
	case context.DeadlineExceeded.Error():
		return true
	case etcdserver.ErrTimeoutDueToLeaderFail.Error(), etcdserver.ErrTimeout.Error():
		return true
	case etcdserver.ErrStopped.Error():
		return true
	case rpctypes.ErrNotCapable.Error():
		return true
	case rpctypes.ErrTooManyRequests.Error():
		return true
	case raft.ErrProposalDropped.Error():
		return true
	case context.Canceled.Error():
		return false
	case grpc.ErrClientConnClosing.Error():
		return false
	}
	if status.Convert(err).Code() == codes.Unavailable {
		return true
	}
	s.lg.Warn("stress run exiting", zap.String("stress-type", "KV"), zap.String("endpoint", s.m.EtcdClientEndpoint), zap.String("error-type", reflect.TypeOf(err).String()), zap.String("error-desc", rpctypes.ErrorDesc(err)), zap.Error(err))
	return false
}
func (s *keyStresser) Pause() map[string]int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.Close()
}
func (s *keyStresser) Close() map[string]int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.cancel()
	s.cli.Close()
	s.wg.Wait()
	s.emu.Lock()
	s.paused = true
	ess := s.ems
	s.ems = make(map[string]int, 100)
	s.emu.Unlock()
	s.lg.Info("stress STOP", zap.String("stress-type", "KV"), zap.String("endpoint", s.m.EtcdClientEndpoint))
	return ess
}
func (s *keyStresser) ModifiedKeys() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return atomic.LoadInt64(&s.atomicModifiedKeys)
}

type stressFunc func(ctx context.Context) (err error, modifiedKeys int64)
type stressEntry struct {
	weight	float64
	f	stressFunc
}
type stressTable struct {
	entries		[]stressEntry
	sumWeights	float64
}

func createStressTable(entries []stressEntry) *stressTable {
	_logClusterCodePath()
	defer _logClusterCodePath()
	st := stressTable{entries: entries}
	for _, entry := range st.entries {
		st.sumWeights += entry.weight
	}
	return &st
}
func (st *stressTable) choose() stressFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v := rand.Float64() * st.sumWeights
	var sum float64
	var idx int
	for i := range st.entries {
		sum += st.entries[i].weight
		if sum >= v {
			idx = i
			break
		}
	}
	return st.entries[idx].f
}
func newStressPut(cli *clientv3.Client, keySuffixRange, keySize int) stressFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ctx context.Context) (error, int64) {
		_, err := cli.Put(ctx, fmt.Sprintf("foo%016x", rand.Intn(keySuffixRange)), string(randBytes(keySize)))
		return err, 1
	}
}
func newStressTxn(cli *clientv3.Client, keyTxnSuffixRange, txnOps int) stressFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keys := make([]string, keyTxnSuffixRange)
	for i := range keys {
		keys[i] = fmt.Sprintf("/k%03d", i)
	}
	return writeTxn(cli, keys, txnOps)
}
func writeTxn(cli *clientv3.Client, keys []string, txnOps int) stressFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ctx context.Context) (error, int64) {
		ks := make(map[string]struct{}, txnOps)
		for len(ks) != txnOps {
			ks[keys[rand.Intn(len(keys))]] = struct{}{}
		}
		selected := make([]string, 0, txnOps)
		for k := range ks {
			selected = append(selected, k)
		}
		com, delOp, putOp := getTxnOps(selected[0], "bar00")
		thenOps := []clientv3.Op{delOp}
		elseOps := []clientv3.Op{putOp}
		for i := 1; i < txnOps; i++ {
			k, v := selected[i], fmt.Sprintf("bar%02d", i)
			com, delOp, putOp = getTxnOps(k, v)
			txnOp := clientv3.OpTxn([]clientv3.Cmp{com}, []clientv3.Op{delOp}, []clientv3.Op{putOp})
			thenOps = append(thenOps, txnOp)
			elseOps = append(elseOps, txnOp)
		}
		_, err := cli.Txn(ctx).If(com).Then(thenOps...).Else(elseOps...).Commit()
		return err, int64(txnOps)
	}
}
func getTxnOps(k, v string) (cmp clientv3.Cmp, dop clientv3.Op, pop clientv3.Op) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmp = clientv3.Compare(clientv3.Version(k), ">", 0)
	dop = clientv3.OpDelete(k)
	pop = clientv3.OpPut(k, v)
	return cmp, dop, pop
}
func newStressRange(cli *clientv3.Client, keySuffixRange int) stressFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ctx context.Context) (error, int64) {
		_, err := cli.Get(ctx, fmt.Sprintf("foo%016x", rand.Intn(keySuffixRange)))
		return err, 0
	}
}
func newStressRangeInterval(cli *clientv3.Client, keySuffixRange int) stressFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ctx context.Context) (error, int64) {
		start := rand.Intn(keySuffixRange)
		end := start + 500
		_, err := cli.Get(ctx, fmt.Sprintf("foo%016x", start), clientv3.WithRange(fmt.Sprintf("foo%016x", end)))
		return err, 0
	}
}
func newStressDelete(cli *clientv3.Client, keySuffixRange int) stressFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ctx context.Context) (error, int64) {
		_, err := cli.Delete(ctx, fmt.Sprintf("foo%016x", rand.Intn(keySuffixRange)))
		return err, 1
	}
}
func newStressDeleteInterval(cli *clientv3.Client, keySuffixRange int) stressFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ctx context.Context) (error, int64) {
		start := rand.Intn(keySuffixRange)
		end := start + 500
		resp, err := cli.Delete(ctx, fmt.Sprintf("foo%016x", start), clientv3.WithRange(fmt.Sprintf("foo%016x", end)))
		if err == nil {
			return nil, resp.Deleted
		}
		return err, 0
	}
}
