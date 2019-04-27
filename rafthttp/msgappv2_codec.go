package rafthttp

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
	"github.com/coreos/etcd/etcdserver/stats"
	"github.com/coreos/etcd/pkg/pbutil"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/raft/raftpb"
)

const (
	msgTypeLinkHeartbeat	uint8	= 0
	msgTypeAppEntries	uint8	= 1
	msgTypeApp		uint8	= 2
	msgAppV2BufSize			= 1024 * 1024
)

type msgAppV2Encoder struct {
	w		io.Writer
	fs		*stats.FollowerStats
	term		uint64
	index		uint64
	buf		[]byte
	uint64buf	[]byte
	uint8buf	[]byte
}

func newMsgAppV2Encoder(w io.Writer, fs *stats.FollowerStats) *msgAppV2Encoder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &msgAppV2Encoder{w: w, fs: fs, buf: make([]byte, msgAppV2BufSize), uint64buf: make([]byte, 8), uint8buf: make([]byte, 1)}
}
func (enc *msgAppV2Encoder) encode(m *raftpb.Message) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	start := time.Now()
	switch {
	case isLinkHeartbeatMessage(m):
		enc.uint8buf[0] = byte(msgTypeLinkHeartbeat)
		if _, err := enc.w.Write(enc.uint8buf); err != nil {
			return err
		}
	case enc.index == m.Index && enc.term == m.LogTerm && m.LogTerm == m.Term:
		enc.uint8buf[0] = byte(msgTypeAppEntries)
		if _, err := enc.w.Write(enc.uint8buf); err != nil {
			return err
		}
		binary.BigEndian.PutUint64(enc.uint64buf, uint64(len(m.Entries)))
		if _, err := enc.w.Write(enc.uint64buf); err != nil {
			return err
		}
		for i := 0; i < len(m.Entries); i++ {
			binary.BigEndian.PutUint64(enc.uint64buf, uint64(m.Entries[i].Size()))
			if _, err := enc.w.Write(enc.uint64buf); err != nil {
				return err
			}
			if n := m.Entries[i].Size(); n < msgAppV2BufSize {
				if _, err := m.Entries[i].MarshalTo(enc.buf); err != nil {
					return err
				}
				if _, err := enc.w.Write(enc.buf[:n]); err != nil {
					return err
				}
			} else {
				if _, err := enc.w.Write(pbutil.MustMarshal(&m.Entries[i])); err != nil {
					return err
				}
			}
			enc.index++
		}
		binary.BigEndian.PutUint64(enc.uint64buf, m.Commit)
		if _, err := enc.w.Write(enc.uint64buf); err != nil {
			return err
		}
		enc.fs.Succ(time.Since(start))
	default:
		if err := binary.Write(enc.w, binary.BigEndian, msgTypeApp); err != nil {
			return err
		}
		if err := binary.Write(enc.w, binary.BigEndian, uint64(m.Size())); err != nil {
			return err
		}
		if _, err := enc.w.Write(pbutil.MustMarshal(m)); err != nil {
			return err
		}
		enc.term = m.Term
		enc.index = m.Index
		if l := len(m.Entries); l > 0 {
			enc.index = m.Entries[l-1].Index
		}
		enc.fs.Succ(time.Since(start))
	}
	return nil
}

type msgAppV2Decoder struct {
	r		io.Reader
	local, remote	types.ID
	term		uint64
	index		uint64
	buf		[]byte
	uint64buf	[]byte
	uint8buf	[]byte
}

func newMsgAppV2Decoder(r io.Reader, local, remote types.ID) *msgAppV2Decoder {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &msgAppV2Decoder{r: r, local: local, remote: remote, buf: make([]byte, msgAppV2BufSize), uint64buf: make([]byte, 8), uint8buf: make([]byte, 1)}
}
func (dec *msgAppV2Decoder) decode() (raftpb.Message, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		m	raftpb.Message
		typ	uint8
	)
	if _, err := io.ReadFull(dec.r, dec.uint8buf); err != nil {
		return m, err
	}
	typ = uint8(dec.uint8buf[0])
	switch typ {
	case msgTypeLinkHeartbeat:
		return linkHeartbeatMessage, nil
	case msgTypeAppEntries:
		m = raftpb.Message{Type: raftpb.MsgApp, From: uint64(dec.remote), To: uint64(dec.local), Term: dec.term, LogTerm: dec.term, Index: dec.index}
		if _, err := io.ReadFull(dec.r, dec.uint64buf); err != nil {
			return m, err
		}
		l := binary.BigEndian.Uint64(dec.uint64buf)
		m.Entries = make([]raftpb.Entry, int(l))
		for i := 0; i < int(l); i++ {
			if _, err := io.ReadFull(dec.r, dec.uint64buf); err != nil {
				return m, err
			}
			size := binary.BigEndian.Uint64(dec.uint64buf)
			var buf []byte
			if size < msgAppV2BufSize {
				buf = dec.buf[:size]
				if _, err := io.ReadFull(dec.r, buf); err != nil {
					return m, err
				}
			} else {
				buf = make([]byte, int(size))
				if _, err := io.ReadFull(dec.r, buf); err != nil {
					return m, err
				}
			}
			dec.index++
			pbutil.MustUnmarshal(&m.Entries[i], buf)
		}
		if _, err := io.ReadFull(dec.r, dec.uint64buf); err != nil {
			return m, err
		}
		m.Commit = binary.BigEndian.Uint64(dec.uint64buf)
	case msgTypeApp:
		var size uint64
		if err := binary.Read(dec.r, binary.BigEndian, &size); err != nil {
			return m, err
		}
		buf := make([]byte, int(size))
		if _, err := io.ReadFull(dec.r, buf); err != nil {
			return m, err
		}
		pbutil.MustUnmarshal(&m, buf)
		dec.term = m.Term
		dec.index = m.Index
		if l := len(m.Entries); l > 0 {
			dec.index = m.Entries[l-1].Index
		}
	default:
		return m, fmt.Errorf("failed to parse type %d in msgappv2 stream", typ)
	}
	return m, nil
}
