package client

import (
	"go.etcd.io/etcd/pkg/srv"
)

type Discoverer interface {
	Discover(domain string, serviceName string) ([]string, error)
}
type srvDiscover struct{}

func NewSRVDiscover() Discoverer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &srvDiscover{}
}
func (d *srvDiscover) Discover(domain string, serviceName string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	srvs, err := srv.GetClient("etcd-client", domain, serviceName)
	if err != nil {
		return nil, err
	}
	return srvs.Endpoints, nil
}
