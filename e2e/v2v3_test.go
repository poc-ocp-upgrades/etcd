package e2e

func addV2Args(args []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return append(args, "--experimental-enable-v2v3", "v2/")
}
