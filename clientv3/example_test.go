package clientv3_test

import (
	"context"
	"log"
	"os"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/transport"
	"google.golang.org/grpc/grpclog"
)

var (
	dialTimeout	= 5 * time.Second
	requestTimeout	= 10 * time.Second
	endpoints	= []string{"localhost:2379", "localhost:22379", "localhost:32379"}
)

func Example() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	_, err = cli.Put(context.TODO(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
}
func ExampleConfig_withTLS() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tlsInfo := transport.TLSInfo{CertFile: "/tmp/test-certs/test-name-1.pem", KeyFile: "/tmp/test-certs/test-name-1-key.pem", TrustedCAFile: "/tmp/test-certs/trusted-ca.pem"}
	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		log.Fatal(err)
	}
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout, TLS: tlsConfig})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	_, err = cli.Put(context.TODO(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
}
