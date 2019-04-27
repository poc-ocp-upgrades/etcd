package integration

import (
	"io"
	"os"
	"path/filepath"
	"github.com/coreos/etcd/pkg/transport"
)

func copyTLSFiles(ti transport.TLSInfo, dst string) (transport.TLSInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ci := transport.TLSInfo{KeyFile: filepath.Join(dst, "server-key.pem"), CertFile: filepath.Join(dst, "server.pem"), TrustedCAFile: filepath.Join(dst, "etcd-root-ca.pem"), ClientCertAuth: ti.ClientCertAuth}
	if err := copyFile(ti.KeyFile, ci.KeyFile); err != nil {
		return transport.TLSInfo{}, err
	}
	if err := copyFile(ti.CertFile, ci.CertFile); err != nil {
		return transport.TLSInfo{}, err
	}
	if err := copyFile(ti.TrustedCAFile, ci.TrustedCAFile); err != nil {
		return transport.TLSInfo{}, err
	}
	return ci, nil
}
func copyFile(src, dst string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()
	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()
	if _, err = io.Copy(w, f); err != nil {
		return err
	}
	return w.Sync()
}
