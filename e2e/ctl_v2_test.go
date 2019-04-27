package e2e

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/pkg/fileutil"
	"github.com/coreos/etcd/pkg/testutil"
)

func TestCtlV2Set(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Set(t, &configNoTLS, false)
}
func TestCtlV2SetQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Set(t, &configNoTLS, true)
}
func TestCtlV2SetClientTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Set(t, &configClientTLS, false)
}
func TestCtlV2SetPeerTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Set(t, &configPeerTLS, false)
}
func TestCtlV2SetTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Set(t, &configTLS, false)
}
func testCtlV2Set(t *testing.T, cfg *etcdProcessClusterConfig, quorum bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, cfg, quorum)
	defer func() {
		if errC := epc.Close(); errC != nil {
			t.Fatalf("error closing etcd processes (%v)", errC)
		}
	}()
	key, value := "foo", "bar"
	if err := etcdctlSet(epc, key, value); err != nil {
		t.Fatalf("failed set (%v)", err)
	}
	if err := etcdctlGet(epc, key, value, quorum); err != nil {
		t.Fatalf("failed get (%v)", err)
	}
}
func TestCtlV2Mk(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Mk(t, &configNoTLS, false)
}
func TestCtlV2MkQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Mk(t, &configNoTLS, true)
}
func TestCtlV2MkTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Mk(t, &configTLS, false)
}
func testCtlV2Mk(t *testing.T, cfg *etcdProcessClusterConfig, quorum bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, cfg, quorum)
	defer func() {
		if errC := epc.Close(); errC != nil {
			t.Fatalf("error closing etcd processes (%v)", errC)
		}
	}()
	key, value := "foo", "bar"
	if err := etcdctlMk(epc, key, value, true); err != nil {
		t.Fatalf("failed mk (%v)", err)
	}
	if err := etcdctlMk(epc, key, value, false); err != nil {
		t.Fatalf("failed mk (%v)", err)
	}
	if err := etcdctlGet(epc, key, value, quorum); err != nil {
		t.Fatalf("failed get (%v)", err)
	}
}
func TestCtlV2Rm(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Rm(t, &configNoTLS)
}
func TestCtlV2RmTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Rm(t, &configTLS)
}
func testCtlV2Rm(t *testing.T, cfg *etcdProcessClusterConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, cfg, true)
	defer func() {
		if errC := epc.Close(); errC != nil {
			t.Fatalf("error closing etcd processes (%v)", errC)
		}
	}()
	key, value := "foo", "bar"
	if err := etcdctlSet(epc, key, value); err != nil {
		t.Fatalf("failed set (%v)", err)
	}
	if err := etcdctlRm(epc, key, value, true); err != nil {
		t.Fatalf("failed rm (%v)", err)
	}
	if err := etcdctlRm(epc, key, value, false); err != nil {
		t.Fatalf("failed rm (%v)", err)
	}
}
func TestCtlV2Ls(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Ls(t, &configNoTLS, false)
}
func TestCtlV2LsQuorum(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Ls(t, &configNoTLS, true)
}
func TestCtlV2LsTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Ls(t, &configTLS, false)
}
func testCtlV2Ls(t *testing.T, cfg *etcdProcessClusterConfig, quorum bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, cfg, quorum)
	defer func() {
		if errC := epc.Close(); errC != nil {
			t.Fatalf("error closing etcd processes (%v)", errC)
		}
	}()
	key, value := "foo", "bar"
	if err := etcdctlSet(epc, key, value); err != nil {
		t.Fatalf("failed set (%v)", err)
	}
	if err := etcdctlLs(epc, key, quorum); err != nil {
		t.Fatalf("failed ls (%v)", err)
	}
}
func TestCtlV2Watch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Watch(t, &configNoTLS, false)
}
func TestCtlV2WatchTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Watch(t, &configTLS, false)
}
func testCtlV2Watch(t *testing.T, cfg *etcdProcessClusterConfig, noSync bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, cfg, true)
	defer func() {
		if errC := epc.Close(); errC != nil {
			t.Fatalf("error closing etcd processes (%v)", errC)
		}
	}()
	key, value := "foo", "bar"
	errc := etcdctlWatch(epc, key, value, noSync)
	if err := etcdctlSet(epc, key, value); err != nil {
		t.Fatalf("failed set (%v)", err)
	}
	select {
	case err := <-errc:
		if err != nil {
			t.Fatalf("failed watch (%v)", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatalf("watch timed out")
	}
}
func TestCtlV2GetRoleUser(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, &configNoTLS, false)
	defer func() {
		if err := epc.Close(); err != nil {
			t.Fatalf("error closing etcd processes (%v)", err)
		}
	}()
	if err := etcdctlRoleAdd(epc, "foo"); err != nil {
		t.Fatalf("failed to add role (%v)", err)
	}
	if err := etcdctlUserAdd(epc, "username", "password"); err != nil {
		t.Fatalf("failed to add user (%v)", err)
	}
	if err := etcdctlUserGrant(epc, "username", "foo"); err != nil {
		t.Fatalf("failed to grant role (%v)", err)
	}
	if err := etcdctlUserGet(epc, "username"); err != nil {
		t.Fatalf("failed to get user (%v)", err)
	}
	regrantArgs := etcdctlPrefixArgs(epc)
	regrantArgs = append(regrantArgs, "user", "grant", "--roles", "foo", "username")
	if err := spawnWithExpect(regrantArgs, "duplicate"); err != nil {
		t.Fatalf("missing duplicate error on double grant role (%v)", err)
	}
}
func TestCtlV2UserListUsername(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2UserList(t, "username")
}
func TestCtlV2UserListRoot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2UserList(t, "root")
}
func testCtlV2UserList(t *testing.T, username string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, &configNoTLS, false)
	defer func() {
		if err := epc.Close(); err != nil {
			t.Fatalf("error closing etcd processes (%v)", err)
		}
	}()
	if err := etcdctlUserAdd(epc, username, "password"); err != nil {
		t.Fatalf("failed to add user (%v)", err)
	}
	if err := etcdctlUserList(epc, username); err != nil {
		t.Fatalf("failed to list users (%v)", err)
	}
}
func TestCtlV2RoleList(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, &configNoTLS, false)
	defer func() {
		if err := epc.Close(); err != nil {
			t.Fatalf("error closing etcd processes (%v)", err)
		}
	}()
	if err := etcdctlRoleAdd(epc, "foo"); err != nil {
		t.Fatalf("failed to add role (%v)", err)
	}
	if err := etcdctlRoleList(epc, "foo"); err != nil {
		t.Fatalf("failed to list roles (%v)", err)
	}
}
func TestCtlV2Backup(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Backup(t, 0, false)
}
func TestCtlV2BackupSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Backup(t, 1, false)
}
func TestCtlV2BackupV3(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Backup(t, 0, true)
}
func TestCtlV2BackupV3Snapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtlV2Backup(t, 1, true)
}
func testCtlV2Backup(t *testing.T, snapCount int, v3 bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	backupDir, err := ioutil.TempDir("", "testbackup0.etcd")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(backupDir)
	etcdCfg := configNoTLS
	etcdCfg.snapCount = snapCount
	epc1 := setupEtcdctlTest(t, &etcdCfg, false)
	if err := ctlV3Put(ctlCtx{t: t, epc: epc1}, "v3key", "123", ""); err != nil {
		t.Fatal(err)
	}
	if err := etcdctlSet(epc1, "foo1", "bar"); err != nil {
		t.Fatal(err)
	}
	if v3 {
		if err := epc1.Stop(); err != nil {
			t.Fatal(err)
		}
	}
	if err := etcdctlBackup(epc1, epc1.procs[0].Config().dataDirPath, backupDir, v3); err != nil {
		t.Fatal(err)
	}
	if err := epc1.Close(); err != nil {
		t.Fatalf("error closing etcd processes (%v)", err)
	}
	cfg2 := configNoTLS
	cfg2.dataDirPath = backupDir
	cfg2.keepDataDir = true
	cfg2.forceNewCluster = true
	epc2 := setupEtcdctlTest(t, &cfg2, false)
	if err := etcdctlGet(epc2, "foo1", "bar", false); err != nil {
		t.Fatal(err)
	}
	ctx2 := ctlCtx{t: t, epc: epc2}
	if v3 {
		if err := ctlV3Get(ctx2, []string{"v3key"}, kv{"v3key", "123"}); err != nil {
			t.Fatal(err)
		}
	} else {
		if err := ctlV3Get(ctx2, []string{"v3key"}); err != nil {
			t.Fatal(err)
		}
	}
	if err := etcdctlSet(epc2, "foo2", "bar"); err != nil {
		t.Fatal(err)
	}
	if err := etcdctlGet(epc2, "foo2", "bar", false); err != nil {
		t.Fatal(err)
	}
	if err := epc2.Close(); err != nil {
		t.Fatalf("error closing etcd processes (%v)", err)
	}
}
func TestCtlV2AuthWithCommonName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	copiedCfg := configClientTLS
	copiedCfg.clientCertAuthEnabled = true
	epc := setupEtcdctlTest(t, &copiedCfg, false)
	defer func() {
		if err := epc.Close(); err != nil {
			t.Fatalf("error closing etcd processes (%v)", err)
		}
	}()
	if err := etcdctlRoleAdd(epc, "testrole"); err != nil {
		t.Fatalf("failed to add role (%v)", err)
	}
	if err := etcdctlRoleGrant(epc, "testrole", "--rw", "--path=/foo"); err != nil {
		t.Fatalf("failed to grant role (%v)", err)
	}
	if err := etcdctlUserAdd(epc, "root", "123"); err != nil {
		t.Fatalf("failed to add user (%v)", err)
	}
	if err := etcdctlUserAdd(epc, "Autogenerated CA", "123"); err != nil {
		t.Fatalf("failed to add user (%v)", err)
	}
	if err := etcdctlUserGrant(epc, "Autogenerated CA", "testrole"); err != nil {
		t.Fatalf("failed to grant role (%v)", err)
	}
	if err := etcdctlAuthEnable(epc); err != nil {
		t.Fatalf("failed to enable auth (%v)", err)
	}
	if err := etcdctlSet(epc, "foo", "bar"); err != nil {
		t.Fatalf("failed to write (%v)", err)
	}
}
func TestCtlV2ClusterHealth(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer testutil.AfterTest(t)
	epc := setupEtcdctlTest(t, &configNoTLS, true)
	defer func() {
		if err := epc.Close(); err != nil {
			t.Fatalf("error closing etcd processes (%v)", err)
		}
	}()
	if err := etcdctlClusterHealth(epc, "cluster is healthy"); err != nil {
		t.Fatalf("cluster-health expected to be healthy (%v)", err)
	}
	epc.procs[0].Stop()
	for i := 0; i < 3; i++ {
		err := etcdctlClusterHealth(epc, "cluster is degraded")
		if err == nil {
			break
		} else if i == 2 {
			t.Fatalf("cluster-health expected to be degraded (%v)", err)
		}
		time.Sleep(time.Second)
	}
	epc.procs[1].Stop()
	if err := etcdctlClusterHealth(epc, "cluster is unavailable"); err != nil {
		t.Fatalf("cluster-health expected to be unavailable (%v)", err)
	}
	epc.procs[0], epc.procs[1] = nil, nil
}
func etcdctlPrefixArgs(clus *etcdProcessCluster) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	endpoints := strings.Join(clus.EndpointsV2(), ",")
	cmdArgs := []string{ctlBinPath, "--endpoints", endpoints}
	if clus.cfg.clientTLS == clientTLS {
		cmdArgs = append(cmdArgs, "--ca-file", caPath, "--cert-file", certPath, "--key-file", privateKeyPath)
	}
	return cmdArgs
}
func etcdctlClusterHealth(clus *etcdProcessCluster, val string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "cluster-health")
	return spawnWithExpect(cmdArgs, val)
}
func etcdctlSet(clus *etcdProcessCluster, key, value string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "set", key, value)
	return spawnWithExpect(cmdArgs, value)
}
func etcdctlMk(clus *etcdProcessCluster, key, value string, first bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "mk", key, value)
	if first {
		return spawnWithExpect(cmdArgs, value)
	}
	return spawnWithExpect(cmdArgs, "Error:  105: Key already exists")
}
func etcdctlGet(clus *etcdProcessCluster, key, value string, quorum bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "get", key)
	if quorum {
		cmdArgs = append(cmdArgs, "--quorum")
	}
	return spawnWithExpect(cmdArgs, value)
}
func etcdctlRm(clus *etcdProcessCluster, key, value string, first bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "rm", key)
	if first {
		return spawnWithExpect(cmdArgs, "PrevNode.Value: "+value)
	}
	return spawnWithExpect(cmdArgs, "Error:  100: Key not found")
}
func etcdctlLs(clus *etcdProcessCluster, key string, quorum bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "ls")
	if quorum {
		cmdArgs = append(cmdArgs, "--quorum")
	}
	return spawnWithExpect(cmdArgs, key)
}
func etcdctlWatch(clus *etcdProcessCluster, key, value string, noSync bool) <-chan error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "watch", "--after-index=1", key)
	if noSync {
		cmdArgs = append(cmdArgs, "--no-sync")
	}
	errc := make(chan error, 1)
	go func() {
		errc <- spawnWithExpect(cmdArgs, value)
	}()
	return errc
}
func etcdctlRoleAdd(clus *etcdProcessCluster, role string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "role", "add", role)
	return spawnWithExpect(cmdArgs, role)
}
func etcdctlRoleGrant(clus *etcdProcessCluster, role string, perms ...string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "role", "grant")
	cmdArgs = append(cmdArgs, perms...)
	cmdArgs = append(cmdArgs, role)
	return spawnWithExpect(cmdArgs, role)
}
func etcdctlRoleList(clus *etcdProcessCluster, expectedRole string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "role", "list")
	return spawnWithExpect(cmdArgs, expectedRole)
}
func etcdctlUserAdd(clus *etcdProcessCluster, user, pass string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "user", "add", user+":"+pass)
	return spawnWithExpect(cmdArgs, "User "+user+" created")
}
func etcdctlUserGrant(clus *etcdProcessCluster, user, role string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "user", "grant", "--roles", role, user)
	return spawnWithExpect(cmdArgs, "User "+user+" updated")
}
func etcdctlUserGet(clus *etcdProcessCluster, user string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "user", "get", user)
	return spawnWithExpect(cmdArgs, "User: "+user)
}
func etcdctlUserList(clus *etcdProcessCluster, expectedUser string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "user", "list")
	return spawnWithExpect(cmdArgs, expectedUser)
}
func etcdctlAuthEnable(clus *etcdProcessCluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "auth", "enable")
	return spawnWithExpect(cmdArgs, "Authentication Enabled")
}
func etcdctlBackup(clus *etcdProcessCluster, dataDir, backupDir string, v3 bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(etcdctlPrefixArgs(clus), "backup", "--data-dir", dataDir, "--backup-dir", backupDir)
	if v3 {
		cmdArgs = append(cmdArgs, "--with-v3")
	}
	proc, err := spawnCmd(cmdArgs)
	if err != nil {
		return err
	}
	return proc.Close()
}
func mustEtcdctl(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !fileutil.Exist(binDir + "/etcdctl") {
		t.Fatalf("could not find etcdctl binary")
	}
}
func setupEtcdctlTest(t *testing.T, cfg *etcdProcessClusterConfig, quorum bool) *etcdProcessCluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mustEtcdctl(t)
	if !quorum {
		cfg = configStandalone(*cfg)
	}
	epc, err := newEtcdProcessCluster(cfg)
	if err != nil {
		t.Fatalf("could not start etcd process cluster (%v)", err)
	}
	return epc
}
