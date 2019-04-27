package agent

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLAndPort(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	addr := "https://127.0.0.1:2379"
	urlAddr, port, err := getURLAndPort(addr)
	if err != nil {
		t.Fatal(err)
	}
	exp := &url.URL{Scheme: "https", Host: "127.0.0.1:2379"}
	if !reflect.DeepEqual(urlAddr, exp) {
		t.Fatalf("expected %+v, got %+v", exp, urlAddr)
	}
	if port != 2379 {
		t.Fatalf("port expected 2379, got %d", port)
	}
}
