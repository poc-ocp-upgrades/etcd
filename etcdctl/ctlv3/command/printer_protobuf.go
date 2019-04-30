package command

import (
	"fmt"
	"os"
	v3 "go.etcd.io/etcd/clientv3"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	mvccpb "go.etcd.io/etcd/mvcc/mvccpb"
)

type pbPrinter struct{ printer }
type pbMarshal interface{ Marshal() ([]byte, error) }

func newPBPrinter() printer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &pbPrinter{&printerRPC{newPrinterUnsupported("protobuf"), printPB}}
}
func (p *pbPrinter) Watch(r v3.WatchResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	evs := make([]*mvccpb.Event, len(r.Events))
	for i, ev := range r.Events {
		evs[i] = (*mvccpb.Event)(ev)
	}
	wr := pb.WatchResponse{Header: &r.Header, Events: evs, CompactRevision: r.CompactRevision, Canceled: r.Canceled, Created: r.Created}
	printPB(&wr)
}
func printPB(v interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m, ok := v.(pbMarshal)
	if !ok {
		ExitWithError(ExitBadFeature, fmt.Errorf("marshal unsupported for type %T (%v)", v, v))
	}
	b, err := m.Marshal()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	fmt.Print(string(b))
}
