package store

import (
	"testing"
)

func TestIsHidden(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	watch := "/"
	key := "/_foo"
	hidden := isHidden(watch, key)
	if !hidden {
		t.Fatalf("%v should be hidden to %v\n", key, watch)
	}
	watch = "/_foo"
	hidden = isHidden(watch, key)
	if hidden {
		t.Fatalf("%v should not be hidden to %v\n", key, watch)
	}
	key = "/_foo/foo"
	hidden = isHidden(watch, key)
	if hidden {
		t.Fatalf("%v should not be hidden to %v\n", key, watch)
	}
	key = "/_foo/_foo"
	hidden = isHidden(watch, key)
	if !hidden {
		t.Fatalf("%v should be hidden to %v\n", key, watch)
	}
	watch = "_foo/foo"
	key = "/_foo/"
	hidden = isHidden(watch, key)
	if hidden {
		t.Fatalf("%v should not be hidden to %v\n", key, watch)
	}
}
