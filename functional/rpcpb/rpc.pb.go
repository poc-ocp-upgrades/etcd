package rpcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"
import io "io"

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2

type Operation int32

const (
	Operation_NOT_STARTED                                 Operation = 0
	Operation_INITIAL_START_ETCD                          Operation = 10
	Operation_RESTART_ETCD                                Operation = 11
	Operation_SIGTERM_ETCD                                Operation = 20
	Operation_SIGQUIT_ETCD_AND_REMOVE_DATA                Operation = 21
	Operation_SAVE_SNAPSHOT                               Operation = 30
	Operation_RESTORE_RESTART_FROM_SNAPSHOT               Operation = 31
	Operation_RESTART_FROM_SNAPSHOT                       Operation = 32
	Operation_SIGQUIT_ETCD_AND_ARCHIVE_DATA               Operation = 40
	Operation_SIGQUIT_ETCD_AND_REMOVE_DATA_AND_STOP_AGENT Operation = 41
	Operation_BLACKHOLE_PEER_PORT_TX_RX                   Operation = 100
	Operation_UNBLACKHOLE_PEER_PORT_TX_RX                 Operation = 101
	Operation_DELAY_PEER_PORT_TX_RX                       Operation = 200
	Operation_UNDELAY_PEER_PORT_TX_RX                     Operation = 201
)

var Operation_name = map[int32]string{0: "NOT_STARTED", 10: "INITIAL_START_ETCD", 11: "RESTART_ETCD", 20: "SIGTERM_ETCD", 21: "SIGQUIT_ETCD_AND_REMOVE_DATA", 30: "SAVE_SNAPSHOT", 31: "RESTORE_RESTART_FROM_SNAPSHOT", 32: "RESTART_FROM_SNAPSHOT", 40: "SIGQUIT_ETCD_AND_ARCHIVE_DATA", 41: "SIGQUIT_ETCD_AND_REMOVE_DATA_AND_STOP_AGENT", 100: "BLACKHOLE_PEER_PORT_TX_RX", 101: "UNBLACKHOLE_PEER_PORT_TX_RX", 200: "DELAY_PEER_PORT_TX_RX", 201: "UNDELAY_PEER_PORT_TX_RX"}
var Operation_value = map[string]int32{"NOT_STARTED": 0, "INITIAL_START_ETCD": 10, "RESTART_ETCD": 11, "SIGTERM_ETCD": 20, "SIGQUIT_ETCD_AND_REMOVE_DATA": 21, "SAVE_SNAPSHOT": 30, "RESTORE_RESTART_FROM_SNAPSHOT": 31, "RESTART_FROM_SNAPSHOT": 32, "SIGQUIT_ETCD_AND_ARCHIVE_DATA": 40, "SIGQUIT_ETCD_AND_REMOVE_DATA_AND_STOP_AGENT": 41, "BLACKHOLE_PEER_PORT_TX_RX": 100, "UNBLACKHOLE_PEER_PORT_TX_RX": 101, "DELAY_PEER_PORT_TX_RX": 200, "UNDELAY_PEER_PORT_TX_RX": 201}

func (x Operation) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{0}
}

type Case int32

const (
	Case_SIGTERM_ONE_FOLLOWER                                               Case = 0
	Case_SIGTERM_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT                        Case = 1
	Case_SIGTERM_LEADER                                                     Case = 2
	Case_SIGTERM_LEADER_UNTIL_TRIGGER_SNAPSHOT                              Case = 3
	Case_SIGTERM_QUORUM                                                     Case = 4
	Case_SIGTERM_ALL                                                        Case = 5
	Case_SIGQUIT_AND_REMOVE_ONE_FOLLOWER                                    Case = 10
	Case_SIGQUIT_AND_REMOVE_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT             Case = 11
	Case_SIGQUIT_AND_REMOVE_LEADER                                          Case = 12
	Case_SIGQUIT_AND_REMOVE_LEADER_UNTIL_TRIGGER_SNAPSHOT                   Case = 13
	Case_SIGQUIT_AND_REMOVE_QUORUM_AND_RESTORE_LEADER_SNAPSHOT_FROM_SCRATCH Case = 14
	Case_BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER                             Case = 100
	Case_BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT      Case = 101
	Case_BLACKHOLE_PEER_PORT_TX_RX_LEADER                                   Case = 102
	Case_BLACKHOLE_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT            Case = 103
	Case_BLACKHOLE_PEER_PORT_TX_RX_QUORUM                                   Case = 104
	Case_BLACKHOLE_PEER_PORT_TX_RX_ALL                                      Case = 105
	Case_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER                                 Case = 200
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER                          Case = 201
	Case_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT          Case = 202
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT   Case = 203
	Case_DELAY_PEER_PORT_TX_RX_LEADER                                       Case = 204
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_LEADER                                Case = 205
	Case_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT                Case = 206
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT         Case = 207
	Case_DELAY_PEER_PORT_TX_RX_QUORUM                                       Case = 208
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_QUORUM                                Case = 209
	Case_DELAY_PEER_PORT_TX_RX_ALL                                          Case = 210
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_ALL                                   Case = 211
	Case_NO_FAIL_WITH_STRESS                                                Case = 300
	Case_NO_FAIL_WITH_NO_STRESS_FOR_LIVENESS                                Case = 301
	Case_FAILPOINTS                                                         Case = 400
	Case_EXTERNAL                                                           Case = 500
)

var Case_name = map[int32]string{0: "SIGTERM_ONE_FOLLOWER", 1: "SIGTERM_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT", 2: "SIGTERM_LEADER", 3: "SIGTERM_LEADER_UNTIL_TRIGGER_SNAPSHOT", 4: "SIGTERM_QUORUM", 5: "SIGTERM_ALL", 10: "SIGQUIT_AND_REMOVE_ONE_FOLLOWER", 11: "SIGQUIT_AND_REMOVE_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT", 12: "SIGQUIT_AND_REMOVE_LEADER", 13: "SIGQUIT_AND_REMOVE_LEADER_UNTIL_TRIGGER_SNAPSHOT", 14: "SIGQUIT_AND_REMOVE_QUORUM_AND_RESTORE_LEADER_SNAPSHOT_FROM_SCRATCH", 100: "BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER", 101: "BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT", 102: "BLACKHOLE_PEER_PORT_TX_RX_LEADER", 103: "BLACKHOLE_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT", 104: "BLACKHOLE_PEER_PORT_TX_RX_QUORUM", 105: "BLACKHOLE_PEER_PORT_TX_RX_ALL", 200: "DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER", 201: "RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER", 202: "DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT", 203: "RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT", 204: "DELAY_PEER_PORT_TX_RX_LEADER", 205: "RANDOM_DELAY_PEER_PORT_TX_RX_LEADER", 206: "DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT", 207: "RANDOM_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT", 208: "DELAY_PEER_PORT_TX_RX_QUORUM", 209: "RANDOM_DELAY_PEER_PORT_TX_RX_QUORUM", 210: "DELAY_PEER_PORT_TX_RX_ALL", 211: "RANDOM_DELAY_PEER_PORT_TX_RX_ALL", 300: "NO_FAIL_WITH_STRESS", 301: "NO_FAIL_WITH_NO_STRESS_FOR_LIVENESS", 400: "FAILPOINTS", 500: "EXTERNAL"}
var Case_value = map[string]int32{"SIGTERM_ONE_FOLLOWER": 0, "SIGTERM_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT": 1, "SIGTERM_LEADER": 2, "SIGTERM_LEADER_UNTIL_TRIGGER_SNAPSHOT": 3, "SIGTERM_QUORUM": 4, "SIGTERM_ALL": 5, "SIGQUIT_AND_REMOVE_ONE_FOLLOWER": 10, "SIGQUIT_AND_REMOVE_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT": 11, "SIGQUIT_AND_REMOVE_LEADER": 12, "SIGQUIT_AND_REMOVE_LEADER_UNTIL_TRIGGER_SNAPSHOT": 13, "SIGQUIT_AND_REMOVE_QUORUM_AND_RESTORE_LEADER_SNAPSHOT_FROM_SCRATCH": 14, "BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER": 100, "BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT": 101, "BLACKHOLE_PEER_PORT_TX_RX_LEADER": 102, "BLACKHOLE_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT": 103, "BLACKHOLE_PEER_PORT_TX_RX_QUORUM": 104, "BLACKHOLE_PEER_PORT_TX_RX_ALL": 105, "DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER": 200, "RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER": 201, "DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT": 202, "RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT": 203, "DELAY_PEER_PORT_TX_RX_LEADER": 204, "RANDOM_DELAY_PEER_PORT_TX_RX_LEADER": 205, "DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT": 206, "RANDOM_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT": 207, "DELAY_PEER_PORT_TX_RX_QUORUM": 208, "RANDOM_DELAY_PEER_PORT_TX_RX_QUORUM": 209, "DELAY_PEER_PORT_TX_RX_ALL": 210, "RANDOM_DELAY_PEER_PORT_TX_RX_ALL": 211, "NO_FAIL_WITH_STRESS": 300, "NO_FAIL_WITH_NO_STRESS_FOR_LIVENESS": 301, "FAILPOINTS": 400, "EXTERNAL": 500}

func (x Case) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(Case_name, int32(x))
}
func (Case) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{1}
}

type Stresser int32

const (
	Stresser_KV                Stresser = 0
	Stresser_LEASE             Stresser = 1
	Stresser_ELECTION_RUNNER   Stresser = 2
	Stresser_WATCH_RUNNER      Stresser = 3
	Stresser_LOCK_RACER_RUNNER Stresser = 4
	Stresser_LEASE_RUNNER      Stresser = 5
)

var Stresser_name = map[int32]string{0: "KV", 1: "LEASE", 2: "ELECTION_RUNNER", 3: "WATCH_RUNNER", 4: "LOCK_RACER_RUNNER", 5: "LEASE_RUNNER"}
var Stresser_value = map[string]int32{"KV": 0, "LEASE": 1, "ELECTION_RUNNER": 2, "WATCH_RUNNER": 3, "LOCK_RACER_RUNNER": 4, "LEASE_RUNNER": 5}

func (x Stresser) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(Stresser_name, int32(x))
}
func (Stresser) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{2}
}

type Checker int32

const (
	Checker_KV_HASH      Checker = 0
	Checker_LEASE_EXPIRE Checker = 1
	Checker_RUNNER       Checker = 2
	Checker_NO_CHECK     Checker = 3
)

var Checker_name = map[int32]string{0: "KV_HASH", 1: "LEASE_EXPIRE", 2: "RUNNER", 3: "NO_CHECK"}
var Checker_value = map[string]int32{"KV_HASH": 0, "LEASE_EXPIRE": 1, "RUNNER": 2, "NO_CHECK": 3}

func (x Checker) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(Checker_name, int32(x))
}
func (Checker) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{3}
}

type Request struct {
	Operation Operation `protobuf:"varint,1,opt,name=Operation,proto3,enum=rpcpb.Operation" json:"Operation,omitempty"`
	Member    *Member   `protobuf:"bytes,2,opt,name=Member" json:"Member,omitempty"`
	Tester    *Tester   `protobuf:"bytes,3,opt,name=Tester" json:"Tester,omitempty"`
}

func (m *Request) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Request{}
}
func (m *Request) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Request) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Request) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{0}
}

