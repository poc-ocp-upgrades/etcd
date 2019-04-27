package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"github.com/coreos/etcd/pkg/pathutil"
	"github.com/ugorji/go/codec"
)

const (
	ErrorCodeKeyNotFound		= 100
	ErrorCodeTestFailed		= 101
	ErrorCodeNotFile		= 102
	ErrorCodeNotDir			= 104
	ErrorCodeNodeExist		= 105
	ErrorCodeRootROnly		= 107
	ErrorCodeDirNotEmpty		= 108
	ErrorCodeUnauthorized		= 110
	ErrorCodePrevValueRequired	= 201
	ErrorCodeTTLNaN			= 202
	ErrorCodeIndexNaN		= 203
	ErrorCodeInvalidField		= 209
	ErrorCodeInvalidForm		= 210
	ErrorCodeRaftInternal		= 300
	ErrorCodeLeaderElect		= 301
	ErrorCodeWatcherCleared		= 400
	ErrorCodeEventIndexCleared	= 401
)

type Error struct {
	Code	int	`json:"errorCode"`
	Message	string	`json:"message"`
	Cause	string	`json:"cause"`
	Index	uint64	`json:"index"`
}

func (e Error) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%v: %v (%v) [%v]", e.Code, e.Message, e.Cause, e.Index)
}

var (
	ErrInvalidJSON	= errors.New("client: response is invalid json. The endpoint is probably not valid etcd cluster endpoint.")
	ErrEmptyBody	= errors.New("client: response body is empty")
)

type PrevExistType string

const (
	PrevIgnore	= PrevExistType("")
	PrevExist	= PrevExistType("true")
	PrevNoExist	= PrevExistType("false")
)

var (
	defaultV2KeysPrefix = "/v2/keys"
)

func NewKeysAPI(c Client) KeysAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewKeysAPIWithPrefix(c, defaultV2KeysPrefix)
}
func NewKeysAPIWithPrefix(c Client, p string) KeysAPI {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &httpKeysAPI{client: c, prefix: p}
}

type KeysAPI interface {
	Get(ctx context.Context, key string, opts *GetOptions) (*Response, error)
	Set(ctx context.Context, key, value string, opts *SetOptions) (*Response, error)
	Delete(ctx context.Context, key string, opts *DeleteOptions) (*Response, error)
	Create(ctx context.Context, key, value string) (*Response, error)
	CreateInOrder(ctx context.Context, dir, value string, opts *CreateInOrderOptions) (*Response, error)
	Update(ctx context.Context, key, value string) (*Response, error)
	Watcher(key string, opts *WatcherOptions) Watcher
}
type WatcherOptions struct {
	AfterIndex	uint64
	Recursive	bool
}
type CreateInOrderOptions struct{ TTL time.Duration }
type SetOptions struct {
	PrevValue		string
	PrevIndex		uint64
	PrevExist		PrevExistType
	TTL			time.Duration
	Refresh			bool
	Dir			bool
	NoValueOnSuccess	bool
}
type GetOptions struct {
	Recursive	bool
	Sort		bool
	Quorum		bool
}
type DeleteOptions struct {
	PrevValue	string
	PrevIndex	uint64
	Recursive	bool
	Dir		bool
}
type Watcher interface {
	Next(context.Context) (*Response, error)
}
type Response struct {
	Action		string	`json:"action"`
	Node		*Node	`json:"node"`
	PrevNode	*Node	`json:"prevNode"`
	Index		uint64	`json:"-"`
	ClusterID	string	`json:"-"`
}
type Node struct {
	Key		string		`json:"key"`
	Dir		bool		`json:"dir,omitempty"`
	Value		string		`json:"value"`
	Nodes		Nodes		`json:"nodes"`
	CreatedIndex	uint64		`json:"createdIndex"`
	ModifiedIndex	uint64		`json:"modifiedIndex"`
	Expiration	*time.Time	`json:"expiration,omitempty"`
	TTL		int64		`json:"ttl,omitempty"`
}

func (n *Node) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("{Key: %s, CreatedIndex: %d, ModifiedIndex: %d, TTL: %d}", n.Key, n.CreatedIndex, n.ModifiedIndex, n.TTL)
}
func (n *Node) TTLDuration() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Duration(n.TTL) * time.Second
}

type Nodes []*Node

func (ns Nodes) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(ns)
}
func (ns Nodes) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ns[i].Key < ns[j].Key
}
func (ns Nodes) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns[i], ns[j] = ns[j], ns[i]
}

type httpKeysAPI struct {
	client	httpClient
	prefix	string
}

