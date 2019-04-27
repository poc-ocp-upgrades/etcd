package clientv3_test

import (
	"context"
	"fmt"
	"log"
	"github.com/coreos/etcd/clientv3"
)

func ExampleMaintenance_status() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, ep := range endpoints {
		cli, err := clientv3.New(clientv3.Config{Endpoints: []string{ep}, DialTimeout: dialTimeout})
		if err != nil {
			log.Fatal(err)
		}
		defer cli.Close()
		resp, err := cli.Status(context.Background(), ep)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("endpoint: %s / Leader: %v\n", ep, resp.Header.MemberId == resp.Leader)
	}
}
func ExampleMaintenance_defragment() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, ep := range endpoints {
		cli, err := clientv3.New(clientv3.Config{Endpoints: []string{ep}, DialTimeout: dialTimeout})
		if err != nil {
			log.Fatal(err)
		}
		defer cli.Close()
		if _, err = cli.Defragment(context.TODO(), ep); err != nil {
			log.Fatal(err)
		}
	}
}
