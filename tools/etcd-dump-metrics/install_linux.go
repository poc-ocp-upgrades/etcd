package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"
	"go.etcd.io/etcd/pkg/fileutil"
)

const downloadURL = `https://storage.googleapis.com/etcd/%s/etcd-%s-linux-amd64.tar.gz`

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
	tarPath := filepath.Join(dir, "etcd.tar.gz")
	if err = ioutil.WriteFile(tarPath, d, fileutil.PrivateFileMode); err != nil {
		return "", err
	}
	if err = exec.Command("bash", "-c", fmt.Sprintf("tar xzvf %s -C %s --strip-components=1", tarPath, dir)).Run(); err != nil {
		return "", err
	}
	return filepath.Join(dir, "etcd"), nil
}
