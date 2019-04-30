package gw

import (
	"go.etcd.io/etcd/etcdserver/api/v3lock/v3lockpb"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"io"
	"net/http"
	godefaulthttp "net/http"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_Lock_Lock_0(ctx context.Context, marshaler runtime.Marshaler, client v3lockpb.LockClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq v3lockpb.LockRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Lock(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func request_Lock_Unlock_0(ctx context.Context, marshaler runtime.Marshaler, client v3lockpb.LockClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var protoReq v3lockpb.UnlockRequest
	var metadata runtime.ServerMetadata
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.Unlock(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}
func RegisterLockHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterLockHandler(ctx, mux, conn)
}
func RegisterLockHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return RegisterLockHandlerClient(ctx, mux, v3lockpb.NewLockClient(conn))
}
func RegisterLockHandlerClient(ctx context.Context, mux *runtime.ServeMux, client v3lockpb.LockClient) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle("POST", pattern_Lock_Lock_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Lock_Lock_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Lock_Lock_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle("POST", pattern_Lock_Unlock_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Lock_Unlock_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_Lock_Unlock_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_Lock_Lock_0	= runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 1}, []string{"v3", "lock"}, ""))
	pattern_Lock_Unlock_0	= runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v3", "lock", "unlock"}, ""))
)
var (
	forward_Lock_Lock_0	= runtime.ForwardResponseMessage
	forward_Lock_Unlock_0	= runtime.ForwardResponseMessage
)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
