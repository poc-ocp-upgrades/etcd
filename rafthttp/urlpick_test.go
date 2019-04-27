package rafthttp

import (
	"net/url"
	"testing"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestURLPickerPickTwice(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	picker := mustNewURLPicker(t, []string{"http://127.0.0.1:2380", "http://127.0.0.1:7001"})
	u := picker.pick()
	urlmap := map[url.URL]bool{{Scheme: "http", Host: "127.0.0.1:2380"}: true, {Scheme: "http", Host: "127.0.0.1:7001"}: true}
	if !urlmap[u] {
		t.Errorf("url picked = %+v, want a possible url in %+v", u, urlmap)
	}
	uu := picker.pick()
	if u != uu {
		t.Errorf("url picked = %+v, want %+v", uu, u)
	}
}
func TestURLPickerUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	picker := mustNewURLPicker(t, []string{"http://127.0.0.1:2380", "http://127.0.0.1:7001"})
	picker.update(testutil.MustNewURLs(t, []string{"http://localhost:2380", "http://localhost:7001"}))
	u := picker.pick()
	urlmap := map[url.URL]bool{{Scheme: "http", Host: "localhost:2380"}: true, {Scheme: "http", Host: "localhost:7001"}: true}
	if !urlmap[u] {
		t.Errorf("url picked = %+v, want a possible url in %+v", u, urlmap)
	}
}
func TestURLPickerUnreachable(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	picker := mustNewURLPicker(t, []string{"http://127.0.0.1:2380", "http://127.0.0.1:7001"})
	u := picker.pick()
	picker.unreachable(u)
	uu := picker.pick()
	if u == uu {
		t.Errorf("url picked = %+v, want other possible urls", uu)
	}
}
func mustNewURLPicker(t *testing.T, us []string) *urlPicker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	urls := testutil.MustNewURLs(t, us)
	return newURLPicker(urls)
}
