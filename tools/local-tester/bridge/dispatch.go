package main

import (
	"io"
	"math/rand"
	"sync"
	"time"
)

var (
	dispatchPoolDelay	= 100 * time.Millisecond
	dispatchPacketBytes	= 32
)

type dispatcher interface {
	Copy(io.Writer, fetchFunc) error
}
type fetchFunc func() ([]byte, error)
type dispatcherPool struct {
	mu	sync.Mutex
	q	[]dispatchPacket
}
type dispatchPacket struct {
	buf	[]byte
	out	io.Writer
}

func newDispatcherPool() dispatcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := &dispatcherPool{}
	go d.writeLoop()
	return d
}
func (d *dispatcherPool) writeLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		time.Sleep(dispatchPoolDelay)
		d.flush()
	}
}
func (d *dispatcherPool) flush() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d.mu.Lock()
	pkts := d.q
	d.q = nil
	d.mu.Unlock()
	if len(pkts) == 0 {
		return
	}
	pktmap := make(map[io.Writer][]dispatchPacket)
	outs := []io.Writer{}
	for _, pkt := range pkts {
		opkts, ok := pktmap[pkt.out]
		if !ok {
			outs = append(outs, pkt.out)
		}
		pktmap[pkt.out] = append(opkts, pkt)
	}
	for len(outs) != 0 {
		r := rand.Intn(len(outs))
		rpkts := pktmap[outs[r]]
		rpkts[0].out.Write(rpkts[0].buf)
		rpkts = rpkts[1:]
		if len(rpkts) == 0 {
			delete(pktmap, outs[r])
			outs = append(outs[:r], outs[r+1:]...)
		} else {
			pktmap[outs[r]] = rpkts
		}
	}
}
func (d *dispatcherPool) Copy(w io.Writer, f fetchFunc) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		b, err := f()
		if err != nil {
			return err
		}
		pkts := []dispatchPacket{}
		for len(b) > 0 {
			pkt := b
			if len(b) > dispatchPacketBytes {
				pkt = pkt[:dispatchPacketBytes]
				b = b[dispatchPacketBytes:]
			} else {
				b = nil
			}
			pkts = append(pkts, dispatchPacket{pkt, w})
		}
		d.mu.Lock()
		d.q = append(d.q, pkts...)
		d.mu.Unlock()
	}
}

type dispatcherImmediate struct{}

func newDispatcherImmediate() dispatcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &dispatcherImmediate{}
}
func (d *dispatcherImmediate) Copy(w io.Writer, f fetchFunc) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		b, err := f()
		if err != nil {
			return err
		}
		if _, err := w.Write(b); err != nil {
			return err
		}
	}
}
