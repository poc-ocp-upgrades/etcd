package adapter

import (
	"context"
	"errors"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
)

var errAlreadySentHeader = errors.New("adapter: already sent header")

type ws2wc struct{ wserv pb.WatchServer }

func WatchServerToWatchClient(wserv pb.WatchServer) pb.WatchClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ws2wc{wserv}
}
func (s *ws2wc) Watch(ctx context.Context, opts ...grpc.CallOption) (pb.Watch_WatchClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cs := newPipeStream(ctx, func(ss chanServerStream) error {
		return s.wserv.Watch(&ws2wcServerStream{ss})
	})
	return &ws2wcClientStream{cs}, nil
}

type ws2wcClientStream struct{ chanClientStream }
type ws2wcServerStream struct{ chanServerStream }

func (s *ws2wcClientStream) Send(wr *pb.WatchRequest) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.SendMsg(wr)
}
func (s *ws2wcClientStream) Recv() (*pb.WatchResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var v interface{}
	if err := s.RecvMsg(&v); err != nil {
		return nil, err
	}
	return v.(*pb.WatchResponse), nil
}
func (s *ws2wcServerStream) Send(wr *pb.WatchResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.SendMsg(wr)
}
func (s *ws2wcServerStream) Recv() (*pb.WatchRequest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var v interface{}
	if err := s.RecvMsg(&v); err != nil {
		return nil, err
	}
	return v.(*pb.WatchRequest), nil
}
