package clientv3

import pb "github.com/coreos/etcd/etcdserver/etcdserverpb"

type opType int

const (
	tRange	opType	= iota + 1
	tPut
	tDeleteRange
	tTxn
)

var (
	noPrefixEnd = []byte{0}
)

type Op struct {
	t		opType
	key		[]byte
	end		[]byte
	limit		int64
	sort		*SortOption
	serializable	bool
	keysOnly	bool
	countOnly	bool
	minModRev	int64
	maxModRev	int64
	minCreateRev	int64
	maxCreateRev	int64
	rev		int64
	prevKV		bool
	ignoreValue	bool
	ignoreLease	bool
	progressNotify	bool
	createdNotify	bool
	filterPut	bool
	filterDelete	bool
	val		[]byte
	leaseID		LeaseID
	cmps		[]Cmp
	thenOps		[]Op
	elseOps		[]Op
}

func (op Op) IsTxn() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.t == tTxn
}
func (op Op) Txn() ([]Cmp, []Op, []Op) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.cmps, op.thenOps, op.elseOps
}
func (op Op) KeyBytes() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.key
}
func (op *Op) WithKeyBytes(key []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	op.key = key
}
func (op Op) RangeBytes() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.end
}
func (op Op) Rev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.rev
}
func (op Op) IsPut() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.t == tPut
}
func (op Op) IsGet() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.t == tRange
}
func (op Op) IsDelete() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.t == tDeleteRange
}
func (op Op) IsSerializable() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.serializable == true
}
func (op Op) IsKeysOnly() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.keysOnly == true
}
func (op Op) IsCountOnly() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.countOnly == true
}
func (op Op) MinModRev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.minModRev
}
func (op Op) MaxModRev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.maxModRev
}
func (op Op) MinCreateRev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.minCreateRev
}
func (op Op) MaxCreateRev() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.maxCreateRev
}
func (op *Op) WithRangeBytes(end []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	op.end = end
}
func (op Op) ValueBytes() []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return op.val
}
func (op *Op) WithValueBytes(v []byte) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	op.val = v
}
func (op Op) toRangeRequest() *pb.RangeRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if op.t != tRange {
		panic("op.t != tRange")
	}
	r := &pb.RangeRequest{Key: op.key, RangeEnd: op.end, Limit: op.limit, Revision: op.rev, Serializable: op.serializable, KeysOnly: op.keysOnly, CountOnly: op.countOnly, MinModRevision: op.minModRev, MaxModRevision: op.maxModRev, MinCreateRevision: op.minCreateRev, MaxCreateRevision: op.maxCreateRev}
	if op.sort != nil {
		r.SortOrder = pb.RangeRequest_SortOrder(op.sort.Order)
		r.SortTarget = pb.RangeRequest_SortTarget(op.sort.Target)
	}
	return r
}
func (op Op) toTxnRequest() *pb.TxnRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	thenOps := make([]*pb.RequestOp, len(op.thenOps))
	for i, tOp := range op.thenOps {
		thenOps[i] = tOp.toRequestOp()
	}
	elseOps := make([]*pb.RequestOp, len(op.elseOps))
	for i, eOp := range op.elseOps {
		elseOps[i] = eOp.toRequestOp()
	}
	cmps := make([]*pb.Compare, len(op.cmps))
	for i := range op.cmps {
		cmps[i] = (*pb.Compare)(&op.cmps[i])
	}
	return &pb.TxnRequest{Compare: cmps, Success: thenOps, Failure: elseOps}
}
func (op Op) toRequestOp() *pb.RequestOp {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch op.t {
	case tRange:
		return &pb.RequestOp{Request: &pb.RequestOp_RequestRange{RequestRange: op.toRangeRequest()}}
	case tPut:
		r := &pb.PutRequest{Key: op.key, Value: op.val, Lease: int64(op.leaseID), PrevKv: op.prevKV, IgnoreValue: op.ignoreValue, IgnoreLease: op.ignoreLease}
		return &pb.RequestOp{Request: &pb.RequestOp_RequestPut{RequestPut: r}}
	case tDeleteRange:
		r := &pb.DeleteRangeRequest{Key: op.key, RangeEnd: op.end, PrevKv: op.prevKV}
		return &pb.RequestOp{Request: &pb.RequestOp_RequestDeleteRange{RequestDeleteRange: r}}
	case tTxn:
		return &pb.RequestOp{Request: &pb.RequestOp_RequestTxn{RequestTxn: op.toTxnRequest()}}
	default:
		panic("Unknown Op")
	}
}
func (op Op) isWrite() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if op.t == tTxn {
		for _, tOp := range op.thenOps {
			if tOp.isWrite() {
				return true
			}
		}
		for _, tOp := range op.elseOps {
			if tOp.isWrite() {
				return true
			}
		}
		return false
	}
	return op.t != tRange
}
func OpGet(key string, opts ...OpOption) Op {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := Op{t: tRange, key: []byte(key)}
	ret.applyOpts(opts)
	return ret
}
func OpDelete(key string, opts ...OpOption) Op {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := Op{t: tDeleteRange, key: []byte(key)}
	ret.applyOpts(opts)
	switch {
	case ret.leaseID != 0:
		panic("unexpected lease in delete")
	case ret.limit != 0:
		panic("unexpected limit in delete")
	case ret.rev != 0:
		panic("unexpected revision in delete")
	case ret.sort != nil:
		panic("unexpected sort in delete")
	case ret.serializable:
		panic("unexpected serializable in delete")
	case ret.countOnly:
		panic("unexpected countOnly in delete")
	case ret.minModRev != 0, ret.maxModRev != 0:
		panic("unexpected mod revision filter in delete")
	case ret.minCreateRev != 0, ret.maxCreateRev != 0:
		panic("unexpected create revision filter in delete")
	case ret.filterDelete, ret.filterPut:
		panic("unexpected filter in delete")
	case ret.createdNotify:
		panic("unexpected createdNotify in delete")
	}
	return ret
}
func OpPut(key, val string, opts ...OpOption) Op {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := Op{t: tPut, key: []byte(key), val: []byte(val)}
	ret.applyOpts(opts)
	switch {
	case ret.end != nil:
		panic("unexpected range in put")
	case ret.limit != 0:
		panic("unexpected limit in put")
	case ret.rev != 0:
		panic("unexpected revision in put")
	case ret.sort != nil:
		panic("unexpected sort in put")
	case ret.serializable:
		panic("unexpected serializable in put")
	case ret.countOnly:
		panic("unexpected countOnly in put")
	case ret.minModRev != 0, ret.maxModRev != 0:
		panic("unexpected mod revision filter in put")
	case ret.minCreateRev != 0, ret.maxCreateRev != 0:
		panic("unexpected create revision filter in put")
	case ret.filterDelete, ret.filterPut:
		panic("unexpected filter in put")
	case ret.createdNotify:
		panic("unexpected createdNotify in put")
	}
	return ret
}
func OpTxn(cmps []Cmp, thenOps []Op, elseOps []Op) Op {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Op{t: tTxn, cmps: cmps, thenOps: thenOps, elseOps: elseOps}
}
func opWatch(key string, opts ...OpOption) Op {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := Op{t: tRange, key: []byte(key)}
	ret.applyOpts(opts)
	switch {
	case ret.leaseID != 0:
		panic("unexpected lease in watch")
	case ret.limit != 0:
		panic("unexpected limit in watch")
	case ret.sort != nil:
		panic("unexpected sort in watch")
	case ret.serializable:
		panic("unexpected serializable in watch")
	case ret.countOnly:
		panic("unexpected countOnly in watch")
	case ret.minModRev != 0, ret.maxModRev != 0:
		panic("unexpected mod revision filter in watch")
	case ret.minCreateRev != 0, ret.maxCreateRev != 0:
		panic("unexpected create revision filter in watch")
	}
	return ret
}
func (op *Op) applyOpts(opts []OpOption) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, opt := range opts {
		opt(op)
	}
}

