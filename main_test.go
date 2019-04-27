package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
)

func TestMain(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if strings.HasSuffix(os.Args[0], "etcd.test") {
		return
	}
	notifier := make(chan os.Signal, 1)
	signal.Notify(notifier, syscall.SIGINT, syscall.SIGTERM)
	go main()
	<-notifier
}
