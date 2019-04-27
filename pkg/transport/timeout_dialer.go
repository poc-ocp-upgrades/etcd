package transport

import (
	"net"
	"time"
)

type rwTimeoutDialer struct {
	wtimeoutd	time.Duration
	rdtimeoutd	time.Duration
	net.Dialer
}

func (d *rwTimeoutDialer) Dial(network, address string) (net.Conn, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := d.Dialer.Dial(network, address)
	tconn := &timeoutConn{rdtimeoutd: d.rdtimeoutd, wtimeoutd: d.wtimeoutd, Conn: conn}
	return tconn, err
}
