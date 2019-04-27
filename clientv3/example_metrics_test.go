package clientv3_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"github.com/coreos/etcd/clientv3"
	grpcprom "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func ExampleClient_metrics() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialOptions: []grpc.DialOption{grpc.WithUnaryInterceptor(grpcprom.UnaryClientInterceptor), grpc.WithStreamInterceptor(grpcprom.StreamClientInterceptor)}})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	cli.Get(context.TODO(), "test_key")
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		http.Serve(ln, promhttp.Handler())
	}()
	defer func() {
		ln.Close()
		<-donec
	}()
	url := "http://" + ln.Addr().String() + "/metrics"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("fetch error: %v", err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalf("fetch error: reading %s: %v", url, err)
	}
	for _, l := range strings.Split(string(b), "\n") {
		if strings.Contains(l, `grpc_client_started_total{grpc_method="Range"`) {
			fmt.Println(l)
			break
		}
	}
}