type SnapshotInfo struct {
	MemberName        string   `protobuf:"bytes,1,opt,name=MemberName,proto3" json:"MemberName,omitempty"`
	MemberClientURLs  []string `protobuf:"bytes,2,rep,name=MemberClientURLs" json:"MemberClientURLs,omitempty"`
	SnapshotPath      string   `protobuf:"bytes,3,opt,name=SnapshotPath,proto3" json:"SnapshotPath,omitempty"`
	SnapshotFileSize  string   `protobuf:"bytes,4,opt,name=SnapshotFileSize,proto3" json:"SnapshotFileSize,omitempty"`
	SnapshotTotalSize string   `protobuf:"bytes,5,opt,name=SnapshotTotalSize,proto3" json:"SnapshotTotalSize,omitempty"`
	SnapshotTotalKey  int64    `protobuf:"varint,6,opt,name=SnapshotTotalKey,proto3" json:"SnapshotTotalKey,omitempty"`
	SnapshotHash      int64    `protobuf:"varint,7,opt,name=SnapshotHash,proto3" json:"SnapshotHash,omitempty"`
	SnapshotRevision  int64    `protobuf:"varint,8,opt,name=SnapshotRevision,proto3" json:"SnapshotRevision,omitempty"`
	Took              string   `protobuf:"bytes,9,opt,name=Took,proto3" json:"Took,omitempty"`
}

func (m *SnapshotInfo) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = SnapshotInfo{}
}
func (m *SnapshotInfo) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*SnapshotInfo) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*SnapshotInfo) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{1}
}

type Response struct {
	Success      bool          `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Status       string        `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Member       *Member       `protobuf:"bytes,3,opt,name=Member" json:"Member,omitempty"`
	SnapshotInfo *SnapshotInfo `protobuf:"bytes,4,opt,name=SnapshotInfo" json:"SnapshotInfo,omitempty"`
}

func (m *Response) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Response{}
}
func (m *Response) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Response) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Response) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{2}
}

type Member struct {
	EtcdExecPath          string        `protobuf:"bytes,1,opt,name=EtcdExecPath,proto3" json:"EtcdExecPath,omitempty" yaml:"etcd-exec-path"`
	AgentAddr             string        `protobuf:"bytes,11,opt,name=AgentAddr,proto3" json:"AgentAddr,omitempty" yaml:"agent-addr"`
	FailpointHTTPAddr     string        `protobuf:"bytes,12,opt,name=FailpointHTTPAddr,proto3" json:"FailpointHTTPAddr,omitempty" yaml:"failpoint-http-addr"`
	BaseDir               string        `protobuf:"bytes,101,opt,name=BaseDir,proto3" json:"BaseDir,omitempty" yaml:"base-dir"`
	EtcdLogPath           string        `protobuf:"bytes,102,opt,name=EtcdLogPath,proto3" json:"EtcdLogPath,omitempty" yaml:"etcd-log-path"`
	EtcdClientProxy       bool          `protobuf:"varint,201,opt,name=EtcdClientProxy,proto3" json:"EtcdClientProxy,omitempty" yaml:"etcd-client-proxy"`
	EtcdPeerProxy         bool          `protobuf:"varint,202,opt,name=EtcdPeerProxy,proto3" json:"EtcdPeerProxy,omitempty" yaml:"etcd-peer-proxy"`
	EtcdClientEndpoint    string        `protobuf:"bytes,301,opt,name=EtcdClientEndpoint,proto3" json:"EtcdClientEndpoint,omitempty" yaml:"etcd-client-endpoint"`
	Etcd                  *Etcd         `protobuf:"bytes,302,opt,name=Etcd" json:"Etcd,omitempty" yaml:"etcd"`
	EtcdOnSnapshotRestore *Etcd         `protobuf:"bytes,303,opt,name=EtcdOnSnapshotRestore" json:"EtcdOnSnapshotRestore,omitempty"`
	ClientCertData        string        `protobuf:"bytes,401,opt,name=ClientCertData,proto3" json:"ClientCertData,omitempty" yaml:"client-cert-data"`
	ClientCertPath        string        `protobuf:"bytes,402,opt,name=ClientCertPath,proto3" json:"ClientCertPath,omitempty" yaml:"client-cert-path"`
	ClientKeyData         string        `protobuf:"bytes,403,opt,name=ClientKeyData,proto3" json:"ClientKeyData,omitempty" yaml:"client-key-data"`
	ClientKeyPath         string        `protobuf:"bytes,404,opt,name=ClientKeyPath,proto3" json:"ClientKeyPath,omitempty" yaml:"client-key-path"`
	ClientTrustedCAData   string        `protobuf:"bytes,405,opt,name=ClientTrustedCAData,proto3" json:"ClientTrustedCAData,omitempty" yaml:"client-trusted-ca-data"`
	ClientTrustedCAPath   string        `protobuf:"bytes,406,opt,name=ClientTrustedCAPath,proto3" json:"ClientTrustedCAPath,omitempty" yaml:"client-trusted-ca-path"`
	PeerCertData          string        `protobuf:"bytes,501,opt,name=PeerCertData,proto3" json:"PeerCertData,omitempty" yaml:"peer-cert-data"`
	PeerCertPath          string        `protobuf:"bytes,502,opt,name=PeerCertPath,proto3" json:"PeerCertPath,omitempty" yaml:"peer-cert-path"`
	PeerKeyData           string        `protobuf:"bytes,503,opt,name=PeerKeyData,proto3" json:"PeerKeyData,omitempty" yaml:"peer-key-data"`
	PeerKeyPath           string        `protobuf:"bytes,504,opt,name=PeerKeyPath,proto3" json:"PeerKeyPath,omitempty" yaml:"peer-key-path"`
	PeerTrustedCAData     string        `protobuf:"bytes,505,opt,name=PeerTrustedCAData,proto3" json:"PeerTrustedCAData,omitempty" yaml:"peer-trusted-ca-data"`
	PeerTrustedCAPath     string        `protobuf:"bytes,506,opt,name=PeerTrustedCAPath,proto3" json:"PeerTrustedCAPath,omitempty" yaml:"peer-trusted-ca-path"`
	SnapshotPath          string        `protobuf:"bytes,601,opt,name=SnapshotPath,proto3" json:"SnapshotPath,omitempty" yaml:"snapshot-path"`
	SnapshotInfo          *SnapshotInfo `protobuf:"bytes,602,opt,name=SnapshotInfo" json:"SnapshotInfo,omitempty"`
}

func (m *Member) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Member{}
}
func (m *Member) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Member) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Member) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{3}
}

type Tester struct {
	DataDir                 string   `protobuf:"bytes,1,opt,name=DataDir,proto3" json:"DataDir,omitempty" yaml:"data-dir"`
	Network                 string   `protobuf:"bytes,2,opt,name=Network,proto3" json:"Network,omitempty" yaml:"network"`
	Addr                    string   `protobuf:"bytes,3,opt,name=Addr,proto3" json:"Addr,omitempty" yaml:"addr"`
	DelayLatencyMs          uint32   `protobuf:"varint,11,opt,name=DelayLatencyMs,proto3" json:"DelayLatencyMs,omitempty" yaml:"delay-latency-ms"`
	DelayLatencyMsRv        uint32   `protobuf:"varint,12,opt,name=DelayLatencyMsRv,proto3" json:"DelayLatencyMsRv,omitempty" yaml:"delay-latency-ms-rv"`
	UpdatedDelayLatencyMs   uint32   `protobuf:"varint,13,opt,name=UpdatedDelayLatencyMs,proto3" json:"UpdatedDelayLatencyMs,omitempty" yaml:"updated-delay-latency-ms"`
	RoundLimit              int32    `protobuf:"varint,21,opt,name=RoundLimit,proto3" json:"RoundLimit,omitempty" yaml:"round-limit"`
	ExitOnCaseFail          bool     `protobuf:"varint,22,opt,name=ExitOnCaseFail,proto3" json:"ExitOnCaseFail,omitempty" yaml:"exit-on-failure"`
	EnablePprof             bool     `protobuf:"varint,23,opt,name=EnablePprof,proto3" json:"EnablePprof,omitempty" yaml:"enable-pprof"`
	CaseDelayMs             uint32   `protobuf:"varint,31,opt,name=CaseDelayMs,proto3" json:"CaseDelayMs,omitempty" yaml:"case-delay-ms"`
	CaseShuffle             bool     `protobuf:"varint,32,opt,name=CaseShuffle,proto3" json:"CaseShuffle,omitempty" yaml:"case-shuffle"`
	Cases                   []string `protobuf:"bytes,33,rep,name=Cases" json:"Cases,omitempty" yaml:"cases"`
	FailpointCommands       []string `protobuf:"bytes,34,rep,name=FailpointCommands" json:"FailpointCommands,omitempty" yaml:"failpoint-commands"`
	RunnerExecPath          string   `protobuf:"bytes,41,opt,name=RunnerExecPath,proto3" json:"RunnerExecPath,omitempty" yaml:"runner-exec-path"`
	ExternalExecPath        string   `protobuf:"bytes,42,opt,name=ExternalExecPath,proto3" json:"ExternalExecPath,omitempty" yaml:"external-exec-path"`
	Stressers               []string `protobuf:"bytes,101,rep,name=Stressers" json:"Stressers,omitempty" yaml:"stressers"`
	Checkers                []string `protobuf:"bytes,102,rep,name=Checkers" json:"Checkers,omitempty" yaml:"checkers"`
	StressKeySize           int32    `protobuf:"varint,201,opt,name=StressKeySize,proto3" json:"StressKeySize,omitempty" yaml:"stress-key-size"`
	StressKeySizeLarge      int32    `protobuf:"varint,202,opt,name=StressKeySizeLarge,proto3" json:"StressKeySizeLarge,omitempty" yaml:"stress-key-size-large"`
	StressKeySuffixRange    int32    `protobuf:"varint,203,opt,name=StressKeySuffixRange,proto3" json:"StressKeySuffixRange,omitempty" yaml:"stress-key-suffix-range"`
	StressKeySuffixRangeTxn int32    `protobuf:"varint,204,opt,name=StressKeySuffixRangeTxn,proto3" json:"StressKeySuffixRangeTxn,omitempty" yaml:"stress-key-suffix-range-txn"`
	StressKeyTxnOps         int32    `protobuf:"varint,205,opt,name=StressKeyTxnOps,proto3" json:"StressKeyTxnOps,omitempty" yaml:"stress-key-txn-ops"`
	StressClients           int32    `protobuf:"varint,301,opt,name=StressClients,proto3" json:"StressClients,omitempty" yaml:"stress-clients"`
	StressQPS               int32    `protobuf:"varint,302,opt,name=StressQPS,proto3" json:"StressQPS,omitempty" yaml:"stress-qps"`
}

func (m *Tester) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Tester{}
}
func (m *Tester) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Tester) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Tester) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{4}
}

