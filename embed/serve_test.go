package embed

import (
	"io/ioutil"
	"os"
	"testing"
	"github.com/coreos/etcd/auth"
)

func TestStartEtcdWrongToken(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tdir, err := ioutil.TempDir(os.TempDir(), "token-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tdir)
	cfg := NewConfig()
	cfg.Dir = tdir
	cfg.AuthToken = "wrong-token"
	if _, err = StartEtcd(cfg); err != auth.ErrInvalidAuthOpts {
		t.Fatalf("expected %v, got %v", auth.ErrInvalidAuthOpts, err)
	}
}
