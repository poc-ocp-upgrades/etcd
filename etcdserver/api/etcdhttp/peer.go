package etcdhttp

import (
	"encoding/json"
	"net/http"
	"go.etcd.io/etcd/etcdserver"
	"go.etcd.io/etcd/etcdserver/api"
	"go.etcd.io/etcd/etcdserver/api/rafthttp"
	"go.etcd.io/etcd/lease/leasehttp"
	"go.uber.org/zap"
)

const (
	peerMembersPrefix = "/members"
)

func NewPeerHandler(lg *zap.Logger, s etcdserver.ServerPeer) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newPeerHandler(lg, s.Cluster(), s.RaftHandler(), s.LeaseHandler())
}
func newPeerHandler(lg *zap.Logger, cluster api.Cluster, raftHandler http.Handler, leaseHandler http.Handler) http.Handler {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mh := &peerMembersHandler{lg: lg, cluster: cluster}
	mux := http.NewServeMux()
	mux.HandleFunc("/", http.NotFound)
	mux.Handle(rafthttp.RaftPrefix, raftHandler)
	mux.Handle(rafthttp.RaftPrefix+"/", raftHandler)
	mux.Handle(peerMembersPrefix, mh)
	if leaseHandler != nil {
		mux.Handle(leasehttp.LeasePrefix, leaseHandler)
		mux.Handle(leasehttp.LeaseInternalPrefix, leaseHandler)
	}
	mux.HandleFunc(versionPath, versionHandler(cluster, serveVersion))
	return mux
}

type peerMembersHandler struct {
	lg	*zap.Logger
	cluster	api.Cluster
}

func (h *peerMembersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !allowMethod(w, r, "GET") {
		return
	}
	w.Header().Set("X-Etcd-Cluster-ID", h.cluster.ID().String())
	if r.URL.Path != peerMembersPrefix {
		http.Error(w, "bad path", http.StatusBadRequest)
		return
	}
	ms := h.cluster.Members()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ms); err != nil {
		if h.lg != nil {
			h.lg.Warn("failed to encode membership members", zap.Error(err))
		} else {
			plog.Warningf("failed to encode members response (%v)", err)
		}
	}
}
