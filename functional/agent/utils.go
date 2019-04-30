package agent

import (
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
	"go.etcd.io/etcd/pkg/fileutil"
)

func archive(baseDir, etcdLogPath, dataDir string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dir := filepath.Join(baseDir, "etcd-failure-archive", time.Now().Format(time.RFC3339))
	if existDir(dir) {
		dir = filepath.Join(baseDir, "etcd-failure-archive", time.Now().Add(time.Second).Format(time.RFC3339))
	}
	if err := fileutil.TouchDirAll(dir); err != nil {
		return err
	}
	if err := os.Rename(etcdLogPath, filepath.Join(dir, "etcd.log")); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	if err := os.Rename(dataDir, filepath.Join(dir, filepath.Base(dataDir))); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
func existDir(fpath string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	st, err := os.Stat(fpath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	} else {
		return st.IsDir()
	}
	return false
}
func getURLAndPort(addr string) (urlAddr *url.URL, port int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	urlAddr, err = url.Parse(addr)
	if err != nil {
		return nil, -1, err
	}
	var s string
	_, s, err = net.SplitHostPort(urlAddr.Host)
	if err != nil {
		return nil, -1, err
	}
	port, err = strconv.Atoi(s)
	if err != nil {
		return nil, -1, err
	}
	return urlAddr, port, err
}
func cleanPageCache() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := exec.Command("/bin/sh", "-c", `echo "echo 1 > /proc/sys/vm/drop_caches" | sudo sh`)
	return cmd.Run()
}
