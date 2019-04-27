package e2e

import "testing"

func TestCtlV3Defrag(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, defragTest)
}
func maintenanceInitKeys(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var kvs = []kv{{"key", "val1"}, {"key", "val2"}, {"key", "val3"}}
	for i := range kvs {
		if err := ctlV3Put(cx, kvs[i].key, kvs[i].val, ""); err != nil {
			cx.t.Fatal(err)
		}
	}
}
func defragTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	maintenanceInitKeys(cx)
	if err := ctlV3Compact(cx, 4, cx.compactPhysical); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Defrag(cx); err != nil {
		cx.t.Fatalf("defragTest ctlV3Defrag error (%v)", err)
	}
}
func ctlV3Defrag(cx ctlCtx) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "defrag")
	lines := make([]string, cx.epc.cfg.clusterSize)
	for i := range lines {
		lines[i] = "Finished defragmenting etcd member"
	}
	return spawnWithExpects(cmdArgs, lines...)
}
