package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/pkg/expect"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestCtlV3Snapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, snapshotTest)
}
func snapshotTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	maintenanceInitKeys(cx)
	leaseID, err := ctlV3LeaseGrant(cx, 100)
	if err != nil {
		cx.t.Fatalf("snapshot: ctlV3LeaseGrant error (%v)", err)
	}
	if err = ctlV3Put(cx, "withlease", "withlease", leaseID); err != nil {
		cx.t.Fatalf("snapshot: ctlV3Put error (%v)", err)
	}
	fpath := "test.snapshot"
	defer os.RemoveAll(fpath)
	if err = ctlV3SnapshotSave(cx, fpath); err != nil {
		cx.t.Fatalf("snapshotTest ctlV3SnapshotSave error (%v)", err)
	}
	st, err := getSnapshotStatus(cx, fpath)
	if err != nil {
		cx.t.Fatalf("snapshotTest getSnapshotStatus error (%v)", err)
	}
	if st.Revision != 5 {
		cx.t.Fatalf("expected 4, got %d", st.Revision)
	}
	if st.TotalKey < 4 {
		cx.t.Fatalf("expected at least 4, got %d", st.TotalKey)
	}
}
func TestCtlV3SnapshotCorrupt(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, snapshotCorruptTest)
}
func snapshotCorruptTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fpath := "test.snapshot"
	defer os.RemoveAll(fpath)
	if err := ctlV3SnapshotSave(cx, fpath); err != nil {
		cx.t.Fatalf("snapshotTest ctlV3SnapshotSave error (%v)", err)
	}
	f, oerr := os.OpenFile(fpath, os.O_WRONLY, 0)
	if oerr != nil {
		cx.t.Fatal(oerr)
	}
	if _, err := f.Write(make([]byte, 512)); err != nil {
		cx.t.Fatal(err)
	}
	f.Close()
	defer os.RemoveAll("snap.etcd")
	serr := spawnWithExpect(append(cx.PrefixArgs(), "snapshot", "restore", "--data-dir", "snap.etcd", fpath), "expected sha256")
	if serr != nil {
		cx.t.Fatal(serr)
	}
}
func TestCtlV3SnapshotStatusBeforeRestore(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, snapshotStatusBeforeRestoreTest)
}
func snapshotStatusBeforeRestoreTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fpath := "test.snapshot"
	defer os.RemoveAll(fpath)
	if err := ctlV3SnapshotSave(cx, fpath); err != nil {
		cx.t.Fatalf("snapshotTest ctlV3SnapshotSave error (%v)", err)
	}
	_, err := getSnapshotStatus(cx, fpath)
	if err != nil {
		cx.t.Fatalf("snapshotTest getSnapshotStatus error (%v)", err)
	}
	defer os.RemoveAll("snap.etcd")
	serr := spawnWithExpect(append(cx.PrefixArgs(), "snapshot", "restore", "--data-dir", "snap.etcd", fpath), "added member")
	if serr != nil {
		cx.t.Fatal(serr)
	}
}
func ctlV3SnapshotSave(cx ctlCtx, fpath string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "snapshot", "save", fpath)
	return spawnWithExpect(cmdArgs, fmt.Sprintf("Snapshot saved at %s", fpath))
}

type snapshotStatus struct {
	Hash		uint32	`json:"hash"`
	Revision	int64	`json:"revision"`
	TotalKey	int	`json:"totalKey"`
	TotalSize	int64	`json:"totalSize"`
}

