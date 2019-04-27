package tcpproxy

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"
	"github.com/coreos/pkg/capnslog"
)

var (
	plog = capnslog.NewPackageLogger("github.com/coreos/etcd", "proxy/tcpproxy")
)

type remote struct {
	mu		sync.Mutex
	srv		*net.SRV
	addr		string
	inactive	bool
}

func (r *remote) inactivate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.inactive = true
}
func (r *remote) tryReactivate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := net.Dial("tcp", r.addr)
	if err != nil {
		return err
	}
	conn.Close()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.inactive = false
	return nil
}
func (r *remote) isActive() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.mu.Lock()
	defer r.mu.Unlock()
	return !r.inactive
}

type TCPProxy struct {
	Listener	net.Listener
	Endpoints	[]*net.SRV
	MonitorInterval	time.Duration
	donec		chan struct{}
	mu		sync.Mutex
	remotes		[]*remote
	pickCount	int
}

func (tp *TCPProxy) Run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tp.donec = make(chan struct{})
	if tp.MonitorInterval == 0 {
		tp.MonitorInterval = 5 * time.Minute
	}
	for _, srv := range tp.Endpoints {
		addr := fmt.Sprintf("%s:%d", srv.Target, srv.Port)
		tp.remotes = append(tp.remotes, &remote{srv: srv, addr: addr})
	}
	eps := []string{}
	for _, ep := range tp.Endpoints {
		eps = append(eps, fmt.Sprintf("%s:%d", ep.Target, ep.Port))
	}
	plog.Printf("ready to proxy client requests to %+v", eps)
	go tp.runMonitor()
	for {
		in, err := tp.Listener.Accept()
		if err != nil {
			return err
		}
		go tp.serve(in)
	}
}
func (tp *TCPProxy) pick() *remote {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var weighted []*remote
	var unweighted []*remote
	bestPr := uint16(65535)
	w := 0
	for _, r := range tp.remotes {
		switch {
		case !r.isActive():
		case r.srv.Priority < bestPr:
			bestPr = r.srv.Priority
			w = 0
			weighted = nil
			unweighted = []*remote{r}
			fallthrough
		case r.srv.Priority == bestPr:
			if r.srv.Weight > 0 {
				weighted = append(weighted, r)
				w += int(r.srv.Weight)
			} else {
				unweighted = append(unweighted, r)
			}
		}
	}
	if weighted != nil {
		if len(unweighted) > 0 && rand.Intn(100) == 1 {
			r := unweighted[tp.pickCount%len(unweighted)]
			tp.pickCount++
			return r
		}
		choose := rand.Intn(w)
		for i := 0; i < len(weighted); i++ {
			choose -= int(weighted[i].srv.Weight)
			if choose <= 0 {
				return weighted[i]
			}
		}
	}
	if unweighted != nil {
		for i := 0; i < len(tp.remotes); i++ {
			picked := tp.remotes[tp.pickCount%len(tp.remotes)]
			tp.pickCount++
			if picked.isActive() {
				return picked
			}
		}
	}
	return nil
}
func (tp *TCPProxy) serve(in net.Conn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		err	error
		out	net.Conn
	)
	for {
		tp.mu.Lock()
		remote := tp.pick()
		tp.mu.Unlock()
		if remote == nil {
			break
		}
		out, err = net.Dial("tcp", remote.addr)
		if err == nil {
			break
		}
		remote.inactivate()
		plog.Warningf("deactivated endpoint [%s] due to %v for %v", remote.addr, err, tp.MonitorInterval)
	}
	if out == nil {
		in.Close()
		return
	}
	go func() {
		io.Copy(in, out)
		in.Close()
		out.Close()
	}()
	io.Copy(out, in)
	out.Close()
	in.Close()
}
func (tp *TCPProxy) runMonitor() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		select {
		case <-time.After(tp.MonitorInterval):
			tp.mu.Lock()
			for _, rem := range tp.remotes {
				if rem.isActive() {
					continue
				}
				go func(r *remote) {
					if err := r.tryReactivate(); err != nil {
						plog.Warningf("failed to activate endpoint [%s] due to %v (stay inactive for another %v)", r.addr, err, tp.MonitorInterval)
					} else {
						plog.Printf("activated %s", r.addr)
					}
				}(rem)
			}
			tp.mu.Unlock()
		case <-tp.donec:
			return
		}
	}
}
func (tp *TCPProxy) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tp.Listener.Close()
	close(tp.donec)
}
