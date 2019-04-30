package systemd

import "net"

func DialJournal() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := net.Dial("unixgram", "/run/systemd/journal/socket")
	if conn != nil {
		defer conn.Close()
	}
	return err
}
