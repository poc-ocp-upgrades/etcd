package rafthttp

import (
	"bytes"
	"reflect"
	"testing"
	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
)

func TestMsgAppV2(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []raftpb.Message{linkHeartbeatMessage, {Type: raftpb.MsgApp, From: 1, To: 2, Term: 1, LogTerm: 1, Index: 0, Entries: []raftpb.Entry{{Term: 1, Index: 1, Data: []byte("some data")}, {Term: 1, Index: 2, Data: []byte("some data")}, {Term: 1, Index: 3, Data: []byte("some data")}}}, {Type: raftpb.MsgApp, From: 1, To: 2, Term: 1, LogTerm: 1, Index: 3, Entries: []raftpb.Entry{{Term: 1, Index: 4, Data: []byte("some data")}}}, linkHeartbeatMessage, {Type: raftpb.MsgApp, From: 1, To: 2, Term: 1, LogTerm: 1, Index: 4, Entries: []raftpb.Entry{{Term: 1, Index: 5, Data: []byte("some data")}}}, {Type: raftpb.MsgApp, From: 1, To: 2, Term: 3, LogTerm: 1, Index: 5, Entries: []raftpb.Entry{{Term: 3, Index: 6, Data: []byte("some data")}}}, linkHeartbeatMessage, {Type: raftpb.MsgApp, From: 1, To: 2, Term: 3, LogTerm: 2, Index: 6, Entries: []raftpb.Entry{{Term: 3, Index: 7, Data: []byte("some data")}}}, {Type: raftpb.MsgApp, From: 1, To: 2, Term: 3, LogTerm: 2, Index: 7, Entries: nil}, linkHeartbeatMessage}
	b := &bytes.Buffer{}
	enc := newMsgAppV2Encoder(b, &stats.FollowerStats{})
	dec := newMsgAppV2Decoder(b, types.ID(2), types.ID(1))
	for i, tt := range tests {
		if err := enc.encode(&tt); err != nil {
			t.Errorf("#%d: unexpected encode message error: %v", i, err)
			continue
		}
		m, err := dec.decode()
		if err != nil {
			t.Errorf("#%d: unexpected decode message error: %v", i, err)
			continue
		}
		if !reflect.DeepEqual(m, tt) {
			t.Errorf("#%d: message = %+v, want %+v", i, m, tt)
		}
	}
}
