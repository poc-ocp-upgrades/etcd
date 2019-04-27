package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/coreos/etcd/version"
)

type actionAssertingHTTPClient struct {
	t	*testing.T
	num	int
	act	httpAction
	resp	http.Response
	body	[]byte
	err	error
}

func (a *actionAssertingHTTPClient) Do(_ context.Context, act httpAction) (*http.Response, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !reflect.DeepEqual(a.act, act) {
		a.t.Errorf("#%d: unexpected httpAction: want=%#v got=%#v", a.num, a.act, act)
	}
	return &a.resp, a.body, a.err
}

type staticHTTPClient struct {
	resp	http.Response
	body	[]byte
	err	error
}

func (s *staticHTTPClient) Do(context.Context, httpAction) (*http.Response, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &s.resp, s.body, s.err
}

type staticHTTPAction struct{ request http.Request }

func (s *staticHTTPAction) HTTPRequest(url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &s.request
}

type staticHTTPResponse struct {
	resp	http.Response
	body	[]byte
	err	error
}
type multiStaticHTTPClient struct {
	responses	[]staticHTTPResponse
	cur		int
}

func (s *multiStaticHTTPClient) Do(context.Context, httpAction) (*http.Response, []byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	r := s.responses[s.cur]
	s.cur++
	return &r.resp, r.body, r.err
}
func newStaticHTTPClientFactory(responses []staticHTTPResponse) httpClientFactory {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cur int
	return func(url.URL) httpClient {
		r := responses[cur]
		cur++
		return &staticHTTPClient{resp: r.resp, body: r.body, err: r.err}
	}
}

type fakeTransport struct {
	respchan	chan *http.Response
	errchan		chan error
	startCancel	chan struct{}
	finishCancel	chan struct{}
}

func newFakeTransport() *fakeTransport {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &fakeTransport{respchan: make(chan *http.Response, 1), errchan: make(chan error, 1), startCancel: make(chan struct{}, 1), finishCancel: make(chan struct{}, 1)}
}
func (t *fakeTransport) CancelRequest(*http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.startCancel <- struct{}{}
}

type fakeAction struct{}

func (a *fakeAction) HTTPRequest(url.URL) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &http.Request{}
}
func TestSimpleHTTPClientDoSuccess(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := newFakeTransport()
	c := &simpleHTTPClient{transport: tr}
	tr.respchan <- &http.Response{StatusCode: http.StatusTeapot, Body: ioutil.NopCloser(strings.NewReader("foo"))}
	resp, body, err := c.Do(context.Background(), &fakeAction{})
	if err != nil {
		t.Fatalf("incorrect error value: want=nil got=%v", err)
	}
	wantCode := http.StatusTeapot
	if wantCode != resp.StatusCode {
		t.Fatalf("invalid response code: want=%d got=%d", wantCode, resp.StatusCode)
	}
	wantBody := []byte("foo")
	if !reflect.DeepEqual(wantBody, body) {
		t.Fatalf("invalid response body: want=%q got=%q", wantBody, body)
	}
}
func TestSimpleHTTPClientDoError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := newFakeTransport()
	c := &simpleHTTPClient{transport: tr}
	tr.errchan <- errors.New("fixture")
	_, _, err := c.Do(context.Background(), &fakeAction{})
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}
func TestSimpleHTTPClientDoCancelContext(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := newFakeTransport()
	c := &simpleHTTPClient{transport: tr}
	tr.startCancel <- struct{}{}
	tr.finishCancel <- struct{}{}
	_, _, err := c.Do(context.Background(), &fakeAction{})
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
}

type checkableReadCloser struct {
	io.ReadCloser
	closed	bool
}

func (c *checkableReadCloser) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !c.closed {
		c.closed = true
		return c.ReadCloser.Close()
	}
	return nil
}
func TestSimpleHTTPClientDoCancelContextResponseBodyClosed(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := newFakeTransport()
	c := &simpleHTTPClient{transport: tr}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	body := &checkableReadCloser{ReadCloser: ioutil.NopCloser(strings.NewReader("foo"))}
	go func() {
		testutil.WaitSchedule()
		tr.respchan <- &http.Response{Body: body}
	}()
	_, _, err := c.Do(ctx, &fakeAction{})
	if err == nil {
		t.Fatalf("expected non-nil error, got nil")
	}
	if !body.closed {
		t.Fatalf("expected closed body")
	}
}

