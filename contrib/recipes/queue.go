package recipe

import (
	"context"
	v3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type Queue struct {
	client		*v3.Client
	ctx		context.Context
	keyPrefix	string
}

func NewQueue(client *v3.Client, keyPrefix string) *Queue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Queue{client, context.TODO(), keyPrefix}
}
func (q *Queue) Enqueue(val string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := newUniqueKV(q.client, q.keyPrefix, val)
	return err
}
func (q *Queue) Dequeue() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := q.client.Get(q.ctx, q.keyPrefix, v3.WithFirstRev()...)
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
	ev, err := WaitPrefixEvents(q.client, q.keyPrefix, resp.Header.Revision, []mvccpb.Event_EventType{mvccpb.PUT})
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
