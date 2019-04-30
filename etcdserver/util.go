package etcdserver

import (
	"fmt"
	"reflect"
	"strings"
	"time"
	"github.com/golang/protobuf/proto"
	"go.etcd.io/etcd/etcdserver/api/membership"
	"go.etcd.io/etcd/etcdserver/api/rafthttp"
	pb "go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/pkg/types"
	"go.uber.org/zap"
)

func isConnectedToQuorumSince(transport rafthttp.Transporter, since time.Time, self types.ID, members []*membership.Member) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return numConnectedSince(transport, since, self, members) >= (len(members)/2)+1
}
func isConnectedSince(transport rafthttp.Transporter, since time.Time, remote types.ID) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := transport.ActiveSince(remote)
	return !t.IsZero() && t.Before(since)
}
func isConnectedFullySince(transport rafthttp.Transporter, since time.Time, self types.ID, members []*membership.Member) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return numConnectedSince(transport, since, self, members) == len(members)
}
func numConnectedSince(transport rafthttp.Transporter, since time.Time, self types.ID, members []*membership.Member) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	connectedNum := 0
	for _, m := range members {
		if m.ID == self || isConnectedSince(transport, since, m.ID) {
			connectedNum++
		}
	}
	return connectedNum
}
func longestConnected(tp rafthttp.Transporter, membs []types.ID) (types.ID, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var longest types.ID
	var oldest time.Time
	for _, id := range membs {
		tm := tp.ActiveSince(id)
		if tm.IsZero() {
			continue
		}
		if oldest.IsZero() {
			oldest = tm
			longest = id
		}
		if tm.Before(oldest) {
			oldest = tm
			longest = id
		}
	}
	if uint64(longest) == 0 {
		return longest, false
	}
	return longest, true
}

type notifier struct {
	c	chan struct{}
	err	error
}

func newNotifier() *notifier {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &notifier{c: make(chan struct{})}
}
func (nc *notifier) notify(err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nc.err = err
	close(nc.c)
}
func warnOfExpensiveRequest(lg *zap.Logger, now time.Time, reqStringer fmt.Stringer, respMsg proto.Message, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var resp string
	if !isNil(respMsg) {
		resp = fmt.Sprintf("size:%d", proto.Size(respMsg))
	}
	warnOfExpensiveGenericRequest(lg, now, reqStringer, "", resp, err)
}
func warnOfExpensiveReadOnlyTxnRequest(lg *zap.Logger, now time.Time, r *pb.TxnRequest, txnResponse *pb.TxnResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	reqStringer := pb.NewLoggableTxnRequest(r)
	var resp string
	if !isNil(txnResponse) {
		var resps []string
		for _, r := range txnResponse.Responses {
			switch op := r.Response.(type) {
			case *pb.ResponseOp_ResponseRange:
				resps = append(resps, fmt.Sprintf("range_response_count:%d", len(op.ResponseRange.Kvs)))
			default:
			}
		}
		resp = fmt.Sprintf("responses:<%s> size:%d", strings.Join(resps, " "), proto.Size(txnResponse))
	}
	warnOfExpensiveGenericRequest(lg, now, reqStringer, "read-only range ", resp, err)
}
func warnOfExpensiveReadOnlyRangeRequest(lg *zap.Logger, now time.Time, reqStringer fmt.Stringer, rangeResponse *pb.RangeResponse, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var resp string
	if !isNil(rangeResponse) {
		resp = fmt.Sprintf("range_response_count:%d size:%d", len(rangeResponse.Kvs), proto.Size(rangeResponse))
	}
	warnOfExpensiveGenericRequest(lg, now, reqStringer, "read-only range ", resp, err)
}
func warnOfExpensiveGenericRequest(lg *zap.Logger, now time.Time, reqStringer fmt.Stringer, prefix string, resp string, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	d := time.Since(now)
	if d > warnApplyDuration {
		if lg != nil {
			lg.Warn("apply request took too long", zap.Duration("took", d), zap.Duration("expected-duration", warnApplyDuration), zap.String("prefix", prefix), zap.String("request", reqStringer.String()), zap.String("response", resp), zap.Error(err))
		} else {
			var result string
			if err != nil {
				result = fmt.Sprintf("error:%v", err)
			} else {
				result = resp
			}
			plog.Warningf("%srequest %q with result %q took too long (%v) to execute", prefix, reqStringer.String(), result, d)
		}
		slowApplies.Inc()
	}
}
func isNil(msg proto.Message) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return msg == nil || reflect.ValueOf(msg).IsNil()
}
