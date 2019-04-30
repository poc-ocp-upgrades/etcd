package flags

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type SelectiveStringValue struct {
	v	string
	valids	map[string]struct{}
}

func (ss *SelectiveStringValue) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, ok := ss.valids[s]; ok {
		ss.v = s
		return nil
	}
	return errors.New("invalid value")
}
func (ss *SelectiveStringValue) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ss.v
}
func (ss *SelectiveStringValue) Valids() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := make([]string, 0, len(ss.valids))
	for k := range ss.valids {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}
func NewSelectiveStringValue(valids ...string) *SelectiveStringValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	vm := make(map[string]struct{})
	for _, v := range valids {
		vm[v] = struct{}{}
	}
	return &SelectiveStringValue{valids: vm, v: valids[0]}
}

type SelectiveStringsValue struct {
	vs	[]string
	valids	map[string]struct{}
}

func (ss *SelectiveStringsValue) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	vs := strings.Split(s, ",")
	for i := range vs {
		if _, ok := ss.valids[vs[i]]; ok {
			ss.vs = append(ss.vs, vs[i])
		} else {
			return fmt.Errorf("invalid value %q", vs[i])
		}
	}
	sort.Strings(ss.vs)
	return nil
}
func (ss *SelectiveStringsValue) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Join(ss.vs, ",")
}
func (ss *SelectiveStringsValue) Valids() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := make([]string, 0, len(ss.valids))
	for k := range ss.valids {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}
func NewSelectiveStringsValue(valids ...string) *SelectiveStringsValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	vm := make(map[string]struct{})
	for _, v := range valids {
		vm[v] = struct{}{}
	}
	return &SelectiveStringsValue{valids: vm, vs: []string{}}
}
