package etcdmain

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"time"
	"go.etcd.io/etcd/proxy/tcpproxy"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	gatewayListenAddr		string
	gatewayEndpoints		[]string
	gatewayDNSCluster		string
	gatewayDNSClusterServiceName	string
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
	rootCmd.AddCommand(newGatewayCommand())
}
func newGatewayCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lpc := &cobra.Command{Use: "gateway <subcommand>", Short: "gateway related command"}
	lpc.AddCommand(newGatewayStartCommand())
	return lpc
}
func newGatewayStartCommand() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := cobra.Command{Use: "start", Short: "start the gateway", Run: startGateway}
	cmd.Flags().StringVar(&gatewayListenAddr, "listen-addr", "127.0.0.1:23790", "listen address")
	cmd.Flags().StringVar(&gatewayDNSCluster, "discovery-srv", "", "DNS domain used to bootstrap initial cluster")
	cmd.Flags().StringVar(&gatewayDNSClusterServiceName, "discovery-srv-name", "", "service name to query when using DNS discovery")
	cmd.Flags().BoolVar(&gatewayInsecureDiscovery, "insecure-discovery", false, "accept insecure SRV records")
	cmd.Flags().StringVar(&gatewayCA, "trusted-ca-file", "", "path to the client server TLS CA file.")
	cmd.Flags().StringSliceVar(&gatewayEndpoints, "endpoints", []string{"127.0.0.1:2379"}, "comma separated etcd cluster endpoints")
	cmd.Flags().DurationVar(&getewayRetryDelay, "retry-delay", time.Minute, "duration of delay before retrying failed endpoints")
	return &cmd
}
func stripSchema(eps []string) []string {
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
	var lg *zap.Logger
	lg, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	srvs := discoverEndpoints(lg, gatewayDNSCluster, gatewayCA, gatewayInsecureDiscovery, gatewayDNSClusterServiceName)
	if len(srvs.Endpoints) == 0 {
		srvs.Endpoints = gatewayEndpoints
	}
	srvs.Endpoints = stripSchema(srvs.Endpoints)
	if len(srvs.SRVs) == 0 {
		for _, ep := range srvs.Endpoints {
			h, p, serr := net.SplitHostPort(ep)
			if serr != nil {
				fmt.Printf("error parsing endpoint %q", ep)
				os.Exit(1)
			}
			var port uint16
			fmt.Sscanf(p, "%d", &port)
			srvs.SRVs = append(srvs.SRVs, &net.SRV{Target: h, Port: port})
		}
	}
	if len(srvs.Endpoints) == 0 {
		fmt.Println("no endpoints found")
		os.Exit(1)
	}
	var l net.Listener
	l, err = net.Listen("tcp", gatewayListenAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	tp := tcpproxy.TCPProxy{Logger: lg, Listener: l, Endpoints: srvs.SRVs, MonitorInterval: getewayRetryDelay}
	notifySystemd(lg)
	tp.Run()
}
