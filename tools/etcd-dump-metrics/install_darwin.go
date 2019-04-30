package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"go.etcd.io/etcd/pkg/fileutil"
)

const downloadURL = `https://storage.googleapis.com/etcd/%s/etcd-%s-darwin-amd64.zip`

func install(ver, dir string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ep := fmt.Sprintf(downloadURL, ver, ver)
	resp, err := http.Get(ep)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	zipPath := filepath.Join(dir, "etcd.zip")
	if err = ioutil.WriteFile(zipPath, d, fileutil.PrivateFileMode); err != nil {
		return "", err
	}
	if err = exec.Command("bash", "-c", fmt.Sprintf("unzip %s -d %s", zipPath, dir)).Run(); err != nil {
		return "", err
	}
	bp1 := filepath.Join(dir, fmt.Sprintf("etcd-%s-darwin-amd64", ver), "etcd")
	bp2 := filepath.Join(dir, "etcd")
	return bp2, os.Rename(bp1, bp2)
}
