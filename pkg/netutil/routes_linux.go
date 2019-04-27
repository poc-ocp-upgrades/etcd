package netutil

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sort"
	"syscall"
	"github.com/coreos/etcd/pkg/cpuutil"
)

var errNoDefaultRoute = fmt.Errorf("could not find default route")
var errNoDefaultHost = fmt.Errorf("could not find default host")
var errNoDefaultInterface = fmt.Errorf("could not find default interface")

func GetDefaultHost() (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rmsgs, rerr := getDefaultRoutes()
	if rerr != nil {
		return "", rerr
	}
	if rmsg, ok := rmsgs[syscall.AF_INET]; ok {
		if host, err := chooseHost(syscall.AF_INET, rmsg); host != "" || err != nil {
			return host, err
		}
		delete(rmsgs, syscall.AF_INET)
	}
	var families []int
	for family := range rmsgs {
		families = append(families, int(family))
	}
	sort.Ints(families)
	for _, f := range families {
		family := uint8(f)
		if host, err := chooseHost(family, rmsgs[family]); host != "" || err != nil {
			return host, err
		}
	}
	return "", errNoDefaultHost
}
func chooseHost(family uint8, rmsg *syscall.NetlinkMessage) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	host, oif, err := parsePREFSRC(rmsg)
	if host != "" || err != nil {
		return host, err
	}
	ifmsg, ierr := getIfaceAddr(oif, family)
	if ierr != nil {
		return "", ierr
	}
	attrs, aerr := syscall.ParseNetlinkRouteAttr(ifmsg)
	if aerr != nil {
		return "", aerr
	}
	for _, attr := range attrs {
		if attr.Attr.Type == syscall.RTA_DST {
			return net.IP(attr.Value).String(), nil
		}
	}
	return "", nil
}
func getDefaultRoutes() (map[uint8]*syscall.NetlinkMessage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dat, err := syscall.NetlinkRIB(syscall.RTM_GETROUTE, syscall.AF_UNSPEC)
	if err != nil {
		return nil, err
	}
	msgs, msgErr := syscall.ParseNetlinkMessage(dat)
	if msgErr != nil {
		return nil, msgErr
	}
	routes := make(map[uint8]*syscall.NetlinkMessage)
	rtmsg := syscall.RtMsg{}
	for _, m := range msgs {
		if m.Header.Type != syscall.RTM_NEWROUTE {
			continue
		}
		buf := bytes.NewBuffer(m.Data[:syscall.SizeofRtMsg])
		if rerr := binary.Read(buf, cpuutil.ByteOrder(), &rtmsg); rerr != nil {
			continue
		}
		if rtmsg.Dst_len == 0 && rtmsg.Table == syscall.RT_TABLE_MAIN {
			msg := m
			routes[rtmsg.Family] = &msg
		}
	}
	if len(routes) > 0 {
		return routes, nil
	}
	return nil, errNoDefaultRoute
}
func getIfaceAddr(idx uint32, family uint8) (*syscall.NetlinkMessage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dat, err := syscall.NetlinkRIB(syscall.RTM_GETADDR, int(family))
	if err != nil {
		return nil, err
	}
	msgs, msgErr := syscall.ParseNetlinkMessage(dat)
	if msgErr != nil {
		return nil, msgErr
	}
	ifaddrmsg := syscall.IfAddrmsg{}
	for _, m := range msgs {
		if m.Header.Type != syscall.RTM_NEWADDR {
			continue
		}
		buf := bytes.NewBuffer(m.Data[:syscall.SizeofIfAddrmsg])
		if rerr := binary.Read(buf, cpuutil.ByteOrder(), &ifaddrmsg); rerr != nil {
			continue
		}
		if ifaddrmsg.Index == idx {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("could not find address for interface index %v", idx)
}
func getIfaceLink(idx uint32) (*syscall.NetlinkMessage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dat, err := syscall.NetlinkRIB(syscall.RTM_GETLINK, syscall.AF_UNSPEC)
	if err != nil {
		return nil, err
	}
	msgs, msgErr := syscall.ParseNetlinkMessage(dat)
	if msgErr != nil {
		return nil, msgErr
	}
	ifinfomsg := syscall.IfInfomsg{}
	for _, m := range msgs {
		if m.Header.Type != syscall.RTM_NEWLINK {
			continue
		}
		buf := bytes.NewBuffer(m.Data[:syscall.SizeofIfInfomsg])
		if rerr := binary.Read(buf, cpuutil.ByteOrder(), &ifinfomsg); rerr != nil {
			continue
		}
		if ifinfomsg.Index == int32(idx) {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("could not find link for interface index %v", idx)
}
func GetDefaultInterfaces() (map[string]uint8, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	interfaces := make(map[string]uint8)
	rmsgs, rerr := getDefaultRoutes()
	if rerr != nil {
		return interfaces, rerr
	}
	for family, rmsg := range rmsgs {
		_, oif, err := parsePREFSRC(rmsg)
		if err != nil {
			return interfaces, err
		}
		ifmsg, ierr := getIfaceLink(oif)
		if ierr != nil {
			return interfaces, ierr
		}
		attrs, aerr := syscall.ParseNetlinkRouteAttr(ifmsg)
		if aerr != nil {
			return interfaces, aerr
		}
		for _, attr := range attrs {
			if attr.Attr.Type == syscall.IFLA_IFNAME {
				interfaces[string(attr.Value[:len(attr.Value)-1])] += family
			}
		}
	}
	if len(interfaces) > 0 {
		return interfaces, nil
	}
	return interfaces, errNoDefaultInterface
}
func parsePREFSRC(m *syscall.NetlinkMessage) (host string, oif uint32, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var attrs []syscall.NetlinkRouteAttr
	attrs, err = syscall.ParseNetlinkRouteAttr(m)
	if err != nil {
		return "", 0, err
	}
	for _, attr := range attrs {
		if attr.Attr.Type == syscall.RTA_PREFSRC {
			host = net.IP(attr.Value).String()
		}
		if attr.Attr.Type == syscall.RTA_OIF {
			oif = cpuutil.ByteOrder().Uint32(attr.Value)
		}
		if host != "" && oif != uint32(0) {
			break
		}
	}
	if oif == 0 {
		err = errNoDefaultRoute
	}
	return host, oif, err
}
