package grpcproxy

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getAuthTokenFromClient(ctx context.Context) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		ts, ok := md["token"]
		if ok {
			return ts[0]
		}
	}
	return ""
}
func withClientAuthToken(ctx context.Context, ctxWithToken context.Context) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	token := getAuthTokenFromClient(ctxWithToken)
	if token != "" {
		ctx = context.WithValue(ctx, "token", token)
	}
	return ctx
}

type proxyTokenCredential struct{ token string }

func (cred *proxyTokenCredential) RequireTransportSecurity() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (cred *proxyTokenCredential) GetRequestMetadata(ctx context.Context, s ...string) (map[string]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return map[string]string{"token": cred.token}, nil
}
func AuthUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	token := getAuthTokenFromClient(ctx)
	if token != "" {
		tokenCred := &proxyTokenCredential{token}
		opts = append(opts, grpc.PerRPCCredentials(tokenCred))
	}
	return invoker(ctx, method, req, reply, cc, opts...)
}
func AuthStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tokenif := ctx.Value("token")
	if tokenif != nil {
		tokenCred := &proxyTokenCredential{tokenif.(string)}
		opts = append(opts, grpc.PerRPCCredentials(tokenCred))
	}
	return streamer(ctx, desc, cc, method, opts...)
}
