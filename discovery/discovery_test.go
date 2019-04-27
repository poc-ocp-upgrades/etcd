package discovery

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"
	"github.com/coreos/etcd/client"
	"github.com/jonboulle/clockwork"
)

const (
	maxRetryInTest = 3
)

func TestNewProxyFuncUnset(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pf, err := newProxyFunc("")
	if pf != nil {
		t.Fatal("unexpected non-nil proxyFunc")
	}
	if err != nil {
		t.Fatalf("unexpected non-nil err: %v", err)
	}
}
func TestNewProxyFuncBad(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []string{"%%", "http://foo.com/%1"}
	for i, in := range tests {
		pf, err := newProxyFunc(in)
		if pf != nil {
			t.Errorf("#%d: unexpected non-nil proxyFunc", i)
		}
		if err == nil {
			t.Errorf("#%d: unexpected nil err", i)
		}
	}
}
func TestNewProxyFunc(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := map[string]string{"bar.com": "http://bar.com", "http://disco.foo.bar": "http://disco.foo.bar"}
	for in, w := range tests {
		pf, err := newProxyFunc(in)
		if pf == nil {
			t.Errorf("%s: unexpected nil proxyFunc", in)
			continue
		}
		if err != nil {
			t.Errorf("%s: unexpected non-nil err: %v", in, err)
			continue
		}
		g, err := pf(&http.Request{})
		if err != nil {
			t.Errorf("%s: unexpected non-nil err: %v", in, err)
		}
		if g.String() != w {
			t.Errorf("%s: proxyURL=%q, want %q", in, g, w)
		}
	}
}
func TestCheckCluster(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := "/prefix/1000"
	self := "/1000/1"
	tests := []struct {
		nodes	[]*client.Node
		index	uint64
		werr	error
		wsize	int
	}{{[]*client.Node{{Key: "/1000/_config/size", Value: "3", CreatedIndex: 1}, {Key: "/1000/_config/"}, {Key: self, CreatedIndex: 2}, {Key: "/1000/2", CreatedIndex: 3}, {Key: "/1000/3", CreatedIndex: 4}, {Key: "/1000/4", CreatedIndex: 5}}, 5, nil, 3}, {[]*client.Node{{Key: "/1000/_config/size", Value: "3", CreatedIndex: 1}, {Key: "/1000/_config/"}, {Key: "/1000/2", CreatedIndex: 2}, {Key: "/1000/3", CreatedIndex: 3}, {Key: self, CreatedIndex: 4}, {Key: "/1000/4", CreatedIndex: 5}}, 5, nil, 3}, {[]*client.Node{{Key: "/1000/_config/size", Value: "3", CreatedIndex: 1}, {Key: "/1000/_config/"}, {Key: "/1000/2", CreatedIndex: 2}, {Key: "/1000/3", CreatedIndex: 3}, {Key: "/1000/4", CreatedIndex: 4}, {Key: self, CreatedIndex: 5}}, 5, ErrFullCluster, 3}, {[]*client.Node{{Key: "/1000/_config/size", Value: "3", CreatedIndex: 1}, {Key: "/1000/_config/"}, {Key: "/1000/2", CreatedIndex: 2}, {Key: "/1000/3", CreatedIndex: 3}}, 3, nil, 3}, {[]*client.Node{{Key: "/1000/_config/size", Value: "3", CreatedIndex: 1}, {Key: "/1000/_config/"}, {Key: "/1000/2", CreatedIndex: 2}, {Key: "/1000/3", CreatedIndex: 3}, {Key: "/1000/4", CreatedIndex: 4}}, 3, ErrFullCluster, 3}, {[]*client.Node{{Key: "/1000/_config/size", Value: "bad", CreatedIndex: 1}}, 0, ErrBadSizeKey, 0}, {[]*client.Node{}, 0, ErrSizeNotFound, 0}}
	for i, tt := range tests {
		var rs []*client.Response
		if len(tt.nodes) > 0 {
			rs = append(rs, &client.Response{Node: tt.nodes[0], Index: tt.index})
			rs = append(rs, &client.Response{Node: &client.Node{Key: cluster, Nodes: tt.nodes[1:]}, Index: tt.index})
		}
		c := &clientWithResp{rs: rs}
		dBase := discovery{cluster: cluster, id: 1, c: c}
		cRetry := &clientWithRetry{failTimes: 3}
		cRetry.rs = rs
		fc := clockwork.NewFakeClock()
		dRetry := discovery{cluster: cluster, id: 1, c: cRetry, clock: fc}
		for _, d := range []discovery{dBase, dRetry} {
			go func() {
				for i := uint(1); i <= maxRetryInTest; i++ {
					fc.BlockUntil(1)
					fc.Advance(time.Second * (0x1 << i))
				}
			}()
			ns, size, index, err := d.checkCluster()
			if err != tt.werr {
				t.Errorf("#%d: err = %v, want %v", i, err, tt.werr)
			}
			if reflect.DeepEqual(ns, tt.nodes) {
				t.Errorf("#%d: nodes = %v, want %v", i, ns, tt.nodes)
			}
			if size != tt.wsize {
				t.Errorf("#%d: size = %v, want %d", i, size, tt.wsize)
			}
			if index != tt.index {
				t.Errorf("#%d: index = %v, want %d", i, index, tt.index)
			}
		}
	}
}
func TestWaitNodes(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	all := []*client.Node{0: {Key: "/1000/1", CreatedIndex: 2}, 1: {Key: "/1000/2", CreatedIndex: 3}, 2: {Key: "/1000/3", CreatedIndex: 4}}
	tests := []struct {
		nodes	[]*client.Node
		rs	[]*client.Response
	}{{all, []*client.Response{}}, {all[:1], []*client.Response{{Node: &client.Node{Key: "/1000/2", CreatedIndex: 3}}, {Node: &client.Node{Key: "/1000/3", CreatedIndex: 4}}}}, {all[:2], []*client.Response{{Node: &client.Node{Key: "/1000/3", CreatedIndex: 4}}}}, {append(all, &client.Node{Key: "/1000/4", CreatedIndex: 5}), []*client.Response{{Node: &client.Node{Key: "/1000/3", CreatedIndex: 4}}}}}
	for i, tt := range tests {
		c := &clientWithResp{rs: nil, w: &watcherWithResp{rs: tt.rs}}
		dBase := &discovery{cluster: "1000", c: c}
		var retryScanResp []*client.Response
		if len(tt.nodes) > 0 {
			retryScanResp = append(retryScanResp, &client.Response{Node: &client.Node{Key: "1000", Value: strconv.Itoa(3)}})
			retryScanResp = append(retryScanResp, &client.Response{Node: &client.Node{Nodes: tt.nodes}})
		}
		cRetry := &clientWithResp{rs: retryScanResp, w: &watcherWithRetry{rs: tt.rs, failTimes: 2}}
		fc := clockwork.NewFakeClock()
		dRetry := &discovery{cluster: "1000", c: cRetry, clock: fc}
		for _, d := range []*discovery{dBase, dRetry} {
			go func() {
				for i := uint(1); i <= maxRetryInTest; i++ {
					fc.BlockUntil(1)
					fc.Advance(time.Second * (0x1 << i))
				}
			}()
			g, err := d.waitNodes(tt.nodes, 3, 0)
			if err != nil {
				t.Errorf("#%d: err = %v, want %v", i, err, nil)
			}
			if !reflect.DeepEqual(g, all) {
				t.Errorf("#%d: all = %v, want %v", i, g, all)
			}
		}
	}
}
func TestCreateSelf(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rs := []*client.Response{{Node: &client.Node{Key: "1000/1", CreatedIndex: 2}}}
	w := &watcherWithResp{rs: rs}
	errw := &watcherWithErr{err: errors.New("watch err")}
	c := &clientWithResp{rs: rs, w: w}
	errc := &clientWithErr{err: errors.New("create err"), w: w}
	errdupc := &clientWithErr{err: client.Error{Code: client.ErrorCodeNodeExist}}
	errwc := &clientWithResp{rs: rs, w: errw}
	tests := []struct {
		c	client.KeysAPI
		werr	error
	}{{c, nil}, {errc, errc.err}, {errwc, errw.err}, {errdupc, ErrDuplicateID}}
	for i, tt := range tests {
		d := discovery{cluster: "1000", c: tt.c}
		if err := d.createSelf(""); err != tt.werr {
			t.Errorf("#%d: err = %v, want %v", i, err, nil)
		}
	}
}
func TestNodesToCluster(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		nodes		[]*client.Node
		size		int
		wcluster	string
		werr		error
	}{{[]*client.Node{0: {Key: "/1000/1", Value: "1=http://1.1.1.1:2380", CreatedIndex: 1}, 1: {Key: "/1000/2", Value: "2=http://2.2.2.2:2380", CreatedIndex: 2}, 2: {Key: "/1000/3", Value: "3=http://3.3.3.3:2380", CreatedIndex: 3}}, 3, "1=http://1.1.1.1:2380,2=http://2.2.2.2:2380,3=http://3.3.3.3:2380", nil}, {[]*client.Node{0: {Key: "/1000/1", Value: "1=http://1.1.1.1:2380", CreatedIndex: 1}, 1: {Key: "/1000/2", Value: "2=http://2.2.2.2:2380", CreatedIndex: 2}, 2: {Key: "/1000/3", Value: "2=http://3.3.3.3:2380", CreatedIndex: 3}}, 3, "1=http://1.1.1.1:2380,2=http://2.2.2.2:2380,2=http://3.3.3.3:2380", ErrDuplicateName}, {[]*client.Node{0: {Key: "/1000/1", Value: "1=1.1.1.1:2380", CreatedIndex: 1}, 1: {Key: "/1000/2", Value: "2=http://2.2.2.2:2380", CreatedIndex: 2}, 2: {Key: "/1000/3", Value: "2=http://3.3.3.3:2380", CreatedIndex: 3}}, 3, "1=1.1.1.1:2380,2=http://2.2.2.2:2380,2=http://3.3.3.3:2380", ErrInvalidURL}}
	for i, tt := range tests {
		cluster, err := nodesToCluster(tt.nodes, tt.size)
		if err != tt.werr {
			t.Errorf("#%d: err = %v, want %v", i, err, tt.werr)
		}
		if !reflect.DeepEqual(cluster, tt.wcluster) {
			t.Errorf("#%d: cluster = %v, want %v", i, cluster, tt.wcluster)
		}
	}
}
func TestSortableNodes(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := []*client.Node{0: {CreatedIndex: 5}, 1: {CreatedIndex: 1}, 2: {CreatedIndex: 3}, 3: {CreatedIndex: 4}}
	for i := 0; i < 10000; i++ {
		ns = append(ns, &client.Node{CreatedIndex: uint64(rand.Int31())})
	}
	sns := sortableNodes{ns}
	sort.Sort(sns)
	var cis []int
	for _, n := range sns.Nodes {
		cis = append(cis, int(n.CreatedIndex))
	}
	if !sort.IntsAreSorted(cis) {
		t.Errorf("isSorted = %v, want %v", sort.IntsAreSorted(cis), true)
	}
	cis = make([]int, 0)
	for _, n := range ns {
		cis = append(cis, int(n.CreatedIndex))
	}
	if !sort.IntsAreSorted(cis) {
		t.Errorf("isSorted = %v, want %v", sort.IntsAreSorted(cis), true)
	}
}
func TestRetryFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nRetries = maxRetryInTest
	defer func() {
		nRetries = math.MaxUint32
	}()
	cluster := "1000"
	c := &clientWithRetry{failTimes: 4}
	fc := clockwork.NewFakeClock()
	d := discovery{cluster: cluster, id: 1, c: c, clock: fc}
	go func() {
		for i := uint(1); i <= maxRetryInTest; i++ {
			fc.BlockUntil(1)
			fc.Advance(time.Second * (0x1 << i))
		}
	}()
	if _, _, _, err := d.checkCluster(); err != ErrTooManyRetries {
		t.Errorf("err = %v, want %v", err, ErrTooManyRetries)
	}
}

