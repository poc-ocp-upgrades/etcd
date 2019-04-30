package e2e

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"go.etcd.io/etcd/pkg/expect"
)

func waitReadyExpectProc(exproc *expect.ExpectProcess, readyStrs []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := 0
	matchSet := func(l string) bool {
		for _, s := range readyStrs {
			if strings.Contains(l, s) {
				c++
				break
			}
		}
		return c == len(readyStrs)
	}
	_, err := exproc.ExpectFunc(matchSet)
	return err
}
func spawnWithExpect(args []string, expected string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return spawnWithExpects(args, []string{expected}...)
}
func spawnWithExpects(args []string, xs ...string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := spawnWithExpectLines(args, xs...)
	return err
}
func spawnWithExpectLines(args []string, xs ...string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	proc, err := spawnCmd(args)
	if err != nil {
		return nil, err
	}
	var (
		lines		[]string
		lineFunc	= func(txt string) bool {
			return true
		}
	)
	for _, txt := range xs {
		for {
			l, lerr := proc.ExpectFunc(lineFunc)
			if lerr != nil {
				proc.Close()
				return nil, fmt.Errorf("%v (expected %q, got %q)", lerr, txt, lines)
			}
			lines = append(lines, l)
			if strings.Contains(l, txt) {
				break
			}
		}
	}
	perr := proc.Close()
	if len(xs) == 0 && proc.LineCount() != noOutputLineCount {
		return nil, fmt.Errorf("unexpected output (got lines %q, line count %d)", lines, proc.LineCount())
	}
	return lines, perr
}
func randomLeaseID() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
}
func dataMarshal(data interface{}) (d string, e error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	m, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(m), nil
}
func closeWithTimeout(p *expect.ExpectProcess, d time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errc := make(chan error, 1)
	go func() {
		errc <- p.Close()
	}()
	select {
	case err := <-errc:
		return err
	case <-time.After(d):
		p.Stop()
		closeWithTimeout(p, time.Second)
	}
	return fmt.Errorf("took longer than %v to Close process %+v", d, p)
}
func toTLS(s string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return strings.Replace(s, "http://", "https://", 1)
}
