package clientv3_test

import (
	"context"
	"fmt"
	"log"
	"github.com/coreos/etcd/clientv3"
)

func ExampleCluster_memberList() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	resp, err := cli.MemberList(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("members:", len(resp.Members))
}
func ExampleCluster_memberAdd() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints[:2], DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	peerURLs := endpoints[2:]
	mresp, err := cli.MemberAdd(context.Background(), peerURLs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("added member.PeerURLs:", mresp.Member.PeerURLs)
}
func ExampleCluster_memberRemove() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints[1:], DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	resp, err := cli.MemberList(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	_, err = cli.MemberRemove(context.Background(), resp.Members[0].ID)
	if err != nil {
		log.Fatal(err)
	}
}
func ExampleCluster_memberUpdate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	resp, err := cli.MemberList(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	peerURLs := []string{"http://localhost:12380"}
	_, err = cli.MemberUpdate(context.Background(), resp.Members[0].ID, peerURLs)
	if err != nil {
		log.Fatal(err)
	}
}
