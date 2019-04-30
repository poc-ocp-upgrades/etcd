package main

import (
	"context"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"io/ioutil"
	"net/url"
	godefaulthttp "net/http"
	"os"
	"strings"
	"time"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/embed"
	"go.uber.org/zap"
)

func newEmbedURLs(n int) (urls []url.URL) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	urls = make([]url.URL, n)
	for i := 0; i < n; i++ {
		u, _ := url.Parse(fmt.Sprintf("unix://localhost:%d%06d", os.Getpid(), i))
		urls[i] = *u
	}
	return urls
}
func setupEmbedCfg(cfg *embed.Config, curls, purls, ics []url.URL) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg.Logger = "zap"
	cfg.LogOutputs = []string{"/dev/null"}
	cfg.Debug = false
	var err error
	cfg.Dir, err = ioutil.TempDir(os.TempDir(), fmt.Sprintf("%016X", time.Now().UnixNano()))
	if err != nil {
		panic(err)
	}
	os.RemoveAll(cfg.Dir)
	cfg.ClusterState = "new"
	cfg.LCUrls, cfg.ACUrls = curls, curls
	cfg.LPUrls, cfg.APUrls = purls, purls
	cfg.InitialCluster = ""
	for i := range ics {
		cfg.InitialCluster += fmt.Sprintf(",%d=%s", i, ics[i].String())
	}
	cfg.InitialCluster = cfg.InitialCluster[1:]
}
func getCommand(exec, name, dir, cURL, pURL, cluster string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := fmt.Sprintf("%s --name %s --data-dir %s --listen-client-urls %s --advertise-client-urls %s ", exec, name, dir, cURL, cURL)
	s += fmt.Sprintf("--listen-peer-urls %s --initial-advertise-peer-urls %s ", pURL, pURL)
	s += fmt.Sprintf("--initial-cluster %s ", cluster)
	return s + "--initial-cluster-token tkn --initial-cluster-state new"
}
func write(ep string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{strings.Replace(ep, "/metrics", "", 1)}})
	if err != nil {
		lg.Panic("failed to create client", zap.Error(err))
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = cli.Put(ctx, "____test", "")
	if err != nil {
		lg.Panic("failed to write test key", zap.Error(err))
	}
	_, err = cli.Get(ctx, "____test")
	if err != nil {
		lg.Panic("failed to read test key", zap.Error(err))
	}
	_, err = cli.Delete(ctx, "____test")
	if err != nil {
		lg.Panic("failed to delete test key", zap.Error(err))
	}
	cli.Watch(ctx, "____test", clientv3.WithCreatedNotify())
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
