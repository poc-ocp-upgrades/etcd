package picker

import "fmt"

type Policy uint8

const (
	custom			Policy	= iota
	RoundrobinBalanced	Policy	= iota
)

func (p Policy) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch p {
	case custom:
		panic("'custom' picker policy is not supported yet")
	case RoundrobinBalanced:
		return "etcd-client-roundrobin-balanced"
	default:
		panic(fmt.Errorf("invalid balancer picker policy (%d)", p))
	}
}