type clientWithResp struct {
	rs	[]*client.Response
	w	client.Watcher
	client.KeysAPI
}

func (c *clientWithResp) Create(ctx context.Context, key string, value string) (*client.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(c.rs) == 0 {
		return &client.Response{}, nil
	}
	r := c.rs[0]
	c.rs = c.rs[1:]
	return r, nil
}
func (c *clientWithResp) Get(ctx context.Context, key string, opts *client.GetOptions) (*client.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(c.rs) == 0 {
		return &client.Response{}, &client.Error{Code: client.ErrorCodeKeyNotFound}
	}
	r := c.rs[0]
	c.rs = append(c.rs[1:], r)
	return r, nil
}
func (c *clientWithResp) Watcher(key string, opts *client.WatcherOptions) client.Watcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.w
}

type clientWithErr struct {
	err	error
	w	client.Watcher
	client.KeysAPI
}

func (c *clientWithErr) Create(ctx context.Context, key string, value string) (*client.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &client.Response{}, c.err
}
func (c *clientWithErr) Get(ctx context.Context, key string, opts *client.GetOptions) (*client.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &client.Response{}, c.err
}
func (c *clientWithErr) Watcher(key string, opts *client.WatcherOptions) client.Watcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.w
}

type watcherWithResp struct {
	client.KeysAPI
	rs	[]*client.Response
}

