package flags

import (
	"flag"
	"sort"
	"strings"
)

type StringsValue sort.StringSlice

func (ss *StringsValue) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*ss = strings.Split(s, ",")
	return nil
}
func (ss *StringsValue) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Join(*ss, ",")
}
func NewStringsValue(s string) (ss *StringsValue) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s == "" {
		return &StringsValue{}
	}
	ss = new(StringsValue)
	if err := ss.Set(s); err != nil {
		plog.Panicf("new StringsValue should never fail: %v", err)
	}
	return ss
}
func StringsFromFlag(fs *flag.FlagSet, flagName string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string(*fs.Lookup(flagName).Value.(*StringsValue))
}
