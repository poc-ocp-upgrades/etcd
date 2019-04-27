package concurrency

import (
	"context"
	"math"
	v3 "github.com/coreos/etcd/clientv3"
)

type STM interface {
	Get(key ...string) string
	Put(key, val string, opts ...v3.OpOption)
	Rev(key string) int64
	Del(key string)
	commit() *v3.TxnResponse
	reset()
}
type Isolation int

const (
	SerializableSnapshot	Isolation	= iota
	Serializable
	RepeatableReads
	ReadCommitted
)

type stmError struct{ err error }
type stmOptions struct {
	iso		Isolation
	ctx		context.Context
	prefetch	[]string
}
type stmOption func(*stmOptions)

func WithIsolation(lvl Isolation) stmOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(so *stmOptions) {
		so.iso = lvl
	}
}
func WithAbortContext(ctx context.Context) stmOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(so *stmOptions) {
		so.ctx = ctx
	}
}
func WithPrefetch(keys ...string) stmOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(so *stmOptions) {
		so.prefetch = append(so.prefetch, keys...)
	}
}
func NewSTM(c *v3.Client, apply func(STM) error, so ...stmOption) (*v3.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := &stmOptions{ctx: c.Ctx()}
	for _, f := range so {
		f(opts)
	}
	if len(opts.prefetch) != 0 {
		f := apply
		apply = func(s STM) error {
			s.Get(opts.prefetch...)
			return f(s)
		}
	}
	return runSTM(mkSTM(c, opts), apply)
}
func mkSTM(c *v3.Client, opts *stmOptions) STM {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch opts.iso {
	case SerializableSnapshot:
		s := &stmSerializable{stm: stm{client: c, ctx: opts.ctx}, prefetch: make(map[string]*v3.GetResponse)}
		s.conflicts = func() []v3.Cmp {
			return append(s.rset.cmps(), s.wset.cmps(s.rset.first()+1)...)
		}
		return s
	case Serializable:
		s := &stmSerializable{stm: stm{client: c, ctx: opts.ctx}, prefetch: make(map[string]*v3.GetResponse)}
		s.conflicts = func() []v3.Cmp {
			return s.rset.cmps()
		}
		return s
	case RepeatableReads:
		s := &stm{client: c, ctx: opts.ctx, getOpts: []v3.OpOption{v3.WithSerializable()}}
		s.conflicts = func() []v3.Cmp {
			return s.rset.cmps()
		}
		return s
	case ReadCommitted:
		s := &stm{client: c, ctx: opts.ctx, getOpts: []v3.OpOption{v3.WithSerializable()}}
		s.conflicts = func() []v3.Cmp {
			return nil
		}
		return s
	default:
		panic("unsupported stm")
	}
}

type stmResponse struct {
	resp	*v3.TxnResponse
	err	error
}

func runSTM(s STM, apply func(STM) error) (*v3.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	outc := make(chan stmResponse, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				e, ok := r.(stmError)
				if !ok {
					panic(r)
				}
				outc <- stmResponse{nil, e.err}
			}
		}()
		var out stmResponse
		for {
			s.reset()
			if out.err = apply(s); out.err != nil {
				break
			}
			if out.resp = s.commit(); out.resp != nil {
				break
			}
		}
		outc <- out
	}()
	r := <-outc
	return r.resp, r.err
}

type stm struct {
	client		*v3.Client
	ctx		context.Context
	rset		readSet
	wset		writeSet
	getOpts		[]v3.OpOption
	conflicts	func() []v3.Cmp
}
type stmPut struct {
	val	string
	op	v3.Op
}
type readSet map[string]*v3.GetResponse

func (rs readSet) add(keys []string, txnresp *v3.TxnResponse) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i, resp := range txnresp.Responses {
		rs[keys[i]] = (*v3.GetResponse)(resp.GetResponseRange())
	}
}
func (rs readSet) first() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := int64(math.MaxInt64 - 1)
	for _, resp := range rs {
		if rev := resp.Header.Revision; rev < ret {
			ret = rev
		}
	}
	return ret
}
func (rs readSet) cmps() []v3.Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmps := make([]v3.Cmp, 0, len(rs))
	for k, rk := range rs {
		cmps = append(cmps, isKeyCurrent(k, rk))
	}
	return cmps
}

type writeSet map[string]stmPut

