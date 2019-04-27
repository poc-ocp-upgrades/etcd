package flags

import (
	"flag"
	"os"
	"testing"
)

func TestSetFlagsFromEnv(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fs := flag.NewFlagSet("testing", flag.ExitOnError)
	fs.String("a", "", "")
	fs.String("b", "", "")
	fs.String("c", "", "")
	fs.Parse([]string{})
	os.Clearenv()
	os.Setenv("ETCD_A", "foo")
	if err := fs.Set("b", "bar"); err != nil {
		t.Fatal(err)
	}
	os.Setenv("ETCD_C", "woof")
	if err := fs.Set("c", "quack"); err != nil {
		t.Fatal(err)
	}
	for f, want := range map[string]string{"a": "", "b": "bar", "c": "quack"} {
		if got := fs.Lookup(f).Value.String(); got != want {
			t.Fatalf("flag %q=%q, want %q", f, got, want)
		}
	}
	err := SetFlagsFromEnv("ETCD", fs)
	if err != nil {
		t.Errorf("err=%v, want nil", err)
	}
	for f, want := range map[string]string{"a": "foo", "b": "bar", "c": "quack"} {
		if got := fs.Lookup(f).Value.String(); got != want {
			t.Errorf("flag %q=%q, want %q", f, got, want)
		}
	}
}
func TestSetFlagsFromEnvBad(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fs := flag.NewFlagSet("testing", flag.ExitOnError)
	fs.Int("x", 0, "")
	os.Setenv("ETCD_X", "not_a_number")
	if err := SetFlagsFromEnv("ETCD", fs); err == nil {
		t.Errorf("err=nil, want != nil")
	}
}
func TestSetFlagsFromEnvParsingError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fs := flag.NewFlagSet("etcd", flag.ContinueOnError)
	var tickMs uint
	fs.UintVar(&tickMs, "heartbeat-interval", 0, "Time (in milliseconds) of a heartbeat interval.")
	if oerr := os.Setenv("ETCD_HEARTBEAT_INTERVAL", "100 # ms"); oerr != nil {
		t.Fatal(oerr)
	}
	defer os.Unsetenv("ETCD_HEARTBEAT_INTERVAL")
	if serr := SetFlagsFromEnv("ETCD", fs); serr.Error() != `invalid value "100 # ms" for ETCD_HEARTBEAT_INTERVAL: strconv.ParseUint: parsing "100 # ms": invalid syntax` {
		t.Fatalf("expected parsing error, got %v", serr)
	}
}