type OpOption func(*Op)

func WithLease(leaseID LeaseID) OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.leaseID = leaseID
	}
}
func WithLimit(n int64) OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.limit = n
	}
}
func WithRev(rev int64) OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.rev = rev
	}
}
func WithSort(target SortTarget, order SortOrder) OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		if target == SortByKey && order == SortAscend {
			order = SortNone
		}
		op.sort = &SortOption{target, order}
	}
}
func GetPrefixRangeEnd(prefix string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(getPrefix([]byte(prefix)))
}
func getPrefix(key []byte) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	end := make([]byte, len(key))
	copy(end, key)
	for i := len(end) - 1; i >= 0; i-- {
		if end[i] < 0xff {
			end[i] = end[i] + 1
			end = end[:i+1]
			return end
		}
	}
	return noPrefixEnd
}
func WithPrefix() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		if len(op.key) == 0 {
			op.key, op.end = []byte{0}, []byte{0}
			return
		}
		op.end = getPrefix(op.key)
	}
}
func WithRange(endKey string) OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.end = []byte(endKey)
	}
}
func WithFromKey() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return WithRange("\x00")
}
func WithSerializable() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.serializable = true
	}
}
func WithKeysOnly() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.keysOnly = true
	}
}
func WithCountOnly() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.countOnly = true
	}
}
func WithMinModRev(rev int64) OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.minModRev = rev
	}
}
func WithMaxModRev(rev int64) OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.maxModRev = rev
	}
}
func WithMinCreateRev(rev int64) OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.minCreateRev = rev
	}
}
func WithMaxCreateRev(rev int64) OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.maxCreateRev = rev
	}
}
func WithFirstCreate() []OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return withTop(SortByCreateRevision, SortAscend)
}
func WithLastCreate() []OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return withTop(SortByCreateRevision, SortDescend)
}
func WithFirstKey() []OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return withTop(SortByKey, SortAscend)
}
func WithLastKey() []OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return withTop(SortByKey, SortDescend)
}
func WithFirstRev() []OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return withTop(SortByModRevision, SortAscend)
}
func WithLastRev() []OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return withTop(SortByModRevision, SortDescend)
}
func withTop(target SortTarget, order SortOrder) []OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []OpOption{WithPrefix(), WithSort(target, order), WithLimit(1)}
}
func WithProgressNotify() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.progressNotify = true
	}
}
func WithCreatedNotify() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.createdNotify = true
	}
}
func WithFilterPut() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.filterPut = true
	}
}
func WithFilterDelete() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.filterDelete = true
	}
}
func WithPrevKV() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.prevKV = true
	}
}
func WithIgnoreValue() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.ignoreValue = true
	}
}
func WithIgnoreLease() OpOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *Op) {
		op.ignoreLease = true
	}
}

type LeaseOp struct {
	id		LeaseID
	attachedKeys	bool
}
type LeaseOption func(*LeaseOp)

func (op *LeaseOp) applyOpts(opts []LeaseOption) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, opt := range opts {
		opt(op)
	}
}
func WithAttachedKeys() LeaseOption {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(op *LeaseOp) {
		op.attachedKeys = true
	}
}
func toLeaseTimeToLiveRequest(id LeaseID, opts ...LeaseOption) *pb.LeaseTimeToLiveRequest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := &LeaseOp{id: id}
	ret.applyOpts(opts)
	return &pb.LeaseTimeToLiveRequest{ID: int64(id), Keys: ret.attachedKeys}
}
