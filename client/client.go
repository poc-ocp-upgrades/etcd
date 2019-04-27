package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"sync"
	"time"
	"github.com/coreos/etcd/version"
)

var (
	ErrNoEndpoints			= errors.New("client: no endpoints available")
	ErrTooManyRedirects		= errors.New("client: too many redirects")
	ErrClusterUnavailable		= errors.New("client: etcd cluster is unavailable or misconfigured")
	ErrNoLeaderEndpoint		= errors.New("client: no leader endpoint available")
	errTooManyRedirectChecks	= errors.New("client: too many redirect checks")
	oneShotCtxValue			interface{}
)
var DefaultRequestTimeout = 5 * time.Second
var DefaultTransport CancelableTransport = &http.Transport{Proxy: http.ProxyFromEnvironment, Dial: (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).Dial, TLSHandshakeTimeout: 10 * time.Second}

type EndpointSelectionMode int

const (
	EndpointSelectionRandom	EndpointSelectionMode	= iota
	EndpointSelectionPrioritizeLeader
)

type Config struct {
	Endpoints		[]string
	Transport		CancelableTransport
	CheckRedirect		CheckRedirectFunc
	Username		string
	Password		string
	HeaderTimeoutPerRequest	time.Duration
	SelectionMode		EndpointSelectionMode
}

func (cfg *Config) transport() CancelableTransport {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if cfg.Transport == nil {
		return DefaultTransport
	}
	return cfg.Transport
}
func (cfg *Config) checkRedirect() CheckRedirectFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if cfg.CheckRedirect == nil {
		return DefaultCheckRedirect
	}
	return cfg.CheckRedirect
}

type CancelableTransport interface {
	http.RoundTripper
	CancelRequest(req *http.Request)
}
type CheckRedirectFunc func(via int) error

var DefaultCheckRedirect CheckRedirectFunc = func(via int) error {
	if via > 10 {
		return ErrTooManyRedirects
	}
	return nil
}

type Client interface {
	Sync(context.Context) error
	AutoSync(context.Context, time.Duration) error
	Endpoints() []string
	SetEndpoints(eps []string) error
	GetVersion(ctx context.Context) (*version.Versions, error)
	httpClient
}

func New(cfg Config) (Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := &httpClusterClient{clientFactory: newHTTPClientFactory(cfg.transport(), cfg.checkRedirect(), cfg.HeaderTimeoutPerRequest), rand: rand.New(rand.NewSource(int64(time.Now().Nanosecond()))), selectionMode: cfg.SelectionMode}
	if cfg.Username != "" {
		c.credentials = &credentials{username: cfg.Username, password: cfg.Password}
	}
	if err := c.SetEndpoints(cfg.Endpoints); err != nil {
		return nil, err
	}
	return c, nil
}

type httpClient interface {
	Do(context.Context, httpAction) (*http.Response, []byte, error)
}

func newHTTPClientFactory(tr CancelableTransport, cr CheckRedirectFunc, headerTimeout time.Duration) httpClientFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(ep url.URL) httpClient {
		return &redirectFollowingHTTPClient{checkRedirect: cr, client: &simpleHTTPClient{transport: tr, endpoint: ep, headerTimeout: headerTimeout}}
	}
}

type credentials struct {
	username	string
	password	string
}
type httpClientFactory func(url.URL) httpClient
type httpAction interface{ HTTPRequest(url.URL) *http.Request }
type httpClusterClient struct {
	clientFactory	httpClientFactory
	endpoints	[]url.URL
	pinned		int
	credentials	*credentials
	sync.RWMutex
	rand		*rand.Rand
	selectionMode	EndpointSelectionMode
}

