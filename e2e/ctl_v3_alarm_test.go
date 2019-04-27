package e2e

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"
	"github.com/coreos/etcd/clientv3"
)

func TestCtlV3Alarm(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCtl(t, alarmTest, withQuota(int64(13*os.Getpagesize())))
}
func alarmTest(cx ctlCtx) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	smallbuf := strings.Repeat("a", 64)
	if err := ctlV3Put(cx, "1st_test", smallbuf, ""); err != nil {
		cx.t.Fatal(err)
	}
	buf := strings.Repeat("b", int(os.Getpagesize()))
	for {
		if err := ctlV3Put(cx, "2nd_test", buf, ""); err != nil {
			if !strings.Contains(err.Error(), "etcdserver: mvcc: database space exceeded") {
				cx.t.Fatal(err)
			}
			break
		}
	}
	if err := ctlV3Alarm(cx, "list", "alarm:NOSPACE"); err != nil {
		cx.t.Fatal(err)
	}
	if err := cURLGet(cx.epc, cURLReq{endpoint: "/health", expected: `{"health":"false"}`}); err != nil {
		cx.t.Fatalf("failed get with curl (%v)", err)
	}
	if err := ctlV3Put(cx, "3rd_test", smallbuf, ""); err != nil {
		if !strings.Contains(err.Error(), "etcdserver: mvcc: database space exceeded") {
			cx.t.Fatal(err)
		}
	}
	eps := cx.epc.EndpointsV3()
	cli, err := clientv3.New(clientv3.Config{Endpoints: eps, DialTimeout: 3 * time.Second})
	if err != nil {
		cx.t.Fatal(err)
	}
	defer cli.Close()
	sresp, err := cli.Status(context.TODO(), eps[0])
	if err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Compact(cx, sresp.Header.Revision, true); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Defrag(cx); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Alarm(cx, "disarm", "alarm:NOSPACE"); err != nil {
		cx.t.Fatal(err)
	}
	if err := ctlV3Put(cx, "4th_test", smallbuf, ""); err != nil {
		cx.t.Fatal(err)
	}
}
func ctlV3Alarm(cx ctlCtx, cmd string, as ...string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmdArgs := append(cx.PrefixArgs(), "alarm", cmd)
	return spawnWithExpects(cmdArgs, as...)
}
