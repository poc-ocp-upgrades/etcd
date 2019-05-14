package etcdmain

import (
	"fmt"
	"github.com/coreos/etcd/pkg/srv"
	"github.com/coreos/etcd/pkg/transport"
	"os"
)

func discoverEndpoints(dns string, ca string, insecure bool) (s srv.SRVClients) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dns == "" {
		return s
	}
	srvs, err := srv.GetClient("etcd-client", dns)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	endpoints := srvs.Endpoints
	plog.Infof("discovered the cluster %s from %s", endpoints, dns)
	if insecure {
		return *srvs
	}
	tlsInfo := transport.TLSInfo{TrustedCAFile: ca, ServerName: dns}
	plog.Infof("validating discovered endpoints %v", endpoints)
	endpoints, err = transport.ValidateSecureEndpoints(tlsInfo, endpoints)
	if err != nil {
		plog.Warningf("%v", err)
	}
	plog.Infof("using discovered endpoints %v", endpoints)
	eps := make(map[string]struct{})
	for _, ep := range endpoints {
		eps[ep] = struct{}{}
	}
	for i := range srvs.Endpoints {
		if _, ok := eps[srvs.Endpoints[i]]; !ok {
			continue
		}
		s.Endpoints = append(s.Endpoints, srvs.Endpoints[i])
		s.SRVs = append(s.SRVs, srvs.SRVs[i])
	}
	return s
}