func (ws writeSet) get(keys ...string) *stmPut {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, key := range keys {
		if wv, ok := ws[key]; ok {
			return &wv
		}
	}
	return nil
}
func (ws writeSet) cmps(rev int64) []v3.Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmps := make([]v3.Cmp, 0, len(ws))
	for key := range ws {
		cmps = append(cmps, v3.Compare(v3.ModRevision(key), "<", rev))
	}
	return cmps
}
func (ws writeSet) puts() []v3.Op {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	puts := make([]v3.Op, 0, len(ws))
	for _, v := range ws {
		puts = append(puts, v.op)
	}
	return puts
}
func (s *stm) Get(keys ...string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if wv := s.wset.get(keys...); wv != nil {
		return wv.val
	}
	return respToValue(s.fetch(keys...))
}
func (s *stm) Put(key, val string, opts ...v3.OpOption) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.wset[key] = stmPut{val, v3.OpPut(key, val, opts...)}
}
func (s *stm) Del(key string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.wset[key] = stmPut{"", v3.OpDelete(key)}
}
func (s *stm) Rev(key string) int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if resp := s.fetch(key); resp != nil && len(resp.Kvs) != 0 {
		return resp.Kvs[0].ModRevision
	}
	return 0
}
func (s *stm) commit() *v3.TxnResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	txnresp, err := s.client.Txn(s.ctx).If(s.conflicts()...).Then(s.wset.puts()...).Commit()
	if err != nil {
		panic(stmError{err})
	}
	if txnresp.Succeeded {
		return txnresp
	}
	return nil
}
func (s *stm) fetch(keys ...string) *v3.GetResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(keys) == 0 {
		return nil
	}
	ops := make([]v3.Op, len(keys))
	for i, key := range keys {
		if resp, ok := s.rset[key]; ok {
			return resp
		}
		ops[i] = v3.OpGet(key, s.getOpts...)
	}
	txnresp, err := s.client.Txn(s.ctx).Then(ops...).Commit()
	if err != nil {
		panic(stmError{err})
	}
	s.rset.add(keys, txnresp)
	return (*v3.GetResponse)(txnresp.Responses[0].GetResponseRange())
}
func (s *stm) reset() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.rset = make(map[string]*v3.GetResponse)
	s.wset = make(map[string]stmPut)
}

type stmSerializable struct {
	stm
	prefetch	map[string]*v3.GetResponse
}

func (s *stmSerializable) Get(keys ...string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if wv := s.wset.get(keys...); wv != nil {
		return wv.val
	}
	firstRead := len(s.rset) == 0
	for _, key := range keys {
		if resp, ok := s.prefetch[key]; ok {
			delete(s.prefetch, key)
			s.rset[key] = resp
		}
	}
	resp := s.stm.fetch(keys...)
	if firstRead {
		s.getOpts = []v3.OpOption{v3.WithRev(resp.Header.Revision), v3.WithSerializable()}
	}
	return respToValue(resp)
}
func (s *stmSerializable) Rev(key string) int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.Get(key)
	return s.stm.Rev(key)
}
func (s *stmSerializable) gets() ([]string, []v3.Op) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	keys := make([]string, 0, len(s.rset))
	ops := make([]v3.Op, 0, len(s.rset))
	for k := range s.rset {
		keys = append(keys, k)
		ops = append(ops, v3.OpGet(k))
	}
	return keys, ops
}
func (s *stmSerializable) commit() *v3.TxnResponse {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	keys, getops := s.gets()
	txn := s.client.Txn(s.ctx).If(s.conflicts()...).Then(s.wset.puts()...)
	txnresp, err := txn.Else(getops...).Commit()
	if err != nil {
		panic(stmError{err})
	}
	if txnresp.Succeeded {
		return txnresp
	}
	s.rset.add(keys, txnresp)
	s.prefetch = s.rset
	s.getOpts = nil
	return nil
}
func isKeyCurrent(k string, r *v3.GetResponse) v3.Cmp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(r.Kvs) != 0 {
		return v3.Compare(v3.ModRevision(k), "=", r.Kvs[0].ModRevision)
	}
	return v3.Compare(v3.ModRevision(k), "=", 0)
}
func respToValue(resp *v3.GetResponse) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if resp == nil || len(resp.Kvs) == 0 {
		return ""
	}
	return string(resp.Kvs[0].Value)
}
func NewSTMRepeatable(ctx context.Context, c *v3.Client, apply func(STM) error) (*v3.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewSTM(c, apply, WithAbortContext(ctx), WithIsolation(RepeatableReads))
}
func NewSTMSerializable(ctx context.Context, c *v3.Client, apply func(STM) error) (*v3.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewSTM(c, apply, WithAbortContext(ctx), WithIsolation(Serializable))
}
func NewSTMReadCommitted(ctx context.Context, c *v3.Client, apply func(STM) error) (*v3.TxnResponse, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewSTM(c, apply, WithAbortContext(ctx), WithIsolation(ReadCommitted))
}
