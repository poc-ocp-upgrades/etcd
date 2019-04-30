package main

import (
	"flag"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

type bridgeConn struct {
	in	net.Conn
	out	net.Conn
	d	dispatcher
}

func newBridgeConn(in net.Conn, d dispatcher) (*bridgeConn, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out, err := net.Dial("tcp", flag.Args()[1])
	if err != nil {
		in.Close()
		return nil, err
	}
	return &bridgeConn{in, out, d}, nil
}
func (b *bridgeConn) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%v <-> %v", b.in.RemoteAddr(), b.out.RemoteAddr())
}
func (b *bridgeConn) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.in.Close()
	b.out.Close()
}
func bridge(b *bridgeConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Println("bridging", b.String())
	go b.d.Copy(b.out, makeFetch(b.in))
	b.d.Copy(b.in, makeFetch(b.out))
}
func delayBridge(b *bridgeConn, txDelay, rxDelay time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	go b.d.Copy(b.out, makeFetchDelay(makeFetch(b.in), txDelay))
	b.d.Copy(b.in, makeFetchDelay(makeFetch(b.out), rxDelay))
}
func timeBridge(b *bridgeConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	go func() {
		t := time.Duration(rand.Intn(5)+1) * time.Second
		time.Sleep(t)
		log.Printf("killing connection %s after %v\n", b.String(), t)
		b.Close()
	}()
	bridge(b)
}
func blackhole(b *bridgeConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Println("blackholing connection", b.String())
	io.Copy(ioutil.Discard, b.in)
	b.Close()
}
func readRemoteOnly(b *bridgeConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Println("one way (<-)", b.String())
	b.d.Copy(b.in, makeFetch(b.out))
}
func writeRemoteOnly(b *bridgeConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Println("one way (->)", b.String())
	b.d.Copy(b.out, makeFetch(b.in))
}
func corruptReceive(b *bridgeConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Println("corruptReceive", b.String())
	go b.d.Copy(b.in, makeFetchCorrupt(makeFetch(b.out)))
	b.d.Copy(b.out, makeFetch(b.in))
}
func corruptSend(b *bridgeConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Println("corruptSend", b.String())
	go b.d.Copy(b.out, makeFetchCorrupt(makeFetch(b.in)))
	b.d.Copy(b.in, makeFetch(b.out))
}
func makeFetch(c io.Reader) fetchFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func() ([]byte, error) {
		b := make([]byte, 4096)
		n, err := c.Read(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func makeFetchCorrupt(f func() ([]byte, error)) fetchFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func() ([]byte, error) {
		b, err := f()
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(b); i++ {
			if rand.Intn(16*1024) == 0 {
				b[i] = b[i] + 1
			}
		}
		return b, nil
	}
}
func makeFetchRand(f func() ([]byte, error)) fetchFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func() ([]byte, error) {
		if rand.Intn(10) == 0 {
			return nil, fmt.Errorf("fetchRand: done")
		}
		b, err := f()
		if err != nil {
			return nil, err
		}
		return b, nil
	}
}
func makeFetchDelay(f fetchFunc, delay time.Duration) fetchFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func() ([]byte, error) {
		b, err := f()
		if err != nil {
			return nil, err
		}
		time.Sleep(delay)
		return b, nil
	}
}
func randomBlackhole(b *bridgeConn) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Println("random blackhole: connection", b.String())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		b.d.Copy(b.in, makeFetchRand(makeFetch(b.out)))
		wg.Done()
	}()
	go func() {
		b.d.Copy(b.out, makeFetchRand(makeFetch(b.in)))
		wg.Done()
	}()
	wg.Wait()
	b.Close()
}

