package cmd

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"strings"
	"github.com/bgentry/speakeasy"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/report"
	"google.golang.org/grpc/grpclog"
)

var (
	dialTotal	int
	leaderEps	[]string
	globalUserName	string
	globalPassword	string
)

func mustFindLeaderEndpoints(c *clientv3.Client) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resp, lerr := c.MemberList(context.TODO())
	if lerr != nil {
		fmt.Fprintf(os.Stderr, "failed to get a member list: %s\n", lerr)
		os.Exit(1)
	}
	leaderId := uint64(0)
	for _, ep := range c.Endpoints() {
		if sresp, serr := c.Status(context.TODO(), ep); serr == nil {
			leaderId = sresp.Leader
			break
		}
	}
	for _, m := range resp.Members {
		if m.ID == leaderId {
			leaderEps = m.ClientURLs
			return
		}
	}
	fmt.Fprintf(os.Stderr, "failed to find a leader endpoint\n")
	os.Exit(1)
}
func getUsernamePassword(usernameFlag string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if globalUserName != "" && globalPassword != "" {
		return globalUserName, globalPassword, nil
	}
	colon := strings.Index(usernameFlag, ":")
	if colon == -1 {
		password, err := speakeasy.Ask("Password: ")
		if err != nil {
			return "", "", err
		}
		globalUserName = usernameFlag
		globalPassword = password
	} else {
		globalUserName = usernameFlag[:colon]
		globalPassword = usernameFlag[colon+1:]
	}
	return globalUserName, globalPassword, nil
}
func mustCreateConn() *clientv3.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	connEndpoints := leaderEps
	if len(connEndpoints) == 0 {
		connEndpoints = []string{endpoints[dialTotal%len(endpoints)]}
		dialTotal++
	}
	cfg := clientv3.Config{Endpoints: connEndpoints, DialTimeout: dialTimeout}
	if !tls.Empty() || tls.TrustedCAFile != "" {
		cfgtls, err := tls.ClientConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad tls config: %v\n", err)
			os.Exit(1)
		}
		cfg.TLS = cfgtls
	}
	if len(user) != 0 {
		username, password, err := getUsernamePassword(user)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad user information: %s %v\n", user, err)
			os.Exit(1)
		}
		cfg.Username = username
		cfg.Password = password
	}
	client, err := clientv3.New(cfg)
	if targetLeader && len(leaderEps) == 0 {
		mustFindLeaderEndpoints(client)
		client.Close()
		return mustCreateConn()
	}
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "dial error: %v\n", err)
		os.Exit(1)
	}
	return client
}
func mustCreateClients(totalClients, totalConns uint) []*clientv3.Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	conns := make([]*clientv3.Client, totalConns)
	for i := range conns {
		conns[i] = mustCreateConn()
	}
	clients := make([]*clientv3.Client, totalClients)
	for i := range clients {
		clients[i] = conns[i%int(totalConns)]
	}
	return clients
}
func mustRandBytes(n int) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rb := make([]byte, n)
	_, err := rand.Read(rb)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate value: %v\n", err)
		os.Exit(1)
	}
	return rb
}
func newReport() report.Report {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := "%4.4f"
	if precise {
		p = "%g"
	}
	if sample {
		return report.NewReportSample(p)
	}
	return report.NewReport(p)
}
func newWeightedReport() report.Report {
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := "%4.4f"
	if precise {
		p = "%g"
	}
	if sample {
		return report.NewReportSample(p)
	}
	return report.NewWeightedReport(report.NewReport(p), p)
}
