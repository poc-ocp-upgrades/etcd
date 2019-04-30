package picker

import (
	"context"
	"google.golang.org/grpc/balancer"
)

func NewErr(err error) Picker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &errPicker{err: err}
}

type errPicker struct{ err error }

func (p *errPicker) Pick(context.Context, balancer.PickOptions) (balancer.SubConn, func(balancer.DoneInfo), error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil, p.err
}
