package logutil

import (
	"fmt"
	"sync"
	"time"
	"github.com/coreos/pkg/capnslog"
)

var (
	defaultMergePeriod	= time.Second
	defaultTimeOutputScale	= 10 * time.Millisecond
	outputInterval		= time.Second
)

type line struct {
	level	capnslog.LogLevel
	str	string
}

func (l line) append(s string) line {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return line{level: l.level, str: l.str + " " + s}
}

type status struct {
	period	time.Duration
	start	time.Time
	count	int
}

func (s *status) isInMergePeriod(now time.Time) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.period == 0 || s.start.Add(s.period).After(now)
}
func (s *status) isEmpty() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return s.count == 0
}
func (s *status) summary(now time.Time) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ts := s.start.Round(defaultTimeOutputScale)
	took := now.Round(defaultTimeOutputScale).Sub(ts)
	return fmt.Sprintf("[merged %d repeated lines in %s]", s.count, took)
}
func (s *status) reset(now time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.start = now
	s.count = 0
}

type MergeLogger struct {
	*capnslog.PackageLogger
	mu	sync.Mutex
	statusm	map[line]*status
}

func NewMergeLogger(logger *capnslog.PackageLogger) *MergeLogger {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l := &MergeLogger{PackageLogger: logger, statusm: make(map[line]*status)}
	go l.outputLoop()
	return l
}
func (l *MergeLogger) MergeInfo(entries ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.merge(line{level: capnslog.INFO, str: fmt.Sprint(entries...)})
}
func (l *MergeLogger) MergeInfof(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.merge(line{level: capnslog.INFO, str: fmt.Sprintf(format, args...)})
}
func (l *MergeLogger) MergeNotice(entries ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.merge(line{level: capnslog.NOTICE, str: fmt.Sprint(entries...)})
}
func (l *MergeLogger) MergeNoticef(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.merge(line{level: capnslog.NOTICE, str: fmt.Sprintf(format, args...)})
}
func (l *MergeLogger) MergeWarning(entries ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.merge(line{level: capnslog.WARNING, str: fmt.Sprint(entries...)})
}
func (l *MergeLogger) MergeWarningf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.merge(line{level: capnslog.WARNING, str: fmt.Sprintf(format, args...)})
}
func (l *MergeLogger) MergeError(entries ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.merge(line{level: capnslog.ERROR, str: fmt.Sprint(entries...)})
}
func (l *MergeLogger) MergeErrorf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.merge(line{level: capnslog.ERROR, str: fmt.Sprintf(format, args...)})
}
func (l *MergeLogger) merge(ln line) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l.mu.Lock()
	if status, ok := l.statusm[ln]; ok {
		status.count++
		l.mu.Unlock()
		return
	}
	l.statusm[ln] = &status{period: defaultMergePeriod, start: time.Now()}
	l.mu.Unlock()
	l.PackageLogger.Logf(ln.level, ln.str)
}
func (l *MergeLogger) outputLoop() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for now := range time.Tick(outputInterval) {
		var outputs []line
		l.mu.Lock()
		for ln, status := range l.statusm {
			if status.isInMergePeriod(now) {
				continue
			}
			if status.isEmpty() {
				delete(l.statusm, ln)
				continue
			}
			outputs = append(outputs, ln.append(status.summary(now)))
			status.reset(now)
		}
		l.mu.Unlock()
		for _, o := range outputs {
			l.PackageLogger.Logf(o.level, o.str)
		}
	}
}
