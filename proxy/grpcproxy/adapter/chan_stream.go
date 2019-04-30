package adapter

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type chanServerStream struct {
	headerc		chan<- metadata.MD
	trailerc	chan<- metadata.MD
	grpc.Stream
	headers	[]metadata.MD
}

func (ss *chanServerStream) SendHeader(md metadata.MD) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ss.headerc == nil {
		return errAlreadySentHeader
	}
	outmd := make(map[string][]string)
	for _, h := range append(ss.headers, md) {
		for k, v := range h {
			outmd[k] = v
		}
	}
	select {
	case ss.headerc <- outmd:
		ss.headerc = nil
		ss.headers = nil
		return nil
	case <-ss.Context().Done():
	}
	return ss.Context().Err()
}
func (ss *chanServerStream) SetHeader(md metadata.MD) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ss.headerc == nil {
		return errAlreadySentHeader
	}
	ss.headers = append(ss.headers, md)
	return nil
}
func (ss *chanServerStream) SetTrailer(md metadata.MD) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss.trailerc <- md
}

type chanClientStream struct {
	headerc		<-chan metadata.MD
	trailerc	<-chan metadata.MD
	*chanStream
}

func (cs *chanClientStream) Header() (metadata.MD, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case md := <-cs.headerc:
		return md, nil
	case <-cs.Context().Done():
	}
	return nil, cs.Context().Err()
}
func (cs *chanClientStream) Trailer() metadata.MD {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case md := <-cs.trailerc:
		return md
	case <-cs.Context().Done():
		return nil
	}
}
func (cs *chanClientStream) CloseSend() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(cs.chanStream.sendc)
	return nil
}

type chanStream struct {
	recvc	<-chan interface{}
	sendc	chan<- interface{}
	ctx	context.Context
	cancel	context.CancelFunc
}

func (s *chanStream) Context() context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.ctx
}
func (s *chanStream) SendMsg(m interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case s.sendc <- m:
		if err, ok := m.(error); ok {
			return err
		}
		return nil
	case <-s.ctx.Done():
	}
	return s.ctx.Err()
}
func (s *chanStream) RecvMsg(m interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v := m.(*interface{})
	for {
		select {
		case msg, ok := <-s.recvc:
			if !ok {
				return grpc.ErrClientConnClosing
			}
			if err, ok := msg.(error); ok {
				return err
			}
			*v = msg
			return nil
		case <-s.ctx.Done():
		}
		if len(s.recvc) == 0 {
			break
		}
	}
	return s.ctx.Err()
}
func newPipeStream(ctx context.Context, ssHandler func(chanServerStream) error) chanClientStream {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ch1, ch2 := make(chan interface{}, 1), make(chan interface{})
	headerc, trailerc := make(chan metadata.MD, 1), make(chan metadata.MD, 1)
	cctx, ccancel := context.WithCancel(ctx)
	cli := &chanStream{recvc: ch1, sendc: ch2, ctx: cctx, cancel: ccancel}
	cs := chanClientStream{headerc, trailerc, cli}
	sctx, scancel := context.WithCancel(ctx)
	srv := &chanStream{recvc: ch2, sendc: ch1, ctx: sctx, cancel: scancel}
	ss := chanServerStream{headerc, trailerc, srv, nil}
	go func() {
		if err := ssHandler(ss); err != nil {
			select {
			case srv.sendc <- err:
			case <-sctx.Done():
			case <-cctx.Done():
			}
		}
		scancel()
		ccancel()
	}()
	return cs
}
