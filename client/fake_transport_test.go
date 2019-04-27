package client

import (
	"errors"
	"net/http"
)

func (t *fakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case resp := <-t.respchan:
		return resp, nil
	case err := <-t.errchan:
		return nil, err
	case <-t.startCancel:
	case <-req.Cancel:
	}
	select {
	case resp := <-t.respchan:
		return resp, nil
	case <-t.finishCancel:
		return nil, errors.New("cancelled")
	}
}
