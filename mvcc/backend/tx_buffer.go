package backend

import (
	"bytes"
	"sort"
)

type txBuffer struct{ buckets map[string]*bucketBuffer }

func (txb *txBuffer) reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for k, v := range txb.buckets {
		if v.used == 0 {
			delete(txb.buckets, k)
		}
		v.used = 0
	}
}

type txWriteBuffer struct {
	txBuffer
	seq	bool
}

func (txw *txWriteBuffer) put(bucket, k, v []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	txw.seq = false
	txw.putSeq(bucket, k, v)
}
func (txw *txWriteBuffer) putSeq(bucket, k, v []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, ok := txw.buckets[string(bucket)]
	if !ok {
		b = newBucketBuffer()
		txw.buckets[string(bucket)] = b
	}
	b.add(k, v)
}
func (txw *txWriteBuffer) writeback(txr *txReadBuffer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for k, wb := range txw.buckets {
		rb, ok := txr.buckets[k]
		if !ok {
			delete(txw.buckets, k)
			txr.buckets[k] = wb
			continue
		}
		if !txw.seq && wb.used > 1 {
			sort.Sort(wb)
		}
		rb.merge(wb)
	}
	txw.reset()
}

type txReadBuffer struct{ txBuffer }

func (txr *txReadBuffer) Range(bucketName, key, endKey []byte, limit int64) ([][]byte, [][]byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b := txr.buckets[string(bucketName)]; b != nil {
		return b.Range(key, endKey, limit)
	}
	return nil, nil
}
func (txr *txReadBuffer) ForEach(bucketName []byte, visitor func(k, v []byte) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if b := txr.buckets[string(bucketName)]; b != nil {
		return b.ForEach(visitor)
	}
	return nil
}

type kv struct {
	key	[]byte
	val	[]byte
}
type bucketBuffer struct {
	buf	[]kv
	used	int
}

func newBucketBuffer() *bucketBuffer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &bucketBuffer{buf: make([]kv, 512), used: 0}
}
func (bb *bucketBuffer) Range(key, endKey []byte, limit int64) (keys [][]byte, vals [][]byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := func(i int) bool {
		return bytes.Compare(bb.buf[i].key, key) >= 0
	}
	idx := sort.Search(bb.used, f)
	if idx < 0 {
		return nil, nil
	}
	if len(endKey) == 0 {
		if bytes.Equal(key, bb.buf[idx].key) {
			keys = append(keys, bb.buf[idx].key)
			vals = append(vals, bb.buf[idx].val)
		}
		return keys, vals
	}
	if bytes.Compare(endKey, bb.buf[idx].key) <= 0 {
		return nil, nil
	}
	for i := idx; i < bb.used && int64(len(keys)) < limit; i++ {
		if bytes.Compare(endKey, bb.buf[i].key) <= 0 {
			break
		}
		keys = append(keys, bb.buf[i].key)
		vals = append(vals, bb.buf[i].val)
	}
	return keys, vals
}
func (bb *bucketBuffer) ForEach(visitor func(k, v []byte) error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < bb.used; i++ {
		if err := visitor(bb.buf[i].key, bb.buf[i].val); err != nil {
			return err
		}
	}
	return nil
}
func (bb *bucketBuffer) add(k, v []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bb.buf[bb.used].key, bb.buf[bb.used].val = k, v
	bb.used++
	if bb.used == len(bb.buf) {
		buf := make([]kv, (3*len(bb.buf))/2)
		copy(buf, bb.buf)
		bb.buf = buf
	}
}
func (bb *bucketBuffer) merge(bbsrc *bucketBuffer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := 0; i < bbsrc.used; i++ {
		bb.add(bbsrc.buf[i].key, bbsrc.buf[i].val)
	}
	if bb.used == bbsrc.used {
		return
	}
	if bytes.Compare(bb.buf[(bb.used-bbsrc.used)-1].key, bbsrc.buf[0].key) < 0 {
		return
	}
	sort.Stable(bb)
	widx := 0
	for ridx := 1; ridx < bb.used; ridx++ {
		if !bytes.Equal(bb.buf[ridx].key, bb.buf[widx].key) {
			widx++
		}
		bb.buf[widx] = bb.buf[ridx]
	}
	bb.used = widx + 1
}
func (bb *bucketBuffer) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bb.used
}
func (bb *bucketBuffer) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bytes.Compare(bb.buf[i].key, bb.buf[j].key) < 0
}
func (bb *bucketBuffer) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bb.buf[i], bb.buf[j] = bb.buf[j], bb.buf[i]
}
