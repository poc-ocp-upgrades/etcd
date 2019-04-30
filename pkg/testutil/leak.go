package testutil

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"
)

func CheckLeakedGoroutine() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if testing.Short() {
		return false
	}
	gs := interestingGoroutines()
	if len(gs) == 0 {
		return false
	}
	stackCount := make(map[string]int)
	re := regexp.MustCompile(`\(0[0-9a-fx, ]*\)`)
	for _, g := range gs {
		normalized := string(re.ReplaceAll([]byte(g), []byte("(...)")))
		stackCount[normalized]++
	}
	fmt.Fprintf(os.Stderr, "Too many goroutines running after all test(s).\n")
	for stack, count := range stackCount {
		fmt.Fprintf(os.Stderr, "%d instances of:\n%s\n", count, stack)
	}
	return true
}
func CheckAfterTest(d time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	http.DefaultTransport.(*http.Transport).CloseIdleConnections()
	if testing.Short() {
		return nil
	}
	var bad string
	badSubstring := map[string]string{").writeLoop(": "a Transport", "created by net/http/httptest.(*Server).Start": "an httptest.Server", "timeoutHandler": "a TimeoutHandler", "net.(*netFD).connect(": "a timing out dial", ").noteClientGone(": "a closenotifier sender", ").readLoop(": "a Transport", ".grpc": "a gRPC resource"}
	var stacks string
	begin := time.Now()
	for time.Since(begin) < d {
		bad = ""
		stacks = strings.Join(interestingGoroutines(), "\n\n")
		for substr, what := range badSubstring {
			if strings.Contains(stacks, substr) {
				bad = what
			}
		}
		if bad == "" {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return fmt.Errorf("appears to have leaked %s:\n%s", bad, stacks)
}
func AfterTest(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := CheckAfterTest(300 * time.Millisecond); err != nil {
		t.Errorf("Test %v", err)
	}
}
func interestingGoroutines() (gs []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buf := make([]byte, 2<<20)
	buf = buf[:runtime.Stack(buf, true)]
	for _, g := range strings.Split(string(buf), "\n\n") {
		sl := strings.SplitN(g, "\n", 2)
		if len(sl) != 2 {
			continue
		}
		stack := strings.TrimSpace(sl[1])
		if stack == "" || strings.Contains(stack, "sync.(*WaitGroup).Done") || strings.Contains(stack, "os.(*file).close") || strings.Contains(stack, "created by os/signal.init") || strings.Contains(stack, "runtime/panic.go") || strings.Contains(stack, "created by testing.RunTests") || strings.Contains(stack, "testing.Main(") || strings.Contains(stack, "runtime.goexit") || strings.Contains(stack, "go.etcd.io/etcd/pkg/testutil.interestingGoroutines") || strings.Contains(stack, "go.etcd.io/etcd/pkg/logutil.(*MergeLogger).outputLoop") || strings.Contains(stack, "github.com/golang/glog.(*loggingT).flushDaemon") || strings.Contains(stack, "created by runtime.gc") || strings.Contains(stack, "runtime.MHeap_Scavenger") {
			continue
		}
		gs = append(gs, stack)
	}
	sort.Strings(gs)
	return gs
}
