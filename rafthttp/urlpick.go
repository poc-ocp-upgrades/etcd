package rafthttp

import (
	"net/url"
	"sync"
	"github.com/coreos/etcd/pkg/types"
)

type urlPicker struct {
	mu	sync.Mutex
	urls	types.URLs
	picked	int
}

func newURLPicker(urls types.URLs) *urlPicker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &urlPicker{urls: urls}
}
func (p *urlPicker) update(urls types.URLs) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.mu.Lock()
	defer p.mu.Unlock()
	p.urls = urls
	p.picked = 0
}
func (p *urlPicker) pick() url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.urls[p.picked]
}
func (p *urlPicker) unreachable(u url.URL) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.mu.Lock()
	defer p.mu.Unlock()
	if u == p.urls[p.picked] {
		p.picked = (p.picked + 1) % len(p.urls)
	}
}
