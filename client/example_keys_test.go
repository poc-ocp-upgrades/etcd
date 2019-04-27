package client_test

import (
	"context"
	"fmt"
	"log"
	"sort"
	"github.com/coreos/etcd/client"
)

func ExampleKeysAPI_directory() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, err := client.New(client.Config{Endpoints: exampleEndpoints, Transport: exampleTransport})
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	o := client.SetOptions{Dir: true}
	resp, err := kapi.Set(context.Background(), "/myNodes", "", &o)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = kapi.Set(context.Background(), "/myNodes/key1", "value1", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = kapi.Set(context.Background(), "/myNodes/key2", "value2", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = kapi.Get(context.Background(), "/myNodes", nil)
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(resp.Node.Nodes)
	for _, n := range resp.Node.Nodes {
		fmt.Printf("Key: %q, Value: %q\n", n.Key, n.Value)
	}
}
func ExampleKeysAPI_setget() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, err := client.New(client.Config{Endpoints: exampleEndpoints, Transport: exampleTransport})
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)
	resp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err = kapi.Get(context.Background(), "/foo", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
}
