package command

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/mirror"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/spf13/cobra"
)

var (
	mminsecureTr	bool
	mmcert		string
	mmkey		string
	mmcacert	string
	mmprefix	string
	mmdestprefix	string
	mmnodestprefix	bool
)

func NewMakeMirrorCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &cobra.Command{Use: "make-mirror [options] <destination>", Short: "Makes a mirror at the destination etcd cluster", Run: makeMirrorCommandFunc}
	c.Flags().StringVar(&mmprefix, "prefix", "", "Key-value prefix to mirror")
	c.Flags().StringVar(&mmdestprefix, "dest-prefix", "", "destination prefix to mirror a prefix to a different prefix in the destination cluster")
	c.Flags().BoolVar(&mmnodestprefix, "no-dest-prefix", false, "mirror key-values to the root of the destination cluster")
	c.Flags().StringVar(&mmcert, "dest-cert", "", "Identify secure client using this TLS certificate file for the destination cluster")
	c.Flags().StringVar(&mmkey, "dest-key", "", "Identify secure client using this TLS key file")
	c.Flags().StringVar(&mmcacert, "dest-cacert", "", "Verify certificates of TLS enabled secure servers using this CA bundle")
	c.Flags().BoolVar(&mminsecureTr, "dest-insecure-transport", true, "Disable transport security for client connections")
	return c
}
func makeMirrorCommandFunc(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) != 1 {
		ExitWithError(ExitBadArgs, errors.New("make-mirror takes one destination argument."))
	}
	dialTimeout := dialTimeoutFromCmd(cmd)
	keepAliveTime := keepAliveTimeFromCmd(cmd)
	keepAliveTimeout := keepAliveTimeoutFromCmd(cmd)
	sec := &secureCfg{cert: mmcert, key: mmkey, cacert: mmcacert, insecureTransport: mminsecureTr}
	cc := &clientConfig{endpoints: []string{args[0]}, dialTimeout: dialTimeout, keepAliveTime: keepAliveTime, keepAliveTimeout: keepAliveTimeout, scfg: sec, acfg: nil}
	dc := cc.mustClient()
	c := mustClientFromCmd(cmd)
	err := makeMirror(context.TODO(), c, dc)
	ExitWithError(ExitError, err)
}
func makeMirror(ctx context.Context, c *clientv3.Client, dc *clientv3.Client) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	total := int64(0)
	go func() {
		for {
			time.Sleep(30 * time.Second)
			fmt.Println(atomic.LoadInt64(&total))
		}
	}()
	s := mirror.NewSyncer(c, mmprefix, 0)
	rc, errc := s.SyncBase(ctx)
	if mmnodestprefix && len(mmdestprefix) > 0 {
		ExitWithError(ExitBadArgs, fmt.Errorf("`--dest-prefix` and `--no-dest-prefix` cannot be set at the same time, choose one."))
	}
	if !mmnodestprefix && len(mmdestprefix) == 0 {
		mmdestprefix = mmprefix
	}
	for r := range rc {
		for _, kv := range r.Kvs {
			_, err := dc.Put(ctx, modifyPrefix(string(kv.Key)), string(kv.Value))
			if err != nil {
				return err
			}
			atomic.AddInt64(&total, 1)
		}
	}
	err := <-errc
	if err != nil {
		return err
	}
	wc := s.SyncUpdates(ctx)
	for wr := range wc {
		if wr.CompactRevision != 0 {
			return rpctypes.ErrCompacted
		}
		var lastRev int64
		ops := []clientv3.Op{}
		for _, ev := range wr.Events {
			nextRev := ev.Kv.ModRevision
			if lastRev != 0 && nextRev > lastRev {
				_, err := dc.Txn(ctx).Then(ops...).Commit()
				if err != nil {
					return err
				}
				ops = []clientv3.Op{}
			}
			lastRev = nextRev
			switch ev.Type {
			case mvccpb.PUT:
				ops = append(ops, clientv3.OpPut(modifyPrefix(string(ev.Kv.Key)), string(ev.Kv.Value)))
				atomic.AddInt64(&total, 1)
			case mvccpb.DELETE:
				ops = append(ops, clientv3.OpDelete(modifyPrefix(string(ev.Kv.Key))))
				atomic.AddInt64(&total, 1)
			default:
				panic("unexpected event type")
			}
		}
		if len(ops) != 0 {
			_, err := dc.Txn(ctx).Then(ops...).Commit()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func modifyPrefix(key string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Replace(key, mmprefix, mmdestprefix, 1)
}
