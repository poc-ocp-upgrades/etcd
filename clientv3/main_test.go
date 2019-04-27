package clientv3_test

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/auth"
	"github.com/coreos/etcd/integration"
	"github.com/coreos/etcd/pkg/testutil"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	auth.BcryptCost = bcrypt.MinCost
}
func TestMain(m *testing.M) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
		cfg := integration.ClusterConfig{Size: 3}
		clus := integration.NewClusterV3(nil, &cfg)
		endpoints = make([]string, 3)
		for i := range endpoints {
			endpoints[i] = clus.Client(i).Endpoints()[0]
		}
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
