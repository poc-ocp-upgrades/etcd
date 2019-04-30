package stringutil

import (
	"math/rand"
	"time"
)

func UniqueStrings(slen uint, n int) (ss []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	exist := make(map[string]struct{})
	ss = make([]string, 0, n)
	for len(ss) < n {
		s := randString(slen)
		if _, ok := exist[s]; !ok {
			ss = append(ss, s)
			exist[s] = struct{}{}
		}
	}
	return ss
}
func RandomStrings(slen uint, n int) (ss []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss = make([]string, 0, n)
	for i := 0; i < n; i++ {
		ss = append(ss, randString(slen))
	}
	return ss
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randString(l uint) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rand.Seed(time.Now().UnixNano())
	s := make([]byte, l)
	for i := 0; i < int(l); i++ {
		s[i] = chars[rand.Intn(len(chars))]
	}
	return string(s)
}
