package picker

import (
	"google.golang.org/grpc/balancer"
)

type Picker interface{ balancer.Picker }
