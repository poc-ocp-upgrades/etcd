package client

import "fmt"

type ClusterError struct{ Errors []error }

func (ce *ClusterError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := ErrClusterUnavailable.Error()
	for i, e := range ce.Errors {
		s += fmt.Sprintf("; error #%d: %s\n", i, e)
	}
	return s
}
func (ce *ClusterError) Detail() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := ""
	for i, e := range ce.Errors {
		s += fmt.Sprintf("error #%d: %s\n", i, e)
	}
	return s
}
