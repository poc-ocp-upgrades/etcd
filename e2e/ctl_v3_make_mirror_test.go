package e2e

import (
	"fmt"
	"testing"
	"time"
)

func TestCtlV3MakeMirror(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, makeMirrorTest)
}
func TestCtlV3MakeMirrorModifyDestPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, makeMirrorModifyDestPrefixTest)
}
func TestCtlV3MakeMirrorNoDestPrefix(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, makeMirrorNoDestPrefixTest)
}
func makeMirrorTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		flags	= []string{}
		kvs	= []kv{{"key1", "val1"}, {"key2", "val2"}, {"key3", "val3"}}
		kvs2	= []kvExec{{key: "key1", val: "val1"}, {key: "key2", val: "val2"}, {key: "key3", val: "val3"}}
		prefix	= "key"
	)
	testMirrorCommand(cx, flags, kvs, kvs2, prefix, prefix)
}
func makeMirrorModifyDestPrefixTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		flags		= []string{"--prefix", "o_", "--dest-prefix", "d_"}
		kvs		= []kv{{"o_key1", "val1"}, {"o_key2", "val2"}, {"o_key3", "val3"}}
		kvs2		= []kvExec{{key: "d_key1", val: "val1"}, {key: "d_key2", val: "val2"}, {key: "d_key3", val: "val3"}}
		srcprefix	= "o_"
		destprefix	= "d_"
	)
	testMirrorCommand(cx, flags, kvs, kvs2, srcprefix, destprefix)
}
func makeMirrorNoDestPrefixTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		flags		= []string{"--prefix", "o_", "--no-dest-prefix"}
		kvs		= []kv{{"o_key1", "val1"}, {"o_key2", "val2"}, {"o_key3", "val3"}}
		kvs2		= []kvExec{{key: "key1", val: "val1"}, {key: "key2", val: "val2"}, {key: "key3", val: "val3"}}
		srcprefix	= "o_"
		destprefix	= "key"
	)
	testMirrorCommand(cx, flags, kvs, kvs2, srcprefix, destprefix)
}
func testMirrorCommand(cx ctlCtx, flags []string, sourcekvs []kv, destkvs []kvExec, srcprefix, destprefix string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mirrorcfg := configAutoTLS
	mirrorcfg.clusterSize = 1
	mirrorcfg.basePort = 10000
	mirrorctx := ctlCtx{t: cx.t, cfg: mirrorcfg, dialTimeout: 7 * time.Second}
	mirrorepc, err := newEtcdProcessCluster(&mirrorctx.cfg)
	if err != nil {
		cx.t.Fatalf("could not start etcd process cluster (%v)", err)
	}
	mirrorctx.epc = mirrorepc
	defer func() {
		if err = mirrorctx.epc.Close(); err != nil {
			cx.t.Fatalf("error closing etcd processes (%v)", err)
		}
	}()
	cmdArgs := append(cx.PrefixArgs(), "make-mirror")
	cmdArgs = append(cmdArgs, flags...)
	cmdArgs = append(cmdArgs, fmt.Sprintf("localhost:%d", mirrorcfg.basePort))
	proc, err := spawnCmd(cmdArgs)
	if err != nil {
		cx.t.Fatal(err)
	}
	defer func() {
		err = proc.Stop()
		if err != nil {
			cx.t.Fatal(err)
		}
	}()
	for i := range sourcekvs {
		if err = ctlV3Put(cx, sourcekvs[i].key, sourcekvs[i].val, ""); err != nil {
			cx.t.Fatal(err)
		}
	}
	if err = ctlV3Get(cx, []string{srcprefix, "--prefix"}, sourcekvs...); err != nil {
		cx.t.Fatal(err)
	}
	if err = ctlV3Watch(mirrorctx, []string{destprefix, "--rev", "1", "--prefix"}, destkvs...); err != nil {
		cx.t.Fatal(err)
	}
}