type blockingBody struct{ c chan struct{} }

func (bb *blockingBody) Read(p []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	<-bb.c
	return 0, errors.New("closed")
}
func (bb *blockingBody) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(bb.c)
	return nil
}
func TestSimpleHTTPClientDoCancelContextResponseBodyClosedWithBlockingBody(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := newFakeTransport()
	c := &simpleHTTPClient{transport: tr}
	ctx, cancel := context.WithCancel(context.Background())
	body := &checkableReadCloser{ReadCloser: &blockingBody{c: make(chan struct{})}}
	go func() {
		tr.respchan <- &http.Response{Body: body}
		time.Sleep(2 * time.Millisecond)
		cancel()
	}()
	_, _, err := c.Do(ctx, &fakeAction{})
	if err != context.Canceled {
		t.Fatalf("expected %+v, got %+v", context.Canceled, err)
	}
	if !body.closed {
		t.Fatalf("expected closed body")
	}
}
func TestSimpleHTTPClientDoCancelContextWaitForRoundTrip(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := newFakeTransport()
	c := &simpleHTTPClient{transport: tr}
	donechan := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c.Do(ctx, &fakeAction{})
		close(donechan)
	}()
	cancel()
	select {
	case <-donechan:
		t.Fatalf("simpleHTTPClient.Do should not have exited yet")
	default:
	}
	tr.finishCancel <- struct{}{}
	select {
	case <-donechan:
		return
	case <-time.After(time.Second):
		t.Fatalf("simpleHTTPClient.Do did not exit within 1s")
	}
}
func TestSimpleHTTPClientDoHeaderTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr := newFakeTransport()
	tr.finishCancel <- struct{}{}
	c := &simpleHTTPClient{transport: tr, headerTimeout: time.Millisecond}
	errc := make(chan error)
	go func() {
		_, _, err := c.Do(context.Background(), &fakeAction{})
		errc <- err
	}()
	select {
	case err := <-errc:
		if err == nil {
			t.Fatalf("expected non-nil error, got nil")
		}
	case <-time.After(time.Second):
		t.Fatalf("unexpected timeout when waiting for the test to finish")
	}
}
func TestHTTPClusterClientDo(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeErr := errors.New("fake!")
	fakeURL := url.URL{}
	tests := []struct {
		client		*httpClusterClient
		ctx		context.Context
		wantCode	int
		wantErr		error
		wantPinned	int
	}{{client: &httpClusterClient{endpoints: []url.URL{fakeURL, fakeURL}, clientFactory: newStaticHTTPClientFactory([]staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusTeapot}}, {err: fakeErr}}), rand: rand.New(rand.NewSource(0))}, wantCode: http.StatusTeapot}, {client: &httpClusterClient{endpoints: []url.URL{fakeURL, fakeURL}, clientFactory: newStaticHTTPClientFactory([]staticHTTPResponse{{err: fakeErr}, {resp: http.Response{StatusCode: http.StatusTeapot}}}), rand: rand.New(rand.NewSource(0))}, wantCode: http.StatusTeapot, wantPinned: 1}, {client: &httpClusterClient{endpoints: []url.URL{fakeURL, fakeURL}, clientFactory: newStaticHTTPClientFactory([]staticHTTPResponse{{err: context.Canceled}, {resp: http.Response{StatusCode: http.StatusTeapot}}}), rand: rand.New(rand.NewSource(0))}, wantErr: context.Canceled}, {client: &httpClusterClient{endpoints: []url.URL{}, clientFactory: newHTTPClientFactory(nil, nil, 0), rand: rand.New(rand.NewSource(0))}, wantErr: ErrNoEndpoints}, {client: &httpClusterClient{endpoints: []url.URL{fakeURL, fakeURL}, clientFactory: newStaticHTTPClientFactory([]staticHTTPResponse{{err: fakeErr}, {err: fakeErr}}), rand: rand.New(rand.NewSource(0))}, wantErr: &ClusterError{Errors: []error{fakeErr, fakeErr}}}, {client: &httpClusterClient{endpoints: []url.URL{fakeURL, fakeURL}, clientFactory: newStaticHTTPClientFactory([]staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusBadGateway}}, {resp: http.Response{StatusCode: http.StatusTeapot}}}), rand: rand.New(rand.NewSource(0))}, wantCode: http.StatusTeapot, wantPinned: 1}, {client: &httpClusterClient{endpoints: []url.URL{fakeURL, fakeURL}, clientFactory: newStaticHTTPClientFactory([]staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusBadGateway}}, {resp: http.Response{StatusCode: http.StatusTeapot}}}), rand: rand.New(rand.NewSource(0))}, ctx: context.WithValue(context.Background(), &oneShotCtxValue, &oneShotCtxValue), wantErr: fmt.Errorf("client: etcd member  returns server error [Bad Gateway]"), wantPinned: 1}}
	for i, tt := range tests {
		if tt.ctx == nil {
			tt.ctx = context.Background()
		}
		resp, _, err := tt.client.Do(tt.ctx, nil)
		if !reflect.DeepEqual(tt.wantErr, err) {
			t.Errorf("#%d: got err=%v, want=%v", i, err, tt.wantErr)
			continue
		}
		if resp == nil {
			if tt.wantCode != 0 {
				t.Errorf("#%d: resp is nil, want=%d", i, tt.wantCode)
				continue
			}
		} else if resp.StatusCode != tt.wantCode {
			t.Errorf("#%d: resp code=%d, want=%d", i, resp.StatusCode, tt.wantCode)
			continue
		}
		if tt.client.pinned != tt.wantPinned {
			t.Errorf("#%d: pinned=%d, want=%d", i, tt.client.pinned, tt.wantPinned)
		}
	}
}
func TestHTTPClusterClientDoDeadlineExceedContext(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeURL := url.URL{}
	tr := newFakeTransport()
	tr.finishCancel <- struct{}{}
	c := &httpClusterClient{clientFactory: newHTTPClientFactory(tr, DefaultCheckRedirect, 0), endpoints: []url.URL{fakeURL}}
	errc := make(chan error)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		_, _, err := c.Do(ctx, &fakeAction{})
		errc <- err
	}()
	select {
	case err := <-errc:
		if err != context.DeadlineExceeded {
			t.Errorf("err = %+v, want %+v", err, context.DeadlineExceeded)
		}
	case <-time.After(time.Second):
		t.Fatalf("unexpected timeout when waiting for request to deadline exceed")
	}
}

