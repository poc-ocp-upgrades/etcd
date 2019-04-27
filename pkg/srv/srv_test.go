package srv

import (
	"errors"
	"net"
	"reflect"
	"strings"
	"testing"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestSRVGetCluster(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer func() {
		lookupSRV = net.LookupSRV
		resolveTCPAddr = net.ResolveTCPAddr
	}()
	name := "dnsClusterTest"
	dns := map[string]string{"1.example.com.:2480": "10.0.0.1:2480", "2.example.com.:2480": "10.0.0.2:2480", "3.example.com.:2480": "10.0.0.3:2480", "4.example.com.:2380": "10.0.0.3:2380"}
	srvAll := []*net.SRV{{Target: "1.example.com.", Port: 2480}, {Target: "2.example.com.", Port: 2480}, {Target: "3.example.com.", Port: 2480}}
	tests := []struct {
		withSSL		[]*net.SRV
		withoutSSL	[]*net.SRV
		urls		[]string
		expected	string
	}{{[]*net.SRV{}, []*net.SRV{}, nil, ""}, {srvAll, []*net.SRV{}, nil, "0=https://1.example.com:2480,1=https://2.example.com:2480,2=https://3.example.com:2480"}, {srvAll, []*net.SRV{{Target: "4.example.com.", Port: 2380}}, nil, "0=https://1.example.com:2480,1=https://2.example.com:2480,2=https://3.example.com:2480,3=http://4.example.com:2380"}, {srvAll, []*net.SRV{{Target: "4.example.com.", Port: 2380}}, []string{"https://10.0.0.1:2480"}, "dnsClusterTest=https://1.example.com:2480,0=https://2.example.com:2480,1=https://3.example.com:2480,2=http://4.example.com:2380"}, {srvAll, nil, []string{"https://10.0.0.1:2480"}, "dnsClusterTest=https://1.example.com:2480,0=https://2.example.com:2480,1=https://3.example.com:2480"}, {nil, srvAll, []string{"https://10.0.0.1:2480"}, "0=http://2.example.com:2480,1=http://3.example.com:2480"}}
	resolveTCPAddr = func(network, addr string) (*net.TCPAddr, error) {
		if strings.Contains(addr, "10.0.0.") {
			return net.ResolveTCPAddr(network, addr)
		}
		if dns[addr] == "" {
			return nil, errors.New("missing dns record")
		}
		return net.ResolveTCPAddr(network, dns[addr])
	}
	for i, tt := range tests {
		lookupSRV = func(service string, proto string, domain string) (string, []*net.SRV, error) {
			if service == "etcd-server-ssl" {
				return "", tt.withSSL, nil
			}
			if service == "etcd-server" {
				return "", tt.withoutSSL, nil
			}
			return "", nil, errors.New("Unknown service in mock")
		}
		urls := testutil.MustNewURLs(t, tt.urls)
		str, err := GetCluster("etcd-server", name, "example.com", urls)
		if err != nil {
			t.Fatalf("%d: err: %#v", i, err)
		}
		if strings.Join(str, ",") != tt.expected {
			t.Errorf("#%d: cluster = %s, want %s", i, str, tt.expected)
		}
	}
}
func TestSRVDiscover(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer func() {
		lookupSRV = net.LookupSRV
	}()
	tests := []struct {
		withSSL		[]*net.SRV
		withoutSSL	[]*net.SRV
		expected	[]string
	}{{[]*net.SRV{}, []*net.SRV{}, []string{}}, {[]*net.SRV{{Target: "10.0.0.1", Port: 2480}, {Target: "10.0.0.2", Port: 2480}, {Target: "10.0.0.3", Port: 2480}}, []*net.SRV{}, []string{"https://10.0.0.1:2480", "https://10.0.0.2:2480", "https://10.0.0.3:2480"}}, {[]*net.SRV{{Target: "10.0.0.1", Port: 2480}, {Target: "10.0.0.2", Port: 2480}, {Target: "10.0.0.3", Port: 2480}}, []*net.SRV{{Target: "10.0.0.1", Port: 7001}}, []string{"https://10.0.0.1:2480", "https://10.0.0.2:2480", "https://10.0.0.3:2480", "http://10.0.0.1:7001"}}, {[]*net.SRV{{Target: "10.0.0.1", Port: 2480}, {Target: "10.0.0.2", Port: 2480}, {Target: "10.0.0.3", Port: 2480}}, []*net.SRV{{Target: "10.0.0.1", Port: 7001}}, []string{"https://10.0.0.1:2480", "https://10.0.0.2:2480", "https://10.0.0.3:2480", "http://10.0.0.1:7001"}}, {[]*net.SRV{{Target: "a.example.com", Port: 2480}, {Target: "b.example.com", Port: 2480}, {Target: "c.example.com", Port: 2480}}, []*net.SRV{}, []string{"https://a.example.com:2480", "https://b.example.com:2480", "https://c.example.com:2480"}}}
	for i, tt := range tests {
		lookupSRV = func(service string, proto string, domain string) (string, []*net.SRV, error) {
			if service == "etcd-client-ssl" {
				return "", tt.withSSL, nil
			}
			if service == "etcd-client" {
				return "", tt.withoutSSL, nil
			}
			return "", nil, errors.New("Unknown service in mock")
		}
		srvs, err := GetClient("etcd-client", "example.com")
		if err != nil {
			t.Fatalf("%d: err: %#v", i, err)
		}
		if !reflect.DeepEqual(srvs.Endpoints, tt.expected) {
			t.Errorf("#%d: endpoints = %v, want %v", i, srvs.Endpoints, tt.expected)
		}
	}
}
