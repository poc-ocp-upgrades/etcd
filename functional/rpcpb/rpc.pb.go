package rpcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"
import binary "encoding/binary"
import io "io"

var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

const _ = proto.ProtoPackageIsVersion2

type StresserType int32

const (
	StresserType_KV_WRITE_SMALL		StresserType	= 0
	StresserType_KV_WRITE_LARGE		StresserType	= 1
	StresserType_KV_READ_ONE_KEY		StresserType	= 2
	StresserType_KV_READ_RANGE		StresserType	= 3
	StresserType_KV_DELETE_ONE_KEY		StresserType	= 4
	StresserType_KV_DELETE_RANGE		StresserType	= 5
	StresserType_KV_TXN_WRITE_DELETE	StresserType	= 6
	StresserType_LEASE			StresserType	= 10
	StresserType_ELECTION_RUNNER		StresserType	= 20
	StresserType_WATCH_RUNNER		StresserType	= 31
	StresserType_LOCK_RACER_RUNNER		StresserType	= 41
	StresserType_LEASE_RUNNER		StresserType	= 51
)

var StresserType_name = map[int32]string{0: "KV_WRITE_SMALL", 1: "KV_WRITE_LARGE", 2: "KV_READ_ONE_KEY", 3: "KV_READ_RANGE", 4: "KV_DELETE_ONE_KEY", 5: "KV_DELETE_RANGE", 6: "KV_TXN_WRITE_DELETE", 10: "LEASE", 20: "ELECTION_RUNNER", 31: "WATCH_RUNNER", 41: "LOCK_RACER_RUNNER", 51: "LEASE_RUNNER"}
var StresserType_value = map[string]int32{"KV_WRITE_SMALL": 0, "KV_WRITE_LARGE": 1, "KV_READ_ONE_KEY": 2, "KV_READ_RANGE": 3, "KV_DELETE_ONE_KEY": 4, "KV_DELETE_RANGE": 5, "KV_TXN_WRITE_DELETE": 6, "LEASE": 10, "ELECTION_RUNNER": 20, "WATCH_RUNNER": 31, "LOCK_RACER_RUNNER": 41, "LEASE_RUNNER": 51}

func (x StresserType) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.EnumName(StresserType_name, int32(x))
}
func (StresserType) EnumDescriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{0}
}

type Checker int32

const (
	Checker_KV_HASH		Checker	= 0
	Checker_LEASE_EXPIRE	Checker	= 1
	Checker_RUNNER		Checker	= 2
	Checker_NO_CHECK	Checker	= 3
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
	return fileDescriptorRpc, []int{1}
}

type Operation int32

const (
	Operation_NOT_STARTED					Operation	= 0
	Operation_INITIAL_START_ETCD				Operation	= 10
	Operation_RESTART_ETCD					Operation	= 11
	Operation_SIGTERM_ETCD					Operation	= 20
	Operation_SIGQUIT_ETCD_AND_REMOVE_DATA			Operation	= 21
	Operation_SAVE_SNAPSHOT					Operation	= 30
	Operation_RESTORE_RESTART_FROM_SNAPSHOT			Operation	= 31
	Operation_RESTART_FROM_SNAPSHOT				Operation	= 32
	Operation_SIGQUIT_ETCD_AND_ARCHIVE_DATA			Operation	= 40
	Operation_SIGQUIT_ETCD_AND_REMOVE_DATA_AND_STOP_AGENT	Operation	= 41
	Operation_BLACKHOLE_PEER_PORT_TX_RX			Operation	= 100
	Operation_UNBLACKHOLE_PEER_PORT_TX_RX			Operation	= 101
	Operation_DELAY_PEER_PORT_TX_RX				Operation	= 200
	Operation_UNDELAY_PEER_PORT_TX_RX			Operation	= 201
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
	return fileDescriptorRpc, []int{2}
}

type Case int32

const (
	Case_SIGTERM_ONE_FOLLOWER						Case	= 0
	Case_SIGTERM_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT			Case	= 1
	Case_SIGTERM_LEADER							Case	= 2
	Case_SIGTERM_LEADER_UNTIL_TRIGGER_SNAPSHOT				Case	= 3
	Case_SIGTERM_QUORUM							Case	= 4
	Case_SIGTERM_ALL							Case	= 5
	Case_SIGQUIT_AND_REMOVE_ONE_FOLLOWER					Case	= 10
	Case_SIGQUIT_AND_REMOVE_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT		Case	= 11
	Case_SIGQUIT_AND_REMOVE_LEADER						Case	= 12
	Case_SIGQUIT_AND_REMOVE_LEADER_UNTIL_TRIGGER_SNAPSHOT			Case	= 13
	Case_SIGQUIT_AND_REMOVE_QUORUM_AND_RESTORE_LEADER_SNAPSHOT_FROM_SCRATCH	Case	= 14
	Case_BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER				Case	= 100
	Case_BLACKHOLE_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT	Case	= 101
	Case_BLACKHOLE_PEER_PORT_TX_RX_LEADER					Case	= 102
	Case_BLACKHOLE_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT		Case	= 103
	Case_BLACKHOLE_PEER_PORT_TX_RX_QUORUM					Case	= 104
	Case_BLACKHOLE_PEER_PORT_TX_RX_ALL					Case	= 105
	Case_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER					Case	= 200
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER				Case	= 201
	Case_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT		Case	= 202
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_ONE_FOLLOWER_UNTIL_TRIGGER_SNAPSHOT	Case	= 203
	Case_DELAY_PEER_PORT_TX_RX_LEADER					Case	= 204
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_LEADER				Case	= 205
	Case_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT		Case	= 206
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_LEADER_UNTIL_TRIGGER_SNAPSHOT		Case	= 207
	Case_DELAY_PEER_PORT_TX_RX_QUORUM					Case	= 208
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_QUORUM				Case	= 209
	Case_DELAY_PEER_PORT_TX_RX_ALL						Case	= 210
	Case_RANDOM_DELAY_PEER_PORT_TX_RX_ALL					Case	= 211
	Case_NO_FAIL_WITH_STRESS						Case	= 300
	Case_NO_FAIL_WITH_NO_STRESS_FOR_LIVENESS				Case	= 301
	Case_FAILPOINTS								Case	= 400
	Case_EXTERNAL								Case	= 500
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
	return fileDescriptorRpc, []int{3}
}

