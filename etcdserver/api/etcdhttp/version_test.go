package etcdhttp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/coreos/etcd/version"
)

func TestServeVersion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatalf("error creating request: %v", err)
	}
	rw := httptest.NewRecorder()
	serveVersion(rw, req, "2.1.0")
	if rw.Code != http.StatusOK {
		t.Errorf("code=%d, want %d", rw.Code, http.StatusOK)
	}
	vs := version.Versions{Server: version.Version, Cluster: "2.1.0"}
	w, err := json.Marshal(&vs)
	if err != nil {
		t.Fatal(err)
	}
	if g := rw.Body.String(); g != string(w) {
		t.Fatalf("body = %q, want %q", g, string(w))
	}
	if ct := rw.HeaderMap.Get("Content-Type"); ct != "application/json" {
		t.Errorf("contet-type header = %s, want %s", ct, "application/json")
	}
}
func TestServeVersionFails(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range []string{"CONNECT", "TRACE", "PUT", "POST", "HEAD"} {
		req, err := http.NewRequest(m, "", nil)
		if err != nil {
			t.Fatalf("error creating request: %v", err)
		}
		rw := httptest.NewRecorder()
		serveVersion(rw, req, "2.1.0")
		if rw.Code != http.StatusMethodNotAllowed {
			t.Errorf("method %s: code=%d, want %d", m, rw.Code, http.StatusMethodNotAllowed)
		}
	}
}