func (w *watcherWithResp) Next(context.Context) (*client.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(w.rs) == 0 {
		return &client.Response{}, nil
	}
	r := w.rs[0]
	w.rs = w.rs[1:]
	return r, nil
}

type watcherWithErr struct{ err error }

func (w *watcherWithErr) Next(context.Context) (*client.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &client.Response{}, w.err
}

type clientWithRetry struct {
	clientWithResp
	failCount	int
	failTimes	int
}

func (c *clientWithRetry) Create(ctx context.Context, key string, value string) (*client.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.failCount < c.failTimes {
		c.failCount++
		return nil, &client.ClusterError{Errors: []error{context.DeadlineExceeded}}
	}
	return c.clientWithResp.Create(ctx, key, value)
}
func (c *clientWithRetry) Get(ctx context.Context, key string, opts *client.GetOptions) (*client.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.failCount < c.failTimes {
		c.failCount++
		return nil, &client.ClusterError{Errors: []error{context.DeadlineExceeded}}
	}
	return c.clientWithResp.Get(ctx, key, opts)
}

type watcherWithRetry struct {
	rs		[]*client.Response
	failCount	int
	failTimes	int
}

func (w *watcherWithRetry) Next(context.Context) (*client.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if w.failCount < w.failTimes {
		w.failCount++
		return nil, &client.ClusterError{Errors: []error{context.DeadlineExceeded}}
	}
	if len(w.rs) == 0 {
		return &client.Response{}, nil
	}
	r := w.rs[0]
	w.rs = w.rs[1:]
	return r, nil
}
