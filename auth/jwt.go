package auth

import (
	"context"
	"crypto/rsa"
	"io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
)

type tokenJWT struct {
	signMethod	string
	signKey		*rsa.PrivateKey
	verifyKey	*rsa.PublicKey
}

func (t *tokenJWT) enable() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *tokenJWT) disable() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *tokenJWT) invalidateUser(string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *tokenJWT) genTokenPrefix() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "", nil
}
func (t *tokenJWT) info(ctx context.Context, token string, rev uint64) (*AuthInfo, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		username	string
		revision	uint64
	)
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return t.verifyKey, nil
	})
	switch err.(type) {
	case nil:
		if !parsed.Valid {
			plog.Warningf("invalid jwt token: %s", token)
			return nil, false
		}
		claims := parsed.Claims.(jwt.MapClaims)
		username = claims["username"].(string)
		revision = uint64(claims["revision"].(float64))
	default:
		plog.Warningf("failed to parse jwt token: %s", err)
		return nil, false
	}
	return &AuthInfo{Username: username, Revision: revision}, true
}
func (t *tokenJWT) assign(ctx context.Context, username string, revision uint64) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tk := jwt.NewWithClaims(jwt.GetSigningMethod(t.signMethod), jwt.MapClaims{"username": username, "revision": revision})
	token, err := tk.SignedString(t.signKey)
	if err != nil {
		plog.Debugf("failed to sign jwt token: %s", err)
		return "", err
	}
	plog.Debugf("jwt token: %s", token)
	return token, err
}
func prepareOpts(opts map[string]string) (jwtSignMethod, jwtPubKeyPath, jwtPrivKeyPath string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for k, v := range opts {
		switch k {
		case "sign-method":
			jwtSignMethod = v
		case "pub-key":
			jwtPubKeyPath = v
		case "priv-key":
			jwtPrivKeyPath = v
		default:
			plog.Errorf("unknown token specific option: %s", k)
			return "", "", "", ErrInvalidAuthOpts
		}
	}
	if len(jwtSignMethod) == 0 {
		return "", "", "", ErrInvalidAuthOpts
	}
	return jwtSignMethod, jwtPubKeyPath, jwtPrivKeyPath, nil
}
func newTokenProviderJWT(opts map[string]string) (*tokenJWT, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	jwtSignMethod, jwtPubKeyPath, jwtPrivKeyPath, err := prepareOpts(opts)
	if err != nil {
		return nil, ErrInvalidAuthOpts
	}
	t := &tokenJWT{}
	t.signMethod = jwtSignMethod
	verifyBytes, err := ioutil.ReadFile(jwtPubKeyPath)
	if err != nil {
		plog.Errorf("failed to read public key (%s) for jwt: %s", jwtPubKeyPath, err)
		return nil, err
	}
	t.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		plog.Errorf("failed to parse public key (%s): %s", jwtPubKeyPath, err)
		return nil, err
	}
	signBytes, err := ioutil.ReadFile(jwtPrivKeyPath)
	if err != nil {
		plog.Errorf("failed to read private key (%s) for jwt: %s", jwtPrivKeyPath, err)
		return nil, err
	}
	t.signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		plog.Errorf("failed to parse private key (%s): %s", jwtPrivKeyPath, err)
		return nil, err
	}
	return t, nil
}
