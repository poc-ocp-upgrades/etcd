package clientv3

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
	"google.golang.org/grpc/grpclog"
)

func TestLogger(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	buf := new(bytes.Buffer)
	l := NewLogger(grpclog.NewLoggerV2WithVerbosity(buf, buf, buf, 10))
	l.Infof("hello world!")
	if !strings.Contains(buf.String(), "hello world!") {
		t.Fatalf("expected 'hello world!', got %q", buf.String())
	}
	buf.Reset()
	l.Lvl(10).Infof("Level 10")
	l.Lvl(30).Infof("Level 30")
	if !strings.Contains(buf.String(), "Level 10") {
		t.Fatalf("expected 'Level 10', got %q", buf.String())
	}
	if strings.Contains(buf.String(), "Level 30") {
		t.Fatalf("unexpected 'Level 30', got %q", buf.String())
	}
	buf.Reset()
	l = NewLogger(grpclog.NewLoggerV2(ioutil.Discard, ioutil.Discard, ioutil.Discard))
	l.Infof("ignore this")
	if len(buf.Bytes()) > 0 {
		t.Fatalf("unexpected logs %q", buf.String())
	}
}
