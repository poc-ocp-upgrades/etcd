package transport

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewTimeoutTransport(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr, err := NewTimeoutTransport(TLSInfo{}, time.Hour, time.Hour, time.Hour)
	if err != nil {
		t.Fatalf("unexpected NewTimeoutTransport error: %v", err)
	}
	remoteAddr := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.RemoteAddr))
	}
	srv := httptest.NewServer(http.HandlerFunc(remoteAddr))
	defer srv.Close()
	conn, err := tr.Dial("tcp", srv.Listener.Addr().String())
	if err != nil {
		t.Fatalf("unexpected dial error: %v", err)
	}
	defer conn.Close()
	tconn, ok := conn.(*timeoutConn)
	if !ok {
		t.Fatalf("failed to dial out *timeoutConn")
	}
	if tconn.rdtimeoutd != time.Hour {
		t.Errorf("read timeout = %s, want %s", tconn.rdtimeoutd, time.Hour)
	}
	if tconn.wtimeoutd != time.Hour {
		t.Errorf("write timeout = %s, want %s", tconn.wtimeoutd, time.Hour)
	}
	req, err := http.NewRequest("GET", srv.URL, nil)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	resp, err := tr.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	addr0, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	resp, err = tr.RoundTrip(req)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	addr1, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	if bytes.Equal(addr0, addr1) {
		t.Errorf("addr0 = %s addr1= %s, want not equal", string(addr0), string(addr1))
	}
}