type Etcd struct {
	Name                string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty" yaml:"name"`
	DataDir             string   `protobuf:"bytes,2,opt,name=DataDir,proto3" json:"DataDir,omitempty" yaml:"data-dir"`
	WALDir              string   `protobuf:"bytes,3,opt,name=WALDir,proto3" json:"WALDir,omitempty" yaml:"wal-dir"`
	HeartbeatIntervalMs int64    `protobuf:"varint,11,opt,name=HeartbeatIntervalMs,proto3" json:"HeartbeatIntervalMs,omitempty" yaml:"heartbeat-interval"`
	ElectionTimeoutMs   int64    `protobuf:"varint,12,opt,name=ElectionTimeoutMs,proto3" json:"ElectionTimeoutMs,omitempty" yaml:"election-timeout"`
	ListenClientURLs    []string `protobuf:"bytes,21,rep,name=ListenClientURLs" json:"ListenClientURLs,omitempty" yaml:"listen-client-urls"`
	AdvertiseClientURLs []string `protobuf:"bytes,22,rep,name=AdvertiseClientURLs" json:"AdvertiseClientURLs,omitempty" yaml:"advertise-client-urls"`
	ClientAutoTLS       bool     `protobuf:"varint,23,opt,name=ClientAutoTLS,proto3" json:"ClientAutoTLS,omitempty" yaml:"auto-tls"`
	ClientCertAuth      bool     `protobuf:"varint,24,opt,name=ClientCertAuth,proto3" json:"ClientCertAuth,omitempty" yaml:"client-cert-auth"`
	ClientCertFile      string   `protobuf:"bytes,25,opt,name=ClientCertFile,proto3" json:"ClientCertFile,omitempty" yaml:"cert-file"`
	ClientKeyFile       string   `protobuf:"bytes,26,opt,name=ClientKeyFile,proto3" json:"ClientKeyFile,omitempty" yaml:"key-file"`
	ClientTrustedCAFile string   `protobuf:"bytes,27,opt,name=ClientTrustedCAFile,proto3" json:"ClientTrustedCAFile,omitempty" yaml:"trusted-ca-file"`
	ListenPeerURLs      []string `protobuf:"bytes,31,rep,name=ListenPeerURLs" json:"ListenPeerURLs,omitempty" yaml:"listen-peer-urls"`
	AdvertisePeerURLs   []string `protobuf:"bytes,32,rep,name=AdvertisePeerURLs" json:"AdvertisePeerURLs,omitempty" yaml:"initial-advertise-peer-urls"`
	PeerAutoTLS         bool     `protobuf:"varint,33,opt,name=PeerAutoTLS,proto3" json:"PeerAutoTLS,omitempty" yaml:"peer-auto-tls"`
	PeerClientCertAuth  bool     `protobuf:"varint,34,opt,name=PeerClientCertAuth,proto3" json:"PeerClientCertAuth,omitempty" yaml:"peer-client-cert-auth"`
	PeerCertFile        string   `protobuf:"bytes,35,opt,name=PeerCertFile,proto3" json:"PeerCertFile,omitempty" yaml:"peer-cert-file"`
	PeerKeyFile         string   `protobuf:"bytes,36,opt,name=PeerKeyFile,proto3" json:"PeerKeyFile,omitempty" yaml:"peer-key-file"`
	PeerTrustedCAFile   string   `protobuf:"bytes,37,opt,name=PeerTrustedCAFile,proto3" json:"PeerTrustedCAFile,omitempty" yaml:"peer-trusted-ca-file"`
	InitialCluster      string   `protobuf:"bytes,41,opt,name=InitialCluster,proto3" json:"InitialCluster,omitempty" yaml:"initial-cluster"`
	InitialClusterState string   `protobuf:"bytes,42,opt,name=InitialClusterState,proto3" json:"InitialClusterState,omitempty" yaml:"initial-cluster-state"`
	InitialClusterToken string   `protobuf:"bytes,43,opt,name=InitialClusterToken,proto3" json:"InitialClusterToken,omitempty" yaml:"initial-cluster-token"`
	SnapshotCount       int64    `protobuf:"varint,51,opt,name=SnapshotCount,proto3" json:"SnapshotCount,omitempty" yaml:"snapshot-count"`
	QuotaBackendBytes   int64    `protobuf:"varint,52,opt,name=QuotaBackendBytes,proto3" json:"QuotaBackendBytes,omitempty" yaml:"quota-backend-bytes"`
	PreVote             bool     `protobuf:"varint,63,opt,name=PreVote,proto3" json:"PreVote,omitempty" yaml:"pre-vote"`
	InitialCorruptCheck bool     `protobuf:"varint,64,opt,name=InitialCorruptCheck,proto3" json:"InitialCorruptCheck,omitempty" yaml:"initial-corrupt-check"`
}

func (m *Etcd) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Etcd{}
}
func (m *Etcd) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Etcd) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Etcd) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{5}
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterType((*Request)(nil), "rpcpb.Request")
	proto.RegisterType((*SnapshotInfo)(nil), "rpcpb.SnapshotInfo")
	proto.RegisterType((*Response)(nil), "rpcpb.Response")
	proto.RegisterType((*Member)(nil), "rpcpb.Member")
	proto.RegisterType((*Tester)(nil), "rpcpb.Tester")
	proto.RegisterType((*Etcd)(nil), "rpcpb.Etcd")
	proto.RegisterEnum("rpcpb.Operation", Operation_name, Operation_value)
	proto.RegisterEnum("rpcpb.Case", Case_name, Case_value)
	proto.RegisterEnum("rpcpb.Stresser", Stresser_name, Stresser_value)
	proto.RegisterEnum("rpcpb.Checker", Checker_name, Checker_value)
}

var _ context.Context
var _ grpc.ClientConn

const _ = grpc.SupportPackageIsVersion4

type TransportClient interface {
	Transport(ctx context.Context, opts ...grpc.CallOption) (Transport_TransportClient, error)
}
type transportClient struct{ cc *grpc.ClientConn }

func NewTransportClient(cc *grpc.ClientConn) TransportClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &transportClient{cc}
}
func (c *transportClient) Transport(ctx context.Context, opts ...grpc.CallOption) (Transport_TransportClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stream, err := grpc.NewClientStream(ctx, &_Transport_serviceDesc.Streams[0], c.cc, "/rpcpb.Transport/Transport", opts...)
	if err != nil {
		return nil, err
	}
	x := &transportTransportClient{stream}
	return x, nil
}

type Transport_TransportClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}
type transportTransportClient struct{ grpc.ClientStream }

func (x *transportTransportClient) Send(m *Request) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return x.ClientStream.SendMsg(m)
}
func (x *transportTransportClient) Recv() (*Response, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

type TransportServer interface {
	Transport(Transport_TransportServer) error
}

func RegisterTransportServer(s *grpc.Server, srv TransportServer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.RegisterService(&_Transport_serviceDesc, srv)
}
func _Transport_Transport_Handler(srv interface{}, stream grpc.ServerStream) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return srv.(TransportServer).Transport(&transportTransportServer{stream})
}

type Transport_TransportServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}
type transportTransportServer struct{ grpc.ServerStream }

