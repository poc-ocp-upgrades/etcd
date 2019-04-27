package wal

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/wal/walpb"
)

type corruptFunc func(string, int64) error

func TestRepairTruncate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	corruptf := func(p string, offset int64) error {
		f, err := openLast(p)
		if err != nil {
			return err
		}
		defer f.Close()
		return f.Truncate(offset - 4)
	}
	testRepair(t, makeEnts(10), corruptf, 9)
}
func testRepair(t *testing.T, ents [][]raftpb.Entry, corrupt corruptFunc, expectedEnts int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p, err := ioutil.TempDir(os.TempDir(), "waltest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(p)
	w, err := Create(p, nil)
	defer func() {
		if err = w.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Fatal(err)
	}
	for _, es := range ents {
		if err = w.Save(raftpb.HardState{}, es); err != nil {
			t.Fatal(err)
		}
	}
	offset, err := w.tail().Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatal(err)
	}
	w.Close()
	err = corrupt(p, offset)
	if err != nil {
		t.Fatal(err)
	}
	w, err = Open(p, walpb.Snapshot{})
	if err != nil {
		t.Fatal(err)
	}
	_, _, _, err = w.ReadAll()
	if err != io.ErrUnexpectedEOF {
		t.Fatalf("err = %v, want error %v", err, io.ErrUnexpectedEOF)
	}
	w.Close()
	if ok := Repair(p); !ok {
		t.Fatalf("fix = %t, want %t", ok, true)
	}
	w, err = Open(p, walpb.Snapshot{})
	if err != nil {
		t.Fatal(err)
	}
	_, _, walEnts, err := w.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(walEnts) != expectedEnts {
		t.Fatalf("len(ents) = %d, want %d", len(walEnts), expectedEnts)
	}
	for i := 1; i <= 10; i++ {
		es := []raftpb.Entry{{Index: uint64(expectedEnts + i)}}
		if err = w.Save(raftpb.HardState{}, es); err != nil {
			t.Fatal(err)
		}
	}
	w.Close()
	w, err = Open(p, walpb.Snapshot{})
	if err != nil {
		t.Fatal(err)
	}
	_, _, walEnts, err = w.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(walEnts) != expectedEnts+10 {
		t.Fatalf("len(ents) = %d, want %d", len(walEnts), expectedEnts+10)
	}
}
func makeEnts(ents int) (ret [][]raftpb.Entry) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 1; i <= ents; i++ {
		ret = append(ret, []raftpb.Entry{{Index: uint64(i)}})
	}
	return ret
}
func TestRepairWriteTearLast(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	corruptf := func(p string, offset int64) error {
		f, err := openLast(p)
		if err != nil {
			return err
		}
		defer f.Close()
		if offset < 1024 {
			return fmt.Errorf("got offset %d, expected >1024", offset)
		}
		if terr := f.Truncate(1024); terr != nil {
			return terr
		}
		if terr := f.Truncate(offset); terr != nil {
			return terr
		}
		return nil
	}
	testRepair(t, makeEnts(50), corruptf, 40)
}
func TestRepairWriteTearMiddle(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	corruptf := func(p string, offset int64) error {
		f, err := openLast(p)
		if err != nil {
			return err
		}
		defer f.Close()
		_, werr := f.WriteAt(make([]byte, 512), 4096+512)
		return werr
	}
	ents := makeEnts(5)
	dat := make([]byte, 4096)
	for i := range dat {
		dat[i] = byte(i)
	}
	for i := range ents {
		ents[i][0].Data = dat
	}
	testRepair(t, ents, corruptf, 1)
}
