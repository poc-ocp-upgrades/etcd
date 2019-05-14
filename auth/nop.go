package auth

import (
	"context"
)

type tokenNop struct{}

func (t *tokenNop) enable() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *tokenNop) disable() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *tokenNop) invalidateUser(string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *tokenNop) genTokenPrefix() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "", nil
}
func (t *tokenNop) info(ctx context.Context, token string, rev uint64) (*AuthInfo, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, false
}
func (t *tokenNop) assign(ctx context.Context, username string, revision uint64) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "", ErrAuthFailed
}
func newTokenProviderNop() (*tokenNop, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &tokenNop{}, nil
}
