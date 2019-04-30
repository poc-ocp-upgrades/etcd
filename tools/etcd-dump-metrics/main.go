package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"go.etcd.io/etcd/embed"
	"go.uber.org/zap"
)

var lg *zap.Logger

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	lg, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	addr := flag.String("addr", "", "etcd metrics URL to fetch from (empty to use current git branch)")
	downloadVer := flag.String("download-ver", "", "etcd binary version to download and fetch metrics from")
	debug := flag.Bool("debug", false, "true to enable debug logging")
	flag.Parse()
	if *addr != "" && *downloadVer != "" {
		panic("specify either 'addr' or 'download-ver'")
	}
	if *debug {
		lg = zap.NewExample()
	}
	ep := *addr
	if ep == "" {
		if *downloadVer != "" {
			ver := *downloadVer
			d, err := ioutil.TempDir(os.TempDir(), ver)
			if err != nil {
				panic(err)
			}
			defer os.RemoveAll(d)
			var bp string
			bp, err = install(ver, d)
			if err != nil {
				panic(err)
			}
			ep = "http://localhost:2379/metrics"
			cluster := "s1=http://localhost:2380,s2=http://localhost:22380"
			d1 := filepath.Join(d, "s1")
			d2 := filepath.Join(d, "s2")
			os.RemoveAll(d1)
			os.RemoveAll(d2)
			type run struct {
				err	error
				cmd	*exec.Cmd
			}
			rc := make(chan run)
			cs1 := getCommand(bp, "s1", d1, "http://localhost:2379", "http://localhost:2380", cluster)
			cmd1 := exec.Command("bash", "-c", cs1)
			go func() {
				if *debug {
					cmd1.Stderr = os.Stderr
				}
				if cerr := cmd1.Start(); cerr != nil {
					lg.Warn("failed to start first process", zap.Error(cerr))
					rc <- run{err: cerr}
					return
				}
				lg.Debug("started first process")
				rc <- run{cmd: cmd1}
			}()
			cs2 := getCommand(bp, "s2", d2, "http://localhost:22379", "http://localhost:22380", cluster)
			cmd2 := exec.Command("bash", "-c", cs2)
			go func() {
				if *debug {
					cmd2.Stderr = os.Stderr
				}
				if cerr := cmd2.Start(); cerr != nil {
					lg.Warn("failed to start second process", zap.Error(cerr))
					rc <- run{err: cerr}
					return
				}
				lg.Debug("started second process")
				rc <- run{cmd: cmd2}
			}()
			rc1 := <-rc
			if rc1.err != nil {
				panic(rc1.err)
			}
			rc2 := <-rc
			if rc2.err != nil {
				panic(rc2.err)
			}
			defer func() {
				lg.Debug("killing processes")
				rc1.cmd.Process.Kill()
				rc2.cmd.Process.Kill()
				rc1.cmd.Wait()
				rc2.cmd.Wait()
				lg.Debug("killed processes")
			}()
			lg.Debug("waiting")
			time.Sleep(7 * time.Second)
			lg.Debug("started 2-node etcd cluster")
		} else {
			uss := newEmbedURLs(4)
			ep = uss[0].String() + "/metrics"
			cfgs := []*embed.Config{embed.NewConfig(), embed.NewConfig()}
			cfgs[0].Name, cfgs[1].Name = "0", "1"
			setupEmbedCfg(cfgs[0], []url.URL{uss[0]}, []url.URL{uss[1]}, []url.URL{uss[1], uss[3]})
			setupEmbedCfg(cfgs[1], []url.URL{uss[2]}, []url.URL{uss[3]}, []url.URL{uss[1], uss[3]})
			type embedAndError struct {
				ec	*embed.Etcd
				err	error
			}
			ech := make(chan embedAndError)
			for _, cfg := range cfgs {
				go func(c *embed.Config) {
					e, err := embed.StartEtcd(c)
					if err != nil {
						ech <- embedAndError{err: err}
						return
					}
					<-e.Server.ReadyNotify()
					ech <- embedAndError{ec: e}
				}(cfg)
			}
			for range cfgs {
				ev := <-ech
				if ev.err != nil {
					lg.Panic("failed to start embedded etcd", zap.Error(ev.err))
				}
				defer ev.ec.Close()
			}
			lg.Debug("waiting")
			time.Sleep(7 * time.Second)
			lg.Debug("started 2-node embedded etcd cluster")
		}
	}
	write(ep)
	lg.Debug("fetching metrics", zap.String("endpoint", ep))
	fmt.Println(getMetrics(ep))
}
