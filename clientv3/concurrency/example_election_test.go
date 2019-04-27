package concurrency_test

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

func ExampleElection_Campaign() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	e1 := concurrency.NewElection(s1, "/my-election/")
	s2, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()
	e2 := concurrency.NewElection(s2, "/my-election/")
	var wg sync.WaitGroup
	wg.Add(2)
	electc := make(chan *concurrency.Election, 2)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		if err := e1.Campaign(context.Background(), "e1"); err != nil {
			log.Fatal(err)
		}
		electc <- e1
	}()
	go func() {
		defer wg.Done()
		if err := e2.Campaign(context.Background(), "e2"); err != nil {
			log.Fatal(err)
		}
		electc <- e2
	}()
	cctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	e := <-electc
	fmt.Println("completed first election with", string((<-e.Observe(cctx)).Kvs[0].Value))
	if err := e.Resign(context.TODO()); err != nil {
		log.Fatal(err)
	}
	e = <-electc
	fmt.Println("completed second election with", string((<-e.Observe(cctx)).Kvs[0].Value))
	wg.Wait()
}
