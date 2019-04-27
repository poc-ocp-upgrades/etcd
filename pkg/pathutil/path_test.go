package pathutil

import "testing"

func TestCanonicalURLPath(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		p	string
		wp	string
	}{{"/a", "/a"}, {"", "/"}, {"a", "/a"}, {"//a", "/a"}, {"/a/.", "/a"}, {"/a/..", "/"}, {"/a/", "/a/"}, {"/a//", "/a/"}}
	for i, tt := range tests {
		if g := CanonicalURLPath(tt.p); g != tt.wp {
			t.Errorf("#%d: canonical path = %s, want %s", i, g, tt.wp)
		}
	}
}
