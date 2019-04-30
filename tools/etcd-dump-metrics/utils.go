package main

import "sort"

func aggSort(ss []string) (sorted []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	set := make(map[string]struct{})
	for _, s := range ss {
		set[s] = struct{}{}
	}
	sorted = make([]string, 0, len(set))
	for k := range set {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}
func sortMap(set map[string]struct{}) (sorted []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sorted = make([]string, 0, len(set))
	for k := range set {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	return sorted
}
