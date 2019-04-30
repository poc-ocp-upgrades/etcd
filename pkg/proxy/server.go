package proxy

import (
	"fmt"
	"io"
	mrand "math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
	"go.etcd.io/etcd/pkg/transport"
	humanize "github.com/dustin/go-humanize"
	"go.uber.org/zap"
)

var (
	defaultDialTimeout	= 3 * time.Second
	defaultBufferSize	= 48 * 1024
	defaultRetryInterval	= 10 * time.Millisecond
	defaultLogger		*zap.Logger
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	defaultLogger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

type Server interface {
	From() string
	To() string
	Ready() <-chan struct{}
	Done() <-chan struct{}
	Error() <-chan error
	Close() error
	PauseAccept()
	UnpauseAccept()
	DelayAccept(latency, rv time.Duration)
	UndelayAccept()
	LatencyAccept() time.Duration
	DelayTx(latency, rv time.Duration)
	UndelayTx()
	LatencyTx() time.Duration
	DelayRx(latency, rv time.Duration)
	UndelayRx()
	LatencyRx() time.Duration
	ModifyTx(f func(data []byte) []byte)
	UnmodifyTx()
	ModifyRx(f func(data []byte) []byte)
	UnmodifyRx()
	BlackholeTx()
	UnblackholeTx()
	BlackholeRx()
	UnblackholeRx()
	PauseTx()
	UnpauseTx()
	PauseRx()
	UnpauseRx()
	ResetListener() error
}
type ServerConfig struct {
	Logger		*zap.Logger
	From		url.URL
	To		url.URL
	TLSInfo		transport.TLSInfo
	DialTimeout	time.Duration
	BufferSize	int
	RetryInterval	time.Duration
}
type server struct {
	lg		*zap.Logger
	from		url.URL
	fromPort	int
	to		url.URL
	toPort		int
	tlsInfo		transport.TLSInfo
	dialTimeout	time.Duration
	bufferSize	int
	retryInterval	time.Duration
	readyc		chan struct{}
	donec		chan struct{}
	errc		chan error
	closeOnce	sync.Once
	closeWg		sync.WaitGroup
	listenerMu	sync.RWMutex
	listener	net.Listener
	pauseAcceptMu	sync.Mutex
	pauseAcceptc	chan struct{}
	latencyAcceptMu	sync.RWMutex
	latencyAccept	time.Duration
	modifyTxMu	sync.RWMutex
	modifyTx	func(data []byte) []byte
	modifyRxMu	sync.RWMutex
	modifyRx	func(data []byte) []byte
	pauseTxMu	sync.Mutex
	pauseTxc	chan struct{}
	pauseRxMu	sync.Mutex
	pauseRxc	chan struct{}
	latencyTxMu	sync.RWMutex
	latencyTx	time.Duration
	latencyRxMu	sync.RWMutex
	latencyRx	time.Duration
}

func NewServer(cfg ServerConfig) Server {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := &server{lg: cfg.Logger, from: cfg.From, to: cfg.To, tlsInfo: cfg.TLSInfo, dialTimeout: cfg.DialTimeout, bufferSize: cfg.BufferSize, retryInterval: cfg.RetryInterval, readyc: make(chan struct{}), donec: make(chan struct{}), errc: make(chan error, 16), pauseAcceptc: make(chan struct{}), pauseTxc: make(chan struct{}), pauseRxc: make(chan struct{})}
	_, fromPort, err := net.SplitHostPort(cfg.From.Host)
	if err == nil {
		s.fromPort, _ = strconv.Atoi(fromPort)
	}
	var toPort string
	_, toPort, err = net.SplitHostPort(cfg.To.Host)
	if err == nil {
		s.toPort, _ = strconv.Atoi(toPort)
	}
	if s.dialTimeout == 0 {
		s.dialTimeout = defaultDialTimeout
	}
	if s.bufferSize == 0 {
		s.bufferSize = defaultBufferSize
	}
	if s.retryInterval == 0 {
		s.retryInterval = defaultRetryInterval
	}
	if s.lg == nil {
		s.lg = defaultLogger
	}
	close(s.pauseAcceptc)
	close(s.pauseTxc)
	close(s.pauseRxc)
	if strings.HasPrefix(s.from.Scheme, "http") {
		s.from.Scheme = "tcp"
	}
	if strings.HasPrefix(s.to.Scheme, "http") {
		s.to.Scheme = "tcp"
	}
	addr := fmt.Sprintf(":%d", s.fromPort)
	if s.fromPort == 0 {
		addr = s.from.Host
	}
	var ln net.Listener
	if !s.tlsInfo.Empty() {
		ln, err = transport.NewListener(addr, s.from.Scheme, &s.tlsInfo)
	} else {
		ln, err = net.Listen(s.from.Scheme, addr)
	}
	if err != nil {
		s.errc <- err
		s.Close()
		return s
	}
	s.listener = ln
	s.closeWg.Add(1)
	go s.listenAndServe()
	s.lg.Info("started proxying", zap.String("from", s.From()), zap.String("to", s.To()))
	return s
}
func (s *server) From() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s://%s", s.from.Scheme, s.from.Host)
}
func (s *server) To() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s://%s", s.to.Scheme, s.to.Host)
}
func (s *server) listenAndServe() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer s.closeWg.Done()
	s.lg.Info("proxy is listening on", zap.String("from", s.From()))
	close(s.readyc)
	for {
		s.pauseAcceptMu.Lock()
		pausec := s.pauseAcceptc
		s.pauseAcceptMu.Unlock()
		select {
		case <-pausec:
		case <-s.donec:
			return
		}
		s.latencyAcceptMu.RLock()
		lat := s.latencyAccept
		s.latencyAcceptMu.RUnlock()
		if lat > 0 {
			select {
			case <-time.After(lat):
			case <-s.donec:
				return
			}
		}
		s.listenerMu.RLock()
		ln := s.listener
		s.listenerMu.RUnlock()
		in, err := ln.Accept()
		if err != nil {
			select {
			case s.errc <- err:
				select {
				case <-s.donec:
					return
				default:
				}
			case <-s.donec:
				return
			}
			s.lg.Debug("listener accept error", zap.Error(err))
			if strings.HasSuffix(err.Error(), "use of closed network connection") {
				select {
				case <-time.After(s.retryInterval):
				case <-s.donec:
					return
				}
				s.lg.Debug("listener is closed; retry listening on", zap.String("from", s.From()))
				if err = s.ResetListener(); err != nil {
					select {
					case s.errc <- err:
						select {
						case <-s.donec:
							return
						default:
						}
					case <-s.donec:
						return
					}
					s.lg.Warn("failed to reset listener", zap.Error(err))
				}
			}
			continue
		}
		var out net.Conn
		if !s.tlsInfo.Empty() {
			var tp *http.Transport
			tp, err = transport.NewTransport(s.tlsInfo, s.dialTimeout)
			if err != nil {
				select {
				case s.errc <- err:
					select {
					case <-s.donec:
						return
					default:
					}
				case <-s.donec:
					return
				}
				continue
			}
			out, err = tp.Dial(s.to.Scheme, s.to.Host)
		} else {
			out, err = net.Dial(s.to.Scheme, s.to.Host)
		}
		if err != nil {
			select {
			case s.errc <- err:
				select {
				case <-s.donec:
					return
				default:
				}
			case <-s.donec:
				return
			}
			s.lg.Debug("failed to dial", zap.Error(err))
			continue
		}
		go func() {
			s.transmit(out, in)
			out.Close()
			in.Close()
		}()
		go func() {
			s.receive(in, out)
			in.Close()
			out.Close()
		}()
	}
}
func (s *server) transmit(dst io.Writer, src io.Reader) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.ioCopy(dst, src, proxyTx)
}
func (s *server) receive(dst io.Writer, src io.Reader) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.ioCopy(dst, src, proxyRx)
}

