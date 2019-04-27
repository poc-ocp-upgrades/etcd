package etcdserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
	"github.com/coreos/etcd/etcdserver/membership"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/version"
	"github.com/coreos/go-semver/semver"
)

func isMemberBootstrapped(cl *membership.RaftCluster, member string, rt http.RoundTripper, timeout time.Duration) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rcl, err := getClusterFromRemotePeers(getRemotePeerURLs(cl, member), timeout, false, rt)
	if err != nil {
		return false
	}
	id := cl.MemberByName(member).ID
	m := rcl.Member(id)
	if m == nil {
		return false
	}
	if len(m.ClientURLs) > 0 {
		return true
	}
	return false
}
func GetClusterFromRemotePeers(urls []string, rt http.RoundTripper) (*membership.RaftCluster, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return getClusterFromRemotePeers(urls, 10*time.Second, true, rt)
}
func getClusterFromRemotePeers(urls []string, timeout time.Duration, logerr bool, rt http.RoundTripper) (*membership.RaftCluster, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := &http.Client{Transport: rt, Timeout: timeout}
	for _, u := range urls {
		resp, err := cc.Get(u + "/members")
		if err != nil {
			if logerr {
				plog.Warningf("could not get cluster response from %s: %v", u, err)
			}
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			if logerr {
				plog.Warningf("could not read the body of cluster response: %v", err)
			}
			continue
		}
		var membs []*membership.Member
		if err = json.Unmarshal(b, &membs); err != nil {
			if logerr {
				plog.Warningf("could not unmarshal cluster response: %v", err)
			}
			continue
		}
		id, err := types.IDFromString(resp.Header.Get("X-Etcd-Cluster-ID"))
		if err != nil {
			if logerr {
				plog.Warningf("could not parse the cluster ID from cluster res: %v", err)
			}
			continue
		}
		if len(membs) > 0 {
			return membership.NewClusterFromMembers("", id, membs), nil
		}
		return nil, fmt.Errorf("failed to get raft cluster member(s) from the given urls.")
	}
	return nil, fmt.Errorf("could not retrieve cluster information from the given urls")
}
func getRemotePeerURLs(cl *membership.RaftCluster, local string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	us := make([]string, 0)
	for _, m := range cl.Members() {
		if m.Name == local {
			continue
		}
		us = append(us, m.PeerURLs...)
	}
	sort.Strings(us)
	return us
}
func getVersions(cl *membership.RaftCluster, local types.ID, rt http.RoundTripper) map[string]*version.Versions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	members := cl.Members()
	vers := make(map[string]*version.Versions)
	for _, m := range members {
		if m.ID == local {
			cv := "not_decided"
			if cl.Version() != nil {
				cv = cl.Version().String()
			}
			vers[m.ID.String()] = &version.Versions{Server: version.Version, Cluster: cv}
			continue
		}
		ver, err := getVersion(m, rt)
		if err != nil {
			plog.Warningf("cannot get the version of member %s (%v)", m.ID, err)
			vers[m.ID.String()] = nil
		} else {
			vers[m.ID.String()] = ver
		}
	}
	return vers
}
func decideClusterVersion(vers map[string]*version.Versions) *semver.Version {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cv *semver.Version
	lv := semver.Must(semver.NewVersion(version.Version))
	for mid, ver := range vers {
		if ver == nil {
			return nil
		}
		v, err := semver.NewVersion(ver.Server)
		if err != nil {
			plog.Errorf("cannot understand the version of member %s (%v)", mid, err)
			return nil
		}
		if lv.LessThan(*v) {
			plog.Warningf("the local etcd version %s is not up-to-date", lv.String())
			plog.Warningf("member %s has a higher version %s", mid, ver.Server)
		}
		if cv == nil {
			cv = v
		} else if v.LessThan(*cv) {
			cv = v
		}
	}
	return cv
}
func isCompatibleWithCluster(cl *membership.RaftCluster, local types.ID, rt http.RoundTripper) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	vers := getVersions(cl, local, rt)
	minV := semver.Must(semver.NewVersion(version.MinClusterVersion))
	maxV := semver.Must(semver.NewVersion(version.Version))
	maxV = &semver.Version{Major: maxV.Major, Minor: maxV.Minor}
	return isCompatibleWithVers(vers, local, minV, maxV)
}
func isCompatibleWithVers(vers map[string]*version.Versions, local types.ID, minV, maxV *semver.Version) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ok bool
	for id, v := range vers {
		if id == local.String() {
			continue
		}
		if v == nil {
			continue
		}
		clusterv, err := semver.NewVersion(v.Cluster)
		if err != nil {
			plog.Errorf("cannot understand the cluster version of member %s (%v)", id, err)
			continue
		}
		if clusterv.LessThan(*minV) {
			plog.Warningf("the running cluster version(%v) is lower than the minimal cluster version(%v) supported", clusterv.String(), minV.String())
			return false
		}
		if maxV.LessThan(*clusterv) {
			plog.Warningf("the running cluster version(%v) is higher than the maximum cluster version(%v) supported", clusterv.String(), maxV.String())
			return false
		}
		ok = true
	}
	return ok
}
func getVersion(m *membership.Member, rt http.RoundTripper) (*version.Versions, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := &http.Client{Transport: rt}
	var (
		err	error
		resp	*http.Response
	)
	for _, u := range m.PeerURLs {
		resp, err = cc.Get(u + "/version")
		if err != nil {
			plog.Warningf("failed to reach the peerURL(%s) of member %s (%v)", u, m.ID, err)
			continue
		}
		var b []byte
		b, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			plog.Warningf("failed to read out the response body from the peerURL(%s) of member %s (%v)", u, m.ID, err)
			continue
		}
		var vers version.Versions
		if err = json.Unmarshal(b, &vers); err != nil {
			plog.Warningf("failed to unmarshal the response body got from the peerURL(%s) of member %s (%v)", u, m.ID, err)
			continue
		}
		return &vers, nil
	}
	return nil, err
}
