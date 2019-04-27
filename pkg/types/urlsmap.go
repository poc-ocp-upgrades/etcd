package types

import (
	"fmt"
	"sort"
	"strings"
)

type URLsMap map[string]URLs

func NewURLsMap(s string) (URLsMap, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := parse(s)
	cl := URLsMap{}
	for name, urls := range m {
		us, err := NewURLs(urls)
		if err != nil {
			return nil, err
		}
		cl[name] = us
	}
	return cl, nil
}
func NewURLsMapFromStringMap(m map[string]string, sep string) (URLsMap, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	um := URLsMap{}
	for k, v := range m {
		um[k], err = NewURLs(strings.Split(v, sep))
		if err != nil {
			return nil, err
		}
	}
	return um, nil
}
func (c URLsMap) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var pairs []string
	for name, urls := range c {
		for _, url := range urls {
			pairs = append(pairs, fmt.Sprintf("%s=%s", name, url.String()))
		}
	}
	sort.Strings(pairs)
	return strings.Join(pairs, ",")
}
func (c URLsMap) URLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var urls []string
	for _, us := range c {
		for _, u := range us {
			urls = append(urls, u.String())
		}
	}
	sort.Strings(urls)
	return urls
}
func (c URLsMap) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(c)
}
func parse(s string) map[string][]string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := make(map[string][]string)
	for s != "" {
		key := s
		if i := strings.IndexAny(key, ","); i >= 0 {
			key, s = key[:i], key[i+1:]
		} else {
			s = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}
		m[key] = append(m[key], value)
	}
	return m
}
