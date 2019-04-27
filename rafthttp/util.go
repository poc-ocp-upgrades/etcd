package rafthttp

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/version"
	"github.com/coreos/go-semver/semver"
)

var (
	errMemberRemoved	= fmt.Errorf("the member has been permanently removed from the cluster")
	errMemberNotFound	= fmt.Errorf("member not found")
)

func NewListener(u url.URL, tlsinfo *transport.TLSInfo) (net.Listener, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return transport.NewTimeoutListener(u.Host, u.Scheme, tlsinfo, ConnReadTimeout, ConnWriteTimeout)
}
func NewRoundTripper(tlsInfo transport.TLSInfo, dialTimeout time.Duration) (http.RoundTripper, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return transport.NewTimeoutTransport(tlsInfo, dialTimeout, 0, 0)
}
func newStreamRoundTripper(tlsInfo transport.TLSInfo, dialTimeout time.Duration) (http.RoundTripper, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return transport.NewTimeoutTransport(tlsInfo, dialTimeout, ConnReadTimeout, ConnWriteTimeout)
}
func createPostRequest(u url.URL, path string, body io.Reader, ct string, urls types.URLs, from, cid types.ID) *http.Request {
	_logClusterCodePath()
	defer _logClusterCodePath()
	uu := u
	uu.Path = path
	req, err := http.NewRequest("POST", uu.String(), body)
	if err != nil {
		plog.Panicf("unexpected new request error (%v)", err)
	}
	req.Header.Set("Content-Type", ct)
	req.Header.Set("X-Server-From", from.String())
	req.Header.Set("X-Server-Version", version.Version)
	req.Header.Set("X-Min-Cluster-Version", version.MinClusterVersion)
	req.Header.Set("X-Etcd-Cluster-ID", cid.String())
	setPeerURLsHeader(req, urls)
	return req
}
func checkPostResponse(resp *http.Response, body []byte, req *http.Request, to types.ID) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch resp.StatusCode {
	case http.StatusPreconditionFailed:
		switch strings.TrimSuffix(string(body), "\n") {
		case errIncompatibleVersion.Error():
			plog.Errorf("request sent was ignored by peer %s (server version incompatible)", to)
			return errIncompatibleVersion
		case errClusterIDMismatch.Error():
			plog.Errorf("request sent was ignored (cluster ID mismatch: remote[%s]=%s, local=%s)", to, resp.Header.Get("X-Etcd-Cluster-ID"), req.Header.Get("X-Etcd-Cluster-ID"))
			return errClusterIDMismatch
		default:
			return fmt.Errorf("unhandled error %q when precondition failed", string(body))
		}
	case http.StatusForbidden:
		return errMemberRemoved
	case http.StatusNoContent:
		return nil
	default:
		return fmt.Errorf("unexpected http status %s while posting to %q", http.StatusText(resp.StatusCode), req.URL.String())
	}
}
func reportCriticalError(err error, errc chan<- error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	select {
	case errc <- err:
	default:
	}
}
func compareMajorMinorVersion(a, b *semver.Version) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	na := &semver.Version{Major: a.Major, Minor: a.Minor}
	nb := &semver.Version{Major: b.Major, Minor: b.Minor}
	switch {
	case na.LessThan(*nb):
		return -1
	case nb.LessThan(*na):
		return 1
	default:
		return 0
	}
}
func serverVersion(h http.Header) *semver.Version {
	_logClusterCodePath()
	defer _logClusterCodePath()
	verStr := h.Get("X-Server-Version")
	if verStr == "" {
		verStr = "2.0.0"
	}
	return semver.Must(semver.NewVersion(verStr))
}
func minClusterVersion(h http.Header) *semver.Version {
	_logClusterCodePath()
	defer _logClusterCodePath()
	verStr := h.Get("X-Min-Cluster-Version")
	if verStr == "" {
		verStr = "2.0.0"
	}
	return semver.Must(semver.NewVersion(verStr))
}
func checkVersionCompability(name string, server, minCluster *semver.Version) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localServer := semver.Must(semver.NewVersion(version.Version))
	localMinCluster := semver.Must(semver.NewVersion(version.MinClusterVersion))
	if compareMajorMinorVersion(server, localMinCluster) == -1 {
		return fmt.Errorf("remote version is too low: remote[%s]=%s, local=%s", name, server, localServer)
	}
	if compareMajorMinorVersion(minCluster, localServer) == 1 {
		return fmt.Errorf("local version is too low: remote[%s]=%s, local=%s", name, server, localServer)
	}
	return nil
}
func setPeerURLsHeader(req *http.Request, urls types.URLs) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if urls == nil {
		return
	}
	peerURLs := make([]string, urls.Len())
	for i := range urls {
		peerURLs[i] = urls[i].String()
	}
	req.Header.Set("X-PeerURLs", strings.Join(peerURLs, ","))
}
func addRemoteFromRequest(tr Transporter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if from, err := types.IDFromString(r.Header.Get("X-Server-From")); err == nil {
		if urls := r.Header.Get("X-PeerURLs"); urls != "" {
			tr.AddRemote(from, strings.Split(urls, ","))
		}
	}
}
