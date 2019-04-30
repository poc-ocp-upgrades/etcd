package mockserver

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sync"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

type MockServer struct {
	ln		net.Listener
	Network		string
	Address		string
	GrpcServer	*grpc.Server
}

func (ms *MockServer) ResolverAddress() resolver.Address {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch ms.Network {
	case "unix":
		return resolver.Address{Addr: fmt.Sprintf("unix://%s", ms.Address)}
	case "tcp":
		return resolver.Address{Addr: ms.Address}
	default:
		panic("illegal network type: " + ms.Network)
	}
}

type MockServers struct {
	mu	sync.RWMutex
	Servers	[]*MockServer
	wg	sync.WaitGroup
}

func StartMockServers(count int) (ms *MockServers, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return StartMockServersOnNetwork(count, "tcp")
}
func StartMockServersOnNetwork(count int, network string) (ms *MockServers, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch network {
	case "tcp":
		return startMockServersTcp(count)
	case "unix":
		return startMockServersUnix(count)
	default:
		return nil, fmt.Errorf("unsupported network type: %s", network)
	}
}
func startMockServersTcp(count int) (ms *MockServers, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	addrs := make([]string, 0, count)
	for i := 0; i < count; i++ {
		addrs = append(addrs, "localhost:0")
	}
	return startMockServers("tcp", addrs)
}
func startMockServersUnix(count int) (ms *MockServers, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir := os.TempDir()
	addrs := make([]string, 0, count)
	for i := 0; i < count; i++ {
		f, err := ioutil.TempFile(dir, "etcd-unix-so-")
		if err != nil {
			return nil, fmt.Errorf("failed to allocate temp file for unix socket: %v", err)
		}
		fn := f.Name()
		err = os.Remove(fn)
		if err != nil {
			return nil, fmt.Errorf("failed to remove temp file before creating unix socket: %v", err)
		}
		addrs = append(addrs, fn)
	}
	return startMockServers("unix", addrs)
}
func startMockServers(network string, addrs []string) (ms *MockServers, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ms = &MockServers{Servers: make([]*MockServer, len(addrs)), wg: sync.WaitGroup{}}
	defer func() {
		if err != nil {
			ms.Stop()
		}
	}()
	for idx, addr := range addrs {
		ln, err := net.Listen(network, addr)
		if err != nil {
			return nil, fmt.Errorf("failed to listen %v", err)
		}
		ms.Servers[idx] = &MockServer{ln: ln, Network: network, Address: ln.Addr().String()}
		ms.StartAt(idx)
	}
	return ms, nil
}
func (ms *MockServers) StartAt(idx int) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if ms.Servers[idx].ln == nil {
		ms.Servers[idx].ln, err = net.Listen(ms.Servers[idx].Network, ms.Servers[idx].Address)
		if err != nil {
			return fmt.Errorf("failed to listen %v", err)
		}
	}
	svr := grpc.NewServer()
	pb.RegisterKVServer(svr, &mockKVServer{})
	ms.Servers[idx].GrpcServer = svr
	ms.wg.Add(1)
	go func(svr *grpc.Server, l net.Listener) {
		svr.Serve(l)
	}(ms.Servers[idx].GrpcServer, ms.Servers[idx].ln)
	return nil
}
func (ms *MockServers) StopAt(idx int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ms.mu.Lock()
	defer ms.mu.Unlock()
	if ms.Servers[idx].ln == nil {
		return
	}
	ms.Servers[idx].GrpcServer.Stop()
	ms.Servers[idx].GrpcServer = nil
	ms.Servers[idx].ln = nil
	ms.wg.Done()
}
func (ms *MockServers) Stop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for idx := range ms.Servers {
		ms.StopAt(idx)
	}
	ms.wg.Wait()
}

type mockKVServer struct{}

func (m *mockKVServer) Range(context.Context, *pb.RangeRequest) (*pb.RangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pb.RangeResponse{}, nil
}
func (m *mockKVServer) Put(context.Context, *pb.PutRequest) (*pb.PutResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pb.PutResponse{}, nil
}
func (m *mockKVServer) DeleteRange(context.Context, *pb.DeleteRangeRequest) (*pb.DeleteRangeResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pb.DeleteRangeResponse{}, nil
}
func (m *mockKVServer) Txn(context.Context, *pb.TxnRequest) (*pb.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pb.TxnResponse{}, nil
}
func (m *mockKVServer) Compact(context.Context, *pb.CompactionRequest) (*pb.CompactionResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pb.CompactionResponse{}, nil
}
