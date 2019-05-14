package recipe

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type DoubleBarrier struct {
	s     *concurrency.Session
	ctx   context.Context
	key   string
	count int
	myKey *EphemeralKV
}

func NewDoubleBarrier(s *concurrency.Session, key string, count int) *DoubleBarrier {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DoubleBarrier{s: s, ctx: context.TODO(), key: key, count: count}
}
func (b *DoubleBarrier) Enter() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := b.s.Client()
	ek, err := newUniqueEphemeralKey(b.s, b.key+"/waiters")
	if err != nil {
		return err
	}
	b.myKey = ek
	resp, err := client.Get(b.ctx, b.key+"/waiters", clientv3.WithPrefix())
	if err != nil {
		return err
	}
	if len(resp.Kvs) > b.count {
		return ErrTooManyClients
	}
	if len(resp.Kvs) == b.count {
		_, err = client.Put(b.ctx, b.key+"/ready", "")
		return err
	}
	_, err = WaitEvents(client, b.key+"/ready", ek.Revision(), []mvccpb.Event_EventType{mvccpb.PUT})
	return err
}
func (b *DoubleBarrier) Leave() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := b.s.Client()
	resp, err := client.Get(b.ctx, b.key+"/waiters", clientv3.WithPrefix())
	if err != nil {
		return err
	}
	if len(resp.Kvs) == 0 {
		return nil
	}
	lowest, highest := resp.Kvs[0], resp.Kvs[0]
	for _, k := range resp.Kvs {
		if k.ModRevision < lowest.ModRevision {
			lowest = k
		}
		if k.ModRevision > highest.ModRevision {
			highest = k
		}
	}
	isLowest := string(lowest.Key) == b.myKey.Key()
	if len(resp.Kvs) == 1 {
		if _, err = client.Delete(b.ctx, b.key+"/ready"); err != nil {
			return err
		}
		return b.myKey.Delete()
	}
	if isLowest {
		_, err = WaitEvents(client, string(highest.Key), highest.ModRevision, []mvccpb.Event_EventType{mvccpb.DELETE})
		if err != nil {
			return err
		}
		return b.Leave()
	}
	if err = b.myKey.Delete(); err != nil {
		return err
	}
	key := string(lowest.Key)
	_, err = WaitEvents(client, key, lowest.ModRevision, []mvccpb.Event_EventType{mvccpb.DELETE})
	if err != nil {
		return err
	}
	return b.Leave()
}