func (c *httpClusterClient) getLeaderEndpoint(ctx context.Context, eps []url.URL) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ceps := make([]url.URL, len(eps))
	copy(ceps, eps)
	clientCopy := &httpClusterClient{clientFactory: c.clientFactory, credentials: c.credentials, rand: c.rand, pinned: 0, endpoints: ceps}
	mAPI := NewMembersAPI(clientCopy)
	leader, err := mAPI.Leader(ctx)
	if err != nil {
		return "", err
	}
	if len(leader.ClientURLs) == 0 {
		return "", ErrNoLeaderEndpoint
	}
	return leader.ClientURLs[0], nil
}
func (c *httpClusterClient) parseEndpoints(eps []string) ([]url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(eps) == 0 {
		return []url.URL{}, ErrNoEndpoints
	}
	neps := make([]url.URL, len(eps))
	for i, ep := range eps {
		u, err := url.Parse(ep)
		if err != nil {
			return []url.URL{}, err
		}
		neps[i] = *u
	}
	return neps, nil
}
func (c *httpClusterClient) SetEndpoints(eps []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	neps, err := c.parseEndpoints(eps)
	if err != nil {
		return err
	}
	c.Lock()
	defer c.Unlock()
	c.endpoints = shuffleEndpoints(c.rand, neps)
	c.pinned = 0
	return nil
}
func (c *httpClusterClient) Do(ctx context.Context, act httpAction) (*http.Response, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := act
	c.RLock()
	leps := len(c.endpoints)
	eps := make([]url.URL, leps)
	n := copy(eps, c.endpoints)
	pinned := c.pinned
	if c.credentials != nil {
		action = &authedAction{act: act, credentials: *c.credentials}
	}
	c.RUnlock()
	if leps == 0 {
		return nil, nil, ErrNoEndpoints
	}
	if leps != n {
		return nil, nil, errors.New("unable to pick endpoint: copy failed")
	}
	var resp *http.Response
	var body []byte
	var err error
	cerr := &ClusterError{}
	isOneShot := ctx.Value(&oneShotCtxValue) != nil
	for i := pinned; i < leps+pinned; i++ {
		k := i % leps
		hc := c.clientFactory(eps[k])
		resp, body, err = hc.Do(ctx, action)
		if err != nil {
			cerr.Errors = append(cerr.Errors, err)
			if err == ctx.Err() {
				return nil, nil, ctx.Err()
			}
			if err == context.Canceled || err == context.DeadlineExceeded {
				return nil, nil, err
			}
		} else if resp.StatusCode/100 == 5 {
			switch resp.StatusCode {
			case http.StatusInternalServerError, http.StatusServiceUnavailable:
				cerr.Errors = append(cerr.Errors, fmt.Errorf("client: etcd member %s has no leader", eps[k].String()))
			default:
				cerr.Errors = append(cerr.Errors, fmt.Errorf("client: etcd member %s returns server error [%s]", eps[k].String(), http.StatusText(resp.StatusCode)))
			}
			err = cerr.Errors[0]
		}
		if err != nil {
			if !isOneShot {
				continue
			}
			c.Lock()
			c.pinned = (k + 1) % leps
			c.Unlock()
			return nil, nil, err
		}
		if k != pinned {
			c.Lock()
			c.pinned = k
			c.Unlock()
		}
		return resp, body, nil
	}
	return nil, nil, cerr
}
func (c *httpClusterClient) Endpoints() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.RLock()
	defer c.RUnlock()
	eps := make([]string, len(c.endpoints))
	for i, ep := range c.endpoints {
		eps[i] = ep.String()
	}
	return eps
}
func (c *httpClusterClient) Sync(ctx context.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mAPI := NewMembersAPI(c)
	ms, err := mAPI.List(ctx)
	if err != nil {
		return err
	}
	var eps []string
	for _, m := range ms {
		eps = append(eps, m.ClientURLs...)
	}
	neps, err := c.parseEndpoints(eps)
	if err != nil {
		return err
	}
	npin := 0
	switch c.selectionMode {
	case EndpointSelectionRandom:
		c.RLock()
		eq := endpointsEqual(c.endpoints, neps)
		c.RUnlock()
		if eq {
			return nil
		}
		neps = shuffleEndpoints(c.rand, neps)
	case EndpointSelectionPrioritizeLeader:
		nle, err := c.getLeaderEndpoint(ctx, neps)
		if err != nil {
			return ErrNoLeaderEndpoint
		}
		for i, n := range neps {
			if n.String() == nle {
				npin = i
				break
			}
		}
	default:
		return fmt.Errorf("invalid endpoint selection mode: %d", c.selectionMode)
	}
	c.Lock()
	defer c.Unlock()
	c.endpoints = neps
	c.pinned = npin
	return nil
}
func (c *httpClusterClient) AutoSync(ctx context.Context, interval time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		err := c.Sync(ctx)
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}
func (c *httpClusterClient) GetVersion(ctx context.Context) (*version.Versions, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	act := &getAction{Prefix: "/version"}
	resp, body, err := c.Do(ctx, act)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK:
		if len(body) == 0 {
			return nil, ErrEmptyBody
		}
		var vresp version.Versions
		if err := json.Unmarshal(body, &vresp); err != nil {
			return nil, ErrInvalidJSON
		}
		return &vresp, nil
	default:
		var etcdErr Error
		if err := json.Unmarshal(body, &etcdErr); err != nil {
			return nil, ErrInvalidJSON
		}
		return nil, etcdErr
	}
}

