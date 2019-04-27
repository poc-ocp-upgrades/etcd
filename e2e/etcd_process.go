package e2e

import (
	"fmt"
	"net/url"
	"os"
	"github.com/coreos/etcd/pkg/expect"
	"github.com/coreos/etcd/pkg/fileutil"
)

var (
	etcdServerReadyLines	= []string{"enabled capabilities for version", "published"}
	binPath			string
	ctlBinPath		string
)

type etcdProcess interface {
	EndpointsV2() []string
	EndpointsV3() []string
	EndpointsMetrics() []string
	Start() error
	Restart() error
	Stop() error
	Close() error
	WithStopSignal(sig os.Signal) os.Signal
	Config() *etcdServerProcessConfig
}
type etcdServerProcess struct {
	cfg	*etcdServerProcessConfig
	proc	*expect.ExpectProcess
	donec	chan struct{}
}
type etcdServerProcessConfig struct {
	execPath	string
	args		[]string
	tlsArgs		[]string
	dataDirPath	string
	keepDataDir	bool
	name		string
	purl		url.URL
	acurl		string
	murl		string
	initialToken	string
	initialCluster	string
}

func newEtcdServerProcess(cfg *etcdServerProcessConfig) (*etcdServerProcess, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !fileutil.Exist(cfg.execPath) {
		return nil, fmt.Errorf("could not find etcd binary")
	}
	if !cfg.keepDataDir {
		if err := os.RemoveAll(cfg.dataDirPath); err != nil {
			return nil, err
		}
	}
	return &etcdServerProcess{cfg: cfg, donec: make(chan struct{})}, nil
}
func (ep *etcdServerProcess) EndpointsV2() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{ep.cfg.acurl}
}
func (ep *etcdServerProcess) EndpointsV3() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ep.EndpointsV2()
}
func (ep *etcdServerProcess) EndpointsMetrics() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{ep.cfg.murl}
}
func (ep *etcdServerProcess) Start() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ep.proc != nil {
		panic("already started")
	}
	proc, err := spawnCmd(append([]string{ep.cfg.execPath}, ep.cfg.args...))
	if err != nil {
		return err
	}
	ep.proc = proc
	return ep.waitReady()
}
func (ep *etcdServerProcess) Restart() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ep.Stop(); err != nil {
		return err
	}
	ep.donec = make(chan struct{})
	return ep.Start()
}
func (ep *etcdServerProcess) Stop() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ep == nil || ep.proc == nil {
		return nil
	}
	if err := ep.proc.Stop(); err != nil {
		return err
	}
	ep.proc = nil
	<-ep.donec
	ep.donec = make(chan struct{})
	if ep.cfg.purl.Scheme == "unix" || ep.cfg.purl.Scheme == "unixs" {
		os.Remove(ep.cfg.purl.Host + ep.cfg.purl.Path)
	}
	return nil
}
func (ep *etcdServerProcess) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ep.Stop(); err != nil {
		return err
	}
	return os.RemoveAll(ep.cfg.dataDirPath)
}
func (ep *etcdServerProcess) WithStopSignal(sig os.Signal) os.Signal {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := ep.proc.StopSignal
	ep.proc.StopSignal = sig
	return ret
}
func (ep *etcdServerProcess) waitReady() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(ep.donec)
	return waitReadyExpectProc(ep.proc, etcdServerReadyLines)
}
func (ep *etcdServerProcess) Config() *etcdServerProcessConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ep.cfg
}