type Request struct {
	Operation	Operation	`protobuf:"varint,1,opt,name=Operation,proto3,enum=rpcpb.Operation" json:"Operation,omitempty"`
	Member		*Member		`protobuf:"bytes,2,opt,name=Member" json:"Member,omitempty"`
	Tester		*Tester		`protobuf:"bytes,3,opt,name=Tester" json:"Tester,omitempty"`
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
	MemberName		string		`protobuf:"bytes,1,opt,name=MemberName,proto3" json:"MemberName,omitempty"`
	MemberClientURLs	[]string	`protobuf:"bytes,2,rep,name=MemberClientURLs" json:"MemberClientURLs,omitempty"`
	SnapshotPath		string		`protobuf:"bytes,3,opt,name=SnapshotPath,proto3" json:"SnapshotPath,omitempty"`
	SnapshotFileSize	string		`protobuf:"bytes,4,opt,name=SnapshotFileSize,proto3" json:"SnapshotFileSize,omitempty"`
	SnapshotTotalSize	string		`protobuf:"bytes,5,opt,name=SnapshotTotalSize,proto3" json:"SnapshotTotalSize,omitempty"`
	SnapshotTotalKey	int64		`protobuf:"varint,6,opt,name=SnapshotTotalKey,proto3" json:"SnapshotTotalKey,omitempty"`
	SnapshotHash		int64		`protobuf:"varint,7,opt,name=SnapshotHash,proto3" json:"SnapshotHash,omitempty"`
	SnapshotRevision	int64		`protobuf:"varint,8,opt,name=SnapshotRevision,proto3" json:"SnapshotRevision,omitempty"`
	Took			string		`protobuf:"bytes,9,opt,name=Took,proto3" json:"Took,omitempty"`
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
	Success		bool		`protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Status		string		`protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
	Member		*Member		`protobuf:"bytes,3,opt,name=Member" json:"Member,omitempty"`
	SnapshotInfo	*SnapshotInfo	`protobuf:"bytes,4,opt,name=SnapshotInfo" json:"SnapshotInfo,omitempty"`
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
	EtcdExec		string		`protobuf:"bytes,1,opt,name=EtcdExec,proto3" json:"EtcdExec,omitempty" yaml:"etcd-exec"`
	AgentAddr		string		`protobuf:"bytes,11,opt,name=AgentAddr,proto3" json:"AgentAddr,omitempty" yaml:"agent-addr"`
	FailpointHTTPAddr	string		`protobuf:"bytes,12,opt,name=FailpointHTTPAddr,proto3" json:"FailpointHTTPAddr,omitempty" yaml:"failpoint-http-addr"`
	BaseDir			string		`protobuf:"bytes,101,opt,name=BaseDir,proto3" json:"BaseDir,omitempty" yaml:"base-dir"`
	EtcdClientProxy		bool		`protobuf:"varint,201,opt,name=EtcdClientProxy,proto3" json:"EtcdClientProxy,omitempty" yaml:"etcd-client-proxy"`
	EtcdPeerProxy		bool		`protobuf:"varint,202,opt,name=EtcdPeerProxy,proto3" json:"EtcdPeerProxy,omitempty" yaml:"etcd-peer-proxy"`
	EtcdClientEndpoint	string		`protobuf:"bytes,301,opt,name=EtcdClientEndpoint,proto3" json:"EtcdClientEndpoint,omitempty" yaml:"etcd-client-endpoint"`
	Etcd			*Etcd		`protobuf:"bytes,302,opt,name=Etcd" json:"Etcd,omitempty" yaml:"etcd"`
	EtcdOnSnapshotRestore	*Etcd		`protobuf:"bytes,303,opt,name=EtcdOnSnapshotRestore" json:"EtcdOnSnapshotRestore,omitempty"`
	ClientCertData		string		`protobuf:"bytes,401,opt,name=ClientCertData,proto3" json:"ClientCertData,omitempty" yaml:"client-cert-data"`
	ClientCertPath		string		`protobuf:"bytes,402,opt,name=ClientCertPath,proto3" json:"ClientCertPath,omitempty" yaml:"client-cert-path"`
	ClientKeyData		string		`protobuf:"bytes,403,opt,name=ClientKeyData,proto3" json:"ClientKeyData,omitempty" yaml:"client-key-data"`
	ClientKeyPath		string		`protobuf:"bytes,404,opt,name=ClientKeyPath,proto3" json:"ClientKeyPath,omitempty" yaml:"client-key-path"`
	ClientTrustedCAData	string		`protobuf:"bytes,405,opt,name=ClientTrustedCAData,proto3" json:"ClientTrustedCAData,omitempty" yaml:"client-trusted-ca-data"`
	ClientTrustedCAPath	string		`protobuf:"bytes,406,opt,name=ClientTrustedCAPath,proto3" json:"ClientTrustedCAPath,omitempty" yaml:"client-trusted-ca-path"`
	PeerCertData		string		`protobuf:"bytes,501,opt,name=PeerCertData,proto3" json:"PeerCertData,omitempty" yaml:"peer-cert-data"`
	PeerCertPath		string		`protobuf:"bytes,502,opt,name=PeerCertPath,proto3" json:"PeerCertPath,omitempty" yaml:"peer-cert-path"`
	PeerKeyData		string		`protobuf:"bytes,503,opt,name=PeerKeyData,proto3" json:"PeerKeyData,omitempty" yaml:"peer-key-data"`
	PeerKeyPath		string		`protobuf:"bytes,504,opt,name=PeerKeyPath,proto3" json:"PeerKeyPath,omitempty" yaml:"peer-key-path"`
	PeerTrustedCAData	string		`protobuf:"bytes,505,opt,name=PeerTrustedCAData,proto3" json:"PeerTrustedCAData,omitempty" yaml:"peer-trusted-ca-data"`
	PeerTrustedCAPath	string		`protobuf:"bytes,506,opt,name=PeerTrustedCAPath,proto3" json:"PeerTrustedCAPath,omitempty" yaml:"peer-trusted-ca-path"`
	SnapshotPath		string		`protobuf:"bytes,601,opt,name=SnapshotPath,proto3" json:"SnapshotPath,omitempty" yaml:"snapshot-path"`
	SnapshotInfo		*SnapshotInfo	`protobuf:"bytes,602,opt,name=SnapshotInfo" json:"SnapshotInfo,omitempty"`
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
	DataDir			string		`protobuf:"bytes,1,opt,name=DataDir,proto3" json:"DataDir,omitempty" yaml:"data-dir"`
	Network			string		`protobuf:"bytes,2,opt,name=Network,proto3" json:"Network,omitempty" yaml:"network"`
	Addr			string		`protobuf:"bytes,3,opt,name=Addr,proto3" json:"Addr,omitempty" yaml:"addr"`
	DelayLatencyMs		uint32		`protobuf:"varint,11,opt,name=DelayLatencyMs,proto3" json:"DelayLatencyMs,omitempty" yaml:"delay-latency-ms"`
	DelayLatencyMsRv	uint32		`protobuf:"varint,12,opt,name=DelayLatencyMsRv,proto3" json:"DelayLatencyMsRv,omitempty" yaml:"delay-latency-ms-rv"`
	UpdatedDelayLatencyMs	uint32		`protobuf:"varint,13,opt,name=UpdatedDelayLatencyMs,proto3" json:"UpdatedDelayLatencyMs,omitempty" yaml:"updated-delay-latency-ms"`
	RoundLimit		int32		`protobuf:"varint,21,opt,name=RoundLimit,proto3" json:"RoundLimit,omitempty" yaml:"round-limit"`
	ExitOnCaseFail		bool		`protobuf:"varint,22,opt,name=ExitOnCaseFail,proto3" json:"ExitOnCaseFail,omitempty" yaml:"exit-on-failure"`
	EnablePprof		bool		`protobuf:"varint,23,opt,name=EnablePprof,proto3" json:"EnablePprof,omitempty" yaml:"enable-pprof"`
	CaseDelayMs		uint32		`protobuf:"varint,31,opt,name=CaseDelayMs,proto3" json:"CaseDelayMs,omitempty" yaml:"case-delay-ms"`
	CaseShuffle		bool		`protobuf:"varint,32,opt,name=CaseShuffle,proto3" json:"CaseShuffle,omitempty" yaml:"case-shuffle"`
	Cases			[]string	`protobuf:"bytes,33,rep,name=Cases" json:"Cases,omitempty" yaml:"cases"`
	FailpointCommands	[]string	`protobuf:"bytes,34,rep,name=FailpointCommands" json:"FailpointCommands,omitempty" yaml:"failpoint-commands"`
	RunnerExecPath		string		`protobuf:"bytes,41,opt,name=RunnerExecPath,proto3" json:"RunnerExecPath,omitempty" yaml:"runner-exec-path"`
	ExternalExecPath	string		`protobuf:"bytes,42,opt,name=ExternalExecPath,proto3" json:"ExternalExecPath,omitempty" yaml:"external-exec-path"`
	Stressers		[]*Stresser	`protobuf:"bytes,101,rep,name=Stressers" json:"Stressers,omitempty" yaml:"stressers"`
	Checkers		[]string	`protobuf:"bytes,102,rep,name=Checkers" json:"Checkers,omitempty" yaml:"checkers"`
	StressKeySize		int32		`protobuf:"varint,201,opt,name=StressKeySize,proto3" json:"StressKeySize,omitempty" yaml:"stress-key-size"`
	StressKeySizeLarge	int32		`protobuf:"varint,202,opt,name=StressKeySizeLarge,proto3" json:"StressKeySizeLarge,omitempty" yaml:"stress-key-size-large"`
	StressKeySuffixRange	int32		`protobuf:"varint,203,opt,name=StressKeySuffixRange,proto3" json:"StressKeySuffixRange,omitempty" yaml:"stress-key-suffix-range"`
	StressKeySuffixRangeTxn	int32		`protobuf:"varint,204,opt,name=StressKeySuffixRangeTxn,proto3" json:"StressKeySuffixRangeTxn,omitempty" yaml:"stress-key-suffix-range-txn"`
	StressKeyTxnOps		int32		`protobuf:"varint,205,opt,name=StressKeyTxnOps,proto3" json:"StressKeyTxnOps,omitempty" yaml:"stress-key-txn-ops"`
	StressClients		int32		`protobuf:"varint,301,opt,name=StressClients,proto3" json:"StressClients,omitempty" yaml:"stress-clients"`
	StressQPS		int32		`protobuf:"varint,302,opt,name=StressQPS,proto3" json:"StressQPS,omitempty" yaml:"stress-qps"`
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

type Stresser struct {
	Type	string	`protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty" yaml:"type"`
	Weight	float64	`protobuf:"fixed64,2,opt,name=Weight,proto3" json:"Weight,omitempty" yaml:"weight"`
}

func (m *Stresser) Reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*m = Stresser{}
}
func (m *Stresser) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return proto.CompactTextString(m)
}
func (*Stresser) ProtoMessage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (*Stresser) Descriptor() ([]byte, []int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fileDescriptorRpc, []int{5}
}

