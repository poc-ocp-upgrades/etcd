package ioutil

import (
	"bytes"
	"io"
	"testing"
)

type readerNilCloser struct{ io.Reader }

func (rc *readerNilCloser) Close() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func TestExactReadCloserExpectEOF(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buf := bytes.NewBuffer(make([]byte, 10))
	rc := NewExactReadCloser(&readerNilCloser{buf}, 1)
	if _, err := rc.Read(make([]byte, 10)); err != ErrExpectEOF {
		t.Fatalf("expected %v, got %v", ErrExpectEOF, err)
	}
}
func TestExactReadCloserShort(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buf := bytes.NewBuffer(make([]byte, 5))
	rc := NewExactReadCloser(&readerNilCloser{buf}, 10)
	if _, err := rc.Read(make([]byte, 10)); err != nil {
		t.Fatalf("Read expected nil err, got %v", err)
	}
	if err := rc.Close(); err != ErrShortRead {
		t.Fatalf("Close expected %v, got %v", ErrShortRead, err)
	}
}
