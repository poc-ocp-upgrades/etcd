package flags

import (
	"errors"
	"flag"
	"sort"
	"strings"
)

func NewStringsFlag(valids ...string) *StringsFlag {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &StringsFlag{Values: valids, val: valids[0]}
}

type StringsFlag struct {
	Values []string
	val    string
}

func (ss *StringsFlag) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, v := range ss.Values {
		if s == v {
			ss.val = s
			return nil
		}
	}
	return errors.New("invalid value")
}
func (ss *StringsFlag) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ss.val
}

type StringsValueV2 sort.StringSlice

func (ss *StringsValueV2) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*ss = strings.Split(s, ",")
	return nil
}
func (ss *StringsValueV2) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Join(*ss, ",")
}
func NewStringsValueV2(s string) (ss *StringsValueV2) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s == "" {
		return &StringsValueV2{}
	}
	ss = new(StringsValueV2)
	if err := ss.Set(s); err != nil {
		plog.Panicf("new StringsValueV2 should never fail: %v", err)
	}
	return ss
}
func StringsFromFlagV2(fs *flag.FlagSet, flagName string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string(*fs.Lookup(flagName).Value.(*StringsValueV2))
}
