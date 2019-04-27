package e2e

import (
	"fmt"
	"os"
	"testing"
	"github.com/coreos/etcd/clientv3"
)

func TestCtlV3AuthEnable(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authEnableTest)
}
func TestCtlV3AuthDisable(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authDisableTest)
}
func TestCtlV3AuthWriteKey(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authCredWriteKeyTest)
}
func TestCtlV3AuthRoleUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authRoleUpdateTest)
}
func TestCtlV3AuthUserDeleteDuringOps(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authUserDeleteDuringOpsTest)
}
func TestCtlV3AuthRoleRevokeDuringOps(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authRoleRevokeDuringOpsTest)
}
func TestCtlV3AuthTxn(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestTxn)
}
func TestCtlV3AuthPrefixPerm(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestPrefixPerm)
}
func TestCtlV3AuthMemberAdd(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestMemberAdd)
}
func TestCtlV3AuthMemberRemove(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestMemberRemove, withQuorum(), withNoStrictReconfig())
}
func TestCtlV3AuthMemberUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestMemberUpdate)
}
func TestCtlV3AuthCertCN(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestCertCN, withCfg(configClientTLSCertAuth))
}
func TestCtlV3AuthRevokeWithDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestRevokeWithDelete)
}
func TestCtlV3AuthInvalidMgmt(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestInvalidMgmt)
}
func TestCtlV3AuthFromKeyPerm(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestFromKeyPerm)
}
func TestCtlV3AuthAndWatch(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestWatch)
}
func TestCtlV3AuthLeaseTestKeepAlive(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authLeaseTestKeepAlive)
}
func TestCtlV3AuthLeaseTestTimeToLiveExpired(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authLeaseTestTimeToLiveExpired)
}
func TestCtlV3AuthRoleGet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestRoleGet)
}
func TestCtlV3AuthUserGet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestUserGet)
}
func TestCtlV3AuthRoleList(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestRoleList)
}
func TestCtlV3AuthDefrag(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestDefrag)
}
func TestCtlV3AuthEndpointHealth(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestEndpointHealth, withQuorum())
}
func TestCtlV3AuthSnapshot(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestSnapshot)
}
func TestCtlV3AuthCertCNAndUsername(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, authTestCertCNAndUsername, withCfg(configClientTLSCertAuth))
}
func authEnableTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
}
func authEnable(cx ctlCtx) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3User(cx, []string{"add", "root", "--interactive=false"}, "User root created", []string{"root"}); err != nil {
		return fmt.Errorf("failed to create root user %v", err)
	}
	if err := ctlV3User(cx, []string{"grant-role", "root", "root"}, "Role root is granted to user root", nil); err != nil {
		return fmt.Errorf("failed to grant root user root role %v", err)
	}
	if err := ctlV3AuthEnable(cx); err != nil {
		return fmt.Errorf("authEnableTest ctlV3AuthEnable error (%v)", err)
	}
	return nil
}
func ctlV3AuthEnable(cx ctlCtx) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "auth", "enable")
	return spawnWithExpect(cmdArgs, "Authentication Enabled")
}
func authDisableTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3Put(cx, "hoo", "a", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3PutFailPerm(cx, "hoo", "bar"); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3AuthDisable(cx); err != nil {
		cx.t.Fatalf("authDisableTest ctlV3AuthDisable error (%v)", err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Put(cx, "hoo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "", ""
	if err := ctlV3Put(cx, "hoo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Get(cx, []string{"hoo"}, []kv{{"hoo", "bar"}}...); err != nil {
		cx.t.Fatal(err)
	}
}
func ctlV3AuthDisable(cx ctlCtx) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "auth", "disable")
	return spawnWithExpect(cmdArgs, "Authentication Disabled")
}
func authCredWriteKeyTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3Put(cx, "foo", "a", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	if err := ctlV3Put(cx, "foo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Get(cx, []string{"foo"}, []kv{{"foo", "bar"}}...); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "a", "b"
	if err := ctlV3PutFailAuth(cx, "foo", "bar"); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Get(cx, []string{"foo"}, []kv{{"foo", "bar"}}...); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Put(cx, "foo", "bar2", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Get(cx, []string{"foo"}, []kv{{"foo", "bar2"}}...); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "badpass"
	if err := ctlV3PutFailAuth(cx, "foo", "baz"); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Get(cx, []string{"foo"}, []kv{{"foo", "bar2"}}...); err != nil {
		cx.t.Fatal(err)
	}
}
func authRoleUpdateTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3Put(cx, "foo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3PutFailPerm(cx, "hoo", "bar"); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, "hoo", "", false}); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Put(cx, "hoo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Get(cx, []string{"hoo"}, []kv{{"hoo", "bar"}}...); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3RoleRevokePermission(cx, "test-role", "hoo", "", false); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3PutFailPerm(cx, "hoo", "bar"); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Get(cx, []string{"foo"}, []kv{{"foo", "bar"}}...); err != nil {
		cx.t.Fatal(err)
	}
}
func authUserDeleteDuringOpsTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3Put(cx, "foo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Put(cx, "foo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Get(cx, []string{"foo"}, []kv{{"foo", "bar"}}...); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	err := ctlV3User(cx, []string{"delete", "test-user"}, "User test-user deleted", []string{})
	if err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3PutFailAuth(cx, "foo", "baz"); err != nil {
		cx.t.Fatal(err)
	}
}
func authRoleRevokeDuringOpsTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3Put(cx, "foo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Put(cx, "foo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Get(cx, []string{"foo"}, []kv{{"foo", "bar"}}...); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3Role(cx, []string{"add", "test-role2"}, "Role test-role2 created"); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3RoleGrantPermission(cx, "test-role2", grantingPerm{true, true, "hoo", "", false}); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3User(cx, []string{"grant-role", "test-user", "test-role2"}, "Role test-role2 is granted to user test-user", nil); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Put(cx, "hoo", "bar", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Get(cx, []string{"hoo"}, []kv{{"hoo", "bar"}}...); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	err := ctlV3User(cx, []string{"revoke-role", "test-user", "test-role"}, "Role test-role is revoked from user test-user", []string{})
	if err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3PutFailPerm(cx, "foo", "baz"); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Put(cx, "hoo", "bar2", ""); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Get(cx, []string{"hoo"}, []kv{{"hoo", "bar2"}}...); err != nil {
		cx.t.Fatal(err)
	}
}
func ctlV3PutFailAuth(cx ctlCtx, key, val string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return spawnWithExpect(append(cx.PrefixArgs(), "put", key, val), "authentication failed")
}
func ctlV3PutFailPerm(cx ctlCtx, key, val string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return spawnWithExpect(append(cx.PrefixArgs(), "put", key, val), "permission denied")
}
func authSetupTestUser(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := ctlV3User(cx, []string{"add", "test-user", "--interactive=false"}, "User test-user created", []string{"pass"}); err != nil {
		cx.t.Fatal(err)
	}
	if err := spawnWithExpect(append(cx.PrefixArgs(), "role", "add", "test-role"), "Role test-role created"); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3User(cx, []string{"grant-role", "test-user", "test-role"}, "Role test-role is granted to user test-user", nil); err != nil {
		cx.t.Fatal(err)
	}
	cmd := append(cx.PrefixArgs(), "role", "grant-permission", "test-role", "readwrite", "foo")
	if err := spawnWithExpect(cmd, "Role test-role updated"); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestTxn(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	keys := []string{"c1", "s1", "f1"}
	grantedKeys := []string{"c2", "s2", "f2"}
	for _, key := range keys {
		if err := ctlV3Put(cx, key, "v", ""); err != nil {
			cx.t.Fatal(err)
		}
	}
	for _, key := range grantedKeys {
		if err := ctlV3Put(cx, key, "v", ""); err != nil {
			cx.t.Fatal(err)
		}
	}
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	cx.user, cx.pass = "root", "root"
	for _, key := range grantedKeys {
		if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, key, "", false}); err != nil {
			cx.t.Fatal(err)
		}
	}
	cx.interactive = true
	cx.user, cx.pass = "test-user", "pass"
	rqs := txnRequests{compare: []string{`version("c2") = "1"`}, ifSucess: []string{"get s2"}, ifFail: []string{"get f2"}, results: []string{"SUCCESS", "s2", "v"}}
	if err := ctlV3Txn(cx, rqs); err != nil {
		cx.t.Fatal(err)
	}
	rqs = txnRequests{compare: []string{`version("c1") = "1"`}, ifSucess: []string{"get s2"}, ifFail: []string{"get f2"}, results: []string{"Error: etcdserver: permission denied"}}
	if err := ctlV3Txn(cx, rqs); err != nil {
		cx.t.Fatal(err)
	}
	rqs = txnRequests{compare: []string{`version("c2") = "1"`}, ifSucess: []string{"get s1"}, ifFail: []string{"get f2"}, results: []string{"Error: etcdserver: permission denied"}}
	if err := ctlV3Txn(cx, rqs); err != nil {
		cx.t.Fatal(err)
	}
	rqs = txnRequests{compare: []string{`version("c2") = "1"`}, ifSucess: []string{"get s2"}, ifFail: []string{"get f1"}, results: []string{"Error: etcdserver: permission denied"}}
	if err := ctlV3Txn(cx, rqs); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestPrefixPerm(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	prefix := "/prefix/"
	cx.user, cx.pass = "root", "root"
	if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, prefix, "", true}); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("%s%d", prefix, i)
		if err := ctlV3Put(cx, key, "val", ""); err != nil {
			cx.t.Fatal(err)
		}
	}
	if err := ctlV3PutFailPerm(cx, clientv3.GetPrefixRangeEnd(prefix), "baz"); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, "", "", true}); err != nil {
		cx.t.Fatal(err)
	}
	prefix2 := "/prefix2/"
	cx.user, cx.pass = "test-user", "pass"
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("%s%d", prefix2, i)
		if err := ctlV3Put(cx, key, "val", ""); err != nil {
			cx.t.Fatal(err)
		}
	}
}
func authTestMemberAdd(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	peerURL := fmt.Sprintf("http://localhost:%d", etcdProcessBasePort+11)
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3MemberAdd(cx, peerURL); err == nil {
		cx.t.Fatalf("ordinary user must not be allowed to add a member")
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3MemberAdd(cx, peerURL); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestMemberRemove(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	ep, memIDToRemove, clusterID := cx.memberToRemove()
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3MemberRemove(cx, ep, memIDToRemove, clusterID); err == nil {
		cx.t.Fatalf("ordinary user must not be allowed to remove a member")
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3MemberRemove(cx, ep, memIDToRemove, clusterID); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestMemberUpdate(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	mr, err := getMemberList(cx)
	if err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	peerURL := fmt.Sprintf("http://localhost:%d", etcdProcessBasePort+11)
	memberID := fmt.Sprintf("%x", mr.Members[0].ID)
	if err = ctlV3MemberUpdate(cx, memberID, peerURL); err == nil {
		cx.t.Fatalf("ordinary user must not be allowed to update a member")
	}
	cx.user, cx.pass = "root", "root"
	if err = ctlV3MemberUpdate(cx, memberID, peerURL); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestCertCN(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3User(cx, []string{"add", "example.com", "--interactive=false"}, "User example.com created", []string{""}); err != nil {
		cx.t.Fatal(err)
	}
	if err := spawnWithExpect(append(cx.PrefixArgs(), "role", "add", "test-role"), "Role test-role created"); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3User(cx, []string{"grant-role", "example.com", "test-role"}, "Role test-role is granted to user example.com", nil); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, "hoo", "", false}); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "", ""
	if err := ctlV3Put(cx, "hoo", "bar", ""); err != nil {
		cx.t.Error(err)
	}
	cx.user, cx.pass = "", ""
	if err := ctlV3PutFailPerm(cx, "baz", "bar"); err != nil {
		cx.t.Error(err)
	}
}
func authTestRevokeWithDelete(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	cx.user, cx.pass = "root", "root"
	if err := ctlV3Role(cx, []string{"add", "test-role2"}, "Role test-role2 created"); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3User(cx, []string{"grant-role", "test-user", "test-role2"}, "Role test-role2 is granted to user test-user", nil); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3User(cx, []string{"get", "test-user"}, "Roles: test-role test-role2", nil); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Role(cx, []string{"delete", "test-role2"}, "Role test-role2 deleted"); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3User(cx, []string{"get", "test-user"}, "Roles: test-role", nil); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestInvalidMgmt(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Role(cx, []string{"delete", "root"}, "Error: etcdserver: invalid auth management"); err == nil {
		cx.t.Fatal("deleting the role root must not be allowed")
	}
	if err := ctlV3User(cx, []string{"revoke-role", "root", "root"}, "Error: etcdserver: invalid auth management", []string{}); err == nil {
		cx.t.Fatal("revoking the role root from the user root must not be allowed")
	}
}
func authTestFromKeyPerm(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	cx.user, cx.pass = "root", "root"
	if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, "z", "\x00", false}); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("z%d", i)
		if err := ctlV3Put(cx, key, "val", ""); err != nil {
			cx.t.Fatal(err)
		}
	}
	largeKey := ""
	for i := 0; i < 10; i++ {
		largeKey += "\xff"
		if err := ctlV3Put(cx, largeKey, "val", ""); err != nil {
			cx.t.Fatal(err)
		}
	}
	if err := ctlV3PutFailPerm(cx, "x", "baz"); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3RoleRevokePermission(cx, "test-role", "z", "", true); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("z%d", i)
		if err := ctlV3PutFailPerm(cx, key, "val"); err != nil {
			cx.t.Fatal(err)
		}
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, "", "\x00", false}); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("z%d", i)
		if err := ctlV3Put(cx, key, "val", ""); err != nil {
			cx.t.Fatal(err)
		}
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3RoleRevokePermission(cx, "test-role", "", "", true); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("z%d", i)
		if err := ctlV3PutFailPerm(cx, key, "val"); err != nil {
			cx.t.Fatal(err)
		}
	}
}
func authLeaseTestKeepAlive(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	leaseID, err := ctlV3LeaseGrant(cx, 10)
	if err != nil {
		cx.t.Fatalf("leaseTestKeepAlive: ctlV3LeaseGrant error (%v)", err)
	}
	if err := ctlV3Put(cx, "key", "val", leaseID); err != nil {
		cx.t.Fatalf("leaseTestKeepAlive: ctlV3Put error (%v)", err)
	}
	if err := ctlV3LeaseKeepAlive(cx, leaseID); err != nil {
		cx.t.Fatalf("leaseTestKeepAlive: ctlV3LeaseKeepAlive error (%v)", err)
	}
	if err := ctlV3Get(cx, []string{"key"}, kv{"key", "val"}); err != nil {
		cx.t.Fatalf("leaseTestKeepAlive: ctlV3Get error (%v)", err)
	}
}
func authLeaseTestTimeToLiveExpired(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	ttl := 3
	if err := leaseTestTimeToLiveExpire(cx, ttl); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestWatch(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, "key", "key4", false}); err != nil {
		cx.t.Fatal(err)
	}
	tests := []struct {
		puts	[]kv
		args	[]string
		wkv	[]kvExec
		want	bool
	}{{[]kv{{"key", "value"}}, []string{"key", "--rev", "1"}, []kvExec{{key: "key", val: "value"}}, true}, {[]kv{{"key1", "val1"}, {"key3", "val3"}, {"key2", "val2"}}, []string{"key", "key3", "--rev", "1"}, []kvExec{{key: "key1", val: "val1"}, {key: "key2", val: "val2"}}, true}, {[]kv{}, []string{"key5", "--rev", "1"}, []kvExec{}, false}, {[]kv{}, []string{"key", "key6", "--rev", "1"}, []kvExec{}, false}}
	cx.user, cx.pass = "test-user", "pass"
	for i, tt := range tests {
		donec := make(chan struct{})
		go func(i int, puts []kv) {
			defer close(donec)
			for j := range puts {
				if err := ctlV3Put(cx, puts[j].key, puts[j].val, ""); err != nil {
					cx.t.Fatalf("watchTest #%d-%d: ctlV3Put error (%v)", i, j, err)
				}
			}
		}(i, tt.puts)
		var err error
		if tt.want {
			err = ctlV3Watch(cx, tt.args, tt.wkv...)
		} else {
			err = ctlV3WatchFailPerm(cx, tt.args)
		}
		if err != nil {
			if cx.dialTimeout > 0 && !isGRPCTimedout(err) {
				cx.t.Errorf("watchTest #%d: ctlV3Watch error (%v)", i, err)
			}
		}
		<-donec
	}
}
func authTestRoleGet(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	expected := []string{"Role test-role", "KV Read:", "foo", "KV Write:", "foo"}
	if err := spawnWithExpects(append(cx.PrefixArgs(), "role", "get", "test-role"), expected...); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := spawnWithExpects(append(cx.PrefixArgs(), "role", "get", "test-role"), expected...); err != nil {
		cx.t.Fatal(err)
	}
	expected = []string{"Error: etcdserver: permission denied"}
	if err := spawnWithExpects(append(cx.PrefixArgs(), "role", "get", "root"), expected...); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestUserGet(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	expected := []string{"User: test-user", "Roles: test-role"}
	if err := spawnWithExpects(append(cx.PrefixArgs(), "user", "get", "test-user"), expected...); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := spawnWithExpects(append(cx.PrefixArgs(), "user", "get", "test-user"), expected...); err != nil {
		cx.t.Fatal(err)
	}
	expected = []string{"Error: etcdserver: permission denied"}
	if err := spawnWithExpects(append(cx.PrefixArgs(), "user", "get", "root"), expected...); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestRoleList(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	if err := spawnWithExpect(append(cx.PrefixArgs(), "role", "list"), "test-role"); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestDefrag(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	maintenanceInitKeys(cx)
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Defrag(cx); err == nil {
		cx.t.Fatal("ordinary user should not be able to issue a defrag request")
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3Defrag(cx); err != nil {
		cx.t.Fatal(err)
	}
}
func authTestSnapshot(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	maintenanceInitKeys(cx)
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	fpath := "test.snapshot"
	defer os.RemoveAll(fpath)
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3SnapshotSave(cx, fpath); err == nil {
		cx.t.Fatal("ordinary user should not be able to save a snapshot")
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3SnapshotSave(cx, fpath); err != nil {
		cx.t.Fatalf("snapshotTest ctlV3SnapshotSave error (%v)", err)
	}
	st, err := getSnapshotStatus(cx, fpath)
	if err != nil {
		cx.t.Fatalf("snapshotTest getSnapshotStatus error (%v)", err)
	}
	if st.Revision != 4 {
		cx.t.Fatalf("expected 4, got %d", st.Revision)
	}
	if st.TotalKey < 3 {
		cx.t.Fatalf("expected at least 3, got %d", st.TotalKey)
	}
}
func authTestEndpointHealth(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	if err := ctlV3EndpointHealth(cx); err != nil {
		cx.t.Fatalf("endpointStatusTest ctlV3EndpointHealth error (%v)", err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3EndpointHealth(cx); err != nil {
		cx.t.Fatalf("endpointStatusTest ctlV3EndpointHealth error (%v)", err)
	}
	cx.user, cx.pass = "root", "root"
	if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, "health", "", false}); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3EndpointHealth(cx); err != nil {
		cx.t.Fatalf("endpointStatusTest ctlV3EndpointHealth error (%v)", err)
	}
}
func authTestCertCNAndUsername(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := authEnable(cx); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "root", "root"
	authSetupTestUser(cx)
	if err := ctlV3User(cx, []string{"add", "example.com", "--interactive=false"}, "User example.com created", []string{""}); err != nil {
		cx.t.Fatal(err)
	}
	if err := spawnWithExpect(append(cx.PrefixArgs(), "role", "add", "test-role-cn"), "Role test-role-cn created"); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3User(cx, []string{"grant-role", "example.com", "test-role-cn"}, "Role test-role-cn is granted to user example.com", nil); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3RoleGrantPermission(cx, "test-role-cn", grantingPerm{true, true, "hoo", "", false}); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3RoleGrantPermission(cx, "test-role", grantingPerm{true, true, "bar", "", false}); err != nil {
		cx.t.Fatal(err)
	}
	cx.user, cx.pass = "", ""
	if err := ctlV3Put(cx, "hoo", "bar", ""); err != nil {
		cx.t.Error(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3Put(cx, "bar", "bar", ""); err != nil {
		cx.t.Error(err)
	}
	cx.user, cx.pass = "", ""
	if err := ctlV3PutFailPerm(cx, "baz", "bar"); err != nil {
		cx.t.Error(err)
	}
	cx.user, cx.pass = "test-user", "pass"
	if err := ctlV3PutFailPerm(cx, "baz", "bar"); err != nil {
		cx.t.Error(err)
	}
}
