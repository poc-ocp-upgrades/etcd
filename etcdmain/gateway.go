package etcdmain

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"time"
	"github.com/coreos/etcd/proxy/tcpproxy"
	"github.com/spf13/cobra"
)

var (
	gatewayListenAddr		string
	gatewayEndpoints		[]string
	gatewayDNSCluster		string
	gatewayInsecureDiscovery	bool
	getewayRetryDelay		time.Duration
	gatewayCA			string
)
var (
	rootCmd = &cobra.Command{Use: "etcd", Short: "etcd server", SuggestFor: []string{"etcd"}}
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd.AddCommand(newGatewayCommand())
}
func newGatewayCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	lpc := &cobra.Command{Use: "gateway <subcommand>", Short: "gateway related command"}
	lpc.AddCommand(newGatewayStartCommand())
	return lpc
}
func newGatewayStartCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := cobra.Command{Use: "start", Short: "start the gateway", Run: startGateway}
	cmd.Flags().StringVar(&gatewayListenAddr, "listen-addr", "127.0.0.1:23790", "listen address")
	cmd.Flags().StringVar(&gatewayDNSCluster, "discovery-srv", "", "DNS domain used to bootstrap initial cluster")
	cmd.Flags().BoolVar(&gatewayInsecureDiscovery, "insecure-discovery", false, "accept insecure SRV records")
	cmd.Flags().StringVar(&gatewayCA, "trusted-ca-file", "", "path to the client server TLS CA file.")
	cmd.Flags().StringSliceVar(&gatewayEndpoints, "endpoints", []string{"127.0.0.1:2379"}, "comma separated etcd cluster endpoints")
	cmd.Flags().DurationVar(&getewayRetryDelay, "retry-delay", time.Minute, "duration of delay before retrying failed endpoints")
	return &cmd
}
func stripSchema(eps []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var endpoints []string
	for _, ep := range eps {
		if u, err := url.Parse(ep); err == nil && u.Host != "" {
			ep = u.Host
		}
		endpoints = append(endpoints, ep)
	}
	return endpoints
}
func startGateway(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	srvs := discoverEndpoints(gatewayDNSCluster, gatewayCA, gatewayInsecureDiscovery)
	if len(srvs.Endpoints) == 0 {
		srvs.Endpoints = gatewayEndpoints
	}
	srvs.Endpoints = stripSchema(srvs.Endpoints)
	if len(srvs.SRVs) == 0 {
		for _, ep := range srvs.Endpoints {
			h, p, err := net.SplitHostPort(ep)
			if err != nil {
				plog.Fatalf("error parsing endpoint %q", ep)
			}
			var port uint16
			fmt.Sscanf(p, "%d", &port)
			srvs.SRVs = append(srvs.SRVs, &net.SRV{Target: h, Port: port})
		}
	}
	if len(srvs.Endpoints) == 0 {
		plog.Fatalf("no endpoints found")
	}
	l, err := net.Listen("tcp", gatewayListenAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	tp := tcpproxy.TCPProxy{Listener: l, Endpoints: srvs.SRVs, MonitorInterval: getewayRetryDelay}
	notifySystemd()
	tp.Run()
}
