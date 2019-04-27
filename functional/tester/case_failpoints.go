package tester

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"github.com/coreos/etcd/functional/rpcpb"
)

type failpointStats struct {
	mu	sync.Mutex
	crashes	map[string]int
}

var fpStats failpointStats

func failpointFailures(clus *Cluster) (ret []Case, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var fps []string
	fps, err = failpointPaths(clus.Members[0].FailpointHTTPAddr)
	if err != nil {
		return nil, err
	}
	for _, fp := range fps {
		if len(fp) == 0 {
			continue
		}
		fpFails := casesFromFailpoint(fp, clus.Tester.FailpointCommands)
		for i, fpf := range fpFails {
			if strings.Contains(fp, "Snap") {
				fpFails[i] = &caseUntilSnapshot{desc: fpf.Desc(), rpcpbCase: rpcpb.Case_FAILPOINTS, Case: fpf}
			} else {
				fpFails[i] = &caseDelay{Case: fpf, delayDuration: clus.GetCaseDelayDuration()}
			}
		}
		ret = append(ret, fpFails...)
	}
	fpStats.crashes = make(map[string]int)
	return ret, err
}
func failpointPaths(endpoint string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		return nil, rerr
	}
	var fps []string
	for _, l := range strings.Split(string(body), "\n") {
		fp := strings.Split(l, "=")[0]
		fps = append(fps, fp)
	}
	return fps, nil
}
func casesFromFailpoint(fp string, failpointCommands []string) (fs []Case) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	recov := makeRecoverFailpoint(fp)
	for _, fcmd := range failpointCommands {
		inject := makeInjectFailpoint(fp, fcmd)
		fs = append(fs, []Case{&caseFollower{caseByFunc: caseByFunc{desc: fmt.Sprintf("failpoint %q (one: %q)", fp, fcmd), rpcpbCase: rpcpb.Case_FAILPOINTS, injectMember: inject, recoverMember: recov}, last: -1, lead: -1}, &caseLeader{caseByFunc: caseByFunc{desc: fmt.Sprintf("failpoint %q (leader: %q)", fp, fcmd), rpcpbCase: rpcpb.Case_FAILPOINTS, injectMember: inject, recoverMember: recov}, last: -1, lead: -1}, &caseQuorum{caseByFunc: caseByFunc{desc: fmt.Sprintf("failpoint %q (quorum: %q)", fp, fcmd), rpcpbCase: rpcpb.Case_FAILPOINTS, injectMember: inject, recoverMember: recov}, injected: make(map[int]struct{})}, &caseAll{desc: fmt.Sprintf("failpoint %q (all: %q)", fp, fcmd), rpcpbCase: rpcpb.Case_FAILPOINTS, injectMember: inject, recoverMember: recov}}...)
	}
	return fs
}
func makeInjectFailpoint(fp, val string) injectMemberFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(clus *Cluster, idx int) (err error) {
		return putFailpoint(clus.Members[idx].FailpointHTTPAddr, fp, val)
	}
}
func makeRecoverFailpoint(fp string) recoverMemberFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(clus *Cluster, idx int) error {
		if err := delFailpoint(clus.Members[idx].FailpointHTTPAddr, fp); err == nil {
			return nil
		}
		fpStats.mu.Lock()
		fpStats.crashes[fp]++
		fpStats.mu.Unlock()
		return recover_SIGTERM_ETCD(clus, idx)
	}
}
func putFailpoint(ep, fp, val string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	req, _ := http.NewRequest(http.MethodPut, ep+"/"+fp, strings.NewReader(val))
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("failed to PUT %s=%s at %s (%v)", fp, val, ep, resp.Status)
	}
	return nil
}
func delFailpoint(ep, fp string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	req, _ := http.NewRequest(http.MethodDelete, ep+"/"+fp, strings.NewReader(""))
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("failed to DELETE %s at %s (%v)", fp, ep, resp.Status)
	}
	return nil
}
