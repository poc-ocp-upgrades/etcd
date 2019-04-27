package mockstore

import (
	"time"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/store"
)

type StoreRecorder struct {
	store.Store
	testutil.Recorder
}
type storeRecorder struct {
	store.Store
	testutil.Recorder
}

func NewNop() store.Store {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &storeRecorder{Recorder: &testutil.RecorderBuffered{}}
}
func NewRecorder() *StoreRecorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sr := &storeRecorder{Recorder: &testutil.RecorderBuffered{}}
	return &StoreRecorder{Store: sr, Recorder: sr.Recorder}
}
func NewRecorderStream() *StoreRecorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sr := &storeRecorder{Recorder: testutil.NewRecorderStream()}
	return &StoreRecorder{Store: sr, Recorder: sr.Recorder}
}
func (s *storeRecorder) Version() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (s *storeRecorder) Index() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 0
}
func (s *storeRecorder) Get(path string, recursive, sorted bool) (*store.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "Get", Params: []interface{}{path, recursive, sorted}})
	return &store.Event{}, nil
}
func (s *storeRecorder) Set(path string, dir bool, val string, expireOpts store.TTLOptionSet) (*store.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "Set", Params: []interface{}{path, dir, val, expireOpts}})
	return &store.Event{}, nil
}
func (s *storeRecorder) Update(path, val string, expireOpts store.TTLOptionSet) (*store.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "Update", Params: []interface{}{path, val, expireOpts}})
	return &store.Event{}, nil
}
func (s *storeRecorder) Create(path string, dir bool, val string, uniq bool, expireOpts store.TTLOptionSet) (*store.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "Create", Params: []interface{}{path, dir, val, uniq, expireOpts}})
	return &store.Event{}, nil
}
func (s *storeRecorder) CompareAndSwap(path, prevVal string, prevIdx uint64, val string, expireOpts store.TTLOptionSet) (*store.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "CompareAndSwap", Params: []interface{}{path, prevVal, prevIdx, val, expireOpts}})
	return &store.Event{}, nil
}
func (s *storeRecorder) Delete(path string, dir, recursive bool) (*store.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "Delete", Params: []interface{}{path, dir, recursive}})
	return &store.Event{}, nil
}
func (s *storeRecorder) CompareAndDelete(path, prevVal string, prevIdx uint64) (*store.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "CompareAndDelete", Params: []interface{}{path, prevVal, prevIdx}})
	return &store.Event{}, nil
}
func (s *storeRecorder) Watch(_ string, _, _ bool, _ uint64) (store.Watcher, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "Watch"})
	return store.NewNopWatcher(), nil
}
func (s *storeRecorder) Save() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "Save"})
	return nil, nil
}
func (s *storeRecorder) Recovery(b []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "Recovery"})
	return nil
}
func (s *storeRecorder) SaveNoCopy() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "SaveNoCopy"})
	return nil, nil
}
func (s *storeRecorder) Clone() store.Store {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "Clone"})
	return s
}
func (s *storeRecorder) JsonStats() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *storeRecorder) DeleteExpiredKeys(cutoff time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "DeleteExpiredKeys", Params: []interface{}{cutoff}})
}
func (s *storeRecorder) HasTTLKeys() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Record(testutil.Action{Name: "HasTTLKeys"})
	return true
}

type errStoreRecorder struct {
	storeRecorder
	err	error
}

func NewErrRecorder(err error) *StoreRecorder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sr := &errStoreRecorder{err: err}
	sr.Recorder = &testutil.RecorderBuffered{}
	return &StoreRecorder{Store: sr, Recorder: sr.Recorder}
}
func (s *errStoreRecorder) Get(path string, recursive, sorted bool) (*store.Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.storeRecorder.Get(path, recursive, sorted)
	return nil, s.err
}
func (s *errStoreRecorder) Watch(path string, recursive, sorted bool, index uint64) (store.Watcher, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.storeRecorder.Watch(path, recursive, sorted, index)
	return nil, s.err
}
