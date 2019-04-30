package logutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"go.etcd.io/etcd/pkg/systemd"
	"github.com/coreos/go-systemd/journal"
	"go.uber.org/zap/zapcore"
)

func NewJournalWriter(wr io.Writer) (io.Writer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &journalWriter{Writer: wr}, systemd.DialJournal()
}

type journalWriter struct{ io.Writer }
type logLine struct {
	Level	string	`json:"level"`
	Caller	string	`json:"caller"`
}

func (w *journalWriter) Write(p []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	line := &logLine{}
	if err := json.NewDecoder(bytes.NewReader(p)).Decode(line); err != nil {
		return 0, err
	}
	var pri journal.Priority
	switch line.Level {
	case zapcore.DebugLevel.String():
		pri = journal.PriDebug
	case zapcore.InfoLevel.String():
		pri = journal.PriInfo
	case zapcore.WarnLevel.String():
		pri = journal.PriWarning
	case zapcore.ErrorLevel.String():
		pri = journal.PriErr
	case zapcore.DPanicLevel.String():
		pri = journal.PriCrit
	case zapcore.PanicLevel.String():
		pri = journal.PriCrit
	case zapcore.FatalLevel.String():
		pri = journal.PriCrit
	default:
		panic(fmt.Errorf("unknown log level: %q", line.Level))
	}
	err := journal.Send(string(p), pri, map[string]string{"PACKAGE": filepath.Dir(line.Caller), "SYSLOG_IDENTIFIER": filepath.Base(os.Args[0])})
	if err != nil {
		return w.Writer.Write(p)
	}
	return 0, nil
}
