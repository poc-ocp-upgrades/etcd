package etcdserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
	"go.etcd.io/etcd/etcdserver/api/membership"
	"go.etcd.io/etcd/pkg/types"
	"go.etcd.io/etcd/version"
	"github.com/coreos/go-semver/semver"
	"go.uber.org/zap"
)

func isMemberBootstrapped(lg *zap.Logger, cl *membership.RaftCluster, member string, rt http.RoundTripper, timeout time.Duration) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rcl, err := getClusterFromRemotePeers(lg, getRemotePeerURLs(cl, member), timeout, false, rt)
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
func GetClusterFromRemotePeers(lg *zap.Logger, urls []string, rt http.RoundTripper) (*membership.RaftCluster, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return getClusterFromRemotePeers(lg, urls, 10*time.Second, true, rt)
}
func getClusterFromRemotePeers(lg *zap.Logger, urls []string, timeout time.Duration, logerr bool, rt http.RoundTripper) (*membership.RaftCluster, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := &http.Client{Transport: rt, Timeout: timeout}
	for _, u := range urls {
		addr := u + "/members"
		resp, err := cc.Get(addr)
		if err != nil {
			if logerr {
				if lg != nil {
					lg.Warn("failed to get cluster response", zap.String("address", addr), zap.Error(err))
				} else {
					plog.Warningf("could not get cluster response from %s: %v", u, err)
				}
			}
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			if logerr {
				if lg != nil {
					lg.Warn("failed to read body of cluster response", zap.String("address", addr), zap.Error(err))
				} else {
					plog.Warningf("could not read the body of cluster response: %v", err)
				}
			}
			continue
		}
		var membs []*membership.Member
		if err = json.Unmarshal(b, &membs); err != nil {
			if logerr {
				if lg != nil {
					lg.Warn("failed to unmarshal cluster response", zap.String("address", addr), zap.Error(err))
				} else {
					plog.Warningf("could not unmarshal cluster response: %v", err)
				}
			}
			continue
		}
		id, err := types.IDFromString(resp.Header.Get("X-Etcd-Cluster-ID"))
		if err != nil {
			if logerr {
				if lg != nil {
					lg.Warn("failed to parse cluster ID", zap.String("address", addr), zap.String("header", resp.Header.Get("X-Etcd-Cluster-ID")), zap.Error(err))
				} else {
					plog.Warningf("could not parse the cluster ID from cluster res: %v", err)
				}
			}
			continue
		}
		if len(membs) > 0 {
			return membership.NewClusterFromMembers(lg, "", id, membs), nil
		}
		return nil, fmt.Errorf("failed to get raft cluster member(s) from the given URLs")
	}
	return nil, fmt.Errorf("could not retrieve cluster information from the given URLs")
}
func getRemotePeerURLs(cl *membership.RaftCluster, local string) []string {
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
func getVersions(lg *zap.Logger, cl *membership.RaftCluster, local types.ID, rt http.RoundTripper) map[string]*version.Versions {
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
		ver, err := getVersion(lg, m, rt)
		if err != nil {
			if lg != nil {
				lg.Warn("failed to get version", zap.String("remote-member-id", m.ID.String()), zap.Error(err))
			} else {
				plog.Warningf("cannot get the version of member %s (%v)", m.ID, err)
			}
			vers[m.ID.String()] = nil
		} else {
			vers[m.ID.String()] = ver
		}
	}
	return vers
}
func decideClusterVersion(lg *zap.Logger, vers map[string]*version.Versions) *semver.Version {
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
			if lg != nil {
				lg.Warn("failed to parse server version of remote member", zap.String("remote-member-id", mid), zap.String("remote-member-version", ver.Server), zap.Error(err))
			} else {
				plog.Errorf("cannot understand the version of member %s (%v)", mid, err)
			}
			return nil
		}
		if lv.LessThan(*v) {
			if lg != nil {
				lg.Warn("leader found higher-versioned member", zap.String("local-member-version", lv.String()), zap.String("remote-member-id", mid), zap.String("remote-member-version", ver.Server))
			} else {
				plog.Warningf("the local etcd version %s is not up-to-date", lv.String())
				plog.Warningf("member %s has a higher version %s", mid, ver.Server)
			}
		}
		if cv == nil {
			cv = v
		} else if v.LessThan(*cv) {
			cv = v
		}
	}
	return cv
}
func isCompatibleWithCluster(lg *zap.Logger, cl *membership.RaftCluster, local types.ID, rt http.RoundTripper) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	vers := getVersions(lg, cl, local, rt)
	minV := semver.Must(semver.NewVersion(version.MinClusterVersion))
	maxV := semver.Must(semver.NewVersion(version.Version))
	maxV = &semver.Version{Major: maxV.Major, Minor: maxV.Minor}
	return isCompatibleWithVers(lg, vers, local, minV, maxV)
}
func isCompatibleWithVers(lg *zap.Logger, vers map[string]*version.Versions, local types.ID, minV, maxV *semver.Version) bool {
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
			if lg != nil {
				lg.Warn("failed to parse cluster version of remote member", zap.String("remote-member-id", id), zap.String("remote-member-cluster-version", v.Cluster), zap.Error(err))
			} else {
				plog.Errorf("cannot understand the cluster version of member %s (%v)", id, err)
			}
			continue
		}
		if clusterv.LessThan(*minV) {
			if lg != nil {
				lg.Warn("cluster version of remote member is not compatible; too low", zap.String("remote-member-id", id), zap.String("remote-member-cluster-version", clusterv.String()), zap.String("minimum-cluster-version-supported", minV.String()))
			} else {
				plog.Warningf("the running cluster version(%v) is lower than the minimal cluster version(%v) supported", clusterv.String(), minV.String())
			}
			return false
		}
		if maxV.LessThan(*clusterv) {
			if lg != nil {
				lg.Warn("cluster version of remote member is not compatible; too high", zap.String("remote-member-id", id), zap.String("remote-member-cluster-version", clusterv.String()), zap.String("minimum-cluster-version-supported", minV.String()))
			} else {
				plog.Warningf("the running cluster version(%v) is higher than the maximum cluster version(%v) supported", clusterv.String(), maxV.String())
			}
			return false
		}
		ok = true
	}
	return ok
}
func getVersion(lg *zap.Logger, m *membership.Member, rt http.RoundTripper) (*version.Versions, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cc := &http.Client{Transport: rt}
	var (
		err	error
		resp	*http.Response
	)
	for _, u := range m.PeerURLs {
		addr := u + "/version"
		resp, err = cc.Get(addr)
		if err != nil {
			if lg != nil {
				lg.Warn("failed to reach the peer URL", zap.String("address", addr), zap.String("remote-member-id", m.ID.String()), zap.Error(err))
			} else {
				plog.Warningf("failed to reach the peerURL(%s) of member %s (%v)", u, m.ID, err)
			}
			continue
		}
		var b []byte
		b, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			if lg != nil {
				lg.Warn("failed to read body of response", zap.String("address", addr), zap.String("remote-member-id", m.ID.String()), zap.Error(err))
			} else {
				plog.Warningf("failed to read out the response body from the peerURL(%s) of member %s (%v)", u, m.ID, err)
			}
			continue
		}
		var vers version.Versions
		if err = json.Unmarshal(b, &vers); err != nil {
			if lg != nil {
				lg.Warn("failed to unmarshal response", zap.String("address", addr), zap.String("remote-member-id", m.ID.String()), zap.Error(err))
			} else {
				plog.Warningf("failed to unmarshal the response body got from the peerURL(%s) of member %s (%v)", u, m.ID, err)
			}
			continue
		}
		return &vers, nil
	}
	return nil, err
}
