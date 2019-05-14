package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	cURLDebug = false
)

func EnablecURLDebug() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cURLDebug = true
}
func DisablecURLDebug() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cURLDebug = false
}
func printcURL(req *http.Request) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !cURLDebug {
		return nil
	}
	var (
		command string
		b       []byte
		err     error
	)
	if req.URL != nil {
		command = fmt.Sprintf("curl -X %s %s", req.Method, req.URL.String())
	}
	if req.Body != nil {
		b, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		command += fmt.Sprintf(" -d %q", string(b))
	}
	fmt.Fprintf(os.Stderr, "cURL Command: %s\n", command)
	body := bytes.NewBuffer(b)
	req.Body = ioutil.NopCloser(body)
	return nil
}
