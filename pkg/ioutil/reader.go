package ioutil

import "io"

func NewLimitedBufferReader(r io.Reader, n int) io.Reader {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &limitedBufferReader{r: r, n: n}
}

type limitedBufferReader struct {
	r	io.Reader
	n	int
}

func (r *limitedBufferReader) Read(p []byte) (n int, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	np := p
	if len(np) > r.n {
		np = np[:r.n]
	}
	return r.r.Read(np)
}
