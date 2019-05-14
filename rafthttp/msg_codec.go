package rafthttp

import (
	"encoding/binary"
	"errors"
	"github.com/coreos/etcd/pkg/pbutil"
	"github.com/coreos/etcd/raft/raftpb"
	"io"
)

type messageEncoder struct{ w io.Writer }

func (enc *messageEncoder) encode(m *raftpb.Message) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := binary.Write(enc.w, binary.BigEndian, uint64(m.Size())); err != nil {
		return err
	}
	_, err := enc.w.Write(pbutil.MustMarshal(m))
	return err
}

type messageDecoder struct{ r io.Reader }

var (
	readBytesLimit     uint64 = 512 * 1024 * 1024
	ErrExceedSizeLimit        = errors.New("rafthttp: error limit exceeded")
)

func (dec *messageDecoder) decode() (raftpb.Message, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return dec.decodeLimit(readBytesLimit)
}
func (dec *messageDecoder) decodeLimit(numBytes uint64) (raftpb.Message, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var m raftpb.Message
	var l uint64
	if err := binary.Read(dec.r, binary.BigEndian, &l); err != nil {
		return m, err
	}
	if l > numBytes {
		return m, ErrExceedSizeLimit
	}
	buf := make([]byte, int(l))
	if _, err := io.ReadFull(dec.r, buf); err != nil {
		return m, err
	}
	return m, m.Unmarshal(buf)
}
