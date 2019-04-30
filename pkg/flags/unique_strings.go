package flags

import (
	"flag"
	"sort"
	"strings"
)

type UniqueStringsValue struct{ Values map[string]struct{} }

func (us *UniqueStringsValue) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	us.Values = make(map[string]struct{})
	for _, v := range strings.Split(s, ",") {
		us.Values[v] = struct{}{}
	}
	return nil
}
func (us *UniqueStringsValue) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Join(us.stringSlice(), ",")
}
func (us *UniqueStringsValue) stringSlice() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss := make([]string, 0, len(us.Values))
	for v := range us.Values {
		ss = append(ss, v)
	}
	sort.Strings(ss)
	return ss
}
func NewUniqueStringsValue(s string) (us *UniqueStringsValue) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	us = &UniqueStringsValue{Values: make(map[string]struct{})}
	if s == "" {
		return us
	}
	if err := us.Set(s); err != nil {
		plog.Panicf("new UniqueStringsValue should never fail: %v", err)
	}
	return us
}
func UniqueStringsFromFlag(fs *flag.FlagSet, flagName string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (*fs.Lookup(flagName).Value.(*UniqueStringsValue)).stringSlice()
}
func UniqueStringsMapFromFlag(fs *flag.FlagSet, flagName string) map[string]struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (*fs.Lookup(flagName).Value.(*UniqueStringsValue)).Values
}
