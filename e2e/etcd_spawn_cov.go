package e2e

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"github.com/coreos/etcd/pkg/expect"
	"github.com/coreos/etcd/pkg/fileutil"
	"github.com/coreos/etcd/pkg/flags"
)

const noOutputLineCount = 2

func spawnCmd(args []string) (*expect.ExpectProcess, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if args[0] == binPath {
		return spawnEtcd(args)
	}
	if args[0] == ctlBinPath || args[0] == ctlBinPath+"3" {
		env := []string{"ETCDCTL_ARGS=" + strings.Join(args, "\xe7\xcd")}
		if args[0] == ctlBinPath+"3" {
			env = append(env, "ETCDCTL_API=3")
		}
		covArgs, err := getCovArgs()
		if err != nil {
			return nil, err
		}
		env = append(env, os.Environ()...)
		ep, err := expect.NewExpectWithEnv(binDir+"/etcdctl_test", covArgs, env)
		if err != nil {
			return nil, err
		}
		ep.StopSignal = syscall.SIGTERM
		return ep, nil
	}
	return expect.NewExpect(args[0], args[1:]...)
}
func spawnEtcd(args []string) (*expect.ExpectProcess, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	covArgs, err := getCovArgs()
	if err != nil {
		return nil, err
	}
	var env []string
	if args[1] == "grpc-proxy" {
		env = append(os.Environ(), "ETCDCOV_ARGS="+strings.Join(args, "\xe7\xcd"))
	} else {
		env = args2env(args[1:])
	}
	ep, err := expect.NewExpectWithEnv(binDir+"/etcd_test", covArgs, env)
	if err != nil {
		return nil, err
	}
	ep.StopSignal = syscall.SIGTERM
	return ep, nil
}
func getCovArgs() ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	coverPath := os.Getenv("COVERDIR")
	if !filepath.IsAbs(coverPath) {
		coverPath = filepath.Join("..", coverPath)
	}
	if !fileutil.Exist(coverPath) {
		return nil, fmt.Errorf("could not find coverage folder")
	}
	covArgs := []string{fmt.Sprintf("-test.coverprofile=e2e.%v.coverprofile", time.Now().UnixNano()), "-test.outputdir=" + coverPath}
	return covArgs, nil
}
func args2env(args []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var covEnvs []string
	for i := range args {
		if !strings.HasPrefix(args[i], "--") {
			continue
		}
		flag := strings.Split(args[i], "--")[1]
		val := "true"
		if strings.Contains(args[i], "=") {
			split := strings.Split(flag, "=")
			flag = split[0]
			val = split[1]
		}
		if i+1 < len(args) {
			if !strings.HasPrefix(args[i+1], "--") {
				val = args[i+1]
			}
		}
		covEnvs = append(covEnvs, flags.FlagToEnv("ETCD", flag)+"="+val)
	}
	return covEnvs
}
