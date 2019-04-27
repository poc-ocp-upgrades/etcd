package netutil

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"sort"
	"time"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/pkg/capnslog"
)

var (
	plog		= capnslog.NewPackageLogger("github.com/coreos/etcd", "pkg/netutil")
	resolveTCPAddr	= resolveTCPAddrDefault
)

const retryInterval = time.Second

func resolveTCPAddrDefault(ctx context.Context, addr string) (*net.TCPAddr, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	host, port, serr := net.SplitHostPort(addr)
	if serr != nil {
		return nil, serr
	}
	portnum, perr := net.DefaultResolver.LookupPort(ctx, "tcp", port)
	if perr != nil {
		return nil, perr
	}
	var ips []net.IPAddr
	if ip := net.ParseIP(host); ip != nil {
		ips = []net.IPAddr{{IP: ip}}
	} else {
		ipss, err := net.DefaultResolver.LookupIPAddr(ctx, host)
		if err != nil {
			return nil, err
		}
		ips = ipss
	}
	ip := ips[0]
	return &net.TCPAddr{IP: ip.IP, Port: portnum, Zone: ip.Zone}, nil
}
func resolveTCPAddrs(ctx context.Context, urls [][]url.URL) ([][]url.URL, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newurls := make([][]url.URL, 0)
	for _, us := range urls {
		nus := make([]url.URL, len(us))
		for i, u := range us {
			nu, err := url.Parse(u.String())
			if err != nil {
				return nil, fmt.Errorf("failed to parse %q (%v)", u.String(), err)
			}
			nus[i] = *nu
		}
		for i, u := range nus {
			h, err := resolveURL(ctx, u)
			if err != nil {
				return nil, fmt.Errorf("failed to resolve %q (%v)", u.String(), err)
			}
			if h != "" {
				nus[i].Host = h
			}
		}
		newurls = append(newurls, nus)
	}
	return newurls, nil
}
func resolveURL(ctx context.Context, u url.URL) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if u.Scheme == "unix" || u.Scheme == "unixs" {
		return "", nil
	}
	host, _, err := net.SplitHostPort(u.Host)
	if err != nil {
		plog.Errorf("could not parse url %s during tcp resolving", u.Host)
		return "", err
	}
	if host == "localhost" || net.ParseIP(host) != nil {
		return "", nil
	}
	for ctx.Err() == nil {
		tcpAddr, err := resolveTCPAddr(ctx, u.Host)
		if err == nil {
			plog.Infof("resolving %s to %s", u.Host, tcpAddr.String())
			return tcpAddr.String(), nil
		}
		plog.Warningf("failed resolving host %s (%v); retrying in %v", u.Host, err, retryInterval)
		select {
		case <-ctx.Done():
			plog.Errorf("could not resolve host %s", u.Host)
			return "", err
		case <-time.After(retryInterval):
		}
	}
	return "", ctx.Err()
}
func urlsEqual(ctx context.Context, a []url.URL, b []url.URL) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(a) != len(b) {
		return false, fmt.Errorf("len(%q) != len(%q)", urlsToStrings(a), urlsToStrings(b))
	}
	urls, err := resolveTCPAddrs(ctx, [][]url.URL{a, b})
	if err != nil {
		return false, err
	}
	preva, prevb := a, b
	a, b = urls[0], urls[1]
	sort.Sort(types.URLs(a))
	sort.Sort(types.URLs(b))
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false, fmt.Errorf("%q(resolved from %q) != %q(resolved from %q)", a[i].String(), preva[i].String(), b[i].String(), prevb[i].String())
		}
	}
	return true, nil
}
func URLStringsEqual(ctx context.Context, a []string, b []string) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(a) != len(b) {
		return false, fmt.Errorf("len(%q) != len(%q)", a, b)
	}
	urlsA := make([]url.URL, 0)
	for _, str := range a {
		u, err := url.Parse(str)
		if err != nil {
			return false, fmt.Errorf("failed to parse %q", str)
		}
		urlsA = append(urlsA, *u)
	}
	urlsB := make([]url.URL, 0)
	for _, str := range b {
		u, err := url.Parse(str)
		if err != nil {
			return false, fmt.Errorf("failed to parse %q", str)
		}
		urlsB = append(urlsB, *u)
	}
	return urlsEqual(ctx, urlsA, urlsB)
}
func urlsToStrings(us []url.URL) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rs := make([]string, len(us))
	for i := range us {
		rs[i] = us[i].String()
	}
	return rs
}
func IsNetworkTimeoutError(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	nerr, ok := err.(net.Error)
	return ok && nerr.Timeout()
}
