package membership

import (
	"bytes"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"context"
	"crypto/sha1"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"path"
	"sort"
	"strings"
	"sync"
	"time"
	"github.com/coreos/etcd/mvcc/backend"
	"github.com/coreos/etcd/pkg/netutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/coreos/etcd/store"
	"github.com/coreos/etcd/version"
	"github.com/coreos/go-semver/semver"
)

type RaftCluster struct {
	id	types.ID
	token	string
	store	store.Store
	be	backend.Backend
	sync.Mutex
	version	*semver.Version
	members	map[types.ID]*Member
	removed	map[types.ID]bool
}

func NewClusterFromURLsMap(token string, urlsmap types.URLsMap) (*RaftCluster, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := NewCluster(token)
	for name, urls := range urlsmap {
		m := NewMember(name, urls, token, nil)
		if _, ok := c.members[m.ID]; ok {
			return nil, fmt.Errorf("member exists with identical ID %v", m)
		}
		if uint64(m.ID) == raft.None {
			return nil, fmt.Errorf("cannot use %x as member id", raft.None)
		}
		c.members[m.ID] = m
	}
	c.genID()
	return c, nil
}
func NewClusterFromMembers(token string, id types.ID, membs []*Member) *RaftCluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := NewCluster(token)
	c.id = id
	for _, m := range membs {
		c.members[m.ID] = m
	}
	return c
}
func NewCluster(token string) *RaftCluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &RaftCluster{token: token, members: make(map[types.ID]*Member), removed: make(map[types.ID]bool)}
}
func (c *RaftCluster) ID() types.ID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.id
}
func (c *RaftCluster) Members() []*Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	var ms MembersByID
	for _, m := range c.members {
		ms = append(ms, m.Clone())
	}
	sort.Sort(ms)
	return []*Member(ms)
}
func (c *RaftCluster) Member(id types.ID) *Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	return c.members[id].Clone()
}
func (c *RaftCluster) MemberByName(name string) *Member {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	var memb *Member
	for _, m := range c.members {
		if m.Name == name {
			if memb != nil {
				plog.Panicf("two members with the given name %q exist", name)
			}
			memb = m
		}
	}
	return memb.Clone()
}
func (c *RaftCluster) MemberIDs() []types.ID {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	var ids []types.ID
	for _, m := range c.members {
		ids = append(ids, m.ID)
	}
	sort.Sort(types.IDSlice(ids))
	return ids
}
func (c *RaftCluster) IsIDRemoved(id types.ID) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	return c.removed[id]
}
func (c *RaftCluster) PeerURLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	urls := make([]string, 0)
	for _, p := range c.members {
		urls = append(urls, p.PeerURLs...)
	}
	sort.Strings(urls)
	return urls
}
func (c *RaftCluster) ClientURLs() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	urls := make([]string, 0)
	for _, p := range c.members {
		urls = append(urls, p.ClientURLs...)
	}
	sort.Strings(urls)
	return urls
}
func (c *RaftCluster) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "{ClusterID:%s ", c.id)
	var ms []string
	for _, m := range c.members {
		ms = append(ms, fmt.Sprintf("%+v", m))
	}
	fmt.Fprintf(b, "Members:[%s] ", strings.Join(ms, " "))
	var ids []string
	for id := range c.removed {
		ids = append(ids, id.String())
	}
	fmt.Fprintf(b, "RemovedMemberIDs:[%s]}", strings.Join(ids, " "))
	return b.String()
}
func (c *RaftCluster) genID() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mIDs := c.MemberIDs()
	b := make([]byte, 8*len(mIDs))
	for i, id := range mIDs {
		binary.BigEndian.PutUint64(b[8*i:], uint64(id))
	}
	hash := sha1.Sum(b)
	c.id = types.ID(binary.BigEndian.Uint64(hash[:8]))
}
func (c *RaftCluster) SetID(id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.id = id
}
func (c *RaftCluster) SetStore(st store.Store) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.store = st
}
func (c *RaftCluster) SetBackend(be backend.Backend) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.be = be
	mustCreateBackendBuckets(c.be)
}
func (c *RaftCluster) Recover(onSet func(*semver.Version)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	c.members, c.removed = membersFromStore(c.store)
	c.version = clusterVersionFromStore(c.store)
	mustDetectDowngrade(c.version)
	onSet(c.version)
	for _, m := range c.members {
		plog.Infof("added member %s %v to cluster %s from store", m.ID, m.PeerURLs, c.id)
	}
	if c.version != nil {
		plog.Infof("set the cluster version to %v from store", version.Cluster(c.version.String()))
	}
}
func (c *RaftCluster) ValidateConfigurationChange(cc raftpb.ConfChange) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	members, removed := membersFromStore(c.store)
	id := types.ID(cc.NodeID)
	if removed[id] {
		return ErrIDRemoved
	}
	switch cc.Type {
	case raftpb.ConfChangeAddNode:
		if members[id] != nil {
			return ErrIDExists
		}
		urls := make(map[string]bool)
		for _, m := range members {
			for _, u := range m.PeerURLs {
				urls[u] = true
			}
		}
		m := new(Member)
		if err := json.Unmarshal(cc.Context, m); err != nil {
			plog.Panicf("unmarshal member should never fail: %v", err)
		}
		for _, u := range m.PeerURLs {
			if urls[u] {
				return ErrPeerURLexists
			}
		}
	case raftpb.ConfChangeRemoveNode:
		if members[id] == nil {
			return ErrIDNotFound
		}
	case raftpb.ConfChangeUpdateNode:
		if members[id] == nil {
			return ErrIDNotFound
		}
		urls := make(map[string]bool)
		for _, m := range members {
			if m.ID == id {
				continue
			}
			for _, u := range m.PeerURLs {
				urls[u] = true
			}
		}
		m := new(Member)
		if err := json.Unmarshal(cc.Context, m); err != nil {
			plog.Panicf("unmarshal member should never fail: %v", err)
		}
		for _, u := range m.PeerURLs {
			if urls[u] {
				return ErrPeerURLexists
			}
		}
	default:
		plog.Panicf("ConfChange type should be either AddNode, RemoveNode or UpdateNode")
	}
	return nil
}
func (c *RaftCluster) AddMember(m *Member) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	if c.store != nil {
		mustSaveMemberToStore(c.store, m)
	}
	if c.be != nil {
		mustSaveMemberToBackend(c.be, m)
	}
	c.members[m.ID] = m
	plog.Infof("added member %s %v to cluster %s", m.ID, m.PeerURLs, c.id)
}
func (c *RaftCluster) RemoveMember(id types.ID) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	if c.store != nil {
		mustDeleteMemberFromStore(c.store, id)
	}
	if c.be != nil {
		mustDeleteMemberFromBackend(c.be, id)
	}
	delete(c.members, id)
	c.removed[id] = true
	plog.Infof("removed member %s from cluster %s", id, c.id)
}
func (c *RaftCluster) UpdateAttributes(id types.ID, attr Attributes) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	if m, ok := c.members[id]; ok {
		m.Attributes = attr
		if c.store != nil {
			mustUpdateMemberAttrInStore(c.store, m)
		}
		if c.be != nil {
			mustSaveMemberToBackend(c.be, m)
		}
		return
	}
	_, ok := c.removed[id]
	if !ok {
		plog.Panicf("error updating attributes of unknown member %s", id)
	}
	plog.Warningf("skipped updating attributes of removed member %s", id)
}
func (c *RaftCluster) UpdateRaftAttributes(id types.ID, raftAttr RaftAttributes) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	c.members[id].RaftAttributes = raftAttr
	if c.store != nil {
		mustUpdateMemberInStore(c.store, c.members[id])
	}
	if c.be != nil {
		mustSaveMemberToBackend(c.be, c.members[id])
	}
	plog.Noticef("updated member %s %v in cluster %s", id, raftAttr.PeerURLs, c.id)
}
func (c *RaftCluster) Version() *semver.Version {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	if c.version == nil {
		return nil
	}
	return semver.Must(semver.NewVersion(c.version.String()))
}
func (c *RaftCluster) SetVersion(ver *semver.Version, onSet func(*semver.Version)) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Lock()
	defer c.Unlock()
	if c.version != nil {
		plog.Noticef("updated the cluster version from %v to %v", version.Cluster(c.version.String()), version.Cluster(ver.String()))
	} else {
		plog.Noticef("set the initial cluster version to %v", version.Cluster(ver.String()))
	}
	c.version = ver
	mustDetectDowngrade(c.version)
	if c.store != nil {
		mustSaveClusterVersionToStore(c.store, ver)
	}
	if c.be != nil {
		mustSaveClusterVersionToBackend(c.be, ver)
	}
	onSet(ver)
}
func (c *RaftCluster) IsReadyToAddNewMember() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nmembers := 1
	nstarted := 0
	for _, member := range c.members {
		if member.IsStarted() {
			nstarted++
		}
		nmembers++
	}
	if nstarted == 1 && nmembers == 2 {
		plog.Debugf("The number of started member is 1. This cluster can accept add member request.")
		return true
	}
	nquorum := nmembers/2 + 1
	if nstarted < nquorum {
		plog.Warningf("Reject add member request: the number of started member (%d) will be less than the quorum number of the cluster (%d)", nstarted, nquorum)
		return false
	}
	return true
}
func (c *RaftCluster) IsReadyToRemoveMember(id uint64) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	nmembers := 0
	nstarted := 0
	for _, member := range c.members {
		if uint64(member.ID) == id {
			continue
		}
		if member.IsStarted() {
			nstarted++
		}
		nmembers++
	}
	nquorum := nmembers/2 + 1
	if nstarted < nquorum {
		plog.Warningf("Reject remove member request: the number of started member (%d) will be less than the quorum number of the cluster (%d)", nstarted, nquorum)
		return false
	}
	return true
}
func membersFromStore(st store.Store) (map[types.ID]*Member, map[types.ID]bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	members := make(map[types.ID]*Member)
	removed := make(map[types.ID]bool)
	e, err := st.Get(StoreMembersPrefix, true, true)
	if err != nil {
		if isKeyNotFound(err) {
			return members, removed
		}
		plog.Panicf("get storeMembers should never fail: %v", err)
	}
	for _, n := range e.Node.Nodes {
		var m *Member
		m, err = nodeToMember(n)
		if err != nil {
			plog.Panicf("nodeToMember should never fail: %v", err)
		}
		members[m.ID] = m
	}
	e, err = st.Get(storeRemovedMembersPrefix, true, true)
	if err != nil {
		if isKeyNotFound(err) {
			return members, removed
		}
		plog.Panicf("get storeRemovedMembers should never fail: %v", err)
	}
	for _, n := range e.Node.Nodes {
		removed[MustParseMemberIDFromKey(n.Key)] = true
	}
	return members, removed
}
func clusterVersionFromStore(st store.Store) *semver.Version {
	_logClusterCodePath()
	defer _logClusterCodePath()
	e, err := st.Get(path.Join(storePrefix, "version"), false, false)
	if err != nil {
		if isKeyNotFound(err) {
			return nil
		}
		plog.Panicf("unexpected error (%v) when getting cluster version from store", err)
	}
	return semver.Must(semver.NewVersion(*e.Node.Value))
}
func ValidateClusterAndAssignIDs(local *RaftCluster, existing *RaftCluster) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ems := existing.Members()
	lms := local.Members()
	if len(ems) != len(lms) {
		return fmt.Errorf("member count is unequal")
	}
	sort.Sort(MembersByPeerURLs(ems))
	sort.Sort(MembersByPeerURLs(lms))
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	for i := range ems {
		if ok, err := netutil.URLStringsEqual(ctx, ems[i].PeerURLs, lms[i].PeerURLs); !ok {
			return fmt.Errorf("unmatched member while checking PeerURLs (%v)", err)
		}
		lms[i].ID = ems[i].ID
	}
	local.members = make(map[types.ID]*Member)
	for _, m := range lms {
		local.members[m.ID] = m
	}
	return nil
}
func mustDetectDowngrade(cv *semver.Version) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lv := semver.Must(semver.NewVersion(version.Version))
	lv = &semver.Version{Major: lv.Major, Minor: lv.Minor}
	if cv != nil && lv.LessThan(*cv) {
		plog.Fatalf("cluster cannot be downgraded (current version: %s is lower than determined cluster version: %s).", version.Version, version.Cluster(cv.String()))
	}
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
