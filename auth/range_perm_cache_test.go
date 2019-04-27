package auth

import (
	"testing"
	"github.com/coreos/etcd/auth/authpb"
	"github.com/coreos/etcd/pkg/adt"
)

func TestRangePermission(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		perms	[]adt.Interval
		begin	[]byte
		end	[]byte
		want	bool
	}{{[]adt.Interval{adt.NewBytesAffineInterval([]byte("a"), []byte("c")), adt.NewBytesAffineInterval([]byte("x"), []byte("z"))}, []byte("a"), []byte("z"), false}, {[]adt.Interval{adt.NewBytesAffineInterval([]byte("a"), []byte("f")), adt.NewBytesAffineInterval([]byte("c"), []byte("d")), adt.NewBytesAffineInterval([]byte("f"), []byte("z"))}, []byte("a"), []byte("z"), true}, {[]adt.Interval{adt.NewBytesAffineInterval([]byte("a"), []byte("d")), adt.NewBytesAffineInterval([]byte("a"), []byte("b")), adt.NewBytesAffineInterval([]byte("c"), []byte("f"))}, []byte("a"), []byte("f"), true}}
	for i, tt := range tests {
		readPerms := &adt.IntervalTree{}
		for _, p := range tt.perms {
			readPerms.Insert(p, struct{}{})
		}
		result := checkKeyInterval(&unifiedRangePermissions{readPerms: readPerms}, tt.begin, tt.end, authpb.READ)
		if result != tt.want {
			t.Errorf("#%d: result=%t, want=%t", i, result, tt.want)
		}
	}
}
