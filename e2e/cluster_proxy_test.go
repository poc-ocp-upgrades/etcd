package e2e

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"github.com/coreos/etcd/pkg/expect"
)

type proxyEtcdProcess struct {
	etcdProc	etcdProcess
	proxyV2		*proxyV2Proc
	proxyV3		*proxyV3Proc
}

func newEtcdProcess(cfg *etcdServerProcessConfig) (etcdProcess, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newProxyEtcdProcess(cfg)
}
func newProxyEtcdProcess(cfg *etcdServerProcessConfig) (*proxyEtcdProcess, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ep, err := newEtcdServerProcess(cfg)
	if err != nil {
		return nil, err
	}
	pep := &proxyEtcdProcess{etcdProc: ep, proxyV2: newProxyV2Proc(cfg), proxyV3: newProxyV3Proc(cfg)}
	return pep, nil
}
func (p *proxyEtcdProcess) Config() *etcdServerProcessConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.etcdProc.Config()
}
func (p *proxyEtcdProcess) EndpointsV2() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.proxyV2.endpoints()
}
func (p *proxyEtcdProcess) EndpointsV3() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.proxyV3.endpoints()
}
func (p *proxyEtcdProcess) EndpointsMetrics() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	panic("not implemented; proxy doesn't provide health information")
}
func (p *proxyEtcdProcess) Start() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := p.etcdProc.Start(); err != nil {
		return err
	}
	if err := p.proxyV2.Start(); err != nil {
		return err
	}
	return p.proxyV3.Start()
}
func (p *proxyEtcdProcess) Restart() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := p.etcdProc.Restart(); err != nil {
		return err
	}
	if err := p.proxyV2.Restart(); err != nil {
		return err
	}
	return p.proxyV3.Restart()
}
func (p *proxyEtcdProcess) Stop() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := p.proxyV2.Stop()
	if v3err := p.proxyV3.Stop(); err == nil {
		err = v3err
	}
	if eerr := p.etcdProc.Stop(); eerr != nil && err == nil {
		if !strings.Contains(eerr.Error(), "exit status 2") {
			err = eerr
		}
	}
	return err
}
func (p *proxyEtcdProcess) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := p.proxyV2.Close()
	if v3err := p.proxyV3.Close(); err == nil {
		err = v3err
	}
	if eerr := p.etcdProc.Close(); eerr != nil && err == nil {
		if !strings.Contains(eerr.Error(), "exit status 2") {
			err = eerr
		}
	}
	return err
}
func (p *proxyEtcdProcess) WithStopSignal(sig os.Signal) os.Signal {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p.proxyV3.WithStopSignal(sig)
	p.proxyV3.WithStopSignal(sig)
	return p.etcdProc.WithStopSignal(sig)
}

type proxyProc struct {
	execPath	string
	args		[]string
	ep		string
	murl		string
	donec		chan struct{}
	proc		*expect.ExpectProcess
}

func (pp *proxyProc) endpoints() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []string{pp.ep}
}
func (pp *proxyProc) start() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pp.proc != nil {
		panic("already started")
	}
	proc, err := spawnCmd(append([]string{pp.execPath}, pp.args...))
	if err != nil {
		return err
	}
	pp.proc = proc
	return nil
}
func (pp *proxyProc) waitReady(readyStr string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer close(pp.donec)
	return waitReadyExpectProc(pp.proc, []string{readyStr})
}
func (pp *proxyProc) Stop() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pp.proc == nil {
		return nil
	}
	if err := pp.proc.Stop(); err != nil && !strings.Contains(err.Error(), "exit status 1") {
		return err
	}
	pp.proc = nil
	<-pp.donec
	pp.donec = make(chan struct{})
	return nil
}
func (pp *proxyProc) WithStopSignal(sig os.Signal) os.Signal {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := pp.proc.StopSignal
	pp.proc.StopSignal = sig
	return ret
}
func (pp *proxyProc) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pp.Stop()
}

type proxyV2Proc struct {
	proxyProc
	dataDir	string
}

func proxyListenURL(cfg *etcdServerProcessConfig, portOffset int) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	u, err := url.Parse(cfg.acurl)
	if err != nil {
		panic(err)
	}
	host, port, _ := net.SplitHostPort(u.Host)
	p, _ := strconv.ParseInt(port, 10, 16)
	u.Host = fmt.Sprintf("%s:%d", host, int(p)+portOffset)
	return u.String()
}
func newProxyV2Proc(cfg *etcdServerProcessConfig) *proxyV2Proc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	listenAddr := proxyListenURL(cfg, 2)
	name := fmt.Sprintf("testname-proxy-%p", cfg)
	args := []string{"--name", name, "--proxy", "on", "--listen-client-urls", listenAddr, "--initial-cluster", cfg.name + "=" + cfg.purl.String()}
	return &proxyV2Proc{proxyProc{execPath: cfg.execPath, args: append(args, cfg.tlsArgs...), ep: listenAddr, donec: make(chan struct{})}, name + ".etcd"}
}
func (v2p *proxyV2Proc) Start() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	os.RemoveAll(v2p.dataDir)
	if err := v2p.start(); err != nil {
		return err
	}
	return v2p.waitReady("httpproxy: endpoints found")
}
func (v2p *proxyV2Proc) Restart() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := v2p.Stop(); err != nil {
		return err
	}
	return v2p.Start()
}
func (v2p *proxyV2Proc) Stop() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := v2p.proxyProc.Stop(); err != nil {
		return err
	}
	return os.RemoveAll(v2p.dataDir)
}

type proxyV3Proc struct{ proxyProc }

func newProxyV3Proc(cfg *etcdServerProcessConfig) *proxyV3Proc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	listenAddr := proxyListenURL(cfg, 3)
	args := []string{"grpc-proxy", "start", "--listen-addr", strings.Split(listenAddr, "/")[2], "--endpoints", cfg.acurl, "--advertise-client-url", ""}
	murl := ""
	if cfg.murl != "" {
		murl = proxyListenURL(cfg, 4)
		args = append(args, "--metrics-addr", murl)
	}
	tlsArgs := []string{}
	for i := 0; i < len(cfg.tlsArgs); i++ {
		switch cfg.tlsArgs[i] {
		case "--cert-file":
			tlsArgs = append(tlsArgs, "--cert", cfg.tlsArgs[i+1], "--cert-file", cfg.tlsArgs[i+1])
			i++
		case "--key-file":
			tlsArgs = append(tlsArgs, "--key", cfg.tlsArgs[i+1], "--key-file", cfg.tlsArgs[i+1])
			i++
		case "--ca-file":
			tlsArgs = append(tlsArgs, "--cacert", cfg.tlsArgs[i+1], "--trusted-ca-file", cfg.tlsArgs[i+1])
			i++
		case "--auto-tls":
			tlsArgs = append(tlsArgs, "--auto-tls", "--insecure-skip-tls-verify")
		case "--peer-ca-file", "--peer-cert-file", "--peer-key-file":
			i++
		case "--client-cert-auth", "--peer-auto-tls":
		default:
			tlsArgs = append(tlsArgs, cfg.tlsArgs[i])
		}
	}
	return &proxyV3Proc{proxyProc{execPath: cfg.execPath, args: append(args, tlsArgs...), ep: listenAddr, murl: murl, donec: make(chan struct{})}}
}
func (v3p *proxyV3Proc) Restart() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := v3p.Stop(); err != nil {
		return err
	}
	return v3p.Start()
}
func (v3p *proxyV3Proc) Start() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := v3p.start(); err != nil {
		return err
	}
	return v3p.waitReady("listening for grpc-proxy client requests")
}
