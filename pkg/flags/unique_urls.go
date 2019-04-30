package flags

import (
	"flag"
	"net/url"
	"sort"
	"strings"
	"go.etcd.io/etcd/pkg/types"
)

type UniqueURLs struct {
	Values	map[string]struct{}
	uss	[]url.URL
	Allowed	map[string]struct{}
}

func (us *UniqueURLs) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, ok := us.Values[s]; ok {
		return nil
	}
	if _, ok := us.Allowed[s]; ok {
		us.Values[s] = struct{}{}
		return nil
	}
	ss, err := types.NewURLs(strings.Split(s, ","))
	if err != nil {
		return err
	}
	us.Values = make(map[string]struct{})
	us.uss = make([]url.URL, 0)
	for _, v := range ss {
		us.Values[v.String()] = struct{}{}
		us.uss = append(us.uss, v)
	}
	return nil
}
func (us *UniqueURLs) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	all := make([]string, 0, len(us.Values))
	for u := range us.Values {
		all = append(all, u)
	}
	sort.Strings(all)
	return strings.Join(all, ",")
}
func NewUniqueURLsWithExceptions(s string, exceptions ...string) *UniqueURLs {
	_logClusterCodePath()
	defer _logClusterCodePath()
	us := &UniqueURLs{Values: make(map[string]struct{}), Allowed: make(map[string]struct{})}
	for _, v := range exceptions {
		us.Allowed[v] = struct{}{}
	}
	if s == "" {
		return us
	}
	if err := us.Set(s); err != nil {
		plog.Panicf("new UniqueURLs should never fail: %v", err)
	}
	return us
}
func UniqueURLsFromFlag(fs *flag.FlagSet, urlsFlagName string) []url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (*fs.Lookup(urlsFlagName).Value.(*UniqueURLs)).uss
}
func UniqueURLsMapFromFlag(fs *flag.FlagSet, urlsFlagName string) map[string]struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (*fs.Lookup(urlsFlagName).Value.(*UniqueURLs)).Values
}