func (x *transportTransportServer) Send(m *Response) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return x.ServerStream.SendMsg(m)
}
func (x *transportTransportServer) Recv() (*Request, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Transport_serviceDesc = grpc.ServiceDesc{ServiceName: "rpcpb.Transport", HandlerType: (*TransportServer)(nil), Methods: []grpc.MethodDesc{}, Streams: []grpc.StreamDesc{{StreamName: "Transport", Handler: _Transport_Transport_Handler, ServerStreams: true, ClientStreams: true}}, Metadata: "rpcpb/rpc.proto"}

func (m *Request) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Operation != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Operation))
	}
	if m.Member != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Member.Size()))
		n1, err := m.Member.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Tester != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Tester.Size()))
		n2, err := m.Tester.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *SnapshotInfo) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *SnapshotInfo) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MemberName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.MemberName)))
		i += copy(dAtA[i:], m.MemberName)
	}
	if len(m.MemberClientURLs) > 0 {
		for _, s := range m.MemberClientURLs {
			dAtA[i] = 0x12
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.SnapshotPath) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.SnapshotPath)))
		i += copy(dAtA[i:], m.SnapshotPath)
	}
	if len(m.SnapshotFileSize) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.SnapshotFileSize)))
		i += copy(dAtA[i:], m.SnapshotFileSize)
	}
	if len(m.SnapshotTotalSize) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.SnapshotTotalSize)))
		i += copy(dAtA[i:], m.SnapshotTotalSize)
	}
	if m.SnapshotTotalKey != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SnapshotTotalKey))
	}
	if m.SnapshotHash != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SnapshotHash))
	}
	if m.SnapshotRevision != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SnapshotRevision))
	}
	if len(m.Took) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Took)))
		i += copy(dAtA[i:], m.Took)
	}
	return i, nil
}
func (m *Response) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if m.Success {
		dAtA[i] = 0x8
		i++
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Status) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	if m.Member != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Member.Size()))
		n3, err := m.Member.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.SnapshotInfo != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SnapshotInfo.Size()))
		n4, err := m.SnapshotInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *Member) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.EtcdExecPath) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.EtcdExecPath)))
		i += copy(dAtA[i:], m.EtcdExecPath)
	}
	if len(m.AgentAddr) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.AgentAddr)))
		i += copy(dAtA[i:], m.AgentAddr)
	}
	if len(m.FailpointHTTPAddr) > 0 {
		dAtA[i] = 0x62
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.FailpointHTTPAddr)))
		i += copy(dAtA[i:], m.FailpointHTTPAddr)
	}
	if len(m.BaseDir) > 0 {
		dAtA[i] = 0xaa
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.BaseDir)))
		i += copy(dAtA[i:], m.BaseDir)
	}
	if len(m.EtcdLogPath) > 0 {
		dAtA[i] = 0xb2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.EtcdLogPath)))
		i += copy(dAtA[i:], m.EtcdLogPath)
	}
	if m.EtcdClientProxy {
		dAtA[i] = 0xc8
		i++
		dAtA[i] = 0xc
		i++
		if m.EtcdClientProxy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.EtcdPeerProxy {
		dAtA[i] = 0xd0
		i++
		dAtA[i] = 0xc
		i++
		if m.EtcdPeerProxy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.EtcdClientEndpoint) > 0 {
		dAtA[i] = 0xea
		i++
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.EtcdClientEndpoint)))
		i += copy(dAtA[i:], m.EtcdClientEndpoint)
	}
	if m.Etcd != nil {
		dAtA[i] = 0xf2
		i++
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.Etcd.Size()))
		n5, err := m.Etcd.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.EtcdOnSnapshotRestore != nil {
		dAtA[i] = 0xfa
		i++
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.EtcdOnSnapshotRestore.Size()))
		n6, err := m.EtcdOnSnapshotRestore.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.ClientCertData) > 0 {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x19
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientCertData)))
		i += copy(dAtA[i:], m.ClientCertData)
	}
	if len(m.ClientCertPath) > 0 {
		dAtA[i] = 0x92
		i++
		dAtA[i] = 0x19
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientCertPath)))
		i += copy(dAtA[i:], m.ClientCertPath)
	}
	if len(m.ClientKeyData) > 0 {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x19
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientKeyData)))
		i += copy(dAtA[i:], m.ClientKeyData)
	}
	if len(m.ClientKeyPath) > 0 {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x19
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientKeyPath)))
		i += copy(dAtA[i:], m.ClientKeyPath)
	}
	if len(m.ClientTrustedCAData) > 0 {
		dAtA[i] = 0xaa
		i++
		dAtA[i] = 0x19
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientTrustedCAData)))
		i += copy(dAtA[i:], m.ClientTrustedCAData)
	}
	if len(m.ClientTrustedCAPath) > 0 {
		dAtA[i] = 0xb2
		i++
		dAtA[i] = 0x19
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientTrustedCAPath)))
		i += copy(dAtA[i:], m.ClientTrustedCAPath)
	}
	if len(m.PeerCertData) > 0 {
		dAtA[i] = 0xaa
		i++
		dAtA[i] = 0x1f
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerCertData)))
		i += copy(dAtA[i:], m.PeerCertData)
	}
	if len(m.PeerCertPath) > 0 {
		dAtA[i] = 0xb2
		i++
		dAtA[i] = 0x1f
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerCertPath)))
		i += copy(dAtA[i:], m.PeerCertPath)
	}
	if len(m.PeerKeyData) > 0 {
		dAtA[i] = 0xba
		i++
		dAtA[i] = 0x1f
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerKeyData)))
		i += copy(dAtA[i:], m.PeerKeyData)
	}
	if len(m.PeerKeyPath) > 0 {
		dAtA[i] = 0xc2
		i++
		dAtA[i] = 0x1f
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerKeyPath)))
		i += copy(dAtA[i:], m.PeerKeyPath)
	}
	if len(m.PeerTrustedCAData) > 0 {
		dAtA[i] = 0xca
		i++
		dAtA[i] = 0x1f
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerTrustedCAData)))
		i += copy(dAtA[i:], m.PeerTrustedCAData)
	}
	if len(m.PeerTrustedCAPath) > 0 {
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0x1f
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerTrustedCAPath)))
		i += copy(dAtA[i:], m.PeerTrustedCAPath)
	}
	if len(m.SnapshotPath) > 0 {
		dAtA[i] = 0xca
		i++
		dAtA[i] = 0x25
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.SnapshotPath)))
		i += copy(dAtA[i:], m.SnapshotPath)
	}
	if m.SnapshotInfo != nil {
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0x25
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SnapshotInfo.Size()))
		n7, err := m.SnapshotInfo.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func (m *Tester) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Tester) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.DataDir) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.DataDir)))
		i += copy(dAtA[i:], m.DataDir)
	}
	if len(m.Network) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Network)))
		i += copy(dAtA[i:], m.Network)
	}
	if len(m.Addr) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Addr)))
		i += copy(dAtA[i:], m.Addr)
	}
	if m.DelayLatencyMs != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.DelayLatencyMs))
	}
	if m.DelayLatencyMsRv != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.DelayLatencyMsRv))
	}
	if m.UpdatedDelayLatencyMs != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.UpdatedDelayLatencyMs))
	}
	if m.RoundLimit != 0 {
		dAtA[i] = 0xa8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.RoundLimit))
	}
	if m.ExitOnCaseFail {
		dAtA[i] = 0xb0
		i++
		dAtA[i] = 0x1
		i++
		if m.ExitOnCaseFail {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.EnablePprof {
		dAtA[i] = 0xb8
		i++
		dAtA[i] = 0x1
		i++
		if m.EnablePprof {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.CaseDelayMs != 0 {
		dAtA[i] = 0xf8
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.CaseDelayMs))
	}
	if m.CaseShuffle {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x2
		i++
		if m.CaseShuffle {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Cases) > 0 {
		for _, s := range m.Cases {
			dAtA[i] = 0x8a
			i++
			dAtA[i] = 0x2
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.FailpointCommands) > 0 {
		for _, s := range m.FailpointCommands {
			dAtA[i] = 0x92
			i++
			dAtA[i] = 0x2
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.RunnerExecPath) > 0 {
		dAtA[i] = 0xca
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.RunnerExecPath)))
		i += copy(dAtA[i:], m.RunnerExecPath)
	}
	if len(m.ExternalExecPath) > 0 {
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ExternalExecPath)))
		i += copy(dAtA[i:], m.ExternalExecPath)
	}
	if len(m.Stressers) > 0 {
		for _, s := range m.Stressers {
			dAtA[i] = 0xaa
			i++
			dAtA[i] = 0x6
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Checkers) > 0 {
		for _, s := range m.Checkers {
			dAtA[i] = 0xb2
			i++
			dAtA[i] = 0x6
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.StressKeySize != 0 {
		dAtA[i] = 0xc8
		i++
		dAtA[i] = 0xc
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.StressKeySize))
	}
	if m.StressKeySizeLarge != 0 {
		dAtA[i] = 0xd0
		i++
		dAtA[i] = 0xc
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.StressKeySizeLarge))
	}
	if m.StressKeySuffixRange != 0 {
		dAtA[i] = 0xd8
		i++
		dAtA[i] = 0xc
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.StressKeySuffixRange))
	}
	if m.StressKeySuffixRangeTxn != 0 {
		dAtA[i] = 0xe0
		i++
		dAtA[i] = 0xc
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.StressKeySuffixRangeTxn))
	}
	if m.StressKeyTxnOps != 0 {
		dAtA[i] = 0xe8
		i++
		dAtA[i] = 0xc
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.StressKeyTxnOps))
	}
	if m.StressClients != 0 {
		dAtA[i] = 0xe8
		i++
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.StressClients))
	}
	if m.StressQPS != 0 {
		dAtA[i] = 0xf0
		i++
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.StressQPS))
	}
	return i, nil
}
func (m *Etcd) Marshal() (dAtA []byte, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}
func (m *Etcd) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.DataDir) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.DataDir)))
		i += copy(dAtA[i:], m.DataDir)
	}
	if len(m.WALDir) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.WALDir)))
		i += copy(dAtA[i:], m.WALDir)
	}
	if m.HeartbeatIntervalMs != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.HeartbeatIntervalMs))
	}
	if m.ElectionTimeoutMs != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.ElectionTimeoutMs))
	}
	if len(m.ListenClientURLs) > 0 {
		for _, s := range m.ListenClientURLs {
			dAtA[i] = 0xaa
			i++
			dAtA[i] = 0x1
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.AdvertiseClientURLs) > 0 {
		for _, s := range m.AdvertiseClientURLs {
			dAtA[i] = 0xb2
			i++
			dAtA[i] = 0x1
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.ClientAutoTLS {
		dAtA[i] = 0xb8
		i++
		dAtA[i] = 0x1
		i++
		if m.ClientAutoTLS {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ClientCertAuth {
		dAtA[i] = 0xc0
		i++
		dAtA[i] = 0x1
		i++
		if m.ClientCertAuth {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.ClientCertFile) > 0 {
		dAtA[i] = 0xca
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientCertFile)))
		i += copy(dAtA[i:], m.ClientCertFile)
	}
	if len(m.ClientKeyFile) > 0 {
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientKeyFile)))
		i += copy(dAtA[i:], m.ClientKeyFile)
	}
	if len(m.ClientTrustedCAFile) > 0 {
		dAtA[i] = 0xda
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.ClientTrustedCAFile)))
		i += copy(dAtA[i:], m.ClientTrustedCAFile)
	}
	if len(m.ListenPeerURLs) > 0 {
		for _, s := range m.ListenPeerURLs {
			dAtA[i] = 0xfa
			i++
			dAtA[i] = 0x1
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.AdvertisePeerURLs) > 0 {
		for _, s := range m.AdvertisePeerURLs {
			dAtA[i] = 0x82
			i++
			dAtA[i] = 0x2
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.PeerAutoTLS {
		dAtA[i] = 0x88
		i++
		dAtA[i] = 0x2
		i++
		if m.PeerAutoTLS {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.PeerClientCertAuth {
		dAtA[i] = 0x90
		i++
		dAtA[i] = 0x2
		i++
		if m.PeerClientCertAuth {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.PeerCertFile) > 0 {
		dAtA[i] = 0x9a
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerCertFile)))
		i += copy(dAtA[i:], m.PeerCertFile)
	}
	if len(m.PeerKeyFile) > 0 {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerKeyFile)))
		i += copy(dAtA[i:], m.PeerKeyFile)
	}
	if len(m.PeerTrustedCAFile) > 0 {
		dAtA[i] = 0xaa
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.PeerTrustedCAFile)))
		i += copy(dAtA[i:], m.PeerTrustedCAFile)
	}
	if len(m.InitialCluster) > 0 {
		dAtA[i] = 0xca
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.InitialCluster)))
		i += copy(dAtA[i:], m.InitialCluster)
	}
	if len(m.InitialClusterState) > 0 {
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.InitialClusterState)))
		i += copy(dAtA[i:], m.InitialClusterState)
	}
	if len(m.InitialClusterToken) > 0 {
		dAtA[i] = 0xda
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.InitialClusterToken)))
		i += copy(dAtA[i:], m.InitialClusterToken)
	}
	if m.SnapshotCount != 0 {
		dAtA[i] = 0x98
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.SnapshotCount))
	}
	if m.QuotaBackendBytes != 0 {
		dAtA[i] = 0xa0
		i++
		dAtA[i] = 0x3
		i++
		i = encodeVarintRpc(dAtA, i, uint64(m.QuotaBackendBytes))
	}
	if m.PreVote {
		dAtA[i] = 0xf8
		i++
		dAtA[i] = 0x3
		i++
		if m.PreVote {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.InitialCorruptCheck {
		dAtA[i] = 0x80
		i++
		dAtA[i] = 0x4
		i++
		if m.InitialCorruptCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}
func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Operation != 0 {
		n += 1 + sovRpc(uint64(m.Operation))
	}
	if m.Member != nil {
		l = m.Member.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Tester != nil {
		l = m.Tester.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *SnapshotInfo) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.MemberName)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.MemberClientURLs) > 0 {
		for _, s := range m.MemberClientURLs {
			l = len(s)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	l = len(m.SnapshotPath)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.SnapshotFileSize)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.SnapshotTotalSize)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.SnapshotTotalKey != 0 {
		n += 1 + sovRpc(uint64(m.SnapshotTotalKey))
	}
	if m.SnapshotHash != 0 {
		n += 1 + sovRpc(uint64(m.SnapshotHash))
	}
	if m.SnapshotRevision != 0 {
		n += 1 + sovRpc(uint64(m.SnapshotRevision))
	}
	l = len(m.Took)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *Response) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Member != nil {
		l = m.Member.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.SnapshotInfo != nil {
		l = m.SnapshotInfo.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *Member) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.EtcdExecPath)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.AgentAddr)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.FailpointHTTPAddr)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.BaseDir)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.EtcdLogPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	if m.EtcdClientProxy {
		n += 3
	}
	if m.EtcdPeerProxy {
		n += 3
	}
	l = len(m.EtcdClientEndpoint)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	if m.Etcd != nil {
		l = m.Etcd.Size()
		n += 2 + l + sovRpc(uint64(l))
	}
	if m.EtcdOnSnapshotRestore != nil {
		l = m.EtcdOnSnapshotRestore.Size()
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.ClientCertData)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.ClientCertPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.ClientKeyData)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.ClientKeyPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.ClientTrustedCAData)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.ClientTrustedCAPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.PeerCertData)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.PeerCertPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.PeerKeyData)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.PeerKeyPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.PeerTrustedCAData)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.PeerTrustedCAPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.SnapshotPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	if m.SnapshotInfo != nil {
		l = m.SnapshotInfo.Size()
		n += 2 + l + sovRpc(uint64(l))
	}
	return n
}
func (m *Tester) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.DataDir)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.DelayLatencyMs != 0 {
		n += 1 + sovRpc(uint64(m.DelayLatencyMs))
	}
	if m.DelayLatencyMsRv != 0 {
		n += 1 + sovRpc(uint64(m.DelayLatencyMsRv))
	}
	if m.UpdatedDelayLatencyMs != 0 {
		n += 1 + sovRpc(uint64(m.UpdatedDelayLatencyMs))
	}
	if m.RoundLimit != 0 {
		n += 2 + sovRpc(uint64(m.RoundLimit))
	}
	if m.ExitOnCaseFail {
		n += 3
	}
	if m.EnablePprof {
		n += 3
	}
	if m.CaseDelayMs != 0 {
		n += 2 + sovRpc(uint64(m.CaseDelayMs))
	}
	if m.CaseShuffle {
		n += 3
	}
	if len(m.Cases) > 0 {
		for _, s := range m.Cases {
			l = len(s)
			n += 2 + l + sovRpc(uint64(l))
		}
	}
	if len(m.FailpointCommands) > 0 {
		for _, s := range m.FailpointCommands {
			l = len(s)
			n += 2 + l + sovRpc(uint64(l))
		}
	}
	l = len(m.RunnerExecPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.ExternalExecPath)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	if len(m.Stressers) > 0 {
		for _, s := range m.Stressers {
			l = len(s)
			n += 2 + l + sovRpc(uint64(l))
		}
	}
	if len(m.Checkers) > 0 {
		for _, s := range m.Checkers {
			l = len(s)
			n += 2 + l + sovRpc(uint64(l))
		}
	}
	if m.StressKeySize != 0 {
		n += 2 + sovRpc(uint64(m.StressKeySize))
	}
	if m.StressKeySizeLarge != 0 {
		n += 2 + sovRpc(uint64(m.StressKeySizeLarge))
	}
	if m.StressKeySuffixRange != 0 {
		n += 2 + sovRpc(uint64(m.StressKeySuffixRange))
	}
	if m.StressKeySuffixRangeTxn != 0 {
		n += 2 + sovRpc(uint64(m.StressKeySuffixRangeTxn))
	}
	if m.StressKeyTxnOps != 0 {
		n += 2 + sovRpc(uint64(m.StressKeyTxnOps))
	}
	if m.StressClients != 0 {
		n += 2 + sovRpc(uint64(m.StressClients))
	}
	if m.StressQPS != 0 {
		n += 2 + sovRpc(uint64(m.StressQPS))
	}
	return n
}
func (m *Etcd) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.DataDir)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	l = len(m.WALDir)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.HeartbeatIntervalMs != 0 {
		n += 1 + sovRpc(uint64(m.HeartbeatIntervalMs))
	}
	if m.ElectionTimeoutMs != 0 {
		n += 1 + sovRpc(uint64(m.ElectionTimeoutMs))
	}
	if len(m.ListenClientURLs) > 0 {
		for _, s := range m.ListenClientURLs {
			l = len(s)
			n += 2 + l + sovRpc(uint64(l))
		}
	}
	if len(m.AdvertiseClientURLs) > 0 {
		for _, s := range m.AdvertiseClientURLs {
			l = len(s)
			n += 2 + l + sovRpc(uint64(l))
		}
	}
	if m.ClientAutoTLS {
		n += 3
	}
	if m.ClientCertAuth {
		n += 3
	}
	l = len(m.ClientCertFile)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.ClientKeyFile)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.ClientTrustedCAFile)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	if len(m.ListenPeerURLs) > 0 {
		for _, s := range m.ListenPeerURLs {
			l = len(s)
			n += 2 + l + sovRpc(uint64(l))
		}
	}
	if len(m.AdvertisePeerURLs) > 0 {
		for _, s := range m.AdvertisePeerURLs {
			l = len(s)
			n += 2 + l + sovRpc(uint64(l))
		}
	}
	if m.PeerAutoTLS {
		n += 3
	}
	if m.PeerClientCertAuth {
		n += 3
	}
	l = len(m.PeerCertFile)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.PeerKeyFile)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.PeerTrustedCAFile)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.InitialCluster)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.InitialClusterState)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	l = len(m.InitialClusterToken)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	if m.SnapshotCount != 0 {
		n += 2 + sovRpc(uint64(m.SnapshotCount))
	}
	if m.QuotaBackendBytes != 0 {
		n += 2 + sovRpc(uint64(m.QuotaBackendBytes))
	}
	if m.PreVote {
		n += 3
	}
	if m.InitialCorruptCheck {
		n += 3
	}
	return n
}
func sovRpc(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRpc(x uint64) (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			m.Operation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operation |= (Operation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Member == nil {
				m.Member = &Member{}
			}
			if err := m.Member.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tester", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tester == nil {
				m.Tester = &Tester{}
			}
			if err := m.Tester.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SnapshotInfo) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SnapshotInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MemberName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberClientURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MemberClientURLs = append(m.MemberClientURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotFileSize", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotFileSize = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotTotalSize", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotTotalSize = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotTotalKey", wireType)
			}
			m.SnapshotTotalKey = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SnapshotTotalKey |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotHash", wireType)
			}
			m.SnapshotHash = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SnapshotHash |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotRevision", wireType)
			}
			m.SnapshotRevision = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SnapshotRevision |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Took", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Took = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Member == nil {
				m.Member = &Member{}
			}
			if err := m.Member.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SnapshotInfo == nil {
				m.SnapshotInfo = &SnapshotInfo{}
			}
			if err := m.SnapshotInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Member) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdExecPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EtcdExecPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailpointHTTPAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FailpointHTTPAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 101:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 102:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdLogPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EtcdLogPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 201:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdClientProxy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EtcdClientProxy = bool(v != 0)
		case 202:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdPeerProxy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EtcdPeerProxy = bool(v != 0)
		case 301:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdClientEndpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EtcdClientEndpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 302:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Etcd", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Etcd == nil {
				m.Etcd = &Etcd{}
			}
			if err := m.Etcd.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 303:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdOnSnapshotRestore", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EtcdOnSnapshotRestore == nil {
				m.EtcdOnSnapshotRestore = &Etcd{}
			}
			if err := m.EtcdOnSnapshotRestore.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 401:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientCertData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientCertData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 402:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientCertPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientCertPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 403:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientKeyData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientKeyData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 404:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientKeyPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientKeyPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 405:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientTrustedCAData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientTrustedCAData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 406:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientTrustedCAPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientTrustedCAPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 501:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerCertData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerCertData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 502:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerCertPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerCertPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 503:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerKeyData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerKeyData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 504:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerKeyPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerKeyPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 505:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerTrustedCAData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerTrustedCAData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 506:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerTrustedCAPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerTrustedCAPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 601:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 602:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SnapshotInfo == nil {
				m.SnapshotInfo = &SnapshotInfo{}
			}
			if err := m.SnapshotInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Tester) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Tester: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tester: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayLatencyMs", wireType)
			}
			m.DelayLatencyMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelayLatencyMs |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayLatencyMsRv", wireType)
			}
			m.DelayLatencyMsRv = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelayLatencyMsRv |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedDelayLatencyMs", wireType)
			}
			m.UpdatedDelayLatencyMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedDelayLatencyMs |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoundLimit", wireType)
			}
			m.RoundLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoundLimit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitOnCaseFail", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ExitOnCaseFail = bool(v != 0)
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnablePprof", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnablePprof = bool(v != 0)
		case 31:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaseDelayMs", wireType)
			}
			m.CaseDelayMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CaseDelayMs |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 32:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CaseShuffle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CaseShuffle = bool(v != 0)
		case 33:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cases", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cases = append(m.Cases, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 34:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailpointCommands", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FailpointCommands = append(m.FailpointCommands, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 41:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunnerExecPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunnerExecPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 42:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalExecPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalExecPath = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 101:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stressers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stressers = append(m.Stressers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 102:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Checkers = append(m.Checkers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 201:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StressKeySize", wireType)
			}
			m.StressKeySize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StressKeySize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 202:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StressKeySizeLarge", wireType)
			}
			m.StressKeySizeLarge = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StressKeySizeLarge |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 203:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StressKeySuffixRange", wireType)
			}
			m.StressKeySuffixRange = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StressKeySuffixRange |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 204:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StressKeySuffixRangeTxn", wireType)
			}
			m.StressKeySuffixRangeTxn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StressKeySuffixRangeTxn |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 205:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StressKeyTxnOps", wireType)
			}
			m.StressKeyTxnOps = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StressKeyTxnOps |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 301:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StressClients", wireType)
			}
			m.StressClients = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StressClients |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 302:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StressQPS", wireType)
			}
			m.StressQPS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StressQPS |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Etcd) Unmarshal(dAtA []byte) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Etcd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Etcd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WALDir", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WALDir = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeartbeatIntervalMs", wireType)
			}
			m.HeartbeatIntervalMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HeartbeatIntervalMs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ElectionTimeoutMs", wireType)
			}
			m.ElectionTimeoutMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ElectionTimeoutMs |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenClientURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenClientURLs = append(m.ListenClientURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdvertiseClientURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdvertiseClientURLs = append(m.AdvertiseClientURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientAutoTLS", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ClientAutoTLS = bool(v != 0)
		case 24:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientCertAuth", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ClientCertAuth = bool(v != 0)
		case 25:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientCertFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientCertFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 26:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientKeyFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientKeyFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 27:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientTrustedCAFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientTrustedCAFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 31:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenPeerURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenPeerURLs = append(m.ListenPeerURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 32:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdvertisePeerURLs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdvertisePeerURLs = append(m.AdvertisePeerURLs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 33:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerAutoTLS", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PeerAutoTLS = bool(v != 0)
		case 34:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerClientCertAuth", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PeerClientCertAuth = bool(v != 0)
		case 35:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerCertFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerCertFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 36:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerKeyFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerKeyFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 37:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerTrustedCAFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerTrustedCAFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 41:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialCluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialCluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 42:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClusterState", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClusterState = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 43:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClusterToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClusterToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 51:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotCount", wireType)
			}
			m.SnapshotCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SnapshotCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 52:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuotaBackendBytes", wireType)
			}
			m.QuotaBackendBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuotaBackendBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 63:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreVote", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PreVote = bool(v != 0)
		case 64:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialCorruptCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InitialCorruptCheck = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRpc(dAtA []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRpc
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRpc(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc   = fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("rpcpb/rpc.proto", fileDescriptorRpc)
}

var fileDescriptorRpc = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x59, 0xdb, 0x73, 0xdb, 0xc6, 0xf5, 0x16, 0x44, 0x5d, 0x57, 0x37, 0x68, 0x65, 0xd9, 0xf0, 0x4d, 0x90, 0xe1, 0x38, 0x3f, 0x59, 0x09, 0xec, 0xfc, 0xec, 0x4c, 0x2e, 0x4e, 0x13, 0x07, 0xa4, 0x20, 0x8b, 0x15, 0x44, 0xd2, 0x4b, 0xc8, 0x76, 0x9e, 0x38, 0x10, 0xb9, 0x92, 0x30, 0xa6, 0x00, 0x06, 0x58, 0x2a, 0x52, 0xfe, 0x81, 0xbe, 0xf6, 0x3e, 0xed, 0x4c, 0x9f, 0xfa, 0xdc, 0xb4, 0xff, 0x86, 0x73, 0x6b, 0xd3, 0xf6, 0xa9, 0xed, 0x0c, 0xa7, 0x4d, 0x5f, 0xfa, 0xd4, 0x07, 0x4e, 0x6f, 0xe9, 0x53, 0x67, 0x77, 0x01, 0x71, 0x01, 0x90, 0x92, 0x9e, 0xa4, 0x3d, 0xe7, 0xfb, 0xbe, 0x3d, 0xbb, 0x67, 0xb1, 0xe7, 0x00, 0x04, 0x73, 0x41, 0xab, 0xde, 0xda, 0xb9, 0x1b, 0xb4, 0xea, 0x77, 0x5a, 0x81, 0x4f, 0x7c, 0x38, 0xca, 0x0c, 0x57, 0xf4, 0x3d, 0x97, 0xec, 0xb7, 0x77, 0xee, 0xd4, 0xfd, 0x83, 0xbb, 0x7b, 0xfe, 0x9e, 0x7f, 0x97, 0x79, 0x77, 0xda, 0xbb, 0x6c, 0xc4, 0x06, 0xec, 0x3f, 0xce, 0xd2, 0xbe, 0x23, 0x81, 0x71, 0x84, 0x3f, 0x6c, 0xe3, 0x90, 0xc0, 0x3b, 0x60, 0xb2, 0xdc, 0xc2, 0x81, 0x43, 0x5c, 0xdf, 0x53, 0xa4, 0x65, 0x69, 0x65, 0xf6, 0x9e, 0x7c, 0x87, 0xa9, 0xde, 0x39, 0xb1, 0xa3, 0x1e, 0x04, 0xde, 0x02, 0x63, 0x5b, 0xf8, 0x60, 0x07, 0x07, 0xca, 0xf0, 0xb2, 0xb4, 0x32, 0x75, 0x6f, 0x26, 0x02, 0x73, 0x23, 0x8a, 0x9c, 0x14, 0x66, 0xe3, 0x90, 0xe0, 0x40, 0xc9, 0x25, 0x60, 0xdc, 0x88, 0x22, 0xa7, 0xf6, 0xb7, 0x61, 0x30, 0x5d, 0xf5, 0x9c, 0x56, 0xb8, 0xef, 0x93, 0xa2, 0xb7, 0xeb, 0xc3, 0x25, 0x00, 0xb8, 0x42, 0xc9, 0x39, 0xc0, 0x2c, 0x9e, 0x49, 0x24, 0x58, 0xe0, 0x2a, 0x90, 0xf9, 0xa8, 0xd0, 0x74, 0xb1, 0x47, 0xb6, 0x91, 0x15, 0x2a, 0xc3, 0xcb, 0xb9, 0x95, 0x49, 0x94, 0xb1, 0x43, 0xad, 0xa7, 0x5d, 0x71, 0xc8, 0x3e, 0x8b, 0x64, 0x12, 0x25, 0x6c, 0x54, 0x2f, 0x1e, 0xaf, 0xbb, 0x4d, 0x5c, 0x75, 0x3f, 0xc6, 0xca, 0x08, 0xc3, 0x65, 0xec, 0xf0, 0x55, 0x30, 0x1f, 0xdb, 0x6c, 0x9f, 0x38, 0x4d, 0x06, 0x1e, 0x65, 0xe0, 0xac, 0x43, 0x54, 0x66, 0xc6, 0x4d, 0x7c, 0xac, 0x8c, 0x2d, 0x4b, 0x2b, 0x39, 0x94, 0xb1, 0x8b, 0x91, 0x6e, 0x38, 0xe1, 0xbe, 0x32, 0xce, 0x70, 0x09, 0x9b, 0xa8, 0x87, 0xf0, 0xa1, 0x1b, 0xd2, 0x7c, 0x4d, 0x24, 0xf5, 0x62, 0x3b, 0x84, 0x60, 0xc4, 0xf6, 0xfd, 0xe7, 0xca, 0x24, 0x0b, 0x8e, 0xfd, 0xaf, 0xfd, 0x4c, 0x02, 0x13, 0x08, 0x87, 0x2d, 0xdf, 0x0b, 0x31, 0x54, 0xc0, 0x78, 0xb5, 0x5d, 0xaf, 0xe3, 0x30, 0x64, 0x7b, 0x3c, 0x81, 0xe2, 0x21, 0xbc, 0x08, 0xc6, 0xaa, 0xc4, 0x21, 0xed, 0x90, 0xe5, 0x77, 0x12, 0x45, 0x23, 0x21, 0xef, 0xb9, 0xd3, 0xf2, 0xfe, 0x66, 0x32, 0x9f, 0x6c, 0x2f, 0xa7, 0xee, 0x2d, 0x44, 0x60, 0xd1, 0x85, 0x12, 0x40, 0xed, 0x4f, 0xd3, 0xf1, 0x04, 0xf0, 0x5d, 0x30, 0x6d, 0x92, 0x7a, 0xc3, 0x3c, 0xc2, 0x75, 0x96, 0x37, 0x76, 0x0a, 0xf2, 0x97, 0xbb, 0x1d, 0x75, 0xf1, 0xd8, 0x39, 0x68, 0x3e, 0xd0, 0x30, 0xa9, 0x37, 0x74, 0x7c, 0x84, 0xeb, 0x7a, 0xcb, 0x21, 0xfb, 0x1a, 0x4a, 0xc0, 0xe1, 0x7d, 0x30, 0x69, 0xec, 0x61, 0x8f, 0x18, 0x8d, 0x46, 0xa0, 0x4c, 0x31, 0xee, 0x62, 0xb7, 0xa3, 0xce, 0x73, 0xae, 0x43, 0x5d, 0xba, 0xd3, 0x68, 0x04, 0x1a, 0xea, 0xe1, 0xa0, 0x05, 0xe6, 0xd7, 0x1d, 0xb7, 0xd9, 0xf2, 0x5d, 0x8f, 0x6c, 0xd8, 0x76, 0x85, 0x91, 0xa7, 0x19, 0x79, 0xa9, 0xdb, 0x51, 0xaf, 0x70, 0xf2, 0x6e, 0x0c, 0xd1, 0xf7, 0x09, 0x69, 0x45, 0x2a, 0x59, 0x22, 0xd4, 0xc1, 0x78, 0xde, 0x09, 0xf1, 0x9a, 0x1b, 0x28, 0x98, 0x69, 0x2c, 0x74, 0x3b, 0xea, 0x1c, 0xd7, 0xd8, 0x71, 0x42, 0xac, 0x37, 0xdc, 0x40, 0x43, 0x31, 0x06, 0x3e, 0x00, 0x53, 0x74, 0x05, 0x96, 0xbf, 0xc7, 0xd6, 0xbb, 0xcb, 0x28, 0x4a, 0xb7, 0xa3, 0x5e, 0x10, 0xd6, 0xdb, 0xf4, 0xf7, 0xa2, 0xe5, 0x8a, 0x60, 0xf8, 0x08, 0xcc, 0xd1, 0x21, 0x3f, 0xf6, 0x95, 0xc0, 0x3f, 0x3a, 0x56, 0x3e, 0x65, 0x29, 0xcd, 0x5f, 0xeb, 0x76, 0x54, 0x45, 0x10, 0xa8, 0x33, 0x88, 0xde, 0xa2, 0x18, 0x0d, 0xa5, 0x59, 0xd0, 0x00, 0x33, 0xd4, 0x54, 0xc1, 0x38, 0xe0, 0x32, 0x9f, 0x71, 0x99, 0x2b, 0xdd, 0x8e, 0x7a, 0x51, 0x90, 0x69, 0x61, 0x1c, 0xc4, 0x22, 0x49, 0x06, 0xac, 0x00, 0xd8, 0x53, 0x35, 0xbd, 0x06, 0xdb, 0x14, 0xe5, 0x13, 0x76, 0x90, 0xf2, 0x6a, 0xb7, 0xa3, 0x5e, 0xcd, 0x86, 0x83, 0x23, 0x98, 0x86, 0xfa, 0x70, 0xe1, 0xff, 0x83, 0x11, 0x6a, 0x55, 0x7e, 0xc9, 0x2f, 0x9b, 0xa9, 0xe8, 0x1c, 0x51, 0x5b, 0x7e, 0xae, 0xdb, 0x51, 0xa7, 0x7a, 0x82, 0x1a, 0x62, 0x50, 0x98, 0x07, 0x8b, 0xf4, 0x6f, 0xd9, 0xeb, 0x3d, 0x15, 0x21, 0xf1, 0x03, 0xac, 0xfc, 0x2a, 0xab, 0x81, 0xfa, 0x43, 0xe1, 0x1a, 0x98, 0xe5, 0x81, 0x14, 0x70, 0x40, 0xd6, 0x1c, 0xe2, 0x28, 0xdf, 0x63, 0x97, 0x47, 0xfe, 0x6a, 0xb7, 0xa3, 0x5e, 0xe2, 0x73, 0x46, 0xf1, 0xd7, 0x71, 0x40, 0xf4, 0x86, 0x43, 0x1c, 0x0d, 0xa5, 0x38, 0x49, 0x15, 0x96, 0xd9, 0xef, 0x9f, 0xaa, 0xc2, 0xb3, 0x9b, 0xe2, 0xd0, 0xbc, 0x70, 0xcb, 0x26, 0x3e, 0x66, 0xa1, 0xfc, 0x80, 0x8b, 0x08, 0x79, 0x89, 0x44, 0x9e, 0xe3, 0xe3, 0x28, 0x92, 0x24, 0x23, 0x21, 0xc1, 0xe2, 0xf8, 0xe1, 0x69, 0x12, 0x3c, 0x8c, 0x24, 0x03, 0xda, 0x60, 0x81, 0x1b, 0xec, 0xa0, 0x1d, 0x12, 0xdc, 0x28, 0x18, 0x2c, 0x96, 0x1f, 0x71, 0xa1, 0x1b, 0xdd, 0x8e, 0x7a, 0x3d, 0x21, 0x44, 0x38, 0x4c, 0xaf, 0x3b, 0x51, 0x48, 0xfd, 0xe8, 0x7d, 0x54, 0x59, 0x78, 0x3f, 0x3e, 0x87, 0x2a, 0x8f, 0xb2, 0x1f, 0x1d, 0xbe, 0x07, 0xa6, 0xe9, 0x99, 0x3c, 0xc9, 0xdd, 0x3f, 0x73, 0xe9, 0x0b, 0x84, 0x9d, 0x61, 0x21, 0x73, 0x09, 0xbc, 0xc8, 0x67, 0xe1, 0xfc, 0xeb, 0x14, 0x7e, 0x74, 0x01, 0x89, 0x78, 0xf8, 0x0e, 0x98, 0xa2, 0xe3, 0x38, 0x5f, 0xff, 0xce, 0xa5, 0x9f, 0x67, 0x46, 0xef, 0x65, 0x4b, 0x44, 0x0b, 0x64, 0x36, 0xf7, 0x7f, 0x06, 0x93, 0xa3, 0xcb, 0x40, 0x40, 0xc3, 0x12, 0x98, 0xa7, 0xc3, 0x64, 0x8e, 0xbe, 0xc9, 0xa5, 0x9f, 0x3f, 0x26, 0x91, 0xc9, 0x50, 0x96, 0x9a, 0xd1, 0x63, 0x21, 0xfd, 0xf7, 0x4c, 0x3d, 0x1e, 0x59, 0x96, 0x4a, 0x6f, 0xf6, 0x44, 0x45, 0xfe, 0xc3, 0x48, 0x7a, 0x75, 0x61, 0xe4, 0x8e, 0x37, 0x36, 0x51, 0xac, 0xdf, 0x4a, 0x15, 0x97, 0x3f, 0x9e, 0xbb, 0xba, 0xfc, 0x7c, 0x3a, 0xee, 0x47, 0xe8, 0xdd, 0x4c, 0xd7, 0x46, 0xef, 0x66, 0x29, 0x7d, 0x37, 0xd3, 0x8d, 0x88, 0xee, 0xe6, 0x08, 0x03, 0x5f, 0x05, 0xe3, 0x25, 0x4c, 0x3e, 0xf2, 0x83, 0xe7, 0xbc, 0x20, 0xe6, 0x61, 0xb7, 0xa3, 0xce, 0x72, 0xb8, 0xc7, 0x1d, 0x1a, 0x8a, 0x21, 0xf0, 0x26, 0x18, 0x61, 0x95, 0x83, 0x6f, 0x91, 0x70, 0x43, 0xf1, 0x52, 0xc1, 0x9c, 0xb0, 0x00, 0x66, 0xd7, 0x70, 0xd3, 0x39, 0xb6, 0x1c, 0x82, 0xbd, 0xfa, 0xf1, 0x56, 0xc8, 0xaa, 0xd4, 0x8c, 0x78, 0x2d, 0x34, 0xa8, 0x5f, 0x6f, 0x72, 0x80, 0x7e, 0x10, 0x6a, 0x28, 0x45, 0x81, 0xdf, 0x06, 0x72, 0xd2, 0x82, 0x0e, 0x59, 0xbd, 0x9a, 0x11, 0xeb, 0x55, 0x5a, 0x46, 0x0f, 0x0e, 0x35, 0x94, 0xe1, 0xc1, 0x0f, 0xc0, 0xe2, 0x76, 0xab, 0xe1, 0x10, 0xdc, 0x48, 0xc5, 0x35, 0xc3, 0x04, 0x6f, 0x76, 0x3b, 0xaa, 0xca, 0x05, 0xdb, 0x1c, 0xa6, 0x67, 0xe3, 0xeb, 0xaf, 0x00, 0xdf, 0x00, 0x00, 0xf9, 0x6d, 0xaf, 0x61, 0xb9, 0x07, 0x2e, 0x51, 0x16, 0x97, 0xa5, 0x95, 0xd1, 0xfc, 0xc5, 0x6e, 0x47, 0x85, 0x5c, 0x2f, 0xa0, 0x3e, 0xbd, 0x49, 0x9d, 0x1a, 0x12, 0x90, 0x30, 0x0f, 0x66, 0xcd, 0x23, 0x97, 0x94, 0xbd, 0x82, 0x13, 0x62, 0x5a, 0x60, 0x95, 0x8b, 0x99, 0x6a, 0x74, 0xe4, 0x12, 0xdd, 0xf7, 0x74, 0x5a, 0x94, 0xdb, 0x01, 0xd6, 0x50, 0x8a, 0x01, 0xdf, 0x06, 0x53, 0xa6, 0xe7, 0xec, 0x34, 0x71, 0xa5, 0x15, 0xf8, 0xbb, 0xca, 0x25, 0x26, 0x70, 0xa9, 0xdb, 0x51, 0x17, 0x22, 0x01, 0xe6, 0xd4, 0x5b, 0xd4, 0x4b, 0xab, 0x6a, 0x0f, 0x4b, 0x2b, 0x32, 0x95, 0x61, 0x8b, 0xd9, 0x0a, 0x15, 0x95, 0xed, 0x83, 0x70, 0x4c, 0xeb, 0xac, 0x88, 0xb3, 0x4d, 0xa0, 0x8b, 0x17, 0xc1, 0x74, 0x5a, 0x3a, 0xac, 0xee, 0xb7, 0x77, 0x77, 0x9b, 0x58, 0x59, 0x4e, 0x4f, 0xcb, 0xb8, 0x21, 0xf7, 0x46, 0xd4, 0x08, 0x0b, 0x5f, 0x06, 0xa3, 0x74, 0x18, 0x2a, 0x37, 0x68, 0x4b, 0x9b, 0x97, 0xbb, 0x1d, 0x75, 0xba, 0x47, 0x0a, 0x35, 0xc4, 0xdd, 0x70, 0x53, 0xe8, 0x56, 0x0a, 0xfe, 0xc1, 0x81, 0xe3, 0x35, 0x42, 0x45, 0x63, 0x9c, 0xeb, 0xdd, 0x8e, 0x7a, 0x39, 0xdd, 0xad, 0xd4, 0x23, 0x8c, 0xd8, 0xac, 0xc4, 0x3c, 0x7a, 0x1c, 0x51, 0xdb, 0xf3, 0x70, 0x70, 0xd2, 0x70, 0xdd, 0x4e, 0x57, 0xa9, 0x80, 0xf9, 0xc5, 0x96, 0x2b, 0x45, 0x81, 0x45, 0x20, 0x9b, 0x47, 0x04, 0x07, 0x9e, 0xd3, 0x3c, 0x91, 0x59, 0x65, 0x32, 0x42, 0x40, 0x38, 0x42, 0x88, 0x42, 0x19, 0x1a, 0xbc, 0x07, 0x26, 0xab, 0x24, 0xc0, 0x61, 0x88, 0x83, 0x50, 0xc1, 0x6c, 0x51, 0x17, 0xba, 0x1d, 0x55, 0x8e, 0x2e, 0x88, 0xd8, 0xa5, 0xa1, 0x1e, 0x0c, 0xde, 0x05, 0x13, 0x85, 0x7d, 0x5c, 0x7f, 0x4e, 0x29, 0xbb, 0x8c, 0x22, 0x3c, 0xd5, 0xf5, 0xc8, 0xa3, 0xa1, 0x13, 0x10, 0x2d, 0x89, 0x9c, 0xbd, 0x89, 0x8f, 0x59, 0x1f, 0xcf, 0x9a, 0xa6, 0x51, 0xf1, 0x7c, 0xf1, 0x99, 0xd8, 0x55, 0x1b, 0xba, 0x1f, 0x63, 0x0d, 0x25, 0x19, 0xf0, 0x31, 0x80, 0x09, 0x83, 0xe5, 0x04, 0x7b, 0x98, 0x77, 0x4d, 0xa3, 0xf9, 0xe5, 0x6e, 0x47, 0xbd, 0xd6, 0x57, 0x47, 0x6f, 0x52, 0x9c, 0x86, 0xfa, 0x90, 0xe1, 0x53, 0x70, 0xa1, 0x67, 0x6d, 0xef, 0xee, 0xba, 0x47, 0xc8, 0xf1, 0xf6, 0xb0, 0xf2, 0x39, 0x17, 0xd5, 0xba, 0x1d, 0x75, 0x29, 0x2b, 0xca, 0x80, 0x7a, 0x40, 0x91, 0x1a, 0xea, 0x2b, 0x00, 0x1d, 0x70, 0xa9, 0x9f, 0xdd, 0x3e, 0xf2, 0x94, 0x2f, 0xb8, 0xf6, 0xcb, 0xdd, 0x8e, 0xaa, 0x9d, 0xaa, 0xad, 0x93, 0x23, 0x4f, 0x43, 0x83, 0x74, 0xe0, 0x06, 0x98, 0x3b, 0x71, 0xd9, 0x47, 0x5e, 0xb9, 0x15, 0x2a, 0x5f, 0x72, 0x69, 0xe1, 0x04, 0x08, 0xd2, 0xe4, 0xc8, 0xd3, 0xfd, 0x56, 0xa8, 0xa1, 0x34, 0x0d, 0xbe, 0x1f, 0xe7, 0x86, 0x17, 0xf7, 0x90, 0x77, 0x90, 0xa3, 0x62, 0x01, 0x8e, 0x74, 0x78, 0x5b, 0x10, 0x9e, 0xa4, 0x26, 0x22, 0xc0, 0xd7, 0xe3, 0x23, 0xf4, 0xb8, 0x52, 0xe5, 0xbd, 0xe3, 0xa8, 0xf8, 0x0e, 0x10, 0xb1, 0x3f, 0x6c, 0xf5, 0x0e, 0xd1, 0xe3, 0x4a, 0x55, 0xfb, 0x66, 0x86, 0x77, 0x9b, 0xf4, 0x16, 0xef, 0xbd, 0x7e, 0x8a, 0xb7, 0xb8, 0xe7, 0x1c, 0x60, 0x0d, 0x31, 0xa7, 0x58, 0x47, 0x86, 0xcf, 0x51, 0x47, 0x56, 0xc1, 0xd8, 0x53, 0xc3, 0xa2, 0xe8, 0x5c, 0xba, 0x8c, 0x7c, 0xe4, 0x34, 0x39, 0x38, 0x42, 0xc0, 0x32, 0x58, 0xd8, 0xc0, 0x4e, 0x40, 0x76, 0xb0, 0x43, 0x8a, 0x1e, 0xc1, 0xc1, 0xa1, 0xd3, 0x8c, 0xaa, 0x44, 0x4e, 0xdc, 0xcd, 0xfd, 0x18, 0xa4, 0xbb, 0x11, 0x4a, 0x43, 0xfd, 0x98, 0xb0, 0x08, 0xe6, 0xcd, 0x26, 0xae, 0xd3, 0x17, 0x78, 0xdb, 0x3d, 0xc0, 0x7e, 0x9b, 0x6c, 0x85, 0xac, 0x5a, 0xe4, 0xc4, 0xa7, 0x1c, 0x47, 0x10, 0x9d, 0x70, 0x8c, 0x86, 0xb2, 0x2c, 0xfa, 0xa0, 0x5b, 0x6e, 0x48, 0xb0, 0x27, 0xbc, 0x80, 0x2f, 0xa6, 0x6f, 0x9e, 0x26, 0x43, 0xc4, 0x2d, 0x7e, 0x3b, 0x68, 0x86, 0x1a, 0xca, 0xd0, 0x20, 0x02, 0x0b, 0x46, 0xe3, 0x10, 0x07, 0xc4, 0x0d, 0xb1, 0xa0, 0x76, 0x91, 0xa9, 0x09, 0x0f, 0x90, 0x13, 0x83, 0x92, 0x82, 0xfd, 0xc8, 0xf0, 0xed, 0xb8, 0xd5, 0x35, 0xda, 0xc4, 0xb7, 0xad, 0x6a, 0x74, 0xeb, 0x0b, 0xb9, 0x71, 0xda, 0xc4, 0xd7, 0x09, 0x15, 0x48, 0x22, 0xe9, 0x3d, 0xd8, 0x6b, 0xbd, 0x8d, 0x36, 0xd9, 0x57, 0x14, 0xc6, 0x1d, 0xd0, 0xad, 0x3b, 0xed, 0x54, 0xb7, 0x4e, 0x29, 0xf0, 0x5b, 0xa2, 0xc8, 0xba, 0xdb, 0xc4, 0xca, 0x65, 0x96, 0x6e, 0xe1, 0x06, 0x63, 0xec, 0x5d, 0x97, 0x5e, 0xfe, 0x29, 0x6c, 0x2f, 0xfa, 0x4d, 0x7c, 0xcc, 0xc8, 0x57, 0xd2, 0x27, 0x8b, 0x3e, 0x39, 0x9c, 0x9b, 0x44, 0x42, 0x2b, 0xd3, 0x4a, 0x33, 0x81, 0xab, 0xe9, 0x46, 0x5f, 0x68, 0xd3, 0xb8, 0x4e, 0x3f, 0x1a, 0xdd, 0x0b, 0x9e, 0x2e, 0xda, 0xc3, 0xb1, 0xac, 0xa8, 0x2c, 0x2b, 0xc2, 0x5e, 0x44, 0x39, 0x66, 0xbd, 0x1f, 0x4f, 0x48, 0x8a, 0x02, 0x6d, 0x30, 0x7f, 0x92, 0xa2, 0x13, 0x9d, 0x65, 0xa6, 0x23, 0xdc, 0x36, 0xae, 0xe7, 0x12, 0xd7, 0x69, 0xea, 0xbd, 0x2c, 0x0b, 0x92, 0x59, 0x01, 0x5a, 0x9a, 0xe9, 0xff, 0x71, 0x7e, 0x6f, 0xb0, 0x1c, 0xa5, 0xfb, 0xe3, 0x5e, 0x92, 0x45, 0x30, 0x7d, 0x41, 0x65, 0x9d, 0x7a, 0x32, 0xcd, 0x1a, 0x93, 0x10, 0x0e, 0x1c, 0x6f, 0xef, 0x33, 0xb9, 0xee, 0xc3, 0xa5, 0x1d, 0x6d, 0xdc, 0xfb, 0xb3, 0xfd, 0xbe, 0x39, 0xf8, 0x55, 0x81, 0x6f, 0x77, 0x02, 0x1e, 0x2f, 0x26, 0x4e, 0xf7, 0x4b, 0x03, 0x9b, 0x7d, 0x4e, 0x16, 0xc1, 0x70, 0x2b, 0xd5, 0x9c, 0x33, 0x85, 0x5b, 0x67, 0xf5, 0xe6, 0x5c, 0x28, 0xcb, 0xa4, 0x1d, 0x57, 0x91, 0xa7, 0xa2, 0xd0, 0x6c, 0xb3, 0x2f, 0x77, 0xb7, 0xd3, 0x67, 0x27, 0x4e, 0x55, 0x9d, 0x03, 0x34, 0x94, 0x62, 0xd0, 0x27, 0x3a, 0x69, 0xa9, 0x12, 0x87, 0xe0, 0xa8, 0x11, 0x10, 0x36, 0x38, 0x25, 0xa4, 0x87, 0x14, 0xa6, 0xa1, 0x7e, 0xe4, 0xac, 0xa6, 0xed, 0x3f, 0xc7, 0x9e, 0xf2, 0xca, 0x59, 0x9a, 0x84, 0xc2, 0x32, 0x9a, 0x8c, 0x0c, 0x1f, 0x82, 0x99, 0xf8, 0xf5, 0xa0, 0xe0, 0xb7, 0x3d, 0xa2, 0xdc, 0x67, 0x77, 0xa1, 0x58, 0x60, 0xe2, 0xf7, 0x90, 0x3a, 0xf5, 0xd3, 0x02, 0x23, 0xe2, 0xa1, 0x05, 0xe6, 0x1f, 0xb7, 0x7d, 0xe2, 0xe4, 0x9d, 0xfa, 0x73, 0xec, 0x35, 0xf2, 0xc7, 0x04, 0x87, 0xca, 0xeb, 0x4c, 0x44, 0x68, 0xbf, 0x3f, 0xa4, 0x10, 0x7d, 0x87, 0x63, 0xf4, 0x1d, 0x0a, 0xd2, 0x50, 0x96, 0x48, 0x4b, 0x49, 0x25, 0xc0, 0x4f, 0x7c, 0x82, 0x95, 0x87, 0xe9, 0xeb, 0xaa, 0x15, 0x60, 0xfd, 0xd0, 0xa7, 0xbb, 0x13, 0x63, 0xc4, 0x1d, 0xf1, 0x83, 0xa0, 0xdd, 0x22, 0xac, 0xab, 0x51, 0xde, 0x4f, 0x1f, 0xe3, 0x93, 0x1d, 0xe1, 0x28, 0x9d, 0xf5, 0x41, 0xc2, 0x8e, 0x08, 0xe4, 0xd5, 0x9f, 0xe6, 0x84, 0xef, 0xc0, 0x70, 0x0e, 0x4c, 0x95, 0xca, 0x76, 0xad, 0x6a, 0x1b, 0xc8, 0x36, 0xd7, 0xe4, 0x21, 0x78, 0x11, 0xc0, 0x62, 0xa9, 0x68, 0x17, 0x0d, 0x8b, 0x1b, 0x6b, 0xa6, 0x5d, 0x58, 0x93, 0x01, 0x94, 0xc1, 0x34, 0x32, 0x05, 0xcb, 0x14, 0xb5, 0x54, 0x8b, 0x8f, 0x6c, 0x13, 0x6d, 0x71, 0xcb, 0x05, 0xb8, 0x0c, 0xae, 0x55, 0x8b, 0x8f, 0x1e, 0x6f, 0x17, 0x39, 0xa6, 0x66, 0x94, 0xd6, 0x6a, 0xc8, 0xdc, 0x2a, 0x3f, 0x31, 0x6b, 0x6b, 0x86, 0x6d, 0xc8, 0x8b, 0x70, 0x1e, 0xcc, 0x54, 0x8d, 0x27, 0x66, 0xad, 0x5a, 0x32, 0x2a, 0xd5, 0x8d, 0xb2, 0x2d, 0x2f, 0xc1, 0x1b, 0xe0, 0x3a, 0x15, 0x2e, 0x23, 0xb3, 0x16, 0x4f, 0xb0, 0x8e, 0xca, 0x5b, 0x3d, 0x88, 0x0a, 0x2f, 0x83, 0xc5, 0xfe, 0xae, 0x65, 0xca, 0xce, 0x4c, 0x69, 0xa0, 0xc2, 0x46, 0x31, 0x9e, 0x73, 0x05, 0xde, 0x05, 0xaf, 0x9c, 0x16, 0x15, 0x1b, 0x57, 0xed, 0x72, 0xa5, 0x66, 0x3c, 0x32, 0x4b, 0xb6, 0x7c, 0x1b, 0x5e, 0x07, 0x97, 0xf3, 0x96, 0x51, 0xd8, 0xdc, 0x28, 0x5b, 0x66, 0xad, 0x62, 0x9a, 0xa8, 0x56, 0x29, 0x23, 0xbb, 0x66, 0x3f, 0xab, 0xa1, 0x67, 0x72, 0x03, 0xaa, 0xe0, 0xea, 0x76, 0x69, 0x30, 0x00, 0xc3, 0x2b, 0x60, 0x71, 0xcd, 0xb4, 0x8c, 0x0f, 0x32, 0xae, 0x17, 0x12, 0xbc, 0x06, 0x2e, 0x6d, 0x97, 0xfa, 0x7b, 0x3f, 0x95, 0x56, 0xff, 0x0e, 0xc0, 0x08, 0xed, 0xfb, 0xa1, 0x02, 0x2e, 0xc4, 0x7b, 0x5b, 0x2e, 0x99, 0xb5, 0xf5, 0xb2, 0x65, 0x95, 0x9f, 0x9a, 0x48, 0x1e, 0x8a, 0x56, 0x93, 0xf1, 0xd4, 0xb6, 0x4b, 0x76, 0xd1, 0xaa, 0xd9, 0xa8, 0xf8, 0xe8, 0x91, 0x89, 0x7a, 0x3b, 0x24, 0x41, 0x08, 0x66, 0x63, 0x82, 0x65, 0x1a, 0x6b, 0x26, 0x92, 0x87, 0xe1, 0x6d, 0x70, 0x2b, 0x69, 0x1b, 0x44, 0xcf, 0x89, 0xf4, 0xc7, 0xdb, 0x65, 0xb4, 0xbd, 0x25, 0x8f, 0xd0, 0x43, 0x13, 0xdb, 0x0c, 0xcb, 0x92, 0x47, 0xe1, 0x4d, 0xa0, 0xc6, 0x5b, 0x2c, 0xec, 0x6e, 0x22, 0x72, 0x00, 0x1f, 0x80, 0x37, 0xce, 0x00, 0x0d, 0x8a, 0x62, 0x8a, 0xa6, 0xa4, 0x0f, 0x37, 0x5a, 0xcf, 0x34, 0x7c, 0x1d, 0xbc, 0x36, 0xd0, 0x3d, 0x48, 0x74, 0x06, 0xae, 0x83, 0x7c, 0x1f, 0x16, 0x5f, 0x65, 0x64, 0xe1, 0xe7, 0x32, 0x12, 0x8a, 0xa9, 0xd1, 0x21, 0x2c, 0x20, 0xc3, 0x2e, 0x6c, 0xc8, 0xb3, 0x70, 0x15, 0xbc, 0x3c, 0xf0, 0x38, 0x24, 0x37, 0xa1, 0x01, 0x0d, 0xf0, 0xee, 0xf9, 0xb0, 0x83, 0xc2, 0xc6, 0xf0, 0x25, 0xb0, 0x3c, 0x58, 0x22, 0xda, 0x92, 0x5d, 0xf8, 0x0e, 0x78, 0xf3, 0x2c, 0xd4, 0xa0, 0x29, 0xf6, 0x4e, 0x9f, 0x22, 0x3a, 0x06, 0xfb, 0xf4, 0xd9, 0x1b, 0x8c, 0xa2, 0x07, 0xc3, 0x85, 0xff, 0x07, 0xb4, 0xbe, 0x87, 0x3d, 0xb9, 0x2d, 0x2f, 0x24, 0x78, 0x07, 0xdc, 0x46, 0x46, 0x69, 0xad, 0xbc, 0x55, 0x3b, 0x07, 0xfe, 0x53, 0x09, 0xbe, 0x07, 0xde, 0x3e, 0x1b, 0x38, 0x68, 0x81, 0x9f, 0x49, 0xd0, 0x04, 0xef, 0x9f, 0x7b, 0xbe, 0x41, 0x32, 0x9f, 0x4b, 0xf0, 0x06, 0xb8, 0xd6, 0x9f, 0x1f, 0xe5, 0xe1, 0x0b, 0x09, 0xae, 0x80, 0x9b, 0xa7, 0xce, 0x14, 0x21, 0xbf, 0x94, 0xe0, 0x5b, 0xe0, 0xfe, 0x69, 0x90, 0x41, 0x61, 0xfc, 0x5a, 0x82, 0x0f, 0xc1, 0x83, 0x73, 0xcc, 0x31, 0x48, 0xe0, 0x37, 0xa7, 0xac, 0x23, 0x4a, 0xf6, 0x57, 0x67, 0xaf, 0x23, 0x42, 0xfe, 0x56, 0x82, 0x4b, 0xe0, 0x72, 0x7f, 0x08, 0x3d, 0x13, 0xbf, 0x93, 0xe0, 0x2d, 0xb0, 0x7c, 0xaa, 0x12, 0x85, 0xfd, 0x5e, 0x82, 0x0a, 0x58, 0x28, 0x95, 0x6b, 0xeb, 0x46, 0xd1, 0xaa, 0x3d, 0x2d, 0xda, 0x1b, 0xb5, 0xaa, 0x8d, 0xcc, 0x6a, 0x55, 0xfe, 0xc5, 0x30, 0x0d, 0x25, 0xe1, 0x29, 0x95, 0x23, 0x67, 0x6d, 0xbd, 0x8c, 0x6a, 0x56, 0xf1, 0x89, 0x59, 0xa2, 0xc8, 0x4f, 0x86, 0xe1, 0x1c, 0x00, 0x14, 0x56, 0x29, 0x17, 0x4b, 0x76, 0x55, 0xfe, 0x6e, 0x0e, 0xce, 0x80, 0x09, 0xf3, 0x99, 0x6d, 0xa2, 0x92, 0x61, 0xc9, 0xff, 0xc8, 0xad, 0x1e, 0x80, 0x89, 0xf8, 0xd3, 0x02, 0x1c, 0x03, 0xc3, 0x9b, 0x4f, 0xe4, 0x21, 0x38, 0x09, 0x46, 0x2d, 0xd3, 0xa8, 0x9a, 0xb2, 0x04, 0x17, 0xc0, 0x9c, 0x69, 0x99, 0x05, 0xbb, 0x58, 0x2e, 0xd5, 0xd0, 0x76, 0xa9, 0xc4, 0x2e, 0x4f, 0x19, 0x4c, 0x3f, 0xa5, 0x4f, 0x7e, 0x6c, 0xc9, 0xc1, 0x45, 0x30, 0x6f, 0x95, 0x0b, 0x9b, 0x35, 0x64, 0x14, 0x4c, 0x14, 0x9b, 0x47, 0x28, 0x90, 0x09, 0xc5, 0x96, 0xd1, 0xd5, 0x3c, 0x18, 0x8f, 0xbe, 0x4b, 0xc0, 0x29, 0x30, 0xbe, 0xf9, 0xa4, 0xb6, 0x61, 0x54, 0x37, 0xe4, 0xa1, 0x1e, 0xd2, 0x7c, 0x56, 0x29, 0x22, 0x3a, 0x33, 0x00, 0x63, 0x27, 0x13, 0x4e, 0x83, 0x89, 0x52, 0xb9, 0x56, 0xd8, 0x30, 0x0b, 0x9b, 0x72, 0xee, 0xde, 0x43, 0x30, 0x69, 0x07, 0x8e, 0x17, 0xb6, 0xfc, 0x80, 0xc0, 0x7b, 0xe2, 0x60, 0x36, 0xfa, 0x3a, 0x1a, 0xfd, 0xe0, 0x7b, 0x65, 0xee, 0x64, 0xcc, 0x7f, 0x0b, 0xd4, 0x86, 0x56, 0xa4, 0xd7, 0xa4, 0xfc, 0x85, 0x17, 0x7f, 0x59, 0x1a, 0x7a, 0xf1, 0xf5, 0x92, 0xf4, 0xd5, 0xd7, 0x4b, 0xd2, 0x9f, 0xbf, 0x5e, 0x92, 0x7e, 0xf2, 0xd7, 0xa5, 0xa1, 0x9d, 0x31, 0xf6, 0x83, 0xf1, 0xfd, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x9f, 0x8c, 0x37, 0x79, 0x1e, 0x00, 0x00}
