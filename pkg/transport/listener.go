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
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
	"go.etcd.io/etcd/pkg/tlsutil"
	"go.uber.org/zap"
)

func NewListener(addr, scheme string, tlsinfo *TLSInfo) (l net.Listener, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if l, err = newListener(addr, scheme); err != nil {
		return nil, err
	}
	return wrapTLS(scheme, tlsinfo, l)
}
func newListener(addr string, scheme string) (net.Listener, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme == "unix" || scheme == "unixs" {
		return NewUnixListener(addr)
	}
	return net.Listen("tcp", addr)
}
func wrapTLS(scheme string, tlsinfo *TLSInfo, l net.Listener) (net.Listener, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if scheme != "https" && scheme != "unixs" {
		return l, nil
	}
	return newTLSListener(l, tlsinfo, checkSAN)
}

type TLSInfo struct {
	CertFile		string
	KeyFile			string
	TrustedCAFile		string
	ClientCertAuth		bool
	CRLFile			string
	InsecureSkipVerify	bool
	ServerName		string
	HandshakeFailure	func(*tls.Conn, error)
	CipherSuites		[]uint16
	selfCert		bool
	parseFunc		func([]byte, []byte) (tls.Certificate, error)
	AllowedCN		string
	Logger			*zap.Logger
	EmptyCN			bool
}

func (info TLSInfo) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("cert = %s, key = %s, trusted-ca = %s, client-cert-auth = %v, crl-file = %s", info.CertFile, info.KeyFile, info.TrustedCAFile, info.ClientCertAuth, info.CRLFile)
}
func (info TLSInfo) Empty() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return info.CertFile == "" && info.KeyFile == ""
}
func SelfCert(lg *zap.Logger, dirpath string, hosts []string) (info TLSInfo, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err = os.MkdirAll(dirpath, 0700); err != nil {
		return
	}
	info.Logger = lg
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
		if info.Logger != nil {
			info.Logger.Warn("cannot generate random number", zap.Error(err))
		}
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
		if info.Logger != nil {
			info.Logger.Warn("cannot generate ECDSA key", zap.Error(err))
		}
		return
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &priv.PublicKey, priv)
	if err != nil {
		if info.Logger != nil {
			info.Logger.Warn("cannot generate x509 certificate", zap.Error(err))
		}
		return
	}
	certOut, err := os.Create(certPath)
	if err != nil {
		info.Logger.Warn("cannot cert file", zap.String("path", certPath), zap.Error(err))
		return
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()
	if info.Logger != nil {
		info.Logger.Info("created cert file", zap.String("path", certPath))
	}
	b, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return
	}
	keyOut, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		if info.Logger != nil {
			info.Logger.Warn("cannot key file", zap.String("path", keyPath), zap.Error(err))
		}
		return
	}
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
	keyOut.Close()
	if info.Logger != nil {
		info.Logger.Info("created key file", zap.String("path", keyPath))
	}
	return SelfCert(lg, dirpath, hosts)
}
func (info TLSInfo) baseConfig() (*tls.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if info.KeyFile == "" || info.CertFile == "" {
		return nil, fmt.Errorf("KeyFile and CertFile must both be present[key: %v, cert: %v]", info.KeyFile, info.CertFile)
	}
	if info.Logger == nil {
		info.Logger = zap.NewNop()
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
	cfg.GetCertificate = func(clientHello *tls.ClientHelloInfo) (cert *tls.Certificate, err error) {
		cert, err = tlsutil.NewCert(info.CertFile, info.KeyFile, info.parseFunc)
		if os.IsNotExist(err) {
			if info.Logger != nil {
				info.Logger.Warn("failed to find peer cert files", zap.String("cert-file", info.CertFile), zap.String("key-file", info.KeyFile), zap.Error(err))
			}
		} else if err != nil {
			if info.Logger != nil {
				info.Logger.Warn("failed to create peer certificate", zap.String("cert-file", info.CertFile), zap.String("key-file", info.KeyFile), zap.Error(err))
			}
		}
		return cert, err
	}
	cfg.GetClientCertificate = func(unused *tls.CertificateRequestInfo) (cert *tls.Certificate, err error) {
		cert, err = tlsutil.NewCert(info.CertFile, info.KeyFile, info.parseFunc)
		if os.IsNotExist(err) {
			if info.Logger != nil {
				info.Logger.Warn("failed to find client cert files", zap.String("cert-file", info.CertFile), zap.String("key-file", info.KeyFile), zap.Error(err))
			}
		} else if err != nil {
			if info.Logger != nil {
				info.Logger.Warn("failed to create client certificate", zap.String("cert-file", info.CertFile), zap.String("key-file", info.KeyFile), zap.Error(err))
			}
		}
		return cert, err
	}
	return cfg, nil
}
func (info TLSInfo) cafiles() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cs := make([]string, 0)
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
	if info.TrustedCAFile != "" || info.ClientCertAuth {
		cfg.ClientAuth = tls.RequireAndVerifyClientCert
	}
	cs := info.cafiles()
	if len(cs) > 0 {
		cp, err := tlsutil.NewCertPool(cs)
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
	cs := info.cafiles()
	if len(cs) > 0 {
		cfg.RootCAs, err = tlsutil.NewCertPool(cs)
		if err != nil {
			return nil, err
		}
	}
	if info.selfCert {
		cfg.InsecureSkipVerify = true
	}
	if info.EmptyCN {
		hasNonEmptyCN := false
		cn := ""
		tlsutil.NewCert(info.CertFile, info.KeyFile, func(certPEMBlock []byte, keyPEMBlock []byte) (tls.Certificate, error) {
			var block *pem.Block
			block, _ = pem.Decode(certPEMBlock)
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return tls.Certificate{}, err
			}
			if len(cert.Subject.CommonName) != 0 {
				hasNonEmptyCN = true
				cn = cert.Subject.CommonName
			}
			return tls.X509KeyPair(certPEMBlock, keyPEMBlock)
		})
		if hasNonEmptyCN {
			return nil, fmt.Errorf("cert has non empty Common Name (%s)", cn)
		}
	}
	return cfg, nil
}
func IsClosedConnError(err error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return err != nil && strings.Contains(err.Error(), "closed")
}
