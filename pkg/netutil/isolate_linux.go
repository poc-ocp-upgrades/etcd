package netutil

import (
	"fmt"
	"os/exec"
)

func DropPort(port int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdStr := fmt.Sprintf("sudo iptables -A OUTPUT -p tcp --destination-port %d -j DROP", port)
	if _, err := exec.Command("/bin/sh", "-c", cmdStr).Output(); err != nil {
		return err
	}
	cmdStr = fmt.Sprintf("sudo iptables -A INPUT -p tcp --destination-port %d -j DROP", port)
	_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	return err
}
func RecoverPort(port int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdStr := fmt.Sprintf("sudo iptables -D OUTPUT -p tcp --destination-port %d -j DROP", port)
	if _, err := exec.Command("/bin/sh", "-c", cmdStr).Output(); err != nil {
		return err
	}
	cmdStr = fmt.Sprintf("sudo iptables -D INPUT -p tcp --destination-port %d -j DROP", port)
	_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()
	return err
}
func SetLatency(ms, rv int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ifces, err := GetDefaultInterfaces()
	if err != nil {
		return err
	}
	if rv > ms {
		rv = 1
	}
	for ifce := range ifces {
		cmdStr := fmt.Sprintf("sudo tc qdisc add dev %s root netem delay %dms %dms distribution normal", ifce, ms, rv)
		_, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
		if err != nil {
			cmdStr = fmt.Sprintf("sudo tc qdisc change dev %s root netem delay %dms %dms distribution normal", ifce, ms, rv)
			_, err = exec.Command("/bin/sh", "-c", cmdStr).Output()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func RemoveLatency() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ifces, err := GetDefaultInterfaces()
	if err != nil {
		return err
	}
	for ifce := range ifces {
		_, err = exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo tc qdisc del dev %s root netem", ifce)).Output()
		if err != nil {
			return err
		}
	}
	return nil
}
