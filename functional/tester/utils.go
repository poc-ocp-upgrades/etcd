package tester

import (
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"strings"
)

func isValidURL(u string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := url.Parse(u)
	return err == nil
}
func getPort(addr string) (port string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	urlAddr, err := url.Parse(addr)
	if err != nil {
		return "", err
	}
	_, port, err = net.SplitHostPort(urlAddr.Host)
	if err != nil {
		return "", err
	}
	return port, nil
}
func getSameValue(vals map[string]int64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var rv int64
	for _, v := range vals {
		if rv == 0 {
			rv = v
		}
		if rv != v {
			return false
		}
	}
	return true
}
func max(n1, n2 int64) int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if n1 > n2 {
		return n1
	}
	return n2
}
func errsToError(errs []error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(errs) == 0 {
		return nil
	}
	stringArr := make([]string, len(errs))
	for i, err := range errs {
		stringArr[i] = err.Error()
	}
	return fmt.Errorf(strings.Join(stringArr, ", "))
}
func randBytes(size int) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = byte(int('a') + rand.Intn(26))
	}
	return data
}