func (k *httpKeysAPI) Set(ctx context.Context, key, val string, opts *SetOptions) (*Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	act := &setAction{Prefix: k.prefix, Key: key, Value: val}
	if opts != nil {
		act.PrevValue = opts.PrevValue
		act.PrevIndex = opts.PrevIndex
		act.PrevExist = opts.PrevExist
		act.TTL = opts.TTL
		act.Refresh = opts.Refresh
		act.Dir = opts.Dir
		act.NoValueOnSuccess = opts.NoValueOnSuccess
	}
	doCtx := ctx
	if act.PrevExist == PrevNoExist {
		doCtx = context.WithValue(doCtx, &oneShotCtxValue, &oneShotCtxValue)
	}
	resp, body, err := k.client.Do(doCtx, act)
	if err != nil {
		return nil, err
	}
	return unmarshalHTTPResponse(resp.StatusCode, resp.Header, body)
}
func (k *httpKeysAPI) Create(ctx context.Context, key, val string) (*Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return k.Set(ctx, key, val, &SetOptions{PrevExist: PrevNoExist})
}
func (k *httpKeysAPI) CreateInOrder(ctx context.Context, dir, val string, opts *CreateInOrderOptions) (*Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	act := &createInOrderAction{Prefix: k.prefix, Dir: dir, Value: val}
	if opts != nil {
		act.TTL = opts.TTL
	}
	resp, body, err := k.client.Do(ctx, act)
	if err != nil {
		return nil, err
	}
	return unmarshalHTTPResponse(resp.StatusCode, resp.Header, body)
}
func (k *httpKeysAPI) Update(ctx context.Context, key, val string) (*Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return k.Set(ctx, key, val, &SetOptions{PrevExist: PrevExist})
}
func (k *httpKeysAPI) Delete(ctx context.Context, key string, opts *DeleteOptions) (*Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	act := &deleteAction{Prefix: k.prefix, Key: key}
	if opts != nil {
		act.PrevValue = opts.PrevValue
		act.PrevIndex = opts.PrevIndex
		act.Dir = opts.Dir
		act.Recursive = opts.Recursive
	}
	doCtx := context.WithValue(ctx, &oneShotCtxValue, &oneShotCtxValue)
	resp, body, err := k.client.Do(doCtx, act)
	if err != nil {
		return nil, err
	}
	return unmarshalHTTPResponse(resp.StatusCode, resp.Header, body)
}
func (k *httpKeysAPI) Get(ctx context.Context, key string, opts *GetOptions) (*Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	act := &getAction{Prefix: k.prefix, Key: key}
	if opts != nil {
		act.Recursive = opts.Recursive
		act.Sorted = opts.Sort
		act.Quorum = opts.Quorum
	}
	resp, body, err := k.client.Do(ctx, act)
	if err != nil {
		return nil, err
	}
	return unmarshalHTTPResponse(resp.StatusCode, resp.Header, body)
}
func (k *httpKeysAPI) Watcher(key string, opts *WatcherOptions) Watcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	act := waitAction{Prefix: k.prefix, Key: key}
	if opts != nil {
		act.Recursive = opts.Recursive
		if opts.AfterIndex > 0 {
			act.WaitIndex = opts.AfterIndex + 1
		}
	}
	return &httpWatcher{client: k.client, nextWait: act}
}

type httpWatcher struct {
	client		httpClient
	nextWait	waitAction
}

func (hw *httpWatcher) Next(ctx context.Context) (*Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		httpresp, body, err := hw.client.Do(ctx, &hw.nextWait)
		if err != nil {
			return nil, err
		}
		resp, err := unmarshalHTTPResponse(httpresp.StatusCode, httpresp.Header, body)
		if err != nil {
			if err == ErrEmptyBody {
				continue
			}
			return nil, err
		}
		hw.nextWait.WaitIndex = resp.Node.ModifiedIndex + 1
		return resp, nil
	}
}
func v2KeysURL(ep url.URL, prefix, key string) *url.URL {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if prefix != "" && prefix[0] != '/' {
		prefix = "/" + prefix
	}
	if key != "" && key[0] != '/' {
		key = "/" + key
	}
	ep.Path = pathutil.CanonicalURLPath(ep.Path + prefix + key)
	return &ep
}

type getAction struct {
	Prefix		string
	Key		string
	Recursive	bool
	Sorted		bool
	Quorum		bool
}

