package flags

type IgnoredFlag struct{ Name string }

func (f *IgnoredFlag) IsBoolFlag() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (f *IgnoredFlag) Set(s string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	plog.Warningf(`flag "-%s" is no longer supported - ignoring.`, f.Name)
	return nil
}
func (f *IgnoredFlag) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ""
}
