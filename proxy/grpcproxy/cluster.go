package grpcproxy

import (
	"context"
	"fmt"
	"os"
	"sync"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/naming"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"golang.org/x/time/rate"
	gnaming "google.golang.org/grpc/naming"
)

const resolveRetryRate = 1

type clusterProxy struct {
	clus	clientv3.Cluster
	ctx	context.Context
	gr	*naming.GRPCResolver
	advaddr	string
	prefix	string
	umu	sync.RWMutex
	umap	map[string]gnaming.Update
}

func NewClusterProxy(c *clientv3.Client, advaddr string, prefix string) (pb.ClusterServer, <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cp := &clusterProxy{clus: c.Cluster, ctx: c.Ctx(), gr: &naming.GRPCResolver{Client: c}, advaddr: advaddr, prefix: prefix, umap: make(map[string]gnaming.Update)}
	donec := make(chan struct{})
	if advaddr != "" && prefix != "" {
		go func() {
			defer close(donec)
			cp.resolve(prefix)
		}()
		return cp, donec
	}
	close(donec)
	return cp, donec
}
func (cp *clusterProxy) resolve(prefix string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rm := rate.NewLimiter(rate.Limit(resolveRetryRate), resolveRetryRate)
	for rm.Wait(cp.ctx) == nil {
		wa, err := cp.gr.Resolve(prefix)
		if err != nil {
			plog.Warningf("failed to resolve %q (%v)", prefix, err)
			continue
		}
		cp.monitor(wa)
	}
}
func (cp *clusterProxy) monitor(wa gnaming.Watcher) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for cp.ctx.Err() == nil {
		ups, err := wa.Next()
		if err != nil {
			plog.Warningf("clusterProxy watcher error (%v)", err)
			if rpctypes.ErrorDesc(err) == naming.ErrWatcherClosed.Error() {
				return
			}
		}
		cp.umu.Lock()
		for i := range ups {
			switch ups[i].Op {
			case gnaming.Add:
				cp.umap[ups[i].Addr] = *ups[i]
			case gnaming.Delete:
				delete(cp.umap, ups[i].Addr)
			}
		}
		cp.umu.Unlock()
	}
}
func (cp *clusterProxy) MemberAdd(ctx context.Context, r *pb.MemberAddRequest) (*pb.MemberAddResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mresp, err := cp.clus.MemberAdd(ctx, r.PeerURLs)
	if err != nil {
		return nil, err
	}
	resp := (pb.MemberAddResponse)(*mresp)
	return &resp, err
}
func (cp *clusterProxy) MemberRemove(ctx context.Context, r *pb.MemberRemoveRequest) (*pb.MemberRemoveResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mresp, err := cp.clus.MemberRemove(ctx, r.ID)
	if err != nil {
		return nil, err
	}
	resp := (pb.MemberRemoveResponse)(*mresp)
	return &resp, err
}
func (cp *clusterProxy) MemberUpdate(ctx context.Context, r *pb.MemberUpdateRequest) (*pb.MemberUpdateResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mresp, err := cp.clus.MemberUpdate(ctx, r.ID, r.PeerURLs)
	if err != nil {
		return nil, err
	}
	resp := (pb.MemberUpdateResponse)(*mresp)
	return &resp, err
}
func (cp *clusterProxy) membersFromUpdates() ([]*pb.Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cp.umu.RLock()
	defer cp.umu.RUnlock()
	mbs := make([]*pb.Member, 0, len(cp.umap))
	for addr, upt := range cp.umap {
		m, err := decodeMeta(fmt.Sprint(upt.Metadata))
		if err != nil {
			return nil, err
		}
		mbs = append(mbs, &pb.Member{Name: m.Name, ClientURLs: []string{addr}})
	}
	return mbs, nil
}
func (cp *clusterProxy) MemberList(ctx context.Context, r *pb.MemberListRequest) (*pb.MemberListResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if cp.advaddr != "" {
		if cp.prefix != "" {
			mbs, err := cp.membersFromUpdates()
			if err != nil {
				return nil, err
			}
			if len(mbs) > 0 {
				return &pb.MemberListResponse{Members: mbs}, nil
			}
		}
		hostname, _ := os.Hostname()
		return &pb.MemberListResponse{Members: []*pb.Member{{Name: hostname, ClientURLs: []string{cp.advaddr}}}}, nil
	}
	mresp, err := cp.clus.MemberList(ctx)
	if err != nil {
		return nil, err
	}
	resp := (pb.MemberListResponse)(*mresp)
	return &resp, err
}
