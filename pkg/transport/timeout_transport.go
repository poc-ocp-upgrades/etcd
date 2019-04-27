package transport

import (
	"net"
	"net/http"
	"time"
)

func NewTimeoutTransport(info TLSInfo, dialtimeoutd, rdtimeoutd, wtimeoutd time.Duration) (*http.Transport, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tr, err := NewTransport(info, dialtimeoutd)
	if err != nil {
		return nil, err
	}
	if rdtimeoutd != 0 || wtimeoutd != 0 {
		tr.MaxIdleConnsPerHost = -1
	} else {
		tr.MaxIdleConnsPerHost = 1024
	}
	tr.Dial = (&rwTimeoutDialer{Dialer: net.Dialer{Timeout: dialtimeoutd, KeepAlive: 30 * time.Second}, rdtimeoutd: rdtimeoutd, wtimeoutd: wtimeoutd}).Dial
	return tr, nil
}
