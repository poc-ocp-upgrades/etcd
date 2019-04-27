package adt_test

import (
	"fmt"
	"github.com/coreos/etcd/pkg/adt"
)

func Example() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ivt := &adt.IntervalTree{}
	ivt.Insert(adt.NewInt64Interval(1, 3), 123)
	ivt.Insert(adt.NewInt64Interval(9, 13), 456)
	ivt.Insert(adt.NewInt64Interval(7, 20), 789)
	rs := ivt.Stab(adt.NewInt64Point(10))
	for _, v := range rs {
		fmt.Printf("Overlapping range: %+v\n", v)
	}
}
