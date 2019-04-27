package store

import (
	"path"
	"sort"
	"time"
	etcdErr "github.com/coreos/etcd/error"
	"github.com/jonboulle/clockwork"
)

const (
	CompareMatch	= iota
	CompareIndexNotMatch
	CompareValueNotMatch
	CompareNotMatch
)

var Permanent time.Time

type node struct {
	Path		string
	CreatedIndex	uint64
	ModifiedIndex	uint64
	Parent		*node	`json:"-"`
	ExpireTime	time.Time
	Value		string
	Children	map[string]*node
	store		*store
}

func newKV(store *store, nodePath string, value string, createdIndex uint64, parent *node, expireTime time.Time) *node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &node{Path: nodePath, CreatedIndex: createdIndex, ModifiedIndex: createdIndex, Parent: parent, store: store, ExpireTime: expireTime, Value: value}
}
func newDir(store *store, nodePath string, createdIndex uint64, parent *node, expireTime time.Time) *node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &node{Path: nodePath, CreatedIndex: createdIndex, ModifiedIndex: createdIndex, Parent: parent, ExpireTime: expireTime, Children: make(map[string]*node), store: store}
}
func (n *node) IsHidden() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, name := path.Split(n.Path)
	return name[0] == '_'
}
func (n *node) IsPermanent() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.ExpireTime.IsZero()
}
func (n *node) IsDir() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return n.Children != nil
}
func (n *node) Read() (string, *etcdErr.Error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n.IsDir() {
		return "", etcdErr.NewError(etcdErr.EcodeNotFile, "", n.store.CurrentIndex)
	}
	return n.Value, nil
}
func (n *node) Write(value string, index uint64) *etcdErr.Error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n.IsDir() {
		return etcdErr.NewError(etcdErr.EcodeNotFile, "", n.store.CurrentIndex)
	}
	n.Value = value
	n.ModifiedIndex = index
	return nil
}
func (n *node) expirationAndTTL(clock clockwork.Clock) (*time.Time, int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !n.IsPermanent() {
		ttlN := n.ExpireTime.Sub(clock.Now())
		ttl := ttlN / time.Second
		if (ttlN % time.Second) > 0 {
			ttl++
		}
		t := n.ExpireTime.UTC()
		return &t, int64(ttl)
	}
	return nil, 0
}
func (n *node) List() ([]*node, *etcdErr.Error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !n.IsDir() {
		return nil, etcdErr.NewError(etcdErr.EcodeNotDir, "", n.store.CurrentIndex)
	}
	nodes := make([]*node, len(n.Children))
	i := 0
	for _, node := range n.Children {
		nodes[i] = node
		i++
	}
	return nodes, nil
}
func (n *node) GetChild(name string) (*node, *etcdErr.Error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !n.IsDir() {
		return nil, etcdErr.NewError(etcdErr.EcodeNotDir, n.Path, n.store.CurrentIndex)
	}
	child, ok := n.Children[name]
	if ok {
		return child, nil
	}
	return nil, nil
}
func (n *node) Add(child *node) *etcdErr.Error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !n.IsDir() {
		return etcdErr.NewError(etcdErr.EcodeNotDir, "", n.store.CurrentIndex)
	}
	_, name := path.Split(child.Path)
	if _, ok := n.Children[name]; ok {
		return etcdErr.NewError(etcdErr.EcodeNodeExist, "", n.store.CurrentIndex)
	}
	n.Children[name] = child
	return nil
}
func (n *node) Remove(dir, recursive bool, callback func(path string)) *etcdErr.Error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !n.IsDir() {
		_, name := path.Split(n.Path)
		if n.Parent != nil && n.Parent.Children[name] == n {
			delete(n.Parent.Children, name)
		}
		if callback != nil {
			callback(n.Path)
		}
		if !n.IsPermanent() {
			n.store.ttlKeyHeap.remove(n)
		}
		return nil
	}
	if !dir {
		return etcdErr.NewError(etcdErr.EcodeNotFile, n.Path, n.store.CurrentIndex)
	}
	if len(n.Children) != 0 && !recursive {
		return etcdErr.NewError(etcdErr.EcodeDirNotEmpty, n.Path, n.store.CurrentIndex)
	}
	for _, child := range n.Children {
		child.Remove(true, true, callback)
	}
	_, name := path.Split(n.Path)
	if n.Parent != nil && n.Parent.Children[name] == n {
		delete(n.Parent.Children, name)
		if callback != nil {
			callback(n.Path)
		}
		if !n.IsPermanent() {
			n.store.ttlKeyHeap.remove(n)
		}
	}
	return nil
}
func (n *node) Repr(recursive, sorted bool, clock clockwork.Clock) *NodeExtern {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n.IsDir() {
		node := &NodeExtern{Key: n.Path, Dir: true, ModifiedIndex: n.ModifiedIndex, CreatedIndex: n.CreatedIndex}
		node.Expiration, node.TTL = n.expirationAndTTL(clock)
		if !recursive {
			return node
		}
		children, _ := n.List()
		node.Nodes = make(NodeExterns, len(children))
		i := 0
		for _, child := range children {
			if child.IsHidden() {
				continue
			}
			node.Nodes[i] = child.Repr(recursive, sorted, clock)
			i++
		}
		node.Nodes = node.Nodes[:i]
		if sorted {
			sort.Sort(node.Nodes)
		}
		return node
	}
	value := n.Value
	node := &NodeExtern{Key: n.Path, Value: &value, ModifiedIndex: n.ModifiedIndex, CreatedIndex: n.CreatedIndex}
	node.Expiration, node.TTL = n.expirationAndTTL(clock)
	return node
}
func (n *node) UpdateTTL(expireTime time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !n.IsPermanent() {
		if expireTime.IsZero() {
			n.ExpireTime = expireTime
			n.store.ttlKeyHeap.remove(n)
			return
		}
		n.ExpireTime = expireTime
		n.store.ttlKeyHeap.update(n)
		return
	}
	if expireTime.IsZero() {
		return
	}
	n.ExpireTime = expireTime
	n.store.ttlKeyHeap.push(n)
}
func (n *node) Compare(prevValue string, prevIndex uint64) (ok bool, which int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	indexMatch := (prevIndex == 0 || n.ModifiedIndex == prevIndex)
	valueMatch := (prevValue == "" || n.Value == prevValue)
	ok = valueMatch && indexMatch
	switch {
	case valueMatch && indexMatch:
		which = CompareMatch
	case indexMatch && !valueMatch:
		which = CompareValueNotMatch
	case valueMatch && !indexMatch:
		which = CompareIndexNotMatch
	default:
		which = CompareNotMatch
	}
	return ok, which
}
func (n *node) Clone() *node {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !n.IsDir() {
		newkv := newKV(n.store, n.Path, n.Value, n.CreatedIndex, n.Parent, n.ExpireTime)
		newkv.ModifiedIndex = n.ModifiedIndex
		return newkv
	}
	clone := newDir(n.store, n.Path, n.CreatedIndex, n.Parent, n.ExpireTime)
	clone.ModifiedIndex = n.ModifiedIndex
	for key, child := range n.Children {
		clone.Children[key] = child.Clone()
	}
	return clone
}
func (n *node) recoverAndclean() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n.IsDir() {
		for _, child := range n.Children {
			child.Parent = n
			child.store = n.store
			child.recoverAndclean()
		}
	}
	if !n.ExpireTime.IsZero() {
		n.store.ttlKeyHeap.push(n)
	}
}
