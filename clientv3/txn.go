package clientv3

import (
	"context"
	"sync"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"google.golang.org/grpc"
)

type Txn interface {
	If(cs ...Cmp) Txn
	Then(ops ...Op) Txn
	Else(ops ...Op) Txn
	Commit() (*TxnResponse, error)
}
type txn struct {
	kv		*kv
	ctx		context.Context
	mu		sync.Mutex
	cif		bool
	cthen		bool
	celse		bool
	isWrite		bool
	cmps		[]*pb.Compare
	sus		[]*pb.RequestOp
	fas		[]*pb.RequestOp
	callOpts	[]grpc.CallOption
}

func (txn *txn) If(cs ...Cmp) Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.mu.Lock()
	defer txn.mu.Unlock()
	if txn.cif {
		panic("cannot call If twice!")
	}
	if txn.cthen {
		panic("cannot call If after Then!")
	}
	if txn.celse {
		panic("cannot call If after Else!")
	}
	txn.cif = true
	for i := range cs {
		txn.cmps = append(txn.cmps, (*pb.Compare)(&cs[i]))
	}
	return txn
}
func (txn *txn) Then(ops ...Op) Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.mu.Lock()
	defer txn.mu.Unlock()
	if txn.cthen {
		panic("cannot call Then twice!")
	}
	if txn.celse {
		panic("cannot call Then after Else!")
	}
	txn.cthen = true
	for _, op := range ops {
		txn.isWrite = txn.isWrite || op.isWrite()
		txn.sus = append(txn.sus, op.toRequestOp())
	}
	return txn
}
func (txn *txn) Else(ops ...Op) Txn {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.mu.Lock()
	defer txn.mu.Unlock()
	if txn.celse {
		panic("cannot call Else twice!")
	}
	txn.celse = true
	for _, op := range ops {
		txn.isWrite = txn.isWrite || op.isWrite()
		txn.fas = append(txn.fas, op.toRequestOp())
	}
	return txn
}
func (txn *txn) Commit() (*TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	txn.mu.Lock()
	defer txn.mu.Unlock()
	r := &pb.TxnRequest{Compare: txn.cmps, Success: txn.sus, Failure: txn.fas}
	var resp *pb.TxnResponse
	var err error
	resp, err = txn.kv.remote.Txn(txn.ctx, r, txn.callOpts...)
	if err != nil {
		return nil, toErr(txn.ctx, err)
	}
	return (*TxnResponse)(resp), nil
}
