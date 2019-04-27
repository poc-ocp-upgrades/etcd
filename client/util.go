package client

import (
	"regexp"
)

var (
	roleNotFoundRegExp	*regexp.Regexp
	userNotFoundRegExp	*regexp.Regexp
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	roleNotFoundRegExp = regexp.MustCompile("auth: Role .* does not exist.")
	userNotFoundRegExp = regexp.MustCompile("auth: User .* does not exist.")
}
func IsKeyNotFound(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if cErr, ok := err.(Error); ok {
		return cErr.Code == ErrorCodeKeyNotFound
	}
	return false
}
func IsRoleNotFound(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ae, ok := err.(authError); ok {
		return roleNotFoundRegExp.MatchString(ae.Message)
	}
	return false
}
func IsUserNotFound(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ae, ok := err.(authError); ok {
		return userNotFoundRegExp.MatchString(ae.Message)
	}
	return false
}