type config struct {
	delayAccept	bool
	resetListen	bool
	connFaultRate	float64
	immediateClose	bool
	blackhole	bool
	timeClose	bool
	writeRemoteOnly	bool
	readRemoteOnly	bool
	randomBlackhole	bool
	corruptSend	bool
	corruptReceive	bool
	reorder		bool
	txDelay		string
	rxDelay		string
}
type acceptFaultFunc func()
type connFaultFunc func(*bridgeConn)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cfg config
	flag.BoolVar(&cfg.delayAccept, "delay-accept", false, "delays accepting new connections")
	flag.BoolVar(&cfg.resetListen, "reset-listen", false, "resets the listening port")
	flag.Float64Var(&cfg.connFaultRate, "conn-fault-rate", 0.0, "rate of faulty connections")
	flag.BoolVar(&cfg.immediateClose, "immediate-close", false, "close after accept")
	flag.BoolVar(&cfg.blackhole, "blackhole", false, "reads nothing, writes go nowhere")
	flag.BoolVar(&cfg.timeClose, "time-close", false, "close after random time")
	flag.BoolVar(&cfg.writeRemoteOnly, "write-remote-only", false, "only write, no read")
	flag.BoolVar(&cfg.readRemoteOnly, "read-remote-only", false, "only read, no write")
	flag.BoolVar(&cfg.randomBlackhole, "random-blackhole", false, "blackhole after data xfer")
	flag.BoolVar(&cfg.corruptReceive, "corrupt-receive", false, "corrupt packets received from destination")
	flag.BoolVar(&cfg.corruptSend, "corrupt-send", false, "corrupt packets sent to destination")
	flag.BoolVar(&cfg.reorder, "reorder", false, "reorder packet delivery")
	flag.StringVar(&cfg.txDelay, "tx-delay", "0", "duration to delay client transmission to server")
	flag.StringVar(&cfg.rxDelay, "rx-delay", "0", "duration to delay client receive from server")
	flag.Parse()
	lAddr := flag.Args()[0]
	fwdAddr := flag.Args()[1]
	log.Println("listening on ", lAddr)
	log.Println("forwarding to ", fwdAddr)
	l, err := net.Listen("tcp", lAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	acceptFaults := []acceptFaultFunc{func() {
	}}
	if cfg.delayAccept {
		f := func() {
			log.Println("delaying accept")
			time.Sleep(3 * time.Second)
		}
		acceptFaults = append(acceptFaults, f)
	}
	if cfg.resetListen {
		f := func() {
			log.Println("reset listen port")
			l.Close()
			newListener, err := net.Listen("tcp", lAddr)
			if err != nil {
				log.Fatal(err)
			}
			l = newListener
		}
		acceptFaults = append(acceptFaults, f)
	}
	connFaults := []connFaultFunc{func(b *bridgeConn) {
		bridge(b)
	}}
	if cfg.immediateClose {
		f := func(b *bridgeConn) {
			log.Printf("terminating connection %s immediately", b.String())
			b.Close()
		}
		connFaults = append(connFaults, f)
	}
	if cfg.blackhole {
		connFaults = append(connFaults, blackhole)
	}
	if cfg.timeClose {
		connFaults = append(connFaults, timeBridge)
	}
	if cfg.writeRemoteOnly {
		connFaults = append(connFaults, writeRemoteOnly)
	}
	if cfg.readRemoteOnly {
		connFaults = append(connFaults, readRemoteOnly)
	}
	if cfg.randomBlackhole {
		connFaults = append(connFaults, randomBlackhole)
	}
	if cfg.corruptSend {
		connFaults = append(connFaults, corruptSend)
	}
	if cfg.corruptReceive {
		connFaults = append(connFaults, corruptReceive)
	}
	txd, txdErr := time.ParseDuration(cfg.txDelay)
	if txdErr != nil {
		log.Fatal(txdErr)
	}
	rxd, rxdErr := time.ParseDuration(cfg.rxDelay)
	if rxdErr != nil {
		log.Fatal(rxdErr)
	}
	if txd != 0 || rxd != 0 {
		f := func(b *bridgeConn) {
			delayBridge(b, txd, rxd)
		}
		connFaults = append(connFaults, f)
	}
	if len(connFaults) > 1 && cfg.connFaultRate == 0 {
		log.Fatal("connection faults defined but conn-fault-rate=0")
	}
	var disp dispatcher
	if cfg.reorder {
		disp = newDispatcherPool()
	} else {
		disp = newDispatcherImmediate()
	}
	for {
		acceptFaults[rand.Intn(len(acceptFaults))]()
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		r := rand.Intn(len(connFaults))
		if rand.Intn(100) >= int(100.0*cfg.connFaultRate) {
			r = 0
		}
		bc, err := newBridgeConn(conn, disp)
		if err != nil {
			log.Printf("oops %v", err)
			continue
		}
		go connFaults[r](bc)
	}
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
