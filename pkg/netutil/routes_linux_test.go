package netutil

import "testing"

func TestGetDefaultInterface(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ifc, err := GetDefaultInterfaces()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("default network interfaces: %+v\n", ifc)
}
func TestGetDefaultHost(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ip, err := GetDefaultHost()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("default ip: %v", ip)
}
