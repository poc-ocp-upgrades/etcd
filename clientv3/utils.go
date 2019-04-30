package clientv3

import (
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func jitterUp(duration time.Duration, jitter float64) time.Duration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	multiplier := jitter * (rand.Float64()*2 - 1)
	return time.Duration(float64(duration) * (1 + multiplier))
}
func isOpFuncCalled(op string, opts []OpOption) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, opt := range opts {
		v := reflect.ValueOf(opt)
		if v.Kind() == reflect.Func {
			if opFunc := runtime.FuncForPC(v.Pointer()); opFunc != nil {
				if strings.Contains(opFunc.Name(), op) {
					return true
				}
			}
		}
	}
	return false
}