func (g *getAction) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2KeysURL(ep, g.Prefix, g.Key)
	params := u.Query()
	params.Set("recursive", strconv.FormatBool(g.Recursive))
	params.Set("sorted", strconv.FormatBool(g.Sorted))
	params.Set("quorum", strconv.FormatBool(g.Quorum))
	u.RawQuery = params.Encode()
	req, _ := http.NewRequest("GET", u.String(), nil)
	return req
}

type waitAction struct {
	Prefix		string
	Key		string
	WaitIndex	uint64
	Recursive	bool
}

func (w *waitAction) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2KeysURL(ep, w.Prefix, w.Key)
	params := u.Query()
	params.Set("wait", "true")
	params.Set("waitIndex", strconv.FormatUint(w.WaitIndex, 10))
	params.Set("recursive", strconv.FormatBool(w.Recursive))
	u.RawQuery = params.Encode()
	req, _ := http.NewRequest("GET", u.String(), nil)
	return req
}

type setAction struct {
	Prefix			string
	Key			string
	Value			string
	PrevValue		string
	PrevIndex		uint64
	PrevExist		PrevExistType
	TTL			time.Duration
	Refresh			bool
	Dir			bool
	NoValueOnSuccess	bool
}

func (a *setAction) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2KeysURL(ep, a.Prefix, a.Key)
	params := u.Query()
	form := url.Values{}
	if a.Dir {
		params.Set("dir", strconv.FormatBool(a.Dir))
	} else {
		if a.PrevValue != "" {
			params.Set("prevValue", a.PrevValue)
		}
		form.Add("value", a.Value)
	}
	if a.PrevIndex != 0 {
		params.Set("prevIndex", strconv.FormatUint(a.PrevIndex, 10))
	}
	if a.PrevExist != PrevIgnore {
		params.Set("prevExist", string(a.PrevExist))
	}
	if a.TTL > 0 {
		form.Add("ttl", strconv.FormatUint(uint64(a.TTL.Seconds()), 10))
	}
	if a.Refresh {
		form.Add("refresh", "true")
	}
	if a.NoValueOnSuccess {
		params.Set("noValueOnSuccess", strconv.FormatBool(a.NoValueOnSuccess))
	}
	u.RawQuery = params.Encode()
	body := strings.NewReader(form.Encode())
	req, _ := http.NewRequest("PUT", u.String(), body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}

type deleteAction struct {
	Prefix		string
	Key		string
	PrevValue	string
	PrevIndex	uint64
	Dir		bool
	Recursive	bool
}

func (a *deleteAction) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2KeysURL(ep, a.Prefix, a.Key)
	params := u.Query()
	if a.PrevValue != "" {
		params.Set("prevValue", a.PrevValue)
	}
	if a.PrevIndex != 0 {
		params.Set("prevIndex", strconv.FormatUint(a.PrevIndex, 10))
	}
	if a.Dir {
		params.Set("dir", "true")
	}
	if a.Recursive {
		params.Set("recursive", "true")
	}
	u.RawQuery = params.Encode()
	req, _ := http.NewRequest("DELETE", u.String(), nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}

type createInOrderAction struct {
	Prefix	string
	Dir	string
	Value	string
	TTL	time.Duration
}

func (a *createInOrderAction) HTTPRequest(ep url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u := v2KeysURL(ep, a.Prefix, a.Dir)
	form := url.Values{}
	form.Add("value", a.Value)
	if a.TTL > 0 {
		form.Add("ttl", strconv.FormatUint(uint64(a.TTL.Seconds()), 10))
	}
	body := strings.NewReader(form.Encode())
	req, _ := http.NewRequest("POST", u.String(), body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}
func unmarshalHTTPResponse(code int, header http.Header, body []byte) (res *Response, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch code {
	case http.StatusOK, http.StatusCreated:
		if len(body) == 0 {
			return nil, ErrEmptyBody
		}
		res, err = unmarshalSuccessfulKeysResponse(header, body)
	default:
		err = unmarshalFailedKeysResponse(body)
	}
	return res, err
}
func unmarshalSuccessfulKeysResponse(header http.Header, body []byte) (*Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var res Response
	err := codec.NewDecoderBytes(body, new(codec.JsonHandle)).Decode(&res)
	if err != nil {
		return nil, ErrInvalidJSON
	}
	if header.Get("X-Etcd-Index") != "" {
		res.Index, err = strconv.ParseUint(header.Get("X-Etcd-Index"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	res.ClusterID = header.Get("X-Etcd-Cluster-ID")
	return &res, nil
}
func unmarshalFailedKeysResponse(body []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var etcdErr Error
	if err := json.Unmarshal(body, &etcdErr); err != nil {
		return ErrInvalidJSON
	}
	return etcdErr
}
