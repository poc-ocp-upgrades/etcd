package namespace

import (
	"bytes"
	"testing"
)

func TestPrefixInterval(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		pfx	string
		key	[]byte
		end	[]byte
		wKey	[]byte
		wEnd	[]byte
	}{{pfx: "pfx/", key: []byte("a"), wKey: []byte("pfx/a")}, {pfx: "pfx/", key: []byte("abc"), end: []byte("def"), wKey: []byte("pfx/abc"), wEnd: []byte("pfx/def")}, {pfx: "pfx/", key: []byte("abc"), end: []byte{0}, wKey: []byte("pfx/abc"), wEnd: []byte("pfx0")}, {pfx: "\xff\xff", key: []byte("abc"), end: []byte{0}, wKey: []byte("\xff\xffabc"), wEnd: []byte{0}}}
	for i, tt := range tests {
		pfxKey, pfxEnd := prefixInterval(tt.pfx, tt.key, tt.end)
		if !bytes.Equal(pfxKey, tt.wKey) {
			t.Errorf("#%d: expected key=%q, got key=%q", i, tt.wKey, pfxKey)
		}
		if !bytes.Equal(pfxEnd, tt.wEnd) {
			t.Errorf("#%d: expected end=%q, got end=%q", i, tt.wEnd, pfxEnd)
		}
	}
}
