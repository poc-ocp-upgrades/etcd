package main

import "os"

func existFileOrDir(name string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := os.Stat(name)
	return err == nil
}
