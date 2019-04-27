package rafthttp

import (
	"bytes"
	"encoding/binary"
	"io"
	"net/http"
	"reflect"
	"testing"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/version"
	"github.com/coreos/go-semver/semver"
)

func TestEntry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []raftpb.Entry{{}, {Term: 1, Index: 1}, {Term: 1, Index: 1, Data: []byte("some data")}}
	for i, tt := range tests {
		b := &bytes.Buffer{}
		if err := writeEntryTo(b, &tt); err != nil {
			t.Errorf("#%d: unexpected write ents error: %v", i, err)
			continue
		}
		var ent raftpb.Entry
		if err := readEntryFrom(b, &ent); err != nil {
			t.Errorf("#%d: unexpected read ents error: %v", i, err)
			continue
		}
		if !reflect.DeepEqual(ent, tt) {
			t.Errorf("#%d: ent = %+v, want %+v", i, ent, tt)
		}
	}
}
func TestCompareMajorMinorVersion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		va, vb	*semver.Version
		w	int
	}{{semver.Must(semver.NewVersion("2.1.0")), semver.Must(semver.NewVersion("2.1.0")), 0}, {semver.Must(semver.NewVersion("2.0.0")), semver.Must(semver.NewVersion("2.1.0")), -1}, {semver.Must(semver.NewVersion("2.2.0")), semver.Must(semver.NewVersion("2.1.0")), 1}, {semver.Must(semver.NewVersion("2.1.1")), semver.Must(semver.NewVersion("2.1.0")), 0}, {semver.Must(semver.NewVersion("2.1.0-alpha.0")), semver.Must(semver.NewVersion("2.1.0")), 0}}
	for i, tt := range tests {
		if g := compareMajorMinorVersion(tt.va, tt.vb); g != tt.w {
			t.Errorf("#%d: compare = %d, want %d", i, g, tt.w)
		}
	}
}
func TestServerVersion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		h	http.Header
		wv	*semver.Version
	}{{http.Header{}, semver.Must(semver.NewVersion("2.0.0"))}, {http.Header{"X-Server-Version": []string{"2.1.0"}}, semver.Must(semver.NewVersion("2.1.0"))}, {http.Header{"X-Server-Version": []string{"2.1.0-alpha.0+git"}}, semver.Must(semver.NewVersion("2.1.0-alpha.0+git"))}}
	for i, tt := range tests {
		v := serverVersion(tt.h)
		if v.String() != tt.wv.String() {
			t.Errorf("#%d: version = %s, want %s", i, v, tt.wv)
		}
	}
}
func TestMinClusterVersion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tests := []struct {
		h	http.Header
		wv	*semver.Version
	}{{http.Header{}, semver.Must(semver.NewVersion("2.0.0"))}, {http.Header{"X-Min-Cluster-Version": []string{"2.1.0"}}, semver.Must(semver.NewVersion("2.1.0"))}, {http.Header{"X-Min-Cluster-Version": []string{"2.1.0-alpha.0+git"}}, semver.Must(semver.NewVersion("2.1.0-alpha.0+git"))}}
	for i, tt := range tests {
		v := minClusterVersion(tt.h)
		if v.String() != tt.wv.String() {
			t.Errorf("#%d: version = %s, want %s", i, v, tt.wv)
		}
	}
}
func TestCheckVersionCompatibility(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ls := semver.Must(semver.NewVersion(version.Version))
	lmc := semver.Must(semver.NewVersion(version.MinClusterVersion))
	tests := []struct {
		server		*semver.Version
		minCluster	*semver.Version
		wok		bool
	}{{ls, lmc, true}, {lmc, &semver.Version{}, true}, {&semver.Version{Major: ls.Major + 1}, ls, true}, {&semver.Version{Major: lmc.Major - 1}, &semver.Version{}, false}, {&semver.Version{Major: ls.Major + 1, Minor: 1}, &semver.Version{Major: ls.Major + 1}, false}}
	for i, tt := range tests {
		err := checkVersionCompability("", tt.server, tt.minCluster)
		if ok := err == nil; ok != tt.wok {
			t.Errorf("#%d: ok = %v, want %v", i, ok, tt.wok)
		}
	}
}
func writeEntryTo(w io.Writer, ent *raftpb.Entry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := ent.Size()
	if err := binary.Write(w, binary.BigEndian, uint64(size)); err != nil {
		return err
	}
	b, err := ent.Marshal()
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}
func readEntryFrom(r io.Reader, ent *raftpb.Entry) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l uint64
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}
	buf := make([]byte, int(l))
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	return ent.Unmarshal(buf)
}
