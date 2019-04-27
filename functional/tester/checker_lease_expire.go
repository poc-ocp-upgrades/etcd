package tester

import (
	"context"
	"fmt"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type leaseExpireChecker struct {
	ctype	rpcpb.Checker
	lg	*zap.Logger
	m	*rpcpb.Member
	ls	*leaseStresser
	cli	*clientv3.Client
}

func newLeaseExpireChecker(ls *leaseStresser) Checker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &leaseExpireChecker{ctype: rpcpb.Checker_LEASE_EXPIRE, lg: ls.lg, m: ls.m, ls: ls}
}
func (lc *leaseExpireChecker) Type() rpcpb.Checker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return lc.ctype
}
func (lc *leaseExpireChecker) EtcdClientEndpoints() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{lc.m.EtcdClientEndpoint}
}
func (lc *leaseExpireChecker) Check() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if lc.ls == nil {
		return nil
	}
	if lc.ls != nil && (lc.ls.revokedLeases == nil || lc.ls.aliveLeases == nil || lc.ls.shortLivedLeases == nil) {
		return nil
	}
	cli, err := lc.m.CreateEtcdClient(grpc.WithBackoffMaxDelay(time.Second))
	if err != nil {
		return fmt.Errorf("%v (%q)", err, lc.m.EtcdClientEndpoint)
	}
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	lc.cli = cli
	if err := lc.check(true, lc.ls.revokedLeases.leases); err != nil {
		return err
	}
	if err := lc.check(false, lc.ls.aliveLeases.leases); err != nil {
		return err
	}
	return lc.checkShortLivedLeases()
}

const leaseExpireCheckerTimeout = 10 * time.Second

func (lc *leaseExpireChecker) checkShortLivedLeases() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithTimeout(context.Background(), leaseExpireCheckerTimeout)
	errc := make(chan error)
	defer cancel()
	for leaseID := range lc.ls.shortLivedLeases.leases {
		go func(id int64) {
			errc <- lc.checkShortLivedLease(ctx, id)
		}(leaseID)
	}
	var errs []error
	for range lc.ls.shortLivedLeases.leases {
		if err := <-errc; err != nil {
			errs = append(errs, err)
		}
	}
	return errsToError(errs)
}
func (lc *leaseExpireChecker) checkShortLivedLease(ctx context.Context, leaseID int64) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var resp *clientv3.LeaseTimeToLiveResponse
	for i := 0; i < retries; i++ {
		resp, err = lc.getLeaseByID(ctx, leaseID)
		if (err == nil && resp.TTL == -1) || (err != nil && rpctypes.Error(err) == rpctypes.ErrLeaseNotFound) {
			return nil
		}
		if err != nil {
			lc.lg.Debug("retrying; Lease TimeToLive failed", zap.Int("retries", i), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Error(err))
			continue
		}
		if resp.TTL > 0 {
			dur := time.Duration(resp.TTL) * time.Second
			lc.lg.Debug("lease has not been expired, wait until expire", zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Int64("ttl", resp.TTL), zap.Duration("wait-duration", dur))
			time.Sleep(dur)
		} else {
			lc.lg.Debug("lease expired but not yet revoked", zap.Int("retries", i), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Int64("ttl", resp.TTL), zap.Duration("wait-duration", time.Second))
			time.Sleep(time.Second)
		}
		if err = lc.checkLease(ctx, false, leaseID); err != nil {
			continue
		}
		return nil
	}
	return err
}
func (lc *leaseExpireChecker) checkLease(ctx context.Context, expired bool, leaseID int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keysExpired, err := lc.hasKeysAttachedToLeaseExpired(ctx, leaseID)
	if err != nil {
		lc.lg.Warn("hasKeysAttachedToLeaseExpired failed", zap.String("endpoint", lc.m.EtcdClientEndpoint), zap.Error(err))
		return err
	}
	leaseExpired, err := lc.hasLeaseExpired(ctx, leaseID)
	if err != nil {
		lc.lg.Warn("hasLeaseExpired failed", zap.String("endpoint", lc.m.EtcdClientEndpoint), zap.Error(err))
		return err
	}
	if leaseExpired != keysExpired {
		return fmt.Errorf("lease %v expiration mismatch (lease expired=%v, keys expired=%v)", leaseID, leaseExpired, keysExpired)
	}
	if leaseExpired != expired {
		return fmt.Errorf("lease %v expected expired=%v, got %v", leaseID, expired, leaseExpired)
	}
	return nil
}
func (lc *leaseExpireChecker) check(expired bool, leases map[int64]time.Time) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithTimeout(context.Background(), leaseExpireCheckerTimeout)
	defer cancel()
	for leaseID := range leases {
		if err := lc.checkLease(ctx, expired, leaseID); err != nil {
			return err
		}
	}
	return nil
}
func (lc *leaseExpireChecker) getLeaseByID(ctx context.Context, leaseID int64) (*clientv3.LeaseTimeToLiveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return lc.cli.TimeToLive(ctx, clientv3.LeaseID(leaseID), clientv3.WithAttachedKeys())
}
func (lc *leaseExpireChecker) hasLeaseExpired(ctx context.Context, leaseID int64) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for ctx.Err() == nil {
		resp, err := lc.getLeaseByID(ctx, leaseID)
		if err != nil {
			if rpctypes.Error(err) == rpctypes.ErrLeaseNotFound {
				return true, nil
			}
		} else {
			return resp.TTL == -1, nil
		}
		lc.lg.Warn("hasLeaseExpired getLeaseByID failed", zap.String("endpoint", lc.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Error(err))
	}
	return false, ctx.Err()
}
func (lc *leaseExpireChecker) hasKeysAttachedToLeaseExpired(ctx context.Context, leaseID int64) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := lc.cli.Get(ctx, fmt.Sprintf("%d", leaseID), clientv3.WithPrefix())
	if err != nil {
		lc.lg.Warn("hasKeysAttachedToLeaseExpired failed", zap.String("endpoint", lc.m.EtcdClientEndpoint), zap.String("lease-id", fmt.Sprintf("%016x", leaseID)), zap.Error(err))
		return false, err
	}
	return len(resp.Kvs) == 0, nil
}
