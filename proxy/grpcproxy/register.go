package grpcproxy

import (
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/coreos/etcd/clientv3/naming"
	"golang.org/x/time/rate"
	gnaming "google.golang.org/grpc/naming"
	"os"
)

const registerRetryRate = 1

func Register(c *clientv3.Client, prefix string, addr string, ttl int) <-chan struct{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rm := rate.NewLimiter(rate.Limit(registerRetryRate), registerRetryRate)
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		for rm.Wait(c.Ctx()) == nil {
			ss, err := registerSession(c, prefix, addr, ttl)
			if err != nil {
				plog.Warningf("failed to create a session %v", err)
				continue
			}
			select {
			case <-c.Ctx().Done():
				ss.Close()
				return
			case <-ss.Done():
				plog.Warning("session expired; possible network partition or server restart")
				plog.Warning("creating a new session to rejoin")
				continue
			}
		}
	}()
	return donec
}
func registerSession(c *clientv3.Client, prefix string, addr string, ttl int) (*concurrency.Session, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ss, err := concurrency.NewSession(c, concurrency.WithTTL(ttl))
	if err != nil {
		return nil, err
	}
	gr := &naming.GRPCResolver{Client: c}
	if err = gr.Update(c.Ctx(), prefix, gnaming.Update{Op: gnaming.Add, Addr: addr, Metadata: getMeta()}, clientv3.WithLease(ss.Lease())); err != nil {
		return nil, err
	}
	plog.Infof("registered %q with %d-second lease", addr, ttl)
	return ss, nil
}

type meta struct {
	Name string `json:"name"`
}

func getMeta() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hostname, _ := os.Hostname()
	bts, _ := json.Marshal(meta{Name: hostname})
	return string(bts)
}
func decodeMeta(s string) (meta, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := meta{}
	err := json.Unmarshal([]byte(s), &m)
	return m, err
}
