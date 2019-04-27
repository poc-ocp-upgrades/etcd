package namespace

import (
	"bytes"
	"context"
	"github.com/coreos/etcd/clientv3"
)

type leasePrefix struct {
	clientv3.Lease
	pfx	[]byte
}

func NewLease(l clientv3.Lease, prefix string) clientv3.Lease {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &leasePrefix{l, []byte(prefix)}
}
func (l *leasePrefix) TimeToLive(ctx context.Context, id clientv3.LeaseID, opts ...clientv3.LeaseOption) (*clientv3.LeaseTimeToLiveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := l.Lease.TimeToLive(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	if len(resp.Keys) > 0 {
		var outKeys [][]byte
		for i := range resp.Keys {
			if len(resp.Keys[i]) < len(l.pfx) {
				continue
			}
			if !bytes.Equal(resp.Keys[i][:len(l.pfx)], l.pfx) {
				continue
			}
			outKeys = append(outKeys, resp.Keys[i][len(l.pfx):])
		}
		resp.Keys = outKeys
	}
	return resp, nil
}
