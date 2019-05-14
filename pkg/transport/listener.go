package transport

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/coreos/etcd/pkg/tlsutil"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func NewListener(addr, scheme string, tlsinfo *TLSInfo) (l net.Listener, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if l, err = newListener(addr, scheme); err != nil {
		return nil, err
	}
	return wrapTLS(addr, scheme, tlsinfo, l)
}
func newListener(addr string, scheme string) (net.Listener, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "unix" || scheme == "unixs" {
		return NewUnixListener(addr)
	}
	return net.Listen("tcp", addr)
}
func wrapTLS(addr, scheme string, tlsinfo *TLSInfo, l net.Listener) (net.Listener, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme != "https" && scheme != "unixs" {
		return l, nil
	}
	return newTLSListener(l, tlsinfo, checkSAN)
}

type TLSInfo struct {
	CertFile           string
	KeyFile            string
	CAFile             string
	TrustedCAFile      string
	ClientCertAuth     bool
	CRLFile            string
	InsecureSkipVerify bool
	ServerName         string
	HandshakeFailure   func(*tls.Conn, error)
	CipherSuites       []uint16
	selfCert           bool
	parseFunc          func([]byte, []byte) (tls.Certificate, error)
	AllowedCN          string
}

func (info TLSInfo) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("cert = %s, key = %s, ca = %s, trusted-ca = %s, client-cert-auth = %v, crl-file = %s", info.CertFile, info.KeyFile, info.CAFile, info.TrustedCAFile, info.ClientCertAuth, info.CRLFile)
}
func (info TLSInfo) Empty() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return info.CertFile == "" && info.KeyFile == ""
}
func SelfCert(dirpath string, hosts []string) (info TLSInfo, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err = os.MkdirAll(dirpath, 0700); err != nil {
		return
	}
	certPath := filepath.Join(dirpath, "cert.pem")
	keyPath := filepath.Join(dirpath, "key.pem")
	_, errcert := os.Stat(certPath)
	_, errkey := os.Stat(keyPath)
	if errcert == nil && errkey == nil {
		info.CertFile = certPath
		info.KeyFile = keyPath
		info.selfCert = true
		return
	}
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return
	}
	tmpl := x509.Certificate{SerialNumber: serialNumber, Subject: pkix.Name{Organization: []string{"etcd"}}, NotBefore: time.Now(), NotAfter: time.Now().Add(365 * (24 * time.Hour)), KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, BasicConstraintsValid: true}
	for _, host := range hosts {
		h, _, _ := net.SplitHostPort(host)
		if ip := net.ParseIP(h); ip != nil {
			tmpl.IPAddresses = append(tmpl.IPAddresses, ip)
		} else {
			tmpl.DNSNames = append(tmpl.DNSNames, h)
		}
	}
	priv, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		return
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &priv.PublicKey, priv)
	if err != nil {
		return
	}
	certOut, err := os.Create(certPath)
	if err != nil {
		return
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()
	b, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return
	}
	keyOut, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return
	}
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
	keyOut.Close()
	return SelfCert(dirpath, hosts)
}
func (info TLSInfo) baseConfig() (*tls.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if info.KeyFile == "" || info.CertFile == "" {
		return nil, fmt.Errorf("KeyFile and CertFile must both be present[key: %v, cert: %v]", info.KeyFile, info.CertFile)
	}
	_, err := tlsutil.NewCert(info.CertFile, info.KeyFile, info.parseFunc)
	if err != nil {
		return nil, err
	}
	cfg := &tls.Config{MinVersion: tls.VersionTLS12, ServerName: info.ServerName}
	if len(info.CipherSuites) > 0 {
		cfg.CipherSuites = info.CipherSuites
	}
	if info.AllowedCN != "" {
		cfg.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			for _, chains := range verifiedChains {
				if len(chains) != 0 {
					if info.AllowedCN == chains[0].Subject.CommonName {
						return nil
					}
				}
			}
			return errors.New("CommonName authentication failed")
		}
	}
	cfg.GetCertificate = func(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		return tlsutil.NewCert(info.CertFile, info.KeyFile, info.parseFunc)
	}
	cfg.GetClientCertificate = func(unused *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		return tlsutil.NewCert(info.CertFile, info.KeyFile, info.parseFunc)
	}
	return cfg, nil
}
func (info TLSInfo) cafiles() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cs := make([]string, 0)
	if info.CAFile != "" {
		cs = append(cs, info.CAFile)
	}
	if info.TrustedCAFile != "" {
		cs = append(cs, info.TrustedCAFile)
	}
	return cs
}
func (info TLSInfo) ServerConfig() (*tls.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg, err := info.baseConfig()
	if err != nil {
		return nil, err
	}
	cfg.ClientAuth = tls.NoClientCert
	if info.CAFile != "" || info.ClientCertAuth {
		cfg.ClientAuth = tls.RequireAndVerifyClientCert
	}
	CAFiles := info.cafiles()
	if len(CAFiles) > 0 {
		cp, err := tlsutil.NewCertPool(CAFiles)
		if err != nil {
			return nil, err
		}
		cfg.ClientCAs = cp
	}
	cfg.NextProtos = []string{"h2"}
	return cfg, nil
}
func (info TLSInfo) ClientConfig() (*tls.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cfg *tls.Config
	var err error
	if !info.Empty() {
		cfg, err = info.baseConfig()
		if err != nil {
			return nil, err
		}
	} else {
		cfg = &tls.Config{ServerName: info.ServerName}
	}
	cfg.InsecureSkipVerify = info.InsecureSkipVerify
	CAFiles := info.cafiles()
	if len(CAFiles) > 0 {
		cfg.RootCAs, err = tlsutil.NewCertPool(CAFiles)
		if err != nil {
			return nil, err
		}
	}
	if info.selfCert {
		cfg.InsecureSkipVerify = true
	}
	return cfg, nil
}
func IsClosedConnError(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return err != nil && strings.Contains(err.Error(), "closed")
}
