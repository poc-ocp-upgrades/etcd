package balancer

import (
	"go.etcd.io/etcd/clientv3/balancer/picker"
	"go.uber.org/zap"
)

type Config struct {
	Policy	picker.Policy
	Name	string
	Logger	*zap.Logger
}
