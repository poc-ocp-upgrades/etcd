package auth

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	letters				= "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	defaultSimpleTokenLength	= 16
)

var (
	simpleTokenTTL			= 5 * time.Minute
	simpleTokenTTLResolution	= 1 * time.Second
)

type simpleTokenTTLKeeper struct {
	tokens		map[string]time.Time
	donec		chan struct{}
	stopc		chan struct{}
	deleteTokenFunc	func(string)
	mu		*sync.Mutex
}

func (tm *simpleTokenTTLKeeper) stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case tm.stopc <- struct{}{}:
	case <-tm.donec:
	}
	<-tm.donec
}
func (tm *simpleTokenTTLKeeper) addSimpleToken(token string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tm.tokens[token] = time.Now().Add(simpleTokenTTL)
}
func (tm *simpleTokenTTLKeeper) resetSimpleToken(token string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, ok := tm.tokens[token]; ok {
		tm.tokens[token] = time.Now().Add(simpleTokenTTL)
	}
}
func (tm *simpleTokenTTLKeeper) deleteSimpleToken(token string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delete(tm.tokens, token)
}
func (tm *simpleTokenTTLKeeper) run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tokenTicker := time.NewTicker(simpleTokenTTLResolution)
	defer func() {
		tokenTicker.Stop()
		close(tm.donec)
	}()
	for {
		select {
		case <-tokenTicker.C:
			nowtime := time.Now()
			tm.mu.Lock()
			for t, tokenendtime := range tm.tokens {
				if nowtime.After(tokenendtime) {
					tm.deleteTokenFunc(t)
					delete(tm.tokens, t)
				}
			}
			tm.mu.Unlock()
		case <-tm.stopc:
			return
		}
	}
}

type tokenSimple struct {
	indexWaiter		func(uint64) <-chan struct{}
	simpleTokenKeeper	*simpleTokenTTLKeeper
	simpleTokensMu		sync.Mutex
	simpleTokens		map[string]string
}

func (t *tokenSimple) genTokenPrefix() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := make([]byte, defaultSimpleTokenLength)
	for i := 0; i < defaultSimpleTokenLength; i++ {
		bInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[bInt.Int64()]
	}
	return string(ret), nil
}
func (t *tokenSimple) assignSimpleTokenToUser(username, token string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.simpleTokensMu.Lock()
	defer t.simpleTokensMu.Unlock()
	if t.simpleTokenKeeper == nil {
		return
	}
	_, ok := t.simpleTokens[token]
	if ok {
		plog.Panicf("token %s is alredy used", token)
	}
	t.simpleTokens[token] = username
	t.simpleTokenKeeper.addSimpleToken(token)
}
func (t *tokenSimple) invalidateUser(username string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if t.simpleTokenKeeper == nil {
		return
	}
	t.simpleTokensMu.Lock()
	for token, name := range t.simpleTokens {
		if strings.Compare(name, username) == 0 {
			delete(t.simpleTokens, token)
			t.simpleTokenKeeper.deleteSimpleToken(token)
		}
	}
	t.simpleTokensMu.Unlock()
}
func (t *tokenSimple) enable() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	delf := func(tk string) {
		if username, ok := t.simpleTokens[tk]; ok {
			plog.Infof("deleting token %s for user %s", tk, username)
			delete(t.simpleTokens, tk)
		}
	}
	t.simpleTokenKeeper = &simpleTokenTTLKeeper{tokens: make(map[string]time.Time), donec: make(chan struct{}), stopc: make(chan struct{}), deleteTokenFunc: delf, mu: &t.simpleTokensMu}
	go t.simpleTokenKeeper.run()
}
func (t *tokenSimple) disable() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.simpleTokensMu.Lock()
	tk := t.simpleTokenKeeper
	t.simpleTokenKeeper = nil
	t.simpleTokens = make(map[string]string)
	t.simpleTokensMu.Unlock()
	if tk != nil {
		tk.stop()
	}
}
func (t *tokenSimple) info(ctx context.Context, token string, revision uint64) (*AuthInfo, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !t.isValidSimpleToken(ctx, token) {
		return nil, false
	}
	t.simpleTokensMu.Lock()
	username, ok := t.simpleTokens[token]
	if ok && t.simpleTokenKeeper != nil {
		t.simpleTokenKeeper.resetSimpleToken(token)
	}
	t.simpleTokensMu.Unlock()
	return &AuthInfo{Username: username, Revision: revision}, ok
}
func (t *tokenSimple) assign(ctx context.Context, username string, rev uint64) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	index := ctx.Value(AuthenticateParamIndex{}).(uint64)
	simpleTokenPrefix := ctx.Value(AuthenticateParamSimpleTokenPrefix{}).(string)
	token := fmt.Sprintf("%s.%d", simpleTokenPrefix, index)
	t.assignSimpleTokenToUser(username, token)
	return token, nil
}
func (t *tokenSimple) isValidSimpleToken(ctx context.Context, token string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	splitted := strings.Split(token, ".")
	if len(splitted) != 2 {
		return false
	}
	index, err := strconv.Atoi(splitted[1])
	if err != nil {
		return false
	}
	select {
	case <-t.indexWaiter(uint64(index)):
		return true
	case <-ctx.Done():
	}
	return false
}
func newTokenProviderSimple(indexWaiter func(uint64) <-chan struct{}) *tokenSimple {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &tokenSimple{simpleTokens: make(map[string]string), indexWaiter: indexWaiter}
}
