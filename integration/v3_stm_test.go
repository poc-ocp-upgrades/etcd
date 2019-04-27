package integration

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	v3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestSTMConflict(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := NewClusterV3(t, &ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	etcdc := clus.RandClient()
	keys := make([]string, 5)
	for i := 0; i < len(keys); i++ {
		keys[i] = fmt.Sprintf("foo-%d", i)
		if _, err := etcdc.Put(context.TODO(), keys[i], "100"); err != nil {
			t.Fatalf("could not make key (%v)", err)
		}
	}
	errc := make(chan error)
	for i := range keys {
		curEtcdc := clus.RandClient()
		srcKey := keys[i]
		applyf := func(stm concurrency.STM) error {
			src := stm.Get(srcKey)
			dstKey := srcKey
			for dstKey == srcKey {
				dstKey = keys[rand.Intn(len(keys))]
			}
			dst := stm.Get(dstKey)
			srcV, _ := strconv.ParseInt(src, 10, 64)
			dstV, _ := strconv.ParseInt(dst, 10, 64)
			if srcV == 0 {
				return nil
			}
			xfer := int64(rand.Intn(int(srcV)) / 2)
			stm.Put(srcKey, fmt.Sprintf("%d", srcV-xfer))
			stm.Put(dstKey, fmt.Sprintf("%d", dstV+xfer))
			return nil
		}
		go func() {
			iso := concurrency.WithIsolation(concurrency.RepeatableReads)
			_, err := concurrency.NewSTM(curEtcdc, applyf, iso)
			errc <- err
		}()
	}
	for range keys {
		if err := <-errc; err != nil {
			t.Fatalf("apply failed (%v)", err)
		}
	}
	sum := 0
	for _, oldkey := range keys {
		rk, err := etcdc.Get(context.TODO(), oldkey)
		if err != nil {
			t.Fatalf("couldn't fetch key %s (%v)", oldkey, err)
		}
		v, _ := strconv.ParseInt(string(rk.Kvs[0].Value), 10, 64)
		sum += int(v)
	}
	if sum != len(keys)*100 {
		t.Fatalf("bad sum. got %d, expected %d", sum, len(keys)*100)
	}
}
func TestSTMPutNewKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	etcdc := clus.RandClient()
	applyf := func(stm concurrency.STM) error {
		stm.Put("foo", "bar")
		return nil
	}
	iso := concurrency.WithIsolation(concurrency.RepeatableReads)
	if _, err := concurrency.NewSTM(etcdc, applyf, iso); err != nil {
		t.Fatalf("error on stm txn (%v)", err)
	}
	resp, err := etcdc.Get(context.TODO(), "foo")
	if err != nil {
		t.Fatalf("error fetching key (%v)", err)
	}
	if string(resp.Kvs[0].Value) != "bar" {
		t.Fatalf("bad value. got %+v, expected 'bar' value", resp)
	}
}
func TestSTMAbort(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	etcdc := clus.RandClient()
	ctx, cancel := context.WithCancel(context.TODO())
	applyf := func(stm concurrency.STM) error {
		stm.Put("foo", "baz")
		cancel()
		stm.Put("foo", "bap")
		return nil
	}
	iso := concurrency.WithIsolation(concurrency.RepeatableReads)
	sctx := concurrency.WithAbortContext(ctx)
	if _, err := concurrency.NewSTM(etcdc, applyf, iso, sctx); err == nil {
		t.Fatalf("no error on stm txn")
	}
	resp, err := etcdc.Get(context.TODO(), "foo")
	if err != nil {
		t.Fatalf("error fetching key (%v)", err)
	}
	if len(resp.Kvs) != 0 {
		t.Fatalf("bad value. got %+v, expected nothing", resp)
	}
}
func TestSTMSerialize(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := NewClusterV3(t, &ClusterConfig{Size: 3})
	defer clus.Terminate(t)
	etcdc := clus.RandClient()
	keys := make([]string, 5)
	for i := 0; i < len(keys); i++ {
		keys[i] = fmt.Sprintf("foo-%d", i)
	}
	updatec := make(chan struct{})
	go func() {
		defer close(updatec)
		for i := 0; i < 5; i++ {
			s := fmt.Sprintf("%d", i)
			ops := []v3.Op{}
			for _, k := range keys {
				ops = append(ops, v3.OpPut(k, s))
			}
			if _, err := etcdc.Txn(context.TODO()).Then(ops...).Commit(); err != nil {
				t.Fatalf("couldn't put keys (%v)", err)
			}
			updatec <- struct{}{}
		}
	}()
	errc := make(chan error)
	for range updatec {
		curEtcdc := clus.RandClient()
		applyf := func(stm concurrency.STM) error {
			vs := []string{}
			for i := range keys {
				vs = append(vs, stm.Get(keys[i]))
			}
			for i := range vs {
				if vs[0] != vs[i] {
					return fmt.Errorf("got vs[%d] = %v, want %v", i, vs[i], vs[0])
				}
			}
			return nil
		}
		go func() {
			iso := concurrency.WithIsolation(concurrency.Serializable)
			_, err := concurrency.NewSTM(curEtcdc, applyf, iso)
			errc <- err
		}()
	}
	for i := 0; i < 5; i++ {
		if err := <-errc; err != nil {
			t.Error(err)
		}
	}
}
func TestSTMApplyOnConcurrentDeletion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	etcdc := clus.RandClient()
	if _, err := etcdc.Put(context.TODO(), "foo", "bar"); err != nil {
		t.Fatal(err)
	}
	donec, readyc := make(chan struct{}), make(chan struct{})
	go func() {
		<-readyc
		if _, err := etcdc.Delete(context.TODO(), "foo"); err != nil {
			t.Fatal(err)
		}
		close(donec)
	}()
	try := 0
	applyf := func(stm concurrency.STM) error {
		try++
		stm.Get("foo")
		if try == 1 {
			close(readyc)
			<-donec
		}
		stm.Put("foo2", "bar2")
		return nil
	}
	iso := concurrency.WithIsolation(concurrency.RepeatableReads)
	if _, err := concurrency.NewSTM(etcdc, applyf, iso); err != nil {
		t.Fatalf("error on stm txn (%v)", err)
	}
	if try != 2 {
		t.Fatalf("STM apply expected to run twice, got %d", try)
	}
	resp, err := etcdc.Get(context.TODO(), "foo2")
	if err != nil {
		t.Fatalf("error fetching key (%v)", err)
	}
	if string(resp.Kvs[0].Value) != "bar2" {
		t.Fatalf("bad value. got %+v, expected 'bar2' value", resp)
	}
}
func TestSTMSerializableSnapshotPut(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clus := NewClusterV3(t, &ClusterConfig{Size: 1})
	defer clus.Terminate(t)
	cli := clus.Client(0)
	_, err := cli.Put(context.TODO(), "a", "0")
	testutil.AssertNil(t, err)
	tries := 0
	applyf := func(stm concurrency.STM) error {
		if tries > 2 {
			return fmt.Errorf("too many retries")
		}
		tries++
		stm.Get("a")
		stm.Put("b", "1")
		return nil
	}
	iso := concurrency.WithIsolation(concurrency.SerializableSnapshot)
	_, err = concurrency.NewSTM(cli, applyf, iso)
	testutil.AssertNil(t, err)
	_, err = concurrency.NewSTM(cli, applyf, iso)
	testutil.AssertNil(t, err)
	resp, err := cli.Get(context.TODO(), "b")
	testutil.AssertNil(t, err)
	if resp.Kvs[0].Version != 2 {
		t.Fatalf("bad version. got %+v, expected version 2", resp)
	}
}
