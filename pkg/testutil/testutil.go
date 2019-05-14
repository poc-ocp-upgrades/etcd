package testutil

import (
	"net/url"
	"runtime"
	"testing"
	"time"
)

func WaitSchedule() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	time.Sleep(10 * time.Millisecond)
}
func MustNewURLs(t *testing.T, urls []string) []url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if urls == nil {
		return nil
	}
	var us []url.URL
	for _, url := range urls {
		u := MustNewURL(t, url)
		us = append(us, *u)
	}
	return us
}
func MustNewURL(t *testing.T, s string) *url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("parse %v error: %v", s, err)
	}
	return u
}
func FatalStack(t *testing.T, s string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stackTrace := make([]byte, 1024*1024)
	n := runtime.Stack(stackTrace, true)
	t.Error(string(stackTrace[:n]))
	t.Fatalf(s)
}

type ConditionFunc func() (bool, error)

func Poll(interval time.Duration, timeout time.Duration, condition ConditionFunc) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	timeoutCh := time.After(timeout)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-timeoutCh:
			return false, nil
		case <-ticker.C:
			success, err := condition()
			if err != nil {
				return false, err
			}
			if success {
				return true, nil
			}
		}
	}
}
