package auth

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	optSignMethod	= "sign-method"
	optPublicKey	= "pub-key"
	optPrivateKey	= "priv-key"
	optTTL		= "ttl"
)

var knownOptions = map[string]bool{optSignMethod: true, optPublicKey: true, optPrivateKey: true, optTTL: true}
var (
	DefaultTTL = 5 * time.Minute
)

type jwtOptions struct {
	SignMethod	jwt.SigningMethod
	PublicKey	[]byte
	PrivateKey	[]byte
	TTL		time.Duration
}

func (opts *jwtOptions) ParseWithDefaults(optMap map[string]string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if opts.TTL == 0 && optMap[optTTL] == "" {
		opts.TTL = DefaultTTL
	}
	return opts.Parse(optMap)
}
func (opts *jwtOptions) Parse(optMap map[string]string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	if ttl := optMap[optTTL]; ttl != "" {
		opts.TTL, err = time.ParseDuration(ttl)
		if err != nil {
			return err
		}
	}
	if file := optMap[optPublicKey]; file != "" {
		opts.PublicKey, err = ioutil.ReadFile(file)
		if err != nil {
			return err
		}
	}
	if file := optMap[optPrivateKey]; file != "" {
		opts.PrivateKey, err = ioutil.ReadFile(file)
		if err != nil {
			return err
		}
	}
	method := optMap[optSignMethod]
	opts.SignMethod = jwt.GetSigningMethod(method)
	if opts.SignMethod == nil {
		return ErrInvalidAuthMethod
	}
	return nil
}
func (opts *jwtOptions) Key() (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch opts.SignMethod.(type) {
	case *jwt.SigningMethodRSA, *jwt.SigningMethodRSAPSS:
		return opts.rsaKey()
	case *jwt.SigningMethodECDSA:
		return opts.ecKey()
	case *jwt.SigningMethodHMAC:
		return opts.hmacKey()
	default:
		return nil, fmt.Errorf("unsupported signing method: %T", opts.SignMethod)
	}
}
func (opts *jwtOptions) hmacKey() (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(opts.PrivateKey) == 0 {
		return nil, ErrMissingKey
	}
	return opts.PrivateKey, nil
}
func (opts *jwtOptions) rsaKey() (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		priv	*rsa.PrivateKey
		pub	*rsa.PublicKey
		err	error
	)
	if len(opts.PrivateKey) > 0 {
		priv, err = jwt.ParseRSAPrivateKeyFromPEM(opts.PrivateKey)
		if err != nil {
			return nil, err
		}
	}
	if len(opts.PublicKey) > 0 {
		pub, err = jwt.ParseRSAPublicKeyFromPEM(opts.PublicKey)
		if err != nil {
			return nil, err
		}
	}
	if priv == nil {
		if pub == nil {
			return nil, ErrMissingKey
		}
		return pub, nil
	}
	if pub != nil && pub.E != priv.E && pub.N.Cmp(priv.N) != 0 {
		return nil, ErrKeyMismatch
	}
	return priv, nil
}
func (opts *jwtOptions) ecKey() (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		priv	*ecdsa.PrivateKey
		pub	*ecdsa.PublicKey
		err	error
	)
	if len(opts.PrivateKey) > 0 {
		priv, err = jwt.ParseECPrivateKeyFromPEM(opts.PrivateKey)
		if err != nil {
			return nil, err
		}
	}
	if len(opts.PublicKey) > 0 {
		pub, err = jwt.ParseECPublicKeyFromPEM(opts.PublicKey)
		if err != nil {
			return nil, err
		}
	}
	if priv == nil {
		if pub == nil {
			return nil, ErrMissingKey
		}
		return pub, nil
	}
	if pub != nil && pub.Curve != priv.Curve && pub.X.Cmp(priv.X) != 0 && pub.Y.Cmp(priv.Y) != 0 {
		return nil, ErrKeyMismatch
	}
	return priv, nil
}
