package ordering

import (
	"errors"
	"github.com/coreos/etcd/clientv3"
	"sync"
	"time"
)

type OrderViolationFunc func(op clientv3.Op, resp clientv3.OpResponse, prevRev int64) error

var ErrNoGreaterRev = errors.New("etcdclient: no cluster members have a revision higher than the previously received revision")

func NewOrderViolationSwitchEndpointClosure(c clientv3.Client) OrderViolationFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var mu sync.Mutex
	violationCount := 0
	return func(op clientv3.Op, resp clientv3.OpResponse, prevRev int64) error {
		if violationCount > len(c.Endpoints()) {
			return ErrNoGreaterRev
		}
		mu.Lock()
		defer mu.Unlock()
		eps := c.Endpoints()
		c.SetEndpoints(eps[violationCount%len(eps)])
		time.Sleep(1 * time.Second)
		c.SetEndpoints(eps...)
		violationCount++
		return nil
	}
}
