package rpcpb

import (
	"reflect"
	"testing"
)

func TestEtcdFlags(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg := &Etcd{Name: "s1", DataDir: "/tmp/etcd-agent-data-1/etcd.data", WALDir: "/tmp/etcd-agent-data-1/etcd.data/member/wal", HeartbeatIntervalMs: 100, ElectionTimeoutMs: 1000, ListenClientURLs: []string{"https://127.0.0.1:1379"}, AdvertiseClientURLs: []string{"https://127.0.0.1:13790"}, ClientAutoTLS: true, ClientCertAuth: false, ClientCertFile: "", ClientKeyFile: "", ClientTrustedCAFile: "", ListenPeerURLs: []string{"https://127.0.0.1:1380"}, AdvertisePeerURLs: []string{"https://127.0.0.1:13800"}, PeerAutoTLS: true, PeerClientCertAuth: false, PeerCertFile: "", PeerKeyFile: "", PeerTrustedCAFile: "", InitialCluster: "s1=https://127.0.0.1:13800,s2=https://127.0.0.1:23800,s3=https://127.0.0.1:33800", InitialClusterState: "new", InitialClusterToken: "tkn", SnapshotCount: 10000, QuotaBackendBytes: 10740000000, PreVote: true, InitialCorruptCheck: true}
	exp := []string{"--name=s1", "--data-dir=/tmp/etcd-agent-data-1/etcd.data", "--wal-dir=/tmp/etcd-agent-data-1/etcd.data/member/wal", "--heartbeat-interval=100", "--election-timeout=1000", "--listen-client-urls=https://127.0.0.1:1379", "--advertise-client-urls=https://127.0.0.1:13790", "--auto-tls=true", "--client-cert-auth=false", "--listen-peer-urls=https://127.0.0.1:1380", "--initial-advertise-peer-urls=https://127.0.0.1:13800", "--peer-auto-tls=true", "--peer-client-cert-auth=false", "--initial-cluster=s1=https://127.0.0.1:13800,s2=https://127.0.0.1:23800,s3=https://127.0.0.1:33800", "--initial-cluster-state=new", "--initial-cluster-token=tkn", "--snapshot-count=10000", "--quota-backend-bytes=10740000000", "--pre-vote=true", "--experimental-initial-corrupt-check=true"}
	fs := cfg.Flags()
	if !reflect.DeepEqual(exp, fs) {
		t.Fatalf("expected %q, got %q", exp, fs)
	}
}
