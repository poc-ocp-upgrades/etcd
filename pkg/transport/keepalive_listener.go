package transport

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type keepAliveConn interface {
	SetKeepAlive(bool) error
	SetKeepAlivePeriod(d time.Duration) error
}

func NewKeepAliveListener(l net.Listener, scheme string, tlscfg *tls.Config) (net.Listener, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "https" {
		if tlscfg == nil {
			return nil, fmt.Errorf("cannot listen on TLS for given listener: KeyFile and CertFile are not presented")
		}
		return newTLSKeepaliveListener(l, tlscfg), nil
	}
	return &keepaliveListener{Listener: l}, nil
}

type keepaliveListener struct{ net.Listener }

func (kln *keepaliveListener) Accept() (net.Conn, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, err := kln.Listener.Accept()
	if err != nil {
		return nil, err
	}
	kac := c.(keepAliveConn)
	kac.SetKeepAlive(true)
	kac.SetKeepAlivePeriod(30 * time.Second)
	return c, nil
}

type tlsKeepaliveListener struct {
	net.Listener
	config	*tls.Config
}

func (l *tlsKeepaliveListener) Accept() (c net.Conn, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, err = l.Listener.Accept()
	if err != nil {
		return
	}
	kac := c.(keepAliveConn)
	kac.SetKeepAlive(true)
	kac.SetKeepAlivePeriod(30 * time.Second)
	c = tls.Server(c, l.config)
	return c, nil
}
func newTLSKeepaliveListener(inner net.Listener, config *tls.Config) net.Listener {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := &tlsKeepaliveListener{}
	l.Listener = inner
	l.config = config
	return l
}
