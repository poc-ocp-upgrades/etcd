package tester

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"
	"github.com/coreos/etcd/functional/rpcpb"
	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"
)

func read(lg *zap.Logger, fpath string) (*Cluster, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bts, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	lg.Info("opened configuration file", zap.String("path", fpath))
	clus := &Cluster{lg: lg}
	if err = yaml.Unmarshal(bts, clus); err != nil {
		return nil, err
	}
	if len(clus.Members) < 3 {
		return nil, fmt.Errorf("len(clus.Members) expects at least 3, got %d", len(clus.Members))
	}
	for i, mem := range clus.Members {
		if mem.BaseDir == "" {
			return nil, fmt.Errorf("BaseDir cannot be empty (got %q)", mem.BaseDir)
		}
		if mem.EtcdLogPath == "" {
			return nil, fmt.Errorf("EtcdLogPath cannot be empty (got %q)", mem.EtcdLogPath)
		}
		if mem.Etcd.Name == "" {
			return nil, fmt.Errorf("'--name' cannot be empty (got %+v)", mem)
		}
		if mem.Etcd.DataDir == "" {
			return nil, fmt.Errorf("'--data-dir' cannot be empty (got %+v)", mem)
		}
		if mem.Etcd.SnapshotCount == 0 {
			return nil, fmt.Errorf("'--snapshot-count' cannot be 0 (got %+v)", mem.Etcd.SnapshotCount)
		}
		if mem.Etcd.DataDir == "" {
			return nil, fmt.Errorf("'--data-dir' cannot be empty (got %q)", mem.Etcd.DataDir)
		}
		if mem.Etcd.WALDir == "" {
			clus.Members[i].Etcd.WALDir = filepath.Join(mem.Etcd.DataDir, "member", "wal")
		}
		switch mem.Etcd.InitialClusterState {
		case "new":
		case "existing":
		default:
			return nil, fmt.Errorf("'--initial-cluster-state' got %q", mem.Etcd.InitialClusterState)
		}
		if mem.Etcd.HeartbeatIntervalMs == 0 {
			return nil, fmt.Errorf("'--heartbeat-interval' cannot be 0 (got %+v)", mem.Etcd)
		}
		if mem.Etcd.ElectionTimeoutMs == 0 {
			return nil, fmt.Errorf("'--election-timeout' cannot be 0 (got %+v)", mem.Etcd)
		}
		if int64(clus.Tester.DelayLatencyMs) <= mem.Etcd.ElectionTimeoutMs {
			return nil, fmt.Errorf("delay latency %d ms must be greater than election timeout %d ms", clus.Tester.DelayLatencyMs, mem.Etcd.ElectionTimeoutMs)
		}
		port := ""
		listenClientPorts := make([]string, len(clus.Members))
		for i, u := range mem.Etcd.ListenClientURLs {
			if !isValidURL(u) {
				return nil, fmt.Errorf("'--listen-client-urls' has valid URL %q", u)
			}
			listenClientPorts[i], err = getPort(u)
			if err != nil {
				return nil, fmt.Errorf("'--listen-client-urls' has no port %q", u)
			}
		}
		for i, u := range mem.Etcd.AdvertiseClientURLs {
			if !isValidURL(u) {
				return nil, fmt.Errorf("'--advertise-client-urls' has valid URL %q", u)
			}
			port, err = getPort(u)
			if err != nil {
				return nil, fmt.Errorf("'--advertise-client-urls' has no port %q", u)
			}
			if mem.EtcdClientProxy && listenClientPorts[i] == port {
				return nil, fmt.Errorf("clus.Members[%d] requires client port proxy, but advertise port %q conflicts with listener port %q", i, port, listenClientPorts[i])
			}
		}
		listenPeerPorts := make([]string, len(clus.Members))
		for i, u := range mem.Etcd.ListenPeerURLs {
			if !isValidURL(u) {
				return nil, fmt.Errorf("'--listen-peer-urls' has valid URL %q", u)
			}
			listenPeerPorts[i], err = getPort(u)
			if err != nil {
				return nil, fmt.Errorf("'--listen-peer-urls' has no port %q", u)
			}
		}
		for j, u := range mem.Etcd.AdvertisePeerURLs {
			if !isValidURL(u) {
				return nil, fmt.Errorf("'--initial-advertise-peer-urls' has valid URL %q", u)
			}
			port, err = getPort(u)
			if err != nil {
				return nil, fmt.Errorf("'--initial-advertise-peer-urls' has no port %q", u)
			}
			if mem.EtcdPeerProxy && listenPeerPorts[j] == port {
				return nil, fmt.Errorf("clus.Members[%d] requires peer port proxy, but advertise port %q conflicts with listener port %q", i, port, listenPeerPorts[j])
			}
		}
		if !strings.HasPrefix(mem.EtcdLogPath, mem.BaseDir) {
			return nil, fmt.Errorf("EtcdLogPath must be prefixed with BaseDir (got %q)", mem.EtcdLogPath)
		}
		if !strings.HasPrefix(mem.Etcd.DataDir, mem.BaseDir) {
			return nil, fmt.Errorf("Etcd.DataDir must be prefixed with BaseDir (got %q)", mem.Etcd.DataDir)
		}
		if !strings.HasPrefix(mem.Etcd.WALDir, mem.BaseDir) {
			return nil, fmt.Errorf("Etcd.WALDir must be prefixed with BaseDir (got %q)", mem.Etcd.WALDir)
		}
		if mem.Etcd.PeerAutoTLS && mem.Etcd.PeerCertFile != "" {
			return nil, fmt.Errorf("Etcd.PeerAutoTLS 'true', but Etcd.PeerCertFile is %q", mem.Etcd.PeerCertFile)
		}
		if mem.Etcd.PeerAutoTLS && mem.Etcd.PeerKeyFile != "" {
			return nil, fmt.Errorf("Etcd.PeerAutoTLS 'true', but Etcd.PeerKeyFile is %q", mem.Etcd.PeerKeyFile)
		}
		if mem.Etcd.PeerAutoTLS && mem.Etcd.PeerTrustedCAFile != "" {
			return nil, fmt.Errorf("Etcd.PeerAutoTLS 'true', but Etcd.PeerTrustedCAFile is %q", mem.Etcd.PeerTrustedCAFile)
		}
		if mem.Etcd.ClientAutoTLS && mem.Etcd.ClientCertFile != "" {
			return nil, fmt.Errorf("Etcd.ClientAutoTLS 'true', but Etcd.ClientCertFile is %q", mem.Etcd.ClientCertFile)
		}
		if mem.Etcd.ClientAutoTLS && mem.Etcd.ClientKeyFile != "" {
			return nil, fmt.Errorf("Etcd.ClientAutoTLS 'true', but Etcd.ClientKeyFile is %q", mem.Etcd.ClientKeyFile)
		}
		if mem.Etcd.ClientAutoTLS && mem.Etcd.ClientTrustedCAFile != "" {
			return nil, fmt.Errorf("Etcd.ClientAutoTLS 'true', but Etcd.ClientTrustedCAFile is %q", mem.Etcd.ClientTrustedCAFile)
		}
		if mem.Etcd.PeerClientCertAuth && mem.Etcd.PeerCertFile == "" {
			return nil, fmt.Errorf("Etcd.PeerClientCertAuth 'true', but Etcd.PeerCertFile is %q", mem.Etcd.PeerCertFile)
		}
		if mem.Etcd.PeerClientCertAuth && mem.Etcd.PeerKeyFile == "" {
			return nil, fmt.Errorf("Etcd.PeerClientCertAuth 'true', but Etcd.PeerKeyFile is %q", mem.Etcd.PeerCertFile)
		}
		if mem.Etcd.PeerClientCertAuth && mem.Etcd.PeerTrustedCAFile == "" {
			return nil, fmt.Errorf("Etcd.PeerClientCertAuth 'true', but Etcd.PeerTrustedCAFile is %q", mem.Etcd.PeerCertFile)
		}
		if !mem.Etcd.PeerClientCertAuth && mem.Etcd.PeerCertFile != "" {
			return nil, fmt.Errorf("Etcd.PeerClientCertAuth 'false', but Etcd.PeerCertFile is %q", mem.Etcd.PeerCertFile)
		}
		if !mem.Etcd.PeerClientCertAuth && mem.Etcd.PeerKeyFile != "" {
			return nil, fmt.Errorf("Etcd.PeerClientCertAuth 'false', but Etcd.PeerKeyFile is %q", mem.Etcd.PeerCertFile)
		}
		if !mem.Etcd.PeerClientCertAuth && mem.Etcd.PeerTrustedCAFile != "" {
			return nil, fmt.Errorf("Etcd.PeerClientCertAuth 'false', but Etcd.PeerTrustedCAFile is %q", mem.Etcd.PeerTrustedCAFile)
		}
		if mem.Etcd.PeerClientCertAuth && mem.Etcd.PeerAutoTLS {
			return nil, fmt.Errorf("Etcd.PeerClientCertAuth and Etcd.PeerAutoTLS cannot be both 'true'")
		}
		if (mem.Etcd.PeerCertFile == "") != (mem.Etcd.PeerKeyFile == "") {
			return nil, fmt.Errorf("Both Etcd.PeerCertFile %q and Etcd.PeerKeyFile %q must be either empty or non-empty", mem.Etcd.PeerCertFile, mem.Etcd.PeerKeyFile)
		}
		if mem.Etcd.ClientCertAuth && mem.Etcd.ClientAutoTLS {
			return nil, fmt.Errorf("Etcd.ClientCertAuth and Etcd.ClientAutoTLS cannot be both 'true'")
		}
		if mem.Etcd.ClientCertAuth && mem.Etcd.ClientCertFile == "" {
			return nil, fmt.Errorf("Etcd.ClientCertAuth 'true', but Etcd.ClientCertFile is %q", mem.Etcd.PeerCertFile)
		}
		if mem.Etcd.ClientCertAuth && mem.Etcd.ClientKeyFile == "" {
			return nil, fmt.Errorf("Etcd.ClientCertAuth 'true', but Etcd.ClientKeyFile is %q", mem.Etcd.PeerCertFile)
		}
		if mem.Etcd.ClientCertAuth && mem.Etcd.ClientTrustedCAFile == "" {
			return nil, fmt.Errorf("Etcd.ClientCertAuth 'true', but Etcd.ClientTrustedCAFile is %q", mem.Etcd.ClientTrustedCAFile)
		}
		if !mem.Etcd.ClientCertAuth && mem.Etcd.ClientCertFile != "" {
			return nil, fmt.Errorf("Etcd.ClientCertAuth 'false', but Etcd.ClientCertFile is %q", mem.Etcd.PeerCertFile)
		}
		if !mem.Etcd.ClientCertAuth && mem.Etcd.ClientKeyFile != "" {
			return nil, fmt.Errorf("Etcd.ClientCertAuth 'false', but Etcd.ClientKeyFile is %q", mem.Etcd.PeerCertFile)
		}
		if !mem.Etcd.ClientCertAuth && mem.Etcd.ClientTrustedCAFile != "" {
			return nil, fmt.Errorf("Etcd.ClientCertAuth 'false', but Etcd.ClientTrustedCAFile is %q", mem.Etcd.PeerCertFile)
		}
		if (mem.Etcd.ClientCertFile == "") != (mem.Etcd.ClientKeyFile == "") {
			return nil, fmt.Errorf("Both Etcd.ClientCertFile %q and Etcd.ClientKeyFile %q must be either empty or non-empty", mem.Etcd.ClientCertFile, mem.Etcd.ClientKeyFile)
		}
		peerTLS := mem.Etcd.PeerAutoTLS || (mem.Etcd.PeerClientCertAuth && mem.Etcd.PeerCertFile != "" && mem.Etcd.PeerKeyFile != "" && mem.Etcd.PeerTrustedCAFile != "")
		if peerTLS {
			for _, cu := range mem.Etcd.ListenPeerURLs {
				var u *url.URL
				u, err = url.Parse(cu)
				if err != nil {
					return nil, err
				}
				if u.Scheme != "https" {
					return nil, fmt.Errorf("peer TLS is enabled with wrong scheme %q", cu)
				}
			}
			for _, cu := range mem.Etcd.AdvertisePeerURLs {
				var u *url.URL
				u, err = url.Parse(cu)
				if err != nil {
					return nil, err
				}
				if u.Scheme != "https" {
					return nil, fmt.Errorf("peer TLS is enabled with wrong scheme %q", cu)
				}
			}
			clus.Members[i].PeerCertPath = mem.Etcd.PeerCertFile
			if mem.Etcd.PeerCertFile != "" {
				var data []byte
				data, err = ioutil.ReadFile(mem.Etcd.PeerCertFile)
				if err != nil {
					return nil, fmt.Errorf("failed to read %q (%v)", mem.Etcd.PeerCertFile, err)
				}
				clus.Members[i].PeerCertData = string(data)
			}
			clus.Members[i].PeerKeyPath = mem.Etcd.PeerKeyFile
			if mem.Etcd.PeerKeyFile != "" {
				var data []byte
				data, err = ioutil.ReadFile(mem.Etcd.PeerKeyFile)
				if err != nil {
					return nil, fmt.Errorf("failed to read %q (%v)", mem.Etcd.PeerKeyFile, err)
				}
				clus.Members[i].PeerCertData = string(data)
			}
			clus.Members[i].PeerTrustedCAPath = mem.Etcd.PeerTrustedCAFile
			if mem.Etcd.PeerTrustedCAFile != "" {
				var data []byte
				data, err = ioutil.ReadFile(mem.Etcd.PeerTrustedCAFile)
				if err != nil {
					return nil, fmt.Errorf("failed to read %q (%v)", mem.Etcd.PeerTrustedCAFile, err)
				}
				clus.Members[i].PeerCertData = string(data)
			}
		}
		clientTLS := mem.Etcd.ClientAutoTLS || (mem.Etcd.ClientCertAuth && mem.Etcd.ClientCertFile != "" && mem.Etcd.ClientKeyFile != "" && mem.Etcd.ClientTrustedCAFile != "")
		if clientTLS {
			for _, cu := range mem.Etcd.ListenClientURLs {
				var u *url.URL
				u, err = url.Parse(cu)
				if err != nil {
					return nil, err
				}
				if u.Scheme != "https" {
					return nil, fmt.Errorf("client TLS is enabled with wrong scheme %q", cu)
				}
			}
			for _, cu := range mem.Etcd.AdvertiseClientURLs {
				var u *url.URL
				u, err = url.Parse(cu)
				if err != nil {
					return nil, err
				}
				if u.Scheme != "https" {
					return nil, fmt.Errorf("client TLS is enabled with wrong scheme %q", cu)
				}
			}
			clus.Members[i].ClientCertPath = mem.Etcd.ClientCertFile
			if mem.Etcd.ClientCertFile != "" {
				var data []byte
				data, err = ioutil.ReadFile(mem.Etcd.ClientCertFile)
				if err != nil {
					return nil, fmt.Errorf("failed to read %q (%v)", mem.Etcd.ClientCertFile, err)
				}
				clus.Members[i].ClientCertData = string(data)
			}
			clus.Members[i].ClientKeyPath = mem.Etcd.ClientKeyFile
			if mem.Etcd.ClientKeyFile != "" {
				var data []byte
				data, err = ioutil.ReadFile(mem.Etcd.ClientKeyFile)
				if err != nil {
					return nil, fmt.Errorf("failed to read %q (%v)", mem.Etcd.ClientKeyFile, err)
				}
				clus.Members[i].ClientCertData = string(data)
			}
			clus.Members[i].ClientTrustedCAPath = mem.Etcd.ClientTrustedCAFile
			if mem.Etcd.ClientTrustedCAFile != "" {
				var data []byte
				data, err = ioutil.ReadFile(mem.Etcd.ClientTrustedCAFile)
				if err != nil {
					return nil, fmt.Errorf("failed to read %q (%v)", mem.Etcd.ClientTrustedCAFile, err)
				}
				clus.Members[i].ClientCertData = string(data)
			}
		}
	}
	if len(clus.Tester.Cases) == 0 {
		return nil, errors.New("Cases not found")
	}
	if clus.Tester.DelayLatencyMs <= clus.Tester.DelayLatencyMsRv*5 {
		return nil, fmt.Errorf("delay latency %d ms must be greater than 5x of delay latency random variable %d ms", clus.Tester.DelayLatencyMs, clus.Tester.DelayLatencyMsRv)
	}
	if clus.Tester.UpdatedDelayLatencyMs == 0 {
		clus.Tester.UpdatedDelayLatencyMs = clus.Tester.DelayLatencyMs
	}
	for _, v := range clus.Tester.Cases {
		if _, ok := rpcpb.Case_value[v]; !ok {
			return nil, fmt.Errorf("%q is not defined in 'rpcpb.Case_value'", v)
		}
	}
	for _, v := range clus.Tester.Stressers {
		if _, ok := rpcpb.Stresser_value[v]; !ok {
			return nil, fmt.Errorf("Stresser is unknown; got %q", v)
		}
	}
	for _, v := range clus.Tester.Checkers {
		if _, ok := rpcpb.Checker_value[v]; !ok {
			return nil, fmt.Errorf("Checker is unknown; got %q", v)
		}
	}
	if clus.Tester.StressKeySuffixRangeTxn > 100 {
		return nil, fmt.Errorf("StressKeySuffixRangeTxn maximum value is 100, got %v", clus.Tester.StressKeySuffixRangeTxn)
	}
	if clus.Tester.StressKeyTxnOps > 64 {
		return nil, fmt.Errorf("StressKeyTxnOps maximum value is 64, got %v", clus.Tester.StressKeyTxnOps)
	}
	return clus, err
}
