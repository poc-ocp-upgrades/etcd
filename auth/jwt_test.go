package auth

import (
	"context"
	"fmt"
	"testing"
)

const (
	jwtPubKey	= "../integration/fixtures/server.crt"
	jwtPrivKey	= "../integration/fixtures/server.key.insecure"
)

func TestJWTInfo(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := map[string]string{"pub-key": jwtPubKey, "priv-key": jwtPrivKey, "sign-method": "RS256"}
	jwt, err := newTokenProviderJWT(opts)
	if err != nil {
		t.Fatal(err)
	}
	token, aerr := jwt.assign(context.TODO(), "abc", 123)
	if aerr != nil {
		t.Fatal(err)
	}
	ai, ok := jwt.info(context.TODO(), token, 123)
	if !ok {
		t.Fatalf("failed to authenticate with token %s", token)
	}
	if ai.Revision != 123 {
		t.Fatalf("expected revision 123, got %d", ai.Revision)
	}
	ai, ok = jwt.info(context.TODO(), "aaa", 120)
	if ok || ai != nil {
		t.Fatalf("expected aaa to fail to authenticate, got %+v", ai)
	}
}
func TestJWTBad(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := map[string]string{"pub-key": jwtPubKey, "priv-key": jwtPrivKey, "sign-method": "RS256"}
	opts["pub-key"] = jwtPrivKey
	if _, err := newTokenProviderJWT(opts); err == nil {
		t.Fatalf("expected failure on missing public key")
	}
	opts["pub-key"] = jwtPubKey
	opts["priv-key"] = jwtPubKey
	if _, err := newTokenProviderJWT(opts); err == nil {
		t.Fatalf("expected failure on missing public key")
	}
	opts["priv-key"] = jwtPrivKey
	delete(opts, "sign-method")
	if _, err := newTokenProviderJWT(opts); err == nil {
		t.Fatal("expected error on missing option")
	}
	opts["sign-method"] = "RS256"
	opts["pub-key"] = "whatever"
	if _, err := newTokenProviderJWT(opts); err == nil {
		t.Fatalf("expected failure on missing public key")
	}
	opts["pub-key"] = jwtPubKey
	opts["priv-key"] = "whatever"
	if _, err := newTokenProviderJWT(opts); err == nil {
		t.Fatalf("expeceted failure on missing private key")
	}
	opts["priv-key"] = jwtPrivKey
}
func testJWTOpts() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s,pub-key=%s,priv-key=%s,sign-method=RS256", tokenTypeJWT, jwtPubKey, jwtPrivKey)
}
