package main

import (
	"os"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if strings.HasSuffix(os.Args[0], "etcdctl.test") {
		return
	}
	main()
}
