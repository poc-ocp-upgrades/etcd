package tester

import (
	"fmt"
	"os/exec"
	"go.etcd.io/etcd/functional/rpcpb"
)

type caseExternal struct {
	Case
	desc		string
	rpcpbCase	rpcpb.Case
	scriptPath	string
}

func (c *caseExternal) Inject(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return exec.Command(c.scriptPath, "enable", fmt.Sprintf("%d", clus.rd)).Run()
}
func (c *caseExternal) Recover(clus *Cluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return exec.Command(c.scriptPath, "disable", fmt.Sprintf("%d", clus.rd)).Run()
}
func (c *caseExternal) Desc() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.desc
}
func (c *caseExternal) TestCase() rpcpb.Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.rpcpbCase
}
func new_Case_EXTERNAL(scriptPath string) Case {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &caseExternal{desc: fmt.Sprintf("external fault injector (script: %q)", scriptPath), rpcpbCase: rpcpb.Case_EXTERNAL, scriptPath: scriptPath}
}
