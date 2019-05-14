package store

import (
	"encoding/json"
	"fmt"
	etcdErr "github.com/coreos/etcd/error"
	"github.com/coreos/etcd/pkg/types"
	"github.com/jonboulle/clockwork"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

const defaultVersion = 2

var minExpireTime time.Time

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	minExpireTime, _ = time.Parse(time.RFC3339, "2000-01-01T00:00:00Z")
}

type Store interface {
	Version() int
	Index() uint64
	Get(nodePath string, recursive, sorted bool) (*Event, error)
	Set(nodePath string, dir bool, value string, expireOpts TTLOptionSet) (*Event, error)
	Update(nodePath string, newValue string, expireOpts TTLOptionSet) (*Event, error)
	Create(nodePath string, dir bool, value string, unique bool, expireOpts TTLOptionSet) (*Event, error)
	CompareAndSwap(nodePath string, prevValue string, prevIndex uint64, value string, expireOpts TTLOptionSet) (*Event, error)
	Delete(nodePath string, dir, recursive bool) (*Event, error)
	CompareAndDelete(nodePath string, prevValue string, prevIndex uint64) (*Event, error)
	Watch(prefix string, recursive, stream bool, sinceIndex uint64) (Watcher, error)
	Save() ([]byte, error)
	Recovery(state []byte) error
	Clone() Store
	SaveNoCopy() ([]byte, error)
	JsonStats() []byte
	DeleteExpiredKeys(cutoff time.Time)
	HasTTLKeys() bool
}
type TTLOptionSet struct {
	ExpireTime time.Time
	Refresh    bool
}
type store struct {
	Root           *node
	WatcherHub     *watcherHub
	CurrentIndex   uint64
	Stats          *Stats
	CurrentVersion int
	ttlKeyHeap     *ttlKeyHeap
	worldLock      sync.RWMutex
	clock          clockwork.Clock
	readonlySet    types.Set
}

