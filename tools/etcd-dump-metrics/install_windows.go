package main

import "errors"

func install(ver, dir string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "", errors.New("windows install is not supported yet")
}
