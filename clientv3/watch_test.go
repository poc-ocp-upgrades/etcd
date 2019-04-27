package clientv3

import (
	"testing"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func TestEvent(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		ev		*Event
		isCreate	bool
		isModify	bool
	}{{ev: &Event{Type: EventTypePut, Kv: &mvccpb.KeyValue{CreateRevision: 3, ModRevision: 3}}, isCreate: true}, {ev: &Event{Type: EventTypePut, Kv: &mvccpb.KeyValue{CreateRevision: 3, ModRevision: 4}}, isModify: true}}
	for i, tt := range tests {
		if tt.isCreate && !tt.ev.IsCreate() {
			t.Errorf("#%d: event should be Create event", i)
		}
		if tt.isModify && !tt.ev.IsModify() {
			t.Errorf("#%d: event should be Modify event", i)
		}
	}
}
