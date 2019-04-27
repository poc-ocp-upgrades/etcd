package clientv3_test

import (
	"context"
	"fmt"
	"log"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
)

func ExampleKV_put() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
}
func ExampleKV_putErrorHandling() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Put(ctx, "", "sample_value")
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			fmt.Printf("ctx is canceled by another routine: %v\n", err)
		case context.DeadlineExceeded:
			fmt.Printf("ctx is attached with a deadline is exceeded: %v\n", err)
		case rpctypes.ErrEmptyKey:
			fmt.Printf("client-side error: %v\n", err)
		default:
			fmt.Printf("bad cluster endpoints, which are not etcd servers: %v\n", err)
		}
	}
}
func ExampleKV_get() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	_, err = cli.Put(context.TODO(), "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
func ExampleKV_getWithRev() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	presp, err := cli.Put(context.TODO(), "foo", "bar1")
	if err != nil {
		log.Fatal(err)
	}
	_, err = cli.Put(context.TODO(), "foo", "bar2")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "foo", clientv3.WithRev(presp.Header.Revision))
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
func ExampleKV_getSortedPrefix() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	for i := range make([]int, 3) {
		ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
		_, err = cli.Put(ctx, fmt.Sprintf("key_%d", i), "value")
		cancel()
		if err != nil {
			log.Fatal(err)
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "key", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
func ExampleKV_delete() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	gresp, err := cli.Get(ctx, "key", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	dresp, err := cli.Delete(ctx, "key", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all keys:", int64(len(gresp.Kvs)) == dresp.Deleted)
}
func ExampleKV_compact() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	compRev := resp.Header.Revision
	ctx, cancel = context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Compact(ctx, compRev)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
}
func ExampleKV_txn() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	kvc := clientv3.NewKV(cli)
	_, err = kvc.Put(context.TODO(), "key", "xyz")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = kvc.Txn(ctx).If(clientv3.Compare(clientv3.Value("key"), ">", "abc")).Then(clientv3.OpPut("key", "XYZ")).Else(clientv3.OpPut("key", "ABC")).Commit()
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	gresp, err := kvc.Get(context.TODO(), "key")
	cancel()
	if err != nil {
		log.Fatal(err)
	}
	for _, ev := range gresp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
func ExampleKV_do() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints, DialTimeout: dialTimeout})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	ops := []clientv3.Op{clientv3.OpPut("put-key", "123"), clientv3.OpGet("put-key"), clientv3.OpPut("put-key", "456")}
	for _, op := range ops {
		if _, err := cli.Do(context.TODO(), op); err != nil {
			log.Fatal(err)
		}
	}
}