func New(namespaces ...string) Store {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := newStore(namespaces...)
	s.clock = clockwork.NewRealClock()
	return s
}
func newStore(namespaces ...string) *store {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := new(store)
	s.CurrentVersion = defaultVersion
	s.Root = newDir(s, "/", s.CurrentIndex, nil, Permanent)
	for _, namespace := range namespaces {
		s.Root.Add(newDir(s, namespace, s.CurrentIndex, s.Root, Permanent))
	}
	s.Stats = newStats()
	s.WatcherHub = newWatchHub(1000)
	s.ttlKeyHeap = newTtlKeyHeap()
	s.readonlySet = types.NewUnsafeSet(append(namespaces, "/")...)
	return s
}
func (s *store) Version() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.CurrentVersion
}
func (s *store) Index() uint64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.worldLock.RLock()
	defer s.worldLock.RUnlock()
	return s.CurrentIndex
}
func (s *store) Get(nodePath string, recursive, sorted bool) (*Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err *etcdErr.Error
	s.worldLock.RLock()
	defer s.worldLock.RUnlock()
	defer func() {
		if err == nil {
			s.Stats.Inc(GetSuccess)
			if recursive {
				reportReadSuccess(GetRecursive)
			} else {
				reportReadSuccess(Get)
			}
			return
		}
		s.Stats.Inc(GetFail)
		if recursive {
			reportReadFailure(GetRecursive)
		} else {
			reportReadFailure(Get)
		}
	}()
	n, err := s.internalGet(nodePath)
	if err != nil {
		return nil, err
	}
	e := newEvent(Get, nodePath, n.ModifiedIndex, n.CreatedIndex)
	e.EtcdIndex = s.CurrentIndex
	e.Node.loadInternalNode(n, recursive, sorted, s.clock)
	return e, nil
}
func (s *store) Create(nodePath string, dir bool, value string, unique bool, expireOpts TTLOptionSet) (*Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err *etcdErr.Error
	s.worldLock.Lock()
	defer s.worldLock.Unlock()
	defer func() {
		if err == nil {
			s.Stats.Inc(CreateSuccess)
			reportWriteSuccess(Create)
			return
		}
		s.Stats.Inc(CreateFail)
		reportWriteFailure(Create)
	}()
	e, err := s.internalCreate(nodePath, dir, value, unique, false, expireOpts.ExpireTime, Create)
	if err != nil {
		return nil, err
	}
	e.EtcdIndex = s.CurrentIndex
	s.WatcherHub.notify(e)
	return e, nil
}
func (s *store) Set(nodePath string, dir bool, value string, expireOpts TTLOptionSet) (*Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err *etcdErr.Error
	s.worldLock.Lock()
	defer s.worldLock.Unlock()
	defer func() {
		if err == nil {
			s.Stats.Inc(SetSuccess)
			reportWriteSuccess(Set)
			return
		}
		s.Stats.Inc(SetFail)
		reportWriteFailure(Set)
	}()
	n, getErr := s.internalGet(nodePath)
	if getErr != nil && getErr.ErrorCode != etcdErr.EcodeKeyNotFound {
		err = getErr
		return nil, err
	}
	if expireOpts.Refresh {
		if getErr != nil {
			err = getErr
			return nil, err
		} else {
			value = n.Value
		}
	}
	e, err := s.internalCreate(nodePath, dir, value, false, true, expireOpts.ExpireTime, Set)
	if err != nil {
		return nil, err
	}
	e.EtcdIndex = s.CurrentIndex
	if getErr == nil {
		prev := newEvent(Get, nodePath, n.ModifiedIndex, n.CreatedIndex)
		prev.Node.loadInternalNode(n, false, false, s.clock)
		e.PrevNode = prev.Node
	}
	if !expireOpts.Refresh {
		s.WatcherHub.notify(e)
	} else {
		e.SetRefresh()
		s.WatcherHub.add(e)
	}
	return e, nil
}
func getCompareFailCause(n *node, which int, prevValue string, prevIndex uint64) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch which {
	case CompareIndexNotMatch:
		return fmt.Sprintf("[%v != %v]", prevIndex, n.ModifiedIndex)
	case CompareValueNotMatch:
		return fmt.Sprintf("[%v != %v]", prevValue, n.Value)
	default:
		return fmt.Sprintf("[%v != %v] [%v != %v]", prevValue, n.Value, prevIndex, n.ModifiedIndex)
	}
}
func (s *store) CompareAndSwap(nodePath string, prevValue string, prevIndex uint64, value string, expireOpts TTLOptionSet) (*Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err *etcdErr.Error
	s.worldLock.Lock()
	defer s.worldLock.Unlock()
	defer func() {
		if err == nil {
			s.Stats.Inc(CompareAndSwapSuccess)
			reportWriteSuccess(CompareAndSwap)
			return
		}
		s.Stats.Inc(CompareAndSwapFail)
		reportWriteFailure(CompareAndSwap)
	}()
	nodePath = path.Clean(path.Join("/", nodePath))
	if s.readonlySet.Contains(nodePath) {
		return nil, etcdErr.NewError(etcdErr.EcodeRootROnly, "/", s.CurrentIndex)
	}
	n, err := s.internalGet(nodePath)
	if err != nil {
		return nil, err
	}
	if n.IsDir() {
		err = etcdErr.NewError(etcdErr.EcodeNotFile, nodePath, s.CurrentIndex)
		return nil, err
	}
	if ok, which := n.Compare(prevValue, prevIndex); !ok {
		cause := getCompareFailCause(n, which, prevValue, prevIndex)
		err = etcdErr.NewError(etcdErr.EcodeTestFailed, cause, s.CurrentIndex)
		return nil, err
	}
	if expireOpts.Refresh {
		value = n.Value
	}
	s.CurrentIndex++
	e := newEvent(CompareAndSwap, nodePath, s.CurrentIndex, n.CreatedIndex)
	e.EtcdIndex = s.CurrentIndex
	e.PrevNode = n.Repr(false, false, s.clock)
	eNode := e.Node
	n.Write(value, s.CurrentIndex)
	n.UpdateTTL(expireOpts.ExpireTime)
	valueCopy := value
	eNode.Value = &valueCopy
	eNode.Expiration, eNode.TTL = n.expirationAndTTL(s.clock)
	if !expireOpts.Refresh {
		s.WatcherHub.notify(e)
	} else {
		e.SetRefresh()
		s.WatcherHub.add(e)
	}
	return e, nil
}
func (s *store) Delete(nodePath string, dir, recursive bool) (*Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err *etcdErr.Error
	s.worldLock.Lock()
	defer s.worldLock.Unlock()
	defer func() {
		if err == nil {
			s.Stats.Inc(DeleteSuccess)
			reportWriteSuccess(Delete)
			return
		}
		s.Stats.Inc(DeleteFail)
		reportWriteFailure(Delete)
	}()
	nodePath = path.Clean(path.Join("/", nodePath))
	if s.readonlySet.Contains(nodePath) {
		return nil, etcdErr.NewError(etcdErr.EcodeRootROnly, "/", s.CurrentIndex)
	}
	if recursive {
		dir = true
	}
	n, err := s.internalGet(nodePath)
	if err != nil {
		return nil, err
	}
	nextIndex := s.CurrentIndex + 1
	e := newEvent(Delete, nodePath, nextIndex, n.CreatedIndex)
	e.EtcdIndex = nextIndex
	e.PrevNode = n.Repr(false, false, s.clock)
	eNode := e.Node
	if n.IsDir() {
		eNode.Dir = true
	}
	callback := func(path string) {
		s.WatcherHub.notifyWatchers(e, path, true)
	}
	err = n.Remove(dir, recursive, callback)
	if err != nil {
		return nil, err
	}
	s.CurrentIndex++
	s.WatcherHub.notify(e)
	return e, nil
}
func (s *store) CompareAndDelete(nodePath string, prevValue string, prevIndex uint64) (*Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err *etcdErr.Error
	s.worldLock.Lock()
	defer s.worldLock.Unlock()
	defer func() {
		if err == nil {
			s.Stats.Inc(CompareAndDeleteSuccess)
			reportWriteSuccess(CompareAndDelete)
			return
		}
		s.Stats.Inc(CompareAndDeleteFail)
		reportWriteFailure(CompareAndDelete)
	}()
	nodePath = path.Clean(path.Join("/", nodePath))
	n, err := s.internalGet(nodePath)
	if err != nil {
		return nil, err
	}
	if n.IsDir() {
		return nil, etcdErr.NewError(etcdErr.EcodeNotFile, nodePath, s.CurrentIndex)
	}
	if ok, which := n.Compare(prevValue, prevIndex); !ok {
		cause := getCompareFailCause(n, which, prevValue, prevIndex)
		return nil, etcdErr.NewError(etcdErr.EcodeTestFailed, cause, s.CurrentIndex)
	}
	s.CurrentIndex++
	e := newEvent(CompareAndDelete, nodePath, s.CurrentIndex, n.CreatedIndex)
	e.EtcdIndex = s.CurrentIndex
	e.PrevNode = n.Repr(false, false, s.clock)
	callback := func(path string) {
		s.WatcherHub.notifyWatchers(e, path, true)
	}
	err = n.Remove(false, false, callback)
	if err != nil {
		return nil, err
	}
	s.WatcherHub.notify(e)
	return e, nil
}
func (s *store) Watch(key string, recursive, stream bool, sinceIndex uint64) (Watcher, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.worldLock.RLock()
	defer s.worldLock.RUnlock()
	key = path.Clean(path.Join("/", key))
	if sinceIndex == 0 {
		sinceIndex = s.CurrentIndex + 1
	}
	w, err := s.WatcherHub.watch(key, recursive, stream, sinceIndex, s.CurrentIndex)
	if err != nil {
		return nil, err
	}
	return w, nil
}
func (s *store) walk(nodePath string, walkFunc func(prev *node, component string) (*node, *etcdErr.Error)) (*node, *etcdErr.Error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	components := strings.Split(nodePath, "/")
	curr := s.Root
	var err *etcdErr.Error
	for i := 1; i < len(components); i++ {
		if len(components[i]) == 0 {
			return curr, nil
		}
		curr, err = walkFunc(curr, components[i])
		if err != nil {
			return nil, err
		}
	}
	return curr, nil
}
func (s *store) Update(nodePath string, newValue string, expireOpts TTLOptionSet) (*Event, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err *etcdErr.Error
	s.worldLock.Lock()
	defer s.worldLock.Unlock()
	defer func() {
		if err == nil {
			s.Stats.Inc(UpdateSuccess)
			reportWriteSuccess(Update)
			return
		}
		s.Stats.Inc(UpdateFail)
		reportWriteFailure(Update)
	}()
	nodePath = path.Clean(path.Join("/", nodePath))
	if s.readonlySet.Contains(nodePath) {
		return nil, etcdErr.NewError(etcdErr.EcodeRootROnly, "/", s.CurrentIndex)
	}
	currIndex, nextIndex := s.CurrentIndex, s.CurrentIndex+1
	n, err := s.internalGet(nodePath)
	if err != nil {
		return nil, err
	}
	if n.IsDir() && len(newValue) != 0 {
		return nil, etcdErr.NewError(etcdErr.EcodeNotFile, nodePath, currIndex)
	}
	if expireOpts.Refresh {
		newValue = n.Value
	}
	e := newEvent(Update, nodePath, nextIndex, n.CreatedIndex)
	e.EtcdIndex = nextIndex
	e.PrevNode = n.Repr(false, false, s.clock)
	eNode := e.Node
	n.Write(newValue, nextIndex)
	if n.IsDir() {
		eNode.Dir = true
	} else {
		newValueCopy := newValue
		eNode.Value = &newValueCopy
	}
	n.UpdateTTL(expireOpts.ExpireTime)
	eNode.Expiration, eNode.TTL = n.expirationAndTTL(s.clock)
	if !expireOpts.Refresh {
		s.WatcherHub.notify(e)
	} else {
		e.SetRefresh()
		s.WatcherHub.add(e)
	}
	s.CurrentIndex = nextIndex
	return e, nil
}
func (s *store) internalCreate(nodePath string, dir bool, value string, unique, replace bool, expireTime time.Time, action string) (*Event, *etcdErr.Error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	currIndex, nextIndex := s.CurrentIndex, s.CurrentIndex+1
	if unique {
		nodePath += "/" + fmt.Sprintf("%020s", strconv.FormatUint(nextIndex, 10))
	}
	nodePath = path.Clean(path.Join("/", nodePath))
	if s.readonlySet.Contains(nodePath) {
		return nil, etcdErr.NewError(etcdErr.EcodeRootROnly, "/", currIndex)
	}
	if expireTime.Before(minExpireTime) {
		expireTime = Permanent
	}
	dirName, nodeName := path.Split(nodePath)
	d, err := s.walk(dirName, s.checkDir)
	if err != nil {
		s.Stats.Inc(SetFail)
		reportWriteFailure(action)
		err.Index = currIndex
		return nil, err
	}
	e := newEvent(action, nodePath, nextIndex, nextIndex)
	eNode := e.Node
	n, _ := d.GetChild(nodeName)
	if n != nil {
		if replace {
			if n.IsDir() {
				return nil, etcdErr.NewError(etcdErr.EcodeNotFile, nodePath, currIndex)
			}
			e.PrevNode = n.Repr(false, false, s.clock)
			n.Remove(false, false, nil)
		} else {
			return nil, etcdErr.NewError(etcdErr.EcodeNodeExist, nodePath, currIndex)
		}
	}
	if !dir {
		valueCopy := value
		eNode.Value = &valueCopy
		n = newKV(s, nodePath, value, nextIndex, d, expireTime)
	} else {
		eNode.Dir = true
		n = newDir(s, nodePath, nextIndex, d, expireTime)
	}
	d.Add(n)
	if !n.IsPermanent() {
		s.ttlKeyHeap.push(n)
		eNode.Expiration, eNode.TTL = n.expirationAndTTL(s.clock)
	}
	s.CurrentIndex = nextIndex
	return e, nil
}
func (s *store) internalGet(nodePath string) (*node, *etcdErr.Error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nodePath = path.Clean(path.Join("/", nodePath))
	walkFunc := func(parent *node, name string) (*node, *etcdErr.Error) {
		if !parent.IsDir() {
			err := etcdErr.NewError(etcdErr.EcodeNotDir, parent.Path, s.CurrentIndex)
			return nil, err
		}
		child, ok := parent.Children[name]
		if ok {
			return child, nil
		}
		return nil, etcdErr.NewError(etcdErr.EcodeKeyNotFound, path.Join(parent.Path, name), s.CurrentIndex)
	}
	f, err := s.walk(nodePath, walkFunc)
	if err != nil {
		return nil, err
	}
	return f, nil
}
func (s *store) DeleteExpiredKeys(cutoff time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.worldLock.Lock()
	defer s.worldLock.Unlock()
	for {
		node := s.ttlKeyHeap.top()
		if node == nil || node.ExpireTime.After(cutoff) {
			break
		}
		s.CurrentIndex++
		e := newEvent(Expire, node.Path, s.CurrentIndex, node.CreatedIndex)
		e.EtcdIndex = s.CurrentIndex
		e.PrevNode = node.Repr(false, false, s.clock)
		if node.IsDir() {
			e.Node.Dir = true
		}
		callback := func(path string) {
			s.WatcherHub.notifyWatchers(e, path, true)
		}
		s.ttlKeyHeap.pop()
		node.Remove(true, true, callback)
		reportExpiredKey()
		s.Stats.Inc(ExpireCount)
		s.WatcherHub.notify(e)
	}
}
func (s *store) checkDir(parent *node, dirName string) (*node, *etcdErr.Error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	node, ok := parent.Children[dirName]
	if ok {
		if node.IsDir() {
			return node, nil
		}
		return nil, etcdErr.NewError(etcdErr.EcodeNotDir, node.Path, s.CurrentIndex)
	}
	n := newDir(s, path.Join(parent.Path, dirName), s.CurrentIndex+1, parent, Permanent)
	parent.Children[dirName] = n
	return n, nil
}
func (s *store) Save() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(s.Clone())
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (s *store) SaveNoCopy() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (s *store) Clone() Store {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.worldLock.Lock()
	clonedStore := newStore()
	clonedStore.CurrentIndex = s.CurrentIndex
	clonedStore.Root = s.Root.Clone()
	clonedStore.WatcherHub = s.WatcherHub.clone()
	clonedStore.Stats = s.Stats.clone()
	clonedStore.CurrentVersion = s.CurrentVersion
	s.worldLock.Unlock()
	return clonedStore
}
func (s *store) Recovery(state []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.worldLock.Lock()
	defer s.worldLock.Unlock()
	err := json.Unmarshal(state, s)
	if err != nil {
		return err
	}
	s.ttlKeyHeap = newTtlKeyHeap()
	s.Root.recoverAndclean()
	return nil
}
func (s *store) JsonStats() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Stats.Watchers = uint64(s.WatcherHub.count)
	return s.Stats.toJson()
}
func (s *store) HasTTLKeys() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.worldLock.RLock()
	defer s.worldLock.RUnlock()
	return s.ttlKeyHeap.Len() != 0
}
