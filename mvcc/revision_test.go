package mvcc

import (
	"bytes"
	"math"
	"reflect"
	"testing"
)

func TestRevision(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []revision{{}, {main: 1, sub: 0}, {main: 1, sub: 1}, {main: 2, sub: 0}, {main: math.MaxInt64, sub: math.MaxInt64}}
	bs := make([][]byte, len(tests))
	for i, tt := range tests {
		b := newRevBytes()
		revToBytes(tt, b)
		bs[i] = b
		if grev := bytesToRev(b); !reflect.DeepEqual(grev, tt) {
			t.Errorf("#%d: revision = %+v, want %+v", i, grev, tt)
		}
	}
	for i := 0; i < len(tests)-1; i++ {
		if bytes.Compare(bs[i], bs[i+1]) >= 0 {
			t.Errorf("#%d: %v (%+v) should be smaller than %v (%+v)", i, bs[i], tests[i], bs[i+1], tests[i+1])
		}
	}
}
