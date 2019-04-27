package testutil

import (
	"fmt"
	"os"
	"testing"
)

var ranSample = false

func TestMain(m *testing.M) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m.Run()
	isLeaked := CheckLeakedGoroutine()
	if ranSample && !isLeaked {
		fmt.Fprintln(os.Stderr, "expected leaky goroutines but none is detected")
		os.Exit(1)
	}
	os.Exit(0)
}
func TestSample(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer AfterTest(t)
	ranSample = true
	for range make([]struct{}, 100) {
		go func() {
			select {}
		}()
	}
}
