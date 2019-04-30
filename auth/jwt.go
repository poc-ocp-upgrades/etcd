package auth

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"errors"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type tokenJWT struct {
	lg		*zap.Logger
	signMethod	jwt.SigningMethod
	key		interface{}
	ttl		time.Duration
	verifyOnly	bool
}

func (t *tokenJWT) enable() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *tokenJWT) disable() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *tokenJWT) invalidateUser(string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (t *tokenJWT) genTokenPrefix() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "", nil
}
func (t *tokenJWT) info(ctx context.Context, token string, rev uint64) (*AuthInfo, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		username	string
		revision	uint64
	)
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != t.signMethod.Alg() {
			return nil, errors.New("invalid signing method")
		}
		switch k := t.key.(type) {
		case *rsa.PrivateKey:
			return &k.PublicKey, nil
		case *ecdsa.PrivateKey:
			return &k.PublicKey, nil
		default:
			return t.key, nil
		}
	})
	if err != nil {
		if t.lg != nil {
			t.lg.Warn("failed to parse a JWT token", zap.String("token", token), zap.Error(err))
		} else {
			plog.Warningf("failed to parse jwt token: %s", err)
		}
		return nil, false
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !parsed.Valid || !ok {
		if t.lg != nil {
			t.lg.Warn("invalid JWT token", zap.String("token", token))
		} else {
			plog.Warningf("invalid jwt token: %s", token)
		}
		return nil, false
	}
	username = claims["username"].(string)
	revision = uint64(claims["revision"].(float64))
	return &AuthInfo{Username: username, Revision: revision}, true
}
func (t *tokenJWT) assign(ctx context.Context, username string, revision uint64) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.verifyOnly {
		return "", ErrVerifyOnly
	}
	tk := jwt.NewWithClaims(t.signMethod, jwt.MapClaims{"username": username, "revision": revision, "exp": time.Now().Add(t.ttl).Unix()})
	token, err := tk.SignedString(t.key)
	if err != nil {
		if t.lg != nil {
			t.lg.Warn("failed to sign a JWT token", zap.String("user-name", username), zap.Uint64("revision", revision), zap.Error(err))
		} else {
			plog.Debugf("failed to sign jwt token: %s", err)
		}
		return "", err
	}
	if t.lg != nil {
		t.lg.Info("created/assigned a new JWT token", zap.String("user-name", username), zap.Uint64("revision", revision), zap.String("token", token))
	} else {
		plog.Debugf("jwt token: %s", token)
	}
	return token, err
}
func newTokenProviderJWT(lg *zap.Logger, optMap map[string]string) (*tokenJWT, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	var opts jwtOptions
	err = opts.ParseWithDefaults(optMap)
	if err != nil {
		if lg != nil {
			lg.Warn("problem loading JWT options", zap.Error(err))
		} else {
			plog.Errorf("problem loading JWT options: %s", err)
		}
		return nil, ErrInvalidAuthOpts
	}
	var keys = make([]string, 0, len(optMap))
	for k := range optMap {
		if !knownOptions[k] {
			keys = append(keys, k)
		}
	}
	if len(keys) > 0 {
		if lg != nil {
			lg.Warn("unknown JWT options", zap.Strings("keys", keys))
		} else {
			plog.Warningf("unknown JWT options: %v", keys)
		}
	}
	key, err := opts.Key()
	if err != nil {
		return nil, err
	}
	t := &tokenJWT{lg: lg, ttl: opts.TTL, signMethod: opts.SignMethod, key: key}
	switch t.signMethod.(type) {
	case *jwt.SigningMethodECDSA:
		if _, ok := t.key.(*ecdsa.PublicKey); ok {
			t.verifyOnly = true
		}
	case *jwt.SigningMethodRSA, *jwt.SigningMethodRSAPSS:
		if _, ok := t.key.(*rsa.PublicKey); ok {
			t.verifyOnly = true
		}
	}
	return t, nil
}