type roundTripResponse struct {
	resp	*http.Response
	err	error
}
type simpleHTTPClient struct {
	transport	CancelableTransport
	endpoint	url.URL
	headerTimeout	time.Duration
}

func (c *simpleHTTPClient) Do(ctx context.Context, act httpAction) (*http.Response, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req := act.HTTPRequest(c.endpoint)
	if err := printcURL(req); err != nil {
		return nil, nil, err
	}
	isWait := false
	if req != nil && req.URL != nil {
		ws := req.URL.Query().Get("wait")
		if len(ws) != 0 {
			var err error
			isWait, err = strconv.ParseBool(ws)
			if err != nil {
				return nil, nil, fmt.Errorf("wrong wait value %s (%v for %+v)", ws, err, req)
			}
		}
	}
	var hctx context.Context
	var hcancel context.CancelFunc
	if !isWait && c.headerTimeout > 0 {
		hctx, hcancel = context.WithTimeout(ctx, c.headerTimeout)
	} else {
		hctx, hcancel = context.WithCancel(ctx)
	}
	defer hcancel()
	reqcancel := requestCanceler(c.transport, req)
	rtchan := make(chan roundTripResponse, 1)
	go func() {
		resp, err := c.transport.RoundTrip(req)
		rtchan <- roundTripResponse{resp: resp, err: err}
		close(rtchan)
	}()
	var resp *http.Response
	var err error
	select {
	case rtresp := <-rtchan:
		resp, err = rtresp.resp, rtresp.err
	case <-hctx.Done():
		reqcancel()
		rtresp := <-rtchan
		resp = rtresp.resp
		switch {
		case ctx.Err() != nil:
			err = ctx.Err()
		case hctx.Err() != nil:
			err = fmt.Errorf("client: endpoint %s exceeded header timeout", c.endpoint.String())
		default:
			panic("failed to get error from context")
		}
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return nil, nil, err
	}
	var body []byte
	done := make(chan struct{})
	go func() {
		body, err = ioutil.ReadAll(resp.Body)
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		resp.Body.Close()
		<-done
		return nil, nil, ctx.Err()
	case <-done:
	}
	return resp, body, err
}

type authedAction struct {
	act		httpAction
	credentials	credentials
}

func (a *authedAction) HTTPRequest(url url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := a.act.HTTPRequest(url)
	r.SetBasicAuth(a.credentials.username, a.credentials.password)
	return r
}

type redirectFollowingHTTPClient struct {
	client		httpClient
	checkRedirect	CheckRedirectFunc
}

func (r *redirectFollowingHTTPClient) Do(ctx context.Context, act httpAction) (*http.Response, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	next := act
	for i := 0; i < 100; i++ {
		if i > 0 {
			if err := r.checkRedirect(i); err != nil {
				return nil, nil, err
			}
		}
		resp, body, err := r.client.Do(ctx, next)
		if err != nil {
			return nil, nil, err
		}
		if resp.StatusCode/100 == 3 {
			hdr := resp.Header.Get("Location")
			if hdr == "" {
				return nil, nil, fmt.Errorf("Location header not set")
			}
			loc, err := url.Parse(hdr)
			if err != nil {
				return nil, nil, fmt.Errorf("Location header not valid URL: %s", hdr)
			}
			next = &redirectedHTTPAction{action: act, location: *loc}
			continue
		}
		return resp, body, nil
	}
	return nil, nil, errTooManyRedirectChecks
}

type redirectedHTTPAction struct {
	action		httpAction
	location	url.URL
}

func (r *redirectedHTTPAction) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	orig := r.action.HTTPRequest(ep)
	orig.URL = &r.location
	return orig
}
func shuffleEndpoints(r *rand.Rand, eps []url.URL) []url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	n := len(eps)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		j := r.Intn(i + 1)
		p[i] = p[j]
		p[j] = i
	}
	neps := make([]url.URL, n)
	for i, k := range p {
		neps[i] = eps[k]
	}
	return neps
}
func endpointsEqual(left, right []url.URL) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(left) != len(right) {
		return false
	}
	sLeft := make([]string, len(left))
	sRight := make([]string, len(right))
	for i, l := range left {
		sLeft[i] = l.String()
	}
	for i, r := range right {
		sRight[i] = r.String()
	}
	sort.Strings(sLeft)
	sort.Strings(sRight)
	for i := range sLeft {
		if sLeft[i] != sRight[i] {
			return false
		}
	}
	return true
}
