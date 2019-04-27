package flags

import (
	"strings"
	"github.com/coreos/etcd/pkg/types"
)

type URLsValue types.URLs

func (us *URLsValue) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	strs := strings.Split(s, ",")
	nus, err := types.NewURLs(strs)
	if err != nil {
		return err
	}
	*us = URLsValue(nus)
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
func NewURLsValue(init string) *URLsValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v := &URLsValue{}
	if err := v.Set(init); err != nil {
		plog.Panicf("new URLsValue should never fail: %v", err)
	}
	return v
}
