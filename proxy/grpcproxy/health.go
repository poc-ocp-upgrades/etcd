package grpcproxy

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/etcdhttp"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"net/http"
	"time"
)

func HandleHealth(mux *http.ServeMux, c *clientv3.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mux.Handle(etcdhttp.PathHealth, etcdhttp.NewHealthHandler(func() etcdhttp.Health {
		return checkHealth(c)
	}))
}
func checkHealth(c *clientv3.Client) etcdhttp.Health {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h := etcdhttp.Health{Health: "false"}
	ctx, cancel := context.WithTimeout(c.Ctx(), time.Second)
	_, err := c.Get(ctx, "a")
	cancel()
	if err == nil || err == rpctypes.ErrPermissionDenied {
		h.Health = "true"
	}
	return h
}
