package transport

import (
	"errors"
	"net"
	"sync"
	"time"
)

var (
	ErrNotTCP = errors.New("only tcp connections have keepalive")
)

func LimitListener(l net.Listener, n int) net.Listener {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &limitListener{l, make(chan struct{}, n)}
}

type limitListener struct {
	net.Listener
	sem chan struct{}
}

func (l *limitListener) acquire() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.sem <- struct{}{}
}
func (l *limitListener) release() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	<-l.sem
}
func (l *limitListener) Accept() (net.Conn, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.acquire()
	c, err := l.Listener.Accept()
	if err != nil {
		l.release()
		return nil, err
	}
	return &limitListenerConn{Conn: c, release: l.release}, nil
}

type limitListenerConn struct {
	net.Conn
	releaseOnce sync.Once
	release     func()
}

func (l *limitListenerConn) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := l.Conn.Close()
	l.releaseOnce.Do(l.release)
	return err
}
func (l *limitListenerConn) SetKeepAlive(doKeepAlive bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tcpc, ok := l.Conn.(*net.TCPConn)
	if !ok {
		return ErrNotTCP
	}
	return tcpc.SetKeepAlive(doKeepAlive)
}
func (l *limitListenerConn) SetKeepAlivePeriod(d time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tcpc, ok := l.Conn.(*net.TCPConn)
	if !ok {
		return ErrNotTCP
	}
	return tcpc.SetKeepAlivePeriod(d)
}
