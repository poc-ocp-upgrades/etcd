package tester

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/transport"
)

type keyStresser struct {
	stype			rpcpb.Stresser
	lg			*zap.Logger
	m			*rpcpb.Member
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
	var stressEntries = []stressEntry{{weight: 0.7, f: newStressPut(s.cli, s.keySuffixRange, s.keySize)}, {weight: 0.7 * float32(s.keySize) / float32(s.keyLargeSize), f: newStressPut(s.cli, s.keySuffixRange, s.keyLargeSize)}, {weight: 0.07, f: newStressRange(s.cli, s.keySuffixRange)}, {weight: 0.07, f: newStressRangeInterval(s.cli, s.keySuffixRange)}, {weight: 0.07, f: newStressDelete(s.cli, s.keySuffixRange)}, {weight: 0.07, f: newStressDeleteInterval(s.cli, s.keySuffixRange)}}
	if s.keyTxnSuffixRange > 0 {
		stressEntries[0].weight = 0.35
		stressEntries = append(stressEntries, stressEntry{weight: 0.35, f: newStressTxn(s.cli, s.keyTxnSuffixRange, s.keyTxnOps)})
	}
	s.stressTable = createStressTable(stressEntries)
	s.emu.Lock()
	s.paused = false
	s.ems = make(map[string]int, 100)
	s.emu.Unlock()
	for i := 0; i < s.clientsN; i++ {
		go s.run()
	}
	s.lg.Info("stress START", zap.String("stress-type", s.stype.String()), zap.String("endpoint", s.m.EtcdClientEndpoint))
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
		switch rpctypes.ErrorDesc(err) {
		case context.DeadlineExceeded.Error():
		case etcdserver.ErrTimeoutDueToLeaderFail.Error(), etcdserver.ErrTimeout.Error():
		case etcdserver.ErrStopped.Error():
		case transport.ErrConnClosing.Desc:
		case rpctypes.ErrNotCapable.Error():
		case rpctypes.ErrTooManyRequests.Error():
		case context.Canceled.Error():
			return
		case grpc.ErrClientConnClosing.Error():
			return
		default:
			s.lg.Warn("stress run exiting", zap.String("stress-type", s.stype.String()), zap.String("endpoint", s.m.EtcdClientEndpoint), zap.String("error-type", reflect.TypeOf(err).String()), zap.Error(err))
			return
		}
		s.emu.Lock()
		if !s.paused {
			s.ems[err.Error()]++
		}
		s.emu.Unlock()
	}
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
	s.lg.Info("stress STOP", zap.String("stress-type", s.stype.String()), zap.String("endpoint", s.m.EtcdClientEndpoint))
	return ess
}
func (s *keyStresser) ModifiedKeys() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return atomic.LoadInt64(&s.atomicModifiedKeys)
}

type stressFunc func(ctx context.Context) (err error, modifiedKeys int64)
type stressEntry struct {
	weight	float32
	f	stressFunc
}
type stressTable struct {
	entries		[]stressEntry
	sumWeights	float32
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
	v := rand.Float32() * st.sumWeights
	var sum float32
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
