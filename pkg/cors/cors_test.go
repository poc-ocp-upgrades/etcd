package cors

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCORSInfo(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		s	string
		winfo	CORSInfo
		ws	string
	}{{"", CORSInfo{}, ""}, {"http://127.0.0.1", CORSInfo{"http://127.0.0.1": true}, "http://127.0.0.1"}, {"*", CORSInfo{"*": true}, "*"}, {" http://127.0.0.1 ", CORSInfo{"http://127.0.0.1": true}, "http://127.0.0.1"}, {"http://127.0.0.1,http://127.0.0.2", CORSInfo{"http://127.0.0.1": true, "http://127.0.0.2": true}, "http://127.0.0.1,http://127.0.0.2"}}
	for i, tt := range tests {
		info := CORSInfo{}
		if err := info.Set(tt.s); err != nil {
			t.Errorf("#%d: set error = %v, want nil", i, err)
		}
		if !reflect.DeepEqual(info, tt.winfo) {
			t.Errorf("#%d: info = %v, want %v", i, info, tt.winfo)
		}
		if g := info.String(); g != tt.ws {
			t.Errorf("#%d: info string = %s, want %s", i, g, tt.ws)
		}
	}
}
func TestCORSInfoOriginAllowed(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		set		string
		origin		string
		wallowed	bool
	}{{"http://127.0.0.1,http://127.0.0.2", "http://127.0.0.1", true}, {"http://127.0.0.1,http://127.0.0.2", "http://127.0.0.2", true}, {"http://127.0.0.1,http://127.0.0.2", "*", false}, {"http://127.0.0.1,http://127.0.0.2", "http://127.0.0.3", false}, {"*", "*", true}, {"*", "http://127.0.0.1", true}}
	for i, tt := range tests {
		info := CORSInfo{}
		if err := info.Set(tt.set); err != nil {
			t.Errorf("#%d: set error = %v, want nil", i, err)
		}
		if g := info.OriginAllowed(tt.origin); g != tt.wallowed {
			t.Errorf("#%d: allowed = %v, want %v", i, g, tt.wallowed)
		}
	}
}
func TestCORSHandler(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	info := &CORSInfo{}
	if err := info.Set("http://127.0.0.1,http://127.0.0.2"); err != nil {
		t.Fatalf("unexpected set error: %v", err)
	}
	h := &CORSHandler{Handler: http.NotFoundHandler(), Info: info}
	header := func(origin string) http.Header {
		return http.Header{"Access-Control-Allow-Methods": []string{"POST, GET, OPTIONS, PUT, DELETE"}, "Access-Control-Allow-Origin": []string{origin}, "Access-Control-Allow-Headers": []string{"accept, content-type, authorization"}}
	}
	tests := []struct {
		method	string
		origin	string
		wcode	int
		wheader	http.Header
	}{{"GET", "http://127.0.0.1", http.StatusNotFound, header("http://127.0.0.1")}, {"GET", "http://127.0.0.2", http.StatusNotFound, header("http://127.0.0.2")}, {"GET", "http://127.0.0.3", http.StatusNotFound, http.Header{}}, {"OPTIONS", "http://127.0.0.1", http.StatusOK, header("http://127.0.0.1")}}
	for i, tt := range tests {
		rr := httptest.NewRecorder()
		req := &http.Request{Method: tt.method, Header: http.Header{"Origin": []string{tt.origin}}}
		h.ServeHTTP(rr, req)
		if rr.Code != tt.wcode {
			t.Errorf("#%d: code = %v, want %v", i, rr.Code, tt.wcode)
		}
		rr.HeaderMap.Del("Content-Type")
		rr.HeaderMap.Del("X-Content-Type-Options")
		if !reflect.DeepEqual(rr.HeaderMap, tt.wheader) {
			t.Errorf("#%d: header = %+v, want %+v", i, rr.HeaderMap, tt.wheader)
		}
	}
}
