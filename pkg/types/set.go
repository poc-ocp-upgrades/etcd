package types

import (
	"reflect"
	"sort"
	"sync"
)

type Set interface {
	Add(string)
	Remove(string)
	Contains(string) bool
	Equals(Set) bool
	Length() int
	Values() []string
	Copy() Set
	Sub(Set) Set
}

func NewUnsafeSet(values ...string) *unsafeSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	set := &unsafeSet{make(map[string]struct{})}
	for _, v := range values {
		set.Add(v)
	}
	return set
}
func NewThreadsafeSet(values ...string) *tsafeSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	us := NewUnsafeSet(values...)
	return &tsafeSet{us, sync.RWMutex{}}
}

type unsafeSet struct{ d map[string]struct{} }

func (us *unsafeSet) Add(value string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	us.d[value] = struct{}{}
}
func (us *unsafeSet) Remove(value string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	delete(us.d, value)
}
func (us *unsafeSet) Contains(value string) (exists bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, exists = us.d[value]
	return exists
}
func (us *unsafeSet) ContainsAll(values []string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, s := range values {
		if !us.Contains(s) {
			return false
		}
	}
	return true
}
func (us *unsafeSet) Equals(other Set) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	v1 := sort.StringSlice(us.Values())
	v2 := sort.StringSlice(other.Values())
	v1.Sort()
	v2.Sort()
	return reflect.DeepEqual(v1, v2)
}
func (us *unsafeSet) Length() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(us.d)
}
func (us *unsafeSet) Values() (values []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	values = make([]string, 0)
	for val := range us.d {
		values = append(values, val)
	}
	return values
}
func (us *unsafeSet) Copy() Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cp := NewUnsafeSet()
	for val := range us.d {
		cp.Add(val)
	}
	return cp
}
func (us *unsafeSet) Sub(other Set) Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	oValues := other.Values()
	result := us.Copy().(*unsafeSet)
	for _, val := range oValues {
		if _, ok := result.d[val]; !ok {
			continue
		}
		delete(result.d, val)
	}
	return result
}

type tsafeSet struct {
	us	*unsafeSet
	m	sync.RWMutex
}

func (ts *tsafeSet) Add(value string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ts.m.Lock()
	defer ts.m.Unlock()
	ts.us.Add(value)
}
func (ts *tsafeSet) Remove(value string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ts.m.Lock()
	defer ts.m.Unlock()
	ts.us.Remove(value)
}
func (ts *tsafeSet) Contains(value string) (exists bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ts.m.RLock()
	defer ts.m.RUnlock()
	return ts.us.Contains(value)
}
func (ts *tsafeSet) Equals(other Set) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ts.m.RLock()
	defer ts.m.RUnlock()
	return ts.us.Equals(other)
}
func (ts *tsafeSet) Length() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ts.m.RLock()
	defer ts.m.RUnlock()
	return ts.us.Length()
}
func (ts *tsafeSet) Values() (values []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ts.m.RLock()
	defer ts.m.RUnlock()
	return ts.us.Values()
}
func (ts *tsafeSet) Copy() Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ts.m.RLock()
	defer ts.m.RUnlock()
	usResult := ts.us.Copy().(*unsafeSet)
	return &tsafeSet{usResult, sync.RWMutex{}}
}
func (ts *tsafeSet) Sub(other Set) Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ts.m.RLock()
	defer ts.m.RUnlock()
	usResult := ts.us.Sub(other).(*unsafeSet)
	return &tsafeSet{usResult, sync.RWMutex{}}
}
