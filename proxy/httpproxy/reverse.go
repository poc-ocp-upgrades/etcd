package httpproxy

import (
	"bytes"
	"context"
	"fmt"
	"github.com/coreos/etcd/etcdserver/api/v2http/httptypes"
	"github.com/coreos/pkg/capnslog"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"
)

var (
	plog             = capnslog.NewPackageLogger("github.com/coreos/etcd", "proxy/httpproxy")
	singleHopHeaders = []string{"Connection", "Keep-Alive", "Proxy-Authenticate", "Proxy-Authorization", "Te", "Trailers", "Transfer-Encoding", "Upgrade"}
)

func removeSingleHopHeaders(hdrs *http.Header) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, h := range singleHopHeaders {
		hdrs.Del(h)
	}
}

type reverseProxy struct {
	director  *director
	transport http.RoundTripper
}

func (p *reverseProxy) ServeHTTP(rw http.ResponseWriter, clientreq *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	reportIncomingRequest(clientreq)
	proxyreq := new(http.Request)
	*proxyreq = *clientreq
	startTime := time.Now()
	var (
		proxybody []byte
		err       error
	)
	if clientreq.Body != nil {
		proxybody, err = ioutil.ReadAll(clientreq.Body)
		if err != nil {
			msg := fmt.Sprintf("failed to read request body: %v", err)
			plog.Println(msg)
			e := httptypes.NewHTTPError(http.StatusInternalServerError, "httpproxy: "+msg)
			if we := e.WriteTo(rw); we != nil {
				plog.Debugf("error writing HTTPError (%v) to %s", we, clientreq.RemoteAddr)
			}
			return
		}
	}
	proxyreq.Header = make(http.Header)
	copyHeader(proxyreq.Header, clientreq.Header)
	normalizeRequest(proxyreq)
	removeSingleHopHeaders(&proxyreq.Header)
	maybeSetForwardedFor(proxyreq)
	endpoints := p.director.endpoints()
	if len(endpoints) == 0 {
		msg := "zero endpoints currently available"
		reportRequestDropped(clientreq, zeroEndpoints)
		plog.Println(msg)
		e := httptypes.NewHTTPError(http.StatusServiceUnavailable, "httpproxy: "+msg)
		if we := e.WriteTo(rw); we != nil {
			plog.Debugf("error writing HTTPError (%v) to %s", we, clientreq.RemoteAddr)
		}
		return
	}
	var requestClosed int32
	completeCh := make(chan bool, 1)
	closeNotifier, ok := rw.(http.CloseNotifier)
	ctx, cancel := context.WithCancel(context.Background())
	proxyreq = proxyreq.WithContext(ctx)
	defer cancel()
	if ok {
		closeCh := closeNotifier.CloseNotify()
		go func() {
			select {
			case <-closeCh:
				atomic.StoreInt32(&requestClosed, 1)
				plog.Printf("client %v closed request prematurely", clientreq.RemoteAddr)
				cancel()
			case <-completeCh:
			}
		}()
		defer func() {
			completeCh <- true
		}()
	}
	var res *http.Response
	for _, ep := range endpoints {
		if proxybody != nil {
			proxyreq.Body = ioutil.NopCloser(bytes.NewBuffer(proxybody))
		}
		redirectRequest(proxyreq, ep.URL)
		res, err = p.transport.RoundTrip(proxyreq)
		if atomic.LoadInt32(&requestClosed) == 1 {
			return
		}
		if err != nil {
			reportRequestDropped(clientreq, failedSendingRequest)
			plog.Printf("failed to direct request to %s: %v", ep.URL.String(), err)
			ep.Failed()
			continue
		}
		break
	}
	if res == nil {
		msg := fmt.Sprintf("unable to get response from %d endpoint(s)", len(endpoints))
		reportRequestDropped(clientreq, failedGettingResponse)
		plog.Println(msg)
		e := httptypes.NewHTTPError(http.StatusBadGateway, "httpproxy: "+msg)
		if we := e.WriteTo(rw); we != nil {
			plog.Debugf("error writing HTTPError (%v) to %s", we, clientreq.RemoteAddr)
		}
		return
	}
	defer res.Body.Close()
	reportRequestHandled(clientreq, res, startTime)
	removeSingleHopHeaders(&res.Header)
	copyHeader(rw.Header(), res.Header)
	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
}
func copyHeader(dst, src http.Header) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
func redirectRequest(req *http.Request, loc url.URL) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req.URL.Scheme = loc.Scheme
	req.URL.Host = loc.Host
}
func normalizeRequest(req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	req.Proto = "HTTP/1.1"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.Close = false
}
func maybeSetForwardedFor(req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clientIP, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return
	}
	if prior, ok := req.Header["X-Forwarded-For"]; ok {
		clientIP = strings.Join(prior, ", ") + ", " + clientIP
	}
	req.Header.Set("X-Forwarded-For", clientIP)
}
