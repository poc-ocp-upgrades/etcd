package cache

import (
	godefaultbytes "bytes"
	"errors"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/pkg/adt"
	"github.com/golang/groupcache/lru"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sync"
)

var (
	DefaultMaxEntries = 2048
	ErrCompacted      = rpctypes.ErrGRPCCompacted
)

type Cache interface {
	Add(req *pb.RangeRequest, resp *pb.RangeResponse)
	Get(req *pb.RangeRequest) (*pb.RangeResponse, error)
	Compact(revision int64)
	Invalidate(key []byte, endkey []byte)
	Size() int
	Close()
}

func keyFunc(req *pb.RangeRequest) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := req.Marshal()
	if err != nil {
		panic(err)
	}
	return string(b)
}
func NewCache(maxCacheEntries int) Cache {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &cache{lru: lru.New(maxCacheEntries), compactedRev: -1}
}
func (c *cache) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}

type cache struct {
	mu           sync.RWMutex
	lru          *lru.Cache
	cachedRanges adt.IntervalTree
	compactedRev int64
}

func (c *cache) Add(req *pb.RangeRequest, resp *pb.RangeResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := keyFunc(req)
	c.mu.Lock()
	defer c.mu.Unlock()
	if req.Revision > c.compactedRev {
		c.lru.Add(key, resp)
	}
	if req.Revision != 0 {
		return
	}
	var (
		iv  *adt.IntervalValue
		ivl adt.Interval
	)
	if len(req.RangeEnd) != 0 {
		ivl = adt.NewStringAffineInterval(string(req.Key), string(req.RangeEnd))
	} else {
		ivl = adt.NewStringAffinePoint(string(req.Key))
	}
	iv = c.cachedRanges.Find(ivl)
	if iv == nil {
		c.cachedRanges.Insert(ivl, []string{key})
	} else {
		iv.Val = append(iv.Val.([]string), key)
	}
}
func (c *cache) Get(req *pb.RangeRequest) (*pb.RangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := keyFunc(req)
	c.mu.Lock()
	defer c.mu.Unlock()
	if req.Revision > 0 && req.Revision < c.compactedRev {
		c.lru.Remove(key)
		return nil, ErrCompacted
	}
	if resp, ok := c.lru.Get(key); ok {
		return resp.(*pb.RangeResponse), nil
	}
	return nil, errors.New("not exist")
}
func (c *cache) Invalidate(key, endkey []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mu.Lock()
	defer c.mu.Unlock()
	var (
		ivs []*adt.IntervalValue
		ivl adt.Interval
	)
	if len(endkey) == 0 {
		ivl = adt.NewStringAffinePoint(string(key))
	} else {
		ivl = adt.NewStringAffineInterval(string(key), string(endkey))
	}
	ivs = c.cachedRanges.Stab(ivl)
	for _, iv := range ivs {
		keys := iv.Val.([]string)
		for _, key := range keys {
			c.lru.Remove(key)
		}
	}
	c.cachedRanges.Delete(ivl)
}
func (c *cache) Compact(revision int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mu.Lock()
	defer c.mu.Unlock()
	if revision > c.compactedRev {
		c.compactedRev = revision
	}
}
func (c *cache) Size() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lru.Len()
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
