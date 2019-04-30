package etcdserver

import "sync"

type AccessController struct {
	corsMu		sync.RWMutex
	CORS		map[string]struct{}
	hostWhitelistMu	sync.RWMutex
	HostWhitelist	map[string]struct{}
}

func NewAccessController() *AccessController {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &AccessController{CORS: map[string]struct{}{"*": {}}, HostWhitelist: map[string]struct{}{"*": {}}}
}
func (ac *AccessController) OriginAllowed(origin string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ac.corsMu.RLock()
	defer ac.corsMu.RUnlock()
	if len(ac.CORS) == 0 {
		return true
	}
	_, ok := ac.CORS["*"]
	if ok {
		return true
	}
	_, ok = ac.CORS[origin]
	return ok
}
func (ac *AccessController) IsHostWhitelisted(host string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ac.hostWhitelistMu.RLock()
	defer ac.hostWhitelistMu.RUnlock()
	if len(ac.HostWhitelist) == 0 {
		return true
	}
	_, ok := ac.HostWhitelist["*"]
	if ok {
		return true
	}
	_, ok = ac.HostWhitelist[host]
	return ok
}
