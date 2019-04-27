package recipe

import (
	"context"
	"fmt"
	v3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type PriorityQueue struct {
	client	*v3.Client
	ctx	context.Context
	key	string
}

func NewPriorityQueue(client *v3.Client, key string) *PriorityQueue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PriorityQueue{client, context.TODO(), key + "/"}
}
func (q *PriorityQueue) Enqueue(val string, pr uint16) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := fmt.Sprintf("%s%05d", q.key, pr)
	_, err := newSequentialKV(q.client, prefix, val)
	return err
}
func (q *PriorityQueue) Dequeue() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := q.client.Get(q.ctx, q.key, v3.WithFirstKey()...)
	if err != nil {
		return "", err
	}
	kv, err := claimFirstKey(q.client, resp.Kvs)
	if err != nil {
		return "", err
	} else if kv != nil {
		return string(kv.Value), nil
	} else if resp.More {
		return q.Dequeue()
	}
	ev, err := WaitPrefixEvents(q.client, q.key, resp.Header.Revision, []mvccpb.Event_EventType{mvccpb.PUT})
	if err != nil {
		return "", err
	}
	ok, err := deleteRevKey(q.client, string(ev.Kv.Key), ev.Kv.ModRevision)
	if err != nil {
		return "", err
	} else if !ok {
		return q.Dequeue()
	}
	return string(ev.Kv.Value), err
}