type proxyType uint8

const (
	proxyTx	proxyType	= iota
	proxyRx
)

func (s *server) ioCopy(dst io.Writer, src io.Reader, ptype proxyType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buf := make([]byte, s.bufferSize)
	for {
		nr1, err := src.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			if strings.HasSuffix(err.Error(), "read: connection reset by peer") {
				return
			}
			if strings.HasSuffix(err.Error(), "use of closed network connection") {
				return
			}
			select {
			case s.errc <- err:
				select {
				case <-s.donec:
					return
				default:
				}
			case <-s.donec:
				return
			}
			s.lg.Debug("failed to read", zap.Error(err))
			return
		}
		if nr1 == 0 {
			return
		}
		data := buf[:nr1]
		switch ptype {
		case proxyTx:
			s.modifyTxMu.RLock()
			if s.modifyTx != nil {
				data = s.modifyTx(data)
			}
			s.modifyTxMu.RUnlock()
		case proxyRx:
			s.modifyRxMu.RLock()
			if s.modifyRx != nil {
				data = s.modifyRx(data)
			}
			s.modifyRxMu.RUnlock()
		default:
			panic("unknown proxy type")
		}
		nr2 := len(data)
		switch ptype {
		case proxyTx:
			s.lg.Debug("modified tx", zap.String("data-received", humanize.Bytes(uint64(nr1))), zap.String("data-modified", humanize.Bytes(uint64(nr2))), zap.String("from", s.From()), zap.String("to", s.To()))
		case proxyRx:
			s.lg.Debug("modified rx", zap.String("data-received", humanize.Bytes(uint64(nr1))), zap.String("data-modified", humanize.Bytes(uint64(nr2))), zap.String("from", s.To()), zap.String("to", s.From()))
		default:
			panic("unknown proxy type")
		}
		var pausec chan struct{}
		switch ptype {
		case proxyTx:
			s.pauseTxMu.Lock()
			pausec = s.pauseTxc
			s.pauseTxMu.Unlock()
		case proxyRx:
			s.pauseRxMu.Lock()
			pausec = s.pauseRxc
			s.pauseRxMu.Unlock()
		default:
			panic("unknown proxy type")
		}
		select {
		case <-pausec:
		case <-s.donec:
			return
		}
		if nr2 == 0 {
			continue
		}
		var lat time.Duration
		switch ptype {
		case proxyTx:
			s.latencyTxMu.RLock()
			lat = s.latencyTx
			s.latencyTxMu.RUnlock()
		case proxyRx:
			s.latencyRxMu.RLock()
			lat = s.latencyRx
			s.latencyRxMu.RUnlock()
		default:
			panic("unknown proxy type")
		}
		if lat > 0 {
			select {
			case <-time.After(lat):
			case <-s.donec:
				return
			}
		}
		var nw int
		nw, err = dst.Write(data)
		if err != nil {
			if err == io.EOF {
				return
			}
			select {
			case s.errc <- err:
				select {
				case <-s.donec:
					return
				default:
				}
			case <-s.donec:
				return
			}
			switch ptype {
			case proxyTx:
				s.lg.Debug("write fail on tx", zap.Error(err))
			case proxyRx:
				s.lg.Debug("write fail on rx", zap.Error(err))
			default:
				panic("unknown proxy type")
			}
			return
		}
		if nr2 != nw {
			select {
			case s.errc <- io.ErrShortWrite:
				select {
				case <-s.donec:
					return
				default:
				}
			case <-s.donec:
				return
			}
			switch ptype {
			case proxyTx:
				s.lg.Debug("write fail on tx; read/write bytes are different", zap.Int("read-bytes", nr1), zap.Int("write-bytes", nw), zap.Error(io.ErrShortWrite))
			case proxyRx:
				s.lg.Debug("write fail on rx; read/write bytes are different", zap.Int("read-bytes", nr1), zap.Int("write-bytes", nw), zap.Error(io.ErrShortWrite))
			default:
				panic("unknown proxy type")
			}
			return
		}
		switch ptype {
		case proxyTx:
			s.lg.Debug("transmitted", zap.String("data-size", humanize.Bytes(uint64(nr1))), zap.String("from", s.From()), zap.String("to", s.To()))
		case proxyRx:
			s.lg.Debug("received", zap.String("data-size", humanize.Bytes(uint64(nr1))), zap.String("from", s.To()), zap.String("to", s.From()))
		default:
			panic("unknown proxy type")
		}
	}
}
func (s *server) Ready() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.readyc
}
func (s *server) Done() <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.donec
}
func (s *server) Error() <-chan error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.errc
}
func (s *server) Close() (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.closeOnce.Do(func() {
		close(s.donec)
		s.listenerMu.Lock()
		if s.listener != nil {
			err = s.listener.Close()
			s.lg.Info("closed proxy listener", zap.String("from", s.From()), zap.String("to", s.To()))
		}
		s.lg.Sync()
		s.listenerMu.Unlock()
	})
	s.closeWg.Wait()
	return err
}
func (s *server) PauseAccept() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.pauseAcceptMu.Lock()
	s.pauseAcceptc = make(chan struct{})
	s.pauseAcceptMu.Unlock()
	s.lg.Info("paused accept", zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) UnpauseAccept() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.pauseAcceptMu.Lock()
	select {
	case <-s.pauseAcceptc:
	case <-s.donec:
		s.pauseAcceptMu.Unlock()
		return
	default:
		close(s.pauseAcceptc)
	}
	s.pauseAcceptMu.Unlock()
	s.lg.Info("unpaused accept", zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) DelayAccept(latency, rv time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if latency <= 0 {
		return
	}
	d := computeLatency(latency, rv)
	s.latencyAcceptMu.Lock()
	s.latencyAccept = d
	s.latencyAcceptMu.Unlock()
	s.lg.Info("set accept latency", zap.Duration("latency", d), zap.Duration("given-latency", latency), zap.Duration("given-latency-random-variable", rv), zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) UndelayAccept() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.latencyAcceptMu.Lock()
	d := s.latencyAccept
	s.latencyAccept = 0
	s.latencyAcceptMu.Unlock()
	s.lg.Info("removed accept latency", zap.Duration("latency", d), zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) LatencyAccept() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.latencyAcceptMu.RLock()
	d := s.latencyAccept
	s.latencyAcceptMu.RUnlock()
	return d
}
func (s *server) DelayTx(latency, rv time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if latency <= 0 {
		return
	}
	d := computeLatency(latency, rv)
	s.latencyTxMu.Lock()
	s.latencyTx = d
	s.latencyTxMu.Unlock()
	s.lg.Info("set transmit latency", zap.Duration("latency", d), zap.Duration("given-latency", latency), zap.Duration("given-latency-random-variable", rv), zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) UndelayTx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.latencyTxMu.Lock()
	d := s.latencyTx
	s.latencyTx = 0
	s.latencyTxMu.Unlock()
	s.lg.Info("removed transmit latency", zap.Duration("latency", d), zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) LatencyTx() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.latencyTxMu.RLock()
	d := s.latencyTx
	s.latencyTxMu.RUnlock()
	return d
}
func (s *server) DelayRx(latency, rv time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if latency <= 0 {
		return
	}
	d := computeLatency(latency, rv)
	s.latencyRxMu.Lock()
	s.latencyRx = d
	s.latencyRxMu.Unlock()
	s.lg.Info("set receive latency", zap.Duration("latency", d), zap.Duration("given-latency", latency), zap.Duration("given-latency-random-variable", rv), zap.String("from", s.To()), zap.String("to", s.From()))
}
func (s *server) UndelayRx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.latencyRxMu.Lock()
	d := s.latencyRx
	s.latencyRx = 0
	s.latencyRxMu.Unlock()
	s.lg.Info("removed receive latency", zap.Duration("latency", d), zap.String("from", s.To()), zap.String("to", s.From()))
}
func (s *server) LatencyRx() time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.latencyRxMu.RLock()
	d := s.latencyRx
	s.latencyRxMu.RUnlock()
	return d
}
func computeLatency(lat, rv time.Duration) time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if rv == 0 {
		return lat
	}
	if rv < 0 {
		rv *= -1
	}
	if rv > lat {
		rv = lat / 10
	}
	now := time.Now()
	mrand.Seed(int64(now.Nanosecond()))
	sign := 1
	if now.Second()%2 == 0 {
		sign = -1
	}
	return lat + time.Duration(int64(sign)*mrand.Int63n(rv.Nanoseconds()))
}
func (s *server) ModifyTx(f func([]byte) []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.modifyTxMu.Lock()
	s.modifyTx = f
	s.modifyTxMu.Unlock()
	s.lg.Info("modifying tx", zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) UnmodifyTx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.modifyTxMu.Lock()
	s.modifyTx = nil
	s.modifyTxMu.Unlock()
	s.lg.Info("unmodifyed tx", zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) ModifyRx(f func([]byte) []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.modifyRxMu.Lock()
	s.modifyRx = f
	s.modifyRxMu.Unlock()
	s.lg.Info("modifying rx", zap.String("from", s.To()), zap.String("to", s.From()))
}
func (s *server) UnmodifyRx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.modifyRxMu.Lock()
	s.modifyRx = nil
	s.modifyRxMu.Unlock()
	s.lg.Info("unmodifyed rx", zap.String("from", s.To()), zap.String("to", s.From()))
}
func (s *server) BlackholeTx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.ModifyTx(func([]byte) []byte {
		return nil
	})
	s.lg.Info("blackholed tx", zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) UnblackholeTx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.UnmodifyTx()
	s.lg.Info("unblackholed tx", zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) BlackholeRx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.ModifyRx(func([]byte) []byte {
		return nil
	})
	s.lg.Info("blackholed rx", zap.String("from", s.To()), zap.String("to", s.From()))
}
func (s *server) UnblackholeRx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.UnmodifyRx()
	s.lg.Info("unblackholed rx", zap.String("from", s.To()), zap.String("to", s.From()))
}
func (s *server) PauseTx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.pauseTxMu.Lock()
	s.pauseTxc = make(chan struct{})
	s.pauseTxMu.Unlock()
	s.lg.Info("paused tx", zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) UnpauseTx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.pauseTxMu.Lock()
	select {
	case <-s.pauseTxc:
	case <-s.donec:
		s.pauseTxMu.Unlock()
		return
	default:
		close(s.pauseTxc)
	}
	s.pauseTxMu.Unlock()
	s.lg.Info("unpaused tx", zap.String("from", s.From()), zap.String("to", s.To()))
}
func (s *server) PauseRx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.pauseRxMu.Lock()
	s.pauseRxc = make(chan struct{})
	s.pauseRxMu.Unlock()
	s.lg.Info("paused rx", zap.String("from", s.To()), zap.String("to", s.From()))
}
func (s *server) UnpauseRx() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.pauseRxMu.Lock()
	select {
	case <-s.pauseRxc:
	case <-s.donec:
		s.pauseRxMu.Unlock()
		return
	default:
		close(s.pauseRxc)
	}
	s.pauseRxMu.Unlock()
	s.lg.Info("unpaused rx", zap.String("from", s.To()), zap.String("to", s.From()))
}
func (s *server) ResetListener() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.listenerMu.Lock()
	defer s.listenerMu.Unlock()
	if err := s.listener.Close(); err != nil {
		if !strings.HasSuffix(err.Error(), "use of closed network connection") {
			return err
		}
	}
	var ln net.Listener
	var err error
	if !s.tlsInfo.Empty() {
		ln, err = transport.NewListener(s.from.Host, s.from.Scheme, &s.tlsInfo)
	} else {
		ln, err = net.Listen(s.from.Scheme, s.from.Host)
	}
	if err != nil {
		return err
	}
	s.listener = ln
	s.lg.Info("reset listener on", zap.String("from", s.From()))
	return nil
}