func getSnapshotStatus(cx ctlCtx, fpath string) (snapshotStatus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "--write-out", "json", "snapshot", "status", fpath)
	proc, err := spawnCmd(cmdArgs)
	if err != nil {
		return snapshotStatus{}, err
	}
	var txt string
	txt, err = proc.Expect("totalKey")
	if err != nil {
		return snapshotStatus{}, err
	}
	if err = proc.Close(); err != nil {
		return snapshotStatus{}, err
	}
	resp := snapshotStatus{}
	dec := json.NewDecoder(strings.NewReader(txt))
	if err := dec.Decode(&resp); err == io.EOF {
		return snapshotStatus{}, err
	}
	return resp, nil
}
func TestIssue6361(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	mustEtcdctl(t)
	os.Setenv("ETCDCTL_API", "3")
	defer os.Unsetenv("ETCDCTL_API")
	epc, err := newEtcdProcessCluster(&etcdProcessClusterConfig{clusterSize: 1, initialToken: "new", keepDataDir: true})
	if err != nil {
		t.Fatalf("could not start etcd process cluster (%v)", err)
	}
	defer func() {
		if errC := epc.Close(); errC != nil {
			t.Fatalf("error closing etcd processes (%v)", errC)
		}
	}()
	dialTimeout := 7 * time.Second
	prefixArgs := []string{ctlBinPath, "--endpoints", strings.Join(epc.EndpointsV3(), ","), "--dial-timeout", dialTimeout.String()}
	kvs := []kv{{"foo1", "val1"}, {"foo2", "val2"}, {"foo3", "val3"}}
	for i := range kvs {
		if err = spawnWithExpect(append(prefixArgs, "put", kvs[i].key, kvs[i].val), "OK"); err != nil {
			t.Fatal(err)
		}
	}
	fpath := filepath.Join(os.TempDir(), "test.snapshot")
	defer os.RemoveAll(fpath)
	if err = spawnWithExpect(append(prefixArgs, "snapshot", "save", fpath), fmt.Sprintf("Snapshot saved at %s", fpath)); err != nil {
		t.Fatal(err)
	}
	if err = epc.procs[0].Stop(); err != nil {
		t.Fatal(err)
	}
	newDataDir := filepath.Join(os.TempDir(), "test.data")
	defer os.RemoveAll(newDataDir)
	err = spawnWithExpect([]string{ctlBinPath, "snapshot", "restore", fpath, "--name", epc.procs[0].Config().name, "--initial-cluster", epc.procs[0].Config().initialCluster, "--initial-cluster-token", epc.procs[0].Config().initialToken, "--initial-advertise-peer-urls", epc.procs[0].Config().purl.String(), "--data-dir", newDataDir}, "membership: added member")
	if err != nil {
		t.Fatal(err)
	}
	epc.procs[0].Config().dataDirPath = newDataDir
	for i := range epc.procs[0].Config().args {
		if epc.procs[0].Config().args[i] == "--data-dir" {
			epc.procs[0].Config().args[i+1] = newDataDir
		}
	}
	if err = epc.procs[0].Restart(); err != nil {
		t.Fatal(err)
	}
	for i := range kvs {
		if err = spawnWithExpect(append(prefixArgs, "get", kvs[i].key), kvs[i].val); err != nil {
			t.Fatal(err)
		}
	}
	clientURL := fmt.Sprintf("http://localhost:%d", etcdProcessBasePort+30)
	peerURL := fmt.Sprintf("http://localhost:%d", etcdProcessBasePort+31)
	err = spawnWithExpect(append(prefixArgs, "member", "add", "newmember", fmt.Sprintf("--peer-urls=%s", peerURL)), " added to cluster ")
	if err != nil {
		t.Fatal(err)
	}
	var newDataDir2 string
	newDataDir2, err = ioutil.TempDir("", "newdata2")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(newDataDir2)
	name2 := "infra2"
	initialCluster2 := epc.procs[0].Config().initialCluster + fmt.Sprintf(",%s=%s", name2, peerURL)
	var nepc *expect.ExpectProcess
	nepc, err = spawnCmd([]string{epc.procs[0].Config().execPath, "--name", name2, "--listen-client-urls", clientURL, "--advertise-client-urls", clientURL, "--listen-peer-urls", peerURL, "--initial-advertise-peer-urls", peerURL, "--initial-cluster", initialCluster2, "--initial-cluster-state", "existing", "--data-dir", newDataDir2})
	if err != nil {
		t.Fatal(err)
	}
	if _, err = nepc.Expect("enabled capabilities for version"); err != nil {
		t.Fatal(err)
	}
	prefixArgs = []string{ctlBinPath, "--endpoints", clientURL, "--dial-timeout", dialTimeout.String()}
	for i := range kvs {
		if err = spawnWithExpect(append(prefixArgs, "get", kvs[i].key), kvs[i].val); err != nil {
			t.Fatal(err)
		}
	}
	if err = nepc.Stop(); err != nil {
		t.Fatal(err)
	}
}
