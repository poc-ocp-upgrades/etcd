package auth

import (
	"context"
	"testing"
)

func TestSimpleTokenDisabled(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	initialState := newTokenProviderSimple(dummyIndexWaiter)
	explicitlyDisabled := newTokenProviderSimple(dummyIndexWaiter)
	explicitlyDisabled.enable()
	explicitlyDisabled.disable()
	for _, tp := range []*tokenSimple{initialState, explicitlyDisabled} {
		ctx := context.WithValue(context.WithValue(context.TODO(), AuthenticateParamIndex{}, uint64(1)), AuthenticateParamSimpleTokenPrefix{}, "dummy")
		token, err := tp.assign(ctx, "user1", 0)
		if err != nil {
			t.Fatal(err)
		}
		authInfo, ok := tp.info(ctx, token, 0)
		if ok {
			t.Errorf("expected (true, \"user1\") got (%t, %s)", ok, authInfo.Username)
		}
		tp.invalidateUser("user1")
	}
}
func TestSimpleTokenAssign(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tp := newTokenProviderSimple(dummyIndexWaiter)
	tp.enable()
	ctx := context.WithValue(context.WithValue(context.TODO(), AuthenticateParamIndex{}, uint64(1)), AuthenticateParamSimpleTokenPrefix{}, "dummy")
	token, err := tp.assign(ctx, "user1", 0)
	if err != nil {
		t.Fatal(err)
	}
	authInfo, ok := tp.info(ctx, token, 0)
	if !ok || authInfo.Username != "user1" {
		t.Errorf("expected (true, \"token2\") got (%t, %s)", ok, authInfo.Username)
	}
	tp.invalidateUser("user1")
	_, ok = tp.info(context.TODO(), token, 0)
	if ok {
		t.Errorf("expected ok == false after user is invalidated")
	}
}
