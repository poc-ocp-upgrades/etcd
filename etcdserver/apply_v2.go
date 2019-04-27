package etcdserver

import (
	"encoding/json"
	"path"
	"time"
	"github.com/coreos/etcd/etcdserver/api"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/pbutil"
	"github.com/coreos/etcd/store"
	"github.com/coreos/go-semver/semver"
)

type ApplierV2 interface {
	Delete(r *RequestV2) Response
	Post(r *RequestV2) Response
	Put(r *RequestV2) Response
	QGet(r *RequestV2) Response
	Sync(r *RequestV2) Response
}

func NewApplierV2(s store.Store, c *membership.RaftCluster) ApplierV2 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &applierV2store{store: s, cluster: c}
}

type applierV2store struct {
	store	store.Store
	cluster	*membership.RaftCluster
}

func (a *applierV2store) Delete(r *RequestV2) Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case r.PrevIndex > 0 || r.PrevValue != "":
		return toResponse(a.store.CompareAndDelete(r.Path, r.PrevValue, r.PrevIndex))
	default:
		return toResponse(a.store.Delete(r.Path, r.Dir, r.Recursive))
	}
}
func (a *applierV2store) Post(r *RequestV2) Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return toResponse(a.store.Create(r.Path, r.Dir, r.Val, true, r.TTLOptions()))
}
func (a *applierV2store) Put(r *RequestV2) Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ttlOptions := r.TTLOptions()
	exists, existsSet := pbutil.GetBool(r.PrevExist)
	switch {
	case existsSet:
		if exists {
			if r.PrevIndex == 0 && r.PrevValue == "" {
				return toResponse(a.store.Update(r.Path, r.Val, ttlOptions))
			}
			return toResponse(a.store.CompareAndSwap(r.Path, r.PrevValue, r.PrevIndex, r.Val, ttlOptions))
		}
		return toResponse(a.store.Create(r.Path, r.Dir, r.Val, false, ttlOptions))
	case r.PrevIndex > 0 || r.PrevValue != "":
		return toResponse(a.store.CompareAndSwap(r.Path, r.PrevValue, r.PrevIndex, r.Val, ttlOptions))
	default:
		if storeMemberAttributeRegexp.MatchString(r.Path) {
			id := membership.MustParseMemberIDFromKey(path.Dir(r.Path))
			var attr membership.Attributes
			if err := json.Unmarshal([]byte(r.Val), &attr); err != nil {
				plog.Panicf("unmarshal %s should never fail: %v", r.Val, err)
			}
			if a.cluster != nil {
				a.cluster.UpdateAttributes(id, attr)
			}
			return Response{}
		}
		if r.Path == membership.StoreClusterVersionKey() {
			if a.cluster != nil {
				a.cluster.SetVersion(semver.Must(semver.NewVersion(r.Val)), api.UpdateCapability)
			}
			return Response{}
		}
		return toResponse(a.store.Set(r.Path, r.Dir, r.Val, ttlOptions))
	}
}
func (a *applierV2store) QGet(r *RequestV2) Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return toResponse(a.store.Get(r.Path, r.Recursive, r.Sorted))
}
func (a *applierV2store) Sync(r *RequestV2) Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a.store.DeleteExpiredKeys(time.Unix(0, r.Time))
	return Response{}
}
func (s *EtcdServer) applyV2Request(r *RequestV2) Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer warnOfExpensiveRequest(time.Now(), r, nil, nil)
	switch r.Method {
	case "POST":
		return s.applyV2.Post(r)
	case "PUT":
		return s.applyV2.Put(r)
	case "DELETE":
		return s.applyV2.Delete(r)
	case "QGET":
		return s.applyV2.QGet(r)
	case "SYNC":
		return s.applyV2.Sync(r)
	default:
		return Response{Err: ErrUnknownMethod}
	}
}
func (r *RequestV2) TTLOptions() store.TTLOptionSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	refresh, _ := pbutil.GetBool(r.Refresh)
	ttlOptions := store.TTLOptionSet{Refresh: refresh}
	if r.Expiration != 0 {
		ttlOptions.ExpireTime = time.Unix(0, r.Expiration)
	}
	return ttlOptions
}
func toResponse(ev *store.Event, err error) Response {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Response{Event: ev, Err: err}
}
