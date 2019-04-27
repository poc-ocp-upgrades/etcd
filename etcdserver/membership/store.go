package membership

import (
	"encoding/json"
	"fmt"
	"path"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/store"
	"github.com/coreos/go-semver/semver"
)

const (
	attributesSuffix	= "attributes"
	raftAttributesSuffix	= "raftAttributes"
	storePrefix		= "/0"
)

var (
	membersBucketName		= []byte("members")
	membersRemovedBucketName	= []byte("members_removed")
	clusterBucketName		= []byte("cluster")
	StoreMembersPrefix		= path.Join(storePrefix, "members")
	storeRemovedMembersPrefix	= path.Join(storePrefix, "removed_members")
)

func mustSaveMemberToBackend(be backend.Backend, m *Member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mkey := backendMemberKey(m.ID)
	mvalue, err := json.Marshal(m)
	if err != nil {
		plog.Panicf("marshal raftAttributes should never fail: %v", err)
	}
	tx := be.BatchTx()
	tx.Lock()
	tx.UnsafePut(membersBucketName, mkey, mvalue)
	tx.Unlock()
}
func mustDeleteMemberFromBackend(be backend.Backend, id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mkey := backendMemberKey(id)
	tx := be.BatchTx()
	tx.Lock()
	tx.UnsafeDelete(membersBucketName, mkey)
	tx.UnsafePut(membersRemovedBucketName, mkey, []byte("removed"))
	tx.Unlock()
}
func mustSaveClusterVersionToBackend(be backend.Backend, ver *semver.Version) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ckey := backendClusterVersionKey()
	tx := be.BatchTx()
	tx.Lock()
	defer tx.Unlock()
	tx.UnsafePut(clusterBucketName, ckey, []byte(ver.String()))
}
func mustSaveMemberToStore(s store.Store, m *Member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(m.RaftAttributes)
	if err != nil {
		plog.Panicf("marshal raftAttributes should never fail: %v", err)
	}
	p := path.Join(MemberStoreKey(m.ID), raftAttributesSuffix)
	if _, err := s.Create(p, false, string(b), false, store.TTLOptionSet{ExpireTime: store.Permanent}); err != nil {
		plog.Panicf("create raftAttributes should never fail: %v", err)
	}
}
func mustDeleteMemberFromStore(s store.Store, id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, err := s.Delete(MemberStoreKey(id), true, true); err != nil {
		plog.Panicf("delete member should never fail: %v", err)
	}
	if _, err := s.Create(RemovedMemberStoreKey(id), false, "", false, store.TTLOptionSet{ExpireTime: store.Permanent}); err != nil {
		plog.Panicf("create removedMember should never fail: %v", err)
	}
}
func mustUpdateMemberInStore(s store.Store, m *Member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(m.RaftAttributes)
	if err != nil {
		plog.Panicf("marshal raftAttributes should never fail: %v", err)
	}
	p := path.Join(MemberStoreKey(m.ID), raftAttributesSuffix)
	if _, err := s.Update(p, string(b), store.TTLOptionSet{ExpireTime: store.Permanent}); err != nil {
		plog.Panicf("update raftAttributes should never fail: %v", err)
	}
}
func mustUpdateMemberAttrInStore(s store.Store, m *Member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(m.Attributes)
	if err != nil {
		plog.Panicf("marshal raftAttributes should never fail: %v", err)
	}
	p := path.Join(MemberStoreKey(m.ID), attributesSuffix)
	if _, err := s.Set(p, false, string(b), store.TTLOptionSet{ExpireTime: store.Permanent}); err != nil {
		plog.Panicf("update raftAttributes should never fail: %v", err)
	}
}
func mustSaveClusterVersionToStore(s store.Store, ver *semver.Version) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, err := s.Set(StoreClusterVersionKey(), false, ver.String(), store.TTLOptionSet{ExpireTime: store.Permanent}); err != nil {
		plog.Panicf("save cluster version should never fail: %v", err)
	}
}
func nodeToMember(n *store.NodeExtern) (*Member, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := &Member{ID: MustParseMemberIDFromKey(n.Key)}
	attrs := make(map[string][]byte)
	raftAttrKey := path.Join(n.Key, raftAttributesSuffix)
	attrKey := path.Join(n.Key, attributesSuffix)
	for _, nn := range n.Nodes {
		if nn.Key != raftAttrKey && nn.Key != attrKey {
			return nil, fmt.Errorf("unknown key %q", nn.Key)
		}
		attrs[nn.Key] = []byte(*nn.Value)
	}
	if data := attrs[raftAttrKey]; data != nil {
		if err := json.Unmarshal(data, &m.RaftAttributes); err != nil {
			return nil, fmt.Errorf("unmarshal raftAttributes error: %v", err)
		}
	} else {
		return nil, fmt.Errorf("raftAttributes key doesn't exist")
	}
	if data := attrs[attrKey]; data != nil {
		if err := json.Unmarshal(data, &m.Attributes); err != nil {
			return m, fmt.Errorf("unmarshal attributes error: %v", err)
		}
	}
	return m, nil
}
func backendMemberKey(id types.ID) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []byte(id.String())
}
func backendClusterVersionKey() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []byte("clusterVersion")
}
func mustCreateBackendBuckets(be backend.Backend) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	tx := be.BatchTx()
	tx.Lock()
	defer tx.Unlock()
	tx.UnsafeCreateBucket(membersBucketName)
	tx.UnsafeCreateBucket(membersRemovedBucketName)
	tx.UnsafeCreateBucket(clusterBucketName)
}
func MemberStoreKey(id types.ID) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return path.Join(StoreMembersPrefix, id.String())
}
func StoreClusterVersionKey() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return path.Join(storePrefix, "version")
}
func MemberAttributesStorePath(id types.ID) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return path.Join(MemberStoreKey(id), attributesSuffix)
}
func MustParseMemberIDFromKey(key string) types.ID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	id, err := types.IDFromString(path.Base(key))
	if err != nil {
		plog.Panicf("unexpected parse member id error: %v", err)
	}
	return id
}
func RemovedMemberStoreKey(id types.ID) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return path.Join(storeRemovedMembersPrefix, id.String())
}
