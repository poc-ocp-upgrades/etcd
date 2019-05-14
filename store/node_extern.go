package store

import (
	"github.com/jonboulle/clockwork"
	"sort"
	"time"
)

type NodeExtern struct {
	Key           string      `json:"key,omitempty"`
	Value         *string     `json:"value,omitempty"`
	Dir           bool        `json:"dir,omitempty"`
	Expiration    *time.Time  `json:"expiration,omitempty"`
	TTL           int64       `json:"ttl,omitempty"`
	Nodes         NodeExterns `json:"nodes,omitempty"`
	ModifiedIndex uint64      `json:"modifiedIndex,omitempty"`
	CreatedIndex  uint64      `json:"createdIndex,omitempty"`
}

func (eNode *NodeExtern) loadInternalNode(n *node, recursive, sorted bool, clock clockwork.Clock) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n.IsDir() {
		eNode.Dir = true
		children, _ := n.List()
		eNode.Nodes = make(NodeExterns, len(children))
		i := 0
		for _, child := range children {
			if child.IsHidden() {
				continue
			}
			eNode.Nodes[i] = child.Repr(recursive, sorted, clock)
			i++
		}
		eNode.Nodes = eNode.Nodes[:i]
		if sorted {
			sort.Sort(eNode.Nodes)
		}
	} else {
		value, _ := n.Read()
		eNode.Value = &value
	}
	eNode.Expiration, eNode.TTL = n.expirationAndTTL(clock)
}
func (eNode *NodeExtern) Clone() *NodeExtern {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if eNode == nil {
		return nil
	}
	nn := &NodeExtern{Key: eNode.Key, Dir: eNode.Dir, TTL: eNode.TTL, ModifiedIndex: eNode.ModifiedIndex, CreatedIndex: eNode.CreatedIndex}
	if eNode.Value != nil {
		s := *eNode.Value
		nn.Value = &s
	}
	if eNode.Expiration != nil {
		t := *eNode.Expiration
		nn.Expiration = &t
	}
	if eNode.Nodes != nil {
		nn.Nodes = make(NodeExterns, len(eNode.Nodes))
		for i, n := range eNode.Nodes {
			nn.Nodes[i] = n.Clone()
		}
	}
	return nn
}

type NodeExterns []*NodeExtern

func (ns NodeExterns) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(ns)
}
func (ns NodeExterns) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ns[i].Key < ns[j].Key
}
func (ns NodeExterns) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns[i], ns[j] = ns[j], ns[i]
}
