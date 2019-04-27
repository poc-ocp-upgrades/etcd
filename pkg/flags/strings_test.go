package flags

import (
	"testing"
)

func TestStringsSet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		vals	[]string
		val	string
		pass	bool
	}{{[]string{"abc", "def"}, "abc", true}, {[]string{"on", "off", "false"}, "on", true}, {[]string{"abc", "def"}, "ghi", false}, {[]string{"on", "off"}, "", false}}
	for i, tt := range tests {
		sf := NewStringsFlag(tt.vals...)
		if sf.val != tt.vals[0] {
			t.Errorf("#%d: want default val=%v,but got %v", i, tt.vals[0], sf.val)
		}
		err := sf.Set(tt.val)
		if tt.pass != (err == nil) {
			t.Errorf("#%d: want pass=%t, but got err=%v", i, tt.pass, err)
		}
	}
}