type Etcd struct {
	Name			string		`protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty" yaml:"name"`
	DataDir			string		`protobuf:"bytes,2,opt,name=DataDir,proto3" json:"DataDir,omitempty" yaml:"data-dir"`
	WALDir			string		`protobuf:"bytes,3,opt,name=WALDir,proto3" json:"WALDir,omitempty" yaml:"wal-dir"`
	HeartbeatIntervalMs	int64		`protobuf:"varint,11,opt,name=HeartbeatIntervalMs,proto3" json:"HeartbeatIntervalMs,omitempty" yaml:"heartbeat-interval"`
	ElectionTimeoutMs	int64		`protobuf:"varint,12,opt,name=ElectionTimeoutMs,proto3" json:"ElectionTimeoutMs,omitempty" yaml:"election-timeout"`
	ListenClientURLs	[]string	`protobuf:"bytes,21,rep,name=ListenClientURLs" json:"ListenClientURLs,omitempty" yaml:"listen-client-urls"`
	AdvertiseClientURLs	[]string	`protobuf:"bytes,22,rep,name=AdvertiseClientURLs" json:"AdvertiseClientURLs,omitempty" yaml:"advertise-client-urls"`
	ClientAutoTLS		bool		`protobuf:"varint,23,opt,name=ClientAutoTLS,proto3" json:"ClientAutoTLS,omitempty" yaml:"auto-tls"`
	ClientCertAuth		bool		`protobuf:"varint,24,opt,name=ClientCertAuth,proto3" json:"ClientCertAuth,omitempty" yaml:"client-cert-auth"`
	ClientCertFile		string		`protobuf:"bytes,25,opt,name=ClientCertFile,proto3" json:"ClientCertFile,omitempty" yaml:"cert-file"`
	ClientKeyFile		string		`protobuf:"bytes,26,opt,name=ClientKeyFile,proto3" json:"ClientKeyFile,omitempty" yaml:"key-file"`
	ClientTrustedCAFile	string		`protobuf:"bytes,27,opt,name=ClientTrustedCAFile,proto3" json:"ClientTrustedCAFile,omitempty" yaml:"trusted-ca-file"`
	ListenPeerURLs		[]string	`protobuf:"bytes,31,rep,name=ListenPeerURLs" json:"ListenPeerURLs,omitempty" yaml:"listen-peer-urls"`
	AdvertisePeerURLs	[]string	`protobuf:"bytes,32,rep,name=AdvertisePeerURLs" json:"AdvertisePeerURLs,omitempty" yaml:"initial-advertise-peer-urls"`
	PeerAutoTLS		bool		`protobuf:"varint,33,opt,name=PeerAutoTLS,proto3" json:"PeerAutoTLS,omitempty" yaml:"peer-auto-tls"`
	PeerClientCertAuth	bool		`protobuf:"varint,34,opt,name=PeerClientCertAuth,proto3" json:"PeerClientCertAuth,omitempty" yaml:"peer-client-cert-auth"`
	PeerCertFile		string		`protobuf:"bytes,35,opt,name=PeerCertFile,proto3" json:"PeerCertFile,omitempty" yaml:"peer-cert-file"`
	PeerKeyFile		string		`protobuf:"bytes,36,opt,name=PeerKeyFile,proto3" json:"PeerKeyFile,omitempty" yaml:"peer-key-file"`
	PeerTrustedCAFile	string		`protobuf:"bytes,37,opt,name=PeerTrustedCAFile,proto3" json:"PeerTrustedCAFile,omitempty" yaml:"peer-trusted-ca-file"`
	InitialCluster		string		`protobuf:"bytes,41,opt,name=InitialCluster,proto3" json:"InitialCluster,omitempty" yaml:"initial-cluster"`
	InitialClusterState	string		`protobuf:"bytes,42,opt,name=InitialClusterState,proto3" json:"InitialClusterState,omitempty" yaml:"initial-cluster-state"`
	InitialClusterToken	string		`protobuf:"bytes,43,opt,name=InitialClusterToken,proto3" json:"InitialClusterToken,omitempty" yaml:"initial-cluster-token"`
	SnapshotCount		int64		`protobuf:"varint,51,opt,name=SnapshotCount,proto3" json:"SnapshotCount,omitempty" yaml:"snapshot-count"`
	QuotaBackendBytes	int64		`protobuf:"varint,52,opt,name=QuotaBackendBytes,proto3" json:"QuotaBackendBytes,omitempty" yaml:"quota-backend-bytes"`
	PreVote			bool		`protobuf:"varint,63,opt,name=PreVote,proto3" json:"PreVote,omitempty" yaml:"pre-vote"`
	InitialCorruptCheck	bool		`protobuf:"varint,64,opt,name=InitialCorruptCheck,proto3" json:"InitialCorruptCheck,omitempty" yaml:"initial-corrupt-check"`
	Logger			string		`protobuf:"bytes,71,opt,name=Logger,proto3" json:"Logger,omitempty" yaml:"logger"`
	LogOutputs		[]string	`protobuf:"bytes,72,rep,name=LogOutputs" json:"LogOutputs,omitempty" yaml:"log-outputs"`
	Debug			bool		`protobuf:"varint,73,opt,name=Debug,proto3" json:"Debug,omitempty" yaml:"debug"`
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
	return fileDescriptorRpc, []int{6}
}
func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterType((*Request)(nil), "rpcpb.Request")
	proto.RegisterType((*SnapshotInfo)(nil), "rpcpb.SnapshotInfo")
	proto.RegisterType((*Response)(nil), "rpcpb.Response")
	proto.RegisterType((*Member)(nil), "rpcpb.Member")
	proto.RegisterType((*Tester)(nil), "rpcpb.Tester")
	proto.RegisterType((*Stresser)(nil), "rpcpb.Stresser")
	proto.RegisterType((*Etcd)(nil), "rpcpb.Etcd")
	proto.RegisterEnum("rpcpb.StresserType", StresserType_name, StresserType_value)
	proto.RegisterEnum("rpcpb.Checker", Checker_name, Checker_value)
	proto.RegisterEnum("rpcpb.Operation", Operation_name, Operation_value)
	proto.RegisterEnum("rpcpb.Case", Case_name, Case_value)
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
	if len(m.EtcdExec) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.EtcdExec)))
		i += copy(dAtA[i:], m.EtcdExec)
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
		for _, msg := range m.Stressers {
			dAtA[i] = 0xaa
			i++
			dAtA[i] = 0x6
			i++
			i = encodeVarintRpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
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
func (m *Stresser) Marshal() (dAtA []byte, err error) {
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
func (m *Stresser) MarshalTo(dAtA []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.Weight != 0 {
		dAtA[i] = 0x11
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Weight))))
		i += 8
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
	if len(m.Logger) > 0 {
		dAtA[i] = 0xba
		i++
		dAtA[i] = 0x4
		i++
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Logger)))
		i += copy(dAtA[i:], m.Logger)
	}
	if len(m.LogOutputs) > 0 {
		for _, s := range m.LogOutputs {
			dAtA[i] = 0xc2
			i++
			dAtA[i] = 0x4
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
	if m.Debug {
		dAtA[i] = 0xc8
		i++
		dAtA[i] = 0x4
		i++
		if m.Debug {
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
	l = len(m.EtcdExec)
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
		for _, e := range m.Stressers {
			l = e.Size()
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
func (m *Stresser) Size() (n int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Weight != 0 {
		n += 9
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
	l = len(m.Logger)
	if l > 0 {
		n += 2 + l + sovRpc(uint64(l))
	}
	if len(m.LogOutputs) > 0 {
		for _, s := range m.LogOutputs {
			l = len(s)
			n += 2 + l + sovRpc(uint64(l))
		}
	}
	if m.Debug {
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
				return fmt.Errorf("proto: wrong wireType = %d for field EtcdExec", wireType)
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
			m.EtcdExec = string(dAtA[iNdEx:postIndex])
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
			m.Stressers = append(m.Stressers, &Stresser{})
			if err := m.Stressers[len(m.Stressers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *Stresser) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Stresser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stresser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
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
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Weight = float64(math.Float64frombits(v))
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
		case 71:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logger", wireType)
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
			m.Logger = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 72:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LogOutputs", wireType)
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
			m.LogOutputs = append(m.LogOutputs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 73:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debug", wireType)
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
			m.Debug = bool(v != 0)
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
	ErrInvalidLengthRpc	= fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc	= fmt.Errorf("proto: integer overflow")
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proto.RegisterFile("rpcpb/rpc.proto", fileDescriptorRpc)
}

var fileDescriptorRpc = []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x59, 0x4b, 0x73, 0xdc, 0xc6, 0xb5, 0x16, 0x38, 0x24, 0x45, 0x36, 0x5f, 0x60, 0x53, 0x94, 0xa0, 0x17, 0x41, 0x41, 0x96, 0x2f, 0x45, 0x1b, 0x92, 0xaf, 0xe4, 0xf2, 0x43, 0xbe, 0xb6, 0x8c, 0x19, 0x82, 0xe4, 0xdc, 0x01, 0x67, 0x46, 0x3d, 0x20, 0x25, 0xdf, 0x0d, 0x0a, 0x9c, 0x69, 0x92, 0x53, 0x1a, 0x02, 0x63, 0xa0, 0x47, 0x26, 0xbd, 0xbe, 0x55, 0xd9, 0x26, 0xce, 0xa3, 0x92, 0xaa, 0xfc, 0x84, 0x38, 0x59, 0xe4, 0x4f, 0xc8, 0xaf, 0xc4, 0x49, 0x56, 0xc9, 0x62, 0x2a, 0x71, 0x36, 0x59, 0x65, 0x31, 0x95, 0xf7, 0x22, 0x95, 0xea, 0x6e, 0x80, 0xd3, 0x00, 0x66, 0x28, 0xae, 0x48, 0x9c, 0xf3, 0x7d, 0x5f, 0x9f, 0xee, 0xd3, 0xdd, 0xe7, 0x34, 0x09, 0xe6, 0x82, 0x76, 0xbd, 0xbd, 0x7b, 0x37, 0x68, 0xd7, 0xef, 0xb4, 0x03, 0x9f, 0xf8, 0x70, 0x8c, 0x19, 0xae, 0xe8, 0xfb, 0x4d, 0x72, 0xd0, 0xd9, 0xbd, 0x53, 0xf7, 0x0f, 0xef, 0xee, 0xfb, 0xfb, 0xfe, 0x5d, 0xe6, 0xdd, 0xed, 0xec, 0xb1, 0x2f, 0xf6, 0xc1, 0x7e, 0xe3, 0x2c, 0xed, 0x5b, 0x12, 0x38, 0x8f, 0xf0, 0x87, 0x1d, 0x1c, 0x12, 0x78, 0x07, 0x4c, 0x56, 0xda, 0x38, 0x70, 0x49, 0xd3, 0xf7, 0x14, 0x69, 0x59, 0x5a, 0x99, 0xbd, 0x27, 0xdf, 0x61, 0xaa, 0x77, 0x4e, 0xec, 0xa8, 0x0f, 0x81, 0xb7, 0xc0, 0xf8, 0x16, 0x3e, 0xdc, 0xc5, 0x81, 0x32, 0xb2, 0x2c, 0xad, 0x4c, 0xdd, 0x9b, 0x89, 0xc0, 0xdc, 0x88, 0x22, 0x27, 0x85, 0xd9, 0x38, 0x24, 0x38, 0x50, 0x72, 0x09, 0x18, 0x37, 0xa2, 0xc8, 0xa9, 0xfd, 0x69, 0x04, 0x4c, 0xd7, 0x3c, 0xb7, 0x1d, 0x1e, 0xf8, 0xa4, 0xe8, 0xed, 0xf9, 0x70, 0x09, 0x00, 0xae, 0x50, 0x76, 0x0f, 0x31, 0x8b, 0x67, 0x12, 0x09, 0x16, 0xb8, 0x0a, 0x64, 0xfe, 0x55, 0x68, 0x35, 0xb1, 0x47, 0xb6, 0x91, 0x15, 0x2a, 0x23, 0xcb, 0xb9, 0x95, 0x49, 0x94, 0xb1, 0x43, 0xad, 0xaf, 0x5d, 0x75, 0xc9, 0x01, 0x8b, 0x64, 0x12, 0x25, 0x6c, 0x54, 0x2f, 0xfe, 0x5e, 0x6f, 0xb6, 0x70, 0xad, 0xf9, 0x31, 0x56, 0x46, 0x19, 0x2e, 0x63, 0x87, 0xaf, 0x82, 0xf9, 0xd8, 0x66, 0xfb, 0xc4, 0x6d, 0x31, 0xf0, 0x18, 0x03, 0x67, 0x1d, 0xa2, 0x32, 0x33, 0x96, 0xf0, 0xb1, 0x32, 0xbe, 0x2c, 0xad, 0xe4, 0x50, 0xc6, 0x2e, 0x46, 0xba, 0xe9, 0x86, 0x07, 0xca, 0x79, 0x86, 0x4b, 0xd8, 0x44, 0x3d, 0x84, 0x9f, 0x35, 0x43, 0x9a, 0xaf, 0x89, 0xa4, 0x5e, 0x6c, 0x87, 0x10, 0x8c, 0xda, 0xbe, 0xff, 0x54, 0x99, 0x64, 0xc1, 0xb1, 0xdf, 0xb5, 0x1f, 0x4b, 0x60, 0x02, 0xe1, 0xb0, 0xed, 0x7b, 0x21, 0x86, 0x0a, 0x38, 0x5f, 0xeb, 0xd4, 0xeb, 0x38, 0x0c, 0xd9, 0x1a, 0x4f, 0xa0, 0xf8, 0x13, 0x5e, 0x04, 0xe3, 0x35, 0xe2, 0x92, 0x4e, 0xc8, 0xf2, 0x3b, 0x89, 0xa2, 0x2f, 0x21, 0xef, 0xb9, 0xd3, 0xf2, 0xfe, 0x66, 0x32, 0x9f, 0x6c, 0x2d, 0xa7, 0xee, 0x2d, 0x44, 0x60, 0xd1, 0x85, 0x12, 0x40, 0xed, 0x93, 0xe9, 0x78, 0x00, 0xf8, 0x1a, 0x98, 0x30, 0x49, 0xbd, 0x61, 0x1e, 0xe1, 0x3a, 0xdf, 0x01, 0xf9, 0x0b, 0xbd, 0xae, 0x2a, 0x1f, 0xbb, 0x87, 0xad, 0x07, 0x1a, 0x26, 0xf5, 0x86, 0x8e, 0x8f, 0x70, 0x5d, 0x43, 0x27, 0x28, 0x78, 0x1f, 0x4c, 0x1a, 0xfb, 0xd8, 0x23, 0x46, 0xa3, 0x11, 0x28, 0x53, 0x8c, 0xb2, 0xd8, 0xeb, 0xaa, 0xf3, 0x9c, 0xe2, 0x52, 0x97, 0xee, 0x36, 0x1a, 0x81, 0x86, 0xfa, 0x38, 0x68, 0x81, 0xf9, 0x75, 0xb7, 0xd9, 0x6a, 0xfb, 0x4d, 0x8f, 0x6c, 0xda, 0x76, 0x95, 0x91, 0xa7, 0x19, 0x79, 0xa9, 0xd7, 0x55, 0xaf, 0x70, 0xf2, 0x5e, 0x0c, 0xd1, 0x0f, 0x08, 0x69, 0x47, 0x2a, 0x59, 0x22, 0xd4, 0xc1, 0xf9, 0xbc, 0x1b, 0xe2, 0xb5, 0x66, 0xa0, 0x60, 0xa6, 0xb1, 0xd0, 0xeb, 0xaa, 0x73, 0x5c, 0x63, 0xd7, 0x0d, 0xb1, 0xde, 0x68, 0x06, 0x1a, 0x8a, 0x31, 0x70, 0x03, 0xcc, 0xd1, 0xe8, 0xf9, 0x6e, 0xad, 0x06, 0xfe, 0xd1, 0xb1, 0xf2, 0x19, 0xcb, 0x44, 0xfe, 0x5a, 0xaf, 0xab, 0x2a, 0xc2, 0x5c, 0xeb, 0x0c, 0xa2, 0xb7, 0x29, 0x46, 0x43, 0x69, 0x16, 0x34, 0xc0, 0x0c, 0x35, 0x55, 0x31, 0x0e, 0xb8, 0xcc, 0xe7, 0x5c, 0xe6, 0x4a, 0xaf, 0xab, 0x5e, 0x14, 0x64, 0xda, 0x18, 0x07, 0xb1, 0x48, 0x92, 0x01, 0xab, 0x00, 0xf6, 0x55, 0x4d, 0xaf, 0xc1, 0x26, 0xa6, 0x7c, 0xca, 0xf2, 0x9f, 0x57, 0x7b, 0x5d, 0xf5, 0x6a, 0x36, 0x1c, 0x1c, 0xc1, 0x34, 0x34, 0x80, 0x0b, 0xff, 0x1b, 0x8c, 0x52, 0xab, 0xf2, 0x53, 0x7e, 0x47, 0x4c, 0x45, 0xe9, 0xa7, 0xb6, 0xfc, 0x5c, 0xaf, 0xab, 0x4e, 0xf5, 0x05, 0x35, 0xc4, 0xa0, 0x30, 0x0f, 0x16, 0xe9, 0xcf, 0x8a, 0xd7, 0xdf, 0xcc, 0x21, 0xf1, 0x03, 0xac, 0xfc, 0x2c, 0xab, 0x81, 0x06, 0x43, 0xe1, 0x1a, 0x98, 0xe5, 0x81, 0x14, 0x70, 0x40, 0xd6, 0x5c, 0xe2, 0x2a, 0xdf, 0x61, 0x67, 0x3e, 0x7f, 0xb5, 0xd7, 0x55, 0x2f, 0xf1, 0x31, 0xa3, 0xf8, 0xeb, 0x38, 0x20, 0x7a, 0xc3, 0x25, 0xae, 0x86, 0x52, 0x9c, 0xa4, 0x0a, 0xbb, 0x38, 0x3e, 0x39, 0x55, 0xa5, 0xed, 0x92, 0x83, 0x84, 0x0a, 0xbb, 0x58, 0x0c, 0x30, 0xc3, 0x2d, 0x25, 0x7c, 0xcc, 0x42, 0xf9, 0x2e, 0x17, 0x11, 0xf2, 0x12, 0x89, 0x3c, 0xc5, 0xc7, 0x51, 0x24, 0x49, 0x46, 0x42, 0x82, 0xc5, 0xf1, 0xbd, 0xd3, 0x24, 0x78, 0x18, 0x49, 0x06, 0xb4, 0xc1, 0x02, 0x37, 0xd8, 0x41, 0x27, 0x24, 0xb8, 0x51, 0x30, 0x58, 0x2c, 0xdf, 0xe7, 0x42, 0x37, 0x7a, 0x5d, 0xf5, 0x7a, 0x42, 0x88, 0x70, 0x98, 0x5e, 0x77, 0xa3, 0x90, 0x06, 0xd1, 0x07, 0xa8, 0xb2, 0xf0, 0x7e, 0x70, 0x06, 0x55, 0x1e, 0xe5, 0x20, 0x3a, 0x7c, 0x0f, 0x4c, 0xd3, 0x3d, 0x79, 0x92, 0xbb, 0xbf, 0x72, 0xb9, 0xcb, 0xbd, 0xae, 0xba, 0xc8, 0xe5, 0xd8, 0x1e, 0x16, 0x32, 0x97, 0xc0, 0x8b, 0x7c, 0x16, 0xce, 0xdf, 0x4e, 0xe1, 0xf3, 0x30, 0x12, 0x78, 0xf8, 0x0e, 0x98, 0xa2, 0xdf, 0x71, 0xbe, 0xfe, 0xce, 0xe9, 0x4a, 0xaf, 0xab, 0x5e, 0x10, 0xe8, 0xfd, 0x6c, 0x89, 0x68, 0x81, 0xcc, 0xc6, 0xfe, 0xc7, 0x70, 0x32, 0x1f, 0x5a, 0x44, 0xc3, 0x32, 0x98, 0xa7, 0x9f, 0xc9, 0x1c, 0xfd, 0x33, 0x97, 0x3e, 0x7f, 0x4c, 0x22, 0x93, 0xa1, 0x2c, 0x35, 0xa3, 0xc7, 0x42, 0xfa, 0xd7, 0x0b, 0xf5, 0x78, 0x64, 0x59, 0x2a, 0x7c, 0x37, 0x55, 0x48, 0x7f, 0x3b, 0x9a, 0x9e, 0x5d, 0x18, 0xb9, 0xe3, 0x85, 0x4d, 0xd4, 0xd8, 0xb7, 0x52, 0x35, 0xe1, 0x77, 0x67, 0x2e, 0x0a, 0x3f, 0x9f, 0x8e, 0xdb, 0x08, 0x7a, 0xbf, 0xd2, 0xb9, 0xd1, 0xfb, 0x55, 0x4a, 0xdf, 0xaf, 0x74, 0x21, 0xa2, 0xfb, 0x35, 0xc2, 0xc0, 0x57, 0xc1, 0xf9, 0x32, 0x26, 0x1f, 0xf9, 0xc1, 0x53, 0x5e, 0xc7, 0xf2, 0xb0, 0xd7, 0x55, 0x67, 0x39, 0xdc, 0xe3, 0x0e, 0x0d, 0xc5, 0x10, 0x78, 0x13, 0x8c, 0xb2, 0xdb, 0x9f, 0x2f, 0x91, 0x70, 0x43, 0xf1, 0xeb, 0x9e, 0x39, 0x61, 0x01, 0xcc, 0xae, 0xe1, 0x96, 0x7b, 0x6c, 0xb9, 0x04, 0x7b, 0xf5, 0xe3, 0xad, 0x90, 0x55, 0x9a, 0x19, 0xf1, 0x5a, 0x68, 0x50, 0xbf, 0xde, 0xe2, 0x00, 0xfd, 0x30, 0xd4, 0x50, 0x8a, 0x02, 0xff, 0x17, 0xc8, 0x49, 0x0b, 0x7a, 0xc6, 0x6a, 0xce, 0x8c, 0x58, 0x73, 0xd2, 0x32, 0x7a, 0xf0, 0x4c, 0x43, 0x19, 0x1e, 0xfc, 0x00, 0x2c, 0x6e, 0xb7, 0x1b, 0x2e, 0xc1, 0x8d, 0x54, 0x5c, 0x33, 0x4c, 0xf0, 0x66, 0xaf, 0xab, 0xaa, 0x5c, 0xb0, 0xc3, 0x61, 0x7a, 0x36, 0xbe, 0xc1, 0x0a, 0xf0, 0x0d, 0x00, 0x90, 0xdf, 0xf1, 0x1a, 0x56, 0xf3, 0xb0, 0x49, 0x94, 0xc5, 0x65, 0x69, 0x65, 0x2c, 0x7f, 0xb1, 0xd7, 0x55, 0x21, 0xd7, 0x0b, 0xa8, 0x4f, 0x6f, 0x51, 0xa7, 0x86, 0x04, 0x24, 0xcc, 0x83, 0x59, 0xf3, 0xa8, 0x49, 0x2a, 0x5e, 0xc1, 0x0d, 0x31, 0x2d, 0x92, 0xca, 0xc5, 0x4c, 0x35, 0x3a, 0x6a, 0x12, 0xdd, 0xf7, 0x74, 0x5a, 0x58, 0x3b, 0x01, 0xd6, 0x50, 0x8a, 0x01, 0xdf, 0x06, 0x53, 0xa6, 0xe7, 0xee, 0xb6, 0x70, 0xb5, 0x1d, 0xf8, 0x7b, 0xca, 0x25, 0x26, 0x70, 0xa9, 0xd7, 0x55, 0x17, 0x22, 0x01, 0xe6, 0xd4, 0xdb, 0xd4, 0xab, 0x21, 0x11, 0x0b, 0x1f, 0x80, 0x29, 0x2a, 0xc3, 0x26, 0xb3, 0x15, 0x2a, 0x2a, 0x5b, 0x07, 0x61, 0x9b, 0xd6, 0x59, 0x21, 0x66, 0x8b, 0x40, 0x27, 0x2f, 0x82, 0xe9, 0xb0, 0xf4, 0xb3, 0x76, 0xd0, 0xd9, 0xdb, 0x6b, 0x61, 0x65, 0x39, 0x3d, 0x2c, 0xe3, 0x86, 0xdc, 0x1b, 0x51, 0x23, 0x2c, 0x7c, 0x19, 0x8c, 0xd1, 0xcf, 0x50, 0xb9, 0x41, 0x3b, 0xd1, 0xbc, 0xdc, 0xeb, 0xaa, 0xd3, 0x7d, 0x52, 0xa8, 0x21, 0xee, 0x86, 0x25, 0xa1, 0xe3, 0x28, 0xf8, 0x87, 0x87, 0xae, 0xd7, 0x08, 0x15, 0x8d, 0x71, 0xae, 0xf7, 0xba, 0xea, 0xe5, 0x74, 0xc7, 0x51, 0x8f, 0x30, 0x62, 0xc3, 0x11, 0xf3, 0xe8, 0x76, 0x44, 0x1d, 0xcf, 0xc3, 0x01, 0xed, 0x80, 0xd8, 0xb1, 0xbc, 0x9d, 0xae, 0x52, 0x01, 0xf3, 0xb3, 0x6e, 0x29, 0xae, 0x52, 0x49, 0x0a, 0x2c, 0x02, 0xd9, 0x3c, 0x22, 0x38, 0xf0, 0xdc, 0xd6, 0x89, 0xcc, 0x2a, 0x93, 0x11, 0x02, 0xc2, 0x11, 0x42, 0x14, 0xca, 0xd0, 0x60, 0x01, 0x4c, 0xd6, 0x48, 0x80, 0xc3, 0x10, 0x07, 0xa1, 0x82, 0x97, 0x73, 0x2b, 0x53, 0xf7, 0xe6, 0xe2, 0x13, 0x1e, 0xd9, 0xc5, 0x3e, 0x2e, 0x8c, 0xb1, 0x1a, 0xea, 0xf3, 0xe0, 0x5d, 0x30, 0x51, 0x38, 0xc0, 0xf5, 0xa7, 0x54, 0x63, 0x8f, 0x2d, 0x8c, 0x70, 0xcc, 0xeb, 0x91, 0x47, 0x43, 0x27, 0x20, 0x5a, 0x23, 0x39, 0xbb, 0x84, 0x8f, 0x59, 0x3f, 0xce, 0xba, 0xa8, 0x31, 0x71, 0xc3, 0xf1, 0x91, 0xd8, 0xdd, 0x1b, 0x36, 0x3f, 0xc6, 0x1a, 0x4a, 0x32, 0xe0, 0x23, 0x00, 0x13, 0x06, 0xcb, 0x0d, 0xf6, 0x31, 0x6f, 0xa3, 0xc6, 0xf2, 0xcb, 0xbd, 0xae, 0x7a, 0x6d, 0xa0, 0x8e, 0xde, 0xa2, 0x38, 0x0d, 0x0d, 0x20, 0xc3, 0xc7, 0xe0, 0x42, 0xdf, 0xda, 0xd9, 0xdb, 0x6b, 0x1e, 0x21, 0xd7, 0xdb, 0xc7, 0xca, 0x17, 0x5c, 0x54, 0xeb, 0x75, 0xd5, 0xa5, 0xac, 0x28, 0x03, 0xea, 0x01, 0x45, 0x6a, 0x68, 0xa0, 0x00, 0x74, 0xc1, 0xa5, 0x41, 0x76, 0xfb, 0xc8, 0x53, 0xbe, 0xe4, 0xda, 0x2f, 0xf7, 0xba, 0xaa, 0x76, 0xaa, 0xb6, 0x4e, 0x8e, 0x3c, 0x0d, 0x0d, 0xd3, 0x81, 0x9b, 0x60, 0xee, 0xc4, 0x65, 0x1f, 0x79, 0x95, 0x76, 0xa8, 0x7c, 0xc5, 0xa5, 0x85, 0x2d, 0x21, 0x48, 0x93, 0x23, 0x4f, 0xf7, 0xdb, 0xa1, 0x86, 0xd2, 0x34, 0xf8, 0x7e, 0x9c, 0x1b, 0x5e, 0xed, 0x43, 0xde, 0x52, 0x8e, 0x89, 0x15, 0x39, 0xd2, 0xe1, 0x7d, 0x42, 0x78, 0x92, 0x9a, 0x88, 0x00, 0x5f, 0x8f, 0xf7, 0xd4, 0xa3, 0x6a, 0x8d, 0x37, 0x93, 0x63, 0x62, 0x63, 0x1f, 0xb1, 0x3f, 0x6c, 0xf7, 0x37, 0xd1, 0xa3, 0x6a, 0x4d, 0xfb, 0x3f, 0x30, 0x11, 0xef, 0x28, 0x7a, 0xb3, 0xdb, 0xc7, 0xed, 0xe8, 0x25, 0x29, 0xde, 0xec, 0xe4, 0xb8, 0x8d, 0x35, 0xc4, 0x9c, 0xf0, 0x36, 0x18, 0x7f, 0x8c, 0x9b, 0xfb, 0x07, 0x84, 0xd5, 0x0a, 0x29, 0x3f, 0xdf, 0xeb, 0xaa, 0x33, 0x1c, 0xf6, 0x11, 0xb3, 0x6b, 0x28, 0x02, 0x68, 0xff, 0x3f, 0xc7, 0x5b, 0x5b, 0x2a, 0xdc, 0x7f, 0xa2, 0x8a, 0xc2, 0x9e, 0x7b, 0x48, 0x85, 0xd9, 0x6b, 0x55, 0x28, 0x5a, 0x23, 0x67, 0x28, 0x5a, 0xab, 0x60, 0xfc, 0xb1, 0x61, 0x51, 0x74, 0x2e, 0x5d, 0xb3, 0x3e, 0x72, 0x5b, 0x1c, 0x1c, 0x21, 0x60, 0x05, 0x2c, 0x6c, 0x62, 0x37, 0x20, 0xbb, 0xd8, 0x25, 0x45, 0x8f, 0xe0, 0xe0, 0x99, 0xdb, 0x8a, 0x4a, 0x52, 0x4e, 0xcc, 0xd4, 0x41, 0x0c, 0xd2, 0x9b, 0x11, 0x4a, 0x43, 0x83, 0x98, 0xb0, 0x08, 0xe6, 0xcd, 0x16, 0xae, 0xd3, 0x47, 0xbe, 0xdd, 0x3c, 0xc4, 0x7e, 0x87, 0x6c, 0x85, 0xac, 0x34, 0xe5, 0xc4, 0x2b, 0x05, 0x47, 0x10, 0x9d, 0x70, 0x8c, 0x86, 0xb2, 0x2c, 0x7a, 0xab, 0x58, 0xcd, 0x90, 0x60, 0x4f, 0x78, 0xa4, 0x2f, 0xa6, 0xaf, 0xb9, 0x16, 0x43, 0xc4, 0xef, 0x89, 0x4e, 0xd0, 0x0a, 0x35, 0x94, 0xa1, 0x41, 0x04, 0x16, 0x8c, 0xc6, 0x33, 0x1c, 0x90, 0x66, 0x88, 0x05, 0xb5, 0x8b, 0x4c, 0x4d, 0x38, 0x9c, 0x6e, 0x0c, 0x4a, 0x0a, 0x0e, 0x22, 0xc3, 0xb7, 0xe3, 0xbe, 0xda, 0xe8, 0x10, 0xdf, 0xb6, 0x6a, 0x51, 0x89, 0x11, 0x72, 0xe3, 0x76, 0x88, 0xaf, 0x13, 0x2a, 0x90, 0x44, 0xd2, 0x4b, 0xb7, 0xdf, 0xe7, 0x1b, 0x1d, 0x72, 0xa0, 0x28, 0x8c, 0x3b, 0xe4, 0x69, 0xe0, 0x76, 0x52, 0x4f, 0x03, 0x4a, 0x81, 0xff, 0x23, 0x8a, 0xac, 0x37, 0x5b, 0x58, 0xb9, 0x9c, 0x7e, 0xe5, 0x32, 0xf6, 0x5e, 0x93, 0x56, 0x9a, 0x14, 0xb6, 0x1f, 0x7d, 0x09, 0x1f, 0x33, 0xf2, 0x95, 0xf4, 0xce, 0xa2, 0xa7, 0x92, 0x73, 0x93, 0x48, 0x68, 0x65, 0xfa, 0x76, 0x26, 0x70, 0x35, 0xfd, 0xaa, 0x10, 0x7a, 0x42, 0xae, 0x33, 0x88, 0x46, 0xd7, 0x82, 0xa7, 0x8b, 0x36, 0x8c, 0x2c, 0x2b, 0x2a, 0xcb, 0x8a, 0xb0, 0x16, 0x51, 0x8e, 0x59, 0xa3, 0xc9, 0x13, 0x92, 0xa2, 0x40, 0x1b, 0xcc, 0x9f, 0xa4, 0xe8, 0x44, 0x67, 0x99, 0xe9, 0x08, 0x37, 0x59, 0xd3, 0x6b, 0x92, 0xa6, 0xdb, 0xd2, 0xfb, 0x59, 0x16, 0x24, 0xb3, 0x02, 0xb4, 0x0f, 0xa0, 0xbf, 0xc7, 0xf9, 0xbd, 0xc1, 0x72, 0x94, 0x6e, 0xc6, 0xfb, 0x49, 0x16, 0xc1, 0xf4, 0x35, 0xcc, 0x9e, 0x05, 0xc9, 0x34, 0x6b, 0x4c, 0x42, 0xd8, 0x70, 0xfc, 0x2d, 0x91, 0xc9, 0xf5, 0x00, 0x2e, 0x6d, 0x9f, 0xe3, 0x87, 0x06, 0x5b, 0xef, 0x9b, 0xc3, 0xdf, 0x25, 0x7c, 0xb9, 0x13, 0xf0, 0x78, 0x32, 0x71, 0xba, 0x5f, 0x1a, 0xfa, 0xb2, 0xe0, 0x64, 0x11, 0x0c, 0xb7, 0x52, 0x2f, 0x01, 0xa6, 0x70, 0xeb, 0x45, 0x0f, 0x01, 0x2e, 0x94, 0x65, 0xd2, 0xf6, 0xae, 0xc8, 0x53, 0x51, 0x68, 0x75, 0xd8, 0x5f, 0xf7, 0x6e, 0xa7, 0xf7, 0x4e, 0x9c, 0xaa, 0x3a, 0x07, 0x68, 0x28, 0xc5, 0xa0, 0x27, 0x3a, 0x69, 0xa9, 0x11, 0x97, 0xe0, 0xa8, 0xeb, 0x10, 0x16, 0x38, 0x25, 0xa4, 0x87, 0x14, 0xa6, 0xa1, 0x41, 0xe4, 0xac, 0xa6, 0xed, 0x3f, 0xc5, 0x9e, 0xf2, 0xca, 0x8b, 0x34, 0x09, 0x85, 0x65, 0x34, 0x19, 0x19, 0x3e, 0x04, 0x33, 0xf1, 0x5b, 0xa4, 0xe0, 0x77, 0x3c, 0xa2, 0xdc, 0x67, 0x77, 0xa1, 0x58, 0xbc, 0xe2, 0x47, 0x4f, 0x9d, 0xfa, 0x69, 0xf1, 0x12, 0xf1, 0xd0, 0x02, 0xf3, 0x8f, 0x3a, 0x3e, 0x71, 0xf3, 0x6e, 0xfd, 0x29, 0xf6, 0x1a, 0xf9, 0x63, 0x82, 0x43, 0xe5, 0x75, 0x26, 0x22, 0xf4, 0xfa, 0x1f, 0x52, 0x88, 0xbe, 0xcb, 0x31, 0xfa, 0x2e, 0x05, 0x69, 0x28, 0x4b, 0xa4, 0xa5, 0xa4, 0x1a, 0xe0, 0x1d, 0x9f, 0x60, 0xe5, 0x61, 0xfa, 0xba, 0x6a, 0x07, 0x58, 0x7f, 0xe6, 0xd3, 0xd5, 0x89, 0x31, 0xe2, 0x8a, 0xf8, 0x41, 0xd0, 0x69, 0x13, 0xd6, 0x31, 0x29, 0xef, 0xa7, 0xb7, 0xf1, 0xc9, 0x8a, 0x70, 0x94, 0xce, 0x7a, 0x2c, 0x61, 0x45, 0x04, 0x32, 0x2d, 0x93, 0x96, 0xbf, 0xbf, 0x8f, 0x03, 0x65, 0x83, 0x2d, 0xac, 0x50, 0x26, 0x5b, 0xcc, 0xae, 0xa1, 0x08, 0x40, 0xdf, 0x0f, 0x96, 0xbf, 0x5f, 0xe9, 0x90, 0x76, 0x87, 0x84, 0xca, 0x26, 0x3b, 0xcf, 0xc2, 0xfb, 0xa1, 0xe5, 0xef, 0xeb, 0x3e, 0x77, 0x6a, 0x48, 0x40, 0xd2, 0x4e, 0x7a, 0x0d, 0xef, 0x76, 0xf6, 0x95, 0x22, 0x0b, 0x54, 0xe8, 0xa4, 0x1b, 0xd4, 0xac, 0x21, 0xee, 0x5e, 0xfd, 0xb7, 0x04, 0xa6, 0xe3, 0x1a, 0xcf, 0x4a, 0x38, 0x04, 0xb3, 0xa5, 0x1d, 0xe7, 0x31, 0x2a, 0xda, 0xa6, 0x53, 0xdb, 0x32, 0x2c, 0x4b, 0x3e, 0x97, 0xb0, 0x59, 0x06, 0xda, 0x30, 0x65, 0x09, 0x2e, 0x80, 0xb9, 0xd2, 0x8e, 0x83, 0x4c, 0x63, 0xcd, 0xa9, 0x94, 0x4d, 0xa7, 0x64, 0x7e, 0x20, 0x8f, 0xc0, 0x79, 0x30, 0x13, 0x1b, 0x91, 0x51, 0xde, 0x30, 0xe5, 0x1c, 0x5c, 0x04, 0xf3, 0xa5, 0x1d, 0x67, 0xcd, 0xb4, 0x4c, 0xdb, 0x3c, 0x41, 0x8e, 0x46, 0xf4, 0xc8, 0xcc, 0xb1, 0x63, 0xf0, 0x12, 0x58, 0x28, 0xed, 0x38, 0xf6, 0x93, 0x72, 0x34, 0x16, 0x77, 0xcb, 0xe3, 0x70, 0x12, 0x8c, 0x59, 0xa6, 0x51, 0x33, 0x65, 0x40, 0x89, 0xa6, 0x65, 0x16, 0xec, 0x62, 0xa5, 0xec, 0xa0, 0xed, 0x72, 0xd9, 0x44, 0xf2, 0x05, 0x28, 0x83, 0xe9, 0xc7, 0x86, 0x5d, 0xd8, 0x8c, 0x2d, 0x2a, 0x1d, 0xd6, 0xaa, 0x14, 0x4a, 0x0e, 0x32, 0x0a, 0x26, 0x8a, 0xcd, 0xb7, 0x29, 0x90, 0x09, 0xc5, 0x96, 0xfb, 0xab, 0x79, 0x70, 0x3e, 0xea, 0x81, 0xe1, 0x14, 0x38, 0x5f, 0xda, 0x71, 0x36, 0x8d, 0xda, 0xa6, 0x7c, 0xae, 0x8f, 0x34, 0x9f, 0x54, 0x8b, 0x88, 0xce, 0x18, 0x80, 0xf1, 0x88, 0x35, 0x02, 0xa7, 0xc1, 0x44, 0xb9, 0xe2, 0x14, 0x36, 0xcd, 0x42, 0x49, 0xce, 0xad, 0xfe, 0x28, 0x27, 0xfc, 0xed, 0x1f, 0xce, 0x81, 0xa9, 0x72, 0xc5, 0x76, 0x6a, 0xb6, 0x81, 0x6c, 0x73, 0x4d, 0x3e, 0x07, 0x2f, 0x02, 0x58, 0x2c, 0x17, 0xed, 0xa2, 0x61, 0x71, 0xa3, 0x63, 0xda, 0x85, 0x35, 0x19, 0xd0, 0x21, 0x90, 0x29, 0x58, 0xa6, 0xa8, 0xa5, 0x56, 0xdc, 0xb0, 0x4d, 0xb4, 0xc5, 0x2d, 0x17, 0xe0, 0x32, 0xb8, 0x56, 0x2b, 0x6e, 0x3c, 0xda, 0x2e, 0x72, 0x8c, 0x63, 0x94, 0xd7, 0x1c, 0x64, 0x6e, 0x55, 0x76, 0x4c, 0x67, 0xcd, 0xb0, 0x0d, 0x79, 0x91, 0xae, 0x79, 0xcd, 0xd8, 0x31, 0x9d, 0x5a, 0xd9, 0xa8, 0xd6, 0x36, 0x2b, 0xb6, 0xbc, 0x04, 0x6f, 0x80, 0xeb, 0x54, 0xb8, 0x82, 0x4c, 0x27, 0x1e, 0x60, 0x1d, 0x55, 0xb6, 0xfa, 0x10, 0x15, 0x5e, 0x06, 0x8b, 0x83, 0x5d, 0xcb, 0x94, 0x9d, 0x19, 0xd2, 0x40, 0x85, 0xcd, 0x62, 0x3c, 0xe6, 0x0a, 0xbc, 0x0b, 0x5e, 0x39, 0x2d, 0x2a, 0xf6, 0x5d, 0xb3, 0x2b, 0x55, 0xc7, 0xd8, 0x30, 0xcb, 0xb6, 0x7c, 0x1b, 0x5e, 0x07, 0x97, 0xf3, 0x96, 0x51, 0x28, 0x6d, 0x56, 0x2c, 0xd3, 0xa9, 0x9a, 0x26, 0x72, 0xaa, 0x15, 0x64, 0x3b, 0xf6, 0x13, 0x07, 0x3d, 0x91, 0x1b, 0x50, 0x05, 0x57, 0xb7, 0xcb, 0xc3, 0x01, 0x18, 0x5e, 0x01, 0x8b, 0x6b, 0xa6, 0x65, 0x7c, 0x90, 0x71, 0x3d, 0x97, 0xe0, 0x35, 0x70, 0x69, 0xbb, 0x3c, 0xd8, 0xfb, 0x99, 0xb4, 0xfa, 0x67, 0x00, 0x46, 0xe9, 0xa3, 0x11, 0x2a, 0xe0, 0x42, 0xbc, 0xb6, 0x74, 0x1b, 0xae, 0x57, 0x2c, 0xab, 0xf2, 0xd8, 0x44, 0xf2, 0xb9, 0x68, 0x36, 0x19, 0x8f, 0xb3, 0x5d, 0xb6, 0x8b, 0x96, 0x63, 0xa3, 0xe2, 0xc6, 0x86, 0x89, 0xfa, 0x2b, 0x24, 0xd1, 0xf3, 0x10, 0x13, 0x2c, 0xd3, 0x58, 0x63, 0x3b, 0xe2, 0x36, 0xb8, 0x95, 0xb4, 0x0d, 0xa3, 0xe7, 0x44, 0xfa, 0xa3, 0xed, 0x0a, 0xda, 0xde, 0x92, 0x47, 0xe9, 0xa6, 0x89, 0x6d, 0xf4, 0xcc, 0x8d, 0xc1, 0x9b, 0x40, 0x8d, 0x97, 0x58, 0x58, 0xdd, 0x44, 0xe4, 0x00, 0x3e, 0x00, 0x6f, 0xbc, 0x00, 0x34, 0x2c, 0x8a, 0x29, 0x9a, 0x92, 0x01, 0xdc, 0x68, 0x3e, 0xd3, 0xf0, 0x75, 0xf0, 0xda, 0x50, 0xf7, 0x30, 0xd1, 0x19, 0xb8, 0x0e, 0xf2, 0x03, 0x58, 0x7c, 0x96, 0x91, 0x85, 0xef, 0xcb, 0x48, 0x28, 0xa6, 0x46, 0x9b, 0xb0, 0x80, 0xe8, 0x29, 0x96, 0x67, 0xe1, 0x2a, 0x78, 0x79, 0xe8, 0x76, 0x48, 0x2e, 0x42, 0x03, 0x1a, 0xe0, 0xdd, 0xb3, 0x61, 0x87, 0x85, 0x8d, 0xe1, 0x4b, 0x60, 0x79, 0xb8, 0x44, 0xb4, 0x24, 0x7b, 0xf0, 0x1d, 0xf0, 0xe6, 0x8b, 0x50, 0xc3, 0x86, 0xd8, 0x3f, 0x7d, 0x88, 0x68, 0x1b, 0x1c, 0xd0, 0xb3, 0x37, 0x1c, 0x45, 0x37, 0x46, 0x13, 0xfe, 0x17, 0xd0, 0x06, 0x6e, 0xf6, 0xe4, 0xb2, 0x3c, 0x97, 0xe0, 0x1d, 0x70, 0x1b, 0x19, 0xe5, 0xb5, 0xca, 0x96, 0x73, 0x06, 0xfc, 0x67, 0x12, 0x7c, 0x0f, 0xbc, 0xfd, 0x62, 0xe0, 0xb0, 0x09, 0x7e, 0x2e, 0x41, 0x13, 0xbc, 0x7f, 0xe6, 0xf1, 0x86, 0xc9, 0x7c, 0x21, 0xc1, 0x1b, 0xe0, 0xda, 0x60, 0x7e, 0x94, 0x87, 0x2f, 0x25, 0xb8, 0x02, 0x6e, 0x9e, 0x3a, 0x52, 0x84, 0xfc, 0x4a, 0x82, 0x6f, 0x81, 0xfb, 0xa7, 0x41, 0x86, 0x85, 0xf1, 0x0b, 0x09, 0x3e, 0x04, 0x0f, 0xce, 0x30, 0xc6, 0x30, 0x81, 0x5f, 0x9e, 0x32, 0x8f, 0x28, 0xd9, 0x5f, 0xbf, 0x78, 0x1e, 0x11, 0xf2, 0x57, 0x12, 0x5c, 0x02, 0x97, 0x07, 0x43, 0xe8, 0x9e, 0xf8, 0xb5, 0x04, 0x6f, 0x81, 0xe5, 0x53, 0x95, 0x28, 0xec, 0x37, 0x12, 0x54, 0xc0, 0x42, 0xb9, 0xe2, 0xac, 0x1b, 0x45, 0xcb, 0x79, 0x5c, 0xb4, 0x37, 0x9d, 0x9a, 0x8d, 0xcc, 0x5a, 0x4d, 0xfe, 0xc9, 0x08, 0x0d, 0x25, 0xe1, 0x29, 0x57, 0x22, 0xa7, 0xb3, 0x5e, 0x41, 0x8e, 0x55, 0xdc, 0x31, 0xcb, 0x14, 0xf9, 0xe9, 0x08, 0x9c, 0x03, 0x80, 0xc2, 0xaa, 0x95, 0x62, 0xd9, 0xae, 0xc9, 0xdf, 0xce, 0xc1, 0x19, 0x30, 0x61, 0x3e, 0xb1, 0x4d, 0x54, 0x36, 0x2c, 0xf9, 0x2f, 0xb9, 0x7b, 0x0f, 0xc1, 0xa4, 0x1d, 0xb8, 0x5e, 0xd8, 0xf6, 0x03, 0x02, 0xef, 0x89, 0x1f, 0xb3, 0xd1, 0x5f, 0xb1, 0xa2, 0xff, 0x98, 0x5f, 0x99, 0x3b, 0xf9, 0xe6, 0xff, 0x4c, 0xd5, 0xce, 0xad, 0x48, 0xaf, 0x49, 0xf9, 0x0b, 0xcf, 0xff, 0xb0, 0x74, 0xee, 0xf9, 0x37, 0x4b, 0xd2, 0xd7, 0xdf, 0x2c, 0x49, 0xbf, 0xff, 0x66, 0x49, 0xfa, 0xe1, 0x1f, 0x97, 0xce, 0xed, 0x8e, 0xb3, 0xff, 0xb8, 0xdf, 0xff, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x49, 0xef, 0x9b, 0xba, 0x1f, 0x00, 0x00}
