package tester

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
)

const (
	defaultTTL	= 120
	defaultTTLShort	= 2
)

type leaseStresser struct {
	stype			rpcpb.Stresser
	lg			*zap.Logger
	m			*rpcpb.Member
	cli			*clientv3.Client
	ctx			context.Context
	cancel			func()
	rateLimiter		*rate.Limiter
	atomicModifiedKey	int64
	numLeases		int
	keysPerLease		int
	aliveLeases		*atomicLeases
	revokedLeases		*atomicLeases
	shortLivedLeases	*atomicLeases
	runWg			sync.WaitGroup
	aliveWg			sync.WaitGroup
}
type atomicLeases struct {
	rwLock	sync.RWMutex
	leases	map[int64]time.Time
}

func (al *atomicLeases) add(leaseID int64, t time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	al.rwLock.Lock()
	al.leases[leaseID] = t
	al.rwLock.Unlock()
}
func (al *atomicLeases) update(leaseID int64, t time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	al.rwLock.Lock()
	_, ok := al.leases[leaseID]
	if ok {
		al.leases[leaseID] = t
	}
	al.rwLock.Unlock()
}
func (al *atomicLeases) read(leaseID int64) (rv time.Time, ok bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	al.rwLock.RLock()
	rv, ok = al.leases[leaseID]
	al.rwLock.RUnlock()
	return rv, ok
}
func (al *atomicLeases) remove(leaseID int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	al.rwLock.Lock()
	delete(al.leases, leaseID)
	al.rwLock.Unlock()
}
func (al *atomicLeases) getLeasesMap() map[int64]time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	leasesCopy := make(map[int64]time.Time)
	al.rwLock.RLock()
	for k, v := range al.leases {
		leasesCopy[k] = v
	}
	al.rwLock.RUnlock()
	return leasesCopy
}
func (ls *leaseStresser) setupOnce() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ls.aliveLeases != nil {
		return nil
	}
	if ls.numLeases == 0 {
		panic("expect numLeases to be set")
	}
	if ls.keysPerLease == 0 {
		panic("expect keysPerLease to be set")
	}
	ls.aliveLeases = &atomicLeases{leases: make(map[int64]time.Time)}
	return nil
}
func (ls *leaseStresser) Stress() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ls.lg.Info("stress START", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint))
	if err := ls.setupOnce(); err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	ls.ctx = ctx
	ls.cancel = cancel
	cli, err := ls.m.CreateEtcdClient(grpc.WithBackoffMaxDelay(1 * time.Second))
	if err != nil {
		return fmt.Errorf("%v (%s)", err, ls.m.EtcdClientEndpoint)
	}
	ls.cli = cli
	ls.revokedLeases = &atomicLeases{leases: make(map[int64]time.Time)}
	ls.shortLivedLeases = &atomicLeases{leases: make(map[int64]time.Time)}
	ls.runWg.Add(1)
	go ls.run()
	return nil
}
func (ls *leaseStresser) run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer ls.runWg.Done()
	ls.restartKeepAlives()
	for {
		err := ls.rateLimiter.WaitN(ls.ctx, 2*ls.numLeases*ls.keysPerLease)
		if err == context.Canceled {
			return
		}
		ls.lg.Debug("stress creating leases", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint))
		ls.createLeases()
		ls.lg.Debug("stress created leases", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint))
		ls.lg.Debug("stress dropped leases", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint))
		ls.randomlyDropLeases()
		ls.lg.Debug("stress dropped leases", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint))
	}
}
func (ls *leaseStresser) restartKeepAlives() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for leaseID := range ls.aliveLeases.getLeasesMap() {
		ls.aliveWg.Add(1)
		go func(id int64) {
			ls.keepLeaseAlive(id)
		}(leaseID)
	}
}
func (ls *leaseStresser) createLeases() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ls.createAliveLeases()
	ls.createShortLivedLeases()
}
func (ls *leaseStresser) createAliveLeases() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	neededLeases := ls.numLeases - len(ls.aliveLeases.getLeasesMap())
	var wg sync.WaitGroup
	for i := 0; i < neededLeases; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			leaseID, err := ls.createLeaseWithKeys(defaultTTL)
			if err != nil {
				ls.lg.Debug("createLeaseWithKeys failed", zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.Error(err))
				return
			}
			ls.aliveLeases.add(leaseID, time.Now())
			ls.aliveWg.Add(1)
			go ls.keepLeaseAlive(leaseID)
		}()
	}
	wg.Wait()
}
func (ls *leaseStresser) createShortLivedLeases() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	neededLeases := ls.numLeases - len(ls.shortLivedLeases.getLeasesMap())
	var wg sync.WaitGroup
	for i := 0; i < neededLeases; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			leaseID, err := ls.createLeaseWithKeys(defaultTTLShort)
			if err != nil {
				return
			}
			ls.shortLivedLeases.add(leaseID, time.Now())
		}()
	}
	wg.Wait()
}
func (ls *leaseStresser) createLeaseWithKeys(ttl int64) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	leaseID, err := ls.createLease(ttl)
	if err != nil {
		ls.lg.Debug("createLease failed", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.Error(err))
		return -1, err
	}
	ls.lg.Debug("createLease created lease", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)))
	if err := ls.attachKeysWithLease(leaseID); err != nil {
		return -1, err
	}
	return leaseID, nil
}
func (ls *leaseStresser) randomlyDropLeases() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wg sync.WaitGroup
	for l := range ls.aliveLeases.getLeasesMap() {
		wg.Add(1)
		go func(leaseID int64) {
			defer wg.Done()
			dropped, err := ls.randomlyDropLease(leaseID)
			if err != nil {
				ls.lg.Debug("randomlyDropLease failed", zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Error(err))
				ls.aliveLeases.remove(leaseID)
				return
			}
			if !dropped {
				return
			}
			ls.lg.Debug("randomlyDropLease dropped a lease", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)))
			ls.revokedLeases.add(leaseID, time.Now())
			ls.aliveLeases.remove(leaseID)
		}(l)
	}
	wg.Wait()
}
func (ls *leaseStresser) createLease(ttl int64) (int64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := ls.cli.Grant(ls.ctx, ttl)
	if err != nil {
		return -1, err
	}
	return int64(resp.ID), nil
}
func (ls *leaseStresser) keepLeaseAlive(leaseID int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer ls.aliveWg.Done()
	ctx, cancel := context.WithCancel(ls.ctx)
	stream, err := ls.cli.KeepAlive(ctx, clientv3.LeaseID(leaseID))
	defer func() {
		cancel()
	}()
	for {
		select {
		case <-time.After(500 * time.Millisecond):
		case <-ls.ctx.Done():
			ls.lg.Debug("keepLeaseAlive context canceled", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Error(ls.ctx.Err()))
			renewTime, ok := ls.aliveLeases.read(leaseID)
			if ok && renewTime.Add(defaultTTL/2*time.Second).Before(time.Now()) {
				ls.aliveLeases.remove(leaseID)
				ls.lg.Debug("keepLeaseAlive lease has not been renewed, dropped it", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)))
			}
			return
		}
		if err != nil {
			ls.lg.Debug("keepLeaseAlive lease creates stream error", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Error(err))
			cancel()
			ctx, cancel = context.WithCancel(ls.ctx)
			stream, err = ls.cli.KeepAlive(ctx, clientv3.LeaseID(leaseID))
			cancel()
			continue
		}
		if err != nil {
			ls.lg.Debug("keepLeaseAlive failed to receive lease keepalive response", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Error(err))
			continue
		}
		ls.lg.Debug("keepLeaseAlive waiting on lease stream", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)))
		leaseRenewTime := time.Now()
		respRC := <-stream
		if respRC == nil {
			ls.lg.Debug("keepLeaseAlive received nil lease keepalive response", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)))
			continue
		}
		if respRC.TTL <= 0 {
			ls.lg.Debug("keepLeaseAlive stream received lease keepalive response TTL <= 0", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Int64("ttl", respRC.TTL))
			ls.aliveLeases.remove(leaseID)
			return
		}
		ls.lg.Debug("keepLeaseAlive renewed a lease", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)))
		ls.aliveLeases.update(leaseID, leaseRenewTime)
	}
}
func (ls *leaseStresser) attachKeysWithLease(leaseID int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var txnPuts []clientv3.Op
	for j := 0; j < ls.keysPerLease; j++ {
		txnput := clientv3.OpPut(fmt.Sprintf("%d%s%d", leaseID, "_", j), fmt.Sprintf("bar"), clientv3.WithLease(clientv3.LeaseID(leaseID)))
		txnPuts = append(txnPuts, txnput)
	}
	for ls.ctx.Err() == nil {
		_, err := ls.cli.Txn(ls.ctx).Then(txnPuts...).Commit()
		if err == nil {
			atomic.AddInt64(&ls.atomicModifiedKey, 2*int64(ls.keysPerLease))
			return nil
		}
		if rpctypes.Error(err) == rpctypes.ErrLeaseNotFound {
			return err
		}
	}
	return ls.ctx.Err()
}
func (ls *leaseStresser) randomlyDropLease(leaseID int64) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rand.Intn(2) != 0 {
		return false, nil
	}
	for ls.ctx.Err() == nil {
		_, err := ls.cli.Revoke(ls.ctx, clientv3.LeaseID(leaseID))
		if err == nil || rpctypes.Error(err) == rpctypes.ErrLeaseNotFound {
			return true, nil
		}
	}
	ls.lg.Debug("randomlyDropLease error", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Error(ls.ctx.Err()))
	return false, ls.ctx.Err()
}
func (ls *leaseStresser) Pause() map[string]int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ls.Close()
}
func (ls *leaseStresser) Close() map[string]int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ls.cancel()
	ls.runWg.Wait()
	ls.aliveWg.Wait()
	ls.cli.Close()
	ls.lg.Info("stress STOP", zap.String("stress-type", ls.stype.String()), zap.String("endpoint", ls.m.EtcdClientEndpoint))
	return nil
}
func (ls *leaseStresser) ModifiedKeys() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return atomic.LoadInt64(&ls.atomicModifiedKey)
}