type fakeCancelContext struct{}

var fakeCancelContextError = errors.New("fake context canceled")

func (f fakeCancelContext) Deadline() (time.Time, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Time{}, false
}
func (f fakeCancelContext) Done() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := make(chan struct{}, 1)
	d <- struct{}{}
	return d
}
func (f fakeCancelContext) Err() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fakeCancelContextError
}
func (f fakeCancelContext) Value(key interface{}) interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return 1
}
func withTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return parent, func() {
		parent = nil
	}
}
func TestHTTPClusterClientDoCanceledContext(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeURL := url.URL{}
	tr := newFakeTransport()
	tr.finishCancel <- struct{}{}
	c := &httpClusterClient{clientFactory: newHTTPClientFactory(tr, DefaultCheckRedirect, 0), endpoints: []url.URL{fakeURL}}
	errc := make(chan error)
	go func() {
		ctx, cancel := withTimeout(fakeCancelContext{}, time.Millisecond)
		cancel()
		_, _, err := c.Do(ctx, &fakeAction{})
		errc <- err
	}()
	select {
	case err := <-errc:
		if err != fakeCancelContextError {
			t.Errorf("err = %+v, want %+v", err, fakeCancelContextError)
		}
	case <-time.After(time.Second):
		t.Fatalf("unexpected timeout when waiting for request to fake context canceled")
	}
}
func TestRedirectedHTTPAction(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	act := &redirectedHTTPAction{action: &staticHTTPAction{request: http.Request{Method: "DELETE", URL: &url.URL{Scheme: "https", Host: "foo.example.com", Path: "/ping"}}}, location: url.URL{Scheme: "https", Host: "bar.example.com", Path: "/pong"}}
	want := &http.Request{Method: "DELETE", URL: &url.URL{Scheme: "https", Host: "bar.example.com", Path: "/pong"}}
	got := act.HTTPRequest(url.URL{Scheme: "http", Host: "baz.example.com", Path: "/pang"})
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("HTTPRequest is %#v, want %#v", want, got)
	}
}
func TestRedirectFollowingHTTPClient(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		checkRedirect	CheckRedirectFunc
		client		httpClient
		wantCode	int
		wantErr		error
	}{{checkRedirect: func(int) error {
		return ErrTooManyRedirects
	}, client: &multiStaticHTTPClient{responses: []staticHTTPResponse{{err: errors.New("fail!")}}}, wantErr: errors.New("fail!")}, {checkRedirect: func(int) error {
		return ErrTooManyRedirects
	}, client: &multiStaticHTTPClient{responses: []staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusTeapot}}}}, wantCode: http.StatusTeapot}, {checkRedirect: func(via int) error {
		if via >= 2 {
			return ErrTooManyRedirects
		}
		return nil
	}, client: &multiStaticHTTPClient{responses: []staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusTemporaryRedirect, Header: http.Header{"Location": []string{"http://example.com"}}}}, {resp: http.Response{StatusCode: http.StatusTeapot}}}}, wantCode: http.StatusTeapot}, {checkRedirect: func(via int) error {
		if via >= 3 {
			return ErrTooManyRedirects
		}
		return nil
	}, client: &multiStaticHTTPClient{responses: []staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusTemporaryRedirect, Header: http.Header{"Location": []string{"http://example.com"}}}}, {resp: http.Response{StatusCode: http.StatusTemporaryRedirect, Header: http.Header{"Location": []string{"http://example.com"}}}}, {resp: http.Response{StatusCode: http.StatusTeapot}}}}, wantCode: http.StatusTeapot}, {checkRedirect: func(via int) error {
		if via >= 2 {
			return ErrTooManyRedirects
		}
		return nil
	}, client: &multiStaticHTTPClient{responses: []staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusTemporaryRedirect, Header: http.Header{"Location": []string{"http://example.com"}}}}, {resp: http.Response{StatusCode: http.StatusTemporaryRedirect, Header: http.Header{"Location": []string{"http://example.com"}}}}, {resp: http.Response{StatusCode: http.StatusTeapot}}}}, wantErr: ErrTooManyRedirects}, {checkRedirect: func(int) error {
		return ErrTooManyRedirects
	}, client: &multiStaticHTTPClient{responses: []staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusTemporaryRedirect}}}}, wantErr: errors.New("Location header not set")}, {checkRedirect: func(int) error {
		return ErrTooManyRedirects
	}, client: &multiStaticHTTPClient{responses: []staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusTemporaryRedirect, Header: http.Header{"Location": []string{":"}}}}}}, wantErr: errors.New("Location header not valid URL: :")}, {checkRedirect: func(int) error {
		return nil
	}, client: &staticHTTPClient{resp: http.Response{StatusCode: http.StatusTemporaryRedirect, Header: http.Header{"Location": []string{"http://example.com"}}}}, wantErr: errTooManyRedirectChecks}}
	for i, tt := range tests {
		client := &redirectFollowingHTTPClient{client: tt.client, checkRedirect: tt.checkRedirect}
		resp, _, err := client.Do(context.Background(), nil)
		if !reflect.DeepEqual(tt.wantErr, err) {
			t.Errorf("#%d: got err=%v, want=%v", i, err, tt.wantErr)
			continue
		}
		if resp == nil {
			if tt.wantCode != 0 {
				t.Errorf("#%d: resp is nil, want=%d", i, tt.wantCode)
			}
			continue
		}
		if resp.StatusCode != tt.wantCode {
			t.Errorf("#%d: resp code=%d, want=%d", i, resp.StatusCode, tt.wantCode)
			continue
		}
	}
}
func TestDefaultCheckRedirect(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		num	int
		err	error
	}{{0, nil}, {5, nil}, {10, nil}, {11, ErrTooManyRedirects}, {29, ErrTooManyRedirects}}
	for i, tt := range tests {
		err := DefaultCheckRedirect(tt.num)
		if !reflect.DeepEqual(tt.err, err) {
			t.Errorf("#%d: want=%#v got=%#v", i, tt.err, err)
		}
	}
}
func TestHTTPClusterClientSync(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cf := newStaticHTTPClientFactory([]staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]},{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}})
	hc := &httpClusterClient{clientFactory: cf, rand: rand.New(rand.NewSource(0))}
	err := hc.SetEndpoints([]string{"http://127.0.0.1:2379"})
	if err != nil {
		t.Fatalf("unexpected error during setup: %#v", err)
	}
	want := []string{"http://127.0.0.1:2379"}
	got := hc.Endpoints()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("incorrect endpoints: want=%#v got=%#v", want, got)
	}
	err = hc.Sync(context.Background())
	if err != nil {
		t.Fatalf("unexpected error during Sync: %#v", err)
	}
	want = []string{"http://127.0.0.1:2379", "http://127.0.0.1:4001", "http://127.0.0.1:4002", "http://127.0.0.1:4003"}
	got = hc.Endpoints()
	sort.Sort(sort.StringSlice(got))
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("incorrect endpoints post-Sync: want=%#v got=%#v", want, got)
	}
	err = hc.SetEndpoints([]string{"http://127.0.0.1:4009"})
	if err != nil {
		t.Fatalf("unexpected error during reset: %#v", err)
	}
	want = []string{"http://127.0.0.1:4009"}
	got = hc.Endpoints()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("incorrect endpoints post-reset: want=%#v got=%#v", want, got)
	}
}
func TestHTTPClusterClientSyncFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cf := newStaticHTTPClientFactory([]staticHTTPResponse{{err: errors.New("fail!")}})
	hc := &httpClusterClient{clientFactory: cf, rand: rand.New(rand.NewSource(0))}
	err := hc.SetEndpoints([]string{"http://127.0.0.1:2379"})
	if err != nil {
		t.Fatalf("unexpected error during setup: %#v", err)
	}
	want := []string{"http://127.0.0.1:2379"}
	got := hc.Endpoints()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("incorrect endpoints: want=%#v got=%#v", want, got)
	}
	err = hc.Sync(context.Background())
	if err == nil {
		t.Fatalf("got nil error during Sync")
	}
	got = hc.Endpoints()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("incorrect endpoints after failed Sync: want=%#v got=%#v", want, got)
	}
}
func TestHTTPClusterClientAutoSyncCancelContext(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cf := newStaticHTTPClientFactory([]staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]},{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}})
	hc := &httpClusterClient{clientFactory: cf, rand: rand.New(rand.NewSource(0))}
	err := hc.SetEndpoints([]string{"http://127.0.0.1:2379"})
	if err != nil {
		t.Fatalf("unexpected error during setup: %#v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err = hc.AutoSync(ctx, time.Hour)
	if err != context.Canceled {
		t.Fatalf("incorrect error value: want=%v got=%v", context.Canceled, err)
	}
}
func TestHTTPClusterClientAutoSyncFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cf := newStaticHTTPClientFactory([]staticHTTPResponse{{err: errors.New("fail!")}})
	hc := &httpClusterClient{clientFactory: cf, rand: rand.New(rand.NewSource(0))}
	err := hc.SetEndpoints([]string{"http://127.0.0.1:2379"})
	if err != nil {
		t.Fatalf("unexpected error during setup: %#v", err)
	}
	err = hc.AutoSync(context.Background(), time.Hour)
	if !strings.HasPrefix(err.Error(), ErrClusterUnavailable.Error()) {
		t.Fatalf("incorrect error value: want=%v got=%v", ErrClusterUnavailable, err)
	}
}
func TestHTTPClusterClientGetVersion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	body := []byte(`{"etcdserver":"2.3.2","etcdcluster":"2.3.0"}`)
	cf := newStaticHTTPClientFactory([]staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Length": []string{"44"}}}, body: body}})
	hc := &httpClusterClient{clientFactory: cf, rand: rand.New(rand.NewSource(0))}
	err := hc.SetEndpoints([]string{"http://127.0.0.1:4003", "http://127.0.0.1:2379", "http://127.0.0.1:4001", "http://127.0.0.1:4002"})
	if err != nil {
		t.Fatalf("unexpected error during setup: %#v", err)
	}
	actual, err := hc.GetVersion(context.Background())
	if err != nil {
		t.Errorf("non-nil error: %#v", err)
	}
	expected := version.Versions{Server: "2.3.2", Cluster: "2.3.0"}
	if !reflect.DeepEqual(&expected, actual) {
		t.Errorf("incorrect Response: want=%#v got=%#v", expected, actual)
	}
}
func TestHTTPClusterClientSyncPinEndpoint(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cf := newStaticHTTPClientFactory([]staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]},{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}, {resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]},{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}, {resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]},{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}})
	hc := &httpClusterClient{clientFactory: cf, rand: rand.New(rand.NewSource(0))}
	err := hc.SetEndpoints([]string{"http://127.0.0.1:4003", "http://127.0.0.1:2379", "http://127.0.0.1:4001", "http://127.0.0.1:4002"})
	if err != nil {
		t.Fatalf("unexpected error during setup: %#v", err)
	}
	pinnedEndpoint := hc.endpoints[hc.pinned]
	for i := 0; i < 3; i++ {
		err = hc.Sync(context.Background())
		if err != nil {
			t.Fatalf("#%d: unexpected error during Sync: %#v", i, err)
		}
		if g := hc.endpoints[hc.pinned]; g != pinnedEndpoint {
			t.Errorf("#%d: pinned endpoint = %v, want %v", i, g, pinnedEndpoint)
		}
	}
}
func TestHTTPClusterClientSyncUnpinEndpoint(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cf := newStaticHTTPClientFactory([]staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]},{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}, {resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}, {resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]},{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}})
	hc := &httpClusterClient{clientFactory: cf, rand: rand.New(rand.NewSource(0))}
	err := hc.SetEndpoints([]string{"http://127.0.0.1:4003", "http://127.0.0.1:2379", "http://127.0.0.1:4001", "http://127.0.0.1:4002"})
	if err != nil {
		t.Fatalf("unexpected error during setup: %#v", err)
	}
	wants := []string{"http://127.0.0.1:2379", "http://127.0.0.1:4001", "http://127.0.0.1:4002"}
	for i := 0; i < 3; i++ {
		err = hc.Sync(context.Background())
		if err != nil {
			t.Fatalf("#%d: unexpected error during Sync: %#v", i, err)
		}
		if g := hc.endpoints[hc.pinned]; g.String() != wants[i] {
			t.Errorf("#%d: pinned endpoint = %v, want %v", i, g, wants[i])
		}
	}
}
func TestHTTPClusterClientSyncPinLeaderEndpoint(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cf := newStaticHTTPClientFactory([]staticHTTPResponse{{resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]},{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}, {resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]}`)}, {resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"members":[{"id":"2745e2525fce8fe","peerURLs":["http://127.0.0.1:7003"],"name":"node3","clientURLs":["http://127.0.0.1:4003"]},{"id":"42134f434382925","peerURLs":["http://127.0.0.1:2380","http://127.0.0.1:7001"],"name":"node1","clientURLs":["http://127.0.0.1:2379","http://127.0.0.1:4001"]},{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}]}`)}, {resp: http.Response{StatusCode: http.StatusOK, Header: http.Header{"Content-Type": []string{"application/json"}}}, body: []byte(`{"id":"94088180e21eb87b","peerURLs":["http://127.0.0.1:7002"],"name":"node2","clientURLs":["http://127.0.0.1:4002"]}`)}})
	hc := &httpClusterClient{clientFactory: cf, rand: rand.New(rand.NewSource(0)), selectionMode: EndpointSelectionPrioritizeLeader, endpoints: []url.URL{{}}}
	wants := []string{"http://127.0.0.1:4003", "http://127.0.0.1:4002"}
	for i, want := range wants {
		err := hc.Sync(context.Background())
		if err != nil {
			t.Fatalf("#%d: unexpected error during Sync: %#v", i, err)
		}
		pinned := hc.endpoints[hc.pinned].String()
		if pinned != want {
			t.Errorf("#%d: pinned endpoint = %v, want %v", i, pinned, want)
		}
	}
}
func TestHTTPClusterClientResetFail(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := [][]string{{}, {":"}}
	for i, tt := range tests {
		hc := &httpClusterClient{rand: rand.New(rand.NewSource(0))}
		err := hc.SetEndpoints(tt)
		if err == nil {
			t.Errorf("#%d: expected non-nil error", i)
		}
	}
}
func TestHTTPClusterClientResetPinRandom(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	round := 2000
	pinNum := 0
	for i := 0; i < round; i++ {
		hc := &httpClusterClient{rand: rand.New(rand.NewSource(int64(i)))}
		err := hc.SetEndpoints([]string{"http://127.0.0.1:4001", "http://127.0.0.1:4002", "http://127.0.0.1:4003"})
		if err != nil {
			t.Fatalf("#%d: reset error (%v)", i, err)
		}
		if hc.endpoints[hc.pinned].String() == "http://127.0.0.1:4001" {
			pinNum++
		}
	}
	min := 1.0/3.0 - 0.05
	max := 1.0/3.0 + 0.05
	if ratio := float64(pinNum) / float64(round); ratio > max || ratio < min {
		t.Errorf("pinned ratio = %v, want [%v, %v]", ratio, min, max)
	}
}
