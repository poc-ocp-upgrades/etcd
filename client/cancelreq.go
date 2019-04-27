package client

import "net/http"

func requestCanceler(tr CancelableTransport, req *http.Request) func() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ch := make(chan struct{})
	req.Cancel = ch
	return func() {
		close(ch)
	}
}
