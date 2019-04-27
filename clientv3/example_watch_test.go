package clientv3_test

import (
	"context"
	"fmt"
	"log"
	"github.com/coreos/etcd/clientv3"
)

func ExampleWatcher_watch() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	rch := cli.Watch(context.Background(), "foo")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
func ExampleWatcher_watchWithPrefix() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	rch := cli.Watch(context.Background(), "foo", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
func ExampleWatcher_watchWithRange() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	rch := cli.Watch(context.Background(), "foo1", clientv3.WithRange("foo4"))
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
func ExampleWatcher_watchWithProgressNotify() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	rch := cli.Watch(context.Background(), "foo", clientv3.WithProgressNotify())
	wresp := <-rch
	fmt.Printf("wresp.Header.Revision: %d\n", wresp.Header.Revision)
	fmt.Println("wresp.IsProgressNotify:", wresp.IsProgressNotify())
}
