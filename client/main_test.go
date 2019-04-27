package client_test

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/pkg/transport"
)

var exampleEndpoints []string
var exampleTransport *http.Transport

func TestMain(m *testing.M) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	useCluster, hasRunArg := false, false
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-test.run=") {
			exp := strings.Split(arg, "=")[1]
			match, err := regexp.MatchString(exp, "Example")
			useCluster = (err == nil && match) || strings.Contains(exp, "Example")
			hasRunArg = true
			break
		}
	}
	if !hasRunArg {
		os.Args = append(os.Args, "-test.run=Test")
	}
	var v int
	if useCluster {
		tr, trerr := transport.NewTransport(transport.TLSInfo{}, time.Second)
		if trerr != nil {
			fmt.Fprintf(os.Stderr, "%v", trerr)
			os.Exit(1)
		}
		cfg := integration.ClusterConfig{Size: 1}
		clus := integration.NewClusterV3(nil, &cfg)
		exampleEndpoints = []string{clus.Members[0].URL()}
		exampleTransport = tr
		v = m.Run()
		clus.Terminate(nil)
		if err := testutil.CheckAfterTest(time.Second); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
	} else {
		v = m.Run()
	}
	if v == 0 && testutil.CheckLeakedGoroutine() {
		os.Exit(1)
	}
	os.Exit(v)
}
