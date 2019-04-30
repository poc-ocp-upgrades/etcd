package flags

import (
	"flag"
	"net/url"
	"strings"
	"go.etcd.io/etcd/pkg/types"
)

type URLsValue types.URLs

func (us *URLsValue) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss, err := types.NewURLs(strings.Split(s, ","))
	if err != nil {
		return err
	}
	*us = URLsValue(ss)
	return nil
}
func (us *URLsValue) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	all := make([]string, len(*us))
	for i, u := range *us {
		all[i] = u.String()
	}
	return strings.Join(all, ",")
}
func NewURLsValue(s string) *URLsValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s == "" {
		return &URLsValue{}
	}
	v := &URLsValue{}
	if err := v.Set(s); err != nil {
		plog.Panicf("new URLsValue should never fail: %v", err)
	}
	return v
}
func URLsFromFlag(fs *flag.FlagSet, urlsFlagName string) []url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []url.URL(*fs.Lookup(urlsFlagName).Value.(*URLsValue))
}
