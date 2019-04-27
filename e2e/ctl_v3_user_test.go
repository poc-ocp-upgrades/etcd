package e2e

import "testing"

func TestCtlV3UserAdd(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, userAddTest)
}
func TestCtlV3UserAddNoTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, userAddTest, withCfg(configNoTLS))
}
func TestCtlV3UserAddClientTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, userAddTest, withCfg(configClientTLS))
}
func TestCtlV3UserAddPeerTLS(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, userAddTest, withCfg(configPeerTLS))
}
func TestCtlV3UserAddTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, userAddTest, withDialTimeout(0))
}
func TestCtlV3UserDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, userDelTest)
}
func TestCtlV3UserPasswd(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, userPasswdTest)
}

type userCmdDesc struct {
	args		[]string
	expectedStr	string
	stdIn		[]string
}

func userAddTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdSet := []userCmdDesc{{args: []string{"add", "username", "--interactive=false"}, expectedStr: "User username created", stdIn: []string{"password"}}, {args: []string{"add", "usertest:password"}, expectedStr: "User usertest created", stdIn: []string{}}, {args: []string{"add", ":password"}, expectedStr: "empty user name is not allowed.", stdIn: []string{}}, {args: []string{"add", "username", "--interactive=false"}, expectedStr: "user name already exists", stdIn: []string{"password"}}}
	for i, cmd := range cmdSet {
		if err := ctlV3User(cx, cmd.args, cmd.expectedStr, cmd.stdIn); err != nil {
			if cx.dialTimeout > 0 && !isGRPCTimedout(err) {
				cx.t.Fatalf("userAddTest #%d: ctlV3User error (%v)", i, err)
			}
		}
	}
}
func userDelTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdSet := []userCmdDesc{{args: []string{"add", "username", "--interactive=false"}, expectedStr: "User username created", stdIn: []string{"password"}}, {args: []string{"delete", "username"}, expectedStr: "User username deleted"}, {args: []string{"delete", "username"}, expectedStr: "user name not found"}}
	for i, cmd := range cmdSet {
		if err := ctlV3User(cx, cmd.args, cmd.expectedStr, cmd.stdIn); err != nil {
			cx.t.Fatalf("userDelTest #%d: ctlV3User error (%v)", i, err)
		}
	}
}
func userPasswdTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdSet := []userCmdDesc{{args: []string{"add", "username", "--interactive=false"}, expectedStr: "User username created", stdIn: []string{"password"}}, {args: []string{"passwd", "username", "--interactive=false"}, expectedStr: "Password updated", stdIn: []string{"password1"}}}
	for i, cmd := range cmdSet {
		if err := ctlV3User(cx, cmd.args, cmd.expectedStr, cmd.stdIn); err != nil {
			cx.t.Fatalf("userPasswdTest #%d: ctlV3User error (%v)", i, err)
		}
	}
}
func ctlV3User(cx ctlCtx, args []string, expStr string, stdIn []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "user")
	cmdArgs = append(cmdArgs, args...)
	proc, err := spawnCmd(cmdArgs)
	if err != nil {
		return err
	}
	for _, s := range stdIn {
		if err = proc.Send(s + "\r"); err != nil {
			return err
		}
	}
	_, err = proc.Expect(expStr)
	return err
}
